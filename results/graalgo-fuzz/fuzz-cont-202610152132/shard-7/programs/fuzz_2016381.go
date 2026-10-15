package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    v3 := make(map[uint32]int)
    if (true && (true || (false || (!true)))) {
        var v4 uint16 = 65535
        fmt.Println(v4)
        _ = v4
    } else {
        var v5 uint64 = 18446744073709551615
        fmt.Println(v5)
        _ = v5
    }
    if ("" != "hello") {
        var v6 string = "hello"
        if ((!(-360 == -238)) && false) {
            var v7 uint8 = 197
            _ = v7
        }
        _ = v6
    }
    switch "\n\t" {
    case "\n\t":
        var v8 int32 = 1
        _ = v8
    case "":
        var v9 uint32 = 574
        _ = v9
    case "hello":
        if true {
            v3[4294967295] = 9223372036854775807
            var v10 uint16 = 0
            var v11 int = 1
            _ = v10
            _ = v11
        }
    default:
        var v12 uint16 = 1
        _ = v12
    }
    var v13 int16 = 0
    switch "" {
    case "\n\t":
        var v14 uint64 = 4294967295
        _ = v14
    case "a":
        var v15 uint32 = 1
        _ = v15
    case "":
        switch ("" + "\n\t") {
        case "hello":
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        }
    default:
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := v3[42]
            v18 := v1[-1]
            var v19 int = v17
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
    }
    fmt.Println(v13)
    v13 += 42
    switch 0 {
    case 2147483647:
        switch "\n\t" {
        case "hello":
            var v20 int16 = 0
            _ = v20
        case "foo bar":
            var v21 bool = false
            _ = v21
        }
    case 32767:
        v13 = (v13 * v13)
    case -1:
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 int = 42
            var v24 uint = 65535
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    var v25 bool = true
    var v26 uint64 = 515
    v25 = (("\n\t" + ("\n\t" + "a")) < ("hello" + "hello"))
    v27 := GenS1{F0: 42}
    v28 := GenS1{F0: 0}
    v29 := GenS1{F0: 65}
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    var v30 bool = ((1 >= 65535) && v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
