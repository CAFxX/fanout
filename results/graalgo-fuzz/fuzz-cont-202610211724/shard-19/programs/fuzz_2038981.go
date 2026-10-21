package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int8 }
type GenS2 struct { F0 string; F1 int16; F2 int32 }
func genFn0(p0 int64) uint8 {
    p0 = p0
    p0 += (p0 >> ((0) & 63))
    for fi7 := 0; fi7 < 20; fi7++ {
        switch (fi7 % (((fi7 / (((fi7 & (fi7 * (-(fi7))))) | 1))) | 1)) {
        case 343:
            var v8 bool = (!(-183 <= 42))
            _ = v8
        case 4294967295:
            p0 = (p0 / (((p0 * 839)) | 1))
        case 127:
            v9 := []int{fi7, 32767}
            _ = v9
        default:
            var v10 uint16 = 1
            _ = v10
        }
        p0 = p0
        _ = fi7
    }
    p0 = p0
    var v11 int16 = -1
    _ = p0
    _ = v11
    return 255
}
func genFn1(p0 uint64, p1 uint32, p2 int16) uint16 {
    var v12 int32 = 2147483647
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-422, 794}
    v3 := make(map[uint32]string)
    var v4 float32 = 0.1
    v5 := make(map[uint32]string)
    var v6 int = -9223372036854775808
    v6 -= (v6 - (v6 % ((v6) | 1)))
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 bool = false
        v14 = true
        if ((float64(v4) == 0.1) && false) {
            v15 := GenS1{F0: 42, F1: 65535, F2: 1}
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    v5[1] = ("foo bar" + (("a" + "") + (("" + ("foo bar" + ("" + ("a" + ("foo bar" + ("hello" + "\n\t")))))) + "foo bar")))
    var v16 int8 = 127
    if (!false) {
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(v4, v16, v6)
            _ = fi17
        }
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(v4, v4)
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 int8 = v16
                switch fi19 {
                case 283:
                    if false {
                        var v21 string = "foo bar"
                        _ = v21
                    }
                default:
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            v1[""] = 65535
                            fmt.Println(v4, fi19, v16)
                        }
                    }
                }
                _ = fi19
                _ = v20
            }
            if true {
                fmt.Println(v16)
                v22 := GenS1{F0: 42, F1: 2, F2: 81}
                if (!true) {
                    var v23 uint = 4294967295
                    v6 = (fi18 / ((128) | 1))
                    v2 = append(v2, 127)
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        v25 := GenS1{F0: 1, F1: 1, F2: 0}
                        v26 := GenS1{F0: 2147483648, F1: 4294967295, F2: (v16 - (v16 | (v16 | (-(v16)))))}
                        v6 += (^(fi18))
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS2{F0: "\n\t", F1: -1, F2: 0}
                    fmt.Println(v6, v16, fi18)
                    _ = v27
                }
                _ = v22
            }
            _ = fi18
        }
    } else {
        v28 := GenS1{F0: 517, F1: 1, F2: -46}
        var v29 uint32 = 1
        switch 4294967295 {
        case -592:
            if false {
                v3[v29] = ("foo bar" + "\n\t")
                v30 := GenS2{F0: "", F1: 42, F2: 42}
                v6 = v6
                _ = v30
            }
        case -845:
            v31 := GenS2{F0: "hello", F1: 1, F2: 835}
            _ = v31
        case 671:
            if len(v2) > 0 {
                v2[0] = v6
            }
        default:
            if len(v2) > 0 {
                v2[0] = v6
            }
        }
        _ = v28
        _ = v29
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        fmt.Println(v6)
        var v33 string = ("a" + ("hello" + ""))
        _ = fi32
        _ = v33
    }
    v16 = 100
    var v34 int32 = 188
    var v35 bool = false
    v16 = v16
    switch -525 {
    case 936:
        var v36 uint64 = 4294967295
        _ = v36
    }
    if true {
        v37 := GenS2{F0: "", F1: 1, F2: v34}
        v6 += v6
        _ = v37
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = fi38
    }
    v39 := GenS1{F0: 1, F1: 677, F2: 71}
    var v40 uint8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
