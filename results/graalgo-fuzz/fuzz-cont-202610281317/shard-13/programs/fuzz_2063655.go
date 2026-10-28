package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 uint; F1 int; F2 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint; F1 int64; F2 uint64 }
func genFn0(p0 int64, p1 uint64) string {
    p1 = p1
    v5 := []int{}
    if (-870 != 0) {
        var v6 int16 = -92
        var v7 int64 = 1
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: (p0 | (p0 / ((0) | 1))), F1: 42, F2: -1}
        if (false && ((!((((true || false) && (!true)) && false) || true)) && false)) {
            var v9 int8 = 100
            _ = v9
        } else {
            if true {
                var v10 uint64 = p1
                _ = v10
            }
            _ = len(v5)
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            p0 = 42
            if true {
                v5 = append(v5, (fi11 * (fi11 % ((4294967295) | 1))))
                fmt.Println(fi11)
                fmt.Println(p0)
            } else {
                var v12 int8 = 11
                p0 = (p0 | p0)
                switch (("hello" + "\n\t") + "hello") {
                case "a":
                    fmt.Println(v12, fi11)
                default:
                    fmt.Println(fi11, v12, p1)
                }
                _ = v12
            }
            _ = fi11
        }
        _ = v8
    }
    switch "" {
    case "a":
        v13 := GenS2{F0: 255, F1: -792, F2: 42}
        _ = v13
    default:
        var v14 int32 = 0
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v5
    return (("" + "\n\t") + ("a" + ("\n\t" + "hello")))
}

func main() {
    v1 := []int{127, 0, 9223372036854775807, -1}
    var v2 float32 = -2.25
    var v3 float32 = ((0.0 / v2) + v2)
    v4 := []string{"hello"}
    var v15 float64 = 0.1
    fmt.Println(v15, v15, v2)
    if false {
        v3 += (v2 / v3)
        if false {
            _ = len(v4)
            v16 := GenS2{F0: 255, F1: 4294967295, F2: -32768}
            _ = v16
        } else {
            var v17 uint = 73
            _ = v17
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            _ = len(v4)
            _ = fi18
        }
    }
    switch 32767 {
    case -954:
        if true {
            fmt.Println(v15)
            switch -833 {
            case -471:
                v4 = append(v4, "")
            default:
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS2{F0: 0, F1: (fi19 >> ((2) & 63)), F2: 1}
                    _ = fi19
                    _ = v20
                }
            }
            v15 = v15
        } else {
            switch -1 {
            case -9223372036854775808:
                var v21 int64 = 1
                _ = v21
            }
        }
    case -9223372036854775808:
        var v22 int32 = 0
        _ = v22
    default:
        var v23 string = "a"
        _ = v23
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v24 float32 = (((((1e10 - v3) * (0.1 - v3)) + 0.1) * (float32(3.14159) * v3)) * float32(0.0))
        _ = v24
    }
    if true {
        v25 := GenS3{F0: 1, F1: 65535}
        fmt.Println(v3, v15, v15)
        _ = v25
    } else {
        switch ("\n\t" + "a") {
        case "hello":
            v3 -= v2
        default:
            switch 9223372036854775807 {
            case 1:
                var v26 uint8 = 2
                _ = v26
            case 9223372036854775807:
                v2 += (1.5 - -2.25)
            case 42:
                var v27 string = ""
                _ = v27
            default:
                var v28 uint64 = 42
                _ = v28
            }
        }
        var v29 uint8 = 2
        _ = v29
    }
    var v30 uint8 = 128
    fmt.Println(v3, v30, v30)
    for fi31 := 0; fi31 < 4; fi31++ {
        if ((fi31 % ((fi31) | 1)) > (fi31 / ((fi31) | 1))) {
            v15 = ((0.1 * v15) * 1e10)
            v30 += v30
        }
        v32 := GenS4{F0: 1, F1: -9223372036854775808, F2: 198}
        _ = fi31
        _ = v32
    }
    v30 -= v30
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
