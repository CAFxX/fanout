package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint8; F1 int32; F2 string }

func main() {
    var v1 uint32 = 705
    var v2 uint8 = 42
    v3 := make(map[int]uint32)
    switch 128 {
    case 32767:
        v2 = 128
    }
    v1 = 890
    var v4 float64 = (float64(float32((-2.25 + (0.0 / -2.25)))) + ((0.0 - -2.25) * (1.5 - 0.1)))
    if true {
        v3[42] = v1
    }
    fmt.Println(v2)
    if (42 <= 0) {
        if true {
            for fi5 := 0; fi5 < 3; fi5++ {
                var v6 uint8 = 0
                v7 := GenS1{F0: 9223372036854775807, F1: ((("\n\t" + ("\n\t" + ("foo bar" + ("foo bar" + (("foo bar" + "") + ("a" + "")))))) + "") + "a")}
                v8 := GenS1{F0: fi5, F1: ("\n\t" + "hello")}
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
            for fi9 := 0; fi9 < 18; fi9++ {
                if true {
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        switch ("\n\t" + ("" + "hello")) {
                        case "a":
                            v4 -= v4
                        case "hello":
                            v10 := []int{127}
                            _ = v10
                        case "":
                            v11 := []int{(^(fi9))}
                            _ = v11
                        }
                    case "a":
                        var v12 int64 = 42
                        _ = v12
                    case "hello":
                        v13 := []int{(fi9 - fi9), fi9, -1, (^(fi9))}
                        _ = v13
                    }
                } else {
                    var v14 uint32 = 943
                    _ = v14
                }
                v4 = float64(0.0)
                v4 = v4
                _ = fi9
            }
            v4 -= (3.14159 / 1.5)
        } else {
            for fi15 := 0; fi15 < 7; fi15++ {
                v2 = (v2 & (v2 % ((v2) | 1)))
                _ = fi15
            }
        }
    } else {
        var v16 uint64 = 904
        v17 := GenS1{F0: -338, F1: "foo bar"}
        var v18 int64 = -377
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint64 = 473
    v20 := GenS2{F0: (v2 >> ((65535) & 63)), F1: 1, F2: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
