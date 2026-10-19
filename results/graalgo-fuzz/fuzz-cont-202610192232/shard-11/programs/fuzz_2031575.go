package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
func genFn0() uint32 {
    var v5 int64 = 1
    v5 = v5
    _ = v5
    return 65535
}
func genFn1(p0 int8, p1 int16, p2 int64) uint16 {
    var v6 int64 = p2
    var v7 int32 = 0
    v8 := []int{274, -2147483648, 4294967295, 1}
    for fi9 := 0; fi9 < 9; fi9++ {
        v8 = append(v8, fi9)
        _ = fi9
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]int)
    var v3 uint64 = 0
    v4 := []int{}
    for fi10 := 0; fi10 < 6; fi10++ {
        if (false || ((!true) || false)) {
            var v11 int64 = 319
            _ = v11
        }
        _ = fi10
    }
    _ = len(v2)
    if (!(-55 == 55)) {
        fmt.Println(v1, v1)
        v12 := GenS1{F0: genFn1(-103, 0, 0)}
        _ = v12
    } else {
        if (!(!false)) {
            var v13 uint8 = 45
            _ = v13
        }
        var v14 uint16 = genFn1(-1, 1, 1)
        _ = v14
    }
    switch ("a" + "\n\t") {
    case "hello":
        v1 -= 3.14159
    case "foo bar":
        v1 = 3.14159
    default:
        v3 = v3
    }
    if (!true) {
        v1 = (((float32(0.0) - (v1 + float32(1e10))) + 0.1) * 1.5)
    }
    switch -1 {
    case 42:
        v15 := GenS2{F0: -1}
        _ = v15
    case 127:
        var v16 int64 = 0
        _ = v16
    case 9223372036854775807:
        v17 := GenS2{F0: 0}
        _ = v17
    default:
        if false {
            fmt.Println(v3, v3)
            var v18 string = "foo bar"
            v18 = "hello"
            _ = v18
        } else {
            var v19 int8 = 0
            for fi20 := 0; fi20 < 16; fi20++ {
                fmt.Println(v1, v19)
                var v21 uint8 = 127
                var v22 float64 = -2.25
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v19
        }
    }
    var v23 uint64 = 396
    fmt.Println(v3)
    if (uint8(v3) == 0) {
        var v24 uint16 = 671
        _ = v24
    } else {
        v23 = 23
        _ = len(v2)
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        if (!(true && (true || false))) {
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 uint = 18446744073709551615
                _ = fi26
                _ = v27
            }
        }
        _ = fi25
    }
    fmt.Println(v23, v1, v23)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
