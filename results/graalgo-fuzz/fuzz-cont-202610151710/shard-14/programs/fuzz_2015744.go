package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint; F1 string; F2 int }
type GenS6 struct { F0 string; F1 int; F2 int8 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 string; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -1
    var v3 uint64 = 582
    v4 := []string{"a"}
    v5 := make(map[string]uint32)
    var v6 uint64 = (v3 % ((v3) | 1))
    v7 := GenS1{F0: (("foo bar" + ("" + "foo bar")) + ("hello" + ("hello" + ""))), F1: (2 >= 255), F2: 65535}
    v8 := GenS1{F0: "hello", F1: false, F2: 627}
    var v9 uint64 = (v3 >> ((18446744073709551615) & 63))
    switch -1 {
    case -364:
        v10 := GenS1{F0: "hello", F1: true, F2: 957}
        _ = v10
    case 9223372036854775807:
        if len(v4) > 0 {
            v4[0] = ("a" + "hello")
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch "foo bar" {
    case "hello":
        for fi11 := 0; fi11 < 15; fi11++ {
            v2 -= v2
            v12 := v1[fi11]
            v13 := GenS1{F0: "", F1: true, F2: 1}
            _ = fi11
            _ = v12
            _ = v13
        }
    case "":
        delete(v1, -2147483648)
    case "\n\t":
        v4 = append(v4, "\n\t")
    default:
        if true {
            v4 = append(v4, "a")
            v14 := GenS2{F0: -128}
            _ = v14
        }
    }
    switch ("" + "") {
    case "":
        switch -9223372036854775808 {
        case -2147483648:
            if (-236 >= 0) {
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 uint = 1
                    switch "a" {
                    case "\n\t":
                        if true {
                            var v17 int8 = 127
                            _ = v17
                        }
                    case "":
                        for fi18 := 0; fi18 < 16; fi18++ {
                            v19 := GenS3{F0: 2147483647, F1: 507}
                            v6 = (v9 - v3)
                            _ = fi18
                            _ = v19
                        }
                    case "hello":
                        var v20 int = -129
                        _ = v20
                    }
                    v2 += (v2 + (-(v2)))
                    _ = fi15
                    _ = v16
                }
            }
        case 856:
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 string = "foo bar"
                v22 = v22
                _ = fi21
                _ = v22
            }
        case 0:
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := GenS4{F0: 2, F1: 2, F2: (v9 ^ 18446744073709551615)}
                _ = fi23
                _ = v24
            }
        default:
            if true {
                var v25 uint = 18446744073709551615
                v3 = v6
                _ = v25
            } else {
                if false {
                    for fi26 := 0; fi26 < 7; fi26++ {
                        v27 := GenS3{F0: 838, F1: 1}
                        var v28 int16 = -194
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                } else {
                    v29 := GenS5{F0: 42, F1: "hello", F2: 0}
                    if false {
                        switch -882 {
                        case -2147483648:
                            var v30 string = ""
                            _ = v30
                        default:
                            fmt.Println(v3)
                        }
                        fmt.Println(v2, v9)
                        var v31 uint8 = 0
                        _ = v31
                    }
                    var v32 float64 = (-2.25 / (-2.25 - 1.5))
                    _ = v29
                    _ = v32
                }
                if (!((!true) || true)) {
                    v33 := GenS6{F0: ((("hello" + "\n\t") + "foo bar") + "\n\t"), F1: 1, F2: 0}
                    v9 -= (v6 * v3)
                    v34 := GenS3{F0: -358, F1: 947}
                    _ = v33
                    _ = v34
                }
                v2 = (v2 << ((uint(v9)) & 63))
            }
        }
    }
    v9 = 412
    switch 1 {
    case -2147483648:
        switch -9223372036854775808 {
        case 345:
            if true {
                v1[0] = 1
            }
        case -933:
            var v35 string = ""
            _ = v35
        case -2147483648:
            switch ("foo bar" + "") {
            case "foo bar":
                v9 -= (v3 - 2)
            case "hello":
                v2 -= v2
            default:
                _ = len(v1)
            }
        }
    case 9223372036854775807:
        switch 9223372036854775807 {
        case -9223372036854775808:
            fmt.Println(v2)
        case 4294967295:
            v4 = append(v4, (("foo bar" + ("foo bar" + "hello")) + "foo bar"))
        case -63:
            v36 := GenS7{F0: -962, F1: -176}
            _ = v36
        default:
            if (2147483648 > 915) {
                fmt.Println(v6, v2)
                var v37 string = ((("\n\t" + ("" + "hello")) + "") + "hello")
                fmt.Println(v9, v37)
                _ = v37
            }
        }
    }
    switch ((((("" + "") + "foo bar") + (((("\n\t" + "\n\t") + ("" + "")) + "") + ((("hello" + "") + ("hello" + "hello")) + ("" + "\n\t")))) + "\n\t") + ("a" + ("a" + ("a" + "a")))) {
    case "hello":
        switch -10 {
        case 1:
            v4 = append(v4, "hello")
        case -9223372036854775808:
            _ = len(v4)
        case 32767:
            var v38 uint32 = 65535
            _ = v38
        }
    }
    if (128 == -781) {
        if (0 == -119) {
            v39 := GenS4{F0: 42, F1: 0, F2: (v3 >> ((806) & 63))}
            _ = v39
        }
    } else {
        var v40 uint8 = 127
        v41 := GenS7{F0: -373, F1: 9223372036854775807}
        _ = v40
        _ = v41
    }
    v42 := GenS8{F0: ("" + "hello"), F1: 65535}
    var v43 float64 = float64(1e10)
    var v44 int32 = -479
    _ = len(v4)
    for fi45 := 0; fi45 < 13; fi45++ {
        v46 := GenS6{F0: "hello", F1: fi45, F2: 93}
        var v47 uint64 = (v6 ^ 18446744073709551615)
        _ = fi45
        _ = v46
        _ = v47
    }
    v43 += 1e10
    if ((false || (-2.25 == 0.0)) && (!true)) {
        v2 = v2
        var v48 int8 = -128
        _ = v48
    } else {
        v1[4294967295] = 391
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
