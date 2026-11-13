package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint = 2
    var v2 uint32 = 1
    var v3 int16 = 42
    var v4 uint64 = 0
    var v5 uint16 = 42
    var v6 uint16 = (v5 / ((255) | 1))
    if (v6 < v5) {
        v7 := []int{-524, -903, 42}
        _ = v7
    }
    var v8 int32 = 42
    v9 := []int{-9223372036854775808}
    v10 := GenS1{F0: -128, F1: 4294967295}
    for fi11 := 0; fi11 < 19; fi11++ {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        if true {
            if ((fi11 - int(v2)) < (fi11 << ((966) & 63))) {
                if (("foo bar" + "a") < "a") {
                    fmt.Println(v8)
                }
                var v12 bool = true
                var v13 int64 = 42
                _ = v12
                _ = v13
            }
            if (-9223372036854775808 == 9223372036854775807) {
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS1{F0: -128, F1: (v4 | uint64(v2))}
                    v2 -= 42
                    if (!((^(v6)) < v5)) {
                        _ = len(v9)
                        v16 := GenS2{F0: 127}
                        _ = v16
                    } else {
                        var v17 int = (fi11 | (fi14 << (((v1 << ((255) & 63))) & 63)))
                        v5 = 288
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
            }
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(v8, v5, fi18)
            _ = fi19
        }
        var v20 bool = (int16(v8) < (v3 >> (((v1 & (^(v1)))) & 63)))
        v2 = v2
        _ = fi18
        _ = v20
    }
    if len(v9) > 0 {
        v9[0] = 4294967295
    }
    fmt.Println(v3, v3)
    fmt.Println(v3, v4)
    v5 -= uint16(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
