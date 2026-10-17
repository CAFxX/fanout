package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int32 }
func genFn0(p0 int16) int8 {
    v6 := GenS1{F0: -9223372036854775808, F1: 932}
    if false {
        var v7 uint32 = 81
        v8 := GenS1{F0: 0, F1: 255}
        _ = v7
        _ = v8
    }
    var v9 int32 = -1
    p0 = (p0 | p0)
    fmt.Println(p0, v9)
    _ = p0
    _ = v6
    _ = v9
    return -128
}
func genFn1() int {
    var v10 uint64 = 42
    if (2147483647 != 1) {
        fmt.Println(v10)
        fmt.Println(v10)
    } else {
        v10 -= v10
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 int32 = 1
            switch fi11 {
            case 92:
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 uint64 = v10
                    fmt.Println(v12, v12, v10)
                    if (("foo bar" + "hello") == "foo bar") {
                        if false {
                            var v15 uint32 = 2147483647
                            var v16 uint16 = 302
                            var v17 uint16 = v16
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        switch -344 {
                        case 128:
                            v18 := GenS1{F0: (^(fi11)), F1: 2}
                            _ = v18
                        case 42:
                            v19 := GenS1{F0: (-(fi11)), F1: 32767}
                            _ = v19
                        default:
                            v20 := GenS1{F0: 32768, F1: 1}
                            _ = v20
                        }
                    } else {
                        var v21 bool = true
                        if true {
                            var v22 uint64 = 18446744073709551615
                            _ = v22
                        } else {
                            v23 := []int{(fi11 >> ((4294967295) & 63)), (fi13 & fi11), fi13}
                            _ = v23
                        }
                        v24 := GenS1{F0: fi13, F1: 90}
                        _ = v21
                        _ = v24
                    }
                    if true {
                        var v25 int = fi11
                        switch "a" {
                        case "hello":
                            v26 := GenS2{F0: 829, F1: v10, F2: ((1.5 + -2.25) <= float64((0.1 / 1.5)))}
                            _ = v26
                        default:
                            v14 += v10
                        }
                        var v27 uint16 = 42
                        _ = v25
                        _ = v27
                    }
                    _ = fi13
                    _ = v14
                }
            case -119:
                var v28 uint = 18446744073709551615
                _ = v28
            }
            for fi29 := 0; fi29 < 19; fi29++ {
                v12 = (v12 & v12)
                _ = fi29
            }
            _ = fi11
            _ = v12
        }
    }
    v30 := GenS3{F0: 65535, F1: 2147483647, F2: 1}
    var v31 int64 = -754
    if false {
        v31 = -9223372036854775808
        var v32 bool = false
        _ = v32
    } else {
        var v33 uint = 2
        _ = v33
    }
    _ = v10
    _ = v30
    _ = v31
    return -1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 981
    v3 := make(map[uint32]uint32)
    v4 := []uint32{}
    v5 := []int{}
    delete(v3, 534)
    switch genFn1() {
    case 128:
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
    case 525:
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 uint = 255
            switch fi34 {
            case 42:
                var v36 int8 = -128
                _ = v36
            case -9223372036854775808:
                v2 = v2
            case 2147483647:
                var v37 int = (fi34 * fi34)
                _ = v37
            }
            _ = len(v4)
            _ = fi34
            _ = v35
        }
    default:
        for fi38 := 0; fi38 < 14; fi38++ {
            var v39 float32 = 0.0
            _ = fi38
            _ = v39
        }
    }
    v40 := GenS4{F0: 708, F1: 32767, F2: 42}
    var v41 float32 = (-2.25 + -2.25)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        if len(v5) > 0 {
            v5[0] = genFn1()
        }
        v43 := GenS4{F0: 678, F1: 65535, F2: -2147483648}
        v2 = v2
        _ = fi42
        _ = v43
    }
    switch "hello" {
    case "a":
        var v44 uint8 = 128
        _ = v44
    case "hello":
        _ = len(v5)
    default:
        if len(v5) > 0 {
            v5[0] = -2147483648
        }
    }
    switch genFn1() {
    case 257:
        if ((42 < 1) && (!false)) {
            var v45 int32 = 0
            _ = len(v4)
            _ = v45
        } else {
            var v46 float64 = (3.14159 - -2.25)
            v41 = v41
            _ = v46
        }
    case 0:
        fmt.Println(v41, v41)
    case 42:
        v2 = v2
    }
    if false {
        v41 -= (((0.0 / v41) * v41) - (v41 - (0.0 + (v41 * (v41 - ((1.5 * v41) * 3.14159))))))
        fmt.Println(v41, v2)
    } else {
        for fi47 := 0; fi47 < 7; fi47++ {
            if len(v5) > 0 {
                v5[0] = genFn1()
            }
            switch (fi47 * 32768) {
            case 4294967295:
                v2 -= 0
            default:
                var v48 string = ""
                _ = v48
            }
            var v49 uint = 0
            _ = fi47
            _ = v49
        }
        v2 = (v2 << ((42) & 63))
    }
    var v50 bool = (true || false)
    for fi51 := 0; fi51 < 16; fi51++ {
        v50 = true
        _ = fi51
    }
    for fi52 := 0; fi52 < 3; fi52++ {
        v50 = (146 > 0)
        _ = fi52
    }
    if false {
        var v53 int8 = 1
        var v54 uint = 897
        _ = len(v1)
        _ = v53
        _ = v54
    } else {
        fmt.Println(v2, v50)
        v55 := v3[65535]
        switch genFn1() {
        case 9223372036854775807:
            var v56 int = 0
            _ = v56
        case 382:
            v50 = v50
        }
        _ = v55
    }
    v41 -= v41
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v50
}
