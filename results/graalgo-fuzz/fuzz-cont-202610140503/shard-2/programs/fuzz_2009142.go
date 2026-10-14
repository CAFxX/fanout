package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int32) bool {
    fmt.Println(p0, p0)
    if true {
        switch 42 {
        case -1:
            v6 := GenS1{F0: 93, F1: ""}
            _ = v6
        default:
            v7 := GenS1{F0: 42, F1: ""}
            _ = v7
        }
        fmt.Println(p1, p0)
        p1 += -2147483648
    } else {
        v8 := []int{}
        if len(v8) > 0 {
            v8[0] = 32767
        }
        _ = v8
    }
    v9 := GenS1{F0: 2, F1: "\n\t"}
    _ = p0
    _ = p1
    _ = v9
    return false
}
func genFn1(p0 int8, p1 uint8) uint16 {
    p1 = p1
    p1 -= 107
    if false {
        if false {
            var v10 uint8 = p1
            _ = v10
        }
        p0 = 79
        switch int(p0) {
        case -2147483648:
            var v11 uint32 = 65535
            _ = v11
        case 855:
            if false {
                v12 := []int{32768, 32768}
                _ = v12
            }
        default:
            if false {
                v13 := []int{}
                var v14 uint = 713
                p1 = (p1 << (((v14 & (^(v14)))) & 63))
                _ = v13
                _ = v14
            }
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 65535
    var v2 int = 42
    var v3 int64 = 712
    v4 := make(map[uint32]string)
    var v5 uint32 = v1
    var v15 float32 = -2.25
    var v16 uint16 = 39
    var v17 uint8 = 127
    for fi18 := 0; fi18 < 2; fi18++ {
        v15 -= float32(1.5)
        v1 -= (^(v1))
        var v19 int16 = 42
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: uint8(v16), F1: (("a" + (("" + "") + ("foo bar" + ("hello" + "a")))) + "foo bar")}
    var v21 float64 = float64(v15)
    v5 -= (v5 ^ (v1 << ((0) & 63)))
    v4[v1] = ("\n\t" + "a")
    v3 += 42
    v22 := GenS1{F0: (v17 / (((v17 * (v17 / ((0) | 1)))) | 1)), F1: ("" + ("a" + "\n\t"))}
    delete(v4, v5)
    for fi23 := 0; fi23 < 10; fi23++ {
        switch "foo bar" {
        case "a":
            for fi24 := 0; fi24 < 6; fi24++ {
                v3 = (v3 << ((568) & 63))
                v21 -= (v21 / v21)
                v25 := GenS2{F0: (^(v16))}
                _ = fi24
                _ = v25
            }
        }
        v26 := v4[(^(v5))]
        var v27 bool = false
        _ = fi23
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
