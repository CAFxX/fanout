package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 bool; F2 string }
type GenS3 struct { F0 int; F1 uint64; F2 int8 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int32 = 42
    var v2 uint32 = 884
    v3 := make(map[string]string)
    var v4 int8 = -128
    v5 := []string{"\n\t", "\n\t", "a"}
    var v6 uint32 = 65535
    switch int(v4) {
    case 9223372036854775807:
        if (true || false) {
            if (277 == 127) {
                for fi7 := 0; fi7 < 3; fi7++ {
                    switch -170 {
                    case -9223372036854775808:
                        v8 := GenS1{F0: 1}
                        _ = v8
                    case -567:
                        for fi9 := 0; fi9 < 11; fi9++ {
                            var v10 uint = 42
                            v11 := GenS2{F0: v4, F1: ((int8(v1) > v4) && (true && false)), F2: "\n\t"}
                            v5 = append(v5, "\n\t")
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    case -857:
                        var v12 int = 1
                        _ = v12
                    default:
                        fmt.Println(v2)
                    }
                    _ = fi7
                }
                switch 42 {
                case 42:
                    switch "\n\t" {
                    case "":
                        var v13 string = ("\n\t" + ("hello" + ("\n\t" + "a")))
                        _ = v13
                    default:
                        if true {
                            _ = len(v3)
                        }
                    }
                }
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS1{F0: -301}
                v1 += v1
                _ = fi14
                _ = v15
            }
            v16 := GenS3{F0: -215, F1: 42, F2: -20}
            _ = v16
        }
    case 1:
        var v17 uint8 = 127
        _ = v17
    default:
        switch 1 {
        case 1:
            v18 := GenS4{F0: true, F1: 2, F2: 255}
            _ = v18
        case 32767:
            v19 := GenS2{F0: (v4 - (-(v4))), F1: false, F2: "hello"}
            _ = v19
        case 32768:
            if (-2147483648 < 763) {
                var v20 int32 = (v1 & v1)
                _ = v20
            }
        }
    }
    if true {
        var v21 int64 = -9223372036854775808
        _ = v21
    } else {
        if (true && true) {
            v6 = (v6 % (((v2 << ((4294967295) & 63))) | 1))
            v1 += v1
        } else {
            v6 += v6
            v6 = 4294967295
            var v22 int8 = v4
            _ = v22
        }
    }
    var v23 bool = (0.0 != 0.0)
    v24 := GenS3{F0: 128, F1: 18446744073709551615, F2: (v4 / ((v4) | 1))}
    v25 := v3[""]
    var v26 uint64 = 2
    fmt.Println(v1, v1)
    var v27 string = v25
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 uint16 = 65535
        v30 := GenS5{F0: 0}
        if ((-9223372036854775808 == 42) && (!(!v23))) {
            var v31 uint32 = v6
            _ = v31
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        for fi33 := 0; fi33 < 14; fi33++ {
            fmt.Println(v26)
            _ = fi33
        }
        _ = fi32
    }
    v2 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
