package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int8; F2 bool }
func genFn0(p0 int32, p1 int) int8 {
    if true {
        var v7 bool = false
        _ = v7
    }
    var v8 float64 = float64(3.14159)
    p1 = (p1 >> ((65535) & 63))
    _ = p0
    _ = p1
    _ = v8
    return -105
}
func genFn1(p0 bool, p1 uint64, p2 uint64) uint16 {
    p1 = p2
    if ("foo bar" != ("hello" + ("\n\t" + ("a" + "foo bar")))) {
        fmt.Println(p2, p2, p0)
    } else {
        switch 0 {
        case 970:
            var v9 uint = 0
            _ = v9
        case 127:
            v10 := GenS1{F0: 611, F1: 65535}
            _ = v10
        }
        v11 := []int{}
        switch "" {
        case "hello":
            switch "foo bar" {
            case "\n\t":
                var v12 uint16 = 815
                _ = v12
            case "foo bar":
                if (-1 >= -943) {
                    if p0 {
                        var v13 int8 = 100
                        v14 := GenS2{F0: 57, F1: 4294967295}
                        for fi15 := 0; fi15 < 13; fi15++ {
                            p0 = false
                            for fi16 := 0; fi16 < 11; fi16++ {
                                if true {
                                    var v17 bool = (2147483647 == 2)
                                    var v18 string = "a"
                                    var v19 uint8 = 93
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                                v11 = append(v11, (fi15 - (fi16 + fi15)))
                                if len(v11) > 0 {
                                    v11[0] = fi16
                                }
                                _ = fi16
                            }
                            for fi20 := 0; fi20 < 7; fi20++ {
                                var v21 string = ("a" + "foo bar")
                                _ = fi20
                                _ = v21
                            }
                            _ = fi15
                        }
                        _ = v13
                        _ = v14
                    } else {
                        v22 := GenS1{F0: 981, F1: 716}
                        p1 -= (p1 >> ((994) & 63))
                        _ = v22
                    }
                }
            case "":
                p2 = (p1 | 18446744073709551615)
            default:
                var v23 uint32 = 1
                _ = v23
            }
        case "foo bar":
            if p0 {
                if p0 {
                    fmt.Println(p0, p1, p1)
                }
                for fi24 := 0; fi24 < 3; fi24++ {
                    p2 -= (p2 >> ((667) & 63))
                    v25 := GenS3{F0: -51, F1: 858}
                    _ = fi24
                    _ = v25
                }
                fmt.Println(p2)
            } else {
                if p0 {
                    v11 = append(v11, 773)
                } else {
                    var v26 uint64 = (p1 >> ((0) & 63))
                    _ = v26
                }
            }
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []string{"\n\t", ("a" + "hello"), "a", (("hello" + ("" + (((("foo bar" + (("a" + "hello") + ("a" + ""))) + ((("a" + (("a" + "foo bar") + (("" + "a") + "a"))) + "\n\t") + (("\n\t" + "a") + ("" + "a")))) + "foo bar") + "\n\t"))) + "foo bar")}
    var v2 int = 142
    var v3 int8 = 1
    var v4 int8 = 26
    v5 := []string{"a", "a", ("a" + "a"), "hello"}
    v6 := make(map[uint32]uint32)
    v27 := GenS4{F0: 1, F1: (v4 / ((v4) | 1)), F2: true}
    v3 = v4
    if true {
        switch "hello" {
        case "hello":
            v28 := GenS1{F0: 2, F1: 2147483647}
            _ = v28
        case "":
            switch ("\n\t" + "a") {
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            case "hello":
                v2 += (v2 + v2)
            case "\n\t":
                v2 -= -25
            }
        }
        v2 -= (v2 << ((255) & 63))
        v2 += v2
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 uint16 = 46
        _ = fi29
        _ = v30
    }
    var v31 int16 = 32767
    v32 := GenS3{F0: v3, F1: 459}
    v3 -= 127
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
