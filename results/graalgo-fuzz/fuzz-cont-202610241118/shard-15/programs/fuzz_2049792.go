package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 int) string {
    if (0 == 4294967295) {
        v5 := GenS1{F0: 9223372036854775807}
        fmt.Println(p0)
        var v6 uint = 639
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        if false {
            v8 := []int{(^(fi7)), 1}
            fmt.Println(p0, p0)
            var v9 int32 = 2147483647
            _ = v8
            _ = v9
        } else {
            var v10 bool = (!false)
            p0 += (^(fi7))
            _ = v10
        }
        p0 = (fi7 * (p0 | fi7))
        if false {
            p0 = (-(fi7))
        }
        _ = fi7
    }
    _ = p0
    return "\n\t"
}
func genFn1(p0 int32, p1 uint8) int64 {
    fmt.Println(p1, p1)
    var v11 uint32 = 65535
    for fi12 := 0; fi12 < 7; fi12++ {
        if false {
            var v13 float32 = 1.5
            v11 = (v11 & 0)
            fmt.Println(v11)
            _ = v13
        }
        _ = fi12
    }
    v14 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v14
    return -457
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int32 = -864
    var v4 string = "hello"
    var v15 uint8 = 2
    fmt.Println(v4, v4)
    fmt.Println(v3, v15, v3)
    var v16 uint64 = 529
    var v17 int32 = -380
    if (0.0 <= 1e10) {
        v18 := GenS2{F0: false, F1: 42}
        if false {
            fmt.Println(v3)
            var v19 float64 = -2.25
            _ = v19
        }
        _ = v18
    }
    var v20 string = (v4 + "foo bar")
    if ((((float64(float32(-2.25)) + (1.5 * 1e10)) == float64(((float32(1e10) - 3.14159) * (float32(0.0) + 1e10)))) || (18446744073709551615 == 4294967295)) || (0.0 == ((-2.25 / 3.14159) - 0.0))) {
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint = 255
            var v23 int = (^(fi21))
            fmt.Println(fi21)
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := []int{}
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        if false {
            var v26 float32 = 3.14159
            v1[fi25] = 2147483647
            var v27 int16 = 42
            _ = v26
            _ = v27
        }
        _ = fi25
    }
    v28 := GenS3{F0: 32767, F1: 560}
    if ((9223372036854775807 >= -12) || ((v15 / ((v15) | 1)) != (^(v15)))) {
        var v29 int16 = -959
        if false {
            v30 := GenS3{F0: 42, F1: 255}
            v31 := GenS3{F0: (v29 / ((v29) | 1)), F1: 129}
            _ = v30
            _ = v31
        }
        _ = v29
    }
    v32 := GenS1{F0: 373}
    var v33 bool = (!true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
