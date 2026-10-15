package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
func genFn0(p0 uint16) uint32 {
    v6 := []int{int(p0), 854, 2147483647}
    fmt.Println(p0)
    v7 := GenS1{F0: 778, F1: (p0 & p0), F2: 2}
    _ = p0
    _ = v6
    _ = v7
    return 796
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []int{v1, -9223372036854775808, (v1 & v1), v1}
    var v3 uint8 = 187
    var v4 uint = 65535
    var v5 int = (v1 & v1)
    v5 = v5
    for fi8 := 0; fi8 < 2; fi8++ {
        v1 -= (^(fi8))
        _ = fi8
    }
    switch ("foo bar" + "foo bar") {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        if true {
            if true {
                fmt.Println(v4, v5, v3)
                v4 -= v4
                for fi10 := 0; fi10 < 18; fi10++ {
                    v1 += (v1 + -9223372036854775808)
                    v5 = fi10
                    _ = fi10
                }
            } else {
                if true {
                    var v11 uint16 = 32767
                    v3 -= 2
                    _ = v11
                } else {
                    switch -2147483648 {
                    case 32767:
                        v12 := GenS1{F0: (v4 | v4), F1: 533, F2: 2147483647}
                        _ = v12
                    default:
                        if (!false) {
                            _ = len(v2)
                            var v13 float32 = 0.0
                            if true {
                                var v14 string = "foo bar"
                                fmt.Println(v13, v3)
                                _ = v14
                            } else {
                                v1 = (fi9 ^ (-(v1)))
                                var v15 int = 28
                                _ = v15
                            }
                            _ = v13
                        }
                    }
                    var v16 uint16 = 763
                    var v17 float32 = 0.0
                    _ = v16
                    _ = v17
                }
                _ = len(v2)
            }
            v3 += (v3 + v3)
            v18 := GenS2{F0: -9223372036854775808}
            _ = v18
        }
        if false {
            switch 1 {
            case 4294967295:
                for fi19 := 0; fi19 < 16; fi19++ {
                    switch "hello" {
                    case "\n\t":
                        v20 := GenS2{F0: 0}
                        _ = v20
                    case "hello":
                        fmt.Println(v4, v5)
                    default:
                        if len(v2) > 0 {
                            v2[0] = v5
                        }
                    }
                    _ = fi19
                }
            case 9223372036854775807:
                switch ("a" + "") {
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    v4 = (v4 * (v4 % (((v4 << ((v4) & 63))) | 1)))
                }
            default:
                v21 := GenS3{F0: false}
                _ = v21
            }
            if false {
                if ((-2.25 == 0.0) && (0.0 < 0.1)) {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        if ((0.1 / -2.25) != 0.1) {
                            v23 := GenS1{F0: 750, F1: 574, F2: 2147483648}
                            if len(v2) > 0 {
                                v2[0] = -1
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v3, v4)
                            v1 += (fi9 ^ fi22)
                            _ = v23
                        } else {
                            fmt.Println(v1, fi22, fi9)
                        }
                        var v24 float64 = 3.14159
                        var v25 int8 = 127
                        _ = fi22
                        _ = v24
                        _ = v25
                    }
                } else {
                    var v26 bool = ((42 <= -2147483648) || (!(!((!true) && (((0.0 * -2.25) + float32((-2.25 / 1e10))) == 0.0)))))
                    _ = v26
                }
                if (v3 != v3) {
                    v5 = v1
                } else {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v4 += (v4 * (v4 >> ((v4) & 63)))
                        _ = fi27
                    }
                }
            }
        }
        _ = fi9
    }
    v28 := GenS4{F0: -1}
    v29 := GenS5{F0: (v1 | 2147483647)}
    var v30 int8 = -128
    if false {
        v31 := GenS4{F0: (v30 | v30)}
        _ = v31
    } else {
        if false {
            v30 += v30
        } else {
            v32 := GenS1{F0: v4, F1: 0, F2: 4294967295}
            v1 = -581
            _ = v32
        }
    }
    v33 := GenS3{F0: (!true)}
    var v34 int32 = 1
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 bool = true
        v2 = append(v2, v1)
        _ = fi35
        _ = v36
    }
    v30 = v30
    var v37 uint8 = (v3 & 2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v5
}
