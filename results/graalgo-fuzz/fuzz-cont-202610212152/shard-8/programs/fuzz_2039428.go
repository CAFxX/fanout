package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
func genFn0() int8 {
    var v7 int8 = 42
    var v8 uint = 1
    v9 := GenS1{F0: 722}
    v8 = v8
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    v1 := []int{182, 1}
    var v2 int = -231
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v5 bool = (-2.25 > 0.0)
    var v6 uint = 787
    v10 := GenS1{F0: 4294967295}
    v2 -= (v2 & 42)
    var v11 int = 127
    v11 = v11
    v5 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ((("hello" + "foo bar") + ((("foo bar" + "hello") + "") + ("" + ""))) + "hello") {
    case "":
        var v12 uint = (^(v6))
        _ = v12
    case "foo bar":
        if (421 != 1) {
            v6 += (v6 | 42)
            _ = len(v1)
            switch "" {
            case "\n\t":
                if v5 {
                    v4[("\n\t" + ((("foo bar" + "") + "a") + "hello"))] = "foo bar"
                    if (!true) {
                        var v13 int32 = 12
                        var v14 uint8 = 128
                        var v15 uint32 = 2
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        var v16 bool = (!v5)
                        v1 = append(v1, v11)
                        for fi17 := 0; fi17 < 16; fi17++ {
                            v18 := v3[(fi17 + v11)]
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    }
                }
            }
        }
    }
    var v19 int32 = 2147483647
    v5 = (v5 || (true || (255 < 32767)))
    if ((!(-247 > v19)) && (false && true)) {
        switch (("foo bar" + ("hello" + "foo bar")) + "\n\t") {
        case "foo bar":
            fmt.Println(v2, v11, v2)
        case "":
            switch "\n\t" {
            case "a":
                var v20 string = ("a" + "\n\t")
                _ = v20
            case "":
                if len(v1) > 0 {
                    v1[0] = v11
                }
            default:
                v21 := GenS2{F0: "\n\t"}
                _ = v21
            }
        }
        var v22 int = v2
        if v5 {
            switch v2 {
            case 356:
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 uint16 = 0
                    _ = fi23
                    _ = v24
                }
            case -1:
                var v25 float64 = float64(0.1)
                _ = v25
            case 127:
                v2 = (-(v2))
            default:
                if (1 >= 42) {
                    fmt.Println(v22, v2, v19)
                } else {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v6 = v6
                        _ = fi26
                    }
                }
            }
        } else {
            v27 := GenS1{F0: 1}
            fmt.Println(v5)
            _ = len(v3)
            _ = v27
        }
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
