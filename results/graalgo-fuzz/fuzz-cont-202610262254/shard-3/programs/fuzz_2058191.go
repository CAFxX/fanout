package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool }
func genFn0() bool {
    if (!true) {
        switch ("foo bar" + "") {
        case "hello":
            v7 := []int{4294967295}
            _ = v7
        case "":
            if true {
                var v8 int16 = 32767
                v8 -= (-(v8))
                switch 42 {
                case 736:
                    var v9 uint32 = 2147483647
                    _ = v9
                case 2147483647:
                    var v10 uint = 42
                    _ = v10
                case 42:
                    fmt.Println(v8)
                }
                _ = v8
            }
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 uint8 = 1
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    var v14 bool = (42 > -1)
    v14 = v14
    var v15 int = 127
    _ = v14
    _ = v15
    return true
}

func main() {
    var v1 float64 = 0.0
    v2 := []int{}
    v3 := make(map[string]int)
    var v4 float64 = v1
    var v5 int16 = -1
    v6 := []int{}
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int = (fi16 ^ fi16)
        _ = fi16
        _ = v17
    }
    if true {
        v18 := v3["hello"]
        var v19 int16 = v5
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS1{F0: 42}
        var v22 uint16 = 1
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 string = ((("foo bar" + ("a" + "")) + "foo bar") + "a")
    v1 = float64(-2.25)
    v24 := GenS1{F0: 629}
    v23 = v23
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v25 bool = true
    switch "" {
    case "\n\t":
        _ = len(v3)
    default:
        v26 := GenS2{F0: v5, F1: ((!v25) || v25)}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
