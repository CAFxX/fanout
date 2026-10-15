package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int8, p1 string) uint64 {
    v7 := GenS1{F0: ((-2.25 + 1e10) > 1.5)}
    v8 := GenS2{F0: 48}
    v9 := []int{32768, 132, -493, -1}
    fmt.Println(p1, p0)
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(p1, p0)
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int16 = 0
            v12 = -32768
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 575
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    var v4 uint8 = 102
    var v5 uint64 = 112
    var v6 string = ("hello" + "a")
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: true}
        v15 := v2[1]
        fmt.Println(v4, v6, v5)
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v6, v5, v5)
    var v16 uint16 = uint16(v4)
    v17 := GenS2{F0: v4}
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := []int{fi18, (fi18 >> ((259) & 63)), 292, (^(fi18))}
        if (!true) {
            if (false && (!false)) {
                for fi20 := 0; fi20 < 11; fi20++ {
                    switch ("hello" + v6) {
                    case "a":
                        var v21 float32 = float32(float64(((1.5 * 1e10) - 0.0)))
                        _ = v21
                    default:
                        v16 += 0
                    }
                    v4 += v4
                    var v22 int64 = -34
                    if false {
                        v22 = -9223372036854775808
                        var v23 uint8 = v4
                        var v24 float32 = 1.5
                        _ = v23
                        _ = v24
                    } else {
                        var v25 bool = true
                        _ = v25
                    }
                    _ = fi20
                    _ = v22
                }
            } else {
                if true {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        var v27 uint = 255
                        v16 = v16
                        if (897 != 2) {
                            var v28 int64 = 405
                            switch v6 {
                            case "a":
                                v29 := GenS2{F0: v4}
                                _ = v29
                            case "foo bar":
                                v27 += 2
                            case "hello":
                                var v30 bool = false
                                _ = v30
                            }
                            v28 = 1
                            v28 = 0
                            delete(v3, 566)
                            _ = v28
                        } else {
                            switch (v6 + v6) {
                            case "hello":
                                _ = len(v19)
                            }
                        }
                        _ = fi26
                        _ = v27
                    }
                }
                v3[fi18] = true
                var v31 int8 = 42
                _ = v31
            }
            var v32 uint16 = v16
            v6 = ((((v6 + "\n\t") + "\n\t") + "foo bar") + (v6 + (("hello" + v6) + ((v6 + v6) + "\n\t"))))
            _ = v32
        } else {
            switch (v6 + v6) {
            case "\n\t":
                for fi33 := 0; fi33 < 3; fi33++ {
                    var v34 uint = 4294967295
                    switch (-(fi33)) {
                    case 0:
                        fmt.Println(v4)
                        var v35 uint64 = (v5 * (v5 << (((^(v34))) & 63)))
                        if (!false) {
                            v34 += (^(v34))
                        }
                        _ = v35
                    case -785:
                        switch v6 {
                        case "\n\t":
                            var v36 int8 = 97
                            _ = v36
                        case "foo bar":
                            var v37 float32 = 3.14159
                            _ = v37
                        case "a":
                            v6 = (v6 + v6)
                        }
                    }
                    _ = fi33
                    _ = v34
                }
            default:
                v38 := GenS3{F0: v16}
                _ = v38
            }
            v39 := GenS4{F0: 65535, F1: (v4 << ((654) & 63))}
            v16 = 338
            _ = v39
        }
        _ = fi18
        _ = v19
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        var v41 int64 = 9223372036854775807
        _ = fi40
        _ = v41
    }
    if false {
        var v42 uint16 = 0
        v6 = v6
        delete(v3, 4294967295)
        _ = v42
    }
    var v43 uint16 = (^(v16))
    for fi44 := 0; fi44 < 11; fi44++ {
        v5 += v5
        _ = fi44
    }
    v45 := GenS5{F0: 0}
    v1[(v6 + v6)] = false
    v6 = "foo bar"
    v5 = v5
    if ((!("\n\t" < (v6 + v6))) && (-1 > 183)) {
        fmt.Println(v16)
    }
    var v46 bool = ((v5 - (^(v5))) > genFn0(0, ((v6 + (("\n\t" + v6) + v6)) + ((v6 + (v6 + (v6 + "foo bar"))) + "hello"))))
    var v47 float64 = 0.0
    var v48 int16 = -591
    v49 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
