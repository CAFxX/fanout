package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
func genFn0() int16 {
    var v6 int16 = -191
    v6 += v6
    v6 -= -707
    _ = v6
    return -1
}

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 65535
    var v3 uint32 = v2
    var v4 int64 = 42
    var v5 float64 = v1
    var v7 bool = false
    for fi8 := 0; fi8 < 17; fi8++ {
        v5 = v5
        v4 = v4
        _ = fi8
    }
    var v9 uint8 = 56
    var v10 uint8 = 2
    v9 = 52
    fmt.Println(v4)
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 uint8 = v10
            if (false || false) {
                var v14 int = fi11
                var v15 uint8 = 250
                _ = v14
                _ = v15
            } else {
                var v16 string = ("\n\t" + "hello")
                if true {
                    var v17 bool = (v2 != 2)
                    v16 = v16
                    var v18 uint = 65535
                    _ = v17
                    _ = v18
                }
                _ = v16
            }
            fmt.Println(v9, fi12, v4)
            _ = fi12
            _ = v13
        }
        v10 += v9
        _ = fi11
    }
    var v19 uint32 = (^(v3))
    v20 := []int{}
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := GenS1{F0: fi21, F1: 65535, F2: 2147483647}
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 4294967295, F1: 4294967295, F2: (v3 >> ((2) & 63))}
    v4 = 0
    v4 = v4
    for fi24 := 0; fi24 < 9; fi24++ {
        v9 -= (v10 << ((4294967295) & 63))
        v3 = v3
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
