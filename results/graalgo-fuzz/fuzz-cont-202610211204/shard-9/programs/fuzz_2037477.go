package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint16; F1 int; F2 string }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int8; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint8 }
func genFn0() bool {
    var v7 float32 = 0.1
    if (18446744073709551615 >= 2) {
        v7 = (float32(-2.25) * ((v7 - v7) / (v7 + v7)))
        v8 := []int{652}
        _ = v8
    }
    v9 := GenS1{F0: 32767, F1: true}
    _ = v7
    _ = v9
    return (42 >= 42)
}

func main() {
    var v1 int8 = 0
    var v2 int64 = 1
    v3 := make(map[uint32]bool)
    v4 := []string{("\n\t" + "\n\t"), (("" + "a") + ("" + (("" + "\n\t") + "a")))}
    var v5 uint32 = 0
    v6 := []int{32767, 2147483647}
    switch "a" {
    case "foo bar":
        v10 := GenS2{F0: 42, F1: true}
        _ = v10
    case "a":
        if (!true) {
            var v11 int64 = v2
            var v12 int = 127
            fmt.Println(v11)
            _ = v11
            _ = v12
        } else {
            v2 -= (v2 << ((258) & 63))
        }
    }
    if ((-2.25 < 1e10) || (622 > 42)) {
        switch "hello" {
        case "":
            v5 = v5
        case "\n\t":
            v13 := GenS1{F0: 32767, F1: (!(!("\n\t" < "\n\t")))}
            _ = v13
        case "a":
            delete(v3, (v5 | (v5 / (((v5 + v5)) | 1))))
        }
        v14 := GenS3{F0: 1, F1: 2147483647, F2: "foo bar"}
        if false {
            for fi15 := 0; fi15 < 13; fi15++ {
                v16 := GenS4{F0: (v1 + 1), F1: v5}
                v17 := GenS4{F0: 10, F1: (v5 << ((uint(v5)) & 63))}
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            var v18 float64 = 0.0
            fmt.Println(v18)
            _ = v18
        }
        _ = v14
    } else {
        if len(v6) > 0 {
            v6[0] = 127
        }
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v20 := GenS2{F0: v5, F1: true}
        _ = fi19
        _ = v20
    }
    if true {
        switch 9223372036854775807 {
        case -1:
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS5{F0: 18446744073709551615, F1: 859}
                if len(v6) > 0 {
                    v6[0] = 42
                }
                for fi23 := 0; fi23 < 7; fi23++ {
                    if true {
                        if true {
                            v1 -= v1
                        } else {
                            var v24 float32 = 0.1
                            v2 = v2
                            switch "foo bar" {
                            case "foo bar":
                                var v25 string = "foo bar"
                                _ = v25
                            default:
                                fmt.Println(v24)
                            }
                            _ = v24
                        }
                    } else {
                        if true {
                            v5 = v5
                            var v26 bool = false
                            v1 += v1
                            _ = v26
                        }
                        var v27 uint64 = uint64(v5)
                        _ = v27
                    }
                    var v28 uint64 = 42
                    fmt.Println(fi21, v28, fi21)
                    v2 -= (^(v2))
                    if true {
                        if (v2 <= v2) {
                            var v29 float64 = 0.1
                            _ = v29
                        } else {
                            v30 := GenS2{F0: 42, F1: false}
                            _ = v30
                        }
                        var v31 int64 = v2
                        switch -1 {
                        case 128:
                            fmt.Println(v5, v28)
                        }
                        _ = v31
                    }
                    if false {
                        v32 := v3[v5]
                        var v33 float32 = 0.0
                        switch fi23 {
                        case 1:
                            v28 += (v28 | (^(v28)))
                            v2 = -1
                        }
                        _ = len(v3)
                        v1 -= (v1 % (((v1 << ((637) & 63))) | 1))
                        v34 := GenS6{F0: (v1 * (v1 % (((v1 / ((74) | 1))) | 1))), F1: 24, F2: (^(v2))}
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = fi23
                    _ = v28
                }
                _ = fi21
                _ = v22
            }
        case -248:
            v35 := GenS4{F0: 1, F1: (v5 % (((v5 % (((v5 >> ((2) & 63))) | 1))) | 1))}
            _ = v35
        case 127:
            var v36 uint64 = 0
            _ = v36
        default:
            if true {
                var v37 uint32 = 2147483648
                v38 := GenS1{F0: 685, F1: true}
                var v39 uint16 = uint16(v37)
                _ = v37
                _ = v38
                _ = v39
            }
        }
        var v40 uint8 = 42
        var v41 uint64 = 1
        _ = v40
        _ = v41
    }
    var v42 int8 = v1
    _ = len(v4)
    v42 += (-(v1))
    switch 42 {
    case 32768:
        v43 := GenS3{F0: 2, F1: -1, F2: "a"}
        _ = v43
    case -9223372036854775808:
        for fi44 := 0; fi44 < 14; fi44++ {
            fmt.Println(fi44, v2, v5)
            v45 := GenS7{F0: 1}
            _ = fi44
            _ = v45
        }
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        fmt.Println(v5, v42)
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
