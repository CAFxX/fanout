package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
func genFn0(p0 int16, p1 uint8, p2 uint8) int64 {
    fmt.Println(p2, p0, p2)
    v7 := GenS1{F0: (p2 + (p2 % (((p1 ^ (p1 % (((p1 >> ((787) & 63))) | 1)))) | 1))), F1: 1}
    p1 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 712
}

func main() {
    var v1 float32 = float32((1e10 - -2.25))
    var v2 int32 = 42
    v3 := []uint32{65535, 770, 558}
    var v4 uint8 = 42
    var v5 uint64 = 4294967295
    var v6 int = 1
    var v8 bool = (false && ((1e10 + 0.0) >= 0.1))
    var v9 bool = false
    if v9 {
        switch (("hello" + ("foo bar" + "a")) + "") {
        case "foo bar":
            for fi10 := 0; fi10 < 17; fi10++ {
                for fi11 := 0; fi11 < 6; fi11++ {
                    v9 = false
                    if false {
                        var v12 float32 = ((v1 - v1) * 3.14159)
                        var v13 float32 = (((-2.25 + v1) / float32(((1.5 + (float64(v1) * 3.14159)) * 0.0))) * v1)
                        _ = v12
                        _ = v13
                    }
                    _ = fi11
                }
                _ = fi10
            }
        case "\n\t":
            v1 -= float32((0.1 / 3.14159))
        default:
            var v14 float64 = -2.25
            _ = v14
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 19; fi16++ {
                fmt.Println(v2)
                _ = fi16
            }
            _ = fi15
        }
    } else {
        switch "foo bar" {
        case "a":
            var v17 float32 = -2.25
            _ = v17
        }
        var v18 uint8 = (v4 * (v4 % ((v4) | 1)))
        for fi19 := 0; fi19 < 9; fi19++ {
            switch "a" {
            case "":
                var v20 uint32 = 310
                _ = v20
            case "a":
                v21 := GenS1{F0: v4, F1: v5}
                _ = v21
            case "\n\t":
                for fi22 := 0; fi22 < 16; fi22++ {
                    switch (("" + ("foo bar" + "hello")) + ("a" + ("a" + "a"))) {
                    case "a":
                        if ((0.1 + 0.1) >= -2.25) {
                            if v9 {
                                v4 = 143
                                if true {
                                    if ((v5 << ((781) & 63)) <= (^(v5))) {
                                        v23 := GenS1{F0: v4, F1: (v5 | v5)}
                                        _ = v23
                                    } else {
                                        v24 := GenS1{F0: v18, F1: 2}
                                        var v25 int16 = 674
                                        if v8 {
                                            var v26 bool = v9
                                            var v27 uint8 = 2
                                            v28 := GenS1{F0: (v18 + v18), F1: v5}
                                            var v29 int64 = -9223372036854775808
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                        _ = v24
                                        _ = v25
                                    }
                                } else {
                                    var v30 uint = 255
                                    v2 += (v2 - (^(v2)))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(fi19, v6, v30)
                                    var v31 int16 = -32768
                                    var v32 int = v6
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                if len(v3) > 0 {
                                    v3[0] = 619
                                }
                            }
                            var v33 uint32 = 1
                            _ = v33
                        } else {
                            var v34 int16 = 434
                            v5 += (v5 ^ (^(v5)))
                            _ = v34
                        }
                    case "hello":
                        _ = len(v3)
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    var v35 uint32 = 407
                    _ = fi22
                    _ = v35
                }
            }
            var v36 int = (fi19 << ((964) & 63))
            v37 := GenS2{F0: 9223372036854775807, F1: 255, F2: 395}
            _ = fi19
            _ = v36
            _ = v37
        }
        _ = v18
    }
    v5 += v5
    fmt.Println(v4)
    v9 = false
    switch "foo bar" {
    case "foo bar":
        switch "a" {
        case "hello":
            for fi38 := 0; fi38 < 5; fi38++ {
                var v39 uint16 = 42
                _ = fi38
                _ = v39
            }
        }
    case "\n\t":
        var v40 uint8 = v4
        _ = v40
    case "hello":
        fmt.Println(v1, v4, v2)
    }
    v41 := GenS1{F0: v4, F1: v5}
    v42 := GenS2{F0: v6, F1: 1, F2: 2}
    for fi43 := 0; fi43 < 10; fi43++ {
        v6 = (fi43 | v6)
        _ = fi43
    }
    var v44 int32 = (v2 & (v2 | (v2 >> ((0) & 63))))
    for fi45 := 0; fi45 < 19; fi45++ {
        v1 += (1e10 / 3.14159)
        _ = fi45
    }
    var v46 uint = 958
    var v47 float64 = -2.25
    switch (v6 >> ((v46) & 63)) {
    case -11:
        v8 = (v8 || v8)
    default:
        v44 = (v2 | (v44 / (((v2 % ((v44) | 1))) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v44
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
