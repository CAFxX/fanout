package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 string }
type GenS2 struct { F0 int64; F1 uint; F2 int8 }
func genFn0() int64 {
    var v6 int = 32767
    var v7 uint = 65535
    var v8 int = (^(v6))
    v9 := GenS1{F0: 4294967295, F1: 0, F2: "\n\t"}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -670
}

func main() {
    var v1 int = 128
    var v2 uint8 = 2
    var v3 uint = 51
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    v10 := GenS2{F0: 9223372036854775807, F1: (v3 | (v3 % ((65535) | 1))), F2: -37}
    v11 := GenS1{F0: 522, F1: (v2 / ((v2) | 1)), F2: ("a" + "a")}
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 string = ("foo bar" + "")
        v14 := GenS2{F0: -9223372036854775808, F1: v3, F2: -117}
        v15 := []int{(-(fi12)), -1, (fi12 - (v1 / ((128) | 1)))}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v2 -= v2
    if true {
        v2 = v2
        var v16 uint32 = 570
        fmt.Println(v16, v3)
        _ = v16
    }
    var v17 uint64 = uint64(v3)
    fmt.Println(v2)
    v17 -= (^(v17))
    var v18 float64 = 1.5
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 int32 = 1
        var v21 uint32 = 65535
        var v22 string = (("" + ("foo bar" + "a")) + "a")
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    v1 = (v1 + (v1 - v1))
    v17 = (v17 << ((29) & 63))
    v18 = v18
    v2 -= v2
    v18 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
