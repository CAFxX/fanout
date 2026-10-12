package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 bool }
type GenS2 struct { F0 uint; F1 bool; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }

func main() {
    v1 := []uint32{0}
    v2 := []string{"a", "a", "foo bar"}
    v3 := make(map[int]string)
    var v4 int = 4294967295
    if (0.1 < 1.5) {
        v5 := GenS1{F0: 804, F1: -116, F2: (!true)}
        v4 = (v4 * (v4 * (v4 | 28)))
        switch "foo bar" {
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            for fi6 := 0; fi6 < 5; fi6++ {
                fmt.Println(fi6, v4)
                if false {
                    fmt.Println(fi6)
                    switch (v4 / ((v4) | 1)) {
                    case 808:
                        for fi7 := 0; fi7 < 16; fi7++ {
                            var v8 string = "a"
                            v9 := GenS1{F0: 2, F1: -323, F2: (!(false && true))}
                            v4 = v4
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                            _ = fi7
                            _ = v8
                            _ = v9
                        }
                    case 844:
                        for fi10 := 0; fi10 < 7; fi10++ {
                            v11 := GenS2{F0: 18446744073709551615, F1: false, F2: 1}
                            _ = fi10
                            _ = v11
                        }
                    default:
                        v4 = -2147483648
                    }
                } else {
                    if (1 > 392) {
                        v12 := GenS3{F0: 0, F1: 1}
                        v4 += (fi6 / (((v4 << ((20) & 63))) | 1))
                        _ = v12
                    }
                    v13 := GenS4{F0: true}
                    fmt.Println(v4)
                    _ = v13
                }
                _ = fi6
            }
        }
        _ = v5
    } else {
        v4 = (v4 >> ((18446744073709551615) & 63))
        fmt.Println(v4)
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        fmt.Println(fi14)
        var v15 string = ("a" + "a")
        _ = fi14
        _ = v15
    }
    v2 = append(v2, "a")
    if (647 != 2) {
        v1 = append(v1, 229)
        v3[993] = "\n\t"
    }
    v16 := GenS3{F0: 42, F1: 255}
    v17 := GenS5{F0: "hello"}
    for fi18 := 0; fi18 < 14; fi18++ {
        fmt.Println(v4, fi18)
        _ = fi18
    }
    v19 := GenS6{F0: -955}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
