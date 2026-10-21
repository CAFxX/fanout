package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int; F1 int64; F2 int8 }
type GenS3 struct { F0 int; F1 int8; F2 int16 }
type GenS4 struct { F0 uint64; F1 int16; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 uint16; F2 uint32 }

func main() {
    var v1 int8 = 127
    var v2 int = -2147483648
    v3 := []string{"\n\t", "a"}
    var v4 int16 = -32768
    var v5 int32 = 1
    v6 := make(map[uint32]uint32)
    if true {
        var v7 uint8 = 127
        var v8 string = "a"
        _ = v7
        _ = v8
    }
    var v9 int64 = 451
    if (4294967295 == 65535) {
        var v10 uint16 = 42
        _ = v10
    }
    fmt.Println(v9, v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v11 := GenS1{F0: (v4 | v4), F1: -20}
    if false {
        _ = len(v6)
        v4 = v4
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int64 = v9
        v1 -= v1
        _ = fi12
        _ = v13
    }
    if true {
        v6[274] = 2
        var v14 int = 9223372036854775807
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS2{F0: fi15, F1: v9, F2: (v1 >> ((4294967295) & 63))}
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v9)
        if false {
            v18 := GenS3{F0: 1, F1: (v1 ^ int8(v4)), F2: 42}
            switch "a" {
            case "\n\t":
                v19 := GenS4{F0: 0, F1: v4, F2: int64(v4)}
                _ = v19
            }
            _ = v18
        } else {
            v20 := GenS3{F0: -9223372036854775808, F1: v1, F2: -670}
            v3 = append(v3, ("foo bar" + "a"))
            _ = v20
        }
        _ = fi17
    }
    if ((!false) && ((v2 & -816) <= (^(v2)))) {
        v6[1] = 0
        _ = len(v3)
        if (v5 < v5) {
            if false {
                switch (("a" + ("foo bar" + ((("foo bar" + ("foo bar" + "a")) + (("\n\t" + "foo bar") + "")) + ""))) + "hello") {
                case "":
                    fmt.Println(v4)
                }
            } else {
                v3 = append(v3, "foo bar")
            }
            v21 := GenS5{F0: -2147483648}
            v22 := GenS6{F0: v2, F1: 32767, F2: 1}
            _ = v21
            _ = v22
        }
    }
    v2 = (v2 ^ (-(v2)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
