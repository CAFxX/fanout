package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint8) uint32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{fi6}
        fmt.Println(p0, fi6)
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 float64 = 0.1
            v10 := GenS1{F0: 2}
            if (-1 < -1) {
                if false {
                    v11 := GenS1{F0: uint64(p1)}
                    _ = v11
                } else {
                    fmt.Println(v9, p0, fi6)
                }
                p0 = p0
                var v12 uint = 42
                _ = v12
            } else {
                fmt.Println(p1, fi8, fi8)
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v13 uint = 0
    _ = p0
    _ = p1
    _ = v13
    return (p0 ^ p0)
}

func main() {
    v1 := []string{"foo bar", "foo bar", "foo bar", "foo bar"}
    v2 := make(map[int]string)
    var v3 int32 = -300
    var v4 uint32 = 65535
    v5 := []uint32{v4}
    var v14 int64 = -1
    v15 := GenS2{F0: 32767}
    v14 = 0
    fmt.Println(v14, v4)
    v4 += (v4 << ((861) & 63))
    fmt.Println(v4)
    v16 := GenS2{F0: 1}
    var v17 int16 = 1000
    var v18 int32 = 1
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 float32 = (0.1 + (3.14159 / 0.1))
        v17 += v17
        var v21 int64 = v14
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 float64 = 3.14159
    fmt.Println(v18)
    switch "\n\t" {
    case "":
        fmt.Println(v3, v14, v4)
    case "a":
        switch ("a" + "a") {
        case "foo bar":
            fmt.Println(v18, v18)
        }
    default:
        var v23 int = -2147483648
        _ = v23
    }
    v17 += (^(v17))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
