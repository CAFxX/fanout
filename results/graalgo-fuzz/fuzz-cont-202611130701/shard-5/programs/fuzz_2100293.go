package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int32; F1 uint16; F2 uint32 }

func main() {
    var v1 uint16 = 104
    var v2 int32 = int32(v1)
    var v3 float64 = 1e10
    var v4 uint32 = 330
    for fi5 := 0; fi5 < 2; fi5++ {
        switch (fi5 | (fi5 | 127)) {
        case 1:
            var v6 uint32 = (v4 + v4)
            _ = v6
        case 270:
            v7 := []int{1, fi5, fi5, (fi5 << ((42) & 63))}
            _ = v7
        case -45:
            v4 = (v4 % ((2147483647) | 1))
        default:
            var v8 uint64 = 0
            _ = v8
        }
        var v9 uint = 2
        var v10 float64 = (float64((-2.25 / -2.25)) - (v3 - v3))
        _ = fi5
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := []int{602}
        if true {
            v13 := GenS1{F0: v1, F1: (v4 ^ (v4 / (((v4 >> ((65535) & 63))) | 1)))}
            if ((9223372036854775807 >= 9223372036854775807) || (-55 < -107)) {
                if (3 > 128) {
                    var v14 int = -1
                    var v15 float32 = (1.5 * 0.0)
                    var v16 int8 = -128
                    _ = v14
                    _ = v15
                    _ = v16
                }
                if ((18446744073709551615 >= 4294967295) && (858 < 18446744073709551615)) {
                    fmt.Println(v4, v1)
                    for fi17 := 0; fi17 < 5; fi17++ {
                        var v18 bool = false
                        v3 += v3
                        _ = fi17
                        _ = v18
                    }
                    var v19 uint32 = (^(v4))
                    _ = v19
                } else {
                    v20 := GenS2{F0: (v4 + 2)}
                    _ = v20
                }
            }
            _ = v13
        } else {
            switch "\n\t" {
            case "a":
                if (42 >= 127) {
                    _ = len(v12)
                    v4 = (v4 / ((uint32(v1)) | 1))
                }
            case "foo bar":
                v3 = ((float64(1.5) / 1.0) + ((v3 - ((((v3 * ((v3 + (v3 / v3)) * v3)) / -2.25) / (3.14159 - v3)) - v3)) + ((v3 + float64(0.0)) * v3)))
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                switch "\n\t" {
                case "a":
                    fmt.Println(v3)
                }
                v22 := GenS2{F0: v4}
                if len(v12) > 0 {
                    v12[0] = (fi11 & (fi11 << ((120) & 63)))
                }
                _ = fi21
                _ = v22
            }
            var v23 int = 887
            _ = v23
        }
        _ = fi11
        _ = v12
    }
    v1 = v1
    for fi24 := 0; fi24 < 15; fi24++ {
        v3 += v3
        var v25 uint32 = (v4 - (^(v4)))
        for fi26 := 0; fi26 < 4; fi26++ {
            if (true || ((0 > int16(v4)) && true)) {
                v27 := []int{(fi24 >> ((uint(v1)) & 63)), (fi24 + (fi26 % (((^(fi24))) | 1)))}
                _ = v27
            } else {
                var v28 int8 = 74
                _ = v28
            }
            v4 = (v25 << ((98) & 63))
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    var v29 int64 = -9223372036854775808
    v30 := []int{0, -2147483648}
    v29 = v29
    if (-414 < 0) {
        var v31 string = ""
        switch "a" {
        case "foo bar":
            for fi32 := 0; fi32 < 8; fi32++ {
                v33 := GenS2{F0: v4}
                _ = fi32
                _ = v33
            }
        case "\n\t":
            var v34 string = (v31 + v31)
            _ = v34
        default:
            v35 := GenS2{F0: (v4 / (((v4 & (v4 % ((v4) | 1)))) | 1))}
            _ = v35
        }
        _ = v31
    }
    v36 := GenS3{F0: (v2 % ((v2) | 1)), F1: 1}
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 uint = 1
        switch "hello" {
        case "foo bar":
            v39 := GenS1{F0: v1, F1: (v4 ^ 4294967295)}
            _ = v39
        }
        v40 := GenS4{F0: 255, F1: (-(v2))}
        _ = fi37
        _ = v38
        _ = v40
    }
    if true {
        for fi41 := 0; fi41 < 13; fi41++ {
            v42 := GenS5{F0: ""}
            _ = fi41
            _ = v42
        }
        var v43 int64 = 1
        v44 := GenS6{F0: -1}
        _ = v43
        _ = v44
    } else {
        var v45 int16 = 1000
        _ = v45
    }
    var v46 int = -9223372036854775808
    fmt.Println(v46, v2, v1)
    for fi47 := 0; fi47 < 1; fi47++ {
        v4 = (^(v4))
        _ = len(v30)
        _ = fi47
    }
    if ((v46 | v46) == 32768) {
        for fi48 := 0; fi48 < 15; fi48++ {
            v29 = -9223372036854775808
            if true {
                v30 = append(v30, fi48)
            } else {
                v49 := GenS7{F0: -128, F1: ("\n\t" + "a")}
                _ = v49
            }
            if true {
                for fi50 := 0; fi50 < 7; fi50++ {
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                    v51 := GenS1{F0: 255, F1: (v4 / (((^(v4))) | 1))}
                    _ = fi50
                    _ = v51
                }
                var v52 uint32 = (v4 | (^(v4)))
                _ = v52
            } else {
                v53 := GenS8{F0: 820, F1: v1, F2: 4294967295}
                var v54 uint32 = (v4 - v4)
                _ = v53
                _ = v54
            }
            _ = fi48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v46
}
