package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int8; F1 string; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 bool }
func genFn0() uint64 {
    if false {
        v7 := GenS1{F0: 42, F1: 245, F2: 2}
        _ = v7
    }
    var v8 uint32 = 42
    _ = v8
    return 42
}
func genFn1(p0 uint16, p1 uint, p2 uint64) uint16 {
    var v9 string = "hello"
    v9 = v9
    fmt.Println(p1, v9, p1)
    var v10 int8 = 20
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (p0 ^ (^(p0)))
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 1
    var v3 int32 = 1
    v4 := []string{(("a" + ("" + "foo bar")) + "hello"), "\n\t", "", ("a" + "")}
    var v5 float32 = (1.5 / 1.0)
    v6 := make(map[string]uint32)
    switch -1 {
    case -1:
        if (false && true) {
            if ((float64((-2.25 + (0.1 * 0.0))) / 1.5) < 0.1) {
                var v11 int = -419
                _ = v11
            } else {
                v12 := GenS2{F0: 1000, F1: 2}
                _ = v12
            }
        }
    case -370:
        v3 = v3
    case 98:
        for fi13 := 0; fi13 < 5; fi13++ {
            switch "a" {
            case "hello":
                var v14 uint8 = uint8(v2)
                _ = v14
            case "foo bar":
                for fi15 := 0; fi15 < 14; fi15++ {
                    v16 := GenS3{F0: 109, F1: 135}
                    _ = fi15
                    _ = v16
                }
            }
            v5 = ((0.1 + ((0.1 / 1.0) - v5)) + v5)
            _ = fi13
        }
    }
    if false {
        var v17 bool = (89 > 171)
        v18 := GenS4{F0: -128, F1: "foo bar", F2: 355}
        _ = v17
        _ = v18
    }
    var v19 float64 = 3.14159
    var v20 uint32 = 65535
    if (!true) {
        _ = len(v4)
        v3 -= (^(v3))
        v2 += 42
    } else {
        for fi21 := 0; fi21 < 12; fi21++ {
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS4{F0: 127, F1: (("hello" + ("foo bar" + ("" + ("hello" + "\n\t")))) + (("foo bar" + "") + "foo bar")), F2: 170}
                v2 = v2
                _ = fi22
                _ = v23
            }
            var v24 float64 = float64(((1e10 * v5) / ((v5 / float32(v19)) * (3.14159 - v5))))
            for fi25 := 0; fi25 < 13; fi25++ {
                v24 -= (1.5 / v19)
                switch (fi25 / ((fi25) | 1)) {
                case 32768:
                    fmt.Println(v24, v5, v24)
                case 9223372036854775807:
                    v24 = v24
                default:
                    var v26 uint64 = 4294967295
                    fmt.Println(v19, v20, v5)
                    if (829 >= int64(v2)) {
                        v27 := GenS5{F0: uint8(v20), F1: fi25}
                        _ = v27
                    } else {
                        var v28 int8 = -128
                        v29 := GenS6{F0: 42}
                        v26 -= (v26 / (((v26 / ((genFn0()) | 1))) | 1))
                        _ = v28
                        _ = v29
                    }
                    _ = v26
                }
                v30 := GenS7{F0: 42, F1: ""}
                if true {
                    var v31 bool = false
                    var v32 uint32 = (v20 | 2)
                    _ = v31
                    _ = v32
                } else {
                    var v33 uint64 = 0
                    v34 := v1[127]
                    _ = v33
                    _ = v34
                }
                if (!true) {
                    var v35 string = (("foo bar" + ("\n\t" + ("\n\t" + "hello"))) + "hello")
                    fmt.Println(v2, v35)
                    delete(v6, ("foo bar" + v35))
                    v36 := GenS8{F0: (-652 != 32767)}
                    v5 -= ((v5 - 1.5) - (float32(v24) / 3.14159))
                    var v37 uint8 = 188
                    var v38 uint32 = v20
                    v24 = 3.14159
                    fmt.Println(v5)
                    _ = v35
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = fi25
                _ = v30
            }
            _ = fi21
            _ = v24
        }
    }
    v2 -= 42
    var v39 float64 = v19
    if false {
        v4 = append(v4, "foo bar")
    } else {
        v39 -= v19
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 bool = true
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        switch (("" + "") + "") {
        case "\n\t":
            var v43 int32 = (v3 >> ((0) & 63))
            _ = v43
        case "hello":
            v44 := GenS3{F0: 128, F1: 2}
            _ = v44
        case "":
            var v45 int8 = -128
            _ = v45
        }
        switch 32767 {
        case 32768:
            if true {
                if true {
                    v19 = v19
                }
            } else {
                for fi46 := 0; fi46 < 17; fi46++ {
                    var v47 int8 = -128
                    _ = fi46
                    _ = v47
                }
            }
        default:
            fmt.Println(v19, v3, v5)
        }
        _ = fi42
    }
    fmt.Println(v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
