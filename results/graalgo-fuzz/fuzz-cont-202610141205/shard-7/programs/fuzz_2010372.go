package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0() int16 {
    var v6 int64 = -1
    var v7 uint16 = 32767
    _ = v6
    _ = v7
    return 1000
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 uint32 = 105
    var v5 int8 = -63
    v4 = v4
    v5 -= v5
    var v8 uint32 = v4
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 string = "a"
            var v12 uint32 = (v4 - v4)
            if (((v11 + (v11 + v11)) + ((v11 + ("" + (v11 + "a"))) + "a")) != v11) {
                v13 := GenS1{F0: 918}
                _ = v13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        if (80 != 77) {
            if false {
                v14 := []int{}
                for fi15 := 0; fi15 < 16; fi15++ {
                    switch "\n\t" {
                    case "hello":
                        v4 = v4
                    }
                    _ = fi15
                }
                var v16 int8 = (v5 & -18)
                _ = v14
                _ = v16
            }
            v17 := GenS2{F0: v4}
            var v18 int8 = -28
            _ = v17
            _ = v18
        } else {
            v19 := GenS1{F0: 18446744073709551615}
            _ = v19
        }
        v8 = (v8 & v8)
        _ = fi9
    }
    if false {
        var v20 uint8 = 1
        v21 := GenS2{F0: (v8 + 1)}
        _ = v20
        _ = v21
    } else {
        switch -828 {
        case 970:
            fmt.Println(v4)
        case 0:
            v1 = (v1 << ((306) & 63))
        case -2147483648:
            v8 = v8
        default:
            var v22 float64 = (1.5 + ((((1e10 - 0.1) - ((3.14159 * (3.14159 + 3.14159)) * (1e10 + (0.0 / (3.14159 - 1.5))))) + ((1e10 + -2.25) - 1e10)) * (float64(((1e10 + 3.14159) * 1e10)) - (1.5 / 1.5))))
            _ = v22
        }
        switch int(v5) {
        case 1:
            v5 = 127
        case 2147483647:
            var v23 float32 = 1e10
            _ = v23
        default:
            var v24 float64 = float64(1e10)
            _ = v24
        }
    }
    var v25 int16 = 42
    var v26 uint16 = 0
    var v27 int32 = (v1 - v1)
    v28 := GenS1{F0: 954}
    var v29 uint32 = 2147483648
    v30 := GenS1{F0: 405}
    fmt.Println(v26, v25, v5)
    switch 452 {
    case 32768:
        switch "a" {
        case "foo bar":
            for fi31 := 0; fi31 < 12; fi31++ {
                var v32 uint64 = 609
                if (true || (0 != 2)) {
                    if ((false && true) && (4294967295 <= 2)) {
                        switch ("" + ("" + "hello")) {
                        case "":
                            for fi33 := 0; fi33 < 19; fi33++ {
                                var v34 float32 = 1.5
                                if true {
                                    v29 += (v8 << ((uint(v4)) & 63))
                                    v35 := GenS2{F0: (v4 + v8)}
                                    _ = v35
                                } else {
                                    v26 = (v26 % ((uint16(v32)) | 1))
                                    switch "\n\t" {
                                    case "hello":
                                        if ((v26 <= v26) || false) {
                                            fmt.Println(v25, v5, v32)
                                        }
                                    }
                                }
                                v36 := GenS1{F0: 788}
                                _ = fi33
                                _ = v34
                                _ = v36
                            }
                        default:
                            fmt.Println(v32)
                        }
                        var v37 uint = 255
                        var v38 uint16 = 42
                        _ = v37
                        _ = v38
                    }
                    v32 = v32
                    v39 := GenS2{F0: v8}
                    _ = v39
                }
                _ = fi31
                _ = v32
            }
        case "\n\t":
            switch 32767 {
            case -2147483648:
                v40 := v3[("foo bar" + "")]
                _ = v40
            case 32768:
                v41 := []int{96}
                _ = v41
            case -140:
                v42 := GenS3{F0: 2}
                _ = v42
            default:
                for fi43 := 0; fi43 < 11; fi43++ {
                    v44 := []int{fi43, 2147483647}
                    v4 = (v8 ^ v4)
                    _ = fi43
                    _ = v44
                }
            }
        case "a":
            fmt.Println(v25, v25)
        default:
            v45 := []int{2147483647, 2147483647}
            _ = v45
        }
    }
    var v46 uint16 = v26
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v46
    _ = v5
    _ = v8
}
