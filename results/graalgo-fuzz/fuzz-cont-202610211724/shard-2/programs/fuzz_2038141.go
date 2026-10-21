package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }

func main() {
    var v1 float32 = (1.5 * 0.1)
    var v2 float64 = 3.14159
    var v3 int32 = -2147483648
    var v4 uint = 42
    v5 := make(map[uint32]uint32)
    var v6 int16 = -1
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint = 2
        fmt.Println(v6, v3)
        var v9 bool = true
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-9223372036854775808, 42}
    _ = len(v5)
    v2 = 0.1
    var v11 bool = true
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            v4 = v4
            delete(v5, 42)
            if false {
                v4 = (v4 + v4)
            } else {
                if v11 {
                    _ = len(v10)
                }
            }
            _ = fi13
        }
        _ = fi12
    }
    if (!(42 >= 1)) {
        v14 := GenS1{F0: 755, F1: 0}
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 uint = v4
            if ((!v11) || v11) {
                v17 := GenS2{F0: "foo bar", F1: 2}
                fmt.Println(v1, v4, v3)
                var v18 bool = v11
                _ = v17
                _ = v18
            } else {
                switch (^(fi15)) {
                case -2147483648:
                    if (42 >= -1) {
                        v16 += v4
                    }
                }
                var v19 float64 = v2
                for fi20 := 0; fi20 < 18; fi20++ {
                    if false {
                        var v21 uint32 = 2
                        fmt.Println(fi15, v2)
                        switch "\n\t" {
                        case "\n\t":
                            switch ("" + ("\n\t" + "a")) {
                            case "a":
                                v22 := GenS3{F0: 18446744073709551615}
                                _ = v22
                            case "hello":
                                v23 := GenS1{F0: -1, F1: 42}
                                _ = v23
                            case "foo bar":
                                var v24 int32 = (v3 | v3)
                                _ = v24
                            }
                            _ = len(v5)
                            v5[(^(v21))] = (v21 | v21)
                        case "":
                            fmt.Println(v19, fi20)
                        }
                        _ = v21
                    }
                    v25 := GenS4{F0: (v16 * (v16 ^ v4))}
                    v6 = v6
                    _ = fi20
                    _ = v25
                }
                _ = v19
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v10 = append(v10, 704)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
