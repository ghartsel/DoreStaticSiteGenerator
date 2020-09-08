General Definitions
=========================================================================

Suspensive Primitives
---------------------

There are two types of primitives.  Those that may cause a task to be preempted, these are called suspending primitives, and those that return to the caller without preemption, called non- suspending primitives.  The suspending primitives are,

+-------------+------------------------------------------------------------+
| PRIMITIVE   | SUSPENDING CONDITION                                       |
+=============+============================================================+
| ENTER_CR    | Critical region is not available.                          |
+-------------+------------------------------------------------------------+
| RECV        | No message is pending to task.                             |
+-------------+------------------------------------------------------------+
| WAIT        | No event is pending to task for requested event condition. |
+-------------+------------------------------------------------------------+

The above suspending primitives have complementing primitives which, when called from another task or interrupt service routine, cause the suspended task to resume.  A task resumes when the suspending condition is satisfied.  The complementing primitives in corresponding order are EXIT_CR, SEND, and SIGNAL.

Suspending primitives may not be called from an interrupt service routine because an interrupt may not be suspended by the kernel.  Suspension supports real-time applications at the task level.

While tasks may be prioritized, primitives do not have a priority attribute.  For example, there is no priority assigned to messages for the SEND and RECV primitives.  All prioritization is considered with respect to tasks.

Common Primitive Attributes
---------------------------

This section describes the kernel interface conventions.  These conventions have been defined to give a consistent interface, which makes application software easier to implement and maintain.

Primitive Access
----------------

The primitive interface to the kernel is invoked as an assembly or C-language subroutine call.  Macros, defined in include files IF_K_M.H and IF_K_M.INC implement the actual access to the kernel.

All kernel primitives return a function status.  This is the status of the execution of the algorithm for the primitive; either SUCCESS or FAILURE or some specialized meaning.  No data associated with the primitive are ever returned as a status.  Data are passed and returned as parameters.

Literals Used with Primitives
-----------------------------

Include files provide a common, portable interface to the kernel.  Include files IF_L.H, IF_L.INC, IF_K_L.H, and IF_K_L.INC contain the literals used with the kernel primitives described in this section.  These include files may be changed to adapt the kernel to the application.  Those literals that may be changed are indicated with an asterisk (*).

These literals provide a method to reference kernel attributes symbolically.  This higher level of abstraction makes the interface more maintainable and portable.

+---------------+--------------------------------------------------------------------------------------------------+
| GENERAL       | LITERAL MEANING                                                                                  |
+===============+==================================================================================================+
| TRUE          | General literal for TRUE conditional.                                                            |
+---------------+--------------------------------------------------------------------------------------------------+
| FALSE         | General literal for FALSE conditional.                                                           |
+---------------+--------------------------------------------------------------------------------------------------+
| DO_FOREVER    | This is used within a task to define the start of the main task execution loop.                  |
+---------------+--------------------------------------------------------------------------------------------------+
| DO_NOTHING    | A statement place-holder, which generates no code.                                               |
+---------------+--------------------------------------------------------------------------------------------------+
| INV_ADDR      | This literal is used to set pointer variables to an invalid address, which may be used for fault |
|               | detection.                                                                                       |
+---------------+--------------------------------------------------------------------------------------------------+
| NOT_USED      | This is a parameter literal to indicate a parameter is not used or has no significance.          |
+---------------+--------------------------------------------------------------------------------------------------+
| SUCCESS       | This value is returned by primitives to indicate no errors occurred.                             |
+---------------+--------------------------------------------------------------------------------------------------+
| FAILURE       | If an error occurred during primitive execution, a FAILURE status is returned.                   |
+---------------+--------------------------------------------------------------------------------------------------+
| TIMEOUT       | This literal is returned by primitives to indicate a timeout condition occurred while waiting    |
|               | for the requested operation.                                                                     |
+---------------+--------------------------------------------------------------------------------------------------+
| NO_TOUT       | If no timeout is desired, this literal is used for the time-specification parameter.             |
+---------------+--------------------------------------------------------------------------------------------------+

Task Attributes
^^^^^^^^^^^^^^^

+-------------------+----------------------------------------------------------------------------+
| LITERAL           | MEANING                                                                    |
+===================+============================================================================+
| MAX_TSK(*)        | For task management operations, this defines the maximum number of         |
|                   | task instances, including multiple-instance tasks, that may be referenced. |
|                   | The range is 0 to 63.                                                      |
+-------------------+----------------------------------------------------------------------------+
| MAX_TNAME         | For task management operations, this defines the limit for task names.     |
|                   | The range is 0 to 63.                                                      |
+-------------------+----------------------------------------------------------------------------+
| INV_TNAME         | Invalid task name indication.                                              |
+-------------------+----------------------------------------------------------------------------+
| INV_TID           | Invalid task identifier.                                                   |
+-------------------+----------------------------------------------------------------------------+
| STKSIZ_L          | (*) Stack sizes:  large                                                    |
| STKSIZ_M          |                   medium                                                   |
| STKSIZ_S          |                   small                                                    |
+-------------------+----------------------------------------------------------------------------+
| MEMSIZ_L          | (*) Dynamic memory sizes:  large                                           |        
| MEMSIZ_M          |                            medium                                          |        
| MEMSIZ_S          |                            small                                           |      
+-------------------+----------------------------------------------------------------------------+
| PRI_HIGH          | Task relative priorities:  high                                            |    
| PRI_NORM          |                            normal                                          |     
| PRI_NONE          |                            none                                            |  
+-------------------+----------------------------------------------------------------------------+
| BANK_0            | Task resident bank location:  bank 0                                       |      
| BANK_1            |                               bank 1                                       |     
| BANK_2            |                               bank 2                                       |    
| BANK_3            |                               bank 3                                       |   
+-------------------+----------------------------------------------------------------------------+

