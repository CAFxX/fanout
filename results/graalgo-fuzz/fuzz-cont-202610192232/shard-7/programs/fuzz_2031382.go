package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 0.0
    var v3 uint = 471
    v4 := GenS1{F0: 128, F1: 32767, F2: 0}
    v3 = v3
    v5 := []int{32768, 42, 547, 32767}
    if (!true) {
        var v6 int64 = 1
        for fi7 := 0; fi7 < 11; fi7++ {
            v6 += 0
            fmt.Println(v2, v3, v6)
            _ = len(v5)
            _ = fi7
        }
        if (!(65535 > 803)) {
            v6 -= v6
        }
        _ = v6
    }
    switch -609 {
    case -275:
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint16 = 782
            v9 -= 32767
            switch (fi8 % ((32767) | 1)) {
            case 973:
                var v10 bool = false
                _ = v10
            case 4294967295:
                if ((1e10 <= 1.5) && false) {
                    fmt.Println(v9, v2)
                    v11 := GenS1{F0: 128, F1: -1, F2: 713}
                    if (!(true && (-1 < 42))) {
                        if false {
                            fmt.Println(fi8, v9)
                            var v12 uint64 = 4294967295
                            _ = v12
                        }
                    } else {
                        v3 += (v3 - uint(v9))
                        var v13 int = (fi8 >> ((0) & 63))
                        v14 := GenS2{F0: -1}
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                } else {
                    fmt.Println(fi8)
                    var v15 uint32 = 343
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v1[(^(v15))] = "foo bar"
                        var v17 uint64 = 2
                        var v18 float64 = v2
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = v15
                }
            case 454:
                v3 = v3
            }
            _ = fi8
            _ = v9
        }
    case 2147483647:
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 uint64 = 2
            _ = fi19
            _ = v20
        }
    case 4294967295:
        for fi21 := 0; fi21 < 7; fi21++ {
            fmt.Println(v2)
            var v22 int32 = int32(v3)
            fmt.Println(v2, v3, v2)
            _ = fi21
            _ = v22
        }
    default:
        if false {
            var v23 int32 = -126
            if false {
                for fi24 := 0; fi24 < 20; fi24++ {
                    switch 661 {
                    case 0:
                        for fi25 := 0; fi25 < 5; fi25++ {
                            v5 = append(v5, (fi24 / ((int(v3)) | 1)))
                            var v26 int64 = -882
                            v23 += v23
                            _ = fi25
                            _ = v26
                        }
                    case 1:
                        var v27 uint32 = 65535
                        _ = v27
                    default:
                        fmt.Println(fi24, v2)
                    }
                    _ = fi24
                }
            }
            var v28 int = -596
            _ = v23
            _ = v28
        } else {
            v2 = (((v2 / (3.14159 + (v2 + v2))) - 3.14159) - v2)
            fmt.Println(v3, v3)
        }
    }
    v29 := GenS2{F0: 563}
    v30 := GenS2{F0: 1}
    if (!true) {
        var v31 string = ("\n\t" + "foo bar")
        if (v3 <= (v3 % ((18446744073709551615) | 1))) {
            _ = len(v5)
            v31 = (v31 + "hello")
            var v32 uint8 = 236
            _ = v32
        } else {
            var v33 bool = (false || false)
            _ = v33
        }
        _ = v31
    }
    var v34 uint16 = 0
    switch 1 {
    case 9223372036854775807:
        switch 0 {
        case 2147483647:
            var v35 string = ("a" + "foo bar")
            _ = v35
        default:
            var v36 uint64 = 564
            _ = v36
        }
    default:
        var v37 float64 = 1.5
        _ = v37
    }
    switch 128 {
    case 4294967295:
        delete(v1, 1)
    case 1:
        v38 := GenS2{F0: -236}
        _ = v38
    }
    var v39 int16 = 0
    v40 := GenS2{F0: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
