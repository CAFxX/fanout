package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 uint) string {
    p0 -= 42
    fmt.Println(p0)
    var v7 uint32 = 723
    _ = p0
    _ = v7
    return "hello"
}

func main() {
    var v1 bool = (765 < 2)
    v2 := make(map[string]uint32)
    var v3 int64 = 804
    v4 := make(map[uint32]uint32)
    var v5 int8 = 0
    var v6 int8 = (v5 | 127)
    v5 = v5
    switch ("hello" + "") {
    case "hello":
        var v8 int = 42
        _ = v8
    case "\n\t":
        var v9 uint16 = 529
        _ = v9
    case "a":
        var v10 uint64 = 0
        _ = v10
    default:
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int8 = v6
            _ = fi11
            _ = v12
        }
    }
    if (0 != 1) {
        v13 := []int{}
        _ = v13
    }
    switch 2147483647 {
    case 9223372036854775807:
        var v14 float32 = (float32(3.14159) * float32((-2.25 + 3.14159)))
        _ = v14
    case -324:
        v15 := GenS1{F0: 943, F1: 1, F2: 0}
        _ = v15
    }
    fmt.Println(v3)
    var v16 int16 = 0
    var v17 string = "a"
    var v18 uint32 = 56
    v19 := GenS1{F0: uint(v18), F1: 47, F2: 906}
    v1 = true
    fmt.Println(v17, v3, v17)
    var v20 int64 = v3
    var v21 uint16 = 919
    if (v1 || v1) {
        v17 = v17
        if (v1 || true) {
            switch -9223372036854775808 {
            case 127:
                switch -9223372036854775808 {
                case 128:
                    switch "" {
                    case "\n\t":
                        switch v17 {
                        case "a":
                            var v22 uint = 1
                            _ = v22
                        }
                    default:
                        v20 = -819
                    }
                default:
                    fmt.Println(v18, v16, v21)
                }
            case 32767:
                if (0 < 128) {
                    fmt.Println(v20)
                    fmt.Println(v5)
                }
            default:
                v4[v18] = (v18 - (v18 | 42))
            }
            switch "hello" {
            case "hello":
                var v23 float32 = float32(float64((float32(0.0) * -2.25)))
                _ = v23
            }
            for fi24 := 0; fi24 < 4; fi24++ {
                v21 = (v21 >> ((4294967295) & 63))
                if ((v1 || ((9223372036854775807 != (v3 + v3)) && false)) || v1) {
                    v17 = ""
                } else {
                    fmt.Println(v18, v21)
                }
                _ = fi24
            }
        } else {
            var v25 string = v17
            var v26 int16 = v16
            v27 := GenS2{F0: 1, F1: (v20 / ((v3) | 1))}
            _ = v25
            _ = v26
            _ = v27
        }
        v20 = (v20 << ((0) & 63))
    } else {
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 uint64 = 642
            v30 := GenS2{F0: 202, F1: (v20 + (v3 * (v20 | int64(fi28))))}
            _ = fi28
            _ = v29
            _ = v30
        }
        v20 += v20
    }
    if (v1 && (!v1)) {
        v18 = v18
        v31 := GenS1{F0: 1, F1: 0, F2: 0}
        v32 := []int{-9223372036854775808, -9223372036854775808}
        _ = v31
        _ = v32
    } else {
        var v33 int64 = (v20 + v20)
        for fi34 := 0; fi34 < 17; fi34++ {
            var v35 int = (-(fi34))
            _ = fi34
            _ = v35
        }
        v16 = v16
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
