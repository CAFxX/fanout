package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int; F1 int16; F2 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := []int{-340, -1, 0}
    var v4 int32 = -2147483648
    v5 := make(map[uint32]string)
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 float64 = (-2.25 * 0.1)
        _ = len(v2)
        if (100 == 127) {
            v9 := GenS1{F0: 42}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    v10 := GenS1{F0: 65535}
    if false {
        var v11 int64 = 652
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS1{F0: 0}
            var v14 uint64 = 1
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 string = ("foo bar" + "")
        _ = v11
        _ = v15
    }
    var v16 float32 = 1e10
    if ((!(42 < 18446744073709551615)) || true) {
        for fi17 := 0; fi17 < 20; fi17++ {
            switch (fi17 / ((fi17) | 1)) {
            case -2147483648:
                v4 = (v4 | (v4 / ((v4) | 1)))
            case 32768:
                v6 -= 42
            case -9223372036854775808:
                var v18 uint16 = 255
                _ = v18
            }
            fmt.Println(v4, v4, v4)
            _ = fi17
        }
        v19 := GenS1{F0: 2}
        _ = v19
    } else {
        if (2147483647 == 2147483647) {
            if (false || ((!(2147483647 <= 312)) || (-317 == 586))) {
                v20 := GenS2{F0: v4, F1: 42, F2: 44}
                _ = v20
            }
            v4 = v4
        }
        if true {
            var v21 uint16 = 32767
            _ = v21
        } else {
            for fi22 := 0; fi22 < 3; fi22++ {
                if false {
                    var v23 int32 = v4
                    _ = v23
                } else {
                    if (42 <= 9223372036854775807) {
                        v4 += -550
                    } else {
                        var v24 uint32 = 682
                        var v25 uint = 255
                        _ = v24
                        _ = v25
                    }
                    var v26 int64 = 1
                    if true {
                        if true {
                            v6 = (v6 & (v6 >> ((909) & 63)))
                            v16 += ((v16 * -2.25) * v16)
                            fmt.Println(v4, v26)
                        }
                        if (22 > (^(v6))) {
                            if len(v3) > 0 {
                                v3[0] = (fi22 << ((1) & 63))
                            }
                            switch -9223372036854775808 {
                            case 827:
                                var v27 float32 = v16
                                _ = v27
                            case 32767:
                                for fi28 := 0; fi28 < 15; fi28++ {
                                    switch ("\n\t" + "\n\t") {
                                    case "foo bar":
                                        var v29 uint8 = (^(v6))
                                        _ = v29
                                    case "\n\t":
                                        v30 := GenS1{F0: uint32(v6)}
                                        _ = v30
                                    }
                                    v31 := GenS3{F0: 65535, F1: 42}
                                    var v32 uint64 = 1
                                    _ = fi28
                                    _ = v31
                                    _ = v32
                                }
                            default:
                                v26 = v26
                            }
                            v33 := GenS2{F0: v4, F1: 601, F2: 2147483647}
                            _ = v33
                        } else {
                            v4 -= 2147483647
                        }
                    } else {
                        if false {
                            v34 := GenS2{F0: (^(v4)), F1: 1000, F2: 65535}
                            _ = v34
                        } else {
                            if ((v6 != 64) || false) {
                                var v35 int64 = (v26 << ((1) & 63))
                                _ = v35
                            }
                            if false {
                                delete(v1, fi22)
                                fmt.Println(v26, v6)
                            }
                        }
                        if true {
                            if true {
                                v6 -= 1
                                v36 := GenS1{F0: 0}
                                v37 := GenS4{F0: 42, F1: fi22}
                                _ = v36
                                _ = v37
                            } else {
                                v38 := GenS5{F0: fi22, F1: 32767, F2: false}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v39 := GenS3{F0: 4294967295, F1: v4}
                                _ = v38
                                _ = v39
                            }
                            var v40 float32 = float32(0.1)
                            v4 -= -814
                            _ = v40
                        } else {
                            if (!(!true)) {
                                var v41 uint = 142
                                _ = len(v2)
                                _ = v41
                            } else {
                                fmt.Println(v4)
                                for fi42 := 0; fi42 < 3; fi42++ {
                                    fmt.Println(v4)
                                    _ = fi42
                                }
                            }
                        }
                        _ = len(v5)
                    }
                    _ = v26
                }
                v16 -= (1.5 - v16)
                switch (fi22 / (((^(fi22))) | 1)) {
                case 32768:
                    v6 = v6
                case 128:
                    if true {
                        v43 := v1[-98]
                        switch (fi22 >> ((4294967295) & 63)) {
                        case 42:
                            for fi44 := 0; fi44 < 14; fi44++ {
                                var v45 int16 = 0
                                fmt.Println(v43, v6)
                                _ = fi44
                                _ = v45
                            }
                        case 2147483647:
                            v5[4294967295] = v43
                        case 9223372036854775807:
                            v46 := GenS1{F0: 2147483648}
                            _ = v46
                        }
                        _ = v43
                    }
                case 1:
                    var v47 int64 = 9223372036854775807
                    _ = v47
                default:
                    v16 -= v16
                }
                _ = fi22
            }
            if true {
                v6 = v6
                var v48 float32 = v16
                if (4294967295 != 4294967295) {
                    var v49 int32 = v4
                    _ = v49
                } else {
                    fmt.Println(v4, v4, v48)
                    v50 := v2[174]
                    _ = v50
                }
                _ = v48
            }
            var v51 uint64 = 1
            _ = v51
        }
    }
    v6 = 42
    v6 = (v6 / ((v6) | 1))
    var v52 float64 = (1.5 / 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
