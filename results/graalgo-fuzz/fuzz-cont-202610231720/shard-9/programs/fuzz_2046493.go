package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int64; F1 int8; F2 uint8 }
func genFn0(p0 uint8) uint32 {
    v7 := []int{int(p0), 127}
    var v8 uint64 = 855
    p0 = 1
    for fi9 := 0; fi9 < 20; fi9++ {
        v7 = append(v7, (fi9 << ((1) & 63)))
        for fi10 := 0; fi10 < 7; fi10++ {
            if ((v8 + 1) <= v8) {
                if false {
                    var v11 float32 = (0.0 / -2.25)
                    _ = v11
                } else {
                    v7 = append(v7, 2147483647)
                }
            }
            _ = fi10
        }
        _ = fi9
    }
    v12 := GenS1{F0: 32768, F1: 42}
    _ = p0
    _ = v12
    _ = v7
    _ = v8
    return 541
}

func main() {
    var v1 int8 = -1
    var v2 int32 = 42
    var v3 string = "a"
    var v4 uint16 = 877
    var v5 int64 = 1
    var v6 float32 = float32(-2.25)
    fmt.Println(v3, v2, v1)
    var v13 uint = 388
    var v14 bool = true
    for fi15 := 0; fi15 < 5; fi15++ {
        v2 += (v2 % ((-483) | 1))
        v16 := []int{fi15}
        v4 += (v4 - (^(v4)))
        _ = fi15
        _ = v16
    }
    var v17 uint = (^(v13))
    var v18 int16 = 0
    var v19 int64 = 0
    for fi20 := 0; fi20 < 15; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 int64 = v5
            v3 = (((((("a" + "hello") + v3) + "hello") + ("\n\t" + v3)) + (v3 + (v3 + (v3 + "\n\t")))) + (v3 + v3))
            if v14 {
                v13 += v13
                var v23 uint8 = 73
                var v24 float32 = v6
                _ = v23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        v25 := []int{}
        _ = fi20
        _ = v25
    }
    fmt.Println(v6, v2)
    v14 = ((v18 >> ((v17) & 63)) > (-(v18)))
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 uint = v17
        v14 = true
        _ = fi26
        _ = v27
    }
    switch (v3 + v3) {
    case "\n\t":
        v28 := GenS2{F0: (v19 - v5), F1: -128, F2: 255}
        _ = v28
    }
    fmt.Println(v17, v5, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
