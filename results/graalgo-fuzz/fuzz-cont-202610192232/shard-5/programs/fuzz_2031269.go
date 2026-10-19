package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64; F2 int32 }

func main() {
    v1 := []int{32767, 128}
    var v2 float32 = 0.1
    var v3 int16 = -32768
    var v4 uint = 1
    v5 := make(map[uint32]bool)
    var v6 int = -757
    v4 = v4
    fmt.Println(v3)
    v6 = 4294967295
    v7 := GenS1{F0: ("\n\t" + (("" + (("\n\t" + ("hello" + "hello")) + ("hello" + ("" + "a")))) + ("" + "\n\t")))}
    fmt.Println(v2, v4)
    var v8 uint = 770
    v6 -= (v6 - (v6 - (v6 << (((v4 & 1)) & 63))))
    v9 := GenS1{F0: "\n\t"}
    var v10 int16 = v3
    if len(v1) > 0 {
        v1[0] = 42
    }
    v11 := GenS1{F0: "foo bar"}
    v12 := GenS1{F0: ""}
    var v13 int32 = 2147483647
    switch "\n\t" {
    case "hello":
        v10 = v3
    case "a":
        if ("" < "hello") {
            for fi14 := 0; fi14 < 6; fi14++ {
                for fi15 := 0; fi15 < 19; fi15++ {
                    if true {
                        if true {
                            var v16 float32 = v2
                            switch fi14 {
                            case 4294967295:
                                if (71 > 255) {
                                    fmt.Println(v8, v4, v8)
                                    if (433 <= 2) {
                                        if false {
                                            v4 -= 127
                                            v3 = v3
                                            v13 = 1
                                            if true {
                                                fmt.Println(v2)
                                                var v17 bool = false
                                                v3 -= (^(v10))
                                                _ = v17
                                            }
                                        }
                                    } else {
                                        var v18 float64 = 3.14159
                                        var v19 uint8 = 128
                                        fmt.Println(v18)
                                        _ = v18
                                        _ = v19
                                    }
                                    switch ("hello" + "\n\t") {
                                    case "":
                                        v13 -= (v13 ^ v13)
                                    default:
                                        if true {
                                            var v20 int = fi15
                                            var v21 int64 = 9223372036854775807
                                            var v22 int8 = -128
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v13 -= v13
                                        }
                                    }
                                } else {
                                    var v23 int8 = 100
                                    _ = v23
                                }
                            case 48:
                                if true {
                                    if (0 > 2147483648) {
                                        v24 := GenS1{F0: "foo bar"}
                                        _ = v24
                                    }
                                    v3 = -1
                                    var v25 string = (((("" + "") + "hello") + "\n\t") + ("foo bar" + ("a" + ("" + "a"))))
                                    delete(v5, 505)
                                    v16 -= v2
                                    var v26 int16 = (v3 - v3)
                                    var v27 int64 = 9223372036854775807
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                            case -2147483648:
                                if (true || (1.5 != 0.1)) {
                                    v10 += (-(v10))
                                }
                            default:
                                v10 -= v10
                            }
                            _ = v16
                        } else {
                            v8 -= v8
                            switch "foo bar" {
                            case "hello":
                                v8 = v8
                            case "\n\t":
                                v28 := GenS1{F0: "hello"}
                                _ = v28
                            default:
                                v29 := v5[2147483648]
                                _ = v29
                            }
                            if true {
                                delete(v5, 65535)
                                v30 := GenS2{F0: 259}
                                _ = v30
                            } else {
                                if (!(0.0 == 0.1)) {
                                    if true {
                                        var v31 int16 = v3
                                        _ = v31
                                    }
                                    var v32 uint64 = 0
                                    _ = v32
                                }
                                if (true || (-1 < (v3 >> (((v8 << (((v8 - v8)) & 63))) & 63)))) {
                                    v33 := GenS2{F0: v8}
                                    var v34 uint16 = 32767
                                    _ = v33
                                    _ = v34
                                } else {
                                    v35 := GenS3{F0: true, F1: 4294967295, F2: (-(v13))}
                                    var v36 string = "foo bar"
                                    if false {
                                        var v37 bool = (1e10 >= float64(float32(1.5)))
                                        _ = v37
                                    }
                                    _ = v35
                                    _ = v36
                                }
                                v38 := GenS2{F0: (v4 + v8)}
                                _ = v38
                            }
                            v39 := v5[65535]
                            v13 -= v13
                            delete(v5, 568)
                            _ = v39
                        }
                        v10 += 1000
                        v40 := GenS2{F0: v4}
                        v5[767] = false
                        _ = v40
                    }
                    _ = fi15
                }
                _ = fi14
            }
            var v41 float64 = 1.5
            _ = v41
        }
    case "":
        _ = len(v5)
    default:
        var v42 uint16 = 255
        _ = v42
    }
    v43 := GenS3{F0: true, F1: 1, F2: 775}
    for fi44 := 0; fi44 < 2; fi44++ {
        var v45 int32 = (v13 ^ (v13 + v13))
        _ = fi44
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
