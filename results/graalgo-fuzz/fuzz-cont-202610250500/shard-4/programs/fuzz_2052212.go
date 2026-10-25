package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 uint32; F1 string }

func main() {
    var v1 float32 = (3.14159 + (-2.25 * (1.5 * 0.0)))
    var v2 uint32 = 2147483648
    var v3 int32 = int32(v2)
    fmt.Println(v2, v1)
    fmt.Println(v2, v2)
    if (!false) {
        var v4 int32 = 1
        var v5 int32 = (v3 + v4)
        v6 := GenS1{F0: 1, F1: -952, F2: -204}
        _ = v4
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 18; fi7++ {
            if true {
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 uint = 65535
                    var v10 int32 = v3
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                for fi11 := 0; fi11 < 20; fi11++ {
                    v3 += v3
                    _ = fi11
                }
                var v12 bool = true
                _ = v12
            }
            var v13 uint = uint(v2)
            for fi14 := 0; fi14 < 17; fi14++ {
                fmt.Println(v1)
                v2 = (v2 & (^(v2)))
                v2 += (^(v2))
                v13 += v13
                _ = fi14
            }
            _ = fi7
            _ = v13
        }
        switch 1 {
        case -9223372036854775808:
            v15 := GenS1{F0: 0, F1: -9223372036854775808, F2: 42}
            _ = v15
        case 1:
            var v16 uint64 = 2
            _ = v16
        default:
            var v17 uint64 = 2
            _ = v17
        }
    }
    v1 = -2.25
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int16 = 134
        _ = fi18
        _ = v19
    }
    if ((v2 * v2) == v2) {
        v2 = v2
    } else {
        v3 = v3
        var v20 uint64 = 432
        v21 := []int{32767, 2147483647}
        _ = v20
        _ = v21
    }
    var v22 int16 = 32767
    if ((502 == 18446744073709551615) && false) {
        v3 -= v3
    } else {
        var v23 string = ("\n\t" + "")
        _ = v23
    }
    fmt.Println(v3, v22, v3)
    for fi24 := 0; fi24 < 10; fi24++ {
        fmt.Println(v3)
        _ = fi24
    }
    var v25 string = ""
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int8 = -32
        var v28 int32 = (-(v3))
        v29 := []int{fi26}
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        if true {
            var v31 string = v25
            if true {
                v32 := GenS1{F0: 94, F1: (fi30 & (fi30 + 32767)), F2: (-(v22))}
                _ = v32
            } else {
                if (!false) {
                    var v33 uint = 0
                    _ = v33
                } else {
                    v34 := []int{9223372036854775807, fi30, (^(fi30))}
                    _ = v34
                }
                v25 = "hello"
                v35 := []int{}
                _ = v35
            }
            _ = v31
        } else {
            if true {
                v36 := []int{-906, (fi30 - fi30)}
                v2 -= 380
                v22 -= (v22 * v22)
                _ = v36
            } else {
                var v37 int16 = (v22 / ((v22) | 1))
                _ = v37
            }
            var v38 int32 = -134
            var v39 uint8 = 255
            _ = v38
            _ = v39
        }
        v40 := GenS1{F0: 0, F1: fi30, F2: v22}
        _ = fi30
        _ = v40
    }
    var v41 int64 = -9223372036854775808
    if true {
        v42 := GenS1{F0: 140, F1: 32768, F2: -32768}
        v2 = 41
        v43 := []int{42, -2147483648, -9223372036854775808, -378}
        _ = v42
        _ = v43
    } else {
        v44 := []int{int(v22), 9223372036854775807, 9223372036854775807}
        var v45 int = 32767
        v46 := GenS2{F0: (^(v2)), F1: v25}
        _ = v44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        fmt.Println(v2)
        v41 = (v41 + v41)
        v25 = "foo bar"
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v41
}
