package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint }
func genFn0() uint64 {
    var v5 float64 = 1.5
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint32 = 937
        var v8 int64 = 808
        v7 = (v7 * 1)
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: ("" + "foo bar"), F1: 1, F2: 1}
    _ = v5
    _ = v9
    return 0
}

func main() {
    var v1 int16 = -32768
    var v2 uint64 = 998
    var v3 uint = 130
    var v4 int8 = -52
    fmt.Println(v3, v4, v4)
    if (v1 < v1) {
        if false {
            var v10 int64 = 938
            var v11 uint = (^(v3))
            _ = v10
            _ = v11
        }
        var v12 bool = true
        v4 += 42
        _ = v12
    }
    var v13 int8 = v4
    for fi14 := 0; fi14 < 16; fi14++ {
        v1 = v1
        if true {
            v15 := GenS1{F0: "\n\t", F1: (v1 & (v1 | v1)), F2: 0}
            _ = v15
        }
        _ = fi14
    }
    v16 := GenS1{F0: ("" + "foo bar"), F1: (v1 >> ((v3) & 63)), F2: 0}
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(v4, fi17, v1)
        fmt.Println(v1)
        _ = fi17
    }
    switch "a" {
    case "":
        if true {
            switch 9223372036854775807 {
            case 127:
                switch "\n\t" {
                case "foo bar":
                    for fi18 := 0; fi18 < 16; fi18++ {
                        if false {
                            for fi19 := 0; fi19 < 8; fi19++ {
                                var v20 uint8 = 227
                                v21 := GenS2{F0: 0}
                                fmt.Println(v13)
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            for fi22 := 0; fi22 < 10; fi22++ {
                                fmt.Println(fi18, v4, v4)
                                v23 := []int{(fi18 >> ((v3) & 63)), (fi18 / (((fi22 / ((-9223372036854775808) | 1))) | 1)), fi18, 42}
                                _ = fi22
                                _ = v23
                            }
                        }
                        v13 = v4
                        for fi24 := 0; fi24 < 9; fi24++ {
                            v4 -= (v13 % ((-1) | 1))
                            switch (fi24 % ((32768) | 1)) {
                            case 128:
                                v3 -= (^(v3))
                            case 4294967295:
                                v25 := []int{(-(fi24)), (-(fi18)), (fi18 % (((^(fi24))) | 1))}
                                _ = v25
                            default:
                                var v26 int64 = 0
                                _ = v26
                            }
                            v27 := GenS3{F0: 1, F1: 292}
                            _ = fi24
                            _ = v27
                        }
                        _ = fi18
                    }
                default:
                    for fi28 := 0; fi28 < 7; fi28++ {
                        var v29 string = "hello"
                        _ = fi28
                        _ = v29
                    }
                }
            case -9223372036854775808:
                var v30 uint16 = 32767
                _ = v30
            }
            var v31 uint8 = 42
            _ = v31
        }
    case "hello":
        var v32 uint16 = 0
        _ = v32
    case "a":
        v33 := []int{}
        _ = v33
    default:
        var v34 int32 = 0
        _ = v34
    }
    var v35 int = -481
    if (false || false) {
        for fi36 := 0; fi36 < 8; fi36++ {
            switch (^(v35)) {
            case -9223372036854775808:
                v37 := []int{}
                _ = v37
            default:
                v2 -= v2
            }
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 float64 = 3.14159
                v40 := []int{(fi36 - v35), (fi38 << (((v3 / (((v3 ^ (v3 ^ (v3 ^ (v3 ^ v3))))) | 1))) & 63)), fi38}
                _ = fi38
                _ = v39
                _ = v40
            }
            var v41 int64 = 0
            _ = fi36
            _ = v41
        }
        var v42 int = v35
        _ = v42
    } else {
        v2 = 42
        v4 += (v13 ^ 22)
        v3 += (v3 * v3)
    }
    v2 = (v2 - v2)
    var v43 int16 = -232
    var v44 int32 = 26
    if false {
        v4 = -75
        var v45 uint = (v3 | v3)
        if (!(false && (v35 > (v35 % ((v35) | 1))))) {
            for fi46 := 0; fi46 < 11; fi46++ {
                v47 := GenS4{F0: -254, F1: v3}
                v45 = 615
                _ = fi46
                _ = v47
            }
            if (!(3.14159 == -2.25)) {
                var v48 string = "\n\t"
                v4 = (v4 >> (((v45 - uint(v2))) & 63))
                _ = v48
            }
            var v49 int32 = (v44 / (((^(v44))) | 1))
            _ = v49
        } else {
            var v50 int16 = v1
            _ = v50
        }
        _ = v45
    } else {
        v51 := []int{v35, v35}
        _ = v51
    }
    var v52 float64 = (1e10 / (-2.25 + float64(1.5)))
    var v53 int16 = v43
    v54 := []int{32767}
    switch (-(v35)) {
    case 4294967295:
        v52 -= float64(0.1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v52
    _ = v53
    _ = v54
}
