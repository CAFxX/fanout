package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int16 }
type GenS2 struct { F0 int; F1 string; F2 int32 }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    var v3 uint16 = 255
    var v4 uint = 709
    v5 := []uint32{841, v2, (v2 - (v2 << ((v4) & 63)))}
    for fi6 := 0; fi6 < 3; fi6++ {
        v4 = v4
        _ = fi6
    }
    v2 += (^(v2))
    fmt.Println(v4, v2, v2)
    v7 := GenS1{F0: uint8(v4), F1: 42, F2: 191}
    v3 = uint16(v4)
    v3 = (v3 * 2)
    fmt.Println(v4, v3, v4)
    v8 := GenS1{F0: 42, F1: 127, F2: 0}
    if (3.14159 == 0.0) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v9 := GenS2{F0: 32768, F1: ("a" + ("hello" + ("" + "a"))), F2: 2147483647}
        _ = v9
    } else {
        var v10 int64 = 42
        var v11 uint8 = 255
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        if false {
            var v13 bool = ((v3 | 988) != 0)
            var v14 int32 = 1
            _ = v13
            _ = v14
        } else {
            fmt.Println(v3, v4)
        }
        delete(v1, ("" + "\n\t"))
        if true {
            if false {
                var v15 bool = false
                _ = v15
            } else {
                if true {
                    v16 := GenS1{F0: 255, F1: 110, F2: 332}
                    switch (fi12 / (((^(fi12))) | 1)) {
                    case -2147483648:
                        v3 -= v3
                    case 1:
                        var v17 float64 = 3.14159
                        _ = v17
                    case -1:
                        if (false || (1e10 <= (3.14159 - 0.0))) {
                            v2 = (v2 ^ v2)
                            fmt.Println(fi12, fi12, v2)
                        }
                    }
                    var v18 int16 = 584
                    _ = v16
                    _ = v18
                } else {
                    v2 = 680
                    v19 := GenS1{F0: 189, F1: -128, F2: 0}
                    v3 = v3
                    _ = v19
                }
                var v20 uint16 = v3
                _ = v20
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                v4 += v4
                fmt.Println(v3, fi12, v2)
                _ = fi21
            }
        } else {
            if (!false) {
                var v22 float32 = ((1.5 - ((0.0 / 1.0) + 1e10)) - -2.25)
                _ = v22
            } else {
                fmt.Println(v2)
                v4 -= v4
                if true {
                    if false {
                        fmt.Println(v3)
                        v3 += v3
                    }
                    for fi23 := 0; fi23 < 2; fi23++ {
                        var v24 int = 128
                        v25 := GenS3{F0: -223, F1: (v4 ^ (v4 ^ (v4 >> ((v4) & 63))))}
                        var v26 bool = (false && (-2147483648 < 0))
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    var v27 int16 = 1
                    _ = v27
                }
            }
            if false {
                for fi28 := 0; fi28 < 11; fi28++ {
                    v2 -= (^(v2))
                    _ = fi28
                }
            }
        }
        _ = fi12
    }
    var v29 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
