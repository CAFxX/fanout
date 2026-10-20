package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint16 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 uint64; F2 int32 }
func genFn0(p0 int16) uint64 {
    if false {
        var v6 uint8 = 119
        v7 := GenS1{F0: 32767, F1: ("hello" + ""), F2: 2}
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{fi8}
            var v10 int64 = 581
            fmt.Println(v10, v6)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v6
        _ = v7
    } else {
        fmt.Println(p0)
        var v11 uint64 = 1
        _ = v11
    }
    v12 := GenS2{F0: 0, F1: 42}
    p0 = 1000
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint8 = 1
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := []int{-9223372036854775808, (fi15 / ((fi13) | 1)), (fi15 << ((375) & 63))}
            _ = fi15
            _ = v16
        }
        v14 = (v14 * 182)
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v12
    return 0
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 float32 = v1
    var v5 float64 = (float64(3.14159) * -2.25)
    var v17 uint64 = 331
    var v18 int64 = 1
    switch -371 {
    case -2147483648:
        fmt.Println(v18, v17)
    default:
        if (2147483647 <= -295) {
            if true {
                v19 := GenS3{F0: v17, F1: (!true)}
                v18 = 1
                switch 461 {
                case 1:
                    v1 -= float32(float64(v1))
                case 42:
                    v20 := GenS4{F0: 127}
                    _ = v20
                case 548:
                    v21 := []int{-1, 927}
                    _ = v21
                default:
                    switch "\n\t" {
                    case "hello":
                        var v22 string = "\n\t"
                        _ = v22
                    case "a":
                        var v23 string = ("" + "a")
                        _ = v23
                    case "":
                        if ((v17 ^ (v17 >> ((4294967295) & 63))) <= genFn0(1000)) {
                            if false {
                                v24 := GenS5{F0: 32767, F1: 2}
                                fmt.Println(v4)
                                _ = v24
                            } else {
                                var v25 int = -1
                                if true {
                                    var v26 uint = 18446744073709551615
                                    if false {
                                        delete(v2, "")
                                    } else {
                                        fmt.Println(v5, v25)
                                        v4 -= v1
                                        v1 = v4
                                    }
                                    v27 := GenS6{F0: 165}
                                    _ = v26
                                    _ = v27
                                }
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v4, v25)
                                case "":
                                    var v28 uint64 = v17
                                    _ = v28
                                case "a":
                                    if (1e10 <= 0.1) {
                                        var v29 int16 = 461
                                        var v30 bool = (-1 < 127)
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                _ = v25
                            }
                        }
                    }
                }
                _ = v19
            } else {
                v31 := GenS4{F0: 42}
                var v32 int8 = 42
                _ = v31
                _ = v32
            }
        }
    }
    var v33 uint64 = v17
    fmt.Println(v1, v5, v1)
    var v34 int64 = int64(v33)
    v2["hello"] = 9223372036854775807
    var v35 uint64 = v33
    v36 := v2[("foo bar" + "hello")]
    for fi37 := 0; fi37 < 19; fi37++ {
        if (false && (!true)) {
            v38 := GenS7{F0: v36}
            _ = v38
        }
        _ = fi37
    }
    if (32767 == 928) {
        switch v36 {
        case 9223372036854775807:
            v39 := GenS6{F0: 0}
            _ = v39
        }
        var v40 int8 = -1
        if false {
            var v41 int16 = -787
            _ = len(v3)
            v42 := v3["\n\t"]
            _ = v41
            _ = v42
        } else {
            v43 := []int{(v36 + v36), 2147483647}
            if ((1 != 1) && (v1 < -2.25)) {
                delete(v2, ("foo bar" + ""))
            } else {
                var v44 uint8 = 0
                switch "a" {
                case "\n\t":
                    _ = len(v2)
                default:
                    v45 := GenS3{F0: (v35 & v35), F1: false}
                    _ = v45
                }
                v4 = v4
                _ = v44
            }
            v1 -= (v4 + v1)
            _ = v43
        }
        _ = v40
    } else {
        if true {
            var v46 int16 = -1
            var v47 int64 = (v34 >> ((18446744073709551615) & 63))
            _ = v46
            _ = v47
        } else {
            v48 := []int{(-(v36)), (v36 | 128), v36, v36}
            if true {
                fmt.Println(v34)
            } else {
                fmt.Println(v1, v33, v36)
            }
            _ = v48
        }
        var v49 uint32 = 2147483648
        _ = v49
    }
    switch v36 {
    case 32768:
        var v50 bool = (true || (false || (((!true) || (v36 < (v36 % ((v36) | 1)))) && false)))
        _ = v50
    case 1:
        if (false || true) {
            v51 := []int{}
            _ = v51
        } else {
            v52 := []int{v36, v36, (^(v36)), v36}
            for fi53 := 0; fi53 < 2; fi53++ {
                v3[""] = v36
                for fi54 := 0; fi54 < 18; fi54++ {
                    var v55 uint16 = 0
                    v56 := GenS2{F0: v18, F1: -128}
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                _ = fi53
            }
            _ = v52
        }
    default:
        if true {
            v57 := GenS8{F0: 571}
            var v58 float64 = (v5 / v5)
            _ = v57
            _ = v58
        } else {
            for fi59 := 0; fi59 < 4; fi59++ {
                v5 -= (((v5 - 3.14159) * (v5 * 1.5)) / -2.25)
                for fi60 := 0; fi60 < 5; fi60++ {
                    var v61 uint8 = 115
                    v62 := GenS9{F0: v35, F1: uint64(v61), F2: 42}
                    v61 -= (v61 & 128)
                    _ = fi60
                    _ = v61
                    _ = v62
                }
                v34 = (v34 & (v34 ^ v18))
                _ = fi59
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
