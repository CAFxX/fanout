package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint) uint8 {
    v7 := GenS1{F0: (false && (986 >= 151))}
    var v8 uint64 = (p0 + (p0 - p0))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int32, p1 int32, p2 uint64) uint64 {
    var v9 int64 = 9223372036854775807
    var v10 float64 = (-2.25 / 1.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (^(p2))
}

func main() {
    v1 := []uint32{4294967295, 2147483647, 2147483647}
    v2 := []string{"hello", "", (("\n\t" + "\n\t") + ("" + ("" + "\n\t")))}
    var v3 float64 = 1.5
    v4 := make(map[string]string)
    v5 := []uint32{42, 0}
    var v6 bool = ((4294967295 > 366) && false)
    var v11 uint = 93
    var v12 uint32 = 1
    for fi13 := 0; fi13 < 12; fi13++ {
        fmt.Println(v12, v12, v3)
        var v14 int8 = 42
        v12 = v12
        _ = fi13
        _ = v14
    }
    v3 += 0.1
    v15 := GenS1{F0: v6}
    if v6 {
        v11 -= v11
        fmt.Println(v3, v6)
        v16 := GenS1{F0: (108 < 74)}
        _ = v16
    } else {
        switch "hello" {
        case "hello":
            var v17 float32 = (0.0 + 1.5)
            _ = v17
        case "foo bar":
            if v6 {
                if (78 < 373) {
                    v6 = v6
                }
                for fi18 := 0; fi18 < 11; fi18++ {
                    v6 = true
                    if v6 {
                        fmt.Println(v3)
                        v19 := GenS1{F0: false}
                        if true {
                            var v20 int64 = 511
                            fmt.Println(v11, v20)
                            if v6 {
                                for fi21 := 0; fi21 < 14; fi21++ {
                                    _ = len(v4)
                                    fmt.Println(v11)
                                    _ = fi21
                                }
                                fmt.Println(v3, v3, v20)
                            }
                            _ = v20
                        }
                        _ = v19
                    }
                    _ = fi18
                }
            }
        default:
            var v22 uint16 = 178
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        if len(v2) > 0 {
            v2[0] = ("foo bar" + "a")
        }
        v11 = (v11 * (v11 & (v11 * v11)))
        v6 = true
        _ = fi23
    }
    switch 32767 {
    case -2147483648:
        for fi24 := 0; fi24 < 5; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                fmt.Println(v6)
                _ = fi25
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, "a")
            _ = fi24
        }
    case -278:
        switch -1 {
        case -717:
            var v26 float32 = ((-2.25 * 1.5) * ((0.0 / 1.0) * (1e10 + 1.5)))
            _ = v26
        }
    }
    if v6 {
        if true {
            v27 := GenS2{F0: genFn1(1, 2147483647, genFn1(-817, -938, genFn1(-2147483648, 0, 943)))}
            for fi28 := 0; fi28 < 9; fi28++ {
                v12 = 443
                fmt.Println(v6, v3)
                _ = fi28
            }
            _ = v27
        } else {
            if (false || (-32768 > 1)) {
                v3 = (1.5 * v3)
            } else {
                var v29 uint64 = 117
                _ = v29
            }
        }
        fmt.Println(v12, v11)
        if (32767 < -1) {
            fmt.Println(v11, v3)
            v30 := GenS2{F0: 4294967295}
            var v31 uint64 = 2
            _ = v30
            _ = v31
        } else {
            if (-9223372036854775808 <= -662) {
                switch "\n\t" {
                case "hello":
                    switch "a" {
                    case "foo bar":
                        var v32 bool = v6
                        _ = v32
                    }
                case "foo bar":
                    var v33 bool = v6
                    _ = v33
                case "\n\t":
                    for fi34 := 0; fi34 < 6; fi34++ {
                        for fi35 := 0; fi35 < 8; fi35++ {
                            fmt.Println(v11, fi35)
                            _ = fi35
                        }
                        _ = len(v4)
                        _ = fi34
                    }
                default:
                    switch 127 {
                    case 591:
                        for fi36 := 0; fi36 < 12; fi36++ {
                            fmt.Println(v3, v12)
                            fmt.Println(v6)
                            _ = fi36
                        }
                    case -405:
                        var v37 uint8 = 74
                        _ = v37
                    case 127:
                        v6 = (v6 || (!true))
                    default:
                        var v38 uint32 = (v12 >> ((v11) & 63))
                        _ = v38
                    }
                }
                if v6 {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    var v39 uint = v11
                    _ = v39
                }
                for fi40 := 0; fi40 < 18; fi40++ {
                    var v41 uint8 = 42
                    _ = fi40
                    _ = v41
                }
            }
            v5 = append(v5, 42)
            v42 := v4["hello"]
            _ = v42
        }
    } else {
        v11 = v11
        for fi43 := 0; fi43 < 14; fi43++ {
            fmt.Println(v11)
            var v44 bool = (-2.25 <= (0.1 / 1e10))
            v44 = (!true)
            _ = fi43
            _ = v44
        }
        v11 = (^(v11))
    }
    v11 = (v11 % ((v11) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
