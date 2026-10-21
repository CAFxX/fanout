package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0(p0 int64, p1 int16) int64 {
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(p1)
        _ = fi7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := []int{fi9}
            var v11 int64 = p0
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 32767
    var v3 int = -9223372036854775808
    var v4 int = v3
    var v5 int8 = 100
    if (float32(float64(0.0)) == 0.1) {
        v4 += -985
    } else {
        var v12 uint = 255
        fmt.Println(v4)
        var v13 bool = (128 <= 42)
        _ = v12
        _ = v13
    }
    var v14 uint16 = 6
    var v15 bool = false
    if (!(1 == v14)) {
        v16 := GenS1{F0: ("a" + "foo bar"), F1: "a"}
        v15 = v15
        v17 := []int{(-(v4)), v4, 128}
        _ = v16
        _ = v17
    } else {
        var v18 int64 = 1
        var v19 bool = ((-(v18)) <= v18)
        v19 = v15
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: "\n\t", F1: ("\n\t" + (("hello" + "foo bar") + "a"))}
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int8 = (v5 + (v5 >> ((2) & 63)))
        v3 = 9223372036854775807
        _ = fi21
        _ = v22
    }
    v3 = -161
    _ = len(v1)
    v23 := GenS1{F0: "foo bar", F1: "hello"}
    fmt.Println(v4, v14)
    v3 += (^(v3))
    v4 = -318
    var v24 uint = 18446744073709551615
    v3 -= int(v14)
    if (!v15) {
        if (255 > 2) {
            v1["hello"] = ""
        } else {
            var v25 uint = v24
            v5 = v5
            _ = v25
        }
        switch "a" {
        case "":
            var v26 bool = (!(float32(0.1) == 3.14159))
            _ = v26
        case "a":
            if true {
                if (v15 && v15) {
                    v27 := GenS2{F0: v15}
                    _ = v27
                } else {
                    if v15 {
                        fmt.Println(v15, v5, v24)
                        v4 -= (-(v4))
                    } else {
                        v28 := GenS3{F0: 128, F1: -1}
                        v2 = v2
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 16; fi29++ {
                        for fi30 := 0; fi30 < 7; fi30++ {
                            switch fi29 {
                            case -9223372036854775808:
                                switch (-(fi29)) {
                                case 517:
                                    var v31 float32 = float32((-2.25 * 0.1))
                                    _ = v31
                                case 1:
                                    var v32 int16 = (v2 ^ (v2 & (v2 / ((v2) | 1))))
                                    _ = v32
                                }
                            case 32768:
                                v4 += int(v24)
                            default:
                                var v33 float32 = -2.25
                                _ = v33
                            }
                            fmt.Println(v2)
                            _ = fi30
                        }
                        _ = fi29
                    }
                }
                fmt.Println(v14, v3)
            } else {
                var v34 int64 = 1
                v5 = v5
                for fi35 := 0; fi35 < 16; fi35++ {
                    var v36 string = "hello"
                    _ = fi35
                    _ = v36
                }
                _ = v34
            }
        case "foo bar":
            v1["hello"] = "a"
        default:
            fmt.Println(v14, v3)
        }
        var v37 int16 = 32767
        _ = v37
    }
    if v15 {
        if (-128 <= -90) {
            fmt.Println(v2, v5)
            v15 = ((v24 & v24) < (v24 | (^(v24))))
            v38 := GenS3{F0: 42, F1: 975}
            _ = v38
        } else {
            if false {
                var v39 int = int(v2)
                var v40 int = v4
                _ = v39
                _ = v40
            } else {
                fmt.Println(v24)
                for fi41 := 0; fi41 < 12; fi41++ {
                    switch fi41 {
                    case 663:
                        var v42 int = v3
                        _ = v42
                    case -1:
                        fmt.Println(v15, fi41, v4)
                    }
                    var v43 int64 = -1
                    v2 = v2
                    _ = fi41
                    _ = v43
                }
            }
            v44 := GenS4{F0: v4, F1: (^(v24))}
            for fi45 := 0; fi45 < 11; fi45++ {
                v46 := []int{}
                _ = len(v1)
                _ = fi45
                _ = v46
            }
            _ = v44
        }
        v47 := GenS3{F0: 8, F1: -1}
        v4 = v3
        _ = v47
    } else {
        v48 := GenS5{F0: v2, F1: -978}
        _ = v48
    }
    v4 -= v4
    for fi49 := 0; fi49 < 1; fi49++ {
        if false {
            v5 += v5
            var v50 uint16 = (v14 * (^(v14)))
            v14 = 255
            _ = v50
        }
        v51 := []int{128, (v4 / (((v3 / (((-(v3))) | 1))) | 1)), (v4 / (((fi49 - (v4 ^ -741))) | 1)), -90}
        _ = fi49
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
