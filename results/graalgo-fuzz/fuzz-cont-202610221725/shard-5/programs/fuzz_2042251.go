package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int16 }

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]uint32)
    v3 := []int{}
    var v4 int64 = 9223372036854775807
    v5 := []string{"foo bar"}
    v6 := []string{"", ("" + (("foo bar" + (("foo bar" + "") + "a")) + "a")), (("\n\t" + "a") + "")}
    v7 := GenS1{F0: 110, F1: 255, F2: -128}
    var v8 uint8 = 2
    v9 := GenS1{F0: 42, F1: (v8 * 255), F2: -128}
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(fi10)
        for fi11 := 0; fi11 < 3; fi11++ {
            if true {
                v12 := GenS2{F0: 833, F1: 65535, F2: int32(v4)}
                if false {
                    var v13 float32 = 0.1
                    _ = v13
                } else {
                    _ = len(v5)
                    if true {
                        v8 = v8
                        switch (fi10 & (-(fi11))) {
                        case -9223372036854775808:
                            switch 9223372036854775807 {
                            case 4294967295:
                                _ = len(v5)
                            default:
                                var v14 int64 = v4
                                _ = v14
                            }
                        case 42:
                            switch "a" {
                            case "\n\t":
                                fmt.Println(fi11, v8)
                            case "":
                                v15 := GenS1{F0: 42, F1: (v8 >> ((1) & 63)), F2: 39}
                                if ((1 > 2) && false) {
                                    v16 := GenS2{F0: 0, F1: 42, F2: -2147483648}
                                    _ = v16
                                } else {
                                    v17 := GenS1{F0: 42, F1: 240, F2: 100}
                                    _ = len(v2)
                                    _ = v17
                                }
                                v8 -= 0
                                switch fi11 {
                                case 32768:
                                    var v18 int = (v1 + fi11)
                                    _ = v18
                                }
                                _ = v15
                            case "hello":
                                if (!(100 != 27)) {
                                    var v19 uint = 1
                                    _ = v19
                                }
                            }
                        case 127:
                            v8 = 178
                        default:
                            fmt.Println(fi10)
                        }
                        v20 := v2["foo bar"]
                        _ = v20
                    } else {
                        v4 = v4
                        v1 = 42
                        var v21 bool = (!((319 == 2147483647) && (1 == 681)))
                        if (!((v8 < 75) && v21)) {
                            _ = len(v5)
                            switch fi10 {
                            case -604:
                                var v22 int = fi11
                                _ = v22
                            case 32767:
                                v23 := GenS1{F0: 18446744073709551615, F1: 212, F2: 100}
                                _ = v23
                            case 0:
                                fmt.Println(v1, fi10)
                                v24 := GenS3{F0: 0}
                                v25 := GenS1{F0: 4294967295, F1: (^(v8)), F2: 112}
                                var v26 uint64 = 42
                                v27 := v2["a"]
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                            default:
                                _ = len(v5)
                            }
                        }
                        _ = v21
                    }
                }
                _ = v12
            }
            v4 += (v4 & v4)
            if false {
                v1 -= (fi11 & (fi11 ^ (^(fi11))))
                fmt.Println(v4, fi11, fi10)
            } else {
                v28 := GenS3{F0: 571}
                v29 := GenS2{F0: 4294967295, F1: 42, F2: -320}
                _ = v28
                _ = v29
            }
            _ = fi11
        }
        _ = fi10
    }
    switch "a" {
    case "":
        var v30 float32 = float32((-2.25 + 0.0))
        _ = v30
    case "hello":
        var v31 uint16 = 65535
        _ = v31
    default:
        var v32 int16 = -1
        _ = v32
    }
    switch (v1 << ((42) & 63)) {
    case -2147483648:
        switch 42 {
        case -68:
            fmt.Println(v1)
        case 32767:
            var v33 int8 = -128
            _ = v33
        default:
            if (!false) {
                v1 -= (v1 & v1)
                v1 = (v1 << ((4294967295) & 63))
            } else {
                v34 := GenS4{F0: 740}
                if len(v6) > 0 {
                    v6[0] = "foo bar"
                }
                _ = v34
            }
        }
    case 128:
        v8 = v8
    default:
        switch (("foo bar" + "a") + "a") {
        case "":
            v1 -= (v1 >> ((2) & 63))
        case "a":
            v4 += v4
        default:
            v35 := GenS5{F0: 127, F1: 32767}
            _ = v35
        }
    }
    var v36 int16 = 1000
    if (!true) {
        var v37 int8 = 42
        v4 += 953
        var v38 int32 = 0
        _ = v37
        _ = v38
    } else {
        var v39 uint = 478
        v4 += 1
        _ = v39
    }
    if false {
        fmt.Println(v1, v1)
        v40 := GenS2{F0: 4294967295, F1: 1, F2: -2147483648}
        _ = v40
    }
    var v41 int8 = 1
    v36 -= (v36 % (((-(v36))) | 1))
    _ = len(v3)
    v8 += (v8 & (^(v8)))
    v36 += (v36 * v36)
    var v42 int64 = int64(v1)
    for fi43 := 0; fi43 < 8; fi43++ {
        v42 = v42
        _ = fi43
    }
    v41 -= v41
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
