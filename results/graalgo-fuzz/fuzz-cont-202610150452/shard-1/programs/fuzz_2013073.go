package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS5 struct { F0 string; F1 int64 }
func genFn0(p0 uint32) int8 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: 2}
        v6 := GenS2{F0: 42, F1: 4294967295}
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    _ = p0
    return 42
}
func genFn1(p0 int64, p1 string, p2 int16) int8 {
    var v9 uint64 = 4294967295
    p2 = (p2 ^ p2)
    fmt.Println(p0, p2)
    if (1 != 1) {
        v9 = (v9 | (^(v9)))
        p2 += (-(p2))
        fmt.Println(p1, p0)
    }
    var v10 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 127
}

func main() {
    var v1 float32 = float32(3.14159)
    v2 := []string{(("a" + "foo bar") + ("" + ("foo bar" + "")))}
    var v3 string = ((("hello" + (((("foo bar" + (((("\n\t" + ("hello" + "hello")) + "\n\t") + (("a" + ("a" + "")) + ("hello" + (("" + ("\n\t" + "")) + "a")))) + "a")) + ("\n\t" + "")) + ("\n\t" + "\n\t")) + "\n\t")) + ("a" + ("hello" + "a"))) + (("a" + ("hello" + "foo bar")) + "a"))
    var v11 float64 = 0.1
    v11 += v11
    switch -2147483648 {
    case 1:
        v12 := GenS3{F0: 0}
        _ = v12
    case -2147483648:
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 float64 = ((float64(v1) - v11) * ((v11 / v11) + -2.25))
            if (100 <= -1) {
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 int8 = -1
                    _ = fi15
                    _ = v16
                }
            } else {
                v1 = 0.0
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi13
            _ = v14
        }
    case 458:
        fmt.Println(v1, v1)
    }
    var v17 uint64 = 1
    for fi18 := 0; fi18 < 20; fi18++ {
        if (!true) {
            switch "hello" {
            case "a":
                var v19 int8 = -128
                _ = v19
            case "hello":
                fmt.Println(v17)
            case "foo bar":
                v1 += (v1 / (((v1 * v1) - v1) / v1))
            default:
                if false {
                    fmt.Println(v17, fi18, v1)
                    switch ("a" + ("\n\t" + v3)) {
                    case "hello":
                        var v20 uint = 970
                        _ = v20
                    }
                    if len(v2) > 0 {
                        v2[0] = v3
                    }
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        }
        v21 := GenS4{F0: 42, F1: 996, F2: 1}
        _ = fi18
        _ = v21
    }
    v11 += v11
    switch v3 {
    case "hello":
        var v22 float64 = 3.14159
        _ = v22
    case "a":
        if false {
            var v23 uint16 = 1
            _ = v23
        }
    case "foo bar":
        switch "foo bar" {
        case "foo bar":
            if false {
                var v24 uint16 = 980
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 string = ("a" + "a")
                    var v27 float64 = (v11 - (-2.25 + v11))
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = v24
            } else {
                for fi28 := 0; fi28 < 9; fi28++ {
                    for fi29 := 0; fi29 < 18; fi29++ {
                        _ = len(v2)
                        v30 := GenS3{F0: 255}
                        if (770 != 65535) {
                            if true {
                                v31 := GenS5{F0: "hello", F1: -692}
                                fmt.Println(v11, v17, v17)
                                _ = v31
                            }
                        } else {
                            v17 += v17
                            switch ((v3 + v3) + (("\n\t" + v3) + "a")) {
                            case "hello":
                                v11 -= ((0.1 / ((-2.25 - 0.0) * v11)) - 0.1)
                            case "foo bar":
                                v32 := GenS5{F0: v3, F1: 982}
                                _ = v32
                            case "":
                                var v33 int16 = int16(v17)
                                _ = v33
                            }
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = fi28
                }
                fmt.Println(v17, v17, v1)
                v2 = append(v2, v3)
            }
        case "\n\t":
            v2 = append(v2, v3)
        default:
            fmt.Println(v11)
        }
    default:
        var v34 uint8 = 128
        _ = v34
    }
    fmt.Println(v11, v3)
    var v35 float32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v35
}
