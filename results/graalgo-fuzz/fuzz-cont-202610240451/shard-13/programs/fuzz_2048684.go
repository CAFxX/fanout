package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint; F1 int16; F2 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 int64 }
func genFn0(p0 int, p1 uint16, p2 uint8) int64 {
    p1 -= (^(p1))
    var v4 float32 = float32(3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -9223372036854775808
}

func main() {
    v1 := []string{"foo bar", "\n\t"}
    v2 := make(map[string]string)
    v3 := []string{"hello", ("foo bar" + ""), ("a" + ("hello" + ""))}
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            if true {
                var v7 bool = (!false)
                switch ("hello" + "a") {
                case "hello":
                    switch "" {
                    case "":
                        if v7 {
                            if v7 {
                                if (v7 || (57 == 28)) {
                                    var v8 int32 = -2147483648
                                    v8 += 0
                                    _ = v8
                                } else {
                                    var v9 int = fi5
                                    _ = v9
                                }
                            }
                            v2["foo bar"] = (("" + "hello") + ("" + ""))
                        } else {
                            fmt.Println(fi6, fi5, fi5)
                        }
                    case "hello":
                        if ((!true) || (0 < 42)) {
                            var v10 float64 = 1.5
                            v7 = v7
                            _ = v10
                        } else {
                            fmt.Println(v7)
                            v11 := GenS1{F0: 2, F1: "hello", F2: "a"}
                            v7 = v7
                            _ = v11
                        }
                    case "\n\t":
                        var v12 int = fi6
                        _ = v12
                    default:
                        v7 = v7
                    }
                case "\n\t":
                    if v7 {
                        switch ("foo bar" + ((("a" + "\n\t") + ("\n\t" + "foo bar")) + ("\n\t" + "\n\t"))) {
                        case "hello":
                            switch fi5 {
                            case 97:
                                v7 = true
                            case -2147483648:
                                switch ("a" + "\n\t") {
                                case "foo bar":
                                    var v13 int8 = 124
                                    _ = v13
                                case "hello":
                                    fmt.Println(fi6, v7, v7)
                                    delete(v2, ("foo bar" + "a"))
                                case "\n\t":
                                    switch fi5 {
                                    case 128:
                                        var v14 int16 = 32767
                                        _ = v14
                                    case 4294967295:
                                        var v15 uint32 = 2147483647
                                        _ = v15
                                    case 42:
                                        var v16 uint64 = 842
                                        _ = v16
                                    }
                                default:
                                    var v17 uint32 = 1
                                    _ = v17
                                }
                            case 32767:
                                v18 := GenS1{F0: 42, F1: ("" + "hello"), F2: ""}
                                _ = v18
                            }
                        case "\n\t":
                            v19 := GenS2{F0: -550, F1: 667}
                            _ = v19
                        case "a":
                            var v20 float32 = 0.0
                            _ = v20
                        }
                        if (-1 < -908) {
                            switch ("" + "foo bar") {
                            case "hello":
                                v21 := GenS1{F0: 1, F1: "hello", F2: "a"}
                                _ = v21
                            case "":
                                var v22 uint16 = 65535
                                _ = v22
                            case "\n\t":
                                if (true && false) {
                                    fmt.Println(fi5, fi5)
                                    var v23 int32 = -1
                                    v24 := GenS3{F0: -244, F1: 1}
                                    _ = v23
                                    _ = v24
                                } else {
                                    _ = len(v2)
                                    if v7 {
                                        v25 := GenS2{F0: 9223372036854775807, F1: 255}
                                        v7 = v7
                                        v7 = v7
                                        _ = v25
                                    } else {
                                        v26 := GenS1{F0: 4294967295, F1: "", F2: "hello"}
                                        v27 := v2["hello"]
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                if ((42 >= 32767) || true) {
                                    var v28 uint32 = 855
                                    v28 = 2
                                    if (42 != 4294967295) {
                                        var v29 int16 = 791
                                        _ = v29
                                    }
                                    _ = v28
                                }
                                if (!(!v7)) {
                                    v7 = v7
                                } else {
                                    v7 = (!false)
                                    if (1 != 42) {
                                        var v30 uint64 = 18446744073709551615
                                        v30 = (v30 % (((v30 >> ((871) & 63))) | 1))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v30
                                    }
                                    v31 := GenS1{F0: 42, F1: "", F2: ""}
                                    _ = v31
                                }
                                fmt.Println(v7)
                            }
                            v32 := GenS1{F0: 18446744073709551615, F1: "hello", F2: "\n\t"}
                            v7 = (!v7)
                            _ = v32
                        } else {
                            var v33 int = fi5
                            _ = v33
                        }
                    }
                default:
                    if v7 {
                        v34 := GenS4{F0: -447, F1: 690}
                        _ = v34
                    } else {
                        var v35 float64 = (3.14159 * 0.0)
                        _ = v35
                    }
                }
                v7 = true
                _ = v7
            }
            _ = fi6
        }
        for fi36 := 0; fi36 < 3; fi36++ {
            var v37 uint32 = 2
            v37 += 1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi36
            _ = v37
        }
        _ = fi5
    }
    var v38 uint16 = 1
    if false {
        _ = len(v2)
        v38 += (v38 ^ (v38 / ((988) | 1)))
        v38 -= v38
    } else {
        v39 := GenS4{F0: 250, F1: 0}
        _ = v39
    }
    var v40 string = (("" + "\n\t") + "\n\t")
    var v41 int = 2147483647
    v40 = v40
    var v42 uint64 = 18446744073709551615
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    if (255 <= 127) {
        v2[(v40 + ("hello" + v40))] = ""
    } else {
        if (2 >= 42) {
            v43 := GenS3{F0: (^(v41)), F1: -29}
            _ = v43
        } else {
            v44 := GenS5{F0: 255, F1: 527, F2: 936}
            switch "hello" {
            case "a":
                if false {
                    for fi45 := 0; fi45 < 2; fi45++ {
                        v2[v40] = v40
                        v41 = fi45
                        v42 -= 42
                        _ = fi45
                    }
                    for fi46 := 0; fi46 < 10; fi46++ {
                        delete(v2, (v40 + ((v40 + v40) + (v40 + (((v40 + v40) + "hello") + "\n\t")))))
                        fmt.Println(fi46, v42)
                        _ = fi46
                    }
                    v40 = "\n\t"
                }
            case "":
                var v47 float64 = float64(float32(((0.1 + ((1e10 - (float64((1.5 * 1e10)) / 1.5)) - (1e10 * 0.0))) / 1e10)))
                _ = v47
            case "hello":
                v48 := GenS6{F0: 255}
                _ = v48
            }
            v49 := GenS7{F0: true, F1: genFn0((-(v41)), 32767, 42)}
            _ = v44
            _ = v49
        }
        var v50 int = v41
        _ = v50
    }
    switch (v41 << ((2) & 63)) {
    case 127:
        for fi51 := 0; fi51 < 18; fi51++ {
            fmt.Println(v40)
            var v52 uint16 = v38
            v53 := GenS3{F0: 32768, F1: 127}
            _ = fi51
            _ = v52
            _ = v53
        }
    case -1:
        for fi54 := 0; fi54 < 10; fi54++ {
            var v55 float64 = 3.14159
            for fi56 := 0; fi56 < 1; fi56++ {
                fmt.Println(v40)
                _ = fi56
            }
            _ = fi54
            _ = v55
        }
    case -479:
        if (true && false) {
            for fi57 := 0; fi57 < 10; fi57++ {
                var v58 int8 = 100
                if true {
                    var v59 uint32 = 862
                    v58 = 0
                    _ = v59
                } else {
                    fmt.Println(v41, fi57)
                    v38 = (v38 & 42)
                }
                _ = fi57
                _ = v58
            }
            var v60 uint16 = (v38 ^ 255)
            var v61 uint16 = (v60 ^ (v38 & (^(v38))))
            _ = v60
            _ = v61
        }
    default:
        var v62 int32 = 42
        _ = v62
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v40
    _ = v41
    _ = v42
}
