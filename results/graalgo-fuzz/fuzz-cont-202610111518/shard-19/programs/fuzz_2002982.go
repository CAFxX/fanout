package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint; F2 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int; F1 int16 }
func genFn0() bool {
    v6 := GenS1{F0: 42}
    v7 := GenS2{F0: (9223372036854775807 > -354), F1: false}
    var v8 int64 = 9223372036854775807
    if true {
        v8 = v8
        v8 -= v8
    }
    _ = v6
    _ = v7
    _ = v8
    return true
}
func genFn1() int8 {
    var v9 string = ("\n\t" + "hello")
    switch 4294967295 {
    case -9223372036854775808:
        v9 = v9
    case 32768:
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint32 = 42
            v9 = (v9 + ((v9 + "hello") + v9))
            _ = fi10
            _ = v11
        }
    default:
        v12 := GenS3{F0: 65535, F1: 0}
        _ = v12
    }
    var v13 bool = (true || true)
    _ = v13
    _ = v9
    return -1
}

func main() {
    var v1 float64 = 1.5
    var v2 int16 = -348
    v3 := []string{(("\n\t" + (("foo bar" + ("a" + "")) + "\n\t")) + "hello"), "", "hello", ""}
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]int)
    v1 = v1
    v2 -= (^(v2))
    var v14 int16 = v2
    v15 := GenS4{F0: 2, F1: (432 >= 18446744073709551615)}
    if true {
        v16 := GenS5{F0: 9223372036854775807}
        v17 := GenS6{F0: 4294967295, F1: 2, F2: (!(442 > 232))}
        _ = v16
        _ = v17
    } else {
        v1 = v1
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 bool = true
            v20 := v5[2]
            fmt.Println(fi18, v14, v19)
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 float64 = v1
        _ = v21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        if false {
            var v23 uint32 = 382
            _ = v23
        }
        _ = fi22
    }
    v14 = (v14 % ((841) | 1))
    v14 = v14
    var v24 int16 = (v14 << ((4294967295) & 63))
    if true {
        switch "foo bar" {
        case "a":
            switch "foo bar" {
            case "hello":
                switch "\n\t" {
                case "":
                    fmt.Println(v2, v2, v2)
                case "\n\t":
                    var v25 bool = (!(4294967295 > 227))
                    _ = v25
                case "a":
                    for fi26 := 0; fi26 < 17; fi26++ {
                        for fi27 := 0; fi27 < 20; fi27++ {
                            delete(v5, 701)
                            _ = fi27
                        }
                        v28 := GenS7{F0: -488}
                        v14 = 1
                        _ = fi26
                        _ = v28
                    }
                default:
                    var v29 string = "foo bar"
                    _ = v29
                }
            case "":
                v2 = v14
            }
        default:
            if true {
                v2 += (^(v24))
                delete(v5, 2)
                v2 = (v14 - v24)
            }
        }
    }
    if true {
        fmt.Println(v2, v24, v14)
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            fmt.Println(fi30, fi30, fi30)
            _ = fi31
        }
        v2 = 0
        _ = fi30
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            var v34 float32 = (1.5 * 1.5)
            _ = fi33
            _ = v34
        }
        if true {
            v35 := GenS6{F0: 69, F1: 450, F2: false}
            _ = v35
        } else {
            fmt.Println(fi32, v24, v24)
            v24 = (v14 | -30)
            for fi36 := 0; fi36 < 1; fi36++ {
                var v37 string = "foo bar"
                v38 := GenS7{F0: 951}
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        if (false || (fi32 > fi32)) {
            delete(v4, "\n\t")
        }
        _ = fi32
    }
    fmt.Println(v14)
    var v39 bool = false
    for fi40 := 0; fi40 < 11; fi40++ {
        if ((!((244 != 0) && false)) && false) {
            v3 = append(v3, "")
            v24 += (v14 | v2)
        } else {
            var v41 int = 1
            var v42 int = (^(fi40))
            _ = v41
            _ = v42
        }
        _ = fi40
    }
    if false {
        delete(v5, 596)
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    } else {
        var v43 float64 = 0.0
        if false {
            v44 := GenS8{F0: 4294967295, F1: 885}
            var v45 int16 = (v2 - (v14 - (v2 & v24)))
            _ = v44
            _ = v45
        }
        _ = v43
    }
    fmt.Println(v14, v14, v39)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
