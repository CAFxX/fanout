package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint8 = 127
    var v2 uint = 255
    v3 := make(map[int]uint32)
    var v4 uint8 = v1
    v5 := make(map[uint32]uint32)
    var v6 int32 = 970
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int16 = 612
        _ = fi7
        _ = v8
    }
    if (true && false) {
        var v9 uint64 = 511
        fmt.Println(v4, v1, v2)
        if ((v6 - (v6 - (v6 ^ v6))) < v6) {
            fmt.Println(v2)
            v10 := []int{}
            _ = v10
        } else {
            v4 -= (v1 & v4)
            v11 := GenS1{F0: ("" + "a")}
            v2 = v2
            _ = v11
        }
        _ = v9
    }
    v2 = v2
    var v12 uint16 = uint16(v1)
    var v13 uint16 = 32767
    var v14 string = ""
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(v1, v4)
        v6 -= -310
        delete(v3, (fi15 - fi15))
        _ = fi15
    }
    v1 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
