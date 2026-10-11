package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }

func main() {
    var v1 float32 = ((1e10 * 0.0) + 1.5)
    var v2 int64 = 469
    var v3 string = (("\n\t" + "hello") + "")
    v4 := make(map[int]bool)
    switch v3 {
    case "":
        v5 := GenS1{F0: 501, F1: (v3 + (((v3 + v3) + "") + "foo bar"))}
        _ = v5
    case "\n\t":
        for fi6 := 0; fi6 < 15; fi6++ {
            v3 = v3
            switch (((v3 + ((v3 + v3) + "hello")) + ((v3 + v3) + v3)) + v3) {
            case "":
                var v7 int16 = 42
                _ = v7
            default:
                switch (fi6 + -2147483648) {
                case -2147483648:
                    for fi8 := 0; fi8 < 9; fi8++ {
                        v9 := GenS1{F0: 42, F1: v3}
                        fmt.Println(fi6)
                        var v10 int8 = 127
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                case 643:
                    var v11 uint = 0
                    _ = v11
                }
            }
            _ = fi6
        }
    }
    v3 = ("" + v3)
    switch v3 {
    case "a":
        v4[155] = true
    case "hello":
        v12 := GenS1{F0: 900, F1: "hello"}
        _ = v12
    case "foo bar":
        var v13 int16 = 32767
        _ = v13
    }
    v2 += -9223372036854775808
    v4[-131] = (0 < 128)
    v2 = v2
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS1{F0: 968, F1: v3}
        _ = fi14
        _ = v15
    }
    var v16 float32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
