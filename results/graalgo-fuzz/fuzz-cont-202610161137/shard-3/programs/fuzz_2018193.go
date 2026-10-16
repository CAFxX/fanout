package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint; F1 string; F2 string }
type GenS3 struct { F0 uint }
func genFn0(p0 int16, p1 int16) int32 {
    var v4 float32 = 1e10
    v4 = -2.25
    switch 9223372036854775807 {
    case 42:
        var v5 uint64 = 1
        _ = v5
    case 0:
        for fi6 := 0; fi6 < 2; fi6++ {
            switch "foo bar" {
            case "\n\t":
                v7 := []int{(fi6 + fi6), (fi6 | (fi6 % (((fi6 | (fi6 >> ((65535) & 63)))) | 1)))}
                _ = v7
            case "a":
                var v8 uint64 = 504
                _ = v8
            default:
                if (true && ((0.1 + (0.0 * float64(((v4 + 3.14159) - v4)))) == (-2.25 * 3.14159))) {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        switch 127 {
                        case 128:
                            switch fi9 {
                            case -1:
                                var v10 int32 = -1
                                _ = v10
                            case 9223372036854775807:
                                var v11 uint = 42
                                _ = v11
                            default:
                                p0 += (p0 + int16(fi9))
                            }
                        }
                        v12 := []int{(^(fi9)), fi6, fi9, fi6}
                        if len(v12) > 0 {
                            v12[0] = (fi6 + (fi9 & int(p0)))
                        }
                        _ = fi9
                        _ = v12
                    }
                    var v13 float32 = float32(1.5)
                    _ = v13
                }
            }
            var v14 uint64 = 2
            _ = fi6
            _ = v14
        }
    }
    if true {
        v4 = 0.0
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1() int16 {
    var v15 uint8 = 2
    var v16 int16 = 32767
    _ = v15
    _ = v16
    return (v16 / (((v16 / (((v16 - (-(v16)))) | 1))) | 1))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 763
    v3 := []string{"\n\t"}
    if false {
        fmt.Println(v2)
        v2 = (v2 * genFn1())
        var v17 uint8 = 1
        _ = v17
    }
    if (1 != 42) {
        v18 := GenS1{F0: 1, F1: 4294967295, F2: (!(!false))}
        _ = len(v1)
        v19 := GenS1{F0: 0, F1: 42, F2: false}
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 6; fi20++ {
            v1["hello"] = true
            var v21 int8 = 92
            _ = fi20
            _ = v21
        }
        v2 = 32767
        if (39 != 127) {
            var v22 bool = true
            _ = v22
        } else {
            v23 := GenS1{F0: 0, F1: 24, F2: (true && true)}
            v2 = (v2 & (v2 / ((32767) | 1)))
            _ = v23
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v24 float64 = (1e10 + (0.0 * 3.14159))
    v25 := GenS2{F0: 2, F1: "hello", F2: "\n\t"}
    if false {
        var v26 float32 = 0.1
        if false {
            fmt.Println(v2, v2)
            v26 = v26
            v1[("a" + ((("a" + "\n\t") + ("" + (((("\n\t" + ("a" + "hello")) + "foo bar") + "\n\t") + ("a" + "hello")))) + "foo bar"))] = true
        }
        _ = v26
    } else {
        var v27 float32 = 1.5
        var v28 uint64 = 171
        _ = v27
        _ = v28
    }
    if false {
        var v29 string = ""
        v30 := GenS1{F0: 9223372036854775807, F1: 65535, F2: false}
        v29 = (v29 + v29)
        _ = v29
        _ = v30
    } else {
        if (genFn0(v2, (-(v2))) <= 0) {
            var v31 int = 32767
            _ = v31
        }
        v32 := GenS1{F0: -9223372036854775808, F1: 727, F2: (genFn0(v2, v2) <= 42)}
        v2 = v2
        _ = v32
    }
    v33 := GenS3{F0: 443}
    var v34 float32 = (((3.14159 - (3.14159 / -2.25)) / 1.0) - float32((v24 - (v24 + ((v24 - 0.0) - v24)))))
    v34 -= -2.25
    var v35 uint16 = 255
    var v36 uint32 = 781
    v35 = v35
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 uint16 = v35
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 uint = 2
            _ = fi39
            _ = v40
        }
        v41 := GenS1{F0: 42, F1: v36, F2: false}
        _ = fi37
        _ = v38
        _ = v41
    }
    _ = len(v1)
    for fi42 := 0; fi42 < 5; fi42++ {
        for fi43 := 0; fi43 < 6; fi43++ {
            var v44 uint = 0
            var v45 int64 = 42
            if len(v3) > 0 {
                v3[0] = ((((("a" + ("a" + "a")) + "a") + "") + "") + ("hello" + "foo bar"))
            }
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = len(v1)
        var v46 int32 = 362
        _ = fi42
        _ = v46
    }
    v47 := GenS3{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v47
}
