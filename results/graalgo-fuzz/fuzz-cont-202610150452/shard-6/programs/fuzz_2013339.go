package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int16 = 43
    var v2 string = ("\n\t" + "\n\t")
    v3 := []uint32{42, 2147483647, 1, 977}
    var v4 uint64 = 0
    var v5 uint16 = 1
    var v6 uint16 = (v5 >> ((2) & 63))
    if (true && true) {
        v3 = append(v3, 42)
        v5 = 994
        fmt.Println(v2, v5)
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 2147483647, F1: 153}
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                if false {
                    var v10 bool = false
                    _ = v10
                } else {
                    switch fi8 {
                    case 32767:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
                switch "\n\t" {
                case "a":
                    var v11 uint16 = v6
                    _ = v11
                case "":
                    fmt.Println(v6, v1)
                    switch -9223372036854775808 {
                    case 0:
                        v6 = 796
                    }
                case "hello":
                    v1 += v1
                }
            }
            var v12 int = int(v1)
            if len(v3) > 0 {
                v3[0] = 2
            }
            _ = fi8
            _ = v12
            _ = v9
        }
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        _ = fi7
    }
    v6 += 2
    if false {
        for fi13 := 0; fi13 < 18; fi13++ {
            fmt.Println(v6)
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS1{F0: 987, F1: 414}
        v2 = "hello"
        if true {
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 uint32 = 2147483647
                fmt.Println(v4, v5)
                _ = fi16
                _ = v17
            }
            if false {
                v18 := GenS1{F0: 2, F1: 18446744073709551615}
                _ = v18
            } else {
                var v19 uint = 4294967295
                v5 -= (v6 & v5)
                _ = v19
            }
        } else {
            if ((-2147483648 >= -1) && ((1e10 / 1.5) == (1e10 - 3.14159))) {
                if true {
                    if false {
                        for fi20 := 0; fi20 < 9; fi20++ {
                            var v21 float64 = -2.25
                            v1 += (v1 << ((114) & 63))
                            _ = fi20
                            _ = v21
                        }
                        v22 := GenS1{F0: uint32(v4), F1: 338}
                        v23 := GenS2{F0: 42, F1: 2, F2: 42}
                        _ = v22
                        _ = v23
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v24 int32 = 42
                _ = v24
            }
            if true {
                var v25 bool = false
                var v26 bool = v25
                var v27 string = ""
                _ = v25
                _ = v26
                _ = v27
            } else {
                v5 += v5
                fmt.Println(v2, fi14)
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                v2 = (v2 + (v2 + ("hello" + "")))
                _ = fi28
            }
        }
        _ = fi14
        _ = v15
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 float64 = 0.1
        _ = fi29
        _ = v30
    }
    if (true || true) {
        var v31 string = ((v2 + (("\n\t" + v2) + v2)) + ("foo bar" + v2))
        _ = v31
    }
    v5 = (v6 | v5)
    if false {
        for fi32 := 0; fi32 < 3; fi32++ {
            v33 := GenS3{F0: (v4 >> ((24) & 63)), F1: -33}
            switch ("hello" + "foo bar") {
            case "\n\t":
                v2 = v2
            case "hello":
                var v34 string = "hello"
                _ = v34
            default:
                var v35 uint = 0
                _ = v35
            }
            _ = fi32
            _ = v33
        }
        switch ("" + (v2 + (v2 + ""))) {
        case "a":
            if false {
                switch -2147483648 {
                case 4294967295:
                    if (!false) {
                        v2 = v2
                        if true {
                            var v36 uint = 146
                            _ = v36
                        } else {
                            v4 = (v4 - (v4 % ((v4) | 1)))
                        }
                    } else {
                        v5 = (v6 << ((4294967295) & 63))
                    }
                case 100:
                    v4 += 502
                case -709:
                    v3 = append(v3, uint32(v4))
                }
                if (!(!true)) {
                    var v37 uint64 = v4
                    var v38 uint16 = v6
                    _ = v37
                    _ = v38
                }
            }
        case "hello":
            if false {
                v39 := GenS3{F0: (v4 / (((v4 & (v4 << ((255) & 63)))) | 1)), F1: -1}
                v40 := GenS3{F0: (v4 / (((v4 - uint64(v6))) | 1)), F1: -1}
                v6 += v5
                _ = v39
                _ = v40
            }
        case "foo bar":
            var v41 int8 = -67
            _ = v41
        }
    }
    var v42 uint = 145
    fmt.Println(v5, v42, v42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
