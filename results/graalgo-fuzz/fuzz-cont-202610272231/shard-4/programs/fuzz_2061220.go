package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int8 }
type GenS3 struct { F0 string; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 string; F1 int64; F2 uint16 }

func main() {
    var v1 uint = 993
    var v2 uint8 = 1
    var v3 bool = false
    v4 := []string{"a", (("" + ((("hello" + "") + "foo bar") + ((("\n\t" + "hello") + "foo bar") + "a"))) + ("foo bar" + "foo bar")), ("a" + ("a" + "")), ("hello" + "foo bar")}
    v3 = (((!v3) || false) && v3)
    fmt.Println(v3, v3)
    _ = len(v4)
    v2 -= 2
    v3 = v3
    v5 := GenS1{F0: 379, F1: 32767}
    fmt.Println(v3, v1)
    var v6 uint32 = 0
    var v7 uint16 = 255
    var v8 string = "hello"
    var v9 int64 = -9223372036854775808
    if true {
        switch 32767 {
        case -637:
            if v3 {
                v10 := GenS2{F0: v7, F1: 100, F2: 1}
                var v11 int64 = -9223372036854775808
                _ = v10
                _ = v11
            } else {
                fmt.Println(v7)
                if v3 {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        var v13 int16 = 32767
                        switch 0 {
                        case 928:
                            if v3 {
                                if ((!(!((312 == v7) && (0.1 <= (1.5 - 0.1))))) && true) {
                                    v6 -= v6
                                    for fi14 := 0; fi14 < 20; fi14++ {
                                        var v15 uint16 = v7
                                        _ = fi14
                                        _ = v15
                                    }
                                    _ = len(v4)
                                } else {
                                    var v16 int32 = 1
                                    v17 := GenS3{F0: v8, F1: (v7 / (((v7 - v7)) | 1)), F2: 1}
                                    _ = v16
                                    _ = v17
                                }
                                v3 = v3
                                var v18 int = 32767
                                _ = v18
                            } else {
                                var v19 int64 = v9
                                v20 := GenS4{F0: (^(v2)), F1: -1}
                                v4 = append(v4, v8)
                                _ = v19
                                _ = v20
                            }
                        case 1:
                            switch (fi12 / (((^(fi12))) | 1)) {
                            case 128:
                                var v21 float64 = (1.5 - -2.25)
                                _ = v21
                            case 127:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        case -773:
                            var v22 int8 = -128
                            _ = v22
                        default:
                            v8 = "\n\t"
                        }
                        _ = fi12
                        _ = v13
                    }
                } else {
                    var v23 string = v8
                    _ = v23
                }
                v3 = v3
            }
        }
        v7 = (v7 / (((v7 & (v7 + v7))) | 1))
    }
    v2 -= v2
    _ = len(v4)
    if (!v3) {
        if (v3 || (false && ((3.14159 < 3.14159) && v3))) {
            var v24 int16 = -83
            _ = v24
        } else {
            var v25 int = -9223372036854775808
            _ = v25
        }
        if (42 != -1) {
            v26 := GenS4{F0: 42, F1: 42}
            if true {
                v27 := GenS4{F0: 0, F1: 0}
                fmt.Println(v6, v8)
                _ = v27
            }
            _ = v26
        } else {
            v28 := GenS5{F0: v8, F1: (v9 >> ((255) & 63)), F2: v7}
            _ = v28
        }
        var v29 uint64 = 4294967295
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
