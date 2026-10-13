package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 int8, p2 int32) uint32 {
    var v5 int32 = p2
    var v6 int64 = -9223372036854775808
    var v7 float64 = float64(1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    v3 := []string{"hello", ("a" + ("a" + ("" + (("foo bar" + "") + "\n\t")))), "a"}
    v4 := make(map[string]bool)
    if (-1 == -128) {
        var v8 int8 = 76
        _ = v8
    }
    var v9 float64 = float64(float32(-2.25))
    fmt.Println(v9)
    fmt.Println(v9)
    switch "a" {
    case "hello":
        v1["\n\t"] = "hello"
    default:
        v2[42] = true
    }
    v9 = (v9 * 0.1)
    v9 = v9
    switch (("hello" + "") + "") {
    case "foo bar":
        v9 -= 1.5
    case "a":
        if (true && true) {
            if (("foo bar" + "a") < "foo bar") {
                if true {
                    var v10 float32 = (1e10 * 0.1)
                    if true {
                        var v11 uint32 = genFn0(-906, 127, 0)
                        for fi12 := 0; fi12 < 5; fi12++ {
                            v13 := v1["hello"]
                            v3 = append(v3, "\n\t")
                            _ = fi12
                            _ = v13
                        }
                        fmt.Println(v10, v11)
                        _ = v11
                    }
                    _ = v10
                }
                if (!(((32767 >= -1) || (!(613 > 1))) && (!(-2147483648 != 1)))) {
                    var v14 float64 = (1e10 + ((0.0 * (v9 / float64(0.1))) - 1e10))
                    v15 := GenS1{F0: 4294967295, F1: genFn0(2147483647, -128, 0), F2: 931}
                    if false {
                        for fi16 := 0; fi16 < 3; fi16++ {
                            fmt.Println(fi16, v14)
                            var v17 int = fi16
                            fmt.Println(v17, v14, v17)
                            _ = fi16
                            _ = v17
                        }
                    }
                    _ = v14
                    _ = v15
                } else {
                    fmt.Println(v9)
                }
            } else {
                v18 := v4["hello"]
                v9 += v9
                _ = v18
            }
            switch "hello" {
            case "\n\t":
                switch ("\n\t" + ("a" + "")) {
                case "\n\t":
                    if (1 >= 42) {
                        v9 -= v9
                        fmt.Println(v9)
                        var v19 int64 = -790
                        _ = v19
                    } else {
                        switch "hello" {
                        case "\n\t":
                            var v20 uint8 = 0
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 12; fi21++ {
                            v22 := GenS2{F0: 59}
                            for fi23 := 0; fi23 < 19; fi23++ {
                                var v24 int8 = 127
                                _ = fi23
                                _ = v24
                            }
                            v9 = (((v9 * v9) - float64((0.1 + -2.25))) * 1e10)
                            _ = fi21
                            _ = v22
                        }
                    }
                case "hello":
                    var v25 int = -761
                    _ = v25
                case "foo bar":
                    delete(v1, "hello")
                }
            default:
                var v26 uint16 = 32767
                _ = v26
            }
            if false {
                var v27 int32 = 2147483647
                v27 = (v27 % (((v27 & (v27 + v27))) | 1))
                v9 = -2.25
                _ = v27
            }
        }
    default:
        v9 = (float64(-2.25) * v9)
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 bool = (160 == 127)
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
