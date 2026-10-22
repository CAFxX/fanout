package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int64; F2 int64 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 string = ((("hello" + ("foo bar" + ("foo bar" + ""))) + ("hello" + ("a" + "hello"))) + (("\n\t" + "\n\t") + "foo bar"))
    v2 := make(map[int]bool)
    var v3 float64 = float64(-2.25)
    var v4 uint64 = 42
    var v5 uint32 = 4294967295
    for fi6 := 0; fi6 < 2; fi6++ {
        if ((0 >= 32767) && ("\n\t" != (v1 + (v1 + "")))) {
            switch (^(fi6)) {
            case -70:
                v7 := GenS1{F0: v5}
                _ = v7
            case 127:
                switch (fi6 ^ -84) {
                case -665:
                    var v8 uint32 = (v5 * v5)
                    _ = v8
                case 128:
                    fmt.Println(v4, v5, v3)
                case 2147483647:
                    v4 = v4
                }
            default:
                v5 += (v5 | (v5 * 2))
            }
            if (true || false) {
                fmt.Println(v5)
                switch v1 {
                case "":
                    for fi9 := 0; fi9 < 20; fi9++ {
                        if false {
                            v4 -= 765
                            v10 := []int{}
                            _ = v10
                        } else {
                            if (v4 < v4) {
                                var v11 uint8 = 234
                                v5 = v5
                                v12 := v2[fi6]
                                _ = v11
                                _ = v12
                            } else {
                                v5 += (v5 + 856)
                                var v13 string = v1
                                _ = v13
                            }
                            switch (("a" + v1) + ("foo bar" + ("" + (v1 + v1)))) {
                            case "":
                                v14 := GenS2{F0: ("foo bar" + ((v1 + v1) + "a")), F1: 474}
                                _ = v14
                            }
                        }
                        var v15 float64 = 1e10
                        v16 := GenS1{F0: (v5 % ((42) | 1))}
                        _ = fi9
                        _ = v15
                        _ = v16
                    }
                case "hello":
                    v5 = (^(v5))
                case "foo bar":
                    fmt.Println(v1)
                default:
                    var v17 int = (^(fi6))
                    _ = v17
                }
            } else {
                v18 := GenS3{F0: 595, F1: -444, F2: 234}
                v4 += 18446744073709551615
                _ = v18
            }
            if false {
                v19 := GenS1{F0: (v5 ^ v5)}
                _ = v19
            }
        }
        _ = len(v2)
        v5 += 750
        _ = fi6
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v1 = v1
        _ = fi20
    }
    v3 -= (float64(1.5) - v3)
    fmt.Println(v5, v5, v4)
    var v21 uint64 = v4
    var v22 int = 42
    var v23 uint = 1
    v24 := GenS4{F0: v21}
    v25 := GenS3{F0: 65535, F1: -194, F2: -817}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
