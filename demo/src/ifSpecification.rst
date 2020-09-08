Interface Specification
=========================================================================

This section gives a detailed description of the kernel primitives used to invoke kernel services.

For each primitive, a functional description is given along with operational considerations particular to the primitive.  This is followed by a description of the actual subroutine call.

The description shows the C and assembly language macro interface with the required parameters and return value.  All parameters are required, although, parameters not applicable for a particular call contains a place-holding literal such as NOT_USED.

For the assembly language interface, all registers are preserved, except the D register, which returns the primitive status.  The returned D-register status has the same meaning as the C- language interface status.

For both C and assembly language, parameters shown in upper case are literal constants, while those in lower case are memory address references.

GETTID Get Task Identifier
----------------------------

DESCRIPTION:

    This primitive is used to get a task's identifier, which is used with other primitives that reference a task.

    Given a logical task name, defined at compile time, this primitive returns the run-time identity of the task.

CONSIDERATIONS:

    The task name is a number between zero and the maximum allowed number of tasks, minus one.  The maximum number of tasks in the kernel is 64.  The task name is assigned in the configuration file.   

    The task identifier is a pointer variable to the task's control data structure, however, the application never needs to reference the data structure in normal operation.

    The task must exist in the kernel's task data area. The task may not exist if this primitive is called from a procedure before tasks are created or if the task was terminated as a result of an error detected by the kernel.  In these cases, INV_ADDR is returned as the task identifier.

    Because a task's identity never changes once it is created, this primitive only needs to be called once, usually during initialization.

    This primitive also checks the task's data structures for consistency.

    This is a non-suspending primitive.

C ACCESS:

    status = GETTID (tname, tid_pp);


PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| tname              | CHAR         | Task logical name                                                     |
+--------------------+--------------+-----------------------------------------------------------------------+
| tid_pp             | BYTE **      | Task id location                                                      |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Task id available

        FAILURE Invalid task name, task does not exist or inconsistent data

ASSEMBLER ACCESS:

    
    %GETTID #TNAME, #tid_p


GETMYTID Get Own Task Identifier
----------------------------------

DESCRIPTION:

    This returns the identity of the currently active task.  It operates the same as GETTID, except the task name is not needed.

CONSIDERATIONS:

    The task that invokes this primitive is the currently active task.

    This primitive also verifies the integrity of the task's data structures.

    Because the task's identity does not change after it is created, this primitive 
    only needs to be called once.

    This is a non-suspending primitive.

C ACCESS:

    status = GETMYTID (tid_pp);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| tid_pp             | BYTE **      | Task id location                                                      |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Task id available

        FAILURE Task data structure fault detected

ASSEMBLER ACCESS:

    %GETMYTID   #tid_p

GET_CRID Get Resource Identifier
----------------------------------

DESCRIPTION:

    This primitive gives a resource's identifier given the resource's name.

    The identifier is used for the ENTER_CR and EXIT_CR primitives, which 
    allow mutually exclusive access to the resource.

CONSIDERATIONS:

    Resource names are assigned in the system configuration file.  The name 
    is a consecutive number, beginning with zero and ending with one less 
    than the number of resources allocated.  Designers may define literals to 
    associate more meaningful names with a resource. 

    The name is used by this primitive to map to the semaphore data 
    structures associated with the resource, however, these data structures 
    never need to be referenced by an application.

    Because a resource identifier does not change once it is created by the 
    kernel, this primitive only needs to be called once.

    This is a non-suspending primitive.

C ACCESS:

    status = GET_CRID (cr_name, crid_p);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| cr_name            | INT          | Resource logical name                                                 |
+--------------------+--------------+-----------------------------------------------------------------------+
| crid_p             | BYTE **      | Resource id location                                                  |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Resource id available

        FAILURE Invalid resource name

ASSEMBLER ACCESS:
    
    %GET_CRID   #crname, #crid_p

ENTER_CR Get Access to Resource
---------------------------------

DESCRIPTION:

    This primitive provides for synchronization on a user-defined resource using a semaphore and guarantees mutually exclusive access to a resource.

CONSIDERATIONS:

    The resource id must first be obtained using the CR_GETID primitive.

    Resource access is managed by a semaphore, and tasks are queued to the 
    semaphore on a first-come-first-serve basis.  If the resource is not locked 
    by another task, the resource is locked and the calling task continues as 
    the active task.  If the resource is not available, the calling task is 
    suspended.  The task remains suspended until all previous tasks have 
    released exclusive access to the resource.

    If the current task has already locked the resource, this primitive has no 
    effect, although an error status is returned.

    This primitive may not be called from an interrupt service routine.

    A resource is released by the EXIT_CR primitive.

