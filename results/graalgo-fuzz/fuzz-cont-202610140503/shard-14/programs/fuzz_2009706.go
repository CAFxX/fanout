package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint; F2 int }

func main() {
    v1 := []uint32{1, 263, 4294967295}
    var v2 uint8 = 128
    v3 := make(map[string]uint32)
    v4 := []string{"a", ("" + "a")}
    v5 := GenS1{F0: 128, F1: (true || (260 == 449)), F2: 1000}
    for fi6 := 0; fi6 < 17; fi6++ {
        v2 = 215
        v2 -= (v2 - (v2 | (v2 * (v2 >> ((255) & 63)))))
        if (!false) {
            v2 = v2
            fmt.Println(fi6, fi6)
            var v7 int16 = 1000
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(v2)
    switch "a" {
    case "":
        for fi8 := 0; fi8 < 17; fi8++ {
            _ = len(v3)
            v9 := GenS1{F0: fi8, F1: true, F2: 934}
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := GenS1{F0: 422, F1: true, F2: -1}
                var v12 uint32 = 2147483647
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    case "foo bar":
        v13 := GenS2{F0: 127}
        _ = v13
    case "a":
        v14 := GenS2{F0: -9223372036854775808}
        _ = v14
    default:
        var v15 uint32 = 2147483648
        _ = v15
    }
    v2 = (^(v2))
    v2 = (v2 + v2)
    var v16 int64 = 9223372036854775807
    v17 := GenS3{F0: 226, F1: 777}
    switch 1 {
    case -1:
        v4 = append(v4, ("a" + ""))
    case 27:
        for fi18 := 0; fi18 < 16; fi18++ {
            v4 = append(v4, "")
            _ = fi18
        }
    default:
        for fi19 := 0; fi19 < 14; fi19++ {
            fmt.Println(v2, fi19, fi19)
            v20 := GenS4{F0: 602, F1: 255, F2: (fi19 / ((fi19) | 1))}
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v2, v16)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v16)
    for fi21 := 0; fi21 < 12; fi21++ {
        v16 += (^(v16))
        if false {
            for fi22 := 0; fi22 < 11; fi22++ {
                fmt.Println(fi22, v2)
                if len(v1) > 0 {
                    v1[0] = uint32(v2)
                }
                _ = fi22
            }
        }
        _ = fi21
    }
    var v23 uint = 18446744073709551615
    v23 = (v23 ^ (v23 - (v23 ^ v23)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
