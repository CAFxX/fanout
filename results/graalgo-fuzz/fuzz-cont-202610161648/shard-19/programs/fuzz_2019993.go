package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int32 }
func genFn0(p0 int16, p1 int8, p2 bool) string {
    if p2 {
        fmt.Println(p0, p1)
        fmt.Println(p0, p1)
        var v4 uint32 = 2147483648
        _ = v4
    }
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 float64 = float64(3.14159)
        var v7 string = ""
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return ((("hello" + "\n\t") + "a") + ("foo bar" + ("\n\t" + "")))
}
func genFn1() uint32 {
    v8 := GenS1{F0: 766}
    var v9 int32 = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(fi10, v9)
        _ = fi10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 float64 = 3.14159
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(fi13, fi13)
        _ = fi13
    }
    _ = v8
    _ = v9
    return 504
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    v3 := []uint32{2, 4294967295}
    if true {
        _ = len(v1)
    } else {
        delete(v2, "a")
        var v14 bool = false
        _ = v14
    }
    if (false && false) {
        var v15 int8 = 1
        _ = v15
    }
    if (!false) {
        for fi16 := 0; fi16 < 13; fi16++ {
            _ = len(v2)
            var v17 string = (("\n\t" + "a") + (("hello" + ((((("\n\t" + ("\n\t" + "\n\t")) + ("foo bar" + "hello")) + ("" + "a")) + "") + "a")) + ("\n\t" + (("a" + "\n\t") + ""))))
            delete(v2, "\n\t")
            _ = fi16
            _ = v17
        }
    }
    if (1 > 4294967295) {
        var v18 float64 = (1e10 * 1.5)
        v18 = (((1e10 + (0.0 + (v18 - -2.25))) + (0.0 / ((float64(1.5) - float64(1e10)) - 3.14159))) - v18)
        _ = v18
    } else {
        switch ("" + "") {
        case "\n\t":
            switch "" {
            case "a":
                var v19 uint8 = 2
                _ = v19
            case "\n\t":
                for fi20 := 0; fi20 < 5; fi20++ {
                    switch fi20 {
                    case -1:
                        for fi21 := 0; fi21 < 15; fi21++ {
                            fmt.Println(fi21)
                            _ = fi21
                        }
                    case 4294967295:
                        var v22 float64 = (float64(0.1) + (-2.25 - 0.0))
                        _ = v22
                    case 127:
                        for fi23 := 0; fi23 < 3; fi23++ {
                            var v24 uint32 = 0
                            var v25 uint = 2
                            if ("foo bar" >= ("hello" + ("\n\t" + "\n\t"))) {
                                if false {
                                    var v26 int8 = 75
                                    v27 := GenS1{F0: 893}
                                    v25 = uint(v24)
                                    _ = v26
                                    _ = v27
                                } else {
                                    fmt.Println(v25, fi23, fi23)
                                }
                                if (!false) {
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = v24
                                        }
                                    }
                                    var v28 uint32 = v24
                                    if false {
                                        delete(v2, "foo bar")
                                        var v29 uint16 = 255
                                        _ = v29
                                    } else {
                                        switch -991 {
                                        case 2147483647:
                                            v30 := GenS1{F0: 258}
                                            _ = v30
                                        case 32767:
                                            v2["\n\t"] = false
                                        default:
                                            var v31 uint64 = 2
                                            _ = v31
                                        }
                                        var v32 uint64 = 2
                                        v33 := GenS1{F0: 0}
                                        v3 = append(v3, v28)
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v28
                                }
                                v24 -= v24
                                _ = len(v1)
                            }
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = fi20
                }
            default:
                var v34 uint8 = 150
                _ = v34
            }
        }
        var v35 int16 = 0
        _ = v35
    }
    if (-616 != -262) {
        delete(v2, "foo bar")
        var v36 float64 = (1.5 - 0.0)
        _ = v36
    }
    var v37 uint64 = 4294967295
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v38 int32 = 1
    for fi39 := 0; fi39 < 14; fi39++ {
        if (!false) {
            if (!(((((float64(1e10) / (1.5 * ((0.0 - ((((3.14159 / 3.14159) - 1e10) + float64(1e10)) + 3.14159)) * 0.1))) * (3.14159 / float64(((1.5 / (float32(((0.0 + (0.0 + 1e10)) - 0.0)) / (float32(0.0) + (1.5 + 0.1)))) * ((float32((0.0 * 1.5)) - 1.5) - 0.0))))) / 3.14159) + 1.5) == (-2.25 + ((-2.25 + ((0.0 + ((0.1 * -2.25) * 0.0)) / 1.0)) + 1.5)))) {
                switch ((("foo bar" + "\n\t") + "foo bar") + (("\n\t" + "") + ("\n\t" + "\n\t"))) {
                case "\n\t":
                    var v40 uint8 = 51
                    _ = v40
                }
                v41 := GenS2{F0: 0, F1: v38}
                v37 += v37
                _ = v41
            }
            switch 300 {
            case 612:
                var v42 uint = 844
                _ = v42
            case 9223372036854775807:
                fmt.Println(v37, fi39)
            case 128:
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            }
        }
        _ = len(v2)
        _ = fi39
    }
    for fi43 := 0; fi43 < 10; fi43++ {
        if (255 >= 649) {
            var v44 uint = 18446744073709551615
            var v45 int32 = -651
            _ = len(v1)
            _ = v44
            _ = v45
        }
        var v46 int16 = 32767
        var v47 int16 = (v46 >> ((0) & 63))
        _ = fi43
        _ = v46
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
}
