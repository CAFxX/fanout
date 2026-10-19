package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint64; F2 string }

func main() {
    var v1 int = 0
    var v2 uint = 42
    var v3 int = (v1 >> ((v2) & 63))
    var v4 float64 = -2.25
    var v5 uint32 = 2147483648
    var v6 bool = (false || true)
    if (!((v3 + (v1 - int(v5))) <= (^(v1)))) {
        v1 -= (-(v1))
        if (!(-9223372036854775808 == -1)) {
            for fi7 := 0; fi7 < 15; fi7++ {
                v8 := []int{}
                _ = fi7
                _ = v8
            }
            v3 -= (v3 + v3)
            v9 := GenS1{F0: -1}
            _ = v9
        }
        fmt.Println(v4, v2)
    } else {
        if v6 {
            if (((((float32((float64(((0.0 / 3.14159) - 1.5)) - (float64((0.0 - 1.5)) + v4))) + 0.1) * (1e10 + (0.1 * ((1.5 * 3.14159) - -2.25)))) + 1e10) + ((float32(v4) + 1.5) * float32(3.14159))) > 0.1) {
                for fi10 := 0; fi10 < 3; fi10++ {
                    v4 = v4
                    v4 += v4
                    _ = fi10
                }
                var v11 int64 = 42
                _ = v11
            } else {
                var v12 uint32 = (v5 - (v5 & (v5 ^ v5)))
                v13 := []int{v1}
                v13 = append(v13, v1)
                _ = v12
                _ = v13
            }
            v14 := GenS2{F0: v5, F1: (v3 ^ int(v5))}
            v15 := GenS1{F0: 0}
            _ = v14
            _ = v15
        } else {
            v16 := GenS2{F0: 0, F1: 165}
            if ((false || (v5 >= v5)) && (v2 < v2)) {
                v2 = v2
            }
            _ = v16
        }
        if v6 {
            var v17 string = ""
            v18 := GenS1{F0: -2147483648}
            _ = v17
            _ = v18
        } else {
            if false {
                var v19 float32 = (float32(v4) - ((1.5 * 1.5) * ((0.0 / 1.5) + 0.1)))
                fmt.Println(v1)
                _ = v19
            } else {
                var v20 int16 = 1000
                fmt.Println(v6)
                v1 = v1
                _ = v20
            }
            var v21 int8 = 1
            _ = v21
        }
    }
    v22 := []int{(^(v3)), v1, v3}
    var v23 int32 = 42
    var v24 uint8 = 52
    var v25 uint32 = (v5 << ((0) & 63))
    var v26 uint32 = 0
    fmt.Println(v23, v26)
    var v27 int = (v3 * -622)
    v3 += 32767
    v1 = 127
    v22 = append(v22, -2147483648)
    fmt.Println(v6, v5, v26)
    fmt.Println(v24)
    var v28 uint8 = v24
    var v29 float32 = (((3.14159 - (0.1 * 0.1)) / (3.14159 * 1e10)) / (0.1 - ((0.0 + (3.14159 - (-2.25 / (1e10 + -2.25)))) + 0.1)))
    switch v3 {
    case 0:
        v22 = append(v22, v3)
    case 32768:
        var v30 int8 = -62
        _ = v30
    case 2147483647:
        v31 := GenS3{F0: 4294967295, F1: -9223372036854775808}
        _ = v31
    default:
        switch v3 {
        case 42:
            v32 := GenS1{F0: (v23 << ((uint(v24)) & 63))}
            _ = v32
        case -1:
            for fi33 := 0; fi33 < 9; fi33++ {
                switch v1 {
                case -55:
                    for fi34 := 0; fi34 < 11; fi34++ {
                        v35 := GenS4{F0: fi33}
                        var v36 uint = v2
                        v5 -= 65535
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                case 0:
                    if true {
                        if v6 {
                            var v37 uint8 = v24
                            for fi38 := 0; fi38 < 11; fi38++ {
                                var v39 uint16 = 2
                                v40 := GenS5{F0: 42, F1: uint64(v25), F2: "hello"}
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                            _ = v37
                        }
                    } else {
                        var v41 int32 = v23
                        v42 := GenS5{F0: 0, F1: 2, F2: "\n\t"}
                        _ = v41
                        _ = v42
                    }
                }
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                var v43 uint64 = 263
                _ = fi33
                _ = v43
            }
        case 32767:
            v24 = uint8(v2)
        default:
            switch (^(v27)) {
            case 4294967295:
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            default:
                if v6 {
                    fmt.Println(v2)
                } else {
                    var v44 int = v3
                    _ = v44
                }
            }
        }
    }
    var v45 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
