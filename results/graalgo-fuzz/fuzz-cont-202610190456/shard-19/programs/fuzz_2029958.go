package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
type GenS3 struct { F0 bool }

func main() {
    v1 := []string{"", "foo bar", "hello", "foo bar"}
    v2 := []string{"foo bar"}
    var v3 int64 = 1
    var v4 float32 = 0.1
    var v5 float32 = (v4 + float32(1.5))
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 12; fi7++ {
        if true {
            for fi8 := 0; fi8 < 18; fi8++ {
                switch (("a" + "\n\t") + ("\n\t" + "foo bar")) {
                case "\n\t":
                    var v9 uint64 = 0
                    _ = v9
                case "foo bar":
                    v10 := GenS1{F0: 2147483648}
                    _ = v10
                case "hello":
                    v4 -= v4
                }
                _ = fi8
            }
        } else {
            if len(v1) > 0 {
                v1[0] = (("foo bar" + "\n\t") + "foo bar")
            }
        }
        if true {
            v11 := GenS2{F0: ("hello" + ""), F1: 42, F2: 242}
            v3 -= v3
            _ = v11
        } else {
            v4 = ((-2.25 + 0.0) / (v4 / 1.5))
            if true {
                for fi12 := 0; fi12 < 6; fi12++ {
                    v5 = 1e10
                    var v13 int32 = 0
                    if true {
                        if (true && (876 == 16)) {
                            v5 += 1.5
                        }
                    }
                    _ = fi12
                    _ = v13
                }
                var v14 uint16 = 0
                if false {
                    fmt.Println(fi7, v14, fi7)
                    _ = len(v2)
                }
                _ = v14
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int16 = 42
                fmt.Println(v16, fi7, fi15)
                _ = fi15
                _ = v16
            }
        }
        _ = fi7
    }
    v2 = append(v2, "")
    var v17 uint8 = 128
    v18 := GenS3{F0: (!((429 < -469) || false))}
    var v19 uint = 255
    var v20 uint = (v19 * 1)
    var v21 uint = (v20 * 65535)
    var v22 uint32 = 0
    switch 128 {
    case 128:
        fmt.Println(v17, v21)
    case 32768:
        var v23 int16 = -32768
        _ = v23
    case 32767:
        v24 := GenS3{F0: (2 > 2)}
        _ = v24
    default:
        var v25 string = "hello"
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
