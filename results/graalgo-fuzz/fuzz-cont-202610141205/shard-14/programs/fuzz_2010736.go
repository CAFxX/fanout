package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 9223372036854775807
    var v3 int16 = 1000
    var v4 uint64 = 42
    var v5 uint32 = 2147483647
    v6 := []int{178}
    var v7 float32 = ((0.1 + 1e10) + 3.14159)
    var v8 int16 = 1000
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v9 float64 = (((3.14159 / 1.0) * 0.1) / 0.1)
    v8 = v8
    if (v2 < v2) {
        switch -530 {
        case 128:
            if (v2 != (v2 / ((v2) | 1))) {
                if ((false || ((668 > 42) || false)) && true) {
                    _ = len(v1)
                } else {
                    if ((true || false) && (255 < 42)) {
                        v10 := GenS1{F0: ((v4 == v4) || false)}
                        var v11 uint64 = 4294967295
                        _ = v10
                        _ = v11
                    }
                }
                v8 += v3
            } else {
                v12 := GenS2{F0: 128, F1: (^(v4)), F2: "a"}
                v8 -= -675
                _ = v12
            }
        case 32767:
            switch "foo bar" {
            case "foo bar":
                if false {
                    fmt.Println(v2)
                    for fi13 := 0; fi13 < 14; fi13++ {
                        var v14 uint64 = 309
                        _ = fi13
                        _ = v14
                    }
                }
            case "a":
                var v15 uint16 = 65535
                _ = v15
            case "\n\t":
                switch "hello" {
                case "":
                    if false {
                        if (!(-2147483648 != -1)) {
                            fmt.Println(v8, v7, v5)
                            v16 := GenS3{F0: 128}
                            var v17 string = ""
                            _ = v16
                            _ = v17
                        }
                    } else {
                        v18 := GenS4{F0: 127}
                        fmt.Println(v3)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v18
                    }
                case "foo bar":
                    var v19 uint8 = 129
                    _ = v19
                case "a":
                    var v20 float64 = (-2.25 - 1e10)
                    _ = v20
                }
            default:
                var v21 uint64 = (v4 * 419)
                _ = v21
            }
        default:
            fmt.Println(v9, v4, v8)
        }
        if true {
            for fi22 := 0; fi22 < 4; fi22++ {
                if false {
                    v8 += (v3 + (v3 | 421))
                    v7 += (v7 + 0.1)
                    var v23 string = ""
                    _ = v23
                }
                _ = fi22
            }
            if (true && false) {
                for fi24 := 0; fi24 < 6; fi24++ {
                    if len(v6) > 0 {
                        v6[0] = fi24
                    }
                    _ = fi24
                }
                for fi25 := 0; fi25 < 6; fi25++ {
                    _ = len(v1)
                    var v26 uint64 = (v4 << ((206) & 63))
                    _ = fi25
                    _ = v26
                }
            } else {
                fmt.Println(v3, v3, v8)
                for fi27 := 0; fi27 < 16; fi27++ {
                    fmt.Println(v8, v8)
                    for fi28 := 0; fi28 < 8; fi28++ {
                        delete(v1, "a")
                        fmt.Println(v9, v4, v8)
                        _ = fi28
                    }
                    _ = fi27
                }
                switch "hello" {
                case "":
                    var v29 uint64 = (v4 - (v4 - v4))
                    _ = v29
                case "a":
                    var v30 float64 = (v9 - v9)
                    _ = v30
                case "hello":
                    if false {
                        var v31 uint16 = 32767
                        _ = v31
                    } else {
                        var v32 int32 = 954
                        fmt.Println(v4, v3, v8)
                        var v33 int8 = 56
                        _ = v32
                        _ = v33
                    }
                default:
                    if (!(-2.25 > 0.1)) {
                        if (!(false || ((-(v2)) <= -979))) {
                            switch 128 {
                            case -467:
                                for fi34 := 0; fi34 < 19; fi34++ {
                                    v8 -= (v3 % ((32767) | 1))
                                    _ = fi34
                                }
                            case 128:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        } else {
                            fmt.Println(v9, v7, v4)
                            v35 := GenS4{F0: 9223372036854775807}
                            _ = v35
                        }
                    }
                }
            }
        }
        v5 = v5
    } else {
        fmt.Println(v5)
        var v36 uint = 18446744073709551615
        var v37 uint16 = 255
        _ = v36
        _ = v37
    }
    if true {
        fmt.Println(v9, v3)
        switch ("hello" + "a") {
        case "\n\t":
            _ = len(v6)
        case "foo bar":
            if false {
                var v38 int16 = -405
                v3 += 343
                _ = v38
            }
        default:
            if false {
                switch -917 {
                case 302:
                    v2 = v2
                case 128:
                    v8 += -555
                default:
                    if len(v6) > 0 {
                        v6[0] = 9223372036854775807
                    }
                }
                if false {
                    var v39 float64 = v9
                    fmt.Println(v3)
                    _ = v39
                } else {
                    var v40 int8 = -128
                    v7 = 0.0
                    _ = v40
                }
                v1[("a" + "\n\t")] = "foo bar"
            }
        }
        v41 := GenS1{F0: false}
        _ = v41
    }
    fmt.Println(v3, v8)
    var v42 uint = 0
    var v43 int32 = -952
    fmt.Println(v4)
    v44 := GenS3{F0: 1}
    _ = len(v6)
    if true {
        v9 = v9
        v45 := GenS4{F0: 2147483647}
        v8 = v3
        _ = v45
    }
    for fi46 := 0; fi46 < 3; fi46++ {
        v47 := GenS5{F0: (fi46 | fi46)}
        _ = fi46
        _ = v47
    }
    if (128 <= 2) {
        v42 = v42
        v48 := GenS2{F0: 198, F1: 467, F2: "hello"}
        _ = v48
    } else {
        if (68 != int8(v3)) {
            if false {
                var v49 string = ("a" + "foo bar")
                switch (v49 + ("\n\t" + (v49 + v49))) {
                case "foo bar":
                    if (2 == 32767) {
                        for fi50 := 0; fi50 < 3; fi50++ {
                            var v51 bool = false
                            var v52 uint16 = 1
                            v6 = append(v6, fi50)
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                        v43 = (v43 / (((v43 << ((42) & 63))) | 1))
                    }
                case "\n\t":
                    fmt.Println(v49, v7)
                case "":
                    fmt.Println(v3, v2)
                }
                _ = v49
            } else {
                _ = len(v1)
            }
            for fi53 := 0; fi53 < 12; fi53++ {
                v43 = int32(v5)
                _ = fi53
            }
        }
        if (false || true) {
            v8 = -32768
            fmt.Println(v2)
        }
        fmt.Println(v42)
    }
    switch 32768 {
    case 283:
        fmt.Println(v5, v8, v8)
    case 9223372036854775807:
        v42 = 18446744073709551615
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
