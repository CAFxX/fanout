package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint32; F2 int32 }
func genFn0(p0 uint8) uint16 {
    v5 := GenS1{F0: -1, F1: -566}
    var v6 uint16 = 764
    fmt.Println(v6, p0)
    var v7 int = 1
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = 1e10
    v3 := []string{"a", (("a" + (("\n\t" + "foo bar") + ("foo bar" + ("a" + "a")))) + ("\n\t" + ((("\n\t" + "a") + "foo bar") + ("a" + "hello")))), ""}
    v4 := make(map[uint32]uint32)
    var v8 int16 = 252
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint = 924
        if len(v3) > 0 {
            v3[0] = (("" + "foo bar") + ("a" + "foo bar"))
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            v8 -= -413
            if (!((true && (fi11 != fi9)) || ((false && (127 > (fi11 >> ((4294967295) & 63)))) || true))) {
                var v12 uint = v10
                v1 = v1
                v10 = v10
                _ = len(v4)
                _ = v12
            }
            v13 := GenS2{F0: 0}
            _ = fi11
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    var v14 uint16 = (v1 + v1)
    v14 += v1
    v15 := GenS3{F0: "hello", F1: 1, F2: int32(v14)}
    var v16 float64 = -2.25
    v17 := GenS1{F0: -967, F1: -219}
    var v18 int = 351
    var v19 int8 = 1
    fmt.Println(v14, v1)
    _ = len(v3)
    v8 = (v8 & 1000)
    fmt.Println(v1, v1)
    fmt.Println(v2)
    delete(v4, 2)
    delete(v4, 49)
    v20 := v4[449]
    var v21 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}
