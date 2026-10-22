package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 uint16; F1 bool; F2 int }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64) int32 {
    var v6 uint = 42
    p0 = (p0 ^ p0)
    fmt.Println(p0)
    v7 := []int{-9223372036854775808, -487, -515, -1}
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 bool = false
        if false {
            v10 := GenS1{F0: 9223372036854775807, F1: (true && v9), F2: -1}
            _ = v10
        }
        v11 := GenS1{F0: (p0 & (-(p0))), F1: false, F2: 42}
        _ = fi8
        _ = v11
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 767
    var v2 int8 = 100
    var v3 float32 = (1.5 * 0.0)
    var v4 int16 = v1
    var v5 uint32 = 65535
    switch ("hello" + "hello") {
    case "a":
        for fi12 := 0; fi12 < 8; fi12++ {
            switch ("a" + "") {
            case "a":
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := []int{}
                    fmt.Println(v4, v3, v4)
                    _ = fi13
                    _ = v14
                }
            }
            if false {
                var v15 float32 = ((-2.25 + (((float32(1.5) * (float32(1.5) + v3)) / v3) + 0.0)) / v3)
                var v16 float64 = 0.0
                if (127 <= 129) {
                    var v17 bool = ((v5 | v5) == v5)
                    v18 := []int{-9223372036854775808, fi12, fi12}
                    var v19 uint64 = 551
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    var v20 uint32 = 4294967295
                    fmt.Println(v4)
                    var v21 float32 = (v15 / v3)
                    _ = v20
                    _ = v21
                }
                _ = v15
                _ = v16
            }
            var v22 uint32 = (^(v5))
            _ = fi12
            _ = v22
        }
    case "":
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 float32 = v3
            for fi25 := 0; fi25 < 8; fi25++ {
                v26 := GenS1{F0: 1, F1: (true || true), F2: 0}
                _ = fi25
                _ = v26
            }
            fmt.Println(v24, v3)
            _ = fi23
            _ = v24
        }
    case "\n\t":
        if false {
            v4 -= 667
        } else {
            v27 := GenS1{F0: 42, F1: (205 < 2), F2: 42}
            _ = v27
        }
    default:
        v4 = (v4 & v4)
    }
    v28 := []int{-334, 0, 2147483647}
    v28 = append(v28, int(v5))
    var v29 int32 = -1
    v28 = append(v28, 32768)
    switch 0 {
    case 4294967295:
        v4 -= v1
    case 127:
        v1 = (v4 | v4)
    case -269:
        var v30 int64 = -9223372036854775808
        _ = v30
    }
    fmt.Println(v29)
    v31 := GenS2{F0: 2, F1: true, F2: 158}
    if true {
        for fi32 := 0; fi32 < 7; fi32++ {
            switch "a" {
            case "foo bar":
                switch fi32 {
                case 0:
                    for fi33 := 0; fi33 < 11; fi33++ {
                        v34 := GenS1{F0: 0, F1: false, F2: (^(v29))}
                        fmt.Println(v29, v4)
                        _ = fi33
                        _ = v34
                    }
                default:
                    var v35 int64 = -9223372036854775808
                    _ = v35
                }
            }
            _ = fi32
        }
    } else {
        fmt.Println(v29, v2, v5)
        switch 127 {
        case 128:
            var v36 bool = false
            _ = v36
        }
        var v37 float64 = (0.1 * (3.14159 + 1e10))
        _ = v37
    }
    if true {
        v29 += (v29 % (((v29 / ((0) | 1))) | 1))
        var v38 uint = 4294967295
        _ = v38
    }
    v39 := GenS2{F0: 42, F1: false, F2: 128}
    if false {
        fmt.Println(v4, v2)
        if len(v28) >= 2 {
            v28 = v28[1:len(v28)]
        }
        v40 := GenS3{F0: 1, F1: 42}
        _ = v40
    } else {
        if len(v28) > 0 {
            v28[0] = -712
        }
        v2 = v2
        if false {
            switch 0 {
            case 0:
                fmt.Println(v3)
            case 42:
                if true {
                    v5 = v5
                    if (1 <= 249) {
                        if false {
                            v41 := GenS4{F0: 2}
                            v2 = (v2 % ((0) | 1))
                            _ = v41
                        } else {
                            v42 := GenS4{F0: 0}
                            for fi43 := 0; fi43 < 2; fi43++ {
                                var v44 int16 = (v4 << ((925) & 63))
                                _ = fi43
                                _ = v44
                            }
                            if len(v28) > 0 {
                                v28[0] = 32767
                            }
                            _ = v42
                        }
                    }
                    var v45 bool = (42 == 42)
                    _ = v45
                }
            }
            var v46 int64 = -229
            _ = v46
        } else {
            var v47 float64 = (0.0 - (((0.1 / 3.14159) + 3.14159) + 1.5))
            fmt.Println(v1, v1)
            switch "" {
            case "a":
                v47 = v47
            default:
                var v48 string = "foo bar"
                _ = v48
            }
            _ = v47
        }
    }
    fmt.Println(v29)
    v49 := GenS1{F0: 1, F1: false, F2: 644}
    for fi50 := 0; fi50 < 4; fi50++ {
        v51 := GenS1{F0: 493, F1: false, F2: (v29 ^ (v29 + (v29 / (((v29 + v29)) | 1))))}
        _ = fi50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v49
    _ = v5
}
