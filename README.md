# taskboard
A simple dashboard which uses taskwarrior as backend.

(TODO: put screenshot here)

This is a simple terminal based dashboard I made to learn go, which uses termui library to visualize the tasks created with taskwarrior. Once my taskwarrior database started having a lot of tasks with different projects, it became difficult to get a good overview of everyting, without employing filters.

## Usage

Right now there are no ways to configure it elegantly. Just replace the filepaths and project names in the source code and give it a go (ha!) with `go run taskboard.go`

It currently shows 3 different project side by side, with the simple task IDs. You can keep using regular task warrior commands to create/delete tasks and it should update.

## License

See LICENSE file

