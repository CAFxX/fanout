package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int8 }

func main() {
    var v1 int = 32768
    var v2 uint64 = 0
    v3 := make(map[int]string)
    var v4 int16 = 659
    var v5 float32 = 3.14159
    fmt.Println(v5, v1)
    if (true && true) {
        v6 := v3[42]
        _ = v6
    } else {
        v5 = v5
        fmt.Println(v1, v4)
        v2 = v2
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v2, fi7)
        _ = fi7
    }
    if ((!false) || false) {
        var v8 string = "\n\t"
        _ = v8
    } else {
        v2 += 77
        switch (^(v1)) {
        case 32767:
            fmt.Println(v4, v5, v2)
        }
        v9 := GenS1{F0: 2147483647, F1: 4294967295, F2: 127}
        _ = v9
    }
    switch "foo bar" {
    case "hello":
        var v10 int16 = 1000
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(v5)
            if true {
                v12 := GenS2{F0: 617, F1: 2147483647, F2: 415}
                _ = v12
            }
            _ = fi11
        }
    default:
        v13 := GenS3{F0: v2}
        _ = v13
    }
    v4 -= (v4 / (((v4 + -959)) | 1))
    v1 += (-(v1))
    if false {
        switch "" {
        case "foo bar":
            var v14 uint16 = 515
            _ = v14
        default:
            var v15 string = (("hello" + "hello") + (("hello" + ("foo bar" + "\n\t")) + ("foo bar" + "\n\t")))
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint16 = 2
            switch "" {
            case "hello":
                v5 -= float32(3.14159)
            case "foo bar":
                _ = len(v3)
            default:
                switch int(v2) {
                case -523:
                    v1 = 487
                case 208:
                    v5 -= v5
                case 2147483647:
                    fmt.Println(v2)
                }
            }
            _ = fi16
            _ = v17
        }
    } else {
        v1 -= (v1 + 4294967295)
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        switch "a" {
        case "\n\t":
            if false {
                v19 := v3[1]
                var v20 uint = 2
                _ = v19
                _ = v20
            } else {
                _ = len(v3)
                v21 := GenS4{F0: 1, F1: 42}
                _ = v21
            }
        case "hello":
            v2 += v2
        }
        switch fi18 {
        case 0:
            v1 += v1
        }
        v22 := GenS5{F0: 0, F1: 100}
        _ = fi18
        _ = v22
    }
    var v23 int64 = 604
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
