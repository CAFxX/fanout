package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 int32, p1 uint16) uint8 {
    v7 := GenS1{F0: (4294967295 != 4294967295), F1: 1000}
    var v8 int16 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 75
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 1
    v3 := make(map[string]string)
    v4 := []string{"\n\t", "foo bar"}
    v5 := make(map[string]string)
    var v6 float64 = 3.14159
    var v9 float64 = (v6 * 0.1)
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v10 string = "\n\t"
        _ = v10
    case -9223372036854775808:
        v11 := GenS1{F0: false, F1: -384}
        _ = v11
    case 32768:
        v12 := GenS1{F0: (((0.0 - float32(v9)) * float32((v6 + float64((((-2.25 - 1e10) + 3.14159) - 3.14159))))) >= 0.1), F1: 1}
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 string = "a"
        _ = fi13
        _ = v14
    }
    switch int(v2) {
    case 42:
        var v15 uint16 = 487
        _ = v15
    default:
        for fi16 := 0; fi16 < 1; fi16++ {
            v9 += (float64(1.5) / ((v9 - float64(1.5)) - (v6 * v6)))
            if true {
                fmt.Println(v2)
                v9 += v6
            }
            v3["foo bar"] = "a"
            _ = fi16
        }
    }
    if ((0 != 0) && true) {
        _ = len(v1)
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v4 = append(v4, "a")
        v2 = v2
        _ = fi17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v18 int16 = 32767
    for fi19 := 0; fi19 < 16; fi19++ {
        if false {
            if ("a" == (("hello" + "a") + "a")) {
                fmt.Println(v18)
            } else {
                v4 = append(v4, "foo bar")
                var v20 uint16 = 2
                switch "" {
                case "hello":
                    for fi21 := 0; fi21 < 20; fi21++ {
                        if (!(!(true && (!((791 < uint(v20)) || false))))) {
                            delete(v1, fi21)
                            fmt.Println(fi19)
                        }
                        switch ("" + "") {
                        case "":
                            fmt.Println(v20, v20, v2)
                            var v22 string = ((("foo bar" + "\n\t") + ("\n\t" + "a")) + ("\n\t" + "a"))
                            v2 += (v2 % (((v2 & v2)) | 1))
                            _ = v22
                        }
                        fmt.Println(v9)
                        _ = fi21
                    }
                case "foo bar":
                    if true {
                        var v23 uint = 4294967295
                        _ = v23
                    }
                case "\n\t":
                    var v24 bool = true
                    _ = v24
                }
                _ = v20
            }
            for fi25 := 0; fi25 < 19; fi25++ {
                fmt.Println(v18)
                _ = fi25
            }
        }
        fmt.Println(v9, v9)
        _ = fi19
    }
    v26 := GenS2{F0: 946, F1: 2147483647}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        var v27 uint8 = genFn0(-1, 378)
        for fi28 := 0; fi28 < 9; fi28++ {
            for fi29 := 0; fi29 < 8; fi29++ {
                var v30 bool = ((("foo bar" < "a") && true) && ((-(fi28)) < (fi29 ^ -2147483648)))
                v5["foo bar"] = "hello"
                if (v30 || v30) {
                    if (((v2 % ((v2) | 1)) != -1) || v30) {
                        v4 = append(v4, "")
                        if false {
                            v18 += (v18 % (((v18 - (v18 >> ((42) & 63)))) | 1))
                            var v31 int = 0
                            _ = v31
                        } else {
                            _ = len(v3)
                        }
                        if len(v4) > 0 {
                            v4[0] = (((("\n\t" + ("\n\t" + "hello")) + "foo bar") + ("\n\t" + ("foo bar" + ""))) + "foo bar")
                        }
                    }
                    v32 := GenS2{F0: (fi28 | fi29), F1: 168}
                    switch ((("foo bar" + (((("hello" + "hello") + "\n\t") + "hello") + "foo bar")) + ("foo bar" + "hello")) + "a") {
                    case "hello":
                        v18 = (v18 + v18)
                    }
                    _ = v32
                } else {
                    v18 -= (v18 | -385)
                }
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
        var v33 uint8 = (v27 / ((v27) | 1))
        _ = v27
        _ = v33
    }
    v34 := GenS1{F0: false, F1: (-(v18))}
    for fi35 := 0; fi35 < 19; fi35++ {
        v9 += (0.0 + (0.0 * (-2.25 / v6)))
        var v36 int64 = v2
        delete(v3, "a")
        _ = fi35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
