package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 string; F2 int32 }

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{}
    var v3 uint = 1
    v4 := make(map[int]string)
    _ = len(v2)
    for fi5 := 0; fi5 < 6; fi5++ {
        v3 += v3
        var v6 float64 = (1e10 * 0.0)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: (true && (!(32768 <= 0)))}
    var v8 float32 = (v1 / (((v1 + float32(float64((v1 + (v1 * v1))))) - float32((0.1 + float64(-2.25)))) * 0.1))
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            v1 -= v8
            _ = fi10
        }
        if ((false || (!true)) && false) {
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 bool = (-128 == 1)
                fmt.Println(fi9, v3, v12)
                var v13 uint = 894
                _ = fi11
                _ = v12
                _ = v13
            }
            if false {
                if (1 == 32767) {
                    _ = len(v2)
                    fmt.Println(v8, fi9)
                    for fi14 := 0; fi14 < 2; fi14++ {
                        _ = len(v2)
                        _ = fi14
                    }
                } else {
                    var v15 int8 = 0
                    v3 = v3
                    if true {
                        if false {
                            if (((fi9 & fi9) > (fi9 << (((v3 % ((255) | 1))) & 63))) && (722 >= 65535)) {
                                var v16 uint64 = 847
                                var v17 bool = false
                                _ = v16
                                _ = v17
                            }
                            fmt.Println(v1)
                        } else {
                            delete(v4, fi9)
                            var v18 bool = ((((0 != 0) && (false || (v1 == v8))) && (!(true || ((v1 - float32((float64(float32((3.14159 / 0.1))) / 0.1))) != v1)))) && false)
                            _ = v18
                        }
                    } else {
                        fmt.Println(fi9, v3)
                        switch ((("a" + ("\n\t" + "hello")) + ("hello" + "")) + "a") {
                        case "hello":
                            switch "hello" {
                            case "hello":
                                v19 := GenS2{F0: 127, F1: ("\n\t" + "foo bar"), F2: -1}
                                _ = v19
                            case "\n\t":
                                if false {
                                    if true {
                                        for fi20 := 0; fi20 < 17; fi20++ {
                                            v21 := GenS1{F0: (42 <= 0)}
                                            fmt.Println(fi20, v3, fi9)
                                            var v22 uint8 = 95
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                        var v23 int32 = 2147483647
                                        _ = v23
                                    } else {
                                        v8 -= (v1 - float32(0.1))
                                    }
                                    v3 = 1
                                }
                            default:
                                fmt.Println(v8, fi9, v8)
                            }
                        }
                    }
                    _ = v15
                }
                if false {
                    fmt.Println(v3, v1)
                    v4[fi9] = ("hello" + ((("a" + "hello") + "\n\t") + "foo bar"))
                } else {
                    if true {
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                    } else {
                        v24 := v4[(fi9 ^ 322)]
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 5; fi25++ {
                        fmt.Println(v3)
                        var v26 int32 = -1
                        _ = fi25
                        _ = v26
                    }
                }
                var v27 uint16 = 916
                _ = v27
            }
            v1 = (float32(0.0) - float32(0.0))
        } else {
            fmt.Println(fi9, v1)
            v1 = 0.1
            fmt.Println(fi9, v8)
        }
        if (9223372036854775807 != 1) {
            var v28 int = (fi9 | (fi9 >> ((210) & 63)))
            v28 = v28
            _ = v28
        } else {
            if (1 < 42) {
                fmt.Println(v8, v8)
            }
        }
        _ = fi9
    }
    v3 = v3
    var v29 int16 = 0
    if false {
        v8 = v8
    } else {
        var v30 uint = 966
        for fi31 := 0; fi31 < 1; fi31++ {
            for fi32 := 0; fi32 < 1; fi32++ {
                v29 = (v29 - (v29 - (v29 << ((v30) & 63))))
                _ = fi32
            }
            _ = fi31
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
