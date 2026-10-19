package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string) uint64 {
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint = 4294967295
        p0 = ("foo bar" + ("\n\t" + p0))
        p0 = ("foo bar" + ((("a" + p0) + p0) + p0))
        _ = fi5
        _ = v6
    }
    p0 = "\n\t"
    v7 := GenS1{F0: true}
    v8 := GenS2{F0: 190}
    v9 := []int{1, 0, 128, 9223372036854775807}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int32) uint32 {
    var v10 uint = 79
    v11 := []int{514, 32767, 127, int(p0)}
    switch 42 {
    case 1:
        var v12 uint8 = 128
        _ = v12
    case 32768:
        p0 = (p0 << ((v10) & 63))
    case 2147483647:
        v13 := GenS2{F0: p0}
        _ = v13
    }
    v14 := GenS1{F0: false}
    var v15 uint64 = 0
    _ = p0
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    return 2147483648
}

func main() {
    var v1 uint8 = 1
    var v2 float32 = 0.1
    v3 := make(map[string]uint32)
    var v4 float32 = 1e10
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := []int{fi17, fi17}
            v1 = (v1 % ((255) | 1))
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 4; fi19++ {
            switch "" {
            case "a":
                v20 := GenS2{F0: -270}
                _ = v20
            case "foo bar":
                var v21 uint32 = 1
                v22 := []int{fi16, fi16}
                v3["hello"] = (v21 - v21)
                v2 = (((v2 - v4) / 0.1) - 3.14159)
                v1 += v1
                _ = v21
                _ = v22
            }
            fmt.Println(fi19, fi16, fi19)
            if (!(1 != 9223372036854775807)) {
                v1 -= v1
                v1 = 2
            }
            _ = fi19
        }
        v2 = ((v4 - v4) + v2)
        _ = fi16
    }
    var v23 float64 = 3.14159
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(v2, v4)
        _ = fi24
    }
    v1 += (v1 | 2)
    v2 = 0.0
    v25 := GenS3{F0: 42}
    v26 := []int{0, 476, 2147483647}
    v1 = v1
    _ = len(v3)
    var v27 uint64 = 42
    v3["hello"] = genFn1(-1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
