package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
func genFn0(p0 bool, p1 uint32, p2 uint16) uint {
    v7 := []int{}
    p2 = p2
    var v8 float64 = (((0.1 / ((1e10 / (1e10 - 0.1)) + 0.0)) - 3.14159) - 3.14159)
    switch -9223372036854775808 {
    case -202:
        v8 -= v8
    case -471:
        fmt.Println(p2, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 483
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]int)
    var v3 int16 = 0
    v4 := []string{"", "foo bar"}
    var v5 int16 = 1000
    var v6 string = "hello"
    for fi9 := 0; fi9 < 5; fi9++ {
        fmt.Println(v5, fi9)
        _ = fi9
    }
    v3 = (v3 % ((v5) | 1))
    switch 4294967295 {
    case 4294967295:
        var v10 uint32 = 2
        _ = v10
    }
    if len(v4) > 0 {
        v4[0] = v6
    }
    _ = len(v2)
    switch v6 {
    case "":
        _ = len(v4)
    case "a":
        switch (v6 + v6) {
        case "\n\t":
            var v11 uint8 = 127
            _ = v11
        case "":
            v1 -= v1
        case "a":
            _ = len(v2)
        default:
            if (false && (18446744073709551615 >= 2)) {
                _ = len(v4)
            }
        }
    case "hello":
        delete(v2, ("a" + v6))
    }
    if false {
        v3 += (v5 - v3)
        v6 = "foo bar"
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        if true {
            delete(v2, v6)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v13 := GenS1{F0: -530, F1: fi12, F2: 871}
            _ = v13
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
