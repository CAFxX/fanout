package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 uint32, p1 int32) string {
    if false {
        switch ("hello" + "\n\t") {
        case "":
            v4 := []int{32767, 128}
            _ = v4
        case "hello":
            var v5 string = "hello"
            _ = v5
        }
        if (0 == 42) {
            p1 -= p1
        }
        if (false && false) {
            v6 := GenS1{F0: 9223372036854775807, F1: -436}
            fmt.Println(p1, p1)
            if false {
                fmt.Println(p0, p1)
            }
            _ = v6
        }
    }
    v7 := []int{853}
    _ = p0
    _ = p1
    _ = v7
    return "\n\t"
}

func main() {
    var v1 bool = true
    var v2 string = ""
    var v3 uint64 = 1
    for fi8 := 0; fi8 < 8; fi8++ {
        v1 = v1
        var v9 uint = 255
        _ = fi8
        _ = v9
    }
    var v10 bool = v1
    v11 := []int{967, 32768}
    var v12 uint = 255
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS1{F0: 42, F1: -32768}
        _ = fi13
        _ = v14
    }
    var v15 int64 = 0
    for fi16 := 0; fi16 < 9; fi16++ {
        if true {
            if len(v11) > 0 {
                v11[0] = fi16
            }
            v12 -= 394
            var v17 string = (v2 + v2)
            _ = v17
        } else {
            for fi18 := 0; fi18 < 4; fi18++ {
                if v10 {
                    if (4294967295 != uint64(v12)) {
                        fmt.Println(v2)
                        if (128 == 194) {
                            switch v2 {
                            case "a":
                                fmt.Println(v12, fi18, v12)
                            case "hello":
                                v15 -= v15
                            case "foo bar":
                                var v19 float32 = 1e10
                                _ = v19
                            }
                            if len(v11) > 0 {
                                v11[0] = fi18
                            }
                        }
                        v10 = v10
                    }
                    var v20 int32 = 2147483647
                    _ = v20
                }
                var v21 int = fi18
                switch ((v2 + v2) + ((("a" + (v2 + v2)) + "") + (v2 + ""))) {
                case "a":
                    v11 = append(v11, v21)
                }
                switch "" {
                case "hello":
                    var v22 int64 = v15
                    _ = v22
                case "\n\t":
                    var v23 bool = v10
                    _ = v23
                default:
                    var v24 string = ""
                    _ = v24
                }
                v21 = (-(fi16))
                _ = fi18
                _ = v21
            }
            v25 := GenS1{F0: v15, F1: 32767}
            _ = v25
        }
        if v10 {
            fmt.Println(v3, v3, v15)
        }
        _ = fi16
    }
    var v26 float32 = (((float32(1.5) * 1.5) - -2.25) * 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v26
    _ = v3
}
