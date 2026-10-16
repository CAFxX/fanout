package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int16; F2 int }
func genFn0() int16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int64 = 501
        v6 -= 1
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: "", F1: 100}
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -1
    v3 := make(map[uint32]string)
    var v4 uint8 = 255
    var v8 int32 = 244
    if (genFn0() < 1000) {
        for fi9 := 0; fi9 < 7; fi9++ {
            _ = len(v1)
            _ = fi9
        }
        if ((v2 << ((18446744073709551615) & 63)) == v2) {
            var v10 string = ("hello" + ("hello" + "foo bar"))
            _ = v10
        } else {
            switch ("hello" + "a") {
            case "\n\t":
                switch 128 {
                case 2147483647:
                    _ = len(v3)
                case 0:
                    if (!true) {
                        var v11 uint16 = 0
                        _ = v11
                    }
                case 9223372036854775807:
                    for fi12 := 0; fi12 < 11; fi12++ {
                        v13 := []int{42, (fi12 & fi12), (fi12 * (fi12 << ((42) & 63))), fi12}
                        if (!true) {
                            v2 += 42
                            v13 = append(v13, (fi12 - (^(fi12))))
                            switch 32767 {
                            case 1:
                                var v14 float32 = (-2.25 + 3.14159)
                                _ = v14
                            case 128:
                                var v15 int32 = -141
                                _ = v15
                            case -1000:
                                switch "\n\t" {
                                case "hello":
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                case "":
                                    v16 := GenS2{F0: -9223372036854775808, F1: 2147483647}
                                    _ = v16
                                default:
                                    _ = len(v13)
                                }
                            default:
                                for fi17 := 0; fi17 < 11; fi17++ {
                                    if len(v13) > 0 {
                                        v13[0] = (fi12 << ((65535) & 63))
                                    }
                                    _ = fi17
                                }
                            }
                        }
                        _ = fi12
                        _ = v13
                    }
                }
            case "a":
                fmt.Println(v4, v8)
            case "foo bar":
                v18 := []int{-1, -881}
                _ = v18
            }
        }
    }
    v19 := GenS1{F0: "foo bar", F1: 117}
    fmt.Println(v4, v8, v4)
    delete(v3, 2)
    for fi20 := 0; fi20 < 1; fi20++ {
        if false {
            v21 := GenS1{F0: "", F1: -76}
            _ = v21
        } else {
            for fi22 := 0; fi22 < 10; fi22++ {
                v23 := GenS3{F0: -994}
                switch fi22 {
                case 4294967295:
                    v2 = (v2 | (v2 ^ v2))
                case 9223372036854775807:
                    fmt.Println(fi20, v8, v8)
                case -1:
                    var v24 bool = false
                    _ = v24
                default:
                    fmt.Println(v4, v4, fi22)
                }
                _ = fi22
                _ = v23
            }
            var v25 int64 = -167
            var v26 string = ("" + "\n\t")
            _ = v25
            _ = v26
        }
        v8 -= (v8 ^ 0)
        var v27 uint64 = 0
        _ = fi20
        _ = v27
    }
    if ((3.14159 + (1.5 - float32(1.5))) <= 1.5) {
        var v28 int16 = genFn0()
        if false {
            if (false || (-72 >= 42)) {
                var v29 uint = 18446744073709551615
                _ = v29
            }
            var v30 uint8 = 128
            _ = v30
        }
        var v31 int8 = (v2 % (((v2 / ((v2) | 1))) | 1))
        _ = v28
        _ = v31
    } else {
        for fi32 := 0; fi32 < 5; fi32++ {
            var v33 int16 = -232
            for fi34 := 0; fi34 < 19; fi34++ {
                switch -1 {
                case 0:
                    v2 = v2
                default:
                    fmt.Println(v33)
                }
                v4 += v4
                if true {
                    var v35 float64 = float64(1.5)
                    _ = v35
                } else {
                    var v36 int = fi32
                    _ = v36
                }
                _ = fi34
            }
            fmt.Println(v8)
            _ = fi32
            _ = v33
        }
        for fi37 := 0; fi37 < 8; fi37++ {
            var v38 uint = 255
            if (37 >= 2) {
                v39 := GenS4{F0: ("\n\t" + (("hello" + "foo bar") + ("hello" + "hello"))), F1: 532, F2: 9223372036854775807}
                delete(v3, 65535)
                var v40 float64 = 1.5
                _ = v39
                _ = v40
            }
            v4 += (v4 % ((v4) | 1))
            _ = fi37
            _ = v38
        }
        var v41 int64 = 1
        _ = v41
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v2)
        default:
            v4 = (^(v4))
        }
        v2 = (v2 << ((1) & 63))
        var v43 string = ((("\n\t" + ("a" + (("" + (("" + "foo bar") + "")) + ""))) + "foo bar") + "foo bar")
        _ = fi42
        _ = v43
    }
    v8 = v8
    v8 += v8
    switch -2147483648 {
    case -2147483648:
        var v44 uint32 = 631
        _ = v44
    default:
        v2 = int8(v4)
    }
    switch -2147483648 {
    case 127:
        var v45 int = 4294967295
        _ = v45
    }
    v46 := []int{32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v8
}
