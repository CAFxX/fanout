package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 string; F2 int8 }

func main() {
    var v1 string = ""
    var v2 int32 = -94
    var v3 int64 = 1
    v4 := make(map[uint32]string)
    var v5 uint8 = 127
    v4[2] = v1
    fmt.Println(v2, v5, v3)
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 134, F1: uint(v5)}
            _ = fi7
            _ = v8
        }
        var v9 int32 = 802
        _ = fi6
        _ = v9
    }
    v1 = v1
    switch 128 {
    case 42:
        v10 := GenS2{F0: 2, F1: 18446744073709551615}
        _ = v10
    case 127:
        v11 := GenS3{F0: (v5 + 128)}
        _ = v11
    default:
        var v12 int32 = v2
        _ = v12
    }
    if (!(0.1 <= 1e10)) {
        if ((v2 % (((v2 + v2)) | 1)) > 42) {
            fmt.Println(v1, v5)
            var v13 float32 = 0.0
            _ = v13
        }
        v14 := []int{128, 4294967295}
        _ = v14
    } else {
        v15 := GenS4{F0: 2147483648, F1: 42, F2: (v3 - (v3 & (v3 + v3)))}
        if false {
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int16 = -1
                _ = fi16
                _ = v17
            }
        } else {
            v1 = "a"
            v5 -= (v5 / (((v5 >> ((4294967295) & 63))) | 1))
        }
        _ = v15
    }
    v18 := GenS2{F0: 18446744073709551615, F1: 1}
    v2 = v2
    var v19 string = (("hello" + "") + (((v1 + v1) + "\n\t") + v1))
    delete(v4, 2147483647)
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS3{F0: 166}
        switch 42 {
        case -728:
            fmt.Println(v1)
        case -834:
            switch v19 {
            case "foo bar":
                delete(v4, 2147483648)
            default:
                switch v19 {
                case "foo bar":
                    switch (fi20 / ((fi20) | 1)) {
                    case 212:
                        v22 := []int{(fi20 ^ (fi20 >> ((155) & 63))), 664, fi20, (fi20 ^ fi20)}
                        _ = v22
                    default:
                        if true {
                            delete(v4, 2147483648)
                            v23 := GenS5{F0: (v5 & 115)}
                            fmt.Println(v1)
                            _ = v23
                        } else {
                            var v24 uint16 = 122
                            _ = v24
                        }
                    }
                }
            }
        case 1:
            if (697 >= 19) {
                v4[648] = "\n\t"
                var v25 bool = true
                var v26 uint64 = 0
                _ = v25
                _ = v26
            } else {
                for fi27 := 0; fi27 < 5; fi27++ {
                    if (!((0.1 >= 1e10) || (42 != 26))) {
                        delete(v4, 80)
                        switch (-(fi20)) {
                        case 2147483647:
                            var v28 int16 = 490
                            _ = v28
                        }
                    } else {
                        v29 := GenS1{F0: 42, F1: 0}
                        fmt.Println(v5)
                        v19 = "hello"
                        _ = v29
                    }
                    switch (fi27 >> ((65535) & 63)) {
                    case 42:
                        v30 := GenS6{F0: (true && (("foo bar" + v19) == (v19 + v1))), F1: v1, F2: int8(v5)}
                        v2 = v2
                        v31 := []int{}
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                }
            }
        }
        var v32 float32 = float32((0.1 / (1.5 + 1e10)))
        _ = fi20
        _ = v21
        _ = v32
    }
    if false {
        fmt.Println(v2)
    } else {
        var v33 uint32 = 2
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
