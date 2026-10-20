package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint8) bool {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int = (fi5 % (((fi5 >> ((65535) & 63))) | 1))
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{(fi7 + 541), 9223372036854775807, (fi7 * (^(fi7)))}
        var v9 bool = true
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = (float64(float32(1.5)) * (1e10 - (((0.0 * 1.5) + 0.1) + 1.5)))
    v11 := []int{25, 32767, 32768, -9223372036854775808}
    _ = p0
    _ = v10
    _ = v11
    return true
}

func main() {
    v1 := []uint32{84, 0}
    var v2 uint64 = 0
    var v3 uint = 42
    v4 := make(map[string]uint32)
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v12 float32 = (1.5 * 0.0)
    fmt.Println(v3, v12)
    delete(v4, ("\n\t" + "a"))
    v4["foo bar"] = 446
    delete(v4, "hello")
    v2 += v2
    switch "foo bar" {
    case "hello":
        v13 := GenS1{F0: 1000, F1: 255}
        _ = v13
    case "foo bar":
        v14 := GenS1{F0: 0, F1: 82}
        _ = v14
    case "a":
        switch 32767 {
        case 32768:
            v3 += 30
        case 32767:
            for fi15 := 0; fi15 < 13; fi15++ {
                switch (("foo bar" + "hello") + "\n\t") {
                case "foo bar":
                    v3 -= (v3 << ((v3) & 63))
                }
                _ = fi15
            }
        default:
            var v16 uint64 = v2
            _ = v16
        }
    }
    if (!(503 <= int64(v2))) {
        if false {
            var v17 uint8 = 2
            _ = v17
        }
        fmt.Println(v2)
        v18 := GenS1{F0: -32768, F1: 210}
        _ = v18
    }
    if true {
        delete(v4, ("" + "a"))
        var v19 int = 0
        v20 := GenS1{F0: 1000, F1: 1}
        _ = v19
        _ = v20
    } else {
        v2 = v2
        v21 := GenS1{F0: 42, F1: 254}
        _ = v21
    }
    fmt.Println(v3, v3, v2)
    v2 = v2
    var v22 int64 = -696
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
