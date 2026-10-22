package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
func genFn0() bool {
    var v5 uint64 = 2
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 bool = true
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 float64 = (1.5 + -2.25)
            _ = fi9
            _ = v10
        }
        _ = fi6
    }
    _ = v5
    return false
}
func genFn1() string {
    if false {
        var v11 int = -781
        v12 := GenS1{F0: ("\n\t" + "\n\t"), F1: 2, F2: -2147483648}
        _ = v11
        _ = v12
    } else {
        if false {
            var v13 float64 = float64(3.14159)
            v14 := GenS1{F0: ((("a" + "a") + "a") + ((("hello" + "hello") + "\n\t") + ((("a" + "") + "\n\t") + "\n\t"))), F1: 1, F2: 784}
            var v15 uint64 = 18446744073709551615
            _ = v13
            _ = v14
            _ = v15
        }
    }
    var v16 uint = 2
    _ = v16
    return ""
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]string)
    var v3 uint8 = 35
    v4 := []string{"\n\t", "a"}
    fmt.Println(v1, v3)
    switch "hello" {
    case "foo bar":
        for fi17 := 0; fi17 < 11; fi17++ {
            v4 = append(v4, ("\n\t" + "hello"))
            if ((!false) || true) {
                fmt.Println(v1, v1)
            } else {
                if (true || (!false)) {
                    var v18 float64 = 0.0
                    _ = v18
                } else {
                    var v19 uint8 = 128
                    _ = len(v4)
                    v20 := GenS1{F0: "foo bar", F1: v1, F2: 0}
                    _ = v19
                    _ = v20
                }
            }
            _ = fi17
        }
    case "":
        v3 += v3
    case "hello":
        fmt.Println(v3, v1)
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            if false {
                var v23 int = -1
                _ = v23
            }
            _ = fi22
        }
        fmt.Println(fi21)
        _ = fi21
    }
    switch 2147483647 {
    case -286:
        v2[128] = "\n\t"
    }
    var v24 uint = 255
    v24 = (v24 % (((v1 / (((v24 * v1)) | 1))) | 1))
    var v25 int8 = 42
    v3 -= (^(v3))
    fmt.Println(v25, v24, v1)
    if (false && (32767 < 738)) {
        v25 -= v25
    } else {
        v1 = (^(v24))
        var v26 int8 = (-(v25))
        v3 = v3
        _ = v26
    }
    v24 -= 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
