package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int = 128
    v2 := make(map[string]uint32)
    v3 := []string{"\n\t", (("foo bar" + "\n\t") + "a")}
    var v4 int32 = 2147483647
    for fi5 := 0; fi5 < 5; fi5++ {
        if false {
            fmt.Println(v4, fi5)
        } else {
            switch ("\n\t" + "") {
            case "\n\t":
                for fi6 := 0; fi6 < 2; fi6++ {
                    switch fi6 {
                    case -771:
                        var v7 int = fi6
                        _ = v7
                    case -9223372036854775808:
                        v2[""] = 4294967295
                    }
                    switch (-(fi5)) {
                    case 9223372036854775807:
                        if (false && (v1 >= (v1 / ((32768) | 1)))) {
                            v4 = (v4 >> ((1) & 63))
                            var v8 string = "foo bar"
                            v4 = v4
                            _ = v8
                        }
                    }
                    var v9 int16 = -1
                    v3 = append(v3, "\n\t")
                    v9 += v9
                    _ = fi6
                    _ = v9
                }
            default:
                switch "hello" {
                case "a":
                    var v10 float64 = (1e10 - 0.1)
                    _ = v10
                default:
                    v11 := GenS1{F0: 65535}
                    _ = v11
                }
            }
            var v12 int64 = 42
            _ = len(v2)
            _ = v12
        }
        var v13 int = (v1 * v1)
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = fi5
        _ = v13
    }
    switch ("\n\t" + ("a" + ("hello" + ""))) {
    case "\n\t":
        fmt.Println(v4)
    case "":
        fmt.Println(v1, v4)
    case "a":
        v14 := GenS1{F0: 255}
        _ = v14
    }
    var v15 int = v1
    switch 127 {
    case -406:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 32767:
        var v16 uint32 = 232
        _ = v16
    case 908:
        v15 = 4294967295
    default:
        for fi17 := 0; fi17 < 16; fi17++ {
            v15 += (v1 & 1)
            v15 = (^(v1))
            _ = fi17
        }
    }
    var v18 int64 = 42
    v19 := GenS1{F0: 2}
    _ = len(v2)
    var v20 int32 = v4
    v21 := v2[(("a" + "\n\t") + "foo bar")]
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
