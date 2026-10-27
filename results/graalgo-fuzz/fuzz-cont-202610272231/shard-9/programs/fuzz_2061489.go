package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 int) bool {
    fmt.Println(p0)
    p0 += (p0 | 32767)
    _ = p0
    return true
}
func genFn1() int8 {
    v6 := []int{-9223372036854775808, 4294967295, 0, 2147483647}
    var v7 uint8 = 42
    v8 := GenS1{F0: (v7 << ((0) & 63)), F1: 65535}
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 bool = false
    var v2 float32 = -2.25
    var v3 int32 = 0
    var v4 int16 = 1000
    v5 := []string{"", "\n\t", ("foo bar" + ("hello" + ("\n\t" + ("foo bar" + "\n\t")))), ""}
    switch ("\n\t" + "\n\t") {
    case "":
        fmt.Println(v4, v3, v1)
    default:
        v5 = append(v5, "")
    }
    if (v1 || false) {
        if len(v5) > 0 {
            v5[0] = ("\n\t" + "\n\t")
        }
        v9 := GenS1{F0: 247, F1: 65535}
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: 42, F1: 42}
                if len(v5) > 0 {
                    v5[0] = ""
                }
                _ = fi11
                _ = v12
            }
            switch ("a" + "\n\t") {
            case "\n\t":
                switch 32768 {
                case -104:
                    switch ("\n\t" + ("a" + ("foo bar" + (("hello" + ("" + ("hello" + "\n\t"))) + "\n\t")))) {
                    case "\n\t":
                        var v13 float32 = v2
                        _ = v13
                    default:
                        var v14 int8 = 42
                        _ = v14
                    }
                case -780:
                    for fi15 := 0; fi15 < 18; fi15++ {
                        v16 := GenS1{F0: 255, F1: 65535}
                        _ = fi15
                        _ = v16
                    }
                case -351:
                    if v1 {
                        var v17 int8 = -128
                        var v18 float64 = 0.1
                        _ = v17
                        _ = v18
                    } else {
                        switch -882 {
                        case 42:
                            if v1 {
                                for fi19 := 0; fi19 < 6; fi19++ {
                                    v1 = (1 < -94)
                                    v5 = append(v5, "\n\t")
                                    _ = fi19
                                }
                                for fi20 := 0; fi20 < 16; fi20++ {
                                    fmt.Println(fi20, v2, fi20)
                                    fmt.Println(fi10, v1)
                                    v21 := GenS1{F0: 135, F1: 115}
                                    _ = fi20
                                    _ = v21
                                }
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                fmt.Println(v4, v3, v2)
                                v22 := GenS1{F0: 127, F1: 1}
                                _ = v22
                            }
                        }
                        v1 = (1 >= -1)
                    }
                }
            case "foo bar":
                var v23 int = 4294967295
                _ = v23
            }
            _ = fi10
        }
        _ = v9
    } else {
        if v1 {
            v3 = v3
            v24 := GenS2{F0: 255, F1: (v3 % ((v3) | 1)), F2: 0}
            v25 := GenS1{F0: 1, F1: 114}
            _ = v24
            _ = v25
        } else {
            fmt.Println(v4)
        }
        var v26 uint32 = 2147483647
        _ = v26
    }
    fmt.Println(v4)
    switch int(v3) {
    case -2147483648:
        _ = len(v5)
    case 32767:
        for fi27 := 0; fi27 < 10; fi27++ {
            v5 = append(v5, (("" + "foo bar") + ("\n\t" + ("a" + "foo bar"))))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi27
        }
    default:
        var v28 int = 127
        _ = v28
    }
    v29 := GenS2{F0: 21, F1: v3, F2: -46}
    var v30 int16 = v4
    v1 = v1
    if v1 {
        if (("\n\t" <= "\n\t") || ((v3 & 42) >= v3)) {
            switch 981 {
            case -519:
                var v31 int64 = 1
                _ = v31
            }
            var v32 bool = (431 == 256)
            v4 = 416
            _ = v32
        }
    } else {
        v1 = false
        if v1 {
            for fi33 := 0; fi33 < 10; fi33++ {
                v1 = false
                fmt.Println(v3, v4)
                _ = fi33
            }
        }
        for fi34 := 0; fi34 < 15; fi34++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (!false) {
                v30 = v4
            }
            _ = fi34
        }
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        v5 = append(v5, ("\n\t" + ("foo bar" + "foo bar")))
        _ = fi35
    }
    switch ("foo bar" + (("\n\t" + ("a" + ("a" + (("a" + "a") + "")))) + ("" + "foo bar"))) {
    case "\n\t":
        var v36 bool = (9223372036854775807 <= 446)
        _ = v36
    default:
        fmt.Println(v2)
    }
    if (!(67 >= 255)) {
        for fi37 := 0; fi37 < 2; fi37++ {
            v4 -= (v30 | (v30 | 0))
            _ = fi37
        }
    } else {
        v1 = (1 != 213)
        for fi38 := 0; fi38 < 12; fi38++ {
            v39 := GenS1{F0: 82, F1: 836}
            if ((-96 != 1) && (42 <= v3)) {
                v40 := GenS2{F0: 255, F1: -496, F2: 9223372036854775807}
                _ = v40
            }
            _ = fi38
            _ = v39
        }
    }
    v5 = append(v5, (("" + "hello") + "a"))
    var v41 bool = false
    v4 -= (v4 << ((4294967295) & 63))
    for fi42 := 0; fi42 < 11; fi42++ {
        v30 = 180
        for fi43 := 0; fi43 < 17; fi43++ {
            var v44 uint = 2
            _ = fi43
            _ = v44
        }
        for fi45 := 0; fi45 < 9; fi45++ {
            v46 := GenS2{F0: 128, F1: (v3 & (v3 ^ (v3 ^ v3))), F2: 0}
            _ = fi45
            _ = v46
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v5
}
