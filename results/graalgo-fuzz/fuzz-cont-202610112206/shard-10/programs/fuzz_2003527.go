package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 uint8, p1 int8, p2 int32) string {
    fmt.Println(p2, p1, p0)
    p2 = int32(p1)
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}
func genFn1() uint8 {
    v5 := GenS1{F0: 0}
    switch "hello" {
    case "\n\t":
        if true {
            var v6 int8 = -57
            v7 := []int{1, 32767, 390}
            _ = v6
            _ = v7
        }
    case "hello":
        var v8 uint8 = 0
        _ = v8
    }
    var v9 int64 = -760
    var v10 int32 = 988
    var v11 int8 = -1
    _ = v10
    _ = v11
    _ = v5
    _ = v9
    return 1
}

func main() {
    v1 := []int{-2147483648, 32767}
    var v2 uint32 = 450
    var v3 bool = (!(true || (!true)))
    var v4 uint = 2
    for fi12 := 0; fi12 < 19; fi12++ {
        v3 = (18446744073709551615 > v4)
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            fmt.Println(fi12, fi12)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (2147483647 == -2147483648) {
                switch ("hello" + "foo bar") {
                case "":
                    if v3 {
                        var v14 float32 = float32(0.1)
                        switch "foo bar" {
                        case "hello":
                            v15 := GenS1{F0: -9223372036854775808}
                            _ = v15
                        case "\n\t":
                            switch ("hello" + "") {
                            case "a":
                                v3 = (1 >= 2)
                            case "hello":
                                if (false && true) {
                                    var v16 int = -2147483648
                                    var v17 int64 = -73
                                    _ = v16
                                    _ = v17
                                }
                            case "\n\t":
                                var v18 uint16 = 1
                                _ = v18
                            default:
                                var v19 uint64 = 42
                                _ = v19
                            }
                        default:
                            fmt.Println(fi13, v3)
                        }
                        var v20 int16 = 42
                        _ = v14
                        _ = v20
                    } else {
                        var v21 float32 = float32(1.5)
                        _ = len(v1)
                        if ((("a" + ((("a" + "") + "") + "")) <= "\n\t") && (9223372036854775807 <= 1)) {
                            var v22 uint16 = 2
                            v2 -= uint32(v22)
                            _ = v22
                        }
                        _ = v21
                    }
                case "foo bar":
                    switch "" {
                    case "":
                        switch "hello" {
                        case "foo bar":
                            v3 = v3
                        default:
                            switch (^(fi13)) {
                            case 127:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case 4294967295:
                                v4 = (v4 >> ((v4) & 63))
                            case -925:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v23 := GenS1{F0: 74}
                                _ = v23
                            }
                        }
                    }
                case "a":
                    fmt.Println(fi12, v3, v2)
                }
                v4 -= v4
            }
            _ = fi13
        }
        _ = fi12
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        v2 = (v2 ^ v2)
        if v3 {
            if false {
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 int32 = 336
                    var v27 uint16 = 255
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        } else {
            v3 = (!(-2147483648 == -2147483648))
            var v28 int = fi24
            _ = v28
        }
        _ = fi24
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if ((^(v4)) != 434) {
            if v3 {
                if len(v1) > 0 {
                    v1[0] = 110
                }
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 int16 = -1
                    fmt.Println(v3, v3)
                    v2 += 0
                    _ = fi29
                    _ = v30
                }
            }
            if true {
                switch -490 {
                case -454:
                    for fi31 := 0; fi31 < 9; fi31++ {
                        fmt.Println(v4)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v4)
                        _ = fi31
                    }
                }
                switch 4294967295 {
                case 2147483647:
                    var v32 uint64 = 4294967295
                    _ = v32
                case 0:
                    for fi33 := 0; fi33 < 17; fi33++ {
                        v3 = (v3 || v3)
                        _ = fi33
                    }
                default:
                    v1 = append(v1, 2147483647)
                }
            } else {
                v34 := GenS2{F0: 127, F1: -1}
                if len(v1) > 0 {
                    v1[0] = 127
                }
                v35 := GenS1{F0: 0}
                _ = v34
                _ = v35
            }
            v3 = true
        }
    }
    v36 := GenS3{F0: (v2 | v2)}
    v2 = 4294967295
    if (v2 >= v2) {
        v37 := GenS2{F0: 42, F1: 1}
        var v38 bool = ("\n\t" > "\n\t")
        _ = v37
        _ = v38
    } else {
        if true {
            v4 = v4
            v39 := GenS4{F0: -2147483648, F1: 42, F2: -1}
            v1 = append(v1, -408)
            _ = v39
        }
        v40 := GenS3{F0: (v2 / (((v2 / ((v2) | 1))) | 1))}
        _ = v40
    }
    var v41 float32 = 1.5
    v42 := GenS3{F0: v2}
    v43 := GenS1{F0: 0}
    for fi44 := 0; fi44 < 13; fi44++ {
        if (true || true) {
            for fi45 := 0; fi45 < 16; fi45++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v46 int64 = 42
                _ = fi45
                _ = v46
            }
            fmt.Println(v4, fi44, v2)
        }
        v4 += 18446744073709551615
        fmt.Println(fi44, fi44, v2)
        _ = fi44
    }
    _ = len(v1)
    fmt.Println(v4)
    var v47 int16 = -1
    var v48 uint64 = uint64(v4)
    v48 = 4294967295
    v4 = uint(v2)
    v49 := GenS1{F0: 9223372036854775807}
    v1 = append(v1, 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v47
    _ = v48
    _ = v49
}
