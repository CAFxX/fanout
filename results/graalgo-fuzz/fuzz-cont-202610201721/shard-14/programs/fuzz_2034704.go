package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }
func genFn0(p0 int32, p1 int8) uint32 {
    v5 := []int{4294967295, -9223372036854775808}
    p0 -= 808
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    v1 := []int{}
    var v2 uint8 = 42
    var v3 int16 = 42
    v4 := []int{-9223372036854775808, -863, -790, -930}
    if (genFn0(-2147483648, 0) <= genFn0(-2147483648, 1)) {
        var v6 uint16 = 65535
        if (!false) {
            v2 = (v2 >> ((4294967295) & 63))
            if true {
                var v7 uint8 = 192
                v3 = v3
                if ((!false) || false) {
                    if (false || false) {
                        for fi8 := 0; fi8 < 20; fi8++ {
                            fmt.Println(v2, fi8)
                            v2 -= (^(v2))
                            _ = fi8
                        }
                        _ = len(v4)
                    } else {
                        v6 = v6
                        v7 -= (v2 << ((176) & 63))
                    }
                    for fi9 := 0; fi9 < 9; fi9++ {
                        fmt.Println(v6)
                        var v10 int = 128
                        v6 -= (v6 | 65535)
                        _ = fi9
                        _ = v10
                    }
                }
                _ = v7
            }
        } else {
            for fi11 := 0; fi11 < 20; fi11++ {
                for fi12 := 0; fi12 < 19; fi12++ {
                    _ = len(v4)
                    var v13 bool = (!(!true))
                    v3 = (v3 - (v3 | v3))
                    _ = fi12
                    _ = v13
                }
                if (0.1 < -2.25) {
                    var v14 bool = (0 < 0)
                    v6 = (v6 / (((v6 * v6)) | 1))
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v6 += 0
                        _ = fi15
                    }
                    _ = v14
                }
                fmt.Println(fi11, v3)
                _ = fi11
            }
        }
        _ = v6
    } else {
        for fi16 := 0; fi16 < 11; fi16++ {
            for fi17 := 0; fi17 < 4; fi17++ {
                fmt.Println(fi16)
                if false {
                    if true {
                        var v18 int64 = 1
                        _ = v18
                    } else {
                        fmt.Println(fi16)
                        fmt.Println(v2, fi16, fi17)
                    }
                }
                _ = fi17
            }
            fmt.Println(v2)
            var v19 uint16 = 2
            _ = fi16
            _ = v19
        }
        v2 = 255
        for fi20 := 0; fi20 < 6; fi20++ {
            if (true && true) {
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 uint64 = 15
                    var v23 float64 = -2.25
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            _ = fi20
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        if true {
            fmt.Println(v3, v3, fi24)
        } else {
            v2 = v2
            v3 -= 238
        }
        _ = fi24
    }
    _ = len(v4)
    v2 = 76
    var v25 uint32 = genFn0(-253, -29)
    if true {
        fmt.Println(v3, v3)
        if true {
            for fi26 := 0; fi26 < 14; fi26++ {
                v25 = (v25 / ((genFn0(-515, 127)) | 1))
                _ = fi26
            }
            var v27 int8 = -1
            _ = v27
        }
        v2 = (v2 >> ((2) & 63))
    }
    if (0.0 >= 0.1) {
        _ = len(v4)
        v2 = v2
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v3 = v3
        fmt.Println(v2)
        _ = fi28
    }
    fmt.Println(v25, v25, v2)
    if ((true && false) && (128 != 9223372036854775807)) {
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 int32 = -92
            v25 = v25
            var v31 uint64 = 1
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    v25 = (^(v25))
    if false {
        fmt.Println(v2, v3, v25)
        if (!(v2 <= 127)) {
            v25 += (v25 & (^(v25)))
        } else {
            v32 := GenS1{F0: (("\n\t" + "\n\t") != ("hello" + "foo bar")), F1: v25, F2: (^(v2))}
            _ = v32
        }
    } else {
        fmt.Println(v3, v25)
    }
    var v33 uint = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
}
