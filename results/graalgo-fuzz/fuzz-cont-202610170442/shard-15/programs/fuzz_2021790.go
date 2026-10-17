package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0() int16 {
    var v6 bool = (false && (("\n\t" == "hello") || false))
    v6 = ((v6 || (18446744073709551615 == 18446744073709551615)) || ((!(!true)) || true))
    for fi7 := 0; fi7 < 3; fi7++ {
        if v6 {
            v8 := []int{}
            if (!v6) {
                v6 = true
                v6 = false
            }
            var v9 int64 = 42
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 65, F1: 553}
        v6 = (!false)
        _ = fi7
        _ = v10
    }
    _ = v6
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 1000
    var v3 string = (("\n\t" + "") + "hello")
    var v4 int64 = -1
    var v5 uint16 = 32767
    fmt.Println(v4, v3, v5)
    v2 -= v2
    _ = len(v1)
    for fi11 := 0; fi11 < 5; fi11++ {
        _ = len(v1)
        _ = fi11
    }
    var v12 bool = true
    if false {
        v5 += v5
    } else {
        var v13 int32 = 42
        v2 = v2
        _ = v13
    }
    v1[2147483648] = v12
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 float64 = 0.1
        for fi16 := 0; fi16 < 16; fi16++ {
            v4 += v4
            v17 := GenS2{F0: 1}
            v15 = 1.5
            _ = fi16
            _ = v17
        }
        var v18 uint64 = 326
        _ = fi14
        _ = v15
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
