package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }

func main() {
    var v1 bool = (!((!true) || true))
    var v2 int = 127
    var v3 int8 = -1
    if ((471 == 5) && false) {
        fmt.Println(v2)
    }
    for fi4 := 0; fi4 < 14; fi4++ {
        v2 = (v2 & 127)
        switch "hello" {
        case "":
            switch (((("foo bar" + "\n\t") + (("" + "") + "")) + "foo bar") + "foo bar") {
            case "a":
                var v5 int64 = -1
                _ = v5
            case "foo bar":
                v3 = (-(v3))
            default:
                var v6 uint16 = 32767
                _ = v6
            }
        case "hello":
            if true {
                v1 = (v1 || true)
                if true {
                    fmt.Println(v3, v2)
                    if (v1 && true) {
                        var v7 int16 = 32767
                        v1 = v1
                        for fi8 := 0; fi8 < 5; fi8++ {
                            var v9 int32 = 637
                            v3 = 100
                            v9 = (v9 * 957)
                            _ = fi8
                            _ = v9
                        }
                        _ = v7
                    } else {
                        var v10 uint8 = 255
                        v3 += (v3 / ((-128) | 1))
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v12 := GenS1{F0: 2147483648, F1: 0}
                        _ = fi11
                        _ = v12
                    }
                } else {
                    fmt.Println(fi4, v2)
                    v2 -= 2147483647
                }
                switch "\n\t" {
                case "\n\t":
                    var v13 uint64 = 0
                    _ = v13
                case "":
                    v1 = v1
                case "hello":
                    v14 := []int{42, v2}
                    _ = v14
                }
            }
        case "\n\t":
            fmt.Println(fi4)
        default:
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 bool = (76 <= 105)
                fmt.Println(fi15, fi4, v2)
                v16 = (v1 || v16)
                switch (("\n\t" + "") + (("\n\t" + ("" + "a")) + "")) {
                case "\n\t":
                    var v17 int16 = -32768
                    fmt.Println(v17)
                    switch "\n\t" {
                    case "\n\t":
                        var v18 int8 = v3
                        _ = v18
                    case "a":
                        v19 := GenS1{F0: 304, F1: 902}
                        _ = v19
                    default:
                        v20 := []int{}
                        _ = v20
                    }
                    _ = v17
                case "hello":
                    v21 := GenS1{F0: 704, F1: 83}
                    _ = v21
                case "foo bar":
                    if (!v16) {
                        v22 := []int{(fi15 - -176), v2, (fi15 - (v2 << ((104) & 63))), (-(v2))}
                        v23 := GenS2{F0: 255}
                        switch -1 {
                        case -572:
                            if ((v3 > (v3 ^ (v3 ^ -1))) && (-9223372036854775808 == fi15)) {
                                var v24 float64 = (float64((0.0 + (3.14159 / -2.25))) - (1e10 * 1e10))
                                _ = v24
                            } else {
                                v1 = (false && (125 != 127))
                                if (v16 && (1 <= 2147483647)) {
                                    switch -568 {
                                    case 892:
                                        if false {
                                            v1 = v16
                                        } else {
                                            var v25 uint16 = 834
                                            _ = v25
                                        }
                                    default:
                                        switch 32768 {
                                        case 42:
                                            var v26 uint16 = 65535
                                            _ = v26
                                        default:
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                        }
                                        v16 = v1
                                    }
                                    if len(v22) > 0 {
                                        v22[0] = (fi15 / (((fi15 % (((v2 - 483)) | 1))) | 1))
                                    }
                                } else {
                                    v1 = v1
                                    fmt.Println(v2)
                                    var v27 int64 = -1
                                    _ = v27
                                }
                            }
                        case -678:
                            switch ("a" + "a") {
                            case "foo bar":
                                if len(v22) > 0 {
                                    v22[0] = 32768
                                }
                            default:
                                fmt.Println(fi15, v16, fi15)
                            }
                        default:
                            var v28 int64 = -461
                            _ = v28
                        }
                        _ = v22
                        _ = v23
                    } else {
                        v29 := []int{(-(v2))}
                        v29 = append(v29, fi4)
                        switch ("a" + ("foo bar" + "foo bar")) {
                        case "hello":
                            v30 := GenS3{F0: "a"}
                            _ = v30
                        case "a":
                            var v31 int16 = -1
                            _ = v31
                        case "\n\t":
                            var v32 int = (fi4 % ((4294967295) | 1))
                            _ = v32
                        default:
                            fmt.Println(v1, fi15, v3)
                        }
                        _ = v29
                    }
                }
                v33 := GenS3{F0: (("" + (((("a" + "") + ("\n\t" + ("a" + "a"))) + "\n\t") + ("" + "hello"))) + "\n\t")}
                v34 := GenS2{F0: 2}
                _ = fi15
                _ = v16
                _ = v33
                _ = v34
            }
        }
        _ = fi4
    }
    v35 := GenS1{F0: 562, F1: 2147483647}
    fmt.Println(v1, v3)
    var v36 uint8 = 206
    v3 -= (v3 << ((18446744073709551615) & 63))
    if (2147483647 <= 2147483648) {
        for fi37 := 0; fi37 < 14; fi37++ {
            switch 127 {
            case -9223372036854775808:
                v3 += 100
            case 128:
                v38 := GenS3{F0: "a"}
                _ = v38
            default:
                fmt.Println(fi37, v3)
            }
            _ = fi37
        }
        v39 := []int{(v2 & (v2 + (^(v2))))}
        _ = v39
    } else {
        fmt.Println(v36, v2, v2)
        var v40 uint8 = 222
        var v41 uint32 = 1
        _ = v40
        _ = v41
    }
    v1 = v1
    var v42 int16 = 473
    v43 := GenS4{F0: (v2 & v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v42
    _ = v43
}
