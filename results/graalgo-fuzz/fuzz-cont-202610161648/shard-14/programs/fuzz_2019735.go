package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 bool; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16 }
func genFn0() uint16 {
    var v7 uint16 = 2
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 1}
        v7 -= v7
        var v10 uint8 = 64
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 32768 {
    case -883:
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 float32 = 1.5
                v7 = (v7 << ((2) & 63))
                _ = fi12
                _ = v13
            }
            v7 = 932
            if (!(1 == 4294967295)) {
                v7 += (v7 & v7)
            } else {
                v14 := GenS1{F0: 156}
                _ = v14
            }
            _ = fi11
        }
    case -1:
        v7 = (v7 * v7)
    case 172:
        v15 := GenS2{F0: -133, F1: true, F2: (!((332 >= 2147483647) || false))}
        _ = v15
    default:
        if (341 >= 4294967295) {
            var v16 uint64 = 971
            v7 = 454
            switch 32768 {
            case -9223372036854775808:
                if false {
                    switch (("a" + ("hello" + "foo bar")) + "a") {
                    case "foo bar":
                        var v17 uint32 = 1
                        _ = v17
                    case "hello":
                        v18 := []int{32767, -2147483648, 32768}
                        _ = v18
                    default:
                        if false {
                            v19 := GenS2{F0: 278, F1: (1e10 > (0.0 / (3.14159 + (((0.1 + (0.0 / 3.14159)) + (1.5 + float64((1e10 / 0.1)))) / 1.0)))), F2: false}
                            var v20 int8 = int8(v7)
                            _ = v19
                            _ = v20
                        }
                    }
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v22 := []int{4294967295, fi21, fi21, (fi21 - fi21)}
                        _ = fi21
                        _ = v22
                    }
                    fmt.Println(v7, v7)
                }
            case 9223372036854775807:
                v23 := []int{128, -1, -2147483648, 9223372036854775807}
                _ = v23
            case 42:
                switch -1 {
                case -796:
                    var v24 uint = 18446744073709551615
                    _ = v24
                case -2147483648:
                    v7 = v7
                default:
                    v25 := GenS2{F0: 32767, F1: (0 <= 18446744073709551615), F2: true}
                    _ = v25
                }
            }
            _ = v16
        } else {
            for fi26 := 0; fi26 < 5; fi26++ {
                v27 := GenS1{F0: -2147483648}
                v28 := []int{-679, (fi26 - fi26)}
                _ = fi26
                _ = v27
                _ = v28
            }
            var v29 bool = (2147483647 != int(v7))
            v30 := GenS3{F0: 1000}
            _ = v29
            _ = v30
        }
    }
    if false {
        v7 = v7
    } else {
        var v31 uint = 286
        _ = v31
    }
    var v32 uint16 = (^(v7))
    _ = v32
    _ = v7
    return v32
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -128
    var v3 uint64 = 18446744073709551615
    var v4 int8 = v2
    var v5 uint32 = 65535
    v6 := []uint32{}
    v33 := GenS4{F0: -878}
    v3 += (v3 | (v3 & (v3 << ((uint(v3)) & 63))))
    var v34 int16 = -457
    if (18446744073709551615 == 0) {
        v4 = v4
    } else {
        v35 := GenS3{F0: v34}
        v34 = v34
        for fi36 := 0; fi36 < 4; fi36++ {
            fmt.Println(v34, v5, v2)
            _ = fi36
        }
        _ = v35
    }
    fmt.Println(v3, v4, v5)
    v2 = 42
    if (false || ((42 == 187) || (v3 < v3))) {
        for fi37 := 0; fi37 < 20; fi37++ {
            v38 := GenS2{F0: (v34 * (-(v34))), F1: (1 >= 66), F2: false}
            _ = fi37
            _ = v38
        }
        v39 := GenS3{F0: 1}
        var v40 uint = 2
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        if true {
            v42 := GenS1{F0: -371}
            var v43 uint64 = v3
            _ = v42
            _ = v43
        } else {
            var v44 int = 32768
            _ = v44
        }
        for fi45 := 0; fi45 < 15; fi45++ {
            fmt.Println(v2, v4, v2)
            fmt.Println(v2, v5, v34)
            var v46 bool = true
            fmt.Println(v2, v3, v34)
            var v47 int32 = 1
            v48 := GenS3{F0: -1}
            _ = fi45
            _ = v46
            _ = v47
            _ = v48
        }
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
