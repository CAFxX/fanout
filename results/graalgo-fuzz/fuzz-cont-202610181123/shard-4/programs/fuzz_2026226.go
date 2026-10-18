package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 int64, p1 int16, p2 int64) uint16 {
    p0 = (p2 % ((p0) | 1))
    if false {
        if (259 < 374) {
            v4 := []int{432, 1, 42}
            var v5 uint16 = 825
            _ = v4
            _ = v5
        } else {
            var v6 uint16 = 255
            var v7 string = "a"
            _ = v6
            _ = v7
        }
        var v8 bool = (377 == 4294967295)
        _ = v8
    }
    if false {
        v9 := GenS1{F0: 656}
        v10 := GenS2{F0: 42, F1: 2}
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: (p0 + 9223372036854775807), F1: 63}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 255
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]int)
    v3 := []string{("\n\t" + ("foo bar" + "")), "\n\t", "a", "a"}
    var v12 uint8 = 127
    var v13 int64 = -9223372036854775808
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 uint = 255
    v15 := GenS1{F0: 9}
    v16 := GenS3{F0: ("hello" + ("" + "hello")), F1: genFn0(9223372036854775807, 774, (v13 % (((v13 % (((^(v13))) | 1))) | 1)))}
    var v17 uint16 = 255
    var v18 uint32 = 0
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint = v14
        fmt.Println(v14, v20)
        _ = fi19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
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
}
