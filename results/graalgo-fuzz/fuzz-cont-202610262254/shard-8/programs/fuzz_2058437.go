package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int = 9223372036854775807
    var v2 uint64 = 4294967295
    var v3 uint64 = v2
    var v4 string = ""
    fmt.Println(v4)
    var v5 bool = (662 != 65535)
    if v5 {
        fmt.Println(v5, v5)
        v6 := []int{(v1 / ((v1) | 1)), (v1 | v1), (v1 ^ -286)}
        _ = len(v6)
        _ = v6
    } else {
        if false {
            v1 = 324
            var v7 string = v4
            v8 := []int{1, (v1 - (v1 | (^(v1)))), v1}
            _ = v7
            _ = v8
        }
        v2 += (v2 & (v3 << ((42) & 63)))
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v2 = 18446744073709551615
        for fi10 := 0; fi10 < 8; fi10++ {
            switch 127 {
            case -9223372036854775808:
                switch "a" {
                case "":
                    v5 = (4294967295 == 0)
                case "\n\t":
                    v5 = ((v2 | (v2 % ((v2) | 1))) != v2)
                default:
                    v4 = ((v4 + "a") + v4)
                }
            }
            v11 := GenS1{F0: (v2 << ((4294967295) & 63))}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: (^(v2))}
        _ = fi12
        _ = v13
    }
    var v14 bool = (true && (278 >= 1))
    if (!(!(2 == 0))) {
        fmt.Println(v14)
        var v15 int8 = 42
        fmt.Println(v5, v2)
        _ = v15
    } else {
        var v16 float64 = 0.0
        switch v4 {
        case "a":
            var v17 uint = 185
            _ = v17
        case "":
            var v18 uint64 = 4294967295
            _ = v18
        case "hello":
            v3 = v2
        }
        if true {
            v19 := []int{(v1 * (v1 * 42)), (v1 * (v1 + 956)), (v1 ^ 128)}
            var v20 uint = 4294967295
            var v21 int16 = 41
            _ = v19
            _ = v20
            _ = v21
        } else {
            v14 = (1 != 42)
            var v22 int8 = 0
            _ = v22
        }
        _ = v16
    }
    v5 = false
    for fi23 := 0; fi23 < 11; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            fmt.Println(v2, v4, fi24)
            _ = fi24
        }
        _ = fi23
    }
    v25 := []int{v1, v1, 128, (v1 ^ (^(v1)))}
    v4 = v4
    var v26 float32 = -2.25
    var v27 uint16 = 255
    v1 = v1
    var v28 uint16 = v27
    var v29 uint16 = v27
    if ((!v14) || true) {
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
