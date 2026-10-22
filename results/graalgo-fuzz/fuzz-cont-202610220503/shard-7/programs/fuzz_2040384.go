package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint8) int {
    switch 1 {
    case 303:
        p0 += 42
    case -9223372036854775808:
        var v5 uint = 2
        _ = v5
    }
    p0 = p0
    switch 32767 {
    case -892:
        var v6 int16 = -1
        _ = v6
    case 127:
        v7 := []int{127, 128}
        _ = v7
    }
    switch "" {
    case "hello":
        p0 = 0
    default:
        fmt.Println(p0)
    }
    _ = p0
    return -688
}

func main() {
    v1 := []uint32{805, 305, 2147483648}
    var v2 uint16 = 2
    var v3 uint64 = 4294967295
    var v4 uint = 2
    v8 := GenS1{F0: ((("foo bar" + ("\n\t" + "foo bar")) + "foo bar") + "hello"), F1: "a"}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= v4
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: "foo bar", F1: "\n\t"}
        var v11 uint8 = 127
        if false {
            if true {
                var v12 uint32 = 1
                v3 = v3
                _ = v12
            } else {
                if (1.5 < 1.5) {
                    switch fi9 {
                    case 9223372036854775807:
                        var v13 float64 = float64(3.14159)
                        _ = v13
                    case 0:
                        if (!false) {
                            switch "\n\t" {
                            case "":
                                v14 := GenS2{F0: v3}
                                _ = v14
                            }
                        }
                    }
                    if true {
                        var v15 float64 = 3.14159
                        var v16 int8 = -1
                        _ = v15
                        _ = v16
                    } else {
                        _ = len(v1)
                        v1 = append(v1, 4294967295)
                    }
                    v2 = (v2 + (^(v2)))
                }
                if true {
                    if len(v1) > 0 {
                        v1[0] = 119
                    }
                }
            }
        } else {
            v11 = 2
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint64 = (v3 - 1)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi17
        _ = v18
    }
    v19 := GenS3{F0: (v4 << (((^(v4))) & 63)), F1: 0}
    _ = len(v1)
    fmt.Println(v3)
    if (166 >= 127) {
        if true {
            var v20 uint16 = (v2 >> (((v4 / ((uint(v3)) | 1))) & 63))
            _ = v20
        } else {
            if (4294967295 != 128) {
                var v21 uint8 = uint8(v2)
                var v22 int64 = 244
                if len(v1) > 0 {
                    v1[0] = 2
                }
                _ = v21
                _ = v22
            }
        }
        v2 = (v2 + (v2 - v2))
        for fi23 := 0; fi23 < 2; fi23++ {
            v1 = append(v1, 65535)
            v4 = v4
            var v24 bool = (true || (!false))
            _ = fi23
            _ = v24
        }
    }
    var v25 int32 = 749
    if true {
        v3 += v3
        var v26 bool = true
        switch -9223372036854775808 {
        case 275:
            switch "" {
            case "foo bar":
                v27 := GenS4{F0: 100}
                _ = v27
            default:
                var v28 uint = v4
                _ = v28
            }
        case 32768:
            for fi29 := 0; fi29 < 13; fi29++ {
                if (!v26) {
                    v26 = v26
                    v30 := GenS2{F0: v3}
                    _ = v30
                } else {
                    v3 += v3
                    var v31 string = ""
                    _ = v31
                }
                switch fi29 {
                case -2147483648:
                    for fi32 := 0; fi32 < 5; fi32++ {
                        var v33 string = (("\n\t" + "") + "hello")
                        _ = fi32
                        _ = v33
                    }
                case -1:
                    if len(v1) > 0 {
                        v1[0] = uint32(v4)
                    }
                }
                _ = fi29
            }
        case 127:
            _ = len(v1)
        default:
            v34 := GenS4{F0: 127}
            _ = v34
        }
        _ = v26
    }
    v25 = v25
    v1 = append(v1, 0)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = 1
    fmt.Println(v25, v25)
    v3 = (v3 / ((18446744073709551615) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v8
}
