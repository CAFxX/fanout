package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int8; F1 int8 }
func genFn0(p0 uint16, p1 int32, p2 uint8) uint32 {
    v4 := []int{128, 383, 32768, -341}
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int16 = -32768
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: "a"}
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 0
}

func main() {
    var v1 bool = false
    var v2 uint8 = 209
    v3 := []int{652, 4294967295, 2147483647}
    if v1 {
        v8 := GenS1{F0: "foo bar"}
        if (42 >= 535) {
            for fi9 := 0; fi9 < 14; fi9++ {
                v1 = true
                _ = fi9
            }
            v2 = v2
        }
        v1 = true
        _ = v8
    }
    v2 += (v2 >> ((0) & 63))
    switch ("hello" + "\n\t") {
    case "":
        if len(v3) > 0 {
            v3[0] = -66
        }
    case "a":
        v10 := GenS1{F0: "hello"}
        _ = v10
    case "foo bar":
        for fi11 := 0; fi11 < 17; fi11++ {
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(v2, v1)
                v1 = (!true)
                _ = fi12
            }
            var v13 uint8 = v2
            v2 = v13
            _ = fi11
            _ = v13
        }
    default:
        if false {
            switch -2147483648 {
            case -9223372036854775808:
                if ((0.1 * 0.1) != 1e10) {
                    switch "foo bar" {
                    case "hello":
                        v1 = (!(395 > 4294967295))
                    case "a":
                        v1 = v1
                    case "foo bar":
                        switch -488 {
                        case -9223372036854775808:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 42:
                            _ = len(v3)
                        default:
                            v2 = (v2 & v2)
                        }
                    }
                    _ = len(v3)
                    var v14 string = ("a" + "foo bar")
                    _ = v14
                }
            case 1:
                v1 = v1
            default:
                if (1 < 0) {
                    var v15 int16 = 32767
                    v16 := GenS2{F0: v1}
                    for fi17 := 0; fi17 < 15; fi17++ {
                        fmt.Println(v1, v2)
                        var v18 int16 = v15
                        v19 := GenS1{F0: "a"}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = v15
                    _ = v16
                }
            }
        }
    }
    var v20 uint64 = 547
    v20 += 0
    v21 := GenS3{F0: 1, F1: 407}
    v1 = (!true)
    if v1 {
        if v1 {
            fmt.Println(v20, v2)
            for fi22 := 0; fi22 < 2; fi22++ {
                if len(v3) > 0 {
                    v3[0] = fi22
                }
                _ = fi22
            }
        }
    }
    v23 := GenS4{F0: 62, F1: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
}
