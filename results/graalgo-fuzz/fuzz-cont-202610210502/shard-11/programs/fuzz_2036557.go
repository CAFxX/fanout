package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
type GenS3 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 0
    v3 := []uint32{65535}
    var v4 int = 42
    var v5 uint8 = 2
    v6 := make(map[int]string)
    if (!(32767 > 0)) {
        var v7 uint16 = 2
        var v8 int = (v4 + int(v7))
        switch "a" {
        case "hello":
            v8 = (^(v8))
        }
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 42}
        fmt.Println(v4)
        _ = v9
    }
    v1[4294967295] = (false && ((v2 > 26) && (96 >= -524)))
    v6[-909] = "a"
    var v10 int16 = -528
    v5 = 127
    for fi11 := 0; fi11 < 5; fi11++ {
        v2 = 65535
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v4)
            _ = fi12
        }
        delete(v6, (v4 + 128))
        _ = fi11
    }
    switch (v4 & 128) {
    case 2147483647:
        v5 -= v5
    case 347:
        var v13 bool = (906 == 4294967295)
        _ = v13
    case 9223372036854775807:
        v3 = append(v3, 4294967295)
    default:
        if true {
            switch "\n\t" {
            case "a":
                switch "\n\t" {
                case "hello":
                    v14 := GenS1{F0: (v10 * v10)}
                    _ = v14
                case "foo bar":
                    v15 := GenS1{F0: (v10 % (((v10 >> (((v2 % (((^(v2))) | 1))) & 63))) | 1))}
                    _ = v15
                }
            case "foo bar":
                fmt.Println(v2, v2)
            case "hello":
                v16 := GenS1{F0: (v10 - 32767)}
                _ = v16
            }
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 float32 = float32(3.14159)
                _ = fi17
                _ = v18
            }
        }
    }
    if (0 > 165) {
        fmt.Println(v10, v4, v5)
        v19 := GenS1{F0: 543}
        _ = v19
    } else {
        if ((2147483647 >= -2147483648) && (255 == 0)) {
            var v20 int16 = (-(v10))
            _ = v20
        }
        v5 = 255
    }
    switch (v4 << ((v2) & 63)) {
    case 591:
        if true {
            var v21 float64 = 3.14159
            _ = v21
        } else {
            v22 := GenS2{F0: v10, F1: 65535, F2: (!(!(uint64(v5) >= 42)))}
            v23 := GenS3{F0: 88, F1: 42}
            _ = v22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 bool = true
        _ = fi24
        _ = v25
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v26 int = (v4 & (v4 - v4))
    if true {
        fmt.Println(v10, v2, v2)
        if ((v26 % ((int(v2)) | 1)) < (v4 & (v4 ^ (-(v4))))) {
            var v27 float32 = 0.0
            for fi28 := 0; fi28 < 1; fi28++ {
                if (false && ((^(v5)) < (^(v5)))) {
                    var v29 float64 = float64((v27 - v27))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v2, v5, v29)
                    _ = v29
                } else {
                    v30 := GenS2{F0: (v10 & (v10 % (((-(v10))) | 1))), F1: 904, F2: (v26 != v4)}
                    v31 := GenS3{F0: 39, F1: 2147483647}
                    _ = v30
                    _ = v31
                }
                delete(v6, 77)
                _ = fi28
            }
            _ = v27
        } else {
            var v32 int8 = 0
            var v33 uint = (v2 - v2)
            _ = len(v6)
            _ = v32
            _ = v33
        }
    }
    v34 := GenS3{F0: -128, F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
