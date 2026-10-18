package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8, p1 uint, p2 bool) string {
    if false {
        for fi5 := 0; fi5 < 3; fi5++ {
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := GenS1{F0: fi6, F1: 42}
                fmt.Println(p2, p2)
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 bool = (!false)
                switch "a" {
                case "":
                    fmt.Println(p2)
                case "foo bar":
                    if false {
                        var v11 float32 = float32(0.1)
                        switch (fi9 & -1) {
                        case 32768:
                            v12 := []int{-2147483648, 9223372036854775807, (^(fi9))}
                            _ = v12
                        case -9223372036854775808:
                            p2 = (!((false || v10) && true))
                        case -2147483648:
                            switch "" {
                            case "":
                                if false {
                                    v13 := []int{}
                                    _ = v13
                                } else {
                                    fmt.Println(fi8, v10, p1)
                                }
                            case "foo bar":
                                v14 := []int{83, (fi8 | fi9), (fi8 & 32767), (fi8 >> (((p1 ^ p1)) & 63))}
                                _ = v14
                            case "hello":
                                v15 := []int{-614}
                                _ = v15
                            }
                            p0 = (p0 | p0)
                        }
                        v10 = (v11 < 1.5)
                        _ = v11
                    } else {
                        v16 := []int{fi8, fi8, fi9}
                        fmt.Println(fi9)
                        p2 = p2
                        _ = v16
                    }
                default:
                    var v17 int32 = -2147483648
                    _ = v17
                }
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    }
    if (false && (!true)) {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: fi18, F1: 32767}
            p0 -= 100
            v20 := GenS1{F0: (fi18 * fi18), F1: 191}
            _ = fi18
            _ = v19
            _ = v20
        }
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int32 = 42
    v3 := []uint32{v1, v1}
    var v4 float32 = -2.25
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 bool = (!false)
        if true {
            v23 := GenS2{F0: 18446744073709551615, F1: (v1 + (^(v1)))}
            for fi24 := 0; fi24 < 5; fi24++ {
                v4 += ((-2.25 * float32(1.5)) - float32((3.14159 * float64(1e10))))
                v1 -= 897
                if v22 {
                    v1 = (v1 & (v1 >> ((42) & 63)))
                }
                _ = fi24
            }
            _ = v23
        }
        v25 := GenS1{F0: 128, F1: 823}
        _ = fi21
        _ = v22
        _ = v25
    }
    v2 = v2
    var v26 string = (("" + "a") + "foo bar")
    switch v26 {
    case "a":
        var v27 int64 = 9223372036854775807
        _ = v27
    case "\n\t":
        if false {
            v2 -= -1
            if ((false || (1e10 == (-2.25 * (3.14159 + 0.1)))) || (!true)) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                if false {
                    for fi28 := 0; fi28 < 8; fi28++ {
                        switch -529 {
                        case -141:
                            for fi29 := 0; fi29 < 17; fi29++ {
                                var v30 int = (fi29 << ((1) & 63))
                                _ = fi29
                                _ = v30
                            }
                        }
                        v26 = v26
                        _ = fi28
                    }
                }
            }
            for fi31 := 0; fi31 < 2; fi31++ {
                if false {
                    var v32 string = ("a" + ("hello" + v26))
                    _ = v32
                }
                _ = fi31
            }
        }
    }
    var v33 float64 = 0.0
    switch (v26 + v26) {
    case "a":
        if true {
            var v34 int64 = 1
            _ = v34
        }
    case "foo bar":
        var v35 float64 = (-2.25 - v33)
        _ = v35
    case "\n\t":
        v2 = -1
    default:
        _ = len(v3)
    }
    if true {
        fmt.Println(v4, v26)
    }
    v36 := GenS3{F0: (v1 ^ 2147483647)}
    fmt.Println(v4, v4, v1)
    var v37 uint32 = (v1 & 880)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
}
