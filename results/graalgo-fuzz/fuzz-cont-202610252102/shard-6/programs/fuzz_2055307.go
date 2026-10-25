package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
func genFn0() uint64 {
    v4 := GenS1{F0: 42}
    v5 := []int{}
    var v6 uint = 1
    var v7 float32 = 3.14159
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1() uint32 {
    if true {
        var v8 uint8 = 2
        v9 := []int{0, -1}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v8
        _ = v9
    }
    if false {
        if (1 != 478) {
            var v10 int64 = 42
            v10 = (v10 % ((v10) | 1))
            _ = v10
        }
        if (9223372036854775807 > 9223372036854775807) {
            var v11 int64 = 9223372036854775807
            _ = v11
        } else {
            if true {
                var v12 uint64 = 18446744073709551615
                var v13 uint32 = 42
                switch 128 {
                case -2147483648:
                    var v14 uint8 = 42
                    _ = v14
                case -526:
                    var v15 int16 = 1000
                    _ = v15
                case 32768:
                    v13 = (^(v13))
                default:
                    if (((v12 & (v12 | genFn0())) >= 18446744073709551615) || false) {
                        fmt.Println(v13, v12)
                    }
                }
                _ = v12
                _ = v13
            }
            if (367 > 128) {
                var v16 float64 = float64((-2.25 / -2.25))
                v16 = ((((((1.5 + 0.1) + (1.5 + -2.25)) - v16) * (v16 + (0.1 + (1e10 * (0.0 / v16))))) / v16) + (v16 / 1.5))
                _ = v16
            }
            var v17 int = 0
            _ = v17
        }
        var v18 uint = 806
        _ = v18
    } else {
        v19 := GenS1{F0: 9223372036854775807}
        _ = v19
    }
    if false {
        v20 := []int{-560, 9223372036854775807, 127, 42}
        var v21 int64 = 0
        var v22 float32 = 3.14159
        _ = v20
        _ = v21
        _ = v22
    }
    return 2147483647
}

func main() {
    var v1 float32 = (0.0 / (((float32((1.5 + 3.14159)) + (1e10 + 1e10)) * 1e10) - 1.5))
    var v2 uint = 255
    var v3 float64 = -2.25
    var v23 uint32 = 65535
    fmt.Println(v2, v1, v1)
    switch "" {
    case "hello":
        v23 = (v23 | genFn1())
    case "\n\t":
        v24 := []int{885, 1}
        _ = v24
    }
    v23 += v23
    if true {
        for fi25 := 0; fi25 < 20; fi25++ {
            v2 -= (v2 >> (((^(v2))) & 63))
            _ = fi25
        }
        v26 := GenS1{F0: 42}
        v27 := GenS1{F0: 0}
        _ = v26
        _ = v27
    } else {
        v28 := GenS1{F0: 0}
        for fi29 := 0; fi29 < 2; fi29++ {
            if (-1 > 0) {
                var v30 uint8 = 127
                _ = v30
            } else {
                v31 := []int{}
                _ = v31
            }
            v32 := []int{fi29}
            v2 += (^(v2))
            _ = fi29
            _ = v32
        }
        _ = v28
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            switch ("a" + "") {
            case "foo bar":
                v2 = 425
            case "a":
                switch 32768 {
                case 32768:
                    v2 -= v2
                case 647:
                    v1 = v1
                case 991:
                    if false {
                        switch fi33 {
                        case 179:
                            v1 = v1
                        case 42:
                            if false {
                                var v35 int8 = -1
                                fmt.Println(v1)
                                _ = v35
                            } else {
                                if (false || ((v1 + float32(1e10)) > v1)) {
                                    fmt.Println(v2)
                                } else {
                                    var v36 int16 = 1
                                    v37 := []int{(fi33 >> ((uint(v23)) & 63))}
                                    if (!true) {
                                        if (fi34 >= fi34) {
                                            var v38 string = (("foo bar" + "foo bar") + "foo bar")
                                            _ = v38
                                        }
                                        var v39 int32 = 1
                                        v37 = append(v37, (fi34 >> (((v2 * (v2 ^ 18446744073709551615))) & 63)))
                                        _ = v39
                                    }
                                    _ = v36
                                    _ = v37
                                }
                            }
                        case -9223372036854775808:
                            switch 2147483647 {
                            case 2147483647:
                                var v40 float64 = v3
                                _ = v40
                            case 0:
                                v41 := []int{}
                                _ = v41
                            case -9223372036854775808:
                                var v42 uint32 = v23
                                _ = v42
                            }
                        default:
                            fmt.Println(fi34, v1, fi34)
                        }
                        v43 := []int{fi34, (^(fi33)), -374}
                        switch "a" {
                        case "hello":
                            v23 -= genFn1()
                        }
                        _ = v43
                    } else {
                        var v44 int16 = 42
                        _ = v44
                    }
                default:
                    v23 -= 2147483647
                }
            }
            v45 := []int{fi33}
            _ = fi34
            _ = v45
        }
        _ = fi33
    }
    if false {
        var v46 string = "\n\t"
        v3 = ((0.0 - v3) - v3)
        switch 32767 {
        case -2147483648:
            for fi47 := 0; fi47 < 16; fi47++ {
                v48 := GenS2{F0: 127}
                _ = fi47
                _ = v48
            }
        case -111:
            for fi49 := 0; fi49 < 4; fi49++ {
                v2 += v2
                v46 = ""
                for fi50 := 0; fi50 < 17; fi50++ {
                    v51 := []int{(^(fi50)), fi50, fi49, 9223372036854775807}
                    var v52 int16 = 32767
                    v2 = (^(v2))
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                _ = fi49
            }
        default:
            var v53 float32 = float32((v3 / (0.1 * 1.5)))
            _ = v53
        }
        _ = v46
    }
    v2 += 866
    if (uint8(v2) > 255) {
        v1 += float32(1e10)
        v54 := GenS3{F0: v2}
        _ = v54
    } else {
        var v55 uint64 = 4
        v56 := []int{}
        var v57 uint = v2
        _ = v55
        _ = v56
        _ = v57
    }
    var v58 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v58
}
