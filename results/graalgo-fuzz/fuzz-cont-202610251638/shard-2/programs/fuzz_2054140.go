package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 bool }

func main() {
    var v1 uint64 = 110
    var v2 int = 745
    v3 := []uint32{65535, 42, 2147483647}
    v2 += 458
    v2 = 302
    if len(v3) > 0 {
        v3[0] = 65535
    }
    if true {
        v1 = (v1 + v1)
    }
    var v4 int32 = -1
    for fi5 := 0; fi5 < 3; fi5++ {
        if (42 >= 255) {
            v4 = -2147483648
        } else {
            if (!(false || true)) {
                _ = len(v3)
                fmt.Println(v1)
                _ = len(v3)
            } else {
                v2 -= (v2 + fi5)
            }
        }
        v3 = append(v3, 2147483647)
        v2 -= (fi5 & (-(fi5)))
        _ = fi5
    }
    if ((("\n\t" > ("" + "foo bar")) || (205 == -620)) || (("" + ("hello" + ("foo bar" + "hello"))) < (("foo bar" + "hello") + (("a" + "") + "foo bar")))) {
        var v6 bool = (true || false)
        _ = v6
    }
    if false {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        switch (v2 >> ((102) & 63)) {
        case -808:
            for fi7 := 0; fi7 < 2; fi7++ {
                v8 := GenS1{F0: (v1 * (v1 >> ((406) & 63)))}
                switch (fi7 | (fi7 | 9223372036854775807)) {
                case -812:
                    v9 := GenS1{F0: 18446744073709551615}
                    _ = v9
                case -544:
                    v1 = (v1 + 0)
                default:
                    fmt.Println(v1)
                }
                fmt.Println(fi7)
                _ = fi7
                _ = v8
            }
        case 127:
            v2 = 32767
        default:
            if true {
                v10 := GenS1{F0: (^(v1))}
                if (float64(((0.0 / ((1.5 + 0.1) + ((3.14159 * 1.5) / ((1e10 - 0.0) / 3.14159)))) - 0.1)) == (0.1 - 3.14159)) {
                    var v11 float32 = -2.25
                    _ = v11
                } else {
                    v12 := GenS1{F0: v1}
                    fmt.Println(v1, v4, v4)
                    _ = v12
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v10
            } else {
                for fi13 := 0; fi13 < 1; fi13++ {
                    _ = len(v3)
                    v1 = (v1 + (v1 ^ 370))
                    if ((!false) && true) {
                        var v14 uint32 = 4294967295
                        for fi15 := 0; fi15 < 20; fi15++ {
                            if len(v3) > 0 {
                                v3[0] = uint32(v1)
                            }
                            v14 = v14
                            _ = fi15
                        }
                        switch ("" + "foo bar") {
                        case "hello":
                            switch "" {
                            case "a":
                                v16 := GenS2{F0: v1, F1: false}
                                _ = v16
                            }
                        }
                        _ = v14
                    }
                    _ = fi13
                }
                var v17 uint64 = 279
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            if false {
                var v19 int = (v2 + fi18)
                var v20 int = 314
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v19
                _ = v20
            }
            _ = fi18
        }
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
