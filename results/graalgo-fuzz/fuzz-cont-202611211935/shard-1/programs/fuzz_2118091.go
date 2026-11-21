package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint16 = 0
    v2 := []uint32{}
    var v3 int32 = -2147483648
    _ = len(v2)
    v4 := GenS1{F0: 2}
    if true {
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 int32 = 1
            _ = fi5
            _ = v6
        }
        var v7 uint8 = 195
        if true {
            v8 := GenS1{F0: 209}
            _ = v8
        }
        _ = v7
    } else {
        switch -9223372036854775808 {
        case 1:
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        case 32768:
            for fi9 := 0; fi9 < 4; fi9++ {
                for fi10 := 0; fi10 < 15; fi10++ {
                    fmt.Println(v3)
                    fmt.Println(fi10)
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                    var v11 float64 = 1e10
                    v1 -= (v1 & 42)
                    _ = fi10
                    _ = v11
                }
                var v12 bool = ((18446744073709551615 >= 4294967295) || (true && (1 == 1)))
                _ = len(v2)
                _ = fi9
                _ = v12
            }
        case 128:
            var v13 int16 = 365
            _ = v13
        default:
            var v14 uint8 = 47
            _ = v14
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v15 int64 = 1
        _ = v15
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v3)
    v16 := GenS2{F0: 0}
    for fi17 := 0; fi17 < 18; fi17++ {
        _ = len(v2)
        fmt.Println(v3, fi17, v3)
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = fi17
    }
    v3 = 0
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS3{F0: (42 >= 203), F1: 1}
        var v20 int64 = 44
        v3 = v3
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int8 = -46
    if (true || (false && (float64(3.14159) == 3.14159))) {
        fmt.Println(v21)
        if true {
            _ = len(v2)
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            switch "" {
            case "hello":
                fmt.Println(v21, fi22)
            case "":
                v21 += 127
            case "\n\t":
                v23 := GenS3{F0: true, F1: 255}
                _ = v23
            default:
                fmt.Println(fi22, fi22, fi22)
            }
            v21 -= (v21 & (v21 * 100))
            v21 -= -1
            _ = fi22
        }
    }
    var v24 float32 = (((float32(-2.25) + ((0.0 - (0.1 + ((3.14159 * float32(-2.25)) - 3.14159))) / (0.1 - 1.5))) + 1.5) * 0.0)
    var v25 uint64 = 18446744073709551615
    for fi26 := 0; fi26 < 7; fi26++ {
        fmt.Println(fi26)
        var v27 uint = 492
        v28 := GenS2{F0: (v25 % ((v25) | 1))}
        _ = fi26
        _ = v27
        _ = v28
    }
    if (-673 > -32768) {
        if true {
            fmt.Println(v24, v21)
            var v29 uint64 = (v25 + v25)
            _ = v29
        } else {
            v1 += (^(v1))
            _ = len(v2)
        }
        fmt.Println(v3, v1)
        v2 = append(v2, 0)
    } else {
        switch "a" {
        case "a":
            var v30 uint8 = 78
            _ = v30
        }
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
