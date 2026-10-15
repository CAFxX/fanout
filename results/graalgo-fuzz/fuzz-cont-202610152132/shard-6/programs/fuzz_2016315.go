package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 string; F1 int16; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int16; F2 uint8 }
type GenS5 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint8 }
func genFn0() int8 {
    v7 := []int{340, 1, -1, -9223372036854775808}
    switch 128 {
    case -800:
        v8 := GenS1{F0: 255, F1: 0, F2: (false || false)}
        _ = v8
    case 127:
        var v9 int8 = -1
        _ = v9
    case 287:
        var v10 int32 = 1
        _ = v10
    default:
        var v11 uint64 = 582
        _ = v11
    }
    var v12 int8 = -128
    for fi13 := 0; fi13 < 2; fi13++ {
        if false {
            v12 -= v12
            if false {
                var v14 int32 = 0
                v15 := GenS1{F0: 32767, F1: 397, F2: ("" != "")}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v14
                _ = v15
            } else {
                var v16 uint16 = 930
                _ = v16
            }
            switch (fi13 / ((512) | 1)) {
            case 9223372036854775807:
                if len(v7) > 0 {
                    v7[0] = (fi13 | 127)
                }
            case 4294967295:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            default:
                v12 = (v12 * v12)
            }
        } else {
            if true {
                var v17 float64 = (0.1 / 3.14159)
                if false {
                    var v18 int16 = 1000
                    _ = v18
                } else {
                    v19 := GenS1{F0: 526, F1: 32767, F2: true}
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = v19
                }
                _ = v17
            } else {
                var v20 uint64 = 4294967295
                _ = v20
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := GenS2{F0: ("" + "\n\t"), F1: -32768, F2: v12}
                v12 = v12
                _ = fi21
                _ = v22
            }
        }
        fmt.Println(v12, v12)
        _ = fi13
    }
    _ = v12
    _ = v7
    return v12
}

func main() {
    var v1 float64 = ((-2.25 + ((((0.1 + ((((3.14159 * 0.0) + (float64(-2.25) * 1e10)) * 0.0) + 3.14159)) + 0.1) / float64(1.5)) * 0.1)) * 0.1)
    var v2 int = -9223372036854775808
    var v3 int64 = 0
    var v4 string = "a"
    var v5 int8 = -97
    v6 := make(map[string]uint32)
    var v23 float64 = 1e10
    v24 := []int{}
    var v25 uint = 4294967295
    v5 += -91
    if len(v24) > 0 {
        v24[0] = (v2 + (v2 | (^(v2))))
    }
    var v26 int32 = int32(v3)
    var v27 int64 = -9223372036854775808
    switch "hello" {
    case "":
        var v28 int8 = (v5 - v5)
        _ = v28
    case "a":
        v25 = (^(v25))
    default:
        var v29 int = (v2 ^ (v2 | v2))
        _ = v29
    }
    var v30 uint = v25
    for fi31 := 0; fi31 < 10; fi31++ {
        switch v2 {
        case 2147483647:
            var v32 int8 = (v5 - v5)
            _ = v32
        case -2147483648:
            var v33 uint64 = 4294967295
            _ = v33
        default:
            v34 := GenS2{F0: v4, F1: 32767, F2: v5}
            _ = v34
        }
        if false {
            var v35 string = "hello"
            _ = v35
        }
        _ = fi31
    }
    fmt.Println(v25, v27, v26)
    if true {
        v30 -= (^(v25))
        switch ("\n\t" + "a") {
        case "":
            for fi36 := 0; fi36 < 1; fi36++ {
                v37 := GenS3{F0: 985}
                v5 = 76
                var v38 string = v4
                _ = fi36
                _ = v37
                _ = v38
            }
        }
    } else {
        for fi39 := 0; fi39 < 14; fi39++ {
            v2 = (fi39 >> ((42) & 63))
            _ = fi39
        }
        switch (v2 % (((v2 >> (((v25 / ((v25) | 1))) & 63))) | 1)) {
        case 32768:
            var v40 uint64 = 1
            _ = v40
        case 4294967295:
            v41 := GenS3{F0: 1}
            _ = v41
        case 141:
            v42 := GenS4{F0: (v26 >> (((^(v25))) & 63)), F1: 1, F2: 42}
            _ = v42
        default:
            for fi43 := 0; fi43 < 20; fi43++ {
                v26 += 0
                _ = fi43
            }
        }
        for fi44 := 0; fi44 < 7; fi44++ {
            for fi45 := 0; fi45 < 3; fi45++ {
                delete(v6, "a")
                _ = fi45
            }
            var v46 int16 = 1000
            v27 += (v3 ^ 1)
            _ = fi44
            _ = v46
        }
    }
    if false {
        if ((v25 >> ((698) & 63)) >= 42) {
            v3 = (v27 ^ v27)
            switch v2 {
            case 1:
                fmt.Println(v1, v5)
            case 42:
                v3 += (v3 + 25)
            case 32767:
                if true {
                    v27 += (v3 ^ 1)
                } else {
                    _ = len(v24)
                    v47 := GenS3{F0: 497}
                    var v48 int16 = -32768
                    _ = v47
                    _ = v48
                }
            default:
                fmt.Println(v4, v23)
            }
            v49 := GenS5{F0: genFn0(), F1: 4294967295, F2: 255}
            _ = v49
        }
        switch 127 {
        case -9223372036854775808:
            var v50 float32 = (0.1 / 1.0)
            _ = v50
        default:
            for fi51 := 0; fi51 < 7; fi51++ {
                if false {
                    if true {
                        switch (fi51 % ((v2) | 1)) {
                        case 599:
                            v52 := GenS6{F0: 42, F1: 72}
                            _ = v52
                        case 32768:
                            v23 -= 1.5
                        }
                    }
                } else {
                    v53 := v6["\n\t"]
                    v23 = 0.0
                    _ = v53
                }
                _ = fi51
            }
        }
    } else {
        var v54 uint16 = 1
        _ = v54
    }
    var v55 uint = (^(v30))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
