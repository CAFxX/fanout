package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
func genFn0(p0 bool, p1 uint16, p2 int8) uint8 {
    var v5 bool = (p0 || (1 != 42))
    v6 := GenS1{F0: -111, F1: 270, F2: 255}
    v5 = true
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int32 = 42
        v9 := []int{fi7, fi7, fi7, (fi7 << ((2) & 63))}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint64 = 4294967295
    var v3 uint8 = 128
    v4 := make(map[string]string)
    for fi10 := 0; fi10 < 15; fi10++ {
        v2 = (v2 % (((v2 & 1)) | 1))
        var v11 int16 = 197
        _ = fi10
        _ = v11
    }
    v3 -= (v3 >> ((v1) & 63))
    if true {
        if false {
            switch ("" + "") {
            case "a":
                for fi12 := 0; fi12 < 15; fi12++ {
                    var v13 uint64 = v2
                    _ = fi12
                    _ = v13
                }
            case "foo bar":
                if false {
                    v14 := []int{4294967295, 4294967295, 127, 0}
                    if true {
                        var v15 string = "foo bar"
                        if true {
                            _ = len(v4)
                        }
                        _ = v15
                    }
                    _ = v14
                } else {
                    var v16 uint32 = uint32(v2)
                    _ = v16
                }
            case "hello":
                v17 := GenS1{F0: -1, F1: -1, F2: v1}
                _ = v17
            }
            v3 += 182
            for fi18 := 0; fi18 < 13; fi18++ {
                v3 -= (v3 * v3)
                for fi19 := 0; fi19 < 16; fi19++ {
                    if (821 <= 18446744073709551615) {
                        var v20 uint64 = (v2 * 1)
                        v21 := []int{-751}
                        _ = v20
                        _ = v21
                    }
                    var v22 uint16 = 724
                    _ = fi19
                    _ = v22
                }
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 int16 = -289
                    switch "foo bar" {
                    case "\n\t":
                        var v25 float32 = 1.5
                        fmt.Println(fi23, v24, v25)
                        _ = v25
                    default:
                        v26 := GenS1{F0: 42, F1: (v24 / (((v24 - v24)) | 1)), F2: 4294967295}
                        _ = v26
                    }
                    var v27 uint16 = 975
                    _ = fi23
                    _ = v24
                    _ = v27
                }
                _ = fi18
            }
        } else {
            var v28 int = -243
            var v29 bool = false
            v30 := v4["hello"]
            _ = v28
            _ = v29
            _ = v30
        }
        var v31 float32 = ((float32((-2.25 / 1.5)) + float32(float64(0.0))) * float32((((3.14159 - (3.14159 * 3.14159)) * -2.25) * 0.0)))
        switch 4294967295 {
        case 0:
            var v32 int32 = 2147483647
            _ = v32
        case 32768:
            v33 := GenS1{F0: -1, F1: 32767, F2: (v1 ^ 0)}
            _ = v33
        case 4294967295:
            for fi34 := 0; fi34 < 19; fi34++ {
                var v35 int = fi34
                _ = fi34
                _ = v35
            }
        }
        _ = v31
    } else {
        v36 := GenS1{F0: 1, F1: -32768, F2: (v1 << ((v1) & 63))}
        _ = v36
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        if true {
            for fi38 := 0; fi38 < 10; fi38++ {
                var v39 int64 = 9223372036854775807
                _ = fi38
                _ = v39
            }
            var v40 string = "\n\t"
            v41 := GenS1{F0: 813, F1: -32768, F2: (v1 & v1)}
            _ = v40
            _ = v41
        } else {
            fmt.Println(v3, v1)
        }
        v42 := GenS2{F0: -747, F1: v2, F2: v1}
        _ = fi37
        _ = v42
    }
    var v43 int8 = int8(v2)
    v4["foo bar"] = ("" + "a")
    switch "" {
    case "a":
        for fi44 := 0; fi44 < 8; fi44++ {
            var v45 uint16 = 564
            switch "a" {
            case "":
                fmt.Println(fi44, v1)
            }
            var v46 uint = (v1 - (v1 << ((v1) & 63)))
            _ = fi44
            _ = v45
            _ = v46
        }
    default:
        v47 := GenS1{F0: -1, F1: -608, F2: (v1 - 2)}
        _ = v47
    }
    v3 += v3
    var v48 int8 = v43
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v48
}
