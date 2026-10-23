package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int }
func genFn0(p0 string, p1 uint) bool {
    var v7 uint16 = 255
    var v8 int = -916
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}
func genFn1() int16 {
    if ("hello" > "") {
        if (!false) {
            if (4294967295 != 18446744073709551615) {
                var v9 int16 = 42
                fmt.Println(v9)
                var v10 int16 = v9
                _ = v10
                _ = v9
            } else {
                v11 := GenS1{F0: 2147483647}
                if ((330 < 1) || (4294967295 <= 65535)) {
                    var v12 float32 = float32(-2.25)
                    if (-25 > -1) {
                        v12 = float32(-2.25)
                        v12 = v12
                        if false {
                            v12 = ((0.1 / (float32(1.5) - v12)) * (float32(3.14159) * (v12 * 0.0)))
                            v12 = -2.25
                        } else {
                            fmt.Println(v12)
                        }
                    }
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 int16 = 835
                        v14 = v14
                        if true {
                            v15 := GenS1{F0: 0}
                            if false {
                                v14 += v14
                                v14 += v14
                            } else {
                                fmt.Println(v14)
                            }
                            v16 := GenS1{F0: 0}
                            _ = v15
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                    }
                }
                var v17 int8 = -1
                _ = v11
                _ = v17
            }
            v18 := GenS1{F0: 814}
            var v19 uint64 = 825
            _ = v18
            _ = v19
        } else {
            if (2147483647 >= 65535) {
                switch ("" + "foo bar") {
                case "\n\t":
                    v20 := GenS1{F0: 42}
                    _ = v20
                case "foo bar":
                    v21 := []int{-1, 820, 9223372036854775807}
                    _ = v21
                case "":
                    var v22 float32 = ((1.5 / 1.0) / float32(float64(-2.25)))
                    _ = v22
                }
                var v23 float64 = float64((3.14159 + ((-2.25 * 1e10) * 1.5)))
                fmt.Println(v23)
                _ = v23
            } else {
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 uint8 = 128
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    var v27 int8 = -1
                    v28 := []int{-1}
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
        }
        v29 := []int{127, 32768, 32768}
        _ = v29
    } else {
        v30 := []int{32768, 4294967295, -939, 32767}
        if (138 >= 208) {
            var v31 int64 = 42
            v31 -= (v31 * v31)
            _ = v31
        } else {
            var v32 uint32 = 2147483648
            if true {
                v32 = 169
                switch -889 {
                case -753:
                    var v33 uint64 = 42
                    _ = v33
                case 9223372036854775807:
                    if (true || (false || false)) {
                        v32 += 1
                        for fi34 := 0; fi34 < 17; fi34++ {
                            if false {
                                var v35 int8 = -1
                                fmt.Println(v32, v32, fi34)
                                v36 := GenS2{F0: 776}
                                _ = v35
                                _ = v36
                            }
                            v37 := GenS1{F0: v32}
                            _ = fi34
                            _ = v37
                        }
                    } else {
                        for fi38 := 0; fi38 < 19; fi38++ {
                            for fi39 := 0; fi39 < 14; fi39++ {
                                var v40 float64 = (1e10 / ((3.14159 * (1e10 + 1e10)) - 1.5))
                                _ = fi39
                                _ = v40
                            }
                            var v41 int32 = 1
                            _ = fi38
                            _ = v41
                        }
                        fmt.Println(v32)
                        var v42 uint32 = (v32 % ((2147483648) | 1))
                        _ = v42
                    }
                case 42:
                    v43 := GenS3{F0: 42, F1: 2, F2: 18446744073709551615}
                    _ = v43
                default:
                    var v44 uint64 = 4294967295
                    _ = v44
                }
            } else {
                if (!false) {
                    v45 := GenS3{F0: 255, F1: 0, F2: 65535}
                    _ = v45
                } else {
                    if (3.14159 < float64(0.0)) {
                        var v46 int8 = 127
                        _ = len(v30)
                        _ = v46
                    }
                    v30 = append(v30, 128)
                    v47 := GenS3{F0: 127, F1: 910, F2: 4294967295}
                    _ = v47
                }
                if len(v30) > 0 {
                    v30[0] = -9223372036854775808
                }
            }
            v32 -= (v32 ^ v32)
            _ = v32
        }
        v48 := GenS2{F0: 2}
        _ = v30
        _ = v48
    }
    var v49 string = "\n\t"
    _ = v49
    return 32767
}

func main() {
    var v1 float64 = 1e10
    var v2 float64 = 3.14159
    var v3 float32 = (3.14159 - 3.14159)
    var v4 uint8 = 42
    var v5 string = ("foo bar" + "\n\t")
    v6 := make(map[int]int)
    v2 += v2
    for fi50 := 0; fi50 < 15; fi50++ {
        if (uint64(v4) <= 1) {
            if true {
                var v51 bool = (!false)
                v5 = v5
                _ = v51
            } else {
                fmt.Println(v3)
                for fi52 := 0; fi52 < 7; fi52++ {
                    if false {
                        v5 = v5
                        v53 := GenS2{F0: 2}
                        v54 := GenS3{F0: v4, F1: 2, F2: 862}
                        var v55 int8 = -1
                        v6[337] = (fi52 & fi50)
                        v3 = ((v3 + (v3 + 1.5)) * v3)
                        var v56 uint16 = 0
                        _ = v53
                        _ = v54
                        _ = v55
                        _ = v56
                    } else {
                        v57 := []int{fi50, -9223372036854775808, 4294967295, (fi52 ^ (fi52 / ((810) | 1)))}
                        v58 := GenS4{F0: ("hello" + v5), F1: 0}
                        fmt.Println(v2)
                        _ = v57
                        _ = v58
                    }
                    _ = fi52
                }
                var v59 string = "a"
                _ = v59
            }
            var v60 uint64 = 18446744073709551615
            v61 := []int{fi50, (fi50 << ((35) & 63)), (fi50 >> ((1) & 63)), (fi50 * fi50)}
            _ = v60
            _ = v61
        }
        _ = fi50
    }
    if (0 < 18446744073709551615) {
        v62 := []int{-664, -1, 0, -947}
        _ = v62
    }
    var v63 uint32 = 341
    var v64 uint32 = (^(v63))
    v4 = (v4 & 0)
    var v65 string = ("foo bar" + "foo bar")
    if ((845 == 0) && (0 == uint16(v4))) {
        _ = len(v6)
        var v66 string = "a"
        v67 := GenS5{F0: 127, F1: "\n\t"}
        _ = v66
        _ = v67
    } else {
        _ = len(v6)
    }
    v68 := GenS6{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v68
}
