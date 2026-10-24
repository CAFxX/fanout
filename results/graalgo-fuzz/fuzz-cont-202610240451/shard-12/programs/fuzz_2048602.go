package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 int8 }
type GenS4 struct { F0 int; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint8 }

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]bool)
    var v3 float32 = 0.1
    var v4 uint16 = 255
    var v5 uint32 = 42
    var v6 float32 = 1.5
    v7 := GenS1{F0: v1}
    var v8 uint = 0
    for fi9 := 0; fi9 < 4; fi9++ {
        switch (-(fi9)) {
        case -9223372036854775808:
            switch ("" + "hello") {
            case "a":
                v10 := GenS1{F0: v5}
                _ = v10
            }
        case 4294967295:
            switch "\n\t" {
            case "":
                v1 = v1
            case "a":
                v11 := GenS1{F0: 2147483647}
                _ = v11
            case "foo bar":
                if false {
                    var v12 float64 = 3.14159
                    var v13 int64 = 554
                    _ = v12
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        fmt.Println(v8)
                        var v15 uint8 = 127
                        if true {
                            v8 -= v8
                            if (false && (false && false)) {
                                switch int(v15) {
                                case 472:
                                    v2[(fi9 | 450)] = (-2147483648 >= -2147483648)
                                case -2147483648:
                                    v16 := v2[-9223372036854775808]
                                    _ = v16
                                default:
                                    var v17 uint32 = (v5 * (v5 ^ v5))
                                    _ = v17
                                }
                                var v18 int = (fi14 | fi14)
                                var v19 float32 = v6
                                var v20 int32 = 0
                                var v21 float64 = (-2.25 - 0.0)
                                var v22 float64 = v21
                                _ = v18
                                _ = v19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        } else {
                            if ((v5 - v5) == v1) {
                                if true {
                                    switch fi9 {
                                    case 4294967295:
                                        fmt.Println(v1)
                                    case 9223372036854775807:
                                        v15 = v15
                                    case 125:
                                        var v23 uint8 = v15
                                        _ = v23
                                    default:
                                        fmt.Println(v4)
                                    }
                                    delete(v2, fi9)
                                    var v24 uint = (v8 * (v8 - (^(v8))))
                                    _ = v24
                                } else {
                                    v25 := GenS2{F0: 185}
                                    _ = v25
                                }
                            }
                        }
                        _ = fi14
                        _ = v15
                    }
                    fmt.Println(v5, fi9, v4)
                }
            }
        case 748:
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := GenS3{F0: v4, F1: fi9, F2: 0}
                _ = fi26
                _ = v27
            }
        }
        if (!false) {
            v28 := GenS3{F0: (v4 | v4), F1: (fi9 * (-(fi9))), F2: 42}
            fmt.Println(v5, v8)
            _ = v28
        } else {
            if (-1 < 1) {
                fmt.Println(v1, v4, v1)
                v29 := []int{fi9, -114}
                switch "" {
                case "foo bar":
                    v1 = v1
                case "hello":
                    for fi30 := 0; fi30 < 14; fi30++ {
                        v6 = float32(1e10)
                        _ = fi30
                    }
                }
                _ = v29
            }
            v31 := v2[fi9]
            var v32 int = (fi9 << ((v8) & 63))
            _ = v31
            _ = v32
        }
        switch "a" {
        case "":
            var v33 uint64 = 160
            _ = v33
        default:
            v2[fi9] = (false || true)
        }
        _ = fi9
    }
    v34 := GenS1{F0: 42}
    var v35 uint32 = v1
    v5 = (v5 & v35)
    for fi36 := 0; fi36 < 16; fi36++ {
        v37 := []int{2147483647, fi36}
        v38 := GenS1{F0: 2}
        if true {
            fmt.Println(fi36, v4, v1)
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 string = "hello"
    for fi40 := 0; fi40 < 13; fi40++ {
        var v41 float32 = 1.5
        var v42 bool = (true || ((-32768 < -426) || (1.5 == v3)))
        _ = fi40
        _ = v41
        _ = v42
    }
    v43 := GenS4{F0: 4294967295, F1: 57, F2: 2}
    var v44 uint32 = 65535
    for fi45 := 0; fi45 < 17; fi45++ {
        switch ((v39 + "foo bar") + ("foo bar" + "\n\t")) {
        case "a":
            for fi46 := 0; fi46 < 4; fi46++ {
                if true {
                    v47 := []int{(fi46 & fi46), 32767, fi45}
                    var v48 uint64 = 4294967295
                    fmt.Println(v3, v35)
                    _ = v47
                    _ = v48
                } else {
                    v49 := GenS5{F0: 42, F1: 190}
                    _ = v49
                }
                v4 = (v4 & v4)
                v44 += v35
                v50 := GenS1{F0: uint32(v4)}
                var v51 uint16 = (v4 << ((42) & 63))
                v8 += (v8 ^ (v8 | 65535))
                _ = fi46
                _ = v50
                _ = v51
            }
        case "foo bar":
            v52 := []int{fi45}
            _ = v52
        default:
            v44 = v5
        }
        switch (fi45 | (^(fi45))) {
        case 0:
            v53 := GenS1{F0: (v44 & v35)}
            _ = v53
        case -9223372036854775808:
            var v54 int32 = 2147483647
            _ = v54
        }
        switch v39 {
        case "hello":
            var v55 uint8 = 42
            _ = v55
        case "\n\t":
            var v56 uint32 = (v1 + v35)
            _ = v56
        case "foo bar":
            var v57 uint32 = v35
            _ = v57
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
