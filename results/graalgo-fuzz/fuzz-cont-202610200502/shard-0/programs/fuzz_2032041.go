package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int32 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    v1 := []string{"foo bar", ("hello" + "hello")}
    v2 := make(map[string]uint32)
    var v3 float64 = -2.25
    var v4 uint32 = 546
    var v5 uint64 = 4294967295
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 11; fi7++ {
        switch 2147483647 {
        case 128:
            fmt.Println(v6, v4, v3)
        case -506:
            var v8 int = (fi7 % ((int(v4)) | 1))
            _ = v8
        default:
            if (!false) {
                if true {
                    var v9 uint16 = 42
                    v10 := v2[(("a" + "") + "a")]
                    _ = v10
                    _ = v9
                }
                v6 = (v6 << ((885) & 63))
            }
        }
        _ = fi7
    }
    v11 := GenS1{F0: 42, F1: "hello", F2: 1}
    v6 = (v6 ^ 42)
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            if false {
                v14 := GenS2{F0: 224, F1: "foo bar"}
                var v15 uint16 = 255
                _ = v14
                _ = v15
            }
            var v16 string = "hello"
            v17 := v2["foo bar"]
            var v18 uint64 = (v5 * 288)
            _ = fi13
            _ = v16
            _ = v17
            _ = v18
        }
        if true {
            v6 -= 43
            var v19 int16 = 124
            v5 -= v5
            _ = v19
        }
        _ = fi12
    }
    v4 = v4
    v20 := GenS1{F0: 4294967295, F1: ("a" + "a"), F2: 2147483647}
    v5 = (v5 % (((v5 + v5)) | 1))
    v6 -= v6
    if true {
        v6 += (v6 / ((v6) | 1))
        for fi21 := 0; fi21 < 8; fi21++ {
            if (42 != -2147483648) {
                v22 := GenS2{F0: uint(v5), F1: "\n\t"}
                _ = v22
            } else {
                switch (fi21 & fi21) {
                case -1:
                    v3 = v3
                }
                v5 += 94
            }
            _ = fi21
        }
        v4 -= v4
    } else {
        var v23 uint16 = 1
        var v24 uint8 = (^(v6))
        _ = v23
        _ = v24
    }
    switch (("a" + "foo bar") + "") {
    case "hello":
        var v25 int64 = -9223372036854775808
        _ = v25
    }
    var v26 string = "a"
    delete(v2, "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
