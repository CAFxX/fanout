package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int }

func main() {
    v1 := []uint32{99, 65535}
    var v2 bool = true
    var v3 bool = (-1 == 42)
    v4 := make(map[string]string)
    var v5 int16 = 1
    v5 = v5
    delete(v4, (("foo bar" + "foo bar") + ("hello" + ("hello" + "\n\t"))))
    v6 := GenS1{F0: "foo bar"}
    switch ("hello" + "hello") {
    case "\n\t":
        v3 = (v2 || true)
    case "foo bar":
        v3 = false
    case "a":
        for fi7 := 0; fi7 < 3; fi7++ {
            if ("" == "foo bar") {
                v3 = (!v3)
            } else {
                if (!v2) {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        if (v5 <= 42) {
                            if (42 > 255) {
                                delete(v4, "foo bar")
                                v5 += v5
                                fmt.Println(v3, v2, fi7)
                            } else {
                                if ((fi8 <= 0) || false) {
                                    if false {
                                        var v9 uint32 = 0
                                        if len(v1) > 0 {
                                            v1[0] = 2147483648
                                        }
                                        if (!false) {
                                            fmt.Println(fi8, fi8, v2)
                                        }
                                        _ = v9
                                    }
                                    fmt.Println(fi8)
                                    v2 = true
                                    var v10 float32 = 1e10
                                    v5 -= v5
                                    _ = v10
                                } else {
                                    if (3.14159 < 3.14159) {
                                        var v11 int = fi7
                                        if v3 {
                                            switch fi8 {
                                            case 0:
                                                v3 = v2
                                            case 32768:
                                                v11 = (fi7 * 975)
                                            case 941:
                                                if v2 {
                                                    var v12 float32 = (float32(0.0) * 0.1)
                                                    v5 += v5
                                                    _ = v12
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                            var v13 int16 = (v5 >> ((2) & 63))
                                            _ = v13
                                        }
                                        _ = v11
                                    }
                                }
                                v14 := GenS1{F0: (("" + "foo bar") + ((((("a" + "\n\t") + "hello") + "a") + ("\n\t" + (("\n\t" + ("\n\t" + "a")) + ("a" + "a")))) + "\n\t"))}
                                _ = v14
                            }
                            if ("\n\t" < ("hello" + "foo bar")) {
                                v3 = true
                                var v15 uint64 = 2
                                v1 = append(v1, 65535)
                                var v16 int64 = -9223372036854775808
                                v2 = (int8(fi7) < -70)
                                _ = v15
                                _ = v16
                            }
                            var v17 uint16 = 512
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            _ = v17
                        }
                        var v18 uint8 = 42
                        var v19 uint32 = 4294967295
                        _ = fi8
                        _ = v18
                        _ = v19
                    }
                }
                v2 = false
                v5 = -72
            }
            v4["hello"] = "foo bar"
            var v20 bool = v3
            _ = fi7
            _ = v20
        }
    }
    var v21 bool = (v3 || true)
    var v22 bool = true
    switch ("a" + ("\n\t" + "foo bar")) {
    case "":
        v23 := GenS2{F0: 2147483647, F1: 204}
        _ = v23
    case "a":
        v2 = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
