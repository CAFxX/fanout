package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 int32; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 int16 }
func genFn0(p0 int32, p1 string, p2 int8) uint16 {
    v6 := []int{523, -1, 2147483647}
    p0 = (p0 % ((0) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 int64 = 974
    var v2 float64 = (1.5 * (((-2.25 * 3.14159) + 0.1) / float64((((0.1 - 3.14159) * 0.1) / (1.5 - 3.14159)))))
    var v3 int16 = -32768
    var v4 int64 = -47
    var v5 uint16 = 805
    v1 = v4
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint = 0
        _ = fi7
        _ = v8
    }
    var v9 uint = 255
    var v10 uint = 1
    if (!true) {
        var v11 uint32 = 2
        _ = v11
    }
    v12 := GenS1{F0: 128, F1: 255}
    fmt.Println(v10)
    if false {
        var v13 int8 = 1
        v14 := GenS2{F0: "\n\t", F1: 128, F2: (v3 & v3)}
        _ = v13
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := []int{fi15, fi15}
            if len(v16) > 0 {
                v16[0] = int(v3)
            }
            v10 = 255
            _ = fi15
            _ = v16
        }
        var v17 int32 = -2147483648
        var v18 uint = v10
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 5; fi19++ {
            switch (("" + "hello") + ("foo bar" + "foo bar")) {
            case "a":
                v10 = v9
            case "hello":
                v3 += (v3 & (v3 | (v3 * (v3 & v3))))
            }
            switch (fi19 * (fi19 % (((fi19 - (fi19 << ((4294967295) & 63)))) | 1))) {
            case 2147483647:
                for fi20 := 0; fi20 < 17; fi20++ {
                    if false {
                        v21 := []int{32768, fi20, fi19}
                        _ = v21
                    } else {
                        var v22 uint16 = 2
                        _ = v22
                    }
                    v10 = (v9 << ((v10) & 63))
                    _ = fi20
                }
            case -227:
                v3 = v3
            case -2147483648:
                var v23 float64 = v2
                _ = v23
            }
            if (!true) {
                v10 += (v10 ^ (v9 * v10))
            } else {
                if true {
                    fmt.Println(v1, v5, v3)
                    var v24 float32 = 0.1
                    _ = v24
                } else {
                    v25 := GenS3{F0: -2147483648}
                    _ = v25
                }
            }
            _ = fi19
        }
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 bool = ((true && (!(fi26 != (fi26 >> (((v10 & (v10 ^ v9))) & 63))))) && (true && false))
            if v27 {
                v28 := []int{fi26, fi26}
                if (0.0 < v2) {
                    v29 := GenS4{F0: uint16(v9)}
                    var v30 int = fi26
                    var v31 uint = 1
                    _ = v29
                    _ = v30
                    _ = v31
                }
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                _ = v28
            }
            _ = fi26
            _ = v27
        }
    }
    var v32 int = 9223372036854775807
    var v33 uint64 = 1
    switch v32 {
    case 42:
        v34 := []int{}
        _ = v34
    case -2147483648:
        switch "" {
        case "foo bar":
            for fi35 := 0; fi35 < 8; fi35++ {
                v36 := GenS5{F0: 945, F1: int32(v5), F2: v3}
                v2 = 3.14159
                switch -1 {
                case 0:
                    v9 = v10
                case 128:
                    var v37 float64 = v2
                    _ = v37
                case 2147483647:
                    v38 := GenS6{F0: v3}
                    _ = v38
                default:
                    var v39 int32 = -2147483648
                    _ = v39
                }
                _ = fi35
                _ = v36
            }
        }
    case 127:
        if true {
            v3 += v3
        } else {
            if (!false) {
                fmt.Println(v5, v1)
                var v40 uint8 = uint8(v33)
                v10 = (v10 ^ 255)
                _ = v40
            }
            var v41 uint32 = 4294967295
            _ = v41
        }
    default:
        for fi42 := 0; fi42 < 1; fi42++ {
            v4 = 42
            _ = fi42
        }
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        v44 := []int{893, fi43, (^(v32)), 1}
        _ = fi43
        _ = v44
    }
    v5 = (v5 - 255)
    v45 := []int{(v32 % ((v32) | 1)), (v32 >> ((v10) & 63))}
    switch v32 {
    case -366:
        v46 := GenS7{F0: int8(v5), F1: -100}
        _ = v46
    case 32767:
        v32 = (v32 % (((v32 | v32)) | 1))
    case 9223372036854775807:
        var v47 uint16 = (^(v5))
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v45
    _ = v5
    _ = v9
}
