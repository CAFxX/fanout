package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 bool }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[int]string)
    var v3 float32 = ((-2.25 * (3.14159 + 3.14159)) / (0.0 * (float32(v1) + float32((1e10 - (v1 - (0.0 * v1)))))))
    var v4 int64 = 998
    var v5 int64 = (v4 >> ((18446744073709551615) & 63))
    v6 := GenS1{F0: 136}
    switch 32767 {
    case -401:
        switch "a" {
        case "foo bar":
            v7 := GenS1{F0: 255}
            _ = v7
        case "a":
            var v8 bool = false
            _ = v8
        case "\n\t":
            var v9 float64 = v1
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 7; fi10++ {
            if (true || (true && (17 < 18446744073709551615))) {
                v11 := GenS2{F0: 0}
                _ = v11
            }
            _ = fi10
        }
    }
    if true {
        var v12 float32 = v3
        _ = v12
    }
    v13 := []int{}
    v5 += (v5 * v4)
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint32 = 743
        _ = fi14
        _ = v15
    }
    switch "foo bar" {
    case "\n\t":
        v16 := GenS3{F0: 2147483647}
        _ = v16
    }
    var v17 int32 = 2147483647
    v18 := GenS3{F0: 42}
    for fi19 := 0; fi19 < 19; fi19++ {
        v4 += (v4 >> ((18446744073709551615) & 63))
        if true {
            v17 = (^(v17))
            if (v3 < ((v3 + 1e10) - v3)) {
                if true {
                    if (v4 >= 9223372036854775807) {
                        switch (fi19 - fi19) {
                        case 4294967295:
                            for fi20 := 0; fi20 < 6; fi20++ {
                                if (float32(((float64(3.14159) * float64((1e10 / ((float32((v1 * 0.0)) / (v3 - ((v3 + float32(((v1 + (v1 - 0.0)) + ((v1 * (v1 - v1)) + ((v1 + v1) * float64(v3)))))) / 1.0))) / ((v3 + float32(0.1)) - 3.14159))))) + (3.14159 * (0.1 * (0.0 / (v1 - v1)))))) > v3) {
                                    var v21 int32 = 0
                                    _ = v21
                                } else {
                                    v4 = v4
                                    var v22 int = fi20
                                    var v23 uint8 = 42
                                    v13 = append(v13, fi19)
                                    _ = v22
                                    _ = v23
                                }
                                _ = fi20
                            }
                        case -1:
                            v24 := GenS3{F0: 2}
                            _ = v24
                        }
                        if (!(fi19 < 32768)) {
                            v4 += (v4 << ((18446744073709551615) & 63))
                            var v25 int = fi19
                            var v26 int32 = 1
                            _ = v25
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 16; fi27++ {
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                var v28 uint8 = 127
                                fmt.Println(fi19, v28, fi27)
                                _ = fi27
                                _ = v28
                            }
                            v17 -= 0
                            var v29 uint64 = 18446744073709551615
                            _ = v29
                        }
                        v17 += 754
                    } else {
                        fmt.Println(fi19, v5)
                        v1 -= 1.5
                    }
                    var v30 int64 = (-(v4))
                    fmt.Println(v1, fi19, v4)
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 18; fi31++ {
                        var v32 int32 = (v17 & v17)
                        var v33 int8 = 1
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    v4 = -26
                }
            }
            switch (("" + "a") + ("\n\t" + "hello")) {
            case "hello":
                v1 += v1
            case "\n\t":
                v3 = ((((float32((float64((float32(((1.5 + v1) / v1)) - (v3 * v3))) * float64((((((float32(((-2.25 * (((float64(0.1) / 1e10) * (v1 - v1)) - v1)) * v1)) + v3) * (0.0 + v3)) + (v3 - v3)) * v3) + v3)))) - (v3 / float32((v1 / -2.25)))) + ((-2.25 + -2.25) / (v3 + (v3 + v3)))) + 3.14159) - (0.0 - (-2.25 * 0.1)))
            case "":
                v17 = (v17 * (v17 % ((42) | 1)))
            default:
                var v34 uint32 = 814
                _ = v34
            }
        }
        var v35 float64 = v1
        _ = fi19
        _ = v35
    }
    v36 := GenS4{F0: 1, F1: false}
    var v37 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
