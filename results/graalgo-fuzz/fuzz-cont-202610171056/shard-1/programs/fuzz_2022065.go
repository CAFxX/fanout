package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int; F1 bool }

func main() {
    v1 := []string{("hello" + "hello"), "\n\t", ("a" + ("" + ("foo bar" + "a"))), "\n\t"}
    var v2 int32 = 42
    var v3 int = 32767
    var v4 uint8 = 127
    var v5 uint32 = 799
    var v6 uint16 = 46
    if false {
        v7 := GenS1{F0: "hello"}
        if (false && true) {
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := GenS1{F0: "foo bar"}
                _ = fi8
                _ = v9
            }
            v10 := GenS2{F0: v2}
            _ = v10
        }
        _ = v7
    } else {
        _ = len(v1)
        v5 -= v5
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS3{F0: 42, F1: 1}
            v3 -= (fi11 + fi11)
            for fi13 := 0; fi13 < 4; fi13++ {
                switch (fi13 ^ v3) {
                case -720:
                    v14 := GenS1{F0: "foo bar"}
                    _ = v14
                case 7:
                    switch "\n\t" {
                    case "foo bar":
                        if (4294967295 >= (fi11 * v3)) {
                            var v15 float64 = 3.14159
                            if (30 >= 12) {
                                v3 += fi13
                                v5 = uint32(v6)
                            } else {
                                if true {
                                    v15 = (v15 * 0.1)
                                    fmt.Println(v5, v4)
                                    v15 = (v15 * v15)
                                }
                                fmt.Println(v15)
                            }
                            _ = v15
                        } else {
                            v6 += (v6 & v6)
                        }
                    }
                    var v16 int = fi13
                    _ = v16
                case -2147483648:
                    var v17 float64 = 1e10
                    var v18 bool = (false && true)
                    _ = v17
                    _ = v18
                default:
                    switch 4294967295 {
                    case 32768:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case -253:
                        v3 -= 42
                    default:
                        var v19 int16 = -73
                        _ = v19
                    }
                }
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    }
    if ((((v2 * (v2 * (v2 * -2147483648))) == v2) && (!((^(v6)) >= (v6 >> ((18446744073709551615) & 63))))) || (!(("a" + "foo bar") == "foo bar"))) {
        if (-1 < -1) {
            fmt.Println(v2, v3)
            var v20 float64 = (-2.25 * 0.0)
            if true {
                v1 = append(v1, "")
                if len(v1) > 0 {
                    v1[0] = ("\n\t" + "foo bar")
                }
            } else {
                fmt.Println(v20, v20, v3)
                v2 += -2147483648
                if false {
                    v21 := GenS2{F0: v2}
                    _ = v21
                } else {
                    v1 = append(v1, ((("a" + "a") + "\n\t") + "hello"))
                    if (v3 != -2147483648) {
                        v22 := GenS4{F0: 100, F1: (v6 & (v6 >> ((4294967295) & 63))), F2: v5}
                        if (v4 >= v4) {
                            v3 = (v3 << ((1) & 63))
                            if (false && true) {
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                            }
                            v4 = (^(v4))
                        } else {
                            v23 := GenS4{F0: 100, F1: 10, F2: v5}
                            v20 = float64(1e10)
                            _ = v23
                        }
                        if len(v1) > 0 {
                            v1[0] = ("hello" + (("\n\t" + ("" + ("a" + ("hello" + "hello")))) + ""))
                        }
                        _ = v22
                    }
                    v20 -= v20
                }
            }
            _ = v20
        }
        v3 -= (v3 / (((v3 % ((v3) | 1))) | 1))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 += (v6 >> ((673) & 63))
    for fi24 := 0; fi24 < 20; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            fmt.Println(fi24, v5, fi25)
            if false {
                fmt.Println(fi24)
                v5 += 42
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v26 int16 = 0
                    _ = v26
                }
            } else {
                fmt.Println(v6, v5, fi25)
            }
            _ = fi25
        }
        _ = fi24
    }
    v5 = 4294967295
    if false {
        fmt.Println(v2, v3)
    }
    if true {
        v27 := GenS5{F0: v3, F1: ((0 >= 584) && false)}
        _ = v27
    }
    var v28 string = ("a" + ("\n\t" + "foo bar"))
    v29 := GenS2{F0: v2}
    if true {
        v6 -= (v6 ^ (v6 ^ (v6 & (^(v6)))))
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 int32 = (v2 / ((-2147483648) | 1))
            var v32 bool = true
            _ = fi30
            _ = v31
            _ = v32
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        if len(v1) > 0 {
            v1[0] = (("foo bar" + v28) + v28)
        }
        var v33 uint8 = 124
        v28 = (v28 + v28)
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
