package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int32; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 string; F1 int32; F2 int32 }

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]bool)
    v1 = v1
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{fi6, (fi6 + -1), fi7, int(v4)}
            v1 = 1
            if true {
                v4 += v4
                v8 = append(v8, int(v4))
            }
            var v9 float64 = 3.14159
            v9 = ((1e10 * (v9 / v9)) * v9)
            v2[2147483648] = true
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 uint = 195
    if true {
        v4 = (v4 | 4294967295)
    }
    switch "a" {
    case "foo bar":
        v4 = v4
    case "":
        var v11 uint = v1
        _ = v11
    }
    if false {
        var v12 int64 = -461
        fmt.Println(v1, v1)
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := []int{}
            _ = fi13
            _ = v14
        }
        _ = v12
    } else {
        v15 := GenS1{F0: v4}
        _ = v15
    }
    if (1.5 > 3.14159) {
        var v16 uint32 = 5
        if (42 <= 128) {
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS2{F0: (v4 + (v4 & v4))}
                _ = fi17
                _ = v18
            }
            v4 = 0
            if false {
                var v19 uint8 = 30
                if (42 < uint16(v1)) {
                    v4 = v4
                }
                var v20 string = ("hello" + "foo bar")
                _ = v19
                _ = v20
            } else {
                v21 := GenS2{F0: (^(v4))}
                v22 := GenS3{F0: false, F1: 2147483647, F2: -128}
                _ = v21
                _ = v22
            }
        } else {
            if true {
                for fi23 := 0; fi23 < 19; fi23++ {
                    v24 := GenS4{F0: "\n\t"}
                    switch (-(fi23)) {
                    case 48:
                        var v25 uint16 = 2
                        _ = v25
                    case -588:
                        if false {
                            var v26 uint32 = v16
                            v27 := GenS1{F0: (v4 - (^(v4)))}
                            _ = v26
                            _ = v27
                        }
                    default:
                        v28 := GenS4{F0: "\n\t"}
                        _ = v28
                    }
                    _ = fi23
                    _ = v24
                }
                var v29 uint8 = 2
                _ = v29
            } else {
                v10 = (v1 | 1)
                v30 := GenS4{F0: ("hello" + "\n\t")}
                var v31 uint16 = 255
                _ = v30
                _ = v31
            }
        }
        v32 := GenS2{F0: (v4 >> ((18446744073709551615) & 63))}
        _ = v16
        _ = v32
    }
    if (true && (("\n\t" + ("" + "")) > "foo bar")) {
        v4 += (v4 % ((v4) | 1))
        v10 = 501
    } else {
        v33 := GenS5{F0: true, F1: true}
        _ = v33
    }
    v10 = (^(v10))
    v1 += v10
    switch ("a" + "foo bar") {
    case "\n\t":
        var v34 int = 1
        _ = v34
    case "":
        if (true || false) {
            v35 := GenS3{F0: (false && ((1 <= 2147483647) && (((v10 >> ((v1) & 63)) > (v10 - v1)) && ((9223372036854775807 < 1) && true)))), F1: 223, F2: -128}
            switch -881 {
            case 32767:
                _ = len(v2)
            case 384:
                for fi36 := 0; fi36 < 9; fi36++ {
                    var v37 int32 = 42
                    v1 += v10
                    _ = fi36
                    _ = v37
                }
            default:
                for fi38 := 0; fi38 < 19; fi38++ {
                    v10 -= 255
                    v5[4294967295] = (!false)
                    _ = fi38
                }
            }
            v10 = v10
            _ = v35
        }
    case "hello":
        v1 += 0
    default:
        v39 := []int{-9223372036854775808}
        _ = v39
    }
    if true {
        v10 += v1
        for fi40 := 0; fi40 < 10; fi40++ {
            fmt.Println(v10, v10, v4)
            v41 := []int{}
            _ = fi40
            _ = v41
        }
        var v42 bool = (-1 < 157)
        _ = v42
    }
    _ = len(v5)
    v43 := []int{920, 128, 32767}
    var v44 uint8 = 60
    for fi45 := 0; fi45 < 9; fi45++ {
        v46 := GenS6{F0: (v10 ^ (^(v1))), F1: 91}
        switch 128 {
        case 0:
            switch ("hello" + "hello") {
            case "a":
                _ = len(v3)
            default:
                for fi47 := 0; fi47 < 6; fi47++ {
                    _ = len(v43)
                    fmt.Println(v4, v4, v10)
                    v44 = v44
                    if false {
                        v1 = (v10 * 4294967295)
                        var v48 uint = (v1 | v10)
                        v49 := GenS7{F0: 42, F1: 1}
                        _ = v48
                        _ = v49
                    } else {
                        var v50 int64 = 1
                        if (1 > 4294967295) {
                            var v51 uint16 = 0
                            _ = v51
                        } else {
                            v52 := GenS8{F0: ("foo bar" + ("foo bar" + "foo bar")), F1: -2147483648, F2: 205}
                            var v53 uint = (v1 % ((705) | 1))
                            _ = v52
                            _ = v53
                        }
                        fmt.Println(fi45)
                        if len(v43) > 0 {
                            v43[0] = (fi45 << ((0) & 63))
                        }
                        fmt.Println(v1)
                        _ = v50
                    }
                    var v54 float64 = 3.14159
                    var v55 float64 = (v54 - v54)
                    v1 = 4294967295
                    if len(v43) > 0 {
                        v43[0] = 0
                    }
                    var v56 int8 = 100
                    _ = fi47
                    _ = v54
                    _ = v55
                    _ = v56
                }
            }
        default:
            _ = len(v43)
        }
        for fi57 := 0; fi57 < 4; fi57++ {
            v44 -= uint8(v10)
            _ = fi57
        }
        _ = fi45
        _ = v46
    }
    var v58 uint16 = 2
    var v59 uint32 = 926
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v58
    _ = v59
}
