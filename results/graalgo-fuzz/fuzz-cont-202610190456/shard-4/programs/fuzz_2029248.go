package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 bool; F2 int32 }
type GenS6 struct { F0 int8; F1 string }
func genFn0(p0 uint32, p1 uint) bool {
    if true {
        v6 := GenS1{F0: p0, F1: -402, F2: 32767}
        _ = v6
    }
    switch -222 {
    case 2147483647:
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int = 127
            v9 := []int{4294967295}
            p1 = 308
            _ = fi7
            _ = v8
            _ = v9
        }
    case 32767:
        v10 := []int{-2147483648, 32767, 0, 9223372036854775807}
        _ = v10
    }
    p1 = 255
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 int = -1
    var v4 string = ""
    v5 := make(map[uint32]string)
    fmt.Println(v4, v4)
    v4 = v4
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := []int{fi11, v3, (v3 & 32767)}
        switch 127 {
        case 591:
            for fi13 := 0; fi13 < 15; fi13++ {
                if (1 == 2) {
                    if true {
                        v3 += 0
                    } else {
                        switch v4 {
                        case "\n\t":
                            var v14 float64 = 3.14159
                            var v15 uint16 = 0
                            switch "" {
                            case "hello":
                                fmt.Println(v15, v14, v3)
                            case "\n\t":
                                var v16 float64 = v14
                                _ = v16
                            case "":
                                fmt.Println(v15, fi11)
                            }
                            _ = v14
                            _ = v15
                        case "":
                            switch v3 {
                            case 0:
                                v17 := GenS1{F0: 42, F1: (fi11 * (fi11 - fi11)), F2: 42}
                                _ = v17
                            case -792:
                                v3 = v3
                            }
                        }
                        _ = len(v5)
                        v18 := GenS1{F0: 116, F1: fi11, F2: 0}
                        _ = v18
                    }
                    fmt.Println(v4)
                    var v19 int8 = 127
                    fmt.Println(fi13, v3, v19)
                    if len(v12) > 0 {
                        v12[0] = (fi11 - 793)
                    }
                    v4 = v4
                    _ = v19
                }
                var v20 int = fi11
                if true {
                    var v21 uint32 = 2147483647
                    _ = v21
                }
                _ = fi13
                _ = v20
            }
        }
        switch v3 {
        case 2147483647:
            v4 = v4
        default:
            var v22 int = 2147483647
            _ = v22
        }
        _ = fi11
        _ = v12
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 string = v4
        _ = fi23
        _ = v24
    }
    v3 += v3
    var v25 int = (v3 >> ((65535) & 63))
    v26 := GenS1{F0: 355, F1: (^(v3)), F2: 1000}
    for fi27 := 0; fi27 < 3; fi27++ {
        if ((!(188 >= 2147483647)) || ((!((v3 / (((v3 & 128)) | 1)) >= (-(v25)))) && (-1 > -23))) {
            var v28 string = "hello"
            v25 += 1
            _ = v28
        } else {
            v3 = (fi27 + 4294967295)
            switch "a" {
            case "foo bar":
                switch 0 {
                case 42:
                    v29 := GenS2{F0: 981}
                    _ = v29
                case -9223372036854775808:
                    v30 := GenS3{F0: -780, F1: 4294967295}
                    _ = v30
                case 2147483647:
                    v31 := GenS4{F0: 127}
                    _ = v31
                default:
                    switch ((("a" + "") + v4) + (v4 + "foo bar")) {
                    case "\n\t":
                        var v32 int = 4294967295
                        _ = v32
                    case "hello":
                        fmt.Println(v25, fi27, v3)
                    case "foo bar":
                        if false {
                            if true {
                                v33 := GenS3{F0: 42, F1: 18446744073709551615}
                                for fi34 := 0; fi34 < 2; fi34++ {
                                    v3 += (fi34 / ((fi27) | 1))
                                    _ = fi34
                                }
                                var v35 uint8 = 42
                                _ = v33
                                _ = v35
                            } else {
                                v4 = v4
                                v36 := []int{-801}
                                v37 := GenS4{F0: 100}
                                _ = v36
                                _ = v37
                            }
                        }
                    }
                }
            case "hello":
                _ = len(v1)
            case "":
                switch (v3 | v25) {
                case 4294967295:
                    for fi38 := 0; fi38 < 9; fi38++ {
                        v39 := GenS5{F0: 42, F1: false, F2: 0}
                        v4 = v4
                        _ = fi38
                        _ = v39
                    }
                }
            }
        }
        v25 -= v3
        var v40 int8 = 1
        _ = fi27
        _ = v40
    }
    v25 = v25
    if (638 >= 1) {
        v41 := GenS3{F0: -320, F1: 42}
        var v42 int8 = 100
        _ = v41
        _ = v42
    }
    v43 := GenS4{F0: 34}
    for fi44 := 0; fi44 < 17; fi44++ {
        if (-13 < 112) {
            if true {
                fmt.Println(v3)
            } else {
                v45 := GenS6{F0: 69, F1: ("" + v4)}
                delete(v2, (-(v25)))
                _ = v45
            }
        }
        _ = fi44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
