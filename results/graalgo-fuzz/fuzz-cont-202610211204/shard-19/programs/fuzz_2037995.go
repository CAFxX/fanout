package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint16, p1 int64, p2 int8) uint64 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 float64 = float64(-2.25)
        _ = fi6
        _ = v7
    }
    p2 = 1
    var v8 int32 = 0
    v9 := GenS1{F0: 42, F1: 1}
    if true {
        v10 := []int{}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 bool, p1 int16) uint32 {
    for fi11 := 0; fi11 < 11; fi11++ {
        p1 = p1
        p1 = p1
        _ = fi11
    }
    var v12 bool = p0
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    v1 := []int{}
    v2 := []string{"a", "a", "foo bar"}
    var v3 int16 = 32767
    var v4 int8 = 0
    var v5 float64 = float64(((3.14159 + float32(0.0)) * 1e10))
    switch -727 {
    case 42:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 127:
        var v13 int64 = 9223372036854775807
        _ = v13
    case -287:
        v14 := GenS1{F0: 701, F1: 4294967295}
        _ = v14
    }
    v15 := GenS1{F0: 0, F1: 1}
    var v16 uint32 = genFn1((!(18446744073709551615 < 18446744073709551615)), (^(v3)))
    var v17 int = 128
    switch "a" {
    case "foo bar":
        v18 := GenS2{F0: 0, F1: 42}
        _ = v18
    case "a":
        v5 = -2.25
    case "\n\t":
        v5 = (-2.25 / (v5 - 1e10))
    }
    v19 := GenS2{F0: 0, F1: 2}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v16, v4, v17)
    v5 -= v5
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 int32 = 63
        if false {
            if (720 >= 1) {
                fmt.Println(v5, v16, v3)
                var v22 uint = 2
                var v23 int8 = 127
                _ = v22
                _ = v23
            } else {
                switch 9223372036854775807 {
                case -2147483648:
                    var v24 uint64 = 18446744073709551615
                    _ = v24
                }
                var v25 uint8 = 1
                _ = v25
            }
            v17 = fi20
        } else {
            if ((^(v4)) >= 97) {
                var v26 uint16 = 32767
                fmt.Println(v21, fi20, v3)
                _ = v26
            } else {
                if true {
                    var v27 int64 = 0
                    v16 = v16
                    _ = v27
                }
            }
        }
        _ = fi20
        _ = v21
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 uint16 = 65535
        _ = fi28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
