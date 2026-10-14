package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int32 }
type GenS2 struct { F0 uint; F1 uint16; F2 int16 }
type GenS3 struct { F0 int16; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint; F1 int64 }

func main() {
    var v1 float64 = ((float64(3.14159) * (0.1 / -2.25)) / 1.5)
    var v2 uint = 0
    v3 := []int{9223372036854775807, 42, 1, 32768}
    v4 := make(map[int]bool)
    v5 := make(map[string]int)
    var v6 int8 = -106
    v7 := GenS1{F0: 401, F1: -766, F2: 42}
    if (!(2 > 1)) {
        var v8 float64 = v1
        v9 := GenS1{F0: 434, F1: 9223372036854775807, F2: -1}
        _ = v8
        _ = v9
    } else {
        var v10 int64 = -344
        _ = v10
    }
    if (-1 < 1) {
        if false {
            v3 = append(v3, -2147483648)
        } else {
            switch 9223372036854775807 {
            case 42:
                if true {
                    v11 := GenS1{F0: 1, F1: 32768, F2: 1}
                    var v12 float64 = v1
                    _ = v11
                    _ = v12
                }
            default:
                for fi13 := 0; fi13 < 7; fi13++ {
                    _ = len(v3)
                    _ = fi13
                }
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                fmt.Println(v6, fi14)
                if true {
                    if (!((0.0 / -2.25) != (0.0 * 0.1))) {
                        switch "\n\t" {
                        case "foo bar":
                            delete(v5, "a")
                        }
                        if (true && false) {
                            var v15 uint16 = 65535
                            v2 -= 65535
                            _ = v15
                        }
                    }
                    v1 -= v1
                    v16 := GenS2{F0: v2, F1: 898, F2: 1}
                    _ = v16
                }
                v6 += 100
                _ = fi14
            }
            switch ("foo bar" + "a") {
            case "":
                for fi17 := 0; fi17 < 1; fi17++ {
                    if true {
                        for fi18 := 0; fi18 < 18; fi18++ {
                            switch fi17 {
                            case -2147483648:
                                v19 := GenS1{F0: 4294967295, F1: (fi17 & (fi17 & 0)), F2: -2147483648}
                                _ = v19
                            }
                            fmt.Println(fi18, v1, fi18)
                            v20 := GenS3{F0: 974, F1: (v2 | v2), F2: 682}
                            _ = fi18
                            _ = v20
                        }
                    } else {
                        for fi21 := 0; fi21 < 5; fi21++ {
                            v2 = v2
                            _ = fi21
                        }
                        if (1 != 9223372036854775807) {
                            fmt.Println(v6, fi17, v6)
                            v22 := GenS1{F0: 2147483648, F1: (fi17 / ((fi17) | 1)), F2: 42}
                            v23 := GenS2{F0: (v2 * (v2 * (v2 >> (((v2 ^ v2)) & 63)))), F1: 42, F2: int16(v2)}
                            _ = v22
                            _ = v23
                        }
                    }
                    _ = len(v3)
                    v1 = 0.1
                    _ = fi17
                }
            case "foo bar":
                var v24 string = "hello"
                _ = v24
            default:
                v25 := GenS4{F0: v2, F1: -9223372036854775808}
                _ = v25
            }
        }
        v1 += v1
        v3 = append(v3, 9223372036854775807)
    }
    if false {
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := v4[(fi26 - fi26)]
            var v28 float64 = v1
            var v29 int16 = 0
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
    } else {
        v1 += v1
        v4[-9223372036854775808] = (true || (!true))
    }
    var v30 uint32 = 657
    if false {
        v31 := v5[("" + "hello")]
        v30 += (^(v30))
        _ = v31
    } else {
        v32 := GenS3{F0: 677, F1: 18446744073709551615, F2: 42}
        v33 := GenS4{F0: v2, F1: 1}
        v6 += (v6 - 100)
        _ = v32
        _ = v33
    }
    v1 = v1
    v34 := GenS1{F0: 65535, F1: -2147483648, F2: 28}
    var v35 float64 = v1
    v36 := v5[""]
    v30 = (v30 & v30)
    fmt.Println(v35, v35)
    v37 := GenS1{F0: v30, F1: v36, F2: 813}
    v30 += v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
