package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0() int {
    var v6 uint64 = 0
    fmt.Println(v6)
    _ = v6
    return 32767
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]int)
    var v3 int64 = 42
    var v4 uint = 964
    var v5 uint8 = 42
    v5 += (^(v5))
    v7 := []int{int(v4)}
    if false {
        v5 -= 128
        switch "" {
        case "foo bar":
            v4 += 42
        case "hello":
            for fi8 := 0; fi8 < 16; fi8++ {
                v9 := GenS1{F0: 32767, F1: "foo bar"}
                _ = fi8
                _ = v9
            }
        }
    } else {
        v5 = 29
        var v10 uint = v4
        var v11 int = 9223372036854775807
        _ = v10
        _ = v11
    }
    if true {
        _ = len(v2)
    } else {
        v5 = 128
    }
    v5 = 42
    v4 = v4
    _ = len(v7)
    for fi12 := 0; fi12 < 13; fi12++ {
        if (!(284 >= uint16(v5))) {
            var v13 int64 = v3
            var v14 bool = ("hello" != (("a" + "foo bar") + ("hello" + "\n\t")))
            _ = v13
            _ = v14
        }
        var v15 int32 = -314
        for fi16 := 0; fi16 < 17; fi16++ {
            v5 = (v5 - v5)
            _ = fi16
        }
        _ = fi12
        _ = v15
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 bool = (!(((!false) && false) && true))
        _ = fi17
        _ = v18
    }
    if ((^(v3)) >= 0) {
        var v19 int8 = -1
        _ = v19
    }
    _ = len(v7)
    for fi20 := 0; fi20 < 6; fi20++ {
        fmt.Println(v3, v5)
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
