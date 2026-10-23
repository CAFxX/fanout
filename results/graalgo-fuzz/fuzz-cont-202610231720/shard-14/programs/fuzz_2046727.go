package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint64; F2 int8 }
func genFn0(p0 int8) int16 {
    v4 := GenS1{F0: 68, F1: 0}
    for fi5 := 0; fi5 < 5; fi5++ {
        p0 = 127
        v6 := []int{}
        switch fi5 {
        case 128:
            var v7 uint = 18446744073709551615
            _ = v7
        default:
            var v8 float64 = 0.1
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 33
}
func genFn1(p0 int32) uint8 {
    if (!(!(true || (!(4294967295 != 65535))))) {
        switch -890 {
        case -2147483648:
            p0 += p0
        case -836:
            switch ("\n\t" + "hello") {
            case "\n\t":
                for fi9 := 0; fi9 < 9; fi9++ {
                    switch "foo bar" {
                    case "foo bar":
                        v10 := []int{(fi9 | 280)}
                        _ = v10
                    case "":
                        for fi11 := 0; fi11 < 9; fi11++ {
                            switch -2147483648 {
                            case -115:
                                v12 := []int{fi9, (fi11 << ((4294967295) & 63)), fi11}
                                _ = v12
                            default:
                                if false {
                                    if (true && (!false)) {
                                        switch "a" {
                                        case "":
                                            v13 := []int{32768, 4294967295, fi9, (fi11 / (((fi11 << ((255) & 63))) | 1))}
                                            _ = v13
                                        case "\n\t":
                                            v14 := GenS1{F0: 128, F1: 0}
                                            _ = v14
                                        case "hello":
                                            if false {
                                                p0 = p0
                                                var v15 int = (fi11 ^ int(p0))
                                                _ = v15
                                            } else {
                                                v16 := GenS1{F0: 181, F1: (p0 << ((255) & 63))}
                                                var v17 int32 = -31
                                                var v18 uint = 18446744073709551615
                                                v17 += v17
                                                v19 := []int{(^(fi9)), fi11}
                                                _ = v16
                                                _ = v17
                                                _ = v18
                                                _ = v19
                                            }
                                        }
                                    } else {
                                        var v20 uint16 = 32767
                                        if (893 <= 864) {
                                            if (2 <= 1) {
                                                v20 = v20
                                            } else {
                                                fmt.Println(p0, v20, fi11)
                                                p0 += -1
                                                var v21 bool = ((p0 / (((^(p0))) | 1)) < 1)
                                                v22 := []int{}
                                                var v23 int = (-(fi9))
                                                _ = v21
                                                _ = v22
                                                _ = v23
                                            }
                                            v20 = 32767
                                        } else {
                                            fmt.Println(v20, fi11, fi9)
                                        }
                                        _ = v20
                                    }
                                } else {
                                    fmt.Println(fi9, fi11)
                                }
                            }
                            _ = fi11
                        }
                    default:
                        p0 = 2147483647
                    }
                    _ = fi9
                }
            }
        default:
            switch "hello" {
            case "a":
                if true {
                    var v24 int16 = genFn0(-1)
                    _ = v24
                }
            case "":
                if (24 >= -32768) {
                    var v25 uint = 2
                    _ = v25
                } else {
                    var v26 uint = 1
                    v27 := []int{-1, 2147483647, 127, 32768}
                    _ = v26
                    _ = v27
                }
            case "hello":
                switch "\n\t" {
                case "hello":
                    switch "\n\t" {
                    case "":
                        v28 := GenS2{F0: -32768, F1: 4294967295}
                        _ = v28
                    }
                case "a":
                    p0 = (^(p0))
                case "":
                    v29 := GenS3{F0: 19}
                    _ = v29
                }
            }
        }
    }
    v30 := GenS1{F0: 128, F1: -794}
    v31 := []int{128, 986}
    p0 -= (p0 << ((65535) & 63))
    _ = p0
    _ = v30
    _ = v31
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -9223372036854775808
    var v3 bool = (!true)
    v32 := []int{127, 42, 4294967295}
    v33 := GenS4{F0: ("\n\t" + "\n\t")}
    var v34 uint = 4294967295
    if len(v32) > 0 {
        v32[0] = -1
    }
    var v35 string = "hello"
    _ = len(v1)
    v34 -= (v34 >> ((v34) & 63))
    v36 := GenS5{F0: 2, F1: 297, F2: -1}
    for fi37 := 0; fi37 < 10; fi37++ {
        for fi38 := 0; fi38 < 1; fi38++ {
            var v39 string = v35
            _ = fi38
            _ = v39
        }
        var v40 uint64 = 4294967295
        _ = fi37
        _ = v40
    }
    v34 = (v34 & 42)
    var v41 uint = 2
    fmt.Println(v2)
    var v42 uint64 = 18446744073709551615
    var v43 int16 = 50
    v44 := GenS4{F0: ((("" + "\n\t") + v35) + "foo bar")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
