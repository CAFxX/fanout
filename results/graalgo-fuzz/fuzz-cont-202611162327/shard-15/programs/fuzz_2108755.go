package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 string; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{4294967295, 634}
    var v3 int32 = 893
    var v4 float32 = ((1.5 - (float32(float64(float32(1e10))) * 0.0)) * (0.0 / 1.0))
    v5 := make(map[int]bool)
    v4 += float32(((((1.5 + 1.5) + 3.14159) * 0.1) - -2.25))
    fmt.Println(v3, v3)
    fmt.Println(v3, v3)
    if (42 != 4294967295) {
        var v6 int32 = (^(v3))
        _ = v6
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                switch (fi7 ^ (fi8 * fi8)) {
                case 127:
                    fmt.Println(fi7, fi7, v4)
                case 2147483647:
                    v9 := GenS1{F0: 0, F1: 1, F2: 94}
                    _ = v9
                case 128:
                    v3 += v3
                }
                if ((32767 != 65535) || true) {
                    var v10 uint8 = 255
                    v11 := v1[2147483648]
                    _ = v10
                    _ = v11
                } else {
                    v12 := GenS2{F0: 65535}
                    var v13 uint16 = 42
                    var v14 uint8 = 128
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 uint32 = 160
                _ = fi8
                _ = v15
            }
            _ = fi7
        }
        v4 -= v4
    }
    var v16 uint = 65535
    v17 := GenS2{F0: 0}
    if true {
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 int32 = v3
                _ = fi19
                _ = v20
            }
            v16 = 18446744073709551615
            _ = fi18
        }
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            _ = len(v1)
        }
    } else {
        v3 = 1
        v4 = v4
        fmt.Println(v16)
    }
    v2 = append(v2, 637)
    switch 2147483647 {
    case 4294967295:
        v21 := GenS3{F0: v16, F1: -1}
        _ = v21
    }
    v4 -= v4
    v22 := GenS4{F0: "hello", F1: 99, F2: 47}
    delete(v5, -581)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
