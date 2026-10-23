package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 int8 = -1
    v4 := []uint32{2147483648, 2}
    fmt.Println(v3)
    if (4294967295 > 0) {
        delete(v1, 757)
        if false {
            for fi5 := 0; fi5 < 1; fi5++ {
                for fi6 := 0; fi6 < 3; fi6++ {
                    var v7 float64 = (float64(3.14159) + (1.5 / 0.1))
                    fmt.Println(v7, fi5)
                    _ = fi6
                    _ = v7
                }
                _ = fi5
            }
        }
        v8 := GenS1{F0: 1}
        _ = v8
    }
    var v9 int8 = -1
    if false {
        if (true || (0 == -913)) {
            var v10 string = "\n\t"
            v11 := GenS1{F0: 65535}
            v2 = (v2 / ((769) | 1))
            _ = v10
            _ = v11
        }
        v4 = append(v4, 65535)
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(v9, fi12)
        switch 0 {
        case 42:
            if false {
                for fi13 := 0; fi13 < 10; fi13++ {
                    v3 += v9
                    if false {
                        fmt.Println(v2, fi12)
                    }
                    _ = fi13
                }
            }
        case 4294967295:
            fmt.Println(v2)
        default:
            v9 = 0
        }
        var v14 int8 = (v3 - v3)
        _ = fi12
        _ = v14
    }
    var v15 int8 = (v3 + v9)
    _ = len(v4)
    v1[uint32(v2)] = 546
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS1{F0: uint32(v2)}
            v9 = (v3 >> ((65535) & 63))
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v9 -= (-(v3))
    v19 := GenS1{F0: 202}
    switch ("hello" + ("a" + "\n\t")) {
    case "hello":
        for fi20 := 0; fi20 < 2; fi20++ {
            switch "a" {
            case "\n\t":
                var v21 int = (fi20 % ((32767) | 1))
                _ = v21
            }
            if true {
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 uint64 = 4294967295
                    switch fi20 {
                    case 0:
                        fmt.Println(v3, v23)
                    }
                    fmt.Println(fi22, v15, v2)
                    _ = fi22
                    _ = v23
                }
            }
            _ = fi20
        }
    }
    var v24 bool = (!("a" >= "hello"))
    v24 = false
    var v25 bool = v24
    var v26 int32 = 2147483647
    var v27 float64 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
