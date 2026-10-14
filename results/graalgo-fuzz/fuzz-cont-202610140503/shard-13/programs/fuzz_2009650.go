package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 string; F1 bool; F2 int16 }

func main() {
    v1 := []string{("" + (("foo bar" + ("foo bar" + "a")) + "")), ((((("" + "") + "foo bar") + "a") + "a") + ""), ("hello" + ("" + "a"))}
    var v2 int16 = 42
    var v3 uint32 = 2147483647
    var v4 uint8 = 56
    v5 := make(map[uint32]int)
    v1 = append(v1, "foo bar")
    v3 -= (v3 - (v3 & 272))
    v6 := GenS1{F0: (false || false), F1: ((0 <= 0) || (!(0 > 127))), F2: 18446744073709551615}
    var v7 float64 = 3.14159
    fmt.Println(v2, v3)
    fmt.Println(v7, v7)
    delete(v5, 2147483648)
    var v8 int = 85
    v9 := GenS1{F0: (int32(v4) < -2147483648), F1: true, F2: 858}
    switch (v8 >> ((65535) & 63)) {
    case 209:
        _ = len(v1)
    case 128:
        switch v8 {
        case 141:
            var v10 float64 = float64(3.14159)
            _ = v10
        default:
            if (!(v8 >= int(v4))) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch int(v2) {
                    case 32768:
                        switch v8 {
                        case -2147483648:
                            _ = len(v5)
                        }
                    case -1:
                        v12 := GenS1{F0: (v4 >= v4), F1: (v7 <= v7), F2: 65535}
                        _ = v12
                    case -2147483648:
                        for fi13 := 0; fi13 < 1; fi13++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi13
                        }
                    default:
                        if false {
                            var v14 bool = (0 < 0)
                            for fi15 := 0; fi15 < 19; fi15++ {
                                v8 = (fi15 << ((4294967295) & 63))
                                _ = fi15
                            }
                            _ = v14
                        } else {
                            v16 := GenS1{F0: (v2 > v2), F1: true, F2: 65535}
                            v8 += 4294967295
                            _ = v16
                        }
                    }
                    switch (("foo bar" + "a") + "a") {
                    case "a":
                        if true {
                            v7 = (v7 + v7)
                            var v17 uint16 = 32767
                            _ = v17
                        } else {
                            v8 -= (v8 << ((65535) & 63))
                            v3 = v3
                        }
                    }
                    v18 := GenS2{F0: "hello", F1: (!((float64(3.14159) < (v7 / float64(1e10))) || ((2 < 933) || (!(!true))))), F2: v2}
                    _ = fi11
                    _ = v18
                }
                fmt.Println(v3, v8, v7)
            }
        }
    }
    switch "" {
    case "hello":
        if ((!(!(false && false))) && (65535 == v3)) {
            var v19 uint16 = 2
            fmt.Println(v7)
            switch (v8 & v8) {
            case 32767:
                v20 := GenS2{F0: "\n\t", F1: true, F2: (v2 | (v2 / ((-1) | 1)))}
                _ = v20
            case 9223372036854775807:
                v21 := v5[2147483648]
                _ = v21
            default:
                v3 = 0
            }
            _ = v19
        } else {
            v4 = 128
            v7 += v7
        }
    case "":
        if (v8 <= 9223372036854775807) {
            var v22 bool = (0 < (v8 * (-(v8))))
            _ = v22
        } else {
            var v23 int = 0
            v3 = (v3 & 74)
            _ = v23
        }
    case "\n\t":
        var v24 int8 = 26
        _ = v24
    default:
        var v25 uint64 = 42
        _ = v25
    }
    var v26 string = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
