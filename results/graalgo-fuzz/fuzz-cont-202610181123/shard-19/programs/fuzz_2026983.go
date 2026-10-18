package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32; F1 uint32; F2 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int; F1 int }
func genFn0(p0 int16) bool {
    for fi7 := 0; fi7 < 5; fi7++ {
        switch ("" + "") {
        case "":
            v8 := GenS1{F0: (-(fi7)), F1: 0, F2: fi7}
            _ = v8
        case "foo bar":
            v9 := GenS2{F0: 255}
            _ = v9
        case "hello":
            var v10 float64 = float64(-2.25)
            _ = v10
        }
        _ = fi7
    }
    var v11 int32 = 2147483647
    v12 := []int{595, 127}
    v13 := GenS2{F0: 2}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return (3.14159 <= (0.0 / float32(float64(1e10))))
}

func main() {
    var v1 string = "a"
    var v2 int64 = 871
    v3 := []uint32{0, 1}
    v4 := []int{}
    v5 := make(map[uint32]int)
    var v6 bool = (65535 != 0)
    for fi14 := 0; fi14 < 7; fi14++ {
        v2 -= 9223372036854775807
        _ = len(v3)
        _ = fi14
    }
    var v15 bool = false
    if false {
        switch v1 {
        case "hello":
            var v16 uint32 = 2
            _ = v16
        default:
            var v17 uint = 1
            _ = v17
        }
        v18 := GenS3{F0: -177}
        _ = v18
    }
    v15 = false
    var v19 uint8 = 1
    v2 = 945
    var v20 uint64 = 4294967295
    if (-875 == -32768) {
        v21 := GenS4{F0: (v19 ^ uint8(v20)), F1: 110, F2: 32767}
        _ = v21
    }
    v22 := GenS5{F0: (v1 + ((v1 + "a") + v1))}
    v23 := GenS6{F0: v1}
    for fi24 := 0; fi24 < 1; fi24++ {
        if v6 {
            if false {
                _ = len(v4)
                if (0 > (v2 & (v2 ^ (v2 & 1)))) {
                    v5[42] = 1
                } else {
                    _ = len(v5)
                    v25 := GenS5{F0: v1}
                    _ = v25
                }
                switch (fi24 | (fi24 - fi24)) {
                case 32767:
                    v26 := GenS7{F0: 0, F1: 42, F2: 831}
                    _ = v26
                case 32768:
                    for fi27 := 0; fi27 < 8; fi27++ {
                        var v28 int64 = v2
                        if (!v15) {
                            v29 := GenS8{F0: 0}
                            var v30 bool = false
                            var v31 uint = 65535
                            _ = v29
                            _ = v30
                            _ = v31
                        } else {
                            switch (^(fi27)) {
                            case 32767:
                                switch "foo bar" {
                                case "\n\t":
                                    v32 := GenS4{F0: v19, F1: v19, F2: 0}
                                    _ = v32
                                case "":
                                    var v33 float64 = 1.5
                                    _ = v33
                                default:
                                    v28 = v2
                                }
                            case 2147483647:
                                v6 = (42 != (^(fi27)))
                            }
                            v34 := GenS1{F0: (fi24 * -3), F1: 127, F2: 127}
                            _ = v34
                        }
                        _ = fi27
                        _ = v28
                    }
                default:
                    v35 := GenS9{F0: fi24, F1: (fi24 + int(v20))}
                    _ = v35
                }
            } else {
                for fi36 := 0; fi36 < 12; fi36++ {
                    var v37 uint8 = v19
                    v15 = (v37 <= (v37 % (((^(v37))) | 1)))
                    v38 := GenS7{F0: 0, F1: 1, F2: 255}
                    var v39 float32 = 0.0
                    v19 = v19
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v39
                }
                var v40 uint16 = 42
                for fi41 := 0; fi41 < 12; fi41++ {
                    var v42 int32 = 484
                    v15 = v6
                    _ = fi41
                    _ = v42
                }
                _ = v40
            }
        }
        v43 := v5[uint32(v20)]
        _ = fi24
        _ = v43
    }
    var v44 int64 = (v2 % ((9223372036854775807) | 1))
    switch "a" {
    case "a":
        var v45 int8 = -1
        _ = v45
    case "":
        for fi46 := 0; fi46 < 15; fi46++ {
            switch v1 {
            case "":
                v2 = v44
            case "foo bar":
                fmt.Println(fi46, v2, v44)
            case "hello":
                for fi47 := 0; fi47 < 12; fi47++ {
                    v6 = ((32767 == 255) || v15)
                    _ = fi47
                }
            default:
                if v15 {
                    v1 = ("" + v1)
                    v48 := GenS6{F0: v1}
                    _ = v48
                } else {
                    v5[4294967295] = 2147483647
                }
            }
            _ = fi46
        }
    default:
        v20 = (v20 % ((v20) | 1))
    }
    v49 := GenS1{F0: -9223372036854775808, F1: 66, F2: 2147483647}
    if (v15 && v15) {
        if true {
            var v50 uint = 1
            v6 = ((1.5 == 3.14159) && v15)
            _ = v50
        } else {
            var v51 int = int(v20)
            _ = v51
        }
        var v52 int = 1
        _ = v52
    } else {
        fmt.Println(v6, v6, v2)
        v44 = (v2 | -1)
        var v53 int32 = -2147483648
        _ = v53
    }
    for fi54 := 0; fi54 < 4; fi54++ {
        v2 -= -9223372036854775808
        switch (fi54 - (fi54 >> ((4294967295) & 63))) {
        case 9:
            v15 = (2 == 18446744073709551615)
        default:
            var v55 int32 = -251
            _ = v55
        }
        _ = fi54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v44
    _ = v49
    _ = v5
    _ = v6
}
