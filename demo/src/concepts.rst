Concepts
=========================================================================

The Task Unit
-------------

This section discusses the concept and implementation of a task.  It gives an introduction to the terminology associated with a task and discusses the kernel abstraction supported by the task.

In implementation terms, a task is a run-time instance of the source code.  an example of a kernel task is shown below.  The program gains control of the processor, executes application-specific logic, and explicitly releases control.  A task uses coordination primitives to get and release processing time.

.. code:: c

    INT task_x ()
    {
        /*** variable declarations ***/
        INT status, msgid, msgsiz;
        CHAR *msg_p;

        /***  task initialization  ***/
        if (task_x_init () == SUCCESS)
        {
            /***  task definition ***/
            DO_FOREVER
            {
                /*** suspensive primitive ***/
                status = RECV (&msgid, &msg_p,
                            &msgsiz, SEC_1);
                switch (status)
                {
                    case SUCCESS:  /* msg ready */
                        process_msg (msgid,
                            msgsiz, msg_p);
                        break;
                    case TIMEOUT:  /* no msg    */
                        handle_timeout ();
                        break;
                    case FAILURE:  /* fault     */
                    default:
                        LOG_WARN (LY_0 + SS_0 +
                            LV_U + P0 +  0,
                                 NOT_USED);
                        break;
                }  /* end switch */
            }  /* end forever */
        }
        /***  unrecoverable fault  ***/
        LOG_FATAL (LY_0 + SS_0 + LV_U + P0 + 0, NOT_USED);
    }

The Task and Its Environment
----------------------------

TASK DISCIPLINE
^^^^^^^^^^^^^^^

To design a system, where the work unit is a task, the task discipline imposed by the kernel must be understood. The main difference between operating systems is the task discipline.

In the kernel, once a task gains control of the processor, it runs to completion.  That is, a task retains control until it explicitly releases control by calling a suspending primitive.  This discipline supports tasks that run quickly and is consistent with the requirements of hard-deadline, real-time systems.  The kernel is suited to those systems where only a small amount of time is needed to completely handle an event.  Naturally, the granularity of an event becomes a design issue.

Care should be taken to only use kernel facilities to affect a task's behavior.  For example, by using timing primitives to delay a task, instead of instructions that cause a task to busy-wait, other tasks are allowed to run while the task is waiting.  Kernel fault management guards against tasks that do not release control of the processor.

Although a task is not normally preempted by the kernel or another task, a task may be preempted by a hardware interrupt.  This is transparent to the task, however, and processing always resumes with the interrupted task.

TASK ATTRIBUTES
^^^^^^^^^^^^^^^

Task capabilities and constraints can be modified for each task, individually, by changing a task's attributes.

PRIORITY


Each task may be assigned a priority level in relation to other tasks.  Tasks that have work pending and are ready to run are given processor control according to their priority.

There are three priority levels.  The first task ready at the 
highest priority level is the next task to run.

INSTANCE


The kernel supports the concept of multiple-instance tasks.  A task instance is a task that is "known" by the kernel at run-time.  Usually, each instance is derived from a separate block of source code.  With multiple instance tasks, however, more than one run-time instance may be from the same source code.  This is useful for tasks that perform the same function on different entities, such as a task for each I/O channel or network link.  Rather than have a single task maintain separate data for each entity, each task manages one set of data.  These types of tasks are also more efficient because they do not need to continually map an entity to its data or function.

Care must be taken, however, in coding multiple instance 
tasks.  Such tasks cannot declare and reference C-language 
type static data because each run-time instance 
references these same data.  Dynamic data must be allocated 
by each task to get its own data space, usually at 
initialization time.

MEMORY AREA


Each task has its own stack and dynamic memory areas.

The stack area is required.  Any stack size may be defined, 
depending on expected stack usage.  The stack is used by all 
subroutines called from the task, primitive calls, and 
interrupts that occur while in the TASK_STATE.

Dynamic memory allocation is optional.  If a task allocates 
dynamic memory, there are no restrictions on its use.

