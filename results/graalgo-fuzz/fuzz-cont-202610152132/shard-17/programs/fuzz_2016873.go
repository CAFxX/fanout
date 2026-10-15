package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 string; F2 string }
func genFn0(p0 int32, p1 int) string {
    v5 := []int{}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint32 = 557
        if true {
            var v8 uint32 = (v7 ^ 1)
            if true {
                fmt.Println(v8)
            }
            for fi9 := 0; fi9 < 16; fi9++ {
                switch fi6 {
                case -657:
                    var v10 uint32 = v8
                    _ = v10
                }
                _ = fi9
            }
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    v5 = append(v5, (p1 & p1))
    var v11 int32 = (p0 + 635)
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return ""
}

func main() {
    var v1 int8 = 1
    var v2 int32 = 0
    v3 := []string{"\n\t", "\n\t", (((("a" + "") + "hello") + "") + "a")}
    v4 := []uint32{442}
    v12 := GenS1{F0: (255 > 538)}
    v2 = v2
    var v13 int64 = 0
    v14 := GenS1{F0: (0 == 58)}
    v2 = (v2 * v2)
    v2 -= -1
    _ = len(v4)
    v2 = v2
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int = (fi15 ^ fi15)
        _ = fi15
        _ = v16
    }
    switch 128 {
    case 9223372036854775807:
        v17 := GenS1{F0: true}
        _ = v17
    case -825:
        v2 += (v2 << ((2) & 63))
    case -1:
        if false {
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 uint16 = 42
                var v20 float64 = (-2.25 * 1.5)
                fmt.Println(v19, fi18, v1)
                _ = fi18
                _ = v19
                _ = v20
            }
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi21 := 0; fi21 < 4; fi21++ {
                    switch ("foo bar" + "a") {
                    case "hello":
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case "\n\t":
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                    default:
                        fmt.Println(v1)
                    }
                    switch "a" {
                    case "\n\t":
                        var v22 int32 = (v2 % (((-(v2))) | 1))
                        _ = v22
                    case "":
                        var v23 uint16 = 42
                        _ = v23
                    }
                    switch ("\n\t" + ("\n\t" + "hello")) {
                    case "hello":
                        var v24 uint = 1
                        _ = v24
                    }
                    _ = fi21
                }
            } else {
                v2 = (v2 % (((-(v2))) | 1))
                var v25 string = ((("" + ("foo bar" + "a")) + "hello") + "hello")
                v1 -= (v1 & v1)
                _ = v25
            }
            switch 128 {
            case -2147483648:
                v1 = (v1 | v1)
            case 2147483647:
                v26 := GenS1{F0: (true && true)}
                _ = v26
            case 128:
                if len(v4) > 0 {
                    v4[0] = 811
                }
            }
        } else {
            var v27 int64 = (v13 << ((464) & 63))
            v28 := GenS2{F0: 956, F1: "\n\t", F2: (("foo bar" + "") + "")}
            if false {
                v2 = -817
            } else {
                v3 = append(v3, ("foo bar" + ("foo bar" + ("\n\t" + "a"))))
            }
            _ = v27
            _ = v28
        }
    }
    fmt.Println(v1, v2)
    v13 += v13
    var v29 int64 = 9223372036854775807
    fmt.Println(v13, v29)
    for fi30 := 0; fi30 < 13; fi30++ {
        var v31 uint8 = 241
        _ = fi30
        _ = v31
    }
    v13 = 9223372036854775807
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 string = "hello"
        v34 := GenS2{F0: 0, F1: v33, F2: (v33 + "\n\t")}
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println(v13, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
