package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 int32) int64 {
    switch "a" {
    case "":
        switch 1 {
        case -2147483648:
            switch "foo bar" {
            case "":
                v6 := GenS1{F0: 33, F1: 79}
                _ = v6
            case "a":
                var v7 int32 = (p0 ^ -270)
                _ = v7
            case "\n\t":
                p0 -= 42
            }
        case -768:
            if (1.5 > 0.0) {
                var v8 uint16 = 568
                switch 1 {
                case 312:
                    var v9 int32 = (p0 / ((p0) | 1))
                    _ = v9
                case 32767:
                    p0 -= (p0 * 970)
                }
                _ = v8
            } else {
                v10 := []int{42}
                p0 = p0
                var v11 bool = (false || true)
                _ = v10
                _ = v11
            }
        }
    case "foo bar":
        fmt.Println(p0)
    case "hello":
        v12 := GenS1{F0: 127, F1: 100}
        _ = v12
    }
    var v13 uint32 = 42
    switch ("hello" + "") {
    case "":
        var v14 uint8 = 1
        _ = v14
    }
    v15 := []int{877, 128, 1}
    _ = p0
    _ = v13
    _ = v15
    return 1
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]string)
    v3 := []int{-164, 966, 7, -246}
    var v4 bool = (!(-158 == -1))
    v5 := []string{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = (2 > 1)
    switch 128 {
    case 981:
        var v16 int = 1
        _ = v16
    default:
        v17 := GenS1{F0: 91, F1: -74}
        _ = v17
    }
    fmt.Println(v4, v4)
    var v18 uint64 = 42
    var v19 int = 32768
    delete(v2, "\n\t")
    v19 = 42
    v5 = append(v5, (("foo bar" + "a") + "foo bar"))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
