package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool, p1 int16) string {
    switch -9223372036854775808 {
    case -1:
        var v6 bool = p0
        _ = v6
    case -9223372036854775808:
        if true {
            var v7 uint64 = 1
            v8 := GenS1{F0: 1, F1: -1}
            var v9 uint64 = 468
            _ = v7
            _ = v8
            _ = v9
        } else {
            switch 32768 {
            case 4294967295:
                p0 = (602 < 128)
            case -2147483648:
                var v10 float64 = 0.0
                _ = v10
            case 1:
                var v11 uint32 = 2147483648
                _ = v11
            default:
                v12 := GenS1{F0: 32767, F1: 42}
                _ = v12
            }
            var v13 uint8 = 144
            _ = v13
        }
    case 127:
        p0 = p0
    default:
        v14 := []int{-747, -1, 1}
        _ = v14
    }
    v15 := GenS1{F0: 32767, F1: 1}
    fmt.Println(p1, p1)
    if (false && ((-2.25 * 1e10) >= 1e10)) {
        var v16 int32 = -2147483648
        _ = v16
    }
    v17 := []int{4294967295, 633, 32768, 1}
    _ = p0
    _ = p1
    _ = v15
    _ = v17
    return "hello"
}

func main() {
    var v1 uint16 = 1
    var v2 uint32 = 1
    v3 := make(map[int]string)
    var v4 float32 = -2.25
    var v5 int16 = 32767
    v2 += v2
    var v18 uint = 4294967295
    v2 = v2
    var v19 string = "foo bar"
    v2 = (v2 | (^(v2)))
    var v20 float64 = 3.14159
    switch -9223372036854775808 {
    case 1:
        v4 = (float32((v20 + v20)) + -2.25)
    case 42:
        v21 := GenS2{F0: 1, F1: -1}
        _ = v21
    case -135:
        for fi22 := 0; fi22 < 2; fi22++ {
            v4 -= v4
            _ = fi22
        }
    }
    switch 590 {
    case -2147483648:
        v19 = v19
    case 1:
        for fi23 := 0; fi23 < 16; fi23++ {
            switch (fi23 + fi23) {
            case 32767:
                var v24 uint32 = 1
                _ = v24
            case 127:
                var v25 int8 = 108
                _ = v25
            case -194:
                v19 = ((v19 + ("" + v19)) + ("\n\t" + "foo bar"))
            }
            v26 := GenS3{F0: 40}
            _ = fi23
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
