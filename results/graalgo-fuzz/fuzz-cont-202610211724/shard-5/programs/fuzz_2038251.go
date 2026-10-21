package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int32; F1 bool; F2 int8 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 int16) int8 {
    var v4 uint16 = 1
    v4 += (v4 - v4)
    fmt.Println(p0)
    _ = p0
    _ = v4
    return -128
}
func genFn1() uint16 {
    v5 := GenS1{F0: 42, F1: 42}
    v6 := GenS1{F0: 18446744073709551615, F1: 0}
    _ = v5
    _ = v6
    return 796
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 83
    var v3 int8 = -118
    _ = len(v1)
    var v7 uint16 = 1
    var v8 int = 289
    v1[v8] = false
    if (-2.25 != (0.0 + (0.0 * 0.1))) {
        v3 += v3
        switch "foo bar" {
        case "\n\t":
            v2 = 60
        default:
            var v9 int32 = 42
            _ = v9
        }
    } else {
        switch ("" + ("foo bar" + "a")) {
        case "":
            for fi10 := 0; fi10 < 6; fi10++ {
                v7 = genFn1()
                v7 = (v7 | (v7 + v7))
                switch ("\n\t" + ((("\n\t" + (("a" + "foo bar") + "\n\t")) + ((("hello" + "hello") + "a") + "hello")) + "hello")) {
                case "":
                    v11 := []int{(v8 + 370)}
                    _ = v11
                case "hello":
                    v2 = v2
                case "a":
                    v2 -= (v2 << ((423) & 63))
                default:
                    if (!true) {
                        if true {
                            v12 := GenS2{F0: -1, F1: true, F2: -128}
                            var v13 int16 = 42
                            _ = v12
                            _ = v13
                        }
                    } else {
                        v8 += fi10
                        delete(v1, 42)
                    }
                }
                _ = fi10
            }
        case "hello":
            var v14 int32 = 600
            _ = v14
        }
    }
    v15 := GenS2{F0: 2147483647, F1: false, F2: v3}
    var v16 float64 = 1.5
    v16 = 0.0
    v17 := GenS2{F0: 2147483647, F1: (!(false && (0 > 928))), F2: (v3 << ((550) & 63))}
    v7 = v7
    v18 := GenS3{F0: -873, F1: -2147483648}
    for fi19 := 0; fi19 < 19; fi19++ {
        v3 = (-(v3))
        _ = fi19
    }
    v8 -= -844
    if (!(((1e10 / (3.14159 / ((1.5 / 1.0) / 3.14159))) * 3.14159) <= 3.14159)) {
        v8 = v8
    }
    fmt.Println(v16, v16)
    if false {
        v2 -= v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
