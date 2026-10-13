package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 int16, p1 int8) int32 {
    p1 -= -54
    v5 := GenS1{F0: (("foo bar" + "\n\t") + "hello"), F1: p1}
    var v6 int16 = (p0 ^ 1)
    var v7 uint = 18446744073709551615
    p1 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint = 329
    v2 := []string{"a"}
    var v3 uint64 = 49
    var v4 uint8 = 0
    v3 = (v3 | (v3 % ((18446744073709551615) | 1)))
    var v8 int16 = 0
    fmt.Println(v8, v1, v8)
    var v9 uint16 = 32767
    switch (("foo bar" + "a") + ("foo bar" + "foo bar")) {
    case "hello":
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int8 = 0
            _ = fi10
            _ = v11
        }
    case "":
        var v12 int8 = 127
        _ = v12
    case "a":
        var v13 int = 127
        _ = v13
    default:
        if (!true) {
            if true {
                if false {
                    v14 := GenS1{F0: "foo bar", F1: 85}
                    _ = v14
                }
            }
            if (!((v9 * 1) > 772)) {
                for fi15 := 0; fi15 < 2; fi15++ {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        switch -9223372036854775808 {
                        case -2147483648:
                            var v17 int64 = -889
                            _ = v17
                        case 32767:
                            if ((-637 <= 152) || (!true)) {
                                switch fi15 {
                                case 436:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    if (((!true) || true) || (42 > 4294967295)) {
                                        var v18 float64 = ((((0.0 + 1.5) / -2.25) / (float64(-2.25) - (0.1 - 1.5))) + 0.0)
                                        v1 = (v1 >> ((727) & 63))
                                        if false {
                                            v9 += (v9 ^ 65535)
                                        }
                                        _ = v18
                                    } else {
                                        v19 := GenS1{F0: "\n\t", F1: 127}
                                        v20 := GenS1{F0: ("a" + ""), F1: 100}
                                        var v21 int16 = int16(fi16)
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                }
                                v1 += 4294967295
                                switch (^(fi16)) {
                                case -140:
                                    var v22 bool = (-868 >= -9223372036854775808)
                                    _ = v22
                                case 2147483647:
                                    v8 -= v8
                                default:
                                    v9 = v9
                                }
                            }
                        case 774:
                            fmt.Println(v9, v3)
                        }
                        var v23 int = (fi16 | fi15)
                        var v24 uint = v1
                        _ = fi16
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v8)
                    for fi25 := 0; fi25 < 9; fi25++ {
                        var v26 int16 = v8
                        var v27 int16 = v26
                        fmt.Println(v26)
                        v26 = (v27 | (v26 / ((0) | 1)))
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    _ = fi15
                }
                v3 -= 4294967295
            } else {
                v9 = v9
                var v28 uint64 = (v3 ^ (v3 >> (((^(v1))) & 63)))
                _ = v28
            }
        }
    }
    if true {
        var v29 uint32 = 2147483647
        v2 = append(v2, (("" + "\n\t") + "hello"))
        if ((false && (0 > (v4 >> ((v1) & 63)))) && false) {
            var v30 uint8 = (v4 % ((0) | 1))
            var v31 int64 = -1
            var v32 float64 = float64(1e10)
            _ = v30
            _ = v31
            _ = v32
        }
        _ = v29
    }
    v8 += v8
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 int16 = v8
        _ = fi33
        _ = v34
    }
    v8 = v8
    v3 -= 2
    var v35 int32 = -1
    v8 = v8
    switch ("\n\t" + ("\n\t" + "")) {
    case "a":
        var v36 int = 4294967295
        _ = v36
    case "hello":
        v9 = (v9 * (v9 * 0))
    case "\n\t":
        v37 := GenS1{F0: ("foo bar" + "a"), F1: 127}
        _ = v37
    }
    v8 += (v8 << ((uint(v4)) & 63))
    var v38 uint32 = 42
    v8 += (v8 * -32768)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v8
    _ = v9
}
