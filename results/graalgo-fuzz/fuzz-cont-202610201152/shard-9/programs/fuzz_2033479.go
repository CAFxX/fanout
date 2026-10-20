package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int }

func main() {
    v1 := []uint32{4294967295, 2, 885, 2147483648}
    v2 := make(map[string]bool)
    var v3 uint32 = 1
    var v4 float64 = 0.1
    var v5 uint32 = v3
    var v6 bool = (!(2 <= 0))
    v6 = false
    for fi7 := 0; fi7 < 4; fi7++ {
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            fmt.Println(fi7, v3, v6)
        default:
            if v6 {
                v8 := GenS1{F0: 961, F1: (v5 * (v3 * (v3 * 42)))}
                v9 := GenS1{F0: 2, F1: 2}
                _ = v8
                _ = v9
            }
        }
        var v10 uint64 = 0
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int8 = 0
        if (633 == 255) {
            var v13 uint64 = 1
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v2)
            _ = v13
        } else {
            delete(v2, "")
            fmt.Println(v12, v5)
        }
        if (!true) {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int = -462
                v3 += (^(v3))
                switch fi14 {
                case 545:
                    var v16 bool = v6
                    _ = v16
                case 944:
                    var v17 float32 = (1e10 + (float32(3.14159) + (0.0 - (float32((0.1 - v4)) - 0.1))))
                    _ = v17
                default:
                    if (float64(0.0) < v4) {
                        var v18 int64 = 1
                        var v19 int8 = v12
                        v15 = fi11
                        v5 = v5
                        v6 = ("" <= "\n\t")
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi14
                _ = v15
            }
            var v20 uint = 65535
            _ = v20
        }
        _ = fi11
        _ = v12
    }
    if (1 > 18) {
        if v6 {
            v21 := GenS1{F0: 628, F1: v5}
            if (v6 || (!(!false))) {
                v22 := GenS1{F0: 712, F1: v5}
                _ = v22
            } else {
                switch -2147483648 {
                case 127:
                    fmt.Println(v6)
                case 785:
                    v3 = (v3 << ((636) & 63))
                case 32767:
                    var v23 int = 841
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 3; fi24++ {
                        _ = len(v1)
                        _ = fi24
                    }
                }
            }
            v5 = v5
            _ = v21
        } else {
            var v25 string = "hello"
            for fi26 := 0; fi26 < 10; fi26++ {
                for fi27 := 0; fi27 < 6; fi27++ {
                    v3 += 4294967295
                    var v28 uint16 = 42
                    _ = fi27
                    _ = v28
                }
                if len(v1) > 0 {
                    v1[0] = 863
                }
                v3 = v3
                _ = fi26
            }
            var v29 uint32 = (v5 & 2147483647)
            _ = v25
            _ = v29
        }
    }
    var v30 uint8 = 173
    for fi31 := 0; fi31 < 4; fi31++ {
        fmt.Println(v30, v6)
        if v6 {
            v32 := GenS2{F0: (v30 / (((^(v30))) | 1)), F1: 66, F2: 0}
            if len(v1) > 0 {
                v1[0] = (v5 * v3)
            }
            v33 := GenS3{F0: v30}
            _ = v32
            _ = v33
        } else {
            var v34 int8 = -128
            v1 = append(v1, v3)
            if (42 != 65535) {
                for fi35 := 0; fi35 < 5; fi35++ {
                    v30 = (v30 & v30)
                    _ = fi35
                }
                v6 = false
                if false {
                    for fi36 := 0; fi36 < 3; fi36++ {
                        var v37 int = 42
                        _ = fi36
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 10; fi38++ {
                        v39 := GenS4{F0: true, F1: fi31}
                        _ = fi38
                        _ = v39
                    }
                } else {
                    fmt.Println(v3, v3)
                    var v40 bool = (0.1 != (float32(v4) + float32((0.0 * float64(1e10)))))
                    v2["hello"] = (127 == 65535)
                    _ = v40
                }
            }
            _ = v34
        }
        _ = fi31
    }
    switch "foo bar" {
    case "\n\t":
        v41 := GenS1{F0: 0, F1: (^(v5))}
        _ = v41
    case "":
        v5 -= 0
    case "hello":
        fmt.Println(v4, v30)
    default:
        v42 := GenS2{F0: (v30 - (v30 + (v30 + (^(v30))))), F1: -40, F2: 2147483647}
        _ = v42
    }
    var v43 uint64 = 18446744073709551615
    var v44 int64 = 42
    v5 = 2147483647
    delete(v2, "foo bar")
    for fi45 := 0; fi45 < 10; fi45++ {
        var v46 uint8 = v30
        _ = fi45
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
