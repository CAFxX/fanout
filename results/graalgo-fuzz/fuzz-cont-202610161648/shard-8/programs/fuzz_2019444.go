package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int32; F1 string; F2 int }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 int32; F1 int16; F2 uint8 }
func genFn0(p0 int16) uint {
    v5 := GenS1{F0: 1, F1: -1}
    v6 := []int{2147483647, 330, 498, -446}
    _ = p0
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 bool, p1 int16) uint64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int16 = (p1 ^ -103)
        p0 = (int8(v8) < 42)
        p1 += (-(v8))
        _ = fi7
        _ = v8
    }
    p1 = p1
    v9 := GenS1{F0: 100, F1: -1}
    var v10 int32 = 0
    v11 := []int{2147483647, 128, 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 526
}

func main() {
    var v1 float64 = 1.5
    v2 := []int{}
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]bool)
    v12 := GenS1{F0: 127, F1: 35}
    var v13 int = 4294967295
    fmt.Println(v1, v3)
    var v14 int = (^(v13))
    delete(v4, "")
    v15 := GenS2{F0: ("\n\t" + "")}
    v1 += v1
    v14 = (v13 & v14)
    fmt.Println(v3, v3)
    v14 = v13
    if (2 <= 253) {
        switch "hello" {
        case "\n\t":
            var v16 float32 = 3.14159
            _ = v16
        case "hello":
            switch v13 {
            case -9223372036854775808:
                for fi17 := 0; fi17 < 16; fi17++ {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        if true {
                            if (65 <= int8(v3)) {
                                fmt.Println(fi18)
                                var v19 int16 = -1
                                _ = v19
                            } else {
                                if ((v3 & v3) == -645) {
                                    v14 -= v13
                                    v1 = v1
                                    if false {
                                        if ((fi18 * (fi17 << ((391) & 63))) != (-(v13))) {
                                            var v20 int16 = -1
                                            _ = v20
                                        }
                                    } else {
                                        if true {
                                            v21 := GenS2{F0: "hello"}
                                            _ = v21
                                        } else {
                                            var v22 uint64 = 0
                                            v23 := GenS2{F0: "\n\t"}
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                }
                                v13 = fi18
                                v24 := GenS3{F0: 824, F1: 42, F2: 1}
                                _ = v24
                            }
                            var v25 uint8 = 42
                            if (!(true && (v25 != v25))) {
                                if (!true) {
                                    _ = len(v4)
                                    switch ("hello" + ("a" + "foo bar")) {
                                    case "":
                                        v26 := v4["a"]
                                        _ = v26
                                    }
                                }
                                fmt.Println(v14)
                            }
                            _ = v25
                        } else {
                            fmt.Println(v14, fi18, v3)
                        }
                        v27 := GenS2{F0: ""}
                        var v28 uint = 553
                        _ = fi18
                        _ = v27
                        _ = v28
                    }
                    if (0 != 18446744073709551615) {
                        for fi29 := 0; fi29 < 1; fi29++ {
                            v30 := GenS4{F0: 0, F1: 170, F2: 2147483648}
                            _ = len(v2)
                            _ = fi29
                            _ = v30
                        }
                        var v31 float64 = (v1 * v1)
                        v32 := GenS3{F0: genFn1((!(!(961 != 0))), 1000), F1: 385, F2: -612}
                        _ = v31
                        _ = v32
                    } else {
                        switch "hello" {
                        case "foo bar":
                            v33 := v4[""]
                            _ = v33
                        }
                        v3 += (v3 | v3)
                    }
                    v2 = append(v2, (fi17 + 303))
                    _ = fi17
                }
            case 1:
                if false {
                    fmt.Println(v1)
                } else {
                    if ((true || false) && ("foo bar" <= "")) {
                        for fi34 := 0; fi34 < 8; fi34++ {
                            _ = len(v2)
                            v14 -= (v14 * (fi34 >> ((65535) & 63)))
                            if (false && false) {
                                var v35 uint = 2
                                var v36 uint32 = 2
                                switch "" {
                                case "foo bar":
                                    v37 := GenS5{F0: 2147483647, F1: "\n\t", F2: -9223372036854775808}
                                    _ = v37
                                case "hello":
                                    for fi38 := 0; fi38 < 20; fi38++ {
                                        fmt.Println(fi34)
                                        _ = fi38
                                    }
                                }
                                _ = v35
                                _ = v36
                            } else {
                                for fi39 := 0; fi39 < 11; fi39++ {
                                    var v40 bool = false
                                    _ = fi39
                                    _ = v40
                                }
                            }
                            _ = fi34
                        }
                        var v41 bool = false
                        _ = v41
                    } else {
                        if (1 >= 42) {
                            if len(v2) > 0 {
                                v2[0] = (v13 << ((genFn0(-770)) & 63))
                            }
                        }
                    }
                    v42 := GenS6{F0: 2, F1: 953}
                    _ = v42
                }
            case 4294967295:
                v14 = v14
            default:
                if false {
                    switch (v14 << ((828) & 63)) {
                    case 4294967295:
                        v43 := GenS2{F0: "foo bar"}
                        _ = v43
                    case -159:
                        var v44 uint16 = 65535
                        _ = v44
                    }
                    v45 := GenS7{F0: 2147483647, F1: -551, F2: 127}
                    _ = v45
                } else {
                    v1 = (v1 - v1)
                }
            }
        }
    }
    var v46 bool = (!((v1 <= ((-2.25 + (v1 * 0.0)) / 3.14159)) && (128 == 214)))
    var v47 uint16 = 2
    var v48 int16 = -32768
    var v49 uint16 = (v47 % (((v47 << ((442) & 63))) | 1))
    v46 = (689 != 873)
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
