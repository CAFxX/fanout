package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint64; F2 string }

func main() {
    var v1 bool = true
    v2 := []string{"", ("\n\t" + "hello")}
    var v3 int16 = -32768
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v4 float32 = 0.0
    for fi5 := 0; fi5 < 10; fi5++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v6 uint16 = 65535
        _ = fi5
        _ = v6
    }
    if ((v1 && (!((v3 + -938) != (v3 % ((v3) | 1))))) || v1) {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: -433}
            v2 = append(v2, "a")
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 -= v3
                if (!true) {
                    fmt.Println(fi7, v1)
                } else {
                    v1 = true
                }
            } else {
                var v9 uint64 = 4294967295
                var v10 uint8 = 69
                _ = v10
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    } else {
        if len(v2) > 0 {
            v2[0] = ((("\n\t" + "") + "") + (("a" + "\n\t") + ("hello" + (("a" + "") + "foo bar"))))
        }
        v4 = v4
        var v11 uint = 0
        _ = v11
    }
    var v12 int = 4
    if v1 {
        var v13 int32 = -2147483648
        _ = v13
    } else {
        v12 = v12
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        _ = len(v2)
        for fi15 := 0; fi15 < 15; fi15++ {
            if (!(!false)) {
                v4 += float32(-2.25)
            } else {
                if true {
                    v16 := GenS2{F0: (v3 ^ (v3 + 1000)), F1: 793, F2: ("a" + (("\n\t" + "\n\t") + "hello"))}
                    _ = v16
                }
            }
            _ = fi15
        }
        _ = fi14
    }
    var v17 int16 = v3
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v18 int8 = 42
    if v1 {
        for fi19 := 0; fi19 < 20; fi19++ {
            for fi20 := 0; fi20 < 2; fi20++ {
                if (((2147483647 <= 744) && (832 > 874)) || (538 >= 900)) {
                    var v21 uint64 = 491
                    _ = v21
                }
                v17 = 32767
                switch -905 {
                case 42:
                    var v22 int8 = (v18 | (v18 | (v18 % ((v18) | 1))))
                    _ = v22
                case 9223372036854775807:
                    v4 = v4
                case 532:
                    v2 = append(v2, "a")
                default:
                    v23 := GenS2{F0: 42, F1: 4294967295, F2: "foo bar"}
                    _ = v23
                }
                _ = fi20
            }
            var v24 int8 = v18
            fmt.Println(fi19, fi19)
            _ = fi19
            _ = v24
        }
    } else {
        v17 = 42
    }
    fmt.Println(v17, v3)
    var v25 uint = 42
    for fi26 := 0; fi26 < 14; fi26++ {
        v1 = v1
        fmt.Println(v3, v4)
        v2 = append(v2, "hello")
        _ = fi26
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
