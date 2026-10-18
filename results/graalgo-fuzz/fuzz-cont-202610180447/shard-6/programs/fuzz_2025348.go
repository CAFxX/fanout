package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint8, p1 string, p2 uint) uint8 {
    if true {
        v6 := []int{9223372036854775807, 42}
        _ = v6
    } else {
        var v7 int = -899
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{(^(fi8)), fi8, 2147483647}
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    if ((((false && (p1 <= (p1 + "foo bar"))) && (true && true)) && true) || false) {
        v10 := []int{-351, 134, -1}
        var v11 int8 = -117
        _ = v10
        _ = v11
    }
    var v12 uint16 = 244
    var v13 float64 = (3.14159 * 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return (p0 / ((255) | 1))
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 32767
    var v3 uint = 205
    var v4 uint = v3
    var v5 uint16 = (^(v2))
    var v14 bool = false
    v4 = v3
    v4 = (v4 >> (((v4 | 896)) & 63))
    var v15 float32 = 3.14159
    if (!v14) {
        var v16 int16 = 342
        v4 = v3
        _ = v16
    } else {
        var v17 int64 = 0
        fmt.Println(v14, v15)
        _ = v17
    }
    var v18 uint32 = 4294967295
    if (397 <= -515) {
        _ = len(v1)
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v18 -= v18
        if false {
            var v20 float32 = v15
            for fi21 := 0; fi21 < 16; fi21++ {
                v20 -= 1.5
                _ = len(v1)
                v22 := GenS1{F0: ((((v5 & v2) <= 210) && (593 > 2)) || false)}
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            v23 := []int{-217, -1, (-(fi19))}
            _ = v23
        }
        _ = fi19
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v1[-729] = (fi24 << ((v4) & 63))
        v25 := GenS2{F0: 0}
        var v26 uint8 = genFn0(127, (("foo bar" + ("a" + ("hello" + (("a" + "foo bar") + "hello")))) + ("\n\t" + "a")), v3)
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := []int{fi27, (fi27 % ((32768) | 1))}
        for fi29 := 0; fi29 < 18; fi29++ {
            if len(v28) > 0 {
                v28[0] = (fi29 + (-(fi29)))
            }
            if false {
                v3 += v4
                v30 := GenS1{F0: (-545 < 2147483647)}
                _ = v30
            }
            v5 -= v5
            var v31 int = fi29
            var v32 uint32 = (v18 ^ (v18 | (^(v18))))
            v5 = v5
            var v33 uint8 = 42
            _ = fi29
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
