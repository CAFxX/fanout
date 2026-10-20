package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32, p1 int16) int {
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := []int{127, (fi5 / (((fi5 - (fi6 - (fi6 + (fi6 << ((157) & 63)))))) | 1)), fi5}
            fmt.Println(p0, p1)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := []int{-620, -1}
    fmt.Println(p0)
    v8 = append(v8, 4294967295)
    _ = p0
    _ = p1
    _ = v8
    return 9223372036854775807
}
func genFn1(p0 int32) uint16 {
    for fi9 := 0; fi9 < 15; fi9++ {
        p0 = p0
        _ = fi9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if ((-2.25 + float32(3.14159)) < float32(1.5)) {
            for fi11 := 0; fi11 < 13; fi11++ {
                fmt.Println(fi10, p0)
                _ = fi11
            }
            fmt.Println(p0, fi10)
        }
        fmt.Println(fi10)
        _ = fi10
    }
    _ = p0
    return 65535
}

func main() {
    var v1 bool = true
    var v2 int64 = 1
    v3 := make(map[uint32]bool)
    var v4 int64 = (v2 / ((1) | 1))
    if (!(false || (!v1))) {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint64 = 0
            v14 := GenS1{F0: (fi12 << ((991) & 63)), F1: 0, F2: 255}
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        switch -9223372036854775808 {
        case 128:
            v15 := GenS1{F0: -806, F1: v4, F2: 906}
            _ = v15
        case 4294967295:
            for fi16 := 0; fi16 < 5; fi16++ {
                v4 = 42
                _ = fi16
            }
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            for fi18 := 0; fi18 < 13; fi18++ {
                if false {
                    v2 = int64(fi18)
                } else {
                    v1 = v1
                }
                v19 := GenS1{F0: -9223372036854775808, F1: -1, F2: 667}
                v4 -= v4
                _ = fi18
                _ = v19
            }
            fmt.Println(fi17, v2)
            _ = fi17
        }
    }
    switch genFn0(2147483647, -1) {
    case 13:
        v1 = (673 > 0)
    case -2147483648:
        if v1 {
            var v20 uint = 342
            _ = v20
        } else {
            var v21 uint = 0
            _ = v21
        }
    case 346:
        v1 = ((1 > -2147483648) || (4294967295 >= 0))
    default:
        v22 := GenS1{F0: 42, F1: -9223372036854775808, F2: 115}
        _ = v22
    }
    fmt.Println(v1)
    v23 := GenS1{F0: 533, F1: (v4 << ((33) & 63)), F2: 18446744073709551615}
    v24 := v3[2147483648]
    fmt.Println(v2)
    var v25 int8 = -56
    _ = len(v3)
    switch -635 {
    case 127:
        for fi26 := 0; fi26 < 5; fi26++ {
            v1 = v24
            if (!v24) {
                v1 = v1
                v27 := []int{(-(fi26)), (fi26 >> ((235) & 63))}
                var v28 int8 = (v25 + 1)
                _ = v27
                _ = v28
            }
            _ = fi26
        }
    }
    switch 0 {
    case -983:
        if (!v1) {
            fmt.Println(v24, v24, v2)
        } else {
            v29 := []int{-9223372036854775808, genFn0(65535, 32767), genFn0(42, 42), 1}
            switch "hello" {
            case "\n\t":
                for fi30 := 0; fi30 < 4; fi30++ {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        fmt.Println(v25)
                        switch 626 {
                        case -140:
                            fmt.Println(v1, v2)
                        case 32767:
                            var v32 float32 = 1.5
                            _ = v32
                        }
                        _ = fi31
                    }
                    _ = fi30
                }
            default:
                var v33 uint8 = 65
                _ = v33
            }
            delete(v3, 224)
            _ = v29
        }
    }
    if true {
        v34 := GenS1{F0: int(v25), F1: v2, F2: 939}
        _ = v34
    }
    v2 = v4
    switch "" {
    case "foo bar":
        switch "a" {
        case "hello":
            var v35 int32 = 479
            _ = v35
        }
    }
    v25 -= (v25 / ((v25) | 1))
    var v36 int = -562
    v37 := GenS2{F0: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
