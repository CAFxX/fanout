package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint64, p2 bool) int8 {
    var v6 uint64 = 371
    var v7 uint16 = 0
    if p2 {
        fmt.Println(p1, p1)
    } else {
        v8 := []int{-500, 826}
        fmt.Println(p1)
        fmt.Println(v7, v7)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 bool = false
    v3 := make(map[int]int)
    var v4 uint64 = 23
    var v5 int32 = 1
    if v2 {
        v9 := []int{-2147483648}
        _ = v9
    }
    var v10 string = ("a" + "foo bar")
    for fi11 := 0; fi11 < 7; fi11++ {
        v5 = v5
        _ = fi11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := []int{(fi12 ^ 42), 100}
        _ = fi12
        _ = v13
    }
    var v14 int32 = v5
    var v15 float64 = float64(float32(0.1))
    v16 := GenS1{F0: (v10 + v10), F1: 1, F2: v2}
    switch 4294967295 {
    case 2147483647:
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(v4, v14)
            var v18 int8 = 1
            _ = fi17
            _ = v18
        }
    }
    var v19 int8 = 100
    v19 = genFn0((v19 / ((v19) | 1)), v4, (3.14159 > float32(float64(3.14159))))
    if v2 {
        switch -802 {
        case 9223372036854775807:
            var v20 int16 = v1
            _ = v20
        }
    }
    if v2 {
        if v2 {
            for fi21 := 0; fi21 < 3; fi21++ {
                fmt.Println(v5, v15, v10)
                _ = fi21
            }
            v4 = (v4 >> ((42) & 63))
        }
    } else {
        if true {
            var v22 float32 = 1e10
            var v23 int32 = (^(v5))
            _ = v22
            _ = v23
        } else {
            for fi24 := 0; fi24 < 13; fi24++ {
                fmt.Println(v14, v14, v19)
                for fi25 := 0; fi25 < 12; fi25++ {
                    if (v14 <= (v14 ^ (^(v14)))) {
                        var v26 int16 = int16(v4)
                        var v27 int32 = -2147483648
                        _ = v26
                        _ = v27
                    } else {
                        fmt.Println(v1)
                        v1 -= (v1 + (v1 ^ v1))
                    }
                    v28 := []int{(fi25 - fi25)}
                    v29 := GenS1{F0: (v10 + (v10 + ("hello" + "a"))), F1: 2, F2: ((float32(float64((-2.25 + -2.25))) * -2.25) <= ((-2.25 * 0.1) * 1.5))}
                    _ = fi25
                    _ = v28
                    _ = v29
                }
                _ = fi24
            }
            if (!(286 >= 1)) {
                fmt.Println(v19)
            }
        }
    }
    v30 := GenS2{F0: (v4 ^ (v4 << ((4294967295) & 63)))}
    fmt.Println(v14, v14, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
