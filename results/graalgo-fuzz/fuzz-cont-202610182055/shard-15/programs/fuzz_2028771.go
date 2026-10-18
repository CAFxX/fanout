package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int = 9223372036854775807
    var v3 uint16 = 501
    var v4 uint64 = 2
    var v5 float32 = 1e10
    v6 := []uint32{0, 0}
    if ((float32(0.1) * 3.14159) == ((1e10 / v5) + (float32(((0.1 / -2.25) / 3.14159)) * (v5 + v5)))) {
        if false {
            fmt.Println(v4, v5, v2)
        } else {
            if (!true) {
                for fi7 := 0; fi7 < 4; fi7++ {
                    v4 = uint64(v3)
                    _ = fi7
                }
                fmt.Println(v2)
            } else {
                if false {
                    var v8 uint32 = 2147483648
                    _ = v8
                } else {
                    var v9 uint32 = 1
                    var v10 int8 = -1
                    v11 := GenS1{F0: (v4 % (((^(v4))) | 1)), F1: uint8(v3)}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                var v12 float32 = ((v5 / v5) / 0.1)
                _ = v12
            }
            var v13 bool = true
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            v3 += v3
            var v15 bool = false
            _ = fi14
            _ = v15
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    switch (v2 % ((-2147483648) | 1)) {
    case 9223372036854775807:
        v1["foo bar"] = v2
    case -9223372036854775808:
        for fi16 := 0; fi16 < 9; fi16++ {
            switch "foo bar" {
            case "hello":
                v2 -= (^(fi16))
            case "":
                var v17 int16 = 1000
                _ = v17
            default:
                if true {
                    fmt.Println(fi16, v3, v5)
                }
            }
            v2 += 9223372036854775807
            var v18 uint32 = 4294967295
            _ = fi16
            _ = v18
        }
    case 42:
        fmt.Println(v4, v3)
    }
    v3 += v3
    var v19 int8 = 42
    var v20 int8 = v19
    var v21 int8 = 66
    fmt.Println(v2)
    if true {
        if (v4 <= v4) {
            for fi22 := 0; fi22 < 16; fi22++ {
                v20 += v19
                for fi23 := 0; fi23 < 13; fi23++ {
                    v5 = v5
                    v20 += v20
                    if false {
                        if (4294967295 != (-(fi23))) {
                            fmt.Println(v21, v4, fi22)
                            delete(v1, "\n\t")
                            var v24 uint16 = v3
                            v25 := GenS1{F0: v4, F1: 41}
                            _ = v24
                            _ = v25
                        } else {
                            if false {
                                v1[""] = -34
                                v26 := GenS2{F0: 2}
                                _ = v26
                            }
                        }
                        fmt.Println(fi23, v19)
                        fmt.Println(v2, v20, v4)
                    }
                    _ = fi23
                }
                if true {
                    if len(v6) > 0 {
                        v6[0] = 42
                    }
                    if len(v6) > 0 {
                        v6[0] = 4294967295
                    }
                    if (32767 < (v3 * 2)) {
                        v1["a"] = 843
                    } else {
                        v4 = v4
                        v27 := GenS3{F0: 897, F1: 1}
                        for fi28 := 0; fi28 < 8; fi28++ {
                            var v29 bool = true
                            _ = fi28
                            _ = v29
                        }
                        _ = v27
                    }
                }
                _ = fi22
            }
            v19 -= v21
        } else {
            if true {
                if (!(((v3 | v3) > 485) || (!((true && true) || false)))) {
                    v21 = (v21 / ((v20) | 1))
                    if (((v20 == (v21 << ((215) & 63))) && (!(true || true))) && true) {
                        for fi30 := 0; fi30 < 17; fi30++ {
                            var v31 uint = 526
                            _ = fi30
                            _ = v31
                        }
                        v2 -= (^(v2))
                        v5 = v5
                    }
                }
                if (1 == -1) {
                    v32 := GenS3{F0: 946, F1: -32768}
                    fmt.Println(v4, v5)
                    _ = v32
                } else {
                    v5 = v5
                    fmt.Println(v4, v4, v3)
                }
                fmt.Println(v4, v4)
            } else {
                var v33 int = (v2 + int(v3))
                _ = v33
            }
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int = (-(fi34))
                _ = fi34
                _ = v35
            }
            var v36 uint = 921
            _ = v36
        }
        fmt.Println(v2)
        var v37 uint16 = 0
        _ = v37
    } else {
        var v38 int8 = (v20 + 100)
        var v39 int8 = (v21 + int8(v2))
        _ = v38
        _ = v39
    }
    if false {
        var v40 uint8 = 255
        v41 := v1["foo bar"]
        var v42 bool = false
        _ = v40
        _ = v41
        _ = v42
    }
    var v43 bool = false
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
