package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0(p0 int16, p1 uint8) int32 {
    p1 = p1
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 = 1000
        var v6 int32 = -1
        _ = fi5
        _ = v6
    }
    v7 := []int{2147483647, 32768, 145, 127}
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    return -2147483648
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 2
    v3 := []int{}
    var v4 float64 = 1e10
    v8 := GenS1{F0: (-2.25 < -2.25)}
    v4 += -2.25
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: ((fi9 & (fi9 % (((fi9 * 9223372036854775807)) | 1))) == fi9)}
        _ = fi9
        _ = v10
    }
    if true {
        v11 := GenS2{F0: 371, F1: ""}
        for fi12 := 0; fi12 < 7; fi12++ {
            if (278 <= 279) {
                v13 := GenS2{F0: -9223372036854775808, F1: ""}
                v14 := GenS3{F0: 2147483647, F1: (("\n\t" + ("\n\t" + "")) + "foo bar")}
                _ = v13
                _ = v14
            } else {
                v2 -= (v2 | v2)
                v4 += v4
                if false {
                    v1[("" + "\n\t")] = 42
                    v15 := v1["hello"]
                    _ = v15
                } else {
                    var v16 float64 = 0.0
                    _ = v16
                }
            }
            _ = fi12
        }
        _ = v11
    } else {
        var v17 uint64 = 1
        var v18 float64 = v4
        var v19 string = (("a" + "") + ((("\n\t" + "foo bar") + (("foo bar" + "") + ("a" + ("a" + ("" + "hello"))))) + "hello"))
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 float64 = 1.5
    fmt.Println(v20, v2, v2)
    v21 := GenS4{F0: 440, F1: 0}
    var v22 uint = 1
    var v23 uint32 = uint32(v22)
    var v24 int = 1
    _ = len(v1)
    if len(v3) > 0 {
        v3[0] = 42
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        if (float32(3.14159) <= 1e10) {
            if (!(!false)) {
                v26 := GenS3{F0: (v23 / ((v23) | 1)), F1: (("a" + (((("hello" + "a") + "a") + ("\n\t" + ("a" + ""))) + "foo bar")) + "foo bar")}
                var v27 uint32 = (v23 * (v23 | (v23 << ((v22) & 63))))
                v28 := GenS3{F0: uint32(v2), F1: "\n\t"}
                _ = v26
                _ = v27
                _ = v28
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = len(v3)
                fmt.Println(v20, fi25, v22)
            }
            v22 = v22
            v29 := GenS2{F0: 9223372036854775807, F1: "hello"}
            _ = v29
        }
        var v30 uint8 = 0
        _ = fi25
        _ = v30
    }
    var v31 int = v24
    if (v23 <= v23) {
        v20 = 1e10
        v3 = append(v3, -776)
    } else {
        v22 += (v22 * (v22 % ((451) | 1)))
    }
    var v32 int = (v31 - (^(v31)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v8
}
