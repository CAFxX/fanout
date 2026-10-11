package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 32767
    v3 := []string{"", ""}
    var v4 uint64 = 0
    var v5 uint16 = v2
    var v6 int64 = 42
    var v7 uint32 = 0
    var v8 uint32 = v7
    var v9 int = 359
    fmt.Println(v8, v9, v2)
    v8 -= v8
    if (!((uint8(v2) <= uint8(v4)) || ((false || false) && ((v6 & v6) >= (v6 >> ((620) & 63)))))) {
        for fi10 := 0; fi10 < 14; fi10++ {
            v8 = (^(v8))
            fmt.Println(v7, fi10)
            _ = fi10
        }
        v11 := v1[(("hello" + "hello") + "")]
        v12 := GenS1{F0: ("a" + ("a" + "a"))}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int8 = -1
        fmt.Println(v8, v9)
        var v15 int = (v9 * fi13)
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 string = (("" + "") + (("" + (((("a" + "foo bar") + "hello") + "hello") + "")) + "hello"))
    for fi17 := 0; fi17 < 8; fi17++ {
        fmt.Println(v4)
        for fi18 := 0; fi18 < 18; fi18++ {
            if true {
                if (((v16 + "") + ((v16 + "a") + "hello")) < v16) {
                    v7 -= (v7 / ((uint32(v2)) | 1))
                }
                v19 := GenS1{F0: v16}
                if false {
                    _ = len(v3)
                    var v20 float32 = ((-2.25 * float32((1.5 / 1.0))) * float32(1e10))
                    _ = v20
                } else {
                    if (2147483647 != v7) {
                        v16 = v16
                        _ = len(v3)
                    } else {
                        v9 += (fi18 >> ((18446744073709551615) & 63))
                        if len(v3) > 0 {
                            v3[0] = v16
                        }
                        var v21 string = v16
                        _ = v21
                    }
                    v22 := GenS1{F0: "a"}
                    if false {
                        var v23 string = ""
                        if false {
                            var v24 int32 = -2147483648
                            fmt.Println(v5)
                            var v25 string = ""
                            switch 805 {
                            case 841:
                                v5 -= v5
                            case 4294967295:
                                fmt.Println(v23)
                            default:
                                v24 += v24
                            }
                            v24 -= v24
                            _ = v24
                            _ = v25
                        }
                        _ = v23
                    } else {
                        v5 = (v5 << ((0) & 63))
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v16 = "hello"
                    }
                    var v26 float64 = 1e10
                    _ = len(v3)
                    fmt.Println(v4, v9)
                    v4 += (^(v4))
                    var v27 int = 4294967295
                    _ = v22
                    _ = v26
                    _ = v27
                }
                v3 = append(v3, "a")
                fmt.Println(v5, fi17, v9)
                _ = v19
            } else {
                switch fi17 {
                case 127:
                    var v28 uint = 18446744073709551615
                    _ = v28
                case 464:
                    _ = len(v3)
                default:
                    if ((-540 <= int16(v6)) && true) {
                        switch v16 {
                        case "foo bar":
                            if ((32767 <= 32767) && false) {
                                var v29 uint8 = 1
                                var v30 uint32 = (v8 % ((uint32(v5)) | 1))
                                _ = v29
                                _ = v30
                            }
                        case "hello":
                            v31 := GenS1{F0: (("a" + (("\n\t" + v16) + "a")) + ("a" + "foo bar"))}
                            _ = v31
                        }
                    }
                }
            }
            if true {
                v32 := GenS2{F0: 2147483647, F1: 42}
                delete(v1, ("\n\t" + "a"))
                _ = v32
            } else {
                if (v9 < (v9 / ((fi17) | 1))) {
                    v4 = 635
                    v4 -= v4
                    v8 -= v8
                } else {
                    var v33 int = (v9 | (v9 + (fi18 >> ((18446744073709551615) & 63))))
                    if (((-2.25 + 1.5) == 1.5) && (1000 != -591)) {
                        if (!((uint32(v5) < v8) && true)) {
                            fmt.Println(v6, v5, v33)
                        }
                        var v34 float64 = 1e10
                        _ = v34
                    }
                    switch int(v2) {
                    case 92:
                        v2 = (^(v2))
                        v1["\n\t"] = ((false && true) && false)
                    case -9223372036854775808:
                        fmt.Println(v4)
                    }
                    if (0 < 42) {
                        switch (fi17 * 32767) {
                        case -2147483648:
                            var v35 uint64 = (^(v4))
                            _ = v35
                        case 738:
                            delete(v1, v16)
                        }
                        v9 -= fi18
                    }
                    var v36 string = "hello"
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + (((v36 + "\n\t") + ("" + "foo bar")) + "foo bar"))
                    }
                    var v37 int8 = 42
                    _ = v33
                    _ = v36
                    _ = v37
                }
                var v38 uint = 2
                _ = v38
            }
            _ = fi18
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
