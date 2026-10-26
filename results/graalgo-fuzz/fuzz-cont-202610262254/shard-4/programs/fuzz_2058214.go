package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int8 }

func main() {
    var v1 uint32 = 1
    var v2 uint32 = v1
    v3 := make(map[int]bool)
    var v4 int32 = 2147483647
    var v5 uint8 = 2
    v6 := make(map[int]bool)
    v7 := v6[32767]
    switch "hello" {
    case "\n\t":
        v8 := GenS1{F0: v1}
        _ = v8
    case "a":
        v2 += v2
    }
    v7 = ((v4 > v4) || v7)
    v5 = 127
    v9 := GenS1{F0: (v1 ^ (^(v1)))}
    v10 := GenS1{F0: 1}
    var v11 string = ("hello" + ("a" + "\n\t"))
    switch v11 {
    case "":
        fmt.Println(v7, v2, v1)
    case "a":
        var v12 int32 = 1
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 4; fi13++ {
            switch "hello" {
            case "hello":
                switch v11 {
                case "":
                    for fi14 := 0; fi14 < 14; fi14++ {
                        v7 = (v11 >= v11)
                        if (0 == 0) {
                            v1 += v2
                        }
                        _ = fi14
                    }
                default:
                    var v15 int = (fi13 % ((9223372036854775807) | 1))
                    _ = v15
                }
            case "foo bar":
                if v7 {
                    v2 = 1
                } else {
                    v16 := GenS1{F0: 2147483648}
                    for fi17 := 0; fi17 < 17; fi17++ {
                        v11 = ""
                        _ = fi17
                    }
                    _ = v16
                }
            case "a":
                fmt.Println(v11, v11, v7)
            }
            v1 += (v1 % (((^(v1))) | 1))
            if false {
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 bool = v7
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 18; fi20++ {
                    v11 = v11
                    _ = fi20
                }
            }
            _ = fi13
        }
    default:
        for fi21 := 0; fi21 < 16; fi21++ {
            switch v11 {
            case "hello":
                v2 += 65535
            case "foo bar":
                if false {
                    var v22 int32 = 1
                    _ = v22
                }
            case "\n\t":
                v2 = 2
            default:
                fmt.Println(v5, v4)
            }
            v23 := GenS2{F0: 179, F1: v4}
            if (0 > v5) {
                var v24 string = (("foo bar" + v11) + (v11 + ("foo bar" + "hello")))
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 float64 = 1.5
                    v1 -= 1
                    _ = fi25
                    _ = v26
                }
                v27 := GenS3{F0: 32767, F1: (v1 * 664)}
                _ = v24
                _ = v27
            }
            _ = fi21
            _ = v23
        }
    }
    _ = len(v6)
    if v7 {
        fmt.Println(v1)
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            fmt.Println(v4)
            var v29 uint = 415
            v5 -= v5
            _ = fi28
            _ = v29
        }
        switch -9223372036854775808 {
        case 32767:
            v30 := GenS1{F0: (v2 * v1)}
            _ = v30
        case -1:
            switch -9223372036854775808 {
            case 4294967295:
                var v31 uint8 = v5
                _ = v31
            case -727:
                v11 = (v11 + (v11 + "a"))
            }
        default:
            var v32 uint = 200
            _ = v32
        }
        v33 := GenS4{F0: -9223372036854775808, F1: ""}
        _ = v33
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        v35 := GenS3{F0: -20, F1: v1}
        v36 := GenS5{F0: 1}
        _ = fi34
        _ = v35
        _ = v36
    }
    v2 = (v2 >> ((18446744073709551615) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
