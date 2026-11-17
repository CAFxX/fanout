package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 bool }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    v4 := []uint32{2, 471, 42, 1}
    switch 42 {
    case 32767:
        var v5 string = (("hello" + ("a" + "foo bar")) + ("\n\t" + ("\n\t" + "foo bar")))
        _ = v5
    case 4294967295:
        switch -115 {
        case 42:
            v1 = (v1 * v1)
        case 2147483647:
            var v6 uint = 1
            _ = v6
        default:
            if false {
                v1 = float64((0.1 * 0.0))
            }
        }
    default:
        if true {
            fmt.Println(v1)
            var v7 uint = 4294967295
            fmt.Println(v1)
            _ = v7
        }
    }
    v8 := GenS1{F0: 75, F1: 2, F2: 42}
    _ = len(v4)
    switch 2147483647 {
    case -9223372036854775808:
        var v9 int64 = -307
        _ = v9
    case -588:
        var v10 float64 = v1
        _ = v10
    default:
        var v11 int64 = 42
        _ = v11
    }
    var v12 float32 = 1.5
    v13 := GenS1{F0: 214, F1: 47, F2: 2147483648}
    var v14 string = ((((("a" + ("hello" + "hello")) + "hello") + "\n\t") + "a") + "foo bar")
    v15 := GenS1{F0: 0, F1: 81, F2: 846}
    if true {
        v16 := GenS2{F0: 255, F1: -90, F2: (48 == 2)}
        _ = v16
    }
    switch v14 {
    case "a":
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 int8 = 78
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 string = (v14 + (v14 + v14))
                var v21 string = (("" + v14) + v14)
                if (255 == 0) {
                    var v22 int8 = (v18 | v18)
                    _ = v22
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            v14 = (v14 + "")
            _ = fi17
            _ = v18
        }
    default:
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS1{F0: 2, F1: 8, F2: 733}
            var v25 int16 = 446
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    v26 := GenS2{F0: 224, F1: -1, F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v8
}
