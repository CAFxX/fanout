package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int64; F1 int; F2 string }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 string; F1 bool }
func genFn0() int16 {
    if ((0.1 != (3.14159 / (1.5 + (3.14159 * 0.1)))) && (2 <= 1)) {
        v5 := GenS1{F0: 180, F1: -112, F2: -1}
        _ = v5
    }
    var v6 uint = 4294967295
    v6 = v6
    if false {
        var v7 string = (("hello" + "\n\t") + "hello")
        var v8 float64 = (0.1 + 0.0)
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                switch -2147483648 {
                case 4294967295:
                    v11 := GenS2{F0: ((210 > uint8(v6)) || true)}
                    _ = v11
                case 348:
                    var v12 float32 = float32(v8)
                    _ = v12
                case -364:
                    v13 := []int{}
                    _ = v13
                default:
                    var v14 float64 = (v8 + (-2.25 / v8))
                    _ = v14
                }
                v7 = "\n\t"
                v15 := GenS3{F0: 777, F1: fi9}
                _ = fi10
                _ = v15
            }
            var v16 int = (fi9 >> (((v6 + 2)) & 63))
            v7 = "\n\t"
            _ = fi9
            _ = v16
        }
        _ = v7
        _ = v8
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        if false {
            v6 += 472
            v18 := GenS1{F0: fi17, F1: 42, F2: 42}
            v6 += v6
            _ = v18
        }
        _ = fi17
    }
    _ = v6
    return -1
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint64 = 695
    v3 := make(map[int]bool)
    var v4 uint = 845
    v19 := GenS4{F0: 930, F1: -2147483648, F2: ""}
    if true {
        switch 168 {
        case -607:
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS5{F0: 42, F1: 335}
                _ = fi20
                _ = v21
            }
        case -2147483648:
            for fi22 := 0; fi22 < 17; fi22++ {
                v23 := []int{fi22, 127, fi22}
                v4 = v4
                v4 -= (v4 + (^(v4)))
                _ = fi22
                _ = v23
            }
        case 2147483647:
            v24 := GenS6{F0: (("foo bar" + (("\n\t" + "\n\t") + (("" + "\n\t") + "a"))) + "hello"), F1: ((1.5 / ((float32(0.1) / 1.5) + ((((float32(v1) / (1e10 * 3.14159)) + ((3.14159 - (1e10 - (1.5 - 0.0))) - 1e10)) / 1.0) / (((-2.25 * 3.14159) - 1e10) * float32(float64(-2.25)))))) > 3.14159)}
            _ = v24
        }
    } else {
        var v25 uint32 = 638
        v2 += v2
        var v26 uint8 = 190
        _ = v25
        _ = v26
    }
    _ = len(v3)
    v1 -= (0.0 / (-2.25 + v1))
    v2 += v2
    v4 = (v4 % ((v4) | 1))
    var v27 float32 = (3.14159 + (0.1 * 1.5))
    fmt.Println(v27, v1)
    v1 += float64(0.1)
    v4 -= 18446744073709551615
    for fi28 := 0; fi28 < 15; fi28++ {
        switch "\n\t" {
        case "a":
            v29 := []int{1, (^(fi28))}
            _ = v29
        case "foo bar":
            fmt.Println(v27, v2)
        }
        v4 -= 65535
        if ((0 == 32767) && (-9223372036854775808 <= 762)) {
            v30 := v3[fi28]
            for fi31 := 0; fi31 < 15; fi31++ {
                _ = len(v3)
                delete(v3, fi28)
                _ = fi31
            }
            _ = v30
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
