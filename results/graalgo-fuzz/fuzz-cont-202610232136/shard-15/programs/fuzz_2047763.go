package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }

func main() {
    var v1 uint32 = 2
    var v2 int64 = -1
    var v3 bool = (false || false)
    v4 := []string{("a" + (("foo bar" + "foo bar") + "a"))}
    v5 := []uint32{(v1 ^ 2147483648), (v1 * v1)}
    var v6 uint8 = 1
    switch "\n\t" {
    case "a":
        v6 = (v6 % (((v6 % ((v6) | 1))) | 1))
    }
    switch ("\n\t" + (("\n\t" + "") + "foo bar")) {
    case "hello":
        fmt.Println(v3, v2)
    case "a":
        for fi7 := 0; fi7 < 5; fi7++ {
            if (!(float32(0.1) > ((0.1 / -2.25) + ((float32(0.0) * -2.25) * -2.25)))) {
                for fi8 := 0; fi8 < 5; fi8++ {
                    v6 -= (v6 + v6)
                    _ = fi8
                }
            } else {
                switch "a" {
                case "hello":
                    for fi9 := 0; fi9 < 18; fi9++ {
                        if (4294967295 >= 2) {
                            _ = len(v5)
                        } else {
                            if true {
                                v5 = append(v5, (v1 * v1))
                                switch fi7 {
                                case 1:
                                    v10 := GenS1{F0: v1, F1: 56}
                                    _ = v10
                                case -9223372036854775808:
                                    var v11 float64 = 1.5
                                    _ = v11
                                }
                            }
                        }
                        _ = fi9
                    }
                }
            }
            if (true || false) {
                if (!(!("\n\t" < ("hello" + "")))) {
                    if (fi7 >= fi7) {
                        if v3 {
                            var v12 int = 32768
                            fmt.Println(v2, v6)
                            for fi13 := 0; fi13 < 13; fi13++ {
                                v14 := GenS1{F0: 335, F1: -88}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if (2 <= 0) {
                                    var v15 uint8 = v6
                                    v1 = v1
                                    var v16 uint16 = 2
                                    fmt.Println(v2, v15)
                                    _ = v15
                                    _ = v16
                                }
                                _ = fi13
                                _ = v14
                            }
                            _ = v12
                        }
                    }
                    for fi17 := 0; fi17 < 19; fi17++ {
                        fmt.Println(v2)
                        _ = fi17
                    }
                }
                if true {
                    v6 += (^(v6))
                    switch ("\n\t" + ("foo bar" + "")) {
                    case "":
                        fmt.Println(fi7, v3)
                    case "a":
                        v3 = ((!v3) && v3)
                    }
                    var v18 float32 = 1.5
                    _ = v18
                } else {
                    var v19 uint64 = 42
                    _ = v19
                }
                v1 += v1
            }
            _ = fi7
        }
    }
    switch 0 {
    case -9223372036854775808:
        var v20 int64 = (v2 - (^(v2)))
        _ = v20
    }
    v21 := GenS1{F0: v1, F1: -128}
    v3 = (v3 && v3)
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int = (fi22 | fi22)
        var v24 uint8 = (v6 + 226)
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 uint = uint(v6)
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: v1, F1: 0}
    v1 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
