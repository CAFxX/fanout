package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0(p0 uint8, p1 uint16, p2 bool) string {
    p0 = p0
    p2 = (864 != 2)
    _ = p0
    _ = p1
    _ = p2
    return (("foo bar" + ((("\n\t" + "") + "") + "\n\t")) + "a")
}

func main() {
    var v1 int16 = -32768
    var v2 int = 956
    var v3 uint = 1
    fmt.Println(v3)
    v4 := []int{-260, v2}
    for fi5 := 0; fi5 < 3; fi5++ {
        v1 = v1
        v4 = append(v4, (fi5 % (((v2 + (v2 >> (((v3 * v3)) & 63)))) | 1)))
        switch 127 {
        case 42:
            v3 = (v3 + v3)
        default:
            var v6 int8 = 127
            _ = v6
        }
        _ = fi5
    }
    switch ("" + (("hello" + ("" + "")) + "a")) {
    case "foo bar":
        v7 := GenS1{F0: -59}
        _ = v7
    case "a":
        var v8 float32 = -2.25
        _ = v8
    case "hello":
        v3 = v3
    default:
        v2 -= (v2 ^ (v2 * (v2 ^ -9223372036854775808)))
    }
    var v9 int8 = 13
    var v10 int32 = -578
    var v11 bool = false
    v10 -= v10
    v1 -= (v1 ^ v1)
    if v11 {
        if v11 {
            v2 = (v2 + -2147483648)
            var v12 uint16 = 792
            v4 = append(v4, -66)
            _ = v12
        } else {
            if v11 {
                if len(v4) > 0 {
                    v4[0] = (v2 % ((32767) | 1))
                }
                v10 = (-(v10))
            } else {
                v3 = (v3 * v3)
                if len(v4) > 0 {
                    v4[0] = v2
                }
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                v9 = 0
                _ = fi13
            }
        }
        var v14 int32 = v10
        if (!v11) {
            var v15 int32 = (v10 - 249)
            var v16 int8 = 127
            fmt.Println(v16, v3)
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 string = "foo bar"
                v14 += v14
                _ = fi17
                _ = v18
            }
        }
        _ = v14
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch (v2 + 128) {
        case 916:
            v3 += (^(v3))
        case -823:
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 bool = v11
                var v21 int = (fi19 & int(v1))
                v2 = (v2 * (fi19 / ((int(v9)) | 1)))
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        fmt.Println(v10, v2, v11)
    }
    switch v2 {
    case 514:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -1:
        var v22 string = ("" + "a")
        _ = v22
    case 127:
        switch int(v9) {
        case 42:
            var v23 uint = (v3 * 4294967295)
            _ = v23
        case -9223372036854775808:
            switch 383 {
            case -232:
                v11 = (-2.25 == ((3.14159 / (1.5 / ((0.0 / 1e10) + (1e10 + 0.0)))) * 3.14159))
            case -984:
                v9 = (v9 ^ v9)
            default:
                var v24 uint8 = 128
                _ = v24
            }
        case -1:
            v25 := GenS2{F0: (v9 <= (-(v9))), F1: (v9 / (((v9 / (((-(v9))) | 1))) | 1)), F2: 224}
            _ = v25
        default:
            var v26 uint = (^(v3))
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
