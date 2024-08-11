package models

import "github.com/uadmin/uadmin"

// CommonResponsibilities is a map of common responsibilities that can be reused
var CommonResponsibilities = map[string]struct {
    MenuName        string
    MenuDisplayName string
    MenuIcon        string
}{
    // Level 1

    "View Dashboard": {
        "dashboard", 
        "Dashboard", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M10 19v-5h4v5c0 .55.45 1 1 1h3c.55 0 1-.45 1-1v-7h1.7c.46 0 .68-.57.33-.87L12.67 3.6c-.38-.34-.96-.34-1.34 0l-8.36 7.53c-.34.3-.13.87.33.87H5v7c0 .55.45 1 1 1h3c.55 0 1-.45 1-1"/></svg>`,
    },
    "View Timesheet": {
        "timesheet", 
        "Timesheet", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M19 6h-1V1H6v5H5c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2M8 3h8v3H8zm4 16c-2.76 0-5-2.24-5-5s2.24-5 5-5s5 2.24 5 5s-2.24 5-5 5"/><path fill="#333" d="M12.5 11.5h-1v2.71l1.64 1.64l.71-.71l-1.35-1.35z"/></svg>`,
    },
    "View Payroll": {
        "payroll", 
        "Payroll", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M17.755 14c.78 0 1.467.397 1.87 1H13.5a2.5 2.5 0 0 0-2.5 2.5v4q0 .243.045.472c-2.939-.186-5.136-1.25-6.53-3.207a2.75 2.75 0 0 1-.511-1.596v-.92A2.25 2.25 0 0 1 6.253 14zM12 2.005a5 5 0 1 1 0 10a5 5 0 0 1 0-10M12 17.5a1.5 1.5 0 0 1 1.5-1.5h8a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5h-8a1.5 1.5 0 0 1-1.5-1.5zm10 .5a1 1 0 0 1-1-1h-1a2 2 0 0 0 2 2zm0 2a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1zm-8-3a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2zm1 5a2 2 0 0 0-2-2v1a1 1 0 0 1 1 1zm4.25-2.5a1.75 1.75 0 1 0-3.5 0a1.75 1.75 0 0 0 3.5 0"/></svg>`,
    },
    "View and Request Leave": {
        "leaves", 
        "Leaves", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M21 11.11V7a2 2 0 0 0-2-2h-4V3a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v2H3a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h7.26A7 7 0 1 0 21 11.11M9 3h4v2H9m10 15a5 5 0 0 1-6 0a5 5 0 1 1 6 0m-4-7h1.5v2.82l2.44 1.41l-.75 1.3L15 16.69z"/></svg>`,
    },

    // Level 2

    "View Department": {
        "department", 
        "Department", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 16 16"><path fill="#333" d="M8 2.002a1.998 1.998 0 1 0 0 3.996a1.998 1.998 0 0 0 0-3.996M12.5 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m-9 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M5 7.993A1 1 0 0 1 6 7h4a1 1 0 0 1 1 1v3a3 3 0 0 1-.146.927A3.001 3.001 0 0 1 5 11zM4 8c0-.365.097-.706.268-1H2a1 1 0 0 0-1 1v2.5a2.5 2.5 0 0 0 3.436 2.319A4 4 0 0 1 4 10.999zm8 0v3c0 .655-.157 1.273-.436 1.819A2.5 2.5 0 0 0 15 10.5V8a1 1 0 0 0-1-1h-2.268c.17.294.268.635.268 1"/></svg>`,
    },

    // Level 3

    "View Employee": {
        "employees", 
        "Employees", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M10 2h4a2 2 0 0 1 2 2v2h4a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8c0-1.11.89-2 2-2h4V4c0-1.11.89-2 2-2m4 4V4h-4v2z"/></svg>`,
    },
    "Manage Employee": {
        "employees", 
        "Employees", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M10 2h4a2 2 0 0 1 2 2v2h4a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8c0-1.11.89-2 2-2h4V4c0-1.11.89-2 2-2m4 4V4h-4v2z"/></svg>`,
    },
    "Add Employee": {
        "employees", 
        "Employees", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M10 2h4a2 2 0 0 1 2 2v2h4a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8c0-1.11.89-2 2-2h4V4c0-1.11.89-2 2-2m4 4V4h-4v2z"/></svg>`,
    },
    "Manage Roles": {
        "roles", 
        "Roles", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 640 512"><path fill="#333" d="M224 0a128 128 0 1 1 0 256a128 128 0 1 1 0-256m-45.7 304h91.4c11.8 0 23.4 1.2 34.5 3.3c-2.1 18.5 7.4 35.6 21.8 44.8c-16.6 10.6-26.7 31.6-20 53.3c4 12.9 9.4 25.5 16.4 37.6s15.2 23.1 24.4 33c15.7 16.9 39.6 18.4 57.2 8.7v.9c0 9.2 2.7 18.5 7.9 26.3l-382.2.1C13.3 512 0 498.7 0 482.3C0 383.8 79.8 304 178.3 304M436 218.2c0-7 4.5-13.3 11.3-14.8c10.5-2.4 21.5-3.7 32.7-3.7s22.2 1.3 32.7 3.7c6.8 1.5 11.3 7.8 11.3 14.8v30.6c7.9 3.4 15.4 7.7 22.3 12.8l24.9-14.3c6.1-3.5 13.7-2.7 18.5 2.4c7.6 8.1 14.3 17.2 20.1 27.2s10.3 20.4 13.5 31c2.1 6.7-1.1 13.7-7.2 17.2l-25 14.4c.4 4 .7 8.1.7 12.3s-.2 8.2-.7 12.3l25 14.4c6.1 3.5 9.2 10.5 7.2 17.2c-3.3 10.6-7.8 21-13.5 31s-12.5 19.1-20.1 27.2c-4.8 5.1-12.5 5.9-18.5 2.4L546.3 442c-6.9 5.1-14.3 9.4-22.3 12.8v30.6c0 7-4.5 13.3-11.3 14.8c-10.5 2.4-21.5 3.7-32.7 3.7s-22.2-1.3-32.7-3.7c-6.8-1.5-11.3-7.8-11.3-14.8v-30.5c-8-3.4-15.6-7.7-22.5-12.9l-24.7 14.3c-6.1 3.5-13.7 2.7-18.5-2.4c-7.6-8.1-14.3-17.2-20.1-27.2s-10.3-20.4-13.5-31c-2.1-6.7 1.1-13.7 7.2-17.2l24.8-14.3c-.4-4.1-.7-8.2-.7-12.4s.2-8.3.7-12.4L343.8 325c-6.1-3.5-9.2-10.5-7.2-17.2c3.3-10.6 7.7-21 13.5-31s12.5-19.1 20.1-27.2c4.8-5.1 12.4-5.9 18.5-2.4l24.8 14.3c6.9-5.1 14.5-9.4 22.5-12.9v-30.5zm92.1 133.5a48.1 48.1 0 1 0-96.1 0a48.1 48.1 0 1 0 96.1 0"/></svg>`,
    },
    "Manage Salaries": {
        "salaries", 
        "Salaries", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 576 512"><path fill="#333" d="M64 64C28.7 64 0 92.7 0 128v256c0 35.3 28.7 64 64 64h448c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64zm208 128h224c8.8 0 16 7.2 16 16s-7.2 16-16 16H272c-8.8 0-16-7.2-16-16s7.2-16 16-16m-16 112c0-8.8 7.2-16 16-16h224c8.8 0 16 7.2 16 16s-7.2 16-16 16H272c-8.8 0-16-7.2-16-16m-92-152v13.9c7.5 1.2 14.6 2.9 21.1 4.7c10.7 2.8 17 13.8 14.2 24.5s-13.8 17-24.5 14.2c-11-2.9-21.6-5-31.2-5.2c-7.9-.1-16 1.8-21.5 5c-4.8 2.8-6.2 5.6-6.2 9.3c0 1.8.1 3.5 5.3 6.7c6.3 3.8 15.5 6.7 28.3 10.5l.7.2c11.2 3.4 25.6 7.7 37.1 15c12.9 8.1 24.3 21.3 24.6 41.6c.3 20.9-10.5 36.1-24.8 45c-7.2 4.5-15.2 7.3-23.2 9v13.8c0 11-9 20-20 20s-20-9-20-20v-14.6c-10.3-2.2-20-5.5-28.2-8.4c-2.1-.7-4.1-1.4-6.1-2.1c-10.5-3.5-16.1-14.8-12.6-25.3s14.8-16.1 25.3-12.6c2.5.8 4.9 1.7 7.2 2.4c13.6 4.6 24 8.1 35.1 8.5c8.6.3 16.5-1.6 21.4-4.7c4.1-2.5 6-5.5 5.9-10.5c0-2.9-.8-5-5.9-8.2c-6.3-4-15.4-6.9-28-10.7l-1.7-.5c-10.9-3.3-24.6-7.4-35.6-14C88 251.8 76.1 239 76 218.8c-.1-21.1 11.8-35.7 25.8-43.9c6.9-4.1 14.5-6.8 22.2-8.5v-14c0-11 9-20 20-20s20 9 20 20z"/></svg>`,
    },
    "Manage Deductions": {
        "salaries", 
        "Salaries", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 576 512"><path fill="#333" d="M64 64C28.7 64 0 92.7 0 128v256c0 35.3 28.7 64 64 64h448c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64zm208 128h224c8.8 0 16 7.2 16 16s-7.2 16-16 16H272c-8.8 0-16-7.2-16-16s7.2-16 16-16m-16 112c0-8.8 7.2-16 16-16h224c8.8 0 16 7.2 16 16s-7.2 16-16 16H272c-8.8 0-16-7.2-16-16m-92-152v13.9c7.5 1.2 14.6 2.9 21.1 4.7c10.7 2.8 17 13.8 14.2 24.5s-13.8 17-24.5 14.2c-11-2.9-21.6-5-31.2-5.2c-7.9-.1-16 1.8-21.5 5c-4.8 2.8-6.2 5.6-6.2 9.3c0 1.8.1 3.5 5.3 6.7c6.3 3.8 15.5 6.7 28.3 10.5l.7.2c11.2 3.4 25.6 7.7 37.1 15c12.9 8.1 24.3 21.3 24.6 41.6c.3 20.9-10.5 36.1-24.8 45c-7.2 4.5-15.2 7.3-23.2 9v13.8c0 11-9 20-20 20s-20-9-20-20v-14.6c-10.3-2.2-20-5.5-28.2-8.4c-2.1-.7-4.1-1.4-6.1-2.1c-10.5-3.5-16.1-14.8-12.6-25.3s14.8-16.1 25.3-12.6c2.5.8 4.9 1.7 7.2 2.4c13.6 4.6 24 8.1 35.1 8.5c8.6.3 16.5-1.6 21.4-4.7c4.1-2.5 6-5.5 5.9-10.5c0-2.9-.8-5-5.9-8.2c-6.3-4-15.4-6.9-28-10.7l-1.7-.5c-10.9-3.3-24.6-7.4-35.6-14C88 251.8 76.1 239 76 218.8c-.1-21.1 11.8-35.7 25.8-43.9c6.9-4.1 14.5-6.8 22.2-8.5v-14c0-11 9-20 20-20s20 9 20 20z"/></svg>`,
    },

    // Level 4

    "View Timesheet Logs": {
        "timesheetLogs", 
        "Timesheet Logs", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M19 6h-1V1H6v5H5c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2M8 3h8v3H8zm4 16c-2.76 0-5-2.24-5-5s2.24-5 5-5s5 2.24 5 5s-2.24 5-5 5"/><path fill="#333" d="M12.5 11.5h-1v2.71l1.64 1.64l.71-.71l-1.35-1.35z"/></svg>`,
    },
    "View Payroll Logs": {
        "payrollLogs", 
        "Payroll Logs", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M17.755 14c.78 0 1.467.397 1.87 1H13.5a2.5 2.5 0 0 0-2.5 2.5v4q0 .243.045.472c-2.939-.186-5.136-1.25-6.53-3.207a2.75 2.75 0 0 1-.511-1.596v-.92A2.25 2.25 0 0 1 6.253 14zM12 2.005a5 5 0 1 1 0 10a5 5 0 0 1 0-10M12 17.5a1.5 1.5 0 0 1 1.5-1.5h8a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5h-8a1.5 1.5 0 0 1-1.5-1.5zm10 .5a1 1 0 0 1-1-1h-1a2 2 0 0 0 2 2zm0 2a2 2 0 0 0-2 2h1a1 1 0 0 1 1-1zm-8-3a1 1 0 0 1-1 1v1a2 2 0 0 0 2-2zm1 5a2 2 0 0 0-2-2v1a1 1 0 0 1 1 1zm4.25-2.5a1.75 1.75 0 1 0-3.5 0a1.75 1.75 0 0 0 3.5 0"/></svg>`,
    },
    "View Leave Logs": {
        "leavesLogs", 
        "Leaves Logs", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M21 11.11V7a2 2 0 0 0-2-2h-4V3a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v2H3a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h7.26A7 7 0 1 0 21 11.11M9 3h4v2H9m10 15a5 5 0 0 1-6 0a5 5 0 1 1 6 0m-4-7h1.5v2.82l2.44 1.41l-.75 1.3L15 16.69z"/></svg>`,
    },
    "Manage Identification Format": {
        "identification", 
        "Identification", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="24" viewBox="0 0 384 512"><path fill="#333" d="M336 0H48C21.5 0 0 21.5 0 48v416c0 26.5 21.5 48 48 48h288c26.5 0 48-21.5 48-48V48c0-26.5-21.5-48-48-48M144 32h96c8.8 0 16 7.2 16 16s-7.2 16-16 16h-96c-8.8 0-16-7.2-16-16s7.2-16 16-16m48 128c35.3 0 64 28.7 64 64s-28.7 64-64 64s-64-28.7-64-64s28.7-64 64-64m112 236.8c0 10.6-10 19.2-22.4 19.2H102.4C90 416 80 407.4 80 396.8v-19.2c0-31.8 30.1-57.6 67.2-57.6h5c12.3 5.1 25.7 8 39.8 8s27.6-2.9 39.8-8h5c37.1 0 67.2 25.8 67.2 57.6z"/></svg>`,
    },
    "Manage Company Deductions": {
        "salaries", 
        "Salaries", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 576 512"><path fill="#333" d="M64 64C28.7 64 0 92.7 0 128v256c0 35.3 28.7 64 64 64h448c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64zm208 128h224c8.8 0 16 7.2 16 16s-7.2 16-16 16H272c-8.8 0-16-7.2-16-16s7.2-16 16-16m-16 112c0-8.8 7.2-16 16-16h224c8.8 0 16 7.2 16 16s-7.2 16-16 16H272c-8.8 0-16-7.2-16-16m-92-152v13.9c7.5 1.2 14.6 2.9 21.1 4.7c10.7 2.8 17 13.8 14.2 24.5s-13.8 17-24.5 14.2c-11-2.9-21.6-5-31.2-5.2c-7.9-.1-16 1.8-21.5 5c-4.8 2.8-6.2 5.6-6.2 9.3c0 1.8.1 3.5 5.3 6.7c6.3 3.8 15.5 6.7 28.3 10.5l.7.2c11.2 3.4 25.6 7.7 37.1 15c12.9 8.1 24.3 21.3 24.6 41.6c.3 20.9-10.5 36.1-24.8 45c-7.2 4.5-15.2 7.3-23.2 9v13.8c0 11-9 20-20 20s-20-9-20-20v-14.6c-10.3-2.2-20-5.5-28.2-8.4c-2.1-.7-4.1-1.4-6.1-2.1c-10.5-3.5-16.1-14.8-12.6-25.3s14.8-16.1 25.3-12.6c2.5.8 4.9 1.7 7.2 2.4c13.6 4.6 24 8.1 35.1 8.5c8.6.3 16.5-1.6 21.4-4.7c4.1-2.5 6-5.5 5.9-10.5c0-2.9-.8-5-5.9-8.2c-6.3-4-15.4-6.9-28-10.7l-1.7-.5c-10.9-3.3-24.6-7.4-35.6-14C88 251.8 76.1 239 76 218.8c-.1-21.1 11.8-35.7 25.8-43.9c6.9-4.1 14.5-6.8 22.2-8.5v-14c0-11 9-20 20-20s20 9 20 20z"/></svg>`,
    },

    "View Audit Logs": {
        "auditLogs",
        "Audit Logs", 
        `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 20 20"><g fill="#333"><path d="M10 6.5V2H5.5A1.5 1.5 0 0 0 4 3.5v13A1.5 1.5 0 0 0 5.5 18h4.757a5.5 5.5 0 0 1-1.235-3H8.5a.5.5 0 0 1 0-1h.522a5.5 5.5 0 0 1 .185-1H8.5a.5.5 0 0 1 0-1h1.1q.276-.538.657-1H8.5a.5.5 0 0 1 0-1h2.837c.895-.63 1.986-1 3.163-1a5.5 5.5 0 0 1 1.5.207V8h-4.5A1.5 1.5 0 0 1 10 6.5m-4 4a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0m0 2a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0m.5 1.5a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1"/><path d="M11 6.5V2.25L15.75 7H11.5a.5.5 0 0 1-.5-.5m8 8a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-4-3a.5.5 0 0 0-1 0v3a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 0-1H15z"/></g></svg>`,
    },

    

}

