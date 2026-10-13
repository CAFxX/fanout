package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint64 }
func genFn0(p0 int16) int8 {
    fmt.Println(p0)
    var v5 int = -140
    if true {
        var v6 bool = (1 != 127)
        switch ("\n\t" + "hello") {
        case "a":
            for fi7 := 0; fi7 < 12; fi7++ {
                if (!(!v6)) {
                    if (0 > 18446744073709551615) {
                        switch 42 {
                        case 4294967295:
                            v8 := []int{(fi7 + fi7), (v5 / ((v5) | 1))}
                            _ = v8
                        case -1:
                            v9 := []int{}
                            _ = v9
                        default:
                            var v10 uint16 = 65535
                            _ = v10
                        }
                        var v11 uint32 = 127
                        _ = v11
                    } else {
                        p0 = p0
                    }
                    var v12 int8 = 122
                    p0 = p0
                    _ = v12
                }
                _ = fi7
            }
        case "":
            v5 -= 32768
        case "hello":
            p0 = 0
        }
        _ = v6
    }
    _ = p0
    _ = v5
    return 127
}
func genFn1(p0 bool, p1 uint16, p2 string) bool {
    switch -1 {
    case -1:
        v13 := GenS1{F0: -873}
        _ = v13
    case 548:
        switch 9223372036854775807 {
        case 127:
            p1 -= 0
        case 322:
            for fi14 := 0; fi14 < 16; fi14++ {
                switch ("hello" + p2) {
                case "a":
                    if ((3.14159 - 1.5) <= 3.14159) {
                        switch p2 {
                        case "":
                            p1 = (^(p1))
                        case "hello":
                            p2 = p2
                        case "foo bar":
                            p0 = p0
                        default:
                            for fi15 := 0; fi15 < 15; fi15++ {
                                v16 := []int{}
                                switch (p2 + "a") {
                                case "a":
                                    fmt.Println(p1, fi15, fi14)
                                default:
                                    if false {
                                        if false {
                                            var v17 uint8 = 2
                                            _ = len(v16)
                                            _ = v17
                                        }
                                        var v18 int32 = 0
                                        _ = v18
                                    }
                                    var v19 int = fi15
                                    _ = v19
                                }
                                _ = fi15
                                _ = v16
                            }
                        }
                        if (!p0) {
                            var v20 uint8 = 2
                            if (true && (!(42 == 42))) {
                                p1 -= p1
                                var v21 int16 = -1
                                switch (fi14 ^ (fi14 >> ((18446744073709551615) & 63))) {
                                case -2147483648:
                                    p1 = p1
                                case -524:
                                    for fi22 := 0; fi22 < 13; fi22++ {
                                        var v23 uint8 = 2
                                        v24 := []int{}
                                        var v25 bool = (4294967295 >= 2)
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                case 2147483647:
                                    for fi26 := 0; fi26 < 16; fi26++ {
                                        var v27 uint = 255
                                        v28 := GenS2{F0: (^(v27)), F1: p1}
                                        var v29 uint = v27
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                default:
                                    var v30 uint = 18446744073709551615
                                    _ = v30
                                }
                                _ = v21
                            }
                            _ = v20
                        }
                    }
                default:
                    if p0 {
                        v31 := GenS1{F0: 590}
                        var v32 float32 = 1e10
                        _ = v31
                        _ = v32
                    }
                }
                if (0 != 65535) {
                    v33 := GenS1{F0: 1000}
                    _ = v33
                }
                _ = fi14
            }
        case -2147483648:
            var v34 int = 32768
            _ = v34
        }
    default:
        var v35 uint8 = 108
        _ = v35
    }
    v36 := GenS3{F0: false, F1: -25, F2: 1}
    if p0 {
        p1 = (p1 ^ p1)
    } else {
        var v37 float32 = ((1.5 + 1.5) + 0.0)
        v38 := GenS4{F0: (p1 * p1), F1: 2}
        v39 := []int{9223372036854775807, 42, -2147483648, 128}
        _ = v37
        _ = v38
        _ = v39
    }
    v40 := GenS1{F0: 430}
    for fi41 := 0; fi41 < 10; fi41++ {
        v42 := []int{fi41, fi41, (fi41 % ((-1) | 1)), (fi41 ^ 4294967295)}
        _ = fi41
        _ = v42
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    _ = v40
    return p0
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = -2147483648
    var v3 int16 = 276
    var v4 int64 = -9223372036854775808
    var v43 string = ("foo bar" + "a")
    v44 := []int{32768}
    fmt.Println(v3)
    var v45 int64 = (v4 >> ((18446744073709551615) & 63))
    v45 -= v4
    switch 32767 {
    case -2147483648:
        var v46 uint16 = 65535
        _ = v46
    case -9223372036854775808:
        fmt.Println(v45, v45)
    case 32768:
        fmt.Println(v43, v4)
    }
    var v47 string = "foo bar"
    var v48 float64 = 1.5
    for fi49 := 0; fi49 < 16; fi49++ {
        v3 = (v3 ^ v3)
        _ = fi49
    }
    fmt.Println(v47, v45, v2)
    if (true || true) {
        for fi50 := 0; fi50 < 18; fi50++ {
            fmt.Println(v4, v1)
            _ = fi50
        }
    } else {
        v51 := GenS1{F0: -612}
        fmt.Println(v48, v43, v1)
        _ = v51
    }
    v52 := GenS5{F0: -2147483648, F1: v1}
    v47 = (v43 + ("" + ("" + v43)))
    var v53 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v52
    _ = v53
}
