package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }

func main() {
    var v1 int64 = -1
    var v2 uint32 = 816
    var v3 uint32 = 65535
    v4 := []int{-9223372036854775808, 1}
    v5 := []string{"\n\t", "hello", ""}
    var v6 int = -918
    for fi7 := 0; fi7 < 5; fi7++ {
        v4 = append(v4, fi7)
        _ = fi7
    }
    fmt.Println(v1, v6, v6)
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 bool = (32767 != -356)
        _ = fi8
        _ = v9
    }
    var v10 float32 = 1.5
    switch ("" + (("" + "a") + "\n\t")) {
    case "a":
        var v11 uint32 = (v3 << ((2) & 63))
        _ = v11
    case "foo bar":
        switch 9223372036854775807 {
        case -2147483648:
            if ((!(!false)) || false) {
                if len(v4) > 0 {
                    v4[0] = v6
                }
                var v12 uint64 = 42
                v12 = 875
                _ = v12
            }
        case 32768:
            v13 := GenS1{F0: v6, F1: 209}
            _ = v13
        default:
            switch (("foo bar" + "foo bar") + "a") {
            case "":
                if false {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        var v15 uint64 = 42
                        _ = fi14
                        _ = v15
                    }
                    fmt.Println(v6, v3)
                }
            case "a":
                v1 = (v1 % ((v1) | 1))
            case "hello":
                v16 := GenS1{F0: (v6 * 384), F1: 1}
                _ = v16
            default:
                for fi17 := 0; fi17 < 4; fi17++ {
                    if (true && (634 >= 4294967295)) {
                        if true {
                            var v18 uint64 = 0
                            v1 -= 754
                            _ = v18
                        }
                    }
                    _ = fi17
                }
            }
        }
    default:
        var v19 float32 = v10
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(v3, v10, v3)
        var v21 float32 = float32(1.5)
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 int = -9223372036854775808
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    if true {
        v1 = (v1 * -969)
        if ("" < "\n\t") {
            v24 := GenS1{F0: (v6 + 9223372036854775807), F1: 29}
            _ = v24
        } else {
            var v25 float64 = 0.1
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
