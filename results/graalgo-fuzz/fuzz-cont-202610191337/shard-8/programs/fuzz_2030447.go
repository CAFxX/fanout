package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }

func main() {
    var v1 int16 = 1
    var v2 uint32 = 42
    v3 := make(map[string]string)
    v4 := GenS1{F0: 32767, F1: v1, F2: "a"}
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 float64 = (3.14159 / 1.0)
        v6 -= float64(3.14159)
        _ = fi5
        _ = v6
    }
    v2 = v2
    delete(v3, "foo bar")
    var v7 int = -2147483648
    v2 = 0
    switch (v7 - (v7 - v7)) {
    case 42:
        delete(v3, ("" + "foo bar"))
    default:
        delete(v3, ("foo bar" + ("" + "foo bar")))
    }
    fmt.Println(v1, v2, v7)
    v8 := []int{-1, v7}
    v2 = 0
    v9 := GenS1{F0: 159, F1: (v1 << ((534) & 63)), F2: "a"}
    if ((1.5 > 0.1) || true) {
        _ = len(v3)
        var v10 int32 = 1
        _ = v10
    } else {
        fmt.Println(v2)
    }
    switch "foo bar" {
    case "a":
        _ = len(v3)
    }
    v11 := GenS1{F0: (v1 / (((v1 & v1)) | 1)), F1: v1, F2: ("foo bar" + "")}
    switch v7 {
    case 773:
        v2 -= (v2 % (((v2 % ((691) | 1))) | 1))
    case -2147483648:
        v7 += -809
    default:
        if (0 != 32767) {
            var v12 float32 = 1e10
            v13 := GenS2{F0: -1, F1: ((4294967295 > 0) || false), F2: 54}
            _ = v12
            _ = v13
        }
    }
    var v14 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
