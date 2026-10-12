package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS5 struct { F0 uint32; F1 int; F2 bool }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 int16 = 1
    v2 := []uint32{593}
    v3 := []uint32{2, 65535, 42}
    v4 := []string{"a"}
    var v5 string = (("\n\t" + "foo bar") + "hello")
    v6 := []uint32{}
    if true {
        if len(v2) > 0 {
            v2[0] = 65535
        }
    }
    var v7 uint = 765
    var v8 bool = (0 >= -2147483648)
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int = (fi9 % ((fi9) | 1))
        if v8 {
            var v11 int16 = 24
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    v12 := GenS1{F0: -376, F1: -1}
    if v8 {
        for fi13 := 0; fi13 < 7; fi13++ {
            v5 = v5
            v8 = (!v8)
            for fi14 := 0; fi14 < 12; fi14++ {
                if v8 {
                    v7 += (v7 >> ((255) & 63))
                    var v15 uint8 = 152
                    _ = v15
                }
                v8 = (0 == 128)
                fmt.Println(v7)
                _ = fi14
            }
            _ = fi13
        }
        _ = len(v3)
    }
    if v8 {
        v5 = ""
        fmt.Println(v5, v7)
        for fi16 := 0; fi16 < 18; fi16++ {
            v5 = v5
            _ = fi16
        }
    } else {
        if false {
            for fi17 := 0; fi17 < 12; fi17++ {
                v1 = 42
                v6 = append(v6, 464)
                if (2 < 0) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v18 := GenS1{F0: -1, F1: 0}
                    _ = v18
                }
                _ = fi17
            }
            var v19 float64 = 0.0
            _ = v19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint64 = 0
            fmt.Println(v5, fi20, v1)
            switch (fi20 ^ 32768) {
            case 989:
                v6 = append(v6, 1)
            case 32767:
                var v22 int64 = -1
                _ = v22
            case 42:
                fmt.Println(v21, v5, v1)
            }
            _ = fi20
            _ = v21
        }
        for fi23 := 0; fi23 < 14; fi23++ {
            v8 = (128 >= 127)
            _ = fi23
        }
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        v5 = (v5 + v5)
        var v25 float64 = 3.14159
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS2{F0: (0 == 1), F1: (-(fi26))}
        if len(v3) > 0 {
            v3[0] = 42
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            v7 -= v7
            fmt.Println(fi26, v5)
            fmt.Println(v1, v5)
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    v29 := GenS1{F0: -1, F1: -520}
    v30 := GenS2{F0: v8, F1: 127}
    var v31 int32 = -2147483648
    for fi32 := 0; fi32 < 17; fi32++ {
        fmt.Println(v7)
        v7 = (^(v7))
        switch ("" + (v5 + v5)) {
        case "\n\t":
            var v33 uint = v7
            _ = v33
        }
        _ = fi32
    }
    var v34 bool = (81 == 107)
    switch v5 {
    case "hello":
        var v35 string = (v5 + "foo bar")
        _ = v35
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        if len(v4) > 0 {
            v4[0] = v5
        }
        v31 = int32(v7)
        _ = fi36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        if (v5 > ("a" + ("hello" + ("\n\t" + v5)))) {
            fmt.Println(fi37)
            var v38 int = (^(fi37))
            v39 := GenS3{F0: v31, F1: 1}
            _ = v38
            _ = v39
        }
        switch fi37 {
        case 85:
            switch fi37 {
            case 127:
                switch 403 {
                case -2:
                    v5 = (("hello" + "") + (((v5 + ((v5 + "foo bar") + v5)) + (v5 + v5)) + ("foo bar" + ("a" + "\n\t"))))
                }
            case -2147483648:
                if len(v2) > 0 {
                    v2[0] = 993
                }
            case -304:
                v40 := GenS4{F0: 2, F1: 128, F2: 42}
                _ = v40
            default:
                v1 += v1
            }
        case 128:
            switch ((v5 + "") + (((v5 + (("foo bar" + v5) + (("\n\t" + ("a" + v5)) + v5))) + ((v5 + v5) + "a")) + v5)) {
            case "\n\t":
                var v41 bool = v8
                _ = v41
            case "hello":
                var v42 int16 = v1
                _ = v42
            default:
                var v43 float32 = 1.5
                _ = v43
            }
        case 4294967295:
            if v8 {
                v1 += v1
                fmt.Println(v34, fi37)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                switch "hello" {
                case "\n\t":
                    var v44 int16 = 0
                    _ = v44
                }
                switch 0 {
                case -897:
                    v45 := GenS1{F0: v31, F1: -1}
                    _ = v45
                case 128:
                    var v46 string = v5
                    _ = v46
                case 122:
                    if (32767 >= 65535) {
                        v47 := GenS5{F0: 0, F1: 127, F2: v8}
                        for fi48 := 0; fi48 < 5; fi48++ {
                            fmt.Println(fi37, v31)
                            _ = fi48
                        }
                        v6 = append(v6, 2147483648)
                        _ = v47
                    }
                default:
                    var v49 string = v5
                    _ = v49
                }
            }
        default:
            for fi50 := 0; fi50 < 8; fi50++ {
                _ = len(v4)
                switch (v5 + v5) {
                case "hello":
                    var v51 string = (v5 + (("foo bar" + "foo bar") + (("" + (v5 + (v5 + "hello"))) + v5)))
                    _ = v51
                case "a":
                    switch fi37 {
                    case 0:
                        if true {
                            v5 = v5
                        }
                    case -1:
                        fmt.Println(v34, fi37)
                    case 127:
                        v5 = "foo bar"
                    }
                default:
                    fmt.Println(v31)
                }
                fmt.Println(fi50)
                _ = fi50
            }
        }
        _ = fi37
    }
    if (v8 && (273 < 1)) {
        v52 := GenS6{F0: 18446744073709551615}
        v53 := GenS3{F0: v31, F1: 49}
        v5 = v5
        _ = v52
        _ = v53
    } else {
        var v54 int16 = (v1 - v1)
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
