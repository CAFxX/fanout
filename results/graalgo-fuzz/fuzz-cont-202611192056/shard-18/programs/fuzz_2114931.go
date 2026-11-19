package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int; F1 int8; F2 int }
func genFn0(p0 int32) int16 {
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 -= -1
        for fi7 := 0; fi7 < 16; fi7++ {
            switch "" {
            case "foo bar":
                v8 := GenS1{F0: 9223372036854775807}
                _ = v8
            default:
                fmt.Println(p0)
            }
            fmt.Println(fi7, fi6, fi7)
            _ = fi7
        }
        _ = fi6
    }
    fmt.Println(p0)
    _ = p0
    return 32767
}
func genFn1(p0 int32, p1 int) bool {
    var v9 uint = 791
    p1 -= -923
    v10 := []int{-1, 0}
    if len(v10) > 0 {
        v10[0] = (^(p1))
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return false
}

func main() {
    v1 := []uint32{2147483647}
    var v2 string = "\n\t"
    v3 := make(map[int]string)
    var v4 uint64 = 18446744073709551615
    var v5 float32 = 1e10
    for fi11 := 0; fi11 < 2; fi11++ {
        if ("foo bar" < (v2 + ("hello" + ""))) {
            v12 := GenS1{F0: 0}
            if len(v1) > 0 {
                v1[0] = 65535
            }
            _ = v12
        } else {
            if (42 < -128) {
                var v13 int8 = 1
                fmt.Println(fi11, v13)
                var v14 bool = (1 > 42)
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 string = "\n\t"
                _ = fi15
                _ = v16
            }
            var v17 bool = true
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            if true {
                v4 -= v4
                if true {
                    var v19 bool = false
                    v4 += 18446744073709551615
                    v20 := GenS2{F0: 1}
                    _ = v19
                    _ = v20
                } else {
                    var v21 int64 = 1
                    v4 = (v4 + (v4 + 679))
                    _ = v21
                }
            } else {
                var v22 int = -2147483648
                _ = v22
            }
            _ = len(v3)
            _ = fi18
        }
        _ = fi11
    }
    v23 := GenS2{F0: 4294967295}
    v24 := GenS2{F0: v4}
    fmt.Println(v2, v5, v5)
    var v25 float32 = (1e10 + ((((v5 * (v5 / (float32(((0.0 + (0.0 - (-2.25 + ((3.14159 + 0.0) / (1.5 + 1e10))))) - 0.0)) / 0.1))) + float32(-2.25)) - ((0.0 - (1.5 + v5)) + 3.14159)) - v5))
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    if (!(2 >= 128)) {
        if (false || (true && (!true))) {
            v25 -= (v25 + -2.25)
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 uint32 = 112
                v25 = -2.25
                v27 = (v27 << ((1) & 63))
                _ = fi26
                _ = v27
            }
            switch "" {
            case "a":
                v25 = (3.14159 - 1.5)
            default:
                v28 := GenS3{F0: "", F1: v4}
                _ = v28
            }
        } else {
            switch "foo bar" {
            case "a":
                var v29 int64 = 1
                _ = v29
            default:
                if false {
                    v30 := v3[-1]
                    _ = v30
                }
            }
            v3[4294967295] = v2
            v4 = (v4 - (v4 / ((42) | 1)))
        }
        v2 = v2
        v5 = (v5 / v5)
    }
    switch "hello" {
    case "a":
        v25 = ((v25 + 1.5) / float32((0.0 - 0.1)))
    default:
        var v31 uint32 = 2147483648
        _ = v31
    }
    if false {
        switch ("hello" + "a") {
        case "":
            v32 := GenS3{F0: v2, F1: v4}
            _ = v32
        case "hello":
            var v33 float64 = 1e10
            _ = v33
        case "foo bar":
            var v34 int16 = -298
            _ = v34
        }
        var v35 uint32 = 42
        _ = v35
    }
    var v36 uint8 = 255
    var v37 uint16 = 255
    v38 := GenS4{F0: 32767, F1: -1, F2: 128}
    v4 = (v4 / ((v4) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
