package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int64; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 int64 = 1
    var v2 int32 = 321
    var v3 uint16 = 32767
    v4 := make(map[int]string)
    if true {
        if true {
            var v5 int16 = 1
            _ = v5
        }
        v6 := GenS1{F0: -1, F1: 42}
        var v7 string = ("" + ("a" + "foo bar"))
        _ = v6
        _ = v7
    } else {
        fmt.Println(v1, v3, v3)
        v8 := GenS2{F0: 18, F1: v1, F2: 0}
        _ = v8
    }
    if false {
        if (true && (65535 > 65535)) {
            switch -2147483648 {
            case 626:
                v2 -= v2
            case -2147483648:
                switch 0 {
                case 127:
                    v9 := GenS1{F0: -944, F1: v3}
                    _ = v9
                }
            }
        } else {
            v4[0] = "foo bar"
            if (true || (uint(v3) <= 255)) {
                v10 := GenS2{F0: 2, F1: (-(v1)), F2: 127}
                _ = v10
            }
        }
        if ((4294967295 == 2147483647) && ((42 > 2) || true)) {
            v1 = v1
        }
    }
    var v11 string = "\n\t"
    v12 := GenS2{F0: 2, F1: v1, F2: 1}
    v3 = v3
    v1 = (v1 % ((v1) | 1))
    v11 = v11
    v13 := []int{0, -219, 161}
    v1 -= 0
    v2 = v2
    switch (v11 + v11) {
    case "\n\t":
        _ = len(v4)
    case "":
        var v14 int = 2147483647
        _ = v14
    case "foo bar":
        for fi15 := 0; fi15 < 14; fi15++ {
            for fi16 := 0; fi16 < 10; fi16++ {
                v11 = "foo bar"
                var v17 uint16 = v3
                _ = len(v13)
                _ = fi16
                _ = v17
            }
            v18 := GenS3{F0: 1, F1: 65535}
            _ = fi15
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 1; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                v21 := GenS1{F0: 127, F1: 255}
                if false {
                    v22 := v4[fi20]
                    v4[(fi19 % (((fi20 - 128)) | 1))] = v22
                    var v23 uint = 65535
                    _ = v22
                    _ = v23
                } else {
                    if (v11 > v11) {
                        v3 = (^(v3))
                        var v24 int32 = (^(v2))
                        _ = v24
                    } else {
                        v2 -= v2
                        var v25 float64 = 0.0
                        _ = v25
                    }
                    var v26 int = int(v1)
                    var v27 string = "a"
                    _ = v26
                    _ = v27
                }
                _ = fi20
                _ = v21
            }
            if false {
                fmt.Println(fi19, fi19, v11)
            }
            _ = fi19
        }
    }
    v2 += (^(v2))
    v2 = 42
    if false {
        for fi28 := 0; fi28 < 13; fi28++ {
            _ = len(v4)
            _ = fi28
        }
        v11 = (v11 + v11)
        for fi29 := 0; fi29 < 18; fi29++ {
            v30 := v4[(^(fi29))]
            delete(v4, fi29)
            for fi31 := 0; fi31 < 5; fi31++ {
                var v32 float32 = 3.14159
                fmt.Println(fi31)
                switch fi29 {
                case 127:
                    var v33 int32 = (v2 / ((v2) | 1))
                    _ = v33
                case 1:
                    var v34 uint32 = 1
                    _ = v34
                }
                v13 = append(v13, 42)
                _ = fi31
                _ = v32
            }
            _ = fi29
            _ = v30
        }
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        for fi36 := 0; fi36 < 15; fi36++ {
            v3 -= 853
            v37 := v4[fi36]
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    v4[9223372036854775807] = v11
    if true {
        switch "hello" {
        case "hello":
            var v38 int8 = -1
            _ = v38
        case "\n\t":
            if (255 > 65535) {
                var v39 int64 = v1
                _ = v39
            } else {
                if false {
                    v40 := GenS1{F0: -9223372036854775808, F1: (v3 | 42)}
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 10; fi41++ {
                        var v42 int32 = v2
                        if true {
                            v13 = append(v13, (fi41 << ((2) & 63)))
                        } else {
                            fmt.Println(v42, v2)
                            v43 := v4[(fi41 & 128)]
                            _ = v43
                        }
                        _ = fi41
                        _ = v42
                    }
                }
            }
        }
        v3 -= (v3 & v3)
        for fi44 := 0; fi44 < 11; fi44++ {
            if false {
                _ = len(v13)
                var v45 uint8 = 128
                _ = v45
            }
            var v46 uint8 = 24
            _ = fi44
            _ = v46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
