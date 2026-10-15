package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 int32; F1 string; F2 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 2147483647
    var v3 uint64 = 434
    v4 := make(map[string]string)
    fmt.Println(v2, v3)
    fmt.Println(v3)
    v3 += 4294967295
    var v5 int64 = -9223372036854775808
    fmt.Println(v3, v5)
    if true {
        var v6 uint8 = 1
        _ = v6
    } else {
        v1["\n\t"] = "foo bar"
        v5 = (v5 & v5)
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 int16 = 32767
            v8 -= int16(v3)
            _ = fi7
            _ = v8
        }
    }
    switch -100 {
    case 9223372036854775807:
        v9 := v4[""]
        _ = v9
    case 42:
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                if (!false) {
                    var v12 uint = 2
                    v13 := GenS1{F0: uint(v3)}
                    switch "hello" {
                    case "\n\t":
                        var v14 uint8 = 42
                        _ = v14
                    case "a":
                        if (42 > v3) {
                            v15 := GenS1{F0: v12}
                            _ = v15
                        }
                    }
                    _ = v12
                    _ = v13
                }
                var v16 float32 = ((-2.25 / 1.0) * 1e10)
                switch "foo bar" {
                case "\n\t":
                    v17 := GenS2{F0: 1, F1: "hello", F2: 0}
                    _ = v17
                }
                _ = fi11
                _ = v16
            }
            _ = fi10
        }
    case 4294967295:
        if true {
            fmt.Println(v3, v5, v2)
        } else {
            var v18 bool = ((-128 <= 1) && true)
            var v19 uint16 = 0
            _ = v18
            _ = v19
        }
    }
    switch ((("a" + ("hello" + (("hello" + "hello") + "a"))) + "a") + ("\n\t" + ("" + "\n\t"))) {
    case "a":
        v2 = (v2 | 991)
    default:
        v20 := GenS3{F0: 2147483647, F1: (("hello" + "\n\t") + ("foo bar" + "\n\t")), F2: v2}
        _ = v20
    }
    v2 -= (v2 + v2)
    var v21 bool = (!true)
    var v22 bool = (((true && ((!v21) || (0.1 != 3.14159))) && true) && v21)
    var v23 int32 = 42
    v24 := GenS2{F0: v2, F1: (("\n\t" + ("\n\t" + ((("hello" + "") + "\n\t") + ""))) + ("hello" + "foo bar")), F2: -128}
    delete(v4, "foo bar")
    var v25 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
