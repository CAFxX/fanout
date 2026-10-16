package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    v1 := []int{9223372036854775807}
    var v2 int32 = -194
    v3 := make(map[string]uint32)
    v4 := []string{""}
    var v5 uint = 4294967295
    var v6 int32 = v2
    if (1 > 1) {
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(v2, v6, fi7)
            var v8 uint64 = 593
            var v9 int64 = 9223372036854775807
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v6 -= v6
    var v10 float64 = (0.0 + (float64(float32((-2.25 - float64(1e10)))) * 0.0))
    switch -21 {
    case 617:
        for fi11 := 0; fi11 < 20; fi11++ {
            if ((v6 << (((v5 % (((v5 + v5)) | 1))) & 63)) < (v6 | v2)) {
                if (-128 <= -1) {
                    var v12 string = "a"
                    if false {
                        if true {
                            switch (fi11 | fi11) {
                            case 1:
                                v2 = int32(fi11)
                            }
                            v5 = 745
                        }
                    } else {
                        for fi13 := 0; fi13 < 13; fi13++ {
                            var v14 uint16 = 32767
                            var v15 string = v12
                            fmt.Println(v6, v6, fi13)
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        fmt.Println(v2, v10, v5)
                    }
                    _ = v12
                } else {
                    v2 -= (v2 + (v6 ^ (^(v6))))
                    v3[((("a" + ("\n\t" + "\n\t")) + "foo bar") + "hello")] = 42
                    var v16 uint32 = 2
                    _ = v16
                }
            } else {
                fmt.Println(v6)
                var v17 uint16 = 636
                var v18 int8 = 127
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: 91, F1: ((128 >= 128) && true)}
            _ = fi11
            _ = v19
        }
    }
    if false {
        var v20 uint8 = 127
        v21 := GenS1{F0: 1, F1: true}
        v22 := GenS1{F0: 158, F1: false}
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 int = 2147483647
    v23 -= -804
    v5 = (^(v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
