package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint16, p2 uint8) uint64 {
    v7 := GenS1{F0: int(p1)}
    var v8 uint = 139
    switch -18 {
    case 831:
        var v9 uint32 = 2147483648
        _ = v9
    case 127:
        var v10 int16 = -32768
        _ = v10
    default:
        var v11 int64 = 9223372036854775807
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []string{"a", "hello", "hello"}
    var v2 int32 = 42
    var v3 uint16 = 65535
    v4 := []string{}
    v5 := make(map[string]string)
    var v6 int32 = v2
    if true {
        var v12 int = 4294967295
        v12 -= v12
        if (!true) {
            v12 = v12
            fmt.Println(v12)
        } else {
            v13 := GenS2{F0: (179 < 608), F1: 42, F2: 183}
            _ = v13
        }
        _ = v12
    } else {
        var v14 uint8 = 127
        if false {
            fmt.Println(v6)
            switch "a" {
            case "a":
                v15 := GenS2{F0: (false || (1.5 != -2.25)), F1: 32767, F2: (^(v14))}
                _ = v15
            case "":
                var v16 int64 = 9223372036854775807
                _ = v16
            case "\n\t":
                var v17 uint32 = 355
                _ = v17
            }
            for fi18 := 0; fi18 < 20; fi18++ {
                switch fi18 {
                case 32767:
                    v6 -= v2
                case 176:
                    if ("a" > ("\n\t" + ("hello" + "foo bar"))) {
                        if (!(!((!true) && true))) {
                            if (1e10 <= (((1e10 - (float32((1e10 / 3.14159)) * 1e10)) + 3.14159) * (float32(0.1) / -2.25))) {
                                for fi19 := 0; fi19 < 8; fi19++ {
                                    var v20 int32 = v6
                                    if (false || true) {
                                        if false {
                                            v6 = v20
                                            if len(v4) > 0 {
                                                v4[0] = "a"
                                            }
                                            var v21 int16 = 1
                                            if false {
                                                var v22 int32 = (-(v6))
                                                fmt.Println(fi18, fi18)
                                                _ = v22
                                            }
                                            _ = v21
                                        }
                                    } else {
                                        var v23 uint64 = uint64(v14)
                                        var v24 bool = (42 < 2)
                                        var v25 float32 = 1e10
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    var v26 uint = 1
                                    _ = fi19
                                    _ = v20
                                    _ = v26
                                }
                                switch "" {
                                case "hello":
                                    if true {
                                        for fi27 := 0; fi27 < 8; fi27++ {
                                            v6 = 2147483647
                                            var v28 int64 = -37
                                            _ = fi27
                                            _ = v28
                                        }
                                    } else {
                                        if (!((-66 <= -128) && false)) {
                                            v14 -= (v14 + 42)
                                            fmt.Println(v2)
                                        }
                                    }
                                case "foo bar":
                                    v6 = v2
                                }
                                var v29 int8 = 116
                                _ = v29
                            } else {
                                fmt.Println(v3)
                                switch ("\n\t" + "foo bar") {
                                case "\n\t":
                                    for fi30 := 0; fi30 < 18; fi30++ {
                                        if false {
                                            v31 := GenS3{F0: 4294967295}
                                            _ = v31
                                        }
                                        v32 := GenS3{F0: 192}
                                        _ = fi30
                                        _ = v32
                                    }
                                default:
                                    var v33 string = "\n\t"
                                    _ = v33
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        v34 := GenS2{F0: true, F1: -685, F2: 127}
                        var v35 uint16 = 137
                        _ = v34
                        _ = v35
                    }
                case 1:
                    v6 = v6
                default:
                    for fi36 := 0; fi36 < 20; fi36++ {
                        var v37 int16 = 1000
                        v38 := GenS1{F0: fi18}
                        var v39 bool = (1.5 != (1e10 - 1e10))
                        _ = fi36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
                _ = fi18
            }
        }
        if (255 != 40) {
            var v40 uint8 = v14
            v3 += (v3 + (v3 ^ v3))
            v5["hello"] = "\n\t"
            _ = v40
        }
        _ = v14
    }
    var v41 float32 = 0.1
    _ = len(v4)
    fmt.Println(v41)
    v4 = append(v4, "")
    var v42 bool = (((("" + "\n\t") >= "hello") && true) && false)
    switch ("a" + "foo bar") {
    case "a":
        v3 = v3
    case "\n\t":
        var v43 bool = v42
        _ = v43
    default:
        var v44 float64 = (3.14159 + (3.14159 + -2.25))
        _ = v44
    }
    var v45 int32 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v5
    _ = v6
}
