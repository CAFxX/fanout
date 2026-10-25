package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int32; F2 int64 }
type GenS6 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0(p0 string, p1 int8, p2 uint) int64 {
    for fi5 := 0; fi5 < 6; fi5++ {
        if (!true) {
            switch p0 {
            case "hello":
                p2 = 65535
            case "a":
                var v6 int = (-(fi5))
                _ = v6
            case "foo bar":
                var v7 int32 = 42
                _ = v7
            default:
                switch (p0 + p0) {
                case "\n\t":
                    p2 = p2
                default:
                    switch (((p0 + "") + p0) + "a") {
                    case "":
                        p2 = 18446744073709551615
                    case "\n\t":
                        var v8 uint32 = 2147483648
                        _ = v8
                    case "hello":
                        var v9 bool = false
                        _ = v9
                    default:
                        p1 -= (p1 & (p1 & p1))
                    }
                }
            }
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint8 = 2
                v12 := []int{678}
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            v13 := GenS1{F0: 0, F1: 237, F2: 0}
            _ = v13
        }
        _ = fi5
    }
    switch ("a" + p0) {
    case "foo bar":
        var v14 int32 = -46
        _ = v14
    case "":
        var v15 int8 = (^(p1))
        _ = v15
    default:
        v16 := GenS1{F0: 0, F1: 297, F2: 42}
        _ = v16
    }
    var v17 bool = false
    p0 = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 0
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 255
    v3 := []string{"foo bar"}
    var v4 int64 = -1
    fmt.Println(v2, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 int = (fi18 + fi18)
        _ = fi18
        _ = v19
    }
    v2 = 422
    switch 32767 {
    case 2147483647:
        v4 = v4
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    if false {
        v4 = (v4 % ((v4) | 1))
        v4 += v4
        v2 = v2
    } else {
        v3 = append(v3, "a")
        var v20 bool = true
        _ = v20
    }
    v21 := GenS2{F0: 2, F1: 236}
    if (!(32767 != 790)) {
        if false {
            v1 = v1
            if true {
                for fi22 := 0; fi22 < 9; fi22++ {
                    fmt.Println(fi22, v2, v1)
                    _ = fi22
                }
            }
            var v23 uint16 = 827
            _ = v23
        }
        switch -474 {
        case -9223372036854775808:
            var v24 float32 = 0.1
            _ = v24
        case 32768:
            v25 := GenS1{F0: 42, F1: 0, F2: 4294967295}
            _ = v25
        case 24:
            v1 = v1
        }
    }
    v2 -= v2
    switch "\n\t" {
    case "hello":
        v4 += (v4 << ((v2) & 63))
    case "a":
        v26 := GenS2{F0: 2, F1: 2}
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        if (v2 > (v2 << (((v2 | 255)) & 63))) {
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 int32 = 42
                v30 := GenS1{F0: -2147483648, F1: (^(v1)), F2: 42}
                switch "foo bar" {
                case "foo bar":
                    var v31 int64 = genFn0(("\n\t" + (("hello" + "\n\t") + "a")), 100, (^(v2)))
                    _ = v31
                case "\n\t":
                    if false {
                        var v32 uint64 = 2
                        if false {
                            v4 -= (v4 - -1)
                            var v33 int64 = genFn0("foo bar", -73, (v2 | (v2 | 42)))
                            _ = v33
                        } else {
                            v34 := GenS3{F0: 2147483648}
                            if (v1 < v1) {
                                var v35 int16 = -32768
                                v35 = (-(v35))
                                _ = v35
                            }
                            _ = v34
                        }
                        _ = v32
                    } else {
                        v4 = v4
                    }
                default:
                    if (("hello" + "hello") < ("\n\t" + "foo bar")) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v36 uint16 = 0
                        if ("foo bar" <= ((("foo bar" + "") + ("hello" + "foo bar")) + "a")) {
                            var v37 uint8 = 2
                            _ = v37
                        } else {
                            v38 := GenS4{F0: -1}
                            fmt.Println(v2, v4, v29)
                            _ = v38
                        }
                        _ = v36
                    } else {
                        var v39 float64 = 3.14159
                        v39 += 1e10
                        _ = v39
                    }
                }
                _ = fi28
                _ = v29
                _ = v30
            }
            var v40 int64 = 42
            fmt.Println(v4, v2)
            _ = v40
        } else {
            _ = len(v3)
            for fi41 := 0; fi41 < 14; fi41++ {
                v2 -= v2
                var v42 string = "a"
                v43 := GenS5{F0: (("a" + (((v42 + "hello") + ("\n\t" + ((("\n\t" + v42) + v42) + "\n\t"))) + v42)) + ""), F1: int32(fi41), F2: -1}
                var v44 float32 = (float32((-2.25 / (float64(3.14159) * 3.14159))) / float32(3.14159))
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
            }
        }
        _ = fi27
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        v45 := GenS2{F0: 255, F1: 166}
        _ = v45
    }
    if false {
        switch 127 {
        case -42:
            fmt.Println(v4, v2)
        case -2147483648:
            v46 := GenS6{F0: 255, F1: 110, F2: 255}
            _ = v46
        default:
            var v47 uint8 = 2
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
