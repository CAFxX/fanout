package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int32; F2 int32 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int64; F1 string }
func genFn0(p0 int8) uint {
    p0 = p0
    p0 -= 57
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 uint8 = 1
    v5 = 142
    v6 := GenS1{F0: 42}
    var v7 float32 = float32((0.0 * float64((1.5 + 1.5))))
    v7 += v7
    var v8 bool = ("foo bar" == "a")
    var v9 int64 = -9223372036854775808
    var v10 int32 = 276
    if (v8 || v8) {
        var v11 int = 128
        _ = v11
    } else {
        var v12 float32 = v7
        if v8 {
            fmt.Println(v7, v8)
            if true {
                var v13 uint16 = 927
                if true {
                    var v14 string = ("foo bar" + ("a" + "foo bar"))
                    fmt.Println(v7, v14, v14)
                    _ = v14
                }
                v15 := GenS2{F0: -293, F1: (v10 & int32(v13)), F2: (v10 & (^(v10)))}
                _ = v13
                _ = v15
            } else {
                delete(v2, 634)
                var v16 float32 = float32((-2.25 - 1.5))
                v10 += v10
                _ = v16
            }
            v3["foo bar"] = int(v9)
        } else {
            var v17 int16 = 744
            v18 := GenS3{F0: 468, F1: v10}
            var v19 int32 = v10
            _ = v17
            _ = v18
            _ = v19
        }
        _ = v12
    }
    var v20 bool = (38 != 42)
    var v21 uint = 2
    var v22 uint64 = 232
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if v8 {
        for fi23 := 0; fi23 < 1; fi23++ {
            fmt.Println(fi23, v10)
            _ = fi23
        }
        if false {
            v21 = genFn0(120)
            v1 = append(v1, "hello")
            fmt.Println(v7, v20, v20)
        }
        for fi24 := 0; fi24 < 16; fi24++ {
            v8 = ((!true) || ("foo bar" > ("a" + "")))
            if (false && (!true)) {
                if (("foo bar" + ("hello" + ("a" + "a"))) == "") {
                    v7 += (1.5 - v7)
                }
                for fi25 := 0; fi25 < 13; fi25++ {
                    switch 2147483647 {
                    case 42:
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    case -91:
                        v26 := GenS4{F0: (-(v9)), F1: ((("a" + "a") + ("a" + "a")) + (("foo bar" + "\n\t") + ""))}
                        var v27 int = (fi24 + fi24)
                        if (true || (fi24 < v27)) {
                            v3["foo bar"] = int(v10)
                        }
                        _ = v26
                        _ = v27
                    }
                    _ = fi25
                }
                v28 := v3[("\n\t" + "a")]
                _ = v28
            } else {
                v29 := GenS3{F0: 65535, F1: int32(fi24)}
                var v30 int64 = v9
                v21 = (v21 ^ (v21 << ((genFn0(-1)) & 63)))
                _ = v29
                _ = v30
            }
            v8 = false
            _ = fi24
        }
    }
    switch "a" {
    case "a":
        var v31 uint64 = 21
        _ = v31
    case "foo bar":
        v20 = (42 > -128)
    case "\n\t":
        var v32 int = -946
        _ = v32
    default:
        fmt.Println(v21, v8)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
