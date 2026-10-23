package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    var v4 int = 1
    _ = len(v2)
    var v5 int = (v4 / ((128) | 1))
    v5 = (v5 * v4)
    var v6 int = (-(v5))
    if (255 != 214) {
        v6 += -276
        var v7 int8 = 0
        _ = v7
    }
    fmt.Println(v4, v6)
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{}
        v10 := GenS1{F0: 580}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch ("foo bar" + "foo bar") {
    case "":
        v4 -= 919
    }
    switch 127 {
    case 300:
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: 1}
            v13 := []int{v4, fi11, v4, -1}
            _ = fi11
            _ = v12
            _ = v13
        }
    case 765:
        if (!(true || true)) {
            v14 := GenS1{F0: 0}
            _ = v14
        } else {
            var v15 uint = 255
            var v16 int8 = -1
            _ = v15
            _ = v16
        }
    }
    delete(v2, (v4 / ((9223372036854775807) | 1)))
    switch ("foo bar" + "hello") {
    case "":
        v6 = 32768
    case "foo bar":
        v5 -= (v6 | v6)
    }
    if true {
        if false {
            v6 = (v6 ^ (v5 << ((2) & 63)))
        }
    } else {
        v5 += (v5 * (v5 ^ v6))
        if (!(1e10 < ((float64(-2.25) * 1e10) + (float64(3.14159) + (-2.25 + (((float64(3.14159) / -2.25) / (0.0 - -2.25)) * 1e10)))))) {
            fmt.Println(v4, v4)
            var v17 bool = true
            v18 := GenS2{F0: -535, F1: 39}
            _ = v17
            _ = v18
        }
        v19 := v3[v4]
        _ = v19
    }
    v20 := GenS1{F0: 0}
    var v21 float32 = (0.0 + 1.5)
    fmt.Println(v6, v4)
    switch 1 {
    case 42:
        var v22 string = "foo bar"
        _ = v22
    case 108:
        var v23 int = -1
        _ = v23
    case -226:
        switch 32767 {
        case 2147483647:
            if (("foo bar" + ("" + ((("hello" + "\n\t") + "hello") + ""))) == "foo bar") {
                if true {
                    v24 := v1["\n\t"]
                    v5 -= 0
                    v25 := []int{v6, v6, v4}
                    _ = v24
                    _ = v25
                }
                v26 := GenS1{F0: 42}
                _ = v26
            }
        case -2147483648:
            var v27 uint16 = 1
            _ = v27
        default:
            var v28 int32 = 2147483647
            _ = v28
        }
    default:
        v6 += 127
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
