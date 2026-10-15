package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int8; F2 uint64 }
type GenS4 struct { F0 string; F1 uint16; F2 uint8 }
func genFn0(p0 int32, p1 uint64, p2 int8) int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{int(p0), fi7, fi7}
        _ = fi7
        _ = v8
    }
    v9 := []int{-9223372036854775808, -963, 2147483647}
    var v10 bool = false
    switch "\n\t" {
    case "a":
        var v11 uint16 = 42
        _ = v11
    default:
        fmt.Println(p0, p1)
    }
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (p2 >> ((18446744073709551615) & 63))
}

func main() {
    var v1 float64 = float64((3.14159 / 1.0))
    var v2 int = 309
    v3 := make(map[uint32]uint32)
    var v4 uint = 740
    var v5 float64 = float64(3.14159)
    v6 := []string{"foo bar", "a"}
    v2 = (v2 + 42)
    v12 := GenS1{F0: 1, F1: 1}
    v2 += -1
    for fi13 := 0; fi13 < 15; fi13++ {
        v1 -= v5
        _ = fi13
    }
    v14 := GenS2{F0: -3}
    v15 := GenS1{F0: 32767, F1: 100}
    var v16 int = (v2 * v2)
    if false {
        var v17 uint16 = 585
        _ = v17
    }
    fmt.Println(v1, v5)
    switch ("foo bar" + ("foo bar" + ("" + "a"))) {
    case "foo bar":
        v18 := GenS2{F0: genFn0(415, 849, 127)}
        _ = v18
    default:
        delete(v3, 699)
    }
    v4 = 4294967295
    v19 := GenS3{F0: -9223372036854775808, F1: 100, F2: 4294967295}
    if ((false && (!true)) || true) {
        var v20 uint64 = 2
        switch "a" {
        case "hello":
            fmt.Println(v2)
        default:
            v21 := GenS4{F0: ((("\n\t" + "hello") + "hello") + ("hello" + "\n\t")), F1: 42, F2: 141}
            _ = v21
        }
        if true {
            var v22 float32 = -2.25
            var v23 int8 = -1
            _ = v22
            _ = v23
        }
        _ = v20
    }
    v2 = 467
    if true {
        v24 := GenS1{F0: 906, F1: -1}
        _ = v24
    } else {
        fmt.Println(v2, v1)
        var v25 uint32 = 605
        v4 = v4
        _ = v25
    }
    var v26 uint32 = 2147483648
    v27 := GenS3{F0: 0, F1: 100, F2: 18446744073709551615}
    v28 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
