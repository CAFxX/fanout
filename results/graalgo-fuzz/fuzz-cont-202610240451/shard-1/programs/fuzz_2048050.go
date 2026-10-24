package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 string; F1 int }
func genFn0(p0 int64, p1 int8) string {
    p0 -= 563
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float64 = -2.25
        _ = fi7
        _ = v8
    }
    if false {
        p0 += p0
    }
    v9 := []int{227}
    _ = p0
    _ = p1
    _ = v9
    return ("hello" + "hello")
}
func genFn1(p0 int64, p1 uint32) int32 {
    v10 := GenS1{F0: 255, F1: 255}
    if true {
        for fi11 := 0; fi11 < 3; fi11++ {
            p1 = (^(p1))
            p1 += p1
            var v12 uint32 = 774
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: -994, F1: 65535}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 bool = true
        v16 = (42 >= 0)
        var v17 uint32 = (p1 << ((2) & 63))
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    var v1 uint = 255
    var v2 int16 = 84
    v3 := make(map[string]uint32)
    var v4 uint = v1
    var v5 uint64 = 1
    var v6 int8 = 35
    v18 := GenS2{F0: "a", F1: -2147483648}
    v5 -= 18446744073709551615
    var v19 int16 = (v2 ^ v2)
    var v20 uint16 = 32767
    var v21 bool = true
    v20 += 32767
    v5 = (^(v5))
    if (42 > -2147483648) {
        v22 := []int{int(v4), 42}
        _ = v22
    } else {
        if (true && v21) {
            v23 := GenS2{F0: "", F1: 42}
            v21 = v21
            var v24 int = -1
            _ = v23
            _ = v24
        }
    }
    v25 := []int{2147483647, 4294967295, -1, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
