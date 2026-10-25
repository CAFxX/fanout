package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16; F1 int64; F2 string }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int16 = 177
    var v2 int8 = -16
    v3 := []string{"hello"}
    if false {
        for fi4 := 0; fi4 < 5; fi4++ {
            for fi5 := 0; fi5 < 2; fi5++ {
                v6 := GenS1{F0: (fi5 | fi4), F1: v1}
                var v7 uint8 = 0
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = fi4
        }
        switch -9223372036854775808 {
        case 1:
            var v8 uint8 = 128
            _ = v8
        case -759:
            _ = len(v3)
        case -700:
            var v9 uint8 = 1
            _ = v9
        default:
            var v10 int64 = -9223372036854775808
            _ = v10
        }
        if ((1e10 - (3.14159 * -2.25)) < ((0.0 / (0.1 * (float32((-2.25 - 3.14159)) * float32(float64(((-2.25 * 1.5) / 1.0)))))) * (float32((1e10 + -2.25)) + -2.25))) {
            fmt.Println(v1, v2)
        } else {
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v2 += 100
                    _ = fi12
                }
                _ = fi11
            }
        }
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int32 = -2147483648
            v15 := GenS2{F0: 1, F1: -758, F2: "\n\t"}
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            _ = len(v3)
            if ((!true) && false) {
                v2 = 127
            } else {
                v1 = v1
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                if ((((1.5 * 1.5) * (3.14159 * -2.25)) * 0.1) <= (((0.1 * float32(0.1)) - 3.14159) - float32(1.5))) {
                    fmt.Println(fi16, v1, fi16)
                }
            }
            _ = fi16
        }
        v2 += (^(v2))
    }
    fmt.Println(v1, v1)
    var v17 uint16 = 255
    if (false || true) {
        if len(v3) > 0 {
            v3[0] = "a"
        }
    }
    v17 += 708
    var v18 uint16 = (^(v17))
    for fi19 := 0; fi19 < 10; fi19++ {
        v18 += (v17 / ((v17) | 1))
        v2 = (-(v2))
        _ = fi19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        if (false && (2147483648 > 4294967295)) {
            if (2 == 4294967295) {
                var v21 float64 = (0.0 + 0.1)
                _ = v21
            }
            switch "" {
            case "":
                for fi22 := 0; fi22 < 1; fi22++ {
                    v23 := GenS3{F0: v2}
                    v17 -= v17
                    if (fi20 > (fi20 | 127)) {
                        var v24 string = ((("a" + ("a" + "hello")) + "hello") + "foo bar")
                        fmt.Println(fi20, v1, fi20)
                        if (true && (33 >= 979)) {
                            v3 = append(v3, v24)
                        }
                        _ = v24
                    }
                    _ = fi22
                    _ = v23
                }
            case "foo bar":
                var v25 uint32 = 42
                _ = v25
            }
        }
        v1 = (v1 * 42)
        _ = fi20
    }
    if true {
        for fi26 := 0; fi26 < 11; fi26++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v17 = 2
            fmt.Println(v18, v17, v2)
            _ = fi26
        }
        if true {
            var v27 int16 = v1
            v3 = append(v3, ("hello" + "foo bar"))
            v17 -= (v17 - (v18 | 165))
            _ = v27
        } else {
            v28 := GenS2{F0: 65535, F1: -850, F2: ("\n\t" + "a")}
            v17 -= (v18 + 0)
            if len(v3) > 0 {
                v3[0] = ("\n\t" + "foo bar")
            }
            _ = v28
        }
        fmt.Println(v1, v2)
    } else {
        var v29 uint8 = 149
        _ = v29
    }
    fmt.Println(v17, v17, v2)
    fmt.Println(v18, v1, v17)
    v1 = 685
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
