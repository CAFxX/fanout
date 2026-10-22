package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int }
func genFn0(p0 int8) int8 {
    p0 = (p0 | 42)
    var v7 float32 = 1e10
    v7 += (0.0 * v7)
    _ = p0
    _ = v7
    return p0
}
func genFn1() string {
    var v8 uint32 = 2147483647
    fmt.Println(v8)
    var v9 uint64 = 4294967295
    var v10 uint8 = 127
    v9 -= v9
    _ = v10
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 string = "hello"
    var v2 uint64 = 2
    var v3 float32 = float32(float64(0.1))
    v4 := []string{"hello", v1, (("foo bar" + v1) + v1), v1}
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    v11 := GenS1{F0: -1, F1: 65535}
    if true {
        switch 42 {
        case -365:
            if (!(!false)) {
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 int32 = -246
                    if len(v4) > 0 {
                        v4[0] = ("hello" + v1)
                    }
                    v2 -= (v2 & v2)
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 uint8 = 255
                    if true {
                        switch fi14 {
                        case 42:
                            var v16 string = ("hello" + v1)
                            _ = v16
                        case 557:
                            var v17 int16 = 1000
                            _ = v17
                        case 32767:
                            if true {
                                for fi18 := 0; fi18 < 5; fi18++ {
                                    var v19 int16 = 42
                                    _ = fi18
                                    _ = v19
                                }
                                fmt.Println(v1)
                                switch (fi14 / ((fi14) | 1)) {
                                case 32768:
                                    var v20 int16 = 644
                                    _ = v20
                                default:
                                    v15 = v15
                                }
                            } else {
                                v1 = ((v1 + (v1 + v1)) + "a")
                                fmt.Println(v1, v1)
                            }
                        }
                        var v21 int8 = genFn0(-1)
                        var v22 uint64 = v2
                        _ = v21
                        _ = v22
                    } else {
                        v1 = v1
                        if ((628 < 2) || (uint8(v2) <= 0)) {
                            _ = len(v6)
                            if true {
                                for fi23 := 0; fi23 < 10; fi23++ {
                                    v24 := GenS2{F0: (((v1 + v1) + v1) + v1)}
                                    v25 := GenS1{F0: 253, F1: 987}
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                v26 := GenS2{F0: v1}
                                _ = v26
                            }
                        } else {
                            v15 = v15
                            v27 := GenS2{F0: (v1 + "\n\t")}
                            for fi28 := 0; fi28 < 20; fi28++ {
                                var v29 uint32 = 2147483647
                                v29 += (^(v29))
                                var v30 int8 = int8(v29)
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                            _ = v27
                        }
                    }
                    _ = fi14
                    _ = v15
                }
                fmt.Println(v3, v3)
            } else {
                v31 := GenS2{F0: v1}
                v32 := GenS3{F0: 4294967295, F1: -1}
                _ = v31
                _ = v32
            }
        default:
            var v33 int32 = 266
            _ = v33
        }
        v34 := GenS3{F0: 620, F1: -864}
        fmt.Println(v2, v3, v1)
        _ = v34
    } else {
        if false {
            for fi35 := 0; fi35 < 18; fi35++ {
                v6[fi35] = v1
                v3 = -2.25
                _ = fi35
            }
            v3 += v3
            _ = len(v4)
        } else {
            var v36 float32 = ((v3 / (float32(float64(((0.0 - float32(((0.0 - ((1e10 - 0.1) / 3.14159)) / -2.25))) * v3))) - (v3 + v3))) / (((v3 / 3.14159) - 0.0) / (v3 / v3)))
            if len(v4) > 0 {
                v4[0] = v1
            }
            _ = v36
        }
        v2 -= v2
    }
    v37 := GenS4{F0: 128}
    if (217 > 255) {
        switch (("foo bar" + ("\n\t" + (v1 + v1))) + v1) {
        case "":
            var v38 uint16 = 884
            _ = v38
        case "hello":
            var v39 uint8 = 42
            _ = v39
        case "foo bar":
            v5[785] = -444
        }
        var v40 uint8 = 255
        _ = v40
    }
    var v41 int32 = -1
    _ = len(v4)
    v2 -= v2
    var v42 int16 = -738
    for fi43 := 0; fi43 < 13; fi43++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v44 int16 = -213
        for fi45 := 0; fi45 < 13; fi45++ {
            switch fi45 {
            case 321:
                v44 = (v42 - v42)
            }
            _ = fi45
        }
        _ = fi43
        _ = v44
    }
    v3 -= 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
