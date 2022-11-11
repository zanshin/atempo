# ATEMPO
A visit tracking system for personal websites.

## GOALS
Track specific metrics about visits to the website.

- Visits
  - Daily totals
  - Weekly totals
  - Year-to-date total

- Page views
  - Daily total
  - Weekly total
  - Year-to-date total

- Page ranking
  - Most visited pages in descending order
  - Trend?

- Referrers
  - Most recent referrers
  - Repeat referrers

- Visitor information
  - Browser family
  - Browser version
  - Operating System
  - Device resolution

- Location
  - Visitor location

## TECHNOLOGY
Use Go language for the backend, and to construct the dashboard displaying all the metrics.
JavaScript will be used to collect data on the site/pages being tracked. A MySQL database will
persist all the data.

## ROADMAP
- [X] 0.1.0 - Initial project structure, configuration file, mechanism to read/process configuration items
- [ ] 0.2.0 - Database creation, table creation, initial CRUD operations
- [ ] 0.3.0 - JavaScript to capture data from visited page
- [ ] 0.4.0 - Initial web page(s) to display metrics from database
- [ ] 0.5.0 - Add database indexes and tune queries
- [ ] 0.6.0 -
