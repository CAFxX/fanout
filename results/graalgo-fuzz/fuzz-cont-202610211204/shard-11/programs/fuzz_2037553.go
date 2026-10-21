package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    var v1 uint = 255
    var v2 string = ((("hello" + "a") + ("\n\t" + "")) + "\n\t")
    var v3 uint16 = 492
    v4 := make(map[uint32]string)
    var v5 int16 = -583
    v1 += v1
    fmt.Println(v5)
    switch 4294967295 {
    case -375:
        v6 := v4[2]
        _ = v6
    }
    v2 = v2
    fmt.Println(v1, v1, v2)
    v7 := []int{}
    var v8 float64 = 1.5
    delete(v4, 453)
    _ = len(v7)
    switch (((v2 + v2) + "a") + (((v2 + v2) + v2) + (v2 + (("" + "foo bar") + ("a" + v2))))) {
    case "\n\t":
        v8 += v8
    case "foo bar":
        var v9 float32 = 0.1
        _ = v9
    default:
        if (true || false) {
            var v10 uint = v1
            v11 := GenS1{F0: (v2 + v2), F1: 265}
            v1 += 1
            _ = v10
            _ = v11
        } else {
            v1 -= (v1 / (((v1 & v1)) | 1))
            if (v5 <= 32767) {
                var v12 bool = true
                var v13 uint16 = v3
                var v14 int = 32768
                _ = v12
                _ = v13
                _ = v14
            } else {
                if (!true) {
                    fmt.Println(v5, v8)
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
                var v15 string = ""
                _ = v15
            }
            var v16 uint64 = 1
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint = 1
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        if true {
            v2 = v2
        } else {
            if len(v7) > 0 {
                v7[0] = 32767
            }
        }
        _ = fi19
    }
    var v20 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
