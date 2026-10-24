package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int64 = 779
    var v3 uint32 = (v1 >> ((0) & 63))
    v4 := []uint32{380, (v1 ^ v1), (^(v1))}
    v1 += v3
    var v5 uint32 = (^(v3))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    v2 -= v2
    for fi6 := 0; fi6 < 7; fi6++ {
        if false {
            v3 = (v3 >> ((1) & 63))
            if len(v4) > 0 {
                v4[0] = v3
            }
        } else {
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 uint8 = 2
                _ = fi7
                _ = v8
            }
        }
        _ = fi6
    }
    v3 = v3
    var v9 uint = 42
    var v10 int64 = 9223372036854775807
    if (140 < 42) {
        if ((1 >= -566) && true) {
            var v11 uint32 = (^(v1))
            v3 += (v11 >> ((uint(v5)) & 63))
            _ = v11
        } else {
            v4 = append(v4, v5)
            v3 = 65535
        }
        fmt.Println(v5)
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int = 32768
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: v3}
            v16 := GenS1{F0: (v3 & v1)}
            if false {
                v13 = v13
                if len(v4) > 0 {
                    v4[0] = 2
                }
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
