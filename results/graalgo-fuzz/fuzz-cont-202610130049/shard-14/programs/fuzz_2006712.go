package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 uint16, p1 uint, p2 uint32) uint8 {
    var v5 int8 = -34
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 200
}

func main() {
    var v1 uint16 = 2
    var v2 string = ("hello" + "")
    var v3 float32 = 0.1
    var v4 uint16 = 0
    for fi7 := 0; fi7 < 2; fi7++ {
        if true {
            var v8 uint16 = (^(v4))
            v2 = (v2 + "foo bar")
            var v9 uint64 = 2
            _ = v8
            _ = v9
        } else {
            if true {
                switch (-(fi7)) {
                case 128:
                    var v10 float64 = (1.5 * 3.14159)
                    _ = v10
                case 693:
                    if false {
                        var v11 int32 = 383
                        var v12 float32 = float32((float64(1e10) - 1e10))
                        _ = v11
                        _ = v12
                    }
                }
            } else {
                v13 := []int{}
                _ = v13
            }
        }
        var v14 uint32 = 0
        _ = fi7
        _ = v14
    }
    var v15 bool = (!(((v2 + "a") + "a") >= v2))
    var v16 float64 = 1.5
    switch -837 {
    case 32768:
        fmt.Println(v16)
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint8 = 42
        if false {
            if (v4 != v4) {
                v19 := GenS1{F0: v2, F1: -1}
                _ = v19
            }
        }
        switch (fi17 * (fi17 | -1)) {
        case 858:
            v18 += v18
        }
        _ = fi17
        _ = v18
    }
    var v20 int64 = -9223372036854775808
    fmt.Println(v2)
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v4)
        v1 = (v1 | (v1 | v1))
        _ = fi21
    }
    v15 = v15
    if ((-1 <= 127) || ((1 != 132) && true)) {
        var v22 int32 = 2147483647
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
