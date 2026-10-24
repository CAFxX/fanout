package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }

func main() {
    var v1 int = -262
    v2 := make(map[string]string)
    v3 := []uint32{2147483647, 2147483648, 65535}
    var v4 string = ""
    v5 := make(map[string]int)
    var v6 string = (("a" + v4) + "\n\t")
    for fi7 := 0; fi7 < 16; fi7++ {
        if (v4 > v6) {
            fmt.Println(v1, v4)
        }
        _ = fi7
    }
    var v8 int = (v1 | v1)
    switch v6 {
    case "":
        var v9 uint64 = 825
        _ = v9
    }
    var v10 uint = 2
    for fi11 := 0; fi11 < 4; fi11++ {
        switch v6 {
        case "foo bar":
            _ = len(v5)
        }
        if true {
            var v12 int8 = 127
            switch (v8 * fi11) {
            case 32767:
                var v13 bool = (float32((-2.25 - 3.14159)) <= 0.0)
                _ = v13
            case -1:
                v4 = ((v6 + v6) + "foo bar")
            }
            var v14 int64 = -945
            _ = v12
            _ = v14
        } else {
            if (0 < 1) {
                v4 = (v4 + v4)
            }
            var v15 int32 = 42
            _ = v15
        }
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        _ = fi11
    }
    v10 += 255
    if (!(true && false)) {
        v16 := GenS1{F0: 100}
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 11; fi18++ {
                v8 = 128
                v19 := GenS1{F0: -1}
                v20 := GenS1{F0: 42}
                _ = fi18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 9; fi21++ {
                if false {
                    switch -288 {
                    case 0:
                        v2[v6] = v6
                    case -599:
                        if (65535 <= 203) {
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                            v6 = v4
                        }
                    }
                    v22 := GenS1{F0: -68}
                    _ = v22
                }
                switch "\n\t" {
                case "a":
                    v23 := GenS2{F0: 0, F1: v10, F2: 191}
                    _ = v23
                case "hello":
                    fmt.Println(fi21, v8, v8)
                default:
                    if (!(true && false)) {
                        var v24 uint8 = 255
                        if true {
                            switch v4 {
                            case "hello":
                                fmt.Println(v8, v24, v10)
                            case "a":
                                v24 = v24
                            case "\n\t":
                                v6 = v6
                            }
                        } else {
                            var v25 int64 = 913
                            if false {
                                var v26 float64 = -2.25
                                v3 = append(v3, 449)
                                switch (("foo bar" + "\n\t") + v6) {
                                case "":
                                    if false {
                                        fmt.Println(fi21, fi21)
                                        var v27 int8 = int8(v25)
                                        v28 := GenS1{F0: (v27 % (((v27 - v27)) | 1))}
                                        _ = v27
                                        _ = v28
                                    }
                                default:
                                    fmt.Println(v6, fi21)
                                }
                                _ = v26
                            }
                            _ = v25
                        }
                        fmt.Println(fi17, v6)
                        fmt.Println(v1)
                        v29 := GenS1{F0: 42}
                        _ = v24
                        _ = v29
                    }
                }
                v4 = ((((("a" + ("" + (((v6 + v6) + (((v6 + (v6 + v4)) + "\n\t") + (v4 + ("a" + "")))) + v4))) + v6) + "\n\t") + (v6 + v6)) + v4)
                _ = fi21
            }
            _ = fi17
        }
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
