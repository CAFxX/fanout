package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int64 }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]int)
    var v3 bool = false
    v4 := []string{}
    if ((4294967295 < 1) || v3) {
        v2[("hello" + "hello")] = 928
    } else {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int8 = -75
            if len(v4) > 0 {
                v4[0] = ""
            }
            for fi7 := 0; fi7 < 1; fi7++ {
                v1 = v1
                _ = fi7
            }
            _ = fi5
            _ = v6
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int64 = 42
            if false {
                v9 += (v9 * v9)
            }
            v9 = v9
            _ = fi8
            _ = v9
        }
        var v10 uint32 = 2147483647
        _ = v10
    }
    v1 += v1
    _ = len(v2)
    switch "" {
    case "\n\t":
        v11 := GenS1{F0: 2147483647, F1: "a", F2: 9223372036854775807}
        _ = v11
    case "hello":
        var v12 float64 = v1
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        v1 = v1
        v14 := GenS1{F0: (fi13 - (fi13 / (((-(fi13))) | 1))), F1: "\n\t", F2: 42}
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 822
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int64 = 9223372036854775807
        if (v3 || true) {
            fmt.Println(fi16)
        } else {
            v18 := v2["hello"]
            _ = v18
        }
        var v19 int = (fi16 - (^(fi16)))
        _ = fi16
        _ = v17
        _ = v19
    }
    var v20 uint16 = v15
    if (!true) {
        var v21 int = 32768
        v4 = append(v4, (("foo bar" + "a") + ""))
        switch "foo bar" {
        case "\n\t":
            if (v3 && false) {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                delete(v2, "foo bar")
            }
        }
        _ = v21
    } else {
        _ = len(v4)
    }
    fmt.Println(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
