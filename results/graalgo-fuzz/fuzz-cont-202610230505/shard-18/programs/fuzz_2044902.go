package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int32 }

func main() {
    v1 := []uint32{65535, 4294967295}
    var v2 string = (((("foo bar" + "\n\t") + "") + "foo bar") + ("\n\t" + ("\n\t" + "foo bar")))
    var v3 uint = 2
    var v4 uint = 18446744073709551615
    for fi5 := 0; fi5 < 2; fi5++ {
        v4 = v4
        _ = fi5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        switch "hello" {
        case "hello":
            v2 = "foo bar"
        }
        _ = fi6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        if (0 != 1) {
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
        }
        for fi8 := 0; fi8 < 20; fi8++ {
            if true {
                v9 := GenS1{F0: 0}
                _ = v9
            } else {
                v3 += (v4 | v4)
                v4 = 806
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (1 > -1) {
                    var v10 uint16 = 0
                    fmt.Println(v3, v10)
                    _ = v10
                }
                if false {
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                } else {
                    fmt.Println(v4, fi7)
                    if (-759 < -9223372036854775808) {
                        v2 = ("hello" + "hello")
                        var v11 int64 = -1
                        var v12 uint32 = 42
                        _ = v11
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 840}
                        v14 := GenS1{F0: 635}
                        fmt.Println(v2, fi7, fi8)
                        _ = v13
                        _ = v14
                    }
                    if ((-9223372036854775808 >= -9223372036854775808) && (0 != 255)) {
                        if false {
                            v15 := GenS2{F0: v2, F1: 40}
                            v16 := GenS2{F0: ("hello" + v2), F1: 255}
                            var v17 uint = 18446744073709551615
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        v2 = v2
                    } else {
                        switch "" {
                        case "\n\t":
                            v2 = (v2 + v2)
                            var v18 int32 = -1
                            var v19 float32 = 1.5
                            _ = v18
                            _ = v19
                        default:
                            if (1 < 1) {
                                var v20 bool = (v2 == "a")
                                var v21 int16 = 32767
                                _ = v20
                                _ = v21
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v4 -= (^(v4))
                        var v22 uint64 = 16
                        _ = v22
                    }
                }
            }
            _ = fi8
        }
        _ = fi7
    }
    if (!(false || (-936 >= 0))) {
        v3 += v3
    } else {
        v23 := GenS1{F0: 119}
        for fi24 := 0; fi24 < 12; fi24++ {
            v1 = append(v1, 2)
            var v25 int = -966
            v4 = v3
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    var v26 uint64 = 1
    if false {
        v2 = "a"
    } else {
        if ((0 != -680) && (!(1.5 > 3.14159))) {
            v26 = 1
        }
        if (!true) {
            fmt.Println(v4, v26)
            v1 = append(v1, 42)
        }
    }
    v3 = v3
    v27 := GenS2{F0: v2, F1: 1}
    v28 := GenS3{F0: v2}
    if len(v1) > 0 {
        v1[0] = uint32(v4)
    }
    if true {
        fmt.Println(v4, v26)
    } else {
        switch v2 {
        case "a":
            fmt.Println(v2)
        case "":
            v29 := GenS3{F0: v2}
            _ = v29
        default:
            var v30 int32 = 1
            _ = v30
        }
        v31 := GenS4{F0: v26, F1: 0}
        _ = v31
    }
    v3 += v4
    v26 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
