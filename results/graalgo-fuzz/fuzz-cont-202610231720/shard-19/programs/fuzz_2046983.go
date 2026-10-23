package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int8; F2 uint64 }
type GenS4 struct { F0 int8; F1 int16; F2 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 string, p1 int32) uint64 {
    fmt.Println(p0, p0)
    p0 = "hello"
    p0 = p0
    if (!(true || false)) {
        switch ("hello" + ("hello" + p0)) {
        case "hello":
            v5 := GenS1{F0: false}
            _ = v5
        case "foo bar":
            v6 := GenS1{F0: false}
            _ = v6
        case "":
            fmt.Println(p1)
        }
        if true {
            var v7 int16 = -32768
            _ = v7
        }
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    v1 := []int{127, 32767, 2147483647, -9223372036854775808}
    v2 := make(map[string]string)
    var v3 uint64 = 321
    var v4 uint16 = 65535
    v8 := v2[""]
    v4 = (v4 - v4)
    v9 := GenS2{F0: v3}
    switch 263 {
    case 128:
        if true {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            var v10 bool = (int64(v3) < 9223372036854775807)
            _ = v10
        }
    case 1:
        for fi11 := 0; fi11 < 12; fi11++ {
            switch (((("hello" + "\n\t") + v8) + v8) + v8) {
            case "":
                fmt.Println(fi11, v3)
            case "\n\t":
                v1 = append(v1, 120)
            }
            v8 = "a"
            var v12 int32 = 42
            _ = fi11
            _ = v12
        }
    case -773:
        v13 := GenS1{F0: true}
        _ = v13
    }
    if len(v1) > 0 {
        v1[0] = 128
    }
    switch v8 {
    case "foo bar":
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: 1}
            if true {
                if (-193 <= 0) {
                    v8 = ""
                    for fi16 := 0; fi16 < 11; fi16++ {
                        switch (fi14 >> ((42) & 63)) {
                        case 882:
                            var v17 uint64 = 653
                            _ = v17
                        case 32767:
                            v4 = v4
                        case 42:
                            v3 -= v3
                        }
                        switch "foo bar" {
                        case "foo bar":
                            switch (fi16 / ((1) | 1)) {
                            case 128:
                                if true {
                                    fmt.Println(v4, v3)
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = (fi14 + 127)
                                    }
                                    var v18 int = (fi14 & (-(fi14)))
                                    _ = v18
                                }
                            }
                        case "":
                            v8 = (((((v8 + v8) + "foo bar") + v8) + v8) + v8)
                        case "a":
                            fmt.Println(v8, v3)
                        }
                        var v19 int8 = 0
                        _ = fi16
                        _ = v19
                    }
                    var v20 uint32 = 2147483647
                    _ = v20
                }
                if (1000 > 544) {
                    v3 = (v3 >> ((334) & 63))
                    fmt.Println(v4, fi14, fi14)
                    v8 = (((((v8 + (((v8 + "") + v8) + (("\n\t" + ((v8 + (v8 + ("foo bar" + "\n\t"))) + ((((v8 + ((v8 + ("a" + v8)) + v8)) + "hello") + "foo bar") + v8))) + v8))) + v8) + v8) + (v8 + v8)) + "foo bar")
                } else {
                    var v21 uint8 = 127
                    _ = v21
                }
                v22 := GenS3{F0: 1, F1: 42, F2: v3}
                _ = v22
            } else {
                v2[(("" + "\n\t") + "\n\t")] = "\n\t"
            }
            _ = fi14
            _ = v15
        }
    default:
        v1 = append(v1, 32767)
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 int32 = 42
        _ = fi23
        _ = v24
    }
    v4 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v25 := GenS4{F0: -128, F1: 32767, F2: 656}
    var v26 int16 = 0
    if (v26 == (-(v26))) {
        if (0 != 9223372036854775807) {
            var v27 int = 0
            var v28 int64 = 9223372036854775807
            v29 := GenS5{F0: -1}
            _ = v27
            _ = v28
            _ = v29
        } else {
            var v30 string = (((v8 + v8) + "foo bar") + (v8 + ((((v8 + ("a" + ((v8 + ("foo bar" + v8)) + "a"))) + ((("" + v8) + (v8 + "foo bar")) + v8)) + "foo bar") + v8)))
            v3 += v3
            for fi31 := 0; fi31 < 9; fi31++ {
                v4 += (v4 >> ((42) & 63))
                v4 = v4
                _ = fi31
            }
            _ = v30
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
