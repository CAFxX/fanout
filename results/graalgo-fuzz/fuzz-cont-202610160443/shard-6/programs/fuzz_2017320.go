package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int; F2 int16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int8, p1 int16, p2 uint32) uint32 {
    v5 := GenS1{F0: -9223372036854775808}
    v6 := GenS1{F0: -1}
    var v7 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p2
}

func main() {
    var v1 float32 = 0.0
    var v2 float64 = 1.5
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v8 uint = 2
    if false {
        v9 := GenS1{F0: -9223372036854775808}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            _ = len(v3)
            _ = fi10
        }
        v1 -= v1
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: -656}
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 string = "hello"
            fmt.Println(v1, fi13, v8)
            v1 += (v1 * (1e10 * (3.14159 * v1)))
            if (true || false) {
                var v15 string = (v14 + v14)
                fmt.Println(v2, fi13)
                _ = v15
            } else {
                var v16 uint = 42
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v17 int64 = -1
    if true {
        var v18 uint = v8
        var v19 int32 = 42
        _ = v18
        _ = v19
    }
    var v20 bool = (2147483648 != genFn0(100, 0, 428))
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 string = "\n\t"
        var v23 int8 = 0
        _ = fi21
        _ = v22
        _ = v23
    }
    if (65535 > 158) {
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := GenS2{F0: 32767, F1: -1, F2: -268}
            v8 = v8
            _ = fi24
            _ = v25
        }
        if v20 {
            if ((^(v8)) != 1) {
                var v26 uint16 = 203
                _ = v26
            }
        } else {
            var v27 uint32 = 42
            var v28 int64 = (v17 % ((v17) | 1))
            if (42 > -2147483648) {
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 uint64 = 1
                    _ = len(v3)
                    var v31 uint = 4294967295
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 7; fi32++ {
                    v33 := GenS1{F0: (^(v28))}
                    v17 = (v28 % ((-1) | 1))
                    if false {
                        switch fi32 {
                        case 676:
                            for fi34 := 0; fi34 < 11; fi34++ {
                                if v20 {
                                    var v35 float32 = v1
                                    fmt.Println(v20, v2)
                                    _ = v35
                                } else {
                                    switch (fi34 - fi32) {
                                    case 4294967295:
                                        var v36 int8 = 16
                                        _ = v36
                                    case 0:
                                        v37 := GenS1{F0: (v28 >> ((1) & 63))}
                                        _ = v37
                                    }
                                    v38 := GenS3{F0: -2147483648, F1: 42}
                                    _ = len(v3)
                                    v39 := GenS1{F0: v28}
                                    _ = v38
                                    _ = v39
                                }
                                v17 += v28
                                _ = fi34
                            }
                        case 2147483647:
                            v40 := GenS1{F0: v28}
                            _ = v40
                        case -664:
                            if (v20 || true) {
                                switch "foo bar" {
                                case "":
                                    var v41 uint8 = 128
                                    _ = v41
                                case "a":
                                    v28 -= (v17 + v28)
                                case "hello":
                                    v42 := v4["foo bar"]
                                    _ = v42
                                default:
                                    var v43 float64 = v2
                                    _ = v43
                                }
                                var v44 int8 = 0
                                _ = v44
                            }
                        }
                        var v45 int = fi32
                        if false {
                            for fi46 := 0; fi46 < 12; fi46++ {
                                var v47 bool = v20
                                _ = fi46
                                _ = v47
                            }
                            for fi48 := 0; fi48 < 4; fi48++ {
                                v20 = (false || false)
                                v2 -= (1.5 / v2)
                                _ = fi48
                            }
                            var v49 uint = v8
                            _ = v49
                        }
                        _ = v45
                    } else {
                        if v20 {
                            if v20 {
                                switch (fi32 / ((fi32) | 1)) {
                                case 0:
                                    v2 = v2
                                case -1:
                                    var v50 bool = true
                                    _ = v50
                                case 32768:
                                    v27 = v27
                                default:
                                    if (-1 < 42) {
                                        for fi51 := 0; fi51 < 4; fi51++ {
                                            var v52 string = "hello"
                                            _ = fi51
                                            _ = v52
                                        }
                                        var v53 int8 = 42
                                        v54 := GenS4{F0: v53}
                                        _ = v53
                                        _ = v54
                                    } else {
                                        v2 = v2
                                        var v55 uint = 464
                                        _ = v55
                                    }
                                }
                                v56 := []int{}
                                v57 := GenS3{F0: 605, F1: 127}
                                _ = v56
                                _ = v57
                            }
                            v58 := GenS4{F0: -119}
                            fmt.Println(v27, v27, v17)
                            _ = v58
                        } else {
                            v27 -= v27
                            if v20 {
                                v59 := GenS3{F0: 801, F1: 175}
                                var v60 uint64 = 18446744073709551615
                                _ = v59
                                _ = v60
                            } else {
                                var v61 string = "\n\t"
                                _ = v61
                            }
                        }
                        if true {
                            v27 += (v27 % ((4294967295) | 1))
                            v62 := GenS2{F0: -1, F1: (^(fi32)), F2: 1000}
                            _ = v62
                        }
                        var v63 uint64 = uint64(v8)
                        _ = v63
                    }
                    _ = fi32
                    _ = v33
                }
                if v20 {
                    switch -2147483648 {
                    case 2147483647:
                        var v64 int8 = 100
                        _ = v64
                    }
                }
            } else {
                var v65 int32 = 42
                v8 -= v8
                _ = v65
            }
            _ = v27
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
