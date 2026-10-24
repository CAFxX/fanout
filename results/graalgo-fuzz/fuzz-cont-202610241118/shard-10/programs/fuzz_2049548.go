package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 int8; F1 uint; F2 string }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int8; F1 uint16; F2 int64 }
type GenS5 struct { F0 int16; F1 bool; F2 int16 }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 18446744073709551615
    var v3 uint = v2
    var v4 int = 9223372036854775807
    v5 := make(map[string]int)
    var v6 int64 = -9223372036854775808
    v4 -= -9223372036854775808
    var v7 bool = (!((v6 / ((1) | 1)) == v6))
    if v7 {
        var v8 float64 = 1e10
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                switch v4 {
                case 128:
                    switch (fi9 >> (((v3 >> ((v2) & 63))) & 63)) {
                    case 0:
                        fmt.Println(v8, fi9, v8)
                    default:
                        v1[2147483647] = (v4 / (((fi10 | (fi9 << (((v2 % (((v3 - 42)) | 1))) & 63)))) | 1))
                    }
                case 32767:
                    if (42 < 2) {
                        if v7 {
                            v11 := GenS1{F0: 1, F1: 271, F2: v7}
                            v12 := GenS1{F0: 255, F1: 42, F2: (v3 <= v2)}
                            _ = v11
                            _ = v12
                        } else {
                            v6 = 449
                            if false {
                                var v13 string = ("hello" + "hello")
                                if (v13 >= v13) {
                                    fmt.Println(v2, v2)
                                    switch v13 {
                                    case "":
                                        v14 := GenS1{F0: 1, F1: 0, F2: v7}
                                        _ = v14
                                    default:
                                        v15 := GenS2{F0: -1, F1: (v3 << (((v2 % (((v3 % (((v3 ^ (v2 + (v2 | (^(v2)))))) | 1))) | 1))) & 63)), F2: "a"}
                                        _ = v15
                                    }
                                }
                                _ = v13
                            } else {
                                v8 -= v8
                            }
                        }
                        switch (v4 * fi10) {
                        case 9223372036854775807:
                            v16 := GenS3{F0: 2, F1: 1}
                            _ = v16
                        case -890:
                            var v17 uint8 = 42
                            _ = v17
                        case 381:
                            var v18 int32 = 1
                            v19 := GenS2{F0: 28, F1: v2, F2: "a"}
                            _ = len(v5)
                            _ = v18
                            _ = v19
                        default:
                            if v7 {
                                v1[4294967295] = fi10
                                v4 = 104
                                var v20 uint16 = 2
                                if (!v7) {
                                    v21 := []int{127, fi10}
                                    _ = v21
                                } else {
                                    var v22 int16 = 1000
                                    _ = v22
                                }
                                v23 := []int{v4, fi9}
                                _ = v20
                                _ = v23
                            }
                        }
                    }
                case 1:
                    if v7 {
                        v24 := []int{398, v4, (-(fi10))}
                        var v25 uint64 = 2
                        _ = v24
                        _ = v25
                    }
                }
                _ = fi10
            }
            v26 := []int{v4, v4, (^(fi9))}
            v27 := GenS2{F0: 100, F1: (^(v3)), F2: "hello"}
            _ = fi9
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            v29 := []int{fi28, (v4 << ((0) & 63))}
            _ = fi28
            _ = v29
        }
        _ = v8
    }
    v30 := v1[0]
    for fi31 := 0; fi31 < 10; fi31++ {
        fmt.Println(v30)
        var v32 string = ((("" + (("\n\t" + "foo bar") + "hello")) + "hello") + "a")
        _ = fi31
        _ = v32
    }
    if false {
        for fi33 := 0; fi33 < 4; fi33++ {
            switch v30 {
            case -362:
                var v34 int16 = 0
                _ = v34
            case 525:
                fmt.Println(v30, v6)
            case 204:
                v35 := GenS4{F0: 1, F1: uint16(v2), F2: (v6 << (((v3 << ((v2) & 63))) & 63))}
                _ = v35
            default:
                var v36 int = 443
                _ = v36
            }
            v3 -= 2
            _ = fi33
        }
        fmt.Println(v30, v7, v4)
        if v7 {
            if (0 <= 1) {
                var v37 string = "foo bar"
                _ = v37
            }
        }
    } else {
        var v38 uint64 = 940
        var v39 bool = v7
        v40 := v1[42]
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println(v7, v30, v6)
    v41 := GenS2{F0: 127, F1: (^(v3)), F2: "\n\t"}
    if v7 {
        var v42 uint16 = 65535
        _ = v42
    }
    var v43 int64 = (v6 >> ((697) & 63))
    if v7 {
        v30 -= (v4 % (((v4 + 1)) | 1))
        switch "foo bar" {
        case "foo bar":
            v44 := GenS1{F0: 32767, F1: 255, F2: (!(false && ((float32(1e10) - float32(0.1)) >= 3.14159)))}
            _ = v44
        case "a":
            v45 := GenS5{F0: 0, F1: v7, F2: int16(v4)}
            _ = v45
        }
    } else {
        v6 += (v43 ^ (v6 >> (((v3 / ((v2) | 1))) & 63)))
        v46 := GenS6{F0: 273}
        _ = v46
    }
    for fi47 := 0; fi47 < 17; fi47++ {
        for fi48 := 0; fi48 < 13; fi48++ {
            v49 := GenS4{F0: 109, F1: 2, F2: 374}
            v3 -= v3
            if (v7 || (-2.25 <= 0.0)) {
                v50 := GenS6{F0: int16(v6)}
                if v7 {
                    v51 := GenS5{F0: -677, F1: false, F2: 0}
                    var v52 float64 = ((float64((1.5 * 0.0)) + ((-2.25 * 0.0) / 1e10)) + (3.14159 * 0.1))
                    _ = v51
                    _ = v52
                } else {
                    _ = len(v1)
                }
                v43 += (v43 * 895)
                _ = v50
            } else {
                fmt.Println(v2, fi48, v30)
            }
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    var v53 int64 = 0
    var v54 uint8 = 2
    var v55 int = v30
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v43
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v6
    _ = v7
}
