package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint32; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 string; F1 int16 }
func genFn0(p0 bool, p1 int32, p2 int) uint16 {
    fmt.Println(p1, p2)
    p0 = p0
    var v6 int16 = 1000
    switch p2 {
    case 121:
        v6 -= v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]uint32)
    var v3 uint = uint(v1)
    var v4 int64 = 1
    var v5 uint8 = 127
    v1 = (v1 | 42)
    var v7 float64 = 0.0
    switch 2147483647 {
    case -373:
        fmt.Println(v7, v4, v7)
    case 32768:
        v8 := GenS1{F0: (-(v4)), F1: 2147483647}
        _ = v8
    case 32767:
        switch ("" + "foo bar") {
        case "hello":
            if false {
                switch -818 {
                case 127:
                    for fi9 := 0; fi9 < 2; fi9++ {
                        for fi10 := 0; fi10 < 8; fi10++ {
                            var v11 bool = true
                            switch "foo bar" {
                            case "foo bar":
                                switch "" {
                                case "\n\t":
                                    var v12 int32 = 1
                                    _ = v12
                                case "hello":
                                    var v13 uint64 = 0
                                    _ = v13
                                case "a":
                                    fmt.Println(v1, fi10)
                                default:
                                    var v14 int32 = 0
                                    _ = v14
                                }
                            case "a":
                                v5 -= 24
                            case "hello":
                                v15 := GenS1{F0: (v4 * 1), F1: uint32(v3)}
                                _ = v15
                            default:
                                if v11 {
                                    if (v3 == (^(v3))) {
                                        if v11 {
                                            fmt.Println(v7, v4)
                                            fmt.Println(v5)
                                            v16 := GenS2{F0: v11, F1: 0, F2: 0}
                                            _ = v16
                                        }
                                    } else {
                                        v7 = 3.14159
                                        var v17 uint = 18446744073709551615
                                        switch (fi9 - fi9) {
                                        case 53:
                                            var v18 uint16 = 255
                                            _ = v18
                                        case 42:
                                            v19 := GenS3{F0: 42}
                                            _ = v19
                                        case -4:
                                            v3 = uint(v1)
                                        }
                                        _ = v17
                                    }
                                    var v20 uint32 = 65535
                                    v21 := GenS2{F0: v11, F1: (v20 + v20), F2: 100}
                                    v22 := v2[fi10]
                                    v1 = genFn0(false, 345, fi9)
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                            }
                            v3 = v3
                            _ = fi10
                            _ = v11
                        }
                        _ = fi9
                    }
                case 483:
                    fmt.Println(v1, v4, v7)
                case -703:
                    delete(v2, 821)
                default:
                    var v23 uint32 = 4294967295
                    _ = v23
                }
                var v24 int = 32767
                _ = v24
            } else {
                v4 = -21
                for fi25 := 0; fi25 < 17; fi25++ {
                    v3 = v3
                    v3 -= uint(v1)
                    _ = fi25
                }
            }
        default:
            v3 = v3
        }
    default:
        v26 := GenS4{F0: 2}
        _ = v26
    }
    var v27 int64 = (v4 | (^(v4)))
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS5{F0: 80, F1: (1.5 != 1.5)}
        _ = fi28
        _ = v29
    }
    if (float64((1.5 - (0.0 - 1.5))) == -2.25) {
        var v30 float64 = -2.25
        var v31 string = "a"
        _ = v30
        _ = v31
    } else {
        switch 32768 {
        case 32767:
            if true {
                var v32 uint32 = 2147483648
                _ = v32
            }
        case 9223372036854775807:
            switch "hello" {
            case "hello":
                var v33 float32 = (0.0 * -2.25)
                _ = v33
            }
        }
        fmt.Println(v7, v1)
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 string = "\n\t"
            _ = fi34
            _ = v35
        }
    }
    var v36 uint64 = 2
    v37 := GenS6{F0: 127}
    var v38 bool = true
    for fi39 := 0; fi39 < 3; fi39++ {
        for fi40 := 0; fi40 < 14; fi40++ {
            switch 2147483647 {
            case -9223372036854775808:
                v1 = v1
            case 32767:
                fmt.Println(v1)
            default:
                v5 -= v5
            }
            var v41 float32 = 1.5
            _ = fi40
            _ = v41
        }
        v5 -= v5
        var v42 bool = (true && (v38 && (0 >= 571)))
        _ = fi39
        _ = v42
    }
    if v38 {
        var v43 int = 413
        _ = v43
    } else {
        var v44 uint16 = uint16(v3)
        var v45 bool = (((^(v44)) < (^(v44))) || (624 >= 42))
        _ = v44
        _ = v45
    }
    v3 = (v3 % ((4294967295) | 1))
    delete(v2, -262)
    v46 := GenS7{F0: v4, F1: (v1 + (v1 - v1))}
    v47 := GenS8{F0: "hello", F1: -727}
    var v48 float64 = v7
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v7
}
