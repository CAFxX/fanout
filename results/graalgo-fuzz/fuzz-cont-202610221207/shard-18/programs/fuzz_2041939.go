package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint32, p1 uint32, p2 bool) int32 {
    p1 -= (^(p0))
    var v7 bool = (!false)
    p0 += 2147483648
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 int16 = 946
    var v4 float64 = 3.14159
    var v5 int16 = 32767
    var v6 int32 = 1
    v6 -= v6
    var v9 int64 = 42
    fmt.Println(v9)
    if (0 < 127) {
        switch "hello" {
        case "":
            v3 += (v5 + (^(v3)))
        case "hello":
            v9 = (v9 - 1)
        case "foo bar":
            if true {
                v5 += 1000
            } else {
                v9 = (v9 >> ((255) & 63))
            }
        }
        fmt.Println(v5, v3)
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := v1[(^(fi10))]
            v9 -= v9
            var v12 int16 = (v3 / (((v3 | 42)) | 1))
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v6, v3, v9)
    for fi13 := 0; fi13 < 1; fi13++ {
        if false {
            v14 := v1[fi13]
            _ = v14
        } else {
            for fi15 := 0; fi15 < 2; fi15++ {
                v6 -= (v6 % (((v6 | (v6 - 260))) | 1))
                v5 -= v3
                _ = fi15
            }
            v16 := GenS1{F0: 42}
            _ = v16
        }
        v17 := GenS1{F0: 42}
        var v18 int16 = 1
        _ = fi13
        _ = v17
        _ = v18
    }
    v9 -= v9
    v9 = -9223372036854775808
    v19 := GenS1{F0: 0}
    if (1e10 == (1.5 * ((3.14159 / float32(3.14159)) - 1.5))) {
        v4 = 0.1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
