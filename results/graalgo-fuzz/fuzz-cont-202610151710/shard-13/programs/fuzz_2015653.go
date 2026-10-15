package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
func genFn0(p0 bool, p1 string) int16 {
    if p0 {
        for fi4 := 0; fi4 < 16; fi4++ {
            switch p1 {
            case "foo bar":
                p1 = p1
            case "a":
                var v5 string = (("\n\t" + "a") + "foo bar")
                _ = v5
            default:
                if p0 {
                    fmt.Println(fi4, p0, fi4)
                    p0 = (p1 >= p1)
                    var v6 uint = 2
                    _ = v6
                } else {
                    var v7 string = "a"
                    _ = v7
                }
            }
            switch p1 {
            case "hello":
                var v8 uint16 = 32767
                _ = v8
            case "":
                var v9 int16 = 32767
                _ = v9
            }
            switch (("foo bar" + p1) + ((p1 + p1) + (p1 + p1))) {
            case "foo bar":
                var v10 int32 = 576
                _ = v10
            case "\n\t":
                v11 := []int{(fi4 / (((fi4 * 0)) | 1)), fi4, fi4, (fi4 & fi4)}
                _ = v11
            default:
                var v12 uint32 = 2147483647
                _ = v12
            }
            _ = fi4
        }
    } else {
        v13 := []int{4294967295, 358, 4294967295, -9223372036854775808}
        _ = v13
    }
    v14 := GenS1{F0: 1, F1: 722, F2: 377}
    var v15 bool = false
    p1 = p1
    for fi16 := 0; fi16 < 4; fi16++ {
        p0 = p0
        if true {
            fmt.Println(fi16, p0, fi16)
            v17 := []int{(^(fi16)), fi16}
            _ = v17
        } else {
            p1 = (p1 + ((p1 + "a") + "hello"))
            var v18 float32 = (1.5 - (-2.25 * (1e10 * 0.0)))
            v19 := []int{(fi16 / ((fi16) | 1)), fi16, 32767, fi16}
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 32767
}

func main() {
    var v1 int32 = -2147483648
    var v2 float32 = (float32(-2.25) - 0.1)
    var v3 bool = (1 >= 42)
    v3 = ((v1 + -1) < v1)
    v20 := GenS2{F0: ((100 <= 1) || (v2 < (v2 / ((v2 * ((-2.25 + 1e10) * v2)) + v2))))}
    fmt.Println(v3, v3)
    fmt.Println(v1, v3)
    v3 = (!v3)
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := []int{(fi21 ^ fi21), fi21, (fi21 ^ (^(fi21)))}
        v1 = 1
        _ = fi21
        _ = v22
    }
    fmt.Println(v1, v1)
    fmt.Println(v1, v1, v3)
    switch (("" + "foo bar") + ((("a" + "a") + "") + "hello")) {
    case "\n\t":
        v3 = (!v3)
    case "":
        var v23 uint32 = 2147483647
        _ = v23
    default:
        v2 = float32(0.1)
    }
    if (!(47 >= 8)) {
        v24 := GenS1{F0: 0, F1: 2147483647, F2: 42}
        _ = v24
    }
    v1 = (v1 << ((42) & 63))
    v3 = false
    v3 = (!false)
    if (1000 <= 564) {
        if v3 {
            v25 := GenS3{F0: (v1 - (v1 / ((1) | 1))), F1: 429, F2: 493}
            v26 := []int{98, 128, 9223372036854775807, -2147483648}
            _ = v25
            _ = v26
        }
    }
    v27 := []int{-704, 32767, 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
}
