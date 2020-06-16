<!---
#######################################
## Weather data REST API
##
## Format: markdown (md)
## Latest versions should be placed as first
##
## Notation: 00.01.02
##      - 00: stable released version
##      - 01: new features
##      - 02: bug fixes and small changes
##
## Updating schema (mandatory):
##      <empty_line>
##      <version> (dd/mm/rrrr)
##      ----------------------
##      * <item>
##      * <item>
##      <empty_line>
##
## Useful tutorial: https://en.support.wordpress.com/markdown-quick-reference/
##
#######################################
-->
0.2.0 (16.06.2020)
---------------------
    - Added cache mechanism
    - fixed OpenWeather API unmarshalling error
    
0.1.1 (16.06.2020)
---------------------
    - Added .travis.yml config 
    - Improved production Dockerfile
    
0.1.0 (15.06.2020)
---------------------
    - Implemented first version of REST API
    - added some tests
    - added Makefile
    - added dev and prod Dockerfiles
    
0.0.0 (11.06.2020)
---------------------
    - Initialised repository, added README and .gitignore
