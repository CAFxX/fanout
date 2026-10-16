package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 100
    var v3 int64 = -9223372036854775808
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 int16 = 981
        var v6 float32 = (3.14159 * (1.5 + -2.25))
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint64 = 996
    v8 := []int{-2147483648, 9223372036854775807}
    v3 += -784
    if true {
        v9 := GenS1{F0: 253}
        var v10 bool = true
        _ = v10
        _ = v9
    } else {
        v11 := GenS2{F0: -822}
        _ = v11
    }
    v1["foo bar"] = ""
    var v12 int16 = 0
    v2 = v2
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v13 int16 = v12
    v2 = v2
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 float64 = 3.14159
        var v16 uint32 = 0
        if false {
            for fi17 := 0; fi17 < 5; fi17++ {
                switch fi14 {
                case 0:
                    v18 := GenS2{F0: 9223372036854775807}
                    _ = v18
                }
                v19 := GenS1{F0: 42}
                _ = fi17
                _ = v19
            }
            if ((v3 / (((v3 & v3)) | 1)) == v3) {
                for fi20 := 0; fi20 < 14; fi20++ {
                    v21 := GenS2{F0: 861}
                    _ = fi20
                    _ = v21
                }
            } else {
                if true {
                    v15 = float64(1.5)
                    switch fi14 {
                    case -9223372036854775808:
                        if true {
                            for fi22 := 0; fi22 < 12; fi22++ {
                                v7 += (v7 & (v7 * (v7 & v7)))
                                fmt.Println(fi22, v15)
                                _ = fi22
                            }
                            var v23 uint = 0
                            if true {
                                switch (-(fi14)) {
                                case -804:
                                    if true {
                                        var v24 int64 = 1
                                        _ = v24
                                    }
                                default:
                                    if false {
                                        for fi25 := 0; fi25 < 15; fi25++ {
                                            v3 = v3
                                            v26 := v1[("hello" + (((("\n\t" + ("\n\t" + ("\n\t" + "\n\t"))) + "a") + "a") + ("" + "\n\t")))]
                                            fmt.Println(v3, v23, v16)
                                            v8 = append(v8, int(v23))
                                            v27 := GenS3{F0: v7, F1: (v3 / ((0) | 1))}
                                            _ = fi25
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                }
                                v28 := GenS4{F0: (-(v12)), F1: (v13 & 42)}
                                var v29 uint16 = 65535
                                _ = v28
                                _ = v29
                            }
                            _ = v23
                        }
                    case -706:
                        if (("" + ("a" + ("" + "\n\t"))) <= "foo bar") {
                            _ = len(v1)
                        } else {
                            for fi30 := 0; fi30 < 16; fi30++ {
                                fmt.Println(v12, v2, v15)
                                switch (fi14 * fi30) {
                                case -2147483648:
                                    _ = len(v1)
                                case -1:
                                    v15 += v15
                                }
                                var v31 uint64 = (v7 << ((859) & 63))
                                switch (fi30 & -9223372036854775808) {
                                case -1:
                                    if len(v8) > 0 {
                                        v8[0] = 9223372036854775807
                                    }
                                case 127:
                                    if len(v8) > 0 {
                                        v8[0] = (fi30 * fi14)
                                    }
                                case 42:
                                    var v32 int = 32767
                                    _ = v32
                                }
                                v2 = (^(v2))
                                fmt.Println(v15)
                                _ = fi30
                                _ = v31
                            }
                            v7 = 0
                        }
                    default:
                        v13 += v12
                    }
                } else {
                    var v33 uint32 = v16
                    var v34 string = "foo bar"
                    v13 = v12
                    _ = v33
                    _ = v34
                }
                if ((0 > -92) && ((!false) && (!true))) {
                    fmt.Println(v12)
                    var v35 float64 = -2.25
                    _ = v35
                } else {
                    var v36 uint16 = 65535
                    _ = v36
                }
                v37 := v1[("" + "a")]
                _ = v37
            }
        } else {
            delete(v1, ("a" + ""))
            var v38 float64 = 0.1
            _ = v38
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v39 float64 = -2.25
    switch 128 {
    case 42:
        var v40 string = "foo bar"
        _ = v40
    case 32767:
        fmt.Println(v13, v13, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v7
    _ = v8
}
