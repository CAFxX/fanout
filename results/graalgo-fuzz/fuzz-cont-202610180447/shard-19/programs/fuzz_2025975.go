package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 string; F1 int; F2 bool }
type GenS7 struct { F0 int; F1 int64; F2 int32 }

func main() {
    var v1 bool = true
    var v2 uint16 = 32767
    v3 := []uint32{42, 2147483648}
    v4 := []string{("a" + "\n\t")}
    fmt.Println(v1, v1)
    if (((!(18446744073709551615 >= 2)) || (0 <= 545)) && (v2 != 42)) {
        v3 = append(v3, 2147483648)
    } else {
        switch ("a" + ((("a" + "a") + ("hello" + "a")) + (("foo bar" + "foo bar") + ("\n\t" + "foo bar")))) {
        case "a":
            v1 = v1
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            v5 := GenS1{F0: 2}
            _ = v5
        default:
            var v6 uint8 = 34
            _ = v6
        }
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    v2 -= v2
    for fi7 := 0; fi7 < 3; fi7++ {
        if ((v1 && ((1.5 - (1e10 / 1e10)) != ((0.1 * (-2.25 - 0.0)) - 0.1))) && (42 > 4294967295)) {
            var v8 int64 = -1
            v8 = (v8 >> ((468) & 63))
            v1 = (v1 && v1)
            _ = v8
        }
        v2 -= 0
        switch "a" {
        case "hello":
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: 2}
                var v11 float64 = 3.14159
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = fi7
    }
    switch 127 {
    case 9223372036854775807:
        if ((v2 % ((v2) | 1)) >= v2) {
            var v12 uint32 = 2147483648
            if ("a" != "foo bar") {
                v1 = true
                v2 -= 761
                for fi13 := 0; fi13 < 4; fi13++ {
                    switch "a" {
                    case "hello":
                        var v14 int8 = -128
                        _ = v14
                    default:
                        fmt.Println(v12, v1, v1)
                    }
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 int16 = -32768
                        switch fi13 {
                        case -1:
                            fmt.Println(v1)
                        case 127:
                            var v17 bool = (v16 <= (v16 ^ (^(v16))))
                            var v18 int = fi13
                            fmt.Println(fi15, v17, fi13)
                            _ = v17
                            _ = v18
                        default:
                            v16 = (^(v16))
                        }
                        var v19 float64 = 0.1
                        _ = fi15
                        _ = v16
                        _ = v19
                    }
                    if (1000 > -555) {
                        var v20 bool = false
                        _ = v20
                    }
                    _ = fi13
                }
            } else {
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS1{F0: 42}
                    v23 := GenS1{F0: 235}
                    v24 := GenS1{F0: 233}
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
            _ = v12
        } else {
            if len(v4) > 0 {
                v4[0] = "a"
            }
        }
    case -963:
        v2 = 2
    }
    if v1 {
        if (v1 || (0.1 > 3.14159)) {
            var v25 bool = true
            v26 := GenS2{F0: 0, F1: 77, F2: 18446744073709551615}
            v27 := GenS3{F0: -9223372036854775808}
            _ = v25
            _ = v26
            _ = v27
        } else {
            v28 := GenS4{F0: 128}
            var v29 int32 = 42
            for fi30 := 0; fi30 < 12; fi30++ {
                v31 := GenS2{F0: (fi30 / (((fi30 * fi30)) | 1)), F1: 18446744073709551615, F2: 0}
                v1 = false
                _ = fi30
                _ = v31
            }
            _ = v28
            _ = v29
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        fmt.Println(v2)
        for fi32 := 0; fi32 < 19; fi32++ {
            for fi33 := 0; fi33 < 7; fi33++ {
                var v34 uint = 42
                _ = len(v4)
                fmt.Println(fi32)
                _ = fi33
                _ = v34
            }
            switch (fi32 & 4294967295) {
            case 1:
                if (!true) {
                    for fi35 := 0; fi35 < 7; fi35++ {
                        v2 = v2
                        v36 := GenS5{F0: (v2 | v2), F1: v2}
                        _ = fi35
                        _ = v36
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                }
            case -1:
                v37 := GenS2{F0: fi32, F1: 18446744073709551615, F2: 65535}
                _ = v37
            case -9223372036854775808:
                var v38 float32 = 0.0
                _ = v38
            }
            var v39 int16 = 1
            _ = fi32
            _ = v39
        }
        v1 = true
    }
    var v40 bool = v1
    v41 := GenS6{F0: "foo bar", F1: 128, F2: true}
    v42 := GenS7{F0: 32768, F1: -9223372036854775808, F2: -2147483648}
    if (908 > 42) {
        fmt.Println(v2, v1, v2)
        v1 = (100 > 63)
    }
    switch -675 {
    case -704:
        v43 := GenS5{F0: 65535, F1: (v2 << ((744) & 63))}
        _ = v43
    default:
        v40 = (!(!(!v40)))
    }
    var v44 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v44
}
