package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 272
    var v3 uint16 = 32767
    var v4 float64 = (3.14159 / -2.25)
    v5 := GenS1{F0: 659, F1: 893, F2: (v3 - (v3 & v3))}
    v6 := []int{2147483647, -2147483648, 678}
    v3 += v2
    fmt.Println(v3)
    v7 := GenS2{F0: 65535, F1: 42}
    if true {
        v8 := GenS1{F0: -207, F1: 0, F2: v2}
        v9 := GenS3{F0: -55}
        fmt.Println(v3)
        _ = v8
        _ = v9
    }
    v1[("foo bar" + "foo bar")] = 0
    v2 -= v3
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint8 = 23
        if true {
            delete(v1, ("\n\t" + "a"))
            fmt.Println(fi10, fi10, v3)
            if (1.5 >= -2.25) {
                if (true || (!(0 > 204))) {
                    if false {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            v2 -= 275
                            var v13 int64 = int64(v11)
                            _ = fi12
                            _ = v13
                        }
                        var v14 float32 = (1e10 - 1e10)
                        delete(v1, "foo bar")
                        _ = v14
                    } else {
                        var v15 float64 = v4
                        v16 := GenS2{F0: 4294967295, F1: 255}
                        _ = v15
                        _ = v16
                    }
                    v4 -= v4
                    var v17 uint = 65535
                    _ = v17
                } else {
                    if (("foo bar" + ("" + "foo bar")) < "\n\t") {
                        v18 := GenS1{F0: 32768, F1: 218, F2: 65535}
                        _ = v18
                    }
                    var v19 int = fi10
                    _ = v19
                }
                v20 := GenS4{F0: 10}
                _ = v20
            }
        } else {
            v2 = 255
            v21 := v1["a"]
            if false {
                v3 = (v2 << ((2) & 63))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = v21
        }
        _ = fi10
        _ = v11
    }
    v2 -= 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
