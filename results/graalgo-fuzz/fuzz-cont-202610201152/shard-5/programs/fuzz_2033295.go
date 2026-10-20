package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int16 }

func main() {
    var v1 int8 = -1
    v2 := make(map[int]string)
    var v3 uint8 = 42
    fmt.Println(v1, v3)
    v3 += 128
    if (true && false) {
        if true {
            v4 := v2[0]
            _ = v4
        } else {
            for fi5 := 0; fi5 < 7; fi5++ {
                switch (fi5 + fi5) {
                case -9223372036854775808:
                    v3 = (^(v3))
                case -284:
                    switch ("hello" + "foo bar") {
                    case "a":
                        switch "" {
                        case "":
                            switch ("\n\t" + "hello") {
                            case "a":
                                for fi6 := 0; fi6 < 19; fi6++ {
                                    v3 -= (v3 & v3)
                                    var v7 int8 = v1
                                    if (0 != 18446744073709551615) {
                                        _ = len(v2)
                                    }
                                    var v8 uint64 = 2
                                    _ = fi6
                                    _ = v7
                                    _ = v8
                                }
                            }
                        case "a":
                            v9 := []int{410, (fi5 | (fi5 * fi5)), fi5, (fi5 * fi5)}
                            _ = v9
                        case "\n\t":
                            v10 := v2[fi5]
                            _ = v10
                        }
                    case "\n\t":
                        if ((41 != 195) || false) {
                            switch "\n\t" {
                            case "foo bar":
                                var v11 string = "\n\t"
                                _ = v11
                            default:
                                var v12 float64 = 1.5
                                _ = v12
                            }
                        } else {
                            switch 2147483647 {
                            case 2147483647:
                                fmt.Println(v1, v3, v1)
                            case 32767:
                                v13 := v2[(fi5 >> ((255) & 63))]
                                _ = v13
                            }
                            v14 := GenS1{F0: (fi5 + fi5)}
                            var v15 bool = (!(!((0 != 0) || false)))
                            _ = v14
                            _ = v15
                        }
                    }
                case 188:
                    v16 := []int{(fi5 % (((-(fi5))) | 1)), (fi5 ^ fi5), (fi5 << ((255) & 63)), -1}
                    _ = v16
                default:
                    v17 := GenS1{F0: -528}
                    _ = v17
                }
                v1 = v1
                var v18 int8 = v1
                _ = fi5
                _ = v18
            }
            v19 := GenS1{F0: 128}
            _ = v19
        }
    }
    v20 := v2[4294967295]
    if (3.14159 < -2.25) {
        var v21 uint32 = 65535
        _ = v21
    }
    v20 = v20
    v22 := GenS1{F0: -471}
    v23 := GenS2{F0: 902, F1: -538}
    if (0 >= -1) {
        for fi24 := 0; fi24 < 16; fi24++ {
            fmt.Println(v20)
            if ((0 <= -897) || true) {
                v25 := GenS1{F0: (fi24 << ((65535) & 63))}
                var v26 float64 = 0.0
                fmt.Println(v3, fi24, v20)
                _ = v25
                _ = v26
            }
            v27 := GenS2{F0: 18446744073709551615, F1: 42}
            _ = fi24
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
}
