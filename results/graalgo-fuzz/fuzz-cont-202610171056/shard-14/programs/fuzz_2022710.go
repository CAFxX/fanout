package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint64; F2 int8 }
func genFn0(p0 int16, p1 int64) uint {
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: 1}
        _ = fi6
        _ = v7
    }
    p0 -= p0
    var v8 uint32 = 2
    v9 := GenS2{F0: true}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint16) int {
    fmt.Println(p0)
    if (2 > p0) {
        v10 := []int{-2147483648}
        switch "" {
        case "foo bar":
            p0 -= (p0 * p0)
        case "a":
            var v11 uint32 = 2147483647
            _ = v11
        default:
            var v12 int = 32767
            _ = v12
        }
        _ = v10
    }
    _ = p0
    return 974
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 9223372036854775807
    var v3 int32 = 0
    var v4 int8 = 42
    var v5 uint8 = 0
    switch (("hello" + "hello") + "a") {
    case "":
        v3 -= (v3 / ((-164) | 1))
    case "a":
        v1[1] = 65535
    case "\n\t":
        delete(v1, 0)
    default:
        v13 := GenS3{F0: 32767, F1: 4294967295, F2: v4}
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS1{F0: 324}
        v3 = -1
        _ = fi14
        _ = v15
    }
    v2 = v2
    var v16 bool = false
    switch "" {
    case "hello":
        var v17 uint8 = v5
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v16 = false
        delete(v1, 419)
        _ = fi18
    }
    delete(v1, 42)
    v19 := GenS3{F0: 4294967295, F1: 4294967295, F2: v4}
    for fi20 := 0; fi20 < 19; fi20++ {
        if true {
            v2 = (v2 * v2)
            if ("foo bar" < "a") {
                fmt.Println(v4, fi20, v3)
                v16 = (((fi20 / (((fi20 % ((int(v5)) | 1))) | 1)) > (fi20 | genFn1(912))) || (false || true))
                for fi21 := 0; fi21 < 10; fi21++ {
                    delete(v1, 1)
                    if ((v5 - v5) != v5) {
                        var v22 float32 = (1.5 - (float32((0.0 - (1.5 + ((float64(1.5) / 1.5) * float64(((1.5 + -2.25) * ((1.5 * 0.1) * -2.25))))))) / 1.5))
                        v23 := GenS1{F0: 657}
                        _ = v22
                        _ = v23
                    }
                    _ = len(v1)
                    _ = fi21
                }
            } else {
                if (!(v2 < v2)) {
                    v4 += -1
                } else {
                    fmt.Println(v3)
                    v16 = v16
                }
            }
            v3 += v3
        }
        var v24 float64 = -2.25
        v1[65535] = 2
        _ = fi20
        _ = v24
    }
    var v25 uint = 42
    v5 = (v5 % ((220) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
