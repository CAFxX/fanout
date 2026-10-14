package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 uint32, p2 int16) uint16 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint32 = (p1 | 4294967295)
        p1 += p1
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 40, F1: 0, F2: 128}
    fmt.Println(p1, p1, p0)
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 65535
}

func main() {
    v1 := []int{-374}
    v2 := []string{"hello", "hello", "hello", "foo bar"}
    v3 := []uint32{}
    var v4 uint = 0
    var v5 string = "a"
    v9 := GenS1{F0: 100, F1: -32768, F2: 0}
    var v10 uint = (v4 / (((v4 << ((42) & 63))) | 1))
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            v10 += v10
            if false {
                fmt.Println(v5)
            } else {
                if false {
                    fmt.Println(v10, v4)
                    var v13 uint8 = 1
                    _ = v13
                } else {
                    fmt.Println(v5, v5, v5)
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                }
                if (!true) {
                    fmt.Println(v5, v5)
                    fmt.Println(v10)
                }
                var v14 int8 = 1
                _ = v14
            }
            var v15 int8 = 51
            _ = fi12
            _ = v15
        }
        _ = fi11
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 int16 = 1000
        if false {
            for fi18 := 0; fi18 < 16; fi18++ {
                if true {
                    v5 = v5
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    fmt.Println(fi18, v4)
                }
                v17 = v17
                v19 := GenS2{F0: (-(v17))}
                _ = fi18
                _ = v19
            }
            var v20 int = fi16
            _ = v20
        }
        var v21 uint16 = genFn0(1, 55, v17)
        _ = fi16
        _ = v17
        _ = v21
    }
    v10 -= v10
    var v22 bool = (!true)
    var v23 int = -474
    v24 := GenS1{F0: 127, F1: 1000, F2: (v23 << ((42) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
