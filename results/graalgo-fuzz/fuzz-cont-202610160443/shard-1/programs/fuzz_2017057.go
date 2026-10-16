package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 bool }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint; F1 bool }
func genFn0(p0 uint32, p1 int16) string {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 bool = false
        v7 = v7
        _ = fi6
        _ = v7
    }
    var v8 int64 = 509
    v9 := GenS1{F0: 255, F1: v8, F2: (27 >= 127)}
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: 1, F1: (v8 * (v8 / (((v8 % ((v8) | 1))) | 1))), F2: true}
        _ = fi10
        _ = v11
    }
    v12 := []int{725, 4294967295, 755, 127}
    _ = p0
    _ = p1
    _ = v12
    _ = v8
    _ = v9
    return ("a" + "")
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]string)
    var v3 uint8 = 1
    v4 := make(map[int]int)
    var v5 int = 127
    v1 = v1
    switch v5 {
    case 127:
        switch v5 {
        case 128:
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 string = v1
                _ = fi13
                _ = v14
            }
        case -9223372036854775808:
            for fi15 := 0; fi15 < 6; fi15++ {
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 uint8 = (v3 & (v3 + 194))
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 float64 = (1.5 - 0.1)
                    var v20 uint8 = (v3 % ((0) | 1))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                switch "\n\t" {
                case "foo bar":
                    v5 -= v5
                case "hello":
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v1 = ("foo bar" + (v1 + v1))
                        _ = fi21
                    }
                case "":
                    var v22 string = "a"
                    _ = v22
                default:
                    var v23 uint16 = 0
                    _ = v23
                }
                _ = fi15
            }
        case 237:
            fmt.Println(v5, v5)
        default:
            for fi24 := 0; fi24 < 1; fi24++ {
                fmt.Println(v5, v1)
                _ = fi24
            }
        }
    case -142:
        fmt.Println(v1, v1, v1)
    }
    _ = len(v2)
    for fi25 := 0; fi25 < 11; fi25++ {
        if (!(1 <= 100)) {
            v26 := []int{(fi25 * (v5 ^ (fi25 ^ fi25))), v5, (fi25 + v5)}
            for fi27 := 0; fi27 < 17; fi27++ {
                if (-1 < 1) {
                    v3 = 24
                    if true {
                        v4[(v5 & fi25)] = fi27
                        if (!((true || ((127 != -1) && (18446744073709551615 >= 255))) && (v1 == v1))) {
                            if ((false && false) || (65535 >= 255)) {
                                var v28 uint = 872
                                var v29 bool = (!(!true))
                                v30 := GenS2{F0: 0, F1: 42}
                                _ = v28
                                _ = v29
                                _ = v30
                            } else {
                                _ = len(v26)
                                v5 -= 32767
                                var v31 int16 = -32768
                                _ = v31
                            }
                            v3 = 42
                            fmt.Println(fi27, fi27)
                            switch v1 {
                            case "":
                                delete(v4, fi27)
                            case "a":
                                var v32 string = v1
                                _ = v32
                            default:
                                switch fi25 {
                                case 32767:
                                    if (false || false) {
                                        fmt.Println(fi27, v3)
                                        v33 := GenS1{F0: 255, F1: -592, F2: true}
                                        _ = v33
                                    }
                                }
                            }
                        }
                    } else {
                        if (!false) {
                            v3 = (v3 / (((v3 ^ 0)) | 1))
                            _ = len(v2)
                            v34 := GenS3{F0: -1, F1: 941, F2: 2}
                            _ = v34
                        }
                        v35 := GenS4{F0: 1, F1: false}
                        var v36 uint64 = 242
                        _ = v35
                        _ = v36
                    }
                } else {
                    if false {
                        v3 -= (^(v3))
                    } else {
                        var v37 uint8 = (^(v3))
                        if true {
                            fmt.Println(v37, v1, fi27)
                        }
                        _ = len(v26)
                        _ = v37
                    }
                }
                if (false && false) {
                    var v38 uint = 255
                    v39 := GenS2{F0: -1, F1: 65535}
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    _ = v38
                    _ = v39
                } else {
                    var v40 uint64 = 512
                    v3 -= v3
                    v1 = "a"
                    _ = v40
                }
                if false {
                    v41 := GenS1{F0: 65535, F1: 1, F2: (v3 < v3)}
                    var v42 string = v1
                    if false {
                        var v43 int64 = 353
                        _ = v43
                    }
                    v26 = append(v26, fi27)
                    _ = v41
                    _ = v42
                } else {
                    var v44 float32 = (1.5 + -2.25)
                    _ = v44
                }
                _ = fi27
            }
            _ = v26
        }
        var v45 uint8 = (^(v3))
        _ = fi25
        _ = v45
    }
    v46 := GenS1{F0: 0, F1: 1, F2: (!true)}
    delete(v4, v5)
    v5 = 1
    var v47 int64 = 776
    v48 := []int{(v5 + v5)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
