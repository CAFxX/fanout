package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint; F1 int; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint }

func main() {
    var v1 int64 = 950
    var v2 bool = (100 == 32)
    v3 := make(map[uint32]uint32)
    v4 := []int{42, 1}
    var v5 uint8 = 0
    v6 := []string{"\n\t"}
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 0, F1: "\n\t"}
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 91
    v6 = append(v6, ("\n\t" + "\n\t"))
    if v2 {
        switch "\n\t" {
        case "":
            v5 = (^(v5))
        case "a":
            for fi10 := 0; fi10 < 14; fi10++ {
                for fi11 := 0; fi11 < 3; fi11++ {
                    v12 := GenS1{F0: 1, F1: ((("foo bar" + "foo bar") + "") + ("hello" + ""))}
                    var v13 uint8 = v5
                    var v14 uint16 = 373
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi10
            }
        case "foo bar":
            v4 = append(v4, 0)
        }
    } else {
        v15 := GenS1{F0: v1, F1: "\n\t"}
        if v2 {
            fmt.Println(v2, v5)
            v9 += (v9 / (((v9 * v9)) | 1))
        }
        v2 = ("a" == ("" + "foo bar"))
        _ = v15
    }
    switch "hello" {
    case "\n\t":
        v16 := GenS1{F0: int64(v9), F1: "a"}
        _ = v16
    case "hello":
        var v17 uint8 = 99
        _ = v17
    }
    v18 := GenS2{F0: 18446744073709551615, F1: 128, F2: 9223372036854775807}
    v19 := GenS3{F0: true}
    fmt.Println(v2)
    if (("foo bar" + "\n\t") >= "\n\t") {
        var v20 float32 = 3.14159
        v21 := GenS3{F0: (65535 > 65535)}
        if len(v4) > 0 {
            v4[0] = -2147483648
        }
        _ = v20
        _ = v21
    }
    switch "" {
    case "\n\t":
        fmt.Println(v2, v1, v2)
    }
    var v22 uint8 = (v5 % (((v5 >> ((42) & 63))) | 1))
    var v23 uint16 = 1
    if (2 <= 65535) {
        v1 -= (v1 >> ((4294967295) & 63))
        v9 = 2
    } else {
        fmt.Println(v1)
        switch 0 {
        case 447:
            var v24 float64 = 1e10
            _ = v24
        }
        if ((v23 * v23) <= v9) {
            var v25 uint = 4294967295
            var v26 int64 = (v1 << (((v25 << ((v25) & 63))) & 63))
            var v27 uint64 = 0
            _ = v25
            _ = v26
            _ = v27
        } else {
            v28 := GenS2{F0: 255, F1: -9223372036854775808, F2: 947}
            for fi29 := 0; fi29 < 6; fi29++ {
                switch (fi29 * 257) {
                case 128:
                    for fi30 := 0; fi30 < 14; fi30++ {
                        if (-779 < -1) {
                            var v31 string = "\n\t"
                            var v32 int8 = 127
                            switch v31 {
                            case "":
                                if (v31 != "a") {
                                    v1 = v1
                                } else {
                                    if ((-34 <= (v32 * (v32 << ((397) & 63)))) && (false || v2)) {
                                        var v33 uint = 706
                                        if len(v6) > 0 {
                                            v6[0] = "hello"
                                        }
                                        _ = v33
                                    }
                                    var v34 uint64 = 1
                                    v35 := GenS4{F0: (v32 + v32), F1: 0, F2: 2}
                                    _ = v34
                                    _ = v35
                                }
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v36 int = (fi29 >> ((907) & 63))
                                _ = v36
                            }
                            _ = v31
                            _ = v32
                        }
                        v22 -= 2
                        _ = fi30
                    }
                case 813:
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v38 := GenS5{F0: 1}
                        _ = fi37
                        _ = v38
                    }
                }
                var v39 int = 0
                v2 = v2
                _ = fi29
                _ = v39
            }
            _ = v28
        }
    }
    if (!v2) {
        fmt.Println(v1, v22)
    } else {
        v2 = v2
    }
    var v40 int16 = -24
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