C ACCESS:

    status = ENTER_CR (cr_id);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| cr_id              | BYTE *       | Resource id                                                           |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Access granted and resource locked

        FAILURE Invalid resource id or task has already acquired resource.

ASSEMBLER ACCESS:

    %ENTER_CR   #crid

EXIT_CR Release Access to Resource
------------------------------------

DESCRIPTION:

    Release control of a resource that was obtained with the ENTER_CR 
    primitive.

CONSIDERATIONS:

    The resource identifier must first be obtained using the CR_GETID 
    primitive, and must be the same identifier used to acquire the resource 
    with ENTER_CR.  Only the task that currently has access to the resource 
    may unlock the resource.

    When the resource is released, if a higher priority task is waiting on the 
    resource, the current task is suspended and the higher priority task 
    becomes the active task.  The suspended task becomes active according to 
    its assigned priority.

    An error status is returned if the resource was not previously locked, an 
    invalid resource id was given or corrupted data structures were detected. 

C ACCESS:

    status = EXIT_CR (cr_id);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| cr_id              | BYTE *       | Resource id                                                           |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Resource released

        FAILURE Invalid resource id, resource not acquired with previous ENTER_CR by this task or corrupted kernel data detected.

ASSEMBLER ACCESS:

    %EXIT_CR    #crid

ENTER_SCR Get Exclusive Access to Processor
---------------------------------------------

DESCRIPTION:

    Obtain mutually exclusive access to the processor.

CONSIDERATIONS:

    This primitive is implemented by disabling maskable interrupts, and is 
    not associated with a particular resource.  Access is released by 
    EXIT_SCR.

    This primitive is recommended for mutually exclusive access of short 
    Duration because kernel time-related function is affected. 

    The processor condition code register is preserved upon return from the 
    EXIT_SCR primitive.

    This is a non-suspending primitive.

C ACCESS:

    status = ENTER_SCR ();

PARAMETERS:  none

RETURN:  none

ASSEMBLER ACCESS:

    %ENTER_CR



EXIT_SCR Release Exclusive Access to Processor

DESCRIPTION:

    This primitive is the complement of ENTER_SCR and releases mutually exclusive control of the processor.

CONSIDERATIONS:

    The processor condition code register at the time ENTER_SCR was called 
    are restored.  

    This primitive must not be called without a previous call to ENTER_SCR.

C ACCESS:

    status = EXIT_SCR ();

PARAMETERS:  none

RETURN:  none

ASSEMBLER ACCESS:

    %EXIT_CR

SIGNAL Signal Event Occurrence
--------------------------------

DESCRIPTION:

    This primitive provides a synchronization mechanism by signaling a task that one or more events occurred.

CONSIDERATIONS:

    The task to be signaled must exist and may be the current task.

    The event_id is a user-defined bit map with each bit corresponding to an 
    Event. Events may be defined between tasks for a total of 16 events per 
    task, or at the system level for a total of 16 system events.  Event 
    agreement between tasks is a design issue.

    If a higher priority task is waiting for the event(s), and all wait criteria are 
    satisfied, the signaling task is suspended.  The task resumes as the active 
    task, according to its assigned priority.

    If all wait criteria of the signaled task are not satisfied, the event is posted 
    to the signaled task, whether or not the task has a WAIT request pending.

    The task identifier of the task to be signaled must be obtained using the 
    GETTID or GETMYTID primitives.

C ACCESS:

    status = SIGNAL (tid, event_id);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| tid                | BYTE *       | Task identifier                                                       |
+--------------------+--------------+-----------------------------------------------------------------------+
| event_id           | WORD         | Event identifier                                                      |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Task signalled

        FAILURE Task does not exist or corrupted data detected
  
ASSEMBLER ACCESS:

    %SIGNAL #taskid, #EVENTID

WAIT Wait for Event Occurrence
--------------------------------

DESCRIPTION:

    This primitive complements SIGNAL and allows a task to wait for one or more events to occur.

CONSIDERATIONS:

    The event_id is a user-defined bit map with each bit corresponding to an 
    event; events may be defined between tasks or at the system level, and 
    agreement between tasks is a design issue.    

    The calling task may request a wait for all specified events to occur (AND-
    conditional), or for any one of the events to occur (OR-conditional).

    A WAIT condition is considered satisfied when, either a logical AND 
    condition was specified and all requested events were received, or a 
    logical OR condition was specified and any of the requested events was 
    received.

    All events are cleared, including those that were not used to complete the 
    wait, when the task is resumed.

    This primitive may not be called from an interrupt service routine.

