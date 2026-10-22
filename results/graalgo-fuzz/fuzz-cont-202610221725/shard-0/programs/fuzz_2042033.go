package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 string; F2 uint }
func genFn0(p0 string, p1 int) uint8 {
    p1 -= (^(p1))
    var v4 int = (p1 / (((p1 | p1)) | 1))
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 uint8 = 0
    var v2 string = "a"
    var v3 uint32 = 783
    v5 := GenS1{F0: 362, F1: -1, F2: v1}
    v1 += 0
    v6 := []int{815, -9223372036854775808}
    _ = len(v6)
    fmt.Println(v1)
    var v7 int = -830
    v3 = 2
    switch v7 {
    case 0:
        switch ("" + v2) {
        case "":
            fmt.Println(v3, v3)
        case "foo bar":
            v2 = (v2 + "\n\t")
        case "a":
            v8 := GenS1{F0: -876, F1: -689, F2: (v1 ^ v1)}
            _ = v8
        default:
            v9 := GenS1{F0: (v7 & (v7 / ((v7) | 1))), F1: 194, F2: (v1 + (v1 - v1))}
            _ = v9
        }
    case 4294967295:
        v1 += genFn0("foo bar", v7)
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if (!false) {
            fmt.Println(v7, fi10)
            v11 := GenS1{F0: (fi10 >> ((0) & 63)), F1: 42, F2: (^(v1))}
            _ = v11
        } else {
            v3 = (^(v3))
            if (true || (0 < (v3 + v3))) {
                fmt.Println(fi10)
                if false {
                    v1 -= 255
                    for fi12 := 0; fi12 < 6; fi12++ {
                        v13 := GenS2{F0: (^(v1))}
                        v7 = (v7 * (fi12 & fi10))
                        v7 += (v7 * v7)
                        _ = fi12
                        _ = v13
                    }
                    if (894 == 0) {
                        v1 = (v1 * v1)
                        v2 = ((("\n\t" + ((v2 + v2) + (v2 + v2))) + v2) + (v2 + v2))
                        _ = len(v6)
                    } else {
                        var v14 int8 = -128
                        if false {
                            var v15 uint = 255
                            v16 := GenS3{F0: "\n\t", F1: v2, F2: v15}
                            if len(v6) > 0 {
                                v6[0] = fi10
                            }
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    }
                }
            } else {
                switch "foo bar" {
                case "hello":
                    for fi17 := 0; fi17 < 10; fi17++ {
                        var v18 int = fi17
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi17
                        _ = v18
                    }
                case "a":
                    fmt.Println(v1, v3, v3)
                case "":
                    for fi19 := 0; fi19 < 11; fi19++ {
                        _ = len(v6)
                        var v20 int64 = -9223372036854775808
                        v2 = (v2 + v2)
                        _ = fi19
                        _ = v20
                    }
                default:
                    v1 = v1
                }
            }
        }
        _ = fi10
    }
    v7 -= v7
    if true {
        var v21 string = "a"
        switch "" {
        case "\n\t":
            var v22 uint8 = (v1 * (v1 + v1))
            _ = v22
        case "hello":
            if len(v6) > 0 {
                v6[0] = int(v3)
            }
        }
        _ = v21
    } else {
        v23 := GenS3{F0: (v2 + v2), F1: v2, F2: 782}
        if (376 != uint16(v3)) {
            fmt.Println(v1)
            if (0 <= 1) {
                for fi24 := 0; fi24 < 17; fi24++ {
                    _ = len(v6)
                    if (true || ((v3 | (v3 + v3)) < 254)) {
                        v6 = append(v6, (v7 | fi24))
                        var v25 int64 = -9223372036854775808
                        _ = v25
                    }
                    _ = fi24
                }
                v26 := GenS1{F0: (v7 << ((65535) & 63)), F1: 42, F2: v1}
                fmt.Println(v1)
                _ = v26
            } else {
                for fi27 := 0; fi27 < 14; fi27++ {
                    if len(v6) > 0 {
                        v6[0] = (fi27 ^ (^(fi27)))
                    }
                    for fi28 := 0; fi28 < 19; fi28++ {
                        var v29 uint8 = 0
                        var v30 string = v2
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = fi27
                }
            }
            fmt.Println(v7, v2)
        }
        _ = v23
    }
    if (1e10 < 0.0) {
        switch 4 {
        case 128:
            _ = len(v6)
        default:
            v2 = ((v2 + (v2 + "foo bar")) + (v2 + (v2 + v2)))
        }
    }
    fmt.Println(v7)
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
