package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int16; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint16 }
func genFn0(p0 string, p1 bool) int64 {
    var v5 uint8 = 19
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: fi6, F1: -1, F2: (p0 + (p0 + (p0 + (p0 + p0))))}
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    if (1e10 >= (1.5 + 1e10)) {
        var v9 int8 = 1
        var v10 string = (p0 + "foo bar")
        _ = v10
        _ = v9
    }
    var v11 uint16 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    _ = v8
    return 1
}
func genFn1() int32 {
    var v12 uint64 = 42
    var v13 string = "hello"
    v14 := []int{-1}
    v15 := GenS1{F0: -419, F1: -32768, F2: "\n\t"}
    switch v13 {
    case "foo bar":
        v13 = v13
    case "":
        var v16 int = -819
        _ = v16
    }
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    var v1 int = 90
    var v2 bool = ((0 >= 1) || true)
    var v3 int8 = 100
    var v4 uint32 = 1
    var v17 uint32 = (^(v4))
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 17; fi19++ {
            if v2 {
                v1 -= fi19
                v20 := []int{}
                _ = len(v20)
                _ = v20
            }
            _ = fi19
        }
        v21 := []int{(v1 - v1), fi18, v1, fi18}
        v1 = 4294967295
        _ = fi18
        _ = v21
    }
    v22 := []int{v1, (v1 >> ((302) & 63)), v1}
    var v23 uint16 = 255
    var v24 int64 = -1
    v25 := GenS1{F0: 451, F1: -1, F2: "\n\t"}
    v2 = (514 <= -1)
    v22 = append(v22, (v1 + (v1 << ((805) & 63))))
    if (v1 == (^(v1))) {
        v17 = v17
        switch ("a" + "\n\t") {
        case "a":
            v26 := GenS2{F0: 34}
            _ = v26
        }
        v24 -= -606
    } else {
        switch v1 {
        case 0:
            switch "" {
            case "a":
                var v27 uint = 255
                _ = v27
            case "hello":
                fmt.Println(v2)
            default:
                v24 += (v24 | v24)
            }
        case -9223372036854775808:
            switch ((("" + "") + "") + ("foo bar" + ((("a" + "foo bar") + "\n\t") + "\n\t"))) {
            case "\n\t":
                v28 := GenS2{F0: uint8(v23)}
                _ = v28
            case "":
                for fi29 := 0; fi29 < 20; fi29++ {
                    if false {
                        fmt.Println(v17, v2, v24)
                    } else {
                        var v30 string = ""
                        for fi31 := 0; fi31 < 14; fi31++ {
                            v32 := GenS3{F0: 1, F1: 0, F2: 84}
                            _ = fi31
                            _ = v32
                        }
                        if false {
                            if false {
                                v33 := GenS4{F0: 18446744073709551615, F1: 42, F2: 255}
                                v4 = 42
                                _ = v33
                            } else {
                                var v34 float64 = -2.25
                                _ = v34
                            }
                            var v35 uint = 4294967295
                            _ = v35
                        } else {
                            var v36 uint16 = 276
                            var v37 int32 = -578
                            _ = v36
                            _ = v37
                        }
                        _ = v30
                    }
                    v22 = append(v22, (fi29 - fi29))
                    _ = fi29
                }
            case "a":
                switch (v1 / (((-(v1))) | 1)) {
                case 128:
                    _ = len(v22)
                case 0:
                    var v38 uint8 = 255
                    _ = v38
                }
            }
        default:
            v39 := GenS1{F0: (v1 ^ v1), F1: -32768, F2: "\n\t"}
            _ = v39
        }
    }
    var v40 int8 = (v3 | 15)
    v41 := GenS5{F0: 128, F1: (^(v23))}
    if v2 {
        if v2 {
            if len(v22) > 0 {
                v22[0] = v1
            }
            if v2 {
                v22 = append(v22, v1)
                v22 = append(v22, (v1 + v1))
                if (((v1 & v1) > (v1 - (v1 / ((v1) | 1)))) && (!((v3 ^ 42) >= 42))) {
                    v1 -= (v1 << ((446) & 63))
                    if false {
                        fmt.Println(v3, v1)
                        fmt.Println(v3)
                    } else {
                        v42 := GenS2{F0: 100}
                        for fi43 := 0; fi43 < 6; fi43++ {
                            switch ("a" + "") {
                            case "hello":
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            case "a":
                                var v44 int = (^(v1))
                                _ = v44
                            case "foo bar":
                                v22 = append(v22, v1)
                            }
                            v45 := GenS6{F0: genFn1(), F1: 41, F2: v23}
                            v3 = v40
                            _ = fi43
                            _ = v45
                        }
                        v22 = append(v22, (v1 | v1))
                        _ = v42
                    }
                }
            }
        } else {
            if false {
                var v46 int8 = (v40 - v40)
                fmt.Println(v17)
                _ = v46
            } else {
                var v47 uint32 = v17
                _ = v47
            }
            for fi48 := 0; fi48 < 9; fi48++ {
                var v49 uint16 = v23
                _ = fi48
                _ = v49
            }
        }
        v50 := GenS1{F0: v1, F1: 42, F2: "foo bar"}
        _ = v50
    } else {
        v3 -= v40
        for fi51 := 0; fi51 < 20; fi51++ {
            var v52 int16 = 42
            if ((false && true) || (!v2)) {
                var v53 int64 = 1
                _ = v53
            }
            _ = fi51
            _ = v52
        }
        v1 = v1
    }
    v54 := GenS1{F0: (v1 - v1), F1: 42, F2: ""}
    v55 := GenS1{F0: v1, F1: -786, F2: "\n\t"}
    v56 := GenS6{F0: -1, F1: 206, F2: 0}
    v23 = (v23 * (v23 | (^(v23))))
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v54
    _ = v55
    _ = v56
}
