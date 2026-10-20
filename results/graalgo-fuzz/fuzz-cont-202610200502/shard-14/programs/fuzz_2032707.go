package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint64 = 177
    var v2 int8 = 22
    v3 := make(map[uint32]bool)
    if false {
        var v4 int8 = 42
        _ = v4
    } else {
        var v5 uint16 = 65535
        var v6 uint = 907
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(fi7)
        var v8 int16 = -32768
        _ = fi7
        _ = v8
    }
    v9 := []int{20, 127}
    var v10 int16 = 42
    v10 = v10
    v1 -= (v1 % (((v1 + (v1 + 2))) | 1))
    var v11 int8 = v2
    var v12 uint16 = 855
    if false {
        if true {
            switch ("\n\t" + ("foo bar" + "hello")) {
            case "\n\t":
                var v13 bool = (2 > 1)
                _ = v13
            default:
                switch 32767 {
                case 563:
                    if true {
                        for fi14 := 0; fi14 < 12; fi14++ {
                            var v15 float64 = 0.0
                            switch "\n\t" {
                            case "\n\t":
                                v11 -= v11
                            default:
                                if (uint(v12) <= 1) {
                                    for fi16 := 0; fi16 < 11; fi16++ {
                                        if true {
                                            v17 := GenS1{F0: false}
                                            v15 = 1.5
                                            _ = v17
                                        } else {
                                            v2 = v2
                                        }
                                        _ = fi16
                                    }
                                }
                            }
                            _ = fi14
                            _ = v15
                        }
                    }
                case 32768:
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v10 -= 0
                        v1 += v1
                        _ = fi18
                    }
                default:
                    var v19 bool = (true && true)
                    _ = v19
                }
            }
        } else {
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS1{F0: ((false || (true || (true && (1 >= -989)))) || true)}
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: (2 != v12)}
            fmt.Println(v2)
            _ = v22
        }
        if false {
            if false {
                var v23 float32 = 1.5
                _ = v23
            }
            v24 := GenS1{F0: (!(!false))}
            _ = v24
        } else {
            if len(v9) > 0 {
                v9[0] = 2147483647
            }
            v12 = 65535
        }
    } else {
        v1 = v1
        v12 -= (v12 / ((v12) | 1))
    }
    _ = len(v9)
    v25 := v3[1]
    var v26 uint = 179
    var v27 int8 = v11
    var v28 string = (("" + "\n\t") + "foo bar")
    v28 = v28
    var v29 int16 = (^(v10))
    v2 = (v11 >> (((^(v26))) & 63))
    v28 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v9
}
