package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int32, p1 uint8, p2 int8) int64 {
    v6 := []int{-283, -2147483648}
    var v7 float32 = 3.14159
    var v8 uint = 818
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -984
}
func genFn1(p0 uint64) string {
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = (p0 >> ((127) & 63))
    _ = p0
    return ((("a" + (("\n\t" + "foo bar") + (("a" + "\n\t") + ""))) + "") + "hello")
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 718
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v5 int = 32767
    if (0.0 <= -2.25) {
        var v9 int8 = 100
        fmt.Println(v5, v9)
        _ = v9
    }
    delete(v3, (v5 ^ (-(v5))))
    v5 = v5
    for fi10 := 0; fi10 < 7; fi10++ {
        v2 = (v2 >> ((2) & 63))
        switch (v5 ^ (^(v5))) {
        case -1:
            if true {
                for fi11 := 0; fi11 < 20; fi11++ {
                    fmt.Println(v1)
                    _ = fi11
                }
            } else {
                var v12 int64 = genFn0(-1, 42, v1)
                var v13 float64 = 0.1
                v14 := []int{(-(v5))}
                _ = v12
                _ = v13
                _ = v14
            }
        case 4294967295:
            _ = len(v4)
        case 32767:
            var v15 int32 = 0
            _ = v15
        default:
            if (0.1 != float32(0.1)) {
                fmt.Println(v5, v2)
                fmt.Println(v1, v5)
                if (false && true) {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        v5 = 127
                        var v17 int16 = 42
                        _ = fi16
                        _ = v17
                    }
                }
            } else {
                var v18 uint16 = 42
                _ = v18
            }
        }
        v2 -= (v2 << ((4294967295) & 63))
        _ = fi10
    }
    v1 += v1
    _ = len(v4)
    var v19 float64 = (((-2.25 / (((1.5 - 0.0) / ((0.1 + -2.25) - 1.5)) + -2.25)) + 0.1) + 3.14159)
    var v20 int = v5
    v21 := GenS1{F0: (("hello" > "foo bar") || false)}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
