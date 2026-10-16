package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8, p1 string, p2 int32) int16 {
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := []int{-498, (fi4 << ((65535) & 63))}
        var v6 uint16 = 657
        p2 -= (^(p2))
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(p1, p2)
        switch fi7 {
        case 4294967295:
            fmt.Println(p2, fi7)
        default:
            p2 = -1
        }
        _ = fi7
    }
    p2 += p2
    var v8 int8 = 88
    var v9 string = (((p1 + "hello") + p1) + p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint64 = (v1 % ((v1) | 1))
    var v10 int8 = 100
    fmt.Println(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        fmt.Println(v10)
        if false {
            v1 = (v1 + v3)
        } else {
            var v12 float64 = 3.14159
            _ = v12
        }
        _ = fi11
    }
    v1 = (v3 >> ((255) & 63))
    v13 := GenS1{F0: 1}
    switch ("" + "\n\t") {
    case "a":
        v14 := GenS2{F0: 6}
        _ = v14
    case "hello":
        for fi15 := 0; fi15 < 8; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                v10 = (v10 | -1)
                _ = fi16
            }
            _ = fi15
        }
    case "foo bar":
        var v17 float32 = float32((0.1 * (0.0 / float64((((1e10 * 0.1) / 1.0) + float32(float64((((-2.25 + 0.0) / ((0.1 + 0.1) / float32(1.5))) * (-2.25 + float32(float64(0.1)))))))))))
        _ = v17
    }
    v1 = (v3 / ((v1) | 1))
    v10 -= 8
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
}
