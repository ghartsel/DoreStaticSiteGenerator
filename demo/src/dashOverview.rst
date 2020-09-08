Dashboard Introduction
=========================================================================

SYSBLD is a program that runs in the DOS environment.  The user is prompted for system information, which is used to create an ASCII source file.  This file is then assembled to produce an object file, which is linked with the kernel and application software to become part of the target system.

Purpose of the Utility
----------------------

The system configuration utility has two purposes.  One is as a method of binding the kernel to the application at run-time, and the other is to provide fault detection early in the implementation phase.

The file created by SYSBLD describes the hardware relevant to the kernel, application resource needs, and the tasks and their attributes.  The kernel uses this information at initialization, and for some special conditions, to configure the hardware, kernel, and application, thus binding the kernel to the application.

SYSBLD is an automated mechanism for creating a source file.  Part of the automation process is to check data consistency and resource availability.

As configuration information is provided, SYSBLD correlates the data and tests against limits.  If SYSBLD is used to create the configuration file, the probability of a resource-related fault occurring at run-time is greatly reduced.

Assembling and Modifying the Output File
----------------------------------------

The following pages show the assembly language source file created by SYSBLD; standard directives and syntax are used.

The file may be created manually with a text editor or the output of SYSBLD may be edited, although this defeats the fault protection advantage of SYSBLD.

If the file is created manually, the only requirement is that PUBLIC data must be declared exactly as shown.  All other parts of the file may be changed to suit the development environment and application.

.. code::

        NAME CONFIG_STT
        P68H11
        RSEG SDATA

    $..\INC\IF_L.INC          REQUIRED INCLUDE FILES
    $..\INC\IF_K_L.INC
    $..\INC\TNAME.INC

    ****************************************
    * PUBLIC DECLARATIONS
    * 
        PUBLIC eram_loc
        PUBLIC apdata,usrstk,osdata,sysstk
        PUBLIC tsk_cfg
        PUBLIC hw_init_ep,sw_init_ep,erom_bk_ep,fault_handler
        PUBLIC r_que_cnt,r_sem_cnt,r_msg_cnt
        PUBLIC u68_tmsk2,u68_init,u68_hprio,u68_option

    ****************************************
    * EXTERNAL DECLARATIONS
    * 
        EXTERN epage_chg DEFINED TASKS AND
        EXTERN hw_init   OTHER ENTRY POINTS
        EXTERN io_init
        EXTERN task_0
        EXTERN task_1

    ****************************************
    * PHYSICAL MEMORY DECLARATIONS
    * 
    eram_loc:
        FDB $4000        DYNAMIC MEMORY LOCATION
        FDB $4000        DYNAMIC MEMORY SIZE

        FDB INV_ADDR     MEMORY DESCRIPTOR TERMINATOR
    ****************************************
    * LOGICAL MEMORY DECLARATIONS
    * 
    apdata:
        FDB $5000        APPLICATION MEMORY PARTITION
        FDB $1000        APPLICATION PARTITION SIZE
    usrstk:
        FDB $6000        TASK STACK PARTITION
        FDB $1000        TASK STACK PARTITION SIZE
    osdata:
        FDB $7000        KERNEL DATA PARTITION
        FDB $0f00        KERNEL DATA PARTITION SIZE
    sysstk:
        FDB $7f00        SUPERVISOR STACK LOCATION
        FDB $0100        SUPERVISOR STACK SIZE

    ****************************************
    * TASK SPECIFICATIONS

    tsk_cfg:
        FDB task_0       ENTRY POINT
        FCB BANK_0       RESIDENT BANK
        FCB TASK_0       NAME
        FDB STKSIZ_S     STACK SIZE OPTION
        FDB NOT_USED     MEMORY SIZE OPTION
        FCB PRI_HIGH     PRIORITY OPTION
        FCB NOT_USED

        FDB task_1
        FCB BANK_0
        FCB TASK_1
        FDB STKSIZ_S
        FDB NOT_USED
        FCB PRI_NORM
        FCB NOT_USED

        FDB INV_ADDR     TASK DESCRIPTOR TERMINATOR
    ****************************************
    * SPECIAL PURPOSE PROCS & TASKS

    erom_bk_ep:
        FDB epage_chg    BANK SWITCH PROCEDURE ENTRY POINT
    hw_init_ep:
        FDB hw_init      H/W INITIALIZATION PROC ENTRY PT
    sw_init_ep:
        FDB sw_init      S/W INITIALIZATION PROC ENTRY PT
    fault_handler:
        FCB TASK_0       FAULT HANDLER TASK NAME

    ****************************************
    * RESOURCE SPECIFICATIONS

    r_que_cnt:
        FDB 2            NUMBER OF QUEUES
    r_sem_cnt:
        FDB 2            NUMBER OF SEMAPHORES
    r_msg_cnt:
        FDB 4            NUMBER OF MESSAGES

    ****************************************
    * PROCESSOR STARTUP CONFIGURATION

    u68_tmsk2:
        FCB $03          INTERNAL REGISTERS:  TMSK2
    u68_init:
        FCB $01                               INIT
    u68_hprio:
        FCB $07                               HPRIO
    u68_option:
        FCB $03                               OPTION

        END

