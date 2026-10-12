package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 uint; F1 bool; F2 string }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 bool, p1 uint64, p2 uint) string {
    var v5 uint16 = 0
    var v6 uint8 = 238
    switch ("a" + ("" + ("" + ""))) {
    case "a":
        var v7 uint8 = 2
        _ = v7
    case "hello":
        v8 := GenS1{F0: v5, F1: -1, F2: p0}
        _ = v8
    default:
        var v9 uint32 = 328
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ""
}

func main() {
    var v1 bool = (1 > 2147483647)
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 uint8 = 42
    for fi10 := 0; fi10 < 10; fi10++ {
        switch fi10 {
        case 1:
            var v11 uint = 2
            _ = v11
        case 9223372036854775807:
            var v12 uint8 = (v4 | 2)
            _ = v12
        case 32768:
            var v13 bool = (true || (!false))
            _ = v13
        }
        v14 := GenS2{F0: 234, F1: v1, F2: "hello"}
        _ = fi10
        _ = v14
    }
    v2[("hello" + "hello")] = 0
    v3[4294967295] = 305
    fmt.Println(v1, v4)
    v15 := GenS2{F0: 18446744073709551615, F1: (-128 < -61), F2: ""}
    switch -2147483648 {
    case 841:
        var v16 float32 = (float32(3.14159) / 1.0)
        _ = v16
    case 32767:
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 uint8 = v4
            if v1 {
                for fi19 := 0; fi19 < 2; fi19++ {
                    if v1 {
                        if (!true) {
                            v20 := []int{fi17, 128, fi17}
                            _ = v20
                        }
                        v1 = v1
                        var v21 uint8 = 2
                        _ = v21
                    } else {
                        fmt.Println(v4, v1, v4)
                    }
                    var v22 float32 = (1.5 * 0.1)
                    var v23 float32 = 1.5
                    _ = fi19
                    _ = v22
                    _ = v23
                }
            } else {
                var v24 uint64 = 115
                _ = v24
            }
            v1 = (((false || v1) || v1) && v1)
            _ = fi17
            _ = v18
        }
    case 2147483647:
        v25 := []int{}
        _ = v25
    default:
        v4 -= (^(v4))
    }
    v26 := GenS2{F0: 0, F1: (v1 || v1), F2: ("\n\t" + "foo bar")}
    var v27 float64 = (((-2.25 - 0.0) + 0.1) * 1e10)
    for fi28 := 0; fi28 < 6; fi28++ {
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 bool = (1 <= -1)
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    fmt.Println(v1, v4, v1)
    var v31 float32 = ((((float32(v27) - (float32(v27) - (1e10 + (3.14159 - -2.25)))) / 1.0) + float32(float64(-2.25))) + (-2.25 + ((0.1 / 1.0) / (1.5 - (3.14159 / 1.0)))))
    var v32 uint64 = 2
    v32 -= 371
    var v33 string = ("hello" + ("\n\t" + "foo bar"))
    var v34 float64 = -2.25
    switch 838 {
    case -739:
        v35 := GenS3{F0: 127, F1: 0}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
