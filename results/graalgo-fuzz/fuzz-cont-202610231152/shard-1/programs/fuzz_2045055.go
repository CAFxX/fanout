package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint64) string {
    var v5 bool = false
    var v6 float32 = 1.5
    v7 := []int{485, 103}
    v8 := GenS1{F0: 42, F1: ""}
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint16 = 65535
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return ("\n\t" + ("" + "hello"))
}

func main() {
    var v1 float32 = (0.0 * float32(float64(((float32(float64(1e10)) + (float32(float64(0.0)) * ((0.1 + 0.0) + 0.0))) * ((-2.25 / 1.0) + -2.25)))))
    var v2 int32 = 0
    var v3 uint16 = 880
    v4 := []string{}
    var v11 string = ("\n\t" + "foo bar")
    v4 = append(v4, v11)
    v1 = v1
    v12 := GenS2{F0: 1, F1: -275}
    switch v11 {
    case "foo bar":
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v13 bool = (-1 != 42)
            _ = v13
        default:
            var v14 int = 42
            _ = v14
        }
    }
    v15 := GenS3{F0: (v2 & 1)}
    for fi16 := 0; fi16 < 10; fi16++ {
        if false {
            fmt.Println(v11, v11)
        }
        if (v11 == v11) {
            if (1000 < 0) {
                v17 := GenS4{F0: 1, F1: 255, F2: 2}
                var v18 int = fi16
                var v19 uint16 = v3
                _ = v17
                _ = v18
                _ = v19
            } else {
                var v20 string = ((("foo bar" + v11) + (v11 + (v11 + "\n\t"))) + "foo bar")
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 uint8 = 4
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
            if true {
                v2 += 0
                var v23 int32 = (^(v2))
                _ = v23
            } else {
                v11 = v11
                if (true && (v11 > "\n\t")) {
                    if ((v3 + (v3 >> ((255) & 63))) <= 548) {
                        v11 = (v11 + v11)
                    } else {
                        var v24 int64 = int64(fi16)
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 string = v11
                        var v27 bool = (!false)
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    if false {
                        var v28 string = v11
                        _ = v28
                    } else {
                        v1 += ((v1 * v1) * 3.14159)
                        v29 := GenS1{F0: v3, F1: ((v11 + "") + v11)}
                        _ = v29
                    }
                }
                v3 = v3
            }
        }
        fmt.Println(v1)
        _ = fi16
    }
    v30 := GenS5{F0: 74}
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 uint = 65535
        var v33 float64 = 1.5
        v33 = float64((float32(v33) * v1))
        _ = fi31
        _ = v32
        _ = v33
    }
    if ((("\n\t" + v11) <= v11) || (78 <= 72)) {
        var v34 int16 = 0
        for fi35 := 0; fi35 < 8; fi35++ {
            v36 := GenS5{F0: 1}
            if (!(false || (false && true))) {
                switch 9223372036854775807 {
                case 42:
                    for fi37 := 0; fi37 < 5; fi37++ {
                        _ = len(v4)
                        if (1e10 >= 1e10) {
                            fmt.Println(v34, v2, v34)
                            var v38 string = (v11 + (v11 + (v11 + v11)))
                            var v39 uint8 = 196
                            _ = v38
                            _ = v39
                        } else {
                            if (!true) {
                                fmt.Println(v2, fi37)
                            }
                        }
                        switch (fi37 ^ (fi37 / (((fi37 * -2147483648)) | 1))) {
                        case 127:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 189:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 9223372036854775807:
                            var v40 uint8 = 127
                            if (((v2 ^ v2) != 0) && (4294967295 == 0)) {
                                v11 = "foo bar"
                                var v41 int64 = -1
                                switch v11 {
                                case "":
                                    v2 -= v2
                                default:
                                    v42 := GenS3{F0: (v2 % ((v2) | 1))}
                                    _ = v42
                                }
                                _ = v41
                            }
                            _ = v40
                        default:
                            v4 = append(v4, "\n\t")
                            var v43 float64 = 1.5
                            _ = v43
                        }
                        _ = fi37
                    }
                case 2147483647:
                    switch v11 {
                    case "hello":
                        switch fi35 {
                        case 937:
                            v44 := GenS4{F0: 42, F1: 0, F2: 2}
                            _ = v44
                        default:
                            var v45 float64 = 3.14159
                            _ = v45
                        }
                    default:
                        for fi46 := 0; fi46 < 18; fi46++ {
                            v34 += v34
                            _ = fi46
                        }
                    }
                }
                fmt.Println(v3, fi35, v34)
            }
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
