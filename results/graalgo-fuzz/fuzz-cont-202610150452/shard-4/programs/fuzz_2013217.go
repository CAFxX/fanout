package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    v1 := []uint32{395}
    var v2 int32 = -42
    v3 := []uint32{1, 4294967295}
    var v4 int16 = 1
    v5 := make(map[int]int)
    v6 := GenS1{F0: -553, F1: -773, F2: -9223372036854775808}
    for fi7 := 0; fi7 < 3; fi7++ {
        switch (("a" + (("a" + "\n\t") + "hello")) + (("a" + "\n\t") + "foo bar")) {
        case "":
            for fi8 := 0; fi8 < 9; fi8++ {
                v2 = (v2 | 1)
                _ = fi8
            }
        }
        v9 := GenS2{F0: -9223372036854775808}
        _ = fi7
        _ = v9
    }
    var v10 float32 = 0.1
    v11 := GenS3{F0: v2, F1: -1}
    v10 += v10
    var v12 int8 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 uint8 = 144
    v14 := v5[-1]
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            v2 = (v2 << ((617) & 63))
            var v17 string = ("foo bar" + "foo bar")
            _ = fi16
            _ = v17
        }
        fmt.Println(fi15, v4)
        _ = fi15
    }
    v12 -= v12
    v18 := GenS1{F0: 42, F1: 838, F2: (v14 / ((v14) | 1))}
    v2 += (v2 % ((2147483647) | 1))
    v14 = (v14 | (v14 * (v14 | (v14 / (((v14 / (((-(v14))) | 1))) | 1)))))
    if len(v3) > 0 {
        v3[0] = 65535
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
