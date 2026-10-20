package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 int; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 606
    var v3 uint64 = 799
    var v4 int64 = -9223372036854775808
    v5 := make(map[int]string)
    fmt.Println(v2)
    var v6 int8 = -128
    var v7 int = 0
    var v8 string = "a"
    fmt.Println(v6)
    var v9 int32 = 1
    for fi10 := 0; fi10 < 10; fi10++ {
        v9 -= v9
        fmt.Println(v8, v8)
        _ = fi10
    }
    if (!(!(v8 != v8))) {
        fmt.Println(v3, v8)
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := []int{v7}
            v13 := GenS1{F0: 2}
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint8 = 2
                if (v4 == v4) {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    v16 := GenS2{F0: 42}
                    if true {
                        fmt.Println(v2, v15, fi11)
                    } else {
                        v17 := GenS3{F0: -2147483648, F1: (v2 * 2147483647)}
                        v7 = fi11
                        _ = v17
                    }
                    _ = len(v12)
                    v7 = int(v9)
                    _ = v16
                } else {
                    fmt.Println(v3, v15)
                }
                if true {
                    v4 = v4
                    v4 += (v4 & v4)
                }
                v18 := GenS4{F0: 2147483647}
                var v19 bool = false
                v19 = ((!(v9 != 564)) && true)
                v5[int(v15)] = (v8 + ((v8 + (("" + v8) + (v8 + "a"))) + "a"))
                _ = fi14
                _ = v15
                _ = v18
                _ = v19
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := GenS5{F0: v3, F1: (953 <= 32767), F2: 2}
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 float32 = float32(float64(0.0))
                switch "" {
                case "foo bar":
                    switch -2147483648 {
                    case 32768:
                        switch ("\n\t" + "a") {
                        case "a":
                            v24 := GenS6{F0: v3, F1: v3}
                            _ = v24
                        case "foo bar":
                            fmt.Println(v7, fi20, v2)
                        case "":
                            _ = len(v5)
                        }
                    }
                case "":
                    v6 = v6
                case "\n\t":
                    v23 -= float32(((3.14159 * (((0.0 / (((0.1 + 1.5) * 1.5) + float64(v23))) + 0.1) + 0.1)) + 0.1))
                    if (!(false && ((true || true) || (1 <= (v2 & (v2 | 624)))))) {
                        v23 = v23
                    }
                    var v25 uint8 = 2
                    if false {
                        v26 := []int{}
                        var v27 string = v8
                        _ = v26
                        _ = v27
                    }
                    _ = v25
                default:
                    v28 := GenS7{F0: (v8 + ("a" + ("" + v8)))}
                    _ = v28
                }
                fmt.Println(v4, v6)
                _ = fi22
                _ = v23
            }
            _ = fi20
            _ = v21
        }
    } else {
        if false {
            for fi29 := 0; fi29 < 12; fi29++ {
                for fi30 := 0; fi30 < 15; fi30++ {
                    v31 := GenS8{F0: fi29, F1: (v3 >> ((18446744073709551615) & 63))}
                    if ((true || (42 == 42)) && true) {
                        if false {
                            if false {
                                v8 = v8
                            } else {
                                if (v4 == (v4 ^ -965)) {
                                    _ = len(v5)
                                }
                            }
                            var v32 float32 = ((1e10 - float32(1e10)) + 3.14159)
                            var v33 uint32 = (v2 - v2)
                            _ = v32
                            _ = v33
                        } else {
                            var v34 uint32 = 4294967295
                            var v35 float64 = (1e10 - (1e10 * 3.14159))
                            v36 := GenS2{F0: 30}
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                    } else {
                        var v37 bool = (((((3.14159 * 0.1) * 1.5) - -2.25) * 0.0) >= (float64(1e10) - 0.0))
                        v8 = v8
                        _ = v37
                    }
                    _ = fi30
                    _ = v31
                }
                var v38 int = v7
                _ = fi29
                _ = v38
            }
        } else {
            v3 += v3
            for fi39 := 0; fi39 < 13; fi39++ {
                v4 += (v4 + -1)
                var v40 int8 = -128
                switch v8 {
                case "hello":
                    v8 = v8
                }
                _ = fi39
                _ = v40
            }
            var v41 string = v8
            _ = v41
        }
    }
    fmt.Println(v2, v6)
    if (!(!false)) {
        if true {
            var v42 int8 = (v6 - v6)
            for fi43 := 0; fi43 < 14; fi43++ {
                var v44 uint32 = 2
                for fi45 := 0; fi45 < 7; fi45++ {
                    var v46 int16 = 0
                    if true {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v9, fi43)
                        }
                    } else {
                        v47 := GenS9{F0: (fi43 + v7), F1: v2, F2: 594}
                        switch (^(v7)) {
                        case -2147483648:
                            if true {
                                fmt.Println(fi45, v8)
                                fmt.Println(v9, v2, v9)
                                v48 := []int{int(v3), (fi45 / (((^(fi45))) | 1))}
                                _ = v48
                            } else {
                                var v49 float64 = ((-2.25 - 1.5) * 3.14159)
                                v9 = v9
                                v50 := GenS1{F0: 65535}
                                _ = v49
                                _ = v50
                            }
                        default:
                            var v51 int16 = v46
                            _ = v51
                        }
                        v52 := GenS6{F0: uint64(v2), F1: v3}
                        _ = v47
                        _ = v52
                    }
                    _ = fi45
                    _ = v46
                }
                fmt.Println(v42)
                _ = fi43
                _ = v44
            }
            switch (v8 + v8) {
            case "hello":
                v9 = 42
            default:
                v53 := GenS9{F0: (v7 & 128), F1: (^(v2)), F2: 255}
                _ = v53
            }
            _ = v42
        } else {
            v54 := GenS4{F0: (v2 + v2)}
            var v55 int16 = -1
            fmt.Println(v7, v3)
            _ = v54
            _ = v55
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
