package main

import (
    "fmt"
)

func genFn0() uint {
    var v6 uint32 = 4294967295
    fmt.Println(v6)
    _ = v6
    return 510
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = float32(1e10)
    var v3 uint8 = 255
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    v2 = (v2 + float32((-2.25 / -2.25)))
    delete(v5, 1)
    v2 -= (1.5 + (0.1 + v2))
    fmt.Println(v2, v2)
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v3)
        v8 := []int{fi7, fi7, fi7, (-(fi7))}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v3 += (v3 << ((genFn0()) & 63))
        if false {
            var v10 int32 = 0
            if true {
                v2 -= v2
            }
            _ = v10
        } else {
            switch ("a" + "\n\t") {
            case "":
                v2 = (0.1 / 1.0)
            }
        }
        _ = fi9
    }
    var v11 uint16 = 793
    var v12 uint64 = 0
    var v13 uint32 = 42
    var v14 uint = 657
    v12 -= (v12 & (v12 ^ v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
