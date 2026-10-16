package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 uint64 }
type GenS4 struct { F0 string; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int32; F1 int32; F2 uint }
type GenS8 struct { F0 uint8 }
func genFn0() int {
    v5 := GenS1{F0: 733, F1: -128}
    var v6 int32 = 786
    var v7 int8 = 100
    v6 += -2147483648
    fmt.Println(v6)
    _ = v5
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint, p1 int64, p2 uint) uint16 {
    v8 := GenS2{F0: -128, F1: 1}
    var v9 float32 = (float32(0.0) / 1.5)
    var v10 int64 = p1
    v11 := GenS2{F0: 1, F1: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 uint32 = 926
    v2 := make(map[string]string)
    var v3 int = 157
    var v4 float32 = ((1e10 + (((0.0 / 1.0) + float32(((-2.25 + 0.1) - (float64((3.14159 + float32((1e10 / 1e10)))) - 3.14159)))) - float32(((0.1 / 3.14159) + (-2.25 + 1e10))))) / (((1e10 / -2.25) / ((0.1 + float32(0.0)) / -2.25)) - 1.5))
    v12 := GenS3{F0: int32(v1), F1: 42, F2: 42}
    if ((205 != 127) && true) {
        switch (-(v3)) {
        case 269:
            var v13 float64 = 1e10
            _ = v13
        default:
            v14 := GenS4{F0: ("" + "a"), F1: -1, F2: 185}
            _ = v14
        }
        v15 := GenS3{F0: 42, F1: 1, F2: 42}
        _ = v15
    }
    v16 := GenS5{F0: v1, F1: 245, F2: 4294967295}
    var v17 int32 = 2147483647
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(v1)
        var v19 uint8 = 127
        var v20 int64 = 9223372036854775807
        _ = fi18
        _ = v19
        _ = v20
    }
    if (2 == genFn1(4294967295, -9223372036854775808, 65535)) {
        var v21 bool = false
        _ = v21
    } else {
        var v22 int32 = 2147483647
        switch ("\n\t" + ("a" + "foo bar")) {
        case "\n\t":
            fmt.Println(v3, v1)
        case "hello":
            v22 = (v17 << ((255) & 63))
        case "foo bar":
            for fi23 := 0; fi23 < 5; fi23++ {
                v1 = v1
                var v24 int32 = v22
                delete(v2, ("a" + "hello"))
                _ = fi23
                _ = v24
            }
        default:
            v1 = (v1 ^ v1)
        }
        var v25 int64 = 0
        _ = v22
        _ = v25
    }
    fmt.Println(v17, v3)
    v26 := GenS5{F0: (v1 & v1), F1: 66, F2: 0}
    if true {
        fmt.Println(v3, v17)
        var v27 uint16 = 1
        _ = v27
    } else {
        v28 := GenS6{F0: -1, F1: 42}
        _ = v28
    }
    var v29 float64 = 1.5
    if (608 <= 205) {
        for fi30 := 0; fi30 < 10; fi30++ {
            v1 = 17
            for fi31 := 0; fi31 < 13; fi31++ {
                switch v3 {
                case 42:
                    var v32 int = genFn0()
                    _ = v32
                case -9223372036854775808:
                    var v33 int16 = -803
                    switch (v3 | 1) {
                    case 68:
                        v4 -= v4
                    }
                    v34 := []int{(fi30 * -734), (fi31 | -9223372036854775808)}
                    _ = v33
                    _ = v34
                case 1:
                    v35 := GenS5{F0: (v1 / ((v1) | 1)), F1: 96, F2: 4294967295}
                    _ = v35
                }
                _ = fi31
            }
            _ = fi30
        }
        var v36 bool = true
        _ = v36
    } else {
        v37 := GenS7{F0: (v17 << ((0) & 63)), F1: (-(v17)), F2: 72}
        switch v3 {
        case -582:
            v38 := GenS8{F0: 1}
            _ = v38
        case 2147483647:
            for fi39 := 0; fi39 < 10; fi39++ {
                v40 := v2["hello"]
                var v41 uint8 = 255
                _ = fi39
                _ = v40
                _ = v41
            }
        case 32768:
            _ = len(v2)
        default:
            var v42 int64 = 1
            _ = v42
        }
        v1 = 163
        _ = v37
    }
    v3 = (-(v3))
    v43 := GenS4{F0: (("a" + ("foo bar" + "hello")) + ""), F1: 2147483647, F2: 128}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v43
}
