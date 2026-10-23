package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 string; F1 int64; F2 int64 }

func main() {
    var v1 uint16 = 816
    var v2 uint = 0
    v3 := []string{"hello"}
    var v4 string = ("foo bar" + ("hello" + ("" + ("a" + "a"))))
    for fi5 := 0; fi5 < 16; fi5++ {
        if false {
            _ = len(v3)
            if true {
                v3 = append(v3, "")
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = v4
                }
            }
            v2 += (^(v2))
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = len(v3)
            switch v4 {
            case "":
                var v6 string = v4
                _ = v6
            case "foo bar":
                if false {
                    if true {
                        var v7 bool = (-9223372036854775808 >= 0)
                        _ = v7
                    }
                }
            default:
                v1 = 2
            }
        }
        _ = fi5
    }
    var v8 float64 = 0.0
    v1 -= (v1 ^ (v1 << ((v2) & 63)))
    var v9 int64 = -977
    v2 += (v2 % (((v2 >> (((v2 / ((v2) | 1))) & 63))) | 1))
    if (false || ((v1 % (((^(v1))) | 1)) >= v1)) {
        if len(v3) > 0 {
            v3[0] = ("a" + (v4 + "a"))
        }
    } else {
        v2 += (v2 ^ v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = v1
    }
    var v10 uint = 1
    var v11 float32 = (((-2.25 + 0.0) * (1.5 / ((1e10 * ((1.5 + (0.1 / 1.5)) + 0.0)) * -2.25))) * 1.5)
    v1 -= (v1 << ((255) & 63))
    fmt.Println(v9, v1)
    v12 := GenS1{F0: 2, F1: 122, F2: v1}
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 bool = false
        var v15 int64 = (v9 >> ((v2) & 63))
        var v16 uint16 = uint16(v10)
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint8 = 127
    switch 0 {
    case 0:
        for fi18 := 0; fi18 < 6; fi18++ {
            if true {
                for fi19 := 0; fi19 < 7; fi19++ {
                    v11 -= float32((3.14159 + (v8 + v8)))
                    _ = fi19
                }
                switch ((("a" + v4) + (v4 + v4)) + v4) {
                case "":
                    v20 := GenS2{F0: ("hello" + v4), F1: int64(v1), F2: (^(v9))}
                    _ = v20
                default:
                    v21 := GenS1{F0: 375, F1: 100, F2: (v1 ^ (v1 | 140))}
                    _ = v21
                }
                var v22 string = v4
                _ = v22
            }
            _ = fi18
        }
    }
    if false {
        v23 := GenS1{F0: 1, F1: 127, F2: v1}
        v24 := GenS1{F0: uint64(v17), F1: 68, F2: uint16(v17)}
        v9 -= -9223372036854775808
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 12; fi25++ {
            v10 = v2
            fmt.Println(v10, v9, v9)
            fmt.Println(v1, v9, v17)
            _ = fi25
        }
        if (!false) {
            var v26 float64 = ((v8 - (v8 * (((v8 - v8) - (1.5 / v8)) + 0.1))) + float64((1e10 - v11)))
            switch 127 {
            case 32767:
                var v27 bool = (((float32(v26) - (0.0 + ((float32((v8 + 0.0)) - 0.1) + (v11 + (v11 + float32(v26)))))) * v11) != (v11 / 1.0))
                _ = v27
            case -454:
                for fi28 := 0; fi28 < 8; fi28++ {
                    for fi29 := 0; fi29 < 14; fi29++ {
                        var v30 int64 = (v9 | 515)
                        v2 += v2
                        fmt.Println(v1, fi28)
                        _ = fi29
                        _ = v30
                    }
                    _ = fi28
                }
            }
            fmt.Println(v10, v10, v8)
            _ = v26
        } else {
            switch -1 {
            case 731:
                var v31 string = v4
                _ = v31
            case -80:
                for fi32 := 0; fi32 < 8; fi32++ {
                    v9 += v9
                    v10 = 2
                    _ = fi32
                }
            }
            v9 += -1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
