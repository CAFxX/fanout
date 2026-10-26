package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 string; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int8 }
func genFn0(p0 uint8, p1 uint) bool {
    p1 -= p1
    v6 := []int{32768}
    _ = p0
    _ = p1
    _ = v6
    return ((true && (-9223372036854775808 == 32768)) && false)
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{435, 567, 2147483647, 2147483648}
    var v3 int32 = (v1 << ((0) & 63))
    var v4 int = 188
    var v5 int8 = -86
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = false
        v8 = (((!false) && true) || (true || true))
        switch "hello" {
        case "foo bar":
            v1 = 465
        case "hello":
            v9 := GenS1{F0: "foo bar", F1: (v3 << ((18446744073709551615) & 63))}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint64 = 304
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(v11, v11, v4)
            _ = fi12
        }
        if ((!false) && true) {
            var v13 uint8 = 1
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    if true {
        v1 = (v1 ^ 378)
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS1{F0: ((("" + "") + "foo bar") + "a"), F1: v3}
            switch fi14 {
            case 803:
                switch (fi14 | (v4 ^ v4)) {
                case 629:
                    _ = len(v2)
                case 128:
                    v16 := GenS1{F0: "\n\t", F1: -1}
                    _ = v16
                default:
                    v3 += 2147483647
                }
            default:
                switch "\n\t" {
                case "":
                    v17 := GenS1{F0: "a", F1: v1}
                    _ = v17
                case "foo bar":
                    v4 -= v4
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = 435
                    }
                }
            }
            _ = fi14
            _ = v15
        }
    }
    var v18 uint16 = 32767
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v19 := GenS2{F0: "", F1: 0, F2: 2147483647}
    var v20 int64 = 1
    v20 -= v20
    var v21 int = (v4 ^ (v4 | int(v1)))
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := GenS3{F0: 0, F1: 2, F2: 127}
        _ = fi22
        _ = v23
    }
    var v24 int32 = int32(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
