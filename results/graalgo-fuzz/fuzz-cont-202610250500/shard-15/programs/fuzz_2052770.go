package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint16 = 65535
    var v4 int32 = 437
    var v5 float32 = 0.0
    var v6 bool = (42 >= -18)
    if v6 {
        v3 = v3
    } else {
        for fi7 := 0; fi7 < 2; fi7++ {
            if (!v6) {
                for fi8 := 0; fi8 < 1; fi8++ {
                    if v6 {
                        if v6 {
                            var v9 int16 = 32767
                            switch fi8 {
                            case -654:
                                v3 -= 102
                            case 128:
                                v10 := v1["hello"]
                                _ = v10
                            }
                            _ = v9
                        } else {
                            if false {
                                fmt.Println(v4)
                                fmt.Println(fi8, v5)
                                var v11 int32 = 0
                                if false {
                                    v6 = v6
                                } else {
                                    _ = len(v2)
                                    v4 = (-(v11))
                                    if v6 {
                                        v4 = (v4 % ((v4) | 1))
                                    } else {
                                        var v12 uint = 65535
                                        _ = v12
                                    }
                                    v13 := GenS1{F0: 1, F1: 2}
                                    if ((v4 % ((v4) | 1)) < (v11 & 42)) {
                                        if false {
                                            v6 = v6
                                        }
                                        delete(v1, "a")
                                        var v14 uint16 = v3
                                        _ = v14
                                    }
                                    var v15 uint16 = v3
                                    _ = v13
                                    _ = v15
                                }
                                if (0 < -32768) {
                                    var v16 int = fi8
                                    _ = v16
                                } else {
                                    if (127 <= 0) {
                                        var v17 int16 = -1
                                        _ = v17
                                    } else {
                                        v6 = v6
                                    }
                                }
                                _ = v11
                            }
                            if (692 != 0) {
                                switch (fi8 | (fi8 - fi8)) {
                                case 32768:
                                    var v18 uint = 899
                                    _ = v18
                                case 0:
                                    v3 -= 65535
                                case 767:
                                    if true {
                                        v19 := GenS2{F0: -1, F1: 128, F2: v4}
                                        v4 = v4
                                        if (("a" + "") > "a") {
                                            fmt.Println(v5, v6)
                                        }
                                        v20 := GenS3{F0: 1, F1: v3}
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v21 := []int{fi7, int(v3), (fi8 | (fi8 - (fi8 / ((fi7) | 1))))}
                                        _ = v21
                                    }
                                }
                                var v22 int32 = -1
                                _ = v22
                            } else {
                                v3 += v3
                                v4 = v4
                                v5 -= (float32(float64(v5)) - v5)
                                v6 = (2147483648 == 1)
                            }
                        }
                        if false {
                            v1["foo bar"] = 1
                            v4 += v4
                        } else {
                            var v23 uint32 = 2
                            fmt.Println(fi7, fi8, v23)
                            _ = v23
                        }
                        v3 = v3
                    } else {
                        var v24 uint8 = 128
                        v6 = (451 <= (v3 << ((802) & 63)))
                        v25 := []int{}
                        _ = v24
                        _ = v25
                    }
                    fmt.Println(v3, fi8)
                    _ = fi8
                }
            } else {
                v4 = -2147483648
                fmt.Println(v4, v4)
            }
            _ = fi7
        }
        v26 := GenS3{F0: 152, F1: v3}
        v6 = (!((1e10 - (float64((3.14159 + 0.0)) + 1.5)) == 0.1))
        _ = v26
    }
    switch -157 {
    case -633:
        v27 := GenS4{F0: -529, F1: 42}
        _ = v27
    case 32768:
        v3 = v3
    case 4294967295:
        var v28 uint32 = 2147483647
        _ = v28
    }
    var v29 uint16 = (v3 & (v3 % ((42) | 1)))
    fmt.Println(v5)
    switch "foo bar" {
    case "a":
        v3 = v29
    case "hello":
        var v30 uint32 = 42
        _ = v30
    case "foo bar":
        switch int(v3) {
        case 738:
            v5 = v5
        case 128:
            var v31 int = 437
            _ = v31
        case -224:
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 uint = 255
                v34 := []int{(fi32 * (fi32 ^ fi32)), (fi32 ^ (-(fi32))), fi32}
                for fi35 := 0; fi35 < 4; fi35++ {
                    fmt.Println(v33)
                    v5 = ((-2.25 / v5) + ((v5 - 3.14159) + ((((((float32(0.1) + 0.1) / v5) + (v5 * (float32(1e10) * float32(-2.25)))) - v5) - float32(0.0)) + float32(0.1))))
                    _ = fi35
                }
                _ = fi32
                _ = v33
                _ = v34
            }
        }
    default:
        v6 = true
    }
    var v36 int16 = -32768
    v37 := GenS2{F0: v36, F1: 42, F2: (v4 * (v4 / ((v4) | 1)))}
    var v38 int8 = 1
    v39 := []int{-489, 128, 91}
    var v40 int8 = v38
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
