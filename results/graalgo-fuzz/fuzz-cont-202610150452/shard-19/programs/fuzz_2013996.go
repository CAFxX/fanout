package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 string; F1 uint; F2 int8 }

func main() {
    var v1 float32 = 0.0
    var v2 string = ""
    var v3 int32 = 42
    fmt.Println(v2, v1)
    switch "hello" {
    case "\n\t":
        fmt.Println(v3)
    case "a":
        v4 := GenS1{F0: 0, F1: v3}
        _ = v4
    }
    var v5 string = ""
    v6 := []int{-9223372036854775808}
    v1 = v1
    for fi7 := 0; fi7 < 11; fi7++ {
        if false {
            switch v5 {
            case "a":
                var v8 uint16 = 2
                _ = v8
            }
            v5 = (v2 + "hello")
            v3 -= (v3 & 447)
        } else {
            v2 = v2
            v1 -= v1
            fmt.Println(fi7, v1, v2)
        }
        _ = fi7
    }
    switch v5 {
    case "":
        v9 := GenS1{F0: 449, F1: v3}
        _ = v9
    case "foo bar":
        if (127 != -1) {
            var v10 uint = 65535
            for fi11 := 0; fi11 < 20; fi11++ {
                v6 = append(v6, -9223372036854775808)
                _ = fi11
            }
            _ = v10
        } else {
            v6 = append(v6, -9223372036854775808)
            fmt.Println(v5, v1)
            for fi12 := 0; fi12 < 18; fi12++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi12
            }
        }
    }
    var v13 float32 = (3.14159 * -2.25)
    v13 = v13
    for fi14 := 0; fi14 < 20; fi14++ {
        if true {
            v3 = (v3 ^ v3)
        } else {
            fmt.Println(v5)
            var v15 int8 = 42
            _ = v15
        }
        v16 := GenS2{F0: v2, F1: 287, F2: -1}
        _ = fi14
        _ = v16
    }
    v17 := GenS2{F0: "\n\t", F1: 65535, F2: 42}
    fmt.Println(v13, v3)
    v18 := GenS2{F0: "\n\t", F1: 215, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
