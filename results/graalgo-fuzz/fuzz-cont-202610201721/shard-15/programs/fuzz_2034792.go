package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int32 }

func main() {
    var v1 bool = false
    var v2 uint64 = 394
    v3 := []int{32768, 127, 115, 0}
    v2 += v2
    v2 -= v2
    _ = len(v3)
    for fi4 := 0; fi4 < 3; fi4++ {
        fmt.Println(v1)
        _ = fi4
    }
    if v1 {
        v1 = v1
        var v5 int32 = -827
        for fi6 := 0; fi6 < 13; fi6++ {
            v2 -= 4294967295
            for fi7 := 0; fi7 < 10; fi7++ {
                fmt.Println(fi6, v5)
                _ = fi7
            }
            _ = fi6
        }
        _ = v5
    }
    if (32767 > 255) {
        for fi8 := 0; fi8 < 12; fi8++ {
            if true {
                v9 := GenS1{F0: int(v2), F1: 590, F2: int32(fi8)}
                v1 = (v2 < v2)
                _ = v9
            } else {
                if (!(!(42 < 1))) {
                    switch ("foo bar" + "foo bar") {
                    case "hello":
                        var v10 int16 = 999
                        _ = v10
                    default:
                        var v11 uint = 211
                        _ = v11
                    }
                } else {
                    var v12 uint8 = 152
                    for fi13 := 0; fi13 < 13; fi13++ {
                        var v14 uint = 255
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                }
                v3 = append(v3, fi8)
                _ = len(v3)
            }
            _ = fi8
        }
    }
    if false {
        v15 := GenS1{F0: 0, F1: 4294967295, F2: -2147483648}
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 int32 = -2147483648
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS1{F0: 289, F1: 2, F2: (v17 + (-(v17)))}
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    var v20 float32 = float32(3.14159)
    if (false && ((v1 || true) || (583 <= 303))) {
        for fi21 := 0; fi21 < 14; fi21++ {
            v3 = append(v3, fi21)
            _ = fi21
        }
        v22 := GenS1{F0: 21, F1: 0, F2: -1}
        _ = v22
    } else {
        switch "" {
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            if v1 {
                var v23 bool = (!(160 >= 0))
                v23 = (true && true)
                _ = v23
            } else {
                if (false || false) {
                    fmt.Println(v2, v1)
                    v2 = v2
                    v20 -= v20
                }
            }
        case "hello":
            v20 -= (-2.25 / 1.5)
        default:
            if (255 == 2) {
                switch "foo bar" {
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if v1 {
                    var v24 int = -2147483648
                    v20 = v20
                    _ = v24
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
