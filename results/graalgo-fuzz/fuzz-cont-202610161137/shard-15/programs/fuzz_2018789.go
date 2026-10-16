package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int; F1 int32; F2 uint64 }

func main() {
    v1 := []uint32{4294967295, 1}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    v4 := []uint32{135, 1, 2147483647}
    _ = len(v1)
    if true {
        if ((2147483647 < -1) && true) {
            var v5 uint64 = 0
            var v6 int16 = 1
            if false {
                for fi7 := 0; fi7 < 17; fi7++ {
                    var v8 int32 = 2147483647
                    _ = fi7
                    _ = v8
                }
                var v9 uint32 = 2147483647
                _ = v9
            }
            _ = v5
            _ = v6
        } else {
            for fi10 := 0; fi10 < 3; fi10++ {
                switch (fi10 >> ((173) & 63)) {
                case -2147483648:
                    var v11 uint64 = 627
                    _ = v11
                }
                switch fi10 {
                case 1:
                    delete(v3, 65535)
                case 483:
                    for fi12 := 0; fi12 < 14; fi12++ {
                        var v13 string = ("a" + ("hello" + ""))
                        var v14 float64 = 1e10
                        var v15 int32 = -860
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                case -2147483648:
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 int64 = 9223372036854775807
                        var v18 uint8 = 241
                        var v19 uint32 = 131
                        var v20 int = 128
                        delete(v2, v19)
                        switch (fi16 % ((fi10) | 1)) {
                        case 42:
                            v20 += v20
                        case 4294967295:
                            if false {
                                switch "\n\t" {
                                case "a":
                                    v3[(v19 >> ((2) & 63))] = v20
                                case "":
                                    v20 = v20
                                    var v21 bool = true
                                    _ = v21
                                case "foo bar":
                                    v22 := GenS1{F0: 0}
                                    _ = v22
                                default:
                                    v1 = append(v1, v19)
                                }
                                var v23 int64 = 42
                                _ = v23
                            } else {
                                v24 := GenS1{F0: -128}
                                _ = v24
                            }
                        }
                        fmt.Println(fi16)
                        if true {
                            v25 := GenS2{F0: v18}
                            _ = v25
                        }
                        fmt.Println(fi16)
                        v26 := GenS3{F0: 460, F1: v17}
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v26
                    }
                default:
                    switch (fi10 % ((2147483647) | 1)) {
                    case 42:
                        _ = len(v4)
                    }
                }
                _ = len(v4)
                _ = fi10
            }
        }
        var v27 int16 = 0
        _ = v27
    }
    v28 := v3[469]
    _ = len(v3)
    var v29 int8 = 0
    _ = len(v1)
    if false {
        var v30 int = v28
        if ("\n\t" > "foo bar") {
            v31 := GenS2{F0: 128}
            _ = len(v1)
            _ = v31
        } else {
            if true {
                _ = len(v3)
                if (255 <= 65535) {
                    v32 := GenS1{F0: -128}
                    v28 += -417
                    _ = v32
                }
                var v33 uint32 = 4294967295
                _ = v33
            } else {
                v29 -= (v29 / ((1) | 1))
            }
        }
        _ = v30
    } else {
        var v34 int64 = 150
        _ = v34
    }
    v2[263] = v28
    for fi35 := 0; fi35 < 1; fi35++ {
        var v36 int32 = -105
        _ = fi35
        _ = v36
    }
    v29 -= (-(v29))
    var v37 uint32 = 65535
    if (1e10 > 3.14159) {
        v28 = v28
        v38 := GenS4{F0: v28, F1: -853, F2: 2}
        v39 := GenS3{F0: 42, F1: -1}
        _ = v38
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v4
}
