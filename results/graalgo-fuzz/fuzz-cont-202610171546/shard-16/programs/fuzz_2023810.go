package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint; F1 uint8; F2 string }
func genFn0(p0 bool, p1 int8, p2 string) uint16 {
    v6 := []int{308, 77, 4294967295, 128}
    if (false || false) {
        p1 -= (-(p1))
        fmt.Println(p2)
        if false {
            var v7 int32 = 2147483647
            v8 := GenS1{F0: 42, F1: 1, F2: 42}
            for fi9 := 0; fi9 < 15; fi9++ {
                fmt.Println(p0)
                _ = fi9
            }
            _ = v7
            _ = v8
        } else {
            switch p2 {
            case "hello":
                v10 := GenS2{F0: 808}
                _ = v10
            case "\n\t":
                if (347 >= 2) {
                    v11 := GenS3{F0: 42, F1: p1, F2: 2}
                    if len(v6) > 0 {
                        v6[0] = 42
                    }
                    _ = v11
                } else {
                    fmt.Println(p0, p0, p1)
                }
            case "":
                var v12 int16 = 1000
                _ = v12
            }
        }
    }
    var v13 int8 = (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"a", "hello", "foo bar", (((("foo bar" + (("a" + ("\n\t" + "\n\t")) + ("" + "\n\t"))) + "hello") + ((("\n\t" + "hello") + "\n\t") + ("hello" + "a"))) + "")}
    var v3 int32 = 42
    var v4 uint = 917
    var v5 int64 = -598
    if false {
        for fi14 := 0; fi14 < 9; fi14++ {
            fmt.Println(fi14, v3, v5)
            if true {
                _ = len(v1)
                fmt.Println(v4, fi14)
                switch 2147483647 {
                case 0:
                    for fi15 := 0; fi15 < 8; fi15++ {
                        switch "" {
                        case "foo bar":
                            switch (^(fi14)) {
                            case 0:
                                var v16 int32 = int32(fi15)
                                _ = v16
                            default:
                                v1[("a" + "\n\t")] = fi14
                            }
                        default:
                            v3 = v3
                        }
                        _ = fi15
                    }
                default:
                    for fi17 := 0; fi17 < 20; fi17++ {
                        var v18 uint = v4
                        v19 := GenS2{F0: 2}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
            }
            v20 := GenS3{F0: 4294967295, F1: 100, F2: 42}
            _ = fi14
            _ = v20
        }
        var v21 uint64 = 0
        _ = v21
    }
    switch "hello" {
    case "":
        v22 := GenS4{F0: false, F1: 418, F2: 55}
        _ = v22
    }
    fmt.Println(v3, v3, v4)
    var v23 int32 = (^(v3))
    fmt.Println(v23, v4, v23)
    delete(v1, ("foo bar" + "\n\t"))
    v5 = v5
    if (false && true) {
        for fi24 := 0; fi24 < 8; fi24++ {
            for fi25 := 0; fi25 < 12; fi25++ {
                if true {
                    switch (-(fi25)) {
                    case 4294967295:
                        var v26 int8 = 42
                        v23 += (v23 >> (((v4 / ((0) | 1))) & 63))
                        v5 += v5
                        _ = v26
                    case -2147483648:
                        v3 += (v23 + v3)
                    }
                    fmt.Println(fi24)
                }
                v3 -= -1
                _ = fi25
            }
            _ = fi24
        }
    } else {
        v27 := GenS1{F0: 853, F1: 42, F2: v4}
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS5{F0: 42, F1: uint8(v4), F2: ("" + "hello")}
            delete(v1, ("foo bar" + "hello"))
            var v30 int = fi28
            _ = fi28
            _ = v29
            _ = v30
        }
        v3 -= v3
        _ = v27
    }
    fmt.Println(v3)
    var v31 float64 = (0.0 - 0.1)
    var v32 uint = v4
    for fi33 := 0; fi33 < 2; fi33++ {
        for fi34 := 0; fi34 < 9; fi34++ {
            if (1.5 >= 3.14159) {
                var v35 uint16 = 42
                v36 := GenS2{F0: 65535}
                v5 -= v5
                _ = v35
                _ = v36
            }
            v31 = 1.5
            var v37 uint8 = 19
            _ = fi34
            _ = v37
        }
        for fi38 := 0; fi38 < 15; fi38++ {
            switch "foo bar" {
            case "\n\t":
                var v39 bool = true
                _ = v39
            case "":
                fmt.Println(v31, v23, v31)
            case "foo bar":
                v32 = v4
            }
            if (!(fi33 >= (fi33 - fi33))) {
                var v40 int64 = v5
                _ = v40
            } else {
                var v41 int64 = (v5 << ((v4) & 63))
                v31 = v31
                v1["foo bar"] = fi38
                v42 := v1[("hello" + "")]
                _ = v41
                _ = v42
            }
            v43 := GenS2{F0: 358}
            _ = fi38
            _ = v43
        }
        _ = fi33
    }
    var v44 int = 4294967295
    v3 = (v23 & v23)
    v45 := GenS3{F0: 42, F1: -1, F2: 18446744073709551615}
    var v46 int16 = 42
    var v47 uint8 = 2
    var v48 int16 = (v46 - (v46 | 1000))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
