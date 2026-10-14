package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
func genFn0(p0 uint32) uint8 {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: "hello", F1: (fi5 - (fi5 & -981))}
        _ = fi5
        _ = v6
    }
    var v7 uint32 = p0
    var v8 uint32 = v7
    v8 += (v7 % ((p0) | 1))
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 int = 127
    var v9 int32 = 2147483647
    var v10 int64 = 1
    switch "a" {
    case "hello":
        var v11 uint = 0
        _ = v11
    case "foo bar":
        var v12 uint = 2
        _ = v12
    default:
        v10 += 9223372036854775807
    }
    v10 -= (v10 >> ((1) & 63))
    v13 := []int{}
    var v14 bool = true
    var v15 int16 = int16(v10)
    v15 = (^(v15))
    for fi16 := 0; fi16 < 9; fi16++ {
        v9 += v9
        _ = fi16
    }
    v17 := v2[2147483648]
    v17 = v14
    switch -1 {
    case -334:
        for fi18 := 0; fi18 < 5; fi18++ {
            v10 += v10
            if (fi18 > 128) {
                for fi19 := 0; fi19 < 5; fi19++ {
                    if (-1 <= -106) {
                        v17 = (fi19 <= fi19)
                        switch (fi19 - (v4 % ((v4) | 1))) {
                        case 32768:
                            var v20 bool = false
                            _ = v20
                        case 279:
                            switch "" {
                            case "a":
                                var v21 bool = ("a" < "foo bar")
                                _ = v21
                            case "\n\t":
                                var v22 int64 = -1
                                _ = v22
                            default:
                                var v23 int = 1
                                _ = v23
                            }
                        case 9223372036854775807:
                            v24 := GenS1{F0: "hello", F1: fi18}
                            _ = v24
                        }
                        if true {
                            var v25 int32 = 1
                            v26 := GenS1{F0: ("foo bar" + "hello"), F1: v4}
                            v14 = false
                            v25 = (v25 * v25)
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS1{F0: "foo bar", F1: 162}
                        _ = v27
                    }
                    _ = fi19
                }
            } else {
                fmt.Println(v4)
                switch v4 {
                case 32768:
                    if false {
                        var v28 int8 = 42
                        var v29 float32 = (0.1 / (3.14159 / -2.25))
                        v13 = append(v13, 398)
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS1{F0: "hello", F1: (fi18 - (-(v4)))}
                        _ = v30
                    }
                case 507:
                    var v31 uint32 = 4294967295
                    _ = v31
                default:
                    for fi32 := 0; fi32 < 15; fi32++ {
                        var v33 bool = v14
                        _ = fi32
                        _ = v33
                    }
                }
                var v34 int16 = v15
                _ = v34
            }
            _ = fi18
        }
    case 32767:
        if len(v13) > 0 {
            v13[0] = 0
        }
    case -646:
        v15 += 1000
    }
    fmt.Println(v14, v17, v17)
    switch ("hello" + ("" + "hello")) {
    case "hello":
        switch 708 {
        case -9223372036854775808:
            switch "hello" {
            case "foo bar":
                var v35 bool = v17
                _ = v35
            case "hello":
                v36 := GenS1{F0: "\n\t", F1: (v4 * (v4 * 9223372036854775807))}
                _ = v36
            case "":
                for fi37 := 0; fi37 < 16; fi37++ {
                    for fi38 := 0; fi38 < 17; fi38++ {
                        var v39 int8 = -54
                        _ = fi38
                        _ = v39
                    }
                    _ = fi37
                }
            default:
                var v40 int8 = 0
                _ = v40
            }
        }
    case "":
        var v41 uint = 42
        _ = v41
    default:
        var v42 int32 = v9
        _ = v42
    }
    v17 = (65535 > 255)
    for fi43 := 0; fi43 < 5; fi43++ {
        v17 = true
        _ = fi43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
