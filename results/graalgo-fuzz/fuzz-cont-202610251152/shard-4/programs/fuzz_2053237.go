package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint16; F1 int16; F2 string }
func genFn0(p0 int8) uint16 {
    if true {
        v7 := GenS1{F0: 42, F1: 42}
        _ = v7
    } else {
        v8 := GenS2{F0: 0, F1: 0}
        var v9 uint = 280
        v9 = v9
        _ = v8
        _ = v9
    }
    var v10 uint64 = 829
    v11 := GenS3{F0: 2, F1: 456, F2: "hello"}
    _ = p0
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 string = "foo bar"
    var v2 float64 = -2.25
    v3 := make(map[int]bool)
    var v4 uint16 = 341
    var v5 int = 4294967295
    var v6 float32 = 1.5
    switch v1 {
    case "\n\t":
        v4 = (v4 % ((v4) | 1))
    case "foo bar":
        var v12 int = -628
        _ = v12
    case "hello":
        if (true || false) {
            v4 = 0
            delete(v3, 2147483647)
            switch v1 {
            case "\n\t":
                switch 9223372036854775807 {
                case 470:
                    var v13 uint32 = 4294967295
                    _ = v13
                }
            }
        }
    }
    var v14 uint = 312
    if false {
        v1 = ""
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        v2 -= v2
        var v16 float64 = (0.0 + v2)
        switch (-(fi15)) {
        case 32768:
            var v17 uint32 = 190
            _ = v17
        case 2147483647:
            v5 -= v5
        }
        _ = fi15
        _ = v16
    }
    v2 += float64((v6 - v6))
    var v18 float64 = v2
    switch 62 {
    case -2147483648:
        v19 := []int{v5, (v5 % ((v5) | 1))}
        _ = v19
    case 439:
        v20 := []int{v5, v5, 32768, (v5 * v5)}
        _ = v20
    case 289:
        v3[v5] = true
    }
    fmt.Println(v5, v6)
    v6 = (v6 + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
