package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint; F1 int64; F2 string }
type GenS3 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int32 = 1
    var v2 int32 = v1
    v3 := make(map[int]int)
    var v4 uint32 = 995
    v5 := GenS1{F0: 0, F1: (!false)}
    v6 := GenS1{F0: (v4 << ((163) & 63)), F1: ((v1 - 42) < (v1 << ((523) & 63)))}
    if false {
        var v7 uint = 42
        if true {
            v8 := GenS1{F0: (^(v4)), F1: ((!true) || (!false))}
            v9 := GenS1{F0: v4, F1: (!(!false))}
            var v10 float64 = -2.25
            _ = v10
            _ = v8
            _ = v9
        }
        v7 -= v7
        _ = v7
    }
    v2 += -781
    v11 := GenS1{F0: (v4 << ((619) & 63)), F1: false}
    var v12 uint16 = 32767
    v1 -= v1
    switch "hello" {
    case "foo bar":
        var v13 uint16 = v12
        _ = v13
    case "":
        var v14 int32 = int32(v4)
        _ = v14
    default:
        v15 := GenS1{F0: v4, F1: (!false)}
        _ = v15
    }
    if ((882 < 464) || false) {
        var v16 uint8 = 42
        _ = v16
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                v2 = (v2 & -480)
                _ = fi18
            }
            if (!true) {
                v19 := GenS1{F0: v4, F1: (!(1 == 1))}
                switch "hello" {
                case "hello":
                    var v20 int64 = -1
                    _ = v20
                }
                _ = v19
            } else {
                v21 := GenS1{F0: (v4 & v4), F1: true}
                var v22 int32 = -418
                v12 = (v12 / (((v12 % ((uint16(v4)) | 1))) | 1))
                _ = v21
                _ = v22
            }
            _ = fi17
        }
        fmt.Println(v4)
        switch (("foo bar" + "foo bar") + "foo bar") {
        case "hello":
            if true {
                var v23 string = (((("a" + ("" + ("" + "\n\t"))) + "foo bar") + "hello") + "\n\t")
                _ = v23
            } else {
                v12 = v12
                switch -1 {
                case 2147483647:
                    v12 -= uint16(v4)
                case 554:
                    v24 := []int{127, 2147483647, 2147483647}
                    _ = v24
                case 4294967295:
                    v25 := GenS2{F0: uint(v12), F1: 0, F2: (("hello" + "a") + "foo bar")}
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 20; fi26++ {
                        if true {
                            switch fi26 {
                            case -1:
                                var v27 int64 = 9223372036854775807
                                _ = v27
                            case -2147483648:
                                fmt.Println(v4, v12)
                            default:
                                var v28 int32 = v1
                                _ = v28
                            }
                        } else {
                            v29 := GenS3{F0: ((-(fi26)) < int(v12)), F1: v4}
                            _ = v29
                        }
                        v30 := GenS2{F0: 1, F1: 456, F2: (("foo bar" + (("" + "foo bar") + ((("hello" + "") + "\n\t") + ("\n\t" + (("" + ((("a" + ("" + (("a" + "\n\t") + ("foo bar" + "hello")))) + "\n\t") + ("a" + "\n\t"))) + "foo bar"))))) + (("\n\t" + "\n\t") + "hello"))}
                        if (819 == -481) {
                            v1 = (v1 * 0)
                        }
                        _ = fi26
                        _ = v30
                    }
                }
            }
        case "":
            var v31 uint64 = 42
            _ = v31
        case "\n\t":
            v32 := []int{}
            _ = v32
        default:
            if (32768 >= 771) {
                var v33 uint16 = v12
                _ = v33
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
