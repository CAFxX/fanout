package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 65535
    var v3 bool = false
    var v4 int16 = 0
    if false {
        v5 := []int{626, 953, 32767}
        v2 -= (^(v2))
        _ = v5
    } else {
        var v6 uint32 = 42
        v3 = (1 != 9223372036854775807)
        _ = v6
    }
    if ((^(v2)) < v2) {
        v2 += (v2 % ((1) | 1))
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                fmt.Println(v2, fi8, v3)
                v9 := []int{fi7, 589}
                v10 := GenS1{F0: 42, F1: 892}
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := []int{(fi7 & (fi7 * fi7)), fi7, 0}
            v3 = v3
            _ = fi7
            _ = v11
        }
        switch 128 {
        case -146:
            if v3 {
                fmt.Println(v2, v3)
                var v12 float64 = -2.25
                var v13 int8 = 42
                _ = v12
                _ = v13
            } else {
                var v14 int16 = v4
                _ = v14
            }
        case 32768:
            v4 = v4
        default:
            v4 = v4
        }
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 string = ("" + "a")
        v17 := GenS1{F0: 0, F1: 42}
        var v18 bool = ((v2 == 65535) || ((v4 & v4) == (v4 / ((-1) | 1))))
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 int64 = 42
    for fi20 := 0; fi20 < 6; fi20++ {
        v19 = (v19 * 9223372036854775807)
        v2 -= 267
        _ = fi20
    }
    var v21 uint8 = uint8(v2)
    fmt.Println(v19)
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS1{F0: -128, F1: v19}
        for fi24 := 0; fi24 < 7; fi24++ {
            v25 := GenS1{F0: 0, F1: (v19 / (((v19 % (((v19 * (v19 / ((-936) | 1)))) | 1))) | 1))}
            v26 := GenS1{F0: -128, F1: (v19 << ((18446744073709551615) & 63))}
            var v27 int = 1
            v28 := GenS2{F0: v4, F1: -1}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
            _ = v28
        }
        _ = fi22
        _ = v23
    }
    if (1 != (v2 % ((65535) | 1))) {
        var v29 int8 = 42
        if false {
            var v30 int16 = v4
            _ = v30
        }
        _ = v29
    }
    v31 := GenS2{F0: (v4 % (((^(v4))) | 1)), F1: 127}
    var v32 int = 128
    v33 := GenS1{F0: -36, F1: int64(v2)}
    if (v32 == (v32 / ((v32) | 1))) {
        var v34 uint = 1
        v35 := GenS1{F0: 127, F1: v19}
        v19 = (v19 * (v19 + -424))
        _ = v34
        _ = v35
    }
    v36 := GenS2{F0: (v4 * (v4 - v4)), F1: 1}
    v21 -= (v21 << ((950) & 63))
    if v3 {
        switch v32 {
        case 4294967295:
            var v37 bool = v3
            _ = v37
        case 1:
            switch (v32 + int(v21)) {
            case 32767:
                fmt.Println(v2, v4)
            }
        case -154:
            for fi38 := 0; fi38 < 18; fi38++ {
                if (!((v32 / ((v32) | 1)) != 57)) {
                    var v39 int64 = (v19 << ((0) & 63))
                    switch v32 {
                    case 127:
                        v40 := []int{(v32 & fi38), v32}
                        _ = v40
                    case -2147483648:
                        var v41 int64 = 0
                        _ = v41
                    }
                    _ = v39
                } else {
                    if false {
                        var v42 uint8 = 128
                        _ = v42
                    } else {
                        var v43 int16 = (v4 + (v4 >> ((2) & 63)))
                        v4 = 0
                        _ = v43
                    }
                    if (!(1.5 != 1e10)) {
                        var v44 int = fi38
                        var v45 bool = false
                        v46 := v1[644]
                        _ = v44
                        _ = v45
                        _ = v46
                    } else {
                        v21 = (v21 | 255)
                        for fi47 := 0; fi47 < 10; fi47++ {
                            v48 := []int{}
                            _ = fi47
                            _ = v48
                        }
                        if false {
                            v49 := GenS3{F0: v19, F1: 235}
                            _ = v49
                        }
                    }
                }
                v3 = v3
                _ = fi38
            }
        default:
            if (1 <= 18446744073709551615) {
                var v50 int8 = int8(v32)
                _ = v50
            } else {
                v32 += (-(v32))
            }
        }
    }
    var v51 float64 = 3.14159
    v2 += 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v51
}
