package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 bool }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0() uint16 {
    var v6 float32 = (1.5 * 1e10)
    fmt.Println(v6)
    v7 := []int{32767}
    _ = v6
    _ = v7
    return 255
}
func genFn1() bool {
    var v8 float64 = (0.0 + 0.1)
    v9 := []int{32767, 32767}
    _ = len(v9)
    var v10 int64 = -1
    _ = v10
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{(("a" + "hello") + "a"), "\n\t", "foo bar", "\n\t"}
    var v3 uint16 = 966
    var v4 int16 = 32767
    var v5 float32 = 3.14159
    fmt.Println(v4, v5, v4)
    if (false || (!((-787 < -257) && (680 != 18446744073709551615)))) {
        if (235 != 156) {
            var v11 int = 4294967295
            var v12 string = ("hello" + "")
            _ = v11
            _ = v12
        } else {
            switch "\n\t" {
            case "foo bar":
                _ = len(v2)
            default:
                v4 = v4
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                if ((3.14159 + (float64(0.0) * 0.0)) <= 1.5) {
                    if true {
                        v14 := v1["\n\t"]
                        if len(v2) > 0 {
                            v2[0] = ((((("" + (("\n\t" + "foo bar") + "foo bar")) + "\n\t") + "hello") + "foo bar") + ((("hello" + "") + ("a" + ("\n\t" + "hello"))) + "\n\t"))
                        }
                        for fi15 := 0; fi15 < 6; fi15++ {
                            v16 := GenS1{F0: 100, F1: 2, F2: v14}
                            v5 += v5
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    } else {
                        v17 := GenS2{F0: 42, F1: 0}
                        var v18 uint64 = 2
                        _ = v17
                        _ = v18
                    }
                }
                _ = fi13
            }
        }
        var v19 float32 = v5
        fmt.Println(v4)
        _ = v19
    } else {
        v5 += v5
        v5 = (v5 - 1.5)
    }
    v4 -= (v4 - (v4 - (^(v4))))
    var v20 string = ""
    var v21 uint32 = 147
    if true {
        var v22 int32 = -703
        for fi23 := 0; fi23 < 19; fi23++ {
            fmt.Println(v5, v20)
            v24 := GenS2{F0: 2147483648, F1: 0}
            v3 -= v3
            _ = fi23
            _ = v24
        }
        v25 := GenS3{F0: v21, F1: 161}
        _ = v22
        _ = v25
    }
    v21 = v21
    fmt.Println(v5, v3, v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
