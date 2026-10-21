package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint32 }
func genFn0(p0 int64) uint32 {
    var v7 uint32 = 528
    v7 = 65535
    var v8 uint32 = (v7 ^ v7)
    v8 -= (v7 / ((v8) | 1))
    if (-2147483648 == 269) {
        v9 := GenS1{F0: 153}
        switch 128 {
        case 128:
            switch ("\n\t" + (("" + "") + "hello")) {
            case "":
                if (false && (406 >= 32767)) {
                    v10 := []int{220, 0}
                    _ = v10
                } else {
                    v11 := []int{333, -2147483648, -756, 434}
                    _ = v11
                }
            case "\n\t":
                v12 := GenS1{F0: 820}
                _ = v12
            default:
                v13 := GenS1{F0: 321}
                _ = v13
            }
        default:
            switch 4294967295 {
            case 32767:
                for fi14 := 0; fi14 < 19; fi14++ {
                    v15 := GenS2{F0: false, F1: 851, F2: v7}
                    fmt.Println(v8)
                    v16 := GenS2{F0: true, F1: 786, F2: v7}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case 562:
                switch "hello" {
                case "a":
                    if (!((191 != -4) || false)) {
                        var v17 bool = true
                        v18 := []int{415, -387, -9223372036854775808}
                        var v19 bool = v17
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                case "":
                    p0 = p0
                case "foo bar":
                    for fi20 := 0; fi20 < 19; fi20++ {
                        v21 := []int{fi20, (fi20 + 32768), fi20}
                        _ = fi20
                        _ = v21
                    }
                }
            case 0:
                var v22 uint64 = 653
                _ = v22
            default:
                p0 += p0
            }
        }
        _ = v9
    }
    _ = p0
    _ = v7
    _ = v8
    return (v8 & v8)
}

func main() {
    v1 := []uint32{}
    var v2 string = "a"
    v3 := make(map[string]uint32)
    var v4 int = 128
    var v5 int32 = -1
    var v6 float32 = (0.0 / float32(float64(1e10)))
    v23 := GenS2{F0: false, F1: 4294967295, F2: 671}
    var v24 int8 = 100
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 int64 = int64(v5)
        var v27 uint32 = 0
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint16 = 322
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 int16 = -1
        _ = fi29
        _ = v30
    }
    v31 := v3["foo bar"]
    switch "" {
    case "hello":
        v32 := GenS1{F0: 2}
        _ = v32
    case "\n\t":
        var v33 uint64 = 1
        _ = v33
    case "a":
        if ((v5 + v5) >= (^(v5))) {
            var v34 bool = true
            _ = v34
        }
    default:
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            switch "\n\t" {
            case "\n\t":
                v24 = v24
            case "":
                for fi35 := 0; fi35 < 18; fi35++ {
                    for fi36 := 0; fi36 < 1; fi36++ {
                        v31 += 2147483648
                        _ = fi36
                    }
                    _ = fi35
                }
            }
            v37 := GenS1{F0: (^(v28))}
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 bool = false
                _ = fi38
                _ = v39
            }
            _ = v37
        }
    }
    v40 := GenS2{F0: false, F1: 18446744073709551615, F2: (v31 << ((255) & 63))}
    v31 -= (^(v31))
    var v41 uint8 = 127
    fmt.Println(v41, v31, v28)
    var v42 uint = 974
    for fi43 := 0; fi43 < 6; fi43++ {
        if true {
            switch v2 {
            case "\n\t":
                v28 = v28
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "a":
                _ = len(v1)
            }
            for fi44 := 0; fi44 < 19; fi44++ {
                v24 = -128
                if (1e10 >= 0.0) {
                    var v45 float64 = 0.0
                    switch 0 {
                    case 32768:
                        switch v4 {
                        case 128:
                            v24 = 1
                        }
                    case -1:
                        var v46 int = (fi43 * (-(fi44)))
                        _ = v46
                    case -9223372036854775808:
                        var v47 uint16 = v28
                        _ = v47
                    default:
                        v5 = v5
                    }
                    v48 := GenS1{F0: v28}
                    v49 := GenS1{F0: (v28 - v28)}
                    _ = v45
                    _ = v48
                    _ = v49
                }
                _ = fi44
            }
        }
        _ = fi43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
