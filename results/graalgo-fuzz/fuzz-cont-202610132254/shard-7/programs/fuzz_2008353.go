package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []int{896, -2147483648, 122}
    var v2 string = ("foo bar" + "hello")
    v3 := make(map[string]int)
    var v4 bool = true
    var v5 uint = 643
    var v6 uint32 = 4294967295
    var v7 int16 = -603
    if (((v2 + v2) + ((v2 + v2) + "\n\t")) <= v2) {
        v3["hello"] = 0
        switch 42 {
        case 490:
            fmt.Println(v2, v6)
        case 0:
            if v4 {
                v8 := GenS1{F0: v2, F1: 127}
                var v9 uint16 = 65535
                _ = v8
                _ = v9
            }
        default:
            switch ("" + "foo bar") {
            case "":
                fmt.Println(v6)
            case "hello":
                for fi10 := 0; fi10 < 5; fi10++ {
                    v5 += v5
                    _ = fi10
                }
            case "foo bar":
                v4 = ((v5 * 0) != v5)
            }
        }
    }
    v2 = ""
    v11 := GenS2{F0: v4}
    switch 42 {
    case -613:
        v5 = (v5 + (v5 / (((^(v5))) | 1)))
    case 504:
        fmt.Println(v2)
    case 4294967295:
        fmt.Println(v7)
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint16 = 436
        _ = fi12
        _ = v13
    }
    var v14 int32 = 0
    _ = len(v3)
    _ = len(v3)
    if v4 {
        v15 := GenS3{F0: 230}
        switch -881 {
        case 0:
            var v16 float32 = 1e10
            _ = v16
        default:
            var v17 string = v2
            _ = v17
        }
        _ = v15
    } else {
        if true {
            if v4 {
                var v18 uint = 18446744073709551615
                fmt.Println(v18)
                _ = v18
            } else {
                var v19 uint32 = v6
                _ = v19
            }
            fmt.Println(v5)
            var v20 uint32 = 42
            _ = v20
        } else {
            v14 -= 1
        }
    }
    v21 := GenS1{F0: ("hello" + ("foo bar" + (v2 + v2))), F1: 7}
    if true {
        v1 = append(v1, -2147483648)
        if v4 {
            delete(v3, "hello")
            fmt.Println(v6)
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        } else {
            if false {
                v14 = v14
            }
            v22 := GenS2{F0: ((34 != 42) || v4)}
            for fi23 := 0; fi23 < 9; fi23++ {
                v6 = (v6 & v6)
                fmt.Println(v4, v4, v2)
                _ = fi23
            }
            _ = v22
        }
    }
    switch 9223372036854775807 {
    case 4294967295:
        switch 2147483647 {
        case 846:
            for fi24 := 0; fi24 < 6; fi24++ {
                _ = len(v3)
                _ = fi24
            }
        case 4294967295:
            var v25 float64 = float64((0.1 + (-2.25 * 1.5)))
            _ = v25
        }
    default:
        fmt.Println(v6, v14)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
