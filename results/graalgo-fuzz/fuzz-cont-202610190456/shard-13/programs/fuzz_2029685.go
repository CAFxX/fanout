package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 int64, p1 int64) bool {
    fmt.Println(p1)
    p1 = 0
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 int32, p1 uint16) uint16 {
    var v5 int8 = -1
    v6 := []int{647}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p1
}

func main() {
    var v1 uint32 = 801
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 float32 = 0.1
    if false {
        fmt.Println(v1, v4)
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 float64 = -2.25
            v8 -= v8
            _ = fi7
            _ = v8
        }
    }
    v1 -= 779
    var v9 int64 = -9223372036854775808
    if ((((2 == 128) && false) || (true || false)) && false) {
        for fi10 := 0; fi10 < 17; fi10++ {
            switch (-(fi10)) {
            case -25:
                v9 -= 155
            case 905:
                var v11 uint64 = 2
                _ = v11
            }
            fmt.Println(v4, v1, fi10)
            v12 := []int{(fi10 / ((fi10) | 1))}
            _ = fi10
            _ = v12
        }
        var v13 uint8 = 127
        var v14 float32 = 1e10
        _ = v13
        _ = v14
    } else {
        switch "" {
        case "hello":
            for fi15 := 0; fi15 < 3; fi15++ {
                fmt.Println(fi15)
                var v16 uint = 0
                _ = fi15
                _ = v16
            }
        case "a":
            v17 := v2[127]
            _ = v17
        default:
            v18 := []int{}
            _ = v18
        }
    }
    v19 := GenS1{F0: 100}
    for fi20 := 0; fi20 < 14; fi20++ {
        v4 = float32(-2.25)
        var v21 int = fi20
        v22 := GenS2{F0: (("a" + "a") + ("foo bar" + ("" + "\n\t"))), F1: ((!(1 == 0)) && (!true)), F2: 0}
        _ = fi20
        _ = v21
        _ = v22
    }
    v1 += v1
    for fi23 := 0; fi23 < 3; fi23++ {
        v4 += v4
        if (!(!true)) {
            v24 := []int{fi23, -2147483648, fi23}
            if (!true) {
                v25 := GenS3{F0: 4294967295, F1: 415}
                v26 := GenS2{F0: "a", F1: ((float32((((float64(0.0) * 0.1) / 0.1) * 0.1)) * (v4 - v4)) != (((v4 / v4) + v4) - v4)), F2: 0}
                _ = v25
                _ = v26
            } else {
                var v27 int8 = 127
                _ = v27
            }
            if false {
                var v28 int = 32767
                var v29 int16 = -32768
                _ = v28
                _ = v29
            } else {
                v2[(fi23 << ((961) & 63))] = 186
                v30 := GenS3{F0: 49, F1: 705}
                switch fi23 {
                case -2147483648:
                    var v31 float64 = (3.14159 * float64(1e10))
                    _ = v31
                }
                _ = v30
            }
            _ = v24
        }
        if true {
            v32 := []int{(fi23 + fi23), 865}
            v33 := v3[fi23]
            fmt.Println(v9, fi23)
            _ = v32
            _ = v33
        } else {
            var v34 int = fi23
            for fi35 := 0; fi35 < 10; fi35++ {
                v4 += (-2.25 + v4)
                switch "hello" {
                case "":
                    var v36 uint32 = v1
                    _ = v36
                }
                _ = fi35
            }
            _ = v34
        }
        _ = fi23
    }
    v2[128] = (v1 ^ (v1 + v1))
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
