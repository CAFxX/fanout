package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint16; F2 int8 }
func genFn0() int16 {
    var v6 int64 = 0
    var v7 int32 = -2147483648
    v8 := []int{-9223372036854775808, 886, -9223372036854775808}
    _ = v6
    _ = v7
    _ = v8
    return 1000
}
func genFn1() bool {
    var v9 int = 128
    v10 := []int{558, v9, (^(v9)), 0}
    _ = v10
    _ = v9
    return (508 >= 4294967295)
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    v5 := make(map[string]int)
    switch -116 {
    case 0:
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS1{F0: 1000}
            _ = fi11
            _ = v12
        }
    default:
        fmt.Println(v1)
    }
    if (false && true) {
        v13 := GenS1{F0: 251}
        switch "hello" {
        case "\n\t":
            v1 -= (^(v1))
        default:
            var v14 int16 = genFn0()
            _ = v14
        }
        _ = v13
    } else {
        v15 := GenS1{F0: 1000}
        v1 -= -1
        _ = v15
    }
    v16 := []int{42, 32767, 303, -2147483648}
    v1 += v1
    v17 := GenS1{F0: 1000}
    if false {
        v1 = -170
        for fi18 := 0; fi18 < 13; fi18++ {
            if (false && true) {
                if (true && false) {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        var v20 float32 = 0.0
                        v1 = (v1 / ((44) | 1))
                        _ = fi19
                        _ = v20
                    }
                    v5["hello"] = 32767
                    var v21 float64 = 0.0
                    _ = v21
                } else {
                    if (!true) {
                        var v22 uint = 1
                        fmt.Println(v1)
                        _ = v22
                    } else {
                        var v23 uint64 = 42
                        v24 := GenS2{F0: 58}
                        fmt.Println(v23)
                        _ = v23
                        _ = v24
                    }
                }
                var v25 uint16 = 1
                if (true && true) {
                    v1 += v1
                }
                _ = v25
            }
            _ = fi18
        }
        var v26 float64 = ((-2.25 - (3.14159 * ((1e10 - float64((0.1 * float32(3.14159)))) * 0.1))) / 1.0)
        _ = v26
    }
    v1 -= v1
    var v27 uint16 = 2
    v27 = (v27 << ((65535) & 63))
    var v28 int64 = 42
    v29 := GenS3{F0: ("\n\t" + "\n\t"), F1: 32767, F2: 127}
    var v30 int8 = 117
    fmt.Println(v1, v28)
    switch "a" {
    case "foo bar":
        var v31 uint16 = (v27 - 1)
        _ = v31
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    v28 = -503
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
