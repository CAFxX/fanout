package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int; F2 uint64 }
type GenS7 struct { F0 int16; F1 int32; F2 int32 }
func genFn0(p0 uint64) int8 {
    v5 := []int{}
    var v6 int32 = 0
    _ = p0
    _ = v5
    _ = v6
    return -101
}
func genFn1(p0 int64, p1 string, p2 int16) uint16 {
    switch -276 {
    case -1:
        v7 := GenS1{F0: 0, F1: 128}
        _ = v7
    case 1:
        fmt.Println(p0, p0)
    default:
        v8 := []int{2147483647, -407, -405}
        _ = v8
    }
    var v9 uint = 0
    var v10 float64 = 1.5
    v10 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 114
}

func main() {
    v1 := []string{"", "hello"}
    var v2 bool = (0 < -101)
    v3 := make(map[uint32]string)
    var v4 bool = (0 > 86)
    v11 := GenS2{F0: 2}
    v2 = true
    _ = len(v3)
    var v12 uint32 = 42
    for fi13 := 0; fi13 < 6; fi13++ {
        if v4 {
            v4 = (!("a" >= "hello"))
            v14 := GenS3{F0: "hello", F1: -38}
            v15 := GenS4{F0: (!v2)}
            _ = v14
            _ = v15
        } else {
            v12 -= v12
            if false {
                var v16 int = fi13
                _ = v16
            }
        }
        _ = fi13
    }
    if (v4 || false) {
        var v17 bool = (!(65535 >= v12))
        _ = v17
    }
    v18 := GenS4{F0: (((4294967295 >= 866) && true) || (32767 < 2))}
    var v19 int64 = 42
    var v20 uint = 294
    v12 -= (^(v12))
    if v2 {
        for fi21 := 0; fi21 < 4; fi21++ {
            switch fi21 {
            case -328:
                v22 := GenS3{F0: (("a" + ("foo bar" + "\n\t")) + "hello"), F1: 0}
                _ = v22
            case 698:
                for fi23 := 0; fi23 < 7; fi23++ {
                    if (-105 > -81) {
                        v20 -= 42
                        var v24 float32 = float32(-2.25)
                        _ = v24
                    }
                    _ = fi23
                }
            case 32768:
                for fi25 := 0; fi25 < 18; fi25++ {
                    var v26 bool = false
                    _ = fi25
                    _ = v26
                }
            default:
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 float64 = (-2.25 + 1e10)
                    v12 -= 708
                    var v29 int32 = 42
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            v2 = v2
            switch fi21 {
            case 439:
                v1 = append(v1, ("foo bar" + ("hello" + "")))
            case 42:
                v30 := GenS5{F0: 141}
                _ = v30
            case 127:
                v2 = v2
            }
            _ = fi21
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            v4 = ((!v2) || (1 <= 149))
            _ = fi31
        }
        _ = len(v3)
    } else {
        for fi32 := 0; fi32 < 1; fi32++ {
            if (-2.25 == float64(0.0)) {
                if v4 {
                    if v2 {
                        if ((3.14159 / 3.14159) == -2.25) {
                            fmt.Println(v2, v4, fi32)
                        }
                        var v33 int32 = -243
                        if v4 {
                            switch (fi32 / ((fi32) | 1)) {
                            case 9223372036854775807:
                                for fi34 := 0; fi34 < 19; fi34++ {
                                    v35 := GenS6{F0: 0, F1: (fi34 / ((fi32) | 1)), F2: 0}
                                    _ = fi34
                                    _ = v35
                                }
                            default:
                                for fi36 := 0; fi36 < 11; fi36++ {
                                    if true {
                                        v37 := GenS7{F0: 824, F1: v33, F2: v33}
                                        fmt.Println(v19, v4, v12)
                                        if true {
                                            if v2 {
                                                v20 += (v20 * 284)
                                            } else {
                                                var v38 float32 = float32((-2.25 * 1.5))
                                                switch "\n\t" {
                                                case "hello":
                                                    delete(v3, 689)
                                                default:
                                                    fmt.Println(v4)
                                                }
                                                _ = v38
                                            }
                                        } else {
                                            if v4 {
                                                var v39 string = ""
                                                _ = v39
                                            }
                                            var v40 uint16 = genFn1(v19, "\n\t", 42)
                                            var v41 int16 = 32767
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = v37
                                    } else {
                                        var v42 float64 = 0.0
                                        var v43 uint16 = 0
                                        var v44 float64 = (v42 - (float64(((float32(v42) + (3.14159 - (0.1 * 3.14159))) * (1.5 * 1e10))) + 1.5))
                                        var v45 uint = 984
                                        v33 += (v33 * int32(fi32))
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                    _ = fi36
                                }
                            }
                            v12 = (v12 * (v12 | (v12 & v12)))
                        }
                        _ = v33
                    }
                } else {
                    switch "" {
                    case "foo bar":
                        if (0 < 74) {
                            var v46 int16 = 32767
                            v20 = (^(v20))
                            switch fi32 {
                            case 787:
                                if v2 {
                                    if true {
                                        v47 := GenS2{F0: v20}
                                        v4 = ("" <= "foo bar")
                                        var v48 uint = (v20 | 255)
                                        _ = v47
                                        _ = v48
                                    }
                                } else {
                                    v49 := v3[(v12 ^ (v12 % (((v12 ^ 4294967295)) | 1)))]
                                    v49 = v49
                                    _ = v49
                                }
                            case 32768:
                                v20 += 18446744073709551615
                            case 330:
                                var v50 int64 = (v19 - -9223372036854775808)
                                _ = v50
                            }
                            _ = v46
                        }
                    }
                }
            } else {
                _ = len(v1)
            }
            _ = fi32
        }
        v2 = false
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v2, v2)
        default:
            for fi51 := 0; fi51 < 9; fi51++ {
                for fi52 := 0; fi52 < 12; fi52++ {
                    v4 = (!(v4 && (v2 || (!v2))))
                    v2 = ((0.0 * (0.1 + (((1e10 + float64(-2.25)) + (-2.25 / 1.5)) / ((((-2.25 / 3.14159) / 1.0) - (0.0 - -2.25)) + float64(-2.25))))) >= 1.5)
                    fmt.Println(v4)
                    switch (^(fi52)) {
                    case -421:
                        var v53 float64 = -2.25
                        _ = v53
                    default:
                        fmt.Println(fi52)
                    }
                    v19 += (^(v19))
                    _ = fi52
                }
                var v54 uint = 2
                var v55 int8 = -10
                _ = fi51
                _ = v54
                _ = v55
            }
        }
    }
    for fi56 := 0; fi56 < 20; fi56++ {
        var v57 int8 = int8(v12)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v58 uint64 = 42
        _ = fi56
        _ = v57
        _ = v58
    }
    v59 := GenS7{F0: 1000, F1: 995, F2: 966}
    delete(v3, (v12 * (^(v12))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v59
}
