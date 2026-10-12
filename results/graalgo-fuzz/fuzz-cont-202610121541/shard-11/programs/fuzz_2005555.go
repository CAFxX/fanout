package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
func genFn0(p0 int16, p1 string, p2 uint8) uint16 {
    var v6 uint64 = 197
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    var v1 bool = true
    var v2 bool = v1
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]string)
    v2 = true
    var v7 string = ((("foo bar" + ("foo bar" + ("hello" + ((("foo bar" + "a") + ("foo bar" + "a")) + "a")))) + "") + "hello")
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float32 = ((-2.25 / 1.5) / (0.1 - 0.0))
        var v10 float64 = ((1e10 / -2.25) + (-2.25 * (((1e10 + 1e10) * (1.5 * 1e10)) / 1.0)))
        if v1 {
            v1 = v1
            if (v2 || true) {
                v11 := GenS1{F0: 127, F1: (fi8 >> ((125) & 63)), F2: 441}
                _ = v11
            } else {
                v1 = v1
            }
            v12 := GenS1{F0: 127, F1: fi8, F2: 895}
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v13 int64 = -1
    v14 := GenS1{F0: 0, F1: 0, F2: 18446744073709551615}
    if false {
        if true {
            for fi15 := 0; fi15 < 18; fi15++ {
                if (!v2) {
                    fmt.Println(v1)
                } else {
                    var v16 int8 = 42
                    v7 = ((v7 + ("\n\t" + v7)) + "a")
                    v7 = ((v7 + "\n\t") + v7)
                    _ = v16
                }
                _ = fi15
            }
        }
    } else {
        v1 = v2
    }
    var v17 int64 = (v13 ^ v13)
    var v18 uint16 = 255
    for fi19 := 0; fi19 < 14; fi19++ {
        v2 = v1
        if v1 {
            v20 := GenS1{F0: 2, F1: fi19, F2: 0}
            var v21 uint = 1
            v7 = v7
            _ = v20
            _ = v21
        }
        _ = fi19
    }
    var v22 int32 = -1
    fmt.Println(v17, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
