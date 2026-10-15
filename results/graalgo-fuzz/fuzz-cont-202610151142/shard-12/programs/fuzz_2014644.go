package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
func genFn0() int8 {
    v5 := []int{}
    var v6 uint16 = 0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int32 = 1
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    if (("a" + ("a" + "")) != "\n\t") {
        var v7 uint = 4294967295
        v2 -= (v2 >> (((v7 | v7)) & 63))
        _ = v7
    } else {
        var v8 uint64 = 4294967295
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                switch ("foo bar" + "\n\t") {
                case "hello":
                    fmt.Println(v1)
                case "":
                    var v11 float64 = float64(-2.25)
                    _ = v11
                default:
                    fmt.Println(fi9)
                }
                v2 = 2147483647
                v8 += v8
                _ = fi10
            }
            var v12 int32 = v2
            _ = fi9
            _ = v12
        }
        _ = v8
    }
    switch 546 {
    case 860:
        v13 := GenS1{F0: 2, F1: 0, F2: ("a" + "")}
        _ = v13
    }
    v2 += (v2 * v2)
    v2 += (v2 | v2)
    v2 = -2147483648
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: 255, F1: 255, F2: "a"}
        _ = fi14
        _ = v15
    }
    fmt.Println(v1, v1)
    switch "hello" {
    case "hello":
        v2 = (v2 % ((814) | 1))
    case "foo bar":
        delete(v3, -505)
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 float64 = ((0.0 + 0.0) / 0.1)
        v2 = v2
        _ = fi16
        _ = v17
    }
    var v18 int64 = -590
    var v19 int = 4294967295
    var v20 float64 = ((-2.25 * (1.5 + 0.1)) / 0.1)
    for fi21 := 0; fi21 < 13; fi21++ {
        if true {
            v22 := GenS1{F0: 255, F1: 1, F2: "a"}
            _ = v22
        }
        v23 := []int{v19, fi21, 9223372036854775807}
        switch "\n\t" {
        case "foo bar":
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 uint32 = (v1 >> ((685) & 63))
                v26 := GenS1{F0: 1, F1: 4294967295, F2: "hello"}
                _ = fi24
                _ = v25
                _ = v26
            }
        case "":
            v23 = append(v23, v19)
        case "a":
            v27 := v3[2147483647]
            _ = v27
        }
        _ = fi21
        _ = v23
    }
    var v28 int8 = 33
    for fi29 := 0; fi29 < 5; fi29++ {
        var v30 uint8 = 128
        for fi31 := 0; fi31 < 5; fi31++ {
            v30 -= v30
            var v32 uint8 = (v30 & v30)
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
}
