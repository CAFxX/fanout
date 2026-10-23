package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int8; F1 uint16; F2 int64 }

func main() {
    var v1 int32 = -1
    var v2 int32 = (v1 & v1)
    var v3 int = -407
    v4 := []int{(-(v3))}
    v5 := make(map[int]string)
    v6 := make(map[int]int)
    fmt.Println(v1, v3)
    switch "a" {
    case "\n\t":
        v7 := v6[(-(v3))]
        _ = v7
    case "hello":
        v8 := GenS1{F0: 1}
        _ = v8
    }
    var v9 int16 = 813
    var v10 int8 = -110
    var v11 uint8 = 0
    for fi12 := 0; fi12 < 5; fi12++ {
        _ = len(v4)
        switch v3 {
        case 172:
            for fi13 := 0; fi13 < 14; fi13++ {
                if (5 <= v2) {
                    v14 := GenS2{F0: 65535, F1: int16(v11)}
                    var v15 uint8 = 128
                    var v16 uint16 = 2
                    if true {
                        if true {
                            _ = len(v6)
                        } else {
                            _ = len(v4)
                            v15 = (v11 / ((v11) | 1))
                            v17 := GenS2{F0: 906, F1: v9}
                            _ = v17
                        }
                    }
                    _ = v14
                    _ = v15
                    _ = v16
                }
                var v18 float64 = (0.1 * 0.0)
                _ = fi13
                _ = v18
            }
        case -9223372036854775808:
            if true {
                for fi19 := 0; fi19 < 4; fi19++ {
                    if false {
                        fmt.Println(v9)
                        v20 := GenS3{F0: (-485 != -9223372036854775808), F1: v11}
                        fmt.Println(v2, fi12)
                        _ = v20
                    }
                    v2 = v2
                    _ = fi19
                }
                fmt.Println(v2, v11, v11)
            }
        case 32768:
            _ = len(v5)
        }
        var v21 uint64 = 18446744073709551615
        _ = fi12
        _ = v21
    }
    v11 = (v11 + 1)
    if ((-1 == -1) || true) {
        v22 := GenS4{F0: (false || true), F1: 2147483647}
        switch (v3 >> ((1) & 63)) {
        case -322:
            for fi23 := 0; fi23 < 13; fi23++ {
                if false {
                    var v24 uint = 18446744073709551615
                    v25 := GenS1{F0: int64(v10)}
                    if false {
                        v26 := GenS4{F0: (!(v2 == (^(v1)))), F1: 42}
                        _ = v26
                    } else {
                        var v27 uint64 = 981
                        v1 = v1
                        var v28 float64 = (0.0 + 1.5)
                        _ = v27
                        _ = v28
                    }
                    _ = v24
                    _ = v25
                } else {
                    v9 -= -455
                    delete(v5, 42)
                }
                _ = fi23
            }
        case 9223372036854775807:
            if false {
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 uint64 = 758
                    var v31 uint64 = v30
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                fmt.Println(v9)
            }
        default:
            var v32 uint8 = v11
            _ = v32
        }
        _ = v22
    }
    v33 := GenS2{F0: 42, F1: (v9 ^ (v9 % (((v9 - v9)) | 1)))}
    var v34 uint32 = 65535
    var v35 int16 = (v9 + (v9 % (((v9 + (v9 ^ (^(v9))))) | 1)))
    v9 += -32768
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 uint8 = v11
        _ = fi36
        _ = v37
    }
    switch ("" + "hello") {
    case "":
        switch "hello" {
        case "a":
            v11 = v11
        case "hello":
            for fi38 := 0; fi38 < 15; fi38++ {
                if ((-(v10)) >= v10) {
                    var v39 bool = false
                    if (true || (!v39)) {
                        switch 0 {
                        case 42:
                            var v40 int = 167
                            _ = v40
                        case -869:
                            v35 = (v35 + -1)
                        case 127:
                            v5[831] = "foo bar"
                        default:
                            fmt.Println(v35, v10)
                        }
                        for fi41 := 0; fi41 < 4; fi41++ {
                            if (!v39) {
                                v39 = v39
                            } else {
                                var v42 uint32 = (v34 - 949)
                                fmt.Println(v11, v11, v34)
                                var v43 uint8 = v11
                                var v44 string = "hello"
                                v5[-90] = ((v44 + v44) + v44)
                                var v45 int8 = v10
                                _ = v42
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            var v46 float64 = -2.25
                            _ = fi41
                            _ = v46
                        }
                    } else {
                        v47 := GenS2{F0: 1, F1: v35}
                        if v39 {
                            var v48 int8 = (^(v10))
                            _ = v48
                        } else {
                            v49 := GenS4{F0: (float64(1.5) != (1e10 - 0.1)), F1: (v1 ^ (v2 | -2147483648))}
                            _ = v49
                        }
                        fmt.Println(v11)
                        _ = v47
                    }
                    if (v11 != 0) {
                        v35 += -32768
                    }
                    _ = v39
                }
                v50 := GenS2{F0: 42, F1: (v35 | 0)}
                switch 0 {
                case 946:
                    if (32767 >= 966) {
                        fmt.Println(fi38)
                        if true {
                            v35 = (v9 << ((1) & 63))
                        } else {
                            for fi51 := 0; fi51 < 2; fi51++ {
                                v52 := GenS4{F0: false, F1: (v2 << ((0) & 63))}
                                fmt.Println(v11)
                                _ = fi51
                                _ = v52
                            }
                        }
                    } else {
                        var v53 int64 = 42
                        _ = v53
                    }
                case 4294967295:
                    if (v11 >= 127) {
                        var v54 float64 = 1.5
                        _ = v54
                    } else {
                        _ = len(v5)
                        switch "foo bar" {
                        case "":
                            var v55 uint32 = v34
                            _ = v55
                        case "\n\t":
                            v56 := GenS5{F0: -138, F1: 0}
                            _ = v56
                        default:
                            v57 := GenS5{F0: -169, F1: v34}
                            _ = v57
                        }
                        var v58 uint8 = 128
                        _ = v58
                    }
                case 357:
                    v9 += (v35 - 1)
                }
                _ = fi38
                _ = v50
            }
        }
    default:
        v59 := GenS6{F0: v10, F1: 65535, F2: 0}
        _ = v59
    }
    var v60 int32 = (v1 + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v9
}
