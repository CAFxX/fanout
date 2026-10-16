package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool, p1 bool, p2 int) uint32 {
    v5 := []int{}
    _ = len(v5)
    var v6 string = ("\n\t" + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 uint16 = 1
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int16 = -1
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            _ = len(v9)
            v1 -= (v1 + v1)
            var v12 uint16 = 468
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    _ = len(v9)
    v13 := GenS1{F0: -22, F1: 1, F2: 78}
    v4 = v4
    v14 := v3["\n\t"]
    var v15 bool = false
    v16 := GenS1{F0: 127, F1: 4294967295, F2: 4294967295}
    v17 := GenS1{F0: -92, F1: 4294967295, F2: 1}
    v14 = (v14 + v14)
    v18 := GenS1{F0: v1, F1: 2, F2: 4294967295}
    var v19 int16 = -32768
    switch v14 {
    case "\n\t":
        var v20 int8 = 46
        _ = v20
    case "hello":
        var v21 uint8 = 42
        _ = v21
    case "":
        switch 285 {
        case -366:
            if v15 {
                fmt.Println(v15, v14, v15)
            } else {
                delete(v2, ("foo bar" + (v14 + v14)))
            }
        case 4294967295:
            if (!v15) {
                for fi22 := 0; fi22 < 12; fi22++ {
                    if (0 > 0) {
                        v23 := GenS2{F0: (^(v4))}
                        _ = len(v3)
                        if true {
                            var v24 int = fi22
                            _ = v24
                        } else {
                            v25 := GenS1{F0: -128, F1: 42, F2: 42}
                            v26 := GenS1{F0: (v1 << ((42) & 63)), F1: 600, F2: 2}
                            var v27 uint8 = 127
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        _ = v23
                    } else {
                        var v28 uint64 = 731
                        for fi29 := 0; fi29 < 18; fi29++ {
                            switch v14 {
                            case "foo bar":
                                fmt.Println(v15, v19)
                            case "":
                                var v30 float64 = (1.5 - (3.14159 + (1.5 + 0.1)))
                                _ = v30
                            default:
                                fmt.Println(v1, v14)
                            }
                            v31 := GenS3{F0: 0}
                            _ = fi29
                            _ = v31
                        }
                        var v32 int64 = -9223372036854775808
                        _ = v28
                        _ = v32
                    }
                    if ((v15 || false) || v15) {
                        if false {
                            v33 := v2[((v14 + ((v14 + "") + v14)) + v14)]
                            var v34 string = v14
                            v14 = (v14 + "")
                            _ = v33
                            _ = v34
                        }
                        v35 := v2[""]
                        _ = v35
                    }
                    fmt.Println(fi22, fi22)
                    _ = fi22
                }
                for fi36 := 0; fi36 < 8; fi36++ {
                    fmt.Println(v19)
                    _ = fi36
                }
                if (v15 || v15) {
                    var v37 uint32 = 4294967295
                    fmt.Println(v19)
                    v19 += v19
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 18; fi38++ {
                        fmt.Println(v14, v19)
                        var v39 uint32 = 2147483648
                        _ = fi38
                        _ = v39
                    }
                }
            }
        }
    default:
        if len(v9) > 0 {
            v9[0] = 128
        }
    }
    fmt.Println(v19, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
