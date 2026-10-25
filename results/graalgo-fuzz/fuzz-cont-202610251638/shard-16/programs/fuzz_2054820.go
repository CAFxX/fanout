package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
type GenS3 struct { F0 uint16; F1 uint64; F2 string }

func main() {
    v1 := []int{42, 1}
    var v2 int32 = 2147483647
    var v3 uint32 = 1
    var v4 bool = true
    var v5 int64 = -788
    var v6 int8 = 97
    switch 833 {
    case 509:
        if (1 > 896) {
            for fi7 := 0; fi7 < 12; fi7++ {
                fmt.Println(v2, v3)
                v8 := GenS1{F0: fi7, F1: (v2 + -1)}
                if (135 >= 0) {
                    var v9 uint32 = (v3 - v3)
                    var v10 float32 = float32(-2.25)
                    if v4 {
                        v5 += (v5 + -9223372036854775808)
                        for fi11 := 0; fi11 < 9; fi11++ {
                            var v12 uint32 = v3
                            _ = fi11
                            _ = v12
                        }
                        v2 = 1
                    } else {
                        v13 := GenS2{F0: 1, F1: v4, F2: ("foo bar" + "")}
                        for fi14 := 0; fi14 < 17; fi14++ {
                            v1 = append(v1, fi7)
                            var v15 int = (-(fi7))
                            _ = fi14
                            _ = v15
                        }
                        _ = v13
                    }
                    _ = v10
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
            fmt.Println(v6)
            if false {
                if v4 {
                    v16 := GenS3{F0: 13, F1: 542, F2: ""}
                    for fi17 := 0; fi17 < 12; fi17++ {
                        for fi18 := 0; fi18 < 11; fi18++ {
                            if v4 {
                                v19 := GenS1{F0: 0, F1: v2}
                                fmt.Println(v5, v5)
                                _ = v19
                            }
                            if len(v1) > 0 {
                                v1[0] = fi18
                            }
                            _ = fi18
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v5, v2)
                        _ = fi17
                    }
                    fmt.Println(v5, v2)
                    _ = v16
                }
            } else {
                if v4 {
                    var v20 float32 = 1.5
                    v4 = v4
                    for fi21 := 0; fi21 < 2; fi21++ {
                        fmt.Println(v3, fi21, fi21)
                        var v22 uint64 = 2
                        fmt.Println(v3, fi21, v4)
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                }
            }
        }
    case 128:
        var v23 bool = v4
        _ = v23
    default:
        if v4 {
            v24 := GenS1{F0: 42, F1: v2}
            v5 = (v5 % ((v5) | 1))
            v1 = append(v1, 42)
            _ = v24
        }
    }
    var v25 bool = ((!((v4 && v4) && v4)) || (!v4))
    for fi26 := 0; fi26 < 9; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            if v25 {
                v3 += 126
                v1 = append(v1, 32767)
            }
            var v28 bool = (!((!true) || (!((0 > 180) && (v3 != (v3 % ((v3) | 1)))))))
            _ = fi27
            _ = v28
        }
        var v29 int16 = 620
        _ = fi26
        _ = v29
    }
    if (!(32767 <= 2)) {
        v30 := GenS1{F0: 127, F1: (v2 + v2)}
        var v31 uint16 = 2
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 int16 = 1000
        v3 -= 278
        _ = fi32
        _ = v33
    }
    switch ((("" + "foo bar") + "hello") + "hello") {
    case "a":
        v6 += (v6 + 0)
    }
    v5 = v5
    var v34 uint16 = 575
    var v35 uint8 = 42
    for fi36 := 0; fi36 < 13; fi36++ {
        var v37 uint8 = 8
        var v38 uint32 = (v3 >> ((255) & 63))
        var v39 int16 = 0
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    v6 += v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
