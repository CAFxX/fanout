package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 string; F1 int8; F2 int }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint8; F1 bool }

func main() {
    var v1 float64 = (3.14159 * 0.1)
    var v2 uint16 = 1
    v3 := []string{("foo bar" + ""), "", "hello"}
    var v4 int32 = 1
    var v5 uint64 = 0
    v1 -= 0.0
    _ = len(v3)
    if true {
        if (!(200 >= 11)) {
            v3 = append(v3, "\n\t")
            v1 += (v1 + (v1 - (3.14159 - v1)))
            var v6 uint16 = v2
            _ = v6
        } else {
            switch 330 {
            case 2147483647:
                for fi7 := 0; fi7 < 19; fi7++ {
                    var v8 uint16 = 32767
                    v1 = 1e10
                    _ = fi7
                    _ = v8
                }
            case 42:
                switch "\n\t" {
                case "foo bar":
                    v9 := GenS1{F0: 127, F1: (!true), F2: 1}
                    _ = v9
                default:
                    v1 = float64((1.5 - 1.5))
                }
            case 32767:
                if true {
                    switch 915 {
                    case -157:
                        if false {
                            v3 = append(v3, "\n\t")
                            for fi10 := 0; fi10 < 13; fi10++ {
                                if (-9223372036854775808 > 42) {
                                    v5 -= v5
                                    v1 = (3.14159 * (v1 + v1))
                                    v5 -= 312
                                } else {
                                    fmt.Println(v5)
                                    v5 = v5
                                    if false {
                                        v4 -= v4
                                        v1 = 3.14159
                                        var v11 float64 = (0.1 - float64(0.1))
                                        _ = v11
                                    }
                                }
                                _ = fi10
                            }
                            v12 := GenS1{F0: 255, F1: true, F2: 255}
                            _ = v12
                        }
                    case 32768:
                        if false {
                            if true {
                                v1 -= (1e10 / (0.1 - v1))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v13 bool = true
                                _ = v13
                            }
                            if (!false) {
                                var v14 bool = (!false)
                                _ = v14
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                } else {
                    v15 := GenS2{F0: v2, F1: 894}
                    _ = v15
                }
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5, v1)
    fmt.Println(v2)
    switch "a" {
    case "":
        v16 := GenS3{F0: "\n\t", F1: -44, F2: 336}
        _ = v16
    case "a":
        fmt.Println(v5, v2)
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v17 := GenS3{F0: ("hello" + (("" + ("a" + "foo bar")) + "hello")), F1: 42, F2: 42}
    for fi18 := 0; fi18 < 19; fi18++ {
        if false {
            v19 := GenS1{F0: 1, F1: true, F2: 2}
            _ = v19
        }
        v20 := GenS4{F0: 1, F1: true}
        _ = fi18
        _ = v20
    }
    switch 0 {
    case 4294967295:
        if len(v3) > 0 {
            v3[0] = ((("" + (("hello" + "") + (("\n\t" + ("\n\t" + "")) + "\n\t"))) + (("foo bar" + "a") + "\n\t")) + "foo bar")
        }
    }
    if (-862 >= 570) {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        v21 := GenS5{F0: 42, F1: false}
        _ = v21
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
