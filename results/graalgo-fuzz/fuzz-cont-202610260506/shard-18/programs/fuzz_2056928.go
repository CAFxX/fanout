package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint32; F2 int16 }
func genFn0(p0 int8, p1 uint8) string {
    var v5 int8 = (^(p0))
    p1 = (p1 << ((18446744073709551615) & 63))
    v6 := []int{32767, 361, 9223372036854775807}
    var v7 uint8 = (^(p1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return ("" + "foo bar")
}

func main() {
    v1 := []string{"a", ("foo bar" + "foo bar")}
    var v2 int8 = 127
    v3 := []string{"foo bar", "a", "foo bar", "hello"}
    var v4 float32 = 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 uint = 2
    v9 := GenS1{F0: (!("foo bar" != ("" + "a")))}
    switch 470 {
    case 2147483647:
        var v10 uint64 = 42
        _ = v10
    case 9223372036854775807:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v8 += v8
    v4 -= 0.1
    var v11 uint = 434
    v8 += (v8 << ((v8) & 63))
    v11 += (v8 * v8)
    v4 += 1.5
    if false {
        v4 -= ((v4 + v4) / (v4 - -2.25))
    }
    if (!false) {
        if len(v3) > 0 {
            v3[0] = ""
        }
        if (("a" + ((("" + "\n\t") + ("" + "hello")) + "foo bar")) > "a") {
            if len(v3) > 0 {
                v3[0] = ("hello" + "\n\t")
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi12 := 0; fi12 < 7; fi12++ {
                for fi13 := 0; fi13 < 5; fi13++ {
                    fmt.Println(v8, v2)
                    v2 = 100
                    _ = fi13
                }
                switch ("" + ("" + "foo bar")) {
                case "":
                    v14 := GenS1{F0: false}
                    _ = v14
                case "a":
                    if false {
                        var v15 uint8 = 255
                        switch ("\n\t" + "") {
                        case "a":
                            v15 = (v15 % (((^(v15))) | 1))
                        case "":
                            v8 += v11
                        case "hello":
                            v11 = 42
                        default:
                            v15 += 134
                        }
                        for fi16 := 0; fi16 < 18; fi16++ {
                            v11 -= (v8 * (v11 / (((v8 & (v8 * (v11 + v8)))) | 1)))
                            if (false && false) {
                                v2 += (v2 / ((v2) | 1))
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                var v17 int = 2147483647
                                _ = v17
                            }
                            if true {
                                var v18 uint16 = 291
                                v19 := GenS1{F0: ((v18 - (v18 & v18)) > (v18 * v18))}
                                v20 := GenS1{F0: true}
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                if ((0.0 >= 0.0) && false) {
                                    v2 += v2
                                    v21 := GenS1{F0: false}
                                    v2 -= (^(v2))
                                    fmt.Println(v4, fi12, fi16)
                                    _ = v21
                                }
                            }
                            _ = fi16
                        }
                        _ = v15
                    }
                case "hello":
                    switch -870 {
                    case 927:
                        if false {
                            for fi22 := 0; fi22 < 1; fi22++ {
                                var v23 string = "hello"
                                var v24 uint16 = 255
                                fmt.Println(v8)
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            v25 := GenS1{F0: (127 <= 127)}
                            _ = v25
                        } else {
                            v3 = append(v3, "foo bar")
                            var v26 int64 = 9223372036854775807
                            _ = v26
                        }
                    }
                default:
                    v11 += (v8 % ((v8) | 1))
                }
                if false {
                    v27 := GenS2{F0: (v2 - v2), F1: 2147483648, F2: -1}
                    _ = v27
                }
                _ = fi12
            }
        }
    }
    v11 -= (v8 | 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
