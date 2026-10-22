package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint }
func genFn0(p0 uint) int32 {
    v7 := GenS1{F0: 42}
    var v8 bool = false
    _ = p0
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 uint64 = 1
    var v2 uint = 65535
    var v3 int64 = -1
    v4 := []uint32{42, 124, 42}
    v5 := []string{"a", (("\n\t" + ("\n\t" + "\n\t")) + ("a" + "\n\t")), "foo bar"}
    v6 := []int{0, 4294967295}
    v2 = 747
    switch 0 {
    case 9223372036854775807:
        switch 2147483647 {
        case 42:
            v3 = (v3 | v3)
        case 283:
            switch "" {
            case "foo bar":
                if true {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        var v10 bool = false
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                        _ = fi9
                        _ = v10
                    }
                }
            }
        case 32767:
            var v11 string = ("" + "\n\t")
            _ = v11
        }
    case 127:
        for fi12 := 0; fi12 < 10; fi12++ {
            switch (fi12 >> ((390) & 63)) {
            case 575:
                v2 += (v2 / ((uint(v1)) | 1))
            }
            v1 = 42
            _ = fi12
        }
    default:
        for fi13 := 0; fi13 < 7; fi13++ {
            v3 += 42
            _ = fi13
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v14 uint = 479
    v15 := GenS1{F0: (-(v3))}
    if (0 != 938) {
        var v16 uint8 = 229
        var v17 int16 = 32767
        fmt.Println(v17, v16, v3)
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(v14)
        v19 := GenS2{F0: -246}
        v20 := GenS2{F0: -2147483648}
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v1, v1)
    for fi21 := 0; fi21 < 12; fi21++ {
        v1 = 453
        var v22 uint32 = 0
        v22 -= (v22 - (v22 / (((^(v22))) | 1)))
        _ = fi21
        _ = v22
    }
    v1 = (v1 % ((v1) | 1))
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 int8 = 0
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 int8 = v24
            _ = fi25
            _ = v26
        }
        var v27 int = (fi23 / (((-(fi23))) | 1))
        _ = fi23
        _ = v24
        _ = v27
    }
    v28 := GenS2{F0: genFn0((^(v2)))}
    v6 = append(v6, -1)
    var v29 float32 = float32(float64(-2.25))
    v3 = (v3 / ((-974) | 1))
    switch "\n\t" {
    case "":
        v14 -= v14
    case "a":
        if false {
            v2 = (v14 % (((v14 >> (((v14 + v2)) & 63))) | 1))
            v1 = (v1 * v1)
            var v30 uint16 = 32767
            _ = v30
        } else {
            if (!true) {
                v31 := GenS3{F0: 255, F1: 1, F2: 126}
                _ = v31
            } else {
                for fi32 := 0; fi32 < 1; fi32++ {
                    switch (fi32 - (fi32 + fi32)) {
                    case 32767:
                        for fi33 := 0; fi33 < 2; fi33++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi33
                        }
                    case 683:
                        var v34 int16 = 993
                        _ = v34
                    default:
                        v3 += v3
                    }
                    _ = fi32
                }
                v35 := GenS4{F0: -32768}
                _ = v35
            }
            switch -213 {
            case -9223372036854775808:
                if false {
                    v36 := GenS5{F0: "", F1: (v2 >> ((v14) & 63))}
                    fmt.Println(v29, v2)
                    switch (("a" + ("\n\t" + ("hello" + ("a" + ("a" + ""))))) + "\n\t") {
                    case "\n\t":
                        if len(v6) > 0 {
                            v6[0] = 4294967295
                        }
                    case "a":
                        v37 := GenS4{F0: 32767}
                        _ = v37
                    }
                    _ = v36
                }
            default:
                v1 += (v1 << (((^(v14))) & 63))
            }
        }
    case "hello":
        var v38 bool = (0 < 65535)
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
