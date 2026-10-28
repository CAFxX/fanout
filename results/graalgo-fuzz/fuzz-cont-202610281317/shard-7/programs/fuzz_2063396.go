package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint32; F1 string; F2 int64 }
func genFn0(p0 string) uint {
    v4 := GenS1{F0: 42}
    v5 := []int{128, 42}
    var v6 uint = 154
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return v6
}
func genFn1() bool {
    v7 := GenS2{F0: 2, F1: 267}
    var v8 uint8 = 15
    var v9 int32 = 881
    _ = v7
    _ = v8
    _ = v9
    return (8 == 1)
}

func main() {
    var v1 int32 = -1
    var v2 int64 = 42
    var v3 int8 = 46
    fmt.Println(v2, v3, v3)
    fmt.Println(v2)
    v1 = v1
    var v10 int8 = (-(v3))
    var v11 float32 = 0.1
    var v12 float32 = v11
    var v13 string = "foo bar"
    if (1 != genFn0(v13)) {
        v3 = (v3 & 1)
        v1 += v1
    } else {
        v12 += 3.14159
        if (!false) {
            v2 -= v2
            if false {
                if ((v3 >> ((65535) & 63)) == v3) {
                    if false {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v15 := []int{}
                            for fi16 := 0; fi16 < 9; fi16++ {
                                v1 -= v1
                                if true {
                                    _ = len(v15)
                                    var v17 uint32 = 11
                                    v11 = v12
                                    _ = v17
                                }
                                if (false && (1 != -1)) {
                                    v15 = append(v15, fi16)
                                }
                                _ = fi16
                            }
                            v18 := GenS3{F0: fi14}
                            _ = fi14
                            _ = v15
                            _ = v18
                        }
                        var v19 int16 = -32768
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 2; fi20++ {
                        var v21 int = (-(fi20))
                        v22 := GenS4{F0: 255, F1: 18446744073709551615, F2: 752}
                        if (v21 > 1) {
                            v1 += int32(v10)
                            switch (v21 << ((2) & 63)) {
                            case 4294967295:
                                var v23 int8 = (v3 << ((genFn0("hello")) & 63))
                                _ = v23
                            }
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    var v24 int8 = v10
                    _ = v24
                } else {
                    switch (((v13 + (v13 + "a")) + ("\n\t" + v13)) + "") {
                    case "a":
                        v25 := []int{2147483647, -688}
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 8; fi26++ {
                        v27 := GenS1{F0: fi26}
                        _ = fi26
                        _ = v27
                    }
                }
                switch 0 {
                case 127:
                    if (100 != (v3 | v3)) {
                        v13 = "hello"
                        var v28 string = ((((v13 + v13) + v13) + "") + v13)
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 4; fi29++ {
                            v3 = (v3 - (v10 & (v10 % ((int8(fi29)) | 1))))
                            for fi30 := 0; fi30 < 3; fi30++ {
                                v11 = float32(((1.5 * 1.5) / ((3.14159 + 1e10) + (1.5 + 1.5))))
                                _ = fi30
                            }
                            _ = fi29
                        }
                        v10 = (v10 / ((v10) | 1))
                    }
                case -2147483648:
                    for fi31 := 0; fi31 < 16; fi31++ {
                        var v32 bool = (!(v2 < v2))
                        _ = fi31
                        _ = v32
                    }
                }
                fmt.Println(v1, v2)
            } else {
                fmt.Println(v10, v11, v11)
                for fi33 := 0; fi33 < 2; fi33++ {
                    if false {
                        for fi34 := 0; fi34 < 20; fi34++ {
                            var v35 string = "foo bar"
                            _ = fi34
                            _ = v35
                        }
                        switch (fi33 & (fi33 & fi33)) {
                        case 759:
                            v36 := GenS3{F0: fi33}
                            _ = v36
                        case -141:
                            fmt.Println(v3)
                        default:
                            var v37 float32 = v12
                            _ = v37
                        }
                        switch (fi33 % (((-(fi33))) | 1)) {
                        case 42:
                            v10 = v10
                        }
                    }
                    v3 = (^(v10))
                    var v38 uint64 = 514
                    _ = fi33
                    _ = v38
                }
                if (0 == 32767) {
                    v39 := []int{433, -349, -9223372036854775808, 1}
                    _ = v39
                }
            }
            for fi40 := 0; fi40 < 1; fi40++ {
                switch (fi40 ^ int(v2)) {
                case -920:
                    fmt.Println(v10, v13, v1)
                }
                if false {
                    v41 := []int{fi40, fi40, fi40, 32768}
                    var v42 uint64 = 182
                    v43 := GenS2{F0: 42, F1: (v42 << ((4294967295) & 63))}
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    if true {
                        fmt.Println(v3)
                        v3 = v10
                        v1 += 1
                    } else {
                        fmt.Println(v11, v3)
                        v44 := []int{0, -263, 128, 2147483647}
                        _ = len(v44)
                        _ = v44
                    }
                    var v45 int64 = (v2 & v2)
                    _ = v45
                }
                if true {
                    var v46 float32 = 3.14159
                    v47 := GenS2{F0: 2, F1: 0}
                    _ = v46
                    _ = v47
                }
                _ = fi40
            }
        }
        switch v13 {
        case "":
            fmt.Println(v11, v3, v10)
        case "\n\t":
            var v48 uint64 = 0
            _ = v48
        }
    }
    v49 := GenS2{F0: 65535, F1: 107}
    if false {
        v50 := GenS5{F0: 563, F1: ("hello" + v13), F2: 42}
        var v51 uint64 = 721
        _ = v50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v49
}
