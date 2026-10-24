package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 uint64 }
type GenS3 struct { F0 string; F1 int64; F2 int }
type GenS4 struct { F0 bool; F1 int64; F2 int }
type GenS5 struct { F0 uint32 }

func main() {
    v1 := []int{}
    var v2 string = "foo bar"
    v3 := []int{128, 127, 9223372036854775807, 582}
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    switch -2147483648 {
    case 206:
        v2 = ("foo bar" + (v2 + "foo bar"))
    case 9223372036854775807:
        v7 := GenS1{F0: 42, F1: 44}
        _ = v7
    }
    if (!(!false)) {
        for fi8 := 0; fi8 < 18; fi8++ {
            _ = len(v4)
            for fi9 := 0; fi9 < 16; fi9++ {
                if false {
                    var v10 int = fi8
                    _ = v10
                } else {
                    v11 := GenS1{F0: 482, F1: 127}
                    _ = v11
                }
                v12 := GenS1{F0: 1, F1: 185}
                fmt.Println(v2)
                var v13 int16 = 0
                fmt.Println(v13, v2)
                _ = fi9
                _ = v12
                _ = v13
            }
            _ = fi8
        }
        var v14 int32 = -1
        _ = v14
    }
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 6; fi16++ {
        if true {
            var v17 bool = true
            _ = v17
        }
        v15 -= 9223372036854775807
        var v18 int16 = -1
        _ = fi16
        _ = v18
    }
    v2 = (v2 + "hello")
    fmt.Println(v15, v15)
    if false {
        fmt.Println(v15)
        if false {
            v19 := GenS1{F0: 1, F1: 2}
            if true {
                delete(v4, -1)
                v20 := GenS2{F0: 130, F1: 0, F2: 0}
                _ = v20
            } else {
                v15 = (v15 & 42)
            }
            _ = v19
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            switch v2 {
            case "":
                v15 = v15
            case "a":
                v22 := GenS3{F0: "foo bar", F1: 0, F2: fi21}
                _ = v22
            case "\n\t":
                var v23 float32 = (((3.14159 + ((0.1 / ((float32(1e10) * (float32(float64(0.1)) / 1.5)) / (0.1 - -2.25))) / float32(((0.0 * 3.14159) - 3.14159)))) + float32(3.14159)) + 0.1)
                _ = v23
            }
            var v24 int16 = 1
            switch (fi21 + (fi21 / ((-1) | 1))) {
            case 2147483647:
                if true {
                    var v25 uint64 = 18446744073709551615
                    if (-1 >= -2147483648) {
                        v24 = (v24 & v24)
                        fmt.Println(v2)
                    }
                    _ = v25
                }
            case -1:
                v26 := v5[400]
                _ = v26
            case 128:
                if ((!false) && ((-2.25 + float64(0.1)) != 0.1)) {
                    var v27 float64 = float64((0.1 * (3.14159 + 1.5)))
                    v1 = append(v1, -1)
                    var v28 int32 = -1
                    _ = v27
                    _ = v28
                }
            default:
                switch fi21 {
                case 2147483647:
                    if false {
                        v6[fi21] = (v2 + "a")
                    }
                case 495:
                    switch "a" {
                    case "\n\t":
                        if true {
                            for fi29 := 0; fi29 < 3; fi29++ {
                                var v30 string = "a"
                                var v31 uint8 = 126
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            switch v2 {
                            case "foo bar":
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    _ = len(v6)
                                    _ = fi32
                                }
                            default:
                                var v33 int32 = 42
                                _ = v33
                            }
                            switch -47 {
                            case 9223372036854775807:
                                v34 := GenS1{F0: 548, F1: 221}
                                _ = v34
                            case 128:
                                if true {
                                    v35 := GenS4{F0: (!("hello" == v2)), F1: int64(fi21), F2: (-(fi21))}
                                    v36 := GenS1{F0: 0, F1: 138}
                                    _ = v35
                                    _ = v36
                                } else {
                                    v15 = v15
                                    var v37 bool = (true && false)
                                    _ = v37
                                }
                            case 42:
                                fmt.Println(v15, v24)
                            }
                        }
                    case "":
                        var v38 bool = (false || false)
                        _ = v38
                    case "foo bar":
                        var v39 uint32 = 2
                        _ = v39
                    }
                default:
                    var v40 string = ""
                    _ = v40
                }
            }
            _ = fi21
            _ = v24
        }
    }
    var v41 int = 32768
    _ = len(v1)
    var v42 uint16 = 255
    for fi43 := 0; fi43 < 17; fi43++ {
        v44 := GenS5{F0: 1}
        for fi45 := 0; fi45 < 10; fi45++ {
            v15 = 0
            fmt.Println(v2)
            var v46 float64 = 1.5
            _ = fi45
            _ = v46
        }
        v15 -= -1
        _ = fi43
        _ = v44
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        var v48 uint32 = 0
        var v49 uint8 = 42
        for fi50 := 0; fi50 < 1; fi50++ {
            v42 = (v42 >> ((65535) & 63))
            var v51 string = v2
            _ = fi50
            _ = v51
        }
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
