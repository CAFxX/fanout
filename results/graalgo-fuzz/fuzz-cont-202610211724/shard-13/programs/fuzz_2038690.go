package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint32 }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]uint32)
    var v3 uint8 = 1
    var v4 uint16 = 88
    var v5 int64 = v1
    var v6 int64 = v1
    switch "foo bar" {
    case "a":
        switch 4294967295 {
        case 0:
            v6 -= v1
        case 4294967295:
            v1 = (v1 / (((v6 * (v5 + v5))) | 1))
        }
    case "":
        v4 = (v4 / ((v4) | 1))
    case "\n\t":
        v7 := v2[42]
        _ = v7
    }
    var v8 int32 = 539
    v9 := GenS1{F0: v3, F1: 127, F2: -1}
    if true {
        var v10 uint = 2
        v11 := []int{128, 4294967295, -592, -2147483648}
        switch ("hello" + "foo bar") {
        case "":
            var v12 int32 = v8
            _ = v12
        case "foo bar":
            v13 := GenS1{F0: 179, F1: 1, F2: 1000}
            _ = v13
        case "a":
            var v14 uint16 = (^(v4))
            _ = v14
        }
        _ = v10
        _ = v11
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 bool = false
        var v17 string = ("foo bar" + "a")
        var v18 float32 = -2.25
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: v3, F1: 372, F2: 1000}
    var v20 uint8 = (v3 / (((v3 << ((267) & 63))) | 1))
    var v21 uint32 = 0
    fmt.Println(v8, v4)
    switch 127 {
    case 0:
        v20 += v20
    default:
        v2[-9223372036854775808] = (v21 >> ((464) & 63))
    }
    if (((!((0.1 + 3.14159) < -2.25)) && (-897 >= int16(v5))) || (true || true)) {
        fmt.Println(v3, v3)
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 int = 0
            _ = fi22
            _ = v23
        }
        switch "" {
        case "hello":
            if (-1 >= 368) {
                for fi24 := 0; fi24 < 4; fi24++ {
                    if false {
                        fmt.Println(v20, fi24, v1)
                        if false {
                            for fi25 := 0; fi25 < 13; fi25++ {
                                v4 = (^(v4))
                                if (!false) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v26 int32 = 386
                                        _ = v26
                                    case "":
                                        v1 = -1
                                    default:
                                        v27 := GenS1{F0: v3, F1: fi24, F2: 1000}
                                        _ = v27
                                    }
                                    v28 := []int{(fi25 - (fi24 - fi25)), 9223372036854775807, (fi25 - 1), 127}
                                    _ = v28
                                }
                                _ = fi25
                            }
                        }
                    } else {
                        var v29 string = "\n\t"
                        var v30 string = "hello"
                        _ = v29
                        _ = v30
                    }
                    fmt.Println(v20, v20)
                    _ = fi24
                }
                var v31 float32 = float32(1e10)
                _ = v31
            } else {
                v2[1] = uint32(v3)
                v32 := GenS2{F0: 498, F1: v21, F2: (v21 | 42)}
                if (true && false) {
                    if true {
                        switch 127 {
                        case 4294967295:
                            v5 += v5
                        default:
                            var v33 uint = 42
                            _ = v33
                        }
                    }
                } else {
                    var v34 int16 = 202
                    switch ("a" + "a") {
                    case "foo bar":
                        var v35 uint32 = (v21 ^ 4294967295)
                        _ = v35
                    }
                    _ = v34
                }
                _ = v32
            }
        case "\n\t":
            for fi36 := 0; fi36 < 16; fi36++ {
                for fi37 := 0; fi37 < 11; fi37++ {
                    switch "foo bar" {
                    case "":
                        if false {
                            if false {
                                switch fi37 {
                                case -524:
                                    if (-1 > -883) {
                                        v38 := []int{(-(fi36)), (fi37 * (fi36 | (fi36 ^ (fi37 << ((1) & 63)))))}
                                        var v39 uint32 = v21
                                        _ = v38
                                        _ = v39
                                    } else {
                                        var v40 float64 = (0.1 + (((((3.14159 / (((0.0 + 1e10) - (((0.0 * 3.14159) / float64(1.5)) + (0.1 / ((-2.25 - float64(float32(-2.25))) - 1e10)))) / 1.5)) * float64(1.5)) * 3.14159) * -2.25) * 1.5))
                                        v8 -= (v8 - v8)
                                        _ = v40
                                    }
                                case 32768:
                                    var v41 float64 = 1e10
                                    _ = v41
                                case 960:
                                    v3 = (v20 << ((2) & 63))
                                default:
                                    switch ("foo bar" + "\n\t") {
                                    case "foo bar":
                                        var v42 float32 = (1.5 * 0.0)
                                        _ = v42
                                    }
                                }
                                v43 := v2[fi36]
                                _ = v43
                            } else {
                                v44 := []int{1}
                                _ = v44
                            }
                            v4 = (v4 % ((65535) | 1))
                            fmt.Println(v1, v20)
                        }
                    }
                    v8 += int32(v21)
                    _ = fi37
                }
                v8 += v8
                _ = fi36
            }
        }
    } else {
        if true {
            v45 := GenS2{F0: -918, F1: (v21 - 2147483648), F2: v21}
            fmt.Println(v1)
            _ = v45
        }
    }
    var v46 int64 = v1
    if ((v8 % (((v8 * v8)) | 1)) >= v8) {
        fmt.Println(v4, v8, v6)
        var v47 int16 = -530
        _ = v47
    } else {
        for fi48 := 0; fi48 < 8; fi48++ {
            var v49 float32 = 0.0
            fmt.Println(v1)
            v8 = (v8 << ((156) & 63))
            _ = fi48
            _ = v49
        }
        var v50 int16 = 1000
        _ = len(v2)
        _ = v50
    }
    switch -9223372036854775808 {
    case -327:
        v21 = (^(v21))
    case 128:
        v51 := GenS2{F0: -376, F1: v21, F2: v21}
        _ = v51
    }
    for fi52 := 0; fi52 < 7; fi52++ {
        v46 = 1
        for fi53 := 0; fi53 < 11; fi53++ {
            var v54 int64 = (v5 % ((v1) | 1))
            if true {
                if false {
                    if true {
                        switch (fi52 << ((378) & 63)) {
                        case -115:
                            fmt.Println(fi53)
                        default:
                            var v55 bool = ((42 < -5) || (!true))
                            _ = v55
                        }
                    }
                    var v56 float32 = 1e10
                    v54 = (v54 << ((1) & 63))
                    _ = v56
                }
                switch fi53 {
                case -9223372036854775808:
                    var v57 int8 = -128
                    _ = v57
                case -722:
                    v1 = v46
                case 807:
                    var v58 string = ("\n\t" + "")
                    _ = v58
                }
            }
            var v59 uint64 = 42
            _ = fi53
            _ = v54
            _ = v59
        }
        var v60 float32 = 0.1
        _ = fi52
        _ = v60
    }
    v1 += v5
    v2[1] = 65535
    var v61 uint = uint(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v61
    _ = v8
    _ = v9
}
