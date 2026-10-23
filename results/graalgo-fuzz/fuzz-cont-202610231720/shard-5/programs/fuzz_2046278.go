package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int16 = 330
    v2 := []uint32{}
    var v3 uint32 = 2147483648
    var v4 int64 = -9223372036854775808
    v5 := []int{909, -731}
    v3 = (v3 * 42)
    switch ("\n\t" + "") {
    case "\n\t":
        var v6 int64 = (v4 / ((v4) | 1))
        _ = v6
    case "a":
        v1 = v1
    case "hello":
        var v7 int32 = -201
        _ = v7
    }
    v1 += (v1 + v1)
    var v8 int16 = v1
    var v9 string = ("\n\t" + ("a" + ""))
    for fi10 := 0; fi10 < 11; fi10++ {
        switch ((v9 + (v9 + (v9 + "\n\t"))) + ((v9 + (v9 + v9)) + (v9 + v9))) {
        case "foo bar":
            v9 = ("hello" + v9)
        case "a":
            var v11 int8 = -128
            _ = v11
        case "hello":
            switch 9223372036854775807 {
            case 32767:
                v5 = append(v5, 9223372036854775807)
            case 2147483647:
                if false {
                    var v12 int16 = (^(v1))
                    if false {
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v9 = ""
                            v1 += v1
                            if (((65535 == 0) || true) || true) {
                                var v14 int8 = 0
                                v1 = (v8 >> ((18446744073709551615) & 63))
                                _ = v14
                            } else {
                                fmt.Println(v12, v1)
                            }
                            _ = fi13
                        }
                    }
                    _ = v12
                }
            }
        }
        var v15 uint = 418
        _ = fi10
        _ = v15
    }
    _ = len(v2)
    var v16 float32 = float32(0.1)
    v16 = 1e10
    for fi17 := 0; fi17 < 5; fi17++ {
        v8 -= v8
        switch fi17 {
        case 128:
            v3 -= v3
        default:
            switch fi17 {
            case 2147483647:
                v4 += 1
            case 9223372036854775807:
                if (!true) {
                    if (!(872 >= 42)) {
                        v16 += (v16 / (v16 / float32(0.1)))
                        v18 := GenS1{F0: 65535, F1: uint64(v3)}
                        _ = v18
                    } else {
                        var v19 bool = true
                        _ = v19
                    }
                    v3 = 42
                    v8 = v1
                } else {
                    v20 := GenS2{F0: 18446744073709551615}
                    fmt.Println(fi17, v9, v8)
                    _ = v20
                }
            default:
                if false {
                    if true {
                        var v21 uint64 = 130
                        _ = len(v5)
                        fmt.Println(v1, v21)
                        _ = v21
                    }
                    if (615 <= 255) {
                        v16 = (v16 * 1.5)
                    } else {
                        for fi22 := 0; fi22 < 16; fi22++ {
                            v4 = v4
                            _ = fi22
                        }
                        v8 -= v8
                    }
                    switch 2147483647 {
                    case -51:
                        if (255 <= 32767) {
                            for fi23 := 0; fi23 < 13; fi23++ {
                                if (!(!(42 >= 0))) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                }
                                fmt.Println(v1, v4)
                                _ = fi23
                            }
                            v24 := GenS2{F0: 404}
                            var v25 int8 = 30
                            _ = v24
                            _ = v25
                        } else {
                            if (!true) {
                                v3 = v3
                                v9 = v9
                            }
                        }
                    case 32767:
                        var v26 float64 = (3.14159 / float64(float32(float64(((((float32((1.5 / ((((1.5 - (float64(1e10) - 3.14159)) - (1e10 + 3.14159)) / 1.5) / (1.5 / 1.0)))) + 1.5) * v16) * (((1e10 * float32(1.5)) * (((v16 * v16) + ((float32(3.14159) * (v16 - float32(-2.25))) * (v16 - v16))) / float32(0.1))) + 0.0)) / (v16 - v16))))))
                        _ = v26
                    case -76:
                        if false {
                            var v27 bool = false
                            v9 = "\n\t"
                            _ = v27
                        }
                    }
                } else {
                    switch v9 {
                    case "\n\t":
                        v8 -= 32767
                    case "hello":
                        if true {
                            for fi28 := 0; fi28 < 1; fi28++ {
                                v4 += 0
                                var v29 uint8 = 182
                                _ = fi28
                                _ = v29
                            }
                        } else {
                            fmt.Println(v4, v9, v1)
                            v1 -= int16(v3)
                        }
                    case "":
                        var v30 string = "foo bar"
                        _ = v30
                    }
                    v8 += v8
                }
            }
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
