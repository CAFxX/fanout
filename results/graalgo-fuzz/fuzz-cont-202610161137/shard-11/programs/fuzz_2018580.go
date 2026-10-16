package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 int64; F1 bool; F2 int32 }
func genFn0(p0 int64, p1 uint64, p2 string) uint16 {
    v4 := []int{127, 261}
    if (!true) {
        if ((true && (p2 != p2)) && ((p1 / (((p1 % (((^(p1))) | 1))) | 1)) >= p1)) {
            switch p2 {
            case "":
                var v5 uint = 4294967295
                _ = v5
            case "a":
                switch "hello" {
                case "foo bar":
                    var v6 uint = 1
                    _ = v6
                default:
                    var v7 uint32 = 0
                    _ = v7
                }
            case "\n\t":
                p0 = -9223372036854775808
            }
            var v8 int64 = 294
            _ = v8
        }
        p2 = "a"
    } else {
        var v9 uint8 = 128
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 715
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]uint32)
    var v3 float64 = 1.5
    var v10 float64 = v3
    if (42 >= 105) {
        v3 -= (v10 * v3)
        v3 = 0.0
        v3 -= (v10 * v3)
    } else {
        v1 = v1
    }
    var v11 float32 = -2.25
    v12 := GenS1{F0: 2, F1: 92}
    for fi13 := 0; fi13 < 18; fi13++ {
        fmt.Println(v11, fi13, v3)
        v14 := GenS2{F0: 190, F1: 709, F2: ""}
        v15 := GenS2{F0: 128, F1: 42, F2: "a"}
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: 0}
    var v17 uint8 = 255
    for fi18 := 0; fi18 < 4; fi18++ {
        if false {
            v11 += (v11 * v11)
        } else {
            v17 += (v17 >> ((uint(v1)) & 63))
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            v3 -= 3.14159
            v20 := []int{(fi18 >> ((2) & 63)), -80, (fi19 + fi18)}
            if (!(!(!true))) {
                v21 := GenS3{F0: 98, F1: (true || true), F2: 0}
                _ = v21
            }
            _ = fi19
            _ = v20
        }
        for fi22 := 0; fi22 < 9; fi22++ {
            switch fi18 {
            case -385:
                var v23 float64 = v3
                _ = v23
            case 127:
                fmt.Println(fi18, v17)
            case 42:
                var v24 uint = 65535
                var v25 uint16 = genFn0(9223372036854775807, 517, "hello")
                _ = v24
                _ = v25
            default:
                var v26 bool = (true || (107 >= 1))
                _ = v26
            }
            v2[(-(fi22))] = (v1 % ((65535) | 1))
            fmt.Println(fi22, v1, v3)
            _ = fi22
        }
        _ = fi18
    }
    fmt.Println(v11, v11, v11)
    var v27 bool = false
    var v28 int8 = 127
    var v29 uint16 = 32767
    v1 = (v1 + v1)
    v29 -= v29
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