C ACCESS:

    status = WAIT (evt_desc, e_logic, tout_val);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| evt_desc           | WORD         | Event(s) specification bit map                                        |
+--------------------+--------------+-----------------------------------------------------------------------+
| e_logic            | INT          | EVT_OR  OR condition                                                  |
|                    |              | EVT_AND AND condition                                                 |
+--------------------+--------------+-----------------------------------------------------------------------+
| tout_val           | INT          | Suspension timeout                                                    |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Requested event(s) occurred

        FAILURE Invalid parameter

        TIMEOUT Timeout occurred before requested event(s)

ASSEMBLER ACCESS:

    %WAIT   #EVENTID, #EVENTLOGIC, #TIMEOUT

SEND Send Message to Task
---------------------------

DESCRIPTION:

    This primitive provides intertask communication, using messages, as a task synchronization mechanism.

CONSIDERATIONS:

    The destination task must exist, and receives the message with the RECV primitive.  The destination task may be the current task.

    If a higher priority task is waiting for a message, the sending task is suspended.  The sending task resumes as the active task according to its assigned priority.

    Multiple messages may be queued to a task, and are serviced with a first-in-first-out discipline.  Also, messages are posted to a task regardless of whether or not a RECV is currently pending. 

    This is an asynchronous operation because the sending task does not wait for a response from the destination task.  End-to-end confirmation of message delivery is done by the application.

    Message identifiers and message content are defined between message origination and destination tasks at design time. 

    No message data is copied during the message transfer, only the pointer to the message is passed to the destination task; it is the responsibility of both tasks to coordinate allocation and freeing of the message data area.  The message location and size parameters refer only to application message data.  Memory does not need to be reserved for message management because this is handled by the kernel.

    The total number of messages that may be active at any given time are defined in the system configuration file.  

C ACCESS:

    status = SEND (tid, msgid, msg_p, msgsiz);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| tid                | BYTE *       | Destination task id                                                   |
+--------------------+--------------+-----------------------------------------------------------------------+
| msgid              | INT          | Message id tag                                                        |
+--------------------+--------------+-----------------------------------------------------------------------+
| msg_p              | CHAR *       | Message location                                                      |
+--------------------+--------------+-----------------------------------------------------------------------+
| msgsiz             | WORD         | Message length                                                        |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Message sent to task

        FAILURE Unable to send message; invalid destination or resources not available

ASSEMBLER ACCESS:

    %SEND   #taskid, #MSGID, #msg_p, #msgsiz

RECV Wait for Message from Task
---------------------------------

DESCRIPTION:

    This primitive complements the SEND synchronization primitive to receive a message from a task.

CONSIDERATIONS:

    The receiving task continues to receive any queued messages as long as they are available, without suspension, until a task of higher priority becomes ready to run.  Care should be taken in system design to ensure control is released to the scheduler within the watchdog timer period.

    When the task resumes after the RECV call, the received message is removed from the task's message queue and any timeout request is canceled. 

    If a timeout occurs before a message is received, the task becomes active according to its assigned priority.

    The message address returned is the location of the message provided by the task that sent the message.  The message area is managed by the sending and receiving tasks.

    Message id and message content agreement between origination and destination tasks is defined at compile time.

    This primitive may not be called from an interrupt service routine.

C ACCESS:

    status = RECV(msgid_p, msg_pp, msgsiz_p, tout_val);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| msgid_p            | INT *        | Address for message id tag                                            |
+--------------------+--------------+-----------------------------------------------------------------------+
| msg_pp             | CHAR **      | Address for message location                                          |
+--------------------+--------------+-----------------------------------------------------------------------+
| msgsiz_p           | WORD *       | Address for message length                                            |
+--------------------+--------------+-----------------------------------------------------------------------+
| tout_val           | INT          | Suspension timeout                                                    |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Message available

        FAILURE Parameter error

        TIMEOUT Timeout occurred      

ASSEMBLER ACCESS:

    %RECV   #msgid_p, #msg_pp, #msgsiz_p, #TIMEOUT

ALERT Request Wakeup after Time Interval
------------------------------------------

DESCRIPTION:

    Request for the task to be signaled after a specified amount of time has elapsed.

CONSIDERATIONS:

    This is a suspending primitive.

    When the timeout occurs, the task is scheduled to run according to its assigned priority. 

    There is no provision to cancel an alert request before the timeout occurs.  

    If NO_TOUT is requested, this primitive has no effect and the task continues as the active task.

    The time interval is associated with the real-time clock interrupt rate and is accurate to within one clock increment.

    The alert request applies to the currently active task and may not be called from an ISR.

    The time interval parameter is expressed in 100-millisecond units.  The minimum time interval that may be requested is 100 milliseconds and the maximum is 6553.5 seconds.  Accuracy is + 100 milliseconds.

