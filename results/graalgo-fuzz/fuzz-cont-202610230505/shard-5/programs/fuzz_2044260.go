package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int) uint64 {
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 int32 = 2147483647
        v6 := []int{(p0 * 32768)}
        p0 = fi4
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = 1
        _ = fi7
        _ = v8
    }
    switch ((("\n\t" + "a") + (("hello" + "hello") + "a")) + "a") {
    case "hello":
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint16 = 65535
            switch ((("hello" + "hello") + "hello") + (("" + "foo bar") + "foo bar")) {
            case "a":
                v11 := []int{p0, (fi9 >> ((42) & 63))}
                _ = v11
            }
            _ = fi9
            _ = v10
        }
    case "a":
        if (42 > 18446744073709551615) {
            v12 := []int{-822, p0}
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = v12
        }
    case "":
        v13 := GenS1{F0: 65535, F1: 0}
        _ = v13
    }
    _ = p0
    return 194
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("a" + "foo bar")}
    v3 := []string{}
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            fmt.Println(fi15)
            v16 := GenS2{F0: 241, F1: 668, F2: 65535}
            var v17 uint = 18446744073709551615
            _ = fi15
            _ = v16
            _ = v17
        }
        if len(v2) > 0 {
            v2[0] = ("\n\t" + "hello")
        }
        _ = fi14
    }
    var v18 uint16 = 331
    var v19 int8 = 1
    var v20 int = 9223372036854775807
    v2 = append(v2, "foo bar")
    var v21 int32 = -2147483648
    var v22 string = ("a" + "hello")
    var v23 int = int(v21)
    v18 = v18
    for fi24 := 0; fi24 < 14; fi24++ {
        if (true && false) {
            switch (fi24 << ((157) & 63)) {
            case -334:
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 int8 = v19
                    if false {
                        var v27 int32 = v21
                        fmt.Println(v27, v23, v22)
                        _ = v27
                    } else {
                        var v28 bool = true
                        v28 = (v26 <= 29)
                        v20 += (v20 % (((fi25 % ((v23) | 1))) | 1))
                        _ = v28
                    }
                    v29 := GenS1{F0: 2147483647, F1: 231}
                    _ = fi25
                    _ = v26
                    _ = v29
                }
            case -682:
                var v30 int64 = -1
                _ = v30
            case 950:
                v31 := GenS3{F0: v21}
                _ = v31
            }
            var v32 uint = 255
            if (-2.25 == float32(-2.25)) {
                v33 := GenS4{F0: (v22 + (("\n\t" + v22) + "")), F1: (-(v23))}
                _ = v33
            } else {
                switch "hello" {
                case "\n\t":
                    v19 += (v19 * 127)
                case "foo bar":
                    fmt.Println(v32, v18, v19)
                default:
                    for fi34 := 0; fi34 < 18; fi34++ {
                        switch ((v22 + "a") + "") {
                        case "foo bar":
                            if false {
                                fmt.Println(fi24, v22, v23)
                                v20 = v20
                                var v35 uint8 = 131
                                _ = v35
                            } else {
                                v2 = append(v2, ("\n\t" + "foo bar"))
                                v18 = uint16(v32)
                            }
                        case "hello":
                            var v36 int = (fi34 & fi34)
                            _ = v36
                        default:
                            v18 = (v18 / (((v18 / ((255) | 1))) | 1))
                        }
                        _ = fi34
                    }
                }
                var v37 uint64 = 2
                v38 := GenS5{F0: 65535}
                _ = v37
                _ = v38
            }
            _ = v32
        }
        if false {
            fmt.Println(v23)
            delete(v1, "foo bar")
        } else {
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 int16 = 0
                if true {
                    _ = len(v2)
                }
                var v41 int16 = int16(v18)
                _ = fi39
                _ = v40
                _ = v41
            }
            v22 = ((v22 + v22) + v22)
            for fi42 := 0; fi42 < 13; fi42++ {
                var v43 uint = 326
                var v44 float64 = ((((0.0 / 1e10) + 3.14159) / -2.25) / ((1.5 / 1.5) + 3.14159))
                v3 = append(v3, (v22 + "a"))
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        _ = fi24
    }
    var v45 int = v20
    var v46 bool = (v22 >= v22)
    delete(v1, (((v22 + "\n\t") + "a") + v22))
    var v47 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v45
    _ = v46
    _ = v47
}
