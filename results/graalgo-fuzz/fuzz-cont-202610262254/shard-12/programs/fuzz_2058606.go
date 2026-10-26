package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 int8 }
func genFn0(p0 uint16, p1 uint, p2 uint64) int16 {
    if false {
        p1 -= (p1 | 64)
        p0 = p0
    } else {
        v7 := []int{2147483647, -579}
        _ = v7
    }
    v8 := []int{42, 128}
    p2 -= 2
    switch ("\n\t" + ("hello" + "")) {
    case "hello":
        switch "a" {
        case "":
            var v9 string = ("foo bar" + "hello")
            _ = v9
        default:
            fmt.Println(p2, p0, p2)
        }
    case "\n\t":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        v8 = append(v8, -912)
    }
    p1 = (p1 >> (((p1 / ((p1) | 1))) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 32767
}

func main() {
    v1 := []uint32{723}
    var v2 uint32 = 672
    var v3 uint32 = (^(v2))
    var v4 uint = 42
    var v5 int = 51
    var v6 float64 = float64(1e10)
    var v10 int32 = -2147483648
    v11 := GenS1{F0: v4, F1: -114, F2: 1000}
    v10 = v10
    if len(v1) > 0 {
        v1[0] = v3
    }
    v12 := GenS2{F0: 413, F1: (^(v4)), F2: v3}
    var v13 string = ""
    v6 = ((v6 - (((v6 * float64((1e10 * (0.1 - 1.5)))) / 1.0) + float64(1.5))) + (v6 * float64(0.0)))
    for fi14 := 0; fi14 < 6; fi14++ {
        if len(v1) > 0 {
            v1[0] = (v3 | 2147483648)
        }
        _ = fi14
    }
    var v15 uint64 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch v5 {
    case 4294967295:
        for fi16 := 0; fi16 < 20; fi16++ {
            v4 += v4
            switch (^(fi16)) {
            case 4294967295:
                v15 = (v15 & 1)
            case 127:
                for fi17 := 0; fi17 < 20; fi17++ {
                    _ = len(v1)
                    var v18 int16 = -1
                    _ = fi17
                    _ = v18
                }
            default:
                var v19 int64 = 9223372036854775807
                _ = v19
            }
            v20 := GenS3{F0: fi16, F1: 26, F2: 0}
            _ = fi16
            _ = v20
        }
    case 42:
        var v21 uint16 = 42
        _ = v21
    default:
        fmt.Println(v15, v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
