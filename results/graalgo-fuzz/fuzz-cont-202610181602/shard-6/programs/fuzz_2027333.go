package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int32 }
func genFn0(p0 int32, p1 uint64, p2 int32) uint16 {
    var v4 float32 = 0.1
    var v5 uint32 = 1
    var v6 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 42
    var v3 string = ("\n\t" + "")
    var v7 uint16 = 65535
    var v8 int = 1
    v3 = ("hello" + v3)
    v9 := []int{}
    v8 += (v8 / ((-1) | 1))
    if false {
        v8 += (v8 ^ -2147483648)
        fmt.Println(v7, v8)
        fmt.Println(v2, v3)
    } else {
        v10 := GenS1{F0: v3, F1: v2, F2: (v2 - (v2 * v2))}
        _ = v10
    }
    v7 = genFn0(v1, (v2 * v2), 42)
    if len(v9) > 0 {
        v9[0] = v8
    }
    if false {
        if false {
            if true {
                if (-1 <= -128) {
                    v2 = v2
                }
                var v11 int = (v8 + (^(v8)))
                _ = v11
            }
            v12 := GenS1{F0: (v3 + "\n\t"), F1: (v2 * v2), F2: v2}
            if (1 != 428) {
                v13 := GenS1{F0: v3, F1: v2, F2: 135}
                _ = len(v9)
                var v14 uint32 = 2147483647
                _ = v13
                _ = v14
            } else {
                v15 := GenS1{F0: v3, F1: (v2 - (^(v2))), F2: (v2 + 953)}
                for fi16 := 0; fi16 < 16; fi16++ {
                    v1 = v1
                    v8 = -105
                    v17 := GenS2{F0: v7, F1: (v1 & v1)}
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            _ = v12
        } else {
            switch "foo bar" {
            case "hello":
                fmt.Println(v1, v2, v7)
            }
            v7 -= 581
            var v18 bool = false
            _ = v18
        }
        _ = len(v9)
        var v19 bool = false
        _ = v19
    }
    v1 = (v1 >> ((255) & 63))
    var v20 uint32 = 125
    if len(v9) > 0 {
        v9[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
