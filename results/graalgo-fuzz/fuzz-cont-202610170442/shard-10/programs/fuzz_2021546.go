package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int8; F2 uint }
func genFn0(p0 int64, p1 int64) int8 {
    if true {
        p1 -= p0
        for fi6 := 0; fi6 < 4; fi6++ {
            for fi7 := 0; fi7 < 2; fi7++ {
                p0 = p0
                v8 := GenS1{F0: (fi6 / ((fi7) | 1))}
                _ = fi7
                _ = v8
            }
            if (-9223372036854775808 != p1) {
                if (!((fi6 >= fi6) && true)) {
                    fmt.Println(p1, p1)
                    p1 -= p0
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v10 := GenS1{F0: fi9}
                        p0 -= 0
                        _ = fi9
                        _ = v10
                    }
                }
                fmt.Println(p0, fi6, p0)
            }
            _ = fi6
        }
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint16 = 0
        _ = fi11
        _ = v12
    }
    v13 := []int{-9223372036854775808, 4294967295}
    v13 = append(v13, 440)
    _ = p0
    _ = p1
    _ = v13
    return 16
}
func genFn1(p0 int32, p1 uint32, p2 uint64) int8 {
    for fi14 := 0; fi14 < 1; fi14++ {
        if false {
            fmt.Println(fi14, p1)
            p2 += 38
        } else {
            switch "hello" {
            case "foo bar":
                p0 += 33
            case "a":
                v15 := GenS1{F0: (fi14 - (fi14 << ((42) & 63)))}
                _ = v15
            case "\n\t":
                v16 := []int{fi14, (fi14 - -2147483648), 32768, fi14}
                _ = v16
            default:
                var v17 float64 = float64((float32(1e10) - (0.0 + (0.1 + -2.25))))
                _ = v17
            }
            var v18 int16 = -146
            _ = v18
        }
        p1 = 2147483648
        _ = fi14
    }
    if true {
        fmt.Println(p2, p1, p0)
        var v19 float64 = 3.14159
        if (false || (!(("foo bar" + "a") <= (("hello" + "foo bar") + ("foo bar" + "hello"))))) {
            p1 = p1
            p2 -= (p2 / ((p2) | 1))
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 float32 = -2.25
                fmt.Println(v21, fi20, p2)
                var v22 uint = 1
                _ = fi20
                _ = v21
                _ = v22
            }
            var v23 string = ("\n\t" + "hello")
            _ = v23
        }
        _ = v19
    } else {
        v24 := GenS1{F0: -1}
        var v25 int8 = 127
        fmt.Println(v25, p2)
        _ = v24
        _ = v25
    }
    var v26 float32 = (1.5 + 1e10)
    var v27 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return 100
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 int64 = -9223372036854775808
    v4 := []uint32{65535, 2147483648, 42}
    v5 := []int{-391, 4294967295, 4294967295}
    v28 := GenS1{F0: 1}
    v29 := GenS1{F0: 0}
    var v30 int8 = genFn0(682, 42)
    v30 -= (-(v30))
    var v31 uint32 = 1
    for fi32 := 0; fi32 < 2; fi32++ {
        v3 += -1
        _ = fi32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        fmt.Println(v3)
        v4 = append(v4, (v31 * (v31 - (v31 * (v31 % (((v31 & v31)) | 1))))))
        _ = fi33
    }
    fmt.Println(v31, v31, v30)
    if false {
        if (!(127 > 4294967295)) {
            v31 += v31
        }
    }
    v5 = append(v5, 590)
    v34 := GenS2{F0: -384, F1: (v30 ^ (v30 & v30)), F2: 951}
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
