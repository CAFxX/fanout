package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int32 }
func genFn0(p0 uint, p1 string) uint {
    var v7 string = p1
    if false {
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(v7, p1, v7)
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return p0
}

func main() {
    var v1 uint = 0
    var v2 uint16 = uint16(v1)
    var v3 uint = (v1 % ((359) | 1))
    var v4 uint8 = 128
    v5 := []string{}
    var v6 float64 = ((0.0 * 1e10) / 1e10)
    _ = len(v5)
    var v9 uint8 = 162
    switch int(v2) {
    case -795:
        if (true && true) {
            var v10 uint8 = (v4 % ((1) | 1))
            fmt.Println(v4, v1)
            v9 -= 128
            _ = v10
        }
    case 9223372036854775807:
        v9 += (v4 - (v9 / (((v4 / ((2) | 1))) | 1)))
    }
    v11 := GenS1{F0: "foo bar", F1: 1000, F2: ("foo bar" <= ("a" + ""))}
    v12 := GenS1{F0: "\n\t", F1: 1, F2: true}
    switch 42 {
    case 912:
        v13 := GenS1{F0: "hello", F1: 32767, F2: false}
        _ = v13
    case 128:
        if true {
            var v14 uint64 = 1
            _ = v14
        } else {
            var v15 float64 = (3.14159 + (((float64((1.5 + (float32((3.14159 / v6)) * 0.0))) / v6) - (0.0 + 1e10)) + (v6 - -2.25)))
            _ = v15
        }
    case 2147483647:
        v16 := GenS1{F0: (("foo bar" + (((("" + ("a" + "hello")) + ("hello" + (("a" + ("hello" + "a")) + "\n\t"))) + "") + "foo bar")) + "foo bar"), F1: 32767, F2: ((((("foo bar" + "") + "a") + ("\n\t" + "foo bar")) < "\n\t") || (4294967295 <= 4294967295))}
        _ = v16
    default:
        if (127 <= 56) {
            if false {
                v1 = genFn0(uint(v9), "")
                for fi17 := 0; fi17 < 7; fi17++ {
                    if false {
                        v1 = (v3 << ((v3) & 63))
                        v18 := GenS2{F0: 0}
                        _ = v18
                    } else {
                        if true {
                            v19 := GenS2{F0: 32767}
                            switch (fi17 ^ -2147483648) {
                            case -2147483648:
                                v2 -= v2
                            }
                            _ = v19
                        } else {
                            v3 = uint(v9)
                            var v20 uint8 = (v9 >> ((v1) & 63))
                            _ = v20
                        }
                        if (18446744073709551615 >= 210) {
                            fmt.Println(v6)
                            fmt.Println(fi17, v6, fi17)
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = len(v5)
                            var v21 int64 = 9223372036854775807
                            _ = v21
                        }
                    }
                    _ = fi17
                }
                switch ("a" + "\n\t") {
                case "foo bar":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "\n\t":
                    v4 -= (v4 ^ 191)
                case "hello":
                    var v22 float64 = v6
                    _ = v22
                }
            } else {
                var v23 uint32 = 0
                _ = v23
            }
            if false {
                for fi24 := 0; fi24 < 2; fi24++ {
                    fmt.Println(v1)
                    v4 = 128
                    var v25 uint8 = v9
                    _ = fi24
                    _ = v25
                }
            } else {
                v26 := GenS3{F0: (1 >= -128)}
                _ = len(v5)
                _ = v26
            }
        } else {
            if (true || (v3 < 656)) {
                v2 -= (v2 | v2)
                fmt.Println(v2, v9, v2)
            }
            switch ("hello" + "") {
            case "a":
                if (!true) {
                    var v27 int = 882
                    if false {
                        v4 = v4
                        v27 = v27
                    } else {
                        for fi28 := 0; fi28 < 8; fi28++ {
                            switch 128 {
                            case 32768:
                                v29 := GenS4{F0: ("\n\t" + "foo bar"), F1: 42}
                                _ = v29
                            }
                            fmt.Println(v1)
                            v5 = append(v5, "hello")
                            _ = fi28
                        }
                    }
                    fmt.Println(v4)
                    _ = v27
                }
            default:
                var v30 int32 = -1
                _ = v30
            }
        }
    }
    var v31 uint64 = 891
    var v32 int64 = int64(v9)
    v33 := GenS2{F0: -328}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
