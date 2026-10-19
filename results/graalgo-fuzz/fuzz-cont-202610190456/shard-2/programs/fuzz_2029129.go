package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 string; F2 uint32 }
func genFn0() int64 {
    var v6 float32 = float32(0.0)
    if (949 < -32768) {
        v7 := GenS1{F0: 128, F1: 42}
        _ = v7
    }
    var v8 string = "a"
    fmt.Println(v6)
    _ = v6
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 967
    v2 := make(map[uint32]uint32)
    var v3 uint = 443
    var v4 string = ("foo bar" + "hello")
    v5 := []uint32{1, 896, 2147483648, uint32(v1)}
    v3 += 79
    for fi9 := 0; fi9 < 2; fi9++ {
        if true {
            v3 += v3
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(fi9, v3, v3)
                v3 += 255
                var v11 int = fi10
                v3 = (v3 << ((v3) & 63))
                _ = fi10
                _ = v11
            }
        }
        var v12 uint8 = 124
        _ = fi9
        _ = v12
    }
    v13 := GenS2{F0: 91}
    v14 := GenS2{F0: 108}
    if false {
        v3 = v3
    } else {
        var v15 int8 = 0
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int64 = 334
            v17 = -535
            _ = fi16
            _ = v17
        }
        var v18 uint8 = 230
        _ = v15
        _ = v18
    }
    fmt.Println(v3, v3, v1)
    fmt.Println(v3)
    var v19 uint8 = 227
    fmt.Println(v1, v1, v19)
    v2[65535] = 2147483648
    for fi20 := 0; fi20 < 20; fi20++ {
        v3 = (v3 >> ((v3) & 63))
        v21 := GenS2{F0: v19}
        v22 := GenS1{F0: v19, F1: 342}
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int32 = -1
    switch (v4 + v4) {
    case "\n\t":
        for fi24 := 0; fi24 < 15; fi24++ {
            v25 := GenS3{F0: v19}
            _ = fi24
            _ = v25
        }
    case "hello":
        switch ("foo bar" + "hello") {
        case "a":
            v5 = append(v5, 4294967295)
        case "\n\t":
            v19 -= v19
        case "foo bar":
            v3 += (v3 << ((v3) & 63))
        }
    case "foo bar":
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 4; fi27++ {
                v19 += v19
                _ = fi27
            }
            if (true && false) {
                v28 := GenS1{F0: (^(v19)), F1: v23}
                v4 = (v4 + "hello")
                fmt.Println(v3)
                _ = v28
            }
            _ = fi26
        }
    default:
        fmt.Println(v4, v23)
    }
    v23 = v23
    switch 32768 {
    case 433:
        for fi29 := 0; fi29 < 10; fi29++ {
            v19 = uint8(v3)
            _ = fi29
        }
    default:
        for fi30 := 0; fi30 < 3; fi30++ {
            v19 = (v19 - v19)
            fmt.Println(v4, v3, v3)
            v19 += (^(v19))
            _ = fi30
        }
    }
    if true {
        v23 -= (v23 / ((v23) | 1))
        for fi31 := 0; fi31 < 3; fi31++ {
            for fi32 := 0; fi32 < 4; fi32++ {
                if (v3 < (v3 | (v3 + (^(v3))))) {
                    var v33 float32 = 0.1
                    switch -9223372036854775808 {
                    case 127:
                        var v34 uint32 = 42
                        _ = v34
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v35 := v2[564]
                    _ = v33
                    _ = v35
                }
                _ = fi32
            }
            _ = fi31
        }
    } else {
        switch v4 {
        case "a":
            for fi36 := 0; fi36 < 4; fi36++ {
                v19 += 138
                var v37 int = (fi36 << ((v3) & 63))
                var v38 bool = true
                _ = fi36
                _ = v37
                _ = v38
            }
        case "foo bar":
            var v39 int32 = v23
            _ = v39
        default:
            v40 := v2[366]
            _ = v40
        }
        var v41 uint32 = 2147483647
        for fi42 := 0; fi42 < 4; fi42++ {
            var v43 uint16 = 1
            v44 := GenS4{F0: v41}
            var v45 bool = ((v19 >= v19) && true)
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        _ = v41
    }
    switch 0 {
    case -1:
        switch v4 {
        case "a":
            switch -883 {
            case 32767:
                var v46 int16 = -32768
                _ = v46
            case 128:
                var v47 int64 = 9223372036854775807
                _ = v47
            default:
                for fi48 := 0; fi48 < 1; fi48++ {
                    v49 := GenS5{F0: true, F1: v4, F2: 4294967295}
                    _ = fi48
                    _ = v49
                }
            }
        case "\n\t":
            var v50 uint32 = 65535
            _ = v50
        case "":
            _ = len(v2)
        default:
            if (v1 == (^(v1))) {
                v51 := GenS5{F0: (false || (!(v4 <= v4))), F1: v4, F2: 4294967295}
                _ = v51
            }
        }
    case 32767:
        if true {
            var v52 int8 = 0
            if len(v5) > 0 {
                v5[0] = uint32(v3)
            }
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
