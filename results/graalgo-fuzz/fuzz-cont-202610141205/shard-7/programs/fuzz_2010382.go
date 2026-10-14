package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int }
type GenS5 struct { F0 int8; F1 int16; F2 uint16 }
func genFn0(p0 uint32) int {
    p0 += (p0 * p0)
    p0 -= p0
    switch "hello" {
    case "\n\t":
        p0 = p0
    case "a":
        v4 := GenS1{F0: 814, F1: (p0 >> ((2) & 63))}
        _ = v4
    case "foo bar":
        v5 := GenS1{F0: 225, F1: 0}
        _ = v5
    }
    var v6 int32 = -2147483648
    _ = p0
    _ = v6
    return -909
}
func genFn1() int {
    if false {
        v7 := GenS1{F0: 0, F1: 2147483648}
        _ = v7
    } else {
        v8 := GenS2{F0: true}
        var v9 uint32 = 4294967295
        v9 -= (v9 >> ((65535) & 63))
        _ = v8
        _ = v9
    }
    var v10 string = "a"
    _ = v10
    return 128
}

func main() {
    var v1 uint8 = 82
    var v2 float32 = 1e10
    v3 := make(map[string]string)
    var v11 int8 = 42
    delete(v3, "")
    fmt.Println(v11, v11)
    fmt.Println(v1)
    v1 += v1
    v2 -= v2
    switch ("a" + "hello") {
    case "foo bar":
        for fi12 := 0; fi12 < 18; fi12++ {
            if true {
                if true {
                    var v13 int8 = v11
                    for fi14 := 0; fi14 < 7; fi14++ {
                        if false {
                            var v15 int16 = -32768
                            var v16 uint64 = 2
                            _ = v15
                            _ = v16
                        } else {
                            if false {
                                var v17 int16 = -32768
                                _ = v17
                            }
                            switch "" {
                            case "":
                                v18 := []int{}
                                _ = v18
                            }
                            v19 := GenS3{F0: (v1 - v1)}
                            _ = v19
                        }
                        _ = fi14
                    }
                    _ = v13
                }
                v20 := GenS2{F0: (("foo bar" > "hello") || (669 < 734))}
                if true {
                    fmt.Println(fi12)
                }
                _ = v20
            }
            v21 := GenS4{F0: v1, F1: 0, F2: (fi12 - (fi12 + (fi12 ^ genFn1())))}
            _ = fi12
            _ = v21
        }
    case "\n\t":
        var v22 float64 = 0.0
        _ = v22
    default:
        var v23 bool = (!(true && true))
        _ = v23
    }
    v24 := []int{120, 32768, 2147483647}
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 float32 = (((float32(0.0) * v2) * v2) - 0.1)
        for fi27 := 0; fi27 < 11; fi27++ {
            v11 += v11
            fmt.Println(v2)
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    switch -9223372036854775808 {
    case 128:
        v28 := GenS3{F0: (v1 + v1)}
        _ = v28
    case -2147483648:
        for fi29 := 0; fi29 < 14; fi29++ {
            var v30 float32 = 1e10
            _ = fi29
            _ = v30
        }
    default:
        var v31 uint64 = 42
        _ = v31
    }
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    if false {
        switch "\n\t" {
        case "foo bar":
            var v32 float32 = 0.0
            _ = v32
        }
        if true {
            v2 += v2
            var v33 uint16 = 42
            v1 += (v1 % ((42) | 1))
            _ = v33
        }
    } else {
        switch -112 {
        case -1:
            for fi34 := 0; fi34 < 9; fi34++ {
                v1 += (v1 << ((0) & 63))
                for fi35 := 0; fi35 < 3; fi35++ {
                    switch genFn0(42) {
                    case 0:
                        switch 671 {
                        case 4294967295:
                            delete(v3, "\n\t")
                        case 674:
                            if len(v24) > 0 {
                                v24[0] = genFn1()
                            }
                        default:
                            var v36 uint64 = 93
                            _ = v36
                        }
                    default:
                        if (((true && true) && true) && (0 > 65535)) {
                            var v37 uint32 = 4294967295
                            _ = v37
                        } else {
                            if (!false) {
                                var v38 uint = 18446744073709551615
                                switch (fi34 << ((1) & 63)) {
                                case 42:
                                    switch 128 {
                                    case -2147483648:
                                        v38 = (^(v38))
                                    }
                                    var v39 int16 = int16(v1)
                                    _ = v39
                                case 32767:
                                    if len(v24) > 0 {
                                        v24[0] = 2147483647
                                    }
                                default:
                                    v1 = v1
                                }
                                _ = v38
                            }
                            fmt.Println(v1)
                            var v40 uint64 = 18446744073709551615
                            _ = v40
                        }
                    }
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                    _ = fi35
                }
                var v41 int64 = 454
                _ = fi34
                _ = v41
            }
        case 127:
            var v42 uint64 = 18446744073709551615
            _ = v42
        default:
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        }
        switch 274 {
        case 42:
            v2 = v2
        case 227:
            v24 = append(v24, -130)
        case 0:
            v24 = append(v24, 127)
        default:
            var v43 int8 = (v11 + 0)
            _ = v43
        }
    }
    var v44 int32 = 0
    v2 = ((0.0 * (1e10 - 1e10)) + float32((0.0 * -2.25)))
    var v45 float64 = 0.0
    v1 = (v1 % ((133) | 1))
    if true {
        v46 := GenS5{F0: (^(v11)), F1: 451, F2: 2}
        _ = v46
    } else {
        var v47 uint32 = 42
        v2 = float32(v45)
        _ = v47
    }
    var v48 int32 = -133
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v44
    _ = v45
    _ = v48
}
