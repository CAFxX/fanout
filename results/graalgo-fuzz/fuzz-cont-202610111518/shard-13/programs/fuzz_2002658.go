package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 uint; F1 int64; F2 uint16 }
func genFn0(p0 int32) uint16 {
    v4 := []int{9223372036854775807, 34, 856}
    _ = len(v4)
    for fi5 := 0; fi5 < 14; fi5++ {
        if false {
            p0 += -1
            for fi6 := 0; fi6 < 2; fi6++ {
                switch ("a" + "hello") {
                case "foo bar":
                    v7 := GenS1{F0: (!true), F1: 531, F2: ("\n\t" + ("\n\t" + ("foo bar" + "foo bar")))}
                    _ = v7
                case "":
                    p0 = (p0 & p0)
                case "\n\t":
                    v4 = append(v4, fi5)
                }
                _ = fi6
            }
            v8 := GenS2{F0: 2, F1: 9223372036854775807, F2: 42}
            _ = v8
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 bool = (false && true)
    v2 := make(map[uint32]string)
    var v3 int16 = 1000
    if true {
        fmt.Println(v1, v1)
    } else {
        switch 4294967295 {
        case -443:
            if ((9223372036854775807 == -9223372036854775808) && v1) {
                fmt.Println(v1)
                if (670 > 700) {
                    v2[200] = "\n\t"
                    v2[621] = ("hello" + "foo bar")
                    for fi9 := 0; fi9 < 18; fi9++ {
                        var v10 int32 = 42
                        for fi11 := 0; fi11 < 8; fi11++ {
                            v12 := GenS2{F0: 746, F1: -9223372036854775808, F2: 2}
                            _ = fi11
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 11; fi13++ {
                            switch "foo bar" {
                            case "hello":
                                var v14 int8 = -1
                                _ = v14
                            default:
                                v3 += (v3 | (v3 + v3))
                            }
                            switch fi13 {
                            case 42:
                                v15 := GenS2{F0: 0, F1: -1, F2: 255}
                                _ = v15
                            case -890:
                                if true {
                                    if v1 {
                                        var v16 int32 = (v10 ^ (^(v10)))
                                        v17 := GenS1{F0: v1, F1: genFn0(v10), F2: "a"}
                                        v18 := GenS1{F0: v1, F1: 1, F2: (("foo bar" + "") + "a")}
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    }
                                } else {
                                    v1 = true
                                }
                            case -2147483648:
                                v19 := []int{(fi13 | -734), 127}
                                _ = v19
                            }
                            _ = fi13
                        }
                        _ = fi9
                        _ = v10
                    }
                } else {
                    fmt.Println(v3, v1)
                }
            }
        }
    }
    delete(v2, 511)
    var v20 uint = 612
    var v21 int64 = 9223372036854775807
    var v22 int64 = v21
    if (309 <= 1) {
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            v23 := GenS2{F0: (v20 % ((v20) | 1)), F1: v22, F2: 32767}
            _ = v23
        }
    }
    var v24 float64 = (-2.25 + 0.1)
    v22 = v22
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
}
