package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{318, 4294967295, 65535, 2}
    v3 := []int{4294967295, -2147483648, -767, 9223372036854775807}
    var v4 uint8 = 0
    var v5 int32 = 42
    v4 += v4
    var v6 int32 = v5
    for fi7 := 0; fi7 < 5; fi7++ {
        if true {
            var v8 uint64 = 776
            _ = v8
        } else {
            fmt.Println(v5)
            for fi9 := 0; fi9 < 11; fi9++ {
                switch "" {
                case "":
                    var v10 uint = 4294967295
                    _ = v10
                }
                var v11 int64 = 882
                var v12 uint64 = 0
                _ = fi9
                _ = v11
                _ = v12
            }
        }
        switch -2147483648 {
        case 4294967295:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v6 = (^(v5))
        _ = fi7
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    if (!true) {
        v6 = v5
        v13 := GenS1{F0: 113, F1: 2, F2: -128}
        var v14 uint32 = 2
        _ = v13
        _ = v14
    }
    v6 -= (-(v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
