package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int8; F1 int; F2 int8 }
func genFn0(p0 uint8, p1 uint8) int32 {
    switch ("" + "\n\t") {
    case "\n\t":
        if (false || false) {
            var v5 uint = 4294967295
            if false {
                v6 := []int{127, -1, -2147483648}
                var v7 int32 = -1
                p0 = 1
                _ = v6
                _ = v7
            } else {
                p1 = (p1 | (p0 & p0))
            }
            _ = v5
        } else {
            var v8 uint16 = 42
            if (true && true) {
                p1 = p1
            }
            _ = v8
        }
    default:
        var v9 uint8 = 128
        _ = v9
    }
    p0 += 128
    var v10 uint16 = 32767
    p1 += 170
    var v11 int8 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 202
}

func main() {
    var v1 int16 = -897
    var v2 int32 = 42
    v3 := make(map[string]uint32)
    var v4 bool = (!false)
    var v12 bool = v4
    switch "hello" {
    case "\n\t":
        var v13 int32 = (v2 | v2)
        _ = v13
    case "hello":
        switch -136 {
        case -1:
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 uint32 = 1
                _ = fi14
                _ = v15
            }
        case 0:
            v1 = v1
        default:
            v2 = 42
        }
    case "a":
        var v16 uint32 = 65535
        _ = v16
    default:
        for fi17 := 0; fi17 < 15; fi17++ {
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := []int{128, fi17, (fi17 & 127)}
                _ = fi18
                _ = v19
            }
            fmt.Println(fi17, v4)
            var v20 int64 = 766
            _ = fi17
            _ = v20
        }
    }
    v4 = v4
    var v21 bool = (309 < 42)
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v1, v2, fi22)
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint32 = 103
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    var v25 float64 = 0.1
    v25 = v25
    for fi26 := 0; fi26 < 12; fi26++ {
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := []int{}
            if (!v4) {
                var v29 float32 = 0.1
                _ = v29
            } else {
                v12 = v21
                switch fi26 {
                case -430:
                    v30 := GenS1{F0: 1, F1: v2, F2: v1}
                    _ = v30
                case -2147483648:
                    v31 := GenS2{F0: "", F1: 684}
                    _ = v31
                case 9223372036854775807:
                    v32 := GenS3{F0: 32767, F1: 42}
                    _ = v32
                default:
                    var v33 int64 = 9223372036854775807
                    _ = v33
                }
                switch fi26 {
                case 32767:
                    if v21 {
                        if len(v28) > 0 {
                            v28[0] = (-(fi26))
                        }
                    } else {
                        var v34 int = fi27
                        _ = v34
                    }
                    v35 := GenS4{F0: -32, F1: (-(v2))}
                    _ = v35
                }
                var v36 uint = 65535
                var v37 uint8 = 25
                v38 := GenS4{F0: 43, F1: (v2 | (v2 - 0))}
                var v39 float32 = 1e10
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                var v40 string = "a"
                var v41 int = int(v2)
                switch (fi27 - (fi26 ^ (fi26 << ((v36) & 63)))) {
                case 2147483647:
                    v42 := GenS3{F0: 436, F1: v36}
                    _ = v42
                case 42:
                    var v43 uint8 = v37
                    _ = v43
                case 4294967295:
                    if v4 {
                        fmt.Println(v36, v25, v39)
                    } else {
                        fmt.Println(v4)
                    }
                default:
                    if true {
                        v44 := GenS5{F0: 127, F1: v41, F2: int8(v37)}
                        var v45 int16 = (v1 - -644)
                        _ = v44
                        _ = v45
                    } else {
                        var v46 uint32 = 609
                        _ = v46
                    }
                }
                _ = v36
                _ = v37
                _ = v38
                _ = v39
                _ = v40
                _ = v41
            }
            v2 = v2
            _ = fi27
            _ = v28
        }
        var v47 int64 = 0
        v12 = (!v4)
        _ = fi26
        _ = v47
    }
    var v48 uint8 = 2
    v1 += v1
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v48
}
