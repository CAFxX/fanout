package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint; F2 int }
func genFn0(p0 bool) int {
    p0 = ((255 == 65535) && (p0 || p0))
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        p0 = (2147483647 > 919)
    default:
        var v7 int8 = 54
        _ = v7
    }
    if true {
        if ((255 <= 4294967295) || (47 <= 2)) {
            p0 = ((!(3 <= 255)) && ((3.14159 + -2.25) > (0.1 + 3.14159)))
            p0 = (1 >= -519)
            var v8 uint16 = 0
            _ = v8
        }
        fmt.Println(p0)
    } else {
        p0 = p0
    }
    _ = p0
    return 42
}

func main() {
    v1 := []uint32{720, 42}
    var v2 int8 = 43
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    v5 := make(map[string]string)
    var v6 int64 = 9223372036854775807
    v6 -= v6
    v6 = int64(v2)
    _ = len(v1)
    if (36 > 42) {
        var v9 float64 = 0.1
        for fi10 := 0; fi10 < 1; fi10++ {
            v6 += v6
            _ = fi10
        }
        _ = v9
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 bool = false
            var v13 int32 = 2147483647
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS1{F0: 289}
            _ = fi14
            _ = v15
        }
    }
    _ = len(v4)
    switch "a" {
    case "a":
        var v16 uint = 1
        _ = v16
    }
    var v17 string = "hello"
    var v18 float64 = 0.0
    _ = len(v4)
    switch "hello" {
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        v19 := GenS2{F0: 127, F1: 65535, F2: 128}
        _ = v19
    }
    fmt.Println(v6, v17, v6)
    switch 2147483647 {
    case -9223372036854775808:
        var v20 int = 127
        _ = v20
    case 42:
        var v21 bool = (!(!false))
        _ = v21
    }
    var v22 uint64 = 151
    fmt.Println(v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
