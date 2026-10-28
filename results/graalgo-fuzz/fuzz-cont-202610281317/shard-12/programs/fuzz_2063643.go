package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 bool }
type GenS2 struct { F0 int8 }
func genFn0(p0 int) bool {
    p0 -= p0
    for fi5 := 0; fi5 < 2; fi5++ {
        p0 = p0
        v6 := []int{(fi5 * (^(p0))), (fi5 >> ((465) & 63)), (p0 ^ fi5)}
        _ = fi5
        _ = v6
    }
    if false {
        var v7 uint64 = 883
        switch p0 {
        case -9223372036854775808:
            p0 += (p0 & p0)
        case 9223372036854775807:
            var v8 uint16 = 720
            _ = v8
        }
        if true {
            p0 = (p0 - (-(p0)))
            p0 -= p0
        } else {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 float64 = (-2.25 / -2.25)
                _ = fi9
                _ = v10
            }
            v11 := []int{p0, p0, p0, p0}
            for fi12 := 0; fi12 < 1; fi12++ {
                switch "foo bar" {
                case "":
                    v13 := GenS1{F0: 125, F1: 1, F2: ("" >= "a")}
                    _ = v13
                default:
                    if (!false) {
                        v7 -= v7
                        var v14 int = fi12
                        _ = v14
                    } else {
                        switch ("hello" + "") {
                        case "foo bar":
                            var v15 string = ("" + ("foo bar" + "hello"))
                            _ = v15
                        case "hello":
                            v16 := GenS1{F0: -1, F1: 0, F2: (false || true)}
                            _ = v16
                        case "a":
                            v17 := GenS1{F0: -610, F1: 32767, F2: (!false)}
                            _ = v17
                        }
                    }
                }
                _ = fi12
            }
            _ = v11
        }
        _ = v7
    } else {
        if (((0 >= 2147483647) || true) || (42 >= 42)) {
            v18 := []int{(p0 << ((255) & 63)), -9223372036854775808, 127}
            var v19 int16 = 99
            if (-164 < (v19 << ((18446744073709551615) & 63))) {
                v20 := GenS1{F0: 0, F1: 255, F2: (49 == 42)}
                p0 += -367
                _ = v20
            } else {
                switch (("a" + "foo bar") + "\n\t") {
                case "a":
                    for fi21 := 0; fi21 < 4; fi21++ {
                        for fi22 := 0; fi22 < 12; fi22++ {
                            var v23 float32 = 1e10
                            _ = fi22
                            _ = v23
                        }
                        var v24 int32 = 551
                        _ = fi21
                        _ = v24
                    }
                case "\n\t":
                    var v25 uint32 = 65535
                    _ = v25
                }
                if (2 > 2147483648) {
                    p0 -= p0
                    v19 -= v19
                    v26 := GenS1{F0: -786, F1: 854, F2: true}
                    _ = v26
                } else {
                    var v27 float64 = (-2.25 + 3.14159)
                    fmt.Println(p0)
                    if true {
                        if false {
                            for fi28 := 0; fi28 < 2; fi28++ {
                                p0 += fi28
                                v27 = -2.25
                                var v29 uint16 = 0
                                _ = fi28
                                _ = v29
                            }
                            var v30 uint32 = 649
                            for fi31 := 0; fi31 < 5; fi31++ {
                                v27 = v27
                                _ = fi31
                            }
                            _ = v30
                        } else {
                            var v32 bool = (false && false)
                            _ = v32
                        }
                        var v33 string = (("\n\t" + (("a" + "") + "a")) + ("foo bar" + "hello"))
                        var v34 int = p0
                        _ = v33
                        _ = v34
                    }
                    _ = v27
                }
            }
            _ = v18
            _ = v19
        }
        v35 := []int{(p0 % ((32768) | 1))}
        for fi36 := 0; fi36 < 3; fi36++ {
            p0 -= -9223372036854775808
            fmt.Println(fi36, fi36)
            _ = fi36
        }
        _ = v35
    }
    var v37 int8 = -84
    _ = p0
    _ = v37
    return true
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 127
    v3 := make(map[string]bool)
    v4 := []string{"hello", ("\n\t" + (("foo bar" + "") + ("a" + "hello")))}
    for fi38 := 0; fi38 < 4; fi38++ {
        delete(v3, "\n\t")
        v2 += v2
        _ = fi38
    }
    v2 = (v2 >> ((v1) & 63))
    var v39 uint16 = 2
    if ((42 != (v2 * (v2 & 162))) || (!(-1 <= 1))) {
        v3["foo bar"] = (v1 < 680)
        var v40 float64 = -2.25
        fmt.Println(v40)
        _ = v40
    }
    v41 := GenS2{F0: -6}
    for fi42 := 0; fi42 < 7; fi42++ {
        fmt.Println(v2)
        _ = fi42
    }
    for fi43 := 0; fi43 < 11; fi43++ {
        var v44 uint = v1
        _ = fi43
        _ = v44
    }
    var v45 uint = (v1 + 931)
    v46 := v3["\n\t"]
    v47 := GenS2{F0: 127}
    var v48 uint = v45
    var v49 float32 = (1.5 / (float32(1e10) / (1.5 / 1.0)))
    v50 := GenS1{F0: -2147483648, F1: (v39 | v39), F2: v46}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
