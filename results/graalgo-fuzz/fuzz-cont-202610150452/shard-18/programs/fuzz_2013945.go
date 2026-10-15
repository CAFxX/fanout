package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
func genFn0(p0 uint64, p1 int8) uint16 {
    var v4 int32 = 2147483647
    v4 += v4
    _ = p0
    _ = p1
    _ = v4
    return 118
}

func main() {
    var v1 int64 = 1
    v2 := []int{32767, 1}
    var v3 string = "\n\t"
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v3)
    case "":
        for fi5 := 0; fi5 < 7; fi5++ {
            v1 = (v1 / (((v1 ^ 0)) | 1))
            for fi6 := 0; fi6 < 18; fi6++ {
                switch fi6 {
                case 9223372036854775807:
                    var v7 int8 = 100
                    _ = v7
                case -2147483648:
                    if ((1 >= 18446744073709551615) || (-1 <= 0)) {
                        if len(v2) > 0 {
                            v2[0] = (fi5 / (((fi5 ^ fi5)) | 1))
                        }
                        v3 = v3
                        switch v3 {
                        case "\n\t":
                            v8 := GenS1{F0: 255, F1: (v3 + v3)}
                            _ = v8
                        }
                    }
                }
                if false {
                    var v9 bool = (!false)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v9, v9, v9)
                    _ = v9
                } else {
                    v10 := GenS1{F0: 42, F1: v3}
                    var v11 string = v3
                    _ = v10
                    _ = v11
                }
                _ = fi6
            }
            if true {
                var v12 bool = false
                _ = v12
            } else {
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 int = (fi5 - (fi5 / ((fi13) | 1)))
                    v15 := GenS1{F0: 57, F1: ("foo bar" + "foo bar")}
                    v3 = (v3 + ("hello" + v3))
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v2 = append(v2, -391)
                var v16 bool = (0 != 42)
                _ = v16
            }
            _ = fi5
        }
    }
    v17 := GenS1{F0: 2, F1: "hello"}
    v3 = v3
    if len(v2) > 0 {
        v2[0] = -102
    }
    v3 = (("a" + v3) + "")
    fmt.Println(v1, v3)
    var v18 uint8 = 191
    var v19 float32 = 1.5
    v20 := GenS1{F0: v18, F1: ("\n\t" + ("\n\t" + v3))}
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int16 = -1
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 127, F1: ((v3 + ((v3 + v3) + v3)) + "foo bar")}
    var v24 string = v3
    if true {
        switch v24 {
        case "foo bar":
            var v25 int16 = 1000
            _ = v25
        case "":
            switch v3 {
            case "a":
                switch v24 {
                case "\n\t":
                    v2 = append(v2, 893)
                }
            case "foo bar":
                v18 += 82
            default:
                if len(v2) > 0 {
                    v2[0] = 32767
                }
            }
        case "a":
            var v26 int64 = v1
            _ = v26
        default:
            v24 = "foo bar"
        }
        if false {
            var v27 int32 = -2147483648
            fmt.Println(v27, v24, v18)
            v3 = (("a" + v24) + "foo bar")
            _ = v27
        }
        if (0.1 > 1e10) {
            fmt.Println(v19, v18, v24)
            v28 := GenS1{F0: (^(v18)), F1: (v24 + ((v24 + ((("\n\t" + v24) + ("foo bar" + (("a" + "") + "\n\t"))) + "a")) + (v3 + "")))}
            _ = v28
        }
    }
    var v29 int = 26
    v24 = "hello"
    v30 := GenS1{F0: v18, F1: ("hello" + "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
}
