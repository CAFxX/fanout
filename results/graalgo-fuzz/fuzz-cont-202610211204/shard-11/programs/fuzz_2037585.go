package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 bool }
type GenS3 struct { F0 bool; F1 int32; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int32; F1 int64; F2 uint }
type GenS8 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (0 > 204)
    v3 := make(map[uint32]string)
    var v4 float64 = 0.0
    for fi5 := 0; fi5 < 15; fi5++ {
        v2 = true
        _ = fi5
    }
    v2 = (-663 > -181)
    v6 := GenS1{F0: 255, F1: (0 > 1), F2: (!(!v2))}
    switch "\n\t" {
    case "hello":
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int8 = -128
            v9 := []int{-1, fi7, fi7}
            var v10 uint8 = 37
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    case "":
        var v11 int = 1
        _ = v11
    default:
        if v2 {
            var v12 uint = 4294967295
            v13 := []int{}
            var v14 float64 = v4
            _ = v12
            _ = v13
            _ = v14
        } else {
            var v15 uint8 = 127
            v2 = (true && v2)
            _ = v15
        }
    }
    var v16 float32 = 1e10
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 float32 = v16
        _ = fi17
        _ = v18
    }
    var v19 float64 = v4
    v4 = v19
    if true {
        v20 := []int{32768, 1, 42, -778}
        var v21 int = 9223372036854775807
        _ = v20
        _ = v21
    }
    v4 += float64(v16)
    switch "\n\t" {
    case "hello":
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS2{F0: 703, F1: 65535, F2: (-2147483648 == 42)}
            v24 := []int{fi22, fi22, fi22, (^(fi22))}
            _ = fi22
            _ = v23
            _ = v24
        }
    case "":
        var v25 int64 = 42
        _ = v25
    default:
        var v26 int = 128
        _ = v26
    }
    fmt.Println(v2, v16)
    v2 = v2
    switch ("\n\t" + "") {
    case "":
        if v2 {
            v27 := GenS3{F0: false, F1: 0, F2: 18446744073709551615}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := GenS4{F0: false, F1: 697}
                v2 = (-10 < 127)
                _ = fi28
                _ = v29
            }
        }
    case "a":
        fmt.Println(v2)
    default:
        for fi30 := 0; fi30 < 4; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                v4 -= v19
                v32 := GenS4{F0: v2, F1: 42}
                if (!true) {
                    v33 := GenS2{F0: 2, F1: 65535, F2: (v2 || (v2 || v2))}
                    if (((523 != 0) && v2) && ((4294967295 > 2) || v2)) {
                        if v2 {
                            var v34 uint64 = 0
                            _ = v34
                        } else {
                            if v2 {
                                fmt.Println(v16, fi30)
                            }
                            v35 := GenS5{F0: 42}
                            _ = v35
                        }
                        if v2 {
                            v2 = (!(v2 || ((("hello" + "\n\t") == ("foo bar" + "foo bar")) && (!(("\n\t" + "\n\t") >= ("" + (("a" + "foo bar") + "hello")))))))
                        } else {
                            if (!(-408 >= 1)) {
                                v19 += 1e10
                                delete(v3, 1)
                                v4 -= float64(-2.25)
                            }
                        }
                    } else {
                        v19 += v19
                        v36 := GenS6{F0: 417, F1: 1}
                        if true {
                            fmt.Println(fi31, v16, v16)
                        } else {
                            if false {
                                if (1 > 127) {
                                    v4 -= v4
                                    var v37 int16 = 32767
                                    fmt.Println(fi30, v37)
                                    _ = v37
                                } else {
                                    v38 := []int{fi31, fi30, fi31}
                                    v16 += v16
                                    _ = v38
                                }
                                v39 := GenS4{F0: (!v2), F1: 0}
                                switch (fi31 >> ((42) & 63)) {
                                case 128:
                                    var v40 float32 = v16
                                    _ = v40
                                case -9223372036854775808:
                                    _ = len(v3)
                                default:
                                    if ((2 <= 65535) || true) {
                                        v19 += v4
                                    }
                                }
                                _ = v39
                            }
                            var v41 int = -100
                            switch (v41 + (v41 - fi30)) {
                            case 442:
                                var v42 uint = 1
                                _ = v42
                            case 4294967295:
                                v43 := GenS4{F0: (4294967295 < 333), F1: 524}
                                _ = v43
                            case -2147483648:
                                v44 := GenS5{F0: 971}
                                _ = v44
                            default:
                                v45 := GenS7{F0: -1, F1: 0, F2: 1}
                                _ = v45
                            }
                            _ = v41
                        }
                        _ = v36
                    }
                    _ = v33
                } else {
                    fmt.Println(v16, v2)
                    v1[-9223372036854775808] = -9223372036854775808
                    v46 := GenS4{F0: v2, F1: 194}
                    _ = v46
                }
                _ = fi31
                _ = v32
            }
            switch (-(fi30)) {
            case 623:
                for fi47 := 0; fi47 < 13; fi47++ {
                    var v48 int8 = 42
                    v49 := GenS8{F0: -159}
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            case -2147483648:
                v4 -= (((((((1e10 + 0.0) / 1.0) - v19) + 1e10) + 1e10) + -2.25) - (v19 + 1.5))
            case 9223372036854775807:
                fmt.Println(v16)
            default:
                if (2 <= 1) {
                    var v50 uint64 = 19
                    v2 = (0 < 2)
                    var v51 uint8 = 1
                    _ = v50
                    _ = v51
                }
            }
            _ = fi30
        }
    }
    v52 := []int{752, 127, -1}
    v53 := v3[558]
    v52 = append(v52, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v6
}
