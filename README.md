# Data Structures Overview

This document provides an overview of several fundamental data structures included in this IDE project. Each data structure is briefly described, highlighting its primary use case and a real-world example of where it might be applied.

## Linked List

- **TL;DR**: A sequence of nodes where each node points to the next, allowing efficient insertion and deletion.
- **Real-world example**: A music playlist, where each song points to the next, allowing you to easily add or remove songs.


## Singly Linked List

- **TL;DR**: A linear collection of elements, where each element (node) points to the next, allowing for efficient element insertion and deletion.
- **Real-world example**: A conga line, where each person (node) holds onto the person in front of them, and new people can join or leave the line easily.

## Doubly Linked List

- **TL;DR**: Similar to a singly linked list but each node has two links: one pointing to the next node and one to the previous. This allows for traversal in both directions.
- **Real-world example**: A playlist that can be played forwards or backwards, where each song has a reference to both the next and the previous song.

## Node

- **TL;DR**: The fundamental unit used to build various data structures like linked lists, trees, and graphs. It typically stores data and references to other nodes.
- **Real-world example**: A family tree chart on a website, where each family member's detail page (node) has links to their relatives' pages.

## Binary Search Tree Node

- **TL;DR**: A node used in binary search trees, where each node has up to two children. It is defined in such a way that the left child's value is less than the parent node's value, and the right child's value is greater.
- **Real-world example**: Organizing a library's bookshelf, where each book (node) is placed based on its ISBN number, allowing for quick searching.

## Binary Tree

- **TL;DR**: A hierarchical structure where each node has up to two children, useful for organized storage and efficient searches.
- **Real-world example**: A family tree, showing parent-child relationships down the generations.

## Graph

- **TL;DR**: Nodes connected by edges, representing complex relationships; can be directed or undirected.
- **Real-world example**: A road map, where locations are nodes and roads are edges connecting the locations.

## Queue

- **TL;DR**: A FIFO (First In, First Out) structure, ideal for managing tasks in sequential order.
- **Real-world example**: A printer queue, where print jobs are completed in the order they're received.

## Stack

- **TL;DR**: A LIFO (Last In, First Out) structure, suitable for tasks where the last added item needs to be accessed first.
- **Real-world example**: Undo functionality in text editors, where the most recent changes are reversed first.

## Heap

- **TL;DR**: A complete binary tree where each node is smaller (or larger) than its children, facilitating efficient priority-based operations.
- **Real-world example**: A task scheduler, prioritizing tasks based on urgency.

## Priority Queue

- **TL;DR**: Similar to a queue but with elements processed according to their priority rather than the order they were added.
- **Real-world example**: An emergency room triage system, where patients' treatment priority is based on the severity of their condition.

## Map

- **TL;DR**: A collection of key-value pairs, allowing fast data retrieval based on unique keys.
- **Real-world example**: A dictionary, where each word (key) is associated with its definition (value).

## Hash (Hash Table/Hash Map)

- **TL;DR**: A data structure that stores key-value pairs. It uses a hash function to compute an index into an array of slots, from which the desired value can be found.
- **Real-world example**: A locker system where each locker (slot in the array) can be quickly accessed using a unique key, allowing for efficient storage and retrieval of items.

## Set

- **TL;DR**: An unordered collection of unique elements, useful for ensuring no duplicates.
- **Real-world example**: A guest list for an event, ensuring no duplicate invitations.

## Pair

- **TL;DR**: Two related data elements grouped together.
- **Real-world example**: Latitude and longitude coordinates, representing a point on the earth.

Each of these data structures plays a crucial role in software development, offering various ways to organize, store, and manage data efficiently. Understanding these structures and their applications is fundamental to designing effective algorithms and systems.
