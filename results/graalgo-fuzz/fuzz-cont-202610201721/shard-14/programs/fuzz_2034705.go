package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 bool) uint64 {
    var v5 int = -510
    fmt.Println(p0)
    var v6 uint32 = 2147483647
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 bool = (!(!true))
    v2 := make(map[uint32]string)
    var v3 uint64 = 233
    var v4 int64 = 42
    v3 = v3
    var v7 uint64 = (v3 ^ 18446744073709551615)
    switch 626 {
    case 128:
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 bool = true
            _ = fi8
            _ = v9
        }
    default:
        v1 = v1
    }
    switch 0 {
    case 128:
        if v1 {
            var v10 uint = uint(v3)
            var v11 uint8 = 42
            if v1 {
                if (v11 != v11) {
                    var v12 uint = (^(v10))
                    if (!(-9223372036854775808 == v4)) {
                        v13 := GenS1{F0: 9223372036854775807, F1: v4}
                        _ = v13
                    }
                    if (-2147483648 > 0) {
                        v14 := GenS2{F0: 9223372036854775807, F1: 233, F2: 42}
                        v12 = v10
                        var v15 int64 = v4
                        _ = v14
                        _ = v15
                    }
                    _ = v12
                } else {
                    if (-1 != 42) {
                        var v16 string = "\n\t"
                        _ = v16
                    }
                    var v17 bool = v1
                    _ = v17
                }
                switch 948 {
                case 1:
                    for fi18 := 0; fi18 < 13; fi18++ {
                        var v19 int16 = -285
                        switch fi18 {
                        case -974:
                            var v20 int16 = v19
                            _ = v20
                        case 0:
                            switch (fi18 >> ((v10) & 63)) {
                            case -736:
                                var v21 bool = (((v19 & v19) < (v19 % ((v19) | 1))) || false)
                                _ = v21
                            case 127:
                                v22 := GenS3{F0: "", F1: 1}
                                _ = v22
                            default:
                                var v23 uint64 = uint64(v11)
                                _ = v23
                            }
                        }
                        if v1 {
                            var v24 string = ("hello" + "hello")
                            _ = v24
                        }
                        _ = fi18
                        _ = v19
                    }
                case 252:
                    v25 := v2[2147483647]
                    _ = v25
                case 0:
                    v4 = v4
                }
            } else {
                v11 += 128
            }
            _ = v10
            _ = v11
        } else {
            if (v1 || v1) {
                for fi26 := 0; fi26 < 17; fi26++ {
                    v4 = (v4 >> ((18446744073709551615) & 63))
                    _ = fi26
                }
            }
            _ = len(v2)
        }
    default:
        var v27 int8 = -71
        _ = v27
    }
    switch 4294967295 {
    case 546:
        var v28 uint8 = 2
        _ = v28
    case 0:
        if ((288 >= 65535) && (-869 >= 1)) {
            switch "\n\t" {
            case "":
                fmt.Println(v4, v7)
            case "\n\t":
                if v1 {
                    var v29 uint8 = 42
                    var v30 string = "a"
                    v31 := GenS1{F0: 2147483647, F1: -9223372036854775808}
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    v1 = false
                    v1 = (719 != v7)
                    v32 := v2[1]
                    _ = v32
                }
            default:
                v33 := GenS1{F0: 1, F1: v4}
                _ = v33
            }
            v34 := GenS3{F0: "", F1: 2}
            _ = v34
        } else {
            v35 := GenS4{F0: 2147483647, F1: v3}
            v3 = v3
            switch "" {
            case "\n\t":
                _ = len(v2)
            }
            _ = v35
        }
    case 128:
        v36 := GenS5{F0: (v3 >> ((42) & 63)), F1: (v3 | genFn0(((v3 - v3) >= v3))), F2: 2147483648}
        _ = v36
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        v38 := GenS1{F0: (fi37 ^ 127), F1: v4}
        _ = fi37
        _ = v38
    }
    switch "a" {
    case "hello":
        var v39 float64 = 0.1
        _ = v39
    case "foo bar":
        switch 627 {
        case 32767:
            v40 := []int{584}
            _ = v40
        }
    default:
        v41 := GenS6{F0: 111}
        _ = v41
    }
    v3 = (v3 % ((4294967295) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
