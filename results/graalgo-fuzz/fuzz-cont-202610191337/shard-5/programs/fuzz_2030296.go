package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 string }

func main() {
    var v1 uint8 = 2
    var v2 float32 = (0.1 - 0.0)
    var v3 uint8 = v1
    v4 := make(map[int]int)
    var v5 float64 = 1e10
    var v6 uint64 = 928
    if false {
        v4[-469] = -2147483648
        v3 += 225
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            v5 = -2.25
            switch -268 {
            case 4294967295:
                fmt.Println(v5, v5, v5)
            case -383:
                for fi8 := 0; fi8 < 18; fi8++ {
                    v9 := []int{(fi8 << ((225) & 63))}
                    _ = fi8
                    _ = v9
                }
            default:
                v10 := []int{128, (^(fi7)), fi7}
                _ = v10
            }
            v1 += v3
            _ = fi7
        }
        var v11 uint = 2
        v3 += 252
        _ = v11
    }
    var v12 int32 = 42
    switch (("hello" + ("a" + ("\n\t" + "\n\t"))) + "a") {
    case "":
        v13 := GenS1{F0: "", F1: 0, F2: "foo bar"}
        _ = v13
    case "\n\t":
        v5 = (v5 + v5)
    case "a":
        for fi14 := 0; fi14 < 9; fi14++ {
            fmt.Println(v1)
            v15 := []int{(fi14 | fi14), fi14}
            var v16 uint64 = (v6 >> ((42) & 63))
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v17 int32 = (v12 - (v12 % ((v12) | 1)))
    if (!(128 == (v1 / ((0) | 1)))) {
        v18 := GenS1{F0: "a", F1: 0, F2: "foo bar"}
        delete(v4, -72)
        fmt.Println(v17)
        _ = v18
    }
    v3 -= (v3 + 255)
    v3 = (^(v1))
    if ((false && (true || (false || false))) && (42 == 0)) {
        var v19 string = "\n\t"
        switch 387 {
        case 241:
            var v20 string = "\n\t"
            _ = v20
        case 32767:
            var v21 int64 = 1
            _ = v21
        default:
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 float64 = (0.1 - ((v5 / (0.0 + v5)) / float64(3.14159)))
                _ = fi22
                _ = v23
            }
        }
        v3 = 42
        _ = v19
    }
    if (65535 <= 65535) {
        v6 = (v6 ^ (v6 | v6))
        var v24 string = "\n\t"
        _ = v24
    }
    var v25 int32 = 42
    fmt.Println(v6, v2)
    var v26 bool = false
    if (("\n\t" < (("foo bar" + (("\n\t" + "hello") + "")) + "")) && true) {
        if v26 {
            v3 = (v3 + v1)
            var v27 bool = (4294967295 >= 2147483648)
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 12; fi28++ {
            v29 := GenS1{F0: ("hello" + "foo bar"), F1: 0, F2: ("a" + "foo bar")}
            for fi30 := 0; fi30 < 3; fi30++ {
                v31 := GenS1{F0: "", F1: -57, F2: ("a" + "")}
                switch fi28 {
                case 689:
                    v12 = (v12 | v25)
                case -180:
                    switch "a" {
                    case "a":
                        var v32 int = fi30
                        _ = v32
                    case "foo bar":
                        v17 += 297
                    case "":
                        v12 = v25
                    }
                default:
                    v1 = (v1 + (v1 | (^(v1))))
                }
                _ = fi30
                _ = v31
            }
            var v33 float64 = v5
            _ = fi28
            _ = v29
            _ = v33
        }
        v17 -= (v17 * v17)
        var v34 int32 = -338
        _ = v34
    }
    fmt.Println(v2, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
