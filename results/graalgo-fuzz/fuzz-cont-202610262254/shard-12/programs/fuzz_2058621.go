package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint32 }
func genFn0(p0 bool, p1 int, p2 uint32) uint32 {
    p1 = p1
    v4 := []int{(p1 >> ((42) & 63)), p1, p1, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}

func main() {
    v1 := []string{}
    var v2 int32 = 0
    var v3 int = 0
    v3 = (v3 % ((397) | 1))
    v2 += -2147483648
    v2 += v2
    var v5 float64 = 1.5
    var v6 uint = 255
    v6 -= v6
    v2 = -111
    switch (v3 ^ (v3 / ((v3) | 1))) {
    case -991:
        v7 := GenS1{F0: 39, F1: 255, F2: (!(genFn0(true, v3, 42) == 645))}
        _ = v7
    case -9223372036854775808:
        if true {
            fmt.Println(v2, v3)
            if true {
                if true {
                    switch (v3 ^ 0) {
                    case 32767:
                        var v8 uint16 = 65535
                        _ = v8
                    case 0:
                        v2 += v2
                    }
                } else {
                    v9 := GenS1{F0: 67, F1: 0, F2: ((-1 > 42) && ((2 >= 1) || false))}
                    switch "hello" {
                    case "foo bar":
                        switch "hello" {
                        case "":
                            fmt.Println(v3, v3, v6)
                        case "a":
                            for fi10 := 0; fi10 < 11; fi10++ {
                                fmt.Println(v6, v5)
                                _ = fi10
                            }
                        case "\n\t":
                            if false {
                                var v11 float64 = (0.1 / v5)
                                if (!(38 < 255)) {
                                    var v12 int8 = 18
                                    v3 = v3
                                    _ = v12
                                }
                                fmt.Println(v5)
                                _ = v11
                            } else {
                                if ((!(-653 >= 1)) || (696 == 2147483647)) {
                                    switch "" {
                                    case "\n\t":
                                        if true {
                                            v13 := GenS2{F0: 1}
                                            var v14 int = (v3 - (v3 | v3))
                                            _ = v13
                                            _ = v14
                                        }
                                    case "hello":
                                        switch 1 {
                                        case 32767:
                                            v15 := GenS3{F0: ""}
                                            _ = v15
                                        case -366:
                                            if (18446744073709551615 <= 42) {
                                                var v16 int64 = 337
                                                _ = v16
                                            } else {
                                                var v17 uint16 = 32767
                                                _ = v17
                                            }
                                        case 4294967295:
                                            if (!false) {
                                                var v18 int8 = 92
                                                v3 = v3
                                                v19 := GenS3{F0: "\n\t"}
                                                _ = v18
                                                _ = v19
                                            } else {
                                                var v20 uint8 = 128
                                                var v21 int8 = 101
                                                _ = v20
                                                _ = v21
                                            }
                                        }
                                    }
                                    var v22 uint8 = 68
                                    for fi23 := 0; fi23 < 7; fi23++ {
                                        for fi24 := 0; fi24 < 13; fi24++ {
                                            v22 -= 144
                                            v3 = 32768
                                            var v25 uint32 = 2147483648
                                            _ = fi24
                                            _ = v25
                                        }
                                        v1 = append(v1, "a")
                                        fmt.Println(v2, v5)
                                        _ = fi23
                                    }
                                    _ = v22
                                } else {
                                    for fi26 := 0; fi26 < 6; fi26++ {
                                        v2 = v2
                                        _ = fi26
                                    }
                                    _ = len(v1)
                                }
                            }
                        default:
                            switch "foo bar" {
                            case "a":
                                if (v6 != 65535) {
                                    _ = len(v1)
                                }
                            case "\n\t":
                                switch (^(v3)) {
                                case 42:
                                    fmt.Println(v2, v5)
                                default:
                                    v27 := GenS2{F0: 332}
                                    _ = v27
                                }
                            }
                        }
                    }
                    switch ("hello" + "") {
                    case "hello":
                        for fi28 := 0; fi28 < 7; fi28++ {
                            var v29 int = -4
                            var v30 uint = (v6 / ((v6) | 1))
                            switch "\n\t" {
                            case "\n\t":
                                var v31 uint32 = 2
                                _ = v31
                            case "foo bar":
                                switch "hello" {
                                case "foo bar":
                                    var v32 int16 = 1000
                                    _ = v32
                                }
                            default:
                                v33 := GenS3{F0: ("hello" + ("a" + (("\n\t" + (("hello" + "a") + (("\n\t" + "\n\t") + (("\n\t" + "a") + "foo bar")))) + (("a" + "a") + "hello"))))}
                                _ = v33
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                    _ = v9
                }
            } else {
                switch "\n\t" {
                case "hello":
                    v3 = v3
                default:
                    v6 = (v6 & v6)
                }
            }
        } else {
            var v34 string = ("" + "hello")
            v2 += 0
            _ = v34
        }
    default:
        if (!(2 != uint32(v6))) {
            var v35 int16 = 1000
            _ = v35
        }
    }
    v1 = append(v1, (("\n\t" + "a") + ("" + "")))
    v36 := GenS4{F0: 425, F1: 792, F2: 65535}
    v37 := GenS1{F0: 77, F1: 42, F2: (874 > 416)}
    fmt.Println(v6, v3)
    v38 := GenS3{F0: (("foo bar" + "hello") + "\n\t")}
    if false {
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    }
    _ = len(v1)
    v39 := GenS2{F0: 18446744073709551615}
    _ = len(v1)
    if (false && (!("hello" > (("hello" + "") + "a")))) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v6 -= v6
        for fi40 := 0; fi40 < 19; fi40++ {
            var v41 uint = (^(v6))
            v6 += (v6 ^ 2)
            fmt.Println(v41, v6, v3)
            _ = fi40
            _ = v41
        }
        var v42 int32 = v2
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v5
    _ = v6
}
