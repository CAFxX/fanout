package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint; F2 uint }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0() uint32 {
    switch "" {
    case "hello":
        if true {
            var v5 uint8 = 127
            v6 := GenS1{F0: 255, F1: true, F2: 65535}
            _ = v5
            _ = v6
        }
    default:
        switch 2147483647 {
        case 0:
            v7 := GenS2{F0: -32768, F1: 65535, F2: 565}
            _ = v7
        case -515:
            var v8 string = "\n\t"
            _ = v8
        }
    }
    var v9 int64 = 42
    v9 -= -9223372036854775808
    v10 := GenS2{F0: 42, F1: 419, F2: 1}
    _ = v10
    _ = v9
    return 2
}
func genFn1(p0 uint16, p1 int64, p2 int) int64 {
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 float64 = 1.5
        p2 = p2
        v12 = (float64(1.5) * v12)
        _ = fi11
        _ = v12
    }
    fmt.Println(p2, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint8 = 255
    var v2 int8 = 1
    var v3 uint8 = (v1 << ((18446744073709551615) & 63))
    var v4 uint = 2
    v2 -= (v2 + (v2 - (-(v2))))
    var v13 bool = false
    if true {
        var v14 uint64 = 1
        var v15 uint = uint(v1)
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: v1, F1: v13, F2: 255}
    v3 = (v3 & v3)
    v17 := []int{0}
    v18 := GenS2{F0: -1, F1: v4, F2: (^(v4))}
    v17 = append(v17, 127)
    v19 := GenS1{F0: v3, F1: ((v4 + (v4 | (^(v4)))) == 0), F2: 4294967295}
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 17; fi21++ {
            v2 = v2
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            _ = fi21
        }
        v17 = append(v17, (fi20 + fi20))
        _ = fi20
    }
    v2 = (-(v2))
    fmt.Println(v2)
    for fi22 := 0; fi22 < 6; fi22++ {
        if len(v17) > 0 {
            v17[0] = fi22
        }
        var v23 uint = uint(v1)
        var v24 float64 = ((3.14159 - (0.1 / 0.1)) - 0.1)
        _ = fi22
        _ = v23
        _ = v24
    }
    v3 += 255
    fmt.Println(v4, v13, v3)
    for fi25 := 0; fi25 < 13; fi25++ {
        switch ("\n\t" + ("foo bar" + "a")) {
        case "":
            var v26 string = "\n\t"
            _ = v26
        }
        switch 128 {
        case -9223372036854775808:
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        case 32767:
            v27 := GenS3{F0: 1, F1: -32768}
            _ = v27
        }
        if (135 >= 4294967295) {
            var v28 int = 103
            _ = v28
        }
        _ = fi25
    }
    v29 := GenS3{F0: 451, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
