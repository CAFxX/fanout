package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    var v5 int16 = v1
    v3[384] = (!false)
    var v6 int = int(v5)
    v1 = v5
    fmt.Println(v1, v6, v1)
    delete(v3, 251)
    switch "" {
    case "":
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := GenS1{F0: 144}
            _ = fi7
            _ = v8
        }
    }
    var v9 float32 = 0.0
    v5 = v5
    if false {
        if (255 < 588) {
            if (1.5 >= 0.0) {
                var v10 uint = 42
                v1 = 0
                _ = v10
            }
        }
    } else {
        var v11 int64 = -1
        _ = v11
    }
    fmt.Println(v6, v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
