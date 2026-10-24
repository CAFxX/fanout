package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0(p0 string) int {
    v7 := []int{127}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return -30
}
func genFn1() uint8 {
    v8 := GenS1{F0: 32767, F1: true}
    var v9 int8 = 42
    v9 -= 0
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int8 = 100
        _ = fi10
        _ = v11
    }
    v9 -= (v9 & 85)
    _ = v8
    _ = v9
    return 205
}

func main() {
    var v1 bool = (!true)
    var v2 float32 = 1.5
    v3 := make(map[string]uint32)
    var v4 float64 = 0.0
    var v5 float32 = v2
    v6 := make(map[uint32]uint32)
    v1 = (false && v1)
    var v12 uint8 = 42
    v4 -= -2.25
    fmt.Println(v1, v2, v12)
    var v13 uint = 1
    fmt.Println(v12, v12)
    var v14 uint8 = genFn1()
    switch genFn0("hello") {
    case 32767:
        var v15 int16 = -455
        _ = v15
    default:
        v16 := []int{4294967295, 4294967295, 32767}
        _ = v16
    }
    if true {
        v1 = (475 < 65535)
    } else {
        v17 := GenS1{F0: 2, F1: (9223372036854775807 == 32768)}
        v4 += (v4 + v4)
        _ = v17
    }
    v18 := GenS1{F0: 2, F1: true}
    v13 = (v13 | 255)
    for fi19 := 0; fi19 < 4; fi19++ {
        v13 = uint(v14)
        fmt.Println(v2, v4)
        _ = fi19
    }
    switch -561 {
    case -623:
        if v1 {
            v13 = 73
            var v20 int16 = -32768
            _ = v20
        } else {
            switch 0 {
            case 1:
                var v21 int64 = 1
                _ = v21
            case 128:
                var v22 uint64 = uint64(v13)
                _ = v22
            case -224:
                switch (("foo bar" + "a") + (("foo bar" + "a") + "\n\t")) {
                case "foo bar":
                    for fi23 := 0; fi23 < 2; fi23++ {
                        v14 -= genFn1()
                        var v24 uint32 = 396
                        _ = fi23
                        _ = v24
                    }
                case "a":
                    v6[2147483648] = 25
                default:
                    var v25 float32 = (0.0 - v5)
                    _ = v25
                }
            default:
                var v26 int16 = -41
                _ = v26
            }
        }
    case -1:
        fmt.Println(v13, v5, v5)
    case 32768:
        var v27 int = 42
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
