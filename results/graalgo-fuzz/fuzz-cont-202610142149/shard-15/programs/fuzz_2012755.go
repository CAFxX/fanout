package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 uint8; F2 uint8 }
func genFn0() int16 {
    var v5 uint8 = 1
    var v6 int8 = 100
    _ = v5
    _ = v6
    return 137
}
func genFn1(p0 uint32) int8 {
    v7 := []int{1000}
    fmt.Println(p0)
    for fi8 := 0; fi8 < 8; fi8++ {
        _ = len(v7)
        _ = fi8
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v7
    return 127
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 string = "foo bar"
    v9 := GenS1{F0: v4, F1: 65535}
    for fi10 := 0; fi10 < 10; fi10++ {
        switch 897 {
        case -399:
            if false {
                v1 = (v1 * ((1.5 / v1) / 1.5))
                if false {
                    switch ("a" + v4) {
                    case "foo bar":
                        for fi11 := 0; fi11 < 13; fi11++ {
                            if false {
                                if (!(("a" == (((v4 + "hello") + (v4 + v4)) + v4)) && false)) {
                                    switch (fi11 ^ (fi11 * (fi10 + (-(fi11))))) {
                                    case -901:
                                        if false {
                                            if true {
                                                v4 = v4
                                                if (false || ((-2.25 * -2.25) >= 0.0)) {
                                                    v12 := GenS1{F0: "\n\t", F1: 1}
                                                    var v13 int16 = 0
                                                    _ = v12
                                                    _ = v13
                                                }
                                            } else {
                                                var v14 uint64 = 2
                                                _ = v14
                                            }
                                            v15 := GenS2{F0: 860, F1: (((1e10 * (v1 + float32((1.5 - 0.0)))) - float32(0.1)) > 1e10)}
                                            v16 := []int{fi11}
                                            _ = v15
                                            _ = v16
                                        }
                                    default:
                                        var v17 float64 = (3.14159 + -2.25)
                                        _ = v17
                                    }
                                } else {
                                    _ = len(v3)
                                }
                                v4 = v4
                                delete(v3, v4)
                            }
                            switch 42 {
                            case -678:
                                switch v4 {
                                case "a":
                                    var v18 float32 = v1
                                    _ = v18
                                default:
                                    var v19 int32 = -1
                                    _ = v19
                                }
                            default:
                                var v20 int = (fi11 << ((408) & 63))
                                _ = v20
                            }
                            _ = fi11
                        }
                    case "\n\t":
                        for fi21 := 0; fi21 < 11; fi21++ {
                            switch (((v4 + v4) + (((v4 + ((v4 + (v4 + "\n\t")) + ((v4 + "\n\t") + "\n\t"))) + (v4 + v4)) + ((v4 + "") + "a"))) + ((v4 + ("hello" + "")) + (v4 + ""))) {
                            case "\n\t":
                                fmt.Println(fi10, v1, fi21)
                            default:
                                fmt.Println(v1)
                            }
                            v4 = v4
                            switch fi10 {
                            case 128:
                                var v22 uint64 = 4294967295
                                _ = v22
                            }
                            var v23 float32 = 0.1
                            v23 -= v1
                            var v24 uint64 = 42
                            _ = fi21
                            _ = v23
                            _ = v24
                        }
                    case "":
                        fmt.Println(v4, v4)
                    }
                    for fi25 := 0; fi25 < 12; fi25++ {
                        _ = len(v3)
                        var v26 int = fi25
                        var v27 uint = 4294967295
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    if false {
                        v4 = v4
                    }
                }
            }
        }
        var v28 uint32 = 2147483648
        switch v4 {
        case "":
            v29 := []int{2147483647}
            _ = v29
        case "\n\t":
            v30 := []int{42, fi10, fi10, -9223372036854775808}
            _ = v30
        default:
            if true {
                _ = len(v3)
                v31 := GenS3{F0: 18446744073709551615, F1: -58}
                v32 := []int{834, fi10, (fi10 << ((42) & 63))}
                _ = v31
                _ = v32
            }
        }
        _ = fi10
        _ = v28
    }
    var v33 int16 = -126
    fmt.Println(v33, v1, v33)
    v34 := []int{860}
    v4 = "a"
    if len(v34) > 0 {
        v34[0] = 4294967295
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        v36 := GenS2{F0: -1, F1: true}
        v37 := GenS2{F0: 42, F1: true}
        _ = fi35
        _ = v36
        _ = v37
    }
    v33 = v33
    if false {
        if (!true) {
            v38 := GenS4{F0: 2, F1: -201}
            v33 = (v33 * v33)
            _ = v38
        } else {
            if true {
                v39 := GenS5{F0: 98}
                v40 := GenS2{F0: 42, F1: true}
                _ = v39
                _ = v40
            } else {
                v41 := GenS2{F0: -177, F1: (42 >= 42)}
                v4 = (((((v4 + v4) + "foo bar") + "a") + "\n\t") + v4)
                _ = v41
            }
            delete(v2, 2147483647)
            for fi42 := 0; fi42 < 1; fi42++ {
                v4 = "foo bar"
                v1 = -2.25
                _ = fi42
            }
        }
    } else {
        v4 = v4
        fmt.Println(v33, v4, v1)
    }
    v43 := GenS5{F0: 72}
    v44 := GenS6{F0: v33, F1: 255, F2: 2}
    var v45 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v9
}
