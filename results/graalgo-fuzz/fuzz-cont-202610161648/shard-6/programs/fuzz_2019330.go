package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
func genFn0(p0 uint, p1 int32) uint64 {
    var v7 uint = p0
    v8 := GenS1{F0: 1}
    v7 -= (p0 << (((v7 ^ p0)) & 63))
    v7 = p0
    var v9 uint32 = 21
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int8 = 93
    var v2 float32 = 1.5
    var v3 int = -1
    var v4 uint = 42
    var v5 float32 = (v2 / 1.0)
    var v6 int32 = -1
    var v10 uint64 = 2
    var v11 float32 = (v5 - 3.14159)
    switch (^(v3)) {
    case 32768:
        fmt.Println(v6, v5)
    case 127:
        fmt.Println(v2, v11, v5)
    case 42:
        v12 := []int{(v3 - 0), (-(v3)), v3, (v3 ^ 32767)}
        _ = v12
    }
    var v13 uint16 = 65535
    if false {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 float64 = float64((v2 + (v11 / v11)))
            v16 := GenS2{F0: true}
            if ((v3 * (fi14 | (v3 + fi14))) <= fi14) {
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 bool = true
                    var v19 int64 = -9223372036854775808
                    var v20 int32 = v6
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 20; fi21++ {
                    v10 -= (v10 - v10)
                    v22 := []int{42, fi14, (v3 + fi14), int(v6)}
                    var v23 uint = (v4 * v4)
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            } else {
                v1 -= v1
                v24 := GenS2{F0: (true && (uint32(v10) == 8))}
                switch "foo bar" {
                case "hello":
                    v25 := []int{}
                    _ = v25
                default:
                    if true {
                        var v26 uint8 = 128
                        _ = v26
                    }
                }
                _ = v24
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        if false {
            fmt.Println(v4)
            for fi27 := 0; fi27 < 7; fi27++ {
                var v28 int32 = (v6 + v6)
                _ = fi27
                _ = v28
            }
        } else {
            var v29 int16 = 0
            _ = v29
        }
    }
    fmt.Println(v2, v10)
    switch "hello" {
    case "\n\t":
        v30 := []int{}
        _ = v30
    }
    var v31 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
