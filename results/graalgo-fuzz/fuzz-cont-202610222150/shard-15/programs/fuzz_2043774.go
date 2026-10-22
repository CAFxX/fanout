package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32; F1 int64; F2 bool }
type GenS3 struct { F0 bool; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint16; F2 int64 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint64 = 0
    var v2 uint = 255
    var v3 int64 = 42
    v4 := []int{32768, 2147483647, 503}
    v5 := []string{"foo bar", "foo bar"}
    v6 := []int{128}
    v7 := GenS1{F0: (v3 ^ v3), F1: (v2 >> (((v2 | v2)) & 63))}
    if ((v2 ^ (v2 % ((241) | 1))) != v2) {
        v2 = (v2 ^ 4294967295)
    } else {
        var v8 float64 = 0.0
        var v9 uint64 = 552
        switch 9223372036854775807 {
        case 32768:
            switch 32767 {
            case 4294967295:
                fmt.Println(v8)
            case -846:
                var v10 uint = 18446744073709551615
                _ = v10
            default:
                fmt.Println(v3)
            }
        case 9223372036854775807:
            switch ("\n\t" + "hello") {
            case "foo bar":
                v1 = v1
            case "a":
                switch "foo bar" {
                case "a":
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v12 := GenS1{F0: (-(v3)), F1: v2}
                        v13 := GenS2{F0: 2147483648, F1: -1, F2: false}
                        var v14 float32 = 0.0
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                default:
                    v15 := GenS1{F0: -709, F1: v2}
                    _ = v15
                }
            default:
                v3 = (v3 + int64(v1))
            }
        default:
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int = fi16
                v3 = (v3 + v3)
                for fi18 := 0; fi18 < 3; fi18++ {
                    switch v17 {
                    case -9223372036854775808:
                        var v19 uint = v2
                        _ = v19
                    }
                    _ = fi18
                }
                _ = fi16
                _ = v17
            }
        }
        _ = v8
        _ = v9
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        if (1e10 <= float32(0.1)) {
            v1 = (v1 << ((v2) & 63))
        } else {
            v21 := GenS3{F0: (42 > -128), F1: (v1 * 4294967295), F2: -2147483648}
            _ = v21
        }
        v2 += (v2 / ((v2) | 1))
        if ((^(v1)) == v1) {
            v22 := GenS1{F0: (v3 * (v3 & v3)), F1: (v2 * v2)}
            if false {
                switch "" {
                case "foo bar":
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 int = fi23
                        _ = fi23
                        _ = v24
                    }
                case "":
                    for fi25 := 0; fi25 < 6; fi25++ {
                        v26 := GenS4{F0: v2}
                        fmt.Println(v2, v1)
                        v3 -= v3
                        v2 = v2
                        _ = fi25
                        _ = v26
                    }
                case "hello":
                    var v27 int64 = (v3 | (v3 + 1))
                    _ = v27
                default:
                    v3 = (v3 | 0)
                }
            }
            _ = v22
        } else {
            switch (("hello" + "hello") + ("" + (("a" + ("foo bar" + "foo bar")) + ("foo bar" + "")))) {
            case "hello":
                if false {
                    for fi28 := 0; fi28 < 5; fi28++ {
                        if true {
                            var v29 string = ""
                            var v30 int32 = 2147483647
                            _ = v29
                            _ = v30
                        }
                        _ = fi28
                    }
                } else {
                    fmt.Println(v2, fi20)
                    var v31 uint16 = 251
                    _ = v31
                }
            default:
                var v32 uint32 = 1
                _ = v32
            }
            var v33 bool = (v3 == 1)
            _ = v33
        }
        _ = fi20
    }
    v1 = v1
    v3 -= (v3 * 9223372036854775807)
    v34 := GenS5{F0: 42, F1: 32767, F2: 489}
    v3 += 9223372036854775807
    v35 := GenS6{F0: -2147483648, F1: 42}
    if false {
        v1 = 1
        var v36 uint64 = (v1 + (v1 & 18446744073709551615))
        v37 := GenS6{F0: 0, F1: -1}
        _ = v36
        _ = v37
    }
    v2 = 4294967295
    v1 = v1
    if ((v3 ^ v3) != v3) {
        switch "" {
        case "hello":
            v4 = append(v4, 768)
        }
        var v38 uint8 = 127
        var v39 uint16 = 32767
        _ = v38
        _ = v39
    } else {
        v1 = v1
        for fi40 := 0; fi40 < 18; fi40++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v1 += 427
            v1 += v1
            _ = fi40
        }
        var v41 uint16 = 0
        _ = v41
    }
    v3 = v3
    var v42 uint = 21
    v42 += v2
    if ((v1 < v1) || false) {
        _ = len(v4)
        _ = len(v6)
    } else {
        var v43 int16 = -1
        fmt.Println(v43, v42)
        if false {
            for fi44 := 0; fi44 < 17; fi44++ {
                v45 := GenS7{F0: (-(v43)), F1: 1}
                _ = fi44
                _ = v45
            }
            var v46 bool = false
            v47 := GenS4{F0: (v42 % ((uint(v1)) | 1))}
            _ = v46
            _ = v47
        }
        _ = v43
    }
    var v48 int32 = -65
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
