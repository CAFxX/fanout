package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }
func genFn0(p0 int32, p1 int8) int16 {
    var v5 float64 = (1e10 * (0.1 * (-2.25 / 0.1)))
    var v6 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 32767
}
func genFn1() uint32 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint32 = 1
        v9 := GenS1{F0: 1000, F1: 18446744073709551615, F2: ("\n\t" + "a")}
        var v10 uint16 = 65535
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := []int{}
        fmt.Println(fi11)
        var v13 bool = false
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int64 = 42
    v15 := GenS1{F0: 94, F1: 4294967295, F2: ("\n\t" + "hello")}
    v16 := []int{32767, 127}
    _ = v14
    _ = v15
    _ = v16
    return 331
}

func main() {
    var v1 uint32 = 1
    var v2 uint = 1
    var v3 int = 141
    v4 := make(map[int]string)
    switch "" {
    case "\n\t":
        var v17 int64 = -844
        _ = v17
    case "":
        var v18 uint64 = 42
        _ = v18
    case "foo bar":
        v3 = 127
    default:
        _ = len(v4)
    }
    v3 += v3
    var v19 float32 = 1e10
    v19 = (v19 / (-2.25 - v19))
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS1{F0: 42, F1: 2, F2: "hello"}
        fmt.Println(fi20, v3, fi20)
        _ = fi20
        _ = v21
    }
    var v22 float64 = float64(((float32(3.14159) * v19) / 1.0))
    v23 := GenS1{F0: 0, F1: 1, F2: ("a" + "")}
    v24 := v4[9223372036854775807]
    var v25 bool = (!(v3 != v3))
    v1 += (v1 / (((v1 + genFn1())) | 1))
    v26 := GenS1{F0: 32767, F1: 42, F2: v24}
    v25 = (0.1 >= float64((1e10 / float32(-2.25))))
    if true {
        if (v19 >= v19) {
            switch (("" + v24) + "hello") {
            case "foo bar":
                var v27 int32 = 1
                _ = v27
            case "\n\t":
                fmt.Println(v3, v24)
            case "a":
                var v28 float64 = v22
                _ = v28
            default:
                switch ("foo bar" + v24) {
                case "hello":
                    var v29 int = v3
                    _ = v29
                default:
                    v24 = (v24 + "a")
                }
            }
            v30 := []int{v3, (v3 >> ((v2) & 63)), -618, (v3 << (((v2 * (v2 % ((v2) | 1)))) & 63))}
            _ = v30
        }
        v31 := GenS1{F0: 1, F1: 1, F2: v24}
        delete(v4, (v3 + v3))
        _ = v31
    }
    v32 := v4[(v3 ^ (-(v3)))]
    v22 = (float64(0.1) / v22)
    var v33 int64 = 1
    v4[int(v1)] = (v24 + v24)
    var v34 string = v24
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
