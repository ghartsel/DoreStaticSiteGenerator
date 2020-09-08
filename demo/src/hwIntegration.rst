Hardware Interface
=========================================================================

Internal RAM
------------

The kernel uses the microcontroller internal RAM for its primary control data.  This allows the kernel to take advantage of the direct addressing mode for memory reference instructions.  The result is increased performance because most of the program execution occurs within the kernel.

After initialization, some of this area is made available to the application.  See the section on memory layout for a description of internal RAM.

Internal Registers
------------------

Only those internal registers listed below are used by the kernel.  All others are available for the application.

The kernel guarantees that the microcontroller's time-protected registers are initialized within the required time limit.

The 64-byte register block may be located on any 4K boundary.  Default register settings may be used or other values may be used for special application requirements.  The settings are defined in the configuration file.

HPRIO  (X03CH)

    default setting:  07H

    This sets the real-time interrupt as the highest priority interrupt source.  This is to assure the accuracy of the kernel timer services.

INIT   (X03DH)

    default setting:  01H

    This positions the 64-byte internal register block at address 1000H.

OPTION (X039H)

    default setting:  03H

    This sets the watchdog (COP) timeout rate to 1.049 seconds, for an 8MHz crystal.

    The watchdog is pulsed by the kernel whenever a task switch occurs, or periodically, if there is no task work pending.  This means that a task must not retain control of the processor for greater than the watchdog timeout period.

TMSK2  (X024H)

    default setting:  03H

    This sets the timer prescale factor to 16X for a real-time interrupt rate of 32.77 milliseconds.

Interrupt Vectors
-----------------

The kernel uses the following interrupt vectors during normal operation.

RTII  (XFFF0H)

    The real-time clock interrupt supports kernel timer services.

SWI   (XFFF6H)

    The software interrupt (TRAP) is used for the application interface to the kernel.

RESET (XFFFEH)

    The reset interrupt is used to vector execution to the start of the kernel initialization sequence.

    If the default interrupt vector table is used, CONFIG_IV, the following 
    interrupts cause a fatal fault to be reported, then vector to the start of the 
    kernel initialization sequence, as with a RESET interrupt.

OPCODE (XFFF8H) 

    Illegal opcode detected.

COP (XFFFAH)

    Watchdog timeout.

CME (XFFFCH)

    Clock monitor failure.

The default interrupt vector table causes all other interrupts to be reported as warning faults, then returns from the interrupt.

Operating Mode
----------------------

The kernel requires that the microcontroller be configured for Expanded Multiplexed Operation (MODE 1).  
Memory locations may be configured for any location allowed by the microcontroller.  An exception is internal RAM, which must always be located at address 0000H.

All microcontroller interfaces not described above are available to the application, with no restriction on their use by the kernel.

Any combination of RAM and ROM options is allowed.  The kernel code, itself, is ROMable.

Banking
-------

The kernel supports up to four memory banks for task code space.  These may be either RAM or ROM memory.

The kernel automatically switches in the bank where the task code resides by calling the user-provided bank switching routine.  The entry point for this routine is defined in the configuration file.