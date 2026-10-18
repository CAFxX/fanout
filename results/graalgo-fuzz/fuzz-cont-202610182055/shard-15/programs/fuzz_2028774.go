package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int16; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0(p0 string) uint8 {
    p0 = p0
    var v5 string = "a"
    v6 := GenS1{F0: -2147483648, F1: 0, F2: 32767}
    v7 := []int{-1}
    if (!false) {
        var v8 bool = (!false)
        for fi9 := 0; fi9 < 18; fi9++ {
            v5 = v5
            var v10 bool = (("" < v5) && (1 > 1000))
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint8 {
    if true {
        v11 := GenS1{F0: 0, F1: 0, F2: -864}
        _ = v11
    }
    var v12 bool = true
    var v13 uint16 = 384
    var v14 string = "foo bar"
    _ = v12
    _ = v13
    _ = v14
    return genFn0(((v14 + (v14 + v14)) + v14))
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]uint32)
    var v3 int16 = -32768
    v4 := []int{493, 0, -683, -2147483648}
    if len(v4) > 0 {
        v4[0] = 32767
    }
    v4 = append(v4, 4294967295)
    delete(v2, 32767)
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint8 = 127
        fmt.Println(v16, v1, v16)
        var v17 int32 = -744
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS2{F0: -903, F1: (fi18 >> ((0) & 63))}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 bool = ((!true) || (76 == genFn1()))
        _ = fi20
        _ = v21
    }
    if true {
        var v22 uint64 = 4294967295
        _ = v22
    } else {
        if false {
            switch "\n\t" {
            case "":
                if (!(!true)) {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        fmt.Println(fi23)
                        _ = fi23
                    }
                    var v24 int = -576
                    v1 = v1
                    _ = v24
                }
            case "hello":
                v3 = 0
            case "\n\t":
                switch 314 {
                case 32767:
                    switch 0 {
                    case 9223372036854775807:
                        switch -767 {
                        case 127:
                            v25 := GenS2{F0: -1, F1: -1}
                            _ = v25
                        case -1:
                            var v26 float32 = 0.1
                            _ = v26
                        default:
                            v3 = 0
                        }
                    default:
                        v3 -= (v3 << ((223) & 63))
                    }
                case 520:
                    if (true && (1 >= 1)) {
                        v1 = v1
                        v27 := GenS2{F0: -986, F1: -326}
                        _ = v27
                    }
                case -9223372036854775808:
                    fmt.Println(v1, v3)
                }
            }
            v28 := GenS1{F0: -2147483648, F1: -524, F2: 978}
            switch (("\n\t" + ((((((("\n\t" + "foo bar") + (("foo bar" + "foo bar") + "\n\t")) + ("hello" + "")) + "\n\t") + "a") + "hello") + "foo bar")) + "") {
            case "\n\t":
                fmt.Println(v3, v3)
            case "":
                var v29 uint = 255
                _ = v29
            case "hello":
                v30 := GenS3{F0: -1, F1: -369, F2: (v1 * v1)}
                _ = v30
            }
            _ = v28
        }
        if len(v4) > 0 {
            v4[0] = 128
        }
        v3 = v3
    }
    v31 := GenS1{F0: 127, F1: 0, F2: (^(v3))}
    if (!true) {
        v32 := v2[32768]
        v33 := GenS2{F0: 42, F1: 0}
        var v34 string = "a"
        _ = v32
        _ = v33
        _ = v34
    } else {
        switch (((("" + "hello") + ("" + "hello")) + "\n\t") + "a") {
        case "":
            var v35 int = -165
            _ = v35
        case "a":
            v3 = v3
        case "hello":
            v36 := GenS4{F0: 2, F1: -9223372036854775808, F2: 794}
            _ = v36
        default:
            if false {
                if len(v4) > 0 {
                    v4[0] = 1
                }
                var v37 int64 = int64(v3)
                switch (("hello" + "") + "foo bar") {
                case "":
                    var v38 uint32 = 196
                    _ = v38
                }
                _ = v37
            } else {
                v1 = (v1 & (v1 * 0))
                if len(v4) > 0 {
                    v4[0] = -767
                }
                var v39 string = ""
                _ = v39
            }
        }
        if (true && (false && true)) {
            fmt.Println(v3, v3)
        } else {
            if (!true) {
                v40 := GenS4{F0: 0, F1: -1, F2: 4294967295}
                v41 := v2[-151]
                v42 := v2[int(v1)]
                _ = v40
                _ = v41
                _ = v42
            }
            var v43 int16 = (v3 * (v3 ^ 32767))
            _ = v43
        }
        for fi44 := 0; fi44 < 7; fi44++ {
            var v45 bool = true
            fmt.Println(v3, v45, v45)
            _ = fi44
            _ = v45
        }
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
