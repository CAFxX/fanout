package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int32; F2 string }
type GenS5 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint64 = 1
    var v2 bool = false
    var v3 uint8 = 255
    for fi4 := 0; fi4 < 18; fi4++ {
        v2 = (127 <= 42)
        v5 := []int{fi4}
        v1 -= v1
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        v1 = 870
        switch fi6 {
        case 1:
            v7 := []int{(fi6 ^ (fi6 & (^(fi6)))), -1}
            _ = v7
        case 127:
            if true {
                var v8 int32 = 42
                _ = v8
            } else {
                v2 = v2
                v9 := []int{}
                fmt.Println(fi6, fi6, v2)
                _ = v9
            }
        }
        var v10 uint8 = uint8(v1)
        _ = fi6
        _ = v10
    }
    v11 := []int{-2147483648, -9223372036854775808, -1, 42}
    v12 := GenS1{F0: int(v3)}
    var v13 uint32 = 2
    switch "" {
    case "a":
        var v14 int64 = -99
        _ = v14
    case "\n\t":
        v1 = v1
    case "hello":
        v15 := GenS2{F0: true, F1: v1, F2: -924}
        _ = v15
    }
    if len(v11) > 0 {
        v11[0] = 128
    }
    switch -196 {
    case -6:
        for fi16 := 0; fi16 < 11; fi16++ {
            switch "hello" {
            case "\n\t":
                switch fi16 {
                case 128:
                    if v2 {
                        fmt.Println(v13, v1)
                    } else {
                        var v17 uint16 = uint16(v13)
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        _ = v17
                    }
                }
            case "foo bar":
                fmt.Println(v13, v3, v2)
            case "":
                v2 = false
            default:
                for fi18 := 0; fi18 < 17; fi18++ {
                    fmt.Println(v1, fi16)
                    if true {
                        if len(v11) > 0 {
                            v11[0] = fi18
                        }
                    } else {
                        var v19 string = (((("" + "a") + "") + (("hello" + "foo bar") + "foo bar")) + "hello")
                        _ = v19
                    }
                    v20 := GenS2{F0: v2, F1: 142, F2: 42}
                    var v21 bool = v2
                    _ = fi18
                    _ = v20
                    _ = v21
                }
            }
            if true {
                if v2 {
                    switch (fi16 + (fi16 * (fi16 ^ -670))) {
                    case 77:
                        v3 += v3
                    case 32767:
                        for fi22 := 0; fi22 < 3; fi22++ {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            _ = fi22
                        }
                    case 42:
                        v23 := GenS3{F0: 18446744073709551615, F1: -1, F2: (v13 & (v13 & v13))}
                        _ = v23
                    }
                }
                var v24 uint64 = 1
                _ = v24
            }
            _ = fi16
        }
    case 32768:
        if (!((42 == 18446744073709551615) || (false || (!true)))) {
            v3 = v3
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
    case 4294967295:
        var v25 int64 = -1
        _ = v25
    default:
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 bool = (!((v2 || (v2 && v2)) && (-9223372036854775808 == -1)))
            _ = fi26
            _ = v27
        }
    }
    v28 := GenS4{F0: (v1 + v1), F1: 574, F2: ""}
    var v29 float32 = 1e10
    var v30 uint16 = 255
    var v31 uint16 = (v30 / (((^(v30))) | 1))
    if false {
        v30 = v30
        for fi32 := 0; fi32 < 8; fi32++ {
            v33 := GenS2{F0: (true && v2), F1: 4294967295, F2: 1}
            v29 = v29
            _ = fi32
            _ = v33
        }
        v34 := GenS5{F0: v13, F1: v1}
        _ = v34
    } else {
        var v35 uint16 = v31
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
