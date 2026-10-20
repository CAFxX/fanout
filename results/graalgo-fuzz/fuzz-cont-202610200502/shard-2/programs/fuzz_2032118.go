package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 string }
func genFn0(p0 uint16) int {
    if true {
        v4 := []int{}
        p0 = (^(p0))
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v5 float64 = 0.0
            _ = len(v4)
            _ = v5
        }
        _ = v4
    }
    switch 0 {
    case 0:
        v6 := GenS1{F0: 208, F1: 37}
        _ = v6
    case 32768:
        switch "hello" {
        case "a":
            switch 127 {
            case -2147483648:
                var v7 string = ""
                _ = v7
            default:
                for fi8 := 0; fi8 < 18; fi8++ {
                    switch (fi8 - fi8) {
                    case 4294967295:
                        fmt.Println(fi8, fi8)
                    case 1:
                        fmt.Println(p0, fi8)
                    case 0:
                        for fi9 := 0; fi9 < 15; fi9++ {
                            var v10 int = (fi8 / ((1) | 1))
                            if false {
                                var v11 uint64 = 0
                                var v12 int8 = 1
                                _ = v11
                                _ = v12
                            } else {
                                if (false || false) {
                                    v13 := []int{fi8, -9223372036854775808}
                                    switch "hello" {
                                    case "hello":
                                        var v14 float32 = 3.14159
                                        _ = v14
                                    default:
                                        p0 += (p0 + (p0 % (((^(p0))) | 1)))
                                        v13 = append(v13, fi8)
                                        v15 := GenS2{F0: 18446744073709551615}
                                        _ = v15
                                    }
                                    _ = v13
                                }
                            }
                            v16 := []int{v10, fi8}
                            _ = fi9
                            _ = v10
                            _ = v16
                        }
                    }
                    var v17 string = ""
                    _ = fi8
                    _ = v17
                }
            }
        case "foo bar":
            p0 -= (^(p0))
        default:
            fmt.Println(p0)
        }
    case 374:
        switch -658 {
        case 127:
            switch "" {
            case "hello":
                v18 := []int{9223372036854775807, 32768}
                _ = v18
            }
        case 42:
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int64 = 1
                var v21 string = ("" + "")
                v22 := []int{42}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        case 90:
            fmt.Println(p0)
        default:
            p0 -= p0
        }
    }
    _ = p0
    return -892
}

func main() {
    v1 := []uint32{1, 2}
    var v2 int32 = -1
    var v3 uint16 = 103
    for fi23 := 0; fi23 < 11; fi23++ {
        v3 = v3
        _ = fi23
    }
    v2 += -85
    var v24 int64 = 1
    v3 -= (v3 + v3)
    v25 := GenS3{F0: 42, F1: "\n\t"}
    fmt.Println(v2, v24, v3)
    if false {
        v26 := GenS1{F0: 18446744073709551615, F1: 134}
        _ = v26
    } else {
        v2 = 42
        var v27 bool = ((true && (!true)) && true)
        _ = v27
    }
    if true {
        _ = len(v1)
    }
    fmt.Println(v2, v24)
    if (v3 >= v3) {
        if false {
            if false {
                for fi28 := 0; fi28 < 8; fi28++ {
                    var v29 int64 = (v24 << ((734) & 63))
                    var v30 int32 = v2
                    v1 = append(v1, 973)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v31 := GenS1{F0: 971, F1: 7}
                v2 = v2
                _ = v31
            } else {
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
            }
        }
    } else {
        var v32 uint = 255
        if false {
            v32 = 0
            if false {
                var v33 bool = false
                switch 32767 {
                case -942:
                    switch 0 {
                    case 37:
                        var v34 int = -316
                        _ = v34
                    }
                }
                _ = v33
            }
        }
        _ = v32
    }
    v24 = v24
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
