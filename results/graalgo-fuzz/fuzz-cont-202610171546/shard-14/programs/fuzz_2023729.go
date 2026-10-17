package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int32) int32 {
    switch "foo bar" {
    case "\n\t":
        v5 := []int{-1, -1}
        _ = v5
    default:
        switch 14 {
        case 900:
            v6 := []int{2147483647, 916}
            _ = v6
        }
    }
    v7 := GenS1{F0: 2, F1: 42}
    if false {
        v8 := []int{-270, 1, 127}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 19; fi9++ {
            p0 = (p0 | 2147483647)
            _ = fi9
        }
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 bool) uint64 {
    if p0 {
        p0 = p0
    } else {
        fmt.Println(p0)
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(p0, p0)
        v11 := GenS1{F0: 726, F1: 32767}
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 255
    var v13 uint64 = 4294967295
    _ = p0
    _ = v12
    _ = v13
    return (v13 % ((v13) | 1))
}

func main() {
    var v1 string = "\n\t"
    v2 := []int{32767, -329, 2147483647}
    var v3 uint8 = 0
    var v4 int = -33
    var v14 int8 = 127
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    switch "" {
    case "hello":
        if true {
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 int = v4
                var v17 int = (v16 | 0)
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    case "":
        v4 = v4
    default:
        var v18 uint8 = v3
        _ = v18
    }
    var v19 uint8 = 0
    switch v1 {
    case "a":
        v20 := GenS1{F0: 2147483647, F1: 814}
        _ = v20
    case "":
        _ = len(v2)
    }
    var v21 string = v1
    var v22 uint8 = 1
    v4 -= v4
    var v23 uint = 2
    v1 = ((v21 + (v1 + (v21 + (v1 + ((((v1 + ((("hello" + (("" + v21) + ("foo bar" + "a"))) + (v1 + v21)) + v21)) + v1) + (("\n\t" + (v1 + v1)) + "hello")) + v1))))) + (("foo bar" + v21) + ""))
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(v19, v14, v4)
        fmt.Println(v3)
        if len(v2) > 0 {
            v2[0] = (fi24 ^ (fi24 * v4))
        }
        _ = fi24
    }
    var v25 int64 = 434
    switch ((v1 + (v21 + v21)) + (v21 + "hello")) {
    case "a":
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 string = v21
            var v28 uint16 = 615
            _ = len(v2)
            _ = fi26
            _ = v27
            _ = v28
        }
    case "foo bar":
        fmt.Println(v3, v22)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
