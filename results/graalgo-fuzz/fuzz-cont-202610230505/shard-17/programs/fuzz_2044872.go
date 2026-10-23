package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 uint8) int64 {
    v6 := []int{32767, 9223372036854775807, -462, -9223372036854775808}
    var v7 int16 = -189
    for fi8 := 0; fi8 < 13; fi8++ {
        _ = len(v6)
        var v9 float32 = (1.5 - -2.25)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 255, F2: 301}
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 762
    var v2 string = ""
    var v3 uint16 = 32767
    var v4 uint = 42
    var v5 float32 = 3.14159
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: -894, F1: 247, F2: 0}
        v1 += v1
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 int = fi14
            var v16 int8 = 100
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 uint = v4
            v19 := []int{(fi17 >> (((v4 + v4)) & 63)), fi17}
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := []int{(fi13 & fi13)}
        _ = fi13
        _ = v20
    }
    fmt.Println(v3)
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 float32 = v5
        _ = fi21
        _ = v22
    }
    var v23 uint32 = 65535
    for fi24 := 0; fi24 < 4; fi24++ {
        if (841 <= v3) {
            v25 := []int{int(v3), 778, (fi24 ^ 42)}
            if len(v25) > 0 {
                v25[0] = fi24
            }
            _ = v25
        }
        _ = fi24
    }
    v26 := []int{0}
    fmt.Println(v3)
    var v27 uint = v4
    v28 := GenS2{F0: 9223372036854775807, F1: -652}
    var v29 int8 = -113
    var v30 uint = v27
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
