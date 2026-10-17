package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int }
type GenS4 struct { F0 int32; F1 int; F2 string }
type GenS5 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS7 struct { F0 int8 }

func main() {
    v1 := []uint32{401, 4294967295, 1, 103}
    var v2 uint32 = 992
    var v3 uint16 = 2
    var v4 uint = 42
    var v5 uint16 = v3
    if true {
        switch "a" {
        case "foo bar":
            if true {
                var v6 uint = (v4 % ((v4) | 1))
                _ = v6
            } else {
                switch -9223372036854775808 {
                case 32768:
                    switch "a" {
                    case "a":
                        v7 := GenS1{F0: 0}
                        _ = v7
                    case "\n\t":
                        var v8 uint16 = 2
                        _ = v8
                    case "hello":
                        var v9 uint32 = (v2 >> ((v4) & 63))
                        _ = v9
                    default:
                        v10 := GenS1{F0: 0}
                        _ = v10
                    }
                case 127:
                    v2 += (v2 / ((397) | 1))
                case -693:
                    var v11 int8 = -96
                    _ = v11
                }
                if len(v1) > 0 {
                    v1[0] = v2
                }
            }
        case "\n\t":
            if true {
                v12 := GenS1{F0: 0}
                switch "a" {
                case "\n\t":
                    switch (("" + "a") + "\n\t") {
                    case "a":
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                    case "":
                        v2 = (^(v2))
                    }
                case "a":
                    v13 := GenS2{F0: -9223372036854775808, F1: 2147483647}
                    _ = v13
                case "":
                    fmt.Println(v4, v5)
                }
                v14 := GenS3{F0: 2147483648, F1: (v5 + 65535), F2: 352}
                _ = v12
                _ = v14
            } else {
                var v15 uint64 = 18446744073709551615
                var v16 int64 = 821
                v3 -= v5
                _ = v15
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 string = ""
            v5 = 2
            v19 := GenS2{F0: (fi17 - (fi17 % (((fi17 << ((v4) & 63))) | 1))), F1: 42}
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 8; fi20++ {
            if false {
                if (false || true) {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v2 = (^(v2))
                        switch (fi21 & fi20) {
                        case 4294967295:
                            var v22 uint16 = uint16(v4)
                            _ = v22
                        case 9223372036854775807:
                            fmt.Println(v3, v5, v3)
                        default:
                            v23 := GenS4{F0: int32(v4), F1: 0, F2: "foo bar"}
                            _ = v23
                        }
                        v4 = v4
                        _ = fi21
                    }
                    v5 += v5
                } else {
                    var v24 int = 127
                    fmt.Println(v2, fi20, v24)
                    _ = v24
                }
                switch -2147483648 {
                case 128:
                    if false {
                        v4 += (v4 + 918)
                        for fi25 := 0; fi25 < 10; fi25++ {
                            v2 -= v2
                            var v26 float32 = (3.14159 - (1.5 - 0.0))
                            _ = fi25
                            _ = v26
                        }
                        if (v5 >= (v3 + (v5 ^ (v5 >> ((0) & 63))))) {
                            v27 := GenS3{F0: 1, F1: v5, F2: (fi20 << ((v4) & 63))}
                            fmt.Println(fi20, v3, v3)
                            v28 := GenS5{F0: (v5 ^ (v5 / ((65535) | 1))), F1: -76, F2: (v2 % (((v2 * (v2 + 859))) | 1))}
                            _ = v27
                            _ = v28
                        }
                    } else {
                        var v29 string = "\n\t"
                        _ = v29
                    }
                default:
                    v2 -= uint32(v3)
                }
                v30 := GenS6{F0: (v2 >> (((v4 & (v4 + (v4 / ((502) | 1))))) & 63)), F1: v3, F2: v2}
                _ = v30
            }
            _ = fi20
        }
    }
    fmt.Println(v2, v5, v3)
    v1 = append(v1, v2)
    v5 = v3
    v31 := GenS6{F0: (v2 * v2), F1: v3, F2: 65535}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 += 0
    v5 = (v5 + (v5 >> (((^(v4))) & 63)))
    switch "foo bar" {
    case "hello":
        if false {
            v3 = 0
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v4, v2, v2)
        }
    case "foo bar":
        v32 := GenS7{F0: -128}
        _ = v32
    default:
        for fi33 := 0; fi33 < 18; fi33++ {
            fmt.Println(fi33, v5)
            _ = fi33
        }
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 int64 = 1
            v3 -= 65535
            var v37 int32 = 42
            v38 := GenS7{F0: -1}
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        var v39 string = (("\n\t" + "foo bar") + ("\n\t" + ""))
        _ = fi34
        _ = v39
    }
    v40 := GenS5{F0: uint16(v2), F1: -1, F2: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v5
}
