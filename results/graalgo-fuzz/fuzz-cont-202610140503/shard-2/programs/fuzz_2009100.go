package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint32) uint8 {
    p0 -= 2147483647
    var v7 int64 = -1
    fmt.Println(v7, p0)
    var v8 uint = 2
    _ = p0
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{4294967295}
    v3 := make(map[string]string)
    var v4 int32 = 734
    v5 := make(map[int]string)
    var v6 uint8 = 128
    if len(v2) > 0 {
        v2[0] = 959
    }
    switch (v1 + v1) {
    case "foo bar":
        var v9 uint16 = 42
        _ = v9
    }
    if (!((1 <= 80) || true)) {
        switch -974 {
        case 9223372036854775807:
            v10 := GenS1{F0: v4, F1: 2147483647, F2: 283}
            _ = v10
        case 0:
            var v11 float64 = 0.0
            _ = v11
        case 128:
            fmt.Println(v4)
        default:
            var v12 bool = (18446744073709551615 != 138)
            _ = v12
        }
    } else {
        v6 = genFn0(2)
        var v13 uint64 = 2
        _ = v13
    }
    switch 772 {
    case 4294967295:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 42:
        if false {
            v1 = (("" + "a") + (v1 + v1))
        }
    case 127:
        if false {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v3)
            }
        } else {
            if (!true) {
                var v14 int8 = -63
                if (v1 < "hello") {
                    v5[32767] = (v1 + v1)
                    v15 := GenS1{F0: (v4 ^ 42), F1: 2147483647, F2: 1}
                    if (int8(v4) < (v14 | (v14 & 4))) {
                        for fi16 := 0; fi16 < 5; fi16++ {
                            switch ("\n\t" + (v1 + v1)) {
                            case "hello":
                                v17 := GenS2{F0: fi16}
                                _ = v17
                            case "":
                                delete(v5, fi16)
                            default:
                                for fi18 := 0; fi18 < 16; fi18++ {
                                    var v19 uint32 = 42
                                    _ = len(v2)
                                    v19 = v19
                                    v19 = (v19 & 2147483648)
                                    _ = fi18
                                    _ = v19
                                }
                            }
                            var v20 int64 = 0
                            _ = fi16
                            _ = v20
                        }
                    }
                    _ = v15
                } else {
                    if len(v2) > 0 {
                        v2[0] = -2147483648
                    }
                    var v21 float64 = 1.5
                    var v22 bool = false
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 float64 = float64(1e10)
                    v25 := GenS3{F0: -1}
                    delete(v5, fi23)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = v14
            } else {
                v4 += -2147483648
            }
        }
    }
    if ((!(true && (2 < 0))) && (!false)) {
        v6 += 2
    }
    var v26 int16 = -788
    var v27 uint = 625
    for fi28 := 0; fi28 < 16; fi28++ {
        v27 = (v27 + 2)
        _ = fi28
    }
    var v29 int = -2147483648
    v27 += (v27 - (v27 - v27))
    if (!false) {
        switch v29 {
        case 787:
            for fi30 := 0; fi30 < 3; fi30++ {
                switch v1 {
                case "\n\t":
                    v27 += 4294967295
                case "a":
                    switch -912 {
                    case -505:
                        fmt.Println(v29, v27)
                    case 9223372036854775807:
                        var v31 int8 = 42
                        _ = v31
                    default:
                        v32 := GenS3{F0: 0}
                        _ = v32
                    }
                default:
                    switch fi30 {
                    case -1:
                        fmt.Println(fi30)
                    case 1:
                        v33 := GenS3{F0: int64(fi30)}
                        _ = v33
                    }
                }
                if len(v2) > 0 {
                    v2[0] = fi30
                }
                if false {
                    var v34 bool = ((((v1 + v1) + "hello") <= "\n\t") || false)
                    v27 = 2
                    _ = v34
                }
                _ = fi30
            }
        }
    }
    var v35 uint64 = 4294967295
    v6 = 127
    switch (-(v29)) {
    case 790:
        for fi36 := 0; fi36 < 14; fi36++ {
            var v37 uint16 = 255
            _ = fi36
            _ = v37
        }
    }
    if (false || true) {
        switch 2147483647 {
        case -9223372036854775808:
            if true {
                var v38 string = v1
                _ = v38
            }
        case 1:
            v27 = 969
        case 0:
            var v39 int16 = 908
            _ = v39
        default:
            v40 := GenS4{F0: (v35 ^ 422)}
            _ = v40
        }
        if false {
            v26 = (v26 << ((4294967295) & 63))
        }
    }
    var v41 int16 = v26
    fmt.Println(v41, v1)
    v27 += v27
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
