package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8; F1 uint16 }
func genFn0() string {
    if (1 <= 754) {
        var v4 int = -353
        var v5 int16 = -1
        _ = v4
        _ = v5
    }
    var v6 bool = (1e10 <= 0.1)
    v7 := GenS1{F0: -685}
    _ = v6
    _ = v7
    return ""
}
func genFn1() string {
    var v8 uint64 = 42
    v9 := GenS2{F0: 100, F1: (v8 % ((680) | 1)), F2: 12}
    switch 42 {
    case -195:
        var v10 int8 = 0
        _ = v10
    case 404:
        if (("foo bar" != ("hello" + "foo bar")) && ("a" == "foo bar")) {
            v8 = v8
            v8 += v8
            for fi11 := 0; fi11 < 20; fi11++ {
                switch fi11 {
                case -1:
                    switch (("hello" + "\n\t") + "foo bar") {
                    case "\n\t":
                        var v12 int16 = 42
                        _ = v12
                    case "a":
                        v13 := GenS3{F0: -1, F1: ""}
                        _ = v13
                    case "hello":
                        v8 -= (^(v8))
                    default:
                        if (638 >= 254) {
                            switch (fi11 ^ fi11) {
                            case 1:
                                var v14 uint = 42
                                _ = v14
                            }
                        }
                    }
                }
                _ = fi11
            }
        } else {
            v8 = (v8 % ((v8) | 1))
            v15 := GenS4{F0: 42, F1: false}
            switch (((("\n\t" + ((("" + "hello") + "") + "a")) + ("foo bar" + "\n\t")) + (("" + "\n\t") + ("" + "a"))) + "a") {
            case "":
                v16 := GenS5{F0: 718, F1: -278}
                _ = v16
            case "a":
                for fi17 := 0; fi17 < 6; fi17++ {
                    fmt.Println(fi17)
                    _ = fi17
                }
            case "\n\t":
                switch ("" + ("\n\t" + "hello")) {
                case "":
                    if true {
                        for fi18 := 0; fi18 < 17; fi18++ {
                            var v19 uint64 = v8
                            _ = fi18
                            _ = v19
                        }
                        var v20 uint64 = v8
                        _ = v20
                    }
                case "\n\t":
                    v8 += (v8 + (^(v8)))
                case "a":
                    v21 := []int{-2147483648, 127}
                    _ = v21
                default:
                    v8 = (v8 >> ((uint(v8)) & 63))
                }
            }
            _ = v15
        }
    case 4294967295:
        v8 += v8
    default:
        fmt.Println(v8)
    }
    v8 -= (^(v8))
    v22 := []int{2147483647}
    _ = v22
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    v1 := []int{4294967295, 4294967295}
    var v2 bool = false
    v3 := make(map[int]int)
    v1 = append(v1, 42)
    v2 = true
    for fi23 := 0; fi23 < 11; fi23++ {
        v3[fi23] = (-(fi23))
        _ = fi23
    }
    v2 = ((2147483648 <= 521) || false)
    if (true || v2) {
        v2 = v2
        var v24 uint32 = 2147483648
        var v25 int = 42
        _ = v24
        _ = v25
    } else {
        _ = len(v3)
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
        if false {
            v26 := GenS1{F0: 4294967295}
            var v27 int64 = 1
            _ = v26
            _ = v27
        }
    }
    if (true || (4294967295 <= 1)) {
        v2 = v2
        var v28 int16 = 39
        _ = v28
    }
    var v29 uint16 = 1
    switch 402 {
    case 32767:
        v2 = v2
    case 9223372036854775807:
        v29 = 32767
    default:
        v29 = 65535
    }
    var v30 uint8 = 0
    v31 := GenS6{F0: 724}
    switch ("foo bar" + "a") {
    case "a":
        var v32 string = "foo bar"
        _ = v32
    default:
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 float32 = -2.25
            v2 = v2
            _ = fi33
            _ = v34
        }
    }
    switch -9223372036854775808 {
    case 42:
        var v35 int8 = 100
        _ = v35
    case -9223372036854775808:
        var v36 int = 1
        _ = v36
    case 9223372036854775807:
        for fi37 := 0; fi37 < 9; fi37++ {
            v1 = append(v1, int(v30))
            var v38 int32 = -1
            _ = fi37
            _ = v38
        }
    }
    if false {
        v39 := GenS3{F0: -1, F1: ""}
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
        v1 = append(v1, -2147483648)
        _ = v39
    } else {
        if v2 {
            if (("hello" != "foo bar") && (127 <= 48)) {
                fmt.Println(v30, v29)
            }
        } else {
            var v40 uint64 = 18446744073709551615
            for fi41 := 0; fi41 < 12; fi41++ {
                v42 := GenS5{F0: -769, F1: -41}
                _ = fi41
                _ = v42
            }
            v2 = false
            _ = v40
        }
    }
    v43 := GenS7{F0: (v30 | (v30 & v30)), F1: (v29 % ((v29) | 1))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v43
}
