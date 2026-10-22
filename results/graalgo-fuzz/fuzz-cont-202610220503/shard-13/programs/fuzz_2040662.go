package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint; F1 string; F2 uint64 }
func genFn0() int64 {
    v5 := GenS1{F0: ("foo bar" + "a"), F1: -480}
    v6 := GenS2{F0: 4294967295, F1: "", F2: 574}
    var v7 string = "hello"
    fmt.Println(v7)
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 = v7
        v7 = v7
        var v9 float64 = (0.0 + 3.14159)
        _ = fi8
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 string, p1 uint8) int64 {
    var v10 string = "a"
    switch v10 {
    case "foo bar":
        var v11 int32 = -1
        _ = v11
    }
    switch "a" {
    case "":
        for fi12 := 0; fi12 < 7; fi12++ {
            v10 = p0
            v13 := GenS1{F0: "foo bar", F1: 0}
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        p0 = v10
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 float32 = (-2.25 + (1.5 - (0.0 + -2.25)))
            v10 = ((p0 + ("\n\t" + "foo bar")) + "foo bar")
            v10 = (("a" + ((("a" + ("hello" + ("foo bar" + ((("" + ("foo bar" + p0)) + "foo bar") + p0)))) + p0) + (p0 + ("\n\t" + p0)))) + ((p0 + p0) + v10))
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v17 uint = 8
    _ = p0
    _ = p1
    _ = v10
    _ = v17
    return -272
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -2147483648
    v3 := []string{"hello", "", ("foo bar" + "hello"), ""}
    v4 := make(map[int]string)
    if true {
        v2 = v2
    } else {
        switch v2 {
        case 144:
            var v18 uint8 = 0
            _ = v18
        case 127:
            if true {
                v19 := GenS2{F0: 65535, F1: ("foo bar" + ((("hello" + "foo bar") + ("a" + "a")) + "\n\t")), F2: 0}
                v2 = (v2 << ((255) & 63))
                var v20 uint16 = 365
                _ = v19
                _ = v20
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 uint64 = 607
                    switch ("\n\t" + (("a" + "a") + "\n\t")) {
                    case "":
                        var v23 int16 = -333
                        _ = v23
                    case "hello":
                        v22 = 42
                    case "\n\t":
                        for fi24 := 0; fi24 < 10; fi24++ {
                            var v25 uint32 = 400
                            var v26 bool = true
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    default:
                        v27 := v4[(v2 - 1)]
                        _ = v27
                    }
                    _ = len(v1)
                    _ = fi21
                    _ = v22
                }
            }
        case -399:
            v2 = v2
        }
        for fi28 := 0; fi28 < 2; fi28++ {
            var v29 uint16 = 32767
            _ = fi28
            _ = v29
        }
    }
    var v30 uint64 = 42
    var v31 uint = 18446744073709551615
    var v32 float64 = 3.14159
    if true {
        fmt.Println(v31)
    } else {
        fmt.Println(v31)
    }
    switch "" {
    case "hello":
        _ = len(v1)
    case "\n\t":
        v33 := GenS2{F0: 1, F1: "a", F2: (v30 - v30)}
        _ = v33
    case "foo bar":
        var v34 uint = 377
        _ = v34
    default:
        var v35 uint = v31
        _ = v35
    }
    v30 -= (v30 + uint64(v31))
    v30 = 2
    if (v30 >= (v30 - (v30 >> ((v31) & 63)))) {
        v1["foo bar"] = 2147483647
        v30 -= 521
    } else {
        v36 := GenS1{F0: "hello", F1: -337}
        if true {
            v30 += (^(v30))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v37 := GenS2{F0: (^(v31)), F1: (("hello" + "hello") + "hello"), F2: (v30 % (((v30 / ((42) | 1))) | 1))}
        _ = v36
        _ = v37
    }
    v32 -= 0.1
    _ = len(v4)
    v2 = (v2 / (((v2 << ((0) & 63))) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v4)
    v38 := v4[v2]
    switch v38 {
    case "\n\t":
        fmt.Println(v31, v2, v38)
    case "foo bar":
        for fi39 := 0; fi39 < 4; fi39++ {
            if (65535 >= v31) {
                fmt.Println(v38, v31, fi39)
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v30 = v30
                if len(v3) > 0 {
                    v3[0] = v38
                }
            }
            v30 += v30
            _ = fi39
        }
    }
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v4
}
