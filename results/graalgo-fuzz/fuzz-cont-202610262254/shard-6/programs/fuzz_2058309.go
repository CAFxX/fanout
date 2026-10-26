package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string; F1 string }

func main() {
    var v1 float32 = 0.0
    var v2 float32 = v1
    v3 := []int{}
    var v4 bool = true
    var v5 int32 = 690
    var v6 uint64 = 873
    if v4 {
        if (!(-118 != 127)) {
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: 32767, F1: 0}
                _ = fi7
                _ = v8
            }
            v5 = v5
            v5 += v5
        }
    }
    var v9 uint = 255
    var v10 float64 = 3.14159
    v9 = 868
    for fi11 := 0; fi11 < 13; fi11++ {
        if (42 != 1) {
            v1 += v2
            var v12 uint = 229
            _ = v12
        } else {
            for fi13 := 0; fi13 < 9; fi13++ {
                v6 = (v6 ^ v6)
                _ = fi13
            }
        }
        _ = fi11
    }
    var v14 string = "\n\t"
    switch v14 {
    case "foo bar":
        var v15 float32 = ((v1 + 1.5) / ((1e10 * (1e10 + float32((0.1 * 3.14159)))) * v2))
        _ = v15
    default:
        if ((v5 * v5) >= (v5 << ((v9) & 63))) {
            v16 := GenS1{F0: 0, F1: 176}
            var v17 uint64 = v6
            _ = v16
            _ = v17
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if false {
                fmt.Println(v14, v4)
                v18 := GenS2{F0: v14, F1: v14}
                v4 = v4
                _ = v18
            }
        }
    }
    v14 = v14
    v1 += (v1 + v2)
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS1{F0: int16(fi19), F1: 2147483648}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (v5 > int32(v9)) {
            v21 := GenS1{F0: -788, F1: 2}
            _ = len(v3)
            if true {
                for fi22 := 0; fi22 < 17; fi22++ {
                    if true {
                        switch -1 {
                        case 32768:
                            v14 = v14
                        case -646:
                            var v23 int64 = 1
                            _ = v23
                        }
                        var v24 float64 = 0.0
                        switch ("" + v14) {
                        case "hello":
                            v10 += -2.25
                        default:
                            v24 -= v10
                        }
                        _ = v24
                    } else {
                        var v25 string = v14
                        v2 += float32(1.5)
                        if (v6 >= v6) {
                            _ = len(v3)
                            var v26 uint64 = v6
                            var v27 int16 = 0
                            _ = v26
                            _ = v27
                        }
                        _ = v25
                    }
                    v6 = v6
                    var v28 float64 = (v10 - float64(v2))
                    _ = fi22
                    _ = v28
                }
                for fi29 := 0; fi29 < 15; fi29++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v30 uint8 = 255
                    switch fi29 {
                    case 32768:
                        var v31 uint64 = 42
                        _ = v31
                    default:
                        var v32 uint = v9
                        _ = v32
                    }
                    _ = fi29
                    _ = v30
                }
                v33 := GenS1{F0: 32767, F1: 4294967295}
                _ = v33
            }
            _ = v21
        } else {
            v4 = (v10 > (float64(v2) - v10))
            switch v14 {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "hello":
                if v4 {
                    v1 = (0.1 / v2)
                    switch (fi19 * 9223372036854775807) {
                    case -248:
                        var v34 uint8 = 2
                        _ = v34
                    case -841:
                        for fi35 := 0; fi35 < 14; fi35++ {
                            v36 := GenS1{F0: -32768, F1: 2147483648}
                            v37 := GenS1{F0: 0, F1: 2147483648}
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v38 := GenS2{F0: v14, F1: v14}
                    _ = v38
                }
            default:
                var v39 int8 = 0
                _ = v39
            }
        }
        _ = fi19
        _ = v20
    }
    if (20 >= 37) {
        var v40 bool = (!((v9 << ((v9) & 63)) >= v9))
        var v41 uint32 = 2147483647
        _ = v40
        _ = v41
    }
    if true {
        var v42 int32 = 42
        fmt.Println(v4)
        var v43 int64 = 0
        _ = v42
        _ = v43
    } else {
        var v44 string = ((v14 + v14) + "hello")
        fmt.Println(v1)
        for fi45 := 0; fi45 < 9; fi45++ {
            v46 := GenS1{F0: 1, F1: 316}
            _ = fi45
            _ = v46
        }
        _ = v44
    }
    var v47 uint = v9
    var v48 int16 = 42
    v47 -= 4294967295
    var v49 float64 = (0.0 * 3.14159)
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v9
}
