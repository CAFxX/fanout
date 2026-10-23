package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int32 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        switch fi7 {
        case 330:
            var v8 int64 = -153
            _ = v8
        case 128:
            v9 := []int{128, (fi7 * fi7), 2147483647}
            _ = v9
        case 32767:
            switch "" {
            case "":
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 float32 = ((1e10 * 1e10) + (0.1 + 1.5))
                    switch fi7 {
                    case 128:
                        v12 := []int{fi10}
                        _ = v12
                    }
                    var v13 uint8 = 42
                    _ = fi10
                    _ = v11
                    _ = v13
                }
            }
        default:
            v14 := []int{fi7, (-(fi7))}
            _ = v14
        }
        _ = fi7
    }
    var v15 uint16 = 2
    for fi16 := 0; fi16 < 14; fi16++ {
        switch (("\n\t" + "\n\t") + ((("hello" + "foo bar") + "hello") + "hello")) {
        case "a":
            fmt.Println(fi16, fi16)
        case "":
            switch "hello" {
            case "hello":
                switch (("" + ("hello" + "")) + "\n\t") {
                case "hello":
                    if (!(!(v15 != (v15 << ((18446744073709551615) & 63))))) {
                        v15 = v15
                        v15 = 793
                        var v17 int16 = 0
                        _ = v17
                    }
                default:
                    var v18 int64 = 1
                    _ = v18
                }
            case "a":
                var v19 uint16 = (v15 ^ (v15 << ((18446744073709551615) & 63)))
                _ = v19
            case "\n\t":
                v20 := []int{-64, 127, 32768, (fi16 | (fi16 - fi16))}
                _ = v20
            }
        default:
            var v21 float64 = 1.5
            _ = v21
        }
        var v22 uint32 = 65535
        _ = fi16
        _ = v22
    }
    v15 += (v15 * 32767)
    _ = v15
    return 2147483647
}
func genFn1(p0 uint, p1 uint) string {
    if false {
        var v23 string = "foo bar"
        fmt.Println(p0, p1, p1)
        var v24 uint = p0
        _ = v23
        _ = v24
    }
    var v25 int64 = 1
    p1 = 42
    _ = p0
    _ = p1
    _ = v25
    return "hello"
}

func main() {
    v1 := []string{"a"}
    v2 := []uint32{}
    var v3 int64 = 105
    var v4 string = ("" + "a")
    var v5 uint = 65535
    v6 := []string{}
    v4 = v4
    var v26 string = v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (-73 == 1) {
        var v27 int8 = -1
        _ = v27
    } else {
        var v28 float32 = ((3.14159 * 1e10) + -2.25)
        var v29 int8 = 1
        for fi30 := 0; fi30 < 10; fi30++ {
            if false {
                fmt.Println(v4, v29, v5)
                var v31 uint64 = 553
                v5 -= (v5 - (v5 >> ((v5) & 63)))
                _ = v31
            }
            fmt.Println(v4)
            fmt.Println(v5, v29, v29)
            _ = fi30
        }
        _ = v28
        _ = v29
    }
    v4 = ((v4 + "") + ((("" + "\n\t") + v26) + ""))
    v4 = "hello"
    for fi32 := 0; fi32 < 17; fi32++ {
        for fi33 := 0; fi33 < 15; fi33++ {
            v34 := GenS1{F0: v5, F1: -128}
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    var v35 float64 = float64(-2.25)
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 float64 = v35
        v37 += v37
        _ = fi36
        _ = v37
    }
    v38 := GenS1{F0: (v5 << ((v5) & 63)), F1: -67}
    if true {
        v39 := GenS2{F0: (!true), F1: genFn0()}
        switch 4294967295 {
        case 17:
            var v40 float64 = 1.5
            _ = v40
        case 127:
            v41 := GenS1{F0: v5, F1: 0}
            _ = v41
        case 1:
            v42 := GenS1{F0: (v5 % ((815) | 1)), F1: 4}
            _ = v42
        default:
            var v43 int64 = v3
            _ = v43
        }
        if false {
            for fi44 := 0; fi44 < 7; fi44++ {
                fmt.Println(v5)
                v26 = (v4 + (v26 + "hello"))
                _ = fi44
            }
        }
        _ = v39
    }
    v45 := GenS3{F0: 799, F1: 34, F2: 0}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        v35 -= v35
        var v46 uint = v5
        for fi47 := 0; fi47 < 4; fi47++ {
            switch 0 {
            case -2147483648:
                v48 := GenS4{F0: -466}
                _ = v48
            }
            _ = fi47
        }
        _ = v46
    } else {
        fmt.Println(v4, v5, v26)
        switch 9223372036854775807 {
        case 32768:
            v3 -= (-(v3))
        case 974:
            var v49 uint64 = 2
            _ = v49
        case -2147483648:
            var v50 uint8 = 1
            _ = v50
        }
        var v51 int16 = -1
        _ = v51
    }
    fmt.Println(v35, v35, v5)
    var v52 uint8 = 81
    if len(v6) > 0 {
        v6[0] = ("foo bar" + "foo bar")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v45
    _ = v5
    _ = v52
    _ = v6
}
