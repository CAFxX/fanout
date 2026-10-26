package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 int; F2 int32 }
type GenS7 struct { F0 uint; F1 bool; F2 bool }
type GenS8 struct { F0 int16; F1 uint64; F2 int }
type GenS9 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (-650 > -2147483648)
    v3 := make(map[int]uint32)
    v4 := []string{"", "foo bar", ("a" + ("hello" + "a")), ((("hello" + ("a" + ("\n\t" + ((("" + "") + "hello") + "\n\t")))) + "a") + "foo bar")}
    var v5 uint64 = 2
    var v6 uint8 = 30
    v7 := GenS1{F0: 2147483648, F1: -32768}
    fmt.Println(v2, v5, v6)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int8 = -128
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        delete(v1, 2147483648)
        var v11 uint8 = v6
        _ = fi10
        _ = v11
    }
    v2 = (v2 && (!v2))
    var v12 int16 = -32768
    v6 += v6
    v13 := GenS2{F0: 1}
    v3[42] = 2
    v14 := GenS3{F0: (v12 << ((65535) & 63))}
    for fi15 := 0; fi15 < 12; fi15++ {
        v2 = v2
        var v16 int8 = 95
        _ = fi15
        _ = v16
    }
    v6 -= v6
    switch "hello" {
    case "hello":
        if (!v2) {
            v5 = 186
        }
    case "":
        v17 := GenS4{F0: 0, F1: 42, F2: 65535}
        _ = v17
    case "\n\t":
        fmt.Println(v6, v12)
    }
    v18 := GenS5{F0: 114}
    if v2 {
        fmt.Println(v12, v6, v6)
        switch 0 {
        case 127:
            v6 += 147
        case 128:
            if (41 != 255) {
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(v2, v12)
                    _ = fi19
                }
                var v20 float32 = 1e10
                if v2 {
                    _ = len(v4)
                } else {
                    v21 := GenS6{F0: 42, F1: 754, F2: -1}
                    v22 := GenS6{F0: 255, F1: -79, F2: 199}
                    var v23 uint = 938
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = v20
            } else {
                v24 := GenS7{F0: 18446744073709551615, F1: v2, F2: true}
                _ = v24
            }
        case -472:
            var v25 uint64 = v5
            _ = v25
        }
        v5 += v5
    }
    if true {
        v1[65535] = -2147483648
        fmt.Println(v2)
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS8{F0: (v12 % ((-144) | 1)), F1: v5, F2: fi26}
            v28 := GenS9{F0: (v6 | (v6 + 121)), F1: (v5 % ((2) | 1))}
            _ = fi26
            _ = v27
            _ = v28
        }
    } else {
        fmt.Println(v12, v12)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
