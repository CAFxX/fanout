package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }

func main() {
    var v1 uint16 = 65535
    var v2 float32 = (((0.0 * 0.0) * float32(3.14159)) - 1.5)
    var v3 bool = (((!((2147483647 != 452) || true)) || false) && (false && true))
    v4 := []uint32{4294967295, 51, 2}
    v5 := make(map[int]bool)
    var v6 float64 = 3.14159
    _ = len(v4)
    v3 = false
    v7 := GenS1{F0: (v1 % ((815) | 1)), F1: 255, F2: (("foo bar" + "foo bar") + ("" + "hello"))}
    if (v3 || v3) {
        delete(v5, 855)
        v1 = 583
        if (!(!((v1 * (^(v1))) == v1))) {
            var v8 int16 = -1
            switch -802 {
            case 0:
                var v9 int32 = 1
                _ = v9
            case 42:
                if false {
                    _ = len(v5)
                    var v10 int16 = int16(v1)
                    _ = v10
                }
            case -362:
                v2 = ((v2 / (v2 + v2)) - (1e10 - v2))
            }
            _ = v8
        } else {
            var v11 uint64 = 629
            var v12 uint64 = (v11 * v11)
            var v13 string = ("" + "foo bar")
            _ = v11
            _ = v12
            _ = v13
        }
    }
    v1 -= (v1 | (v1 / ((2) | 1)))
    switch 477 {
    case 128:
        var v14 bool = v3
        _ = v14
    }
    var v15 float64 = ((0.1 / 1e10) * v6)
    for fi16 := 0; fi16 < 9; fi16++ {
        v2 += float32(v15)
        _ = fi16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        if v3 {
            for fi18 := 0; fi18 < 3; fi18++ {
                if true {
                    v1 = (v1 | v1)
                } else {
                    v3 = ((v1 + (^(v1))) >= v1)
                    v6 += v15
                    fmt.Println(v1, v3)
                    switch 128 {
                    case 127:
                        var v19 int64 = 1
                        _ = v19
                    case -358:
                        v20 := GenS1{F0: 0, F1: 469, F2: ("a" + "\n\t")}
                        var v21 int = fi17
                        _ = len(v4)
                        _ = v20
                        _ = v21
                    default:
                        v6 -= v6
                    }
                }
                _ = fi18
            }
            delete(v5, fi17)
        }
        _ = fi17
    }
    var v22 int16 = 760
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