// Role model definition
type Role struct {
    uadmin.Model
    Name  string
    Level uint `uadmin:"min:1;max:4;default_value:1"`
}

// Override the Save method
func (r *Role) Save() {
    uadmin.Save(r)
    generateResponsibilitiesForRole(r)
    SetResponsibilitiesForLevel(r)
}

// Generate responsibilities for a specific role
func generateResponsibilitiesForRole(role *Role) {
    // Predefined list of general responsibilities with custom menu names and SVG icons
    responsibilities := map[string]struct{
        MenuName string
        MenuDisplayName string
        MenuIcon string
    }{
        "Approve Time Sheet": {
            "approvals",
            "Approvals",
            `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2m-8.29 13.29a.996.996 0 0 1-1.41 0L5.71 12.7a.996.996 0 1 1 1.41-1.41L10 14.17l6.88-6.88a.996.996 0 1 1 1.41 1.41z"/></svg>`,
        },
        "Approve Payroll": {
            "approvals",
            "Approvals",
            `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#333" d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2m-8.29 13.29a.996.996 0 0 1-1.41 0L5.71 12.7a.996.996 0 1 1 1.41-1.41L10 14.17l6.88-6.88a.996.996 0 1 1 1.41 1.41z"/></svg>`,
        },
        // Add other general responsibilities here
    }

    for responsibilityName, menuInfo := range responsibilities {
        specificResponsibilityName := responsibilityName + " for " + role.Name
        createResponsibility(specificResponsibilityName, menuInfo.MenuName, menuInfo.MenuDisplayName, menuInfo.MenuIcon)
    }
}

