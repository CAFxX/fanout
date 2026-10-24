package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0() int {
    var v4 string = "a"
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{fi5, fi5, fi5}
        v7 := GenS1{F0: (false && (!false)), F1: v4}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = v4
    return 42
}

func main() {
    var v1 uint = 392
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    if false {
        for fi8 := 0; fi8 < 3; fi8++ {
            switch "foo bar" {
            case "foo bar":
                v1 += (v1 >> (((v1 & v1)) & 63))
            case "":
                v1 = (v1 - (^(v1)))
            case "\n\t":
                v1 = 636
            default:
                v9 := []int{4294967295, (fi8 & fi8)}
                _ = v9
            }
            _ = len(v2)
            _ = fi8
        }
        delete(v3, ("" + "hello"))
    } else {
        v10 := GenS2{F0: 0}
        var v11 int64 = -1
        v12 := GenS3{F0: 0, F1: 0}
        _ = v10
        _ = v11
        _ = v12
    }
    fmt.Println(v1)
    var v13 uint = v1
    if false {
        v14 := GenS2{F0: 2}
        _ = v14
    } else {
        for fi15 := 0; fi15 < 5; fi15++ {
            v1 = 18446744073709551615
            _ = fi15
        }
        v16 := []int{4294967295, 32767}
        var v17 int32 = -2147483648
        _ = v16
        _ = v17
    }
    fmt.Println(v1, v1)
    for fi18 := 0; fi18 < 5; fi18++ {
        v13 -= v1
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS4{F0: 621, F1: 550}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v13 += (v1 / ((v1) | 1))
    delete(v2, 256)
    v1 = (v1 ^ 18446744073709551615)
    v1 = v1
    v21 := []int{128, genFn0(), -851}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
}
