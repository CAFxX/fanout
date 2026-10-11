package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0() uint16 {
    v6 := GenS1{F0: 2147483647, F1: (false || (!(!true)))}
    v7 := GenS2{F0: 0, F1: 0, F2: 127}
    var v8 uint32 = 2
    v9 := []int{-196, 885, 1}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 32767
    v3 := []string{("hello" + ("a" + "")), (("foo bar" + "a") + "")}
    var v4 int64 = -886
    v5 := make(map[uint32]uint32)
    var v10 uint64 = 18446744073709551615
    var v11 uint8 = 127
    var v12 string = "hello"
    if (-1 <= 32767) {
        var v13 uint32 = 42
        fmt.Println(v12, v2, v12)
        for fi14 := 0; fi14 < 11; fi14++ {
            v11 += (v11 | v11)
            if (0.1 >= -2.25) {
                v4 = (-(v4))
                v2 = 32767
                fmt.Println(v10, fi14, v2)
            } else {
                var v15 float32 = 0.1
                var v16 bool = (1 == -85)
                _ = v15
                _ = v16
            }
            _ = fi14
        }
        _ = v13
    }
    v12 = ("foo bar" + ((v12 + "\n\t") + ""))
    delete(v1, v12)
    if true {
        v17 := GenS1{F0: 199, F1: (2147483647 == -2147483648)}
        var v18 uint8 = 255
        _ = v17
        _ = v18
    }
    fmt.Println(v11, v12, v11)
    if (42 == 42) {
        switch 1 {
        case 2147483647:
            switch ("hello" + v12) {
            case "foo bar":
                var v19 int32 = 74
                _ = v19
            default:
                for fi20 := 0; fi20 < 3; fi20++ {
                    v21 := GenS1{F0: 1, F1: (0 < 810)}
                    var v22 float32 = (0.1 + (0.0 * 0.0))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        case 4294967295:
            switch v12 {
            case "":
                v23 := GenS2{F0: 196, F1: 65535, F2: 1}
                _ = v23
            case "hello":
                fmt.Println(v11, v2, v2)
            default:
                var v24 uint16 = 32767
                _ = v24
            }
        case -1:
            v1[((v12 + v12) + v12)] = -1
        default:
            v11 += (v11 << ((255) & 63))
        }
        var v25 float64 = (0.0 * (float64((float32(-2.25) * (1e10 + (1.5 / ((-2.25 * (1.5 * 3.14159)) - 3.14159))))) - 0.0))
        _ = v25
    } else {
        v4 = v4
    }
    var v26 uint64 = 681
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
