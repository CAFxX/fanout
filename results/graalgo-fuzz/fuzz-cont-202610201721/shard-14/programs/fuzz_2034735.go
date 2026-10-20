package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 string) string {
    if true {
        var v6 uint64 = 18446744073709551615
        _ = v6
    }
    fmt.Println(p0)
    var v7 int = 128
    p0 = p0
    _ = p0
    _ = v7
    return ((((p0 + p0) + p0) + (p0 + (p0 + p0))) + "hello")
}
func genFn1(p0 int32) uint32 {
    var v8 int16 = -32768
    var v9 int64 = -1
    fmt.Println(v9, p0, p0)
    var v10 uint16 = 551
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint64 = 949
    var v2 bool = true
    var v3 uint16 = 121
    var v4 int32 = 0
    v5 := []string{("" + "hello")}
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint64 = (v1 % (((v1 >> ((42) & 63))) | 1))
        _ = fi11
        _ = v12
    }
    var v13 uint16 = v3
    var v14 int64 = 1
    switch -1 {
    case 127:
        for fi15 := 0; fi15 < 15; fi15++ {
            v4 -= (v4 << ((2) & 63))
            _ = fi15
        }
    case -740:
        v3 += v13
    case -412:
        for fi16 := 0; fi16 < 3; fi16++ {
            for fi17 := 0; fi17 < 18; fi17++ {
                v13 += (v13 << ((uint(v3)) & 63))
                v14 -= -9223372036854775808
                _ = fi17
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi16
        }
    }
    if true {
        var v18 uint64 = 1
        fmt.Println(v2, v13)
        v18 -= 611
        _ = v18
    } else {
        var v19 int32 = int32(v3)
        _ = v19
    }
    v3 += (v3 * (^(v13)))
    for fi20 := 0; fi20 < 6; fi20++ {
        fmt.Println(fi20, v1)
        var v21 int8 = 100
        v13 = (^(v3))
        _ = fi20
        _ = v21
    }
    var v22 int32 = -2147483648
    var v23 uint32 = 21
    if v2 {
        for fi24 := 0; fi24 < 19; fi24++ {
            v3 += 0
            v25 := GenS1{F0: "foo bar", F1: 4}
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                _ = len(v5)
            default:
                if (4294967295 > 1) {
                    var v26 bool = (!((v13 << ((4294967295) & 63)) != v3))
                    for fi27 := 0; fi27 < 6; fi27++ {
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                        var v28 int32 = v22
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                } else {
                    var v29 int16 = 1
                    _ = v29
                }
            }
            _ = fi24
            _ = v25
        }
        v30 := GenS1{F0: "\n\t", F1: 1}
        if false {
            switch 32768 {
            case 1:
                v31 := GenS1{F0: "foo bar", F1: -89}
                _ = v31
            }
        }
        _ = v30
    } else {
        v32 := GenS2{F0: 255, F1: v4}
        _ = v32
    }
    v13 = 0
    var v33 int32 = v4
    if len(v5) > 0 {
        v5[0] = ("a" + "a")
    }
    var v34 bool = v2
    for fi35 := 0; fi35 < 3; fi35++ {
        switch fi35 {
        case 9223372036854775807:
            v1 += v1
        default:
            v2 = false
        }
        _ = len(v5)
        for fi36 := 0; fi36 < 1; fi36++ {
            if len(v5) > 0 {
                v5[0] = "a"
            }
            _ = fi36
        }
        _ = fi35
    }
    if (("" + (("" + ("a" + "foo bar")) + ((("\n\t" + ("foo bar" + "a")) + "foo bar") + "\n\t"))) == "foo bar") {
        switch 2147483647 {
        case 128:
            fmt.Println(v4)
        default:
            fmt.Println(v23, v3)
        }
        var v37 string = "foo bar"
        var v38 float32 = 3.14159
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        var v40 int = fi39
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
