package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 bool }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    v1 := []string{"\n\t"}
    var v2 uint8 = 145
    var v3 uint = 237
    var v4 uint8 = (v2 ^ (^(v2)))
    v4 += v2
    switch "hello" {
    case "":
        fmt.Println(v4)
    case "hello":
        if (!(true || true)) {
            fmt.Println(v3, v4)
        }
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        for fi5 := 0; fi5 < 5; fi5++ {
            v3 -= (v3 << ((255) & 63))
            _ = fi5
        }
    }
    _ = len(v1)
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            if (true || false) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 -= uint8(v3)
            } else {
                if len(v1) > 0 {
                    v1[0] = (("hello" + "") + "")
                }
                v2 = v4
            }
            _ = len(v1)
            _ = fi7
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi6
    }
    v4 += 255
    var v8 int = -363
    var v9 float64 = -2.25
    var v10 float64 = float64(3.14159)
    for fi11 := 0; fi11 < 4; fi11++ {
        v4 += 255
        v9 += v10
        _ = fi11
    }
    if false {
        if false {
            if true {
                v10 = float64(0.0)
                for fi12 := 0; fi12 < 13; fi12++ {
                    v4 = (v4 >> ((uint(v4)) & 63))
                    var v13 uint8 = 255
                    if false {
                        if len(v1) > 0 {
                            v1[0] = ((("a" + (("\n\t" + ((("" + ("" + "foo bar")) + "\n\t") + "hello")) + "foo bar")) + "") + "\n\t")
                        }
                    }
                    _ = fi12
                    _ = v13
                }
                var v14 int32 = 1
                _ = v14
            }
        }
    }
    v1 = append(v1, "")
    v15 := GenS1{F0: 9223372036854775807, F1: v4, F2: (0 == -32768)}
    var v16 float64 = v9
    v8 = (v8 * (v8 | v8))
    v17 := GenS1{F0: -9223372036854775808, F1: (v2 + (^(v4))), F2: true}
    var v18 uint8 = v2
    v1 = append(v1, (((("a" + ("foo bar" + "foo bar")) + ("a" + "hello")) + ("" + "\n\t")) + (((("hello" + "") + "") + "\n\t") + "\n\t")))
    switch v8 {
    case 2147483647:
        v19 := GenS2{F0: "hello", F1: false}
        _ = v19
    case 330:
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint64 = 0
            switch (fi20 + fi20) {
            case -1:
                var v22 uint8 = (v4 >> ((v3) & 63))
                _ = v22
            case 32767:
                if false {
                    if ((v18 / ((v18) | 1)) == v2) {
                        var v23 int8 = 10
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 5; fi24++ {
                            var v25 float64 = (1.5 * (0.0 / 1.5))
                            v1 = append(v1, "")
                            fmt.Println(v4)
                            fmt.Println(v21, v25)
                            fmt.Println(v2, v21)
                            _ = fi24
                            _ = v25
                        }
                        if true {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            v21 = v21
                            if (uint64(v4) <= (v21 % ((4294967295) | 1))) {
                                v16 = v10
                                fmt.Println(v3)
                            }
                        } else {
                            v26 := GenS1{F0: 187, F1: 1, F2: (1e10 > (v10 - ((v9 * (float64((1e10 + (3.14159 + 0.0))) / (-2.25 - ((1.5 - -2.25) / 3.14159)))) + v10)))}
                            _ = v26
                        }
                    }
                } else {
                    if (561 >= 2) {
                        var v27 float32 = -2.25
                        v8 += (fi20 % ((-696) | 1))
                        if true {
                            if false {
                                var v28 int32 = -2147483648
                                v29 := GenS3{F0: true, F1: 46}
                                _ = v28
                                _ = v29
                            }
                        } else {
                            for fi30 := 0; fi30 < 3; fi30++ {
                                if true {
                                    var v31 int16 = 0
                                    _ = v31
                                }
                                if false {
                                    fmt.Println(v21, v4, v16)
                                    v3 = 1
                                    if false {
                                        v4 -= v4
                                        v4 -= (v4 + v2)
                                        var v32 uint16 = uint16(v3)
                                        _ = v32
                                    }
                                }
                                _ = len(v1)
                                _ = fi30
                            }
                            v10 = (-2.25 + float64(0.0))
                        }
                        _ = v27
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if ((-32768 >= -32768) && false) {
                            v8 += fi20
                            fmt.Println(v2)
                        } else {
                            var v33 uint32 = 9
                            _ = v33
                        }
                        fmt.Println(v9)
                    }
                    v18 += 0
                    if (1.5 != ((v9 - (1.5 / float64(0.1))) / 1e10)) {
                        switch (fi20 / ((112) | 1)) {
                        case 32768:
                            fmt.Println(v4)
                        }
                        var v34 uint16 = 1
                        _ = v34
                    }
                }
            case 42:
                if len(v1) > 0 {
                    v1[0] = ("" + "\n\t")
                }
            default:
                var v35 string = ("a" + "a")
                _ = v35
            }
            v3 = 38
            _ = fi20
            _ = v21
        }
    case -2147483648:
        v8 = -425
    default:
        v10 -= v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
