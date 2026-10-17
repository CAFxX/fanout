package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int16) string {
    var v6 float64 = 0.0
    p0 = 32767
    _ = p0
    _ = v6
    return "\n\t"
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 4294967295
    var v3 uint64 = 2
    var v4 int8 = 42
    var v5 uint64 = v3
    fmt.Println(v3, v1)
    if (127 < 103) {
        var v7 string = "a"
        fmt.Println(v1, v3)
        fmt.Println(v7)
        _ = v7
    } else {
        switch 26 {
        case 42:
            var v8 float32 = 0.1
            _ = v8
        case 1:
            switch -945 {
            case 32767:
                v9 := GenS1{F0: 230, F1: "", F2: 42}
                _ = v9
            case 945:
                switch "a" {
                case "":
                    var v10 bool = true
                    _ = v10
                case "a":
                    for fi11 := 0; fi11 < 15; fi11++ {
                        var v12 uint8 = 127
                        var v13 float32 = 3.14159
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                case "\n\t":
                    var v14 float64 = float64(float32((0.0 - (1e10 / (1.5 + 0.0)))))
                    _ = v14
                default:
                    v2 = (^(v2))
                }
            case 128:
                var v15 uint32 = 2
                _ = v15
            }
        case 128:
            v16 := GenS2{F0: (false && (2147483647 != 498))}
            _ = v16
        }
        if true {
            switch "hello" {
            case "hello":
                v17 := GenS2{F0: false}
                _ = v17
            case "foo bar":
                var v18 uint8 = 20
                _ = v18
            case "a":
                v19 := []int{32768, 32768, -251}
                _ = v19
            }
        } else {
            v20 := GenS1{F0: 1, F1: ("\n\t" + "foo bar"), F2: 0}
            v2 = v2
            _ = v20
        }
        if (true || (65535 <= (v2 + v2))) {
            if ((false && ((!true) || true)) || ((v2 & (v2 + v2)) <= (v2 + 96))) {
                var v21 string = ""
                fmt.Println(v4, v3, v1)
                v4 -= (v4 % (((-(v4))) | 1))
                _ = v21
            } else {
                for fi22 := 0; fi22 < 12; fi22++ {
                    v5 -= 4294967295
                    v23 := []int{(fi22 - fi22), (fi22 | 128), fi22}
                    _ = fi22
                    _ = v23
                }
                var v24 uint = 42
                for fi25 := 0; fi25 < 8; fi25++ {
                    v26 := GenS3{F0: v3}
                    fmt.Println(v5, v3)
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
        }
    }
    v27 := []int{-1, -2147483648, int(v4)}
    if (!((int16(v1) > -41) && (0 > 0))) {
        var v28 int8 = v4
        _ = v28
    }
    var v29 uint64 = 42
    var v30 int = 896
    var v31 bool = (!true)
    if (v31 || v31) {
        v29 = (v3 * (v29 / ((v29) | 1)))
        var v32 int = (v30 << ((4294967295) & 63))
        _ = v32
    }
    var v33 uint16 = 1
    v3 = 42
    var v34 int16 = -243
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
