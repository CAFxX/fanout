package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
func genFn0() int32 {
    var v7 float64 = float64(3.14159)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int8 = 0
        v10 := GenS1{F0: 91, F1: "foo bar"}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 127 {
    case -937:
        v7 = (((float64((0.1 - (3.14159 * -2.25))) * float64(1.5)) - v7) - v7)
    case -9223372036854775808:
        fmt.Println(v7)
    case 530:
        fmt.Println(v7)
    }
    v11 := GenS2{F0: 42}
    v7 = (0.1 + 1e10)
    _ = v11
    _ = v7
    return 42
}
func genFn1() int8 {
    var v12 int16 = -1
    v13 := []int{}
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]uint32)
    var v3 uint32 = v1
    v4 := make(map[string]string)
    var v5 uint32 = (v1 + v1)
    var v6 string = "a"
    for fi14 := 0; fi14 < 16; fi14++ {
        v5 = (^(v3))
        v15 := GenS1{F0: 1, F1: (v6 + ("a" + v6))}
        _ = fi14
        _ = v15
    }
    fmt.Println(v1, v3, v6)
    var v16 int8 = -128
    v2[v6] = v5
    var v17 int64 = -1
    if true {
        if false {
            v18 := []int{-1}
            v5 = v5
            _ = v18
        }
    } else {
        var v19 string = "a"
        switch "\n\t" {
        case "a":
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 int64 = int64(v16)
                v3 = 0
                for fi22 := 0; fi22 < 10; fi22++ {
                    v17 += -1
                    _ = fi22
                }
                _ = fi20
                _ = v21
            }
        case "\n\t":
            var v23 int16 = 0
            _ = v23
        case "hello":
            for fi24 := 0; fi24 < 1; fi24++ {
                delete(v2, ((v19 + v6) + ((v6 + v19) + v19)))
                var v25 bool = (v16 < v16)
                _ = fi24
                _ = v25
            }
        default:
            if (913 == 0) {
                v3 -= (v5 ^ v5)
                switch "\n\t" {
                case "hello":
                    v26 := GenS2{F0: 42}
                    _ = v26
                case "a":
                    var v27 uint32 = 2
                    _ = v27
                case "":
                    switch 316 {
                    case 530:
                        v6 = v6
                    case 1:
                        v16 -= v16
                    default:
                        v3 = (v1 + v3)
                    }
                }
            }
        }
        v3 -= 42
        _ = v19
    }
    v5 = 2
    delete(v4, v6)
    v28 := GenS3{F0: ((((v6 + v6) + ((v6 + v6) + v6)) + "") + "hello")}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
