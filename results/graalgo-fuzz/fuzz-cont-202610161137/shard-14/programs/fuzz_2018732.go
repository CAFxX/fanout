package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int }
type GenS3 struct { F0 string }

func main() {
    var v1 int8 = 90
    var v2 float32 = 3.14159
    var v3 uint16 = 0
    var v4 uint8 = 229
    v5 := []string{}
    if (-568 >= -308) {
        if true {
            var v6 bool = false
            _ = v6
        } else {
            for fi7 := 0; fi7 < 16; fi7++ {
                for fi8 := 0; fi8 < 19; fi8++ {
                    v2 = v2
                    v9 := GenS1{F0: ("foo bar" + "hello"), F1: 0}
                    v10 := GenS2{F0: (v4 << ((18446744073709551615) & 63)), F1: 65535, F2: (^(fi8))}
                    fmt.Println(v4, fi7, v4)
                    v11 := GenS1{F0: ("foo bar" + "foo bar"), F1: 4294967295}
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                _ = fi7
            }
            if false {
                var v12 int = 1
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v3 = 11
                    if true {
                        v2 = v2
                    }
                    var v14 uint16 = (v3 % ((42) | 1))
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 8; fi15++ {
                    v4 = (v4 - uint8(v3))
                    _ = fi15
                }
                if len(v5) > 0 {
                    v5[0] = ""
                }
            }
        }
        v16 := GenS2{F0: (v4 | v4), F1: 0, F2: 4294967295}
        var v17 uint16 = (^(v3))
        _ = v16
        _ = v17
    } else {
        if true {
            switch 32768 {
            case -9223372036854775808:
                v3 += 65535
            default:
                v18 := GenS1{F0: "", F1: 0}
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 uint64 = 932
            _ = fi19
            _ = v20
        }
        if (32767 >= 1) {
            v5 = append(v5, (("\n\t" + "") + "foo bar"))
            v21 := GenS2{F0: 42, F1: 1, F2: -183}
            v22 := GenS2{F0: 128, F1: uint32(v3), F2: -9223372036854775808}
            _ = v21
            _ = v22
        }
    }
    v23 := GenS3{F0: "\n\t"}
    v24 := GenS1{F0: ("" + ""), F1: 1}
    fmt.Println(v1, v1)
    var v25 int64 = -1
    var v26 bool = ((-657 == (v25 * v25)) || false)
    switch "foo bar" {
    case "\n\t":
        v3 -= uint16(v4)
    }
    if true {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v27 uint64 = 42
        v28 := GenS3{F0: ""}
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        _ = len(v5)
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
