package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int16 = 42
    var v2 float32 = -2.25
    var v3 uint64 = 0
    v4 := []int{}
    var v5 uint64 = v3
    v2 = v2
    v3 -= 763
    switch 127 {
    case -662:
        var v6 uint64 = 42
        _ = v6
    default:
        fmt.Println(v1, v3, v5)
    }
    v3 += (v5 + 2)
    for fi7 := 0; fi7 < 8; fi7++ {
        switch -9223372036854775808 {
        case 1:
            v8 := GenS1{F0: 114, F1: (fi7 << ((155) & 63)), F2: 255}
            _ = v8
        case 32768:
            v1 = (v1 << ((0) & 63))
        }
        _ = fi7
    }
    if false {
        var v9 uint16 = 441
        var v10 uint8 = 1
        v4 = append(v4, 1)
        _ = v10
        _ = v9
    } else {
        v5 += v3
        var v11 uint = 1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v11
    }
    var v12 int16 = (v1 % ((v1) | 1))
    if (1 <= 2) {
        v5 = (^(v3))
        fmt.Println(v2, v12, v2)
        v13 := GenS1{F0: 100, F1: 584, F2: 128}
        _ = v13
    }
    v1 = -474
    v12 -= (v1 | v1)
    if false {
        if (!true) {
            fmt.Println(v3)
            if (v3 == (v5 - (v3 ^ v5))) {
                var v14 uint16 = 32767
                v12 -= v1
                _ = v14
            }
            var v15 uint8 = 18
            _ = v15
        }
        var v16 int = -1
        _ = v16
    } else {
        switch "foo bar" {
        case "foo bar":
            switch -822 {
            case 4294967295:
                v17 := GenS1{F0: -128, F1: 128, F2: 42}
                _ = v17
            }
        default:
            fmt.Println(v3, v5, v1)
        }
        if len(v4) > 0 {
            v4[0] = 32767
        }
        if ((!(2147483647 != 0)) || (v12 != -457)) {
            for fi18 := 0; fi18 < 6; fi18++ {
                if (!(!((((v2 - v2) * (float32(0.0) * 0.1)) + (((((v2 / v2) / v2) * (v2 / (3.14159 + v2))) + 1.5) + 0.0)) != v2))) {
                    switch "\n\t" {
                    case "":
                        var v19 float32 = 1.5
                        _ = v19
                    case "hello":
                        if ((2147483648 >= 2147483647) && false) {
                            v20 := GenS2{F0: -90, F1: (("hello" + (("\n\t" + ("" + ("" + "a"))) + (("hello" + ("\n\t" + "hello")) + "\n\t"))) + ("hello" + "hello"))}
                            fmt.Println(v2, v12)
                            var v21 int8 = 0
                            _ = v20
                            _ = v21
                        }
                    case "\n\t":
                        var v22 bool = true
                        _ = v22
                    default:
                        for fi23 := 0; fi23 < 4; fi23++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v3 = (v3 >> ((874) & 63))
                            _ = fi23
                        }
                    }
                    v24 := GenS3{F0: -1}
                    if ((fi18 << ((18446744073709551615) & 63)) < fi18) {
                        switch "a" {
                        case "":
                            fmt.Println(v5)
                        case "\n\t":
                            for fi25 := 0; fi25 < 7; fi25++ {
                                v5 = v5
                                var v26 int32 = -520
                                _ = fi25
                                _ = v26
                            }
                        }
                        if true {
                            v5 -= 980
                            fmt.Println(fi18, v5)
                            for fi27 := 0; fi27 < 3; fi27++ {
                                v28 := GenS1{F0: 1, F1: fi27, F2: 1}
                                v3 = (v3 << ((4294967295) & 63))
                                _ = fi27
                                _ = v28
                            }
                        }
                        fmt.Println(fi18)
                    } else {
                        v2 = ((1.5 - (v2 / (((v2 / v2) * 0.0) - v2))) - v2)
                        v12 = (v12 * (v12 - 42))
                    }
                    _ = v24
                } else {
                    v29 := GenS1{F0: -128, F1: (^(fi18)), F2: 0}
                    _ = v29
                }
                fmt.Println(v3, v1)
                v30 := GenS4{F0: 2147483647}
                _ = fi18
                _ = v30
            }
            if true {
                fmt.Println(v1)
            }
            for fi31 := 0; fi31 < 2; fi31++ {
                v4 = append(v4, (fi31 ^ (fi31 - 42)))
                var v32 int64 = -1
                _ = fi31
                _ = v32
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