C ACCESS:

    status = ALERT (tout_val);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| tout_val           | INT          | Elapsed time value in 100 millisecond increments                      |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT TIMEOUT Requested time has expired

        FAILURE Unable to handle alert request

ASSEMBLER ACCESS:

    %ALERT  #TIMEOUT

GETTIK Get Current System Timer Value
---------------------------------------

DESCRIPTION:

    Get the current system time counter value.

CONSIDERATIONS:

    The value is a continuous, 16-bit, 100 millisecond counter.

C ACCESS:

    status = GETTIK (tikval_p)


PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| tikval_p           | WORD *       | Current timer value location                                          |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Timer value available

        FAILURE Kernel fault

ASSEMBLER ACCESS:

    %GETTIK #tikval_p

Q_GETID Get Queue Identifier
------------------------------

DESCRIPTION:

    Get the linked-list queue identifier, given a queue name.  The identifier is used for queue access primitives.

CONSIDERATIONS:

    The queue is a linked list, with a first-in-first-out discipline.

    The queue name is a value between zero and 255.  These are defined in the configuration file, beginning with zero and progressing sequentially for the number of queues defined.

    The queue control blocks are allocated in the kernel address space.  The only data structure requirement imposed on the application is that the first two bytes of the queued item be reserved for linked list management.  However, this pointer should never be written to by the application.

C ACCESS:

    status = Q_GETID (q_name, qid);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| q_name             | INT          | Queue name                                                            |
+--------------------+--------------+-----------------------------------------------------------------------+
| qid                | CHAR *       | Queue id                                                              |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Queue id available

        FAILURE Invalid queue name

ASSEMBLER ACCESS:

    %Q_GETID    #QNAME, #qid

Q_CLEAR Initialize Queue
--------------------------

DESCRIPTION:

    Initialize the specified queue

CONSIDERATIONS:

    The queue control block pointers are reset to indicate there are no queued items.

C ACCESS:

    status = Q_CLEAR (qid);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| qid                | CHAR *       | Queue id                                                              |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Queue initialized

        FAILURE Invalid queue id

ASSEMBLER ACCESS:

    %Q_CLEAR #qid



Q_GET Get Item from Linked-list Queue

DESCRIPTION:

    Get the next item from a linked list queue.

CONSIDERATIONS:

    Items returned from the queue have a link pointer of type BYTE * as the first two bytes of the returned item.

    Queue discipline is first-in-first-out.

    Exclusive access is guaranteed to the calling task, and the task is not suspended if there are no queued items.

C ACCESS:

    status = Q_GET (qid, item_pp);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| qid                | CHAR *       | Queue id                                                              |
+--------------------+--------------+-----------------------------------------------------------------------+
| item_pp            | CHAR **      | Address for queue item                                                |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Item dequeued

        FAILURE Invalid queue id

        LIMIT   Queue is empty

ASSEMBLER ACCESS:

    %Q_GET #qid, #item_pp

Q_PUT Add Item to Linked-list Queue
-------------------------------------

DESCRIPTION:

    Add an item as the last entry of a linked list queue.

CONSIDERATIONS:

    The queue is a linked list.  Items to be queued must reserve a link pointer of type BYTE * as the first two bytes of the item to be queued.

    Queue discipline is first-in-first-out.

    The calling task is guaranteed exclusive access to the queue, and is not suspended.

    There is no limit on the number of items that may be queued, other than the availability of memory for items to queue.

    An error status is returned if an attempt is made to queue an item more than once to any queue.

C ACCESS:

    status = Q_PUT (qid, item_p);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| qid                | CHAR *       | Queue id                                                              |
+--------------------+--------------+-----------------------------------------------------------------------+
| item_p             | CHAR *       | Queue item location                                                   |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Item queued

        FAILURE Invalid queue id or item is already queued.
     
ASSEMBLER ACCESS:

    %Q_PUT #qid, #item_p

LOCATE_MEM Locate Task Dynamic Memory
---------------------------------------

DESCRIPTION:

    Get the location of a dynamic memory partition allocated to the current task.

CONSIDERATIONS:

    This memory partition is never used by the kernel, unlike the tasks' stack area, and is only made known to the task assigned to the partition.  Once the task locates its memory, there are no restrictions on its use and it may be made available to other tasks.

    This primitive only needs to be called once, preferably at task initialization time.

    Higher level memory management functions are the responsibility of the task; memory is never released through the kernel.

    The memory is defined in the configuration file, by specifying the memory size needed by the task.  The requested size is guaranteed to the task at run-time, although the location is determined by the kernel from available memory.

