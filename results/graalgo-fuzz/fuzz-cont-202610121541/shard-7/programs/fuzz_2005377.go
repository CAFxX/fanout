package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 uint8, p1 int64, p2 uint64) int {
    p0 += p0
    p1 = (p1 - p1)
    p0 += (p0 / ((p0) | 1))
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() uint8 {
    var v4 string = ""
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            if false {
                v4 = ("\n\t" + "hello")
                v7 := []int{4294967295, 42}
                v4 = v4
                if true {
                    var v8 float64 = 0.0
                    _ = v8
                }
                _ = v7
            }
            switch v4 {
            case "\n\t":
                switch v4 {
                case "\n\t":
                    var v9 int32 = 436
                    _ = v9
                case "foo bar":
                    var v10 uint8 = 2
                    _ = v10
                }
            case "":
                v11 := []int{}
                _ = v11
            default:
                v4 = "hello"
            }
            _ = fi6
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            v4 = ((((v4 + v4) + (v4 + v4)) + v4) + v4)
            var v13 float64 = 1e10
            v13 += (v13 - 1e10)
            _ = fi12
            _ = v13
        }
        _ = fi5
    }
    _ = v4
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 907
    v3 := make(map[string]bool)
    v14 := GenS1{F0: "a"}
    fmt.Println(v2)
    v2 = (v2 - -1)
    v2 -= 9223372036854775807
    var v15 uint16 = 1
    switch -628 {
    case 32768:
        v16 := []int{1, v2, v2}
        _ = v16
    case 42:
        switch "a" {
        case "foo bar":
            for fi17 := 0; fi17 < 4; fi17++ {
                v15 += (^(v15))
                _ = fi17
            }
        default:
            if true {
                v2 += v2
                v18 := GenS1{F0: (("a" + "a") + "\n\t")}
                v2 = genFn0(1, 655, 792)
                _ = v18
            } else {
                v19 := GenS1{F0: "\n\t"}
                _ = v19
            }
        }
    }
    v20 := GenS2{F0: 2, F1: -128}
    if true {
        fmt.Println(v15, v2)
    }
    var v21 int32 = 0
    v2 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
