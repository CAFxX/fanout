package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 uint32; F2 int8 }

func main() {
    var v1 int16 = 1
    var v2 int64 = -1
    v3 := make(map[string]uint32)
    var v4 int8 = 100
    v5 := []string{("a" + ("" + ("" + "foo bar"))), "a"}
    var v6 int64 = (v2 * (v2 + v2))
    if (!false) {
        v1 -= (v1 ^ v1)
        for fi7 := 0; fi7 < 12; fi7++ {
            v6 = 281
            _ = fi7
        }
        var v8 int8 = 42
        _ = v8
    }
    v1 += (v1 ^ v1)
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint = 18446744073709551615
        if (2 > 18446744073709551615) {
            var v11 int8 = (v4 << (((v10 | v10)) & 63))
            v2 = v6
            _ = v11
        } else {
            var v12 float32 = -2.25
            v13 := GenS1{F0: 32}
            v14 := GenS2{F0: ""}
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    var v15 int64 = v6
    v6 = (v2 + v2)
    switch 939 {
    case -346:
        v6 -= (v2 | v15)
    case 2147483647:
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(fi16, fi16)
            if true {
                var v17 uint32 = 2
                _ = v17
            } else {
                v2 -= 9223372036854775807
                fmt.Println(v2, v15, fi16)
            }
            if (127 != 128) {
                var v18 string = "a"
                _ = v18
            }
            _ = fi16
        }
    case 127:
        var v19 float64 = 0.1
        _ = v19
    default:
        v6 = (^(v6))
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        if (!true) {
            v21 := GenS1{F0: 255}
            for fi22 := 0; fi22 < 5; fi22++ {
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
                var v23 int16 = (v1 | v1)
                if false {
                    v24 := GenS2{F0: (("hello" + "a") + "foo bar")}
                    if (false && true) {
                        if true {
                            v4 += 0
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if (1 == 675) {
                                v5 = append(v5, "a")
                            }
                            var v25 bool = true
                            var v26 uint32 = 1
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS1{F0: 97}
                        v5 = append(v5, "hello")
                        _ = v27
                    }
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            var v28 int8 = v4
            _ = v21
            _ = v28
        }
        _ = fi20
    }
    fmt.Println(v15, v1, v2)
    v6 -= -9223372036854775808
    v29 := GenS3{F0: (v2 % (((-(v15))) | 1))}
    fmt.Println(v15)
    var v30 uint64 = 796
    v3["\n\t"] = 688
    var v31 uint16 = 35
    _ = len(v3)
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 string = "\n\t"
        v34 := GenS4{F0: 256}
        v15 = -1
        _ = fi32
        _ = v33
        _ = v34
    }
    switch "a" {
    case "\n\t":
        v15 += (v2 & v15)
    case "hello":
        var v35 uint8 = 128
        _ = v35
    case "":
        for fi36 := 0; fi36 < 14; fi36++ {
            if (4294967295 >= (v30 | 2)) {
                v37 := GenS5{F0: 2147483647, F1: 2147483648, F2: 77}
                _ = len(v3)
                delete(v3, "foo bar")
                _ = v37
            }
            _ = len(v3)
            _ = fi36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
