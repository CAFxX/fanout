package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0(p0 int64, p1 uint8) uint64 {
    var v6 string = ""
    var v7 uint = 255
    v8 := []int{-1, -572, -9223372036854775808}
    var v9 int16 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 4
    v2 := []string{"\n\t", "foo bar", ("a" + "hello"), "hello"}
    var v3 uint = 890
    var v4 int8 = -38
    v5 := make(map[uint32]int)
    switch "\n\t" {
    case "a":
        var v10 float64 = ((1e10 - 0.0) * 1.5)
        _ = v10
    case "\n\t":
        var v11 uint32 = 866
        _ = v11
    case "":
        if false {
            var v12 int64 = 0
            _ = v12
        }
    default:
        v4 = v4
    }
    switch "\n\t" {
    case "a":
        v1 = v1
    case "foo bar":
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v4 = v4
        } else {
            v3 -= 42
            delete(v5, 4294967295)
        }
    case "\n\t":
        if (!true) {
            v13 := GenS1{F0: (0 >= 2), F1: 42}
            var v14 int = -47
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        delete(v5, 4294967295)
        _ = fi15
    }
    v16 := v5[2147483648]
    if (0 != 65535) {
        v2 = append(v2, "\n\t")
    }
    v17 := GenS1{F0: (343 >= 42), F1: 200}
    _ = len(v5)
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 float32 = float32(((0.0 * float64(float32(0.1))) + (((((0.0 * ((1e10 * -2.25) - 1e10)) + 1e10) + (1.5 + 1.5)) / 0.1) - ((3.14159 + 0.0) + 1.5))))
        _ = fi18
        _ = v19
    }
    if true {
        var v20 float64 = 0.1
        v21 := GenS1{F0: (false || (!true)), F1: 42}
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int64 = 430
        _ = fi22
        _ = v23
    }
    var v24 bool = (!(false && (false || false)))
    fmt.Println(v16, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
