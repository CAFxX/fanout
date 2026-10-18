package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int16; F2 int32 }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0(p0 uint8) uint32 {
    v7 := []int{128}
    switch ("foo bar" + ("hello" + "")) {
    case "\n\t":
        for fi8 := 0; fi8 < 4; fi8++ {
            if ((141 >= 0) || true) {
                var v9 int16 = int16(fi8)
                v10 := GenS1{F0: 2, F1: (p0 | p0), F2: 100}
                var v11 uint16 = 65535
                _ = v10
                _ = v11
                _ = v9
            } else {
                var v12 int = (fi8 % ((fi8) | 1))
                v12 -= 2147483647
                p0 = (p0 ^ (p0 % ((p0) | 1)))
                _ = v12
            }
            if len(v7) > 0 {
                v7[0] = fi8
            }
            if (p0 > (^(p0))) {
                _ = len(v7)
                switch "" {
                case "\n\t":
                    var v13 uint8 = (p0 ^ p0)
                    _ = v13
                case "hello":
                    fmt.Println(p0)
                case "foo bar":
                    for fi14 := 0; fi14 < 18; fi14++ {
                        if false {
                            fmt.Println(fi14)
                            var v15 uint32 = 2147483648
                            var v16 int32 = -883
                            _ = v15
                            _ = v16
                        } else {
                            var v17 uint8 = (p0 >> ((uint(p0)) & 63))
                            var v18 uint = uint(p0)
                            _ = v17
                            _ = v18
                        }
                        var v19 uint8 = p0
                        var v20 string = ""
                        _ = fi14
                        _ = v19
                        _ = v20
                    }
                default:
                    fmt.Println(p0, p0)
                }
            } else {
                var v21 int8 = 100
                v21 = (^(v21))
                _ = v21
            }
            _ = fi8
        }
    case "":
        v22 := GenS2{F0: 311}
        _ = v22
    }
    _ = p0
    _ = v7
    return 2
}

func main() {
    var v1 int8 = -1
    var v2 uint32 = 65535
    v3 := make(map[int]string)
    var v4 float32 = -2.25
    var v5 int = 802
    var v6 uint = 17
    if true {
        var v23 float64 = -2.25
        _ = len(v3)
        _ = v23
    } else {
        fmt.Println(v6, v1)
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := []int{32768}
            v26 := v3[v5]
            var v27 uint64 = 2
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    }
    var v28 bool = false
    if v28 {
        fmt.Println(v6, v6)
        var v29 int64 = -563
        _ = v29
    } else {
        if (!(42 == 2147483647)) {
            delete(v3, 50)
            var v30 float32 = 0.0
            _ = v30
        }
        var v31 bool = false
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 uint64 = 0
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        fmt.Println(v4)
        v35 := []int{v5, fi34, (^(fi34)), (^(fi34))}
        _ = fi34
        _ = v35
    }
    v3[v5] = ("a" + ("hello" + "foo bar"))
    if (32767 > -32768) {
        v36 := []int{}
        if false {
            if (0.0 < 0.1) {
                var v37 int32 = int32(v1)
                switch (v5 >> (((^(v6))) & 63)) {
                case 0:
                    if len(v36) > 0 {
                        v36[0] = v5
                    }
                case 2147483647:
                    fmt.Println(v1, v28)
                case -1:
                    v5 = 128
                }
                _ = v37
            } else {
                switch 289 {
                case 613:
                    v5 += (v5 % (((v5 / (((v5 * (v5 << ((v6) & 63)))) | 1))) | 1))
                case 1:
                    v38 := GenS1{F0: 18446744073709551615, F1: 2, F2: v1}
                    _ = v38
                }
                var v39 uint8 = 164
                var v40 uint = v6
                _ = v39
                _ = v40
            }
            delete(v3, (v5 & 32768))
        }
        _ = v36
    }
    if false {
        for fi41 := 0; fi41 < 18; fi41++ {
            v42 := []int{fi41, 128, v5, v5}
            _ = fi41
            _ = v42
        }
        fmt.Println(v1, v28, v1)
        for fi43 := 0; fi43 < 15; fi43++ {
            var v44 uint32 = (v2 % ((1) | 1))
            var v45 int = int(v1)
            v46 := GenS1{F0: 964, F1: 42, F2: v1}
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
    } else {
        var v47 int8 = v1
        _ = v47
    }
    v48 := []int{}
    v3[v5] = ""
    if (42 >= 0) {
        v6 += 42
        for fi49 := 0; fi49 < 1; fi49++ {
            fmt.Println(v6, v5)
            v50 := GenS3{F0: 69, F1: -669, F2: -787}
            _ = fi49
            _ = v50
        }
    }
    v51 := GenS4{F0: v1, F1: 1}
    if len(v48) >= 2 {
        v48 = v48[1:len(v48)]
    }
    delete(v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v51
    _ = v6
}
