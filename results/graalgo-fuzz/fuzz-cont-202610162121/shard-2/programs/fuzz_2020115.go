package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 uint16 }

func main() {
    var v1 uint16 = 0
    v2 := []string{}
    v3 := make(map[uint32]string)
    v4 := []int{-9223372036854775808, 42, 1, 4294967295}
    switch "hello" {
    case "":
        v1 = (v1 % ((504) | 1))
    case "foo bar":
        switch 32767 {
        case 42:
            v1 = v1
        case 707:
            v5 := v3[2]
            _ = v5
        }
    case "hello":
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 2; fi7++ {
                if false {
                    var v8 uint32 = 42
                    v4 = append(v4, (-(fi6)))
                    _ = v8
                } else {
                    v1 -= (v1 / (((^(v1))) | 1))
                    v9 := GenS1{F0: (v1 * (v1 + (^(v1)))), F1: 0}
                    _ = v9
                }
                switch (fi7 / (((^(fi7))) | 1)) {
                case 9223372036854775807:
                    switch fi6 {
                    case -487:
                        var v10 uint64 = 0
                        _ = v10
                    default:
                        var v11 uint = 255
                        _ = v11
                    }
                case 382:
                    _ = len(v4)
                default:
                    v1 = v1
                    var v12 uint = uint(v1)
                    _ = v12
                }
                v1 += v1
                _ = fi7
            }
            v1 = v1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi6
        }
    default:
        if len(v2) > 0 {
            v2[0] = ("\n\t" + "a")
        }
    }
    var v13 uint8 = 255
    v14 := GenS2{F0: "", F1: (v1 % (((v1 << ((0) & 63))) | 1))}
    fmt.Println(v13, v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        fmt.Println(v1, v1)
        fmt.Println(v1, v13)
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v15 string = ("hello" + (("" + ("" + "\n\t")) + ""))
            switch v15 {
            case "hello":
                v16 := v3[4294967295]
                _ = v16
            case "a":
                v13 -= v13
            case "":
                fmt.Println(v15)
            }
            _ = v15
        } else {
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 uint32 = 0
                v1 -= v1
                fmt.Println(fi17, v1, fi17)
                _ = fi17
                _ = v18
            }
        }
    }
    var v19 bool = ((-9223372036854775808 < 0) && (-1 == 42))
    if (2 >= 1) {
        fmt.Println(v13, v13, v19)
        v19 = (false || (v13 != 128))
    }
    var v20 int8 = 42
    fmt.Println(v1, v1, v19)
    var v21 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
