package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 int32, p1 int32, p2 string) uint64 {
    p1 = p0
    v6 := GenS1{F0: -390, F1: -448}
    p1 += p0
    v7 := []int{127}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 263
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = (0.0 + 0.0)
    var v3 uint8 = 2
    var v4 bool = (!true)
    var v5 uint = 1
    v4 = v4
    switch ("\n\t" + "") {
    case "hello":
        for fi8 := 0; fi8 < 12; fi8++ {
            if v4 {
                v9 := []int{(fi8 << ((548) & 63)), (fi8 - fi8), fi8}
                _ = v9
            }
            _ = fi8
        }
    case "a":
        v3 = 0
    }
    v5 = (^(v5))
    v5 = (v5 << ((v5) & 63))
    fmt.Println(v4)
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := []int{fi11, (fi10 % (((fi10 + (fi10 | (fi11 >> (((v5 << (((v5 * 42)) & 63))) & 63))))) | 1)), fi10, 499}
            _ = fi11
            _ = v12
        }
        v13 := []int{(^(fi10)), fi10, fi10}
        _ = fi10
        _ = v13
    }
    v4 = (105 <= 32767)
    v4 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
