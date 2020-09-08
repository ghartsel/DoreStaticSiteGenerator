Real-time Systems and Autonomous Agents
=========================================================================

The high-level architecture consists of autonomous agents, which supports decentralized, dynamic, cooperating, and open systems and the cloud, which supports application computing and storage resources.

Autonomous agents define distributed software systems. In particular, agents are networked software programs
that perform specific tasks for a user and have intelligence that permits them to perform parts of their tasks
autonomously. Autonomous agents provide:

• Autonomy: The ability to perform the tasks without the direct involvement of humans and control their actions and internal state.
• Interaction: When appropriate, interact with other software agents and humans to solve problems.
• Responsiveness: Respond in a timely fashion to changes in their environment.
• Proactiveness: Exhibit goal-directed behavior and take the initiative as appropriate.

Cloud computing provides flexible and robust storage and computing resources. This enables dynamic data integration from multiple data sources. Additionally, the cloud offers flexibility and adaptability in the management and deployment of data analysis workflows. The dynamic deployment of software components as Cloud services removes the need for new client applications to be developed and deployed when the user requirements change. Services can be customized to support a distributed real-time system for the management and analysis of data streams generated agents.

Real-time systems commonly refer to application software that has strict timing objectives as a requirement.  Software must react to or trigger an event within a specific time limit to be considered to operate successfully.

Real-time applications include process control, telecommunications, avionics, manufacturing automation, and many others.  What these applications have in common is that they interact with the world external to the system to handle events, which may occur at unpredictable intervals.  It is specifically the asynchronous nature of these events that complicate real-time programming.

The multi-tasking kernel is a general solution for managing asynchronous events.  The kernel provides an environment to the application which, to a large extent, hides the asynchronous nature of the system.  This is achieved by allowing the application developer to think of handling events with the simple context of a task.  The kernel takes care of the details of sharing the processor between tasks and assures that each task has fair and timely access to processing resources.

Just as there are many types of operating systems, there are many types of real-time operating systems.  These range from general-purpose operating systems suitable to a broad range of applications, to highly specialized operating systems, which are usually tailored for hardware or application considerations.  In the choice of an operating system, the tradeoff is usually between feature-richness and efficiency of memory, time, and use.

The kernel is designed for a class of applications commonly referred to as embedded applications.  Within that context, it is feature-rich, because it provides all the facilities needed to support embedded system requirements.  At the same time, this narrow focus allows it to achieve efficiency through prior knowledge of embedded system behavior.

Among the characteristics that distinguish embedded real-time systems from other real-time systems, is that they must operate for extended intervals without intervention.  This places an extraordinary reliability requirement on this class of systems.  Some embedded systems also have critical performance constraints, in that the response time to an event must be guaranteed.  Such systems are commonly known as hard-deadline, embedded systems.

Kernel Features
---------------

The kernel features services, which are efficient and flexible enough to meet the particular needs of hard-deadline, embedded systems.

While the run-time considerations are important, system development and maintenance are equally important.  The complexity of real-time applications extends to all phases of the software life cycle.  The kernel, therefore, includes facilities, which allow developers to work at a higher level of abstraction.  This translates to systems completed sooner and with fewer errors.

The following list summarizes the kernel features:

    • Preemption        Multitasking scheduler for real-time response.
    • Priority Scheduler        Tasks may be prioritized by importance.
    • Multiple-instance     Single-source code instance for multiple tasks.
    • Task Bank Switching   Built-in facility supports bank switching.
    • Hard-deadline     Deterministic scheduler.
    • Fault-tolerant        Active fault detection and isolation.
    • Nested Interrupts     Concurrent interrupts use the common stack.
    • C-language Interface  High-level language support.
    • Rom-able          Position-independent runtime code.
    • Build Utility     Reliable resource and task specification.

The Kernel Model
----------------

The kernel model is a conceptual view of the software.  Specifically, the model is concerned with system resource management.

For the kernel, the resources of interest are processing time and the allocation of program and data space memory to processing time intervals.

The task is the central concept for processor time management.  By allocating task code memory space to a time segment, the kernel defines an entity that can be associated with time; and in real-time systems, time partitioning is an inherent requirement.  Further, by assigning multiple task code memory spaces to different time segments, the kernel achieves multi-tasking.

The work done by a task is independent of the kernel, although, certain rules govern the construction of tasks, to achieve time segmentation.

Effectively, a task is an abstraction of time, which provides a time-associated entity that can be managed and referenced.  In the kernel model, all functional references are in reference to tasks.  These operations that may be performed within the context of a task are called primitives.

In addition to providing a mechanism for operating with tasks, primitives also provide higher abstractions for time and memory.

For time, primitives allow developers to design in terms of clock time, synchronization, and atomic operations.  Collectively, these time-related primitives are called task coordination primitives.

For memory, primitives support abstractions such as queues, messages, and partitions.  Generally, the kernel model supports static, rather than dynamic, memory.  That is, all types of memory and their attributes are defined at design time.  There is no dynamic memory allocation facility at run-time, except in a very specialized case.  Such a memory model is better suited to hard-deadline applications, because system behavior becomes more predictable and simplified, with improved performance.

Autonomous Agents
-----------------

Autonomous agents are independently-running entities that can be added and removed from a system without affecting other components. Agents can also provide mechanisms for reconfiguration at run time without needing to be restarted.

Using agents for data collection and analysis provides the following features:

- Agents can acquire data by capturing packets from a network or any other suitable source. Thus, a system built from a collection of agents can cross the traditional boundaries between host-based and network-based resources.
- Because agents can be stopped and started without independently, agents can be upgraded as increased functionality is needed.
- If agents are implemented as separated processes on a host, each agent can be implemented in the programming language that is best suited for the task.

