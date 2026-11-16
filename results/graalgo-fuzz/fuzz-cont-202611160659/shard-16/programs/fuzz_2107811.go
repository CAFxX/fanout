package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 bool; F1 int; F2 int16 }

func main() {
    var v1 float32 = (3.14159 - (1.5 + (((0.1 * ((0.1 - (0.1 + 1.5)) + (0.1 / 1e10))) + ((3.14159 - -2.25) - 3.14159)) - 0.1)))
    var v2 int = 42
    var v3 uint8 = 42
    var v4 uint = 404
    var v5 int = (v2 / (((v2 << (((^(v4))) & 63))) | 1))
    if false {
        var v6 uint8 = (v3 | v3)
        switch v2 {
        case 128:
            switch "foo bar" {
            case "hello":
                switch "" {
                case "hello":
                    v7 := []int{2147483647, (v2 % (((v5 | v5)) | 1)), v2}
                    _ = v7
                case "foo bar":
                    v8 := GenS1{F0: 1}
                    _ = v8
                default:
                    var v9 int64 = -238
                    _ = v9
                }
            case "":
                for fi10 := 0; fi10 < 7; fi10++ {
                    for fi11 := 0; fi11 < 4; fi11++ {
                        v12 := GenS1{F0: 18446744073709551615}
                        var v13 uint = (v4 >> (((v4 - v4)) & 63))
                        fmt.Println(v3, v6, v2)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    var v14 int16 = 997
                    _ = fi10
                    _ = v14
                }
            case "foo bar":
                v4 = (v4 / ((4294967295) | 1))
            default:
                v4 += v4
            }
        case 32767:
            if ((1 <= uint64(v4)) && false) {
                var v15 int32 = -2147483648
                var v16 float64 = ((1e10 * 1e10) * (1e10 - -2.25))
                _ = v15
                _ = v16
            }
        }
        v17 := GenS1{F0: 4294967295}
        _ = v17
        _ = v6
    } else {
        v2 += v2
    }
    v3 += (v3 | v3)
    for fi18 := 0; fi18 < 11; fi18++ {
        if (42 != 0) {
            fmt.Println(v3)
            v19 := GenS2{F0: v3, F1: v2}
            _ = v19
        } else {
            var v20 int16 = -1
            _ = v20
        }
        var v21 uint8 = v3
        var v22 int16 = 0
        _ = fi18
        _ = v21
        _ = v22
    }
    if true {
        var v23 int16 = -616
        v24 := []int{v5, 634}
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 10; fi26++ {
                _ = len(v24)
                _ = fi26
            }
            v27 := GenS1{F0: 2}
            _ = fi25
            _ = v27
        }
        _ = v23
        _ = v24
    } else {
        var v28 bool = (492 < 426)
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint32 = 2147483648
        v30 = (^(v30))
        _ = fi29
        _ = v30
    }
    v31 := []int{(v5 | v5), (v2 >> ((332) & 63)), (v2 * (v2 >> ((v4) & 63))), v2}
    if (uint(v3) != (v4 >> (((v4 ^ (^(v4)))) & 63))) {
        if (-1 >= 100) {
            for fi32 := 0; fi32 < 20; fi32++ {
                v2 = (v5 >> ((v4) & 63))
                var v33 uint = 344
                var v34 int64 = -1
                _ = fi32
                _ = v33
                _ = v34
            }
        }
    } else {
        var v35 float64 = 1e10
        _ = v35
    }
    v4 = (^(v4))
    if true {
        v36 := GenS1{F0: 308}
        if false {
            var v37 uint = uint(v3)
            for fi38 := 0; fi38 < 10; fi38++ {
                for fi39 := 0; fi39 < 3; fi39++ {
                    _ = len(v31)
                    v3 += (v3 % (((v3 | 2)) | 1))
                    v40 := GenS1{F0: 2}
                    _ = fi39
                    _ = v40
                }
                var v41 int = int(v37)
                _ = fi38
                _ = v41
            }
            _ = v37
        }
        _ = v36
    }
    switch (-(v2)) {
    case 655:
        v42 := GenS3{F0: true, F1: -925, F2: -887}
        _ = v42
    }
    v43 := GenS3{F0: true, F1: (^(v2)), F2: 186}
    var v44 uint16 = 1
    var v45 uint8 = (v3 / (((v3 << ((v4) & 63))) | 1))
    for fi46 := 0; fi46 < 10; fi46++ {
        v5 = -813
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
