package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0() int8 {
    var v6 bool = false
    fmt.Println(v6)
    _ = v6
    return 127
}
func genFn1() int16 {
    var v7 float32 = -2.25
    var v8 int64 = -9223372036854775808
    v7 -= v7
    _ = v7
    _ = v8
    return -32768
}

func main() {
    var v1 int32 = -169
    v2 := []uint32{224, 44, 71}
    v3 := []string{"\n\t", "hello", "hello", "hello"}
    var v4 float32 = 1.5
    var v5 uint8 = 142
    switch "hello" {
    case "a":
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 float64 = 0.0
            v11 := GenS1{F0: 471, F1: 65535}
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: 222, F1: 1}
                if false {
                    var v14 uint16 = 741
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    case "\n\t":
        for fi15 := 0; fi15 < 13; fi15++ {
            if true {
                v5 = 0
                for fi16 := 0; fi16 < 8; fi16++ {
                    if true {
                        v4 = v4
                        v1 += -243
                        var v17 int8 = 127
                        _ = v17
                    } else {
                        var v18 int16 = -32768
                        v19 := GenS1{F0: 2, F1: 51}
                        _ = v18
                        _ = v19
                    }
                    switch (fi15 | 42) {
                    case -9223372036854775808:
                        switch -2147483648 {
                        case 4294967295:
                            if (1 > 356) {
                                v1 = (v1 ^ 867)
                                switch 9223372036854775807 {
                                case 1:
                                    var v20 uint16 = 463
                                    _ = v20
                                }
                                switch fi15 {
                                case 42:
                                    var v21 uint32 = 2147483648
                                    _ = v21
                                case -578:
                                    fmt.Println(fi15, fi15, fi15)
                                    if len(v3) > 0 {
                                        v3[0] = "hello"
                                    }
                                case -2147483648:
                                    v4 = 3.14159
                                }
                            }
                            var v22 uint = 65535
                            if false {
                                switch "foo bar" {
                                case "foo bar":
                                    var v23 uint8 = v5
                                    _ = v23
                                case "":
                                    var v24 float32 = v4
                                    _ = v24
                                case "\n\t":
                                    v1 = -583
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v22
                        default:
                            fmt.Println(v1)
                        }
                    case 9223372036854775807:
                        if false {
                            var v25 uint = 18446744073709551615
                            v1 -= v1
                            switch fi15 {
                            case 376:
                                var v26 int16 = 0
                                _ = v26
                            }
                            _ = v25
                        } else {
                            var v27 int16 = 42
                            if true {
                                v1 += (v1 ^ v1)
                                fmt.Println(v1)
                            }
                            _ = v27
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v4 = v4
                    case 680:
                        v5 = v5
                    }
                    _ = fi16
                }
            } else {
                var v28 int = (fi15 << ((1) & 63))
                _ = v28
            }
            v4 = v4
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi15
        }
    }
    v29 := GenS2{F0: v1, F1: 74}
    var v30 int64 = 556
    v4 = (0.0 - v4)
    v31 := GenS2{F0: 1, F1: 100}
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 bool = ((fi32 < (fi32 * fi32)) && false)
        switch fi32 {
        case -252:
            var v34 uint8 = (v5 + v5)
            _ = v34
        }
        _ = fi32
        _ = v33
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := GenS1{F0: 42, F1: 65535}
        var v37 float32 = 1.5
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 uint = 255
    v30 = int64(v5)
    var v39 int64 = (v30 * (v30 << ((889) & 63)))
    var v40 uint8 = v5
    v41 := GenS2{F0: v1, F1: genFn0()}
    v42 := GenS3{F0: (v39 % ((v39) | 1)), F1: (v1 - -1)}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v40 += v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
