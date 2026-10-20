package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint32) int64 {
    v6 := GenS1{F0: 128}
    v7 := GenS2{F0: 0}
    p0 -= 1
    p0 = (^(p0))
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 string = ("foo bar" + "foo bar")
    var v2 int8 = -100
    v3 := []int{-429, 42, -139, 32768}
    v4 := []int{}
    v5 := make(map[string]int)
    switch v1 {
    case "hello":
        switch (v1 + "a") {
        case "":
            for fi8 := 0; fi8 < 11; fi8++ {
                if (!((!(4294967295 <= 30)) && (65535 >= 491))) {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        var v10 uint32 = 2147483648
                        var v11 int = -639
                        v10 = 2147483648
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
                var v12 uint8 = 128
                _ = fi8
                _ = v12
            }
        case "\n\t":
            v1 = "hello"
        case "hello":
            switch "\n\t" {
            case "hello":
                var v13 uint = 21
                _ = v13
            }
        default:
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 bool = ((v1 + (((("hello" + v1) + v1) + v1) + (v1 + "foo bar"))) <= "foo bar")
                _ = fi14
                _ = v15
            }
        }
    case "a":
        for fi16 := 0; fi16 < 8; fi16++ {
            v1 = v1
            _ = fi16
        }
    case "":
        if true {
            var v17 bool = false
            fmt.Println(v17, v2)
            for fi18 := 0; fi18 < 10; fi18++ {
                for fi19 := 0; fi19 < 5; fi19++ {
                    if len(v4) > 0 {
                        v4[0] = (fi19 >> ((0) & 63))
                    }
                    _ = fi19
                }
                v1 = ("\n\t" + ((("foo bar" + "\n\t") + v1) + "hello"))
                for fi20 := 0; fi20 < 12; fi20++ {
                    v1 = v1
                    _ = fi20
                }
                _ = fi18
            }
            _ = v17
        } else {
            var v21 float32 = (0.1 * ((1.5 - 1e10) * 3.14159))
            switch v1 {
            case "foo bar":
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 uint8 = 207
                    _ = fi22
                    _ = v23
                }
            case "hello":
                v21 += (v21 + v21)
            case "a":
                v1 = v1
            }
            _ = v21
        }
    default:
        var v24 int8 = v2
        _ = v24
    }
    fmt.Println(v2, v2)
    fmt.Println(v2, v1)
    if true {
        for fi25 := 0; fi25 < 4; fi25++ {
            v2 = v2
            switch ((v1 + ("foo bar" + "")) + ((v1 + (v1 + v1)) + (((v1 + (v1 + ((v1 + "") + v1))) + v1) + ""))) {
            case "\n\t":
                v5[v1] = fi25
            default:
                switch (v1 + "foo bar") {
                case "hello":
                    if (genFn0(160) < 0) {
                        v2 = (v2 & -1)
                        fmt.Println(v1, v2, fi25)
                        v26 := GenS3{F0: -9223372036854775808, F1: 0}
                        _ = v26
                    }
                case "":
                    if false {
                        if false {
                            v4 = append(v4, 9223372036854775807)
                        } else {
                            _ = len(v3)
                            if true {
                                v1 = (((v1 + v1) + v1) + v1)
                            } else {
                                var v27 uint16 = 383
                                _ = v27
                            }
                        }
                        fmt.Println(v1)
                    } else {
                        v28 := GenS1{F0: 9223372036854775807}
                        _ = v28
                    }
                case "\n\t":
                    if false {
                        var v29 int = 4294967295
                        fmt.Println(v29)
                        var v30 int8 = (v2 * 1)
                        _ = v29
                        _ = v30
                    } else {
                        if false {
                            delete(v5, v1)
                        } else {
                            v1 = (("" + ((v1 + v1) + v1)) + v1)
                        }
                    }
                }
            }
            var v31 uint16 = 1
            _ = fi25
            _ = v31
        }
        v1 = "\n\t"
        for fi32 := 0; fi32 < 17; fi32++ {
            v33 := GenS1{F0: int(v2)}
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        var v35 uint16 = 65535
        if (((genFn0(972) == 9223372036854775807) || (false || (false && false))) || false) {
            var v36 float64 = 1e10
            _ = len(v5)
            fmt.Println(v35, v36, v36)
            _ = v36
        }
        _ = fi34
        _ = v35
    }
    v1 = (v1 + v1)
    var v37 uint64 = 0
    var v38 string = (v1 + (("hello" + ((("foo bar" + ("" + ((v1 + "a") + (("\n\t" + ("" + v1)) + "foo bar")))) + "hello") + ("hello" + ((v1 + "foo bar") + v1)))) + v1))
    v38 = ((v1 + v38) + v38)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