// Create a responsibility without associating it with a role
func createResponsibility(responsibilityName, menuName, menuDisplayName, menuIcon string) {
    var menu MenuName
    err := uadmin.Get(&menu, "name = ?", menuName)
    if err != nil || menu.ID == 0 {
        menu = MenuName{
            Name:        menuName,
            DisplayName: menuDisplayName,
            MenuIcon:    menuIcon,
        }
        uadmin.Save(&menu)
    } else {
        // Update the menu display name and icon if they're different
        if menu.DisplayName != menuDisplayName || menu.MenuIcon != menuIcon {
            menu.DisplayName = menuDisplayName
            menu.MenuIcon = menuIcon
            uadmin.Save(&menu)
        }
    }

    var responsibility Responsibility
    err = uadmin.Get(&responsibility, "name = ?", responsibilityName)
    if err != nil || responsibility.ID == 0 {
        responsibility = Responsibility{
            Name:        responsibilityName,
            DisplayName: menuDisplayName, // Set DisplayName here
            MenuNameID:  menu.ID,
        }
        uadmin.Save(&responsibility)
    } else {
        // Update DisplayName if it's different
        if responsibility.DisplayName != menuDisplayName {
            responsibility.DisplayName = menuDisplayName
            uadmin.Save(&responsibility)
        }
    }
}
// Set responsibilities based on role level
func SetResponsibilitiesForLevel(role *Role) {
    responsibilitiesByLevel := map[uint][]string{
        1: {"View Dashboard", "View Timesheet", "View Payroll", "View and Request Leave"},
        2: {"View Dashboard", "View Timesheet", "View Payroll", "View and Request Leave", "View Department"},
        3: {"View Dashboard", "View Timesheet", "View Payroll", "View and Request Leave", "View Department", "View Employee", "Manage Employee", "Add Employee", "Manage Roles", "Manage Salaries", "Manage Deductions"},
        4: {"View Dashboard", "View Timesheet Logs", "View Payroll Logs", "View Leave Logs", "View Employee", "Manage Employee", "Add Employee", "Manage Identification Format", "Manage Roles", "Manage Salaries", "Manage Deductions", "Manage Company Deductions", "View Audit Logs"},
    }

    responsibilities, exists := responsibilitiesByLevel[role.Level]
    if !exists {
        return // No responsibilities defined for this level
    }

    var roleResponsibility RoleResponsibility
    err := uadmin.Get(&roleResponsibility, "role_id = ?", role.ID)
    if err != nil || roleResponsibility.ID == 0 {
        roleResponsibility = RoleResponsibility{
            RoleID: role.ID,
        }
    }

    for _, responsibilityName := range responsibilities {
        // Fetch the responsibility data from the common map
        menuInfo, found := CommonResponsibilities[responsibilityName]
        if !found {
            menuInfo = struct {
                MenuName        string
                MenuDisplayName string
                MenuIcon        string
            }{
                 
                MenuName:        responsibilityName,  
                MenuDisplayName: responsibilityName, 
                MenuIcon:        "default-icon",      
            }
        }

        handleResponsibility(&roleResponsibility, responsibilityName, menuInfo.MenuName, menuInfo.MenuDisplayName, menuInfo.MenuIcon)
    }

    // Save the RoleResponsibility
    roleResponsibility.Save()
}

// handleResponsibility handles the creation and association of a responsibility
func handleResponsibility(roleResponsibility *RoleResponsibility, responsibilityName, menuName, menuDisplayName, menuIcon string) {
    var responsibility Responsibility
    err := uadmin.Get(&responsibility, "name = ?", responsibilityName)
    if err != nil || responsibility.ID == 0 {
        createResponsibility(responsibilityName, menuName, menuDisplayName, menuIcon)
        uadmin.Get(&responsibility, "name = ?", responsibilityName)
    }

    // Check if the responsibility is already associated with the role
    for _, r := range roleResponsibility.Responsibility {
        if r.ID == responsibility.ID {
            return 
        }
    }

    // If it's a new responsibility, append it
    roleResponsibility.Responsibility = append(roleResponsibility.Responsibility, responsibility)
}