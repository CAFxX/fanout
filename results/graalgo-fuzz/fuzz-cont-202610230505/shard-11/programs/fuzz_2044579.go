package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 uint16; F2 string }
type GenS7 struct { F0 uint64; F1 int32 }
func genFn0(p0 int64) int {
    v7 := []int{32767}
    v7 = append(v7, 0)
    p0 = (p0 & (p0 - (p0 / ((-9223372036854775808) | 1))))
    _ = p0
    _ = v7
    return 32767
}

func main() {
    var v1 uint16 = 2
    var v2 uint = 2
    v3 := []uint32{1, 622, uint32(v2), 65535}
    v4 := []int{-2147483648, -328, -298}
    v5 := []string{((((("" + "a") + ("\n\t" + ("" + "foo bar"))) + (((("a" + "") + "") + "hello") + "foo bar")) + "hello") + "foo bar")}
    v6 := make(map[string]bool)
    fmt.Println(v1, v2)
    v8 := GenS1{F0: 2147483647, F1: 2}
    v9 := GenS1{F0: 42, F1: (^(v1))}
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            fmt.Println(v1)
            _ = fi11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v1 = (v1 - (^(v1)))
            v13 := GenS1{F0: fi12, F1: v1}
            _ = fi12
            _ = v13
        }
        _ = fi10
    }
    v14 := GenS2{F0: 128}
    for fi15 := 0; fi15 < 19; fi15++ {
        if false {
            if (575 <= 32767) {
                v2 += 255
                v16 := GenS2{F0: fi15}
                _ = v16
            } else {
                v2 += 277
            }
            _ = len(v6)
            fmt.Println(fi15)
        } else {
            if false {
                switch ("hello" + "hello") {
                case "a":
                    var v17 uint8 = 126
                    _ = v17
                case "foo bar":
                    v1 += 740
                case "\n\t":
                    var v18 int = fi15
                    _ = v18
                }
                var v19 int8 = -77
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 int16 = -1
                    if ("" > "") {
                        fmt.Println(v19, v1, v19)
                        if true {
                            v22 := GenS2{F0: 105}
                            var v23 float64 = (3.14159 + 0.1)
                            _ = v22
                            _ = v23
                        } else {
                            switch genFn0(int64(v21)) {
                            case 32768:
                                var v24 int64 = 9223372036854775807
                                _ = v24
                            case 521:
                                fmt.Println(v1, v2)
                            default:
                                if false {
                                    var v25 uint8 = 42
                                    var v26 float32 = ((1.5 - 1.5) + 1e10)
                                    v27 := GenS1{F0: fi20, F1: (v1 ^ (v1 + uint16(v25)))}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                } else {
                                    if (!true) {
                                        v28 := GenS3{F0: v2, F1: -850}
                                        v29 := v6["foo bar"]
                                        v29 = true
                                        _ = v28
                                        _ = v29
                                    }
                                    if (false || true) {
                                        fmt.Println(v1, fi15)
                                        var v30 uint = v2
                                        v31 := GenS4{F0: -839}
                                        _ = v30
                                        _ = v31
                                    } else {
                                        switch "hello" {
                                        case "a":
                                            if false {
                                                var v32 int8 = 127
                                                if false {
                                                    v6["\n\t"] = (true && (131 < -732))
                                                }
                                                v33 := GenS5{F0: 802}
                                                _ = v32
                                                _ = v33
                                            }
                                        case "hello":
                                            v34 := GenS6{F0: 65535, F1: (v1 - (v1 ^ 1)), F2: "hello"}
                                            _ = v34
                                        case "\n\t":
                                            if len(v4) > 0 {
                                                v4[0] = fi15
                                            }
                                        default:
                                            var v35 uint32 = 473
                                            _ = v35
                                        }
                                        if true {
                                            var v36 bool = (!false)
                                            _ = v36
                                        } else {
                                            switch (fi15 ^ (fi20 & (fi15 / (((fi20 ^ (fi20 % (((-(fi20))) | 1)))) | 1)))) {
                                            case 612:
                                                switch fi15 {
                                                case -9223372036854775808:
                                                    v37 := GenS4{F0: -1}
                                                    _ = v37
                                                default:
                                                    if true {
                                                        v1 += 251
                                                        if len(v5) > 0 {
                                                            v5[0] = ("foo bar" + (("hello" + "a") + ""))
                                                        }
                                                    }
                                                }
                                            case 127:
                                                var v38 int32 = int32(v19)
                                                _ = v38
                                            case 1:
                                                fmt.Println(v19, v1)
                                                var v39 int32 = int32(v2)
                                                var v40 float32 = -2.25
                                                if true {
                                                    var v41 uint8 = 127
                                                    _ = v41
                                                }
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        fmt.Println(v21, v21, v19)
                                    }
                                }
                            }
                            var v42 int8 = v19
                            _ = v42
                        }
                        v2 = 1
                    } else {
                        if true {
                            var v43 int = fi20
                            v44 := GenS7{F0: 860, F1: 282}
                            _ = v43
                            _ = v44
                        }
                    }
                    v2 = (v2 ^ v2)
                    switch (fi20 << ((v2) & 63)) {
                    case 0:
                        if false {
                            v45 := GenS5{F0: uint64(v2)}
                            _ = v45
                        } else {
                            var v46 int8 = 71
                            var v47 int64 = 139
                            var v48 int16 = 32767
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                    }
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                switch genFn0(42) {
                case 2147483647:
                    v49 := v6["a"]
                    _ = v49
                case 127:
                    switch ("\n\t" + (("\n\t" + ("" + "\n\t")) + "a")) {
                    case "\n\t":
                        _ = len(v4)
                    }
                }
            }
            var v50 uint16 = (^(v1))
            _ = v50
        }
        _ = fi15
    }
    v2 = v2
    switch ("\n\t" + ("hello" + "hello")) {
    case "\n\t":
        fmt.Println(v1)
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        var v52 float64 = -2.25
        v53 := GenS5{F0: 1}
        _ = fi51
        _ = v52
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
