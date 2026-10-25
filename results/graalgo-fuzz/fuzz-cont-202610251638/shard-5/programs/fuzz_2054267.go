package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint; F2 string }
type GenS3 struct { F0 string; F1 int8 }

func main() {
    var v1 int = -9223372036854775808
    v2 := []string{"\n\t"}
    v3 := make(map[string]bool)
    var v4 int16 = 32767
    v5 := GenS1{F0: 260}
    v6 := GenS1{F0: 42}
    switch (("foo bar" + (("hello" + ("a" + "foo bar")) + ((("a" + (("" + "") + "")) + "") + ""))) + "a") {
    case "foo bar":
        v1 = v1
    case "\n\t":
        if false {
            v1 = v1
            switch (v1 | v1) {
            case -9223372036854775808:
                var v7 uint32 = 1
                _ = v7
            case 42:
                v2 = append(v2, ("a" + "a"))
            default:
                v1 = 1
            }
            v8 := GenS1{F0: 891}
            _ = v8
        } else {
            var v9 float64 = (1e10 / (float64(1e10) / 0.1))
            switch "" {
            case "a":
                var v10 int16 = (v4 | (v4 + v4))
                _ = v10
            default:
                var v11 int32 = 1
                _ = v11
            }
            _ = v9
        }
    default:
        if (!true) {
            if ((v1 >> ((4294967295) & 63)) != v1) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    v4 = (^(v4))
                    v1 = v1
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 int8 = 42
                        if true {
                            var v15 string = "hello"
                            _ = v15
                        }
                        var v16 int = (fi12 / (((fi13 + 9223372036854775807)) | 1))
                        _ = fi13
                        _ = v14
                        _ = v16
                    }
                    _ = fi12
                }
                var v17 bool = (true || (true && (!(true || (128 == 255)))))
                _ = v17
            }
        } else {
            v18 := GenS1{F0: 0}
            v19 := GenS2{F0: 419, F1: 2, F2: (("foo bar" + "") + "a")}
            fmt.Println(v4, v1)
            _ = v18
            _ = v19
        }
    }
    v1 -= (v1 ^ 128)
    v2 = append(v2, "\n\t")
    v1 += (v1 ^ 42)
    v20 := v3["\n\t"]
    fmt.Println(v20, v4)
    v20 = (int32(v4) >= -290)
    v4 = (v4 ^ -32768)
    v21 := GenS3{F0: "a", F1: 1}
    fmt.Println(v4, v20, v1)
    v22 := GenS2{F0: 335, F1: 1, F2: ""}
    if (613 > 760) {
        var v23 uint8 = 100
        v24 := GenS2{F0: 0, F1: 255, F2: "\n\t"}
        var v25 float32 = ((-2.25 * 1.5) + 0.0)
        _ = v23
        _ = v24
        _ = v25
    }
    if v20 {
        var v26 float32 = (-2.25 - float32(0.1))
        _ = v26
    } else {
        v4 += v4
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
