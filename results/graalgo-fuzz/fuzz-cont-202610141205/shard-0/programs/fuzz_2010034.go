package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 138
    var v3 float32 = (0.1 - 1.5)
    v4 := make(map[int]string)
    v5 := []int{-284}
    var v6 uint = (v2 & v2)
    var v7 int = 32768
    fmt.Println(v6, v2, v7)
    v8 := v4[v7]
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 32767}
        _ = fi9
        _ = v10
    }
    v3 = v3
    if true {
        switch "" {
        case "foo bar":
            for fi11 := 0; fi11 < 6; fi11++ {
                v5 = append(v5, v7)
                var v12 uint64 = 18446744073709551615
                _ = fi11
                _ = v12
            }
        case "\n\t":
            for fi13 := 0; fi13 < 12; fi13++ {
                if false {
                    fmt.Println(fi13, v2, v8)
                    for fi14 := 0; fi14 < 2; fi14++ {
                        switch v8 {
                        case "a":
                            v15 := GenS1{F0: 0}
                            _ = v15
                        case "":
                            v16 := GenS1{F0: -196}
                            _ = v16
                        }
                        _ = fi14
                    }
                } else {
                    _ = len(v4)
                    v6 += v2
                    var v17 int8 = 127
                    _ = v17
                }
                _ = fi13
            }
        case "hello":
            if true {
                switch v8 {
                case "\n\t":
                    for fi18 := 0; fi18 < 13; fi18++ {
                        _ = len(v5)
                        _ = fi18
                    }
                case "":
                    var v19 uint64 = 0
                    _ = v19
                default:
                    v20 := v4[v7]
                    _ = v20
                }
            }
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            fmt.Println(v3, v6, v7)
            v22 := GenS1{F0: int16(v2)}
            var v23 uint32 = uint32(v2)
            _ = fi21
            _ = v22
            _ = v23
        }
        fmt.Println(v7, v6, v7)
    } else {
        v6 += v6
        v2 -= 276
        if false {
            var v24 float64 = 0.0
            if (!false) {
                switch -715 {
                case 1:
                    var v25 bool = (4294967295 > 2)
                    _ = v25
                case 32768:
                    for fi26 := 0; fi26 < 11; fi26++ {
                        v6 += 216
                        _ = len(v1)
                        _ = fi26
                    }
                case 128:
                    v27 := GenS1{F0: 32767}
                    _ = v27
                default:
                    fmt.Println(v2, v24, v7)
                }
            }
            fmt.Println(v8, v8, v8)
            _ = v24
        }
    }
    v2 -= v6
    var v28 uint32 = 65535
    if (-499 <= 32767) {
        var v29 int8 = 1
        if false {
            v4[(v7 - (v7 << ((255) & 63)))] = "\n\t"
        } else {
            var v30 uint64 = 18446744073709551615
            switch (v8 + "\n\t") {
            case "\n\t":
                _ = len(v5)
            case "":
                v30 -= (^(v30))
            case "foo bar":
                for fi31 := 0; fi31 < 18; fi31++ {
                    switch v8 {
                    case "":
                        if ((v30 <= 4294967295) && true) {
                            switch (v7 | (v7 ^ -2147483648)) {
                            case -9223372036854775808:
                                if true {
                                    v4[42] = v8
                                } else {
                                    if true {
                                        v32 := GenS1{F0: -951}
                                        _ = v32
                                    }
                                    v3 = (1e10 - v3)
                                    v30 = (v30 - (v30 / (((v30 * 4294967295)) | 1)))
                                }
                            case 0:
                                for fi33 := 0; fi33 < 1; fi33++ {
                                    var v34 bool = false
                                    _ = fi33
                                    _ = v34
                                }
                            case 254:
                                fmt.Println(v28, v29)
                            }
                        } else {
                            var v35 int8 = v29
                            _ = v35
                        }
                    }
                    _ = fi31
                }
            default:
                v2 -= (v6 ^ (^(v2)))
            }
            _ = v30
        }
        _ = v29
    }
    v6 = (v6 | (v6 % ((v2) | 1)))
    v36 := GenS1{F0: -1}
    if (false && (!(!false))) {
        v8 = ((v8 + v8) + (v8 + (v8 + "foo bar")))
    } else {
        if true {
            var v37 bool = (!true)
            _ = v37
        }
        v3 += v3
        fmt.Println(v3, v6)
    }
    if false {
        var v38 int8 = 6
        _ = v38
    } else {
        v8 = "\n\t"
        v8 = ((v8 + "\n\t") + (("a" + (v8 + "hello")) + v8))
        fmt.Println(v2, v28, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
