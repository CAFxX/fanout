package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 int64) string {
    p0 += p0
    p0 = -9223372036854775808
    _ = p0
    return "hello"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 42
    v3 := make(map[string]uint32)
    var v4 uint64 = 2
    var v5 string = ""
    v6 := make(map[uint32]int)
    var v7 int8 = -7
    v5 = "\n\t"
    if false {
        v5 = v5
        v7 = (v7 % (((v7 % (((v7 >> ((4294967295) & 63))) | 1))) | 1))
    }
    v8 := GenS1{F0: v5}
    var v9 int8 = v7
    var v10 int64 = 1
    switch -9223372036854775808 {
    case -1:
        v11 := GenS2{F0: -1, F1: 42}
        _ = v11
    default:
        for fi12 := 0; fi12 < 12; fi12++ {
            v7 = v9
            v4 = 42
            _ = fi12
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        if false {
            v14 := []int{fi13, fi13, (fi13 + 0)}
            v15 := GenS2{F0: -904, F1: 65535}
            v10 = (v10 - (v10 << ((2) & 63)))
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            if true {
                var v17 uint8 = 1
                var v18 float32 = 3.14159
                _ = v17
                _ = v18
            }
            _ = fi16
        }
        _ = fi13
    }
    if (!(0 < -709)) {
        for fi19 := 0; fi19 < 18; fi19++ {
            v9 += v7
            v20 := []int{fi19, fi19, fi19}
            if ((v10 / ((v10) | 1)) != (v10 - (v10 * 33))) {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            }
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            v4 = v4
            v22 := []int{}
            _ = fi21
            _ = v22
        }
        v5 = v5
    } else {
        v10 = v10
        var v23 uint = 255
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
