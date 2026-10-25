package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }
func genFn0(p0 int16) uint16 {
    var v6 int64 = 9223372036854775807
    var v7 int32 = 0
    v8 := GenS1{F0: 18446744073709551615, F1: 42}
    v6 = v6
    if (128 >= 2147483647) {
        var v9 uint = 65535
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 103
}
func genFn1() uint32 {
    var v10 int8 = 0
    var v11 int32 = 0
    _ = v10
    _ = v11
    return 2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -815
    v3 := []string{("hello" + "a"), "a", "a", "\n\t"}
    v4 := []string{"foo bar"}
    v5 := make(map[int]uint32)
    if ((true || true) && (202 <= 128)) {
        var v12 int = 522
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(fi13)
            var v14 int32 = (v2 << ((0) & 63))
            _ = fi13
            _ = v14
        }
        var v15 float64 = ((0.1 - ((1.5 / 1.5) / 3.14159)) - 0.1)
        _ = v12
        _ = v15
    }
    fmt.Println(v2)
    v2 = 42
    switch "hello" {
    case "":
        v16 := GenS2{F0: 618, F1: 32767, F2: 42}
        _ = v16
    default:
        v2 -= v2
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 float64 = (1e10 / ((0.1 * 3.14159) * (0.1 - float64(float32(3.14159)))))
        _ = fi17
        _ = v18
    }
    var v19 int64 = -1
    v19 = -1
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int64 = (-(v19))
        _ = fi20
        _ = v21
    }
    var v22 int = 128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5[880] = 942
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
