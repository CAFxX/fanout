package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint32; F1 int8; F2 bool }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int16) int32 {
    fmt.Println(p0, p0)
    v7 := GenS1{F0: 127, F1: true, F2: p0}
    p1 = p1
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint32 = 257
            switch ((("\n\t" + "a") + "\n\t") + ("\n\t" + "\n\t")) {
            case "a":
                v10 = (v10 ^ v10)
            case "hello":
                var v11 uint8 = 38
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        var v12 int32 = 1
        _ = fi8
        _ = v12
    }
    switch -2147483648 {
    case -225:
        v13 := GenS2{F0: 2147483647, F1: 786}
        _ = v13
    case -204:
        var v14 float32 = (0.0 / 1.0)
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}
func genFn1(p0 string, p1 string) uint32 {
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 int = 32767
        _ = fi15
        _ = v16
    }
    if true {
        if false {
            if (704 > -32768) {
                v17 := []int{-2147483648}
                _ = v17
            } else {
                v18 := GenS3{F0: 18446744073709551615}
                v19 := []int{8, 128, 1}
                _ = v18
                _ = v19
            }
            p1 = p0
            switch p1 {
            case "hello":
                p0 = "a"
            case "\n\t":
                v20 := GenS4{F0: p0, F1: 13}
                _ = v20
            default:
                switch p1 {
                case "foo bar":
                    var v21 int16 = -1
                    _ = v21
                case "hello":
                    switch -2147483648 {
                    case 42:
                        v22 := GenS2{F0: 448, F1: 18446744073709551615}
                        _ = v22
                    case 2147483647:
                        var v23 int64 = 731
                        _ = v23
                    default:
                        p0 = ""
                    }
                }
            }
        }
        var v24 int8 = -104
        p1 = ((p0 + p0) + p1)
        _ = v24
    } else {
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return 613
}

func main() {
    var v1 int16 = -187
    var v2 uint = 2
    var v3 uint8 = 13
    var v4 int32 = 1
    var v5 float32 = (1.5 * 0.0)
    v6 := []string{"foo bar", ("foo bar" + ((("hello" + "hello") + ("" + "foo bar")) + "")), "a"}
    switch "foo bar" {
    case "\n\t":
        var v25 uint32 = 4294967295
        _ = v25
    case "foo bar":
        var v26 uint8 = v3
        _ = v26
    }
    v27 := GenS1{F0: 2147483647, F1: ((v2 / (((^(v2))) | 1)) >= (v2 % (((v2 >> ((759) & 63))) | 1))), F2: (v3 % ((116) | 1))}
    for fi28 := 0; fi28 < 3; fi28++ {
        v5 += ((0.1 / v5) - (v5 + ((float32(1e10) / v5) / 1.0)))
        var v29 int = fi28
        v30 := GenS5{F0: 2, F1: (v3 < (v3 % ((v3) | 1)))}
        _ = fi28
        _ = v29
        _ = v30
    }
    switch 9223372036854775807 {
    case 0:
        if true {
            switch ((("" + "") + "a") + "a") {
            case "foo bar":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            var v31 bool = (127 >= -41)
            if v31 {
                _ = len(v6)
                _ = len(v6)
            }
            _ = v31
        } else {
            v32 := GenS6{F0: 2147483648, F1: 1, F2: false}
            _ = v32
        }
    case 199:
        var v33 uint16 = 32767
        _ = v33
    case 32767:
        if (v3 <= v3) {
            if true {
                v34 := GenS1{F0: 128, F1: false, F2: (^(v3))}
                v6 = append(v6, "\n\t")
                _ = v34
            } else {
                if false {
                    v2 -= (v2 | (v2 << ((v2) & 63)))
                }
                for fi35 := 0; fi35 < 15; fi35++ {
                    if true {
                        v3 -= 2
                        for fi36 := 0; fi36 < 11; fi36++ {
                            var v37 int32 = 2147483647
                            var v38 int64 = -9223372036854775808
                            _ = len(v6)
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    }
                    _ = fi35
                }
                for fi39 := 0; fi39 < 7; fi39++ {
                    var v40 uint = v2
                    _ = fi39
                    _ = v40
                }
            }
            v2 = v2
            fmt.Println(v5, v1)
        }
    default:
        v4 = -1
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        v42 := GenS3{F0: 877}
        _ = fi41
        _ = v42
    }
    var v43 float32 = v5
    for fi44 := 0; fi44 < 11; fi44++ {
        for fi45 := 0; fi45 < 4; fi45++ {
            v43 = v5
            _ = fi45
        }
        v46 := GenS2{F0: 658, F1: 4294967295}
        if (!false) {
            if (898 >= 14) {
                v3 = v3
                var v47 float64 = 3.14159
                _ = v47
            } else {
                switch (^(fi44)) {
                case 1:
                    switch "\n\t" {
                    case "hello":
                        v2 = 18446744073709551615
                    case "\n\t":
                        v6 = append(v6, ("hello" + "hello"))
                    }
                default:
                    var v48 uint64 = 18446744073709551615
                    _ = v48
                }
                v49 := GenS7{F0: 0}
                fmt.Println(v2, v4, fi44)
                _ = v49
            }
            v50 := GenS5{F0: 377, F1: (!true)}
            _ = v50
        }
        _ = fi44
        _ = v46
    }
    fmt.Println(v1)
    var v51 float32 = (1.5 - ((float32(3.14159) / float32(float64(v43))) * (-2.25 / 1.0)))
    var v52 int8 = 127
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
