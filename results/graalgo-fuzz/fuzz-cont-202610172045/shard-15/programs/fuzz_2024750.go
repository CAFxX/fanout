package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 int16 = -183
    v2 := []int{4294967295, -111, -777}
    var v3 uint64 = 42
    for fi4 := 0; fi4 < 12; fi4++ {
        if (v3 != 4294967295) {
            var v5 uint32 = 65535
            _ = v5
        } else {
            v3 += 0
            var v6 string = "hello"
            v3 = v3
            _ = v6
        }
        v7 := GenS1{F0: 32767}
        var v8 int32 = 0
        _ = fi4
        _ = v7
        _ = v8
    }
    switch "hello" {
    case "hello":
        if true {
            v1 -= (v1 % ((v1) | 1))
            v3 = v3
            var v9 int8 = 100
            _ = v9
        }
    default:
        switch ("" + "") {
        case "hello":
            switch ("hello" + "") {
            case "":
                if (1000 >= v1) {
                    fmt.Println(v3)
                    var v10 string = (("a" + "") + ("hello" + ""))
                    _ = v10
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v11 int = 32768
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v11 = fi12
                        v2 = append(v2, v11)
                        var v13 float32 = 1e10
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
            case "a":
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := GenS2{F0: 2, F1: 42}
                    _ = fi14
                    _ = v15
                }
            }
        case "foo bar":
            if true {
                if (!false) {
                    if false {
                        for fi16 := 0; fi16 < 8; fi16++ {
                            if true {
                                v2 = append(v2, 32767)
                            } else {
                                var v17 uint16 = 0
                                for fi18 := 0; fi18 < 17; fi18++ {
                                    var v19 uint32 = 1
                                    _ = fi18
                                    _ = v19
                                }
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    fmt.Println(v17, fi16, v1)
                                    _ = fi20
                                }
                                _ = v17
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = len(v2)
                            _ = fi16
                        }
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            v2 = append(v2, 0)
                        case "":
                            var v21 int32 = 698
                            _ = v21
                        }
                    } else {
                        v1 += (v1 | v1)
                        switch 128 {
                        case 465:
                            var v22 int64 = -9223372036854775808
                            _ = v22
                        default:
                            var v23 float32 = float32(1.5)
                            _ = v23
                        }
                        v1 = (v1 % ((v1) | 1))
                    }
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v3 -= (v3 - (v3 ^ v3))
                        var v25 float32 = 1.5
                        _ = fi24
                        _ = v25
                    }
                    var v26 int16 = v1
                    _ = v26
                }
            }
        case "\n\t":
            var v27 float32 = (0.0 - 1.5)
            _ = v27
        }
    }
    var v28 float64 = 3.14159
    fmt.Println(v1, v1, v28)
    var v29 uint = 42
    if len(v2) > 0 {
        v2[0] = 127
    }
    var v30 uint32 = 42
    var v31 int64 = 0
    v3 = v3
    var v32 int = -749
    switch v32 {
    case 827:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -343:
        if (!(0 <= 127)) {
            v30 -= v30
            v2 = append(v2, (v32 & v32))
        } else {
            v29 = (v29 ^ v29)
            v33 := GenS1{F0: v1}
            fmt.Println(v30, v1, v31)
            _ = v33
        }
    default:
        fmt.Println(v1, v31, v31)
    }
    v29 = v29
    v29 = (v29 ^ v29)
    v34 := GenS2{F0: 15, F1: (v31 / ((-899) | 1))}
    if len(v2) > 0 {
        v2[0] = (-(v32))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
}
