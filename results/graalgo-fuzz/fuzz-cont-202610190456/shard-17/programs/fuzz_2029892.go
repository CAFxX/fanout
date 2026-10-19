package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint32 }
func genFn0(p0 int8) int64 {
    v5 := GenS1{F0: 42}
    v6 := []int{-280, 0}
    _ = p0
    _ = v5
    _ = v6
    return int64(p0)
}

func main() {
    v1 := []uint32{101, 1}
    var v2 uint16 = 65535
    var v3 string = ""
    v4 := make(map[uint32]int)
    var v7 int = -651
    fmt.Println(v3)
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int = (-(fi8))
            var v11 int32 = -678
            v11 = (v11 & v11)
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            v7 += 1
        } else {
            var v12 int16 = 1000
            var v13 uint64 = 528
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    switch (v3 + "") {
    case "":
        v1 = append(v1, 0)
    case "hello":
        v14 := GenS2{F0: 127, F1: -1, F2: 65535}
        _ = v14
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
    }
    v3 = ((((v3 + v3) + ((((v3 + ("foo bar" + ((("a" + v3) + v3) + (v3 + (v3 + ("foo bar" + "\n\t")))))) + (v3 + "\n\t")) + "hello") + ((v3 + v3) + "\n\t"))) + ("foo bar" + v3)) + v3)
    switch v3 {
    case "hello":
        v15 := GenS1{F0: -2147483648}
        _ = v15
    default:
        var v16 int16 = -393
        _ = v16
    }
    fmt.Println(v7, v2, v7)
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 int8 = 42
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(fi17, v3)
            _ = len(v4)
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    v2 = (v2 >> ((2) & 63))
    switch 42 {
    case 4294967295:
        var v20 int16 = 0
        _ = v20
    case 32768:
        v7 += (v7 << ((1) & 63))
    case 296:
        var v21 int = v7
        _ = v21
    default:
        v7 = v7
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
