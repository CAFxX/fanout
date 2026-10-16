package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint }
func genFn0(p0 uint16, p1 uint) uint {
    var v4 float64 = 1e10
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            switch (fi5 | fi6) {
            case 0:
                switch 32768 {
                case 42:
                    fmt.Println(p0, p1)
                default:
                    v7 := GenS1{F0: 1000, F1: 0}
                    _ = v7
                }
            case 2147483647:
                var v8 int = int(p1)
                _ = v8
            case -161:
                fmt.Println(p0, p0)
            }
            _ = fi6
        }
        _ = fi5
    }
    v9 := GenS1{F0: 0, F1: 245}
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return (p1 + 495)
}
func genFn1(p0 uint8) uint32 {
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint16 = 32767
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        p0 = (p0 ^ 116)
        v13 := GenS2{F0: fi12}
        _ = fi12
        _ = v13
    }
    _ = p0
    return 810
}

func main() {
    var v1 int8 = -66
    var v2 float32 = 1.5
    var v3 float64 = 1.5
    if (!(!true)) {
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 uint = 626
            v16 := GenS1{F0: -150, F1: 312}
            if false {
                v17 := GenS3{F0: int32(v1), F1: (^(v15))}
                v18 := GenS2{F0: 32768}
                v19 := GenS3{F0: 1, F1: v15}
                _ = v17
                _ = v18
                _ = v19
            } else {
                fmt.Println(fi14, v15, v15)
                v2 -= v2
                v20 := GenS3{F0: -2147483648, F1: 0}
                _ = v20
            }
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        if true {
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := []int{-366, fi21}
                fmt.Println(fi21)
                var v23 uint64 = 508
                _ = fi21
                _ = v22
                _ = v23
            }
            fmt.Println(v1, v2, v2)
        } else {
            v2 = 3.14159
        }
    }
    v24 := GenS2{F0: -9223372036854775808}
    var v25 float32 = (v2 * ((v2 * float32((v3 - v3))) - 0.1))
    var v26 int32 = 0
    fmt.Println(v2, v2, v1)
    var v27 int8 = v1
    var v28 float32 = (v25 - v2)
    v29 := GenS3{F0: (v26 << ((42) & 63)), F1: genFn0(643, 4294967295)}
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 int8 = v27
        var v32 int64 = 42
        v33 := []int{fi30, fi30}
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    fmt.Println(v28, v27, v25)
    v2 += v28
    if (!(255 <= 127)) {
        for fi34 := 0; fi34 < 1; fi34++ {
            v35 := GenS3{F0: -1, F1: 4294967295}
            var v36 float32 = v2
            fmt.Println(v27)
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    var v37 uint16 = 1
    var v38 uint = 232
    var v39 int32 = (v26 / (((v26 + -2147483648)) | 1))
    v40 := []int{128, 1, -215}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
