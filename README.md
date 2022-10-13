# GearTooth
A small programming language that transpiles to a python-based intermediary language 

### Overview

#### GearTooth Compiler - WIP
The GearTooth compiler takes in a .GT source file and transpires it into a python intermediary format.
This then gets run through a post-procesor to pack it into a .gtil file.

#### GTIL format - STABLE
GTIL stands for GearTooth intermediary language and is a compressed binary format used to store the python intermediary code generated by GearTooth Compiler.
It depends on the GearTooth runtime being installed on the host system or being linked in at compile time

#### GearTooth Runtime - WORKS FOR TESTING
The GearTooth runtime environment is a system to unpack gtil code and run it on a system.
It can be installed as a standalone component or linked in with the program before distribution, and there is an API flag for detecting the runtime mode
FYI: users can extract the gtil code from a linked runtime, so you might want to be aware of this

#### ChainLink - TODO
ChainLink is the solution used to link the GearTooth runtime in with the gtil code at compile time, and is easily extensible using a `link.cl` file.
This allows you to specify information like the version of GearTooth to pack and certain backend libraries and/or APIs to use.

### Credits
- [python-minifier](https://github.com/dflook/python-minifier) - @dflook
