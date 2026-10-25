package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 int8; F2 uint32 }

func main() {
    v1 := []string{(("a" + ("foo bar" + ("hello" + ""))) + "foo bar"), ("hello" + "")}
    var v2 uint64 = 90
    var v3 int16 = 0
    var v4 int64 = -534
    var v5 uint = 42
    if (!true) {
        for fi6 := 0; fi6 < 13; fi6++ {
            switch (("hello" + "foo bar") + ("hello" + ("\n\t" + (("" + ((("\n\t" + ("" + "hello")) + ("hello" + (("a" + ("\n\t" + "\n\t")) + ("\n\t" + ("hello" + "hello"))))) + "hello")) + "foo bar")))) {
            case "hello":
                switch fi6 {
                case 0:
                    fmt.Println(v3)
                case 805:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 128:
                    v4 += (v4 << (((v5 * (v5 % ((4294967295) | 1)))) & 63))
                }
            default:
                if len(v1) > 0 {
                    v1[0] = (((("" + ("foo bar" + "foo bar")) + "a") + ("hello" + (("hello" + "foo bar") + ("" + (("hello" + "\n\t") + "hello"))))) + ("\n\t" + "hello"))
                }
            }
            _ = fi6
        }
    }
    if true {
        var v7 int = -535
        v8 := GenS1{F0: 434, F1: 127, F2: v3}
        var v9 uint32 = 65535
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 int32 = -1
        if (!((v2 * 503) == v2)) {
            var v11 int8 = 42
            v2 += (v2 % (((v2 + (v2 | 737))) | 1))
            _ = v11
        } else {
            var v12 bool = false
            _ = v12
        }
        _ = v10
    }
    fmt.Println(v4, v5, v2)
    var v13 int32 = -639
    switch 836 {
    case 870:
        if (!false) {
            v4 -= (v4 * (-(v4)))
            if (0.1 == 1.5) {
                var v14 uint64 = 2
                _ = v14
            } else {
                v15 := GenS2{F0: 42}
                _ = v15
            }
            switch (("a" + "foo bar") + "a") {
            case "foo bar":
                v16 := GenS1{F0: uint16(v2), F1: 159, F2: (v3 << ((v5) & 63))}
                _ = v16
            case "a":
                var v17 uint32 = 0
                _ = v17
            }
        }
    case 42:
        v2 += v2
    default:
        v4 -= (v4 * (-(v4)))
    }
    v18 := GenS2{F0: 14}
    var v19 int64 = (v4 + int64(v13))
    var v20 uint8 = 42
    switch int(v5) {
    case -260:
        fmt.Println(v19, v13, v5)
    case -2147483648:
        v2 -= 1
    case -1:
        var v21 int32 = (v13 ^ 718)
        _ = v21
    }
    v22 := GenS1{F0: 2, F1: uint8(v5), F2: v3}
    for fi23 := 0; fi23 < 14; fi23++ {
        if (!(v3 != (v3 | 42))) {
            var v24 int16 = v3
            for fi25 := 0; fi25 < 1; fi25++ {
                switch (^(fi25)) {
                case 127:
                    if false {
                        fmt.Println(v13, v2)
                    } else {
                        v26 := GenS3{F0: (v5 ^ (^(v5))), F1: 1}
                        _ = v26
                    }
                case 128:
                    v27 := GenS1{F0: 113, F1: v20, F2: (-(v3))}
                    _ = v27
                case 32768:
                    if false {
                        v28 := GenS2{F0: 100}
                        v29 := GenS4{F0: 4294967295}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (100 >= -1) {
                            var v30 uint64 = 18446744073709551615
                            v19 += 9223372036854775807
                            _ = v30
                        }
                        fmt.Println(v2)
                        _ = v28
                        _ = v29
                    }
                default:
                    if ((!(true && false)) && (false && (int8(v2) <= -100))) {
                        v2 -= 1
                        if (!(-449 <= fi23)) {
                            fmt.Println(v2)
                        } else {
                            if (977 == 320) {
                                v31 := GenS1{F0: 2, F1: (v20 ^ v20), F2: v3}
                                if true {
                                    var v32 uint64 = (v2 * uint64(v5))
                                    _ = v32
                                }
                                var v33 uint8 = v20
                                fmt.Println(v2, v20, fi25)
                                _ = v31
                                _ = v33
                            } else {
                                var v34 uint64 = v2
                                _ = v34
                            }
                            var v35 uint64 = (v2 & 0)
                            fmt.Println(fi23, v13, v13)
                            _ = v35
                        }
                        var v36 string = (("foo bar" + "a") + "foo bar")
                        v37 := GenS5{F0: (fi25 | (-(fi23)))}
                        _ = v36
                        _ = v37
                    } else {
                        v38 := GenS6{F0: true, F1: 0, F2: 1}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v38
                    }
                }
                switch (^(fi23)) {
                case -1:
                    v1 = append(v1, "hello")
                }
                _ = fi25
            }
            _ = v24
        }
        v3 = v3
        _ = fi23
    }
    v4 -= (v19 - 519)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
