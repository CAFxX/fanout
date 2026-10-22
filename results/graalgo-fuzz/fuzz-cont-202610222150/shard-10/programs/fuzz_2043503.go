package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 string }
func genFn0(p0 uint8, p1 uint16) int {
    var v4 uint32 = 344
    p0 = p0
    v5 := GenS1{F0: true}
    var v6 bool = false
    v4 = 638
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 734
}
func genFn1(p0 bool) uint {
    p0 = (888 < 1)
    v7 := []int{-436, 127}
    _ = p0
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    var v3 uint = 65535
    v8 := []int{-2147483648, -9223372036854775808}
    v3 += 237
    if (((0.0 / 1e10) != (1e10 - -2.25)) && ((v2 >= v2) || true)) {
        if false {
            var v9 int16 = -32768
            v10 := GenS1{F0: true}
            _ = v10
            _ = v9
        } else {
            v2 = (v2 / (((v2 % ((1) | 1))) | 1))
        }
    }
    fmt.Println(v2, v3)
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int16 = 42
        _ = fi11
        _ = v12
    }
    if false {
        if (3.14159 <= -2.25) {
            var v13 bool = ("" != "")
            switch "" {
            case "foo bar":
                v2 += v2
            case "":
                for fi14 := 0; fi14 < 3; fi14++ {
                    _ = len(v8)
                    _ = fi14
                }
            case "a":
                v13 = true
            default:
                if v13 {
                    v2 = v2
                    switch -9223372036854775808 {
                    case 32768:
                        var v15 int = 607
                        _ = v15
                    case -2147483648:
                        fmt.Println(v2, v13, v3)
                    }
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v3 += (^(v3))
                        for fi17 := 0; fi17 < 8; fi17++ {
                            switch ("hello" + "\n\t") {
                            case "a":
                                if (v13 || true) {
                                    var v18 int32 = -2147483648
                                    v2 += v2
                                    _ = v18
                                }
                            case "hello":
                                var v19 int = 2147483647
                                _ = v19
                            default:
                                fmt.Println(v2)
                            }
                            _ = fi17
                        }
                        _ = fi16
                    }
                } else {
                    if v13 {
                        var v20 uint8 = 42
                        if v13 {
                            var v21 int = 880
                            v20 = (v20 >> (((v3 >> ((921) & 63))) & 63))
                            var v22 int16 = 1000
                            _ = v21
                            _ = v22
                        }
                        _ = v20
                    }
                }
            }
            _ = v13
        }
        _ = len(v8)
        fmt.Println(v3)
    } else {
        var v23 uint32 = 42
        for fi24 := 0; fi24 < 3; fi24++ {
            v3 = v3
            switch genFn0(15, 0) {
            case 497:
                switch ((("\n\t" + ("a" + "foo bar")) + "foo bar") + "\n\t") {
                case "a":
                    v23 = (v23 << ((v3) & 63))
                }
            case -1:
                v23 -= uint32(v3)
            case 127:
                v3 -= (^(v3))
            default:
                v23 += (v23 - (v23 >> ((v3) & 63)))
            }
            _ = fi24
        }
        _ = v23
    }
    v3 += (v3 + v3)
    for fi25 := 0; fi25 < 7; fi25++ {
        v3 = 65535
        if false {
            v2 += 0
            var v26 int = fi25
            _ = v26
        }
        switch "\n\t" {
        case "a":
            if true {
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 uint64 = 1
                    var v29 uint32 = 0
                    switch ("foo bar" + "") {
                    case "hello":
                        if false {
                            var v30 uint8 = uint8(v29)
                            v31 := GenS2{F0: uint16(v2), F1: "a"}
                            _ = v30
                            _ = v31
                        } else {
                            fmt.Println(fi27)
                            v2 -= (v28 * v2)
                            v8 = append(v8, fi25)
                        }
                    default:
                        if (!true) {
                            v29 = v29
                            var v32 bool = (true && (0 <= 1000))
                            _ = v32
                        }
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                var v33 uint = v3
                _ = v33
            }
        }
        _ = fi25
    }
    v3 = (v3 & v3)
    v2 = (v2 << ((genFn1(((!(true && true)) || false))) & 63))
    for fi34 := 0; fi34 < 9; fi34++ {
        switch ("hello" + "foo bar") {
        case "a":
            if (1 >= 1) {
                for fi35 := 0; fi35 < 19; fi35++ {
                    v1[("foo bar" + ("" + "a"))] = (-(fi35))
                    var v36 int32 = -1
                    _ = fi35
                    _ = v36
                }
            }
        default:
            for fi37 := 0; fi37 < 6; fi37++ {
                fmt.Println(fi37)
                _ = fi37
            }
        }
        v2 -= (v2 >> ((genFn1(true)) & 63))
        for fi38 := 0; fi38 < 8; fi38++ {
            v2 = 205
            _ = len(v1)
            _ = fi38
        }
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
