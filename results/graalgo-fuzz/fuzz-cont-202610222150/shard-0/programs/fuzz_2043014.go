package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint16) string {
    p0 += p0
    v7 := GenS1{F0: 657}
    switch 9223372036854775807 {
    case 127:
        if (-128 > 1) {
            var v8 uint16 = p0
            v9 := []int{}
            _ = v8
            _ = v9
        }
    case 32767:
        fmt.Println(p0)
    case 2147483647:
        fmt.Println(p0)
    }
    var v10 int32 = 0
    _ = p0
    _ = v10
    _ = v7
    return "hello"
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.0
    var v3 float32 = 1e10
    var v4 float64 = ((1e10 - 1e10) * 0.1)
    v5 := make(map[uint32]bool)
    v6 := make(map[string]string)
    if (!(!(448 > -778))) {
        var v11 uint16 = 255
        if false {
            if false {
                v2 = v3
                v3 += (1e10 + (v3 + (float32(v4) * 0.1)))
            } else {
                v12 := GenS2{F0: 0, F1: 18446744073709551615}
                var v13 int8 = 117
                v14 := GenS1{F0: 4294967295}
                _ = v12
                _ = v13
                _ = v14
            }
            if true {
                var v15 int = -9223372036854775808
                _ = v15
            }
        }
        _ = v11
    } else {
        delete(v1, -1)
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := []int{(-(fi16)), fi16, (fi16 & (-(fi16))), 2147483647}
            v4 -= v4
            _ = fi16
            _ = v17
        }
        var v18 uint32 = 2
        _ = v18
    }
    var v19 float64 = (v4 / ((v4 + v4) / (v4 + (float64(1e10) + 3.14159))))
    v20 := GenS2{F0: 42, F1: 463}
    delete(v1, 2147483647)
    switch -1 {
    case 0:
        v2 = (v2 * v3)
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        if false {
            if false {
                v22 := []int{(^(fi21)), fi21, (-(fi21))}
                _ = v22
            }
        } else {
            switch ("" + "a") {
            case "":
                v23 := GenS1{F0: 609}
                _ = v23
            }
        }
        fmt.Println(v3, v2, v19)
        v24 := []int{}
        _ = fi21
        _ = v24
    }
    v19 -= v4
    v19 = (float64(float32(v19)) * 1.5)
    v6[(("hello" + "") + "a")] = ""
    fmt.Println(v4, v2, v3)
    for fi25 := 0; fi25 < 12; fi25++ {
        v19 -= float64(v2)
        fmt.Println(v4)
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
