package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint64 }
func genFn0() int8 {
    var v7 uint16 = 728
    v8 := []int{9223372036854775807}
    v9 := GenS1{F0: -1, F1: 255}
    if true {
        switch "" {
        case "a":
            v7 = (v7 >> ((4294967295) & 63))
        case "foo bar":
            v7 += 2
        default:
            var v10 float32 = (-2.25 * 3.14159)
            _ = v10
        }
        switch -100 {
        case -710:
            if false {
                for fi11 := 0; fi11 < 1; fi11++ {
                    if len(v8) > 0 {
                        v8[0] = -2147483648
                    }
                    v8 = append(v8, 406)
                    _ = fi11
                }
                if len(v8) > 0 {
                    v8[0] = -2147483648
                }
            } else {
                var v12 string = ""
                fmt.Println(v7)
                _ = v12
            }
        default:
            v13 := GenS2{F0: 4294967295}
            _ = v13
        }
    }
    switch 9223372036854775807 {
    case 997:
        var v14 uint16 = v7
        _ = v14
    case 127:
        if true {
            for fi15 := 0; fi15 < 18; fi15++ {
                if true {
                    switch "foo bar" {
                    case "foo bar":
                        switch fi15 {
                        case -1:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        case 127:
                            var v16 int16 = 740
                            _ = v16
                        default:
                            v7 = v7
                        }
                    case "hello":
                        _ = len(v8)
                    default:
                        v7 = 32767
                    }
                    v17 := GenS2{F0: 42}
                    _ = v17
                }
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 uint = 0
                    switch fi18 {
                    case 703:
                        v7 = (v7 - (^(v7)))
                    case 4294967295:
                        v7 = 291
                    case -9223372036854775808:
                        fmt.Println(fi15)
                    }
                    _ = fi18
                    _ = v19
                }
                _ = fi15
            }
            if (1000 >= -162) {
                fmt.Println(v7)
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint32 = 65535
                    v7 -= v7
                    _ = fi20
                    _ = v21
                }
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    var v22 int32 = -665
                    _ = v22
                case "a":
                    v7 -= (v7 * 2)
                case "\n\t":
                    var v23 uint64 = 0
                    _ = v23
                }
            }
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return -128
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -1
    var v3 int8 = 1
    v4 := make(map[int]int)
    v5 := []int{-551, -910, 42, -2147483648}
    var v6 bool = (((3.14159 + float32(-2.25)) == 3.14159) || true)
    delete(v1, (("\n\t" + ("" + "")) + ("" + ("\n\t" + ("\n\t" + "foo bar")))))
    v24 := GenS1{F0: 0, F1: 2}
    v2 += (v2 << ((844) & 63))
    v25 := GenS2{F0: 4294967295}
    var v26 uint32 = 2147483648
    var v27 uint64 = 977
    var v28 int = 9223372036854775807
    if (((v27 - (^(v27))) != 18446744073709551615) || v6) {
        var v29 int64 = v2
        fmt.Println(v26, v6)
        var v30 uint64 = 0
        _ = v29
        _ = v30
    }
    fmt.Println(v2)
    switch "a" {
    case "":
        if (v6 || (false || v6)) {
            for fi31 := 0; fi31 < 9; fi31++ {
                fmt.Println(fi31, fi31, v6)
                _ = fi31
            }
            var v32 bool = (32 == 0)
            var v33 bool = (v26 < v26)
            _ = v32
            _ = v33
        }
    case "a":
        v27 = 2
    default:
        if (!(uint(v27) >= 18446744073709551615)) {
            var v34 uint8 = 36
            _ = v34
        }
    }
    v35 := GenS1{F0: -911, F1: 1}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
