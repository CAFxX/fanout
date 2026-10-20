package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint16; F2 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := []string{}
    var v2 uint16 = 130
    v3 := []string{"\n\t", "\n\t", (("foo bar" + ("" + ("hello" + "hello"))) + ("hello" + "foo bar")), "foo bar"}
    v3 = append(v3, ("a" + "\n\t"))
    switch -238 {
    case -878:
        switch ("hello" + "foo bar") {
        case "\n\t":
            for fi4 := 0; fi4 < 5; fi4++ {
                for fi5 := 0; fi5 < 15; fi5++ {
                    fmt.Println(fi5, fi5, fi5)
                    v2 -= (v2 ^ v2)
                    _ = fi5
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi4
            }
        default:
            v6 := GenS1{F0: 255}
            _ = v6
        }
    case -1:
        var v7 uint = 18446744073709551615
        _ = v7
    case -9223372036854775808:
        for fi8 := 0; fi8 < 19; fi8++ {
            if (309 <= 2147483648) {
                v2 -= v2
                var v9 int16 = -32768
                _ = v9
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi8
        }
    default:
        v10 := GenS2{F0: 4294967295}
        _ = v10
    }
    _ = len(v3)
    var v11 int16 = 32767
    var v12 uint = 675
    v13 := GenS3{F0: "\n\t"}
    v14 := GenS4{F0: -538, F1: v2, F2: 640}
    fmt.Println(v11, v2, v11)
    switch "" {
    case "a":
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(fi15, v11, v2)
            v16 := GenS5{F0: 191}
            _ = len(v1)
            _ = fi15
            _ = v16
        }
    }
    if true {
        if (2147483647 >= 42) {
            v17 := GenS6{F0: int16(v12), F1: 914}
            if (((v11 & 975) > 32767) && true) {
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 int8 = -50
                    if false {
                        var v20 uint64 = 18446744073709551615
                        fmt.Println(fi18, v12)
                        _ = v20
                    } else {
                        v21 := GenS3{F0: "\n\t"}
                        _ = v21
                    }
                    if (int16(v2) > int16(v19)) {
                        var v22 int32 = -1
                        _ = v22
                    } else {
                        var v23 int8 = (^(v19))
                        v19 = int8(v11)
                        v24 := GenS7{F0: 440}
                        _ = v23
                        _ = v24
                    }
                    _ = fi18
                    _ = v19
                }
            }
            v12 = (v12 * v12)
            _ = v17
        }
        v25 := GenS7{F0: 2147483647}
        v2 -= (v2 & (^(v2)))
        _ = v25
    }
    var v26 int = -1
    v27 := GenS1{F0: v12}
    var v28 float32 = (((0.1 / 1.0) - 0.0) + float32(-2.25))
    v29 := GenS8{F0: 127, F1: v2}
    if (v12 == v12) {
        var v30 uint64 = 0
        _ = v30
    } else {
        v12 += v12
    }
    fmt.Println(v26, v2, v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
