package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int }
func genFn0(p0 string, p1 uint, p2 uint8) uint {
    if (42 != 18446744073709551615) {
        switch ("foo bar" + ("a" + p0)) {
        case "hello":
            if false {
                if true {
                    var v5 int64 = -9223372036854775808
                    _ = v5
                } else {
                    var v6 int16 = 465
                    if true {
                        var v7 uint32 = 2
                        var v8 uint8 = (^(p2))
                        switch -620 {
                        case -1:
                            switch p0 {
                            case "\n\t":
                                v9 := GenS1{F0: 18446744073709551615, F1: 1}
                                _ = v9
                            case "a":
                                v10 := GenS1{F0: 741, F1: 42}
                                _ = v10
                            default:
                                fmt.Println(p1, v6, p1)
                            }
                        case 127:
                            for fi11 := 0; fi11 < 19; fi11++ {
                                var v12 bool = true
                                _ = fi11
                                _ = v12
                            }
                        case -787:
                            if false {
                                p0 = p0
                                fmt.Println(v6)
                                var v13 uint64 = 367
                                _ = v13
                            }
                        }
                        _ = v7
                        _ = v8
                    }
                    var v14 uint = 65535
                    _ = v14
                    _ = v6
                }
            }
        case "\n\t":
            v15 := GenS2{F0: -976, F1: 32767}
            _ = v15
        case "":
            if false {
                fmt.Println(p0, p1, p1)
                v16 := []int{-2147483648, 576, 42}
                _ = v16
            } else {
                v17 := []int{}
                var v18 string = "\n\t"
                v18 = "a"
                _ = v17
                _ = v18
            }
        default:
            var v19 uint64 = 4294967295
            _ = v19
        }
        v20 := []int{32767, -742, 127}
        _ = v20
    } else {
        v21 := []int{32767, 9223372036854775807}
        switch p0 {
        case "foo bar":
            var v22 string = ("foo bar" + ("a" + ("a" + "foo bar")))
            _ = v22
        case "a":
            var v23 uint16 = 2
            _ = v23
        default:
            v21 = append(v21, -171)
        }
        p0 = p0
        _ = v21
    }
    if false {
        var v24 int16 = 42
        var v25 float32 = 3.14159
        _ = v24
        _ = v25
    }
    v26 := GenS2{F0: 0, F1: 255}
    if (false || (-501 > 701)) {
        v27 := []int{-9223372036854775808, 184, 214, 929}
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 1
}
func genFn1(p0 int32, p1 uint64) uint {
    p0 += p0
    p1 -= 2
    for fi28 := 0; fi28 < 9; fi28++ {
        if ((("" + "\n\t") + ("hello" + ("foo bar" + "a"))) < "a") {
            var v29 uint32 = 1
            v30 := GenS3{F0: -1, F1: 65535, F2: 255}
            v31 := GenS3{F0: 0, F1: 1, F2: 1}
            _ = v29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint = 255
    var v2 uint16 = 518
    var v3 int16 = -1
    var v4 uint8 = 129
    if (v2 > v2) {
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := []int{fi32, (^(fi32)), fi32, (fi32 | fi32)}
            var v34 int64 = -254
            if false {
                var v35 uint32 = 2147483647
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
                v4 = v4
                _ = v35
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        var v36 int64 = 6
        v37 := []int{863, 32767, 32767}
        _ = v36
        _ = v37
    }
    v38 := GenS4{F0: 42}
    var v39 int16 = v3
    if true {
        v3 += v39
    }
    var v40 uint64 = 536
    v2 -= (v2 & v2)
    v41 := GenS4{F0: 1}
    var v42 float64 = (1e10 / float64(1e10))
    if true {
        if true {
            var v43 uint64 = v40
            v42 = (1.5 - ((v42 / v42) - v42))
            var v44 int64 = 1
            _ = v43
            _ = v44
        } else {
            var v45 string = "foo bar"
            switch 1 {
            case 128:
                v40 -= (v40 >> (((v1 % ((42) | 1))) & 63))
            default:
                var v46 uint8 = (v4 & uint8(v40))
                _ = v46
            }
            _ = v45
        }
    } else {
        var v47 int8 = 126
        _ = v47
    }
    for fi48 := 0; fi48 < 7; fi48++ {
        if (3.14159 == v42) {
            v2 = 255
            for fi49 := 0; fi49 < 12; fi49++ {
                v3 += v39
                _ = fi49
            }
            var v50 float64 = (((v42 - (1e10 * v42)) * (((1e10 / v42) - 1.5) / v42)) / (v42 * v42))
            _ = v50
        } else {
            v51 := GenS3{F0: 76, F1: v2, F2: 983}
            var v52 uint8 = 42
            var v53 int16 = v3
            _ = v51
            _ = v52
            _ = v53
        }
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
