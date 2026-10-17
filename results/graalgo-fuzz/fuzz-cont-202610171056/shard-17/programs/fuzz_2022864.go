package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 int64) int8 {
    v7 := GenS1{F0: 42, F1: 1, F2: ("\n\t" + "a")}
    p0 = (p0 / (((^(p0))) | 1))
    p0 -= -945
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 int8) int8 {
    switch "\n\t" {
    case "\n\t":
        p0 = p0
    case "hello":
        v8 := GenS1{F0: 1, F1: p0, F2: (("\n\t" + "foo bar") + (("\n\t" + "") + "foo bar"))}
        _ = v8
    default:
        p0 = (p0 + -1)
    }
    fmt.Println(p0)
    _ = p0
    return p0
}

func main() {
    var v1 uint64 = 0
    v2 := []int{-1, 128, 127}
    var v3 float32 = float32(((1.5 * 0.0) / float64(1e10)))
    var v4 int = -2147483648
    var v5 uint32 = 999
    v6 := make(map[string]int)
    if (!true) {
        switch "foo bar" {
        case "":
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: -1, F1: genFn1(100), F2: "\n\t"}
                v11 := GenS2{F0: 1}
                var v12 uint = 0
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        default:
            var v13 uint32 = (v5 - (v5 % (((v5 | v5)) | 1)))
            _ = v13
        }
        v3 -= (((v3 - v3) + float32(0.1)) / v3)
    }
    var v14 int32 = 693
    var v15 int = (v4 >> ((uint(v1)) & 63))
    switch 1 {
    case 281:
        var v16 int = (-(v4))
        _ = v16
    case -779:
        var v17 float64 = 0.1
        _ = v17
    default:
        v15 = v15
    }
    if (!(v4 >= -9223372036854775808)) {
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := GenS2{F0: -9223372036854775808}
            v4 -= (v4 - -1)
            switch (("" + "") + "\n\t") {
            case "\n\t":
                if (42 <= 0) {
                    var v20 float32 = ((v3 / float32(0.1)) * (1.5 / v3))
                    _ = v20
                } else {
                    var v21 bool = true
                    v22 := GenS3{F0: 42}
                    _ = v21
                    _ = v22
                }
            default:
                if (1.5 <= 0.0) {
                    v15 += fi18
                    var v23 uint8 = 0
                    _ = v23
                } else {
                    var v24 int64 = 327
                    if (v5 > v5) {
                        var v25 float64 = 3.14159
                        _ = len(v2)
                        _ = v25
                    } else {
                        fmt.Println(fi18)
                        switch (v15 % ((v15) | 1)) {
                        case 42:
                            fmt.Println(v14, v24, v24)
                        case 926:
                            v24 += -1
                        case -1:
                            v3 = (((float32(1.5) * 0.1) * -2.25) * float32(1.5))
                        }
                    }
                    _ = v24
                }
            }
            _ = fi18
            _ = v19
        }
        var v26 float32 = -2.25
        v4 = v4
        _ = v26
    } else {
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 4; fi28++ {
                v15 += 9223372036854775807
                _ = fi28
            }
            _ = fi27
        }
        var v29 uint = 590
        fmt.Println(v15, v4, v5)
        _ = v29
    }
    var v30 int32 = (v14 - 42)
    switch (v15 | v15) {
    case -582:
        var v31 string = "foo bar"
        _ = v31
    case 0:
        var v32 string = "foo bar"
        _ = v32
    case 2147483647:
        v1 -= v1
    default:
        var v33 string = ""
        _ = v33
    }
    v34 := GenS4{F0: "foo bar", F1: 210}
    var v35 uint32 = (v5 / (((v5 | (v5 - v5))) | 1))
    v14 = (-(v30))
    var v36 int16 = 1000
    fmt.Println(v3, v14, v4)
    v37 := GenS3{F0: 2}
    v2 = append(v2, v15)
    var v38 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