C ACCESS:

    status = LOCATE_MEM (mem_pp);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| mem_pp             | CHAR **      | Address for task's dynamic memory location; INV_ADDR is               |
|                    |              | returned if memory has not been allocated for the task.               |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Memory location available

        FAILURE Memory not allocated for task

ASSEMBLER ACCESS:

    %LOCATE_MEM #mem_pp

LOG_FATAL Indicate Fatal Fault Occurrence
-------------------------------------------

DESCRIPTION:

    Log fatal-type fault and initiate recovery.

CONSIDERATIONS:

    If a task has been defined in the configuration file as a fault handler, the task is signaled with EVT_0, to indicate a fatal fault occurred, provided the fault did not occur in the fault handler task.

    If a task detects and reports a fatal-type fault or the kernel detects a fatal-type fault in the task domain, the task is removed from the list of available tasks and is never scheduled to run again.  In this case, this primitive returns to the scheduler and the system continues to run, as much as possible without the affected task.

    If a fatal-type fault is detected in the kernel domain or in hardware on which the kernel is dependent, system integrity cannot be guaranteed and a STOP instruction is executed.  Upon reset, the fault is signaled to the fault handler task, if one was defined.

    In all cases, fault-related information is logged to the fault analysis area for future reference.  This includes fault location, fault-specific data, task and kernel stack areas, and kernel and task states.  These data are preserved until the next fault, even through a system reset.

C ACCESS:

    status = LOG_FATAL (loc, qual);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| loc                | WORD         | Fault location code                                                   |
+--------------------+--------------+-----------------------------------------------------------------------+
| qual               | WORD         | ault qualifier                                                        |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:  Not applicable

ASSEMBLER ACCESS:

    %LOG_FATAL #LOC, #QUAL

LOG_WARN Indicate Non-fatal Fault Occurrence
----------------------------------------------

DESCRIPTION:

    Report a warning-type fault.

CONSIDERATIONS:

    This primitive is used if detected faults are notable but may not compromise the system.

    For warning-type faults, only the fault location and fault-specific data, if any, are logged to the Fault Analysis Area.  If a task has been defined in the configuration file as a fault handler, the task signaled with EVT_1, to indicate a warning-type fault occurred, and the task may query the fault analysis area.

C ACCESS:

    status = LOG_WARN (loc, qual);

PARAMETERS:

+--------------------+--------------+-----------------------------------------------------------------------+
| NAME               | TYPE         | DESCRIPTION                                                           |
+====================+==============+=======================================================================+
| loc                | WORD         | Fault location code                                                   |
+--------------------+--------------+-----------------------------------------------------------------------+
| qual               | WORD         | ault qualifier                                                        |
+--------------------+--------------+-----------------------------------------------------------------------+

RETURN:

                                                      

    INT SUCCESS Fault is logged
          
ASSEMBLER ACCESS:

    %LOG_WARN #LOC, #QUAL


ENTER_SSTATE Enter Supervisory State

DESCRIPTION:

    This primitive is the mechanism for an interrupt service routine invoke supervisory state processing.

CONSIDERATIONS:

    This primitive allows the calling interrupt service routine and nested interrupt service routines to switch to the common system stack.

    The last interrupt service routine to exit the supervisory state returns control to the interrupted task and switches to the task stack.

    For every ENTER_SSTATE call, there must be a matching EXIT_SSTATE call.

C ACCESS:

    ENTER_SSTATE ();

PARAMETERS:  none

RETURN:  none
          
ASSEMBLER ACCESS:

    %ENTER_SSTATE

EXIT_SSTATE Exit Supervisory State
------------------------------------

DESCRIPTION:

    This primitive complements the ENTER_SSTATE primitive, and returns to the task state from the supervisory state.

CONSIDERATIONS:

    For every ENTER_SSTATE call, there must be a matching EXIT_SSTATE call.

    Interrupts may be nested and ENTER_SSTATE primitive calls may be nested.  If EXIT_SSTATE is called by an interrupt service routine that is not the last interrupt service routine pending, processing continues in the supervisory state.

    The last interrupt service routine to exit the supervisory state returns control to the interrupted task and switches to the task stack.

C ACCESS:

    EXIT_SSTATE ();

PARAMETERS:  none

RETURN:  none
          
ASSEMBLER ACCESS:

    %EXIT_SSTATE