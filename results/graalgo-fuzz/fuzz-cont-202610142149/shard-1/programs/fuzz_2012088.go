package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 string, p1 uint) uint8 {
    fmt.Println(p1, p1)
    fmt.Println(p0)
    if ("" != p0) {
        p1 = (p1 | p1)
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 uint8 = 232
    var v4 int16 = 32767
    fmt.Println(v3)
    var v5 int8 = 100
    v6 := GenS1{F0: -198}
    var v7 int16 = v4
    for fi8 := 0; fi8 < 7; fi8++ {
        if false {
            v9 := []int{fi8}
            var v10 int64 = -1
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 bool = (false && ((v3 % ((2) | 1)) != (^(v3))))
                _ = fi11
                _ = v12
            }
            _ = v10
            _ = v9
        } else {
            var v13 uint8 = v3
            _ = len(v2)
            _ = v13
        }
        if (!true) {
            var v14 float64 = -2.25
            _ = v14
        }
        switch ("a" + "hello") {
        case "hello":
            v4 += v7
        case "\n\t":
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(v3)
                _ = fi15
            }
        case "a":
            switch ((((("foo bar" + (("" + "a") + "")) + "") + "foo bar") + (((("hello" + (("a" + (("a" + ("\n\t" + (("" + ("a" + ("" + "a"))) + ("foo bar" + ("" + ("hello" + ("a" + "foo bar"))))))) + ("foo bar" + "a"))) + ("a" + ("a" + "")))) + "") + ("hello" + "")) + ("a" + ""))) + "a") {
            case "hello":
                v16 := GenS1{F0: 89}
                _ = v16
            case "a":
                var v17 int8 = v5
                _ = v17
            default:
                var v18 uint = 1
                _ = v18
            }
        default:
            v19 := GenS2{F0: 0, F1: (v4 & v4)}
            _ = v19
        }
        _ = fi8
    }
    v20 := GenS3{F0: 990, F1: 542}
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 string = "a"
        _ = fi21
        _ = v22
    }
    v5 += 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
