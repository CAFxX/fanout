package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }

func main() {
    var v1 uint32 = 2
    v2 := []string{"a"}
    v3 := make(map[int]string)
    var v4 int = -1
    var v5 string = ""
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int16 = -32768
        v7 = (-(v7))
        switch (v5 + ((v5 + "foo bar") + ("foo bar" + ("" + ((((v5 + v5) + (("a" + "hello") + (v5 + (((("a" + ("foo bar" + (v5 + "\n\t"))) + v5) + v5) + "")))) + v5) + "foo bar"))))) {
        case "":
            v1 += 832
        case "a":
            v8 := v3[(^(fi6))]
            _ = v8
        default:
            v4 = (fi6 % ((fi6) | 1))
        }
        _ = fi6
        _ = v7
    }
    if (!true) {
        fmt.Println(v5)
        for fi9 := 0; fi9 < 4; fi9++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v10 uint32 = v1
            v5 = (v5 + v5)
            _ = fi9
            _ = v10
        }
    }
    if true {
        switch (v4 + v4) {
        case 0:
            var v11 int = -2147483648
            _ = v11
        case -742:
            fmt.Println(v5)
        case 906:
            v12 := v3[v4]
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint32 = 2147483647
        switch fi13 {
        case -2147483648:
            if (4294967295 != 0) {
                var v15 uint16 = 42
                if ((("" + (v5 + (v5 + (("foo bar" + "foo bar") + "foo bar")))) + (v5 + "\n\t")) != "hello") {
                    switch 128 {
                    case -1:
                        if true {
                            for fi16 := 0; fi16 < 5; fi16++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi16
                            }
                        } else {
                            var v17 uint32 = (v1 + (v14 | (^(v14))))
                            var v18 float64 = 0.1
                            v14 += (v1 * 730)
                            _ = v17
                            _ = v18
                        }
                    }
                } else {
                    v14 -= v14
                    for fi19 := 0; fi19 < 11; fi19++ {
                        var v20 bool = (v5 <= v5)
                        _ = fi19
                        _ = v20
                    }
                }
                _ = v15
            }
        }
        var v21 uint16 = 42
        _ = fi13
        _ = v14
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        v4 += fi22
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint16 = 402
            _ = fi23
            _ = v24
        }
        var v25 int64 = 1
        _ = fi22
        _ = v25
    }
    v1 += (v1 & (v1 << ((2) & 63)))
    if (2 > 40) {
        switch v4 {
        case -779:
            v1 -= (v1 >> ((1) & 63))
        case 9223372036854775807:
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 uint = 1
                _ = fi26
                _ = v27
            }
        default:
            v28 := GenS1{F0: 2, F1: v1, F2: (v4 << ((808) & 63))}
            _ = v28
        }
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
