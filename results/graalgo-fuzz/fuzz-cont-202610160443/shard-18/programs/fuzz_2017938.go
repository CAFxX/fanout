package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int, p1 int32, p2 uint) uint8 {
    fmt.Println(p2, p1, p1)
    if false {
        v6 := GenS1{F0: 1}
        p1 = p1
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    var v3 uint64 = 1
    var v4 int16 = 1
    var v5 uint16 = 0
    var v7 uint32 = 0
    v8 := GenS2{F0: (v2 + -1)}
    for fi9 := 0; fi9 < 3; fi9++ {
        v3 = v3
        _ = fi9
    }
    v7 -= v7
    delete(v1, ("hello" + "\n\t"))
    v10 := GenS3{F0: (v3 / (((^(v3))) | 1)), F1: v3}
    switch int(v5) {
    case 32767:
        var v11 uint16 = v5
        _ = v11
    }
    var v12 bool = true
    var v13 uint = 158
    for fi14 := 0; fi14 < 4; fi14++ {
        v7 += v7
        switch (fi14 & (fi14 * (fi14 & (fi14 - fi14)))) {
        case 0:
            v2 += v2
        case 4294967295:
            var v15 string = "a"
            _ = v15
        default:
            var v16 uint32 = 4294967295
            _ = v16
        }
        _ = fi14
    }
    v4 -= (v4 & (v4 << ((v13) & 63)))
    var v17 int16 = v4
    v12 = (1 <= -1)
    var v18 uint16 = (^(v5))
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 uint16 = (v5 & 212)
        fmt.Println(v4)
        _ = fi19
        _ = v20
    }
    v17 -= (v17 | v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
