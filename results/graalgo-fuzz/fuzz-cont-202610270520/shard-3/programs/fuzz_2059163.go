package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 int }

func main() {
    var v1 int32 = 918
    var v2 int = -853
    v3 := []int{}
    var v4 uint64 = 4294967295
    var v5 float32 = 0.1
    var v6 uint = 873
    v4 = 104
    fmt.Println(v6)
    switch (v2 / (((v2 - v2)) | 1)) {
    case 902:
        v5 += v5
    case 32768:
        var v7 uint32 = 345
        _ = v7
    case -1:
        v2 = v2
    default:
        v1 = v1
    }
    if (false && ((float64((v5 * v5)) > 0.0) || false)) {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint64 = (^(v4))
            var v10 uint = 65535
            _ = fi8
            _ = v10
            _ = v9
        }
        if (65535 <= 32767) {
            var v11 string = ("" + "foo bar")
            _ = v11
        }
    } else {
        if (2 == 108) {
            var v12 uint32 = 1
            fmt.Println(v2, v5, v5)
            for fi13 := 0; fi13 < 3; fi13++ {
                _ = len(v3)
                v3 = append(v3, (v2 >> (((v6 ^ v6)) & 63)))
                switch (("hello" + "") + ("a" + "\n\t")) {
                case "\n\t":
                    var v14 int16 = 0
                    _ = v14
                case "hello":
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v12 -= v12
                        _ = fi15
                    }
                case "":
                    v3 = append(v3, (v2 >> ((1) & 63)))
                default:
                    var v16 uint16 = uint16(v12)
                    _ = v16
                }
                _ = fi13
            }
            _ = v12
        }
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        if (155 < 236) {
            v6 += v6
        } else {
            var v18 float64 = -2.25
            _ = v18
        }
        _ = fi17
    }
    var v19 uint8 = 238
    if false {
        var v20 uint64 = 2
        if (2147483648 != 4294967295) {
            var v21 int8 = -1
            _ = v21
        }
        if true {
            for fi22 := 0; fi22 < 12; fi22++ {
                v23 := GenS1{F0: "", F1: 0}
                _ = fi22
                _ = v23
            }
            switch "\n\t" {
            case "":
                switch "foo bar" {
                case "a":
                    switch 32767 {
                    case 32767:
                        v3 = append(v3, (v2 ^ 0))
                    case 832:
                        if (((!("" != "foo bar")) || ((("hello" > "foo bar") || ((-9223372036854775808 <= 405) && (v2 != (v2 | -297)))) || true)) && false) {
                            var v24 int8 = 127
                            v20 -= (v4 << ((446) & 63))
                            v25 := GenS1{F0: "hello", F1: 1}
                            _ = v24
                            _ = v25
                        }
                    case 0:
                        var v26 int8 = -83
                        _ = v26
                    }
                case "foo bar":
                    for fi27 := 0; fi27 < 11; fi27++ {
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v29 := GenS1{F0: "foo bar", F1: 42}
                            _ = fi28
                            _ = v29
                        }
                        v3 = append(v3, 2147483647)
                        _ = fi27
                    }
                default:
                    if false {
                        var v30 uint = (v6 + (^(v6)))
                        if len(v3) > 0 {
                            v3[0] = -1
                        }
                        _ = v30
                    }
                }
            case "hello":
                if true {
                    _ = len(v3)
                    switch (-(v2)) {
                    case 2147483647:
                        var v31 uint16 = 65535
                        _ = v31
                    }
                    v2 -= (^(v2))
                } else {
                    if (2 == 2) {
                        var v32 uint16 = 0
                        var v33 bool = false
                        _ = v32
                        _ = v33
                    }
                    switch "hello" {
                    case "hello":
                        v20 = 42
                    case "":
                        fmt.Println(v4, v20, v5)
                    case "a":
                        for fi34 := 0; fi34 < 15; fi34++ {
                            _ = len(v3)
                            fmt.Println(v5, v4, fi34)
                            _ = fi34
                        }
                    }
                }
            case "foo bar":
                var v35 uint8 = v19
                _ = v35
            }
            v2 = v2
        }
        _ = v20
    }
    var v36 int = (v2 % (((^(v2))) | 1))
    v4 = (^(v4))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v36, v5)
    v37 := GenS2{F0: 1000, F1: (v2 | v36)}
    v38 := GenS1{F0: "\n\t", F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
