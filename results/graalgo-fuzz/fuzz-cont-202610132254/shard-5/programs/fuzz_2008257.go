package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int64 }
func genFn0() uint {
    var v7 float32 = ((3.14159 - 0.0) + (((0.1 - -2.25) - 3.14159) / 1.0))
    fmt.Println(v7)
    _ = v7
    return 65535
}
func genFn1(p0 bool) string {
    switch ("" + "\n\t") {
    case "foo bar":
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int32 = -2147483648
            for fi10 := 0; fi10 < 10; fi10++ {
                fmt.Println(fi10)
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    }
    var v11 int32 = -1
    _ = p0
    _ = v11
    return ""
}

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 175
    v3 := []int{42, 911, -39}
    v4 := make(map[int]bool)
    var v5 float64 = 0.0
    v6 := make(map[uint32]string)
    switch "\n\t" {
    case "\n\t":
        _ = len(v6)
    default:
        v2 -= (^(v2))
    }
    if (("\n\t" <= ("foo bar" + "a")) && false) {
        if true {
            if false {
                v6[0] = ""
            } else {
                v2 = v2
                var v12 float64 = float64(0.0)
                var v13 int64 = 9223372036854775807
                _ = v12
                _ = v13
            }
        } else {
            var v14 bool = (1.5 <= v5)
            v2 = (v2 >> ((18446744073709551615) & 63))
            _ = v14
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS1{F0: 32767}
            v2 -= 1
            _ = fi15
            _ = v16
        }
    } else {
        if true {
            v5 = (v5 - (0.1 + 0.0))
        } else {
            if ("foo bar" <= ("" + ("\n\t" + ("a" + "\n\t")))) {
                v5 = v5
            }
            v1 = v1
        }
        v5 = 0.1
    }
    var v17 bool = false
    if (((v1 + (v1 * float32(((v5 * 0.1) * ((float64(v1) * (float64(v1) * ((v5 / (v5 + 3.14159)) + v5))) * float64(1e10)))))) >= float32(((float64((((v1 - v1) + v1) * (v1 - v1))) + v5) - (v5 - (1.5 + 3.14159))))) && v17) {
        if ((-128 >= -128) || (v17 || v17)) {
            var v18 float64 = (v5 / 1.5)
            fmt.Println(v2, v1, v5)
            _ = v18
        }
        v19 := v6[955]
        if ((v2 - 255) > 128) {
            for fi20 := 0; fi20 < 14; fi20++ {
                switch fi20 {
                case -9223372036854775808:
                    if false {
                        v21 := GenS2{F0: (v2 >> ((303) & 63))}
                        _ = v21
                    } else {
                        switch fi20 {
                        case 32767:
                            v22 := GenS2{F0: 186}
                            _ = v22
                        case -9223372036854775808:
                            switch (fi20 % (((-(fi20))) | 1)) {
                            case 32768:
                                for fi23 := 0; fi23 < 4; fi23++ {
                                    v24 := GenS3{F0: -768, F1: -128, F2: 2}
                                    if len(v3) > 0 {
                                        v3[0] = fi23
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                            default:
                                v19 = (((("\n\t" + (v19 + v19)) + (((v19 + v19) + v19) + ((v19 + v19) + "hello"))) + "\n\t") + (v19 + "foo bar"))
                            }
                        }
                        v1 = (v1 - (1e10 + v1))
                    }
                default:
                    fmt.Println(fi20)
                }
                v25 := GenS2{F0: (v2 + (v2 & (^(v2))))}
                switch 32768 {
                case 9223372036854775807:
                    for fi26 := 0; fi26 < 14; fi26++ {
                        _ = len(v4)
                        v27 := GenS4{F0: (v2 ^ v2)}
                        if (0 >= 371) {
                            var v28 int64 = 734
                            v5 = 1.5
                            _ = v28
                        } else {
                            fmt.Println(fi26)
                            if (true && false) {
                                var v29 float64 = (v5 - -2.25)
                                var v30 int8 = -128
                                _ = v29
                                _ = v30
                            }
                            _ = len(v3)
                        }
                        v31 := GenS4{F0: v2}
                        _ = fi26
                        _ = v27
                        _ = v31
                    }
                case -19:
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v5 += v5
                        _ = len(v6)
                        _ = fi32
                    }
                default:
                    var v33 int8 = 120
                    _ = v33
                }
                _ = fi20
                _ = v25
            }
        }
        _ = v19
    } else {
        var v34 uint64 = 18446744073709551615
        _ = v34
    }
    var v35 bool = v17
    for fi36 := 0; fi36 < 6; fi36++ {
        v37 := GenS5{F0: fi36, F1: -1}
        for fi38 := 0; fi38 < 16; fi38++ {
            _ = len(v3)
            fmt.Println(v1, v17)
            _ = fi38
        }
        if v17 {
            fmt.Println(v1, fi36)
        }
        _ = fi36
        _ = v37
    }
    v39 := GenS6{F0: 1}
    var v40 uint64 = 2
    _ = len(v3)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
