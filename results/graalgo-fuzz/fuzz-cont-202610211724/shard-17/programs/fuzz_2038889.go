package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint32 = 4294967295
    v2 := []string{}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    v1 = v1
    v1 = v1
    switch -2147483648 {
    case -6:
        v1 += v1
    case 0:
        var v5 uint = 378
        _ = v5
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    var v6 uint8 = 172
    var v7 uint16 = 32767
    v1 += 2
    fmt.Println(v1)
    switch ("hello" + "") {
    case "a":
        var v8 bool = false
        _ = v8
    }
    var v9 int16 = -395
    v10 := GenS1{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
