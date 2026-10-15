package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool }

func main() {
    v1 := []uint32{}
    v2 := []int{9223372036854775807, 712}
    v3 := make(map[uint32]int)
    var v4 string = ""
    switch 4294967295 {
    case 32768:
        v4 = "a"
    }
    var v5 int8 = 0
    if (-654 >= -32768) {
        if (-1 < -212) {
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 uint8 = 255
                for fi8 := 0; fi8 < 8; fi8++ {
                    v9 := GenS1{F0: v5, F1: 80, F2: 1000}
                    _ = fi8
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
            fmt.Println(v4, v4)
        }
        var v10 uint16 = 2
        _ = v10
    } else {
        v11 := GenS2{F0: 9223372036854775807, F1: 0, F2: -1}
        _ = v11
    }
    switch v4 {
    case "a":
        if ((!(true || false)) || true) {
            v12 := GenS3{F0: (!true)}
            _ = v12
        }
    case "hello":
        switch 252 {
        case 975:
            v5 += (v5 * v5)
        case -2147483648:
            var v13 bool = true
            _ = v13
        }
    case "\n\t":
        if false {
            if len(v2) > 0 {
                v2[0] = 981
            }
        } else {
            delete(v3, 2)
            delete(v3, 2147483647)
        }
    default:
        if true {
            var v14 float64 = 0.0
            _ = v14
        } else {
            if true {
                var v15 int32 = 574
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 int = fi16
                    v5 = (v5 << ((4294967295) & 63))
                    v15 -= (-(v15))
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                v4 = "foo bar"
                for fi19 := 0; fi19 < 3; fi19++ {
                    if len(v2) > 0 {
                        v2[0] = -9223372036854775808
                    }
                    switch fi19 {
                    case 128:
                        v4 = v4
                        fmt.Println(v5)
                        switch fi19 {
                        case -2147483648:
                            fmt.Println(v4)
                        }
                        _ = len(v3)
                    }
                    var v20 int = 128
                    _ = len(v3)
                    v5 += 61
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        }
    }
    v5 -= (-(v5))
    if false {
        if true {
            v4 = v4
        } else {
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 int16 = 354
                var v23 bool = false
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 uint = 65535
            _ = v24
        }
    } else {
        v5 += v5
    }
    var v25 uint32 = 294
    var v26 int8 = (^(v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
