package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }
func genFn0() int32 {
    v7 := GenS1{F0: 0, F1: ""}
    var v8 uint32 = 0
    var v9 uint64 = 42
    if false {
        var v10 uint32 = 1
        if true {
            v11 := []int{}
            if (-9223372036854775808 <= 349) {
                fmt.Println(v10, v9, v10)
            }
            v12 := GenS2{F0: (v8 << ((2) & 63)), F1: true, F2: -128}
            _ = v11
            _ = v12
        }
        _ = v10
    } else {
        v8 = (^(v8))
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := []int{(fi14 >> ((18446744073709551615) & 63))}
        v16 := GenS2{F0: 2147483648, F1: true, F2: 1}
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}
func genFn1(p0 string, p1 int64) uint8 {
    if (!((404 >= 32768) || false)) {
        fmt.Println(p1, p0)
        v17 := []int{128, 1, 32767}
        _ = v17
    }
    v18 := GenS1{F0: p1, F1: ""}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v18
    return 255
}

func main() {
    var v1 bool = (!(false && true))
    var v2 uint = 0
    var v3 uint8 = 104
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    var v6 int8 = -1
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 float64 = (1.5 * float64(1.5))
        var v21 int8 = 113
        _ = fi19
        _ = v20
        _ = v21
    }
    v3 = (^(v3))
    v22 := []int{9223372036854775807, -483, 9223372036854775807, 32767}
    if len(v22) > 0 {
        v22[0] = -2147483648
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v22 = append(v22, fi23)
        v24 := GenS2{F0: 1, F1: v1, F2: 59}
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := GenS1{F0: -1, F1: "hello"}
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    switch "a" {
    case "foo bar":
        var v27 float32 = (3.14159 + ((0.0 - 0.0) / 1.0))
        _ = v27
    case "a":
        var v28 uint8 = v3
        _ = v28
    default:
        v5[-754] = (!(2147483648 <= 0))
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        switch int(v6) {
        case 128:
            v3 = v3
        case 0:
            v2 += v2
        case 236:
            v6 -= (v6 ^ v6)
        }
        for fi30 := 0; fi30 < 14; fi30++ {
            if ((true && (0 > 42)) || v1) {
                var v31 float64 = (0.1 * (1e10 - float64(0.1)))
                if v1 {
                    fmt.Println(v3)
                }
                _ = len(v4)
                switch ((("\n\t" + "") + "a") + "hello") {
                case "hello":
                    var v32 int16 = -32768
                    _ = v32
                case "a":
                    if false {
                        v22 = append(v22, fi29)
                    } else {
                        var v33 uint8 = genFn1((((("foo bar" + "hello") + "hello") + "hello") + (("a" + "\n\t") + "foo bar")), 1)
                        var v34 int16 = 793
                        _ = v33
                        _ = v34
                    }
                    var v35 float64 = v31
                    _ = v35
                case "":
                    v36 := v5[957]
                    _ = v36
                default:
                    var v37 string = "foo bar"
                    _ = v37
                }
                _ = v31
            }
            fmt.Println(v6, fi29)
            _ = fi30
        }
        v6 = (v6 % ((v6) | 1))
        _ = fi29
    }
    v2 -= 42
    v1 = v1
    v2 = (v2 * (v2 >> ((42) & 63)))
    var v38 uint32 = 0
    if false {
        switch "foo bar" {
        case "a":
            var v39 int8 = (v6 << ((v2) & 63))
            _ = v39
        case "foo bar":
            var v40 int64 = -9223372036854775808
            _ = v40
        case "":
            var v41 uint8 = v3
            _ = v41
        default:
            if v1 {
                v42 := v5[int(v2)]
                _ = v42
            }
        }
    } else {
        fmt.Println(v3)
        delete(v5, 9223372036854775807)
        fmt.Println(v2, v3)
    }
    v3 -= genFn1("\n\t", -9223372036854775808)
    switch 9223372036854775807 {
    case 704:
        switch 127 {
        case 128:
            switch 466 {
            case -564:
                for fi43 := 0; fi43 < 2; fi43++ {
                    v5[fi43] = v1
                    if false {
                        v5[-396] = (-356 == -1)
                    }
                    if (0.0 >= 0.0) {
                        var v44 int32 = 42
                        for fi45 := 0; fi45 < 6; fi45++ {
                            if ((1 >= -496) && (!v1)) {
                                v46 := GenS1{F0: int64(v38), F1: ("\n\t" + ("a" + "foo bar"))}
                                v3 = (v3 | v3)
                                _ = v46
                            } else {
                                var v47 float32 = 1.5
                                _ = len(v22)
                                fmt.Println(v6)
                                fmt.Println(v38, v2, v2)
                                _ = v47
                            }
                            var v48 uint16 = 0
                            v49 := GenS3{F0: 18446744073709551615, F1: fi45}
                            var v50 int32 = (v44 % ((genFn0()) | 1))
                            _ = fi45
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                        for fi51 := 0; fi51 < 18; fi51++ {
                            var v52 int = 9223372036854775807
                            if (v1 || true) {
                                if (!v1) {
                                    v53 := GenS4{F0: 2, F1: 863, F2: "a"}
                                    v1 = false
                                    v54 := GenS2{F0: v38, F1: v1, F2: (v6 - (v6 + (v6 >> (((v2 ^ v2)) & 63))))}
                                    _ = v53
                                    _ = v54
                                }
                            }
                            v4[128] = ("\n\t" + "foo bar")
                            _ = fi51
                            _ = v52
                        }
                        _ = v44
                    } else {
                        var v55 string = ("a" + "hello")
                        fmt.Println(v6)
                        _ = v55
                    }
                    _ = fi43
                }
            case 42:
                fmt.Println(v1, v6, v6)
            case 0:
                var v56 uint8 = (^(v3))
                _ = v56
            default:
                switch 9223372036854775807 {
                case -9223372036854775808:
                    var v57 uint = (^(v2))
                    _ = v57
                default:
                    v22 = append(v22, -1)
                }
            }
        case 9223372036854775807:
            var v58 float64 = -2.25
            _ = v58
        }
    default:
        switch ("a" + "hello") {
        case "foo bar":
            v3 -= 2
        default:
            v2 += v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
