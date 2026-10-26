package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0(p0 int8, p1 string, p2 int) int16 {
    p2 -= 128
    var v7 uint8 = 127
    p1 = (p1 + (p1 + (((p1 + p1) + (p1 + p1)) + ("foo bar" + (p1 + p1)))))
    v7 = 255
    var v8 int64 = 522
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1000
}

func main() {
    v1 := []string{("hello" + "\n\t"), (("a" + "\n\t") + (("\n\t" + ("\n\t" + "foo bar")) + ((("" + "a") + ("foo bar" + ((("" + ("\n\t" + "a")) + "a") + "a"))) + "foo bar"))), "hello", "\n\t"}
    var v2 uint16 = 817
    var v3 uint = 2
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    var v6 uint8 = 255
    for fi9 := 0; fi9 < 11; fi9++ {
        switch 282 {
        case 1:
            fmt.Println(v2, fi9, v6)
        case -2147483648:
            v4[(fi9 % (((fi9 / (((fi9 % (((fi9 + (fi9 ^ fi9))) | 1))) | 1))) | 1))] = ((32767 >= 0) || ((168 >= v6) || ("hello" > "")))
        case 127:
            var v10 int = int(v2)
            _ = v10
        default:
            v11 := GenS1{F0: "\n\t"}
            _ = v11
        }
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            v12 := GenS1{F0: "hello"}
            _ = v12
        }
        v13 := GenS2{F0: 42}
        _ = fi9
        _ = v13
    }
    fmt.Println(v3, v3)
    fmt.Println(v3, v2)
    if (2147483648 == 2) {
        if true {
            v6 -= v6
        }
        _ = len(v1)
        v2 -= (^(v2))
    } else {
        v14 := GenS1{F0: "a"}
        if true {
            for fi15 := 0; fi15 < 3; fi15++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch (^(fi15)) {
                case 64:
                    v16 := GenS2{F0: 32}
                    _ = v16
                default:
                    if (true && (!false)) {
                        for fi17 := 0; fi17 < 15; fi17++ {
                            fmt.Println(v6, v3)
                            fmt.Println(v3)
                            v18 := v4[-446]
                            _ = fi17
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 9; fi19++ {
                            v20 := GenS3{F0: uint(v2), F1: 2}
                            _ = fi19
                            _ = v20
                        }
                        if false {
                            if (!true) {
                                if true {
                                    if (1 != 42) {
                                        if true {
                                            for fi21 := 0; fi21 < 18; fi21++ {
                                                v22 := v5[2147483647]
                                                _ = fi21
                                                _ = v22
                                            }
                                        }
                                        v2 -= 2
                                        var v23 uint64 = 42
                                        _ = v23
                                    }
                                }
                                v3 = (v3 % (((^(v3))) | 1))
                                v4[1] = (true && (v2 <= (v2 | (v2 * (v2 * v2)))))
                            }
                        } else {
                            v6 = (v6 * (v6 & (v6 & 247)))
                            fmt.Println(v6)
                            var v24 int16 = 1
                            _ = v24
                        }
                    } else {
                        v25 := GenS4{F0: v3, F1: 592, F2: 0}
                        fmt.Println(v3)
                        var v26 int = fi15
                        _ = v25
                        _ = v26
                    }
                }
                _ = fi15
            }
            if (2147483647 > 2147483647) {
                if true {
                    switch "a" {
                    case "":
                        v2 = (^(v2))
                    default:
                        var v27 uint = (v3 | uint(v6))
                        _ = v27
                    }
                    if len(v1) > 0 {
                        v1[0] = ("" + "foo bar")
                    }
                    var v28 uint8 = v6
                    _ = v28
                }
            }
        } else {
            fmt.Println(v3, v2)
            var v29 float64 = (1.5 + 3.14159)
            _ = v29
        }
        var v30 int16 = 1000
        _ = v14
        _ = v30
    }
    var v31 bool = ((((-2.25 - (1e10 * (1e10 * (0.1 - 1e10)))) > (1e10 * (1.5 - 0.0))) && (113 == (v6 % (((v6 ^ v6)) | 1)))) || false)
    v1 = append(v1, "foo bar")
    var v32 int8 = -1
    var v33 bool = v31
    fmt.Println(v32, v33, v6)
    var v34 uint64 = 1
    if false {
        _ = len(v4)
    }
    if v31 {
        for fi35 := 0; fi35 < 5; fi35++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            _ = fi35
        }
    } else {
        var v36 int32 = 0
        _ = v36
    }
    v34 = v34
    _ = len(v4)
    v37 := v4[-929]
    v34 = v34
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
