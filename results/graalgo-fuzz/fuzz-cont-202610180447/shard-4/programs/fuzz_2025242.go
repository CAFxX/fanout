package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint16; F1 bool; F2 int }
func genFn0(p0 uint8, p1 bool) int16 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int16 = 1000
        _ = fi6
        _ = v7
    }
    if p1 {
        for fi8 := 0; fi8 < 11; fi8++ {
            if p1 {
                v9 := []int{-44, 9223372036854775807, (fi8 >> ((289) & 63)), (fi8 % ((fi8) | 1))}
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 bool = false
                    v11 = ((32767 < -1) || (uint16(p0) >= 465))
                    if ((^(p0)) < 42) {
                        var v12 int32 = 2147483647
                        var v13 int32 = v12
                        _ = v12
                        _ = v13
                    }
                    if true {
                        p1 = ((p1 && v11) && p1)
                        fmt.Println(p0, p1, v11)
                    }
                    _ = fi10
                    _ = v11
                }
                if p1 {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if (!p1) {
                        v14 := GenS1{F0: 9223372036854775807}
                        if p1 {
                            var v15 string = "foo bar"
                            p0 = p0
                            var v16 float32 = 0.1
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    }
                    p0 += (p0 ^ p0)
                }
                _ = v9
            } else {
                v17 := GenS1{F0: 1}
                for fi18 := 0; fi18 < 12; fi18++ {
                    if p1 {
                        var v19 uint = 1
                        if ((0.0 + 3.14159) <= 0.0) {
                            p1 = ((v19 - v19) <= 0)
                            v19 -= (v19 << ((2) & 63))
                            var v20 int8 = 28
                            _ = v20
                        }
                        switch (fi18 ^ 9223372036854775807) {
                        case 369:
                            v21 := []int{(-(fi8)), fi8, fi18, fi18}
                            _ = v21
                        default:
                            fmt.Println(v19, fi18, p0)
                            v19 = v19
                            v22 := GenS2{F0: true, F1: 255, F2: 184}
                            _ = v22
                        }
                        p1 = (32767 <= 255)
                        _ = v19
                    }
                    _ = fi18
                }
                _ = v17
            }
            v23 := GenS1{F0: (fi8 & (-(fi8)))}
            p1 = p1
            _ = fi8
            _ = v23
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            fmt.Println(p1)
        case "hello":
            v24 := []int{int(p0), 32768}
            _ = v24
        case "a":
            var v25 uint8 = (p0 * (p0 / (((p0 - 0)) | 1)))
            _ = v25
        default:
            fmt.Println(p1, p1)
        }
    }
    v26 := GenS2{F0: false, F1: 88, F2: 2}
    _ = p0
    _ = p1
    _ = v26
    return 0
}

func main() {
    var v1 int32 = 0
    var v2 int64 = 854
    var v3 uint64 = 2
    var v4 uint8 = 2
    var v5 int32 = (^(v1))
    if false {
        v5 -= 0
    }
    var v27 uint32 = 4294967295
    var v28 int = -9223372036854775808
    if true {
        v4 = (v4 % (((v4 + (v4 << ((65535) & 63)))) | 1))
        var v29 int = (v28 / (((v28 % (((v28 << ((713) & 63))) | 1))) | 1))
        _ = v29
    }
    if true {
        var v30 bool = (true && ((^(v4)) == (v4 & v4)))
        for fi31 := 0; fi31 < 13; fi31++ {
            for fi32 := 0; fi32 < 19; fi32++ {
                var v33 int16 = genFn0(v4, (!v30))
                v33 = -32768
                switch "hello" {
                case "hello":
                    v34 := GenS1{F0: (-(fi32))}
                    _ = v34
                default:
                    v35 := GenS3{F0: 0, F1: v30, F2: (v28 << ((42) & 63))}
                    _ = v35
                }
                _ = fi32
                _ = v33
            }
            if v30 {
                v30 = v30
            }
            _ = fi31
        }
        _ = v30
    } else {
        switch "\n\t" {
        case "\n\t":
            v36 := GenS3{F0: 1, F1: false, F2: 4294967295}
            _ = v36
        default:
            var v37 string = (("hello" + "a") + "")
            _ = v37
        }
    }
    fmt.Println(v3, v5, v28)
    var v38 float32 = 3.14159
    fmt.Println(v38)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
