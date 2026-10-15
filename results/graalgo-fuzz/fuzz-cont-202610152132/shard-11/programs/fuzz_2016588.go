package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint8; F2 uint }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint64; F1 int; F2 int16 }
type GenS5 struct { F0 string; F1 uint; F2 uint32 }
func genFn0(p0 int32) string {
    var v6 int32 = p0
    p0 -= 605
    fmt.Println(p0, v6)
    v7 := []int{42, 4294967295, 32768}
    if true {
        v6 -= (^(p0))
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    v1 := []uint32{1, 1, 466}
    var v2 int8 = 100
    v3 := make(map[uint32]string)
    var v4 float64 = 1.5
    var v5 uint = 42
    switch 9223372036854775807 {
    case 128:
        v5 = 42
    }
    if false {
        delete(v3, 65535)
        v9 := v3[42]
        _ = v9
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    if false {
        for fi10 := 0; fi10 < 13; fi10++ {
            if false {
                if (0 < 4294967295) {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        var v12 uint64 = 4294967295
                        _ = fi11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 7; fi13++ {
                        v14 := GenS1{F0: "a"}
                        _ = fi13
                        _ = v14
                    }
                    if (!(1.5 != (v4 + v4))) {
                        var v15 bool = (!false)
                        _ = v15
                    } else {
                        fmt.Println(v5)
                        switch (-(fi10)) {
                        case 42:
                            fmt.Println(fi10, fi10, fi10)
                        case 2147483647:
                            v16 := GenS2{F0: ("a" + "hello"), F1: uint8(v5), F2: (v5 & (v5 >> ((v5) & 63)))}
                            _ = v16
                        }
                        var v17 int32 = 1
                        _ = v17
                    }
                }
            } else {
                if true {
                    v18 := GenS1{F0: "a"}
                    _ = v18
                }
                v19 := v3[1]
                if false {
                    switch (v19 + ("foo bar" + (("\n\t" + v19) + ""))) {
                    case "\n\t":
                        v5 = (v5 | (v5 | (v5 >> (((v5 ^ v5)) & 63))))
                    case "":
                        if false {
                            var v20 bool = (65535 >= 0)
                            v20 = (((v4 < 1.5) && v20) && (v20 || (!(!(833 < 2147483647)))))
                            _ = v20
                        } else {
                            if (true && false) {
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    switch fi10 {
                                    case 637:
                                        v22 := GenS3{F0: -1, F1: v2}
                                        _ = v22
                                    case 706:
                                        if (-1 == 597) {
                                            var v23 bool = (1 != 1)
                                            _ = v23
                                        }
                                    case 32768:
                                        fmt.Println(v19, fi21, fi10)
                                    default:
                                        switch "a" {
                                        case "":
                                            var v24 int = fi10
                                            _ = v24
                                        case "foo bar":
                                            v19 = "foo bar"
                                        default:
                                            v4 -= -2.25
                                        }
                                    }
                                    if false {
                                        var v25 uint64 = 274
                                        fmt.Println(v25)
                                        _ = v25
                                    } else {
                                        var v26 float64 = v4
                                        var v27 float32 = 1.5
                                        switch (fi10 * -293) {
                                        case -471:
                                            v4 += 0.1
                                        }
                                        _ = v26
                                        _ = v27
                                    }
                                    var v28 int16 = 42
                                    _ = fi21
                                    _ = v28
                                }
                            }
                        }
                    default:
                        fmt.Println(v5, fi10)
                    }
                    var v29 float64 = 0.0
                    v4 = v4
                    _ = v29
                }
                _ = v19
            }
            v2 = (v2 ^ (v2 - -110))
            v30 := GenS4{F0: 2, F1: (fi10 >> ((1) & 63)), F2: -218}
            _ = fi10
            _ = v30
        }
        v31 := GenS5{F0: "", F1: v5, F2: 0}
        _ = v31
    }
    v32 := GenS5{F0: (("hello" + "foo bar") + "foo bar"), F1: v5, F2: 112}
    var v33 float32 = 3.14159
    fmt.Println(v2, v4, v2)
    var v34 int16 = 1000
    for fi35 := 0; fi35 < 13; fi35++ {
        switch -834 {
        case 127:
            v36 := GenS5{F0: (("\n\t" + ((("\n\t" + (("\n\t" + "") + ((("foo bar" + "hello") + "\n\t") + "foo bar"))) + "hello") + "foo bar")) + "foo bar"), F1: v5, F2: 2147483648}
            _ = v36
        case -582:
            for fi37 := 0; fi37 < 16; fi37++ {
                var v38 uint32 = 2147483648
                var v39 int16 = (v34 & (v34 | (^(v34))))
                var v40 uint32 = 4294967295
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
        default:
            for fi41 := 0; fi41 < 7; fi41++ {
                v33 += 3.14159
                _ = fi41
            }
        }
        _ = fi35
    }
    v33 = v33
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
