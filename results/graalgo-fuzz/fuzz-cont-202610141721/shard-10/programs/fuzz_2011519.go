package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint32; F2 string }
type GenS4 struct { F0 bool; F1 uint32; F2 int16 }
func genFn0() uint8 {
    var v6 bool = false
    v7 := []int{-186, -110, -227, 128}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 128
}
func genFn1() int {
    var v8 int64 = 1
    if (2147483647 <= -2147483648) {
        if ((1 < 42) && true) {
            v8 = (-(v8))
            var v9 int16 = 1000
            var v10 int32 = -2147483648
            _ = v10
            _ = v9
        }
        var v11 bool = (0.1 > -2.25)
        fmt.Println(v11, v11)
        _ = v11
    } else {
        var v12 int = -988
        _ = v12
    }
    _ = v8
    return 498
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 559
    var v3 bool = (42 > -44)
    var v4 int64 = -88
    var v5 int16 = -384
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS1{F0: 255}
        v15 := GenS2{F0: 114}
        v16 := v1[""]
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v4 += (v4 - 9223372036854775807)
        for fi18 := 0; fi18 < 10; fi18++ {
            v3 = (42 != 504)
            var v19 string = ""
            v2 += (^(v2))
            switch (fi17 - -400) {
            case 4294967295:
                fmt.Println(fi18)
            case 0:
                var v20 int = (fi18 & 9223372036854775807)
                if true {
                    switch (v19 + ("\n\t" + ("" + (v19 + v19)))) {
                    case "a":
                        v1[v19] = true
                        if (int8(v4) < 2) {
                            var v21 uint16 = 0
                            if (2 > 217) {
                                var v22 int = fi17
                                v2 = v2
                                var v23 int = (-(fi17))
                                _ = v22
                                _ = v23
                            }
                            switch v19 {
                            case "foo bar":
                                v21 = v21
                            case "hello":
                                v24 := []int{(fi18 * 9223372036854775807), v20}
                                _ = v24
                            }
                            var v25 uint = 560
                            v26 := GenS1{F0: (v25 * v25)}
                            _ = v21
                            _ = v25
                            _ = v26
                        }
                        switch v19 {
                        case "":
                            v20 = fi18
                        case "foo bar":
                            _ = len(v1)
                        default:
                            var v27 uint8 = genFn0()
                            _ = v27
                        }
                        var v28 bool = (!v3)
                        v19 = v19
                        _ = v28
                    default:
                        v2 -= v2
                        var v29 uint = 18446744073709551615
                        if true {
                            if (0.0 >= (0.0 + 1e10)) {
                                v1["foo bar"] = (1 != 2147483647)
                            } else {
                                var v30 int64 = v4
                                _ = v30
                            }
                            if (false || v3) {
                                v29 = (v29 + (^(v29)))
                                _ = len(v1)
                                if true {
                                    var v31 float32 = 0.1
                                    switch v19 {
                                    case "\n\t":
                                        var v32 int = 4294967295
                                        _ = v32
                                    case "hello":
                                        v31 = v31
                                    }
                                    if (!false) {
                                        switch v19 {
                                        case "hello":
                                            if (true || ((!v3) || true)) {
                                                switch (v19 + "hello") {
                                                case "":
                                                    v33 := GenS3{F0: (v19 + v19), F1: 42, F2: "\n\t"}
                                                    _ = v33
                                                }
                                                var v34 float64 = (3.14159 * (0.0 - 1.5))
                                                var v35 uint = (v29 - (^(v29)))
                                                _ = v34
                                                _ = v35
                                            } else {
                                                switch "" {
                                                case "":
                                                    v29 += 4294967295
                                                case "foo bar":
                                                    var v36 int32 = -2147483648
                                                    _ = v36
                                                case "a":
                                                    var v37 uint64 = v2
                                                    _ = v37
                                                default:
                                                    var v38 int16 = v5
                                                    _ = v38
                                                }
                                                var v39 int = (^(fi18))
                                                var v40 uint8 = genFn0()
                                                _ = v39
                                                _ = v40
                                            }
                                        case "":
                                            if (235 == 2147483647) {
                                                v41 := []int{genFn1()}
                                                v42 := GenS4{F0: true, F1: 0, F2: v5}
                                                v20 += (fi17 << ((v29) & 63))
                                                _ = v41
                                                _ = v42
                                            } else {
                                                v4 += (v4 >> (((v29 % ((v29) | 1))) & 63))
                                                var v43 uint16 = 2
                                                _ = len(v1)
                                                _ = v43
                                            }
                                        case "\n\t":
                                            v44 := []int{32768, (fi18 << (((^(v29))) & 63)), fi17}
                                            _ = v44
                                        default:
                                            v45 := []int{fi17, genFn1(), (fi17 << ((v29) & 63)), v20}
                                            _ = v45
                                        }
                                    }
                                    if ((v31 / 3.14159) < (0.1 * 0.1)) {
                                        v3 = (!v3)
                                    }
                                    _ = v31
                                } else {
                                    _ = len(v1)
                                }
                            }
                        }
                        v46 := []int{(fi18 / ((-120) | 1)), genFn1()}
                        fmt.Println(v19)
                        _ = v29
                        _ = v46
                    }
                    _ = len(v1)
                    v2 += (v2 * (v2 << ((656) & 63)))
                } else {
                    v47 := GenS4{F0: v3, F1: uint32(v2), F2: v5}
                    _ = v47
                }
                v4 = v4
                _ = v20
            }
            if (float32(3.14159) == 3.14159) {
                fmt.Println(v2)
                _ = len(v1)
                fmt.Println(v2)
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v48 := []int{}
    var v49 uint32 = 4294967295
    var v50 string = "foo bar"
    v2 += 2
    var v51 int = 9223372036854775807
    if v3 {
        var v52 int32 = 402
        var v53 int16 = v5
        _ = v52
        _ = v53
    } else {
        fmt.Println(v5, v5)
        fmt.Println(v49, v3)
    }
    for fi54 := 0; fi54 < 9; fi54++ {
        var v55 uint8 = 128
        if false {
            fmt.Println(v2, v55, v2)
            v49 += 2
        }
        _ = fi54
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
