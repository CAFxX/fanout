package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]int)
    var v3 bool = false
    var v4 uint32 = 4294967295
    v5 := []int{}
    v6 := GenS1{F0: 1, F1: -9223372036854775808}
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(fi7, fi8, fi7)
            _ = fi8
        }
        _ = fi7
    }
    switch ((("hello" + ("a" + "hello")) + ("hello" + "\n\t")) + "hello") {
    case "a":
        for fi9 := 0; fi9 < 16; fi9++ {
            if v3 {
                if true {
                    switch "" {
                    case "foo bar":
                        v10 := GenS1{F0: (fi9 ^ (fi9 * fi9)), F1: (fi9 ^ (fi9 / ((fi9) | 1)))}
                        _ = v10
                    case "a":
                        v1 += (v1 % ((37) | 1))
                    default:
                        for fi11 := 0; fi11 < 5; fi11++ {
                            v12 := GenS1{F0: (fi11 / (((^(fi9))) | 1)), F1: (fi9 ^ (fi9 & 2147483647))}
                            _ = fi11
                            _ = v12
                        }
                    }
                } else {
                    v3 = (127 >= -1)
                }
                v5 = append(v5, fi9)
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                v1 = (v1 + v1)
                _ = fi13
            }
            v14 := GenS2{F0: 4294967295}
            _ = fi9
            _ = v14
        }
    case "foo bar":
        v15 := GenS2{F0: 0}
        _ = v15
    }
    v4 += v4
    v3 = (1 <= 2)
    var v16 uint32 = v4
    for fi17 := 0; fi17 < 9; fi17++ {
        if len(v5) > 0 {
            v5[0] = (-(fi17))
        }
        var v18 string = "foo bar"
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
