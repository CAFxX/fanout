package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int32; F2 int }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint16 = 255
    var v2 uint16 = v1
    var v3 uint64 = 677
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]bool)
    v6 := []uint32{802}
    v3 += v3
    switch 1 {
    case 2147483647:
        switch 559 {
        case -1:
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: -1, F1: true}
                fmt.Println(v3, v3)
                for fi9 := 0; fi9 < 14; fi9++ {
                    v3 = v3
                    switch (fi9 - fi9) {
                    case 4294967295:
                        v2 += v2
                    case -2147483648:
                        v10 := GenS2{F0: (fi9 >> ((0) & 63))}
                        _ = v10
                    case 962:
                        v11 := GenS3{F0: 42, F1: 452, F2: fi9}
                        _ = v11
                    }
                    _ = fi9
                }
                _ = fi7
                _ = v8
            }
        default:
            v3 = v3
        }
    case 32768:
        var v12 uint64 = 4294967295
        _ = v12
    default:
        for fi13 := 0; fi13 < 10; fi13++ {
            if ((1.5 / 1.5) == 0.0) {
                v2 -= (v1 % (((v1 / ((v1) | 1))) | 1))
                v14 := GenS2{F0: 173}
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 uint8 = 1
                    switch fi13 {
                    case -817:
                        v17 := GenS4{F0: (!(!false))}
                        _ = v17
                    default:
                        var v18 uint32 = 2
                        _ = v18
                    }
                    v2 = v1
                    _ = fi15
                    _ = v16
                }
                _ = v14
            } else {
                v3 += (v3 + (v3 << ((255) & 63)))
                if true {
                    var v19 string = (("hello" + "") + "foo bar")
                    _ = len(v5)
                    _ = v19
                }
                fmt.Println(v3, v3, v3)
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi13
        }
    }
    v2 = 0
    v1 = (v1 + v2)
    fmt.Println(v2, v2)
    v1 = (v1 % ((uint16(v3)) | 1))
    var v20 float32 = ((((1.5 + (((3.14159 / 0.1) * 0.0) + 3.14159)) * (1.5 * 0.0)) - ((((1.5 - (((3.14159 / -2.25) * (-2.25 - 1.5)) / 3.14159)) / 0.1) + 0.0) - -2.25)) - -2.25)
    v1 = v1
    v5[128] = (false && true)
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
