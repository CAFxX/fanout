package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    var v3 bool = (!true)
    v2 = true
    var v4 uint = 65535
    if true {
        switch -291 {
        case 59:
            switch ("a" + "hello") {
            case "hello":
                if v3 {
                    var v5 int16 = 32767
                    _ = len(v1)
                    _ = v5
                } else {
                    v6 := []int{141, 0, 127}
                    _ = v6
                }
            case "a":
                if true {
                    var v7 int32 = -1
                    _ = v7
                } else {
                    v3 = (v2 && (("hello" != ("a" + "a")) || (true || false)))
                    if ((3.14159 + (1.5 - ((3.14159 - 0.1) / 1e10))) == 1e10) {
                        v8 := GenS1{F0: false, F1: 295, F2: 2}
                        switch 1 {
                        case 32767:
                            v9 := GenS1{F0: v3, F1: 42, F2: 0}
                            _ = v9
                        case 42:
                            for fi10 := 0; fi10 < 11; fi10++ {
                                if v3 {
                                    var v11 bool = (v3 && (!false))
                                    v12 := v1[(-(fi10))]
                                    _ = v11
                                    _ = v12
                                }
                                var v13 int32 = 566
                                _ = fi10
                                _ = v13
                            }
                        case 2147483647:
                            v14 := []int{-1, 0, -2147483648, -1}
                            _ = v14
                        default:
                            v4 = 255
                        }
                        v15 := GenS2{F0: 42}
                        _ = v15
                        _ = v8
                    } else {
                        v16 := GenS2{F0: 0}
                        v4 = v4
                        v2 = (((!v2) || true) || (-1 != 509))
                        _ = v16
                    }
                    v17 := GenS3{F0: 42}
                    _ = v17
                }
            case "":
                v18 := GenS2{F0: -670}
                _ = v18
            default:
                var v19 bool = ((v3 && (!(v2 && false))) && true)
                _ = v19
            }
        case -1:
            v4 += (^(v4))
        case 9223372036854775807:
            v2 = (782 < 1)
        }
        _ = len(v1)
    }
    v20 := GenS1{F0: v3, F1: 2, F2: 1}
    var v21 uint = v4
    if (!v3) {
        if true {
            if (int16(v4) >= 1) {
                v22 := GenS3{F0: 136}
                v23 := GenS4{F0: 0}
                _ = v22
                _ = v23
            }
            v21 += v21
            if true {
                if (0 == 1000) {
                    _ = len(v1)
                    if ((((3.14159 + 1.5) * 1.5) - float64(1e10)) < 3.14159) {
                        for fi24 := 0; fi24 < 8; fi24++ {
                            var v25 float64 = ((0.0 * -2.25) - -2.25)
                            for fi26 := 0; fi26 < 4; fi26++ {
                                if ((v2 && v2) || (-595 != -1)) {
                                    v27 := GenS5{F0: 1}
                                    _ = v27
                                } else {
                                    v28 := []int{}
                                    _ = v28
                                }
                                var v29 uint = (v4 << (((^(v4))) & 63))
                                if (v2 && (!(!v2))) {
                                    var v30 string = ""
                                    fmt.Println(v25, v3, v30)
                                    _ = v30
                                }
                                _ = fi26
                                _ = v29
                            }
                            v4 -= v21
                            _ = fi24
                            _ = v25
                        }
                        delete(v1, 2147483647)
                    }
                } else {
                    switch (("" + (("\n\t" + "hello") + "")) + (("\n\t" + "\n\t") + "a")) {
                    case "hello":
                        v31 := GenS6{F0: v4, F1: v4}
                        _ = v31
                    case "a":
                        var v32 int8 = int8(v21)
                        _ = v32
                    default:
                        v33 := []int{53}
                        _ = v33
                    }
                    v34 := GenS7{F0: 391, F1: (v4 ^ (^(v4)))}
                    for fi35 := 0; fi35 < 6; fi35++ {
                        for fi36 := 0; fi36 < 20; fi36++ {
                            v3 = ((-2.25 / (((3.14159 + float64(float32(-2.25))) * 0.0) + (((float64(0.0) * -2.25) * -2.25) + 0.1))) <= (((1e10 - (0.1 / -2.25)) - 1.5) * 3.14159))
                            _ = fi36
                        }
                        var v37 string = "foo bar"
                        _ = fi35
                        _ = v37
                    }
                    _ = v34
                }
                v38 := []int{4294967295, 1, 127}
                switch -964 {
                case 0:
                    switch 1 {
                    case -9223372036854775808:
                        v4 -= (v4 & (v4 / (((^(v21))) | 1)))
                    case 4294967295:
                        if (-815 <= -2147483648) {
                            v21 = v4
                            if len(v38) > 0 {
                                v38[0] = 4294967295
                            }
                        }
                    case 0:
                        v2 = (2147483648 > 2147483648)
                    default:
                        var v39 uint32 = 42
                        _ = v39
                    }
                }
                _ = v38
            } else {
                v40 := v1[127]
                v41 := GenS2{F0: 2147483647}
                _ = v40
                _ = v41
            }
        }
    }
    v21 = 42
    v2 = v2
    fmt.Println(v3, v3)
    for fi42 := 0; fi42 < 7; fi42++ {
        var v43 int8 = 100
        _ = fi42
        _ = v43
    }
    var v44 int = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v44
}
