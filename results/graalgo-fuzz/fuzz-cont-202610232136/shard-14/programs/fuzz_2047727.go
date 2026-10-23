package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint16 }
func genFn0() int {
    v7 := []int{}
    var v8 int64 = 1
    var v9 int16 = -1
    _ = v7
    _ = v8
    _ = v9
    return -95
}
func genFn1(p0 bool, p1 uint64) int64 {
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := []int{genFn0()}
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi10
        _ = v11
    }
    if (!true) {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: 42, F1: 1}
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 0.1
    var v3 int32 = 0
    var v4 float64 = (((v2 - v2) / 1.5) * v2)
    v5 := make(map[string]uint32)
    v6 := make(map[string]string)
    for fi14 := 0; fi14 < 4; fi14++ {
        if (0 < 0) {
            _ = len(v6)
            v15 := GenS2{F0: genFn0()}
            _ = v15
        } else {
            v16 := GenS1{F0: -1, F1: v3}
            v4 -= v2
            var v17 bool = false
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            if ((0.0 > float32(float64(float32(float64(1.5))))) && true) {
                if (false || (!(4294967295 == 2147483647))) {
                    v4 = float64((0.1 - 0.0))
                    if false {
                        v3 = (v3 | (v3 << ((42) & 63)))
                    }
                } else {
                    if (-790 <= 1) {
                        v19 := GenS2{F0: 2147483647}
                        _ = v19
                    }
                }
                v20 := []int{(^(fi18)), (-(fi18))}
                _ = v20
            } else {
                v3 = v3
                v2 -= (-2.25 + -2.25)
                v4 = float64(1e10)
            }
            var v21 string = "\n\t"
            if (0 >= 37) {
                var v22 string = (v21 + "\n\t")
                v3 = (v3 ^ v3)
                if (v2 != (1e10 - (((v4 * v4) + (1.5 - -2.25)) + 3.14159))) {
                    v2 = float64(3.14159)
                    fmt.Println(v3, fi14, v4)
                }
                _ = v22
            } else {
                v23 := GenS3{F0: 436}
                _ = v23
            }
            v2 = 1.5
            var v24 uint32 = 2147483648
            _ = fi18
            _ = v21
            _ = v24
        }
        _ = fi14
    }
    v25 := []int{904, 32767, genFn0()}
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := GenS4{F0: 2}
        if len(v25) > 0 {
            v25[0] = (fi26 << ((1) & 63))
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v3, v3, v4)
    if ("" > ("" + "a")) {
        if (!(!(true && false))) {
            v28 := GenS5{F0: 4294967295, F1: 42}
            _ = v28
        } else {
            fmt.Println(v3, v2)
        }
        var v29 string = ""
        _ = v29
    }
    _ = len(v25)
    for fi30 := 0; fi30 < 3; fi30++ {
        v6["\n\t"] = ("hello" + "a")
        _ = fi30
    }
    v4 += ((3.14159 * (3.14159 - v4)) * 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
