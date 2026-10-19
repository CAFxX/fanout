package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }

func main() {
    v1 := []uint32{485}
    v2 := []int{1, 61, -332, 1}
    v3 := make(map[uint32]int)
    v4 := []int{-1}
    var v5 bool = false
    fmt.Println(v5)
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi7
        }
        _ = fi6
    }
    var v8 float32 = 3.14159
    var v9 float64 = 0.1
    v10 := GenS1{F0: 65535, F1: -1, F2: 2}
    var v11 uint8 = 255
    if (1 <= 87) {
        fmt.Println(v5, v11)
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        switch "foo bar" {
        case "":
            v11 = 255
        case "hello":
            var v13 float64 = 3.14159
            _ = v13
        }
        _ = fi12
    }
    var v14 int16 = 32767
    var v15 int32 = 42
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint8 = (v11 ^ v11)
        var v18 int64 = 178
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 string = ""
    var v20 int = -9223372036854775808
    if v5 {
        for fi21 := 0; fi21 < 10; fi21++ {
            delete(v3, 1000)
            _ = fi21
        }
        switch (v20 ^ v20) {
        case -193:
            var v22 int8 = -6
            _ = v22
        case 128:
            var v23 uint8 = v11
            _ = v23
        case 4294967295:
            if (((((v19 + v19) + "foo bar") + "") + (v19 + "foo bar")) >= (v19 + "hello")) {
                if true {
                    v24 := GenS1{F0: 0, F1: v15, F2: v11}
                    v5 = v5
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        for fi26 := 0; fi26 < 17; fi26++ {
                            if false {
                                v4 = append(v4, (v20 + -9223372036854775808))
                                v14 = 42
                            } else {
                                var v27 int64 = -1
                                _ = v27
                            }
                            if len(v4) > 0 {
                                v4[0] = 32768
                            }
                            fmt.Println(v9, fi26)
                            _ = fi26
                        }
                        _ = fi25
                    }
                    var v28 uint16 = 42
                    var v29 int64 = -629
                    _ = v28
                    _ = v29
                }
                v30 := GenS1{F0: 0, F1: v15, F2: 42}
                _ = v30
            } else {
                if v5 {
                    if v5 {
                        if false {
                            v19 = v19
                            v14 -= v14
                        } else {
                            if false {
                                v5 = v5
                                v19 = ("a" + v19)
                            } else {
                                var v31 uint16 = 656
                                _ = v31
                            }
                        }
                        if v5 {
                            for fi32 := 0; fi32 < 6; fi32++ {
                                if (0.0 <= v8) {
                                    v1 = append(v1, 2)
                                    v15 -= v15
                                    var v33 string = (v19 + v19)
                                    _ = v33
                                }
                                _ = fi32
                            }
                            v34 := v3[635]
                            _ = v34
                        } else {
                            var v35 int16 = 32767
                            for fi36 := 0; fi36 < 4; fi36++ {
                                for fi37 := 0; fi37 < 14; fi37++ {
                                    v38 := GenS1{F0: 688, F1: (^(v15)), F2: v11}
                                    _ = fi37
                                    _ = v38
                                }
                                v11 = 2
                                _ = fi36
                            }
                            var v39 uint32 = 4294967295
                            _ = v35
                            _ = v39
                        }
                    }
                }
                fmt.Println(v20, v9)
                v40 := v3[4294967295]
                _ = v40
            }
        default:
            for fi41 := 0; fi41 < 4; fi41++ {
                v19 = v19
                _ = len(v4)
                _ = fi41
            }
        }
        v42 := GenS1{F0: 2, F1: (v15 * (v15 % ((v15) | 1))), F2: (v11 % (((v11 + 255)) | 1))}
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
