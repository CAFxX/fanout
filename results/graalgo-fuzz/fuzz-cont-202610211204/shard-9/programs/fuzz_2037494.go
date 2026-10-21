package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 uint64 }
func genFn0(p0 int, p1 string, p2 int) uint32 {
    v5 := GenS1{F0: p1, F1: true, F2: 2}
    var v6 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]bool)
    var v3 uint = 40
    var v4 bool = (false && (!false))
    var v7 uint16 = 1
    if v4 {
        fmt.Println(v7, v4)
        v8 := GenS1{F0: ("\n\t" + ((("a" + "a") + ((("" + "a") + ("hello" + "hello")) + "a")) + "")), F1: v4, F2: 0}
        _ = v8
    }
    switch (("foo bar" + "hello") + ("foo bar" + "a")) {
    case "":
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := []int{fi9, 32767}
            v11 := GenS1{F0: "\n\t", F1: false, F2: 0}
            var v12 uint64 = 612
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    case "hello":
        var v13 int = 545
        _ = v13
    }
    if false {
        var v14 string = "foo bar"
        _ = v14
    }
    v15 := []int{int(v3), 552, -1, 9223372036854775807}
    if (!(int64(v3) > (v1 * (v1 / ((v1) | 1))))) {
        for fi16 := 0; fi16 < 1; fi16++ {
            switch ("" + "") {
            case "hello":
                v3 -= (^(v3))
            case "\n\t":
                if ((false || (1 <= 127)) && (((v4 && (42 <= -1)) || false) || v4)) {
                    v17 := GenS2{F0: v3, F1: "foo bar", F2: 2}
                    fmt.Println(v3)
                    switch "\n\t" {
                    case "\n\t":
                        if ((^(v7)) > (v7 * (^(v7)))) {
                            fmt.Println(fi16)
                            fmt.Println(v4, fi16)
                        }
                    case "":
                        v18 := GenS2{F0: 65535, F1: "a", F2: 395}
                        _ = v18
                    }
                    _ = v17
                }
            case "":
                v19 := GenS2{F0: 0, F1: "hello", F2: 454}
                _ = v19
            default:
                var v20 uint64 = 2
                _ = v20
            }
            if (v4 || true) {
                switch "a" {
                case "a":
                    var v21 float32 = (float32(1e10) + (-2.25 + ((float32((0.0 - ((((1e10 - (0.1 / 1.0)) * ((((((0.1 + -2.25) * 3.14159) * ((3.14159 / (((1e10 * 0.0) + 1.5) / ((1.5 - 0.0) - (0.0 - float64(0.1))))) / 3.14159)) * 1e10) / 3.14159) * (3.14159 + 3.14159))) * float64(0.1)) - float64((0.1 - 1.5))))) + 1.5) / ((1.5 / (1e10 - (1.5 / 0.1))) / float32((-2.25 / (((3.14159 + 3.14159) / (3.14159 - ((-2.25 / 0.1) * (1.5 - 0.0)))) - 0.1)))))))
                    _ = v21
                case "\n\t":
                    v3 = 255
                case "foo bar":
                    v15 = append(v15, fi16)
                default:
                    if (-1 >= 0) {
                        v1 = v1
                        delete(v2, 65535)
                        if (v4 || (v4 || false)) {
                            delete(v2, 65535)
                            _ = len(v2)
                            fmt.Println(v7, fi16)
                        }
                    }
                }
                switch 32767 {
                case 32768:
                    fmt.Println(v7)
                default:
                    v4 = v4
                }
                v7 += (v7 / (((v7 * v7)) | 1))
            } else {
                v3 = 2
                v15 = append(v15, 937)
                v22 := GenS1{F0: "\n\t", F1: (v4 && (-1 != 1)), F2: 2}
                _ = v22
            }
            if (4294967295 == 2147483647) {
                var v23 string = "hello"
                _ = v23
            }
            _ = fi16
        }
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        switch "" {
        case "":
            v7 = (v7 | v7)
        case "foo bar":
            v7 -= v7
        }
    }
    v24 := GenS1{F0: "foo bar", F1: (0 <= 0), F2: 42}
    var v25 uint8 = 210
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
