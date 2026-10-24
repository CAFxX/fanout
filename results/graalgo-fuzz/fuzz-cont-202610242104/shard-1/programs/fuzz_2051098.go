package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 int16, p1 uint16, p2 int32) uint32 {
    var v6 uint64 = 42
    p1 -= (p1 - (p1 / ((32767) | 1)))
    var v7 string = (("\n\t" + "a") + "\n\t")
    p2 = 1
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int64) int8 {
    var v9 bool = false
    fmt.Println(p0)
    _ = p0
    _ = v9
    return 100
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]uint32)
    var v3 int16 = 42
    var v4 uint32 = 276
    v5 := []string{"\n\t", ("foo bar" + ("" + ""))}
    v10 := GenS1{F0: v3}
    v11 := GenS1{F0: 1000}
    if true {
        var v12 float32 = 1.5
        fmt.Println(v12)
        _ = v12
    }
    if ((((1.5 - 1e10) / -2.25) * 3.14159) >= 3.14159) {
        if false {
            fmt.Println(v3, v1)
        }
        v13 := GenS2{F0: "foo bar", F1: 756}
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 string = ("foo bar" + ("foo bar" + "\n\t"))
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        fmt.Println(v1, v4)
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS1{F0: 1000}
            v4 -= (v4 & 4294967295)
            v18 := v2[((("a" + ("a" + "")) + "a") + (("" + ("a" + "")) + "\n\t"))]
            _ = fi16
            _ = v17
            _ = v18
        }
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
    }
    v5 = append(v5, "foo bar")
    fmt.Println(v3)
    v19 := v2[("\n\t" + "hello")]
    if (!false) {
        v19 = 2147483648
        var v20 int16 = (v3 % ((int16(v1)) | 1))
        switch "" {
        case "\n\t":
            for fi21 := 0; fi21 < 10; fi21++ {
                if len(v5) > 0 {
                    v5[0] = ((("foo bar" + "hello") + "\n\t") + "\n\t")
                }
                _ = fi21
            }
        case "a":
            var v22 string = ("hello" + "a")
            _ = v22
        }
        _ = v20
    }
    fmt.Println(v3)
    if (65535 > 575) {
        var v23 int32 = (v1 & 1)
        v4 += v19
        var v24 int32 = (v23 / (((v23 << ((2) & 63))) | 1))
        _ = v23
        _ = v24
    } else {
        if false {
            _ = len(v2)
            for fi25 := 0; fi25 < 20; fi25++ {
                switch (fi25 * 1) {
                case 128:
                    var v26 string = "\n\t"
                    _ = v26
                case 9223372036854775807:
                    delete(v2, "")
                default:
                    var v27 int = (fi25 & fi25)
                    _ = v27
                }
                _ = fi25
            }
        }
        v1 = v1
    }
    v2["foo bar"] = (^(v19))
    v4 = v19
    v28 := v2["hello"]
    v29 := v2["a"]
    v2[(("hello" + "\n\t") + ("foo bar" + "a"))] = (v28 + (v28 + (v29 ^ (v19 / (((v19 * (v4 << ((905) & 63)))) | 1)))))
    v1 = v1
    if (uint8(v28) >= 13) {
        switch "a" {
        case "hello":
            v28 += (^(v4))
        case "":
            v4 = (v29 ^ v4)
        case "\n\t":
            v1 = (-(v1))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
