
CPU Performance
=========================================================================

To estimate system performance, this section gives reference times for various kernel facilities.  All times are for an 8 MHz crystal (500 nanosecond instruction cycle time).

Context Switch Timing
---------------------

The time needed to switch from a task that has released control of the processor to the next ready task is 28 microseconds.

This includes the time needed for data integrity checks and to pulse the watchdog.

This number applies if the next task to run is a high priority task.  For each lower priority level, an additional six microseconds are needed.

When an interrupt occurs, the time needed to switch to the interrupt service routine, regardless of whether the kernel is in the task or supervisor state, is 13 to 44 microseconds.  The variance depends on the instruction in progress when the interrupt occurred.

This timing includes the primitive call to ENTER_SSTATE (the interrupt service routine is considered entered immediately upon return from the primitive call) and is slightly less if the kernel is already in supervisor state.

Primitive Timing
----------------

Fifteen microseconds are needed to access any primitive, and most primitives complete in less than 50 microseconds.

However, a look at the messaging and event management primitives gives an indication of actual, useful work capability for those primitives.

The time between the call to SEND a message and when the RECV call is completed, for two high priority tasks, is 476 microseconds; i.e., about 2,100 messages per second.

For each lower task priority level, about six microseconds more is needed to send and receive a message.  Messages do not have a priority.

Similarly, the time between the call to SIGNAL an event and when the WAIT for the event is completed, for two high priority tasks, is 422 microseconds.  This equates to about 2,369 events per second.  As with messaging, about six microseconds is added for each lower task priority level.
