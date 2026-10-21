package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint64; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 uint32 }
func genFn0(p0 int) uint32 {
    if true {
        var v4 uint32 = 42
        v5 := GenS1{F0: 37}
        _ = v4
        _ = v5
    }
    p0 = p0
    if (42 <= 742) {
        p0 = (p0 + p0)
        var v6 uint16 = 1
        p0 -= 32768
        _ = v6
    }
    var v7 int64 = -9223372036854775808
    v7 += v7
    _ = p0
    _ = v7
    return 65535
}
func genFn1() uint32 {
    var v8 string = "\n\t"
    v8 = "hello"
    var v9 float64 = -2.25
    if true {
        v10 := []int{128, 696}
        var v11 bool = false
        _ = v10
        _ = v11
    }
    switch 42 {
    case 9223372036854775807:
        var v12 bool = ((!false) && (2 <= 141))
        _ = v12
    case 0:
        if true {
            switch ("\n\t" + v8) {
            case "a":
                v8 = "hello"
            case "foo bar":
                v13 := []int{-870, 258, -894}
                _ = v13
            case "\n\t":
                v8 = v8
            }
            if true {
                var v14 uint16 = 255
                var v15 uint32 = 65535
                v9 = 1e10
                _ = v14
                _ = v15
            } else {
                v16 := []int{2147483647, 32768, -548, -1}
                fmt.Println(v8)
                _ = v16
            }
            switch v8 {
            case "":
                switch 0 {
                case -2147483648:
                    v17 := []int{4294967295, 127, -2147483648, 2147483647}
                    _ = v17
                case 32767:
                    if true {
                        v18 := []int{}
                        _ = v18
                    }
                }
            case "hello":
                for fi19 := 0; fi19 < 16; fi19++ {
                    v9 = v9
                    switch "foo bar" {
                    case "foo bar":
                        v20 := []int{fi19}
                        _ = v20
                    case "\n\t":
                        v21 := []int{}
                        _ = v21
                    default:
                        switch 0 {
                        case -1:
                            for fi22 := 0; fi22 < 18; fi22++ {
                                fmt.Println(fi19, fi22)
                                _ = fi22
                            }
                        case -409:
                            var v23 uint64 = 2
                            _ = v23
                        default:
                            v24 := GenS2{F0: true}
                            _ = v24
                        }
                    }
                    _ = fi19
                }
            case "\n\t":
                v25 := GenS3{F0: 32767, F1: 761, F2: -1}
                _ = v25
            }
        } else {
            var v26 uint64 = 0
            v27 := GenS4{F0: -254}
            fmt.Println(v9)
            _ = v26
            _ = v27
        }
    case 1:
        switch v8 {
        case "a":
            for fi28 := 0; fi28 < 3; fi28++ {
                var v29 int64 = 331
                switch ("hello" + v8) {
                case "":
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v8 = "a"
                        v31 := []int{(-(fi30)), 1, (fi30 + 0), 127}
                        var v32 int16 = 1000
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
                for fi33 := 0; fi33 < 18; fi33++ {
                    v8 = v8
                    v9 += ((1.5 * (v9 / (-2.25 + 1e10))) - (0.0 / 1e10))
                    _ = fi33
                }
                _ = fi28
                _ = v29
            }
        case "hello":
            var v34 bool = false
            _ = v34
        case "foo bar":
            if (341 == 255) {
                v35 := GenS4{F0: 128}
                _ = v35
            } else {
                v36 := []int{0}
                var v37 bool = (-128 == 42)
                v9 = 0.0
                _ = v36
                _ = v37
            }
        }
    default:
        switch v8 {
        case "foo bar":
            v38 := []int{1, 396, -9223372036854775808}
            _ = v38
        case "a":
            for fi39 := 0; fi39 < 18; fi39++ {
                var v40 int8 = 42
                v8 = ("foo bar" + ((v8 + (("hello" + ((v8 + ((v8 + v8) + v8)) + (v8 + v8))) + v8)) + v8))
                fmt.Println(fi39, v9)
                _ = fi39
                _ = v40
            }
        case "\n\t":
            var v41 uint64 = 0
            _ = v41
        }
    }
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-9223372036854775808, 1, 1}
    var v3 int8 = -51
    fmt.Println(v3)
    v3 -= v3
    _ = len(v1)
    v42 := GenS2{F0: false}
    if true {
        v3 += v3
        v3 += -1
        var v43 int32 = -428
        _ = v43
    } else {
        if (true && false) {
            delete(v1, 42)
            delete(v1, 2147483648)
        }
        if true {
            var v44 uint16 = 467
            _ = v44
        } else {
            if (!false) {
                v45 := v1[genFn0(32768)]
                v46 := v1[(v45 << ((4294967295) & 63))]
                _ = v45
                _ = v46
            } else {
                if ((-239 >= 2147483647) && (127 > 30)) {
                    if true {
                        var v47 uint16 = 32767
                        _ = len(v1)
                        _ = v47
                    } else {
                        delete(v1, 121)
                        v48 := GenS5{F0: 809}
                        var v49 uint = 65535
                        _ = v48
                        _ = v49
                    }
                } else {
                    if true {
                        v3 = v3
                        v3 = 0
                    }
                    var v50 int = -826
                    _ = v50
                }
                var v51 int32 = 2147483647
                _ = v51
            }
        }
        var v52 uint = 42
        _ = v52
    }
    switch "foo bar" {
    case "":
        delete(v1, genFn0(-2147483648))
    }
    var v53 int32 = -2147483648
    v54 := GenS6{F0: -32768, F1: 545}
    if true {
        fmt.Println(v53, v3)
        fmt.Println(v53, v53)
    } else {
        switch "a" {
        case "foo bar":
            var v55 uint64 = 4294967295
            _ = v55
        case "\n\t":
            var v56 uint64 = 571
            _ = v56
        default:
            for fi57 := 0; fi57 < 19; fi57++ {
                v53 = (v53 ^ -303)
                _ = fi57
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v53
    _ = v54
}
