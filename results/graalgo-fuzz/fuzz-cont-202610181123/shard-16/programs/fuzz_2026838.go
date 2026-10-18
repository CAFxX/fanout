package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint8) uint16 {
    var v6 int = -6
    p0 += (p1 / ((165) | 1))
    _ = p0
    _ = p1
    _ = v6
    return 389
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 0
    v3 := []uint32{0, 42, 0}
    v4 := []string{"foo bar", "a"}
    var v5 uint = 65535
    v5 += (v5 | v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: (v5 | v5), F1: 4294967295}
        _ = fi7
        _ = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    for fi9 := 0; fi9 < 4; fi9++ {
        switch 627 {
        case -616:
            delete(v1, 378)
        }
        var v10 uint8 = 0
        v2 = v2
        _ = fi9
        _ = v10
    }
    var v11 int32 = 2147483647
    v5 += (v5 >> (((v5 / ((v5) | 1))) & 63))
    fmt.Println(v5)
    if true {
        for fi12 := 0; fi12 < 7; fi12++ {
            delete(v1, uint32(v2))
            var v13 float32 = 0.0
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: v5, F1: 42}
        v15 := v1[677]
        _ = v14
        _ = v15
    }
    var v16 uint32 = 2
    for fi17 := 0; fi17 < 10; fi17++ {
        switch "a" {
        case "":
            fmt.Println(v16, v16)
        case "a":
            v18 := GenS1{F0: (v5 | v5), F1: (v5 >> (((^(v5))) & 63))}
            _ = v18
        default:
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS2{F0: 1}
                var v21 float64 = 1e10
                v5 += 0
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        v1[(v16 & v16)] = fi17
        _ = fi17
    }
    var v22 string = ("hello" + "hello")
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 uint64 = (^(v2))
        var v25 float64 = 3.14159
        switch v22 {
        case "":
            for fi26 := 0; fi26 < 20; fi26++ {
                v27 := GenS3{F0: uint16(v24)}
                _ = fi26
                _ = v27
            }
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
