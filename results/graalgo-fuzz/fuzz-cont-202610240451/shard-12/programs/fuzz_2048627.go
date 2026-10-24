package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 uint64 = 42
    v4 := []int{}
    var v5 string = "\n\t"
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 14; fi7++ {
        if (100 <= 100) {
            if true {
                for fi8 := 0; fi8 < 14; fi8++ {
                    switch v5 {
                    case "\n\t":
                        v4 = append(v4, -9223372036854775808)
                    case "hello":
                        var v9 int8 = 42
                        _ = v9
                    case "":
                        var v10 uint64 = 18446744073709551615
                        _ = v10
                    }
                    var v11 int8 = 0
                    v12 := GenS1{F0: 0, F1: -119}
                    _ = fi8
                    _ = v11
                    _ = v12
                }
                v6 -= (v6 / (((v6 - (v6 / (((v6 << ((4294967295) & 63))) | 1)))) | 1))
            }
            delete(v1, (fi7 ^ (fi7 << ((0) & 63))))
        } else {
            var v13 uint32 = 281
            _ = v13
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS2{F0: 2147483647, F1: 42}
            var v16 int8 = 1
            _ = fi14
            _ = v15
            _ = v16
        }
        switch "\n\t" {
        case "a":
            v3 += v3
        default:
            fmt.Println(v6)
        }
        _ = fi7
    }
    var v17 int64 = -9223372036854775808
    v6 += 127
    switch "" {
    case "a":
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                if (255 != 0) {
                    fmt.Println(v5, fi18)
                    fmt.Println(v6, fi18)
                    var v20 float32 = -2.25
                    _ = v20
                }
                v3 = 4294967295
                _ = fi19
            }
            var v21 string = v5
            _ = fi18
            _ = v21
        }
    default:
        if ((37 > -1) || ((1 == 0) && true)) {
            var v22 int64 = (v17 % (((v17 - 0)) | 1))
            v22 -= (^(v17))
            var v23 float64 = 1e10
            _ = v22
            _ = v23
        } else {
            v24 := GenS2{F0: 2147483647, F1: 0}
            _ = v24
        }
    }
    v25 := GenS3{F0: 127}
    var v26 int16 = 0
    _ = len(v4)
    var v27 uint16 = 768
    var v28 float64 = ((((0.1 + 1e10) * (((3.14159 * -2.25) - 3.14159) + (3.14159 - 0.0))) - -2.25) + float64(1.5))
    var v29 int8 = -49
    var v30 bool = false
    var v31 int32 = -1
    switch -471 {
    case 2147483647:
        if (int16(v17) < 1) {
            var v32 uint8 = (v6 << ((2) & 63))
            var v33 uint8 = v6
            _ = v32
            _ = v33
        }
    case 4294967295:
        if v30 {
            for fi34 := 0; fi34 < 18; fi34++ {
                if ((!true) && (!v30)) {
                    _ = len(v4)
                    v6 -= v6
                    var v35 int8 = (v29 >> ((18446744073709551615) & 63))
                    _ = v35
                }
                v29 = v29
                fmt.Println(v29)
                _ = fi34
            }
            if v30 {
                var v36 uint16 = (v27 - 655)
                fmt.Println(v26)
                _ = v36
            }
        }
    case -2147483648:
        for fi37 := 0; fi37 < 6; fi37++ {
            v38 := GenS3{F0: (v29 | v29)}
            v30 = ((1e10 > 3.14159) && v30)
            for fi39 := 0; fi39 < 8; fi39++ {
                var v40 uint32 = 2
                fmt.Println(v3, v29)
                var v41 uint8 = 127
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi37
            _ = v38
        }
    }
    v4 = append(v4, 59)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
