package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 42
    v3 := make(map[uint32]string)
    var v4 int = v2
    v5 := []string{"", ("foo bar" + "a"), "hello", "\n\t"}
    switch (v4 - (v4 % (((^(v4))) | 1))) {
    case 2147483647:
        if false {
            v4 += -101
            fmt.Println(v2)
        }
    case 1:
        var v6 float32 = 1.5
        _ = v6
    case -1:
        if true {
            var v7 int8 = 42
            v4 -= 32768
            if false {
                v4 -= int(v7)
            } else {
                switch 109 {
                case -967:
                    v8 := GenS1{F0: -392, F1: "hello", F2: 2147483648}
                    _ = v8
                case 2147483647:
                    if true {
                        switch (("foo bar" + "a") + "foo bar") {
                        case "hello":
                            v7 = (^(v7))
                        case "foo bar":
                            fmt.Println(v2, v2, v7)
                        case "a":
                            var v9 float32 = (3.14159 * 1e10)
                            _ = v9
                        }
                    }
                }
                v7 -= -42
                var v10 uint64 = 1
                _ = v10
            }
            _ = v7
        } else {
            switch (v4 | (v2 & v4)) {
            case -1:
                switch "" {
                case "\n\t":
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                default:
                    v11 := GenS1{F0: -573, F1: "foo bar", F2: 65535}
                    _ = v11
                }
            }
            v4 -= v2
        }
    }
    var v12 uint64 = 994
    var v13 int = 9223372036854775807
    v14 := GenS1{F0: 1, F1: "a", F2: 4294967295}
    for fi15 := 0; fi15 < 6; fi15++ {
        v3[2147483648] = (((("hello" + ("" + ("foo bar" + ("a" + "foo bar")))) + ("foo bar" + (("a" + ("hello" + ("a" + ""))) + ""))) + ("a" + "a")) + "a")
        v16 := GenS2{F0: "a"}
        var v17 uint64 = v12
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v12)
    for fi18 := 0; fi18 < 6; fi18++ {
        if (!((0 >= 42) && (true && (!true)))) {
            v12 += (v12 + (v12 * v12))
        }
        switch "foo bar" {
        case "":
            switch (v4 + 254) {
            case -805:
                v13 = v4
            case 32767:
                v12 -= (v12 >> ((682) & 63))
            }
        default:
            v19 := GenS3{F0: (!false), F1: 489}
            _ = v19
        }
        _ = fi18
    }
    var v20 bool = true
    var v21 int = v4
    fmt.Println(v21)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) > 0 {
        v5[0] = ("hello" + ("\n\t" + "hello"))
    }
    var v22 int32 = -445
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
