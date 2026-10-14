package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 string; F1 int }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = 0
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 42, F2: 42}
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 866
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    var v6 int16 = 32767
    v2 -= v2
    v10 := GenS1{F0: 18446744073709551615, F1: 115, F2: (v2 >> ((71) & 63))}
    fmt.Println(v2, v6)
    v11 := v4[473]
    var v12 int8 = -125
    switch -1 {
    case 2147483647:
        var v13 string = "\n\t"
        _ = v13
    case 317:
        v11 = -2147483648
    case 9223372036854775807:
        v12 = (v12 & v12)
    }
    v14 := v5[2147483647]
    v15 := GenS2{F0: 2, F1: "a"}
    v14 = v14
    for fi16 := 0; fi16 < 17; fi16++ {
        if true {
            var v17 uint64 = 717
            v14 += (v14 + (^(v14)))
            if false {
                var v18 uint32 = (v14 - (^(v14)))
                fmt.Println(v17)
                if ("a" == "foo bar") {
                    var v19 uint = 537
                    _ = v19
                } else {
                    if (167 <= 365) {
                        for fi20 := 0; fi20 < 14; fi20++ {
                            v21 := GenS3{F0: (v18 << ((669) & 63)), F1: 79}
                            _ = len(v3)
                            delete(v4, (^(v18)))
                            _ = fi20
                            _ = v21
                        }
                    } else {
                        v22 := GenS4{F0: ("a" + ""), F1: (-(fi16))}
                        _ = v22
                    }
                    var v23 uint64 = v17
                    _ = v23
                }
                _ = v18
            } else {
                var v24 int64 = 42
                var v25 float64 = 1.5
                var v26 uint32 = (v14 - 4294967295)
                _ = v24
                _ = v25
                _ = v26
            }
            _ = v17
        } else {
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 uint32 = v14
                if (100 <= v12) {
                    v29 := GenS2{F0: (v14 / ((v14) | 1)), F1: ("\n\t" + ("" + (("\n\t" + "hello") + (("\n\t" + "") + "foo bar"))))}
                    if ((18446744073709551615 > 1) && (-25 >= -9223372036854775808)) {
                        var v30 float32 = -2.25
                        _ = v30
                    } else {
                        var v31 uint64 = 1
                        _ = v31
                    }
                    delete(v5, v28)
                    _ = v29
                }
                v32 := GenS3{F0: 42, F1: int8(v28)}
                _ = fi27
                _ = v28
                _ = v32
            }
            var v33 int16 = (v6 & v6)
            if false {
                v33 = (^(v33))
            } else {
                for fi34 := 0; fi34 < 19; fi34++ {
                    v2 = (^(v2))
                    _ = fi34
                }
            }
            _ = v33
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 uint16 = (v2 | 578)
            var v37 int16 = v6
            v11 = fi16
            _ = fi35
            _ = v36
            _ = v37
        }
        if ((false || false) && true) {
            v38 := v5[848]
            if (!(158 < v2)) {
                if true {
                    var v39 int32 = 42
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    _ = v39
                }
                if true {
                    v6 = (v6 - (^(v6)))
                    var v40 int = v11
                    for fi41 := 0; fi41 < 9; fi41++ {
                        var v42 bool = ("a" > "\n\t")
                        _ = fi41
                        _ = v42
                    }
                    _ = v40
                }
            } else {
                var v43 uint8 = 128
                _ = v43
            }
            _ = v38
        }
        _ = fi16
    }
    if ((v6 - (-(v6))) != v6) {
        switch "hello" {
        case "hello":
            v2 -= (v2 ^ 348)
        }
        v6 -= (^(v6))
        var v44 string = "\n\t"
        _ = v44
    }
    var v45 string = "foo bar"
    switch "a" {
    case "\n\t":
        for fi46 := 0; fi46 < 10; fi46++ {
            v12 = -1
            var v47 uint64 = 18446744073709551615
            fmt.Println(fi46)
            _ = fi46
            _ = v47
        }
    default:
        switch v45 {
        case "":
            v48 := GenS2{F0: v14, F1: (v45 + ((("a" + ("hello" + v45)) + v45) + ((v45 + v45) + v45)))}
            _ = v48
        case "hello":
            v5[(^(v14))] = (v14 + v14)
        }
    }
    var v49 uint32 = v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v5
    _ = v6
}
