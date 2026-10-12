package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint) bool {
    v5 := []int{-703, 614, 0}
    var v6 uint = p0
    fmt.Println(v6, p0)
    var v7 float32 = -2.25
    v7 -= (float32(-2.25) * float32(float64((-2.25 + v7))))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (930 > 65535)
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint16 = 65535
    var v3 int16 = -32768
    v4 := []int{2147483647, 32768, 128}
    _ = len(v4)
    if ((0 > -9223372036854775808) && (4294967295 >= (v1 >> ((1) & 63)))) {
        fmt.Println(v1, v3, v3)
    } else {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: (false && (true && (3.14159 == float64(0.1))))}
            for fi10 := 0; fi10 < 8; fi10++ {
                if len(v4) > 0 {
                    v4[0] = (^(fi8))
                }
                _ = fi10
            }
            v3 = -157
            _ = fi8
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        _ = len(v4)
        _ = fi11
    }
    v12 := GenS1{F0: (((1e10 / 0.1) <= -2.25) || true)}
    var v13 uint16 = v2
    var v14 uint16 = (v2 ^ 2)
    var v15 string = ("foo bar" + "\n\t")
    if len(v4) > 0 {
        v4[0] = 1
    }
    var v16 uint8 = 2
    v15 = ("hello" + (v15 + (v15 + v15)))
    if true {
        if true {
            v13 -= (v2 / ((v13) | 1))
            if true {
                _ = len(v4)
                switch (v15 + "hello") {
                case "":
                    var v17 int16 = (v3 / (((v3 | v3)) | 1))
                    _ = v17
                default:
                    v15 = v15
                }
            }
        }
        v18 := GenS1{F0: ("a" >= v15)}
        switch ((v15 + "hello") + v15) {
        case "\n\t":
            v13 = 2
        case "":
            var v19 uint = 42
            _ = v19
        case "a":
            v4 = append(v4, 4294967295)
        default:
            fmt.Println(v13, v15, v13)
        }
        _ = v18
    } else {
        for fi20 := 0; fi20 < 3; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 string = ("foo bar" + v15)
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        v2 += v14
    }
    if (987 < (v14 / ((0) | 1))) {
        fmt.Println(v14)
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
