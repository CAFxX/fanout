package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int8 }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 string) int8 {
    fmt.Println(p1, p1)
    if true {
        p1 = ("\n\t" + ("a" + "\n\t"))
    }
    var v4 uint16 = 32767
    _ = p0
    _ = p1
    _ = v4
    return 44
}

func main() {
    var v1 float32 = float32(-2.25)
    var v2 int16 = -599
    v3 := make(map[uint32]int)
    var v5 uint16 = 312
    var v6 uint32 = 677
    var v7 int32 = -367
    v8 := GenS1{F0: -129, F1: 1, F2: 1}
    switch (("\n\t" + "\n\t") + "a") {
    case "foo bar":
        v1 += v1
    default:
        v5 = (v5 & 42)
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int = (fi9 - fi9)
        v2 -= v2
        v2 = (v2 - v2)
        _ = fi9
        _ = v10
    }
    v3[967] = 797
    var v11 int64 = -417
    if true {
        var v12 uint16 = 608
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        v1 -= -2.25
        v14 := GenS1{F0: v11, F1: 2, F2: -48}
        switch (fi13 ^ (fi13 ^ (^(fi13)))) {
        case 0:
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 int16 = (v2 * 696)
                _ = fi15
                _ = v16
            }
        case -693:
            v17 := []int{0, int(v11), 1, (fi13 * (fi13 - 0))}
            _ = v17
        case 128:
            if true {
                v18 := GenS2{F0: v11, F1: ("hello" + "")}
                _ = len(v3)
                _ = v18
            } else {
                v1 -= v1
            }
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
