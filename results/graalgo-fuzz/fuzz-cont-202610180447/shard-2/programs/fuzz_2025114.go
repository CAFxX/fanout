package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 uint32, p1 uint8) int32 {
    fmt.Println(p1)
    var v7 int8 = 100
    v8 := []int{2147483647, 4294967295}
    var v9 int16 = 32767
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 733
}
func genFn1(p0 uint64, p1 uint8) int32 {
    p0 += 50
    v10 := []int{127}
    switch 9223372036854775807 {
    case 4294967295:
        p0 = (p0 * p0)
    default:
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: (114 <= 1), F1: (-128 >= 9), F2: -1}
            if true {
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 string = ""
                    v14 = v14
                    v15 := GenS1{F0: (-1 > -32768), F1: true, F2: -1}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            _ = fi11
            _ = v12
        }
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS1{F0: true, F1: true, F2: genFn0(1, (^(p1)))}
        _ = fi16
        _ = v17
    }
    p1 += (p1 / ((255) | 1))
    _ = p0
    _ = p1
    _ = v10
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 2
    var v3 int8 = 99
    var v4 string = (("a" + ("hello" + (("hello" + "\n\t") + "foo bar"))) + "hello")
    var v5 bool = true
    v6 := make(map[int]bool)
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 bool = v5
        v20 := GenS1{F0: (221 <= 191), F1: false, F2: -1}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint8 = 88
    switch 32768 {
    case 195:
        fmt.Println(v2)
    case -351:
        for fi22 := 0; fi22 < 11; fi22++ {
            v23 := GenS1{F0: false, F1: false, F2: int32(fi22)}
            v24 := GenS2{F0: 42}
            v25 := GenS2{F0: -9223372036854775808}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
    default:
        v21 = (^(v21))
    }
    v21 = v21
    v26 := GenS3{F0: 538}
    fmt.Println(v21, v21, v21)
    for fi27 := 0; fi27 < 8; fi27++ {
        v4 = v4
        _ = fi27
    }
    fmt.Println(v2, v2, v5)
    v28 := []int{1, 32768, -2147483648, 127}
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint64 = 4294967295
        _ = fi29
        _ = v30
    }
    v31 := GenS4{F0: 42}
    if (2 >= 2147483647) {
        v3 = (v3 ^ (v3 ^ 23))
        var v32 bool = (v21 < (^(v21)))
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
