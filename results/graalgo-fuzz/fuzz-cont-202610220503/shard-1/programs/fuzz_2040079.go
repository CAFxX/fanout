package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 string }

func main() {
    var v1 uint8 = 131
    var v2 uint = 680
    var v3 uint16 = 2
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := []int{(fi4 << ((v2) & 63)), 32767}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 108, F1: -1, F2: "\n\t"}
        var v8 int8 = -128
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v3, v3)
    if false {
        if false {
            fmt.Println(v1, v2)
        } else {
            fmt.Println(v3, v3)
        }
        for fi9 := 0; fi9 < 7; fi9++ {
            if false {
                var v10 float64 = 1.5
                var v11 int16 = -1
                v12 := []int{}
                _ = v10
                _ = v11
                _ = v12
            } else {
                v2 -= v2
            }
            v13 := []int{42, (fi9 & fi9), 0}
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS1{F0: 2, F1: -1, F2: "\n\t"}
                v3 -= (v3 >> (((v2 << (((v2 / ((v2) | 1))) & 63))) & 63))
                _ = fi14
                _ = v15
            }
            _ = fi9
            _ = v13
        }
    }
    if false {
        switch "" {
        case "foo bar":
            v16 := []int{-964, 32767, 128}
            _ = v16
        default:
            fmt.Println(v3)
        }
        switch 87 {
        case 1:
            if (true || true) {
                v3 += v3
            } else {
                if false {
                    if (255 > v3) {
                        if (false || false) {
                            fmt.Println(v3)
                            fmt.Println(v1, v2, v1)
                            v3 = 255
                        }
                        v17 := GenS1{F0: uint8(v2), F1: 9223372036854775807, F2: "foo bar"}
                        _ = v17
                    } else {
                        switch "foo bar" {
                        case "a":
                            var v18 int64 = 0
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 6; fi19++ {
                            for fi20 := 0; fi20 < 18; fi20++ {
                                v21 := []int{-1}
                                switch "" {
                                case "a":
                                    var v22 uint = uint(v3)
                                    _ = v22
                                case "foo bar":
                                    v21 = append(v21, 42)
                                }
                                _ = fi20
                                _ = v21
                            }
                            _ = fi19
                        }
                        var v23 int32 = 1
                        _ = v23
                    }
                }
                var v24 uint = (^(v2))
                _ = v24
            }
        case -1:
            v3 -= v3
        case 127:
            v1 = v1
        }
    }
    fmt.Println(v2, v2, v3)
    for fi25 := 0; fi25 < 3; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            v3 += 208
            v27 := []int{-976, fi26, fi25}
            var v28 bool = true
            _ = fi26
            _ = v27
            _ = v28
        }
        v1 = v1
        _ = fi25
    }
    var v29 string = (("foo bar" + ("foo bar" + "")) + "\n\t")
    var v30 uint = v2
    if false {
        var v31 int16 = 1
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
