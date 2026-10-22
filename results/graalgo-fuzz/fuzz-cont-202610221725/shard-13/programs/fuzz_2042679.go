package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }

func main() {
    var v1 float64 = (1e10 + 1.5)
    v2 := make(map[int]uint32)
    var v3 float64 = (v1 * v1)
    v4 := make(map[int]string)
    var v5 float32 = (3.14159 / 3.14159)
    var v6 float64 = v1
    v7 := GenS1{F0: 398}
    v5 -= 1e10
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int64 = -1
        v10 := GenS1{F0: 0}
        for fi11 := 0; fi11 < 4; fi11++ {
            switch fi8 {
            case 42:
                if false {
                    var v12 int8 = int8(v9)
                    switch (fi11 / ((-8) | 1)) {
                    case 127:
                        v12 = (^(v12))
                    case -9223372036854775808:
                        v6 = v1
                    default:
                        var v13 uint32 = 2147483648
                        _ = v13
                    }
                    v14 := GenS1{F0: 18446744073709551615}
                    _ = v12
                    _ = v14
                }
            case -9223372036854775808:
                var v15 string = "\n\t"
                _ = v15
            case 1:
                v1 = v3
                fmt.Println(v9, fi8)
                var v16 uint32 = 42
                if true {
                    var v17 uint8 = 163
                    switch -9223372036854775808 {
                    case 127:
                        var v18 string = "a"
                        _ = v18
                    case 2147483647:
                        v17 = v17
                    case -9223372036854775808:
                        var v19 int64 = v9
                        _ = v19
                    }
                    _ = v17
                } else {
                    var v20 uint64 = 1
                    if (309 >= (v20 & v20)) {
                        var v21 int16 = 42
                        _ = v21
                    } else {
                        switch "hello" {
                        case "a":
                            v16 = v16
                        default:
                            fmt.Println(fi11)
                        }
                        switch "hello" {
                        case "":
                            v5 = 3.14159
                        case "a":
                            switch "" {
                            case "foo bar":
                                v9 += (^(v9))
                            case "":
                                var v22 uint64 = (v20 + v20)
                                _ = v22
                            default:
                                fmt.Println(v6, v3)
                            }
                        }
                    }
                    v23 := GenS1{F0: 42}
                    _ = v20
                    _ = v23
                }
                v9 += 1
                _ = v16
            }
            fmt.Println(fi8, v6, v9)
            if true {
                v24 := GenS2{F0: 2, F1: -749, F2: 4294967295}
                v3 += -2.25
                var v25 int32 = 0
                _ = v24
                _ = v25
            }
            _ = fi11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = v6
    fmt.Println(v1)
    v26 := GenS1{F0: 255}
    v3 += v3
    v6 = v1
    switch (("" + "foo bar") + ("foo bar" + "foo bar")) {
    case "":
        v1 -= 0.1
    case "hello":
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 string = (("\n\t" + ("hello" + ("hello" + "\n\t"))) + "foo bar")
            fmt.Println(v3)
            var v29 uint64 = 16
            _ = fi27
            _ = v28
            _ = v29
        }
    default:
        v30 := GenS1{F0: 210}
        _ = v30
    }
    v31 := GenS3{F0: 2147483648, F1: true, F2: 91}
    fmt.Println(v5, v3)
    if (40 >= 9223372036854775807) {
        if (-128 != 2) {
            fmt.Println(v1)
        } else {
            v3 = 1.5
            v3 = (1.5 - ((3.14159 + v3) + 0.0))
        }
    } else {
        v5 = 0.0
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 17; fi33++ {
            v5 = v5
            delete(v2, fi32)
            _ = fi33
        }
        delete(v4, fi32)
        _ = fi32
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v3 = ((-2.25 / (v3 - 0.1)) + 0.0)
        _ = fi34
    }
    var v35 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
