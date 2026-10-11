package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []int{-1, 128}
    v2 := []uint32{2, 0, 42, 168}
    v3 := []string{"hello"}
    v4 := make(map[int]bool)
    var v5 int = 32767
    switch (v5 << ((4294967295) & 63)) {
    case 9223372036854775807:
        v5 += v5
    case -84:
        fmt.Println(v5)
    case -58:
        v5 += (v5 >> ((255) & 63))
    }
    v5 += 664
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        switch (-(v5)) {
        case 4294967295:
            var v8 int = v5
            _ = v8
        case -9223372036854775808:
            fmt.Println(v5, v5, v6)
        case 127:
            var v9 int = v5
            _ = v9
        }
        _ = fi7
    }
    var v10 string = "\n\t"
    v4[-787] = (0.0 <= 3.14159)
    if (581 <= (^(v5))) {
        switch (v5 / (((v5 | v5)) | 1)) {
        case 9223372036854775807:
            if (1 >= 65535) {
                var v11 uint8 = (v6 / (((v6 ^ 2)) | 1))
                for fi12 := 0; fi12 < 7; fi12++ {
                    var v13 float32 = 1.5
                    _ = fi12
                    _ = v13
                }
                _ = v11
            } else {
                var v14 int32 = -955
                v15 := GenS1{F0: 127}
                _ = v14
                _ = v15
            }
        default:
            _ = len(v1)
        }
        fmt.Println(v6, v5, v10)
        v10 = "foo bar"
    }
    var v16 int32 = int32(v6)
    var v17 uint8 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
