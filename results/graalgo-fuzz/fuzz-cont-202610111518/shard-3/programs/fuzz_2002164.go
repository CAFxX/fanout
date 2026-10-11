package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int32; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 uint32 }
func genFn0(p0 string, p1 uint, p2 uint16) int64 {
    p1 += p1
    v7 := GenS1{F0: true}
    var v8 uint16 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 100
}
func genFn1() int64 {
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 string = ""
        v11 := []int{}
        v10 = v10
        _ = fi9
        _ = v10
        _ = v11
    }
    if true {
        if false {
            var v12 string = (("\n\t" + "foo bar") + ("" + "foo bar"))
            var v13 string = v12
            v14 := GenS2{F0: -128}
            _ = v12
            _ = v13
            _ = v14
        } else {
            var v15 bool = (true && (("\n\t" + "\n\t") == ""))
            _ = v15
        }
        v16 := []int{}
        _ = v16
    }
    return 42
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []int{(v1 + v1), (v1 / ((v1) | 1))}
    var v3 uint32 = 42
    var v4 uint16 = 2
    var v5 uint16 = (^(v4))
    v6 := make(map[uint32]uint32)
    var v17 float32 = 1e10
    v18 := GenS2{F0: -128}
    switch (("foo bar" + "hello") + "") {
    case "foo bar":
        for fi19 := 0; fi19 < 14; fi19++ {
            if (0.1 >= float64(v17)) {
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 int8 = 1
                    var v22 uint64 = 1
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            v6[v3] = 2
            _ = fi19
        }
    case "\n\t":
        var v23 float64 = float64(-2.25)
        _ = v23
    default:
        fmt.Println(v1)
    }
    fmt.Println(v4)
    switch (^(v1)) {
    case 0:
        var v24 bool = false
        _ = v24
    case 32768:
        v25 := GenS3{F0: v5}
        _ = v25
    case -2147483648:
        for fi26 := 0; fi26 < 16; fi26++ {
            fmt.Println(v3)
            var v27 uint64 = 42
            _ = fi26
            _ = v27
        }
    default:
        var v28 uint8 = 1
        _ = v28
    }
    v29 := GenS4{F0: 148, F1: 478, F2: 761}
    for fi30 := 0; fi30 < 17; fi30++ {
        if false {
            v4 = v4
            switch (v1 >> ((213) & 63)) {
            case 4294967295:
                v4 = 662
            case 9223372036854775807:
                v31 := GenS5{F0: (true && (128 > 1))}
                _ = v31
            }
        }
        _ = fi30
    }
    var v32 string = ("hello" + "foo bar")
    switch v32 {
    case "a":
        for fi33 := 0; fi33 < 10; fi33++ {
            for fi34 := 0; fi34 < 11; fi34++ {
                v35 := GenS4{F0: v5, F1: -1, F2: -444}
                v36 := GenS6{F0: true}
                v32 = v32
                _ = fi34
                _ = v35
                _ = v36
            }
            if false {
                v2 = append(v2, 754)
                for fi37 := 0; fi37 < 14; fi37++ {
                    var v38 uint = 1
                    if true {
                        v17 = v17
                        var v39 string = "a"
                        var v40 uint64 = 132
                        _ = v39
                        _ = v40
                    }
                    if true {
                        var v41 int64 = 1
                        var v42 uint64 = 2
                        fmt.Println(v5, v17)
                        _ = v41
                        _ = v42
                    }
                    _ = fi37
                    _ = v38
                }
            }
            _ = fi33
        }
    default:
        v1 += v1
    }
    switch (v1 % (((-(v1))) | 1)) {
    case 42:
        v5 += v5
    default:
        v43 := GenS7{F0: (v1 % ((v1) | 1)), F1: v3}
        _ = v43
    }
    fmt.Println(v17, v17, v17)
    for fi44 := 0; fi44 < 9; fi44++ {
        _ = len(v6)
        v17 = 1e10
        var v45 uint = 42
        _ = fi44
        _ = v45
    }
    fmt.Println(v3, v1, v3)
    var v46 uint64 = 18446744073709551615
    if (false || true) {
        var v47 bool = (false && (-32768 != -32768))
        fmt.Println(v1)
        v48 := v6[32]
        _ = v47
        _ = v48
    }
    v3 = v3
    v49 := GenS1{F0: (!(false && false))}
    v50 := v6[v3]
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
