package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint64; F2 int64 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16 }
func genFn0() uint16 {
    v6 := GenS1{F0: 2}
    if false {
        var v7 uint8 = 2
        _ = v7
    } else {
        switch (("hello" + ("hello" + "hello")) + "a") {
        case "foo bar":
            switch 42 {
            case -9223372036854775808:
                v8 := []int{632}
                _ = v8
            case 365:
                v9 := GenS2{F0: -2147483648}
                _ = v9
            }
        case "":
            var v10 uint64 = 295
            _ = v10
        case "\n\t":
            var v11 int32 = -364
            _ = v11
        default:
            var v12 float32 = (-2.25 * ((0.1 + 3.14159) - (float32(0.1) * 1e10)))
            _ = v12
        }
        var v13 uint8 = 212
        switch 127 {
        case -2147483648:
            v13 = 42
        case 32768:
            switch 32767 {
            case -9223372036854775808:
                var v14 string = "foo bar"
                _ = v14
            case 316:
                fmt.Println(v13)
            case 2147483647:
                switch 591 {
                case 9223372036854775807:
                    if (false || true) {
                        var v15 int = 985
                        _ = v15
                    }
                case -175:
                    if (!true) {
                        v16 := GenS3{F0: 0, F1: 18446744073709551615, F2: 771}
                        v13 = v13
                        _ = v16
                    }
                }
            default:
                v13 += v13
            }
        }
        _ = v13
    }
    v17 := []int{727, 127, -2147483648, -588}
    v18 := GenS4{F0: -128, F1: -2147483648}
    switch ("hello" + ("" + "hello")) {
    case "\n\t":
        if ("\n\t" < "foo bar") {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            var v19 string = "hello"
            if len(v17) > 0 {
                v17[0] = -1
            }
            _ = v19
        } else {
            var v20 bool = (!(!(32768 != -802)))
            v20 = (2 >= 18446744073709551615)
            if (v20 || (v20 || (255 >= 0))) {
                var v21 float32 = (((3.14159 / -2.25) / -2.25) - 1.5)
                var v22 float64 = float64(v21)
                var v23 float64 = 1e10
                _ = v21
                _ = v22
                _ = v23
            }
            _ = v20
        }
    case "":
        if len(v17) > 0 {
            v17[0] = -1
        }
    }
    _ = v17
    _ = v18
    _ = v6
    return 32767
}
func genFn1(p0 uint, p1 int64, p2 int64) int64 {
    if (!(!false)) {
        var v24 uint8 = 42
        _ = v24
    } else {
        p2 -= p2
        fmt.Println(p0)
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 int = (fi25 - 1)
            p1 += (p1 | 1)
            p1 = (p1 % (((-(p2))) | 1))
            _ = fi26
            _ = v27
        }
        p2 = p2
        _ = fi25
    }
    p2 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 int8 = 100
    var v4 uint64 = 18446744073709551615
    var v5 uint64 = 483
    v28 := GenS3{F0: 32767, F1: v5, F2: 0}
    v29 := v2["a"]
    for fi30 := 0; fi30 < 3; fi30++ {
        switch 32768 {
        case 127:
            switch (fi30 % ((507) | 1)) {
            case 2147483647:
                var v31 int64 = genFn1(42, -1, 1)
                _ = v31
            }
        }
        fmt.Println(v4)
        _ = fi30
    }
    v3 -= v3
    v4 -= v5
    v32 := GenS4{F0: (-(v3)), F1: 32767}
    v33 := GenS5{F0: (v3 << ((4294967295) & 63))}
    for fi34 := 0; fi34 < 19; fi34++ {
        for fi35 := 0; fi35 < 16; fi35++ {
            v29 -= v29
            v3 += -1
            fmt.Println(v3)
            _ = fi35
        }
        if (true && (!true)) {
            v36 := GenS6{F0: -1}
            var v37 float64 = (((1.5 + ((float64(3.14159) - -2.25) * 0.1)) * 1e10) - float64(float32(0.0)))
            _ = v36
            _ = v37
        }
        _ = fi34
    }
    v4 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
