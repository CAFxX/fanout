package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint8; F2 int }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint16; F1 int; F2 uint64 }
type GenS8 struct { F0 int64 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 1
    v3 := []uint32{v1, v1}
    var v4 bool = ((-2.25 / 1.0) < 0.0)
    v5 := []int{4294967295, 9223372036854775807, 4294967295}
    var v6 uint16 = 65535
    switch ("hello" + "") {
    case "":
        if ((964 > (v6 + (v6 ^ (^(v6))))) && ((^(v1)) == (v1 + v1))) {
            v4 = (false && (-2.25 < 1.5))
            v1 += (v1 << ((uint(v6)) & 63))
        } else {
            if v4 {
                v6 = 783
                v3 = append(v3, (^(v1)))
                switch 2147483647 {
                case 4294967295:
                    if len(v5) > 0 {
                        v5[0] = -609
                    }
                }
            } else {
                var v7 bool = v4
                _ = len(v5)
                if false {
                    var v8 int16 = -135
                    var v9 uint = 1
                    _ = v8
                    _ = v9
                } else {
                    v6 -= v6
                    if (!v7) {
                        v2 -= v2
                    }
                }
                _ = v7
            }
            if true {
                var v10 float32 = 3.14159
                v5 = append(v5, -90)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v10
            } else {
                var v11 float64 = 1.5
                switch (("" + "a") + ("" + (("" + ("" + (((("foo bar" + (("a" + "\n\t") + "")) + ("\n\t" + "foo bar")) + ("foo bar" + ("" + ("foo bar" + "a")))) + "hello"))) + "hello"))) {
                case "hello":
                    var v12 bool = false
                    _ = v12
                }
                _ = v11
            }
        }
    case "hello":
        switch 0 {
        case 485:
            var v13 int8 = -128
            _ = v13
        case 128:
            var v14 bool = ((true && false) && false)
            _ = v14
        case -1:
            var v15 uint16 = (^(v6))
            _ = v15
        default:
            v16 := GenS1{F0: 203, F1: 0}
            _ = v16
        }
    }
    fmt.Println(v2)
    if (v4 || (v1 != (v1 / (((^(v1))) | 1)))) {
        switch -2147483648 {
        case 177:
            switch ("foo bar" + "\n\t") {
            case "":
                var v17 uint8 = 1
                _ = v17
            default:
                if (v4 || v4) {
                    if v4 {
                        v18 := GenS1{F0: 68, F1: 0}
                        if ((!(1e10 == 1e10)) || v4) {
                            var v19 int8 = 0
                            v20 := GenS1{F0: (v2 | v2), F1: 1}
                            v2 += v2
                            _ = v19
                            _ = v20
                        } else {
                            switch -876 {
                            case -186:
                                var v21 int = 42
                                _ = v21
                            case 0:
                                v22 := GenS1{F0: (v2 << ((255) & 63)), F1: 0}
                                _ = v22
                            case -1:
                                for fi23 := 0; fi23 < 3; fi23++ {
                                    v24 := GenS1{F0: v2, F1: 18446744073709551615}
                                    if (1 != (v2 | v2)) {
                                        v1 += v1
                                        fmt.Println(v1)
                                    }
                                    var v25 int8 = 42
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                            }
                            v2 += 255
                        }
                        v26 := GenS1{F0: v2, F1: 4294967295}
                        _ = v18
                        _ = v26
                    }
                } else {
                    v1 = v1
                    v27 := GenS1{F0: (^(v2)), F1: 0}
                    if (0.1 > float32(0.1)) {
                        if len(v5) > 0 {
                            v5[0] = 2147483647
                        }
                    } else {
                        v5 = append(v5, -2147483648)
                    }
                    _ = v27
                }
            }
        }
        switch "\n\t" {
        case "":
            if true {
                if len(v5) > 0 {
                    v5[0] = -9223372036854775808
                }
                v4 = v4
                for fi28 := 0; fi28 < 10; fi28++ {
                    var v29 int8 = -128
                    v6 += v6
                    v4 = v4
                    _ = fi28
                    _ = v29
                }
            } else {
                v3 = append(v3, 1)
                switch 32767 {
                case 1:
                    fmt.Println(v2, v6)
                case 0:
                    v30 := GenS2{F0: ("" != ("a" + "hello"))}
                    _ = v30
                default:
                    v2 -= 128
                }
                var v31 int64 = 0
                _ = v31
            }
        case "\n\t":
            v32 := GenS3{F0: -32768}
            _ = v32
        }
        for fi33 := 0; fi33 < 10; fi33++ {
            v34 := GenS4{F0: 497, F1: v2, F2: -943}
            v2 = (v2 - 2)
            _ = fi33
            _ = v34
        }
    } else {
        v6 = (v6 | 2)
        _ = len(v5)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v35 := GenS5{F0: v2, F1: (717 < 18446744073709551615)}
    if v4 {
        var v36 uint8 = v2
        if ("" != "a") {
            for fi37 := 0; fi37 < 20; fi37++ {
                for fi38 := 0; fi38 < 8; fi38++ {
                    if ((!(-2.25 > 3.14159)) || false) {
                        var v39 uint16 = (v6 / ((v6) | 1))
                        switch "a" {
                        case "hello":
                            v40 := GenS1{F0: uint8(v6), F1: 0}
                            _ = v40
                        }
                        v1 -= v1
                        _ = v39
                    } else {
                        if ((2 >= 42) || false) {
                            v41 := GenS6{F0: v36, F1: 127}
                            _ = v41
                        }
                    }
                    fmt.Println(fi37, v1)
                    _ = fi38
                }
                v36 = v2
                if len(v3) > 0 {
                    v3[0] = v1
                }
                _ = fi37
            }
            var v42 int8 = 100
            _ = v42
        }
        _ = v36
    } else {
        var v43 uint16 = 944
        v6 = v43
        _ = v43
    }
    if (!v4) {
        v1 += (^(v1))
        v4 = true
    }
    v44 := GenS7{F0: (v6 / ((v6) | 1)), F1: 32768, F2: 768}
    v1 = (v1 >> ((42) & 63))
    switch 127 {
    case -1:
        fmt.Println(v4, v6, v6)
    }
    if v4 {
        if false {
            v45 := GenS8{F0: 9223372036854775807}
            _ = v45
        } else {
            fmt.Println(v4)
            fmt.Println(v2, v2, v2)
        }
    } else {
        v6 -= (^(v6))
    }
    if true {
        if len(v3) > 0 {
            v3[0] = (v1 % ((uint32(v2)) | 1))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
