package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int64 }

func main() {
    var v1 uint16 = 0
    var v2 int64 = -1
    var v3 string = "hello"
    var v4 uint64 = 0
    var v5 int = 32767
    v6 := []int{}
    if false {
        fmt.Println(v1)
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(fi7, v3, fi7)
            _ = len(v6)
            _ = fi7
        }
    } else {
        v8 := GenS1{F0: 691, F1: -234}
        if (!true) {
            if (!true) {
                switch "a" {
                case "hello":
                    v9 := GenS1{F0: 586, F1: 179}
                    _ = v9
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    var v10 int32 = 935
                    _ = v10
                }
                fmt.Println(v2)
                var v11 uint16 = 0
                _ = v11
            }
        } else {
            fmt.Println(v5, v2, v3)
        }
        v4 -= v4
        _ = v8
    }
    if (42 <= 1) {
        switch (v5 | (v5 * (v5 >> ((4294967295) & 63)))) {
        case -414:
            var v12 float32 = 0.0
            _ = v12
        case 650:
            if false {
                switch v3 {
                case "foo bar":
                    _ = len(v6)
                case "\n\t":
                    switch (((("" + v3) + v3) + ("a" + ("foo bar" + ("\n\t" + ("\n\t" + v3))))) + "\n\t") {
                    case "hello":
                        if len(v6) > 0 {
                            v6[0] = (v5 * (v5 >> ((4294967295) & 63)))
                        }
                    }
                default:
                    var v13 uint16 = (v1 | 65535)
                    _ = v13
                }
                if false {
                    _ = len(v6)
                } else {
                    v14 := GenS1{F0: 42, F1: 1}
                    switch v5 {
                    case 0:
                        var v15 int64 = (v2 - 1)
                        _ = v15
                    default:
                        v16 := GenS1{F0: (^(v2)), F1: 957}
                        _ = v16
                    }
                    var v17 uint = 0
                    _ = v14
                    _ = v17
                }
                v18 := GenS1{F0: v2, F1: 1}
                _ = v18
            } else {
                var v19 uint8 = 0
                v19 = (^(v19))
                _ = v19
            }
        default:
            if true {
                var v20 uint16 = (v1 - v1)
                _ = v20
            }
        }
        if true {
            for fi21 := 0; fi21 < 2; fi21++ {
                fmt.Println(v5, v5)
                v22 := GenS2{F0: (-(fi21))}
                v23 := GenS3{F0: ((!true) || true), F1: 0}
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    }
    fmt.Println(v4)
    if true {
        switch (v3 + (((((v3 + v3) + v3) + "hello") + "a") + v3)) {
        case "":
            if len(v6) > 0 {
                v6[0] = -1
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = len(v6)
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = len(v6)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
