package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := []string{"foo bar"}
    v4 := make(map[string]bool)
    var v5 int32 = -391
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := v4["foo bar"]
        if false {
            v7 = v7
        }
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 19; fi9++ {
        if false {
            v10 := GenS2{F0: 4294967295}
            v5 -= 2147483647
            var v11 int16 = -32768
            _ = v10
            _ = v11
        }
        _ = len(v3)
        fmt.Println(v5, v5)
        _ = fi9
    }
    fmt.Println(v5)
    v5 = (^(v5))
    if true {
        if (0 <= 42) {
            v5 += (v5 * (v5 >> ((427) & 63)))
            for fi12 := 0; fi12 < 8; fi12++ {
                v5 = 2147483647
                var v13 uint = 42
                _ = fi12
                _ = v13
            }
            fmt.Println(v5)
        }
        v5 -= 43
    } else {
        fmt.Println(v5)
    }
    v14 := GenS2{F0: 0}
    v5 += (v5 ^ v5)
    var v15 int32 = v5
    var v16 string = ("hello" + "a")
    v5 -= (v15 % ((v15) | 1))
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v15, v5, fi17)
        var v18 float64 = 0.1
        for fi19 := 0; fi19 < 8; fi19++ {
            switch "foo bar" {
            case "hello":
                if true {
                    _ = len(v2)
                    v5 = -1
                    var v20 uint16 = 0
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    var v21 int = -9223372036854775808
                    _ = v20
                    _ = v21
                }
            case "foo bar":
                switch (fi19 & fi19) {
                case -20:
                    var v22 uint8 = 128
                    _ = v22
                case 128:
                    v18 -= v18
                case 4294967295:
                    v23 := GenS3{F0: 42, F1: 42, F2: 152}
                    _ = v23
                }
            case "\n\t":
                v24 := v1[(("hello" + ("hello" + v16)) + "")]
                _ = v24
            }
            if len(v3) > 0 {
                v3[0] = v16
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    v15 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
