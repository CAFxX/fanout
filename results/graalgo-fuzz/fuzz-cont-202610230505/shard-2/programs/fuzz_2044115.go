package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 int16 }

func main() {
    v1 := []string{"foo bar", "hello", "\n\t"}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 966
    var v4 uint8 = 255
    v5 := GenS1{F0: -2147483648}
    v1 = append(v1, "hello")
    v3 = (v3 & 185)
    var v6 int8 = -1
    v4 = 255
    var v7 int64 = 82
    switch ((("hello" + ("foo bar" + ("a" + "a"))) + ("a" + "a")) + "\n\t") {
    case "foo bar":
        if false {
            switch "" {
            case "\n\t":
                for fi8 := 0; fi8 < 15; fi8++ {
                    v9 := GenS1{F0: 2147483647}
                    _ = fi8
                    _ = v9
                }
            default:
                if (127 == 7) {
                    v3 -= 0
                } else {
                    v3 = (v3 & (v3 | 42))
                }
            }
            for fi10 := 0; fi10 < 16; fi10++ {
                v4 += 255
                _ = fi10
            }
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := v2[42]
                var v13 uint16 = 65535
                _ = len(v1)
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    default:
        for fi14 := 0; fi14 < 14; fi14++ {
            delete(v2, 551)
            switch 32768 {
            case 42:
                var v15 uint8 = v4
                _ = v15
            }
            v2[0] = 404
            _ = fi14
        }
    }
    v16 := GenS1{F0: -9223372036854775808}
    if (!false) {
        if false {
            fmt.Println(v7, v3, v7)
        } else {
            var v17 string = "\n\t"
            var v18 float32 = (-2.25 + (3.14159 + 0.0))
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: -2147483648}
        if (18446744073709551615 < v3) {
            v7 = -9223372036854775808
        } else {
            if false {
                switch 9223372036854775807 {
                case 914:
                    if ((!(true || true)) && (!false)) {
                        v6 = v6
                        v6 += -128
                        v4 = (v4 % (((v4 & v4)) | 1))
                    } else {
                        var v20 string = ""
                        if len(v1) > 0 {
                            v1[0] = ("\n\t" + (((("foo bar" + "\n\t") + (("foo bar" + (v20 + "foo bar")) + "foo bar")) + ("a" + "")) + ("foo bar" + "a")))
                        }
                        _ = v20
                    }
                case 2147483647:
                    var v21 int8 = int8(v7)
                    _ = v21
                case -601:
                    var v22 uint = 65535
                    _ = v22
                default:
                    v23 := GenS1{F0: -623}
                    _ = v23
                }
                v4 -= v4
            } else {
                v24 := GenS1{F0: 9223372036854775807}
                _ = v24
            }
            var v25 int64 = (v7 | 538)
            _ = len(v2)
            _ = v25
        }
        _ = v19
    } else {
        switch (((((("\n\t" + "hello") + ((("hello" + "a") + "") + "hello")) + ("hello" + "")) + "foo bar") + "a") + "\n\t") {
        case "foo bar":
            for fi26 := 0; fi26 < 18; fi26++ {
                switch fi26 {
                case 609:
                    if (true && true) {
                        v4 = (v4 % ((57) | 1))
                        v3 = v3
                        v27 := GenS2{F0: 42, F1: uint16(v3), F2: 238}
                        _ = v27
                    }
                case -770:
                    for fi28 := 0; fi28 < 17; fi28++ {
                        if (false && (-9223372036854775808 <= -566)) {
                            var v29 uint8 = (v4 % (((v4 + v4)) | 1))
                            var v30 uint64 = 4294967295
                            v31 := GenS1{F0: 9223372036854775807}
                            _ = v29
                            _ = v30
                            _ = v31
                        } else {
                            _ = len(v1)
                        }
                        _ = fi28
                    }
                case 128:
                    for fi32 := 0; fi32 < 18; fi32++ {
                        v3 -= v3
                        _ = fi32
                    }
                default:
                    var v33 float64 = (1e10 * 3.14159)
                    _ = v33
                }
                v2[685] = 0
                var v34 uint16 = 42
                _ = fi26
                _ = v34
            }
        default:
            var v35 uint8 = (v4 % ((0) | 1))
            _ = v35
        }
        fmt.Println(v6, v4, v6)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
