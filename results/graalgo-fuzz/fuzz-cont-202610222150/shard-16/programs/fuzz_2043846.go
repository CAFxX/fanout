package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 int16 = -32768
    if false {
        v4 := GenS1{F0: v3, F1: -9223372036854775808}
        _ = v4
    }
    if true {
        _ = len(v2)
        var v5 int8 = -128
        _ = v5
    } else {
        switch "hello" {
        case "\n\t":
            for fi6 := 0; fi6 < 2; fi6++ {
                fmt.Println(v3)
                for fi7 := 0; fi7 < 6; fi7++ {
                    v8 := GenS1{F0: (v3 * (v3 + v3)), F1: -1}
                    v3 = (v3 & (v3 << ((509) & 63)))
                    _ = len(v1)
                    _ = fi7
                    _ = v8
                }
                _ = fi6
            }
        case "foo bar":
            fmt.Println(v3)
        default:
            fmt.Println(v3)
        }
        var v9 uint16 = 1
        _ = v9
    }
    switch -2147483648 {
    case 32768:
        v10 := GenS1{F0: (-(v3)), F1: 0}
        _ = v10
    case 2147483647:
        if false {
            v3 = (v3 / ((v3) | 1))
            for fi11 := 0; fi11 < 12; fi11++ {
                if (0.1 <= 1.5) {
                    if (42 <= 2) {
                        v3 -= 42
                        var v12 int32 = 42
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v14 := []int{fi13, fi13, fi13, (fi13 * (fi11 * 4294967295))}
                        if false {
                            v3 += v3
                        }
                        _ = fi13
                        _ = v14
                    }
                } else {
                    if true {
                        v3 += (v3 + (v3 & v3))
                        fmt.Println(fi11, fi11)
                        if false {
                            v15 := []int{}
                            _ = v15
                        } else {
                            v3 -= v3
                            if false {
                                v16 := GenS1{F0: (v3 - (-(v3))), F1: 601}
                                _ = v16
                            } else {
                                for fi17 := 0; fi17 < 8; fi17++ {
                                    v18 := GenS1{F0: -80, F1: 0}
                                    _ = fi17
                                    _ = v18
                                }
                                fmt.Println(v3, fi11)
                            }
                            v3 += v3
                        }
                    } else {
                        v3 -= v3
                        v19 := GenS2{F0: 0, F1: true}
                        fmt.Println(fi11, fi11)
                        _ = v19
                    }
                }
                _ = fi11
            }
            switch "a" {
            case "hello":
                v3 = -998
            case "foo bar":
                v20 := GenS1{F0: -538, F1: -1}
                _ = v20
            }
        } else {
            var v21 float32 = 1e10
            for fi22 := 0; fi22 < 2; fi22++ {
                if false {
                    var v23 int32 = -2147483648
                    var v24 uint64 = 185
                    var v25 uint16 = 253
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS1{F0: v3, F1: 267}
                    _ = v26
                }
                _ = fi22
            }
            _ = v21
        }
    }
    var v27 int32 = -1
    for fi28 := 0; fi28 < 8; fi28++ {
        fmt.Println(v27, v3, v3)
        _ = fi28
    }
    switch "\n\t" {
    case "a":
        if true {
            var v29 uint32 = 0
            _ = v29
        } else {
            v30 := []int{}
            v27 = (v27 | -724)
            var v31 int64 = 1
            _ = v30
            _ = v31
        }
    default:
        var v32 uint16 = 2
        _ = v32
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 float32 = float32((1.5 - 0.0))
        v35 := []int{(fi33 + fi33), int(v27), (-(fi33)), -506}
        if true {
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 uint32 = 2147483647
                _ = fi36
                _ = v37
            }
            switch "foo bar" {
            case "a":
                v27 -= (v27 | 1)
            }
        } else {
            var v38 bool = true
            v39 := GenS1{F0: (-(v3)), F1: -21}
            _ = v38
            _ = v39
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    v2["\n\t"] = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
}
