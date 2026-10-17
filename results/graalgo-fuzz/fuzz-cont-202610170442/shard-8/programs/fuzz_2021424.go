package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 string; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int8; F1 uint }
func genFn0(p0 uint, p1 int, p2 int) uint8 {
    v5 := GenS1{F0: 42}
    fmt.Println(p0, p2, p0)
    var v6 int8 = -128
    var v7 uint32 = 2
    var v8 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 173
}

func main() {
    v1 := []string{(("\n\t" + ("foo bar" + ((("" + "\n\t") + "a") + ((("foo bar" + "\n\t") + ("" + (("hello" + "foo bar") + (("\n\t" + ("\n\t" + ("hello" + ""))) + "a")))) + "")))) + "hello"), "", ("hello" + (((("\n\t" + ("foo bar" + ("\n\t" + "hello"))) + "\n\t") + "\n\t") + (("a" + "a") + "a")))}
    var v2 int8 = 127
    var v3 uint64 = 368
    var v4 float64 = (3.14159 * 1e10)
    var v9 string = ("foo bar" + ("\n\t" + "foo bar"))
    var v10 uint64 = (v3 >> ((4294967295) & 63))
    v4 = ((v4 / 1.0) - v4)
    switch v9 {
    case "\n\t":
        if false {
            for fi11 := 0; fi11 < 12; fi11++ {
                fmt.Println(v4, v4, fi11)
                var v12 bool = (true && (((v2 | 92) == 42) || true))
                if len(v1) > 0 {
                    v1[0] = v9
                }
                _ = fi11
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v14 int16 = 1
                _ = fi13
                _ = v14
            }
            if true {
                switch "foo bar" {
                case "":
                    v1 = append(v1, ((v9 + v9) + ((v9 + v9) + (v9 + v9))))
                case "hello":
                    if false {
                        switch (v9 + v9) {
                        case "a":
                            if (69 <= 42) {
                                for fi15 := 0; fi15 < 13; fi15++ {
                                    if (-2.25 <= (v4 * (1.5 * 1.5))) {
                                        v4 += v4
                                        if (v4 <= ((v4 - v4) / v4)) {
                                            v1 = append(v1, v9)
                                        }
                                        switch 32768 {
                                        case 1:
                                            switch (fi15 * fi15) {
                                            case 9223372036854775807:
                                                for fi16 := 0; fi16 < 17; fi16++ {
                                                    v4 = float64(3.14159)
                                                    _ = fi16
                                                }
                                            }
                                        case 127:
                                            v9 = (v9 + v9)
                                        case 970:
                                            for fi17 := 0; fi17 < 13; fi17++ {
                                                var v18 string = ""
                                                _ = fi17
                                                _ = v18
                                            }
                                        default:
                                            for fi19 := 0; fi19 < 4; fi19++ {
                                                fmt.Println(v4, v4, v10)
                                                _ = fi19
                                            }
                                        }
                                    } else {
                                        fmt.Println(v2)
                                        v9 = v9
                                    }
                                    v20 := GenS2{F0: (v9 + (v9 + ("hello" + ("hello" + ((("a" + (("a" + (("foo bar" + (v9 + "hello")) + v9)) + v9)) + "\n\t") + v9))))), F1: v9, F2: uint8(v3)}
                                    _ = fi15
                                    _ = v20
                                }
                                v9 = "a"
                                fmt.Println(v9, v9, v9)
                            }
                        default:
                            if len(v1) > 0 {
                                v1[0] = v9
                            }
                        }
                        var v21 int = 42
                        _ = v21
                    } else {
                        v2 = v2
                        if true {
                            var v22 uint = 18446744073709551615
                            _ = v22
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v9 = v9
                        }
                        switch v9 {
                        case "hello":
                            v4 += 1.5
                        case "\n\t":
                            if len(v1) > 0 {
                                v1[0] = (v9 + ("" + "foo bar"))
                            }
                        }
                    }
                case "\n\t":
                    var v23 uint32 = 0
                    _ = v23
                }
            } else {
                var v24 string = ((v9 + ("\n\t" + v9)) + "")
                v3 = v10
                _ = v24
            }
        }
    case "":
        var v25 int8 = (v2 / (((v2 >> ((4294967295) & 63))) | 1))
        _ = v25
    case "foo bar":
        for fi26 := 0; fi26 < 9; fi26++ {
            v1 = append(v1, ((v9 + (v9 + "a")) + (v9 + v9)))
            _ = fi26
        }
    }
    v3 += 1
    var v27 float64 = v4
    v3 = (v10 - v3)
    v2 += -1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v28 bool = true
    var v29 uint64 = (v3 >> ((0) & 63))
    if v28 {
        var v30 bool = (!false)
        switch "a" {
        case "hello":
            var v31 int64 = 0
            _ = v31
        case "":
            v32 := GenS3{F0: 255}
            _ = v32
        default:
            fmt.Println(v30, v29)
        }
        _ = v30
    }
    v33 := GenS4{F0: 2147483647, F1: v10, F2: 127}
    switch -9223372036854775808 {
    case 32767:
        fmt.Println(v3)
    }
    if (1 <= (v10 % ((v10) | 1))) {
        fmt.Println(v28, v9)
    } else {
        v34 := GenS1{F0: 65535}
        v35 := GenS5{F0: 0, F1: -292}
        _ = len(v1)
        _ = v34
        _ = v35
    }
    v3 = 2
    v29 += (v3 * v10)
    v36 := GenS6{F0: v2, F1: 255}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v9
}
