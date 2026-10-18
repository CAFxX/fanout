package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int8; F1 int8; F2 int64 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS5 struct { F0 int; F1 bool; F2 int8 }
type GenS6 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 1
    v3 := []int{450, 4294967295, -207, 59}
    var v4 bool = false
    var v5 uint16 = (v2 / ((32767) | 1))
    v6 := GenS1{F0: v2, F1: v1, F2: 444}
    if ((648 >= -2147483648) && v4) {
        if len(v3) > 0 {
            v3[0] = -996
        }
        var v7 bool = (!(!(!false)))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v7
    } else {
        _ = len(v3)
        v3 = append(v3, 32768)
    }
    v3 = append(v3, 1)
    fmt.Println(v4, v5, v4)
    fmt.Println(v4)
    _ = len(v3)
    if (42 > -29) {
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := GenS2{F0: 2, F1: 82, F2: v2}
            _ = fi8
            _ = v9
        }
        if false {
            if v4 {
                _ = len(v3)
            }
            var v10 float32 = 1.5
            _ = v10
        }
        var v11 uint8 = 1
        _ = v11
    } else {
        var v12 int8 = -18
        v13 := GenS3{F0: v12, F1: v12, F2: int64(v12)}
        switch 1 {
        case 579:
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS2{F0: 90, F1: 42, F2: (v5 * v5)}
                v5 += 42
                _ = fi14
                _ = v15
            }
        case 4294967295:
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
        default:
            fmt.Println(v1, v12, v5)
        }
        _ = v12
        _ = v13
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v17 int8 = 127
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 221
    var v19 string = "hello"
    if (!v4) {
        var v20 float32 = -2.25
        for fi21 := 0; fi21 < 15; fi21++ {
            switch (fi21 | (fi21 >> ((1) & 63))) {
            case 4294967295:
                var v22 int16 = int16(fi21)
                _ = v22
            case 32768:
                switch (-(fi21)) {
                case 32768:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 32767:
                    if ((3.14159 - (-2.25 - ((-2.25 - 3.14159) - 3.14159))) != 0.0) {
                        if (42 < -1) {
                            v2 += (v2 / ((2) | 1))
                            var v23 int8 = 0
                            _ = len(v3)
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v5 += 65535
                            v25 := GenS1{F0: 2, F1: "hello", F2: -2147483648}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi24
                            _ = v25
                        }
                        v1 = "a"
                    }
                case 42:
                    v20 += v20
                default:
                    v26 := GenS4{F0: 2, F1: -9223372036854775808, F2: (^(v5))}
                    _ = v26
                }
            case 127:
                v27 := GenS2{F0: v18, F1: 42, F2: 42}
                _ = v27
            }
            _ = fi21
        }
        _ = v20
    }
    if v4 {
        if v4 {
            for fi28 := 0; fi28 < 2; fi28++ {
                v5 = 2
                if (v4 || v4) {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v30 := GenS5{F0: (fi28 & -118), F1: v4, F2: -1}
                        if v4 {
                            switch "foo bar" {
                            case "a":
                                if v4 {
                                    v1 = v1
                                    fmt.Println(v1, v18, v4)
                                    var v31 int64 = -1
                                    switch "hello" {
                                    case "foo bar":
                                        v31 += (v31 - int64(v2))
                                    default:
                                        v5 += (v2 % ((65535) | 1))
                                    }
                                    _ = v31
                                }
                            case "hello":
                                fmt.Println(v19)
                            default:
                                fmt.Println(fi28, fi28)
                            }
                        }
                        _ = fi29
                        _ = v30
                    }
                    fmt.Println(v18, fi28, v2)
                    for fi32 := 0; fi32 < 20; fi32++ {
                        if false {
                            switch ((v1 + v1) + (v1 + v1)) {
                            case "a":
                                if v4 {
                                    var v33 uint16 = v5
                                    switch fi32 {
                                    case -984:
                                        fmt.Println(fi32, v5)
                                    }
                                    v34 := GenS6{F0: 18446744073709551615, F1: 0}
                                    _ = v33
                                    _ = v34
                                } else {
                                    v3 = append(v3, fi32)
                                    v35 := GenS4{F0: 2, F1: -9223372036854775808, F2: v2}
                                    _ = v35
                                }
                            case "\n\t":
                                var v36 int = (fi28 / ((205) | 1))
                                _ = v36
                            case "":
                                var v37 int16 = -1
                                _ = v37
                            }
                            v19 = "\n\t"
                            fmt.Println(v2, v4, v5)
                        }
                        _ = fi32
                    }
                }
                _ = fi28
            }
            switch 32767 {
            case 2147483647:
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 bool = v4
                    _ = fi38
                    _ = v39
                }
            case 4294967295:
                v40 := GenS1{F0: 0, F1: (v19 + "a"), F2: 0}
                _ = v40
            case 267:
                for fi41 := 0; fi41 < 2; fi41++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v18 += (v18 * v18)
                    _ = fi41
                }
            }
        } else {
            var v42 uint = 18446744073709551615
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        var v44 uint64 = 399
        var v45 uint8 = v18
        switch ((v1 + v1) + "") {
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            for fi46 := 0; fi46 < 17; fi46++ {
                fmt.Println(v18)
                _ = fi46
            }
        case "a":
            _ = len(v3)
        default:
            v44 = v44
        }
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
