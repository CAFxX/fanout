package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
func genFn0() int {
    var v6 uint32 = 42
    v7 := GenS1{F0: -32768}
    _ = v6
    _ = v7
    return 32768
}
func genFn1() uint16 {
    if (!(42 >= 42)) {
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
        var v9 uint32 = 26
        var v10 int = genFn0()
        _ = v10
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        var v11 bool = true
        _ = v11
    case "a":
        v12 := []int{-1, 127, 42}
        _ = v12
    }
    var v13 int64 = 1
    var v14 float64 = 0.0
    v13 += (v13 >> ((2) & 63))
    _ = v13
    _ = v14
    return 908
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 uint8 = 128
    var v4 float32 = 3.14159
    v5 := make(map[int]int)
    v4 = float32(((-2.25 / float64(-2.25)) - (3.14159 - 1.5)))
    if true {
        v4 -= 1.5
        v15 := []int{-1}
        if len(v15) > 0 {
            v15[0] = 1
        }
        _ = v15
    }
    v16 := GenS1{F0: 173}
    if false {
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS2{F0: 9223372036854775807}
            _ = fi17
            _ = v18
        }
        v3 += (^(v3))
    }
    fmt.Println(v3, v4)
    fmt.Println(v3, v4)
    var v19 string = ("foo bar" + "")
    delete(v1, 127)
    var v20 float64 = 0.0
    switch 42 {
    case -9223372036854775808:
        fmt.Println(v4, v19, v4)
    }
    var v21 bool = (!(-572 < -1))
    fmt.Println(v20, v4)
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 uint8 = 42
        _ = fi22
        _ = v23
    }
    switch genFn0() {
    case 32768:
        var v24 uint32 = 2147483648
        _ = v24
    case 42:
        v25 := []int{genFn0(), 665, genFn0()}
        _ = v25
    case -2147483648:
        fmt.Println(v20, v19, v4)
    default:
        if (18446744073709551615 < 1) {
            switch 9223372036854775807 {
            case -2147483648:
                for fi26 := 0; fi26 < 5; fi26++ {
                    fmt.Println(fi26, v21, v21)
                    var v27 float64 = (3.14159 - 1.5)
                    v20 -= v20
                    _ = fi26
                    _ = v27
                }
            case 128:
                delete(v1, 127)
            case 408:
                v4 -= 0.0
            default:
                v28 := GenS1{F0: 1000}
                _ = v28
            }
        }
    }
    switch 1 {
    case 0:
        fmt.Println(v3, v21)
    case -425:
        v21 = (32768 <= 127)
    case 370:
        if v21 {
            v21 = (float64((1e10 * v4)) <= (-2.25 + 3.14159))
        }
    default:
        for fi29 := 0; fi29 < 7; fi29++ {
            v30 := []int{}
            for fi31 := 0; fi31 < 3; fi31++ {
                if v21 {
                    var v32 uint64 = 0
                    _ = v32
                }
                var v33 int16 = -792
                var v34 int64 = -9223372036854775808
                _ = fi31
                _ = v33
                _ = v34
            }
            _ = fi29
            _ = v30
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
