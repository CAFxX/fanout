package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 bool }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 bool }

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{42}
    v3 := make(map[uint32]int)
    var v4 uint = 4294967295
    v5 := make(map[uint32]string)
    switch "\n\t" {
    case "foo bar":
        v6 := GenS1{F0: 4294967295, F1: 81, F2: (true && false)}
        _ = v6
    }
    if true {
        v1 = v1
        fmt.Println(v4, v4)
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS2{F0: (fi7 != fi7), F1: -32768}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        switch (fi9 * 921) {
        case 32768:
            v4 = 65535
        case 32767:
            for fi10 := 0; fi10 < 3; fi10++ {
                switch "foo bar" {
                case "a":
                    if true {
                        var v11 float64 = ((((-2.25 * 0.0) - (float64((-2.25 + (float32(((0.0 / 3.14159) / 1.0)) + 0.0))) / 1.5)) - 3.14159) / 3.14159)
                        if (4294967295 > 18446744073709551615) {
                            if false {
                                var v12 float64 = v11
                                v13 := GenS1{F0: 670, F1: uint64(v4), F2: false}
                                v11 -= v11
                                if true {
                                    var v14 int64 = 99
                                    fmt.Println(fi10)
                                    fmt.Println(v4, v12, v14)
                                    _ = v14
                                } else {
                                    fmt.Println(v12, fi10)
                                }
                                _ = v12
                                _ = v13
                            }
                            v1 = "foo bar"
                            v11 = 1e10
                        }
                        _ = v11
                    } else {
                        v1 = v1
                    }
                    var v15 uint8 = 42
                    fmt.Println(fi10)
                    v15 += (v15 + 107)
                    _ = v15
                case "\n\t":
                    v4 += 65535
                    var v16 int32 = -316
                    _ = v16
                default:
                    var v17 float32 = ((float32(3.14159) / ((-2.25 / float32(0.1)) + (1e10 / 1.5))) / -2.25)
                    _ = v17
                }
                _ = fi10
            }
        case 9223372036854775807:
            v4 = (v4 >> ((2) & 63))
        }
        v1 = v1
        v4 = v4
        _ = fi9
    }
    v1 = v1
    switch v1 {
    case "":
        if true {
            var v18 int16 = int16(v4)
            switch "\n\t" {
            case "":
                v19 := GenS1{F0: 692, F1: 4294967295, F2: true}
                _ = v19
            }
            v18 += v18
            _ = v18
        } else {
            switch "" {
            case "a":
                fmt.Println(v1, v4)
            }
        }
    case "hello":
        v4 = v4
    default:
        v20 := GenS3{F0: (223 >= 42)}
        _ = v20
    }
    switch (v1 + v1) {
    case "":
        var v21 float64 = 0.0
        _ = v21
    }
    v22 := v5[306]
    fmt.Println(v4)
    v22 = v22
    v4 += (v4 - 1)
    v23 := GenS2{F0: (-1 == 81), F1: int16(v4)}
    fmt.Println(v4, v1)
    if ((452 < 1) && false) {
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 string = ("foo bar" + v22)
            _ = fi24
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 10; fi26++ {
            for fi27 := 0; fi27 < 19; fi27++ {
                var v28 int = (fi27 | 1)
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        v29 := GenS2{F0: (1 > 158), F1: 32767}
        switch 42 {
        case 32767:
            _ = len(v2)
        }
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
