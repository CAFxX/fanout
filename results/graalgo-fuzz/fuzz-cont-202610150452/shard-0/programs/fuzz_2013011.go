package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint; F2 uint16 }

func main() {
    var v1 float32 = (0.1 * 3.14159)
    var v2 int32 = 0
    var v3 string = "a"
    for fi4 := 0; fi4 < 6; fi4++ {
        if ((128 == 0) && (0 != -9223372036854775808)) {
            v1 = ((float32(3.14159) - v1) * v1)
            var v5 string = (("\n\t" + v3) + v3)
            for fi6 := 0; fi6 < 11; fi6++ {
                fmt.Println(v2, fi4, v3)
                switch fi6 {
                case 32768:
                    if false {
                        var v7 bool = true
                        _ = v7
                    } else {
                        var v8 int64 = 9223372036854775807
                        _ = v8
                    }
                case 9223372036854775807:
                    if false {
                        v9 := []int{127, fi4}
                        var v10 string = "foo bar"
                        if (!(42 == 2147483647)) {
                            v11 := GenS1{F0: v5}
                            if (!(-14 <= 100)) {
                                if (1 >= 0) {
                                    switch 127 {
                                    case 42:
                                        v3 = (v10 + "hello")
                                    case 2147483647:
                                        v1 = v1
                                    case -9223372036854775808:
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    }
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    var v12 int8 = 39
                                    _ = v12
                                }
                                fmt.Println(v1)
                                var v13 int8 = -78
                                _ = v13
                            } else {
                                v1 = (v1 + 3.14159)
                            }
                            var v14 int = -1
                            _ = v11
                            _ = v14
                        } else {
                            v15 := GenS1{F0: v5}
                            _ = v15
                        }
                        _ = v10
                        _ = v9
                    } else {
                        if (v1 <= (float32((1.5 + 0.0)) - 3.14159)) {
                            fmt.Println(v5)
                            var v16 int8 = 1
                            v17 := []int{32767, (fi4 - fi6), (fi6 - fi6)}
                            _ = v16
                            _ = v17
                        } else {
                            if false {
                                v3 = ("hello" + v5)
                                var v18 uint8 = 0
                                v1 = v1
                                _ = v18
                            } else {
                                switch (fi4 << ((285) & 63)) {
                                case 0:
                                    v1 = v1
                                case -2147483648:
                                    var v19 int32 = 42
                                    v20 := GenS2{F0: "foo bar"}
                                    if (100 > -128) {
                                        v5 = "a"
                                    } else {
                                        v5 = (v3 + ("hello" + v3))
                                    }
                                    _ = v19
                                    _ = v20
                                case 32767:
                                    v21 := []int{}
                                    _ = v21
                                default:
                                    v3 = v5
                                }
                                v5 = "\n\t"
                            }
                            if (((255 <= 536) && false) || (false || false)) {
                                v22 := []int{fi6, fi6}
                                var v23 uint64 = 4294967295
                                _ = v22
                                _ = v23
                            }
                            v24 := GenS3{F0: ((v3 + v3) + "hello")}
                            _ = v24
                        }
                        v25 := GenS4{F0: 1, F1: 42, F2: 1}
                        _ = v25
                    }
                case 128:
                    if false {
                        var v26 int = 9223372036854775807
                        var v27 uint64 = 339
                        var v28 float64 = 1.5
                        _ = v26
                        _ = v27
                        _ = v28
                    } else {
                        fmt.Println(v1)
                    }
                default:
                    v2 -= int32(fi4)
                }
                var v29 uint = 255
                _ = fi6
                _ = v29
            }
            _ = v5
        }
        _ = fi4
    }
    var v30 int32 = v2
    switch v3 {
    case "foo bar":
        v31 := GenS1{F0: (v3 + (v3 + v3))}
        _ = v31
    default:
        v2 -= -138
    }
    v2 += (v30 & v2)
    if (true || (!false)) {
        if (620 != 1) {
            fmt.Println(v30, v1)
            var v32 int64 = -107
            var v33 float64 = 0.1
            _ = v32
            _ = v33
        }
        v34 := []int{}
        _ = v34
    }
    var v35 int = 127
    v2 -= (v30 / (((v2 - (^(v2)))) | 1))
    v36 := GenS2{F0: "\n\t"}
    var v37 int8 = 42
    v38 := GenS3{F0: v3}
    var v39 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
