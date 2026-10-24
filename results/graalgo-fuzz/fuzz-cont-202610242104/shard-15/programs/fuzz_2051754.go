package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint) uint64 {
    fmt.Println(p0)
    if true {
        var v7 uint16 = 2
        if (p0 != 968) {
            v7 += 255
            fmt.Println(v7, p0)
            switch 42 {
            case 128:
                for fi8 := 0; fi8 < 20; fi8++ {
                    v9 := GenS1{F0: 42, F1: 218, F2: 2}
                    var v10 float64 = 0.0
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case 32768:
                if true {
                    v7 += (v7 | v7)
                    for fi11 := 0; fi11 < 3; fi11++ {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            v13 := GenS1{F0: 0, F1: uint8(v7), F2: 4294967295}
                            p0 += p0
                            v14 := GenS1{F0: 255, F1: 190, F2: 4294967295}
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = fi11
                    }
                    v15 := []int{-2147483648}
                    _ = v15
                } else {
                    var v16 int64 = -1
                    _ = v16
                }
            case 32767:
                fmt.Println(v7, p0)
            }
        } else {
            var v17 int64 = 0
            _ = v17
        }
        _ = v7
    } else {
        p0 = p0
        switch 539 {
        case 42:
            v18 := []int{2147483647}
            _ = v18
        case 128:
            p0 = 2
        case 2147483647:
            var v19 int8 = 77
            _ = v19
        }
        var v20 uint8 = 127
        _ = v20
    }
    var v21 uint8 = 1
    v22 := []int{-291, 0, 4}
    _ = p0
    _ = v21
    _ = v22
    return 42
}

func main() {
    var v1 int8 = 44
    v2 := make(map[uint32]bool)
    v3 := []uint32{756, 4294967295, 2, 747}
    var v4 int64 = 9223372036854775807
    var v5 uint32 = 759
    v6 := make(map[uint32]uint32)
    var v23 uint16 = 42
    v4 = (^(v4))
    if false {
        var v24 uint32 = (v5 | (^(v5)))
        fmt.Println(v5, v24)
        _ = v24
    }
    var v25 uint = 255
    v23 -= v23
    if (false || (-540 > -32768)) {
        if false {
            v26 := GenS1{F0: 0, F1: 39, F2: (^(v5))}
            fmt.Println(v1, v23)
            if (4294967295 > int(v4)) {
                var v27 uint16 = v23
                if (-1 == -118) {
                    delete(v6, v5)
                }
                _ = v27
            } else {
                v28 := GenS2{F0: 32767, F1: 0}
                var v29 bool = (v23 > (v23 + v23))
                _ = v28
                _ = v29
            }
            _ = v26
        } else {
            switch 32768 {
            case 128:
                if len(v3) > 0 {
                    v3[0] = uint32(v23)
                }
            case -2147483648:
                v30 := v2[1]
                _ = v30
            }
            if false {
                v31 := GenS3{F0: 18446744073709551615}
                switch -9223372036854775808 {
                case 283:
                    var v32 uint16 = v23
                    _ = v32
                }
                _ = v31
            } else {
                var v33 uint = v25
                v25 = (v25 ^ 667)
                fmt.Println(v1, v4, v33)
                _ = v33
            }
            var v34 float32 = 1.5
            _ = v34
        }
    } else {
        var v35 string = ("" + "hello")
        v5 += 2147483647
        _ = v35
    }
    var v36 float32 = 0.0
    v4 = -9223372036854775808
    v37 := GenS3{F0: genFn0(2)}
    _ = len(v2)
    var v38 int32 = -360
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v4 = v4
    case "a":
        v39 := GenS1{F0: uint8(v25), F1: uint8(v25), F2: (v5 - 102)}
        _ = v39
    default:
        v40 := GenS2{F0: -32768, F1: 127}
        _ = v40
    }
    v41 := GenS3{F0: 292}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v42 := GenS2{F0: -32768, F1: 255}
    if true {
        for fi43 := 0; fi43 < 7; fi43++ {
            var v44 uint16 = v23
            _ = fi43
            _ = v44
        }
    }
    fmt.Println(v23)
    fmt.Println(v25)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
