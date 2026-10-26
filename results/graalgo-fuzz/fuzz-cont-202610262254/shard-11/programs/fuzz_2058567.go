package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string) uint32 {
    v4 := []int{-477, -1, 772}
    v4 = append(v4, 127)
    var v5 int16 = 42
    switch 694 {
    case 4294967295:
        fmt.Println(p0, p0)
    case -2147483648:
        var v6 bool = (!true)
        _ = v6
    case 127:
        fmt.Println(v5, p0)
    default:
        for fi7 := 0; fi7 < 12; fi7++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v8 string = ""
            v8 = p0
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() int {
    v9 := []int{-976, -457}
    var v10 uint = 424
    v10 = v10
    _ = v10
    _ = v9
    return 394
}

func main() {
    var v1 int8 = -111
    var v2 bool = true
    var v3 float32 = -2.25
    var v11 uint64 = 0
    switch -600 {
    case -839:
        v11 += (^(v11))
    default:
        v12 := GenS1{F0: 955, F1: 2147483647}
        _ = v12
    }
    v13 := GenS2{F0: 2147483648}
    v14 := []int{-2147483648, 127}
    for fi15 := 0; fi15 < 4; fi15++ {
        v11 = (v11 >> ((65535) & 63))
        v16 := GenS2{F0: 65535}
        _ = fi15
        _ = v16
    }
    v3 += v3
    var v17 uint16 = 42
    for fi18 := 0; fi18 < 6; fi18++ {
        switch "hello" {
        case "hello":
            switch "\n\t" {
            case "a":
                for fi19 := 0; fi19 < 6; fi19++ {
                    if len(v14) > 0 {
                        v14[0] = (-(fi18))
                    }
                    v1 -= (v1 % (((v1 ^ v1)) | 1))
                    _ = fi19
                }
            default:
                var v20 uint = 255
                _ = v20
            }
        case "foo bar":
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        default:
            if false {
                v14 = append(v14, (fi18 / ((genFn1()) | 1)))
                if (100 > 61) {
                    switch ((("a" + ("foo bar" + "foo bar")) + ("\n\t" + (("\n\t" + ((((("" + "") + "") + ("hello" + (((("\n\t" + "hello") + "a") + "a") + ""))) + "foo bar") + "foo bar")) + ("a" + (("a" + "foo bar") + "\n\t"))))) + ("foo bar" + "foo bar")) {
                    case "foo bar":
                        v21 := GenS2{F0: genFn0("\n\t")}
                        _ = v21
                    }
                }
                var v22 uint16 = uint16(v11)
                _ = v22
            } else {
                v11 += v11
                if (!("hello" < "a")) {
                    var v23 uint8 = 42
                    var v24 int = fi18
                    _ = v23
                    _ = v24
                }
            }
        }
        v1 += 127
        _ = fi18
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint32 = 2
        var v27 uint8 = 76
        v1 -= v1
        _ = fi25
        _ = v26
        _ = v27
    }
    if (!true) {
        switch 32767 {
        case 1:
            if len(v14) > 0 {
                v14[0] = 4294967295
            }
        case -84:
            v28 := GenS2{F0: 1}
            _ = v28
        default:
            var v29 float32 = v3
            _ = v29
        }
        v30 := GenS2{F0: 521}
        v3 = float32(3.14159)
        _ = v30
    }
    var v31 string = (("\n\t" + "a") + ("hello" + ("foo bar" + "foo bar")))
    v31 = ((((v31 + ("a" + "foo bar")) + v31) + (v31 + "foo bar")) + "")
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v31
}
