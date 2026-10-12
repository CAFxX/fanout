package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int32) uint16 {
    if (788 >= 4294967295) {
        p0 -= p0
        var v5 int32 = (p0 ^ (-(p0)))
        _ = v5
    } else {
        var v6 float32 = 1.5
        var v7 bool = true
        _ = v6
        _ = v7
    }
    p0 += (p0 ^ p0)
    _ = p0
    return 0
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]uint32)
    var v3 float64 = 3.14159
    var v4 int64 = -1
    switch "hello" {
    case "":
        var v8 bool = (-2.25 <= 0.1)
        _ = v8
    }
    var v9 uint16 = genFn0(2147483647)
    var v10 int8 = 59
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            v9 = v9
            v3 -= 1e10
            _ = fi12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        switch ("a" + "a") {
        case "\n\t":
            fmt.Println(v10, v10)
        case "foo bar":
            switch fi13 {
            case 574:
                v14 := []int{2147483647}
                _ = v14
            default:
                v1 += v1
            }
        default:
            v10 = (v10 % ((v10) | 1))
        }
        var v15 string = ("\n\t" + ("hello" + ""))
        if false {
            v16 := []int{fi13}
            _ = v16
        }
        _ = fi13
        _ = v15
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v9, v4)
    default:
        v17 := GenS1{F0: ("a" + "a")}
        _ = v17
    }
    v18 := GenS1{F0: "\n\t"}
    for fi19 := 0; fi19 < 8; fi19++ {
        switch 4294967295 {
        case 719:
            fmt.Println(v10, v9)
        case 9223372036854775807:
            v10 += 0
        case 2147483647:
            var v20 float64 = -2.25
            _ = v20
        }
        _ = fi19
    }
    var v21 int16 = -595
    if ((!true) || ((-(v10)) <= 127)) {
        v22 := GenS1{F0: "\n\t"}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 uint64 = 42
            v25 := GenS1{F0: "hello"}
            if (fi23 != fi23) {
                fmt.Println(v4, v10)
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 int16 = v21
        _ = v26
    }
    var v27 bool = false
    if v27 {
        for fi28 := 0; fi28 < 11; fi28++ {
            _ = len(v2)
            fmt.Println(v9, v10, fi28)
            v29 := GenS1{F0: "hello"}
            _ = fi28
            _ = v29
        }
    }
    v27 = true
    for fi30 := 0; fi30 < 5; fi30++ {
        v27 = false
        var v31 int32 = 324
        _ = fi30
        _ = v31
    }
    v32 := []int{501, 32768, 32768, -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v9
}
