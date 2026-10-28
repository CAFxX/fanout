package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint64; F2 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int8, p1 int, p2 string) uint {
    v7 := []int{}
    p0 -= (p0 * (^(p0)))
    var v8 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (("hello" + "") > "\n\t")
    v3 := []uint32{}
    var v4 uint16 = 255
    v5 := []int{-483, 2147483647, 4294967295}
    var v6 int64 = 1
    if (!v2) {
        switch 32767 {
        case -9223372036854775808:
            v9 := GenS1{F0: true}
            _ = v9
        }
        if (("hello" + "hello") != (((("hello" + "foo bar") + ("foo bar" + "")) + "") + (("\n\t" + "hello") + "foo bar"))) {
            fmt.Println(v4)
            if len(v3) > 0 {
                v3[0] = 42
            }
            fmt.Println(v2, v6, v6)
        } else {
            switch -1 {
            case -1:
                v2 = (v2 || false)
            case 32768:
                var v10 string = ("\n\t" + "\n\t")
                _ = v10
            case 898:
                v2 = v2
            default:
                switch "\n\t" {
                case "a":
                    v11 := GenS2{F0: (("a" + "hello") + ("" + ""))}
                    _ = v11
                }
            }
            var v12 string = "a"
            _ = v12
        }
        if (2 == 100) {
            _ = len(v3)
        }
    } else {
        fmt.Println(v2, v4, v4)
        fmt.Println(v2, v2)
        for fi13 := 0; fi13 < 19; fi13++ {
            v4 += 255
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        switch "a" {
        case "":
            switch 501 {
            case -2147483648:
                _ = len(v1)
            case 2147483647:
                var v15 int16 = 0
                _ = v15
            }
        default:
            var v16 bool = (!(!(true && v2)))
            _ = v16
        }
        if v2 {
            for fi17 := 0; fi17 < 17; fi17++ {
                v4 = v4
                _ = fi17
            }
        } else {
            var v18 float32 = 1.5
            v2 = true
            _ = v18
        }
        _ = fi14
    }
    v19 := GenS3{F0: (0.1 != ((1.5 / (-2.25 + (1.5 - 0.0))) + float32((float64((1.5 * -2.25)) - 0.1)))), F1: 318, F2: v6}
    v20 := GenS4{F0: v6}
    fmt.Println(v2)
    var v21 bool = v2
    var v22 uint32 = 25
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
