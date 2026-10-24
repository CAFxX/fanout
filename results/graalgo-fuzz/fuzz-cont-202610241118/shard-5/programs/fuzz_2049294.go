package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int16; F1 int8; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 string; F2 bool }
type GenS5 struct { F0 int64; F1 uint; F2 uint32 }
func genFn0(p0 uint64, p1 int8, p2 int) uint32 {
    v7 := GenS1{F0: 42, F1: 32767}
    var v8 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 128
    var v3 uint8 = (v2 + 128)
    v4 := []string{("" + ("\n\t" + ("hello" + ""))), "\n\t"}
    var v5 uint32 = 2147483647
    var v6 int16 = -32768
    v5 = (v5 ^ uint32(v3))
    for fi9 := 0; fi9 < 1; fi9++ {
        if false {
            v10 := GenS1{F0: 0, F1: 65535}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 19; fi11++ {
                switch ("foo bar" + ("a" + "foo bar")) {
                case "hello":
                    v12 := GenS2{F0: v6, F1: -36, F2: 2147483647}
                    switch "" {
                    case "a":
                        v6 = (v6 | (v6 % (((v6 + v6)) | 1)))
                    case "foo bar":
                        if ((fi9 * (fi9 + fi11)) > 127) {
                            fmt.Println(v6)
                            v6 = (v6 - (-(v6)))
                            delete(v1, 2147483647)
                        } else {
                            var v13 string = (("" + (("hello" + "a") + "foo bar")) + "a")
                            _ = v13
                        }
                    }
                    _ = v12
                case "\n\t":
                    _ = len(v1)
                default:
                    var v14 uint16 = 593
                    _ = v14
                }
                v3 = v2
                _ = fi11
            }
            fmt.Println(v2, v5, fi9)
            v15 := GenS1{F0: 9223372036854775807, F1: 65535}
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            v3 = 2
            _ = fi16
        }
        _ = fi9
    }
    var v17 uint8 = v2
    if true {
        v18 := GenS2{F0: -47, F1: 0, F2: 1}
        v19 := GenS3{F0: 9223372036854775807}
        var v20 uint16 = uint16(v2)
        _ = v18
        _ = v19
        _ = v20
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v3, v2)
        switch 83 {
        case 885:
            if false {
                switch "foo bar" {
                case "a":
                    v5 = v5
                case "":
                    fmt.Println(v3)
                default:
                    fmt.Println(v3, v6, v6)
                }
            } else {
                var v21 bool = (!true)
                switch -9223372036854775808 {
                case 32768:
                    if false {
                        fmt.Println(v21, v5, v6)
                    } else {
                        if v21 {
                            v22 := GenS2{F0: (v6 + 32767), F1: 1, F2: -2147483648}
                            var v23 string = "hello"
                            _ = v22
                            _ = v23
                        } else {
                            v1[(v5 | 65535)] = (-128 > -90)
                        }
                        for fi24 := 0; fi24 < 1; fi24++ {
                            v25 := GenS4{F0: 4294967295, F1: ("" + "hello"), F2: true}
                            _ = fi24
                            _ = v25
                        }
                        var v26 uint32 = v5
                        _ = v26
                    }
                case -1:
                    v27 := v1[v5]
                    _ = v27
                default:
                    var v28 int8 = 1
                    _ = v28
                }
                _ = v21
            }
        }
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        v17 -= 202
        _ = fi29
    }
    v30 := GenS5{F0: 570, F1: 65535, F2: 1}
    if (1 <= 255) {
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 int16 = (v6 - v6)
            if false {
                if true {
                    fmt.Println(v6, v2)
                } else {
                    _ = len(v4)
                    v33 := GenS4{F0: 1, F1: "\n\t", F2: (!(!true))}
                    _ = v33
                }
            } else {
                var v34 int8 = 92
                fmt.Println(v17, v2)
                if true {
                    var v35 string = ((("hello" + "\n\t") + ("a" + "")) + ("a" + "foo bar"))
                    _ = v35
                }
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        var v36 float32 = float32(1e10)
        _ = v36
    } else {
        v4 = append(v4, "")
        var v37 uint = 730
        _ = v37
    }
    switch 24 {
    case 32768:
        if false {
            _ = len(v4)
        }
    }
    v2 -= 0
    fmt.Println(v2, v17)
    v38 := GenS2{F0: 191, F1: -67, F2: 612}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        var v39 float32 = -2.25
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
