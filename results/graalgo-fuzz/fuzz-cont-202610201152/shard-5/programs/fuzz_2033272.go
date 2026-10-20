package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int16; F1 int16; F2 string }

func main() {
    var v1 int8 = -128
    var v2 float64 = -2.25
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 294
    v5 := make(map[string]bool)
    var v6 uint16 = (v4 / (((v4 * (v4 % ((v4) | 1)))) | 1))
    v7 := []int{-551, -587, 59}
    switch -747 {
    case 127:
        v8 := GenS1{F0: v6, F1: 72, F2: (v6 + v4)}
        _ = v8
    case 57:
        var v9 uint16 = (v6 ^ v6)
        _ = v9
    case -517:
        fmt.Println(v4, v6)
    }
    if true {
        var v10 uint16 = 566
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 int8 = v1
            if (!(v3 != (v3 * (v3 | (v3 << ((0) & 63)))))) {
                v13 := GenS2{F0: ("\n\t" + "hello"), F1: v3}
                _ = v13
            } else {
                for fi14 := 0; fi14 < 11; fi14++ {
                    delete(v5, (("" + ("foo bar" + "\n\t")) + "hello"))
                    _ = fi14
                }
            }
            v1 += 86
            _ = fi11
            _ = v12
        }
        fmt.Println(v10)
        _ = v10
    }
    if false {
        if false {
            fmt.Println(v2, v6)
        } else {
            v6 -= 65535
        }
        v15 := GenS2{F0: "\n\t", F1: 42}
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(v3)
        _ = fi16
    }
    switch 32768 {
    case 127:
        fmt.Println(v3)
    default:
        v4 += v6
    }
    var v17 int = 32768
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            v3 = 569
            v20 := GenS3{F0: -1, F1: -32768, F2: "a"}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 uint = 646
            var v23 uint = (v22 | v22)
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
