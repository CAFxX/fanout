package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 uint16, p1 int64) uint {
    var v4 string = ("\n\t" + "a")
    if ((1.5 + -2.25) > (1.5 * 1.5)) {
        v5 := GenS1{F0: (65535 != 2), F1: -127}
        _ = v5
    }
    var v6 int16 = 0
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 724
}
func genFn1(p0 int8, p1 string, p2 string) string {
    p1 = p1
    fmt.Println(p1)
    switch "foo bar" {
    case "a":
        p1 = (("a" + ("a" + p2)) + (("foo bar" + p2) + p2))
    default:
        var v7 int16 = 0
        _ = v7
    }
    var v8 string = (("foo bar" + (p1 + "foo bar")) + (p1 + ((p1 + p2) + ((p2 + "hello") + "hello"))))
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ""
}

func main() {
    var v1 uint8 = 255
    var v2 bool = true
    var v3 uint = 4294967295
    var v9 float32 = 1.5
    for fi10 := 0; fi10 < 4; fi10++ {
        if v2 {
            switch 127 {
            case 563:
                v3 = v3
            case 127:
                fmt.Println(v2, v1, v9)
            case 128:
                var v11 uint16 = 0
                _ = v11
            }
        } else {
            var v12 uint64 = 601
            _ = v12
        }
        _ = fi10
    }
    if (0 > 161) {
        v3 += v3
    }
    if (v2 || (v2 || v2)) {
        if (157 >= 1) {
            fmt.Println(v2, v1)
            v3 = v3
        } else {
            v13 := GenS1{F0: (0 < -1), F1: -1}
            v14 := GenS1{F0: v2, F1: 1}
            _ = v13
            _ = v14
        }
        var v15 int64 = -918
        _ = v15
    }
    v16 := []int{1, -1}
    var v17 uint8 = 86
    switch "a" {
    case "a":
        if len(v16) > 0 {
            v16[0] = -852
        }
    case "hello":
        if v2 {
            var v18 uint8 = (^(v1))
            _ = v18
        }
    case "\n\t":
        if (1 == 1) {
            var v19 float32 = v9
            _ = v19
        } else {
            var v20 uint32 = 0
            _ = v20
        }
    }
    v21 := GenS1{F0: (!false), F1: 51}
    _ = len(v16)
    for fi22 := 0; fi22 < 16; fi22++ {
        for fi23 := 0; fi23 < 12; fi23++ {
            if v2 {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                switch fi23 {
                case 0:
                    v16 = append(v16, (-(fi23)))
                default:
                    var v24 float64 = (1.5 - (3.14159 * float64(1.5)))
                    _ = v24
                }
                var v25 string = ("foo bar" + ("" + ("" + "")))
                _ = v25
            }
            _ = fi23
        }
        v3 += genFn0(496, 0)
        _ = fi22
    }
    switch "" {
    case "\n\t":
        var v26 string = ("\n\t" + "a")
        _ = v26
    case "":
        var v27 int32 = -79
        _ = v27
    case "a":
        if (!v2) {
            for fi28 := 0; fi28 < 8; fi28++ {
                fmt.Println(v17, fi28)
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi28
            }
        } else {
            var v29 uint8 = (v17 | (^(v1)))
            if (-2.25 != (-2.25 * (0.1 + 1.5))) {
                for fi30 := 0; fi30 < 15; fi30++ {
                    var v31 uint32 = 439
                    v32 := GenS2{F0: 0, F1: (("foo bar" + "hello") + ("" + "foo bar"))}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                if v2 {
                    if len(v16) > 0 {
                        v16[0] = 1
                    }
                } else {
                    for fi33 := 0; fi33 < 15; fi33++ {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        _ = fi33
                    }
                    v16 = append(v16, -641)
                    switch -2147483648 {
                    case -2147483648:
                        var v34 float32 = v9
                        _ = v34
                    case 32768:
                        var v35 string = ""
                        _ = v35
                    default:
                        var v36 uint8 = v29
                        _ = v36
                    }
                }
                var v37 int = 1
                _ = v37
            } else {
                var v38 int32 = -2147483648
                v39 := GenS2{F0: 42, F1: ""}
                for fi40 := 0; fi40 < 7; fi40++ {
                    v2 = (320 <= 4294967295)
                    v41 := GenS1{F0: (!false), F1: -116}
                    _ = fi40
                    _ = v41
                }
                _ = v38
                _ = v39
            }
            var v42 uint16 = uint16(v3)
            _ = v29
            _ = v42
        }
    }
    switch 409 {
    case -665:
        if len(v16) > 0 {
            v16[0] = 2147483647
        }
    case 32767:
        var v43 uint64 = 697
        _ = v43
    }
    v17 += (v17 | v17)
    v44 := GenS1{F0: (!(!((((v9 * (-2.25 + (v9 / 1.0))) - v9) - 0.0) < v9))), F1: -128}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v44
    _ = v9
}
