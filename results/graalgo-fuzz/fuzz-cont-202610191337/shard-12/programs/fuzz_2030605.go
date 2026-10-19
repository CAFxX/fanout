package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int64; F2 uint }
type GenS4 struct { F0 int64; F1 int }
func genFn0(p0 int, p1 uint32) string {
    p0 -= -648
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float32 = ((float32(-2.25) - 0.1) * 3.14159)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return (((("\n\t" + "") + "a") + "foo bar") + "")
}
func genFn1(p0 int32) int {
    p0 -= (-(p0))
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := []int{fi9, (fi9 % (((fi9 + (^(fi9)))) | 1))}
        v11 := GenS1{F0: 724}
        var v12 uint32 = 48
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    return -1
}

func main() {
    var v1 int64 = 449
    var v2 int32 = 0
    var v3 uint8 = 2
    v4 := make(map[string]uint32)
    var v5 string = "\n\t"
    var v6 uint16 = 2
    if (100 < 9) {
        var v13 float32 = (0.0 / ((3.14159 - -2.25) - float32((1.5 * (0.1 * 3.14159)))))
        var v14 int8 = 0
        _ = v13
        _ = v14
    }
    v2 += v2
    fmt.Println(v5, v6)
    if false {
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := GenS2{F0: uint(v6)}
            var v17 float64 = (3.14159 + (1e10 * (1e10 - -2.25)))
            for fi18 := 0; fi18 < 18; fi18++ {
                v6 -= v6
                _ = fi18
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        if false {
            fmt.Println(v2, v5, v1)
            for fi19 := 0; fi19 < 8; fi19++ {
                v3 -= (v3 - 255)
                _ = fi19
            }
        } else {
            var v20 string = v5
            v6 += 318
            for fi21 := 0; fi21 < 19; fi21++ {
                if (v5 > (v20 + v20)) {
                    v1 = v1
                }
                _ = fi21
            }
            _ = v20
        }
    }
    delete(v4, "")
    if false {
        v22 := GenS3{F0: false, F1: (-(v1)), F2: 484}
        switch (v5 + "") {
        case "foo bar":
            for fi23 := 0; fi23 < 13; fi23++ {
                v24 := GenS4{F0: v1, F1: genFn1((v2 >> ((255) & 63)))}
                _ = fi23
                _ = v24
            }
        }
        v25 := []int{1}
        _ = v22
        _ = v25
    }
    var v26 int16 = 0
    var v27 int64 = 9223372036854775807
    switch ((v5 + "\n\t") + (((v5 + v5) + "") + ((v5 + v5) + v5))) {
    case "foo bar":
        v28 := GenS4{F0: (v27 % (((v1 % (((v1 - (^(v1)))) | 1))) | 1)), F1: 128}
        _ = v28
    case "":
        v29 := GenS4{F0: (v27 / (((v1 / (((v27 & v27)) | 1))) | 1)), F1: 2147483647}
        _ = v29
    case "hello":
        var v30 float64 = -2.25
        _ = v30
    default:
        if ((v6 * v6) != (v6 | 42)) {
            fmt.Println(v3, v27, v26)
            v5 = v5
            delete(v4, v5)
        } else {
            v31 := GenS1{F0: 18446744073709551615}
            v1 = (v27 ^ (v1 & (v1 + (-(v1)))))
            var v32 int32 = (v2 >> ((4294967295) & 63))
            _ = v31
            _ = v32
        }
    }
    var v33 uint64 = 4294967295
    var v34 int = -788
    for fi35 := 0; fi35 < 1; fi35++ {
        if (1.5 <= (0.1 * (float32(float64(0.0)) + 3.14159))) {
            v6 += 1
            var v36 uint = 4294967295
            _ = v36
        } else {
            v1 -= v27
            var v37 int16 = 82
            _ = v37
        }
        v6 = 255
        fmt.Println(v34)
        _ = fi35
    }
    var v38 string = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
