package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
func genFn0() int16 {
    v4 := []int{127, 127, 440, 4294967295}
    var v5 uint64 = 951
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 int = -2147483648
    v2 := []uint32{65535}
    var v3 float32 = 1.5
    for fi6 := 0; fi6 < 16; fi6++ {
        v1 = (-(v1))
        _ = fi6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        if len(v2) > 0 {
            v2[0] = 2
        }
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(v3)
            var v9 string = "a"
            v10 := GenS1{F0: (fi8 / ((fi7) | 1)), F1: 42, F2: -1}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v11 uint16 = 2
    v3 = v3
    var v12 uint = 2
    var v13 uint16 = (v11 & (v11 << ((2) & 63)))
    var v14 int64 = 0
    var v15 uint16 = v11
    for fi16 := 0; fi16 < 13; fi16++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch ("\n\t" + ("hello" + "")) {
        case "hello":
            fmt.Println(v1, v3)
        }
        v15 = (^(v11))
        _ = fi16
    }
    v11 = (v15 - 32767)
    switch v1 {
    case 316:
        _ = len(v2)
    case -758:
        switch (("\n\t" + "") + ("a" + "")) {
        case "hello":
            for fi17 := 0; fi17 < 18; fi17++ {
                for fi18 := 0; fi18 < 5; fi18++ {
                    fmt.Println(fi18, v13)
                    v11 = v15
                    _ = fi18
                }
                if ((!("\n\t" <= "a")) || ("foo bar" >= "\n\t")) {
                    if ((^(fi17)) > v1) {
                        switch ("\n\t" + "hello") {
                        case "a":
                            for fi19 := 0; fi19 < 20; fi19++ {
                                if ("a" >= (("a" + "\n\t") + (("a" + "a") + ""))) {
                                    if false {
                                        var v20 uint = v12
                                        v21 := GenS1{F0: (v1 * 32768), F1: 2, F2: -2147483648}
                                        var v22 float32 = (v3 * v3)
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    }
                                    var v23 int64 = (v14 ^ (v14 % ((0) | 1)))
                                    v13 = v11
                                    _ = v23
                                } else {
                                    var v24 uint32 = 2147483647
                                    fmt.Println(v11, v15, v1)
                                    if false {
                                        v14 += v14
                                        v3 = v3
                                    } else {
                                        var v25 int16 = 1
                                        var v26 string = ((((("a" + "foo bar") + "") + "a") + "foo bar") + "foo bar")
                                        fmt.Println(v14, v14, v1)
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = v24
                                }
                                v27 := GenS2{F0: -1}
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                                _ = fi19
                                _ = v27
                            }
                        }
                        for fi28 := 0; fi28 < 15; fi28++ {
                            switch 1 {
                            case 4294967295:
                                v29 := GenS1{F0: 178, F1: uint64(v15), F2: -401}
                                _ = v29
                            default:
                                var v30 int64 = -68
                                _ = v30
                            }
                            _ = len(v2)
                            var v31 int32 = 0
                            _ = fi28
                            _ = v31
                        }
                        v32 := GenS3{F0: 1000}
                        _ = v32
                    }
                    v2 = append(v2, 1)
                    v15 -= (v13 | v13)
                } else {
                    fmt.Println(v3)
                    v15 = 65535
                }
                v14 = (v14 % ((v14) | 1))
                _ = fi17
            }
        case "foo bar":
            for fi33 := 0; fi33 < 8; fi33++ {
                for fi34 := 0; fi34 < 3; fi34++ {
                    _ = len(v2)
                    _ = fi34
                }
                _ = fi33
            }
        case "":
            if (v14 < (v14 >> ((2) & 63))) {
                _ = len(v2)
                _ = len(v2)
            } else {
                if false {
                    for fi35 := 0; fi35 < 11; fi35++ {
                        for fi36 := 0; fi36 < 13; fi36++ {
                            _ = len(v2)
                            _ = fi36
                        }
                        v14 -= v14
                        var v37 int8 = 31
                        _ = fi35
                        _ = v37
                    }
                    if true {
                        fmt.Println(v11)
                        v11 += v13
                    } else {
                        for fi38 := 0; fi38 < 9; fi38++ {
                            var v39 int16 = -32768
                            v15 = v11
                            _ = fi38
                            _ = v39
                        }
                        v3 += v3
                    }
                } else {
                    var v40 int32 = -396
                    _ = v40
                }
                switch (-(v1)) {
                case -843:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v41 float32 = v3
                    _ = v41
                }
            }
        }
    }
    fmt.Println(v1)
    var v42 float64 = 1.5
    v2 = append(v2, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v42
}
