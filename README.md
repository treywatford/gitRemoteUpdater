remoteUpdater
=============

Author: Trey Watford

Email: treyjustinwatford@gmail.com

This is a simple file system scanner written in GO that can be used to update git repository remotes.

To use this program please follow the steps listed below:

1. Run the program from the command line passing in the following values
  1. -root=your_root_directory
  2. -replace=text_to_replace
  3. -new_text=text_to_replace_with
  4. -force=true or false Note: true will not prompt before modification.
2. If you chose -force=false, when prompted press y [ENTER] to replace text or n [ENTER] to leave text unchanged.
3. After the program is complete press any key to exit.

Default values for command line flags:
  * -root=/ for unix -root=C:\ for windows
  * -replace=     #empty string by default
  * -new_text=    #empty string by default
  * -force=false

This program was created during an internship with ACS Technologies. For more information on careers at ACS Technologies please visit http://www.acstechnologies.com/company/careers.
