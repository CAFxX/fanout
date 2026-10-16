package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 int64 }
func genFn0(p0 int64, p1 int32, p2 uint) string {
    v7 := GenS1{F0: "foo bar", F1: (^(p1))}
    var v8 int16 = -1
    if false {
        v9 := GenS2{F0: int64(p2), F1: 0}
        _ = v9
    } else {
        v8 += (v8 << ((p2) & 63))
    }
    v10 := []int{127}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    return "a"
}

func main() {
    var v1 uint32 = 221
    var v2 float64 = -2.25
    var v3 uint16 = 255
    v4 := make(map[uint32]int)
    var v5 uint8 = 255
    v6 := make(map[string]bool)
    if (-413 != 32767) {
        v1 = v1
        v11 := GenS1{F0: "\n\t", F1: -1}
        var v12 uint64 = 0
        _ = v11
        _ = v12
    }
    switch ("\n\t" + "a") {
    case "foo bar":
        v13 := GenS3{F0: 0, F1: "hello", F2: v5}
        _ = v13
    }
    var v14 int16 = 1
    v5 = (v5 >> ((4294967295) & 63))
    v5 = v5
    var v15 int = 127
    if true {
        var v16 uint8 = 2
        _ = v16
    }
    v17 := GenS4{F0: 28, F1: (^(v5)), F2: 0}
    v3 -= (v3 % (((^(v3))) | 1))
    v14 += 0
    fmt.Println(v2, v15)
    for fi18 := 0; fi18 < 9; fi18++ {
        v1 = (v1 / ((v1) | 1))
        _ = fi18
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
