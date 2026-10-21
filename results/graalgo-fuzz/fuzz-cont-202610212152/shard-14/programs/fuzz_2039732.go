package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 uint; F1 string; F2 int16 }
func genFn0() uint16 {
    v7 := GenS1{F0: -1, F1: 0, F2: 42}
    var v8 uint = 529
    _ = v7
    _ = v8
    return 878
}
func genFn1(p0 int, p1 uint64, p2 bool) int16 {
    fmt.Println(p1, p2, p0)
    var v9 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -1
}

func main() {
    var v1 uint64 = 570
    var v2 float64 = 1.5
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v5 uint = 255
    var v6 float32 = ((((1.5 - (1.5 - ((((0.0 - float32(((((0.0 - ((v2 / v2) - (((0.1 + v2) / (3.14159 + v2)) / v2))) * (float64(((((1.5 / float32(v2)) * 3.14159) * float32(((0.0 + 1e10) / v2))) + 0.1)) + (3.14159 * v2))) + v2) * v2))) / 1.0) / 1.5) * 1e10))) + ((float32((float64(-2.25) - (float64(1e10) + ((1e10 * 3.14159) - (v2 * 0.0))))) + 1e10) * 1.5)) * (1.5 + (0.1 - 0.1))) + 3.14159)
    v1 += (v1 >> ((v5) & 63))
    _ = len(v4)
    v10 := GenS1{F0: -98, F1: -128, F2: 127}
    v11 := GenS2{F0: 0}
    var v12 float32 = v6
    switch ("" + "hello") {
    case "a":
        var v13 uint8 = 85
        _ = v13
    case "":
        if false {
            v14 := []int{0, -611, -98, 32768}
            _ = v14
        } else {
            if true {
                v15 := []int{658, -1, -44, 42}
                for fi16 := 0; fi16 < 10; fi16++ {
                    v15 = append(v15, -619)
                    fmt.Println(v12, v6, v12)
                    v17 := v4[((("" + ("hello" + "hello")) + "\n\t") + "")]
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
        }
    default:
        var v18 uint16 = 32767
        _ = v18
    }
    fmt.Println(v1, v1)
    var v19 int = int(v1)
    var v20 float32 = (v6 * v12)
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(fi21, v20, v6)
        _ = fi21
    }
    if true {
        v2 += v2
        v22 := GenS3{F0: 0, F1: false, F2: (0 == 736)}
        fmt.Println(v6)
        _ = v22
    }
    switch 1 {
    case 32768:
        v6 -= v20
    case 9223372036854775807:
        v5 += 65535
    case 128:
        v23 := GenS4{F0: (v5 >> (((v5 >> ((v5) & 63))) & 63)), F1: "", F2: 32767}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
