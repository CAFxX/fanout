package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int; F1 string; F2 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint; F2 int16 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint64) uint32 {
    var v7 string = ("foo bar" + "\n\t")
    if (255 >= 255) {
        v7 = "\n\t"
        p0 = 1
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{}
            _ = fi8
            _ = v9
        }
    } else {
        var v10 uint32 = 0
        var v11 int64 = -368
        if false {
            v12 := GenS1{F0: v7, F1: 1}
            _ = v12
        }
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v7
    return 65535
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]bool)
    var v3 int16 = -1
    v4 := []string{"foo bar"}
    v5 := make(map[string]int)
    var v6 string = "a"
    _ = len(v5)
    v1 = (v1 - v1)
    switch -211 {
    case -648:
        var v13 int64 = 0
        _ = v13
    case 2147483647:
        v14 := GenS2{F0: 772, F1: (v6 + v6), F2: 2}
        _ = v14
    case -923:
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS2{F0: fi15, F1: "hello", F2: 42}
            v17 := GenS1{F0: "", F1: 127}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        if false {
            if (false && (("hello" + v6) >= "foo bar")) {
                switch v6 {
                case "foo bar":
                    var v19 int64 = 42
                    _ = v19
                case "":
                    fmt.Println(v1, v3)
                case "hello":
                    switch ("a" + (("a" + v6) + v6)) {
                    case "":
                        if (-35 < 208) {
                            var v20 bool = ((false || true) || (1 < 0))
                            v21 := GenS1{F0: (v6 + "foo bar"), F1: 255}
                            for fi22 := 0; fi22 < 1; fi22++ {
                                v1 = v1
                                v23 := GenS3{F0: 127, F1: 13}
                                _ = len(v2)
                                _ = fi22
                                _ = v23
                            }
                            _ = v20
                            _ = v21
                        }
                    case "a":
                        v24 := GenS4{F0: 147, F1: 42, F2: (v3 >> ((351) & 63))}
                        _ = v24
                    case "foo bar":
                        fmt.Println(fi18, v1, v3)
                    default:
                        v4 = append(v4, "a")
                    }
                default:
                    var v25 int16 = v3
                    _ = v25
                }
                v26 := GenS5{F0: v6, F1: 65535}
                _ = v26
            }
        } else {
            if ("hello" <= ((("foo bar" + "") + "") + (((v6 + (("\n\t" + v6) + v6)) + ("" + v6)) + ("\n\t" + "\n\t")))) {
                fmt.Println(v6)
                _ = len(v5)
            }
        }
        for fi27 := 0; fi27 < 20; fi27++ {
            if false {
                var v28 float32 = (((1.5 * 3.14159) / 1.0) - (1.5 - (float32(float64(((0.1 + (1e10 - 0.1)) - 3.14159))) - (1.5 - float32(0.1)))))
                v5[v6] = 4294967295
                switch 0 {
                case 9223372036854775807:
                    v1 = (v1 * 127)
                }
                _ = v28
            } else {
                if (0 <= 186) {
                    v29 := GenS6{F0: (v1 + (v1 >> ((0) & 63))), F1: (v1 ^ (v1 >> ((4294967295) & 63)))}
                    if (42 == 832) {
                        fmt.Println(v1, fi27, v6)
                    } else {
                        var v30 uint64 = 18446744073709551615
                        var v31 int8 = v1
                        if len(v4) > 0 {
                            v4[0] = v6
                        }
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS6{F0: (^(v1)), F1: (v1 << ((2) & 63))}
                    _ = v29
                    _ = v32
                } else {
                    var v33 bool = false
                    if ((v1 & v1) >= v1) {
                        var v34 uint = 65535
                        var v35 float32 = float32(-2.25)
                        _ = v34
                        _ = v35
                    } else {
                        v1 += v1
                        var v36 int8 = 19
                        v6 = "hello"
                        _ = v36
                    }
                    var v37 uint64 = 2
                    var v38 bool = true
                    _ = v33
                    _ = v37
                    _ = v38
                }
            }
            if ((3.14159 * 1.5) >= (0.0 + (1.5 / 1.0))) {
                v5[("hello" + "foo bar")] = (fi18 << ((1) & 63))
            }
            _ = len(v4)
            v3 = 42
            _ = fi27
        }
        _ = fi18
    }
    var v39 uint8 = 2
    fmt.Println(v3, v39)
    v40 := GenS1{F0: (v6 + "\n\t"), F1: v39}
    for fi41 := 0; fi41 < 19; fi41++ {
        if (2147483648 <= uint32(v39)) {
            fmt.Println(v3)
            v39 -= (^(v39))
            if false {
                v39 = 0
            }
        } else {
            switch ((v6 + (v6 + ((v6 + "a") + (v6 + (v6 + v6))))) + v6) {
            case "a":
                v1 += v1
            case "hello":
                v3 = (v3 / ((42) | 1))
            }
            delete(v2, (fi41 << ((0) & 63)))
            v2[fi41] = true
        }
        v1 = -29
        _ = fi41
    }
    var v42 uint32 = 529
    if true {
        v43 := GenS2{F0: -2147483648, F1: v6, F2: v39}
        var v44 uint8 = v39
        _ = v43
        _ = v44
    }
    v3 = (v3 >> ((2) & 63))
    var v45 int = 127
    v46 := GenS7{F0: 317}
    var v47 uint64 = 107
    fmt.Println(v1, v47)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
