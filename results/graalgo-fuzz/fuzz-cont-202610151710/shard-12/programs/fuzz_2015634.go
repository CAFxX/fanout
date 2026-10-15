package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }

func main() {
    var v1 uint16 = 42
    v2 := []string{"\n\t", "", "a"}
    var v3 int64 = 1
    var v4 uint32 = 42
    if len(v2) > 0 {
        v2[0] = ("foo bar" + ((("a" + "") + "foo bar") + "\n\t"))
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        _ = fi5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: "foo bar", F1: v3}
        v3 += (v3 ^ (^(v3)))
        _ = fi6
        _ = v7
    }
    v1 = 42
    switch 4294967295 {
    case -9223372036854775808:
        var v8 int8 = 0
        _ = v8
    case 32767:
        v9 := GenS1{F0: "\n\t", F1: int64(v4)}
        _ = v9
    default:
        switch int(v1) {
        case 9223372036854775807:
            switch -1 {
            case -2147483648:
                var v10 float32 = (0.0 * (0.0 * 0.1))
                _ = v10
            case 4294967295:
                var v11 float64 = 1.5
                _ = v11
            default:
                var v12 uint64 = 2
                _ = v12
            }
        case 1:
            fmt.Println(v4, v4, v4)
        default:
            v4 = v4
        }
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        if true {
            var v14 int64 = v3
            v4 = 42
            _ = v14
        } else {
            v15 := GenS1{F0: ("foo bar" + ("\n\t" + ("" + "foo bar"))), F1: v3}
            switch (fi13 ^ (fi13 - (fi13 | fi13))) {
            case 9223372036854775807:
                v2 = append(v2, "foo bar")
            case 42:
                for fi16 := 0; fi16 < 9; fi16++ {
                    v17 := GenS1{F0: "foo bar", F1: (v3 % (((^(v3))) | 1))}
                    _ = fi16
                    _ = v17
                }
            case -855:
                if (!(!false)) {
                    v1 = (v1 - (v1 - v1))
                    var v18 uint16 = 0
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        v1 = (v1 & (v1 >> ((65535) & 63)))
                        _ = fi19
                    }
                }
            }
            _ = v15
        }
        _ = fi13
    }
    var v20 int16 = -1
    var v21 uint64 = 244
    for fi22 := 0; fi22 < 9; fi22++ {
        switch "\n\t" {
        case "foo bar":
            v3 += (v3 % ((v3) | 1))
        case "a":
            v21 = (v21 % ((v21) | 1))
        default:
            if ((v20 / ((v20) | 1)) < v20) {
                v21 += (v21 + v21)
                if len(v2) > 0 {
                    v2[0] = ("" + "hello")
                }
                var v23 float32 = 0.0
                _ = v23
            }
        }
        _ = len(v2)
        _ = fi22
    }
    fmt.Println(v3, v20)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
