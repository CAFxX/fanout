package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int32; F2 int16 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0(p0 int) string {
    var v6 int32 = -2147483648
    var v7 int32 = v6
    var v8 string = "foo bar"
    var v9 uint16 = 0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1(p0 uint8, p1 int8, p2 int16) uint8 {
    fmt.Println(p1)
    switch 0 {
    case -1:
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 float32 = 1.5
            var v12 uint = 44
            fmt.Println(p2)
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        v13 := GenS1{F0: (4294967295 == 267)}
        _ = v13
    }
    var v14 uint32 = 1
    var v15 int32 = 1
    switch 4294967295 {
    case -905:
        v15 = -2147483648
    case 2147483647:
        var v16 int8 = (p1 >> ((42) & 63))
        _ = v16
    case 0:
        v14 = v14
    default:
        v17 := GenS2{F0: 643, F1: v15, F2: (p2 / ((-32768) | 1))}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return p0
}

func main() {
    v1 := []string{}
    var v2 int8 = 1
    v3 := make(map[int]bool)
    var v4 int64 = 1
    var v5 int16 = 0
    var v18 int64 = (v4 % ((-483) | 1))
    var v19 string = "a"
    var v20 uint16 = 255
    switch -926 {
    case 2147483647:
        var v21 bool = true
        _ = v21
    default:
        v4 = 984
    }
    v22 := GenS3{F0: false, F1: 42}
    var v23 int8 = v2
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 int16 = 42
        _ = len(v1)
        if (4294967295 != 4294967295) {
            var v26 uint64 = 4294967295
            _ = v26
        } else {
            if ((!(!(42 <= 241))) && (!(!false))) {
                var v27 uint16 = v20
                _ = v27
            }
        }
        _ = fi24
        _ = v25
    }
    v23 = 100
    var v28 float64 = 0.1
    if true {
        v2 += -128
        var v29 int64 = v4
        _ = v29
    }
    v18 = v18
    var v30 int16 = 32767
    fmt.Println(v30)
    var v31 int16 = (v5 << ((731) & 63))
    v23 -= (^(v23))
    var v32 int16 = -685
    fmt.Println(v4)
    switch -9223372036854775808 {
    case -2147483648:
        switch 222 {
        case 0:
            var v33 int8 = (v2 - -1)
            _ = v33
        case 127:
            var v34 uint64 = 4294967295
            _ = v34
        default:
            var v35 int64 = (v4 * v4)
            _ = v35
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
