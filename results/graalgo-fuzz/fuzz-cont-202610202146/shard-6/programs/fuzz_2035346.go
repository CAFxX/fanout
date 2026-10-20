package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int16; F1 int }

func main() {
    var v1 int16 = 596
    var v2 bool = true
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 int = 4294967295
    var v6 uint64 = 364
    if v2 {
        v2 = (!true)
        switch "" {
        case "\n\t":
            v5 -= v5
        case "":
            switch "\n\t" {
            case "a":
                v7 := GenS1{F0: v1, F1: (v2 || (!(true || true)))}
                _ = v7
            case "hello":
                v6 += (v6 ^ v6)
            default:
                for fi8 := 0; fi8 < 7; fi8++ {
                    if true {
                        v9 := []int{9223372036854775807}
                        _ = v9
                    }
                    var v10 int = 32767
                    _ = fi8
                    _ = v10
                }
            }
        }
        if (!(255 < 708)) {
            switch -1 {
            case 4294967295:
                var v11 int = v5
                _ = v11
            }
            v12 := GenS1{F0: v1, F1: v2}
            v13 := GenS2{F0: (^(v1)), F1: -2147483648}
            _ = v12
            _ = v13
        } else {
            fmt.Println(v5, v2, v5)
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 uint32 = 0
        var v16 float64 = 0.0
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint16 = 0
    fmt.Println(v6)
    if true {
        switch "hello" {
        case "foo bar":
            if false {
                if v2 {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        if ((1 >= 966) || (320 > 2147483647)) {
                            v2 = (1000 != 1000)
                        }
                        if v2 {
                            switch (^(fi18)) {
                            case -2147483648:
                                var v19 string = (("foo bar" + "hello") + (("hello" + (("hello" + "") + "a")) + "a"))
                                _ = v19
                            case 4294967295:
                                if true {
                                    for fi20 := 0; fi20 < 16; fi20++ {
                                        if true {
                                            v21 := GenS2{F0: v1, F1: fi18}
                                            v2 = (365 < 0)
                                            v22 := GenS1{F0: (v1 + 1), F1: (v6 != (v6 - 18446744073709551615))}
                                            v5 = v5
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v23 int64 = 0
                                                _ = v23
                                            case "":
                                                v24 := []int{fi18, (fi18 | v5)}
                                                _ = v24
                                            case "\n\t":
                                                v25 := GenS2{F0: 0, F1: -2147483648}
                                                _ = v25
                                            default:
                                                v17 -= 2
                                            }
                                            v26 := []int{fi18, (v5 - -2147483648)}
                                            _ = v21
                                            _ = v22
                                            _ = v26
                                        }
                                        _ = fi20
                                    }
                                } else {
                                    v1 += v1
                                }
                            default:
                                if (0 < 2147483647) {
                                    var v27 int16 = (v1 >> ((565) & 63))
                                    v28 := GenS2{F0: (v1 | v27), F1: 4294967295}
                                    _ = v27
                                    _ = v28
                                } else {
                                    v2 = (!(!(!v2)))
                                    if false {
                                        v5 = (^(fi18))
                                    } else {
                                        var v29 float64 = 1e10
                                        fmt.Println(v17)
                                        _ = v29
                                    }
                                }
                            }
                            if ((!((1 <= 2147483647) && false)) || (!(42 <= 2147483648))) {
                                var v30 uint16 = (v17 >> ((618) & 63))
                                _ = v30
                            }
                            v31 := GenS1{F0: v1, F1: (!(((v5 % (((fi18 / ((v5) | 1))) | 1)) <= v5) && false))}
                            _ = v31
                        }
                        _ = fi18
                    }
                    var v32 float32 = (float32((0.1 + -2.25)) * (1.5 - ((1e10 + 0.0) * 1e10)))
                    _ = v32
                }
                var v33 bool = (42 <= 0)
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 3; fi34++ {
            for fi35 := 0; fi35 < 8; fi35++ {
                v6 += v6
                _ = fi35
            }
            _ = fi34
        }
    } else {
        var v36 uint8 = 0
        var v37 float64 = 1.5
        _ = v36
        _ = v37
    }
    if (136 <= (v6 >> ((255) & 63))) {
        var v38 int8 = -1
        _ = v38
    }
    v39 := GenS2{F0: 0, F1: v5}
    switch 496 {
    case 32768:
        v5 = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
