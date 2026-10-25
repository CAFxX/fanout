package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint }

func main() {
    var v1 bool = (!(!true))
    var v2 int64 = 1
    v3 := []string{"foo bar"}
    v4 := []uint32{766, 0}
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 13; fi6++ {
        if v1 {
            fmt.Println(fi6, v2, fi6)
            var v7 float32 = 3.14159
            _ = v7
        }
        _ = len(v5)
        v8 := GenS1{F0: 563, F1: 4294967295}
        _ = fi6
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := v5[""]
        _ = fi9
        _ = v10
    }
    v11 := v5[""]
    var v12 int64 = v2
    if true {
        var v13 uint16 = 0
        _ = len(v5)
        v11 = ("hello" + "")
        _ = v13
    } else {
        for fi14 := 0; fi14 < 6; fi14++ {
            v12 = (v2 << ((18446744073709551615) & 63))
            for fi15 := 0; fi15 < 20; fi15++ {
                switch "a" {
                case "\n\t":
                    fmt.Println(v11)
                }
                v16 := GenS1{F0: 791, F1: 346}
                if v1 {
                    fmt.Println(v12)
                    var v17 int8 = 1
                    _ = v17
                } else {
                    var v18 uint16 = 723
                    v19 := GenS1{F0: 4294967295, F1: 2}
                    _ = v18
                    _ = v19
                }
                var v20 float64 = 0.0
                fmt.Println(fi15, fi15)
                _ = fi15
                _ = v16
                _ = v20
            }
            _ = fi14
        }
    }
    if true {
        if v1 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 -= (v12 >> ((2) & 63))
        }
        var v21 uint64 = 1
        _ = v21
    }
    if ((!false) || (v2 > v2)) {
        if (805 <= 4294967295) {
            var v22 uint16 = 983
            fmt.Println(v11, v1)
            _ = v22
        }
        fmt.Println(v12)
    }
    v4 = append(v4, 2147483648)
    var v23 string = v11
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS2{F0: 128, F1: 620, F2: 122}
        var v26 uint8 = 128
        switch fi24 {
        case -499:
            for fi27 := 0; fi27 < 9; fi27++ {
                fmt.Println(fi27, v26)
                _ = fi27
            }
        case 32767:
            var v28 uint8 = 127
            _ = v28
        default:
            for fi29 := 0; fi29 < 5; fi29++ {
                fmt.Println(v2, v11, v2)
                switch (v23 + v11) {
                case "a":
                    fmt.Println(v26, v2, fi29)
                case "":
                    var v30 uint32 = 2147483647
                    _ = v30
                case "hello":
                    fmt.Println(v26, fi29, fi24)
                }
                v23 = v23
                _ = fi29
            }
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
