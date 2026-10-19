package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 string }
type GenS3 struct { F0 int64; F1 bool; F2 int8 }
func genFn0() int {
    if true {
        v5 := GenS1{F0: -969, F1: 1, F2: (-1 != -353)}
        _ = v5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(fi6)
        var v7 int8 = 0
        _ = fi6
        _ = v7
    }
    if true {
        var v8 uint32 = 0
        v8 += v8
        switch -883 {
        case -930:
            v9 := GenS1{F0: 1, F1: 1, F2: (true && true)}
            _ = v9
        }
        _ = v8
    }
    return 0
}

func main() {
    var v1 int64 = -1
    var v2 uint8 = 128
    v3 := []uint32{}
    var v4 int32 = 42
    v2 += v2
    var v10 int8 = 1
    var v11 int32 = (^(v4))
    var v12 string = ("" + ("" + "hello"))
    v13 := GenS1{F0: 2147483647, F1: 713, F2: (103 > (v2 | (v2 >> ((255) & 63))))}
    v10 -= 84
    _ = len(v3)
    if (94 != (v1 | v1)) {
        v11 -= -218
    }
    fmt.Println(v12, v2, v12)
    switch -9223372036854775808 {
    case -9223372036854775808:
        v3 = append(v3, 2)
    case -252:
        for fi14 := 0; fi14 < 13; fi14++ {
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 string = "a"
                v17 := GenS1{F0: (v4 - v4), F1: 2147483647, F2: true}
                switch fi14 {
                case 84:
                    v16 = v12
                case -371:
                    fmt.Println(v10)
                default:
                    fmt.Println(v2)
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v18 := GenS2{F0: -87, F1: (v2 ^ v2), F2: "hello"}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi14
        }
    case 4294967295:
        for fi19 := 0; fi19 < 4; fi19++ {
            v2 += v2
            var v20 uint32 = 65535
            switch (fi19 + fi19) {
            case 0:
                v20 = (v20 * uint32(v2))
            case -9223372036854775808:
                var v21 int8 = (v10 & v10)
                _ = v21
            }
            _ = fi19
            _ = v20
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v22 uint16 = 0
    for fi23 := 0; fi23 < 15; fi23++ {
        if ((v2 / ((v2) | 1)) < v2) {
            v10 -= (v10 % (((^(v10))) | 1))
        }
        v10 = (v10 | (^(v10)))
        _ = fi23
    }
    var v24 int8 = v10
    var v25 uint = 18446744073709551615
    v3 = append(v3, 65535)
    for fi26 := 0; fi26 < 10; fi26++ {
        fmt.Println(fi26)
        var v27 uint = 18446744073709551615
        _ = fi26
        _ = v27
    }
    v28 := GenS3{F0: (v1 * (v1 ^ v1)), F1: ((^(v10)) == (v10 * (v24 % ((0) | 1)))), F2: v24}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
