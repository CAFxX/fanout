package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int }
func genFn0(p0 int32, p1 uint16) uint32 {
    v4 := GenS1{F0: 1, F1: 962}
    p0 += (-(p0))
    p0 += -677
    _ = p0
    _ = p1
    _ = v4
    return 19
}

func main() {
    v1 := []string{"\n\t", "a"}
    v2 := []int{128}
    var v3 bool = (42 != 56)
    v5 := GenS1{F0: 2, F1: genFn0(0, 56)}
    v3 = (128 >= 52)
    v3 = ((1e10 / 3.14159) <= ((1e10 * 1.5) + 3.14159))
    fmt.Println(v3)
    v3 = (true && (!v3))
    switch "\n\t" {
    case "foo bar":
        for fi6 := 0; fi6 < 15; fi6++ {
            if len(v1) > 0 {
                v1[0] = ""
            }
            for fi7 := 0; fi7 < 11; fi7++ {
                v3 = true
                if (42 <= -652) {
                    var v8 int64 = -1
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v8
                } else {
                    switch (^(fi6)) {
                    case -421:
                        if (v3 && v3) {
                            if (548 > 2) {
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v9 uint8 = 126
                                    var v10 uint8 = v9
                                    var v11 string = "a"
                                    switch "hello" {
                                    case "a":
                                        var v12 float32 = 1e10
                                        _ = v12
                                    case "\n\t":
                                        if (v3 || v3) {
                                            switch (-(fi7)) {
                                            case 0:
                                                var v13 uint64 = 4294967295
                                                _ = v13
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = (^(fi7))
                                            }
                                        }
                                    }
                                    _ = v10
                                    _ = v11
                                    _ = v9
                                }
                            }
                            fmt.Println(fi6, v3, fi7)
                            if false {
                                fmt.Println(v3)
                                switch (-(fi7)) {
                                case 128:
                                    _ = len(v2)
                                case 1:
                                    switch "hello" {
                                    case "\n\t":
                                        var v14 int32 = -417
                                        _ = v14
                                    default:
                                        var v15 float64 = (((1.5 + -2.25) + 0.0) - 1e10)
                                        _ = v15
                                    }
                                case -2147483648:
                                    if false {
                                        _ = len(v1)
                                        v16 := GenS1{F0: 65535, F1: 0}
                                        var v17 uint8 = 255
                                        _ = v16
                                        _ = v17
                                    }
                                }
                                v18 := GenS1{F0: 2147483648, F1: 4294967295}
                                _ = v18
                            }
                            v3 = false
                            if v3 {
                                v3 = (!v3)
                                switch ("hello" + "") {
                                case "foo bar":
                                    fmt.Println(v3, fi7, fi6)
                                case "":
                                    var v19 uint64 = 307
                                    _ = v19
                                case "a":
                                    fmt.Println(fi7)
                                }
                                fmt.Println(v3)
                            }
                            v20 := GenS2{F0: 0}
                            _ = v20
                        }
                    case 127:
                        v3 = v3
                    }
                }
                v3 = (!v3)
                fmt.Println(v3, v3, fi6)
                _ = fi7
            }
            if false {
                var v21 int = fi6
                _ = v21
            }
            _ = fi6
        }
    case "a":
        v3 = true
    case "hello":
        var v22 int16 = -543
        _ = v22
    }
    var v23 int = 803
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 16; fi25++ {
            v23 -= v23
            var v26 uint64 = 4294967295
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        if (18446744073709551615 >= 42) {
            v3 = v3
        }
        _ = fi27
    }
    switch v23 {
    case 9223372036854775807:
        for fi28 := 0; fi28 < 1; fi28++ {
            v23 = (v23 << ((65535) & 63))
            _ = fi28
        }
    }
    v23 = v23
    if v3 {
        v23 += -2147483648
        fmt.Println(v23, v23)
        var v29 int32 = 576
        _ = v29
    } else {
        fmt.Println(v3, v3)
        v3 = (float32(1e10) > 3.14159)
        fmt.Println(v23)
    }
    fmt.Println(v3, v3)
    for fi30 := 0; fi30 < 13; fi30++ {
        v1 = append(v1, "hello")
        _ = fi30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        v3 = ((42 <= 4294967295) || v3)
        for fi32 := 0; fi32 < 12; fi32++ {
            switch -2147483648 {
            case -2147483648:
                v23 = -9223372036854775808
            case -9223372036854775808:
                fmt.Println(v3, fi31)
            case 0:
                v3 = (4294967295 != 4294967295)
                var v33 uint = 947
                switch (fi32 | (v23 - 9223372036854775807)) {
                case 143:
                    var v34 uint16 = 297
                    _ = v34
                }
                _ = v33
            }
            _ = fi32
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v5
}
