package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0() uint64 {
    var v7 uint = 255
    fmt.Println(v7)
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := []string{("foo bar" + "a")}
    v2 := []string{"a", (("\n\t" + "foo bar") + "\n\t"), ""}
    var v3 float64 = 1.5
    var v4 uint32 = 2147483648
    var v5 string = (("foo bar" + "") + "")
    var v6 int16 = 32767
    fmt.Println(v3, v4, v5)
    if true {
        if (true || true) {
            var v8 int16 = v6
            _ = v8
        }
    }
    var v9 float64 = v3
    v10 := GenS1{F0: 18446744073709551615, F1: 499, F2: 42}
    if ((!((9223372036854775807 != 857) || false)) && ("foo bar" == (v5 + v5))) {
        v1 = append(v1, v5)
        if ("foo bar" != v5) {
            switch (v5 + v5) {
            case "foo bar":
                var v11 string = (v5 + v5)
                _ = v11
            case "\n\t":
                if (((v4 << ((774) & 63)) >= v4) || true) {
                    fmt.Println(v5)
                } else {
                    if (!(((!true) || false) && true)) {
                        if true {
                            v12 := GenS1{F0: 83, F1: 42, F2: 713}
                            _ = v12
                        } else {
                            _ = len(v1)
                        }
                        var v13 float64 = -2.25
                        _ = v13
                    } else {
                        var v14 int64 = 1
                        v15 := GenS1{F0: 2, F1: 42, F2: genFn0()}
                        var v16 bool = true
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(v3, v5, v5)
                    v17 := GenS2{F0: 1}
                    _ = v17
                }
            case "":
                v3 = v3
            default:
                for fi18 := 0; fi18 < 20; fi18++ {
                    if (0 == 891) {
                        var v19 uint8 = 1
                        if (!(false && true)) {
                            switch fi18 {
                            case -1:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 4294967295:
                                for fi20 := 0; fi20 < 1; fi20++ {
                                    v9 += ((v9 * (((v3 + v9) - v3) + (v9 - 1.5))) * (v9 + v3))
                                    var v21 uint32 = 2147483648
                                    v22 := GenS1{F0: 144, F1: 1, F2: 1}
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                        }
                        var v23 int64 = 0
                        _ = v19
                        _ = v23
                    } else {
                        var v24 int16 = v6
                        var v25 int64 = -9223372036854775808
                        v26 := GenS1{F0: 779, F1: 4294967295, F2: 42}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    _ = fi18
                }
            }
            var v27 string = v5
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v27
        } else {
            fmt.Println(v6, v4, v9)
            for fi28 := 0; fi28 < 14; fi28++ {
                v6 = v6
                var v29 string = "\n\t"
                fmt.Println(fi28, v6)
                _ = fi28
                _ = v29
            }
        }
        switch "" {
        case "":
            if false {
                v3 = (float64(1.5) - (1e10 / -2.25))
                _ = len(v2)
                v6 -= v6
            }
        case "a":
            var v30 float32 = (3.14159 * ((0.1 / (((-2.25 - ((0.1 / (-2.25 - (((((float32(v3) + 0.1) + -2.25) / ((3.14159 - ((1.5 - float32(0.0)) + 1e10)) - 1.5)) * 1.5) / -2.25))) + 0.1)) / 0.1) - (3.14159 + 1.5))) / 0.1))
            _ = v30
        case "hello":
            var v31 uint32 = 419
            _ = v31
        }
    }
    var v32 uint = 53
    if false {
        fmt.Println(v9, v3)
    } else {
        if ((v6 - 32767) <= (v6 | v6)) {
            for fi33 := 0; fi33 < 15; fi33++ {
                fmt.Println(v9, v32)
                _ = fi33
            }
            if (!false) {
                for fi34 := 0; fi34 < 18; fi34++ {
                    v35 := GenS3{F0: 221, F1: 65535}
                    var v36 float32 = float32(0.1)
                    if (32767 == 255) {
                        for fi37 := 0; fi37 < 6; fi37++ {
                            _ = len(v1)
                            v5 = "a"
                            _ = fi37
                        }
                    }
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                switch (v5 + "hello") {
                case "foo bar":
                    for fi38 := 0; fi38 < 8; fi38++ {
                        v4 -= 42
                        var v39 int = -9223372036854775808
                        _ = fi38
                        _ = v39
                    }
                case "\n\t":
                    for fi40 := 0; fi40 < 8; fi40++ {
                        if true {
                            switch (v5 + v5) {
                            case "hello":
                                var v41 int = -9223372036854775808
                                _ = v41
                            case "a":
                                v42 := GenS1{F0: v32, F1: 42, F2: 4294967295}
                                _ = v42
                            case "":
                                var v43 int = fi40
                                _ = v43
                            }
                            v6 -= (v6 >> (((^(v32))) & 63))
                        } else {
                            if false {
                                fmt.Println(v32, v5, v32)
                                for fi44 := 0; fi44 < 16; fi44++ {
                                    var v45 int64 = 42
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi44
                                    _ = v45
                                }
                                var v46 int64 = 42
                                _ = v46
                            }
                        }
                        fmt.Println(v32)
                        _ = fi40
                    }
                default:
                    if (-1 != 1) {
                        v1 = append(v1, "\n\t")
                        fmt.Println(v9, v5, v32)
                    } else {
                        for fi47 := 0; fi47 < 10; fi47++ {
                            var v48 uint16 = 875
                            if true {
                                v4 = (^(v4))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            _ = fi47
                            _ = v48
                        }
                    }
                }
            } else {
                if true {
                    switch 706 {
                    case 2147483647:
                        fmt.Println(v3, v6, v3)
                    case 32767:
                        v6 += v6
                    case 0:
                        fmt.Println(v9)
                    }
                } else {
                    v1 = append(v1, ("foo bar" + ((((v5 + (("a" + "\n\t") + "a")) + (v5 + (v5 + (v5 + (v5 + v5))))) + "hello") + "a")))
                    var v49 uint16 = 2
                    _ = v49
                }
                v4 = (v4 + 162)
                v4 -= 42
            }
            var v50 string = v5
            _ = v50
        } else {
            v2 = append(v2, ((v5 + "foo bar") + ("\n\t" + v5)))
        }
        for fi51 := 0; fi51 < 7; fi51++ {
            var v52 uint64 = 1
            _ = fi51
            _ = v52
        }
    }
    fmt.Println(v5, v4)
    v53 := GenS2{F0: 42}
    v2 = append(v2, "a")
    v32 = v32
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v9
}