Kernel fault management continually audits stack and 
dynamic memory to detect corruption.  Corrupted memory 
causes the task to be deleted.  While this may limit fault 
propagation and increase system availability, depending on 
the importance of the deleted task, it is likely that 
neighboring task memory is also corrupted.
RESIDENT BANK The kernel supports hardware memory bank switching to extend the 64-kilobyte address limitation.  Bank switching only applies to code space memory, however.  If the code is located in ROM, then ROM bank switching is allowed and, if the code is located in RAM, then that part of RAM with code space may be switched.  Bank switching is not supported for RAM data memory because this memory maintains the system context used by the kernel.  

The task unit is a natural unit of consideration for bank 
switching.  This is because context switching already occurs 
at the task level, so context switching to another bank is no 
different than a normal kernel context switch.

When a task is ready to run, the kernel determines the bank 
where the task resides, switches in the appropriate bank and 
resumes running the task.

A task's resident bank is defined in the configuration file.  
The entry point of the user-provided bank switching 
procedure is also defined in the configuration file.

The usual cautions apply in writing the bank switching 
procedure to be able to return to the original bank.  The 
initialization stack area, in internal RAM, is available to 
preserve information across a bank switch.  Additionally, the 
kernel and any common routines must be located at the 
same location in all banks.  Only the task code area and some 
specialized routines may be bank-dependent.

Task Identification and Creation
--------------------------------

IDENTIFICATION
^^^^^^^^^^^^^^

From a programming point of view, tasks are referenced by logical, symbolic names.  This makes the system more maintainable and extensible.  At initialization, a physical task name is assigned to each task, which is only used, directly, by the kernel.  Applications never need to use the physical name, except as a primitive parameter, and then only to make the kernel more efficient.  The physical name may be obtained with the GETTID and GETMYTID primitives.  This is usually done at task initialization because logical and physical task names never change.

CREATION
^^^^^^^^

Tasks are created by the kernel at initialization from configuration file information.  This information defines each task and its attributes, which the kernel uses to allocate task resources.

Once created, a task is usually never deleted.  Two exceptions are 1) upon a system reset, and 2) when the kernel detected a critical fault while in the TASK state.


Kernel States
-------------

This section describes the states of the kernel, for overall resource management, and the states of a task, for task resource management.

State Descriptions
^^^^^^^^^^^^^^^^^^

The kernel may be in any of four states, depending on the work to be done.  State is used primarily to determine how memory resources are to be allocated.

Immediately following a hardware or software reset, the kernel takes control of the processor.  This is the INITIALIZATION state.  In this state, the only application programs that run are the hardware and software routines defined in the system configuration file.

If there is work for a task and the task assigned to the work is available, the kernel enters the TASK state.  The intricacies of task-level processing are described in more detail below.

In the SUPERVISORY state, resources are allocated from the systems resources, rather from task resources.  This state may be entered when a hardware interrupt occurs during the TASK state.

Finally, when there is no task work and no interrupt pending, the kernel enters the MONITOR state and uses system resources.

In the TASK and SUPERVISORY states, the watchdog (COP) monitors execution time.  If the task execution time exceeds the watchdog period, the task is removed from the system and normal processing continues.  The task is never scheduled to run again.  If the execution time exceeds the watchdog period in the SUPERVISORY state, which occurs in an interrupt service routine, a system restart is initiated.  The default watchdog period is 1.049 seconds and is set in the microcontroller's OPTION register.

TASK STATES

All tasks are put in a READY state at initialization.    The first task in the task configuration table is the first task to run.  Each task runs or initializes its local data areas until it calls a suspending primitive.

Task states change from READY to ACTIVE when work becomes available for the task.  Real-time processing occurs when tasks are allocated processor time to perform work.  

In the TASK state, with a task ACTIVE, the stack area is allocated from the task's local memory.  A task runs until it calls a suspending primitive and there is no work to do or a higher priority task has work to do.  When suspended, the task is in a WAIT state, waiting for a semaphore, event, message, or timer primitive.  When a task resumes execution, processing begins at the point where the task was suspended, not necessarily at the task's entry point.

