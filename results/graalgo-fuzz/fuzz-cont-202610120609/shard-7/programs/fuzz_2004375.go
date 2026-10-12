package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint64, p1 uint16) int16 {
    switch 42 {
    case 377:
        var v5 uint16 = 255
        _ = v5
    case -1:
        p1 += p1
    default:
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 7; fi7++ {
                if false {
                    v8 := GenS1{F0: (p0 & p0)}
                    p1 += 598
                    _ = v8
                }
                fmt.Println(p0, fi6, p0)
                _ = fi7
            }
            v9 := []int{}
            _ = fi6
            _ = v9
        }
    }
    var v10 int = 127
    _ = p0
    _ = p1
    _ = v10
    return 1
}
func genFn1(p0 string) uint16 {
    var v11 int16 = 260
    v12 := GenS1{F0: 0}
    var v13 uint16 = 431
    switch -2147483648 {
    case -1:
        var v14 uint64 = 123
        _ = v14
    default:
        v11 += 42
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return v13
}

func main() {
    var v1 uint8 = 0
    var v2 int64 = 0
    var v3 int8 = 1
    v4 := []int{2147483647, -354, -2147483648, 524}
    if (!true) {
        if len(v4) > 0 {
            v4[0] = 32768
        }
        v15 := GenS2{F0: 1, F1: (v2 / (((^(v2))) | 1))}
        _ = v15
    } else {
        switch "a" {
        case "hello":
            var v16 uint = 619
            _ = v16
        case "a":
            v17 := GenS2{F0: 255, F1: (-(v2))}
            _ = v17
        }
        v18 := GenS1{F0: 1}
        v19 := GenS1{F0: 0}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi20
    }
    v2 = (v2 + 784)
    var v21 uint32 = 2
    fmt.Println(v3, v21, v21)
    var v22 uint32 = (v21 & (v21 - (^(v21))))
    var v23 int64 = v2
    v24 := GenS3{F0: 0, F1: 959}
    if true {
        if true {
            for fi25 := 0; fi25 < 7; fi25++ {
                for fi26 := 0; fi26 < 12; fi26++ {
                    v27 := GenS3{F0: 0, F1: 18446744073709551615}
                    var v28 float64 = ((float64(0.0) / 3.14159) - 1.5)
                    var v29 uint8 = uint8(v22)
                    var v30 float32 = (-2.25 + 0.0)
                    if (!(!false)) {
                        if (false || false) {
                            var v31 int = (fi26 << ((255) & 63))
                            if (!false) {
                                var v32 string = (("\n\t" + (("hello" + ((("foo bar" + "foo bar") + ("hello" + "\n\t")) + ("a" + ""))) + "foo bar")) + ("" + "\n\t"))
                                v3 -= 100
                                var v33 uint64 = 0
                                var v34 int32 = 0
                                if len(v4) > 0 {
                                    v4[0] = (^(fi26))
                                }
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            _ = v31
                        } else {
                            var v35 uint = 0
                            _ = v35
                        }
                        v36 := GenS2{F0: 255, F1: 9223372036854775807}
                        var v37 int32 = 1
                        _ = v36
                        _ = v37
                    } else {
                        fmt.Println(fi25, fi26)
                        if len(v4) > 0 {
                            v4[0] = 9223372036854775807
                        }
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                for fi38 := 0; fi38 < 13; fi38++ {
                    fmt.Println(fi25)
                    var v39 int = fi25
                    _ = fi38
                    _ = v39
                }
                v23 = 473
                _ = fi25
            }
            fmt.Println(v23, v22)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
