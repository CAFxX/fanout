package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 uint16 }
func genFn0() bool {
    v4 := []int{-1, 0, 128}
    var v5 int32 = -2147483648
    v5 = (v5 ^ 1)
    _ = len(v4)
    _ = v4
    _ = v5
    return ("hello" >= "")
}
func genFn1(p0 int64) string {
    for fi6 := 0; fi6 < 10; fi6++ {
        if true {
            var v7 uint = 65535
            var v8 int32 = 1
            v7 = 388
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    var v9 int16 = 0
    if true {
        if (false || (159 != 1)) {
            fmt.Println(v9, p0)
            for fi10 := 0; fi10 < 17; fi10++ {
                v9 = (v9 & v9)
                _ = fi10
            }
            if (42 == 2) {
                switch "\n\t" {
                case "foo bar":
                    v11 := GenS1{F0: -9223372036854775808, F1: (v9 >> ((2) & 63)), F2: p0}
                    _ = v11
                case "":
                    v12 := GenS1{F0: p0, F1: -1, F2: (p0 ^ 9223372036854775807)}
                    _ = v12
                case "a":
                    p0 -= (p0 / (((^(p0))) | 1))
                }
                v9 += v9
                for fi13 := 0; fi13 < 10; fi13++ {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v15 := GenS1{F0: p0, F1: v9, F2: p0}
                        v9 += 1000
                        p0 = p0
                        p0 -= (^(p0))
                        _ = fi14
                        _ = v15
                    }
                    v9 += -89
                    _ = fi13
                }
            }
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 int = (fi16 + fi16)
            v18 := GenS2{F0: 1}
            fmt.Println(v17, fi16)
            _ = fi16
            _ = v17
            _ = v18
        }
        if true {
            p0 = p0
            switch (("" + ("\n\t" + "foo bar")) + "a") {
            case "foo bar":
                var v19 uint32 = 786
                _ = v19
            case "":
                v20 := GenS1{F0: 967, F1: (v9 - -688), F2: -9223372036854775808}
                _ = v20
            case "a":
                fmt.Println(p0)
            default:
                if false {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        var v22 float32 = (0.1 / ((3.14159 / 1e10) * (-2.25 + float32(1e10))))
                        v22 += v22
                        _ = fi21
                        _ = v22
                    }
                } else {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        p0 = p0
                        if false {
                            p0 += 42
                            v24 := GenS1{F0: p0, F1: -905, F2: 812}
                            _ = v24
                        } else {
                            var v25 int = -1
                            switch (fi23 + (fi23 & (fi23 ^ (fi23 * (fi23 << ((1) & 63)))))) {
                            case 9223372036854775807:
                                v26 := []int{}
                                _ = v26
                            case 337:
                                var v27 bool = true
                                _ = v27
                            }
                            var v28 int8 = -10
                            _ = v25
                            _ = v28
                        }
                        var v29 uint8 = 128
                        _ = fi23
                        _ = v29
                    }
                }
            }
            var v30 float32 = 1e10
            _ = v30
        } else {
            if (true || true) {
                v9 -= v9
                for fi31 := 0; fi31 < 13; fi31++ {
                    fmt.Println(p0)
                    _ = fi31
                }
            }
            var v32 uint8 = 128
            _ = v32
        }
    } else {
        var v33 uint8 = 127
        v33 = v33
        p0 = -571
        _ = v33
    }
    _ = p0
    _ = v9
    return ""
}

func main() {
    var v1 int8 = 39
    var v2 float64 = 3.14159
    v3 := make(map[uint32]bool)
    v34 := []int{-1, 0, 42}
    var v35 float64 = ((v2 - (-2.25 * ((1e10 * v2) + v2))) + v2)
    if false {
        if len(v34) > 0 {
            v34[0] = -402
        }
        var v36 uint = 0
        _ = v36
    } else {
        v37 := v3[42]
        var v38 float64 = (-2.25 / 1.5)
        _ = v37
        _ = v38
    }
    v39 := GenS3{F0: 0, F1: 0}
    v40 := v3[2147483647]
    var v41 bool = (("\n\t" + "\n\t") == "hello")
    v42 := GenS4{F0: ("\n\t" + "foo bar")}
    delete(v3, 65535)
    v34 = append(v34, 32767)
    var v43 uint16 = 535
    v43 = 255
    v2 = ((v35 - float64(0.0)) + float64(-2.25))
    for fi44 := 0; fi44 < 19; fi44++ {
        v1 = v1
        _ = fi44
    }
    v45 := GenS5{F0: 152, F1: v43}
    v41 = (!v40)
    var v46 float64 = v35
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v45
    _ = v46
}
