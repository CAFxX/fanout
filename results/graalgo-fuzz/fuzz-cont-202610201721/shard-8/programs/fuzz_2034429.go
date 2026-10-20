package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string }

func main() {
    var v1 float64 = 0.0
    var v2 uint16 = 255
    var v3 uint64 = 0
    var v4 uint32 = 2147483648
    v5 := []string{"foo bar", "foo bar", (("hello" + "a") + "hello")}
    var v6 uint64 = v3
    if len(v5) > 0 {
        v5[0] = ("foo bar" + "")
    }
    switch ("foo bar" + "hello") {
    case "a":
        switch -182 {
        case -2147483648:
            v7 := GenS1{F0: 2147483647, F1: -9223372036854775808, F2: 42}
            _ = v7
        case 42:
            fmt.Println(v6, v4, v1)
        }
    case "hello":
        switch 128 {
        case 4294967295:
            var v8 float64 = v1
            _ = v8
        case -652:
            v5 = append(v5, "foo bar")
        case 2147483647:
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 uint = 883
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                if true {
                    if (false || false) {
                        var v11 string = "\n\t"
                        switch 0 {
                        case -9223372036854775808:
                            var v12 uint16 = v2
                            _ = v12
                        case 4294967295:
                            v13 := GenS2{F0: 42, F1: 1000}
                            _ = v13
                        case -2147483648:
                            var v14 int = fi9
                            _ = v14
                        default:
                            v4 -= v4
                        }
                        v10 -= v10
                        _ = v11
                    } else {
                        switch fi9 {
                        case -368:
                            var v15 float64 = (v1 + 0.0)
                            _ = v15
                        case 0:
                            switch (fi9 % (((fi9 * 9223372036854775807)) | 1)) {
                            case 128:
                                for fi16 := 0; fi16 < 14; fi16++ {
                                    v17 := GenS2{F0: 42, F1: 793}
                                    v4 = 61
                                    if len(v5) > 0 {
                                        v5[0] = ("\n\t" + "\n\t")
                                    }
                                    v18 := GenS2{F0: int64(v2), F1: 1000}
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                            case 127:
                                if true {
                                    v3 += v6
                                    v19 := GenS1{F0: -2147483648, F1: -421, F2: -1}
                                    _ = v19
                                } else {
                                    _ = len(v5)
                                }
                            }
                        default:
                            v4 = v4
                        }
                        v1 = (v1 + v1)
                        v10 = (^(v10))
                    }
                } else {
                    var v20 uint8 = 0
                    for fi21 := 0; fi21 < 18; fi21++ {
                        if ((v20 % ((v20) | 1)) != 127) {
                            v3 = (v3 & v3)
                            var v22 uint = v10
                            v23 := GenS3{F0: v2}
                            _ = v22
                            _ = v23
                        } else {
                            v24 := GenS4{F0: (v4 << (((^(v10))) & 63))}
                            _ = v24
                        }
                        v25 := GenS2{F0: int64(v3), F1: -1}
                        _ = fi21
                        _ = v25
                    }
                    _ = v20
                }
                _ = fi9
                _ = v10
            }
        }
    case "foo bar":
        var v26 bool = true
        _ = v26
    }
    switch "a" {
    case "hello":
        fmt.Println(v2, v4)
    case "\n\t":
        for fi27 := 0; fi27 < 12; fi27++ {
            if true {
                var v28 uint16 = (v2 ^ v2)
                _ = v28
            } else {
                for fi29 := 0; fi29 < 10; fi29++ {
                    if (!false) {
                        fmt.Println(fi27)
                        switch (("hello" + "hello") + "a") {
                        case "a":
                            fmt.Println(fi29)
                            var v30 uint8 = uint8(v4)
                            _ = v30
                        }
                    } else {
                        v31 := GenS3{F0: v2}
                        fmt.Println(v2, v1, v1)
                        _ = v31
                    }
                    v3 = v6
                    _ = fi29
                }
            }
            v6 += v3
            for fi32 := 0; fi32 < 17; fi32++ {
                v33 := GenS3{F0: (^(v2))}
                v6 = (v3 / ((v3) | 1))
                if len(v5) > 0 {
                    v5[0] = ""
                }
                _ = fi32
                _ = v33
            }
            _ = fi27
        }
    case "a":
        _ = len(v5)
    default:
        if false {
            v2 += 2
            if true {
                if (false || true) {
                    v5 = append(v5, "a")
                    var v34 int = 2147483647
                    _ = v34
                } else {
                    var v35 int8 = 42
                    v4 = (v4 & (v4 + v4))
                    _ = v35
                }
                var v36 uint = 4294967295
                v4 = (v4 & v4)
                _ = v36
            }
            fmt.Println(v2)
        } else {
            var v37 float32 = float32((v1 + ((1.5 * float64(-2.25)) - v1)))
            _ = v37
        }
    }
    fmt.Println(v4, v2, v4)
    var v38 int = -1
    v39 := GenS5{F0: 32768}
    var v40 uint8 = 127
    if false {
        if false {
            for fi41 := 0; fi41 < 1; fi41++ {
                for fi42 := 0; fi42 < 9; fi42++ {
                    var v43 float64 = (3.14159 - ((v1 - v1) * v1))
                    var v44 string = ""
                    if (false || false) {
                        _ = len(v5)
                    }
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                _ = fi41
            }
            _ = len(v5)
        } else {
            v45 := GenS6{F0: 127, F1: 0}
            var v46 int8 = 100
            var v47 float64 = v1
            _ = v45
            _ = v46
            _ = v47
        }
    }
    if (-2147483648 != -2147483648) {
        switch "" {
        case "foo bar":
            var v48 float32 = (1.5 + -2.25)
            _ = v48
        case "\n\t":
            var v49 int = 1
            _ = v49
        case "":
            v3 = (v6 / ((v3) | 1))
        }
        v40 = (v40 / ((v40) | 1))
    }
    v50 := GenS3{F0: (^(v2))}
    for fi51 := 0; fi51 < 9; fi51++ {
        for fi52 := 0; fi52 < 20; fi52++ {
            if (127 != 0) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v53 int32 = 1
                if len(v5) > 0 {
                    v5[0] = "a"
                }
                _ = v53
            } else {
                v3 = 4294967295
                if len(v5) > 0 {
                    v5[0] = "a"
                }
            }
            v3 = 18446744073709551615
            fmt.Println(v38, fi51)
            _ = fi52
        }
        _ = fi51
    }
    v54 := GenS7{F0: 61}
    v40 = v40
    v55 := GenS8{F0: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v50
    _ = v54
    _ = v55
    _ = v6
}
