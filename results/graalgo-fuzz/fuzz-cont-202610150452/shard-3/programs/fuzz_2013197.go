package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
func genFn0(p0 uint64, p1 bool, p2 string) string {
    for fi6 := 0; fi6 < 14; fi6++ {
        if (((!(1.5 <= (float32(3.14159) * 3.14159))) || (p2 >= "foo bar")) && (506 <= 255)) {
            v7 := []int{(fi6 / (((fi6 | (fi6 & fi6))) | 1)), fi6}
            _ = len(v7)
            p2 = (p2 + "\n\t")
            _ = v7
        } else {
            v8 := GenS1{F0: 9223372036854775807, F1: 241}
            _ = v8
        }
        _ = fi6
    }
    if (!(false && (!(int(p0) <= -620)))) {
        if p1 {
            v9 := GenS2{F0: 32768, F1: 0, F2: 2}
            if ((632 == 1) || (42 >= 42)) {
                if p1 {
                    var v10 uint64 = p0
                    _ = v10
                }
                switch -461 {
                case -926:
                    var v11 int = 431
                    _ = v11
                case 1:
                    v12 := []int{128}
                    _ = v12
                case 9223372036854775807:
                    v13 := []int{127, 128, -9223372036854775808, 0}
                    _ = v13
                }
            }
            _ = v9
        }
        fmt.Println(p2, p0, p2)
        p2 = ("" + "hello")
    }
    var v14 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return p2
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    var v4 string = v1
    var v5 uint8 = 2
    v15 := GenS2{F0: 128, F1: 1, F2: 2}
    var v16 int8 = 42
    fmt.Println(v5, v1)
    delete(v3, -316)
    v17 := []int{0, int(v5)}
    var v18 int64 = -388
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS2{F0: (fi19 << ((42) & 63)), F1: 0, F2: 4294967295}
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: (^(v18)), F1: v5}
    var v22 string = "a"
    _ = len(v2)
    if false {
        v18 += v18
        v4 = "foo bar"
        v1 = "a"
    }
    fmt.Println(v1)
    var v23 int = 2147483647
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 uint8 = (v5 / ((98) | 1))
        for fi26 := 0; fi26 < 16; fi26++ {
            switch ("hello" + v1) {
            case "a":
                switch fi26 {
                case 32767:
                    fmt.Println(v18, v5, v16)
                case 0:
                    v27 := GenS1{F0: (-(v18)), F1: v5}
                    _ = v27
                default:
                    var v28 int8 = v16
                    _ = v28
                }
            default:
                var v29 int64 = (v18 >> ((uint(v5)) & 63))
                fmt.Println(v25, v18, v18)
                switch v22 {
                case "\n\t":
                    var v30 uint8 = (^(v25))
                    _ = v30
                case "foo bar":
                    var v31 float64 = 0.1
                    _ = v31
                }
                _ = v29
            }
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
