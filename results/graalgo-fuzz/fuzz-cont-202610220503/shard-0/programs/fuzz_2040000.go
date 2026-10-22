package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint32 }
func genFn0() uint8 {
    var v7 uint16 = 255
    v8 := GenS1{F0: -81}
    if true {
        if ((0 == 42) || (true && true)) {
            var v9 int16 = 42
            _ = v9
        } else {
            v7 = (^(v7))
            fmt.Println(v7)
            var v10 uint32 = 65535
            _ = v10
        }
        var v11 int16 = 32767
        _ = v11
    }
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint32 = 2
    var v2 uint64 = 1
    var v3 uint32 = v1
    v4 := make(map[string]int)
    var v5 int16 = 42
    var v6 uint = 0
    var v12 uint16 = 65535
    v13 := GenS2{F0: 42, F1: -223}
    _ = len(v4)
    v1 = (v1 >> ((v6) & 63))
    v14 := GenS3{F0: 9223372036854775807, F1: (^(v6))}
    delete(v4, "\n\t")
    fmt.Println(v12, v6)
    var v15 uint = v6
    fmt.Println(v3, v2)
    var v16 uint = uint(v2)
    v17 := GenS2{F0: 2147483647, F1: -112}
    v1 = (^(v1))
    v18 := GenS4{F0: (v3 & v3)}
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := []int{fi19}
        if (true && ((v15 | 1) < v15)) {
            v3 += v1
            v20 = append(v20, fi19)
            v21 := v4["foo bar"]
            _ = v21
        } else {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 uint32 = v3
                v1 = (^(v23))
                _ = fi22
                _ = v23
            }
            fmt.Println(v2, v3)
        }
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
