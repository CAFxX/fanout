package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }

func main() {
    var v1 int8 = 100
    var v2 uint16 = 2
    v3 := make(map[string]bool)
    var v4 uint = 1
    var v5 int64 = -399
    var v6 float64 = 0.1
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int8 = v1
        if (((v4 - (v4 | 255)) <= (v4 / ((790) | 1))) || (0.1 <= 1.5)) {
            v4 = (^(v4))
        } else {
            v9 := GenS1{F0: fi7, F1: "", F2: (v4 & v4)}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    v4 += v4
    switch -2147483648 {
    case 117:
        v10 := []int{-802}
        _ = v10
    case -238:
        v5 = 42
    }
    _ = len(v3)
    if false {
        v6 = (0.0 * (v6 / 0.1))
    } else {
        v1 -= v1
        v11 := GenS1{F0: int(v1), F1: ("" + "foo bar"), F2: (v4 % ((42) | 1))}
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            if (!((-(fi12)) < int(v1))) {
                switch ("a" + (("a" + (("a" + "hello") + ("foo bar" + "hello"))) + ("hello" + "a"))) {
                case "a":
                    v14 := v3[""]
                    _ = v14
                default:
                    fmt.Println(fi12, fi12)
                }
                var v15 uint64 = 98
                v2 -= v2
                _ = v15
            } else {
                v6 = v6
                switch fi12 {
                case 9223372036854775807:
                    var v16 string = ""
                    _ = v16
                case 42:
                    if false {
                        var v17 string = "\n\t"
                        var v18 uint8 = 1
                        _ = v17
                        _ = v18
                    }
                }
                v5 = v5
            }
            v19 := []int{(fi13 - 1), fi13, (fi12 * 32767)}
            delete(v3, ("foo bar" + "\n\t"))
            _ = fi13
            _ = v19
        }
        _ = fi12
    }
    v2 -= (v2 >> ((0) & 63))
    v20 := GenS1{F0: 128, F1: "\n\t", F2: (v4 * 175)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
