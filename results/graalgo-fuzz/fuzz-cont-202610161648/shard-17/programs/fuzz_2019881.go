package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint8; F1 int64; F2 bool }
func genFn0(p0 int8, p1 int32, p2 uint64) bool {
    var v5 uint32 = 4294967295
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 42, F1: 42, F2: -1}
            if ((v5 * 467) <= v5) {
                if (!true) {
                    v9 := []int{}
                    v10 := GenS1{F0: 32767, F1: 100, F2: (p0 | 0)}
                    _ = v10
                    _ = v9
                }
                switch ((("foo bar" + ("\n\t" + ("a" + "\n\t"))) + "\n\t") + "foo bar") {
                case "a":
                    p2 = 4294967295
                    v11 := GenS1{F0: 0, F1: -87, F2: p0}
                    v12 := []int{}
                    _ = v11
                    _ = v12
                default:
                    var v13 int32 = (^(p1))
                    _ = v13
                }
                var v14 int = -827
                _ = v14
            }
            var v15 int32 = p1
            v16 := []int{}
            _ = fi7
            _ = v15
            _ = v16
            _ = v8
        }
        _ = fi6
    }
    var v17 uint8 = 22
    fmt.Println(p2)
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: 1, F1: p0, F2: -128}
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v5
    return (false && (-368 != (-(p1))))
}
func genFn1(p0 int32) uint32 {
    fmt.Println(p0)
    p0 += 42
    if false {
        var v20 string = "a"
        _ = v20
    } else {
        p0 = p0
        v21 := GenS2{F0: 42, F1: -7, F2: 9223372036854775807}
        _ = v21
    }
    var v22 int8 = 127
    for fi23 := 0; fi23 < 5; fi23++ {
        v22 = int8(fi23)
        switch int(p0) {
        case 32768:
            v22 = (v22 % ((42) | 1))
        default:
            var v24 uint64 = 1
            _ = v24
        }
        switch fi23 {
        case 708:
            for fi25 := 0; fi25 < 17; fi25++ {
                v26 := GenS3{F0: 2147483648}
                v27 := []int{-9223372036854775808, (-(fi25))}
                fmt.Println(p0)
                _ = fi25
                _ = v26
                _ = v27
            }
        case -194:
            p0 -= p0
        }
        _ = fi23
    }
    _ = p0
    _ = v22
    return 4294967295
}

func main() {
    var v1 uint16 = 255
    var v2 int64 = -912
    var v3 int64 = v2
    v4 := []string{"foo bar", "hello", "foo bar", ("\n\t" + ("\n\t" + ("hello" + "")))}
    v1 -= v1
    if (224 != 0) {
        v2 -= v2
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v1, v1)
    case "a":
        var v28 uint32 = genFn1(0)
        _ = v28
    case "":
        fmt.Println(v3, v1, v3)
    default:
        switch 127 {
        case 4294967295:
            if true {
                v2 = (v2 + (-(v2)))
            }
        case 42:
            v2 = v3
        default:
            v2 += 9223372036854775807
        }
    }
    v29 := GenS4{F0: 1, F1: (int16(v1) >= 1)}
    if false {
        v2 += v2
    }
    if true {
        v4 = append(v4, "foo bar")
        var v30 int = 197
        _ = v30
    } else {
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 uint64 = 216
            if false {
                for fi33 := 0; fi33 < 6; fi33++ {
                    v34 := GenS5{F0: 1, F1: v3, F2: ("a" > "foo bar")}
                    v32 = (v32 ^ v32)
                    _ = len(v4)
                    _ = fi33
                    _ = v34
                }
                v35 := GenS5{F0: 127, F1: 715, F2: (0.0 != 1e10)}
                var v36 uint16 = (v1 & v1)
                _ = v35
                _ = v36
            }
            _ = fi31
            _ = v32
        }
    }
    var v37 uint8 = 2
    for fi38 := 0; fi38 < 8; fi38++ {
        v37 -= v37
        _ = fi38
    }
    v39 := GenS4{F0: 0, F1: false}
    v40 := GenS3{F0: 2147483648}
    v2 = (v3 | (v2 & v2))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        if (264 <= -32768) {
            switch (("\n\t" + "\n\t") + "hello") {
            case "a":
                v3 += v3
            case "\n\t":
                var v41 uint64 = 4294967295
                _ = v41
            default:
                if (875 > 1) {
                    _ = len(v4)
                    fmt.Println(v3)
                }
            }
        } else {
            for fi42 := 0; fi42 < 4; fi42++ {
                var v43 int = (fi42 | int(v3))
                _ = fi42
                _ = v43
            }
        }
        v44 := GenS4{F0: 127, F1: true}
        _ = v44
    } else {
        switch "a" {
        case "a":
            if true {
                fmt.Println(v37, v1, v1)
            }
        case "foo bar":
            if len(v4) > 0 {
                v4[0] = ""
            }
        }
        fmt.Println(v37, v2)
        switch ("" + ("" + "\n\t")) {
        case "foo bar":
            fmt.Println(v1, v37)
        case "\n\t":
            for fi45 := 0; fi45 < 15; fi45++ {
                fmt.Println(v1, v37, fi45)
                _ = fi45
            }
        default:
            if ((("a" + "") + "foo bar") >= "a") {
                fmt.Println(v2, v3, v2)
                _ = len(v4)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v40
}
