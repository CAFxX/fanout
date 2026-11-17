package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
func genFn0(p0 uint8) bool {
    v7 := []int{4294967295, 337, 4294967295}
    p0 += p0
    p0 = 128
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    p0 = p0
    _ = p0
    _ = v7
    return (-1 < -1)
}

func main() {
    v1 := []string{"foo bar", ("foo bar" + ""), "\n\t"}
    var v2 int32 = 2147483647
    var v3 uint8 = 1
    var v4 int16 = -1
    v5 := make(map[uint32]uint32)
    var v6 bool = (!false)
    _ = len(v1)
    v3 = (v3 - 49)
    delete(v5, 1)
    v6 = v6
    var v9 string = "\n\t"
    var v10 uint32 = 911
    for fi11 := 0; fi11 < 10; fi11++ {
        v10 = 787
        _ = fi11
    }
    v12 := GenS1{F0: 1, F1: v10, F2: -9223372036854775808}
    fmt.Println(v10)
    var v13 int = 0
    var v14 int64 = 0
    v1 = append(v1, "\n\t")
    v14 = -744
    var v15 int8 = -1
    v10 = (v10 << ((148) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
