package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 int32; F1 int16 }
func genFn0(p0 int8, p1 uint8, p2 uint16) uint64 {
    switch ("foo bar" + "hello") {
    case "\n\t":
        v4 := []int{128}
        _ = v4
    case "a":
        v5 := GenS1{F0: 1000, F1: 0}
        _ = v5
    case "hello":
        if (42 != 0) {
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 int = (-(fi6))
                v8 := GenS2{F0: 2}
                v9 := GenS2{F0: (p1 << ((42) & 63))}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            var v10 int16 = -162
            _ = v10
        }
    }
    if (!true) {
        var v11 int = int(p0)
        var v12 int64 = -886
        v13 := GenS3{F0: "hello", F1: 903}
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 int32) int32 {
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := []int{(fi14 << ((0) & 63)), fi14}
        _ = fi14
        _ = v15
    }
    var v16 float32 = 1e10
    v17 := GenS4{F0: 2147483648, F1: 0, F2: 857}
    _ = p0
    _ = v16
    _ = v17
    return (p0 << ((18446744073709551615) & 63))
}

func main() {
    var v1 uint16 = 255
    var v2 int16 = -971
    var v3 int32 = -363
    v18 := []int{}
    if (((0 == 253) && (128 >= 282)) && false) {
        var v19 string = "a"
        v20 := GenS3{F0: v19, F1: 517}
        _ = v19
        _ = v20
    }
    var v21 float64 = -2.25
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    if (!false) {
        switch 9223372036854775807 {
        case 127:
            switch 56 {
            case 4294967295:
                if (100 == 127) {
                    var v22 int8 = 1
                    _ = v22
                } else {
                    v23 := GenS5{F0: 255, F1: -1}
                    v24 := GenS2{F0: 255}
                    _ = v23
                    _ = v24
                }
            case -70:
                v18 = append(v18, -233)
            }
        case 42:
            v25 := GenS3{F0: (("foo bar" + "") + "foo bar"), F1: 18446744073709551615}
            _ = v25
        default:
            var v26 bool = false
            _ = v26
        }
        v27 := GenS2{F0: uint8(v1)}
        _ = v27
    }
    fmt.Println(v3, v21)
    var v28 uint16 = (v1 % (((v1 >> ((0) & 63))) | 1))
    var v29 bool = true
    var v30 uint = uint(v1)
    if len(v18) > 0 {
        v18[0] = 1
    }
    var v31 uint32 = 0
    fmt.Println(v28, v21)
    switch ((("foo bar" + ((("" + "\n\t") + "a") + "hello")) + "hello") + ((("a" + "\n\t") + "hello") + "\n\t")) {
    case "hello":
        var v32 bool = (184 == -864)
        _ = v32
    }
    switch "\n\t" {
    case "hello":
        if ((-922 > 128) || (!v29)) {
            v29 = (-2.25 > 0.1)
        } else {
            var v33 int8 = 42
            fmt.Println(v3, v31)
            var v34 uint = 0
            _ = v33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        var v36 uint16 = uint16(v31)
        v30 += (v30 * (v30 ^ v30))
        if (false && v29) {
            v3 = (v3 / ((v3) | 1))
            v28 += (^(v36))
            var v37 float32 = float32(float64(0.0))
            _ = v37
        }
        _ = fi35
        _ = v36
    }
    fmt.Println(v31, v21, v1)
    v38 := GenS6{F0: (v3 / ((v3) | 1)), F1: 108}
    var v39 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
}
