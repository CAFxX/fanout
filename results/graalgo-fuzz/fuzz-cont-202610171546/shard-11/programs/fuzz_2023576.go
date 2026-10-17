package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int32, p1 string) int {
    switch 32767 {
    case 128:
        switch -238 {
        case 9223372036854775807:
            if ((!(!(true && (2147483648 <= 4294967295)))) && ((true && false) || (!false))) {
                v4 := GenS1{F0: -981, F1: 127}
                var v5 uint32 = 2147483647
                _ = v4
                _ = v5
            } else {
                p0 += (p0 % ((p0) | 1))
            }
        }
    default:
        for fi6 := 0; fi6 < 13; fi6++ {
            switch p1 {
            case "":
                var v7 uint = 42
                _ = v7
            case "\n\t":
                p0 = p0
            default:
                switch (p1 + p1) {
                case "foo bar":
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(fi6)
                    case "":
                        v8 := []int{(fi6 >> ((391) & 63)), -2147483648, fi6, fi6}
                        _ = v8
                    }
                case "":
                    var v9 uint64 = 2
                    _ = v9
                }
            }
            _ = fi6
        }
    }
    v10 := GenS2{F0: -128}
    var v11 int8 = 100
    p0 = (p0 | p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 uint8 = 128
    v2 := []string{(("" + "foo bar") + "foo bar"), "\n\t"}
    v3 := make(map[string]uint32)
    v2 = append(v2, "")
    var v12 int8 = 1
    if len(v2) > 0 {
        v2[0] = ("\n\t" + (("a" + ("hello" + "a")) + ("hello" + (("\n\t" + "hello") + ""))))
    }
    if (int16(v12) == 42) {
        var v13 float32 = 1.5
        if (!((594 < 255) || (true && true))) {
            var v14 uint16 = 32767
            _ = v14
        }
        _ = v13
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        delete(v3, ("foo bar" + ("" + "a")))
        v1 = 0
        _ = fi15
    }
    v1 += (v1 % ((v1) | 1))
    fmt.Println(v1, v12)
    for fi16 := 0; fi16 < 12; fi16++ {
        if false {
            _ = len(v2)
            var v17 int32 = -1
            _ = v17
        }
        v1 += (v1 << ((580) & 63))
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
