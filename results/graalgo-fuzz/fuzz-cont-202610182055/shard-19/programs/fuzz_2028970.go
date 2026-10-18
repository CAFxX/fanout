package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
func genFn0(p0 int8, p1 bool, p2 int32) uint64 {
    switch 930 {
    case -788:
        var v5 float64 = (0.1 - ((1e10 / float64(1.5)) / 3.14159))
        _ = v5
    }
    v6 := GenS1{F0: 74}
    var v7 string = (("hello" + "a") + ((("" + "hello") + (("foo bar" + "hello") + "foo bar")) + "foo bar"))
    v8 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int = 265
    var v2 string = "\n\t"
    var v3 string = v2
    v4 := make(map[uint32]bool)
    v3 = "\n\t"
    var v9 int8 = -79
    v10 := GenS1{F0: 2}
    switch v1 {
    case 481:
        for fi11 := 0; fi11 < 13; fi11++ {
            v9 = (v9 + -128)
            _ = fi11
        }
    case -195:
        v12 := GenS1{F0: 733}
        _ = v12
    case 32768:
        var v13 uint = 65535
        _ = v13
    }
    fmt.Println(v1, v9, v3)
    var v14 float64 = (1.5 / -2.25)
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(v3)
        switch "" {
        case "\n\t":
            v16 := GenS1{F0: 4294967295}
            _ = v16
        default:
            for fi17 := 0; fi17 < 18; fi17++ {
                v2 = ((v2 + v3) + (v2 + "foo bar"))
                _ = fi17
            }
        }
        _ = fi15
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v14 -= v14
        if true {
            v3 = ((v2 + ((v3 + (("a" + (v2 + v3)) + ((v3 + "foo bar") + v2))) + v2)) + "hello")
        } else {
            v19 := GenS1{F0: 636}
            if false {
                var v20 float64 = 3.14159
                _ = v20
            } else {
                v4[644] = (false && true)
            }
            _ = v19
        }
        v21 := GenS2{F0: 9223372036854775807}
        _ = fi18
        _ = v21
    }
    _ = len(v4)
    for fi22 := 0; fi22 < 10; fi22++ {
        v1 -= (fi22 * v1)
        _ = fi22
    }
    v14 = v14
    fmt.Println(v1)
    v3 = ("a" + v2)
    fmt.Println(v3, v2, v1)
    var v23 uint32 = 42
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 uint16 = 2
        fmt.Println(v23, v3)
        v26 := []int{(fi24 * fi24)}
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
