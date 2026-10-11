package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 int }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 string; F1 int8 }

func main() {
    v1 := []int{9223372036854775807}
    var v2 uint64 = 42
    var v3 float32 = 0.0
    var v4 uint32 = 2
    v5 := []int{974}
    var v6 uint32 = (v4 << ((uint(v2)) & 63))
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(v2)
        _ = fi7
    }
    v4 = v6
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 bool = true
    v2 -= (v2 >> ((1) & 63))
    v9 := GenS1{F0: ((42 == 1) && v8)}
    switch 42 {
    case 2:
        v10 := GenS2{F0: 32767, F1: 2, F2: 4294967295}
        _ = v10
    }
    v11 := GenS2{F0: -1, F1: 127, F2: 128}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "foo bar" {
    case "foo bar":
        if false {
            switch 0 {
            case 733:
                var v12 uint32 = (v6 - v6)
                _ = v12
            }
            if v8 {
                var v13 bool = (v3 > v3)
                _ = v13
            } else {
                v14 := GenS3{F0: 2, F1: 130}
                v15 := GenS3{F0: 65535, F1: 1}
                v16 := GenS4{F0: "foo bar", F1: (v2 << ((647) & 63))}
                _ = v14
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 12; fi17++ {
                if len(v1) > 0 {
                    v1[0] = (fi17 << ((4294967295) & 63))
                }
                _ = fi17
            }
        }
    }
    var v18 int16 = -858
    var v19 int64 = 42
    for fi20 := 0; fi20 < 7; fi20++ {
        if v8 {
            v21 := GenS4{F0: "a", F1: v2}
            switch fi20 {
            case 32768:
                if len(v5) > 0 {
                    v5[0] = 791
                }
            case 9223372036854775807:
                if v8 {
                    var v22 int16 = 0
                    if (v4 < v6) {
                        fmt.Println(v4, v3, v3)
                        var v23 uint32 = v4
                        switch "hello" {
                        case "a":
                            if v8 {
                                for fi24 := 0; fi24 < 8; fi24++ {
                                    if len(v1) > 0 {
                                        v1[0] = (-(fi24))
                                    }
                                    var v25 float32 = (v3 + 1e10)
                                    _ = fi24
                                    _ = v25
                                }
                                v3 -= (float32(((((3.14159 / float64(v3)) * 0.1) * 1e10) / 0.1)) - (v3 * -2.25))
                            } else {
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    if (607 <= int32(v22)) {
                                        fmt.Println(v4, v8)
                                        v23 = (v23 & v4)
                                        if len(v1) > 0 {
                                            v1[0] = fi20
                                        }
                                    }
                                    if v8 {
                                        v2 += (v2 - (^(v2)))
                                        var v27 uint8 = 42
                                        _ = len(v5)
                                        _ = v27
                                    }
                                    _ = fi26
                                }
                            }
                        case "":
                            var v28 uint32 = 65535
                            _ = v28
                        }
                        _ = v23
                    } else {
                        v5 = append(v5, 4294967295)
                        var v29 string = ((("foo bar" + (("a" + "a") + "hello")) + "") + "hello")
                        _ = v29
                    }
                    _ = v22
                }
            default:
                fmt.Println(v18)
            }
            _ = v21
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 uint = 2
            var v32 uint = (v31 / (((v31 / (((v31 - 385)) | 1))) | 1))
            switch 42 {
            case -1:
                _ = len(v1)
            case 1:
                var v33 int64 = 42
                _ = v33
            }
            _ = fi30
            _ = v31
            _ = v32
        }
        if true {
            v2 = v2
            switch "a" {
            case "":
                fmt.Println(v8, v4)
            case "a":
                for fi34 := 0; fi34 < 5; fi34++ {
                    var v35 float32 = 0.1
                    v36 := GenS4{F0: (("\n\t" + "a") + ("hello" + ("hello" + "foo bar"))), F1: 213}
                    v1 = append(v1, (fi20 * (fi34 << ((319) & 63))))
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case "foo bar":
                v37 := GenS5{F0: "", F1: 1}
                _ = v37
            }
            v6 += v4
        }
        _ = fi20
    }
    v6 += (v4 << ((0) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
