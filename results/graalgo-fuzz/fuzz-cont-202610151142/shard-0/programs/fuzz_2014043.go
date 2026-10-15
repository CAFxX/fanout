package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int16 }
type GenS2 struct { F0 string; F1 bool; F2 bool }
type GenS3 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint16; F1 int16 }

func main() {
    var v1 uint = 2
    v2 := make(map[string]uint32)
    v3 := []uint32{42, 4294967295}
    if true {
        if (128 <= 133) {
            v1 = 632
            v4 := GenS1{F0: 1, F1: -2147483648, F2: -878}
            _ = v4
        } else {
            if (true && true) {
                for fi5 := 0; fi5 < 5; fi5++ {
                    v1 = (v1 | (v1 >> (((v1 & (v1 & (v1 + (^(v1)))))) & 63)))
                    _ = len(v3)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi5
                }
            } else {
                switch "" {
                case "\n\t":
                    v3 = append(v3, uint32(v1))
                }
                v6 := GenS1{F0: 42, F1: 32768, F2: 338}
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                _ = v6
            }
        }
        if true {
            fmt.Println(v1)
            v1 = 42
            fmt.Println(v1)
        } else {
            v7 := GenS1{F0: 236, F1: 0, F2: -1}
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            switch (fi9 - fi9) {
            case -646:
                v10 := v2["hello"]
                _ = v10
            case 32768:
                v11 := GenS1{F0: 900, F1: (fi8 & (fi8 / (((fi8 << ((65535) & 63))) | 1))), F2: 1000}
                _ = v11
            }
            _ = fi9
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            _ = len(v3)
            var v13 uint8 = 1
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 42, F1: fi8, F2: -1}
        _ = fi8
        _ = v14
    }
    v1 += (v1 - 567)
    v15 := GenS2{F0: "", F1: false, F2: false}
    fmt.Println(v1)
    var v16 int16 = 794
    v17 := GenS3{F0: -36, F1: 4294967295, F2: 0}
    v1 = (v1 >> ((v1) & 63))
    v16 -= (v16 % ((v16) | 1))
    v18 := GenS2{F0: ("" + (("foo bar" + "") + "\n\t")), F1: true, F2: (48 >= 42)}
    fmt.Println(v16, v16)
    for fi19 := 0; fi19 < 9; fi19++ {
        switch fi19 {
        case 538:
            v16 = v16
        case -1:
            v1 = 65535
        case 42:
            var v20 int8 = 1
            _ = v20
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS4{F0: 1, F1: v16}
            _ = fi21
            _ = v22
        }
        _ = fi19
    }
    v1 -= (v1 % ((v1) | 1))
    var v23 float32 = 0.0
    if (!true) {
        if false {
            var v24 uint8 = 2
            if false {
                if true {
                    fmt.Println(v24, v16)
                }
                var v25 float32 = v23
                for fi26 := 0; fi26 < 9; fi26++ {
                    fmt.Println(v23, v16, v24)
                    _ = fi26
                }
                _ = v25
            }
            _ = v24
        }
        if (32767 != 127) {
            var v27 int64 = -9223372036854775808
            switch 128 {
            case -796:
                fmt.Println(v27, v16, v27)
            }
            _ = v27
        } else {
            v16 = 1
        }
    } else {
        v28 := GenS1{F0: 2, F1: 240, F2: (v16 ^ v16)}
        var v29 int64 = 9223372036854775807
        _ = v28
        _ = v29
    }
    v1 -= (v1 + 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
