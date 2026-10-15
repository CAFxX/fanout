package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint16 }
func genFn0() int8 {
    switch 42 {
    case -764:
        var v4 float32 = (0.0 / 3.14159)
        _ = v4
    case 32767:
        var v5 float32 = 1e10
        _ = v5
    }
    if false {
        v6 := GenS1{F0: 65535, F1: 0}
        _ = v6
    } else {
        var v7 float64 = 0.0
        switch ("foo bar" + "hello") {
        case "\n\t":
            if false {
                fmt.Println(v7)
            }
        case "hello":
            fmt.Println(v7)
        default:
            v8 := GenS1{F0: 2, F1: 0}
            _ = v8
        }
        if false {
            if false {
                fmt.Println(v7)
            } else {
                v7 = -2.25
                var v9 uint = 140
                if (((("" + "a") + "\n\t") + "a") >= "foo bar") {
                    if false {
                        v10 := []int{-301, 9223372036854775807, 128}
                        var v11 string = "a"
                        var v12 bool = (65535 >= 1)
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    var v13 uint64 = 4294967295
                    _ = v13
                } else {
                    if false {
                        v9 += (^(v9))
                        for fi14 := 0; fi14 < 18; fi14++ {
                            v15 := GenS1{F0: 2147483647, F1: 65535}
                            fmt.Println(fi14, v9)
                            v16 := GenS2{F0: 255, F1: 0}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 uint = 1
                        _ = v17
                    } else {
                        fmt.Println(v7, v7)
                        v9 -= v9
                        if (42 != 1) {
                            v18 := GenS1{F0: uint32(v9), F1: 42}
                            _ = v18
                        } else {
                            var v19 int32 = 421
                            var v20 uint64 = 42
                            _ = v19
                            _ = v20
                        }
                    }
                    switch "a" {
                    case "\n\t":
                        if true {
                            var v21 bool = false
                            for fi22 := 0; fi22 < 3; fi22++ {
                                v23 := []int{(fi22 / ((fi22) | 1))}
                                var v24 float64 = 1.5
                                for fi25 := 0; fi25 < 11; fi25++ {
                                    var v26 int64 = 9223372036854775807
                                    _ = fi25
                                    _ = v26
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        } else {
                            v7 = 1e10
                            switch "a" {
                            case "\n\t":
                                var v27 float64 = 0.1
                                _ = v27
                            }
                            var v28 uint32 = 42
                            _ = v28
                        }
                    case "hello":
                        var v29 int64 = 937
                        _ = v29
                    case "foo bar":
                        if true {
                            var v30 uint16 = 65535
                            v31 := []int{63}
                            _ = v30
                            _ = v31
                        }
                    default:
                        var v32 float64 = float64((0.0 + -2.25))
                        _ = v32
                    }
                    var v33 int = 2147483647
                    _ = v33
                }
                _ = v9
            }
            for fi34 := 0; fi34 < 15; fi34++ {
                var v35 int16 = -1
                _ = fi34
                _ = v35
            }
            v7 += (v7 * v7)
        } else {
            fmt.Println(v7)
        }
        _ = v7
    }
    return 1
}

func main() {
    var v1 int8 = 0
    var v2 int32 = 0
    var v3 int32 = (^(v2))
    v1 = (v1 ^ 86)
    v3 = 42
    for fi36 := 0; fi36 < 15; fi36++ {
        for fi37 := 0; fi37 < 1; fi37++ {
            v2 = 150
            _ = fi37
        }
        if false {
            v38 := []int{-77, (-(fi36)), (fi36 << ((0) & 63))}
            _ = v38
        }
        _ = fi36
    }
    switch ("foo bar" + "a") {
    case "\n\t":
        v39 := GenS2{F0: 255, F1: 127}
        _ = v39
    case "a":
        v40 := GenS2{F0: 128, F1: -9223372036854775808}
        _ = v40
    case "foo bar":
        if true {
            var v41 bool = false
            _ = v41
        } else {
            if true {
                v42 := GenS3{F0: 127}
                _ = v42
            }
        }
    }
    v1 += genFn0()
    var v43 int8 = v1
    if false {
        switch 0 {
        case -128:
            var v44 uint64 = 4294967295
            _ = v44
        }
        v1 += v43
        v45 := GenS1{F0: 911, F1: 255}
        _ = v45
    }
    fmt.Println(v3, v1)
    v46 := []int{127, -839}
    v3 = (v3 << ((255) & 63))
    v47 := GenS4{F0: -1, F1: 2}
    v43 = (v1 + int8(v2))
    var v48 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v46
    _ = v47
    _ = v48
}
