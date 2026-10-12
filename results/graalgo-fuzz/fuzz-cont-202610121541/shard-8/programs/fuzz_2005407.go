package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16; F1 int32 }
func genFn0(p0 int32, p1 uint16) string {
    v7 := GenS1{F0: -557, F1: "a", F2: 2}
    p1 = p1
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v7
    return ""
}
func genFn1(p0 uint8) uint32 {
    var v8 string = "hello"
    var v9 uint8 = p0
    v10 := GenS2{F0: 577, F1: 100, F2: 1}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 157
}

func main() {
    var v1 bool = (2147483647 >= 42)
    var v2 bool = false
    var v3 uint16 = 798
    var v4 uint8 = 40
    v5 := []int{-922, 181}
    var v6 uint64 = 18446744073709551615
    if (!(255 < 747)) {
        var v11 uint8 = (v4 | v4)
        v12 := GenS1{F0: 4294967295, F1: ((("\n\t" + ("" + "\n\t")) + "foo bar") + ((("" + "\n\t") + (("" + ("hello" + ("\n\t" + "foo bar"))) + "foo bar")) + (("\n\t" + "foo bar") + ""))), F2: 128}
        v1 = true
        _ = v11
        _ = v12
    }
    fmt.Println(v1)
    v4 = (v4 + 128)
    if v1 {
        if ("a" > ("foo bar" + "\n\t")) {
            for fi13 := 0; fi13 < 5; fi13++ {
                v2 = v2
                _ = fi13
            }
            v14 := GenS1{F0: 32767, F1: "hello", F2: 2}
            if true {
                v15 := GenS1{F0: -1, F1: "hello", F2: 687}
                _ = v15
            }
            _ = v14
        } else {
            v6 -= 1
        }
    }
    var v16 uint = uint(v6)
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint16 = 1
        _ = fi17
        _ = v18
    }
    v16 = uint(v3)
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    v3 -= v3
    switch -1 {
    case 128:
        var v19 uint16 = (v3 * (v3 >> (((v16 ^ v16)) & 63)))
        _ = v19
    case 0:
        v3 -= (v3 % (((^(v3))) | 1))
    default:
        for fi20 := 0; fi20 < 17; fi20++ {
            if (127 == 88) {
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 string = "foo bar"
                    v16 = v16
                    var v23 int32 = 0
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                switch "" {
                case "":
                    v3 += (v3 & (v3 & (^(v3))))
                default:
                    if v1 {
                        var v24 uint8 = uint8(v3)
                        for fi25 := 0; fi25 < 6; fi25++ {
                            var v26 uint16 = v3
                            if v1 {
                                v27 := GenS3{F0: 65535, F1: -2147483648}
                                if v2 {
                                    if ((!(2147483647 == 1)) || (!(v6 <= v6))) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        fmt.Println(v16, v2, v6)
                                    }
                                    v24 -= (v4 * (v4 & v4))
                                }
                                _ = v27
                            } else {
                                var v28 bool = v1
                                _ = v28
                            }
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    }
                }
            } else {
                var v29 float32 = ((((1e10 - 1.5) - -2.25) + -2.25) / 1e10)
                fmt.Println(v29)
                v4 = (v4 | uint8(v16))
                _ = v29
            }
            switch ("a" + ((("\n\t" + "a") + ("a" + "hello")) + "a")) {
            case "":
                var v30 float64 = (0.1 / -2.25)
                _ = v30
            }
            _ = fi20
        }
    }
    if len(v5) > 0 {
        v5[0] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
