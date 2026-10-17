package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint16) int64 {
    p0 = p0
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{fi5, (fi5 % (((fi5 % (((^(fi5))) | 1))) | 1))}
        v6 = append(v6, -9223372036854775808)
        p0 = (p0 >> ((0) & 63))
        _ = fi5
        _ = v6
    }
    var v7 int64 = int64(p0)
    _ = p0
    _ = v7
    return (^(v7))
}
func genFn1(p0 bool, p1 uint32, p2 int64) uint16 {
    p2 = (p2 - (p2 | -157))
    for fi8 := 0; fi8 < 18; fi8++ {
        switch "" {
        case "a":
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: 4294967295}
                v11 := []int{fi8}
                var v12 uint64 = 780
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        }
        fmt.Println(p0)
        p2 += p2
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    v4 := []string{(("hello" + "hello") + "hello"), "hello", "", ("\n\t" + "")}
    var v13 uint = 2
    var v14 int32 = 2147483647
    if true {
        switch 702 {
        case 32768:
            for fi15 := 0; fi15 < 16; fi15++ {
                for fi16 := 0; fi16 < 9; fi16++ {
                    v14 = (-(v14))
                    v13 = v13
                    _ = len(v2)
                    _ = fi16
                }
                _ = fi15
            }
        case 4294967295:
            var v17 int32 = v14
            _ = v17
        default:
            _ = len(v2)
        }
    }
    v18 := GenS1{F0: 2}
    v19 := GenS1{F0: 18446744073709551615}
    switch 127 {
    case 1:
        fmt.Println(v13)
    }
    v13 -= (v13 + (v13 >> (((v13 + v13)) & 63)))
    var v20 bool = false
    var v21 uint8 = 1
    var v22 float32 = 1.5
    switch -544 {
    case 999:
        fmt.Println(v22)
    case -9223372036854775808:
        var v23 float64 = 0.0
        _ = v23
    case 1:
        switch 118 {
        case 128:
            v24 := GenS1{F0: 4294967295}
            _ = v24
        case 2147483647:
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 int = fi25
                _ = fi25
                _ = v26
            }
        default:
            v27 := GenS2{F0: 255}
            _ = v27
        }
    default:
        switch "foo bar" {
        case "\n\t":
            switch "" {
            case "foo bar":
                var v28 bool = false
                _ = v28
            }
        case "a":
            v29 := GenS1{F0: 1}
            _ = v29
        case "":
            if v20 {
                var v30 int16 = -1
                _ = v30
            }
        default:
            fmt.Println(v13, v22, v21)
        }
    }
    v4 = append(v4, "a")
    if v20 {
        v22 = float32(3.14159)
        switch "\n\t" {
        case "\n\t":
            var v31 int = 32767
            _ = v31
        case "a":
            fmt.Println(v22, v20)
        case "hello":
            var v32 int32 = -487
            _ = v32
        }
    }
    v33 := GenS3{F0: 802, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v4
}
