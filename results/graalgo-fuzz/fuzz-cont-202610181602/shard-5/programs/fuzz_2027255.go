package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int64; F1 int64; F2 int8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64) int16 {
    var v6 int64 = p0
    v7 := GenS1{F0: 42, F1: 42, F2: 0}
    if (127 != -97) {
        v8 := GenS1{F0: 954, F1: -2147483648, F2: 527}
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return 254
}
func genFn1() int16 {
    v9 := GenS2{F0: 631, F1: "hello"}
    var v10 uint8 = 1
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 string = "hello"
        if true {
            v12 = v12
        } else {
            switch (-(fi11)) {
            case -675:
                if (466 == 522) {
                    for fi13 := 0; fi13 < 4; fi13++ {
                        v10 = (v10 | (v10 | (v10 | 118)))
                        var v14 uint64 = 42
                        var v15 string = (("\n\t" + "") + v12)
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
            default:
                var v16 uint = 18446744073709551615
                _ = v16
            }
        }
        v10 -= v10
        _ = fi11
        _ = v12
    }
    var v17 string = ("a" + "\n\t")
    _ = v10
    _ = v17
    _ = v9
    return 1000
}

func main() {
    v1 := []uint32{262, 0}
    v2 := make(map[string]string)
    var v3 int64 = 166
    v4 := []int{}
    var v5 int16 = 17
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: fi18, F1: "foo bar"}
        fmt.Println(v5, v5, fi18)
        fmt.Println(v3, v3, v3)
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        v3 -= 42
        if false {
            switch fi20 {
            case 1:
                var v21 string = ""
                _ = v21
            case 377:
                var v22 int8 = 117
                _ = v22
            }
            if false {
                for fi23 := 0; fi23 < 18; fi23++ {
                    _ = len(v4)
                    _ = fi23
                }
            }
            if (false && ((2 < 141) || (float32(-2.25) <= ((((3.14159 - 3.14159) / 0.1) + 0.0) + -2.25)))) {
                switch "\n\t" {
                case "a":
                    v3 += 830
                case "\n\t":
                    if false {
                        var v24 float64 = (3.14159 * ((3.14159 * (-2.25 + float64(3.14159))) * 0.1))
                        for fi25 := 0; fi25 < 18; fi25++ {
                            var v26 uint32 = 2
                            v4 = append(v4, fi20)
                            v5 = v5
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    } else {
                        v27 := GenS1{F0: 173, F1: -2147483648, F2: 65535}
                        _ = v27
                    }
                default:
                    var v28 string = ("foo bar" + "hello")
                    _ = v28
                }
                var v29 float64 = 1e10
                _ = v29
            }
        } else {
            fmt.Println(v3, v5)
            var v30 string = "\n\t"
            _ = v30
        }
        _ = fi20
    }
    v31 := GenS2{F0: -1, F1: ("\n\t" + "foo bar")}
    switch 439 {
    case 32768:
        var v32 uint64 = 627
        _ = v32
    case 0:
        v33 := GenS3{F0: v3, F1: 0, F2: 100}
        _ = v33
    case -47:
        var v34 uint32 = 42
        _ = v34
    }
    v3 = v3
    var v35 int32 = 408
    if true {
        fmt.Println(v5, v5)
        v36 := GenS4{F0: 2147483647}
        switch -1 {
        case 944:
            var v37 int64 = (v3 ^ (v3 * (v3 ^ (v3 - (v3 - (v3 / ((v3) | 1)))))))
            _ = v37
        case 32768:
            var v38 int64 = -1
            _ = v38
        default:
            fmt.Println(v35)
        }
        _ = v36
    }
    if false {
        _ = len(v1)
        fmt.Println(v5, v35)
        v3 -= int64(v5)
    }
    var v39 int = 130
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 int8 = -1
        fmt.Println(v39, v41)
        v35 = -1
        _ = fi40
        _ = v41
    }
    var v42 string = "a"
    if true {
        v39 = v39
    } else {
        fmt.Println(v5, v42, v39)
    }
    var v43 int16 = (v5 - (v5 + (v5 % (((v5 >> ((868) & 63))) | 1))))
    v43 -= 0
    var v44 int = (^(v39))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
