package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int8; F1 int64; F2 uint8 }
func genFn0(p0 bool) int32 {
    fmt.Println(p0)
    v6 := []int{}
    _ = p0
    _ = v6
    return -581
}
func genFn1(p0 int) uint8 {
    p0 = (p0 ^ p0)
    if (100 < -128) {
        v7 := GenS1{F0: -2147483648}
        var v8 int8 = -1
        _ = v7
        _ = v8
    } else {
        var v9 float64 = 1e10
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = v10
    return 206
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = false
    var v3 int32 = 291
    var v4 bool = v2
    var v5 uint64 = 42
    v2 = (0 >= 2)
    var v11 uint64 = (^(v5))
    var v12 string = "hello"
    if (false && false) {
        if v4 {
            var v13 int16 = 0
            switch (v12 + ((("a" + v12) + ("a" + (v12 + v12))) + v12)) {
            case "foo bar":
                var v14 float32 = 1e10
                _ = v14
            }
            var v15 uint32 = 147
            _ = v13
            _ = v15
        } else {
            var v16 uint = 0
            v5 -= (v5 << (((v16 / ((uint(v5)) | 1))) & 63))
            _ = v16
        }
        v1[("a" + "\n\t")] = 65535
    } else {
        var v17 int8 = 42
        v18 := []int{32767, 127, -2147483648, -2147483648}
        var v19 int64 = -1
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 int32 = v3
    v21 := GenS2{F0: v12}
    v11 -= 961
    v22 := v1[((("hello" + v12) + v12) + v12)]
    var v23 uint = 42
    v22 -= (v22 << ((65535) & 63))
    if ("\n\t" >= "\n\t") {
        if (!false) {
            fmt.Println(v5, v22)
        } else {
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 int8 = 0
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 int8 = v25
                    var v28 string = v12
                    v29 := GenS1{F0: v20}
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = fi24
                _ = v25
            }
            fmt.Println(v5, v2, v12)
            var v30 uint = (v23 | v23)
            _ = v30
        }
    } else {
        if false {
            fmt.Println(v23, v11, v5)
            switch v12 {
            case "\n\t":
                for fi31 := 0; fi31 < 4; fi31++ {
                    fmt.Println(v2, v20)
                    var v32 int64 = 0
                    _ = fi31
                    _ = v32
                }
            default:
                v33 := v1["a"]
                _ = v33
            }
            fmt.Println(v20, v5, v12)
        } else {
            v34 := GenS3{F0: 0, F1: -9223372036854775808}
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 float32 = (float32((1.5 * ((3.14159 / 0.1) * (1.5 - (1e10 - (((-2.25 / -2.25) * (0.0 + 1.5)) + 3.14159)))))) / ((1.5 * float32(1.5)) + 1.5))
        _ = fi35
        _ = v36
    }
    switch v12 {
    case "\n\t":
        v37 := GenS2{F0: (("foo bar" + v12) + (("" + v12) + ("" + v12)))}
        _ = v37
    }
    fmt.Println(v3, v11)
    v38 := GenS4{F0: -38, F1: -566, F2: 163}
    v20 -= int32(v22)
    var v39 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
