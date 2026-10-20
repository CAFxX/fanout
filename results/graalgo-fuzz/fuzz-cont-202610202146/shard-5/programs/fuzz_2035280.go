package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS4 struct { F0 int8; F1 int; F2 string }

func main() {
    var v1 int64 = -9
    v2 := []int{4294967295}
    var v3 uint64 = 4294967295
    if true {
        var v4 int32 = -1
        v1 = v1
        _ = v4
    } else {
        var v5 uint32 = 0
        v2 = append(v2, 2147483647)
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 float32 = (0.1 / 1.0)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v8 int16 = -32768
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    }
    switch -1 {
    case 128:
        v1 = (v1 * -429)
    case -596:
        v3 += v3
    case 9223372036854775807:
        v3 = v3
    default:
        var v9 uint8 = 127
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    v1 -= v1
    var v10 int32 = 0
    fmt.Println(v10, v10)
    if false {
        switch int(v10) {
        case 2147483647:
            for fi11 := 0; fi11 < 5; fi11++ {
                if true {
                    fmt.Println(v1, fi11, v10)
                }
                _ = fi11
            }
        case 901:
            var v12 uint = 4294967295
            _ = v12
        }
        v3 -= 18446744073709551615
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        if true {
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int = -9223372036854775808
                _ = fi14
                _ = v15
            }
        } else {
            var v16 int64 = (-(v1))
            var v17 int8 = 127
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := GenS1{F0: 4294967295, F1: 65535}
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        }
        if (1.5 >= float64(0.0)) {
            v20 := GenS1{F0: 2, F1: 42}
            switch -2147483648 {
            case 194:
                for fi21 := 0; fi21 < 13; fi21++ {
                    v2 = append(v2, (fi21 - (fi21 + fi21)))
                    _ = fi21
                }
            case 42:
                if (true || (625 < 0)) {
                    v22 := GenS1{F0: 655, F1: 72}
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v3 -= 761
                        var v24 uint32 = 4294967295
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                }
            default:
                if (!true) {
                    _ = len(v2)
                    v10 -= v10
                } else {
                    if false {
                        if true {
                            fmt.Println(v10)
                            v25 := GenS2{F0: v1, F1: 140}
                            if false {
                                var v26 uint64 = v3
                                _ = v26
                            } else {
                                v3 += (v3 * (v3 + (v3 << ((500) & 63))))
                                var v27 uint16 = 0
                                var v28 bool = true
                                _ = v27
                                _ = v28
                            }
                            _ = v25
                        } else {
                            if len(v2) > 0 {
                                v2[0] = (^(fi13))
                            }
                            v10 = (v10 * (v10 * -2147483648))
                        }
                        switch "foo bar" {
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "foo bar":
                            v29 := GenS2{F0: -1, F1: (v10 << ((2) & 63))}
                            _ = v29
                        case "":
                            v30 := GenS1{F0: 786, F1: 146}
                            _ = v30
                        default:
                            var v31 int8 = 127
                            _ = v31
                        }
                    } else {
                        fmt.Println(v3, v10, v3)
                        v1 = int64(fi13)
                        if (!(0.1 <= (0.1 * float32(1e10)))) {
                            if true {
                                var v32 uint32 = 4294967295
                                v33 := GenS1{F0: v32, F1: 231}
                                v32 -= v32
                                _ = v32
                                _ = v33
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                    }
                }
            }
            fmt.Println(v10, fi13, fi13)
            _ = v20
        }
        _ = fi13
    }
    var v34 int16 = 1000
    if false {
        var v35 int64 = v1
        for fi36 := 0; fi36 < 3; fi36++ {
            for fi37 := 0; fi37 < 4; fi37++ {
                v10 -= v10
                fmt.Println(v3, fi36)
                var v38 uint32 = 2
                _ = fi37
                _ = v38
            }
            _ = fi36
        }
        _ = v35
    }
    var v39 int8 = -1
    fmt.Println(v3)
    v40 := GenS3{F0: 198, F1: v34, F2: 2}
    v41 := GenS4{F0: (v39 * v39), F1: 1, F2: ("\n\t" + (("\n\t" + ("" + "")) + "hello"))}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v40
    _ = v41
}
