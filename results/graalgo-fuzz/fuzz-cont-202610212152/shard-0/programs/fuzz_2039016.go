package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }

func main() {
    v1 := []uint32{}
    v2 := []uint32{65535, 0}
    var v3 uint64 = 4294967295
    v4 := make(map[int]string)
    var v5 uint = 65535
    var v6 int8 = 0
    delete(v4, 42)
    var v7 string = "foo bar"
    v8 := GenS1{F0: 2147483648, F1: 701, F2: 1}
    if (false && false) {
        switch 42 {
        case 1:
            if (("a" + v7) >= "") {
                v9 := GenS2{F0: 530}
                var v10 uint16 = 255
                _ = v10
                _ = v9
            }
        case -999:
            v11 := GenS1{F0: 1, F1: 706, F2: 155}
            _ = v11
        case -502:
            v5 -= 4294967295
        default:
            var v12 float32 = 3.14159
            _ = v12
        }
        v13 := GenS3{F0: 42}
        var v14 int32 = 0
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            fmt.Println(v5, v6)
            _ = fi16
        }
        if false {
            _ = len(v4)
            var v17 bool = ((false && (!(!true))) && (false || false))
            _ = v17
        } else {
            var v18 uint8 = 12
            for fi19 := 0; fi19 < 5; fi19++ {
                v18 = (v18 / (((^(v18))) | 1))
                _ = fi19
            }
            v20 := GenS1{F0: 0, F1: 2, F2: fi15}
            _ = v18
            _ = v20
        }
        _ = fi15
    }
    var v21 uint64 = (^(v3))
    for fi22 := 0; fi22 < 19; fi22++ {
        v21 -= (v3 - 0)
        _ = fi22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v21 -= 2
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
