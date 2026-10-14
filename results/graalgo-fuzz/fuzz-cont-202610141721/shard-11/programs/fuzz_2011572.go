package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 string; F1 int16; F2 int64 }
func genFn0(p0 uint, p1 bool) uint32 {
    fmt.Println(p1, p0)
    p0 += (^(p0))
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]string)
    v3 := []uint32{0, 640}
    v1 = v1
    v4 := GenS1{F0: 980, F1: 272}
    switch ("a" + "foo bar") {
    case "foo bar":
        var v5 int16 = 1
        _ = v5
    case "hello":
        switch (("foo bar" + ((("a" + ("hello" + "")) + ((((((("foo bar" + (((("a" + ("hello" + ("hello" + "foo bar"))) + "\n\t") + "") + "hello")) + ("hello" + "\n\t")) + "\n\t") + "") + "hello") + "foo bar") + "hello")) + "")) + "hello") {
        case "a":
            v1 = (v1 % ((v1) | 1))
        case "foo bar":
            var v6 int32 = 0
            _ = v6
        case "hello":
            for fi7 := 0; fi7 < 8; fi7++ {
                v1 += v1
                _ = fi7
            }
        default:
            if false {
                if true {
                    if (!((((float64((1e10 - 1.5)) + 3.14159) + 0.0) <= (((1.5 / 0.1) / -2.25) * 0.0)) && ((true || (int16(v1) != 32767)) || (!false)))) {
                        v8 := GenS1{F0: genFn0(0, true), F1: 255}
                        _ = v8
                    }
                }
                _ = len(v2)
                v9 := GenS1{F0: 181, F1: 42}
                _ = v9
            } else {
                if false {
                    if (!(!true)) {
                        var v10 int16 = 424
                        if true {
                            v11 := GenS2{F0: 56}
                            v12 := GenS2{F0: (v1 >> ((2) & 63))}
                            var v13 int16 = v10
                            _ = v11
                            _ = v12
                            _ = v13
                        } else {
                            var v14 int = 32767
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v14
                        }
                        var v15 uint8 = 2
                        _ = v10
                        _ = v15
                    } else {
                        v1 = 0
                    }
                    fmt.Println(v1)
                }
                for fi16 := 0; fi16 < 5; fi16++ {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi17
                    }
                    v1 = (v1 * v1)
                    if (!true) {
                        switch ("hello" + "") {
                        case "hello":
                            v1 = (v1 % ((v1) | 1))
                        }
                        v1 -= v1
                    }
                    _ = fi16
                }
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(fi18)
                    if false {
                        fmt.Println(v1, v1)
                    }
                    _ = fi18
                }
            }
        }
    default:
        var v19 bool = true
        _ = v19
    }
    var v20 uint16 = 42
    v20 = (v20 * 839)
    fmt.Println(v20, v1)
    fmt.Println(v1, v20)
    delete(v2, ("hello" + (("hello" + (("hello" + ("foo bar" + "a")) + "a")) + "hello")))
    var v21 float32 = (3.14159 / 1.0)
    v20 = (v20 * 65535)
    v22 := v2["hello"]
    if true {
        if true {
            _ = len(v2)
            v1 -= int8(v20)
            for fi23 := 0; fi23 < 13; fi23++ {
                for fi24 := 0; fi24 < 18; fi24++ {
                    switch (fi24 - (fi23 * 1)) {
                    case 42:
                        var v25 int8 = (v1 * v1)
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        v26 := GenS1{F0: 65535, F1: 1}
                        _ = v25
                        _ = v26
                    case 128:
                        var v27 uint8 = 154
                        _ = v27
                    case -464:
                        var v28 int64 = 146
                        _ = v28
                    default:
                        var v29 string = (v22 + "a")
                        _ = v29
                    }
                    if false {
                        var v30 int16 = 0
                        v3 = append(v3, 0)
                        if len(v3) > 0 {
                            v3[0] = 415
                        }
                        _ = len(v2)
                        var v31 string = "\n\t"
                        _ = v30
                        _ = v31
                    }
                    _ = fi24
                }
                for fi32 := 0; fi32 < 11; fi32++ {
                    v33 := GenS2{F0: (v1 / ((33) | 1))}
                    fmt.Println(v20, fi23, fi23)
                    fmt.Println(v21, v1)
                    v21 = -2.25
                    _ = fi32
                    _ = v33
                }
                v34 := GenS1{F0: genFn0(4294967295, true), F1: 1}
                _ = fi23
                _ = v34
            }
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        v35 := GenS3{F0: (-(v1)), F1: (!false)}
        _ = v35
    }
    v36 := GenS3{F0: (v1 | (v1 - 127)), F1: false}
    var v37 int32 = 0
    for fi38 := 0; fi38 < 19; fi38++ {
        v39 := GenS4{F0: v22, F1: 0, F2: 0}
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        switch (fi40 % ((fi40) | 1)) {
        case 358:
            if false {
                v20 = v20
                v21 = v21
            } else {
                v22 = (v22 + v22)
            }
        case 4294967295:
            for fi41 := 0; fi41 < 3; fi41++ {
                fmt.Println(v37, fi41)
                _ = fi41
            }
        }
        var v42 int = (^(fi40))
        _ = fi40
        _ = v42
    }
    v37 -= v37
    var v43 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v43
}
