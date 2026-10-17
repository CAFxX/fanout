package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 bool }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint32; F1 int16; F2 uint16 }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    switch 4294967295 {
    case -2147483648:
        v4 := GenS1{F0: (int64(v3) >= 9223372036854775807)}
        _ = v4
    case -9223372036854775808:
        if (!false) {
            v5 := GenS1{F0: false}
            v6 := GenS2{F0: 1, F1: 2, F2: 114}
            _ = v5
            _ = v6
        }
    default:
        var v7 uint64 = v3
        _ = v7
    }
    v3 -= v3
    v8 := GenS2{F0: -252, F1: 77, F2: (^(v3))}
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 string = "hello"
        if (1 > 42) {
            v10 = ("\n\t" + "foo bar")
        } else {
            if (946 != 877) {
                var v11 int64 = -1
                var v12 int32 = 0
                _ = v11
                _ = v12
            } else {
                var v13 uint16 = 408
                v10 = ("foo bar" + "hello")
                for fi14 := 0; fi14 < 7; fi14++ {
                    v1 = append(v1, fi14)
                    switch "hello" {
                    case "foo bar":
                        v3 = v3
                    default:
                        v15 := GenS1{F0: ((fi14 << ((4294967295) & 63)) == (fi9 | (^(fi14))))}
                        switch v10 {
                        case "\n\t":
                            if (((v13 | 1) >= 32767) && (((3.14159 / (1e10 * 0.1)) / 1e10) > 3.14159)) {
                                fmt.Println(v3, fi14)
                                v13 += 763
                            }
                            v16 := GenS3{F0: 1, F1: -32768}
                            _ = v16
                        case "a":
                            var v17 int8 = 1
                            _ = v17
                        }
                        _ = v15
                    }
                    v18 := GenS4{F0: (v13 << ((773) & 63)), F1: 127, F2: ((v10 + "hello") > "")}
                    _ = fi14
                    _ = v18
                }
                _ = v13
            }
        }
        v10 = v10
        _ = fi9
        _ = v10
    }
    if (340 == 32767) {
        fmt.Println(v3)
        v3 = (v3 + v3)
        v3 = (v3 & (v3 << ((4294967295) & 63)))
    } else {
        if (360 > 65535) {
            for fi19 := 0; fi19 < 16; fi19++ {
                if false {
                    switch (fi19 * 2147483647) {
                    case -2147483648:
                        if (0 == 127) {
                            var v20 float32 = 0.0
                            v21 := GenS4{F0: 849, F1: 1, F2: true}
                            var v22 float32 = (v20 - 1e10)
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    case 9223372036854775807:
                        var v23 bool = (934 < 32767)
                        _ = v23
                    default:
                        var v24 uint16 = 42
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 16; fi25++ {
                        var v26 string = "a"
                        if (!false) {
                            fmt.Println(v26, v26)
                            v1 = append(v1, (fi25 | (fi25 * (^(fi25)))))
                        } else {
                            var v27 int32 = -2147483648
                            _ = v27
                        }
                        if true {
                            var v28 int = fi25
                            var v29 uint8 = 255
                            v30 := GenS5{F0: v26, F1: false}
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            if (255 > 65535) {
                                v31 := GenS2{F0: int32(fi25), F1: 43, F2: v3}
                                v32 := GenS3{F0: 1, F1: -891}
                                v33 := GenS6{F0: 42, F1: 0, F2: 46}
                                var v34 string = v26
                                _ = v31
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            v26 = v26
                        }
                        _ = fi25
                        _ = v26
                    }
                    var v35 int16 = -532
                    _ = v35
                }
                _ = fi19
            }
        } else {
            for fi36 := 0; fi36 < 20; fi36++ {
                if (false && ((32767 == 42) && false)) {
                    var v37 uint = 0
                    var v38 float64 = float64(((1e10 / float32(1e10)) + (3.14159 - (3.14159 + 1.5))))
                    v38 = v38
                    _ = v37
                    _ = v38
                }
                v3 += (v3 - 528)
                _ = fi36
            }
        }
        var v39 uint16 = 0
        _ = v39
    }
    v40 := v2[1]
    for fi41 := 0; fi41 < 3; fi41++ {
        v1 = append(v1, 4294967295)
        v40 = v40
        v42 := GenS3{F0: 255, F1: 1}
        _ = fi41
        _ = v42
    }
    fmt.Println(v40, v40)
    v43 := GenS1{F0: (3.14159 == 0.1)}
    switch ("a" + "") {
    case "\n\t":
        v40 = v40
    case "hello":
        for fi44 := 0; fi44 < 1; fi44++ {
            fmt.Println(v3)
            _ = fi44
        }
    default:
        if v40 {
            var v45 uint16 = 42
            v40 = v40
            var v46 bool = v40
            _ = v45
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 float32 = 0.1
        _ = fi47
        _ = v48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v43
    _ = v8
}
