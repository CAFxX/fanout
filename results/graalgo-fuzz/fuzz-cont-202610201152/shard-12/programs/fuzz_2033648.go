package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16) int64 {
    var v4 int16 = -1
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(fi5, p0)
        _ = fi5
    }
    v6 := []int{-482, -1}
    switch "a" {
    case "\n\t":
        p0 = p0
    case "":
        switch 127 {
        case -809:
            fmt.Println(v4)
        case -2147483648:
            switch 452 {
            case 1:
                var v7 uint8 = 128
                _ = v7
            case -2147483648:
                v6 = append(v6, 128)
            case 42:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                for fi8 := 0; fi8 < 5; fi8++ {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        _ = len(v6)
                        v6 = append(v6, (fi9 + fi9))
                        _ = len(v6)
                        _ = fi9
                    }
                    _ = fi8
                }
            }
        default:
            switch 42 {
            case 9223372036854775807:
                v6 = append(v6, -752)
            case 698:
                var v10 uint = 2
                _ = v10
            case 4294967295:
                v11 := GenS1{F0: "a", F1: -1, F2: 32768}
                _ = v11
            }
        }
    case "hello":
        v4 -= (v4 | v4)
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v4
    _ = v6
    return 41
}
func genFn1() uint8 {
    var v12 uint8 = 128
    v12 = v12
    _ = v12
    return (v12 >> ((2) & 63))
}

func main() {
    var v1 string = (("a" + "foo bar") + "\n\t")
    v2 := []uint32{221, 6}
    v3 := []uint32{2}
    var v13 int64 = genFn0(1)
    for fi14 := 0; fi14 < 2; fi14++ {
        _ = len(v3)
        var v15 uint32 = 52
        _ = fi14
        _ = v15
    }
    var v16 int = 1
    v17 := GenS1{F0: "", F1: 42, F2: (v16 - v16)}
    v3 = append(v3, 4294967295)
    v13 -= (v13 * (v13 * v13))
    v18 := GenS2{F0: 239}
    v16 = (^(v16))
    var v19 uint8 = 255
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
