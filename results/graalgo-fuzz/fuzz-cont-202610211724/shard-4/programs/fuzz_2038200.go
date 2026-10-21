package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint8 = 0
    var v2 float64 = 0.0
    var v3 int16 = 501
    var v4 int8 = 127
    var v5 uint16 = 616
    fmt.Println(v1)
    var v6 bool = true
    v1 = v1
    switch "" {
    case "\n\t":
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint32 = 0
            if v6 {
                v9 := []int{(fi7 | fi7), fi7}
                if len(v9) > 0 {
                    v9[0] = fi7
                }
                var v10 uint32 = (v8 / ((v8) | 1))
                _ = v10
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    case "a":
        for fi11 := 0; fi11 < 17; fi11++ {
            if v6 {
                v12 := []int{-2147483648, fi11}
                if (2147483647 <= -635) {
                    v6 = v6
                } else {
                    if (0 > 18446744073709551615) {
                        v13 := GenS1{F0: "\n\t"}
                        var v14 string = (("foo bar" + "hello") + "hello")
                        _ = v13
                        _ = v14
                    }
                    v3 -= int16(v4)
                    v12 = append(v12, int(v1))
                }
                _ = v12
            }
            v15 := GenS1{F0: (("a" + (("\n\t" + ("hello" + "a")) + ((("hello" + "\n\t") + (("\n\t" + ("\n\t" + "hello")) + ("hello" + "a"))) + (("a" + "") + "\n\t")))) + "\n\t")}
            v5 = 2
            _ = fi11
            _ = v15
        }
    case "foo bar":
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint = 255
            v18 := GenS2{F0: v5, F1: (v17 % ((v17) | 1))}
            v3 = (v3 >> ((uint(v5)) & 63))
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        v19 := GenS1{F0: ("" + ("foo bar" + ("\n\t" + "a")))}
        _ = v19
    }
    v3 = 1
    for fi20 := 0; fi20 < 18; fi20++ {
        if (true && v6) {
            var v21 uint32 = 996
            _ = v21
        }
        _ = fi20
    }
    if (!((-1 <= -9223372036854775808) && v6)) {
        v2 = float64(3.14159)
        v22 := GenS1{F0: "\n\t"}
        switch -9223372036854775808 {
        case -403:
            v3 -= 1
        case 32768:
            for fi23 := 0; fi23 < 1; fi23++ {
                fmt.Println(fi23, v6)
                _ = fi23
            }
        case 128:
            var v24 uint32 = 65535
            _ = v24
        }
        _ = v22
    } else {
        if true {
            for fi25 := 0; fi25 < 18; fi25++ {
                v26 := []int{}
                var v27 int = fi25
                v3 += v3
                _ = fi25
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 12; fi28++ {
                if v6 {
                    fmt.Println(v3, v3)
                }
                if ((!(((fi28 <= fi28) || false) && (false && v6))) && false) {
                    v29 := []int{fi28, fi28, (fi28 | fi28)}
                    var v30 uint8 = v1
                    _ = v29
                    _ = v30
                }
                _ = fi28
            }
            var v31 uint32 = 991
            _ = v31
        } else {
            var v32 int16 = v3
            switch 32767 {
            case 42:
                v4 += -4
            case -1:
                v6 = v6
            case -2147483648:
                v33 := []int{9223372036854775807, -9223372036854775808, 4294967295}
                _ = v33
            default:
                v34 := GenS3{F0: 6, F1: 4294967295, F2: 4294967295}
                _ = v34
            }
            _ = v32
        }
        var v35 float32 = 0.1
        for fi36 := 0; fi36 < 16; fi36++ {
            fmt.Println(v4, v2)
            switch fi36 {
            case -441:
                v3 += -282
            case 2147483647:
                fmt.Println(v3, fi36, v1)
            case -887:
                v37 := GenS3{F0: 2, F1: 106, F2: 1}
                _ = v37
            }
            _ = fi36
        }
        _ = v35
    }
    if true {
        switch 4294967295 {
        case -75:
            v1 -= 245
        default:
            if ((v3 ^ v3) == -66) {
                switch -9223372036854775808 {
                case 564:
                    for fi38 := 0; fi38 < 14; fi38++ {
                        v4 = (-(v4))
                        var v39 int8 = -93
                        switch (^(fi38)) {
                        case 2147483647:
                            fmt.Println(v1, v39)
                        }
                        _ = fi38
                        _ = v39
                    }
                }
            }
        }
        fmt.Println(v3, v6, v6)
        v1 = v1
    } else {
        v40 := []int{-280, 127}
        var v41 uint8 = 255
        _ = v40
        _ = v41
    }
    if false {
        v4 = v4
    } else {
        v4 = 100
        v42 := GenS4{F0: -87}
        _ = v42
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 float32 = (1.5 / -2.25)
            _ = fi44
            _ = v45
        }
        v46 := GenS3{F0: 1, F1: 117, F2: 893}
        _ = fi43
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
