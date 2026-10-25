package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    var v1 int32 = 635
    v2 := []int{156, -9223372036854775808}
    var v3 int16 = -524
    var v4 int = -105
    var v5 float64 = 3.14159
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int64 = 9223372036854775807
        var v9 string = "a"
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 255, F1: (^(v3)), F2: -32}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 float32 = float32(v5)
    fmt.Println(v5)
    switch (v4 - v4) {
    case 9223372036854775807:
        switch v4 {
        case -50:
            if (160 != 9223372036854775807) {
                v11 = v11
                var v12 uint = 18446744073709551615
                if false {
                    if (v12 > (v12 & v12)) {
                        v4 = (v4 % (((v4 + v4)) | 1))
                    } else {
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v14 := GenS2{F0: -1, F1: 5}
                            v11 -= 0.0
                            if (2 >= 2) {
                                var v15 uint16 = 1
                                v4 = v4
                                v1 = v1
                                _ = v15
                            }
                            _ = fi13
                            _ = v14
                        }
                        if (!(!false)) {
                            v11 = (((float32(v5) + v11) + 0.0) + v11)
                            v12 = v12
                        } else {
                            var v16 int8 = 42
                            _ = v16
                        }
                        var v17 float64 = 0.1
                        _ = v17
                    }
                }
                _ = v12
            } else {
                v3 -= int16(v1)
            }
        }
    default:
        v1 = (v1 + (v1 + v1))
    }
    if true {
        for fi18 := 0; fi18 < 7; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                v4 -= fi18
                if (v1 <= -1) {
                    v5 = 0.1
                }
                _ = fi19
            }
            _ = fi18
        }
    }
    var v20 uint8 = 255
    var v21 string = "\n\t"
    switch v4 {
    case 128:
        if (!(18446744073709551615 == 4294967295)) {
            var v22 string = v21
            var v23 int8 = -35
            _ = v22
            _ = v23
        } else {
            if (1 < 0) {
                v24 := GenS1{F0: 534, F1: -32768, F2: 42}
                if false {
                    delete(v6, 9223372036854775807)
                } else {
                    v25 := GenS2{F0: 127, F1: (v20 / ((81) | 1))}
                    v26 := GenS2{F0: 100, F1: v20}
                    switch 0 {
                    case 128:
                        var v27 int = (v4 ^ (-(v4)))
                        _ = v27
                    case 32768:
                        v20 -= 227
                    case 127:
                        v2 = append(v2, v4)
                    }
                    _ = v25
                    _ = v26
                }
                if true {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        for fi29 := 0; fi29 < 2; fi29++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if true {
                                v21 = "foo bar"
                            } else {
                                var v30 uint64 = 0
                                var v31 uint16 = 42
                                _ = v30
                                _ = v31
                            }
                            _ = fi29
                        }
                        var v32 uint = 477
                        _ = fi28
                        _ = v32
                    }
                }
                _ = v24
            } else {
                v33 := GenS2{F0: -73, F1: v20}
                v20 += (v20 & v20)
                for fi34 := 0; fi34 < 13; fi34++ {
                    v35 := GenS1{F0: 871, F1: (v3 & (v3 + v3)), F2: 42}
                    v3 = v3
                    _ = fi34
                    _ = v35
                }
                _ = v33
            }
        }
    case 4294967295:
        fmt.Println(v4)
    }
    if false {
        switch (v21 + (v21 + (((("a" + ("" + "\n\t")) + v21) + ((v21 + "a") + v21)) + ("" + v21)))) {
        case "foo bar":
            var v36 uint16 = 2
            _ = v36
        case "hello":
            var v37 int16 = 0
            _ = v37
        }
        if ((true || false) || (false || ((v21 + (v21 + v21)) >= v21))) {
            var v38 int8 = 127
            for fi39 := 0; fi39 < 17; fi39++ {
                v40 := v6[fi39]
                var v41 bool = (!true)
                switch v21 {
                case "a":
                    var v42 int64 = 9223372036854775807
                    _ = v42
                case "hello":
                    var v43 int16 = v3
                    _ = v43
                }
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = v38
        } else {
            for fi44 := 0; fi44 < 9; fi44++ {
                for fi45 := 0; fi45 < 20; fi45++ {
                    var v46 bool = ((!true) && false)
                    _ = fi45
                    _ = v46
                }
                _ = fi44
            }
        }
    }
    v4 = (v4 + (v4 * v4))
    var v47 int8 = 42
    var v48 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
