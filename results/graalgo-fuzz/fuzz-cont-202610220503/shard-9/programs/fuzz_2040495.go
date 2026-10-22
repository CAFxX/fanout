package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []string{"hello"}
    var v2 int = 9223372036854775807
    var v3 float64 = 3.14159
    var v4 uint16 = 255
    v5 := GenS1{F0: 0, F1: 820}
    var v6 uint64 = 423
    var v7 uint64 = 530
    v3 -= 3.14159
    v8 := GenS2{F0: 2}
    v7 = 18446744073709551615
    v2 = v2
    if (2147483647 >= 42) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            fmt.Println(v4)
            if len(v1) > 0 {
                v1[0] = ""
            }
            v1 = append(v1, "hello")
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v4 = v4
        }
        var v9 string = ("hello" + "")
        _ = v9
    } else {
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS3{F0: 42}
            switch "\n\t" {
            case "\n\t":
                v3 += v3
            default:
                switch "" {
                case "foo bar":
                    var v12 float32 = float32((v3 / v3))
                    _ = v12
                }
            }
            _ = fi10
            _ = v11
        }
    }
    switch (-(v2)) {
    case 767:
        if (636 == 991) {
            fmt.Println(v2, v4)
            v7 -= 1
        }
    case 0:
        switch 32768 {
        case -311:
            for fi13 := 0; fi13 < 19; fi13++ {
                fmt.Println(v7, v2)
                v14 := GenS1{F0: -112, F1: 1}
                _ = fi13
                _ = v14
            }
        default:
            var v15 uint32 = uint32(v4)
            _ = v15
        }
    }
    var v16 int = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
