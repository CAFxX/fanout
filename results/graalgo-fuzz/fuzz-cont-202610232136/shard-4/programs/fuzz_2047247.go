package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 bool }

func main() {
    var v1 bool = true
    var v2 uint16 = 130
    v3 := []string{"", ("\n\t" + "a"), "\n\t"}
    v4 := []uint32{0}
    var v5 int8 = 37
    var v6 float64 = 3.14159
    fmt.Println(v2)
    var v7 int64 = -297
    v6 = -2.25
    v8 := GenS1{F0: 2147483647, F1: 255}
    v9 := GenS1{F0: -97, F1: 0}
    if v1 {
        fmt.Println(v7, v5, v2)
        fmt.Println(v5, v2)
        switch 32767 {
        case -1:
            if false {
                if false {
                    switch (((("hello" + "\n\t") + "foo bar") + (("a" + ("hello" + "hello")) + "hello")) + ("a" + "")) {
                    case "a":
                        switch -1 {
                        case 2147483647:
                            var v10 int16 = -1
                            _ = v10
                        default:
                            for fi11 := 0; fi11 < 17; fi11++ {
                                v1 = v1
                                _ = fi11
                            }
                        }
                    case "":
                        var v12 float64 = 1e10
                        _ = v12
                    default:
                        v5 = (v5 << ((255) & 63))
                    }
                    if len(v3) > 0 {
                        v3[0] = ("" + "foo bar")
                    }
                } else {
                    if v1 {
                        fmt.Println(v6, v1)
                        v13 := GenS1{F0: 854, F1: 2}
                        _ = v13
                    } else {
                        if v1 {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v6 += ((v6 / 1e10) * float64((1e10 - ((3.14159 - (((((float32(((0.1 / 1e10) - ((((v6 - v6) - float64(1.5)) / (v6 - v6)) * (v6 * v6)))) * 1e10) * -2.25) * 1e10) - 3.14159) / float32((v6 * v6)))) / 0.1))))
                    }
                }
                v2 = v2
                v1 = v1
            }
        case 1:
            v14 := GenS2{F0: (v5 >> ((2) & 63))}
            _ = v14
        default:
            fmt.Println(v7, v7)
        }
    }
    v15 := GenS1{F0: 0, F1: 992}
    v16 := GenS2{F0: v5}
    v4 = append(v4, 15)
    var v17 uint64 = 0
    if v1 {
        fmt.Println(v6, v7)
    } else {
        if len(v3) > 0 {
            v3[0] = ""
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            fmt.Println(v17, v1, fi18)
            v19 := GenS1{F0: fi18, F1: 18446744073709551615}
            if v1 {
                switch (fi18 << ((7) & 63)) {
                case 631:
                    var v20 int8 = v5
                    _ = v20
                case -703:
                    var v21 uint = 255
                    _ = v21
                case -526:
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v23 := GenS3{F0: (fi22 & (^(fi22))), F1: v1}
                        _ = fi22
                        _ = v23
                    }
                }
            }
            _ = fi18
            _ = v19
        }
    }
    if (!(-2147483648 > -3)) {
        v1 = (v5 > v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
