package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }

func main() {
    var v1 int8 = 57
    var v2 uint64 = 308
    var v3 int32 = -2147483648
    v4 := []uint32{42, 42, 1}
    v5 := []string{"hello", "foo bar"}
    v6 := []string{("" + ("hello" + "")), "hello"}
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int = (fi7 - fi8)
            var v10 int32 = -2147483648
            if (uint16(v2) >= 65535) {
                v3 += (v10 | (v3 * v3))
                fmt.Println(v9)
                fmt.Println(fi7, v3, v1)
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint8 = 16
        var v12 uint8 = v11
        _ = fi7
        _ = v11
        _ = v12
    }
    v2 = (v2 + (v2 << ((926) & 63)))
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS1{F0: 1}
        _ = fi13
        _ = v14
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 = append(v5, "\n\t")
    v15 := GenS2{F0: v1, F1: 1, F2: false}
    v3 += (v3 * int32(v2))
    switch "" {
    case "a":
        var v16 uint8 = 0
        _ = v16
    case "\n\t":
        fmt.Println(v1, v1)
    default:
        var v17 float64 = (1.5 / (3.14159 + (0.1 - (-2.25 + 3.14159))))
        _ = v17
    }
    if false {
        var v18 bool = (123 >= v1)
        v18 = (250 > 1)
        _ = v18
    } else {
        v3 = (^(v3))
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                fmt.Println(v1)
                var v21 uint = 119
                fmt.Println(fi20, v3)
                if (!(30 >= 2147483647)) {
                    var v22 uint8 = 42
                    if false {
                        if false {
                            v6 = append(v6, "hello")
                        }
                    }
                    if false {
                        v22 -= v22
                        v6 = append(v6, "a")
                    }
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            switch ("a" + ("" + "\n\t")) {
            case "":
                v23 := GenS3{F0: 255}
                _ = v23
            }
            if true {
                if true {
                    if (!((v1 >> ((4294967295) & 63)) >= v1)) {
                        if (true || true) {
                            fmt.Println(v3, v1)
                        }
                    } else {
                        if (false || (1000 < 338)) {
                            for fi24 := 0; fi24 < 7; fi24++ {
                                var v25 uint8 = 120
                                v1 = (v1 % ((v1) | 1))
                                v25 += (v25 + v25)
                                _ = fi24
                                _ = v25
                            }
                        }
                        v1 += (^(v1))
                        v26 := GenS3{F0: 255}
                        _ = v26
                    }
                }
                if false {
                    var v27 uint16 = 65535
                    var v28 bool = false
                    _ = v27
                    _ = v28
                } else {
                    var v29 uint16 = 32767
                    for fi30 := 0; fi30 < 14; fi30++ {
                        var v31 string = "a"
                        var v32 float32 = -2.25
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    switch "a" {
                    case "foo bar":
                        v29 = (^(v29))
                    case "hello":
                        switch fi19 {
                        case 699:
                            for fi33 := 0; fi33 < 7; fi33++ {
                                var v34 uint = 0
                                _ = fi33
                                _ = v34
                            }
                        case 32767:
                            v35 := GenS4{F0: 2}
                            _ = v35
                        }
                    default:
                        var v36 uint32 = 1
                        _ = v36
                    }
                    _ = v29
                }
                for fi37 := 0; fi37 < 10; fi37++ {
                    v1 -= 76
                    _ = fi37
                }
            }
            _ = fi19
        }
    }
    var v38 float64 = 0.0
    var v39 int = -234
    v39 = 42
    switch (v39 & (v39 | v39)) {
    case 9223372036854775807:
        fmt.Println(v3, v39)
    case 32768:
        if false {
            var v40 uint8 = 1
            _ = v40
        }
    case 42:
        if (((^(v2)) < v2) || ((v38 + (v38 - (v38 * v38))) < float64(float32(v38)))) {
            if false {
                v5 = append(v5, "hello")
            } else {
                v3 = (v3 - v3)
                switch v39 {
                case 650:
                    if true {
                        var v41 int64 = -400
                        v39 -= v39
                        var v42 int8 = (v1 >> ((99) & 63))
                        _ = v41
                        _ = v42
                    } else {
                        v3 -= 1
                        v38 = (1e10 + -2.25)
                        v43 := GenS1{F0: -736}
                        _ = v43
                    }
                case -1:
                    var v44 uint32 = 1
                    _ = v44
                }
            }
        } else {
            v45 := GenS4{F0: 42}
            _ = v45
        }
    default:
        v38 += 0.1
    }
    var v46 uint32 = 65535
    if (-1 > 42) {
        if ((0 <= 32767) && false) {
            v39 = 1
        }
        v39 -= (^(v39))
        fmt.Println(v46, v46, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
