package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 bool = true
    var v2 int8 = 100
    v3 := []uint32{716, 0, 2}
    v4 := make(map[uint32]uint32)
    var v5 float64 = 0.0
    var v6 uint32 = 42
    if v1 {
        if v1 {
            var v7 uint64 = 0
            _ = v7
        } else {
            var v8 int32 = 440
            if len(v3) > 0 {
                v3[0] = v6
            }
            for fi9 := 0; fi9 < 7; fi9++ {
                if (-1 > v8) {
                    v6 = v6
                    v10 := GenS1{F0: (v2 % ((127) | 1)), F1: 18}
                    var v11 string = "hello"
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint32 = (v6 | (v6 % ((432) | 1)))
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v14 := v4[961]
                        v12 = (^(v14))
                        _ = fi13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 10; fi15++ {
                        if true {
                            switch fi15 {
                            case -718:
                                delete(v4, (v12 & v6))
                            }
                        } else {
                            var v16 int16 = 32767
                            var v17 uint16 = 1
                            fmt.Println(v5, v12)
                            _ = v16
                            _ = v17
                        }
                        var v18 int64 = 21
                        v19 := GenS1{F0: (v2 / ((1) | 1)), F1: 255}
                        _ = fi15
                        _ = v18
                        _ = v19
                    }
                    _ = v12
                }
                var v20 float32 = 0.1
                if false {
                    var v21 int = 32767
                    var v22 uint16 = 557
                    _ = v21
                    _ = v22
                }
                _ = fi9
                _ = v20
            }
            _ = v8
        }
        v6 -= (v6 ^ (^(v6)))
        v5 -= v5
    } else {
        for fi23 := 0; fi23 < 7; fi23++ {
            if false {
                for fi24 := 0; fi24 < 3; fi24++ {
                    switch "foo bar" {
                    case "foo bar":
                        var v25 bool = (-1 < -454)
                        _ = v25
                    case "a":
                        if true {
                            var v26 int = (fi23 & fi24)
                            _ = v26
                        } else {
                            if v1 {
                                var v27 uint64 = 356
                                _ = v27
                            }
                        }
                    }
                    _ = len(v3)
                    v6 += 42
                    _ = fi24
                }
                var v28 int64 = 20
                _ = v28
            } else {
                v2 = (v2 & -63)
            }
            var v29 string = "\n\t"
            _ = fi23
            _ = v29
        }
    }
    if ((v6 << ((65535) & 63)) <= v6) {
        var v30 int8 = -128
        _ = v30
    } else {
        if true {
            var v31 uint32 = (v6 / (((v6 ^ v6)) | 1))
            _ = v31
        }
        fmt.Println(v6, v2, v2)
    }
    var v32 int = 546
    if false {
        var v33 uint = 255
        fmt.Println(v1, v33, v32)
        v33 = 18446744073709551615
        _ = v33
    } else {
        if false {
            _ = len(v3)
            for fi34 := 0; fi34 < 16; fi34++ {
                var v35 uint32 = (^(v6))
                for fi36 := 0; fi36 < 5; fi36++ {
                    var v37 int64 = 893
                    v38 := GenS1{F0: int8(v6), F1: 0}
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = fi34
                _ = v35
            }
            var v39 uint32 = (v6 << ((42) & 63))
            _ = v39
        } else {
            var v40 int = v32
            _ = v40
        }
        v41 := GenS1{F0: (v2 << ((1) & 63)), F1: 212}
        _ = v41
    }
    switch (v32 << ((255) & 63)) {
    case 127:
        v5 += v5
    }
    if v1 {
        var v42 float64 = (3.14159 * 1.5)
        var v43 int8 = v2
        if len(v3) > 0 {
            v3[0] = v6
        }
        _ = v42
        _ = v43
    }
    v44 := GenS1{F0: (v2 >> ((0) & 63)), F1: 11}
    v2 = (v2 | 127)
    var v45 int32 = -861
    v46 := GenS2{F0: 1}
    fmt.Println(v6, v32, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
