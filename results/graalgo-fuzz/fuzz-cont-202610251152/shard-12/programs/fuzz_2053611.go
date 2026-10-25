package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint8) bool {
    v5 := []int{-839, -9223372036854775808}
    p0 = 255
    _ = p0
    _ = v5
    return true
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 0
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 += (v2 * v2)
        if false {
            v8 := GenS1{F0: -1}
            fmt.Println(fi7, v6, v6)
            _ = v8
        }
        _ = fi7
    }
    switch "\n\t" {
    case "\n\t":
        v6 = (v6 - v6)
    case "hello":
        var v9 uint32 = 2147483647
        _ = v9
    case "foo bar":
        v6 -= (v6 + uint64(v2))
    }
    v6 = 18446744073709551615
    v10 := GenS2{F0: 255}
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 uint64 = 1
        v12 = 0
        _ = fi11
        _ = v12
    }
    var v13 int64 = 1
    v14 := GenS2{F0: 42}
    for fi15 := 0; fi15 < 20; fi15++ {
        v1[fi15] = fi15
        _ = fi15
    }
    if (0 != 827) {
        delete(v4, 623)
    } else {
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 uint32 = 193
            v18 := GenS3{F0: v17}
            if (!true) {
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 float64 = 3.14159
                    var v21 int = 127
                    v22 := v3[(^(v17))]
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                v4[v17] = v17
                delete(v3, (v17 << ((224) & 63)))
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v3[217] = int(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
