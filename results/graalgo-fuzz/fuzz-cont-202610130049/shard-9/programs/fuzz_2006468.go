package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint64; F2 int8 }
type GenS6 struct { F0 int16 }
func genFn0() int32 {
    var v4 int32 = -2147483648
    var v5 int32 = -2147483648
    v4 -= v4
    fmt.Println(v4, v4)
    v6 := GenS1{F0: 42, F1: 18446744073709551615, F2: 0}
    _ = v4
    _ = v5
    _ = v6
    return (v4 & 1)
}

func main() {
    v1 := []string{}
    var v2 float32 = float32(0.0)
    v3 := []string{"a", "", ("\n\t" + (("a" + "hello") + ("hello" + "")))}
    _ = len(v1)
    var v7 int = -474
    switch 32768 {
    case 355:
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    case 86:
        var v8 int8 = 9
        _ = v8
    default:
        v7 -= 1
    }
    v1 = append(v1, "\n\t")
    fmt.Println(v7, v7)
    fmt.Println(v2, v7)
    if true {
        v9 := GenS1{F0: 65535, F1: 147, F2: (v7 >> ((710) & 63))}
        _ = v9
    }
    v1 = append(v1, "foo bar")
    v2 = (v2 + v2)
    v10 := GenS2{F0: 1}
    switch "" {
    case "":
        v11 := GenS3{F0: 18446744073709551615, F1: -9223372036854775808, F2: "foo bar"}
        _ = v11
    case "hello":
        var v12 int16 = -933
        _ = v12
    }
    fmt.Println(v7, v2)
    v2 -= 1.5
    var v13 int16 = 462
    if (true || false) {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        v14 := GenS4{F0: (v7 & (v7 & (v7 & (^(v7)))))}
        _ = v14
    } else {
        fmt.Println(v7)
    }
    if ((!(9223372036854775807 >= v7)) && true) {
        v13 = 1
        v1 = append(v1, "")
    } else {
        var v15 uint16 = 2
        var v16 uint8 = 128
        for fi17 := 0; fi17 < 11; fi17++ {
            v15 += v15
            switch (fi17 & v7) {
            case 32768:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 0:
                v13 += (v13 << ((394) & 63))
            case 1:
                v7 = v7
            }
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    if false {
        if true {
            switch (v7 >> ((2) & 63)) {
            case 449:
                var v18 uint64 = 946
                _ = v18
            case -9223372036854775808:
                v13 = -1
            case 4294967295:
                switch v7 {
                case -9223372036854775808:
                    fmt.Println(v13, v2)
                case 524:
                    if (1 > 1) {
                        v3 = append(v3, (("a" + "\n\t") + ""))
                    }
                case -1:
                    v1 = append(v1, "foo bar")
                default:
                    for fi19 := 0; fi19 < 15; fi19++ {
                        switch ("foo bar" + ("foo bar" + ("\n\t" + "hello"))) {
                        case "a":
                            var v20 uint8 = 171
                            _ = v20
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v13)
                        _ = fi19
                    }
                }
            default:
                v21 := GenS3{F0: 1, F1: 1, F2: (("a" + "\n\t") + "\n\t")}
                _ = v21
            }
            v22 := GenS5{F0: true, F1: 0, F2: -114}
            switch "foo bar" {
            case "\n\t":
                v2 -= v2
            }
            _ = v22
        }
        switch v7 {
        case 42:
            v13 = (v13 | (^(v13)))
        case 139:
            v23 := GenS6{F0: v13}
            _ = v23
        default:
            v7 += (v7 % ((0) | 1))
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            if ((!(((1.5 - 0.1) != 0.0) && true)) || (-583 == -386)) {
                v25 := GenS6{F0: (v13 >> ((65535) & 63))}
                if (-2147483648 != genFn0()) {
                    _ = len(v3)
                    for fi26 := 0; fi26 < 9; fi26++ {
                        fmt.Println(fi24, v13, v7)
                        _ = fi26
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    var v27 uint16 = 32767
                    _ = v27
                }
                _ = v25
            }
            _ = fi24
        }
    } else {
        for fi28 := 0; fi28 < 18; fi28++ {
            if true {
                v29 := GenS2{F0: 666}
                _ = v29
            }
            _ = fi28
        }
    }
    v3 = append(v3, "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v7
}
