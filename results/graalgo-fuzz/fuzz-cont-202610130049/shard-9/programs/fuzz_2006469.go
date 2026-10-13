package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint; F1 string; F2 uint32 }
type GenS6 struct { F0 int16 }
func genFn0(p0 string, p1 int8) int64 {
    p0 = (((((p0 + (p0 + p0)) + "a") + ((((p0 + p0) + (p0 + ((p0 + p0) + ""))) + ((p0 + p0) + p0)) + "a")) + "foo bar") + (p0 + (p0 + p0)))
    switch -393 {
    case 360:
        var v4 float32 = (((1e10 / (-2.25 / 0.1)) + 1e10) + float32((((0.1 / float64((((3.14159 / 3.14159) - 1.5) - 3.14159))) / 1e10) - 1e10)))
        _ = v4
    case -2147483648:
        var v5 float32 = (-2.25 + 1e10)
        _ = v5
    case 127:
        if true {
            var v6 int16 = 32767
            _ = v6
        }
    default:
        if false {
            var v7 float32 = 0.0
            if true {
                v8 := GenS1{F0: p0}
                _ = v8
            }
            v7 -= v7
            _ = v7
        }
    }
    switch p0 {
    case "foo bar":
        v9 := GenS2{F0: 255, F1: 127}
        _ = v9
    default:
        var v10 uint8 = 255
        _ = v10
    }
    var v11 int64 = -9223372036854775808
    var v12 int64 = v11
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 275
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 840
    var v3 int64 = v2
    switch 384 {
    case -1:
        v13 := GenS3{F0: (v3 | genFn0((("foo bar" + "a") + ""), int8(v3)))}
        _ = v13
    }
    if false {
        v14 := GenS4{F0: 0, F1: 216, F2: 4294967295}
        _ = v14
    } else {
        v2 = -9223372036854775808
    }
    _ = len(v1)
    v15 := GenS5{F0: 255, F1: "\n\t", F2: 0}
    switch 42 {
    case 493:
        if (-442 == 0) {
            v16 := []int{393, 32768, 9223372036854775807, -2147483648}
            for fi17 := 0; fi17 < 17; fi17++ {
                v3 -= v2
                _ = fi17
            }
            _ = v16
        } else {
            var v18 uint8 = 1
            switch "foo bar" {
            case "foo bar":
                for fi19 := 0; fi19 < 17; fi19++ {
                    for fi20 := 0; fi20 < 18; fi20++ {
                        var v21 int64 = (v3 ^ -1)
                        v22 := GenS1{F0: (("a" + ("hello" + "a")) + ("foo bar" + "\n\t"))}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v23 := []int{(fi19 / ((32767) | 1))}
                    _ = fi19
                    _ = v23
                }
            case "a":
                v24 := []int{127, -1, 386, -1}
                _ = v24
            }
            _ = v18
        }
    case -804:
        delete(v1, -2147483648)
    case -9223372036854775808:
        var v25 float32 = 0.1
        _ = v25
    default:
        var v26 int16 = 1000
        _ = v26
    }
    v3 = genFn0("foo bar", 0)
    for fi27 := 0; fi27 < 14; fi27++ {
        for fi28 := 0; fi28 < 1; fi28++ {
            if true {
                v29 := []int{(fi27 | fi27), (fi28 & (fi28 >> ((2) & 63))), fi28, fi27}
                _ = v29
            }
            v3 = genFn0("\n\t", 127)
            _ = fi28
        }
        var v30 uint = 597
        if (!true) {
            v30 = v30
            var v31 bool = true
            v30 = v30
            _ = v31
        } else {
            v32 := GenS2{F0: 42, F1: 0}
            v33 := []int{fi27, (fi27 | (fi27 << ((v30) & 63))), (fi27 ^ 32768)}
            _ = v32
            _ = v33
        }
        _ = fi27
        _ = v30
    }
    if true {
        v3 = (v2 / (((v3 - (v3 % (((v3 << ((518) & 63))) | 1)))) | 1))
        v34 := GenS6{F0: 42}
        _ = v34
    }
    v2 += (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