Queue Management
^^^^^^^^^^^^^^^^

+-------------------+----------------------------------------------------------------------------+
| LITERAL           | MEANING                                                                    |
+===================+============================================================================+
| MAX_QUE           | For queue management operations, this defines the maximum number of queues |
|                   | that may be referenced.  The range is 0 to 255.                            |
+-------------------+----------------------------------------------------------------------------+
| INV_QNAME         | Invalid queue name reference.                                              |
+-------------------+----------------------------------------------------------------------------+
| INV_QID           | Invalid queue identifier.                                                  |
+-------------------+----------------------------------------------------------------------------+
| Q_FULL            | Queue is full indication.                                                  |
+-------------------+----------------------------------------------------------------------------+
| Q_EMPTY           | Queue is empty indication.                                                 |
+-------------------+----------------------------------------------------------------------------+

Semaphore Management
^^^^^^^^^^^^^^^^^^^^

+-------------------+----------------------------------------------------------------------------+
| LITERAL           | MEANING                                                                    |
+===================+============================================================================+
| MAX_SEM           | For semaphore related operations, this defines the maximum number of       |
|                   | semaphores that may be referenced.  The range is 0 to 255.                 |
+-------------------+----------------------------------------------------------------------------+
| INV_SEMNAME       | Invalid semaphore name reference.                                          |
+-------------------+----------------------------------------------------------------------------+
| INV_SEMID         | Invalid semaphore identifier.                                              |
+-------------------+----------------------------------------------------------------------------+

Event Management
^^^^^^^^^^^^^^^^

+-------------------+----------------------------------------------------------------------------+
| LITERAL           | MEANING                                                                    |
+===================+============================================================================+
| EVT_OR            | Conditionally wait for any event specified.                                |
+-------------------+----------------------------------------------------------------------------+
| EVT_AND           | Conditionally wait for all events specified.                               |
+-------------------+----------------------------------------------------------------------------+
| EVT_0             | Specific event identifiers; event number 0 to event number 15, the         |
| EVT_1             | maximum number of events.  These symbolic names may be                     |
| EVT_2             | replaced by names meaningful to the application.                           |
| EVT_3             |                                                                            |
| EVT_4             |                                                                            |
| EVT_5             |                                                                            |
| EVT_6             |                                                                            |
| EVT_7             |                                                                            |
| EVT_8             |                                                                            |
| EVT_9             |                                                                            |
| EVT_10            |                                                                            |
| EVT_11            |                                                                            |
| EVT_12            |                                                                            |
| EVT_13            |                                                                            |
| EVT_14            |                                                                            |
| EVT_15            |                                                                            |
+-------------------+----------------------------------------------------------------------------+

Timer Management
^^^^^^^^^^^^^^^^

+-------------------+----------------------------------------------------------------------------+
| LITERAL           | MEANING                                                                    | 
+===================+============================================================================+
| MS_100            | (*) Time intervals:  100 milliseconds.                                     |
+-------------------+----------------------------------------------------------------------------+
| MS_500            | 500 milliseconds.                                                          |
+-------------------+----------------------------------------------------------------------------+
| SEC_1             | One second.                                                                |
+-------------------+----------------------------------------------------------------------------+
| SEC_10            | Ten seconds.                                                               |
+-------------------+----------------------------------------------------------------------------+
| MIN_1             | One minute.                                                                |
+-------------------+----------------------------------------------------------------------------+
| MIN_10            | Ten minutes.                                                               |
+-------------------+----------------------------------------------------------------------------+
| HR_1              | One hour.                                                                  |
+-------------------+----------------------------------------------------------------------------+

.. note::

    The timer management literal values are for a real-time clock interrupt rate of 32.77 milliseconds.  If this rate is changed, these literals must also be changed.  Other time reference symbols may be added, within the limitation of the maximum interval that can be represented in 16 bits.

Fault Management
^^^^^^^^^^^^^^^^

When used as the error location parameter with the LOG_FATAL and LOG_WARN primitives these literals in the bit significance shown below.

+-------+------------------------+
| bit   | meaning                |
+=======+========================+
| 15    | reserved               |
+-------+------------------------+
| 14‑12 | layer identifier       |
+-------+------------------------+
| 11‑10 | subsystem identifier   |
+-------+------------------------+
| 9‑8   | level identifier       |
+-------+------------------------+
| 7‑4   | procedure identifier   |
+-------+------------------------+
| 3‑0   | fault identifier       |
+-------+------------------------+

Here is an example of how the literals may be used with the fault management primitives.  The example shows the announcement of the third fault in layer 3, subsystem 0, level 2, and procedure 2.  

.. code:: c

    LOG_FATAL (LY_3 + SS_0 +LV_L + P2 + 3, NOT_USED)

The error location is logged in the Fault Analysis Area.
