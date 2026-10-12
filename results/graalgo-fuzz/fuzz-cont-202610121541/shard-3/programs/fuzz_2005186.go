package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 uint, p1 bool, p2 bool) int {
    var v7 uint8 = 5
    if true {
        fmt.Println(p0)
        var v8 string = ""
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -147
}

func main() {
    var v1 uint16 = 65535
    var v2 float64 = -2.25
    v3 := []int{}
    v4 := make(map[uint32]string)
    v5 := []uint32{65535}
    var v6 float64 = v2
    v6 -= ((v2 + (((-2.25 - 1.5) * v2) * ((((v6 * v6) - (1e10 * 1e10)) * (v2 + v6)) * v2))) / -2.25)
    v9 := GenS1{F0: 193, F1: 19, F2: 42}
    if true {
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(fi10, v6, v2)
            fmt.Println(fi10, v1, v2)
            v6 += v2
            _ = fi10
        }
    } else {
        v6 = 1e10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v6 -= float64(1.5)
            v2 = v6
            fmt.Println(v2)
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println(v1)
    var v13 bool = false
    var v14 int = 32767
    var v15 uint16 = (v1 | 32767)
    v13 = v13
    delete(v4, 2)
    v16 := GenS2{F0: 255, F1: "hello"}
    var v17 int32 = 1
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 bool = ((42 <= 18446744073709551615) && ((42 > 2) || v13))
        v4[4294967295] = ("\n\t" + ("" + "\n\t"))
        _ = fi18
        _ = v19
    }
    var v20 uint8 = 110
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
