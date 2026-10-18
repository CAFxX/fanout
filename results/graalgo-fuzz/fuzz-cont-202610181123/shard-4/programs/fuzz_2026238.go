package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int }
func genFn0(p0 int32, p1 bool, p2 uint) uint {
    v7 := GenS1{F0: p1, F1: 891, F2: -1000}
    var v8 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 654
}

func main() {
    var v1 float64 = (0.1 + 0.0)
    var v2 float32 = -2.25
    v3 := make(map[string]string)
    var v4 bool = (false || (!false))
    var v5 int16 = -32768
    var v6 int = -535
    var v9 string = ("a" + ("a" + ("" + (("foo bar" + ("" + ("a" + "hello"))) + ""))))
    var v10 uint8 = 1
    v5 = (v5 | -32768)
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS2{F0: 255}
        v13 := GenS1{F0: true, F1: 4294967295, F2: -895}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{2147483647, v6, v6, v6}
    switch v9 {
    case "hello":
        if false {
            if ("" <= "foo bar") {
                var v15 int32 = 286
                if (831 < 2) {
                    v16 := GenS3{F0: 42}
                    _ = v16
                } else {
                    var v17 string = v9
                    _ = v17
                }
                _ = v15
            }
        } else {
            v18 := GenS4{F0: 1, F1: 0, F2: (v6 - v6)}
            _ = v18
        }
    case "\n\t":
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    }
    switch v9 {
    case "a":
        var v19 bool = (true && v4)
        _ = v19
    case "":
        var v20 uint8 = v10
        _ = v20
    }
    switch "a" {
    case "hello":
        v4 = true
    case "\n\t":
        v1 = v1
    default:
        v21 := v3["\n\t"]
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        if len(v14) > 0 {
            v14[0] = (fi22 ^ -588)
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
