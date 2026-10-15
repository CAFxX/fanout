package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 bool }
type GenS2 struct { F0 int; F1 bool; F2 int64 }
type GenS3 struct { F0 uint32; F1 bool; F2 int16 }
func genFn0(p0 uint32, p1 string, p2 int64) uint64 {
    v4 := GenS1{F0: -706, F1: -2147483648, F2: false}
    if (4294967295 >= 255) {
        switch p1 {
        case "\n\t":
            if false {
                p1 = "\n\t"
            }
        case "hello":
            p2 += (p2 ^ p2)
        case "a":
            for fi5 := 0; fi5 < 15; fi5++ {
                v6 := []int{}
                _ = fi5
                _ = v6
            }
        }
        v7 := []int{-243, -194, 1}
        _ = len(v7)
        _ = v7
    }
    p2 = (p2 % (((^(p2))) | 1))
    v8 := []int{int(p2), 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return 1
}
func genFn1(p0 int16) int8 {
    var v9 bool = (true || false)
    for fi10 := 0; fi10 < 2; fi10++ {
        v9 = ((^(p0)) > (p0 ^ p0))
        if v9 {
            if v9 {
                v9 = ((42 != 1) && (-585 < 42))
                switch ("" + ("hello" + "a")) {
                case "foo bar":
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v9 = v9
                        if (!false) {
                            v9 = true
                        }
                        _ = fi11
                    }
                case "hello":
                    v12 := []int{9223372036854775807, (fi10 / ((fi10) | 1)), fi10, fi10}
                    _ = v12
                case "a":
                    var v13 uint16 = 901
                    _ = v13
                }
            } else {
                var v14 int32 = 69
                v15 := []int{fi10, int(v14), 42, (^(fi10))}
                _ = v14
                _ = v15
            }
            var v16 int = -1
            switch "\n\t" {
            case "":
                var v17 uint16 = 964
                _ = v17
            case "foo bar":
                var v18 float32 = 1.5
                _ = v18
            case "a":
                var v19 int64 = -26
                _ = v19
            }
            _ = v16
        } else {
            if (!false) {
                for fi20 := 0; fi20 < 10; fi20++ {
                    v21 := []int{-701, (fi10 % ((fi20) | 1))}
                    var v22 uint16 = 32767
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v23 := GenS1{F0: -32768, F1: int32(p0), F2: (v9 || (!(65535 <= 1)))}
                for fi24 := 0; fi24 < 12; fi24++ {
                    v25 := GenS2{F0: fi10, F1: false, F2: -1}
                    _ = fi24
                    _ = v25
                }
                _ = v23
            } else {
                v26 := GenS1{F0: (^(p0)), F1: -370, F2: true}
                _ = v26
            }
        }
        v27 := []int{}
        _ = fi10
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 uint16 = 943
        p0 -= p0
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        v31 := []int{128, 2147483647, (fi30 | fi30), (fi30 & fi30)}
        _ = fi30
        _ = v31
    }
    v32 := GenS2{F0: 2147483647, F1: (false || (v9 || v9)), F2: 42}
    _ = p0
    _ = v32
    _ = v9
    return 127
}

func main() {
    v1 := []uint32{2147483647}
    var v2 int16 = -32768
    var v3 uint8 = 1
    v2 = v2
    var v33 uint32 = 65535
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 20; fi35++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v33 += (^(v33))
            var v36 int8 = 0
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 20; fi37++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v2, fi37)
            fmt.Println(fi37, v3, v2)
            _ = fi37
        }
        var v38 bool = true
        _ = fi34
        _ = v38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 int16 = v2
        _ = fi39
        _ = v40
    }
    switch 128 {
    case -254:
        var v41 uint = 18446744073709551615
        _ = v41
    case -1:
        fmt.Println(v3)
    case 32767:
        var v42 uint32 = (v33 - v33)
        _ = v42
    default:
        _ = len(v1)
    }
    if true {
        v43 := GenS2{F0: -972, F1: (!(649 > 1)), F2: 0}
        var v44 int16 = v2
        if false {
            fmt.Println(v33, v44, v44)
            v3 += (v3 - (v3 & (v3 & 0)))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            for fi45 := 0; fi45 < 12; fi45++ {
                v2 -= 1
                _ = fi45
            }
            if (2 >= 65535) {
                v46 := GenS3{F0: (v33 | 559), F1: false, F2: (v2 >> ((2) & 63))}
                _ = v46
            } else {
                v44 = -32768
            }
        }
        _ = v43
        _ = v44
    }
    switch "hello" {
    case "\n\t":
        if true {
            var v47 bool = true
            v48 := GenS3{F0: v33, F1: false, F2: -1}
            _ = v47
            _ = v48
        } else {
            fmt.Println(v33, v33, v2)
            if false {
                var v49 int16 = 952
                v33 += v33
                _ = v49
            }
        }
    }
    var v50 int8 = -128
    var v51 float32 = (float32(0.1) * 1e10)
    v3 = uint8(v33)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v50
    _ = v51
}
