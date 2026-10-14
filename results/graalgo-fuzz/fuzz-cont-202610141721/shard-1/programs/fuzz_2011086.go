package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 bool }
func genFn0(p0 uint16, p1 int64, p2 int) uint16 {
    var v5 uint32 = 0
    if false {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 bool = (!false)
            _ = fi6
            _ = v7
        }
    } else {
        v8 := []int{}
        var v9 string = ("hello" + "hello")
        if false {
            v9 = "foo bar"
        }
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        if true {
            if true {
                for fi11 := 0; fi11 < 2; fi11++ {
                    v5 -= v5
                    _ = fi11
                }
                var v12 int = p2
                _ = v12
            }
            v5 += (v5 | v5)
            fmt.Println(v5, fi10)
        }
        _ = fi10
    }
    v13 := []int{p2, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    return (p0 - p0)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 42
    v3 := []string{"hello"}
    var v4 int8 = -1
    delete(v1, 2147483648)
    v14 := GenS1{F0: 1, F1: 18446744073709551615, F2: 42}
    var v15 uint = v2
    if (((42 > genFn0(42, 0, -147)) || true) && ((false && true) || false)) {
        v2 = v15
    }
    var v16 float64 = (1e10 - 1.5)
    v4 -= (v4 ^ v4)
    v1[877] = 2147483648
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint = (v15 * v2)
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 int64 = -931
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    var v21 int16 = -1
    fmt.Println(v15, v16, v4)
    switch "" {
    case "a":
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS2{F0: 979, F1: -2147483648, F2: false}
            _ = fi22
            _ = v23
        }
    case "foo bar":
        var v24 uint16 = 722
        _ = v24
    default:
        var v25 int = -787
        _ = v25
    }
    v4 += 100
    var v26 float64 = 0.0
    switch 32768 {
    case 121:
        v27 := GenS2{F0: 509, F1: -2147483648, F2: false}
        _ = v27
    case 127:
        for fi28 := 0; fi28 < 20; fi28++ {
            v16 = -2.25
            _ = fi28
        }
    }
    fmt.Println(v2, v15, v15)
    fmt.Println(v26, v4)
    var v29 int64 = -9223372036854775808
    var v30 int = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
