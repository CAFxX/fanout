package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 string; F2 int16 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 uint }
func genFn0(p0 int8, p1 int64) int8 {
    fmt.Println(p0, p1)
    p1 -= (p1 + p1)
    p0 += (p0 % (((p0 >> ((1) & 63))) | 1))
    var v5 int8 = p0
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(v5, p1)
        var v7 int32 = -832
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return (p0 >> ((4294967295) & 63))
}
func genFn1() uint8 {
    switch 128 {
    case 127:
        v8 := GenS1{F0: genFn0(0, -1)}
        _ = v8
    case 1:
        if false {
            if true {
                v9 := []int{32767, -1, 128}
                switch 32768 {
                case 727:
                    for fi10 := 0; fi10 < 7; fi10++ {
                        for fi11 := 0; fi11 < 9; fi11++ {
                            fmt.Println(fi11)
                            v12 := GenS1{F0: 56}
                            var v13 int32 = 2147483647
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        var v14 string = "hello"
                        fmt.Println(fi10, fi10)
                        _ = fi10
                        _ = v14
                    }
                case -2147483648:
                    _ = len(v9)
                case 17:
                    var v15 int32 = 0
                    _ = v15
                default:
                    var v16 uint8 = 1
                    _ = v16
                }
                if len(v9) > 0 {
                    v9[0] = 2147483647
                }
                _ = v9
            }
        } else {
            if (-9223372036854775808 < 1) {
                if false {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        var v18 uint64 = 627
                        _ = fi17
                        _ = v18
                    }
                } else {
                    if false {
                        v19 := []int{32767}
                        for fi20 := 0; fi20 < 13; fi20++ {
                            var v21 float32 = 0.0
                            if len(v19) > 0 {
                                v19[0] = fi20
                            }
                            v21 -= 1.5
                            _ = fi20
                            _ = v21
                        }
                        var v22 int32 = -2147483648
                        _ = v19
                        _ = v22
                    } else {
                        var v23 uint = 0
                        v24 := []int{2147483647, 555}
                        var v25 uint = 42
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    if true {
                        if (!false) {
                            var v26 uint = 18446744073709551615
                            v27 := []int{32768, 4294967295}
                            var v28 int64 = -9223372036854775808
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            if (0.0 >= 0.1) {
                                var v29 int64 = -9223372036854775808
                                switch (("foo bar" + "foo bar") + "a") {
                                case "hello":
                                    fmt.Println(v29)
                                case "":
                                    v30 := []int{}
                                    _ = v30
                                }
                                _ = v29
                            }
                            if true {
                                var v31 uint32 = 2147483647
                                _ = v31
                            } else {
                                var v32 uint32 = 2147483647
                                v33 := []int{-1, 32768, -1}
                                fmt.Println(v32)
                                _ = v32
                                _ = v33
                            }
                        }
                        var v34 uint8 = 0
                        var v35 bool = (127 < 128)
                        _ = v34
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 16; fi36++ {
                            var v37 uint16 = 255
                            fmt.Println(fi36)
                            _ = fi36
                            _ = v37
                        }
                        if (true && false) {
                            if true {
                                v38 := GenS1{F0: 81}
                                var v39 int16 = -981
                                _ = v38
                                _ = v39
                            }
                        }
                        var v40 bool = false
                        _ = v40
                    }
                }
                if false {
                    var v41 int16 = -291
                    fmt.Println(v41)
                    _ = v41
                }
                var v42 uint16 = 65535
                _ = v42
            } else {
                for fi43 := 0; fi43 < 14; fi43++ {
                    v44 := GenS2{F0: 2147483647, F1: "\n\t", F2: 1}
                    for fi45 := 0; fi45 < 12; fi45++ {
                        var v46 uint64 = 1
                        v46 = (^(v46))
                        _ = fi45
                        _ = v46
                    }
                    v47 := []int{}
                    _ = fi43
                    _ = v44
                    _ = v47
                }
                var v48 int = 127
                _ = v48
            }
            var v49 float32 = (((3.14159 * (1e10 * 3.14159)) / 1.0) / (1e10 + (1e10 / 1.0)))
            _ = v49
        }
    case 963:
        v50 := []int{1, 448}
        _ = v50
    default:
        var v51 uint = 2
        _ = v51
    }
    switch 32768 {
    case 9223372036854775807:
        v52 := GenS3{F0: 65535, F1: 153}
        _ = v52
    }
    v53 := GenS4{F0: 42, F1: 9223372036854775807}
    _ = v53
    return 205
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int16 = 0
    v3 := make(map[int]bool)
    var v4 int8 = 31
    v2 = 0
    _ = len(v3)
    switch "\n\t" {
    case "":
        if (723 != 0) {
            fmt.Println(v2)
            if true {
                var v54 int64 = 333
                switch (-(v1)) {
                case 4294967295:
                    fmt.Println(v4)
                default:
                    fmt.Println(v1)
                }
                _ = v54
            }
            for fi55 := 0; fi55 < 6; fi55++ {
                v1 = (v1 & fi55)
                if true {
                    if (100 == (v4 - genFn0(v4, -5))) {
                        v2 = v2
                    }
                    var v56 uint64 = 0
                    var v57 int64 = -1
                    _ = v56
                    _ = v57
                }
                v2 -= (v2 >> ((1) & 63))
                _ = fi55
            }
        } else {
            var v58 float32 = (0.1 * (1e10 / 3.14159))
            v58 = (float32(0.0) * 0.1)
            v59 := GenS1{F0: v4}
            _ = v58
            _ = v59
        }
    default:
        if ((v1 % ((1) | 1)) > int(v2)) {
            v60 := GenS5{F0: 147, F1: 1}
            _ = v60
        } else {
            if false {
                v2 = v2
                v1 = (-(v1))
                var v61 uint8 = 65
                _ = v61
            } else {
                delete(v3, v1)
                v4 = v4
                switch (v1 + 4294967295) {
                case 776:
                    v1 = v1
                case -1:
                    switch "a" {
                    case "hello":
                        var v62 int64 = -1
                        _ = v62
                    case "a":
                        if true {
                            v2 = v2
                            var v63 bool = ((0.1 != 0.1) || (4294967295 != 18446744073709551615))
                            if true {
                                v63 = (!false)
                                var v64 uint = 4294967295
                                for fi65 := 0; fi65 < 18; fi65++ {
                                    v3[(fi65 & v1)] = v63
                                    _ = fi65
                                }
                                _ = v64
                            } else {
                                var v66 int64 = 0
                                v1 = (-(v1))
                                _ = v66
                            }
                            _ = v63
                        } else {
                            v67 := GenS6{F0: 65535, F1: -716, F2: 431}
                            _ = v67
                        }
                    case "foo bar":
                        if false {
                            var v68 bool = false
                            v1 -= v1
                            v2 = (^(v2))
                            _ = v68
                        }
                    }
                case -616:
                    v1 -= (v1 >> ((255) & 63))
                }
            }
            v4 = (v4 | 42)
        }
    }
    var v69 int32 = 2147483647
    v4 -= (v4 - 1)
    var v70 float64 = 0.0
    var v71 uint64 = 4294967295
    if false {
        for fi72 := 0; fi72 < 10; fi72++ {
            _ = len(v3)
            var v73 uint8 = 1
            _ = fi72
            _ = v73
        }
        for fi74 := 0; fi74 < 8; fi74++ {
            switch "\n\t" {
            case "foo bar":
                switch v1 {
                case 365:
                    var v75 uint8 = 243
                    _ = v75
                default:
                    v76 := GenS7{F0: (v2 & v2)}
                    _ = v76
                }
            }
            _ = fi74
        }
    }
    v77 := GenS8{F0: genFn0((^(v4)), 42), F1: 172}
    v71 = 304
    if (0 == v4) {
        var v78 string = "\n\t"
        v79 := v3[(v1 | 42)]
        _ = v78
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
    _ = v71
    _ = v77
}