In addition to READY, ACTIVE and WAIT states, a task may be in the DORMANT state.  This occurs if a fault was detected by the kernel during the task's ACTIVE state.  Once in the DORMANT state, the task cannot run again until the system is restarted; this prevents fault propagation.

NON-TASK STATES

The SUPERVISORY state is entered by a specific request to use common system resources.  See the discussion below on handling interrupts, for more detailed information.

When there is no task work or interrupt pending, the kernel enters the MONITOR state.  This switches the active stack to the system memory area, to be ready for an interrupt, and executes a STOP instruction.  Processing resumes with the next interrupt; an interrupt always occurs under normal operating conditions with the Real-Time Interrupt (RTII).  To eventually return to the TASK state, an interrupt service routine must call a primitive that completes a task's wait condition.

Scheduling Policy
^^^^^^^^^^^^^^^^^

Tasks are prioritized in the system configuration table when the system is built.  A task's priority never changes once the system is built, which assumes processing requirements are understood before run-time.  With fixed task priorities, response times are predictable.

The next task to run is determined, whenever a task releases control of the processor, by calling a suspending primitive.  At that time, the highest priority task in the READY state is run.  When multiple equal priority tasks are ready to run, the task that first became ready is the first to run.

Tasks may be preempted by interrupts and the interrupt service routine may call a primitive that causes a waiting task to become READY.  Processing always resumes with the interrupt task.

Initialization and Startup
^^^^^^^^^^^^^^^^^^^^^^^^^^

The kernel does the basic hardware and software initialization needed to run the kernel in the target processor environment.  It then automatically invokes initialization routines, provided by the developer, to initialize application-specific hardware and software.

The application-specific routines are the only initialization routines that need to be provided.  Their entry points are specified in the configuration file when the system is built.

Also, configurable microcontroller parameters that need to be set at initialization are defined in the configuration file.

The kernel provides orderly initialization sequencing by initializing from lower to higher levels of abstraction.  First, the hardware, then kernel, then application hardware and software are initialized; the idea is the hardware must be available for the kernel to run, and the hardware and kernel must be available for the application to run.

Entries in the configuration file provide hooks for the kernel to initialize application hardware and software.  The kernel initializes application hardware before software.

At each step, the integrity of the supporting layer is confirmed before the next layer is initialized.  If resources are not available for a layer to provide the necessary services, subsequent layers are not initialized.

Faults that occur before kernel resource allocation and initialization are complete are considered critical because the integrity of the system cannot be guaranteed.  If a critical fault is detected, a STOP instruction is executed to prevent fault propagation.

Handling Interrupts
^^^^^^^^^^^^^^^^^^^

Conceptually, interrupt service routines are very similar to tasks in that they run as a result of an external stimulus that changes the context of the system.  In the case of tasks, the stimuli may be messages, events, or semaphores, while in the case of an interrupt service routine, the stimulus is a hardware event detected by the microcontroller.

Task-related context switches that change resource allocation needs are managed by the kernel.  However, there is no kernel to manage hardware-initiated context switches in the same way.  Interrupt service routines must, therefore, explicitly, interface with the tasks' kernel to coordinate resource allocation.  A SUPERVISORY state is implemented in the kernel for interrupt management.

SUPERVISORY state is entered by calling the primitive, ENTER_SSTATE.  In SUPERVISORY state, the interrupt uses the stack area from the system's memory area.  Upon exiting from the interrupt, by calling EXIT_SSTATE, the interrupted task resumes at the point of interruption, again, using the task's memory area stack.

The kernel supports nested interrupts.  That is, an interrupt may occur, and call ENTER_SSTATE, while another interrupt service routine is in progress.  Control does not return to the task until EXIT_SSTATE is called by the last interrupt service routine.

Tasks may become READY to run as a result of a primitive called by the interrupt service routine, which completes a task's WAITING primitive.

Kernel Primitives
-----------------

This section discusses the concepts associated with primitives.  A detailed description of each kernel primitive is found in Part 3.

kernel functions are called primitives because they are the most basic function for interacting with tasks and requesting kernel services.

Primitives are called during task or interrupt service routine execution.  While primitive code and static data are in the kernel code and data space, primitives use the calling task or interrupt service routine stack area.
