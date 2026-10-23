package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 uint }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint64, p1 int8, p2 int8) uint16 {
    v4 := GenS1{F0: 0}
    var v5 float64 = -2.25
    if false {
        var v6 uint8 = 0
        v7 := GenS1{F0: 276}
        _ = v6
        _ = v7
    } else {
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 598
}

func main() {
    var v1 uint8 = 0
    v2 := []uint32{2147483648, 0, 2, 239}
    var v3 uint64 = 18446744073709551615
    if true {
        var v8 string = ""
        _ = len(v2)
        _ = v8
    }
    _ = len(v2)
    fmt.Println(v1, v3)
    for fi9 := 0; fi9 < 8; fi9++ {
        if (false && (42 > 815)) {
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 bool = false
                _ = fi10
                _ = v11
            }
            v3 = 4294967295
        }
        _ = fi9
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3)
    var v12 string = "\n\t"
    for fi13 := 0; fi13 < 19; fi13++ {
        if false {
            if false {
                if (15 == 683) {
                    if (2147483647 > 4294967295) {
                        var v14 string = v12
                        _ = v14
                    } else {
                        var v15 float64 = 0.1
                        switch int(v1) {
                        case 32767:
                            if false {
                                v12 = ((v12 + "\n\t") + v12)
                            }
                        }
                        _ = v15
                    }
                }
            }
            v16 := GenS1{F0: 593}
            _ = v16
        }
        _ = fi13
    }
    fmt.Println(v3)
    var v17 uint64 = 1
    if (-68 == 42) {
        var v18 uint8 = v1
        fmt.Println(v3, v3)
        _ = v18
    }
    v19 := GenS2{F0: 42, F1: -629, F2: 18446744073709551615}
    switch -511 {
    case 660:
        if false {
            var v20 bool = (genFn0((v3 / ((v17) | 1)), 9, -33) <= 42)
            for fi21 := 0; fi21 < 20; fi21++ {
                v3 = (v17 << ((18446744073709551615) & 63))
                v2 = append(v2, 0)
                var v22 int32 = 839
                _ = fi21
                _ = v22
            }
            v23 := GenS3{F0: -120, F1: (v3 + 18446744073709551615)}
            _ = v20
            _ = v23
        }
    default:
        switch (("\n\t" + v12) + "a") {
        case "foo bar":
            v12 = "hello"
        case "a":
            for fi24 := 0; fi24 < 11; fi24++ {
                for fi25 := 0; fi25 < 16; fi25++ {
                    if true {
                        if (!(2 != 65535)) {
                            v26 := GenS2{F0: 684, F1: 9223372036854775807, F2: 4294967295}
                            fmt.Println(v12, v3)
                            if len(v2) > 0 {
                                v2[0] = 2147483647
                            }
                            if (0 == 223) {
                                v3 -= v3
                                var v27 float64 = ((0.0 + 0.0) * 1e10)
                                if len(v2) > 0 {
                                    v2[0] = 0
                                }
                                _ = v27
                            }
                            _ = v26
                        }
                        v28 := GenS1{F0: 42}
                        v29 := GenS4{F0: 0}
                        _ = v28
                        _ = v29
                    }
                    if (881 > 255) {
                        fmt.Println(fi25)
                    } else {
                        var v30 uint32 = 497
                        _ = v30
                    }
                    switch (v12 + ((v12 + v12) + v12)) {
                    case "":
                        var v31 uint = 873
                        _ = v31
                    case "\n\t":
                        switch ((v12 + v12) + ("\n\t" + v12)) {
                        case "a":
                            v12 = "hello"
                        case "\n\t":
                            var v32 string = v12
                            _ = v32
                        }
                    default:
                        v12 = ("a" + v12)
                    }
                    _ = fi25
                }
                _ = fi24
            }
        default:
            var v33 int32 = 1
            _ = v33
        }
    }
    if true {
        for fi34 := 0; fi34 < 20; fi34++ {
            v35 := GenS4{F0: -759}
            v36 := GenS5{F0: 153}
            _ = fi34
            _ = v35
            _ = v36
        }
        v1 = v1
    }
    var v37 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v37
}
