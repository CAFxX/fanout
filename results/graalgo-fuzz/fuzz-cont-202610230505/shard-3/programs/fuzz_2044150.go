package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
func genFn0() uint {
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 bool = ((0.1 - 0.1) == 0.1)
            var v8 int32 = 2147483647
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    var v9 uint = 796
    v9 += v9
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 bool = (true && false)
        v12 := []int{}
        v11 = v11
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v9
    return (v9 << ((v9) & 63))
}

func main() {
    v1 := []string{"foo bar", "\n\t", ("\n\t" + ((("" + "\n\t") + ("a" + "hello")) + ((("foo bar" + "") + ("" + ("a" + ""))) + ("" + (("" + "foo bar") + ""))))), ("foo bar" + (("a" + "hello") + ""))}
    var v2 string = "\n\t"
    var v3 uint64 = 244
    var v4 bool = false
    switch ("\n\t" + v2) {
    case "":
        var v13 int16 = -32768
        _ = v13
    default:
        var v14 uint16 = 540
        _ = v14
    }
    v15 := GenS1{F0: 42, F1: 876}
    if v4 {
        switch v2 {
        case "":
            var v16 float64 = (-2.25 - (3.14159 / 3.14159))
            _ = v16
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v17 int32 = -2147483648
        _ = v17
    }
    v18 := GenS1{F0: 127, F1: 2147483648}
    fmt.Println(v4, v4)
    v19 := GenS2{F0: 0}
    fmt.Println(v2, v4, v4)
    v4 = true
    var v20 uint16 = 343
    v2 = ((v2 + "hello") + v2)
    v3 += 2
    if true {
        if true {
            v1 = append(v1, v2)
            for fi21 := 0; fi21 < 1; fi21++ {
                v22 := GenS3{F0: 713}
                _ = fi21
                _ = v22
            }
        } else {
            if v4 {
                var v23 int64 = -586
                switch v2 {
                case "\n\t":
                    for fi24 := 0; fi24 < 18; fi24++ {
                        if (!true) {
                            _ = len(v1)
                            v25 := GenS1{F0: 1, F1: 90}
                            fmt.Println(v20)
                            _ = v25
                        }
                        v20 = v20
                        _ = fi24
                    }
                case "hello":
                    var v26 uint = genFn0()
                    _ = v26
                case "":
                    v4 = (69 < 42)
                default:
                    var v27 uint8 = 2
                    _ = v27
                }
                v4 = false
                _ = v23
            } else {
                v20 = (v20 | v20)
                if v4 {
                    var v28 bool = (2 >= 42)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v29 string = v2
                    _ = v28
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 18; fi30++ {
                for fi31 := 0; fi31 < 13; fi31++ {
                    switch 9223372036854775807 {
                    case 0:
                        if v4 {
                            if (!(1 >= 896)) {
                                v4 = (v4 && (!true))
                                _ = len(v1)
                                v20 = (v20 / (((^(v20))) | 1))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch v2 {
                                case "a":
                                    if len(v1) > 0 {
                                        v1[0] = v2
                                    }
                                case "foo bar":
                                    _ = len(v1)
                                }
                            } else {
                                fmt.Println(v2, fi30)
                                var v32 uint16 = (v20 % ((v20) | 1))
                                _ = v32
                            }
                            _ = len(v1)
                            v33 := GenS3{F0: -2147483648}
                            v34 := GenS4{F0: 100}
                            _ = v33
                            _ = v34
                        }
                    case -9223372036854775808:
                        switch v2 {
                        case "\n\t":
                            _ = len(v1)
                        case "a":
                            var v35 int32 = -200
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                            v36 := GenS5{F0: 818}
                            _ = v35
                            _ = v36
                        case "":
                            var v37 int = fi31
                            _ = v37
                        default:
                            v4 = (((fi31 >> ((42) & 63)) > (fi31 % ((fi31) | 1))) && ((3.14159 < (3.14159 * float32(0.0))) || v4))
                        }
                    case 4294967295:
                        v3 = (v3 / ((2) | 1))
                    }
                    v38 := GenS1{F0: 0, F1: 2147483648}
                    _ = fi31
                    _ = v38
                }
                _ = fi30
            }
        }
        v2 = ("foo bar" + "a")
        v39 := GenS2{F0: -815}
        _ = v39
    }
    v40 := GenS2{F0: -814}
    v1 = append(v1, v2)
    var v41 bool = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
