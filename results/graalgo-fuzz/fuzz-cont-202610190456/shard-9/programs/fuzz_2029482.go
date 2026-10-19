package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 string; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int32; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.0
    var v3 uint16 = 880
    v4 := []uint32{679, 42}
    var v5 int32 = 0
    var v6 uint8 = 128
    if false {
        fmt.Println(v3, v6, v2)
        v7 := GenS1{F0: 42}
        _ = v7
    }
    v2 -= v2
    var v8 uint = 18446744073709551615
    for fi9 := 0; fi9 < 15; fi9++ {
        v1[fi9] = "hello"
        if (!(true && false)) {
            v1[(fi9 | fi9)] = ("\n\t" + ("foo bar" + "a"))
            v8 -= v8
            for fi10 := 0; fi10 < 12; fi10++ {
                fmt.Println(v8, v8)
                _ = fi10
            }
        }
        _ = fi9
    }
    if (false || false) {
        switch (("hello" + "") + ("foo bar" + "a")) {
        case "hello":
            v8 = (v8 + v8)
        case "\n\t":
            var v11 int8 = 42
            _ = v11
        case "":
            if (!((-105 >= -1) && (42 == 42))) {
                var v12 uint64 = 4294967295
                var v13 uint8 = 88
                _ = v12
                _ = v13
            }
        }
        v14 := GenS2{F0: (v3 - 2)}
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v2 = ((v2 * ((0.0 / v2) - v2)) - ((0.0 * ((1.5 - (((3.14159 + v2) - 0.1) / v2)) - v2)) / 3.14159))
        if len(v4) > 0 {
            v4[0] = 2
        }
        _ = fi15
    }
    fmt.Println(v2, v6, v6)
    v5 += v5
    var v16 uint16 = (v3 ^ 490)
    var v17 int64 = 9223372036854775807
    v18 := GenS2{F0: 28}
    v19 := GenS3{F0: 18446744073709551615}
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            if (false && (v2 <= (v2 * ((v2 * float32(1.5)) + (0.0 + v2))))) {
                fmt.Println(v17)
            }
            switch "foo bar" {
            case "hello":
                if (4294967295 > 176) {
                    v22 := GenS4{F0: int32(v17), F1: "a", F2: v3}
                    _ = v22
                }
            default:
                fmt.Println(fi20)
            }
            v23 := GenS5{F0: 212, F1: fi21}
            _ = fi21
            _ = v23
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            if false {
                fmt.Println(v17, v2, v17)
            }
            v16 = v16
            if len(v4) > 0 {
                v4[0] = 44
            }
            _ = fi24
        }
        _ = fi20
    }
    v25 := GenS6{F0: v5, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
