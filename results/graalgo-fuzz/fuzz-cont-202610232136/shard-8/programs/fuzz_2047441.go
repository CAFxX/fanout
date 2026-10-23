package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]string)
    v3 := []int{654}
    var v4 float64 = 1.5
    if true {
        var v5 float64 = v4
        _ = v5
    } else {
        v4 = (v4 / (float64((-2.25 - 0.1)) / v4))
        v6 := GenS1{F0: 127}
        _ = v6
    }
    v1 = (v1 >> ((4294967295) & 63))
    v4 = (v4 * v4)
    var v7 uint8 = 0
    v1 = v1
    switch ("hello" + "\n\t") {
    case "hello":
        if true {
            if true {
                v1 -= v1
                fmt.Println(v4, v7)
            } else {
                fmt.Println(v7, v1, v7)
            }
        }
    case "":
        v8 := GenS1{F0: 100}
        _ = v8
    }
    v1 += (^(v1))
    if true {
        v9 := v2[9223372036854775807]
        if (v4 <= float64(3.14159)) {
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS1{F0: -1}
                v3 = append(v3, -934)
                v4 -= -2.25
                _ = fi10
                _ = v11
            }
            v2[32767] = v9
            v3 = append(v3, 32768)
        }
        fmt.Println(v9)
        _ = v9
    } else {
        var v12 int = 127
        _ = v12
    }
    v13 := GenS2{F0: 2147483647}
    var v14 uint16 = 992
    v15 := GenS2{F0: 414}
    for fi16 := 0; fi16 < 3; fi16++ {
        v4 -= v4
        for fi17 := 0; fi17 < 20; fi17++ {
            fmt.Println(fi16, v14, fi17)
            v4 += (((float64(1.5) + (v4 - v4)) / v4) + v4)
            _ = fi17
        }
        v14 += v1
        _ = fi16
    }
    var v18 float32 = 1e10
    v19 := GenS3{F0: 842, F1: v14, F2: v7}
    for fi20 := 0; fi20 < 20; fi20++ {
        delete(v2, -9223372036854775808)
        v4 = (-2.25 / v4)
        for fi21 := 0; fi21 < 14; fi21++ {
            fmt.Println(fi21, v14)
            if false {
                v4 = v4
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    v4 += v4
                case "hello":
                    var v22 string = ""
                    _ = v22
                case "a":
                    var v23 uint64 = 521
                    if (-527 >= 805) {
                        fmt.Println(v14, v7)
                        var v24 uint = 4294967295
                        switch fi21 {
                        case 32768:
                            if false {
                                v25 := GenS4{F0: 1}
                                v14 -= (v1 >> ((uint(v7)) & 63))
                                if (88 == 401) {
                                    var v26 uint16 = (v14 | v1)
                                    switch "hello" {
                                    case "":
                                        var v27 uint8 = (^(v7))
                                        _ = v27
                                    case "\n\t":
                                        v4 = 0.1
                                    default:
                                        v24 -= 483
                                    }
                                    _ = v26
                                } else {
                                    var v28 string = (("a" + "a") + "a")
                                    _ = v28
                                }
                                _ = v25
                            } else {
                                var v29 string = "\n\t"
                                v7 = (v7 % ((v7) | 1))
                                v4 -= float64(1.5)
                                _ = v29
                            }
                        case -969:
                            fmt.Println(v18)
                        case -9223372036854775808:
                            var v30 int = fi20
                            _ = v30
                        default:
                            fmt.Println(v1, v4, fi20)
                        }
                        _ = v24
                    }
                    _ = v23
                default:
                    v1 -= (v14 + (^(v1)))
                }
                fmt.Println(v14, v7)
            }
            v18 = 0.0
            _ = fi21
        }
        _ = fi20
    }
    var v31 uint16 = 65535
    var v32 int = -791
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v7
}
