package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{2147483647, 2147483647}
    v3 := make(map[string]string)
    v4 := []int{1}
    switch "foo bar" {
    case "\n\t":
        v2 = append(v2, 9223372036854775807)
    case "":
        var v5 int8 = -90
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 7; fi6++ {
            fmt.Println(fi6, v1)
            _ = fi6
        }
    }
    if false {
        if (2 < 2) {
            if (648 > 554) {
                var v7 int16 = -301
                _ = v7
            }
        } else {
            v1 = 767
            v3["\n\t"] = "hello"
            fmt.Println(v1)
        }
        v8 := v3[(("hello" + "a") + "foo bar")]
        fmt.Println(v1, v1)
        _ = v8
    }
    v1 = (^(v1))
    for fi9 := 0; fi9 < 16; fi9++ {
        if len(v4) > 0 {
            v4[0] = fi9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            if (265 > 1) {
                if true {
                    v3["\n\t"] = "\n\t"
                }
            } else {
                var v11 int32 = int32(fi9)
                v1 = v1
                switch 630 {
                case 391:
                    v3["\n\t"] = (("\n\t" + ("hello" + "hello")) + ("" + (("" + ("\n\t" + "foo bar")) + "")))
                case 2147483647:
                    fmt.Println(v1)
                }
                _ = v11
            }
            _ = fi10
        }
        v1 = (v1 * v1)
        _ = fi9
    }
    if (42 <= 65535) {
        var v12 float32 = float32(1.5)
        _ = v12
    } else {
        v1 -= (v1 - v1)
        var v13 bool = false
        for fi14 := 0; fi14 < 11; fi14++ {
            switch ("foo bar" + ("hello" + "hello")) {
            case "":
                if len(v2) > 0 {
                    v2[0] = fi14
                }
            }
            _ = fi14
        }
        _ = v13
    }
    switch 1 {
    case 895:
        var v15 string = "foo bar"
        _ = v15
    case 32768:
        switch 0 {
        case 793:
            v1 += (v1 * (^(v1)))
        }
    default:
        v16 := GenS1{F0: ("\n\t" + ("\n\t" + "")), F1: 2}
        _ = v16
    }
    var v17 uint = 0
    var v18 int8 = 1
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS1{F0: "foo bar", F1: 1}
        var v21 uint8 = 128
        _ = fi19
        _ = v20
        _ = v21
    }
    v1 = v1
    var v22 uint64 = 1
    switch 42 {
    case -723:
        if (0.1 >= (-2.25 * -2.25)) {
            switch 2147483647 {
            case -924:
                var v23 int8 = (v18 % ((-60) | 1))
                _ = v23
            }
            v1 = v1
            v24 := GenS1{F0: "hello", F1: 65535}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 3; fi25++ {
                fmt.Println(v1, v17, v17)
                v26 := GenS2{F0: (-(v1)), F1: 2147483647}
                if true {
                    var v27 uint = (v17 + (v17 >> (((v17 % ((v17) | 1))) & 63)))
                    for fi28 := 0; fi28 < 4; fi28++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v1 = (v1 % ((v1) | 1))
                        if ((true && ((-(v18)) != v18)) && true) {
                            switch fi28 {
                            case -1:
                                fmt.Println(v18)
                            case -179:
                                if ((v18 >> (((v27 >> ((v17) & 63))) & 63)) > (v18 ^ (v18 + v18))) {
                                    if false {
                                        v1 = (^(v1))
                                        v18 = v18
                                        v27 = (v17 * v17)
                                        v18 = (v18 >> (((^(v17))) & 63))
                                        if len(v4) > 0 {
                                            v4[0] = (fi28 << ((v17) & 63))
                                        }
                                        v29 := GenS3{F0: 169, F1: -461}
                                        _ = v29
                                    } else {
                                        if false {
                                            fmt.Println(fi28)
                                        }
                                    }
                                } else {
                                    v18 += v18
                                }
                            }
                            v1 += v1
                            fmt.Println(fi25)
                        }
                        _ = fi28
                    }
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
        }
    default:
        var v30 int = -589
        _ = v30
    }
    v31 := GenS4{F0: (v22 / (((v22 | 754)) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v4
}
