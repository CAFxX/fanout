package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int32, p1 int64, p2 uint32) uint {
    var v5 int32 = p0
    v6 := GenS1{F0: true}
    var v7 int16 = 741
    v8 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 978
}

func main() {
    v1 := []uint32{4294967295, 104}
    var v2 uint32 = 42
    v3 := make(map[uint32]bool)
    var v4 int32 = -304
    v2 = v2
    var v9 uint16 = 475
    v4 -= v4
    var v10 uint16 = 744
    _ = len(v1)
    v2 += (^(v2))
    var v11 float64 = ((1.5 - ((1.5 + 3.14159) + 0.1)) / 1.5)
    _ = len(v3)
    if len(v1) > 0 {
        v1[0] = 65535
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        if (fi12 >= 32768) {
            switch fi12 {
            case 120:
                v4 += (v4 - (v4 & (v4 | -463)))
            case 32767:
                v2 = v2
            default:
                var v13 float64 = v11
                _ = v13
            }
        }
        _ = fi12
    }
    if (!true) {
        fmt.Println(v9)
        var v14 int16 = 0
        v14 -= v14
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 string = ((("a" + ("foo bar" + "foo bar")) + "hello") + ("\n\t" + "foo bar"))
        _ = fi15
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
