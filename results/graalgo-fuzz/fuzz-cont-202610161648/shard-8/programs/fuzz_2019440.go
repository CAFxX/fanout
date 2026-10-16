package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 uint64 = 1
    v2 := []string{("\n\t" + "hello"), "foo bar", (("foo bar" + "a") + (("" + "\n\t") + "hello"))}
    var v3 int8 = 42
    v4 := []string{"", "foo bar", "\n\t"}
    var v5 uint = 255
    var v6 int32 = 1
    switch 303 {
    case 2147483647:
        for fi7 := 0; fi7 < 11; fi7++ {
            v1 = v1
            v3 -= 100
            _ = fi7
        }
    default:
        if (0 < 0) {
            var v8 float32 = 1e10
            _ = v8
        } else {
            fmt.Println(v5, v1)
        }
    }
    v9 := GenS1{F0: (v6 & (v6 + v6)), F1: v3, F2: -9223372036854775808}
    if (!true) {
        var v10 uint32 = 2147483647
        _ = v10
    }
    var v11 uint = uint(v1)
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: 865, F1: 42, F2: (fi12 >> ((v5) & 63))}
        if (42 < 126) {
            _ = len(v4)
            _ = len(v4)
        } else {
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(v6, v11)
                v15 := GenS2{F0: "", F1: -1, F2: 42}
                _ = fi14
                _ = v15
            }
            fmt.Println(v3, v11)
            v11 += v5
        }
        _ = fi12
        _ = v13
    }
    v16 := GenS3{F0: "", F1: -606}
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v1, v11, v5)
        if false {
            v11 += 150
            for fi18 := 0; fi18 < 10; fi18++ {
                switch "" {
                case "a":
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v19 int64 = 815
                        v20 := GenS4{F0: ("hello" + "")}
                        _ = v19
                        _ = v20
                    }
                    if len(v4) > 0 {
                        v4[0] = (("foo bar" + "foo bar") + "\n\t")
                    }
                    var v21 float32 = 0.1
                    _ = v21
                }
                if (0 != (v11 + (v5 >> ((65535) & 63)))) {
                    v22 := GenS4{F0: ("" + "\n\t")}
                    if ((v3 << (((v11 - 19)) & 63)) > 100) {
                        switch fi17 {
                        case -75:
                            v1 = (v1 + v1)
                        case 1:
                            var v23 uint8 = 42
                            _ = v23
                        }
                        var v24 int = fi18
                        v25 := GenS5{F0: 0}
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                }
                v3 -= v3
                _ = fi18
            }
            var v26 bool = (false || false)
            _ = v26
        }
        v6 = v6
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
