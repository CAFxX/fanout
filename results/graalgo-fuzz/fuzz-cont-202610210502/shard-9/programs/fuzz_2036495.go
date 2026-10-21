package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint; F2 uint16 }
func genFn0() int {
    v5 := GenS1{F0: 0, F1: (!(-1 > 412)), F2: 127}
    var v6 int16 = 84
    var v7 bool = (v6 >= (v6 - -32768))
    _ = v5
    _ = v6
    _ = v7
    return 272
}

func main() {
    var v1 int8 = 127
    v2 := make(map[string]int)
    var v3 float32 = 1.5
    v4 := make(map[uint32]string)
    v8 := []int{127}
    if len(v8) > 0 {
        v8[0] = 32767
    }
    v1 = (^(v1))
    if (!((-373 <= 42) && (false && true))) {
        for fi9 := 0; fi9 < 18; fi9++ {
            v3 = 0.0
            if (1 < 653) {
                var v10 int16 = 813
                var v11 float32 = v3
                _ = len(v8)
                _ = v10
                _ = v11
            } else {
                var v12 int32 = -184
                if (true && true) {
                    if (true || true) {
                        for fi13 := 0; fi13 < 10; fi13++ {
                            var v14 bool = true
                            if v14 {
                                v1 += (v1 % (((v1 & v1)) | 1))
                            }
                            if v14 {
                                fmt.Println(v12, fi13)
                                v15 := GenS2{F0: 25, F1: 4294967295, F2: 17}
                                _ = v15
                            } else {
                                var v16 string = "hello"
                                fmt.Println(v3, v1)
                                _ = v16
                            }
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        var v17 bool = true
                        _ = v17
                    }
                    v8 = append(v8, fi9)
                }
                var v18 int16 = -32768
                _ = v12
                _ = v18
            }
            _ = fi9
        }
        v19 := v2["hello"]
        var v20 uint16 = 1
        _ = v19
        _ = v20
    }
    var v21 int8 = -1
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 int16 = -903
        _ = len(v4)
        _ = fi22
        _ = v23
    }
    v3 = 3.14159
    if true {
        fmt.Println(v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}
