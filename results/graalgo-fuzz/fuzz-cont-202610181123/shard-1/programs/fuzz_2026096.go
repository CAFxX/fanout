package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 uint32; F1 int }
func genFn0(p0 int32) uint64 {
    var v7 uint32 = 2147483647
    var v8 uint8 = 27
    v7 = v7
    _ = p0
    _ = v7
    _ = v8
    return uint64(v8)
}

func main() {
    var v1 uint8 = 0
    var v2 int32 = 0
    var v3 int16 = 1
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    v6 := make(map[string]int)
    v2 = 2147483647
    var v9 uint16 = 56
    v10 := []int{}
    _ = len(v6)
    var v11 string = ("foo bar" + "a")
    fmt.Println(v9, v11)
    v12 := v5[0]
    v1 -= 127
    if false {
        var v13 float32 = (float32(1.5) / 3.14159)
        _ = v13
    }
    if (true && ("a" >= v11)) {
        v1 = v1
    } else {
        var v14 int32 = (v2 | (v2 << ((237) & 63)))
        v15 := GenS1{F0: v9}
        _ = v14
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 float64 = (1e10 * 0.0)
            var v18 float64 = v17
            fmt.Println(v18, v9, v11)
            _ = fi16
            _ = v17
            _ = v18
        }
        switch 1 {
        case -9223372036854775808:
            v19 := GenS2{F0: (v1 % (((v1 - 127)) | 1)), F1: (v1 >> ((2) & 63))}
            _ = v19
        case 42:
            if true {
                for fi20 := 0; fi20 < 15; fi20++ {
                    fmt.Println(v9, v2)
                    switch (fi20 % (((fi20 + fi20)) | 1)) {
                    case 128:
                        if true {
                            for fi21 := 0; fi21 < 7; fi21++ {
                                v1 = 255
                                v12 = (v11 + v11)
                                if (42 < 0) {
                                    switch "\n\t" {
                                    case "a":
                                        v22 := GenS3{F0: -128, F1: (v1 % ((v1) | 1))}
                                        _ = v22
                                    case "":
                                        v1 = (^(v1))
                                    default:
                                        switch (fi20 ^ fi20) {
                                        case 2147483647:
                                            var v23 int32 = (v2 - (v2 & v2))
                                            _ = v23
                                        default:
                                            _ = len(v5)
                                            if (4294967295 <= 65535) {
                                                fmt.Println(v12)
                                                if false {
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                    var v24 float64 = (3.14159 / (0.1 - float64(1e10)))
                                                    v6[v12] = fi20
                                                    _ = v24
                                                }
                                                v3 -= v3
                                            }
                                            v2 += v2
                                        }
                                    }
                                    if (!true) {
                                        v25 := GenS4{F0: -9223372036854775808, F1: -895}
                                        if false {
                                            var v26 int16 = v3
                                            _ = v26
                                        }
                                        v27 := GenS1{F0: v9}
                                        _ = v25
                                        _ = v27
                                    } else {
                                        var v28 int32 = v2
                                        _ = v28
                                    }
                                }
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                v2 = 2147483647
                                _ = fi21
                            }
                            for fi29 := 0; fi29 < 2; fi29++ {
                                fmt.Println(v3)
                                _ = fi29
                            }
                            for fi30 := 0; fi30 < 13; fi30++ {
                                v31 := v4[(fi20 + (fi20 ^ (-(fi30))))]
                                delete(v6, v11)
                                var v32 uint64 = 0
                                v33 := v6[v12]
                                if (v32 != genFn0((v2 ^ v2))) {
                                    v12 = v12
                                    v11 = "\n\t"
                                }
                                _ = fi30
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        } else {
                            var v34 int32 = (-(v2))
                            _ = v34
                        }
                    case 42:
                        fmt.Println(v11)
                    case 27:
                        v11 = (v11 + ((((("foo bar" + ("a" + "a")) + "\n\t") + v12) + (v11 + v12)) + "hello"))
                    default:
                        switch ("hello" + ("" + v12)) {
                        case "foo bar":
                            v5[362] = v11
                        default:
                            delete(v5, 1)
                        }
                    }
                    _ = fi20
                }
            }
        default:
            var v35 int8 = 100
            _ = v35
        }
        _ = len(v6)
    }
    if (true && (!false)) {
        if true {
            switch 4294967295 {
            case 311:
                v11 = v12
            case 365:
                var v36 int = 1
                _ = v36
            default:
                switch 253 {
                case 1:
                    switch 2147483647 {
                    case 42:
                        var v37 uint64 = 2
                        _ = v37
                    case 32768:
                        v38 := GenS5{F0: 1, F1: 4294967295}
                        _ = v38
                    }
                case -79:
                    var v39 uint16 = (v9 - v9)
                    _ = v39
                default:
                    var v40 int64 = 1
                    _ = v40
                }
            }
            switch 127 {
            case 2147483647:
                if (true || true) {
                    if len(v10) > 0 {
                        v10[0] = 367
                    }
                    var v41 float64 = float64(3.14159)
                    v1 += (v1 << ((0) & 63))
                    _ = v41
                }
            case 4294967295:
                switch "foo bar" {
                case "a":
                    v1 = (v1 ^ uint8(v9))
                }
            case 285:
                switch "a" {
                case "":
                    fmt.Println(v9)
                case "\n\t":
                    _ = len(v6)
                case "hello":
                    v42 := GenS5{F0: 65535, F1: 42}
                    _ = v42
                default:
                    fmt.Println(v2)
                }
            }
        } else {
            var v43 int32 = v2
            var v44 int16 = 42
            _ = v43
            _ = v44
        }
        if false {
            var v45 uint16 = v9
            v46 := GenS2{F0: (v1 + (v1 + (v1 ^ 133))), F1: v1}
            for fi47 := 0; fi47 < 20; fi47++ {
                switch "foo bar" {
                case "\n\t":
                    v11 = "a"
                }
                var v48 float64 = -2.25
                fmt.Println(v2)
                _ = fi47
                _ = v48
            }
            _ = v45
            _ = v46
        }
        v10 = append(v10, 216)
    }
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
