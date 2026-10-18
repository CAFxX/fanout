package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
func genFn0() int8 {
    v6 := GenS1{F0: 42}
    v7 := GenS2{F0: 1}
    _ = v6
    _ = v7
    return -14
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 9
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 uint64 = 0
    v2 += (v2 >> ((685) & 63))
    v2 = 722
    v8 := GenS3{F0: 2147483647}
    var v9 uint = 2
    if (true || true) {
        v3[865] = "a"
        v10 := []int{128, -9223372036854775808, -9223372036854775808}
        _ = v10
    }
    v2 = (v2 | v2)
    delete(v4, 811)
    v11 := v3[9223372036854775807]
    var v12 int8 = -128
    v5 += v5
    if (true && (v5 <= 0)) {
        v13 := GenS3{F0: 719}
        _ = v13
    } else {
        v9 -= (^(v9))
    }
    v9 -= (v9 + 0)
    if true {
        fmt.Println(v9, v11, v11)
        v14 := []int{}
        for fi15 := 0; fi15 < 7; fi15++ {
            fmt.Println(v9, v5, v2)
            v16 := GenS2{F0: 713}
            switch fi15 {
            case 1:
                v17 := GenS4{F0: 9223372036854775807}
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    if false {
        fmt.Println(v11)
        var v18 float32 = 1e10
        var v19 uint = (v9 << (((v9 | 884)) & 63))
        _ = v18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
