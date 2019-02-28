name: Help
class: center, middle

# Exemples

.footnote[2017]

---

# Agenda

1. Introduction
2. Bullet
3. Array
4. Links
5. Images
6. Graphes

---

# Introduction

 Slide with code :

```yaml
dockerminislide:
  build: .
  ports:
    - "8000:80"
```


```yaml
dockerminislide:
  build: .
*  ports:
    - "8000:80"
```

---

# Bullet

* One
* Two

 .footnote[.red.bold[*] Orange]

---
# PlantUML

Sample with PlantUML:

@startuml
Bob->Alice : hello
@enduml

???
Slide Notes

---

# Array


| Header 1      |     2 header    |   header 3 |
| ------------- |: -------------: | ---------: |
| 1 Online      |        1        |      value |
| Line 2        |        2        |      value |
| 3 Online      |        3        |      value |

---

# Doughnut chart

.twocol-left[
### implicit colors

.chart-doughnut[

| One | Two | Three | Four | Five |
|-----|-----|-------|------|------|
| 1   | 2   | 3     | 4    | 5    |
]
]

.twocol-right[
### explicit colors

.chart-doughnut[

| .skyblue[One] | .plum[Two] | .aluminium[Three] | .chocolate[Four] | .butter[Five] |
|---------------|------------|-------------------|------------------|---------------|
| 1             | 2          | 3                 | 4                | 5             |
]
]

---
# Pie chart

.twocol-left[
### implicit colors

.chart-pie[

| One | Two | Three | Four | Five |
|-----|-----|-------|------|------|
| 1   | 2   | 3     | 4    | 5    |
]
]

.twocol-right[
### explicit colors

.chart-pie[

| .skyblue[One] | .plum[Two] | .aluminium[Three] | .chocolate[Four] | .butter[Five] |
|---------------|------------|-------------------|------------------|---------------|
| 1             | 2          | 3                 | 4                | 5             |
]
]

---
# PolarArea chart

.twocol-left[
### implicit colors

.chart-polararea[

| One | Two | Three | Four | Five |
|-----|-----|-------|------|------|
| 1   | 2   | 3     | 4    | 5    |
]
]

.twocol-right[
### explicit colors

.chart-polararea[

| .skyblue[One] | .plum[Two] | .aluminium[Three] | .chocolate[Four] | .butter[Five] |
|---------------|------------|-------------------|------------------|---------------|
| 1             | 2          | 3                 | 4                | 5             |
]
]

---
# Bar chart

.twocol-left[
### implicit colors

.chart-bar[

|   | One | Two | Three | Four | Five |
|---|-----|-----|-------|------|------|
| A | 1   | 2   | 3     | 4    | 5    |
| B | 2   | 3   | 4     | 5    | 6    |
| C | 3   | 4   | 5     | 6    | 7    |
| D | 4   | 5   | 6     | 7    | 8    |
]
]

.twocol-right[
### explicit colors

.chart-bar[

|               | One | Two | Three | Four | Five |
|---------------|-----|-----|-------|------|------|
| .skyblue[A]   | 1   | 2   | 3     | 4    | 5    |
| .plum[B]      | 2   | 3   | 4     | 5    | 6    |
| .aluminium[C] | 3   | 4   | 5     | 6    | 7    |
| .chocolate[D] | 4   | 5   | 6     | 7    | 8    |
]
]

---
# Line chart

.twocol-left[
### implicit colors

.chart-line[

|   | One | Two | Three | Four | Five |
|---|-----|-----|-------|------|------|
| A | 1   | 2   | 3     | 4    | 5    |
| B | 2   | 3   | 4     | 5    | 6    |
| C | 3   | 4   | 5     | 6    | 7    |
| D | 4   | 5   | 6     | 7    | 8    |
]
]

.twocol-right[
### explicit colors

.chart-line[

|               | One | Two | Three | Four | Five |
|---------------|-----|-----|-------|------|------|
| .skyblue[A]   | 1   | 2   | 3     | 4    | 5    |
| .plum[B]      | 2   | 3   | 4     | 5    | 6    |
| .aluminium[C] | 3   | 4   | 5     | 6    | 7    |
| .chocolate[D] | 4   | 5   | 6     | 7    | 8    |
]
]

---
# Radar chart

.twocol-left[
### implicit colors

.chart-radar[

|   | One | Two | Three | Four | Five |
|---|-----|-----|-------|------|------|
| A | 1   | 2   | 3     | 4    | 5    |
| B | 2   | 3   | 4     | 5    | 6    |
| C | 3   | 4   | 5     | 6    | 7    |
| D | 4   | 5   | 6     | 7    | 8    |
]
]

.twocol-right[
### explicit colors

.chart-radar[

|               | One | Two | Three | Four | Five |
|---------------|-----|-----|-------|------|------|
| .skyblue[A]   | 1   | 2   | 3     | 4    | 5    |
| .plum[B]      | 2   | 3   | 4     | 5    | 6    |
| .aluminium[C] | 3   | 4   | 5     | 6    | 7    |
| .chocolate[D] | 4   | 5   | 6     | 7    | 8    |
]
]

---

# Links

<http://www.orange.com>


[Orange](http://www.orange.com "link to google")

---

 # Images

 ![Google logo](https://www.google.fr/images/srpr/logo11w.png "google logo")

---

# Display slide step by step

1. Message 1   
Mon Message   
--

2. Message 2   
Mon Message  
--

3. Message 3   
Mon Message  

---

# Graphes with Mermaid
## Graph Simple

.mermaid[graph LR
        A-->B
        B-->C
        C-->A
        D-->C]

---

# Graphes avec Mermaid
## Diagramme de sequence



.mermaid[sequenceDiagram
    participant Alice
    participant Bob
    Alice->>John: Hello John, how are you?
    loop Healthcheck
        John->>John: Fight against hypochondria
    end
    Note right of John: Rational thoughts <br/>prevail...
    John-->>Alice: Great!
    John->>Bob: How about you?
    Bob-->>John: Jolly good! ]


---

# Graphes with Mermaid
## Gantt


.mermaid[gantt
    title A Gantt Diagram
    section Section
    A task           :a1, 2014-01-01, 30d
    Another task     :after a1  , 20d
    section Another
    Task in sec      :2014-01-12  , 12d
    anther task      : 24d ]


---
name: how

.left-column[
## How does it work?
### - Markdown
]
.right-column[
A Markdown-formatted chunk of text is transformed into individual slides by JavaScript running in the browser:
```remark
# Slide 1
This is slide 1

---

# Slide 2
This is slide 2
```

.slides[
  .first[
  One
  ]
  .second[
  Two   
  ]
 ]
 
 Regular Markdown rules apply with only a single exception:
 
]