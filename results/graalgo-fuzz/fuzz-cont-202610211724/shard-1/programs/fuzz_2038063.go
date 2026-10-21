package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 uint64; F2 int8 }
func genFn0(p0 int64, p1 string, p2 uint32) string {
    v7 := []int{1, 32767, -9223372036854775808}
    if (((p0 ^ 1) == 42) && (false && false)) {
        for fi8 := 0; fi8 < 20; fi8++ {
            switch p1 {
            case "":
                var v9 bool = (((127 > 42) || (!(uint8(p2) <= 0))) && (true || true))
                _ = v9
            default:
                v10 := GenS1{F0: 42}
                _ = v10
            }
            if true {
                p0 = 9223372036854775807
            } else {
                var v11 float32 = 1.5
                _ = v11
            }
            switch p1 {
            case "":
                if false {
                    p1 = (p1 + ("a" + "foo bar"))
                    fmt.Println(p1)
                    fmt.Println(fi8, p2)
                }
            }
            _ = fi8
        }
        v12 := GenS1{F0: 62}
        v7 = append(v7, 2147483647)
        _ = v12
    }
    var v13 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    return p1
}
func genFn1(p0 int16, p1 bool) int32 {
    if true {
        var v14 int8 = 114
        _ = v14
    } else {
        p1 = ((true || (127 == 160)) && p1)
        var v15 uint64 = 18446744073709551615
        _ = v15
    }
    if (!p1) {
        var v16 uint64 = 439
        p1 = (float64(float32(1.5)) > (1.5 - (3.14159 * ((((0.1 * ((0.0 * 3.14159) - 1e10)) + 3.14159) * 1.5) + float64(-2.25)))))
        _ = v16
    } else {
        v17 := []int{42}
        fmt.Println(p0, p0)
        _ = v17
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483647, 2147483647}
    v3 := []string{"hello"}
    var v4 uint8 = 127
    v5 := make(map[uint32]bool)
    var v6 int = -1
    for fi18 := 0; fi18 < 16; fi18++ {
        if (0 == 179) {
            var v19 uint8 = (^(v4))
            if false {
                switch ("\n\t" + "a") {
                case "hello":
                    var v20 int8 = 1
                    _ = v20
                case "foo bar":
                    switch (v6 + (fi18 % ((fi18) | 1))) {
                    case -78:
                        for fi21 := 0; fi21 < 12; fi21++ {
                            var v22 string = "hello"
                            _ = fi21
                            _ = v22
                        }
                    }
                case "a":
                    var v23 int32 = -191
                    _ = v23
                }
                if len(v3) > 0 {
                    v3[0] = ("hello" + "hello")
                }
                if false {
                    switch 2147483647 {
                    case 42:
                        v24 := GenS2{F0: 568}
                        _ = v24
                    default:
                        v25 := GenS2{F0: 0}
                        _ = v25
                    }
                } else {
                    if (0 >= 4294967295) {
                        v19 = 128
                    } else {
                        fmt.Println(v6)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(fi18)
                    }
                    var v26 float64 = ((0.0 - (3.14159 / 0.1)) / 1.0)
                    _ = v26
                }
            } else {
                switch "foo bar" {
                case "\n\t":
                    switch "" {
                    case "a":
                        for fi27 := 0; fi27 < 10; fi27++ {
                            switch (-(fi18)) {
                            case -1:
                                _ = len(v3)
                            }
                            fmt.Println(v6, v19, fi27)
                            _ = fi27
                        }
                    case "hello":
                        if true {
                            var v28 int64 = 42
                            var v29 uint8 = 0
                            _ = v28
                            _ = v29
                        } else {
                            for fi30 := 0; fi30 < 15; fi30++ {
                                var v31 int32 = -1
                                if true {
                                    if (!true) {
                                        v32 := GenS1{F0: v4}
                                        _ = v32
                                    }
                                } else {
                                    v4 = v4
                                    var v33 uint16 = 670
                                    _ = v33
                                }
                                var v34 uint64 = 272
                                _ = fi30
                                _ = v31
                                _ = v34
                            }
                            for fi35 := 0; fi35 < 11; fi35++ {
                                if (false || false) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v36 := GenS1{F0: v19}
                                        _ = v36
                                    case "hello":
                                        if (v4 != (v19 + (v19 >> ((2) & 63)))) {
                                            fmt.Println(v6)
                                            fmt.Println(fi35)
                                        } else {
                                            fmt.Println(v19, fi18)
                                        }
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    v19 = (^(v19))
                                    v5[42] = true
                                }
                                v37 := GenS3{F0: (32767 == 994)}
                                v38 := GenS4{F0: 589, F1: 2147483648}
                                switch (v6 >> ((618) & 63)) {
                                case 633:
                                    fmt.Println(fi35, fi18, v4)
                                case 319:
                                    if false {
                                        v39 := v1[fi18]
                                        v40 := GenS5{F0: "a"}
                                        _ = v39
                                        _ = v40
                                    }
                                case 42:
                                    v19 = 127
                                default:
                                    fmt.Println(fi35)
                                }
                                _ = fi35
                                _ = v37
                                _ = v38
                            }
                        }
                    default:
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            v4 -= (^(v4))
                        case "hello":
                            var v41 string = (("hello" + "a") + ("a" + "a"))
                            _ = v41
                        default:
                            v42 := GenS2{F0: 4294967295}
                            _ = v42
                        }
                    }
                }
            }
            var v43 int8 = -35
            _ = v19
            _ = v43
        } else {
            _ = len(v2)
        }
        _ = fi18
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        switch fi44 {
        case 128:
            for fi45 := 0; fi45 < 16; fi45++ {
                var v46 int = (v6 - fi44)
                v47 := v5[2147483648]
                _ = fi45
                _ = v46
                _ = v47
            }
        case 32768:
            var v48 uint16 = 255
            _ = v48
        default:
            v5[352] = (fi44 <= (fi44 | 128))
        }
        _ = fi44
    }
    v49 := GenS6{F0: (("a" + ("a" + "")) + "a"), F1: 42, F2: 127}
    v6 += v6
    var v50 int64 = 42
    var v51 uint16 = 65535
    var v52 bool = ((1 >= 65535) && (false || (false && false)))
    fmt.Println(v4)
    var v53 int = 1
    var v54 int = v6
    fmt.Println(v50, v53, v50)
    v55 := GenS1{F0: (v4 ^ v4)}
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
