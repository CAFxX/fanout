package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int; F1 uint8; F2 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16 }

func main() {
    var v1 uint = 876
    var v2 float32 = 1e10
    var v3 uint32 = 124
    var v4 int16 = -371
    if false {
        fmt.Println(v1, v3)
        if false {
            if true {
                if (!false) {
                    v5 := []int{}
                    var v6 int8 = 42
                    if true {
                        v4 = v4
                        for fi7 := 0; fi7 < 1; fi7++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi7
                        }
                    } else {
                        v1 = 23
                        var v8 uint = v1
                        _ = v8
                    }
                    _ = v5
                    _ = v6
                } else {
                    if (v1 == v1) {
                        if false {
                            var v9 int64 = 1
                            v10 := GenS1{F0: (v4 / ((79) | 1)), F1: v9}
                            _ = v10
                            _ = v9
                        }
                        var v11 uint8 = uint8(v1)
                        switch ("" + "") {
                        case "hello":
                            v12 := GenS2{F0: 549, F1: int32(v3)}
                            _ = v12
                        case "a":
                            if ((!(255 <= 63)) && (-585 <= 0)) {
                                var v13 uint = v1
                                var v14 float32 = (((1.5 / ((v2 + 0.0) / v2)) * float32(0.0)) * 1.5)
                                switch "hello" {
                                case "hello":
                                    fmt.Println(v4, v14, v3)
                                case "":
                                    v14 = (v14 / v2)
                                case "a":
                                    v13 -= uint(v3)
                                }
                                _ = v13
                                _ = v14
                            } else {
                                v3 = v3
                            }
                        case "":
                            v15 := []int{32767, -179}
                            _ = v15
                        }
                        _ = v11
                    }
                    if true {
                        v16 := GenS3{F0: (42 >= 1), F1: 1}
                        v17 := []int{128}
                        v1 -= (v1 - v1)
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(v3, v2, v1)
                        if false {
                            v1 += (v1 << ((uint(v3)) & 63))
                            for fi18 := 0; fi18 < 14; fi18++ {
                                v19 := GenS1{F0: v4, F1: -479}
                                fmt.Println(fi18, v1)
                                _ = fi18
                                _ = v19
                            }
                        } else {
                            v20 := GenS4{F0: int(v3), F1: 150, F2: (!true)}
                            if true {
                                fmt.Println(v3, v1, v3)
                                var v21 int16 = (v4 - (v4 % (((v4 >> ((v1) & 63))) | 1)))
                                v3 += (v3 >> ((2) & 63))
                                _ = v21
                            } else {
                                for fi22 := 0; fi22 < 18; fi22++ {
                                    var v23 uint32 = 2
                                    fmt.Println(v4, v2)
                                    _ = fi22
                                    _ = v23
                                }
                                var v24 uint8 = 1
                                _ = v24
                            }
                            v1 += (^(v1))
                            _ = v20
                        }
                        var v25 int64 = 9223372036854775807
                        _ = v25
                    }
                }
            } else {
                v26 := GenS3{F0: ((!(!true)) || true), F1: 0}
                _ = v26
            }
            if true {
                for fi27 := 0; fi27 < 10; fi27++ {
                    var v28 uint32 = 800
                    v3 = (v3 * 4294967295)
                    _ = fi27
                    _ = v28
                }
                var v29 float64 = 3.14159
                _ = v29
            }
        } else {
            var v30 int64 = 734
            v31 := []int{42, -350, 128}
            _ = v30
            _ = v31
        }
    }
    var v32 int32 = 42
    v33 := []int{}
    switch 32767 {
    case 1:
        var v34 uint16 = 0
        _ = v34
    default:
        v1 -= v1
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        v36 := GenS5{F0: (^(v3))}
        var v37 float64 = (1.5 + 1e10)
        v2 -= v2
        _ = fi35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        switch ((("foo bar" + "a") + "foo bar") + "a") {
        case "hello":
            v39 := GenS6{F0: 0}
            _ = v39
        }
        _ = len(v33)
        var v40 int32 = 0
        _ = fi38
        _ = v40
    }
    v41 := GenS4{F0: -1, F1: 128, F2: ((!false) && (18446744073709551615 <= 42))}
    for fi42 := 0; fi42 < 17; fi42++ {
        v3 = (v3 / ((v3) | 1))
        var v43 uint = v1
        var v44 uint32 = 604
        _ = fi42
        _ = v43
        _ = v44
    }
    var v45 float32 = float32(((((((0.0 + ((3.14159 / (-2.25 * (1e10 + 1.5))) * 3.14159)) - 1.5) * (-2.25 / (1e10 * 1.5))) / 1e10) - -2.25) - (1.5 * (1e10 / 1.0))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v45
}
