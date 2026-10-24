package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() string {
    v7 := []int{}
    v7 = append(v7, 932)
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            switch "" {
            case "a":
                fmt.Println(fi8, fi9)
            case "":
                var v10 int32 = 767
                _ = v10
            default:
                var v11 float32 = 0.0
                _ = v11
            }
            if (3.14159 == (float64(1e10) / ((1e10 + 0.1) / (1.5 / 0.1)))) {
                var v12 uint16 = 32767
                _ = v12
            } else {
                var v13 int16 = 0
                v13 += (v13 / (((v13 | v13)) | 1))
                if ((0.1 / 0.1) != (-2.25 + 1e10)) {
                    v13 = (v13 % (((v13 + 32767)) | 1))
                }
                _ = v13
            }
            _ = len(v7)
            var v14 int32 = -2147483648
            _ = fi9
            _ = v14
        }
        var v15 uint64 = 4294967295
        _ = fi8
        _ = v15
    }
    _ = v7
    return ("" + ("" + "a"))
}
func genFn1(p0 uint64) int8 {
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := []int{(-(fi16))}
        _ = len(v17)
        _ = fi16
        _ = v17
    }
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 127
}

func main() {
    v1 := []uint32{4294967295, 65535}
    var v2 bool = true
    v3 := []string{"foo bar", "hello", ("a" + "foo bar"), (("foo bar" + ((("foo bar" + "\n\t") + "") + ("\n\t" + (("\n\t" + ("a" + (((("a" + "") + "hello") + (("foo bar" + (("\n\t" + ("foo bar" + ("a" + ("" + "")))) + "hello")) + ("\n\t" + ""))) + ""))) + "")))) + "a")}
    var v4 uint16 = 42
    var v5 int = 224
    var v6 uint64 = 42
    var v18 int32 = 2147483647
    switch "hello" {
    case "":
        v6 += 2
    case "foo bar":
        v1 = append(v1, 1)
    case "a":
        v2 = (v2 && (9223372036854775807 != v5))
    }
    var v19 int8 = -115
    if (v2 && (int8(v18) <= 107)) {
        var v20 int32 = -2147483648
        _ = v20
    } else {
        v21 := GenS1{F0: v18}
        _ = v21
    }
    var v22 int = -2147483648
    switch ("hello" + "hello") {
    case "foo bar":
        var v23 uint16 = v4
        _ = v23
    }
    fmt.Println(v5, v6, v2)
    fmt.Println(v18, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
