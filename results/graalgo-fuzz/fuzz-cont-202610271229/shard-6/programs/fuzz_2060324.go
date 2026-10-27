package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 int }

func main() {
    var v1 float64 = 0.1
    var v2 bool = true
    var v3 float32 = 0.1
    var v4 float32 = 0.0
    v5 := make(map[uint32]int)
    if false {
        if ((0.1 == 1.5) && (0 == 42)) {
            delete(v5, 4294967295)
            fmt.Println(v2, v4)
            for fi6 := 0; fi6 < 9; fi6++ {
                v2 = true
                v1 = v1
                _ = fi6
            }
        }
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 string = "a"
                fmt.Println(v9, fi7)
                switch (v9 + "\n\t") {
                case "":
                    v10 := GenS1{F0: "hello", F1: (!v2), F2: 42}
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
            v11 := GenS2{F0: (v4 != v3), F1: 4294967295, F2: 2147483647}
            var v12 uint = 2
            _ = fi7
            _ = v11
            _ = v12
        }
        switch 735 {
        case 437:
            if (1e10 >= 3.14159) {
                for fi13 := 0; fi13 < 20; fi13++ {
                    v1 = v1
                    if ((127 >= 69) || (!v2)) {
                        for fi14 := 0; fi14 < 3; fi14++ {
                            v15 := []int{fi14, fi13}
                            _ = fi14
                            _ = v15
                        }
                    }
                    v16 := []int{(fi13 - 237), 0}
                    _ = fi13
                    _ = v16
                }
            } else {
                var v17 int16 = 558
                _ = v17
            }
        case 32768:
            switch 252 {
            case -402:
                v2 = true
            default:
                v18 := GenS2{F0: (("hello" + "a") == "foo bar"), F1: 752, F2: 2147483648}
                _ = v18
            }
        case -804:
            if (-1 >= 32767) {
                v1 = ((((1.5 + (-2.25 - v1)) * (v1 / v1)) + (3.14159 + 0.1)) / v1)
                v3 = 1e10
                v3 = 0.1
            }
        }
    } else {
        for fi19 := 0; fi19 < 9; fi19++ {
            v5[42] = 32768
            v4 = -2.25
            _ = fi19
        }
    }
    v20 := GenS3{F0: 1}
    fmt.Println(v3, v1, v3)
    var v21 uint8 = 127
    v3 = v4
    fmt.Println(v1)
    v2 = (-2147483648 > 4294967295)
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS4{F0: 347, F1: (v21 % (((^(v21))) | 1))}
        if true {
            var v24 uint64 = 2
            var v25 uint32 = 4294967295
            v26 := GenS5{F0: (466 < 18446744073709551615)}
            _ = v24
            _ = v25
            _ = v26
        } else {
            var v27 uint16 = 0
            _ = v27
        }
        if false {
            v28 := []int{-2147483648, -9223372036854775808, (-(fi22))}
            _ = v28
        } else {
            v29 := []int{fi22, 4294967295}
            _ = v29
        }
        _ = fi22
        _ = v23
    }
    v30 := GenS6{F0: 42, F1: (^(v21))}
    var v31 int = 2147483647
    v32 := []int{v31}
    var v33 int32 = -1
    v34 := GenS7{F0: v31}
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        fmt.Println(v21, v1, v1)
    }
    if (v2 || false) {
        delete(v5, 42)
        v33 = v33
        if (-9223372036854775808 <= 9223372036854775807) {
            v1 = float64(v4)
            var v35 int64 = 1
            _ = v35
        }
    }
    v21 -= 255
    var v36 int16 = 0
    v37 := GenS6{F0: 0, F1: (v21 - (v21 ^ v21))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
