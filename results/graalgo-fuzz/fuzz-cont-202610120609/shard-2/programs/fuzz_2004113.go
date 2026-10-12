package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
type GenS2 struct { F0 uint8; F1 string; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 472
    var v3 string = ("" + "")
    var v4 int = 42
    v5 := make(map[uint32]bool)
    var v6 uint64 = 1
    if (42 > 42) {
        v4 = v4
        for fi7 := 0; fi7 < 6; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                if (true || (42 < 1)) {
                    v2 = v2
                    var v9 uint64 = v6
                    _ = v9
                }
                v1[(v2 | v2)] = ("a" + v3)
                _ = fi8
            }
            _ = fi7
        }
        var v10 float64 = 1.5
        _ = v10
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int64 = 148
            v1[(v2 * 210)] = "a"
            var v13 float64 = 0.0
            _ = fi11
            _ = v12
            _ = v13
        }
        fmt.Println(v6, v6)
        var v14 uint16 = 65535
        _ = v14
    }
    v15 := GenS1{F0: v6, F1: 128, F2: v4}
    v16 := GenS1{F0: v6, F1: 1, F2: (-(v4))}
    var v17 int32 = 0
    var v18 uint64 = v6
    var v19 int16 = -76
    fmt.Println(v6)
    v20 := GenS1{F0: 0, F1: 10, F2: v4}
    v21 := GenS2{F0: 234, F1: v3, F2: 2}
    switch 9223372036854775807 {
    case 2147483647:
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := v5[(v2 << ((438) & 63))]
            v24 := GenS3{F0: (v19 * v19)}
            var v25 uint8 = 199
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v4, v3, v19)
    v26 := GenS4{F0: ((((v3 + ((v3 + "") + ((v3 + "hello") + v3))) + v3) + ((v3 + v3) + "hello")) + v3), F1: 169}
    for fi27 := 0; fi27 < 9; fi27++ {
        v2 = v2
        if (v2 != v2) {
            v17 -= (v17 * v17)
            v28 := []int{75, -1, (v4 << ((499) & 63))}
            if true {
                for fi29 := 0; fi29 < 10; fi29++ {
                    v30 := GenS1{F0: (^(v6)), F1: 2, F2: (fi29 / ((-472) | 1))}
                    var v31 int = fi27
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                v2 = (^(v2))
            } else {
                for fi32 := 0; fi32 < 11; fi32++ {
                    if (1e10 > (0.0 / 3.14159)) {
                        switch v3 {
                        case "a":
                            v33 := GenS5{F0: true}
                            _ = v33
                        }
                        switch v3 {
                        case "":
                            v2 -= (v2 >> ((243) & 63))
                        case "\n\t":
                            if false {
                                var v34 float64 = 1.5
                                _ = v34
                            }
                        }
                        var v35 uint8 = 2
                        _ = v35
                    }
                    fmt.Println(v3)
                    if (255 >= 42) {
                        fmt.Println(v18, v18, v19)
                    }
                    _ = fi32
                }
            }
            _ = v28
        }
        _ = fi27
    }
    fmt.Println(v17, v4, v3)
    fmt.Println(v19, v6)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
