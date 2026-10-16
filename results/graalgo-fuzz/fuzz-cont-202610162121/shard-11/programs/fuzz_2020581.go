package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint16 }
func genFn0(p0 int) uint8 {
    var v5 uint16 = 65535
    switch ("foo bar" + "a") {
    case "a":
        switch "foo bar" {
        case "\n\t":
            for fi6 := 0; fi6 < 10; fi6++ {
                if false {
                    if true {
                        fmt.Println(fi6, p0)
                    }
                    for fi7 := 0; fi7 < 2; fi7++ {
                        v5 = (v5 | (v5 & (^(v5))))
                        v8 := GenS1{F0: 65535, F1: 42}
                        switch "hello" {
                        case "hello":
                            if true {
                                var v9 uint32 = 0
                                fmt.Println(v9)
                                _ = v9
                            } else {
                                var v10 float64 = 1.5
                                _ = v10
                            }
                        case "\n\t":
                            v5 -= (v5 + v5)
                        case "foo bar":
                            p0 -= fi6
                        default:
                            p0 -= p0
                            v5 += (v5 & v5)
                            fmt.Println(fi6, fi7, fi7)
                            var v11 uint32 = 42
                            if (!false) {
                                var v12 string = "foo bar"
                                _ = v12
                            }
                            _ = v11
                        }
                        var v13 int16 = 32767
                        _ = fi7
                        _ = v13
                        _ = v8
                    }
                    for fi14 := 0; fi14 < 10; fi14++ {
                        var v15 int16 = 1
                        _ = fi14
                        _ = v15
                    }
                } else {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v5 = (v5 / ((v5) | 1))
                        _ = fi16
                    }
                }
                _ = fi6
            }
        default:
            if false {
                var v17 int64 = 42
                v18 := []int{(p0 * 4294967295), (p0 % (((p0 + (p0 * (p0 | p0)))) | 1)), -25, (p0 >> ((1) & 63))}
                if ((67 >= 498) || (-41 == 100)) {
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    v5 += (v5 / ((0) | 1))
                    var v19 int8 = -128
                    _ = v19
                } else {
                    var v20 int16 = 572
                    _ = v20
                }
                _ = v17
                _ = v18
            }
        }
    case "":
        v5 += 83
    }
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 uint64 = 2
    var v2 uint64 = v1
    var v3 int64 = 1
    var v4 int8 = 100
    v21 := []int{-2147483648}
    if ("a" < "") {
        var v22 uint8 = 42
        _ = v22
    } else {
        v3 = (v3 ^ (v3 << ((0) & 63)))
    }
    switch ("\n\t" + (("hello" + "hello") + "\n\t")) {
    case "foo bar":
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 int64 = v3
            v25 := GenS1{F0: 0, F1: 9223372036854775807}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        if true {
            for fi27 := 0; fi27 < 15; fi27++ {
                v1 += (^(v2))
                fmt.Println(fi27, fi26, fi26)
                _ = fi27
            }
        } else {
            var v28 int16 = 32767
            _ = v28
        }
        switch (("a" + ("" + "hello")) + "hello") {
        case "foo bar":
            switch int(v1) {
            case 2147483647:
                v29 := GenS1{F0: 4294967295, F1: v3}
                _ = v29
            case -566:
                if true {
                    for fi30 := 0; fi30 < 17; fi30++ {
                        var v31 int8 = v4
                        _ = fi30
                        _ = v31
                    }
                    var v32 string = ("" + ("a" + "\n\t"))
                    _ = v32
                } else {
                    v33 := GenS1{F0: 18446744073709551615, F1: (v3 + (v3 - (v3 << ((618) & 63))))}
                    _ = v33
                }
            case 127:
                var v34 int32 = 0
                _ = v34
            }
        default:
            v4 = v4
        }
        var v35 int32 = 1
        _ = fi26
        _ = v35
    }
    v2 = v1
    v36 := GenS2{F0: -1, F1: 483, F2: 32767}
    var v37 int16 = -1
    if (!true) {
        v1 = v1
        switch -1 {
        case -9223372036854775808:
            v3 -= int64(v37)
        case 478:
            switch "a" {
            case "foo bar":
                v1 = v2
            case "a":
                fmt.Println(v1, v1, v3)
            case "":
                v38 := GenS1{F0: 65535, F1: 0}
                _ = v38
            default:
                if (!(688 > -209)) {
                    v37 += (v37 ^ 630)
                    v3 = v3
                }
            }
        default:
            v21 = append(v21, 886)
        }
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
    }
    fmt.Println(v37, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
