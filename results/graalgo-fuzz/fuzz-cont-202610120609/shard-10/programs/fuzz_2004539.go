package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint32 }
type GenS3 struct { F0 string; F1 int }

func main() {
    var v1 uint = 255
    v2 := []string{}
    v3 := []string{("" + "\n\t")}
    v4 := make(map[uint32]string)
    if (!(!false)) {
        switch 32768 {
        case 218:
            for fi5 := 0; fi5 < 1; fi5++ {
                v1 = (v1 * v1)
                _ = fi5
            }
        case -717:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -243:
            var v6 uint64 = 307
            _ = v6
        }
        switch -9223372036854775808 {
        case -1:
            switch "foo bar" {
            case "a":
                v7 := GenS1{F0: 969, F1: 100}
                _ = v7
            case "\n\t":
                v8 := GenS1{F0: 372, F1: 127}
                _ = v8
            case "":
                if false {
                    if (4294967295 <= 0) {
                        v9 := GenS2{F0: -1, F1: -90, F2: 1}
                        if (127 > -2147483648) {
                            v1 = v1
                            if false {
                                v1 = (v1 ^ 0)
                            }
                            if len(v2) > 0 {
                                v2[0] = ((("a" + "") + "foo bar") + "\n\t")
                            }
                        } else {
                            fmt.Println(v1)
                        }
                        v1 = (v1 % (((v1 | v1)) | 1))
                        _ = v9
                    }
                } else {
                    v10 := GenS2{F0: -1, F1: 1, F2: 148}
                    v1 += 862
                    if false {
                        var v11 uint64 = 85
                        var v12 int32 = 88
                        _ = v11
                        _ = v12
                    }
                    _ = v10
                }
            }
        case -9223372036854775808:
            if (1 >= 18446744073709551615) {
                v1 = v1
                var v13 int8 = -128
                _ = v13
            } else {
                v14 := GenS3{F0: "a", F1: 4294967295}
                _ = v14
            }
        case 127:
            fmt.Println(v1)
        }
        if (false && (!true)) {
            var v15 uint8 = 42
            if (-2147483648 != 9223372036854775807) {
                delete(v4, 42)
                v16 := v4[762]
                _ = v16
            } else {
                var v17 bool = true
                var v18 int16 = 32767
                for fi19 := 0; fi19 < 2; fi19++ {
                    if true {
                        v20 := GenS2{F0: -1, F1: 1, F2: 42}
                        var v21 uint32 = 0
                        _ = v20
                        _ = v21
                    } else {
                        v22 := GenS2{F0: 0, F1: 100, F2: 677}
                        fmt.Println(v18, v18, v15)
                        _ = v22
                    }
                    var v23 uint64 = 42
                    v24 := GenS2{F0: 42, F1: 100, F2: 2}
                    _ = fi19
                    _ = v23
                    _ = v24
                }
                _ = v17
                _ = v18
            }
            v1 -= (v1 ^ v1)
            _ = v15
        }
    } else {
        fmt.Println(v1)
        switch "\n\t" {
        case "a":
            v1 = v1
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "":
            v25 := GenS3{F0: "hello", F1: 4294967295}
            _ = v25
        }
    }
    var v26 int64 = -1
    if false {
        if true {
            v26 -= v26
            v27 := GenS1{F0: 4294967295, F1: -9}
            v28 := GenS2{F0: 100, F1: 100, F2: 65535}
            _ = v27
            _ = v28
        } else {
            var v29 bool = (-2.25 <= (0.1 - -2.25))
            _ = v29
        }
    } else {
        v30 := v4[65535]
        _ = v30
    }
    v1 += 255
    var v31 uint64 = 781
    switch 204 {
    case 704:
        if (-2147483648 >= 550) {
            var v32 int = 32768
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v32
        } else {
            v26 = (^(v26))
        }
    case -2147483648:
        v33 := GenS2{F0: -22, F1: 100, F2: 2147483647}
        _ = v33
    case 535:
        v1 = v1
    default:
        v31 = 980
    }
    v34 := GenS2{F0: -128, F1: -1, F2: 4294967295}
    var v35 float32 = (1e10 * float32(((((3.14159 / 3.14159) / -2.25) / 1.0) / 1.0)))
    for fi36 := 0; fi36 < 18; fi36++ {
        var v37 int64 = (v26 + (v26 << (((v1 << ((1) & 63))) & 63)))
        _ = fi36
        _ = v37
    }
    var v38 int64 = 42
    var v39 int8 = 107
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
}
