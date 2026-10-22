package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int16 = -32768
    var v2 int64 = 430
    var v3 uint32 = 2147483647
    v1 = (-(v1))
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := []int{fi4}
        switch fi4 {
        case 698:
            if (-279 < -1) {
                var v6 float32 = -2.25
                _ = v6
            } else {
                if ("" <= "foo bar") {
                    v1 -= (v1 << ((255) & 63))
                    v3 += (v3 >> ((638) & 63))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    if (v1 < v1) {
                        var v7 uint32 = (v3 + 1)
                        _ = v7
                    } else {
                        switch "hello" {
                        case "hello":
                            if false {
                                v2 = (-(v2))
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    _ = len(v5)
                                case "foo bar":
                                    v1 -= v1
                                }
                                v1 -= int16(v2)
                                v2 = -1
                            }
                        case "a":
                            v8 := GenS1{F0: ("a" + "")}
                            _ = v8
                        default:
                            v1 -= 0
                        }
                    }
                    for fi9 := 0; fi9 < 17; fi9++ {
                        if true {
                            fmt.Println(v3, v2)
                            v2 += (v2 / ((1) | 1))
                            var v10 int = (fi4 << ((858) & 63))
                            var v11 uint8 = 127
                            var v12 string = ("\n\t" + ("\n\t" + ("\n\t" + ("a" + ("hello" + "\n\t")))))
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                        _ = fi9
                    }
                }
                v1 = (v1 | 726)
                var v13 float32 = -2.25
                _ = v13
            }
        case 4294967295:
            if len(v5) > 0 {
                v5[0] = (fi4 ^ fi4)
            }
        }
        _ = fi4
        _ = v5
    }
    v14 := GenS1{F0: "foo bar"}
    v3 += v3
    v15 := GenS2{F0: 0}
    fmt.Println(v1, v1, v3)
    var v16 string = (("a" + "hello") + "hello")
    var v17 int64 = v2
    if false {
        switch 9223372036854775807 {
        case 462:
            switch "foo bar" {
            case "":
                v3 += v3
            case "\n\t":
                fmt.Println(v3, v17, v2)
            case "hello":
                var v18 uint = 0
                _ = v18
            default:
                v17 = -1
            }
        case 2147483647:
            var v19 float32 = (0.1 - (float32(-2.25) / (1e10 * (-2.25 / 1e10))))
            _ = v19
        }
        v2 -= v2
        v1 -= (^(v1))
    } else {
        v16 = v16
        var v20 float64 = ((-2.25 - ((0.1 * (1.5 + 0.0)) - 0.1)) * (-2.25 - 3.14159))
        v16 = "a"
        _ = v20
    }
    var v21 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
}
