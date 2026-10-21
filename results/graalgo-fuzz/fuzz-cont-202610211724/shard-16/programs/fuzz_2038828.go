package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 uint16 = 879
    var v4 int64 = 42
    v5 := make(map[uint32]bool)
    var v6 int = 534
    v6 += -368
    var v7 int = (v6 / ((128) | 1))
    if (941 <= 909) {
        v4 += -183
    } else {
        if false {
            v4 += (^(v4))
            var v8 int16 = 42
            for fi9 := 0; fi9 < 15; fi9++ {
                v8 = (v8 >> ((255) & 63))
                for fi10 := 0; fi10 < 14; fi10++ {
                    v6 = (fi10 * (v6 | v7))
                    var v11 float32 = 0.1
                    _ = len(v2)
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
            _ = v8
        } else {
            var v12 uint8 = 53
            if true {
                var v13 string = ""
                switch v13 {
                case "\n\t":
                    var v14 float32 = (3.14159 + 1.5)
                    _ = v14
                default:
                    v13 = "a"
                }
                _ = v13
            } else {
                switch (v7 & v6) {
                case 132:
                    if (2147483647 >= 2147483647) {
                        v2 = append(v2, "a")
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                    }
                case -9223372036854775808:
                    v12 = v12
                case 42:
                    _ = len(v2)
                default:
                    v3 += 657
                }
                fmt.Println(v12, v4)
            }
            for fi15 := 0; fi15 < 5; fi15++ {
                fmt.Println(v4)
                switch "hello" {
                case "a":
                    v16 := GenS1{F0: (1.5 < 0.1), F1: v6}
                    _ = v16
                case "\n\t":
                    var v17 float64 = (3.14159 * -2.25)
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 2; fi18++ {
                        switch "a" {
                        case "a":
                            if false {
                                var v19 int8 = -69
                                _ = v19
                            } else {
                                fmt.Println(v7, v3)
                            }
                        default:
                            var v20 float32 = 1.5
                            _ = v20
                        }
                        v21 := GenS1{F0: true, F1: v7}
                        _ = fi18
                        _ = v21
                    }
                }
                _ = fi15
            }
            _ = v12
        }
        v22 := v1[(("foo bar" + "") + "foo bar")]
        _ = v22
    }
    switch v7 {
    case 0:
        fmt.Println(v7, v3)
    case 842:
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := GenS1{F0: ((v4 ^ v4) != v4), F1: v7}
            for fi25 := 0; fi25 < 13; fi25++ {
                v6 = fi23
                if false {
                    var v26 bool = ((fi25 / ((v7) | 1)) <= fi25)
                    if ((!(((!(-544 >= 0)) && false) || ((42 <= 294) || ((0.1 * 3.14159) > ((3.14159 + 0.1) + (3.14159 - float64(((3.14159 * float32((0.0 + 1e10))) / -2.25)))))))) && true) {
                        if (!(0.1 < 1.5)) {
                            var v27 int8 = -103
                            _ = v27
                        }
                    } else {
                        v26 = v26
                        delete(v5, 230)
                        var v28 bool = (42 < 2)
                        fmt.Println(v7)
                        switch fi25 {
                        case -2147483648:
                            var v29 uint = 2
                            _ = v29
                        default:
                            v28 = v26
                        }
                        v1[(("" + ("a" + "hello")) + ("a" + "a"))] = "foo bar"
                        _ = v28
                    }
                    _ = v26
                }
                var v30 int32 = -521
                _ = fi25
                _ = v30
            }
            var v31 int64 = v4
            _ = fi23
            _ = v24
            _ = v31
        }
    default:
        v7 = -1
    }
    v7 = (^(v6))
    switch ("hello" + ("\n\t" + "hello")) {
    case "hello":
        v6 = -163
    case "a":
        if true {
            v4 = 42
            var v32 int = 1
            _ = v32
        }
    case "\n\t":
        v7 += 32767
    default:
        var v33 int8 = -1
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 int = (fi34 ^ v6)
        _ = fi34
        _ = v35
    }
    v6 += (-(v7))
    _ = len(v2)
    v6 = v7
    var v36 int = (-(v7))
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v37 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
