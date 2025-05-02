        # GLA-Core Sandbox Execution

        This module provides secure and isolated environments to execute user-generated code in Go using Yaegi, WebAssembly, or Docker.

        ## Submodules

        - **yaegi/**: Uses Yaegi Go interpreter with timeout restrictions.
        - **wasm/**: Executes code safely inside a WASM runtime using wazero.
        - **docker/**: Runs code in isolated Docker containers with resource constraints.

        ## Interface

        All execution engines implement a common `Executor` interface for uniform usage across the system.
