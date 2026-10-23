package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0(p0 bool, p1 uint64, p2 uint8) int64 {
    v5 := GenS1{F0: 731, F1: 42, F2: -9223372036854775808}
    var v6 uint32 = 889
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 652
}
func genFn1(p0 int32, p1 int64, p2 uint64) int16 {
    fmt.Println(p1, p1)
    p2 = p2
    v7 := GenS2{F0: 885, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}

func main() {
    var v1 uint = 42
    var v2 int = int(v1)
    var v3 uint32 = 42
    v4 := []string{(((("\n\t" + "foo bar") + "\n\t") + "hello") + "a"), ((("a" + "") + "foo bar") + ("a" + ("a" + "\n\t"))), "\n\t", ("a" + "\n\t")}
    v2 -= v2
    fmt.Println(v3, v3, v1)
    if true {
        v8 := GenS1{F0: 32767, F1: v1, F2: 42}
        v9 := GenS1{F0: 955, F1: 18446744073709551615, F2: 199}
        _ = v8
        _ = v9
    }
    switch "hello" {
    case "hello":
        fmt.Println(v2)
    case "\n\t":
        var v10 bool = ((v3 + 954) >= v3)
        _ = v10
    case "foo bar":
        if (true && true) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            fmt.Println(v1, v2, v1)
        }
    default:
        var v11 int16 = -874
        _ = v11
    }
    if true {
        v3 += v3
        var v12 int32 = 1
        _ = v12
    }
    v3 -= (v3 ^ v3)
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint32 = v3
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            v1 = v1
            v4 = append(v4, ("foo bar" + "a"))
            v2 -= int(v3)
            fmt.Println(v1, fi16, v1)
            if true {
                v4 = append(v4, "\n\t")
                if true {
                    var v17 int16 = 1000
                    if false {
                        fmt.Println(v3)
                    } else {
                        switch (("a" + ("hello" + "")) + ((("hello" + "\n\t") + "a") + (("" + (("\n\t" + ("" + "foo bar")) + ("a" + "a"))) + ((("a" + "hello") + (("hello" + "") + "\n\t")) + (("foo bar" + "a") + ("\n\t" + ("\n\t" + "hello"))))))) {
                        case "a":
                            var v18 uint16 = 224
                            var v19 int64 = 1
                            _ = v18
                            _ = v19
                        case "foo bar":
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                        case "":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                    v20 := GenS1{F0: 688, F1: (v1 ^ (v1 / ((v1) | 1))), F2: 315}
                    _ = v17
                    _ = v20
                }
            } else {
                v21 := GenS2{F0: 4294967295, F1: 341}
                v22 := GenS1{F0: 32767, F1: 924, F2: -9223372036854775808}
                _ = v21
                _ = v22
            }
            v3 = (v3 >> ((442) & 63))
            _ = fi16
        }
        if ((v2 != v2) && false) {
            v2 += (v2 % (((-(fi13))) | 1))
            if true {
                var v23 float32 = (3.14159 - (1.5 + 0.0))
                switch (v2 << ((1) & 63)) {
                case -678:
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v4 = append(v4, "a")
                        var v25 int8 = 1
                        var v26 uint32 = (v3 + uint32(v1))
                        v27 := GenS1{F0: 1, F1: (v1 & (v1 >> (((^(v1))) & 63))), F2: 1}
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                case -9223372036854775808:
                    var v28 bool = (-904 < -1)
                    _ = v28
                case 986:
                    var v29 uint = 18446744073709551615
                    _ = v29
                default:
                    v3 += (v3 * 2147483648)
                }
                _ = v23
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    v31 := GenS2{F0: 95, F1: (v1 << ((v1) & 63))}
                    if true {
                        var v32 int64 = -1
                        _ = v32
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = len(v4)
                    _ = fi30
                    _ = v31
                }
                for fi33 := 0; fi33 < 18; fi33++ {
                    var v34 int32 = 619
                    _ = fi33
                    _ = v34
                }
            }
            v2 = int(v1)
        } else {
            fmt.Println(v1, v1)
            v3 = v3
        }
        _ = fi13
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 += (v2 - (v2 << (((v1 << (((^(v1))) & 63))) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
