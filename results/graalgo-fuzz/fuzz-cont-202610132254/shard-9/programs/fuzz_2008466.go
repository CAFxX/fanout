package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int32; F1 bool; F2 uint32 }
func genFn0(p0 string, p1 uint16) uint16 {
    if (!(4294967295 < 1)) {
        if (((true && false) && true) && ("hello" <= p0)) {
            v4 := []int{0, 32768}
            var v5 int32 = 25
            _ = v4
            _ = v5
        } else {
            var v6 float32 = ((float32(float64((float32(0.0) - ((3.14159 * -2.25) / (-2.25 + 0.1))))) + (-2.25 + 1e10)) - (((3.14159 + 1e10) + float32(-2.25)) * 1.5))
            for fi7 := 0; fi7 < 9; fi7++ {
                if true {
                    for fi8 := 0; fi8 < 20; fi8++ {
                        v6 = 3.14159
                        var v9 uint = 2
                        _ = fi8
                        _ = v9
                    }
                    if (!(!(!(true && true)))) {
                        var v10 int8 = -128
                        var v11 uint16 = p1
                        switch "a" {
                        case "":
                            for fi12 := 0; fi12 < 10; fi12++ {
                                switch -666 {
                                case -133:
                                    var v13 uint8 = 2
                                    _ = v13
                                case 32767:
                                    v10 -= v10
                                default:
                                    var v14 int = fi12
                                    _ = v14
                                }
                                fmt.Println(p1, v10)
                                _ = fi12
                            }
                        case "a":
                            if false {
                                v6 -= v6
                            } else {
                                v15 := GenS1{F0: 671}
                                _ = v15
                            }
                        default:
                            p1 += (p1 ^ p1)
                        }
                        _ = v10
                        _ = v11
                    } else {
                        p0 = ((("a" + p0) + p0) + "\n\t")
                    }
                    var v16 float32 = 1.5
                    _ = v16
                }
                _ = fi7
            }
            if true {
                v6 = 3.14159
            } else {
                p0 = ((("hello" + "foo bar") + p0) + p0)
                var v17 string = (p0 + ("\n\t" + "a"))
                _ = v17
            }
            _ = v6
        }
    }
    var v18 uint32 = 2
    if true {
        for fi19 := 0; fi19 < 20; fi19++ {
            if false {
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 uint64 = 4294967295
                    var v22 int64 = 9223372036854775807
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 int16 = 106
                _ = v23
            } else {
                var v24 uint16 = p1
                if (752 < v18) {
                    switch "hello" {
                    case "hello":
                        switch p0 {
                        case "\n\t":
                            v25 := []int{(fi19 + (fi19 & 639))}
                            _ = v25
                        default:
                            v26 := GenS1{F0: -9223372036854775808}
                            _ = v26
                        }
                    case "\n\t":
                        v27 := GenS1{F0: 9223372036854775807}
                        _ = v27
                    default:
                        var v28 int = fi19
                        _ = v28
                    }
                    var v29 int16 = 1000
                    v24 = (v24 * p1)
                    _ = v29
                } else {
                    var v30 int64 = 42
                    _ = v30
                }
                switch 32767 {
                case -895:
                    v31 := GenS2{F0: -128, F1: fi19}
                    _ = v31
                case 2147483647:
                    fmt.Println(p0, v24, p1)
                }
                _ = v24
            }
            _ = fi19
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            v18 = v18
            _ = fi32
        }
    }
    _ = p0
    _ = p1
    _ = v18
    return (^(p1))
}

func main() {
    var v1 int = 1
    var v2 int = 966
    var v3 uint = 255
    v2 = v1
    v3 -= 42
    v33 := []int{v2}
    var v34 float64 = 0.0
    v35 := GenS3{F0: -613, F1: v3}
    for fi36 := 0; fi36 < 9; fi36++ {
        switch v1 {
        case 128:
            var v37 uint16 = 319
            _ = v37
        case 4294967295:
            var v38 uint32 = 2147483647
            _ = v38
        case 32768:
            for fi39 := 0; fi39 < 7; fi39++ {
                switch "hello" {
                case "\n\t":
                    v1 = (-(v1))
                    fmt.Println(v1, v1)
                    v40 := GenS4{F0: 0, F1: true, F2: 1}
                    _ = v40
                default:
                    v2 = (v2 & 4294967295)
                }
                var v41 bool = true
                var v42 int = (v1 | v1)
                _ = fi39
                _ = v41
                _ = v42
            }
        }
        var v43 int32 = 42
        _ = fi36
        _ = v43
    }
    v44 := GenS1{F0: -1}
    if (!(!(!false))) {
        v3 += 2
        fmt.Println(v2, v1, v1)
        for fi45 := 0; fi45 < 5; fi45++ {
            var v46 float64 = 0.1
            _ = fi45
            _ = v46
        }
    }
    fmt.Println(v34)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v44
}
