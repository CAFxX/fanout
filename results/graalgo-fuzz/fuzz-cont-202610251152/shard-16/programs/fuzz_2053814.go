package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint16) uint8 {
    if (128 != -1) {
        p0 = p0
        if ("foo bar" > ("a" + "foo bar")) {
            p0 -= 847
        }
    } else {
        p0 += p0
    }
    v5 := []int{-2147483648}
    if (1000 >= -76) {
        v6 := GenS1{F0: (false || false), F1: -502}
        p0 += 255
        _ = v6
    }
    fmt.Println(p0)
    v7 := GenS1{F0: false, F1: 959}
    _ = p0
    _ = v5
    _ = v7
    return 146
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (float64((float32((0.0 * float64(float32(float64(((1e10 - -2.25) * 1e10)))))) + (-2.25 * 0.1))) - -2.25)
    var v3 uint64 = 1
    v4 := make(map[string]int)
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := []int{fi8, fi8, (fi8 << ((255) & 63))}
        fmt.Println(fi8, v3, v2)
        _ = fi8
        _ = v9
    }
    v3 = 943
    v10 := []int{615, -9223372036854775808, -739, 42}
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            v2 = 1e10
            if true {
                v13 := GenS2{F0: 101}
                if false {
                    if false {
                        delete(v1, (fi11 + (fi12 % (((fi12 % (((fi11 & 538)) | 1))) | 1))))
                    }
                    if false {
                        _ = len(v4)
                    } else {
                        if len(v10) > 0 {
                            v10[0] = fi12
                        }
                        if ((v2 / v2) == v2) {
                            fmt.Println(v3, v3)
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v2 += v2
                        } else {
                            v2 -= (v2 - v2)
                        }
                        v2 = v2
                        if true {
                            var v14 int32 = 0
                            v3 -= (^(v3))
                            _ = v14
                        }
                        var v15 int = (fi11 & (fi12 << ((42) & 63)))
                        _ = v15
                    }
                }
                v2 = 1e10
                _ = v13
            }
            _ = fi12
        }
        _ = fi11
    }
    if true {
        var v16 int32 = -2147483648
        _ = v16
    }
    v17 := GenS3{F0: 42}
    delete(v4, (((("hello" + (("foo bar" + (("a" + "foo bar") + "a")) + "foo bar")) + "a") + "a") + "a"))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v18 := GenS3{F0: genFn0(1)}
    for fi19 := 0; fi19 < 16; fi19++ {
        v3 = v3
        v20 := GenS4{F0: 0, F1: -128}
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: ("foo bar" < (("" + "foo bar") + "a")), F1: 0}
    v3 = (v3 >> ((4294967295) & 63))
    var v22 uint16 = 1
    fmt.Println(v3, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
