package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 int8 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []string{(("" + ("a" + "foo bar")) + "hello"), ("\n\t" + "\n\t"), "a"}
    v2 := make(map[uint32]uint32)
    var v3 bool = ("" == "hello")
    var v4 int32 = 42
    for fi5 := 0; fi5 < 5; fi5++ {
        v3 = (886 >= 255)
        _ = fi5
    }
    v6 := GenS1{F0: true, F1: 2147483648, F2: 0}
    v7 := GenS2{F0: 100, F1: -1, F2: 42}
    switch "" {
    case "":
        if v3 {
            var v8 float32 = ((0.0 + 1.5) * float32(0.0))
            var v9 uint16 = 132
            _ = v8
            _ = v9
        } else {
            v4 += v4
        }
    default:
        for fi10 := 0; fi10 < 15; fi10++ {
            if v3 {
                v4 = v4
                var v11 bool = v3
                _ = v11
            } else {
                var v12 uint16 = 32767
                v3 = v3
                _ = v12
            }
            var v13 uint16 = 954
            v14 := GenS1{F0: v3, F1: 42, F2: v13}
            _ = fi10
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        switch "\n\t" {
        case "":
            fmt.Println(v4, v4)
        case "\n\t":
            v16 := GenS2{F0: -99, F1: (-(v4)), F2: 0}
            _ = v16
        case "a":
            v4 -= (-(v4))
        }
        v17 := GenS3{F0: v4}
        _ = fi15
        _ = v17
    }
    v3 = (688 > 32767)
    switch 291 {
    case -820:
        if ((3.14159 == 0.1) || false) {
            var v18 uint16 = 998
            _ = v18
        }
    }
    v19 := GenS3{F0: 1}
    switch 32768 {
    case -1:
        v4 = v4
    case 42:
        for fi20 := 0; fi20 < 3; fi20++ {
            v2[42] = 2147483648
            var v21 uint8 = 227
            _ = fi20
            _ = v21
        }
    default:
        fmt.Println(v4, v3)
    }
    v22 := GenS3{F0: 1}
    var v23 uint = 805
    v24 := v2[4294967295]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
