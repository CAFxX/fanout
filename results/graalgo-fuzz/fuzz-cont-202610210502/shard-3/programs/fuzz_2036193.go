package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 uint32, p2 bool) uint32 {
    v7 := GenS1{F0: 1, F1: -2147483648}
    v8 := GenS1{F0: -128, F1: -143}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (^(p1))
}
func genFn1() string {
    switch 339 {
    case 42:
        if (730 >= 42) {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint16 = 1
                v10 -= v10
                fmt.Println(v10)
                _ = fi9
                _ = v10
            }
        } else {
            if ((true || ("hello" > "\n\t")) && true) {
                var v11 int32 = -1
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(fi12, fi12)
                    fmt.Println(fi12, v11)
                    if (802 >= 1) {
                        fmt.Println(fi12)
                    } else {
                        if true {
                            v13 := []int{(fi12 ^ (fi12 | (fi12 % ((fi12) | 1)))), 0, (-(fi12))}
                            _ = v13
                        } else {
                            for fi14 := 0; fi14 < 12; fi14++ {
                                v11 = 424
                                var v15 float32 = 1e10
                                _ = fi14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 10; fi16++ {
                                v17 := GenS1{F0: 16, F1: v11}
                                var v18 uint16 = 1
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            fmt.Println(v11)
                        }
                    }
                    _ = fi12
                }
                var v19 bool = (-9223372036854775808 > 1)
                _ = v11
                _ = v19
            } else {
                if false {
                    v20 := []int{2147483647}
                    switch ("a" + "foo bar") {
                    case "a":
                        var v21 float32 = 3.14159
                        _ = v21
                    default:
                        var v22 bool = false
                        _ = v22
                    }
                    _ = v20
                }
                var v23 uint8 = 153
                _ = v23
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                fmt.Println(fi24)
                _ = fi24
            }
            for fi25 := 0; fi25 < 11; fi25++ {
                v26 := []int{fi25}
                _ = fi25
                _ = v26
            }
        }
    case -9223372036854775808:
        var v27 int32 = 42
        _ = v27
    case -70:
        if false {
            v28 := GenS2{F0: ((true || false) || (!(!false)))}
            _ = v28
        } else {
            v29 := []int{2147483647, 128}
            var v30 uint32 = 230
            fmt.Println(v30)
            _ = v29
            _ = v30
        }
    default:
        switch -449 {
        case 0:
            if false {
                var v31 bool = true
                _ = v31
            }
        case -7:
            var v32 int8 = -1
            _ = v32
        case 419:
            v33 := []int{1}
            _ = v33
        default:
            for fi34 := 0; fi34 < 10; fi34++ {
                v35 := []int{}
                for fi36 := 0; fi36 < 19; fi36++ {
                    fmt.Println(fi36)
                    if len(v35) >= 2 {
                        v35 = v35[1:len(v35)]
                    }
                    _ = fi36
                }
                _ = fi34
                _ = v35
            }
        }
    }
    var v37 bool = ((2 >= 1) || false)
    _ = v37
    return ""
}

func main() {
    var v1 float32 = 0.1
    var v2 float32 = (-2.25 - v1)
    v3 := []string{}
    var v4 float32 = (1.5 * (v2 - float32(3.14159)))
    var v5 uint64 = 0
    var v6 int64 = -9223372036854775808
    for fi38 := 0; fi38 < 18; fi38++ {
        v6 = v6
        var v39 int16 = 1000
        _ = fi38
        _ = v39
    }
    fmt.Println(v6)
    v1 = v1
    v4 = (v1 / (v1 * ((v1 - ((3.14159 * (v1 / float32((0.1 - 0.0)))) + v4)) * 1e10)))
    if len(v3) > 0 {
        v3[0] = ""
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        v2 += 3.14159
        _ = fi40
    }
    if true {
        v4 += -2.25
        var v41 uint32 = 65535
        _ = v41
    } else {
        var v42 uint8 = uint8(v5)
        fmt.Println(v6, v1, v5)
        _ = v42
    }
    var v43 int = 975
    var v44 float32 = v4
    for fi45 := 0; fi45 < 1; fi45++ {
        if true {
            v5 -= (v5 ^ (v5 & 416))
            v1 += v44
            switch (v43 + (-(fi45))) {
            case -2147483648:
                v46 := GenS2{F0: ((0.1 * 0.1) < (v2 + (1e10 / v2)))}
                _ = v46
            case -505:
                v3 = append(v3, ("foo bar" + "a"))
            default:
                var v47 uint16 = 1
                _ = v47
            }
        }
        _ = fi45
    }
    var v48 int = v43
    var v49 int64 = (v6 - (^(v6)))
    for fi50 := 0; fi50 < 2; fi50++ {
        var v51 int = (v48 & v43)
        v4 += (v44 - v2)
        _ = fi50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
