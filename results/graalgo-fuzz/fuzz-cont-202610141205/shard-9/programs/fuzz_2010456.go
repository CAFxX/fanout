package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int; F1 uint8; F2 bool }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint8) int32 {
    switch ("foo bar" + "") {
    case "foo bar":
        for fi7 := 0; fi7 < 2; fi7++ {
            p1 += p1
            p0 = 2
            var v8 bool = (3.14159 >= 1.5)
            _ = fi7
            _ = v8
        }
    case "\n\t":
        v9 := []int{32767}
        _ = v9
    }
    p1 = 30
    _ = p0
    _ = p1
    return -661
}
func genFn1(p0 uint) uint {
    p0 += 255
    if ((-2.25 * 1.5) != 1e10) {
        v10 := []int{128}
        switch -683 {
        case -491:
            if (42 == 1) {
                var v11 int64 = 8
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := GenS1{F0: 53, F1: 365, F2: 376}
                    switch fi12 {
                    case 4294967295:
                        if false {
                            switch "a" {
                            case "foo bar":
                                v14 := GenS2{F0: (("foo bar" + ("a" + ("foo bar" + ("foo bar" + "\n\t")))) + "hello")}
                                _ = v14
                            }
                            p0 = p0
                            switch ("foo bar" + "\n\t") {
                            case "foo bar":
                                for fi15 := 0; fi15 < 2; fi15++ {
                                    if false {
                                        var v16 int32 = -1
                                        _ = v16
                                    }
                                    v11 = v11
                                    p0 += (p0 ^ 320)
                                    _ = fi15
                                }
                            }
                        } else {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        }
                    case 0:
                        var v17 uint = 42
                        _ = v17
                    case 909:
                        if true {
                            fmt.Println(fi12, p0)
                        } else {
                            v18 := GenS1{F0: 628, F1: 379, F2: -886}
                            _ = v18
                        }
                    default:
                        var v19 int = 4294967295
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 12; fi20++ {
                        v21 := GenS1{F0: v11, F1: 839, F2: 1}
                        v22 := GenS2{F0: "\n\t"}
                        v23 := GenS1{F0: v11, F1: 42, F2: 347}
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = fi12
                    _ = v13
                }
                if true {
                    switch 128 {
                    case 0:
                        for fi24 := 0; fi24 < 19; fi24++ {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v11 -= v11
                            _ = fi24
                        }
                    case 9223372036854775807:
                        var v25 bool = false
                        _ = v25
                    case -1:
                        v26 := GenS2{F0: ("\n\t" + "\n\t")}
                        _ = v26
                    }
                    v11 = v11
                    if (!false) {
                        fmt.Println(v11, p0)
                    }
                } else {
                    fmt.Println(v11, v11)
                    p0 = (^(p0))
                    v10 = append(v10, 9223372036854775807)
                }
                _ = v11
            }
        case -9223372036854775808:
            p0 = p0
        }
        fmt.Println(p0)
        _ = v10
    } else {
        var v27 uint8 = 1
        _ = v27
    }
    if (p0 >= (p0 & 42)) {
        if true {
            fmt.Println(p0)
            var v28 uint64 = 907
            p0 -= 42
            _ = v28
        } else {
            p0 = p0
        }
    } else {
        if false {
            v29 := []int{}
            fmt.Println(p0)
            _ = v29
        } else {
            fmt.Println(p0)
            var v30 uint32 = 0
            v31 := []int{611}
            _ = v30
            _ = v31
        }
        fmt.Println(p0)
        if true {
            v32 := GenS3{F0: 42, F1: 645}
            v33 := []int{32767, 760}
            var v34 float64 = float64((float32(0.1) + 0.1))
            _ = v32
            _ = v33
            _ = v34
        }
    }
    _ = p0
    return (p0 & (p0 | (p0 | p0)))
}

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 18446744073709551615
    var v3 uint = 42
    var v4 string = "hello"
    var v5 uint = (^(v3))
    var v6 uint16 = (^(v1))
    fmt.Println(v1)
    v3 = v3
    fmt.Println(v5, v5, v5)
    var v35 float64 = -2.25
    if false {
        v5 = (v5 + (^(v5)))
        var v36 uint16 = (v1 << ((42) & 63))
        v37 := []int{-322, 388}
        _ = v36
        _ = v37
    } else {
        v2 += (^(v2))
        var v38 uint = uint(v1)
        v5 = genFn1(genFn1(genFn1(v5)))
        _ = v38
    }
    v39 := GenS4{F0: 1, F1: 141, F2: (!true)}
    v35 = v35
    v40 := GenS5{F0: 100}
    v41 := GenS2{F0: (v4 + v4)}
    switch 1 {
    case 2147483647:
        v42 := []int{}
        _ = v42
    }
    if true {
        var v43 float64 = float64(1e10)
        _ = v43
    } else {
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
