package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint8 = 0
    var v2 int = -1
    v3 := []string{"foo bar", "\n\t"}
    v4 := []string{"\n\t", ("a" + ("hello" + ("foo bar" + "foo bar"))), ("a" + ("hello" + "foo bar"))}
    v5 := []uint32{4294967295, 2147483647, 42, 980}
    var v6 float32 = (-2.25 - 0.1)
    v7 := GenS1{F0: 65535, F1: 12}
    if true {
        var v8 float64 = (1.5 - 0.1)
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS1{F0: 65535, F1: 1}
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 uint8 = 2
                _ = fi11
                _ = v12
            }
            var v13 uint8 = v1
            _ = fi9
            _ = v10
            _ = v13
        }
        _ = v8
    }
    if true {
        var v14 uint8 = (v1 + v1)
        switch -444 {
        case 4294967295:
            v2 -= 1
        case 127:
            for fi15 := 0; fi15 < 7; fi15++ {
                switch "a" {
                case "\n\t":
                    var v16 int16 = 1
                    _ = v16
                case "foo bar":
                    for fi17 := 0; fi17 < 18; fi17++ {
                        var v18 int8 = -86
                        var v19 float64 = -2.25
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                case "hello":
                    fmt.Println(v6, fi15)
                }
                if (323 == -676) {
                    var v20 uint8 = (v14 << ((18446744073709551615) & 63))
                    v20 = (v20 - v20)
                    _ = v20
                } else {
                    switch 42 {
                    case 692:
                        var v21 float32 = v6
                        _ = v21
                    case 128:
                        fmt.Println(v6, v1, v1)
                    case 1:
                        v14 = 127
                    default:
                        if len(v4) > 0 {
                            v4[0] = "foo bar"
                        }
                    }
                }
                _ = fi15
            }
        }
        _ = v14
    } else {
        fmt.Println(v1, v6)
        var v22 int8 = 42
        if true {
            if ((!(759 >= 42)) && true) {
                var v23 float64 = 0.0
                for fi24 := 0; fi24 < 10; fi24++ {
                    var v25 uint32 = 2
                    var v26 float32 = 1.5
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                var v27 int = (v2 << ((0) & 63))
                _ = v23
                _ = v27
            } else {
                var v28 uint32 = 65535
                if true {
                    v6 -= (1e10 + v6)
                } else {
                    v2 -= v2
                }
                var v29 uint32 = 1
                _ = v28
                _ = v29
            }
            for fi30 := 0; fi30 < 16; fi30++ {
                fmt.Println(v1, v2)
                _ = fi30
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v1 += 0
            if len(v5) > 0 {
                v5[0] = 346
            }
        }
        _ = v22
    }
    var v31 uint32 = 0
    fmt.Println(v1, v1, v31)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v32 int = v2
    v33 := GenS1{F0: 1, F1: -128}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
