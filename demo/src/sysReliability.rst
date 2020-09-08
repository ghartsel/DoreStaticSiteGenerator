System Reliability
=========================================================================

Reliability Issues
------------------

Real-time systems are particularly error-prone because of their inherent complexity; as complexity increases the probability of errors also increases.

This problem is compounded by the requirement for embedded systems to operate without intervention, possibly in the presence of errors.

Systems may be classified as fault-tolerant or fault-intolerant.  Fault-tolerant systems are characterized by anticipating that all errors cannot be removed before run-time, and mechanisms are provided to detect and handle faults when they do occur.  Complex, real-time systems, usually, cannot be tested adequately to guarantee that no faults occur.  Fault intolerant systems, on the other hand, assume that the system is sufficiently tested, no faults occur, and no fault-handling is provided.

The kernel is fault-tolerant.  By implementing fault management, the kernel achieves the objectives of increased system availability and assures algorithm correctness.

Reliability by Design
---------------------

The kernel allows reliability to be designed into a system.

One method of fault management that is implemented in the system design is software layering.  A layered system allows it to be viewed in smaller, logically related, and less complex segments.  This makes the design more manageable and, by extension, more reliable.  Layering is also a facility for the isolation of faults.  If a fault can be isolated to a layer, then only that layer needs to be recovered, and system availability is increased.

The kernel, itself, may be considered a single layer.  However, because it is the most dependent layer faults detected in the kernel usually require complete system recovery.  The kernel may not need to be recovered if faults are detected in an application layer.

The kernel supports reliable system design by its built-in fault management policies and by providing mechanisms for applications to interface to the fault management system.  These policies and mechanisms are described in the sections, below.

Fault Management
----------------

Fault management has three aspects; 1) fault detection, 2) fault handling and reporting, and 3) fault recovery.

The kernel uses its fault analysis area for fault management.  These data are described in detail in a separate section.

FAULT DETECTION
^^^^^^^^^^^^^^^

By active fault detection, the kernel can reduce fault propagation.  This is the primary mechanism that supports fault-tolerant objectives.  The following fault detection methods are used by the kernel.

Memory Audits
#############

The kernel partitions and classifies memory for fault 
management.  If the type of memory is known, the kernel can verify that the data is 
consistent with the memory type.  While this method does not guarantee fault 
detection, common types of memory corruption are detectable.

System and task stacks, kernel data structures and kernel structures in the application memory area, such as 
queues, are continuously audited.

Guarded Primitive Access
########################

Whenever a primitive is called, the kernel test that is a valid primitive, that the 
parameters are consistent and within kernel defined limits, and that resources are 
available.

Software Watchdog
#################

The hardware watchdog (COP) is used to detect tasks and interrupt service 
routines that keep control of the processor for longer than the watchdog period.  
This is an indication of faulty algorithm execution.

The above fault detection methods are used at run-time.  To reduce development time by detecting errors early in the process, system build tools are part of kernel fault management.  This supports the notion that faults detected early in the development cycle are less expensive to correct.

Fault detection before run-time has the added benefit of freeing the kernel from some fault detection at run-time, resulting in a more efficient kernel.

The following fault detection methods are used during implementation.

Primitive Declaration
#####################

The macro declaration of the primitive interface allows most assemblers and 
compilers to detect reference and parameter errors in the primitive call.

System Build Utility
####################

This is the primary tool for fault management before run-time that relates directly 
to the kernel.  The system build utility, described in detail in a separate section, 
tests for overall resource availability and data consistency.

High-Level Language Interface
#############################

The high-level language interface allows applications to be implemented with the 
probability of fewer errors.  

FAULT HANDLING AND RECOVERY
^^^^^^^^^^^^^^^^^^^^^^^^^^^

Once faults are detected they must be handled to achieve fault tolerance.  A completely handled fault includes determination of fault severity, reporting and logging the fault, and initiating recovery action.  Recovery action depends on the type and severity of the fault.

Kernel Fault Handling
#####################

Kernel-detected faults are only acted on if they are determined to be critical to
system operation.  Otherwise, the fault is only reported and logged.

Fault classification, either critical or acceptable, depends on the kernel state and
the scope of the data if it is a data-related fault.  Acceptable faults, from the
kernel's point of view, are those that are limited to a single task.  Such faults occur
in the TASK state or in a task's stack or dynamic data.  Faults in any other
kernel state are classified as critical.

For acceptable faults, the kernel simply deletes the affected task, so the fault is not
repeated or propagated.  The fault is reported to the fault handling task if one is
defined.  It is left to the application to determine the impact of the deleted task and
take appropriate recovery action.

For critical faults, the kernel determines that system integrity cannot be 
guaranteed.  The kernel, therefore, reports the fault, then initiates a kernel restart.  
The restart causes all tasks to be restarted.  The cause of the restart is available in 
the fault analysis area.

Application Fault Handling
##########################

Application-detected faults are independent of the kernel.  Their classification and 
recovery action depend on the application.  Applications use kernel primitives to 
report the fault and initiate recovery.

The LOG_WARN primitive is used to report non-critical faults.  This is useful for 
noteworthy faults that do not affect system operation and is useful for check-
pointing during debugging.  The fault is only reported and logged; no recovery is 
initiated by the kernel.

The LOG_FATAL primitive is used to report faults the application determines 
adversely affect system operation.  The kernel reports the fault and initiates 
recovery by restarting the kernel, which also restarts the tasks.

Both LOG_WARN and LOG_FATAL cause fault-related information to be logged 
in the kernel's fault analysis area.  This information is available at run-time to 
report the fault and is useful during debugging to determine the cause of the fault.

Optionally, the fault may be reported to a fault handling task.  The task name must 
first be defined in the configuration file.  This allows the application to take 
application-dependent recovery action.

For primitive access, the kernel detects some classes of errors before a fault occurs. An example is primitive parameter errors.  The kernel reports these faults in the status returned by the primitive.  The kernel always returns a status to a primitive call.  It is the responsibility of the application to check the return value and take the appropriate action.