package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 bool; F1 int16; F2 uint8 }

func main() {
    v1 := []int{582}
    var v2 string = "foo bar"
    v3 := []string{"\n\t"}
    var v4 string = "\n\t"
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint = 0
        var v7 int16 = 1000
        v4 = v4
        _ = fi5
        _ = v6
        _ = v7
    }
    if false {
        var v8 int32 = 1
        _ = v8
    } else {
        var v9 uint8 = 157
        for fi10 := 0; fi10 < 15; fi10++ {
            v2 = "foo bar"
            _ = fi10
        }
        _ = v9
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v3 = append(v3, (v2 + "foo bar"))
        var v12 uint32 = 1
        v2 = ""
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        if ((!false) && false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v3 = append(v3, "\n\t")
        }
        v1 = append(v1, (fi13 / ((fi13) | 1)))
        _ = fi13
    }
    fmt.Println(v2, v4)
    if true {
        for fi14 := 0; fi14 < 20; fi14++ {
            v4 = ""
            for fi15 := 0; fi15 < 20; fi15++ {
                fmt.Println(fi14)
                var v16 uint16 = 0
                var v17 int8 = -1
                _ = fi15
                _ = v16
                _ = v17
            }
            v1 = append(v1, fi14)
            _ = fi14
        }
        v18 := GenS1{F0: 0, F1: -24}
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(fi19, v4, fi19)
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint8 = 2
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            if (!(-974 > -32768)) {
                if (42 < 0) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v4 = v4
                    if (980 < 344) {
                        v23 := GenS2{F0: false, F1: 42, F2: 42}
                        _ = v23
                    } else {
                        if (!true) {
                            if true {
                                var v24 uint8 = 1
                                _ = v24
                            }
                            v2 = v4
                            var v25 int = (fi22 ^ 1)
                            var v26 bool = true
                            v25 = -348
                            _ = v25
                            _ = v26
                        } else {
                            fmt.Println(fi19, v2, v4)
                            v2 = v4
                            fmt.Println(v4)
                            var v27 int8 = 100
                            _ = v27
                        }
                        v28 := GenS1{F0: (fi22 >> ((518) & 63)), F1: int8(fi19)}
                        _ = v28
                    }
                } else {
                    v2 = ("" + ("foo bar" + "hello"))
                }
                v2 = v2
                v2 = "a"
            }
            if false {
                switch -442 {
                case 9223372036854775807:
                    var v29 int64 = 9223372036854775807
                    _ = v29
                case -646:
                    switch -1 {
                    case 1:
                        fmt.Println(v4, fi22, v4)
                    case -187:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        _ = len(v1)
                    }
                case -727:
                    var v30 int32 = -824
                    _ = v30
                default:
                    v4 = v4
                }
                v2 = "hello"
            }
            fmt.Println(fi22)
            _ = fi22
        }
        _ = fi19
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 uint16 = 543
        _ = fi31
        _ = v32
    }
    var v33 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
