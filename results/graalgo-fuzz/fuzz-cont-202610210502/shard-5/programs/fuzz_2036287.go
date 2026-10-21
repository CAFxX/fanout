package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint32; F1 int }
func genFn0() bool {
    for fi7 := 0; fi7 < 4; fi7++ {
        switch ("foo bar" + ("a" + "")) {
        case "\n\t":
            v8 := []int{(fi7 % (((fi7 | fi7)) | 1)), fi7}
            _ = v8
        default:
            if (748 > 0) {
                fmt.Println(fi7)
                if false {
                    for fi9 := 0; fi9 < 17; fi9++ {
                        if false {
                            v10 := []int{fi9}
                            v10 = append(v10, (fi9 ^ (fi7 + fi7)))
                            _ = v10
                        } else {
                            fmt.Println(fi7, fi9)
                            var v11 uint64 = 42
                            _ = v11
                        }
                        _ = fi9
                    }
                    fmt.Println(fi7)
                    var v12 int8 = 0
                    _ = v12
                }
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := GenS1{F0: 18446744073709551615, F1: 739, F2: fi7}
                    if (!false) {
                        v15 := GenS2{F0: 100, F1: 449}
                        var v16 string = "foo bar"
                        if (false || (v16 < v16)) {
                            if false {
                                v16 = ((v16 + v16) + "hello")
                            } else {
                                v17 := GenS3{F0: 18446744073709551615, F1: 0}
                                fmt.Println(fi7, v16, v16)
                                _ = v17
                            }
                            if true {
                                v18 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: (-(fi13))}
                                v16 = v16
                                var v19 uint8 = 128
                                _ = v18
                                _ = v19
                            } else {
                                v20 := []int{(^(fi7)), (-(fi7))}
                                v21 := GenS2{F0: 5, F1: 178}
                                v22 := GenS1{F0: 566, F1: 18446744073709551615, F2: (fi7 * 32768)}
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        }
                        _ = v15
                        _ = v16
                    }
                    v23 := []int{fi7, fi13, (fi13 >> ((613) & 63)), (-(fi13))}
                    _ = fi13
                    _ = v14
                    _ = v23
                }
            } else {
                v24 := []int{(fi7 << ((254) & 63)), fi7}
                var v25 uint64 = 4294967295
                switch -9223372036854775808 {
                case -9223372036854775808:
                    v26 := GenS2{F0: -128, F1: 1}
                    _ = v26
                }
                _ = v24
                _ = v25
            }
        }
        fmt.Println(fi7)
        _ = fi7
    }
    var v27 uint16 = 255
    v27 = 65535
    var v28 int64 = -9223372036854775808
    fmt.Println(v27, v28)
    _ = v27
    _ = v28
    return false
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 int16 = -209
    var v5 uint64 = 18446744073709551615
    var v6 int32 = 2147483647
    v29 := GenS2{F0: -29, F1: 1}
    _ = len(v3)
    var v30 int16 = v4
    v4 -= (-(v30))
    fmt.Println(v30)
    var v31 uint32 = 42
    if (127 < 255) {
        var v32 uint8 = 42
        _ = v32
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        v34 := GenS4{F0: 2147483648, F1: (fi33 - fi33)}
        _ = fi33
        _ = v34
    }
    v35 := GenS1{F0: (^(v5)), F1: v5, F2: 331}
    var v36 int8 = 42
    var v37 int8 = (-(v36))
    if false {
        var v38 uint32 = v31
        for fi39 := 0; fi39 < 4; fi39++ {
            fmt.Println(v38)
            var v40 string = ("\n\t" + "hello")
            v41 := []int{(fi39 / ((fi39) | 1)), fi39}
            _ = fi39
            _ = v40
            _ = v41
        }
        v36 += (v37 / (((v36 >> ((653) & 63))) | 1))
        _ = v38
    } else {
        v42 := GenS4{F0: v31, F1: -2147483648}
        var v43 uint = 412
        fmt.Println(v5, v6, v6)
        _ = v42
        _ = v43
    }
    fmt.Println(v30, v30, v5)
    for fi44 := 0; fi44 < 1; fi44++ {
        var v45 int8 = v37
        var v46 float64 = -2.25
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 int32 = (v6 << ((899) & 63))
    var v48 int = int(v31)
    v37 = (^(v37))
    var v49 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
