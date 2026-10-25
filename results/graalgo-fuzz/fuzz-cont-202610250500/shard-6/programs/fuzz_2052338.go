package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int16; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int64 }
func genFn0(p0 int64, p1 int8, p2 int8) int32 {
    var v5 uint16 = 58
    p1 = (p1 * 26)
    var v6 uint64 = 250
    v5 -= 510
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 65535}
        switch ("hello" + ("foo bar" + "hello")) {
        case "a":
            var v9 int = (fi7 & (fi7 << ((656) & 63)))
            _ = v9
        default:
            p0 = (p0 | (p0 + (p0 % ((p0) | 1))))
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}
func genFn1() int {
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            if (255 != 18446744073709551615) {
                var v12 uint64 = 878
                fmt.Println(fi11, fi11, fi10)
                v13 := []int{fi10, fi10, 32767, (fi10 ^ -1)}
                _ = v12
                _ = v13
            }
            _ = fi11
        }
        if false {
            switch 1 {
            case 32767:
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 bool = true
                    _ = fi14
                    _ = v15
                }
            case -1:
                var v16 float32 = 3.14159
                _ = v16
            default:
                v17 := GenS1{F0: 4294967295}
                _ = v17
            }
            v18 := GenS1{F0: 2}
            switch ("" + "hello") {
            case "foo bar":
                var v19 int32 = 1
                _ = v19
            default:
                var v20 uint = 429
                _ = v20
            }
            _ = v18
        } else {
            v21 := GenS1{F0: 294}
            _ = v21
        }
        _ = fi10
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 string = "foo bar"
        v24 := GenS1{F0: 348}
        _ = fi22
        _ = v23
        _ = v24
    }
    return 32767
}

func main() {
    v1 := []int{543}
    var v2 int64 = 42
    var v3 string = ("a" + "hello")
    var v4 uint = 42
    fmt.Println(v3, v2)
    switch 9223372036854775807 {
    case -530:
        switch 2147483647 {
        case 524:
            v4 += v4
        case 2147483647:
            v25 := GenS2{F0: v3, F1: 1, F2: 1}
            _ = v25
        case 0:
            var v26 uint32 = 500
            _ = v26
        }
    case -9223372036854775808:
        var v27 int = genFn1()
        _ = v27
    default:
        v28 := GenS3{F0: (v3 + v3)}
        _ = v28
    }
    v29 := GenS1{F0: 1}
    v30 := GenS4{F0: v3, F1: v2}
    v31 := GenS1{F0: 170}
    v32 := GenS2{F0: v3, F1: 0, F2: -389}
    var v33 float64 = (-2.25 * 1.5)
    if len(v1) > 0 {
        v1[0] = -1
    }
    v34 := GenS1{F0: 2}
    for fi35 := 0; fi35 < 6; fi35++ {
        var v36 float32 = 3.14159
        if true {
            v36 += (v36 * 3.14159)
        }
        var v37 float64 = ((float64(float32(((0.0 - (-2.25 * v33)) - 3.14159))) / v33) / -2.25)
        _ = fi35
        _ = v36
        _ = v37
    }
    switch 4294967295 {
    case -956:
        var v38 string = "hello"
        _ = v38
    case 32767:
        v39 := GenS4{F0: "", F1: v2}
        _ = v39
    case 464:
        v40 := GenS3{F0: v3}
        _ = v40
    }
    v4 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
