Kernel Integration
=========================================================================

If the standard tool set is used, only the object files are needed, to link the kernel with the rest of the system.

If other tools are used, the source code files need to be edited to be compatible with the assembler used.  Generally, this only means changing assembler DIRECTIVE statements to those recognized by the assembler.  All include files, both those used by the kernel and those used to interface to the kernel, may also need to be edited.

The C-language include files do not need to be changed, because no vendor-dependent constructs or library functions are used.

After all kernel files have been edited and successfully assembled, the kernel object files may be linked with the rest of the system as if the original object files were used.

The files listed below are provided in assembly and C-language source code.  They must be declared in the order shown.

For C-language source files, the declarations are:

.. code:: c

    #include "if_l.h"
    #include "if_t.h"
    #include "if_k_e.h"
    #include "if_k_l.h"
    #include "if_k_m.h"
    #include "if_kf_l.h"
    #include "if_kf_t.h"

For assembly language, the equivalent declarations are:

.. code:: c

    $IF_L.INC
    $IF_K_E.INC
    $IF_K_L.INC
    $IF_K_M.INC
    $IF_KF_L.INC
    $IF_KF_T.INC

It is a requirement to set the correct include file pathname in either the assembler and compiler tools' command lines, or in the source code declarations.

Configuration File
------------------

The system configuration file provides the link between the kernel and application.  It is a program source file, which needs to be created then assembled.  It may be created using a text editor or the SYSBLD utility.  The example configuration file may serve as a template for a new file.

For a detailed description of the configuration file, refer to the section on building a system.

Interrupt Vector Table
----------------------

If the application uses interrupts, the default interrupt vector table, CONFIG_IV, needs to be changed to add the entry points for application interrupt service routines.  
Care should be taken that interrupt vectors used by the kernel are NOT modified.

The interrupts used by the kernel are the Real Time Interrupt (FFF0H), the software interrupt (FFF6H) and reset vector (FFFE).  The kernel also handles the COP Failure interrupt (FFFA), but this may be handled by the application.  All other interrupt vectors are handled as unexpected interrupts, and their occurrence is logged by the kernel. 

Language Considerations
-----------------------

The kernel implementation allows it to be portable across different development environments.  This is done by using generic language features and constructs defined by well-established standards.  Constructs and features particular to a specific vendor are avoided.

The assembly language source code contains directives for such things as symbols, macro definitions and memory location specifications.  For these directives, standard-defined constructs have been used.  However, their function is common across most assemblers and the directive may be replaced by the corresponding one for the assembler in use.

No compiler vendor-dependent externals, including library functions, are referenced.  All external references are resolved within the kernel, with links to the application through the configuration file.

Many compilers require special initialization code and library functions, depending on how the C code is written.  If the application invokes these, they must be specified in the linker command file.

Because the kernel manages the initialization sequence, any compiler-required initialization must be done in the software initialization procedure, which is accessed through the configuration file entry point, sw_init_ep.  For example, a procedure required by some compilers to initialize constant data at run-time provides this function.  Alternatively, initialize the data in an assembly language file, which eliminates the need for compiler-specific procedures and make the software portable to other compilers.

If bank switching is used to extend the addressing limitation, bank switching must be done through the kernel.  This supports the abstraction that the kernel manages system resources and memory is a resource.  Language support for bank switching cannot be used in the kernel environment.
