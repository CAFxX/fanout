package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 int; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int; F2 uint32 }
type GenS5 struct { F0 int16; F1 int16; F2 int32 }

func main() {
    var v1 uint16 = 206
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 int = -9223372036854775808
    var v5 bool = false
    var v6 int = v4
    v7 := GenS1{F0: 652, F1: v6, F2: 67}
    fmt.Println(v6, v1, v4)
    var v8 uint16 = (v1 - (^(v1)))
    delete(v3, "a")
    v8 += 634
    for fi9 := 0; fi9 < 9; fi9++ {
        v5 = v5
        for fi10 := 0; fi10 < 7; fi10++ {
            v4 += 127
            v6 = fi10
            _ = fi10
        }
        _ = fi9
    }
    switch ("hello" + ("a" + "\n\t")) {
    case "":
        for fi11 := 0; fi11 < 17; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 uint32 = 0
                _ = fi12
                _ = v13
            }
            if v5 {
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := GenS2{F0: -461, F1: 42, F2: 2}
                    _ = fi14
                    _ = v15
                }
            }
            v16 := GenS3{F0: v6, F1: 4294967295, F2: 0}
            _ = fi11
            _ = v16
        }
    }
    if true {
        var v17 float32 = -2.25
        fmt.Println(v17, v4)
        _ = v17
    }
    v4 = v4
    switch (v4 & v4) {
    case -9223372036854775808:
        v18 := []int{v4}
        _ = v18
    case 2147483647:
        var v19 int8 = -1
        _ = v19
    case -812:
        var v20 float64 = (0.1 + 0.0)
        _ = v20
    }
    var v21 uint16 = 601
    switch v4 {
    case 32768:
        var v22 float64 = ((3.14159 / ((3.14159 / 1.0) - 1.5)) + 0.0)
        _ = v22
    case 606:
        var v23 uint64 = 4294967295
        _ = v23
    case 2147483647:
        if (!true) {
            switch 9223372036854775807 {
            case 32767:
                if (918 == 0) {
                    switch v6 {
                    case 4294967295:
                        for fi24 := 0; fi24 < 5; fi24++ {
                            switch 1 {
                            case 899:
                                fmt.Println(v1)
                            default:
                                if (((float32(3.14159) - 0.0) * 1.5) != -2.25) {
                                    switch "hello" {
                                    case "foo bar":
                                        v5 = v5
                                    case "\n\t":
                                        v25 := []int{-606, (^(v6)), 127, (v6 / ((v6) | 1))}
                                        _ = v25
                                    default:
                                        var v26 uint32 = 65535
                                        _ = v26
                                    }
                                    var v27 uint64 = 18446744073709551615
                                    _ = v27
                                }
                            }
                            var v28 float32 = 0.1
                            var v29 uint16 = (^(v1))
                            _ = fi24
                            _ = v28
                            _ = v29
                        }
                    case -105:
                        var v30 uint64 = 163
                        _ = v30
                    case -2147483648:
                        v31 := GenS2{F0: 2147483647, F1: 4294967295, F2: 255}
                        _ = v31
                    default:
                        if (0 != 42) {
                            if v5 {
                                var v32 float64 = 0.0
                                switch (v4 - v4) {
                                case -9223372036854775808:
                                    fmt.Println(v5)
                                case 4294967295:
                                    fmt.Println(v21, v6)
                                }
                                _ = v32
                            } else {
                                v33 := GenS2{F0: 291, F1: 2, F2: 418}
                                var v34 float64 = 1e10
                                _ = v33
                                _ = v34
                            }
                            v35 := GenS2{F0: -2147483648, F1: 2, F2: 0}
                            if v5 {
                                for fi36 := 0; fi36 < 18; fi36++ {
                                    for fi37 := 0; fi37 < 18; fi37++ {
                                        fmt.Println(v4, fi36, v1)
                                        var v38 uint8 = 42
                                        _ = fi37
                                        _ = v38
                                    }
                                    v4 = 32767
                                    if v5 {
                                        for fi39 := 0; fi39 < 2; fi39++ {
                                            v8 += (v21 ^ (v21 + (v1 & (^(v8)))))
                                            delete(v2, (fi39 % ((v4) | 1)))
                                            var v40 uint8 = 1
                                            _ = fi39
                                            _ = v40
                                        }
                                        var v41 float64 = 0.1
                                        var v42 int16 = int16(v21)
                                        _ = v41
                                        _ = v42
                                    } else {
                                        var v43 uint = 431
                                        v44 := GenS4{F0: 127, F1: v4, F2: 310}
                                        v1 = v21
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = fi36
                                }
                                for fi45 := 0; fi45 < 1; fi45++ {
                                    v4 += 128
                                    if (0 <= -193) {
                                        for fi46 := 0; fi46 < 6; fi46++ {
                                            var v47 int8 = int8(v1)
                                            fmt.Println(v21)
                                            v48 := []int{(fi45 << ((18446744073709551615) & 63)), int(v21)}
                                            var v49 uint = 1
                                            _ = fi46
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        }
                                    } else {
                                        switch 32767 {
                                        case 642:
                                            var v50 uint64 = 4294967295
                                            _ = v50
                                        case 948:
                                            v51 := []int{}
                                            _ = v51
                                        case 127:
                                            v6 -= v4
                                        }
                                    }
                                    _ = fi45
                                }
                            }
                            _ = v35
                        } else {
                            v5 = (1 == 18446744073709551615)
                        }
                    }
                }
            case -1:
                switch (v6 / ((-1) | 1)) {
                case 4294967295:
                    v21 = 255
                case -165:
                    v21 = v1
                case -9223372036854775808:
                    if (1 >= 450) {
                        var v52 float32 = 0.0
                        var v53 uint8 = 128
                        v54 := GenS4{F0: (v53 | v53), F1: v6, F2: 535}
                        _ = v52
                        _ = v53
                        _ = v54
                    } else {
                        if v5 {
                            v55 := []int{v4, (v6 - v6)}
                            _ = v55
                        } else {
                            v5 = false
                        }
                        if false {
                            switch "a" {
                            case "a":
                                v6 = v6
                            case "\n\t":
                                fmt.Println(v4, v8)
                            case "hello":
                                v1 = (v8 + v8)
                            default:
                                switch ("hello" + "") {
                                case "hello":
                                    if ("\n\t" <= "hello") {
                                        if (!(-253 <= -32768)) {
                                            v56 := GenS3{F0: 32768, F1: 18446744073709551615, F2: 4294967295}
                                            var v57 int16 = 42
                                            v21 += 32767
                                            _ = v56
                                            _ = v57
                                        }
                                        v58 := GenS5{F0: 42, F1: 0, F2: 42}
                                        fmt.Println(v5)
                                        _ = v58
                                    }
                                case "a":
                                    for fi59 := 0; fi59 < 5; fi59++ {
                                        var v60 uint8 = 255
                                        fmt.Println(v5, v21, v60)
                                        v4 = v6
                                        _ = fi59
                                        _ = v60
                                    }
                                }
                            }
                        } else {
                            var v61 int = v6
                            v62 := []int{v61}
                            _ = v61
                            _ = v62
                        }
                    }
                default:
                    var v63 uint32 = 216
                    _ = v63
                }
            case 42:
                var v64 int32 = 42
                _ = v64
            default:
                var v65 uint32 = 2
                _ = v65
            }
        }
    }
    fmt.Println(v21)
    v66 := []int{v6}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v7
    _ = v8
}
