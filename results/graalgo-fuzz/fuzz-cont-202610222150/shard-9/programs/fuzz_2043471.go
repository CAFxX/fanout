package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int32; F2 int }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 8; fi4++ {
        v2 = 18446744073709551615
        fmt.Println(v2, v2)
        _ = fi4
    }
    if ((0 > v2) && (true || ((-2.25 + (0.1 * 0.0)) == (0.1 * 1e10)))) {
        for fi5 := 0; fi5 < 12; fi5++ {
            fmt.Println(fi5, v2)
            _ = fi5
        }
    }
    _ = len(v1)
    var v6 string = "a"
    switch v6 {
    case "hello":
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 uint16 = 42
            var v9 uint64 = (v2 << ((0) & 63))
            _ = fi7
            _ = v8
            _ = v9
        }
    case "":
        fmt.Println(v6, v6)
    case "a":
        v6 = v6
    }
    v6 = v6
    v6 = v6
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint16 = 42
        v12 := GenS1{F0: 229}
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v6, v2)
    for fi13 := 0; fi13 < 9; fi13++ {
        v2 = v2
        fmt.Println(v2)
        v6 = v6
        _ = fi13
    }
    switch "a" {
    case "hello":
        v3[-465] = true
    case "\n\t":
        var v14 uint64 = v2
        _ = v14
    default:
        if (277 < 190) {
            var v15 uint16 = uint16(v2)
            if (!false) {
                var v16 int16 = -304
                v2 += v2
                _ = v16
            } else {
                v17 := GenS1{F0: 127}
                _ = v17
            }
            fmt.Println(v6)
            _ = v15
        } else {
            switch 128 {
            case -1:
                v18 := GenS2{F0: 933, F1: -886, F2: -1}
                _ = v18
            case 2147483647:
                v6 = (v6 + "")
            case 32768:
                switch 32768 {
                case -958:
                    fmt.Println(v2)
                case -2147483648:
                    fmt.Println(v2, v2)
                case 127:
                    v19 := []int{-332, 32767}
                    _ = v19
                default:
                    v20 := []int{32768, 4294967295, 143, 775}
                    _ = v20
                }
            }
        }
    }
    v21 := GenS3{F0: 108}
    v22 := GenS1{F0: 2}
    v6 = (v6 + v6)
    switch -547 {
    case 32768:
        v23 := []int{-143, 0, 894, 9223372036854775807}
        _ = v23
    case -2147483648:
        v2 -= v2
    case 1:
        v24 := v3[0]
        _ = v24
    default:
        var v25 int8 = -79
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 float32 = 0.1
        var v28 float64 = 0.0
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v6
}
