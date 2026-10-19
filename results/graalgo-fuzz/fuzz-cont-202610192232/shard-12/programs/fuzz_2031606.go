package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 uint; F1 int8; F2 string }
func genFn0(p0 uint8, p1 bool, p2 uint64) uint32 {
    var v5 uint8 = 0
    switch -9223372036854775808 {
    case 4294967295:
        if false {
            var v6 uint32 = 955
            var v7 uint8 = v5
            if true {
                for fi8 := 0; fi8 < 10; fi8++ {
                    switch ("" + "foo bar") {
                    case "hello":
                        v9 := GenS1{F0: 2, F1: -1}
                        _ = v9
                    case "\n\t":
                        var v10 float64 = 1.5
                        _ = v10
                    default:
                        p2 = (p2 << ((18446744073709551615) & 63))
                    }
                    v11 := GenS1{F0: v6, F1: 23}
                    _ = fi8
                    _ = v11
                }
                var v12 uint32 = 0
                _ = v12
            } else {
                var v13 uint = 4294967295
                _ = v13
            }
            _ = v6
            _ = v7
        }
    case 2147483647:
        v14 := []int{-519, 0}
        _ = v14
    }
    v15 := GenS2{F0: 42, F1: (p2 ^ 31), F2: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v5
    return 2147483647
}
func genFn1(p0 uint8, p1 uint8, p2 int16) int {
    var v16 uint = 42
    v17 := GenS3{F0: 1}
    v16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return int(p0)
}

func main() {
    var v1 uint8 = 121
    var v2 uint = 255
    var v3 bool = true
    var v4 uint = v2
    for fi18 := 0; fi18 < 13; fi18++ {
        if false {
            if false {
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 int64 = 811
                    v21 := []int{-750, fi19, fi19, fi19}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                var v22 bool = ((false || false) && true)
                v23 := []int{fi18, (fi18 + genFn1((v1 + 2), v1, 1))}
                _ = v22
                _ = v23
            }
            var v24 int8 = 42
            _ = v24
        }
        if v3 {
            var v25 float32 = 1.5
            var v26 uint8 = 253
            _ = v25
            _ = v26
        } else {
            v4 = v4
            if true {
                switch "\n\t" {
                case "hello":
                    for fi27 := 0; fi27 < 14; fi27++ {
                        v28 := GenS2{F0: 32767, F1: 18446744073709551615, F2: -1}
                        _ = fi27
                        _ = v28
                    }
                case "":
                    v29 := GenS3{F0: 0}
                    _ = v29
                case "foo bar":
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v2 += v2
                        _ = fi30
                    }
                }
                fmt.Println(v3, v3, v2)
            }
            if (!true) {
                if (2147483647 > 318) {
                    v31 := GenS3{F0: -240}
                    switch genFn1((v1 | 127), (v1 ^ 42), 766) {
                    case -2147483648:
                        v32 := []int{-174}
                        _ = v32
                    case 32768:
                        v33 := GenS3{F0: -143}
                        _ = v33
                    }
                    var v34 int32 = 2147483647
                    _ = v31
                    _ = v34
                }
                v3 = (v4 > (v2 + v4))
            } else {
                var v35 int32 = -282
                v4 -= v2
                switch int(v35) {
                case 128:
                    v1 -= (v1 - 44)
                case 32768:
                    v35 = v35
                default:
                    var v36 bool = true
                    _ = v36
                }
                _ = v35
            }
        }
        var v37 bool = v3
        _ = fi18
        _ = v37
    }
    var v38 int32 = -170
    v39 := []int{589, 543, -659}
    if v3 {
        v40 := GenS4{F0: false}
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
        _ = v40
    } else {
        for fi41 := 0; fi41 < 12; fi41++ {
            v38 -= (v38 / (((v38 & v38)) | 1))
            v42 := GenS5{F0: v3, F1: genFn1(v1, uint8(v2), 42)}
            _ = fi41
            _ = v42
        }
        v3 = (v38 == v38)
    }
    v39 = append(v39, 32767)
    if false {
        fmt.Println(v4, v38, v3)
        fmt.Println(v3)
    }
    v43 := GenS1{F0: 2147483647, F1: 32767}
    v44 := GenS6{F0: -128, F1: "hello"}
    v2 += v2
    for fi45 := 0; fi45 < 14; fi45++ {
        var v46 int64 = 819
        _ = fi45
        _ = v46
    }
    var v47 string = "\n\t"
    if v3 {
        v4 += 4294967295
        _ = len(v39)
        v48 := GenS5{F0: ("" != "hello"), F1: 1}
        _ = v48
    }
    for fi49 := 0; fi49 < 20; fi49++ {
        fmt.Println(v47, v47)
        _ = fi49
    }
    v50 := GenS7{F0: (v2 >> ((42) & 63)), F1: -128, F2: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v47
    _ = v50
}
