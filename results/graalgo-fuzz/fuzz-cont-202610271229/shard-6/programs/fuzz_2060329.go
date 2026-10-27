package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int64 }
func genFn0(p0 uint) string {
    var v6 int = -326
    switch "foo bar" {
    case "foo bar":
        if (345 > 1) {
            if false {
                for fi7 := 0; fi7 < 3; fi7++ {
                    v8 := []int{v6, 9223372036854775807, 931}
                    for fi9 := 0; fi9 < 2; fi9++ {
                        switch (fi9 % ((-9223372036854775808) | 1)) {
                        case 128:
                            switch fi9 {
                            case -716:
                                var v10 float64 = (0.0 + 3.14159)
                                _ = v10
                            case -9223372036854775808:
                                var v11 bool = false
                                _ = v11
                            }
                        default:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                        v6 += (fi7 + (fi9 & (fi7 - fi7)))
                        p0 -= (^(p0))
                        _ = fi9
                    }
                    fmt.Println(fi7, v6, v6)
                    _ = fi7
                    _ = v8
                }
                v12 := []int{}
                _ = v12
            }
            fmt.Println(v6, v6)
            v6 += (v6 << (((p0 - p0)) & 63))
        }
    }
    v13 := GenS1{F0: 42, F1: 42, F2: 1}
    _ = p0
    _ = v13
    _ = v6
    return "hello"
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v4 uint8 = 42
    var v5 int64 = 18
    var v14 int8 = -34
    v5 = v5
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 int32 = 42
            _ = fi16
            _ = v17
        }
        v18 := []int{(fi15 << ((65535) & 63))}
        _ = fi15
        _ = v18
    }
    delete(v3, "\n\t")
    fmt.Println(v14)
    var v19 int32 = 897
    v2[0] = 65535
    fmt.Println(v4)
    v20 := v1[0]
    if ("foo bar" < "hello") {
        v5 = v5
    } else {
        v3["\n\t"] = 2147483647
        for fi21 := 0; fi21 < 9; fi21++ {
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 int16 = 42
                v3["foo bar"] = 2147483647
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        fmt.Println(v19)
    }
    var v24 int = int(v4)
    v24 = (v24 % (((v24 | -2147483648)) | 1))
    if ("hello" > ("\n\t" + "a")) {
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 int8 = (v14 / ((v14) | 1))
            _ = fi25
            _ = v26
        }
        if (42 > 42) {
            v19 = -2147483648
            switch "foo bar" {
            case "\n\t":
                var v27 int32 = -1
                _ = v27
            default:
                if false {
                    fmt.Println(v14)
                    v28 := []int{v24}
                    var v29 uint64 = 4294967295
                    _ = v28
                    _ = v29
                }
            }
        } else {
            var v30 uint16 = 42
            fmt.Println(v5, v4, v30)
            v31 := GenS1{F0: 18446744073709551615, F1: (v19 & -1), F2: v5}
            _ = v30
            _ = v31
        }
    }
    v32 := []int{v24, v24, v24}
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 uint = 606
        if (v20 || (v19 >= 42)) {
            v35 := v2[0]
            switch "" {
            case "hello":
                for fi36 := 0; fi36 < 10; fi36++ {
                    v34 = v34
                    _ = fi36
                }
            case "foo bar":
                fmt.Println(fi33, v19, v34)
            default:
                v34 = 500
            }
            _ = v35
        } else {
            v37 := GenS1{F0: 18446744073709551615, F1: (-(v19)), F2: v5}
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 string = "foo bar"
        v40 := GenS1{F0: 4294967295, F1: v19, F2: (v5 + (v5 >> ((4294967295) & 63)))}
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := v3[("\n\t" + "a")]
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v41
    _ = v5
}
