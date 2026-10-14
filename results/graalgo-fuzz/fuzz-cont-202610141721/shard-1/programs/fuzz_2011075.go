package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
type GenS3 struct { F0 bool; F1 int16; F2 int }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint16) int16 {
    p0 += (p0 + p0)
    if ((("" + "\n\t") + ("" + "\n\t")) >= "\n\t") {
        v7 := []int{}
        _ = v7
    }
    _ = p0
    return 1000
}
func genFn1(p0 int8) int16 {
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int16 = -154
        _ = fi8
        _ = v9
    }
    p0 = (p0 | 75)
    var v10 int = 32767
    v11 := []int{v10, (-(v10))}
    _ = p0
    _ = v10
    _ = v11
    return 41
}

func main() {
    var v1 int = -2147483648
    var v2 float32 = 0.1
    var v3 uint = 660
    v4 := []string{}
    var v5 uint8 = 191
    var v6 int32 = 0
    switch v1 {
    case -678:
        for fi12 := 0; fi12 < 15; fi12++ {
            if true {
                v13 := GenS1{F0: 884}
                _ = v13
            }
            _ = fi12
        }
    case 4294967295:
        v2 = 0.0
    default:
        v14 := GenS1{F0: 783}
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 5; fi15++ {
            v2 -= (float32(((float64(3.14159) + float64(v2)) - 0.0)) * ((float32(0.0) + -2.25) * v2))
            _ = fi15
        }
        var v16 float32 = 0.1
        var v17 int32 = v6
        _ = v16
        _ = v17
    } else {
        switch ("a" + "foo bar") {
        case "\n\t":
            for fi18 := 0; fi18 < 13; fi18++ {
                for fi19 := 0; fi19 < 11; fi19++ {
                    v2 += float32(0.1)
                    _ = fi19
                }
                v20 := GenS2{F0: (fi18 >= 32767), F1: int64(v6), F2: -423}
                v2 -= 0.0
                _ = fi18
                _ = v20
            }
        case "foo bar":
            v21 := GenS2{F0: (v5 >= v5), F1: 0, F2: -731}
            _ = v21
        case "hello":
            var v22 uint = (^(v3))
            _ = v22
        }
        if (2 <= 808) {
            v1 += v1
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v1, v2, v2)
        }
        if true {
            if len(v4) > 0 {
                v4[0] = ("a" + "\n\t")
            }
            var v23 bool = (!false)
            _ = v23
        } else {
            var v24 int8 = 1
            for fi25 := 0; fi25 < 9; fi25++ {
                _ = len(v4)
                _ = fi25
            }
            _ = v24
        }
    }
    fmt.Println(v2, v5)
    v5 -= v5
    v5 = (^(v5))
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    if true {
        v26 := GenS2{F0: (0 >= -1), F1: -1, F2: -43}
        _ = v26
    } else {
        var v27 float32 = v2
        _ = v27
    }
    v1 -= v1
    if false {
        v2 = v2
        var v28 int = 32768
        _ = v28
    } else {
        var v29 int64 = -877
        v4 = append(v4, "foo bar")
        _ = v29
    }
    if false {
        _ = len(v4)
        if false {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        } else {
            if (0 <= 0) {
                if (-1 > 100) {
                    if (1e10 == v2) {
                        fmt.Println(v1, v5)
                        v30 := GenS3{F0: (1.5 != 3.14159), F1: 958, F2: (v1 << ((65535) & 63))}
                        var v31 uint = (v3 & 2)
                        _ = v30
                        _ = v31
                    }
                    var v32 string = "a"
                    v33 := GenS1{F0: 255}
                    _ = v32
                    _ = v33
                }
                v1 = (v1 & int(v5))
                var v34 int16 = 1
                _ = v34
            }
            for fi35 := 0; fi35 < 20; fi35++ {
                v3 = (^(v3))
                _ = fi35
            }
            v36 := GenS2{F0: false, F1: 541, F2: 0}
            _ = v36
        }
        v6 = -1
    }
    v4 = append(v4, "\n\t")
    v37 := GenS1{F0: 2}
    v38 := GenS4{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
