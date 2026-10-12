package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 uint16; F2 int8 }
func genFn0(p0 uint16, p1 int8) uint64 {
    if (true || ("" >= ("" + "\n\t"))) {
        v6 := GenS1{F0: "hello"}
        v7 := []int{9223372036854775807}
        v8 := GenS2{F0: -1, F1: p0}
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint16 = p0
    switch "foo bar" {
    case "hello":
        var v10 bool = (42 > 0)
        _ = v10
    }
    v9 = p0
    v11 := GenS2{F0: -462, F1: (p0 ^ 1)}
    _ = p0
    _ = p1
    _ = v11
    _ = v9
    return 33
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 1
    var v3 uint64 = uint64(v2)
    var v4 string = "foo bar"
    v5 := make(map[int]bool)
    v12 := []int{127, 0, 127, -910}
    var v13 int8 = -16
    for fi14 := 0; fi14 < 14; fi14++ {
        v13 = (v13 << ((317) & 63))
        v15 := GenS3{F0: (fi14 | fi14)}
        var v16 string = ("a" + ((v4 + v4) + (v4 + ("a" + (("\n\t" + (v4 + (((("" + ((v4 + (v4 + v4)) + "hello")) + (("a" + (v4 + (v4 + v4))) + v4)) + "foo bar") + "hello"))) + "")))))
        _ = fi14
        _ = v15
        _ = v16
    }
    switch (("a" + "\n\t") + "") {
    case "\n\t":
        var v17 uint = 18446744073709551615
        _ = v17
    default:
        var v18 uint = 42
        _ = v18
    }
    if false {
        var v19 float32 = ((-2.25 / (1e10 / 0.1)) * -2.25)
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int8 = v13
            v22 := GenS1{F0: ""}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v19
    } else {
        if ((104 != 128) && true) {
            if len(v12) > 0 {
                v12[0] = 347
            }
            var v23 int32 = 42
            switch int(v3) {
            case -694:
                v24 := GenS2{F0: 643, F1: 0}
                _ = v24
            }
            _ = v23
        }
    }
    if false {
        v2 += 632
    } else {
        if (176 == 168) {
            v25 := GenS1{F0: "hello"}
            var v26 int8 = v13
            fmt.Println(v13, v2)
            _ = v25
            _ = v26
        } else {
            for fi27 := 0; fi27 < 20; fi27++ {
                v4 = v4
                _ = fi27
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 int8 = (v13 / ((127) | 1))
                v2 = v2
                var v30 float32 = (3.14159 / 1.5)
                _ = fi28
                _ = v29
                _ = v30
            }
        }
    }
    v31 := GenS4{F0: v2, F1: 114, F2: (^(v13))}
    for fi32 := 0; fi32 < 13; fi32++ {
        v33 := v5[1]
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
