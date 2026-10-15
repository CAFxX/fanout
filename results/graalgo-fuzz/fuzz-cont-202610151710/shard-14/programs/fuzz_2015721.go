package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint32; F2 string }
type GenS3 struct { F0 string; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int8; F1 bool }

func main() {
    var v1 uint16 = 255
    var v2 bool = ((42 <= 127) || ((true || true) || false))
    v3 := []string{("foo bar" + "foo bar"), "", ""}
    var v4 string = "hello"
    v5 := make(map[int]uint32)
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 3; fi7++ {
        v4 = (v4 + "foo bar")
        _ = fi7
    }
    v8 := GenS1{F0: v4}
    var v9 uint32 = 238
    v3 = append(v3, (((v4 + (("foo bar" + v4) + v4)) + v4) + (v4 + v4)))
    v10 := v5[-2147483648]
    for fi11 := 0; fi11 < 3; fi11++ {
        switch (^(fi11)) {
        case -1:
            fmt.Println(v10, v2)
        default:
            v3 = append(v3, v4)
        }
        _ = fi11
    }
    v12 := GenS2{F0: 4294967295, F1: (v9 % ((v10) | 1)), F2: v4}
    var v13 int64 = -1
    switch -388 {
    case 32767:
        delete(v5, -1)
    case 129:
        fmt.Println(v10, v10)
    case -2147483648:
        if (!false) {
            if len(v3) > 0 {
                v3[0] = (v4 + ((v4 + "") + "\n\t"))
            }
        } else {
            v14 := GenS3{F0: (v4 + (v4 + v4)), F1: uint32(v1), F2: v1}
            switch "hello" {
            case "a":
                v15 := GenS3{F0: "foo bar", F1: 304, F2: 1}
                _ = v15
            case "\n\t":
                v16 := v5[42]
                _ = v16
            case "hello":
                v1 = v1
            }
            _ = v14
        }
    }
    if (v2 && (v2 || false)) {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS3{F0: (v4 + ((v4 + v4) + "\n\t")), F1: (^(v10)), F2: v1}
            v19 := v5[fi17]
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        switch -844 {
        case 42:
            delete(v5, -561)
        case 2147483647:
            if v2 {
                var v20 float32 = 1e10
                var v21 int64 = v6
                _ = v20
                _ = v21
            } else {
                if v2 {
                    var v22 string = v4
                    for fi23 := 0; fi23 < 5; fi23++ {
                        var v24 uint32 = (v9 * (v10 * v9))
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        }
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := GenS2{F0: fi25, F1: (^(v10)), F2: ((v4 + v4) + v4)}
            _ = fi25
            _ = v26
        }
        if (42 != 18446744073709551615) {
            v27 := GenS2{F0: -460, F1: 2147483648, F2: ""}
            v2 = false
            _ = v27
        } else {
            var v28 uint8 = 9
            if true {
                v3 = append(v3, (((v4 + v4) + (("hello" + v4) + v4)) + (v4 + (v4 + v4))))
                var v29 float64 = -2.25
                _ = v29
            } else {
                v10 -= v9
                if v2 {
                    var v30 uint16 = 677
                    if (-921 >= 860) {
                        fmt.Println(v6, v9)
                        if false {
                            v30 = 42
                            var v31 uint8 = uint8(v9)
                            var v32 uint32 = (^(v9))
                            _ = v31
                            _ = v32
                        } else {
                            if (!(-318 > 1000)) {
                                var v33 string = (v4 + v4)
                                _ = v33
                            } else {
                                v9 = (v9 + (v10 | (^(v9))))
                                v3 = append(v3, (v4 + (v4 + "\n\t")))
                                for fi34 := 0; fi34 < 10; fi34++ {
                                    fmt.Println(v2)
                                    switch (fi34 / (((fi34 - fi34)) | 1)) {
                                    case 2147483647:
                                        v35 := GenS1{F0: "\n\t"}
                                        _ = v35
                                    case -1:
                                        v5[fi34] = 65535
                                    case 127:
                                        v30 = uint16(v10)
                                    default:
                                        fmt.Println(v6, v10)
                                    }
                                    v5[(-(fi34))] = v9
                                    _ = fi34
                                }
                            }
                        }
                    }
                    _ = v30
                }
                var v36 uint8 = 0
                _ = v36
            }
            for fi37 := 0; fi37 < 2; fi37++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v38 uint = 2
                var v39 string = ("a" + v4)
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = v28
        }
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        if (2 != 128) {
            for fi41 := 0; fi41 < 14; fi41++ {
                var v42 int64 = v6
                v10 += v9
                _ = fi41
                _ = v42
            }
            if (!v2) {
                var v43 bool = v2
                v44 := GenS3{F0: v4, F1: 65535, F2: 32767}
                _ = v43
                _ = v44
            }
        } else {
            v13 = v13
            v45 := GenS2{F0: -9223372036854775808, F1: v9, F2: "a"}
            _ = v45
        }
        v46 := GenS4{F0: 14, F1: (!v2)}
        _ = fi40
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
