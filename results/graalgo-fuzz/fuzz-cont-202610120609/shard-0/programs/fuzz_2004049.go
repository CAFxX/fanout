package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint32, p1 uint, p2 int32) int32 {
    v5 := GenS1{F0: "a"}
    if (("" + "\n\t") > (("hello" + "a") + "hello")) {
        switch "\n\t" {
        case "a":
            for fi6 := 0; fi6 < 13; fi6++ {
                v7 := []int{(fi6 * 32767), (fi6 | fi6), fi6}
                _ = fi6
                _ = v7
            }
        case "hello":
            var v8 uint64 = 4294967295
            _ = v8
        default:
            var v9 float32 = (0.0 / float32(-2.25))
            _ = v9
        }
    }
    fmt.Println(p1)
    v10 := []int{-995, 128}
    switch ("\n\t" + ("foo bar" + "foo bar")) {
    case "":
        for fi11 := 0; fi11 < 17; fi11++ {
            switch ((("" + "a") + ("\n\t" + (("\n\t" + ((("" + "foo bar") + "") + "hello")) + ("hello" + "hello")))) + ("" + "\n\t")) {
            case "foo bar":
                switch "" {
                case "":
                    var v12 uint8 = uint8(p0)
                    _ = v12
                case "a":
                    switch fi11 {
                    case 42:
                        var v13 float32 = 0.0
                        _ = v13
                    case -740:
                        var v14 int16 = -1
                        _ = v14
                    case -985:
                        p0 = p0
                    default:
                        var v15 int32 = p2
                        _ = v15
                    }
                }
            case "":
                var v16 uint64 = 240
                _ = v16
            case "\n\t":
                var v17 int = 114
                _ = v17
            default:
                v18 := GenS2{F0: false, F1: (fi11 * -9223372036854775808)}
                _ = v18
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v19 := GenS2{F0: true, F1: (fi11 << ((p1) & 63))}
            _ = fi11
            _ = v19
        }
    case "\n\t":
        p0 = 4294967295
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return 42
}

func main() {
    var v1 int32 = -2147483648
    var v2 string = "hello"
    var v3 float32 = 0.0
    v4 := make(map[string]string)
    v2 = v2
    v2 = v2
    var v20 bool = ((936 == 2) && (v1 != (v1 / (((v1 % (((v1 / ((1) | 1))) | 1))) | 1))))
    var v21 string = v2
    var v22 uint64 = 2
    var v23 uint8 = 128
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := v4[v21]
        _ = fi24
        _ = v25
    }
    switch v2 {
    case "":
        v21 = ((v2 + ((v21 + (("" + "hello") + (v21 + "foo bar"))) + v2)) + (("hello" + "\n\t") + "\n\t"))
    case "hello":
        var v26 int32 = v1
        _ = v26
    case "a":
        for fi27 := 0; fi27 < 18; fi27++ {
            if v20 {
                if (v20 && (1 > -121)) {
                    v28 := []int{(fi27 + fi27), (-(fi27)), 0}
                    fmt.Println(fi27)
                    fmt.Println(v1, fi27, v21)
                    _ = v28
                }
            } else {
                var v29 int16 = -335
                _ = v29
            }
            _ = fi27
        }
    default:
        v30 := GenS3{F0: (v1 ^ v1), F1: (v22 | v22)}
        _ = v30
    }
    v31 := []int{int(v23), 1, 4294967295, 32767}
    var v32 uint = 18446744073709551615
    var v33 int32 = -602
    var v34 bool = v20
    for fi35 := 0; fi35 < 20; fi35++ {
        v23 = v23
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
