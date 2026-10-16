package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
type GenS4 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int8; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 int64; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 127
    var v3 uint8 = (v2 >> ((42) & 63))
    var v4 int = 4294967295
    var v5 int16 = 42
    var v6 int = (v4 >> ((18446744073709551615) & 63))
    v7 := GenS1{F0: 4294967295, F1: 42, F2: 4294967295}
    if true {
        fmt.Println(v4, v3, v6)
    } else {
        v8 := []int{42, (v6 << ((255) & 63)), v6, v4}
        v9 := GenS1{F0: 255, F1: -369, F2: 1}
        _ = v8
        _ = v9
    }
    var v10 int64 = 0
    v6 -= v4
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v4)
            _ = fi12
        }
        if true {
            var v13 uint = 65535
            _ = v13
        } else {
            delete(v1, 981)
            if (-3 <= (v4 & (v4 << ((65535) & 63)))) {
                var v14 int16 = 1000
                _ = v14
            } else {
                switch ("a" + "foo bar") {
                case "foo bar":
                    v3 -= v2
                default:
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v4 += (v4 ^ 0)
                        var v16 int16 = v5
                        _ = fi15
                        _ = v16
                    }
                }
                v1[695] = 2
            }
            v5 = v5
        }
        var v17 bool = true
        _ = fi11
        _ = v17
    }
    v18 := GenS1{F0: 255, F1: (v10 | v10), F2: 262}
    v19 := GenS1{F0: 2, F1: 1, F2: 572}
    fmt.Println(v3, v2)
    if (4294967295 > 555) {
        if false {
            var v20 int16 = v5
            v20 += 1000
            v5 = (v20 * 541)
            _ = v20
        } else {
            v5 += v5
            v21 := v1[4294967295]
            _ = v21
        }
    } else {
        if (!((float64(0.1) + 1e10) < 0.1)) {
            fmt.Println(v5, v4)
        }
        switch v4 {
        case 1:
            var v22 int32 = int32(v3)
            _ = v22
        default:
            v23 := GenS1{F0: 42, F1: -1, F2: 106}
            _ = v23
        }
        v3 -= (v2 - 1)
    }
    switch ("a" + "a") {
    case "foo bar":
        v5 -= (v5 << ((743) & 63))
    case "hello":
        var v24 int64 = -723
        _ = v24
    case "":
        v3 = v3
    }
    v25 := []int{v6}
    v26 := GenS1{F0: 65535, F1: (v10 | v10), F2: 1}
    for fi27 := 0; fi27 < 8; fi27++ {
        if true {
            if (21 <= 587) {
                for fi28 := 0; fi28 < 10; fi28++ {
                    v29 := GenS1{F0: 1, F1: 0, F2: 2147483647}
                    fmt.Println(fi27, fi28)
                    var v30 int = (-(fi28))
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v6 -= v4
                v31 := GenS2{F0: ""}
                _ = v31
            } else {
                for fi32 := 0; fi32 < 3; fi32++ {
                    if false {
                        switch "hello" {
                        case "\n\t":
                            switch "a" {
                            case "\n\t":
                                if true {
                                    v3 -= 0
                                    var v33 int32 = -1
                                    _ = v33
                                } else {
                                    if ((-(v5)) >= (v5 % ((v5) | 1))) {
                                        var v34 uint = 18446744073709551615
                                        delete(v1, 65535)
                                        v35 := GenS1{F0: v34, F1: v10, F2: 0}
                                        var v36 float32 = -2.25
                                        var v37 string = "\n\t"
                                        switch (v37 + v37) {
                                        case "":
                                            if true {
                                                v38 := GenS3{F0: -1, F1: 123, F2: (128 > 18)}
                                                switch 722 {
                                                case 106:
                                                    v25 = append(v25, v6)
                                                case -2147483648:
                                                    if true {
                                                        v37 = "hello"
                                                    }
                                                default:
                                                    if true {
                                                        fmt.Println(v4)
                                                        v4 = (fi32 | fi27)
                                                        var v39 float64 = 1.5
                                                        _ = v39
                                                    }
                                                }
                                                var v40 bool = (!(!true))
                                                _ = v38
                                                _ = v40
                                            } else {
                                                if false {
                                                    delete(v1, 2)
                                                } else {
                                                    v10 = -556
                                                    var v41 uint64 = 590
                                                    _ = v41
                                                }
                                                var v42 float64 = float64(1e10)
                                                _ = v42
                                            }
                                        case "\n\t":
                                            v3 = (v3 << ((v34) & 63))
                                            v43 := GenS2{F0: "foo bar"}
                                            v25 = append(v25, (v6 << ((v34) & 63)))
                                            v44 := GenS4{F0: v5, F1: 4294967295, F2: 29}
                                            _ = v43
                                            _ = v44
                                        case "a":
                                            v1[2] = 2147483647
                                            v1[199] = 2
                                        }
                                        if false {
                                            v2 = 255
                                            fmt.Println(fi27, v4)
                                        } else {
                                            v25 = append(v25, (v4 & v4))
                                            var v45 uint16 = 2
                                            _ = v45
                                        }
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    var v46 int8 = -24
                                    if false {
                                        delete(v1, 2147483648)
                                    }
                                    _ = v46
                                }
                            default:
                                var v47 float64 = 3.14159
                                _ = v47
                            }
                        case "":
                            if (-1 <= (^(v6))) {
                                v6 = (-(v6))
                                switch "a" {
                                case "\n\t":
                                    v4 -= (v6 + v6)
                                default:
                                    v48 := GenS1{F0: 255, F1: v10, F2: 4294967295}
                                    _ = v48
                                }
                                var v49 int16 = 32767
                                v50 := GenS3{F0: -1, F1: 42, F2: true}
                                _ = v49
                                _ = v50
                            } else {
                                v6 -= fi32
                                if (-617 >= 351) {
                                    v51 := GenS3{F0: 761, F1: 251, F2: (!false)}
                                    if true {
                                        if false {
                                            v4 -= (-(fi32))
                                        }
                                        v52 := GenS5{F0: 2, F1: 1, F2: 255}
                                        fmt.Println(v4, v3, v5)
                                        if (669 > 255) {
                                            v6 -= v4
                                        } else {
                                            v6 = (fi32 >> ((2) & 63))
                                            v53 := GenS3{F0: 1, F1: 65535, F2: (!((2 < 685) && true))}
                                            _ = v53
                                        }
                                        _ = v52
                                    }
                                    _ = v51
                                } else {
                                    var v54 uint32 = 1
                                    if ((457 > 99) && (v10 > 53)) {
                                        var v55 uint64 = 738
                                        _ = v55
                                    }
                                    _ = v54
                                }
                                v56 := GenS6{F0: 149}
                                _ = v56
                            }
                        case "hello":
                            var v57 int32 = 2147483647
                            _ = v57
                        }
                        v58 := GenS7{F0: 666, F1: (v10 * -9223372036854775808), F2: v2}
                        _ = v58
                    } else {
                        switch (v4 - v6) {
                        case 4294967295:
                            fmt.Println(fi32)
                        default:
                            v3 = v2
                        }
                        _ = len(v1)
                        v59 := GenS1{F0: 255, F1: (v10 * (v10 ^ (v10 - (v10 | (v10 + (v10 >> ((395) & 63))))))), F2: 376}
                        _ = v59
                    }
                    _ = fi32
                }
            }
            var v60 string = ("" + "a")
            _ = v60
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
