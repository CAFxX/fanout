package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string }
func genFn0() int64 {
    if (65535 <= 912) {
        var v6 bool = true
        v7 := GenS1{F0: 0}
        var v8 int64 = -9223372036854775808
        _ = v6
        _ = v7
        _ = v8
    } else {
        switch 4294967295 {
        case 32768:
            if (3.14159 <= 0.1) {
                if (("" >= "a") && (200 == 0)) {
                    var v9 uint8 = 2
                    if ((!false) || true) {
                        switch "a" {
                        case "\n\t":
                            for fi10 := 0; fi10 < 6; fi10++ {
                                switch 0 {
                                case 1:
                                    if (false && ((1 != -259) || (!true))) {
                                        v11 := []int{}
                                        _ = v11
                                    } else {
                                        v12 := GenS1{F0: -1}
                                        v13 := GenS1{F0: -128}
                                        _ = v12
                                        _ = v13
                                    }
                                case -426:
                                    switch (fi10 % (((fi10 << ((941) & 63))) | 1)) {
                                    case -9223372036854775808:
                                        v14 := GenS1{F0: -83}
                                        _ = v14
                                    case 32767:
                                        if true {
                                            for fi15 := 0; fi15 < 14; fi15++ {
                                                v16 := GenS1{F0: 1}
                                                if true {
                                                    v17 := []int{(fi10 % ((fi10) | 1)), (fi10 - 127)}
                                                    var v18 uint = 652
                                                    _ = v17
                                                    _ = v18
                                                } else {
                                                    v19 := []int{(fi15 << ((1) & 63)), fi10}
                                                    fmt.Println(v9)
                                                    _ = v19
                                                }
                                                if (!(!((false && (3.14159 == (1.5 / 0.1))) || true))) {
                                                    v9 = v9
                                                } else {
                                                    v9 -= v9
                                                    var v20 uint16 = 65535
                                                    _ = v20
                                                }
                                                _ = fi15
                                                _ = v16
                                            }
                                            if false {
                                                fmt.Println(v9, fi10)
                                                if true {
                                                    switch fi10 {
                                                    case 117:
                                                        var v21 int16 = 1
                                                        _ = v21
                                                    case 32768:
                                                        fmt.Println(fi10, v9)
                                                    case 42:
                                                        var v22 int32 = 2147483647
                                                        _ = v22
                                                    default:
                                                        v9 -= (v9 - 0)
                                                    }
                                                }
                                                v23 := []int{fi10, fi10}
                                                _ = v23
                                            } else {
                                                v24 := []int{(^(fi10)), fi10}
                                                var v25 float64 = -2.25
                                                _ = v24
                                                _ = v25
                                            }
                                            var v26 int32 = 735
                                            _ = v26
                                        }
                                    }
                                case 0:
                                    v9 = (^(v9))
                                }
                                _ = fi10
                            }
                        default:
                            v9 += (v9 % (((v9 * v9)) | 1))
                        }
                    }
                    _ = v9
                }
                var v27 int = 9223372036854775807
                v27 = v27
                _ = v27
            }
        case 0:
            v28 := []int{}
            _ = v28
        case -590:
            var v29 uint64 = 715
            _ = v29
        default:
            var v30 int16 = -32768
            _ = v30
        }
        v31 := GenS2{F0: 1}
        var v32 uint32 = 65535
        _ = v31
        _ = v32
    }
    var v33 int16 = 1000
    _ = v33
    return 42
}
func genFn1() int {
    var v34 int8 = 100
    switch 32767 {
    case 32768:
        for fi35 := 0; fi35 < 6; fi35++ {
            for fi36 := 0; fi36 < 2; fi36++ {
                switch "foo bar" {
                case "foo bar":
                    var v37 float32 = 1.5
                    v38 := []int{fi35}
                    _ = v37
                    _ = v38
                }
                var v39 int8 = (v34 % ((-128) | 1))
                _ = fi36
                _ = v39
            }
            _ = fi35
        }
    case -1:
        v40 := GenS2{F0: 255}
        _ = v40
    case -9223372036854775808:
        var v41 float64 = 3.14159
        _ = v41
    default:
        v42 := GenS3{F0: 2147483648, F1: 0, F2: 2}
        _ = v42
    }
    _ = v34
    return 1
}

func main() {
    var v1 uint64 = 1
    var v2 uint64 = v1
    var v3 int16 = -1
    v4 := make(map[int]int)
    var v5 uint = 766
    var v43 int8 = -1
    for fi44 := 0; fi44 < 5; fi44++ {
        var v45 float32 = ((0.0 * 3.14159) * float32(0.1))
        for fi46 := 0; fi46 < 13; fi46++ {
            v47 := GenS4{F0: (v2 << (((^(v5))) & 63))}
            _ = fi46
            _ = v47
        }
        for fi48 := 0; fi48 < 7; fi48++ {
            switch fi48 {
            case 128:
                var v49 float32 = (float32(1.5) + ((1.5 * v45) + (v45 - v45)))
                _ = v49
            }
            var v50 uint8 = 118
            v51 := GenS5{F0: ("" + "hello")}
            _ = fi48
            _ = v50
            _ = v51
        }
        _ = fi44
        _ = v45
    }
    fmt.Println(v5, v5)
    v5 = v5
    v43 = (v43 % ((-73) | 1))
    v3 = v3
    if (true || false) {
        switch ("" + "") {
        case "foo bar":
            var v52 float32 = 3.14159
            _ = v52
        default:
            v1 -= uint64(v5)
        }
        v53 := GenS1{F0: -18}
        v2 = v2
        _ = v53
    }
    for fi54 := 0; fi54 < 1; fi54++ {
        if false {
            for fi55 := 0; fi55 < 13; fi55++ {
                var v56 int8 = v43
                v2 = (v2 & v2)
                _ = fi55
                _ = v56
            }
        }
        v1 = (v2 + (v1 % ((v1) | 1)))
        _ = fi54
    }
    v3 = v3
    var v57 float32 = ((1.5 + (((0.0 + (-2.25 - (-2.25 * 0.0))) - -2.25) * ((1e10 - (1e10 + (0.0 / 3.14159))) * 0.1))) - ((((1e10 + (3.14159 / float32((-2.25 + 0.0)))) / (((1e10 * 0.1) + (0.0 * float32(0.1))) * 3.14159)) - (0.1 - 0.1)) + (0.1 + 1e10)))
    if (!(true && true)) {
        if ((925 >= 4294967295) && (!(0.1 != 0.1))) {
            var v58 uint64 = 1
            _ = v58
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v57
}
