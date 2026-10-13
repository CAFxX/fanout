package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint8 = 0
    v2 := []int{42}
    var v3 int32 = 2147483647
    var v4 float64 = -2.25
    v5 := make(map[uint32]string)
    fmt.Println(v4)
    fmt.Println(v1)
    var v6 int = 707
    if (true || true) {
        if false {
            for fi7 := 0; fi7 < 3; fi7++ {
                switch (v6 - int(v1)) {
                case -9223372036854775808:
                    var v8 uint64 = 567
                    _ = v8
                }
                var v9 uint8 = 64
                v10 := GenS1{F0: 2}
                _ = fi7
                _ = v10
                _ = v9
            }
            var v11 uint8 = v1
            _ = v11
        } else {
            fmt.Println(v4, v3)
            v12 := GenS1{F0: 4294967295}
            _ = v12
        }
    } else {
        v4 = float64(float32((v4 * v4)))
        v4 = v4
        _ = len(v2)
    }
    v5[uint32(v1)] = ("hello" + (("a" + "a") + ("a" + "foo bar")))
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint64 = 4294967295
        v6 = v6
        _ = fi13
        _ = v14
    }
    if len(v2) > 0 {
        v2[0] = v6
    }
    v3 += v3
    for fi15 := 0; fi15 < 17; fi15++ {
        v2 = append(v2, (fi15 * (^(v6))))
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
