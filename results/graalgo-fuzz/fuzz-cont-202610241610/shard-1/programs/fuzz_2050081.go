package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint = 1
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    v4 := []string{"hello", "a", "foo bar", "foo bar"}
    var v5 int8 = 100
    var v6 int8 = 100
    fmt.Println(v5, v5, v1)
    var v7 int = 127
    if false {
        var v8 int = v7
        _ = v8
    } else {
        var v9 int32 = 1
        _ = v9
    }
    delete(v2, "\n\t")
    v10 := GenS1{F0: false, F1: v1}
    v6 -= v5
    v6 += (v6 ^ v5)
    var v11 uint32 = 0
    for fi12 := 0; fi12 < 1; fi12++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            v13 := v2["foo bar"]
            _ = v13
        } else {
            delete(v3, v7)
        }
        _ = fi12
    }
    if ("foo bar" == ("a" + "hello")) {
        fmt.Println(v7, v7)
        v14 := v2[("\n\t" + "hello")]
        switch int(v1) {
        case 32768:
            for fi15 := 0; fi15 < 12; fi15++ {
                for fi16 := 0; fi16 < 1; fi16++ {
                    v1 = v1
                    _ = fi16
                }
                if (true || (!false)) {
                    if (true || false) {
                        v17 := GenS2{F0: false, F1: 42}
                        _ = v17
                    } else {
                        fmt.Println(v7, v1, v7)
                        var v18 string = "hello"
                        for fi19 := 0; fi19 < 4; fi19++ {
                            v18 = "\n\t"
                            v7 -= v7
                            _ = fi19
                        }
                        _ = v18
                    }
                }
                if true {
                    var v20 uint = v1
                    for fi21 := 0; fi21 < 7; fi21++ {
                        switch v7 {
                        case 9223372036854775807:
                            var v22 uint32 = (v11 / ((v14) | 1))
                            _ = v22
                        default:
                            switch v7 {
                            case 2147483647:
                                v11 -= (v11 >> ((v20) & 63))
                                fmt.Println(fi21, v1, v11)
                                v4 = append(v4, "hello")
                            case 32768:
                                fmt.Println(fi15, v5, v14)
                            case -2147483648:
                                v23 := GenS2{F0: (42 <= 1), F1: 42}
                                _ = v23
                            default:
                                v7 -= 2147483647
                            }
                        }
                        v5 += v5
                        _ = fi21
                    }
                    _ = v20
                }
                _ = fi15
            }
        case 0:
            if (false || true) {
                for fi24 := 0; fi24 < 5; fi24++ {
                    for fi25 := 0; fi25 < 4; fi25++ {
                        var v26 bool = true
                        var v27 int16 = -246
                        _ = len(v2)
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    var v28 float32 = (1e10 / 1e10)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi24
                    _ = v28
                }
                for fi29 := 0; fi29 < 11; fi29++ {
                    v1 = (v1 * v1)
                    fmt.Println(v6)
                    for fi30 := 0; fi30 < 11; fi30++ {
                        var v31 bool = (1.5 < (1.5 * 3.14159))
                        v32 := v2[((("" + "hello") + ("hello" + "\n\t")) + ("a" + "\n\t"))]
                        if false {
                            v4 = append(v4, "a")
                            v1 -= v1
                            if false {
                                switch ("\n\t" + "a") {
                                case "foo bar":
                                    v33 := GenS1{F0: v31, F1: (v1 - 604)}
                                    _ = v33
                                case "a":
                                    v14 = v14
                                default:
                                    if (v31 || true) {
                                        var v34 int64 = -9223372036854775808
                                        var v35 float64 = 0.0
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 uint32 = (v11 ^ (v11 * (^(v11))))
                                        var v37 int = -1
                                        v5 -= 42
                                        v38 := GenS2{F0: (((v31 && v31) && v31) || (!(v32 > v14))), F1: 381}
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                }
                            }
                            var v39 bool = (255 >= 285)
                            v40 := GenS1{F0: (0.1 == 1e10), F1: v1}
                            _ = v39
                            _ = v40
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    _ = fi29
                }
            }
        case 127:
            for fi41 := 0; fi41 < 15; fi41++ {
                v42 := GenS2{F0: true, F1: 42}
                _ = fi41
                _ = v42
            }
        }
        _ = v14
    }
    v43 := GenS2{F0: (true || false), F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
