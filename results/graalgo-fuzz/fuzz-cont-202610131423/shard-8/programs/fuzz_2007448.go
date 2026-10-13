package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint16, p1 int16, p2 uint8) uint {
    var v7 string = "hello"
    if true {
        for fi8 := 0; fi8 < 13; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := []int{fi9, fi9}
                p0 = (^(p0))
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(p2, p0)
            v7 = v7
            fmt.Println(p2)
            _ = fi11
        }
    } else {
        v12 := GenS1{F0: 2}
        _ = v12
    }
    p0 += (p0 + 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 907
}

func main() {
    var v1 string = ((((((("hello" + "") + ("hello" + ("\n\t" + "a"))) + (("hello" + "foo bar") + "foo bar")) + "") + "a") + "foo bar") + "hello")
    var v2 int8 = -29
    var v3 bool = (true && (!(138 < 1)))
    var v4 uint = 1
    var v5 float32 = (0.1 - (0.0 * 1e10))
    v6 := make(map[int]uint32)
    v13 := GenS1{F0: 697}
    var v14 int64 = 428
    var v15 int32 = 42
    switch v1 {
    case "\n\t":
        for fi16 := 0; fi16 < 3; fi16++ {
            if false {
                v14 = 335
            }
            if v3 {
                var v17 float32 = v5
                _ = v17
            }
            _ = fi16
        }
    case "a":
        fmt.Println(v15)
    }
    v2 = (^(v2))
    var v18 uint8 = 254
    v15 = (v15 << (((^(v4))) & 63))
    var v19 float64 = (0.0 - ((0.0 * 3.14159) + 1.5))
    v18 += v18
    var v20 string = v1
    var v21 uint8 = (v18 + 127)
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 bool = false
        _ = fi22
        _ = v23
    }
    v24 := GenS2{F0: int16(v14), F1: 2}
    v25 := []int{128, 0, 333}
    var v26 int = int(v2)
    fmt.Println(v4, v2)
    v26 -= -978
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
