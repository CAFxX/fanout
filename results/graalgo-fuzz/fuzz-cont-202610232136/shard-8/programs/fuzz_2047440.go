package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 int32; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]string)
    if (2 == 255) {
        for fi5 := 0; fi5 < 12; fi5++ {
            fmt.Println(v3)
            v6 := GenS1{F0: 307, F1: ((true || false) && ((true || true) && (803 != 42))), F2: 1}
            fmt.Println(v3, fi5)
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            if (65535 < 913) {
                var v8 int8 = -49
                var v9 uint32 = 1
                v8 = (-(v8))
                _ = v8
                _ = v9
            }
            delete(v4, 65535)
            _ = fi7
        }
        var v10 string = ("a" + "hello")
        _ = v10
    } else {
        if true {
            v11 := []int{}
            _ = v11
        } else {
            v3 = v3
            v3 += (v3 + (v3 ^ (v3 - 42)))
            switch "a" {
            case "a":
                v12 := GenS1{F0: 221, F1: (2 >= 42), F2: 32767}
                _ = v12
            case "":
                for fi13 := 0; fi13 < 1; fi13++ {
                    v14 := GenS2{F0: ((!(32767 <= 42)) && false)}
                    _ = fi13
                    _ = v14
                }
            case "hello":
                var v15 uint = 433
                _ = v15
            default:
                v3 = (v3 << ((65535) & 63))
            }
        }
    }
    v3 = v3
    fmt.Println(v3)
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int16 = 516
        v18 := []int{32767}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := []int{32768, 32768, 127}
    var v20 uint32 = 2
    switch -1 {
    case -777:
        v21 := GenS3{F0: 691, F1: v20}
        _ = v21
    case -564:
        v22 := GenS2{F0: ((!(65535 >= v20)) || (!(3.14159 >= float32(-2.25))))}
        _ = v22
    }
    fmt.Println(v20)
    var v23 int16 = -705
    v23 += (v23 << ((255) & 63))
    _ = len(v4)
    var v24 int16 = (^(v23))
    if (v20 >= (v20 | v20)) {
        if true {
            v25 := GenS2{F0: (!false)}
            delete(v2, 2147483647)
            for fi26 := 0; fi26 < 9; fi26++ {
                if false {
                    var v27 bool = (127 != 13)
                    if v27 {
                        v20 = 90
                        var v28 int64 = -9223372036854775808
                        v29 := GenS4{F0: (-(v24))}
                        _ = v28
                        _ = v29
                    } else {
                        switch (fi26 << ((221) & 63)) {
                        case 42:
                            for fi30 := 0; fi30 < 1; fi30++ {
                                if v27 {
                                    v19 = append(v19, (^(fi30)))
                                    v31 := GenS1{F0: -1, F1: false, F2: (fi30 >> ((2) & 63))}
                                    v23 = (^(v23))
                                    _ = v31
                                }
                                if true {
                                    fmt.Println(v20, fi26, fi26)
                                } else {
                                    if len(v19) > 0 {
                                        v19[0] = (fi26 + fi30)
                                    }
                                }
                                _ = fi30
                            }
                        case 32768:
                            for fi32 := 0; fi32 < 19; fi32++ {
                                switch (fi26 / ((int(v23)) | 1)) {
                                case 1:
                                    if v27 {
                                        var v33 float64 = ((float64(1e10) - (((0.0 / 3.14159) + (0.1 + 1.5)) + 1e10)) / -2.25)
                                        var v34 string = (("\n\t" + "foo bar") + ("" + "\n\t"))
                                        var v35 string = (v34 + (("" + ((v34 + v34) + "\n\t")) + "a"))
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                case 32768:
                                    v36 := v4[(v20 % ((2147483647) | 1))]
                                    _ = v36
                                default:
                                    fmt.Println(v27)
                                }
                                var v37 uint = 65535
                                _ = fi32
                                _ = v37
                            }
                        default:
                            v38 := GenS5{F0: (!(("\n\t" > ("\n\t" + "")) || (889 <= 1))), F1: -2147483648, F2: 4294967295}
                            _ = v38
                        }
                    }
                    _ = v27
                }
                _ = fi26
            }
            _ = v25
        }
        v23 += -76
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
