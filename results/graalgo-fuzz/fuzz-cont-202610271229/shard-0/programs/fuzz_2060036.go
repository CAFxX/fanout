package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 65535
    var v3 uint32 = v2
    var v4 int64 = 1
    v5 := make(map[string]int)
    _ = len(v5)
    switch 127 {
    case 127:
        fmt.Println(v2, v4)
    case -9223372036854775808:
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(fi6)
            var v7 float32 = 3.14159
            _ = fi6
            _ = v7
        }
    }
    v4 = (v4 % ((v4) | 1))
    var v8 bool = (((-1 <= 975) && (2 != 910)) && false)
    switch "a" {
    case "foo bar":
        v8 = ((-419 >= 42) || ((32767 <= 1000) || (4294967295 <= 2)))
    case "a":
        v2 = 1
    case "\n\t":
        fmt.Println(v3, v8)
    }
    v3 = v2
    v9 := GenS1{F0: 2147483647, F1: -1, F2: v4}
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            if (fi10 != fi11) {
                v2 = v2
                switch (-(fi10)) {
                case 127:
                    v12 := GenS2{F0: 310}
                    _ = v12
                case 2147483647:
                    v3 = (^(v3))
                }
            }
            v4 = (v4 + v4)
            delete(v5, ((("a" + "\n\t") + (("a" + "a") + "hello")) + ("hello" + "")))
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v4, v8)
            default:
                if v8 {
                    fmt.Println(v8)
                    v3 = v3
                }
            }
            _ = fi11
        }
        _ = fi10
    }
    v5["hello"] = -9223372036854775808
    switch -1 {
    case 2147483647:
        v8 = v8
    case -1:
        switch 4294967295 {
        case 372:
            v3 -= v2
        case 2147483647:
            switch ("foo bar" + ("foo bar" + "foo bar")) {
            case "\n\t":
                for fi13 := 0; fi13 < 17; fi13++ {
                    if v8 {
                        v14 := []int{(^(fi13)), fi13, fi13, fi13}
                        fmt.Println(fi13)
                        _ = v14
                    }
                    _ = fi13
                }
            case "hello":
                v3 -= v3
            case "foo bar":
                for fi15 := 0; fi15 < 11; fi15++ {
                    if ((!(259 > 1000)) || v8) {
                        for fi16 := 0; fi16 < 9; fi16++ {
                            if v8 {
                                fmt.Println(v8, fi16)
                                var v17 uint32 = 382
                                var v18 bool = ((1.5 * float64(float32(1e10))) != (float64(0.0) - -2.25))
                                v8 = (true && (104 == 127))
                                _ = v17
                                _ = v18
                            }
                            _ = len(v5)
                            var v19 uint = 651
                            _ = fi16
                            _ = v19
                        }
                    } else {
                        if (true || ((0 < 333) && true)) {
                            v8 = v8
                        }
                        fmt.Println(v4, fi15, v4)
                    }
                    for fi20 := 0; fi20 < 11; fi20++ {
                        v21 := []int{fi15, 0, 32767, fi20}
                        v8 = ((((3.14159 - float32(3.14159)) * 0.0) + (3.14159 - 0.0)) > -2.25)
                        v4 = 699
                        _ = fi20
                        _ = v21
                    }
                    _ = fi15
                }
            default:
                if v8 {
                    v22 := GenS3{F0: 4294967295, F1: 0}
                    v3 = (v2 & (^(v3)))
                    v23 := GenS2{F0: -2147483648}
                    _ = v22
                    _ = v23
                }
            }
        case 32767:
            switch -2147483648 {
            case -1:
                var v24 int32 = 90
                _ = v24
            }
        }
    }
    var v25 int16 = -32768
    var v26 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
