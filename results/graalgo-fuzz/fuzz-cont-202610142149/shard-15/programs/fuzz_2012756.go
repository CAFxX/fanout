package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 string = ("a" + "foo bar")
    var v2 float64 = 1.5
    var v3 int64 = -9223372036854775808
    var v4 int32 = 42
    var v5 string = v1
    var v6 bool = ("a" != "a")
    v3 += 9223372036854775807
    v2 += v2
    v1 = "foo bar"
    var v7 int32 = v4
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: 278, F1: 143, F2: 127}
            fmt.Println(v7, v2)
            _ = fi9
            _ = v10
        }
        switch ("hello" + ((v5 + "a") + v1)) {
        case "\n\t":
            switch ("" + v1) {
            case "a":
                switch "foo bar" {
                case "\n\t":
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v12 := GenS1{F0: 1, F1: 65535, F2: 100}
                        _ = fi11
                        _ = v12
                    }
                default:
                    v13 := GenS1{F0: 42, F1: 879, F2: 127}
                    _ = v13
                }
            }
        }
        _ = fi8
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        if v6 {
            v15 := []int{fi14}
            var v16 int8 = 1
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS1{F0: 424, F1: 2, F2: 127}
    v7 -= v4
    var v18 int64 = (v3 << ((1) & 63))
    var v19 uint8 = 42
    v20 := GenS2{F0: 4294967295, F1: 100}
    fmt.Println(v5, v2)
    fmt.Println(v6, v1, v7)
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 uint64 = 18446744073709551615
        if false {
            if true {
                var v23 int = fi21
                switch v1 {
                case "foo bar":
                    var v24 int16 = 0
                    _ = v24
                default:
                    v25 := []int{(fi21 - (v23 ^ v23)), 42, (-(fi21))}
                    _ = v25
                }
                for fi26 := 0; fi26 < 14; fi26++ {
                    var v27 int = 227
                    if v6 {
                        switch (fi26 & fi21) {
                        case -1:
                            v23 = (fi26 ^ fi26)
                        case 4294967295:
                            v7 += int32(v22)
                        default:
                            if (360 != 235) {
                                v28 := []int{32768, (v23 - -848), (fi26 * (^(v23))), fi21}
                                var v29 uint16 = 0
                                _ = v28
                                _ = v29
                            } else {
                                fmt.Println(fi26, v6)
                                v6 = ((-462 < 32767) || (!false))
                                fmt.Println(v22, v1)
                            }
                        }
                    } else {
                        if true {
                            v30 := []int{-774, (v27 | (^(fi26)))}
                            var v31 float64 = -2.25
                            if v6 {
                                fmt.Println(v23)
                                var v32 int64 = (v18 << ((768) & 63))
                                v19 -= v19
                                _ = v32
                            }
                            _ = v30
                            _ = v31
                        } else {
                            v33 := GenS3{F0: 42}
                            _ = v33
                        }
                        if v6 {
                            v34 := []int{}
                            var v35 bool = ((float32(((v2 * 1.5) + ((1.5 * 1.5) + (v2 + (v2 + v2))))) * 1e10) != -2.25)
                            _ = v34
                            _ = v35
                        }
                    }
                    switch 127 {
                    case -9223372036854775808:
                        v36 := GenS1{F0: 42, F1: 0, F2: 42}
                        _ = v36
                    case -2147483648:
                        v37 := []int{fi21}
                        _ = v37
                    default:
                        var v38 string = v1
                        var v39 uint16 = uint16(v22)
                        _ = v38
                        _ = v39
                    }
                    var v40 int64 = v3
                    _ = fi26
                    _ = v27
                    _ = v40
                }
                _ = v23
            } else {
                fmt.Println(v18, v3)
            }
            var v41 bool = v6
            _ = v41
        }
        if false {
            v42 := GenS2{F0: 631, F1: -33}
            var v43 int64 = v3
            v6 = (!(v2 >= (1.5 / 1.5)))
            _ = v42
            _ = v43
        } else {
            v44 := []int{}
            _ = v44
        }
        _ = fi21
        _ = v22
    }
    var v45 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
