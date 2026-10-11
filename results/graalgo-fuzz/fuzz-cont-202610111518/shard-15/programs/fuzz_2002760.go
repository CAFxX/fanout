package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64; F2 int32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int }

func main() {
    var v1 bool = false
    var v2 uint16 = 42
    var v3 uint16 = v2
    v4 := make(map[string]int)
    switch "hello" {
    case "\n\t":
        if (!(false && (!(0 >= 42)))) {
            if (false && v1) {
                for fi5 := 0; fi5 < 3; fi5++ {
                    v6 := GenS1{F0: 0}
                    v7 := GenS1{F0: 591}
                    _ = fi5
                    _ = v6
                    _ = v7
                }
                switch "\n\t" {
                case "hello":
                    var v8 string = (((("\n\t" + ("a" + "")) + (("a" + "foo bar") + "a")) + ("\n\t" + ("hello" + "a"))) + ("\n\t" + "foo bar"))
                    _ = v8
                case "foo bar":
                    v4[("foo bar" + ("a" + ((("a" + "hello") + "") + "\n\t")))] = 524
                default:
                    v1 = (100 < -128)
                }
                for fi9 := 0; fi9 < 11; fi9++ {
                    if (0 < 2) {
                        v3 = v2
                        v10 := []int{(fi9 / (((fi9 | -408)) | 1))}
                        _ = v10
                    }
                    fmt.Println(fi9, v3)
                    if (1 < 42) {
                        if (!v1) {
                            var v11 string = "foo bar"
                            v12 := GenS2{F0: -1, F1: 1, F2: -705}
                            switch v11 {
                            case "a":
                                fmt.Println(v1, v1)
                            case "hello":
                                fmt.Println(v1, v1)
                            default:
                                fmt.Println(v2, v11, v2)
                            }
                            _ = v11
                            _ = v12
                        }
                        fmt.Println(v2, v2, v2)
                    }
                    _ = fi9
                }
            }
        } else {
            switch 621 {
            case 880:
                var v13 int16 = 0
                _ = v13
            case 0:
                switch ("a" + ((("" + ("hello" + "")) + (("\n\t" + ("\n\t" + "foo bar")) + ("" + "a"))) + "hello")) {
                case "hello":
                    v3 += (v3 & v2)
                case "foo bar":
                    fmt.Println(v2)
                }
            case 9223372036854775807:
                v14 := GenS3{F0: "\n\t", F1: 100}
                _ = v14
            }
        }
    case "":
        for fi15 := 0; fi15 < 4; fi15++ {
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS4{F0: 473}
                v1 = (!true)
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    }
    var v18 uint64 = 474
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS4{F0: 234}
        _ = fi19
        _ = v20
    }
    switch 32768 {
    case 266:
        var v21 int16 = 0
        _ = v21
    case -2147483648:
        _ = len(v4)
    case 32768:
        if v1 {
            var v22 uint16 = (v3 | 32767)
            _ = v22
        }
    default:
        v2 += (v3 * v2)
    }
    fmt.Println(v3)
    var v23 int = 127
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := []int{-9223372036854775808, (v23 % (((v23 >> ((18446744073709551615) & 63))) | 1))}
        v26 := GenS1{F0: (v18 * 837)}
        _ = fi24
        _ = v25
        _ = v26
    }
    switch "\n\t" {
    case "":
        var v27 uint64 = 18446744073709551615
        _ = v27
    default:
        v28 := []int{(v23 + v23), (v23 % ((-522) | 1)), (v23 % ((v23) | 1))}
        _ = v28
    }
    var v29 float64 = (float64(1e10) + 1e10)
    if false {
        var v30 uint16 = v3
        v30 -= (v2 << ((255) & 63))
        var v31 bool = (!(2 <= 255))
        _ = v30
        _ = v31
    } else {
        delete(v4, "")
        for fi32 := 0; fi32 < 11; fi32++ {
            fmt.Println(v2, v1, v2)
            delete(v4, (("hello" + "hello") + ("hello" + "\n\t")))
            v18 = v18
            _ = fi32
        }
        var v33 uint32 = 721
        _ = v33
    }
    var v34 int64 = -892
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 uint16 = (^(v2))
        v37 := v4["foo bar"]
        fmt.Println(v2)
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v34
    _ = v4
}
