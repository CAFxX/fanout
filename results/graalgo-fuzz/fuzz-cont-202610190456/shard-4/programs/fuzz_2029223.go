package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint) int16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    var v7 int64 = 9223372036854775807
    p0 -= p0
    _ = p0
    _ = v7
    return 0
}

func main() {
    var v1 int32 = -543
    v2 := make(map[string]uint32)
    var v3 bool = false
    var v4 int8 = -128
    v8 := []int{32767}
    var v9 int8 = (-(v4))
    v4 += v9
    fmt.Println(v4)
    for fi10 := 0; fi10 < 15; fi10++ {
        switch "foo bar" {
        case "":
            if (v1 != (v1 << ((186) & 63))) {
                var v11 uint32 = 2147483648
                v12 := v2[("\n\t" + "hello")]
                _ = v11
                _ = v12
            } else {
                switch "hello" {
                case "foo bar":
                    delete(v2, "hello")
                case "":
                    v4 -= (v4 << ((1) & 63))
                default:
                    v13 := GenS1{F0: (v9 >> ((255) & 63)), F1: 42}
                    _ = v13
                }
                v14 := GenS1{F0: -83, F1: 409}
                _ = v14
            }
        case "\n\t":
            v1 -= (v1 / ((v1) | 1))
        case "hello":
            if (v4 >= (v4 + (v9 >> ((452) & 63)))) {
                var v15 uint = 42
                v16 := GenS1{F0: (v9 | (v4 + (v4 | 127))), F1: 903}
                _ = v15
                _ = v16
            } else {
                v9 = (v9 & v4)
            }
        }
        _ = fi10
    }
    delete(v2, ("" + ("" + "\n\t")))
    v17 := GenS1{F0: (v4 << ((4294967295) & 63)), F1: 65535}
    v9 -= 100
    switch "" {
    case "\n\t":
        _ = len(v8)
    case "a":
        if v3 {
            fmt.Println(v9, v4, v1)
            if (("a" + "a") < ((("\n\t" + "\n\t") + (("\n\t" + "hello") + "")) + "")) {
                var v18 uint = 109
                fmt.Println(v18, v9)
                _ = v18
            }
        } else {
            var v19 int = 32768
            v20 := GenS1{F0: v9, F1: 161}
            _ = v19
            _ = v20
        }
    case "foo bar":
        if v3 {
            if (v1 != v1) {
                v2["a"] = 4294967295
                v4 = (v9 % ((v4) | 1))
            } else {
                if len(v8) > 0 {
                    v8[0] = 32767
                }
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    if false {
                        switch -897 {
                        case -826:
                            if false {
                                if false {
                                    _ = len(v8)
                                    v8 = append(v8, 4294967295)
                                }
                                v21 := GenS1{F0: v4, F1: 65535}
                                _ = v21
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    v1 = (v1 & 42)
                                case "":
                                    delete(v2, "a")
                                }
                                var v22 uint32 = 2147483648
                                _ = v22
                            }
                        case 127:
                            for fi23 := 0; fi23 < 2; fi23++ {
                                for fi24 := 0; fi24 < 14; fi24++ {
                                    fmt.Println(v1, fi24, v1)
                                    var v25 int = (fi24 - fi23)
                                    var v26 string = "\n\t"
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi23
                            }
                        default:
                            if len(v8) > 0 {
                                v8[0] = -646
                            }
                        }
                    }
                case "":
                    var v27 uint = 65535
                    _ = v27
                default:
                    var v28 float32 = 3.14159
                    _ = v28
                }
                switch "hello" {
                case "\n\t":
                    if v3 {
                        v9 = (^(v4))
                    }
                case "foo bar":
                    if v3 {
                        if v3 {
                            var v29 uint16 = 32767
                            v30 := GenS2{F0: 297}
                            v8 = append(v8, 740)
                            _ = v29
                            _ = v30
                        }
                    }
                }
            }
        }
    }
    v2["foo bar"] = 0
    v31 := GenS2{F0: (v1 ^ (v1 * (v1 / (((v1 >> ((4294967295) & 63))) | 1))))}
    var v32 uint64 = 665
    v33 := GenS2{F0: v1}
    v1 += v1
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 uint32 = 186
        for fi36 := 0; fi36 < 1; fi36++ {
            delete(v2, "foo bar")
            _ = fi36
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v8
    _ = v9
}
