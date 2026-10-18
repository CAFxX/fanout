package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int16; F2 int64 }
type GenS4 struct { F0 bool }
func genFn0(p0 string) int64 {
    var v7 float64 = -2.25
    v7 += 3.14159
    switch 127 {
    case 42:
        p0 = p0
    case -1:
        v7 = float64(1.5)
    case 0:
        var v8 string = p0
        _ = v8
    default:
        var v9 uint32 = 329
        _ = v9
    }
    _ = p0
    _ = v7
    return -825
}

func main() {
    var v1 int32 = 748
    var v2 uint32 = 65535
    var v3 int32 = (^(v1))
    var v4 int = -544
    v5 := []uint32{}
    var v6 uint8 = uint8(v2)
    if ((v3 >> ((65535) & 63)) < (v3 ^ v3)) {
        if len(v5) > 0 {
            v5[0] = (v2 - v2)
        }
        v10 := GenS1{F0: v3, F1: (32767 > 0), F2: 1}
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(v2, v2)
            v3 += (v1 & -1)
            fmt.Println(v2, v6)
            _ = fi11
        }
        _ = v10
    } else {
        switch (v4 >> ((719) & 63)) {
        case 0:
            var v12 string = ""
            _ = v12
        case 9223372036854775807:
            v13 := GenS1{F0: v3, F1: true, F2: -1}
            _ = v13
        default:
            var v14 uint16 = 40
            _ = v14
        }
    }
    var v15 uint16 = 65535
    v2 += (^(v2))
    v4 = int(v6)
    v1 -= v3
    _ = len(v5)
    switch 32767 {
    case 98:
        fmt.Println(v6, v2)
    case 127:
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 string = "foo bar"
            _ = fi16
            _ = v17
        }
    case 733:
        for fi18 := 0; fi18 < 1; fi18++ {
            if false {
                for fi19 := 0; fi19 < 1; fi19++ {
                    fmt.Println(fi19, v2)
                    fmt.Println(fi18, fi19, v3)
                    _ = fi19
                }
                switch ("\n\t" + ("\n\t" + "")) {
                case "":
                    v4 += (v4 - 32768)
                case "a":
                    if len(v5) > 0 {
                        v5[0] = (^(v2))
                    }
                case "foo bar":
                    v3 -= (v1 + -1)
                default:
                    if false {
                        v3 = (^(v3))
                        switch "a" {
                        case "a":
                            if (!(2147483647 < v2)) {
                                if (0.0 <= 3.14159) {
                                    v15 = v15
                                    v6 -= (^(v6))
                                    v20 := GenS2{F0: -1}
                                    _ = v20
                                }
                            } else {
                                var v21 uint = 4294967295
                                _ = v21
                            }
                        }
                        if false {
                            v22 := GenS3{F0: 107, F1: 1, F2: -783}
                            var v23 uint8 = (v6 ^ v6)
                            _ = v22
                            _ = v23
                        }
                    }
                }
                var v24 int = (fi18 / ((v4) | 1))
                _ = v24
            }
            _ = fi18
        }
    }
    fmt.Println(v3)
    v25 := GenS3{F0: -128, F1: -874, F2: genFn0("\n\t")}
    var v26 int8 = 42
    if ((("foo bar" + ((((("\n\t" + "foo bar") + ("" + "")) + "") + "hello") + (("a" + "foo bar") + "foo bar"))) + ("hello" + "hello")) <= "hello") {
        fmt.Println(v26, v3, v15)
        switch "a" {
        case "":
            fmt.Println(v6, v3, v15)
        default:
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := GenS1{F0: v1, F1: (!((true || false) && (v2 >= v2))), F2: -9223372036854775808}
                _ = fi27
                _ = v28
            }
        }
        if (42 > -674) {
            v5 = append(v5, (v2 * uint32(v6)))
            for fi29 := 0; fi29 < 11; fi29++ {
                var v30 int = (v4 ^ v4)
                _ = fi29
                _ = v30
            }
        } else {
            v31 := GenS3{F0: v26, F1: 1000, F2: 562}
            var v32 float32 = 1e10
            _ = v31
            _ = v32
        }
    }
    switch (^(v4)) {
    case 9223372036854775807:
        if len(v5) > 0 {
            v5[0] = v2
        }
    case 4294967295:
        var v33 string = (("\n\t" + ("a" + "a")) + "\n\t")
        _ = v33
    case 2147483647:
        var v34 float32 = (-2.25 + 0.0)
        _ = v34
    default:
        if len(v5) > 0 {
            v5[0] = v2
        }
    }
    var v35 int32 = (v3 / ((v1) | 1))
    var v36 uint = 255
    v37 := GenS4{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
