package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 174
    v2 := make(map[int]int)
    v3 := []int{9223372036854775807, -178, 32768, 2147483647}
    v2[4294967295] = 9223372036854775807
    if len(v3) > 0 {
        v3[0] = 128
    }
    if true {
        v4 := v2[-1]
        _ = v4
    }
    if true {
        for fi5 := 0; fi5 < 4; fi5++ {
            for fi6 := 0; fi6 < 18; fi6++ {
                if true {
                    fmt.Println(v1, fi6)
                    v1 = 9223372036854775807
                    var v7 int = (fi5 << ((65535) & 63))
                    _ = v7
                } else {
                    var v8 uint32 = 65535
                    var v9 uint16 = 633
                    _ = v8
                    _ = v9
                }
                var v10 int = -2147483648
                if (float32((0.0 / 3.14159)) == ((float32(0.1) / 1.0) / (0.1 - (0.1 - -2.25)))) {
                    var v11 int16 = 42
                    _ = v11
                }
                _ = fi6
                _ = v10
            }
            var v12 uint8 = 1
            _ = fi5
            _ = v12
        }
        v1 -= v1
    }
    v1 += v1
    if ((989 <= -1) || true) {
        var v13 float32 = ((0.0 / 1.0) / 1e10)
        var v14 uint32 = 662
        _ = v13
        _ = v14
    }
    v1 = 9223372036854775807
    for fi15 := 0; fi15 < 5; fi15++ {
        v1 -= (v1 + v1)
        if true {
            var v16 float32 = float32(0.0)
            _ = v16
        } else {
            fmt.Println(fi15)
            var v17 uint64 = 2
            _ = v17
        }
        _ = fi15
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
}
