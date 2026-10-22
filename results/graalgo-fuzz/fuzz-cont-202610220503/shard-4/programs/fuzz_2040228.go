package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{65535}
    v3 := []uint32{2, 1}
    v4 := []uint32{4294967295}
    var v5 int8 = -128
    v6 := GenS1{F0: 255}
    var v7 string = "hello"
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := GenS2{F0: 345}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v7 = (v7 + v7)
    v5 -= (v5 ^ v5)
    v7 = v7
    v5 = v5
    v11 := GenS1{F0: 65535}
    v12 := GenS1{F0: 42}
    v7 = ("foo bar" + v7)
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int32 = -751
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        switch v7 {
        case "foo bar":
            v1[fi15] = (v7 + "\n\t")
        case "":
            v5 = 127
        default:
            v16 := GenS3{F0: int8(fi15), F1: (^(fi15))}
            _ = v16
        }
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 string = "hello"
            switch ("" + v7) {
            case "foo bar":
                v5 += v5
            }
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 uint = 1
            fmt.Println(fi19, v7)
            if true {
                var v21 uint32 = 2147483648
                _ = v21
            }
            var v22 uint32 = 42
            fmt.Println(v20, v7)
            var v23 uint16 = 65535
            _ = fi19
            _ = v20
            _ = v22
            _ = v23
        }
        _ = fi15
    }
    fmt.Println(v7, v7)
    v7 = (v7 + "foo bar")
    v4 = append(v4, 2147483648)
    var v24 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
