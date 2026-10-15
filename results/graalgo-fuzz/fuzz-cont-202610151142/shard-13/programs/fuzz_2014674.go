package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint64; F2 string }
func genFn0(p0 int16) uint16 {
    switch 32767 {
    case 0:
        fmt.Println(p0)
    default:
        var v6 uint32 = 12
        _ = v6
    }
    var v7 uint64 = 0
    _ = p0
    _ = v7
    return 255
}

func main() {
    var v1 string = ("a" + ("" + ""))
    var v2 int8 = -1
    var v3 int32 = 1
    var v4 int = int(v2)
    var v5 float64 = 1e10
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int16 = 1000
        switch (v4 << ((538) & 63)) {
        case -705:
            v10 := []int{}
            _ = v10
        default:
            v5 -= float64((1e10 - 1.5))
        }
        v5 = float64(0.1)
        _ = fi8
        _ = v9
    }
    fmt.Println(v5)
    v1 = v1
    for fi11 := 0; fi11 < 4; fi11++ {
        if true {
            v1 = (v1 + (v1 + v1))
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: (3.14159 == 0.1)}
                switch (v1 + (v1 + (v1 + (v1 + v1)))) {
                case "a":
                    v14 := GenS2{F0: (v2 ^ -30)}
                    _ = v14
                case "foo bar":
                    v15 := GenS3{F0: v2, F1: 1, F2: v1}
                    _ = v15
                default:
                    v5 -= 1.5
                }
                v16 := []int{fi11, fi11, -9223372036854775808}
                v1 = ""
                var v17 float32 = float32(v5)
                _ = fi12
                _ = v13
                _ = v16
                _ = v17
            }
            v3 = (v3 << ((255) & 63))
        } else {
            var v18 bool = (true || (2 != 487))
            _ = v18
        }
        _ = fi11
    }
    var v19 uint64 = 1
    fmt.Println(v19)
    var v20 int = (v4 >> ((994) & 63))
    for fi21 := 0; fi21 < 19; fi21++ {
        v1 = v1
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
