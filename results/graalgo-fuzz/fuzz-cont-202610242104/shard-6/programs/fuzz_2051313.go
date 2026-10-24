package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int8 }
func genFn0() uint32 {
    var v6 bool = true
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{}
        var v9 int16 = -719
        var v10 uint8 = 42
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint8 = 1
    _ = v11
    _ = v6
    return 2
}
func genFn1(p0 uint16, p1 uint8, p2 bool) uint8 {
    if p2 {
        v12 := []int{-226, 128, 32768}
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 uint64 = 42
        fmt.Println(p2, p1)
        _ = fi13
        _ = v14
    }
    v15 := []int{}
    for fi16 := 0; fi16 < 12; fi16++ {
        switch "hello" {
        case "":
            if (!(533 > 0)) {
                v17 := GenS1{F0: 1, F1: 588}
                var v18 uint = 65535
                var v19 int16 = -32768
                _ = v17
                _ = v18
                _ = v19
            } else {
                if (-32768 <= 32767) {
                    var v20 uint32 = 900
                    p2 = p2
                    _ = v20
                } else {
                    if (!(-73 >= -1)) {
                        for fi21 := 0; fi21 < 8; fi21++ {
                            var v22 int16 = -32768
                            _ = fi21
                            _ = v22
                        }
                    } else {
                        switch "foo bar" {
                        case "hello":
                            v23 := GenS1{F0: p0, F1: 1}
                            _ = v23
                        case "a":
                            switch "hello" {
                            case "\n\t":
                                var v24 int = (fi16 >> ((65535) & 63))
                                _ = v24
                            default:
                                var v25 uint8 = p1
                                _ = v25
                            }
                        default:
                            var v26 float32 = (3.14159 / 0.1)
                            _ = v26
                        }
                        v27 := GenS2{F0: -1}
                        _ = v27
                    }
                    v28 := GenS1{F0: 859, F1: 4294967295}
                    _ = v28
                }
            }
        default:
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        }
        _ = fi16
    }
    if len(v15) > 0 {
        v15[0] = 4294967295
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (p1 % ((p1) | 1))
}

func main() {
    var v1 uint = 212
    var v2 bool = (-128 < -128)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    var v29 uint16 = 696
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 11; fi31++ {
            switch fi31 {
            case -9223372036854775808:
                var v32 int8 = 0
                _ = v32
            case 2147483647:
                var v33 int16 = 42
                _ = v33
            case -2147483648:
                if ((v1 + v1) > (v1 % (((v1 - (v1 + (^(v1))))) | 1))) {
                    v29 -= v29
                }
            }
            v1 -= v1
            if v2 {
                v29 -= uint16(v1)
            }
            _ = fi31
        }
        _ = fi30
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 int16 = -507
        _ = fi34
        _ = v35
    }
    fmt.Println(v1, v1)
    for fi36 := 0; fi36 < 19; fi36++ {
        v37 := GenS1{F0: (^(v29)), F1: (v1 >> ((18446744073709551615) & 63))}
        fmt.Println(v1, fi36)
        v1 -= (v1 * v1)
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 uint32 = 0
        var v40 int16 = -820
        _ = fi38
        _ = v39
        _ = v40
    }
    v29 += 769
    var v41 int32 = -1
    var v42 int = 42
    if (0.0 == (0.0 / 1.5)) {
        if (!v2) {
            var v43 uint16 = 65535
            var v44 int = (v42 % ((-2147483648) | 1))
            v45 := []int{v42, v42, v44}
            _ = v43
            _ = v44
            _ = v45
        }
        switch (v42 * 782) {
        case -9223372036854775808:
            v46 := []int{v42, v42, v42}
            _ = v46
        case 209:
            var v47 int16 = 134
            _ = v47
        case -1:
            v2 = (true || v2)
        default:
            fmt.Println(v2, v2)
        }
    } else {
        switch (v42 & 9223372036854775807) {
        case -9223372036854775808:
            v41 -= (v41 >> ((v1) & 63))
        case -103:
            v41 = v41
        default:
            switch "foo bar" {
            case "a":
                for fi48 := 0; fi48 < 18; fi48++ {
                    switch "a" {
                    case "a":
                        v2 = v2
                    default:
                        fmt.Println(v41)
                    }
                    _ = fi48
                }
            }
        }
        v42 = (v42 & v42)
        var v49 float32 = 0.1
        _ = v49
    }
    switch (v42 % (((v42 - v42)) | 1)) {
    case -115:
        v2 = (!(true || true))
    case 127:
        var v50 uint32 = 65535
        _ = v50
    case -28:
        v51 := GenS1{F0: (^(v29)), F1: 42}
        _ = v51
    default:
        delete(v4, 2147483648)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
