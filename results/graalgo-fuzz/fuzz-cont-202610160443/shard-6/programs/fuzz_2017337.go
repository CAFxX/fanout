package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = (3.14159 - 1.5)
    var v2 uint = 255
    v3 := []string{"foo bar", "", (("" + ("foo bar" + "\n\t")) + ""), "\n\t"}
    var v4 uint64 = 42
    v5 := GenS1{F0: 1, F1: 763}
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v2 = v2
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int16 = 1000
        var v8 float32 = float32(float64(1.5))
        _ = fi6
        _ = v7
        _ = v8
    }
    if (true && ((v4 / ((1) | 1)) >= (v4 & (^(v4))))) {
        v9 := GenS1{F0: -1, F1: 1}
        v1 += (v1 + 1.5)
        var v10 int64 = 0
        _ = v10
        _ = v9
    }
    if ((42 <= -1) || (1 != 42)) {
        var v11 uint64 = uint64(v2)
        switch 127 {
        case -2147483648:
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 uint64 = (v11 | v11)
                v13 = v11
                if (true || (!false)) {
                    if true {
                        switch "" {
                        case "":
                            if true {
                                if false {
                                    v14 := GenS2{F0: "a"}
                                    _ = v14
                                }
                                var v15 uint = v2
                                _ = v15
                            }
                        }
                    } else {
                        switch 32768 {
                        case 987:
                            v2 += v2
                        }
                        v16 := GenS1{F0: -15, F1: 32767}
                        switch "a" {
                        case "hello":
                            switch "foo bar" {
                            case "\n\t":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "foo bar":
                                var v17 uint8 = 170
                                _ = v17
                            }
                        case "\n\t":
                            v4 -= (v13 / (((^(v11))) | 1))
                        case "a":
                            var v18 int = -133
                            _ = v18
                        }
                        _ = v16
                    }
                    for fi19 := 0; fi19 < 9; fi19++ {
                        if false {
                            switch "foo bar" {
                            case "a":
                                v3 = append(v3, "\n\t")
                            default:
                                if len(v3) > 0 {
                                    v3[0] = ("\n\t" + "foo bar")
                                }
                            }
                        }
                        _ = fi19
                    }
                }
                _ = fi12
                _ = v13
            }
        case 9223372036854775807:
            v1 += v1
        case 1:
            v2 = (v2 + v2)
        }
        _ = v11
    }
    var v20 int64 = 225
    var v21 bool = (0 == v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
