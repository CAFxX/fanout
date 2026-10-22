package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint16; F2 int32 }

func main() {
    var v1 uint8 = 42
    v2 := []int{42, 127}
    var v3 int16 = 992
    var v4 uint64 = 132
    var v5 uint8 = 192
    switch "\n\t" {
    case "hello":
        switch "foo bar" {
        case "a":
            var v6 int = 127
            _ = v6
        case "hello":
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
        }
    }
    switch 9223372036854775807 {
    case -930:
        v7 := GenS1{F0: "", F1: (v4 << ((18446744073709551615) & 63)), F2: 2}
        _ = v7
    case 32767:
        v8 := GenS1{F0: ("\n\t" + "\n\t"), F1: 2, F2: 255}
        _ = v8
    default:
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint32 = 0
            if ((v4 >> ((42) & 63)) == v4) {
                fmt.Println(v1)
                v11 := GenS2{F0: 2, F1: v10, F2: 2}
                for fi12 := 0; fi12 < 12; fi12++ {
                    v13 := GenS1{F0: "foo bar", F1: (v4 ^ 689), F2: 42}
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
            _ = fi9
            _ = v10
        }
    }
    v5 -= (v5 | v1)
    var v14 uint64 = v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v15 int16 = (-(v3))
    var v16 int = 659
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 string = "a"
        v2 = append(v2, 1)
        v19 := GenS3{F0: 2, F1: (v5 % ((v5) | 1)), F2: 1}
        _ = fi17
        _ = v18
        _ = v19
    }
    if true {
        v4 -= (v14 % (((v14 & 836)) | 1))
    }
    var v20 int8 = -128
    v21 := GenS4{F0: (v5 * uint8(v4)), F1: 477, F2: -1}
    if (!true) {
        v2 = append(v2, 32768)
    }
    var v22 string = "foo bar"
    switch v22 {
    case "\n\t":
        var v23 uint8 = 225
        _ = v23
    case "a":
        v20 += (v20 * (v20 % ((v20) | 1)))
    default:
        var v24 string = ""
        _ = v24
    }
    var v25 int8 = 47
    v2 = append(v2, (v16 >> ((0) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
