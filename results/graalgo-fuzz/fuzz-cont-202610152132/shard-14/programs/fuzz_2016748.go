package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 215
    v3 := make(map[int]int)
    var v4 uint8 = 110
    var v5 uint8 = v4
    switch "a" {
    case "":
        for fi6 := 0; fi6 < 1; fi6++ {
            switch "\n\t" {
            case "\n\t":
                v4 = 1
            case "hello":
                var v7 uint64 = 18446744073709551615
                _ = v7
            }
            _ = fi6
        }
    case "a":
        v8 := GenS1{F0: 759}
        _ = v8
    }
    var v9 uint8 = (v4 / ((v4) | 1))
    v10 := []int{127, -749}
    delete(v1, 426)
    if true {
        if true {
            for fi11 := 0; fi11 < 10; fi11++ {
                v9 += v5
                for fi12 := 0; fi12 < 8; fi12++ {
                    if false {
                        var v13 float64 = (0.0 * -2.25)
                        v14 := GenS1{F0: 563}
                        _ = v13
                        _ = v14
                    }
                    v4 = (v4 + 1)
                    var v15 uint32 = 0
                    fmt.Println(v2, v9, v2)
                    _ = fi12
                    _ = v15
                }
                var v16 int64 = 42
                _ = fi11
                _ = v16
            }
        } else {
            var v17 int32 = v2
            delete(v1, 9223372036854775807)
            _ = v17
        }
        v5 = v9
    }
    v18 := GenS2{F0: -654, F1: 18446744073709551615, F2: 1}
    if ("a" != "") {
        v19 := GenS1{F0: 255}
        _ = v19
    }
    if (("a" == "") || ((false || (true || (!true))) || ((!false) && true))) {
        switch "a" {
        case "\n\t":
            if (1 <= 42) {
                var v20 uint8 = v9
                for fi21 := 0; fi21 < 20; fi21++ {
                    v20 = v5
                    _ = fi21
                }
                fmt.Println(v2, v4)
                _ = v20
            } else {
                v22 := GenS2{F0: 1, F1: 4294967295, F2: 940}
                _ = v22
            }
        case "a":
            var v23 int = 135
            _ = v23
        case "":
            if false {
                for fi24 := 0; fi24 < 7; fi24++ {
                    if true {
                        v4 = (^(v9))
                    }
                    _ = fi24
                }
            }
        default:
            v4 -= (v4 << ((42) & 63))
        }
        var v25 int = 4294967295
        _ = v25
    }
    if ((v4 ^ (v9 ^ 2)) == v4) {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        v26 := GenS1{F0: 32767}
        v27 := GenS1{F0: 1}
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
