package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0(p0 int) int8 {
    switch (p0 ^ p0) {
    case 1:
        var v6 int8 = 89
        _ = v6
    case -602:
        fmt.Println(p0)
    default:
        if (-1 > -128) {
            v7 := []int{p0, p0}
            _ = v7
        } else {
            var v8 float32 = (0.0 - 1.5)
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := []int{fi9, fi9}
                switch "hello" {
                case "\n\t":
                    fmt.Println(fi9)
                case "a":
                    v11 := GenS1{F0: (^(p0)), F1: 1, F2: -1}
                    _ = v11
                }
                var v12 int8 = -57
                _ = fi9
                _ = v10
                _ = v12
            }
            var v13 int16 = 1
            _ = v13
            _ = v8
        }
    }
    var v14 uint64 = 42
    p0 = (p0 << ((18446744073709551615) & 63))
    v15 := []int{p0}
    var v16 uint8 = 1
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return -113
}

func main() {
    var v1 int8 = 42
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 683
    var v5 string = v2
    v17 := GenS2{F0: 1000, F1: 100}
    for fi18 := 0; fi18 < 5; fi18++ {
        v5 = "foo bar"
        var v19 uint32 = (v4 << ((18446744073709551615) & 63))
        var v20 int64 = 9223372036854775807
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := v3[1]
            var v24 uint = 665
            v1 += v1
            _ = fi22
            _ = v23
            _ = v24
        }
        if (943 <= 2) {
            var v25 uint16 = 1
            _ = v25
        }
        switch v2 {
        case "hello":
            for fi26 := 0; fi26 < 3; fi26++ {
                v27 := GenS1{F0: (fi21 % (((fi26 >> ((2) & 63))) | 1)), F1: (v1 & genFn0(1)), F2: 32767}
                fmt.Println(fi26)
                v2 = ("a" + v5)
                _ = fi26
                _ = v27
            }
        case "\n\t":
            if false {
                switch fi21 {
                case -739:
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 uint = 1000
                        _ = len(v3)
                        switch fi21 {
                        case 0:
                            var v30 uint32 = (v4 * v4)
                            _ = v30
                        case 128:
                            if ((v29 - v29) <= v29) {
                                v31 := GenS3{F0: ("\n\t" + "\n\t"), F1: 0}
                                var v32 bool = false
                                _ = v31
                                _ = v32
                            }
                        case -192:
                            v1 -= int8(v4)
                        }
                        var v33 int32 = 786
                        _ = fi28
                        _ = v29
                        _ = v33
                    }
                case 128:
                    if (false || true) {
                        v34 := v3[v4]
                        v4 = (^(v34))
                        if (255 >= 0) {
                            switch fi21 {
                            case 333:
                                var v35 float64 = (1e10 * (1e10 - 3.14159))
                                _ = v35
                            case -103:
                                switch (fi21 % (((fi21 << ((42) & 63))) | 1)) {
                                case -2147483648:
                                    var v36 int16 = 42
                                    _ = v36
                                case 9223372036854775807:
                                    var v37 bool = (-2.25 <= -2.25)
                                    _ = v37
                                default:
                                    for fi38 := 0; fi38 < 12; fi38++ {
                                        fmt.Println(v34)
                                        v4 -= (v4 - (v4 - v34))
                                        var v39 uint8 = 0
                                        _ = fi38
                                        _ = v39
                                    }
                                }
                            case 0:
                                v2 = (v5 + v5)
                            }
                        }
                        _ = v34
                    }
                default:
                    v2 = v5
                }
                v40 := v3[(v4 % (((^(v4))) | 1))]
                _ = v40
            }
        case "":
            var v41 int64 = -9223372036854775808
            _ = v41
        default:
            v42 := []int{908, (fi21 + (^(fi21))), (fi21 % (((fi21 ^ (fi21 | 4294967295))) | 1)), (fi21 % (((fi21 * fi21)) | 1))}
            _ = v42
        }
        _ = fi21
    }
    for fi43 := 0; fi43 < 3; fi43++ {
        if false {
            v44 := []int{fi43}
            v45 := GenS1{F0: 32767, F1: (v1 << ((255) & 63)), F2: 32767}
            _ = v44
            _ = v45
        }
        var v46 int32 = 2147483647
        var v47 uint64 = 18446744073709551615
        _ = fi43
        _ = v46
        _ = v47
    }
    for fi48 := 0; fi48 < 10; fi48++ {
        v5 = "foo bar"
        for fi49 := 0; fi49 < 12; fi49++ {
            v2 = (((v2 + ("" + v5)) + v5) + v2)
            _ = fi49
        }
        _ = fi48
    }
    if false {
        var v50 int = 0
        for fi51 := 0; fi51 < 19; fi51++ {
            var v52 float32 = (1e10 * -2.25)
            v53 := GenS4{F0: 88, F1: 1000}
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = v50
    }
    if ((0 == 42) && true) {
        if false {
            var v54 float32 = 0.0
            var v55 uint16 = 255
            var v56 int16 = -32768
            _ = v54
            _ = v55
            _ = v56
        }
        v57 := []int{-423}
        _ = v57
    }
    if true {
        if true {
            var v58 uint32 = v4
            _ = v58
        }
    }
    v4 = v4
    v59 := []int{431, 2147483647, 32767}
    v5 = v5
    var v60 bool = false
    for fi61 := 0; fi61 < 14; fi61++ {
        if (((-(v1)) >= 100) && v60) {
            delete(v3, v4)
            var v62 float64 = 3.14159
            _ = v62
        } else {
            for fi63 := 0; fi63 < 3; fi63++ {
                var v64 uint32 = 4294967295
                if len(v59) >= 2 {
                    v59 = v59[1:len(v59)]
                }
                _ = fi63
                _ = v64
            }
            v3[v4] = 622
            delete(v3, v4)
        }
        _ = fi61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
}
