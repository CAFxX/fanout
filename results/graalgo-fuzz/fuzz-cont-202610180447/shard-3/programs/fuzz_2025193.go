package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int32; F1 string; F2 int16 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -672
    v3 := make(map[string]string)
    if true {
        for fi4 := 0; fi4 < 17; fi4++ {
            for fi5 := 0; fi5 < 13; fi5++ {
                var v6 uint = 18446744073709551615
                v7 := []int{}
                _ = fi5
                _ = v6
                _ = v7
            }
            fmt.Println(fi4, fi4)
            switch "a" {
            case "a":
                fmt.Println(fi4)
            default:
                var v8 bool = (42 > 9223372036854775807)
                _ = v8
            }
            _ = fi4
        }
        fmt.Println(v2)
    } else {
        switch (^(v2)) {
        case 9223372036854775807:
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 float64 = (3.14159 - 0.1)
                switch fi9 {
                case 584:
                    fmt.Println(v2, v2, v2)
                case 4294967295:
                    var v11 int16 = 1
                    _ = v11
                case -912:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        var v13 int8 = 42
                        v13 -= -77
                        v14 := GenS1{F0: 198}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                v10 -= float64(1.5)
                _ = fi9
                _ = v10
            }
        case 127:
            switch 1 {
            case 32768:
                v2 -= (v2 << ((18446744073709551615) & 63))
            }
        case -9223372036854775808:
            for fi15 := 0; fi15 < 10; fi15++ {
                v2 += (v2 - 128)
                _ = fi15
            }
        }
    }
    v2 = (v2 ^ 44)
    var v16 uint64 = 893
    switch "" {
    case "foo bar":
        for fi17 := 0; fi17 < 15; fi17++ {
            switch ("hello" + (("hello" + "foo bar") + "a")) {
            case "\n\t":
                v2 = (^(v2))
            case "a":
                v3["hello"] = "\n\t"
            case "foo bar":
                if true {
                    v18 := GenS2{F0: ("foo bar" + ("foo bar" + "")), F1: ((0 == 0) && (true && true))}
                    _ = v18
                } else {
                    v2 += 42
                }
            }
            _ = fi17
        }
    case "\n\t":
        var v19 float32 = ((3.14159 - -2.25) - (0.0 + 1.5))
        _ = v19
    case "hello":
        v2 = (v2 + v2)
    }
    var v20 uint32 = 624
    fmt.Println(v16)
    v21 := GenS1{F0: 1}
    v2 -= v2
    v22 := GenS1{F0: 42}
    fmt.Println(v20, v2)
    v20 = (^(v20))
    if false {
        fmt.Println(v16, v20)
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(fi23)
            _ = fi23
        }
    } else {
        v24 := GenS3{F0: 0, F1: "\n\t", F2: 42}
        fmt.Println(v20, v2, v16)
        v20 += v20
        _ = v24
    }
    v25 := GenS4{F0: true, F1: 255, F2: (v16 | (^(v16)))}
    v2 = v2
    v26 := GenS3{F0: 42, F1: "\n\t", F2: 1000}
    switch "\n\t" {
    case "foo bar":
        v2 = (v2 << ((1) & 63))
    case "":
        for fi27 := 0; fi27 < 18; fi27++ {
            switch "hello" {
            case "foo bar":
                v20 = v20
            default:
                var v28 int8 = -128
                _ = v28
            }
            _ = fi27
        }
    case "\n\t":
        _ = len(v3)
    default:
        v20 = (v20 & 4294967295)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
}
