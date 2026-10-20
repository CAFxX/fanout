package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 int8; F1 bool; F2 uint32 }

func main() {
    var v1 int16 = -892
    var v2 uint16 = 32767
    var v3 uint16 = (v2 >> ((4294967295) & 63))
    var v4 uint64 = 802
    var v5 uint16 = 2
    v6 := []string{("\n\t" + ("foo bar" + "")), ("a" + "hello"), (("a" + "") + "")}
    if (true || false) {
        _ = len(v6)
    } else {
        var v7 int16 = 1
        v2 = v2
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v7
    }
    if true {
        fmt.Println(v2)
        switch 127 {
        case -2147483648:
            for fi8 := 0; fi8 < 16; fi8++ {
                switch fi8 {
                case 4294967295:
                    for fi9 := 0; fi9 < 17; fi9++ {
                        v4 += v4
                        if (!(true && true)) {
                            v1 -= (v1 - -32768)
                            v2 = v5
                        }
                        _ = fi9
                    }
                case 774:
                    v10 := GenS1{F0: (-(v1)), F1: 107}
                    _ = v10
                case 42:
                    for fi11 := 0; fi11 < 16; fi11++ {
                        fmt.Println(v5)
                        if true {
                            v12 := GenS2{F0: 100, F1: 0, F2: true}
                            _ = v12
                        }
                        _ = fi11
                    }
                default:
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v5, v2)
                    default:
                        v5 += v2
                    }
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                fmt.Println(v1, v2)
                _ = fi8
            }
        case 9223372036854775807:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case 128:
            if false {
                var v13 uint = 65535
                _ = v13
            } else {
                var v14 int16 = (v1 & -32768)
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 4; fi15++ {
                for fi16 := 0; fi16 < 10; fi16++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v17 := GenS1{F0: (v1 % ((v1) | 1)), F1: 42}
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
        }
        v5 = v2
    } else {
        if (!true) {
            v5 = v2
            var v18 bool = (32767 < 42)
            for fi19 := 0; fi19 < 3; fi19++ {
                fmt.Println(v18, fi19, v1)
                _ = fi19
            }
            _ = v18
        }
        fmt.Println(v1)
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v1 -= 1000
        }
    }
    var v20 string = "a"
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS1{F0: (v1 + (v1 ^ (^(v1)))), F1: 100}
        var v23 string = "a"
        _ = fi21
        _ = v22
        _ = v23
    }
    v6 = append(v6, "")
    v24 := GenS3{F0: 100, F1: ((v1 < v1) && true), F2: 65535}
    var v25 bool = (!(302 != 18446744073709551615))
    var v26 int64 = 9223372036854775807
    var v27 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
