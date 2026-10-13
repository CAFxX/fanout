package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 int16, p1 uint8) bool {
    p1 -= 128
    switch 4294967295 {
    case 42:
        var v5 uint = 4294967295
        _ = v5
    case 127:
        switch 669 {
        case -1:
            for fi6 := 0; fi6 < 5; fi6++ {
                fmt.Println(p1, fi6)
                var v7 uint8 = 255
                switch fi6 {
                case 4294967295:
                    v8 := GenS1{F0: 2147483648, F1: 1}
                    _ = v8
                }
                _ = fi6
                _ = v7
            }
        default:
            if true {
                v9 := []int{32768, -2147483648}
                _ = v9
            }
        }
    default:
        fmt.Println(p0)
    }
    v10 := []int{0, 886, -9223372036854775808, 127}
    var v11 int8 = -1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return (65535 >= 2)
}

func main() {
    var v1 int = -256
    v2 := []string{((((("\n\t" + (("\n\t" + "hello") + "hello")) + "foo bar") + "hello") + "\n\t") + ""), ("" + ((("a" + "a") + "a") + "foo bar")), ""}
    var v3 uint8 = 0
    v4 := []string{("" + ("hello" + "")), ("a" + ("hello" + ("" + "hello"))), "foo bar"}
    switch v1 {
    case 524:
        var v12 int16 = -1
        _ = v12
    case 0:
        for fi13 := 0; fi13 < 12; fi13++ {
            _ = len(v4)
            var v14 float32 = 0.1
            _ = fi13
            _ = v14
        }
    case -9223372036854775808:
        if (true || false) {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            v15 := GenS1{F0: 1, F1: -14}
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(fi16)
            v1 -= fi16
            _ = fi16
        }
    }
    var v17 uint8 = v3
    var v18 uint = 0
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint8 = 255
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 65535, F1: -87}
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v18)
        for fi23 := 0; fi23 < 7; fi23++ {
            switch ("a" + "foo bar") {
            case "a":
                fmt.Println(fi22)
            case "":
                switch "\n\t" {
                case "a":
                    switch "" {
                    case "hello":
                        v17 -= 127
                    case "\n\t":
                        v17 -= v3
                    case "a":
                        v18 = 255
                    default:
                        _ = len(v4)
                    }
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = ((("\n\t" + "a") + "hello") + ("" + "hello"))
                    }
                case "hello":
                    v2 = append(v2, "hello")
                default:
                    v2 = append(v2, ((((("hello" + "a") + "\n\t") + "a") + "\n\t") + "a"))
                }
            default:
                v18 -= v18
            }
            _ = fi23
        }
        if (149 < 239) {
            var v24 uint8 = v17
            _ = v24
        } else {
            switch (fi22 + -9223372036854775808) {
            case 0:
                for fi25 := 0; fi25 < 3; fi25++ {
                    var v26 int16 = -32768
                    _ = fi25
                    _ = v26
                }
            case 128:
                var v27 uint32 = 0
                _ = v27
            case 127:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                switch ((("hello" + ("hello" + ("foo bar" + "foo bar"))) + ("\n\t" + "a")) + "\n\t") {
                case "hello":
                    v2 = append(v2, "")
                }
            }
            var v28 float32 = 1.5
            _ = v28
        }
        _ = fi22
    }
    switch v1 {
    case 32767:
        var v29 uint32 = 949
        _ = v29
    default:
        for fi30 := 0; fi30 < 2; fi30++ {
            for fi31 := 0; fi31 < 6; fi31++ {
                v32 := GenS1{F0: 886, F1: 84}
                fmt.Println(fi30, v1, fi30)
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
    }
    v2 = append(v2, "\n\t")
    v17 = (v17 + (v17 * (v17 << ((18446744073709551615) & 63))))
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
