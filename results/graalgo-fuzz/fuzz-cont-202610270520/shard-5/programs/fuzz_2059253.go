package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint; F1 string }

func main() {
    v1 := []string{"hello"}
    v2 := []string{}
    var v3 bool = (!(-1 == 1))
    v4 := make(map[string]bool)
    v5 := []string{"\n\t", "", "foo bar"}
    v6 := GenS1{F0: 178, F1: v3}
    for fi7 := 0; fi7 < 9; fi7++ {
        _ = len(v1)
        if (!(float64(0.1) >= (0.1 * float64(0.1)))) {
            var v8 string = "foo bar"
            v9 := GenS1{F0: 32, F1: v3}
            if v3 {
                var v10 float32 = (0.1 + -2.25)
                v11 := GenS2{F0: 0, F1: 4294967295}
                _ = v10
                _ = v11
            } else {
                v12 := GenS3{F0: 2, F1: v8}
                _ = v12
            }
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    var v13 int = 577
    var v14 int64 = -9223372036854775808
    switch (("\n\t" + "hello") + "\n\t") {
    case "hello":
        if (v3 || v3) {
            var v15 uint8 = 55
            fmt.Println(v15)
            _ = v15
        }
    case "\n\t":
        if (443 != 722) {
            v4["foo bar"] = (681 < 255)
            for fi16 := 0; fi16 < 10; fi16++ {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v18 := v4[("\n\t" + "a")]
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        }
    case "a":
        v14 = v14
    default:
        v3 = false
    }
    v19 := GenS2{F0: -32768, F1: 2}
    fmt.Println(v14)
    if (!false) {
        v20 := GenS2{F0: 810, F1: 2147483647}
        v21 := GenS1{F0: 255, F1: (!v3)}
        for fi22 := 0; fi22 < 11; fi22++ {
            v23 := v4["foo bar"]
            var v24 int16 = -32768
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = v20
        _ = v21
    } else {
        v3 = (0 <= 42)
        v13 -= 9223372036854775807
    }
    v3 = true
    for fi25 := 0; fi25 < 20; fi25++ {
        if (v3 && (!(v3 || true))) {
            v13 += v13
        } else {
            fmt.Println(v13, v13, v3)
            if v3 {
                v3 = (!((-2.25 >= 1e10) || v3))
                v3 = (266 >= 2147483648)
            }
        }
        _ = fi25
    }
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "hello")
    }
    var v26 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
