package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 int64 }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
func genFn0(p0 uint64, p1 string) string {
    var v5 int32 = 577
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 string = ""
        fmt.Println(p0)
        _ = fi6
        _ = v7
    }
    var v8 uint = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return "\n\t"
}

func main() {
    var v1 int16 = -337
    var v2 uint64 = 4294967295
    var v3 uint8 = 2
    v4 := make(map[string]bool)
    for fi9 := 0; fi9 < 13; fi9++ {
        v2 += 4294967295
        v1 = v1
        _ = fi9
    }
    v10 := GenS1{F0: "hello", F1: (v1 | -32768), F2: (v2 % ((1) | 1))}
    if (!true) {
        v11 := GenS1{F0: (("\n\t" + ("\n\t" + ((((("" + "a") + ((("foo bar" + "\n\t") + "\n\t") + "\n\t")) + "a") + ("a" + "")) + ""))) + "hello"), F1: 380, F2: (v2 << ((2) & 63))}
        fmt.Println(v1)
        _ = v11
    } else {
        if (633 != 2147483647) {
            if false {
                if ((v2 ^ v2) == 110) {
                    var v12 string = "hello"
                    for fi13 := 0; fi13 < 6; fi13++ {
                        var v14 int64 = -277
                        _ = fi13
                        _ = v14
                    }
                    v15 := GenS2{F0: 65535, F1: -1, F2: -894}
                    _ = v12
                    _ = v15
                }
                for fi16 := 0; fi16 < 6; fi16++ {
                    if (-1 > int64(fi16)) {
                        fmt.Println(v1, v3)
                        v3 -= v3
                    }
                    _ = fi16
                }
            } else {
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 string = ""
                    var v19 uint64 = v2
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                if false {
                    var v20 int = -9223372036854775808
                    v21 := GenS1{F0: "foo bar", F1: v1, F2: 522}
                    _ = v20
                    _ = v21
                } else {
                    v1 = 281
                    v22 := []int{}
                    _ = v22
                }
                v1 = v1
            }
            if true {
                v3 += (v3 ^ (v3 << ((926) & 63)))
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                _ = len(v4)
                var v24 uint = 65535
                _ = fi23
                _ = v24
            }
        }
        v25 := v4[(("\n\t" + ("" + "foo bar")) + "")]
        var v26 int64 = 42
        _ = v25
        _ = v26
    }
    v3 -= 255
    v3 += uint8(v2)
    v2 += v2
    for fi27 := 0; fi27 < 16; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 string = "hello"
            v30 := GenS1{F0: (v29 + "a"), F1: 105, F2: (v2 | (v2 - (v2 + 2)))}
            _ = fi28
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 11; fi31++ {
            if (1 > 1) {
                var v32 float64 = 0.1
                v2 += v2
                v33 := v4["hello"]
                _ = v32
                _ = v33
            }
            _ = fi31
        }
        fmt.Println(v1, v3)
        _ = fi27
    }
    v34 := []int{1, 4294967295, 32768}
    v35 := GenS3{F0: 1, F1: ("foo bar" + "hello"), F2: v1}
    if false {
        if true {
            switch ((("hello" + "hello") + ("foo bar" + (("" + "hello") + ""))) + ("hello" + ("\n\t" + (("" + "") + ("" + ("" + "\n\t")))))) {
            case "foo bar":
                _ = len(v34)
            default:
                var v36 int32 = -1
                _ = v36
            }
            switch 1 {
            case -2147483648:
                var v37 uint64 = 18446744073709551615
                _ = v37
            default:
                switch "" {
                case "a":
                    fmt.Println(v3, v2)
                case "":
                    v1 = -1
                case "\n\t":
                    if len(v34) >= 2 {
                        v34 = v34[1:len(v34)]
                    }
                }
            }
            var v38 uint64 = v2
            _ = v38
        }
    } else {
        _ = len(v4)
        v34 = append(v34, -123)
    }
    v39 := GenS3{F0: 92, F1: "foo bar", F2: v1}
    var v40 int64 = -9223372036854775808
    var v41 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
