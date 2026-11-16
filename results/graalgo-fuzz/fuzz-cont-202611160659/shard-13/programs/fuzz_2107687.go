package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 int16 = 715
    var v2 float32 = 1e10
    var v3 int16 = v1
    v4 := make(map[uint32]string)
    var v5 int16 = v3
    var v6 int32 = 154
    if true {
        switch ("foo bar" + ("a" + "")) {
        case "":
            for fi7 := 0; fi7 < 13; fi7++ {
                if false {
                    v4[29] = ("foo bar" + "hello")
                    v2 -= v2
                    var v8 int16 = v5
                    _ = v8
                } else {
                    var v9 int64 = 9223372036854775807
                    v10 := v4[65535]
                    _ = v10
                    _ = v9
                }
                v3 -= (v3 & 42)
                var v11 int64 = -9223372036854775808
                _ = fi7
                _ = v11
            }
        case "\n\t":
            v12 := GenS1{F0: (^(v1))}
            _ = v12
        }
    }
    var v13 uint = 1
    v14 := GenS1{F0: v5}
    v5 += (^(v1))
    var v15 bool = (false && false)
    var v16 uint8 = 128
    v1 = (v1 * v3)
    if (!false) {
        switch 128 {
        case 303:
            v17 := []int{207}
            _ = v17
        default:
            v18 := []int{172, -474, 32767}
            _ = v18
        }
        v16 = 42
    }
    v5 = v3
    var v19 uint8 = v16
    v20 := GenS2{F0: (65535 > uint(v16))}
    var v21 int32 = v6
    if (1 != 255) {
        for fi22 := 0; fi22 < 3; fi22++ {
            switch (fi22 >> ((18446744073709551615) & 63)) {
            case -151:
                if (119 != 4294967295) {
                    if false {
                        v21 = v21
                        v23 := GenS2{F0: false}
                        _ = v23
                    } else {
                        v24 := GenS3{F0: -1, F1: ((((("foo bar" + "") + "\n\t") + "\n\t") + "\n\t") + ("hello" + "hello"))}
                        v25 := []int{(fi22 & (^(fi22)))}
                        v1 = (v3 | -1)
                        _ = v24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v27 := GenS2{F0: v15}
                        v28 := GenS4{F0: 983, F1: 198}
                        var v29 float64 = -2.25
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                } else {
                    v30 := []int{(fi22 | fi22), 307}
                    v21 = -1
                    switch "foo bar" {
                    case "a":
                        if true {
                            fmt.Println(v5, v19)
                            var v31 uint64 = 2
                            _ = len(v4)
                            _ = v31
                        } else {
                            for fi32 := 0; fi32 < 3; fi32++ {
                                var v33 int8 = 69
                                _ = fi32
                                _ = v33
                            }
                            delete(v4, 2147483648)
                        }
                    default:
                        v3 = 913
                    }
                    _ = v30
                }
            case 333:
                v34 := []int{}
                _ = v34
            case -1:
                v35 := GenS2{F0: v15}
                _ = v35
            }
            _ = fi22
        }
        v36 := []int{}
        _ = v36
    }
    if (v3 >= (v5 + (v5 + 32767))) {
        fmt.Println(v21, v19, v6)
    } else {
        v6 = v6
    }
    _ = len(v4)
    fmt.Println(v21, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
