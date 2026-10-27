package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 string; F1 int64; F2 bool }

func main() {
    var v1 int64 = 215
    var v2 uint64 = 2
    v3 := []string{"hello", ("" + "\n\t"), ("foo bar" + ("hello" + ("\n\t" + ((("" + "\n\t") + "") + ("hello" + "foo bar"))))), "hello"}
    var v4 int16 = 0
    v5 := GenS1{F0: 255}
    var v6 int32 = 1
    switch "hello" {
    case "a":
        v7 := GenS1{F0: 251}
        _ = v7
    case "":
        var v8 int64 = v1
        _ = v8
    case "foo bar":
        v1 += v1
    default:
        v3 = append(v3, "foo bar")
    }
    if false {
        v9 := GenS1{F0: 0}
        if ("a" == (("" + "a") + ("\n\t" + "\n\t"))) {
            v10 := GenS1{F0: 172}
            _ = v10
        } else {
            var v11 uint64 = (v2 ^ v2)
            for fi12 := 0; fi12 < 11; fi12++ {
                if false {
                    var v13 int = (fi12 | fi12)
                    _ = v13
                }
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 uint8 = 128
                    v11 -= (v2 << ((0) & 63))
                    v4 -= (v4 | v4)
                    _ = fi14
                    _ = v15
                }
                _ = fi12
            }
            _ = v11
        }
        _ = len(v3)
        _ = v9
    } else {
        v6 = 1
    }
    fmt.Println(v4)
    fmt.Println(v1, v2, v2)
    for fi16 := 0; fi16 < 13; fi16++ {
        if (522 <= 255) {
            v17 := GenS1{F0: 75}
            v18 := GenS1{F0: 127}
            v19 := GenS1{F0: 0}
            _ = v17
            _ = v18
            _ = v19
        }
        fmt.Println(v1, v6, v2)
        _ = fi16
    }
    fmt.Println(v1, v6, v6)
    v20 := GenS1{F0: 2}
    if (!true) {
        var v21 uint = 130
        if (((0.1 - 3.14159) * 1e10) > 1.5) {
            v22 := GenS2{F0: v6, F1: 404, F2: (v21 + 42)}
            if len(v3) > 0 {
                v3[0] = "a"
            }
            _ = len(v3)
            _ = v22
        } else {
            var v23 float32 = 3.14159
            if (!false) {
                var v24 string = "hello"
                if (false && false) {
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v25 uint = (v21 >> (((v21 >> (((v21 + 18446744073709551615)) & 63))) & 63))
                        _ = v25
                    }
                    if true {
                        var v26 string = ""
                        var v27 float64 = 1.5
                        _ = v26
                        _ = v27
                    } else {
                        v24 = "\n\t"
                        fmt.Println(v4, v21)
                        for fi28 := 0; fi28 < 4; fi28++ {
                            v29 := GenS3{F0: "foo bar", F1: (fi28 + (fi28 | fi28))}
                            var v30 int64 = -1
                            if len(v3) > 0 {
                                v3[0] = ((v24 + v24) + "")
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    if (false && false) {
                        v3 = append(v3, v24)
                        v23 = v23
                    } else {
                        v31 := GenS4{F0: v24, F1: (v1 - v1), F2: true}
                        _ = len(v3)
                        switch v24 {
                        case "\n\t":
                            fmt.Println(v2)
                        case "foo bar":
                            var v32 uint32 = 4294967295
                            _ = v32
                        case "a":
                            var v33 uint64 = v2
                            _ = v33
                        default:
                            var v34 uint64 = (v2 + v2)
                            _ = v34
                        }
                        _ = v31
                    }
                }
                if true {
                    var v35 int64 = (v1 | (^(v1)))
                    if true {
                        switch 42 {
                        case 0:
                            var v36 int = 203
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 12; fi37++ {
                            v4 += int16(fi37)
                            var v38 uint64 = v2
                            v21 = v21
                            _ = fi37
                            _ = v38
                        }
                        v21 += (^(v21))
                    }
                    for fi39 := 0; fi39 < 20; fi39++ {
                        var v40 int32 = -290
                        fmt.Println(v40, v40)
                        _ = fi39
                        _ = v40
                    }
                    _ = v35
                }
                _ = v24
            }
            _ = v23
        }
        _ = v21
    }
    if ((!((!true) && true)) && (771 < 65535)) {
        v4 = v4
        v41 := GenS3{F0: "hello", F1: 814}
        v42 := GenS2{F0: (v6 >> ((114) & 63)), F1: 709, F2: 42}
        _ = v41
        _ = v42
    }
    fmt.Println(v6, v6)
    var v43 float64 = (0.0 - 1e10)
    for fi44 := 0; fi44 < 10; fi44++ {
        var v45 int = -771
        fmt.Println(v4, fi44, v45)
        _ = fi44
        _ = v45
    }
    v6 = (v6 % ((v6) | 1))
    var v46 int64 = (v1 << ((1) & 63))
    v46 -= (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v43
    _ = v46
    _ = v5
    _ = v6
}
