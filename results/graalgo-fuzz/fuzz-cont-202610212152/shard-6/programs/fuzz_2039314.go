package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint }
func genFn0(p0 string, p1 int64) string {
    switch p0 {
    case "\n\t":
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 int8 = -1
                _ = fi6
                _ = v7
            }
            p1 += (p1 * 1)
        } else {
            fmt.Println(p1)
            var v8 int = 32768
            if (0 == 42) {
                p1 += (p1 ^ 9223372036854775807)
                var v9 float32 = ((-2.25 + (((3.14159 + 0.1) * 3.14159) / 1.0)) / float32(float64(0.1)))
                _ = v9
            }
            _ = v8
        }
    case "foo bar":
        v10 := []int{600}
        _ = v10
    case "a":
        p1 = -855
    default:
        p1 -= (p1 * p1)
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := GenS1{F0: 0, F1: p0, F2: 1}
        p0 = (p0 + "hello")
        _ = fi11
        _ = v12
    }
    if false {
        var v13 int = -859
        v14 := []int{(^(v13)), v13, (v13 * (-(v13)))}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        p0 = ("foo bar" + ("" + p0))
        _ = fi15
    }
    _ = p0
    _ = p1
    return ""
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = -2147483648
    var v3 int32 = (v2 % ((2147483647) | 1))
    var v4 float32 = ((0.0 - 1.5) * ((1e10 - float32(3.14159)) - 1.5))
    v5 := make(map[int]uint32)
    var v16 uint64 = 1
    v17 := GenS2{F0: 32767}
    v18 := []int{127, -9223372036854775808, 32767}
    var v19 uint8 = 29
    v20 := v5[4294967295]
    v18 = append(v18, 4294967295)
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            fmt.Println(v19)
            v23 := GenS3{F0: v3, F1: 65535}
            var v24 uint16 = 498
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    v3 = v2
    delete(v5, -9223372036854775808)
    fmt.Println(v1)
    if (true && false) {
        if (((0 == -9223372036854775808) && false) && (v16 >= v16)) {
            if len(v18) > 0 {
                v18[0] = -318
            }
            for fi25 := 0; fi25 < 8; fi25++ {
                v1 = v1
                _ = fi25
            }
            var v26 int64 = 9223372036854775807
            _ = v26
        }
        v4 += (1e10 - float32(1e10))
        var v27 int32 = -2147483648
        _ = v27
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        for fi29 := 0; fi29 < 6; fi29++ {
            v4 = 3.14159
            var v30 float64 = 1.5
            var v31 int8 = -109
            _ = fi29
            _ = v30
            _ = v31
        }
        if (88 == v19) {
            _ = len(v18)
            var v32 int = 4294967295
            _ = v32
        } else {
            var v33 int16 = int16(v20)
            _ = v33
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
