package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int32; F2 int8 }

func main() {
    var v1 uint = 42
    var v2 int32 = 760
    var v3 float32 = (1.5 + (((0.1 * (1e10 / (-2.25 - 1e10))) * (3.14159 / 1e10)) - 1e10))
    var v4 uint64 = 1
    v2 = v2
    var v5 uint8 = 42
    fmt.Println(v4)
    v4 = v4
    if false {
        fmt.Println(v4)
        var v6 uint8 = v5
        _ = v6
    }
    if false {
        if true {
            if true {
                var v7 float64 = 1.5
                _ = v7
            }
        } else {
            if true {
                v8 := []int{127}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v8
            } else {
                fmt.Println(v2, v2, v1)
            }
            fmt.Println(v5, v5, v3)
            var v9 bool = ((v5 % ((105) | 1)) >= v5)
            _ = v9
        }
    }
    v10 := GenS1{F0: -662}
    if ((0 < 0) && false) {
        fmt.Println(v3, v4)
        if false {
            var v11 uint = uint(v5)
            v12 := []int{127, 32767, 437}
            var v13 uint32 = 557
            _ = v11
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 8; fi14++ {
                if (!(1 != 100)) {
                    v15 := []int{(^(fi14)), (^(fi14)), -457}
                    _ = v15
                }
                _ = fi14
            }
            var v16 int8 = 42
            _ = v16
        }
    } else {
        if ("foo bar" >= "foo bar") {
            if (!(!true)) {
                if (-128 < -128) {
                    switch "a" {
                    case "":
                        v17 := GenS1{F0: 807}
                        _ = v17
                    case "a":
                        var v18 int8 = -1
                        _ = v18
                    case "hello":
                        if false {
                            if false {
                                v19 := GenS1{F0: int64(v5)}
                                if true {
                                    if true {
                                        v20 := GenS2{F0: ""}
                                        _ = v20
                                    }
                                }
                                _ = v19
                            } else {
                                var v21 uint8 = v5
                                var v22 float32 = v3
                                fmt.Println(v1)
                                _ = v21
                                _ = v22
                            }
                            if (965 != (v4 & v4)) {
                                v23 := GenS2{F0: "hello"}
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    fmt.Println(v3)
                                    for fi25 := 0; fi25 < 2; fi25++ {
                                        if (float64(float32((1.5 - 1e10))) > (float64((0.1 + (1.5 * (-2.25 * v3)))) * ((0.0 / (-2.25 + 3.14159)) * 0.1))) {
                                            var v26 float64 = 0.1
                                            var v27 uint8 = 2
                                            _ = v26
                                            _ = v27
                                        }
                                        var v28 bool = (!true)
                                        v29 := []int{128, (fi25 % ((fi24) | 1)), (fi25 << ((v1) & 63)), fi25}
                                        _ = fi25
                                        _ = v28
                                        _ = v29
                                    }
                                    if false {
                                        v30 := GenS3{F0: ((-9223372036854775808 <= 42) && (2147483647 >= (fi24 | fi24)))}
                                        var v31 float64 = (((float64((float32(-2.25) / (((v3 + ((v3 + (v3 * (v3 - (0.0 + ((float32(0.1) + (v3 * v3)) - float32(0.0)))))) + ((0.1 * (float32(float64(float32(1.5))) - (v3 - v3))) * float32(float64((((0.1 + v3) * (v3 * ((v3 / 1e10) - (float32((1.5 + (((1.5 - (((0.0 * (-2.25 + -2.25)) - float64(v3)) - 1.5)) / (1e10 * float64(float32(1.5)))) / ((0.0 / 3.14159) / float64(v3))))) + 0.0)))) * v3)))))) * (v3 / v3)) / (((3.14159 - 0.1) - v3) - ((v3 - (-2.25 - 1e10)) + float32(0.1)))))) - ((float64((float32(3.14159) * float32(0.1))) * (0.1 / ((1.5 * 3.14159) / 1.0))) + 0.0)) / float64(float32(1e10))) + float64(v3))
                                        v32 := []int{fi24}
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    } else {
                                        v5 -= v5
                                    }
                                    _ = fi24
                                }
                                _ = v23
                            }
                        } else {
                            var v33 int16 = 32767
                            for fi34 := 0; fi34 < 13; fi34++ {
                                if ((!false) && true) {
                                    switch fi34 {
                                    case 792:
                                        var v35 string = ""
                                        _ = v35
                                    }
                                } else {
                                    var v36 uint32 = 2147483647
                                    var v37 uint8 = (v5 + v5)
                                    _ = v36
                                    _ = v37
                                }
                                _ = fi34
                            }
                            v38 := GenS1{F0: 1}
                            _ = v33
                            _ = v38
                        }
                    default:
                        if (true && true) {
                            var v39 int16 = -32768
                            for fi40 := 0; fi40 < 19; fi40++ {
                                v41 := GenS4{F0: "a", F1: (v2 * -578), F2: 1}
                                _ = fi40
                                _ = v41
                            }
                            var v42 string = "foo bar"
                            _ = v39
                            _ = v42
                        } else {
                            var v43 int64 = 0
                            fmt.Println(v1, v3)
                            _ = v43
                        }
                    }
                    fmt.Println(v2, v3, v4)
                    var v44 uint8 = 0
                    _ = v44
                } else {
                    for fi45 := 0; fi45 < 20; fi45++ {
                        v46 := GenS3{F0: ((4294967295 < 288) && (uint8(v4) >= 2))}
                        _ = fi45
                        _ = v46
                    }
                    var v47 uint32 = 65535
                    v3 += (0.0 - v3)
                    _ = v47
                }
                var v48 uint64 = (^(v4))
                _ = v48
            } else {
                var v49 uint8 = (v5 + 151)
                _ = v49
            }
            v5 = (v5 % (((v5 ^ (v5 / ((0) | 1)))) | 1))
            if (-938 < 0) {
                v1 -= (v1 / ((202) | 1))
                var v50 string = (("" + "foo bar") + "\n\t")
                v3 = float32(0.1)
                _ = v50
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                var v51 uint = (^(v1))
                _ = v51
            case "":
                var v52 float32 = (-2.25 / v3)
                _ = v52
            case "foo bar":
                var v53 uint16 = 2
                _ = v53
            default:
                var v54 uint16 = 444
                _ = v54
            }
            v2 -= (v2 & 1)
            if (-1 >= 0) {
                v55 := []int{520, -286, 9223372036854775807, int(v4)}
                _ = v55
            }
        }
    }
    v5 += 1
    fmt.Println(v5, v4)
    if false {
        var v56 uint64 = v4
        var v57 uint16 = 847
        _ = v56
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
