package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 string; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int32; F1 bool; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 int8 = 42
    v5 := GenS1{F0: -2147483648, F1: -553}
    v6 := GenS2{F0: 0}
    v7 := GenS3{F0: (^(v4)), F1: "a", F2: (("\n\t" < "a") && (32768 >= 42))}
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v1 -= v1
            v10 := GenS2{F0: -338}
            v11 := GenS2{F0: 1}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int16 = -385
        v12 = v12
        _ = fi8
        _ = v12
    }
    if (true || false) {
        fmt.Println(v1, v4)
    } else {
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 int64 = 0
            _ = fi13
            _ = v14
        }
        fmt.Println(v4)
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 int32 = 42
            _ = fi15
            _ = v16
        }
    }
    if (75 > 18446744073709551615) {
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := []int{fi17}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 5; fi20++ {
                switch -1 {
                case 128:
                    v21 := GenS4{F0: (-(v4))}
                    _ = v21
                case 0:
                    switch fi19 {
                    case 2147483647:
                        var v22 int64 = 42
                        _ = v22
                    case 32767:
                        if (true && (v1 > v1)) {
                            if false {
                                v23 := GenS5{F0: "hello", F1: 4294967295}
                                if true {
                                    switch (("hello" + "") + "foo bar") {
                                    case "hello":
                                        if (145 >= 183) {
                                            v4 += v4
                                            v24 := GenS6{F0: 411, F1: false, F2: 0}
                                            _ = v24
                                        } else {
                                            switch (("hello" + "") + "foo bar") {
                                            case "\n\t":
                                                v25 := GenS7{F0: false}
                                                _ = v25
                                            default:
                                                var v26 string = "a"
                                                _ = v26
                                            }
                                        }
                                    default:
                                        if false {
                                            var v27 uint = 1
                                            _ = v27
                                        } else {
                                            v4 -= -56
                                        }
                                    }
                                    v4 -= (v4 ^ 100)
                                    v4 = (v4 / (((v4 << ((42) & 63))) | 1))
                                }
                                fmt.Println(v1, fi20)
                                v28 := []int{}
                                _ = v23
                                _ = v28
                            }
                            v29 := GenS6{F0: 498, F1: (fi19 > (fi20 + 58)), F2: -622}
                            _ = v29
                        } else {
                            var v30 uint32 = 132
                            _ = v30
                        }
                    default:
                        v31 := GenS2{F0: 2147483647}
                        _ = v31
                    }
                }
                _ = fi20
            }
            _ = fi19
        }
        v4 = 45
    }
    var v32 uint64 = 0
    v33 := []int{}
    if (1000 == 42) {
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 float64 = 3.14159
            if ("\n\t" == "hello") {
                if true {
                    fmt.Println(v4, v32, fi34)
                    fmt.Println(v4, v35, v35)
                }
                if (false || false) {
                    for fi36 := 0; fi36 < 11; fi36++ {
                        if (0 != 1000) {
                            v37 := GenS1{F0: -1, F1: 42}
                            switch fi36 {
                            case -1:
                                if (!true) {
                                    v38 := GenS1{F0: 625, F1: 0}
                                    v39 := GenS7{F0: false}
                                    v40 := GenS3{F0: 96, F1: "", F2: false}
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                } else {
                                    if len(v33) >= 2 {
                                        v33 = v33[1:len(v33)]
                                    }
                                }
                            case 9223372036854775807:
                                var v41 int64 = 9223372036854775807
                                _ = v41
                            }
                            _ = v37
                        } else {
                            if true {
                                if (!(4294967295 <= 42)) {
                                    switch "\n\t" {
                                    case "":
                                        v33 = append(v33, fi34)
                                    default:
                                        var v42 uint32 = 673
                                        _ = v42
                                    }
                                }
                                var v43 int16 = 853
                                _ = v43
                            } else {
                                v44 := GenS2{F0: -653}
                                _ = v44
                            }
                            switch ("" + "\n\t") {
                            case "\n\t":
                                v4 = v4
                            }
                            _ = len(v33)
                        }
                        _ = fi36
                    }
                } else {
                    v35 += 0.1
                }
                _ = len(v3)
            }
            _ = fi34
            _ = v35
        }
        fmt.Println(v32, v4, v4)
        v33 = append(v33, 302)
    }
    var v45 uint32 = 65535
    switch -1 {
    case 9223372036854775807:
        v46 := GenS8{F0: (v45 & v45), F1: 42}
        _ = v46
    case -148:
        var v47 int16 = 1
        _ = v47
    case 42:
        v45 -= v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
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
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
