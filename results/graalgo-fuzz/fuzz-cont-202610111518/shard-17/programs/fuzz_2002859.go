package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int64; F1 uint64; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string }
func genFn0() uint8 {
    var v7 uint8 = 164
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := []int{(fi8 + (^(fi9))), fi9, fi8}
            v7 -= 2
            if (0 != 1) {
                var v11 uint32 = 705
                _ = v11
            }
            if true {
                v10 = append(v10, 78)
                if false {
                    _ = len(v10)
                    var v12 string = ("" + "")
                    _ = v12
                }
            } else {
                v7 -= 193
                v7 = (v7 % (((v7 + v7)) | 1))
                fmt.Println(v7)
            }
            _ = fi9
            _ = v10
        }
        v13 := GenS1{F0: -1, F1: ("" + ("a" + "foo bar")), F2: fi8}
        _ = fi8
        _ = v13
    }
    if true {
        var v14 uint64 = 42
        v15 := GenS1{F0: 0, F1: "", F2: -2147483648}
        _ = v14
        _ = v15
    } else {
        fmt.Println(v7)
        v7 -= 1
    }
    v16 := []int{-809}
    _ = v16
    _ = v7
    return v7
}

func main() {
    var v1 int = -570
    var v2 int32 = -371
    var v3 uint32 = 4294967295
    var v4 int16 = 42
    v5 := []string{"", (("hello" + "hello") + "hello")}
    var v6 string = "foo bar"
    v6 = (v6 + "")
    for fi17 := 0; fi17 < 20; fi17++ {
        v3 = 65535
        if (1 == 0) {
            v18 := GenS2{F0: (v6 + ("\n\t" + "hello")), F1: 42}
            if true {
                v1 += (fi17 ^ (-(v1)))
                switch (v6 + v6) {
                case "\n\t":
                    switch fi17 {
                    case 4294967295:
                        switch (("a" + (("\n\t" + ("hello" + "foo bar")) + "hello")) + v6) {
                        case "\n\t":
                            var v19 uint8 = 127
                            _ = v19
                        case "hello":
                            var v20 int16 = v4
                            _ = v20
                        default:
                            if true {
                                v1 += (fi17 + v1)
                            }
                        }
                    case 2147483647:
                        v2 -= (v2 << ((18446744073709551615) & 63))
                    }
                }
            }
            switch fi17 {
            case 2147483647:
                var v21 float64 = -2.25
                _ = v21
            case -818:
                var v22 uint16 = 547
                _ = v22
            case -9223372036854775808:
                v1 += (v1 % ((fi17) | 1))
            }
            _ = v18
        } else {
            for fi23 := 0; fi23 < 12; fi23++ {
                v6 = ("a" + v6)
                switch v1 {
                case -2147483648:
                    var v24 float32 = -2.25
                    _ = v24
                default:
                    if false {
                        v4 = -1
                    }
                }
                var v25 bool = true
                _ = fi23
                _ = v25
            }
        }
        _ = fi17
    }
    switch v6 {
    case "a":
        v26 := GenS1{F0: -1, F1: v6, F2: (v1 >> ((821) & 63))}
        _ = v26
    case "foo bar":
        v27 := GenS1{F0: 0, F1: v6, F2: v1}
        _ = v27
    case "":
        v6 = "foo bar"
    }
    switch v6 {
    case "\n\t":
        v3 += (v3 << ((255) & 63))
    case "hello":
        _ = len(v5)
    case "foo bar":
        fmt.Println(v3)
    default:
        switch v1 {
        case 42:
            fmt.Println(v2)
        case 495:
            if true {
                var v28 uint = 42
                if (18446744073709551615 >= (v28 * (^(v28)))) {
                    var v29 int64 = -419
                    v2 = -239
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v29
                } else {
                    v2 = v2
                }
                _ = v28
            } else {
                var v30 float32 = 0.1
                var v31 int = v1
                switch v31 {
                case 156:
                    fmt.Println(v30)
                }
                _ = v30
                _ = v31
            }
        case 152:
            for fi32 := 0; fi32 < 8; fi32++ {
                switch (fi32 / ((fi32) | 1)) {
                case 732:
                    var v33 uint16 = 2
                    _ = v33
                }
                _ = fi32
            }
        }
    }
    v34 := GenS3{F0: 42, F1: v1}
    v2 -= -1
    fmt.Println(v3)
    for fi35 := 0; fi35 < 7; fi35++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if ((v6 + v6) < "hello") {
            if false {
                v36 := GenS3{F0: 42, F1: (v1 / (((fi35 / ((32768) | 1))) | 1))}
                fmt.Println(v1)
                for fi37 := 0; fi37 < 9; fi37++ {
                    if false {
                        var v38 bool = false
                        _ = v38
                    }
                    _ = fi37
                }
                _ = v36
            }
            if (true || ((v2 / (((-(v2))) | 1)) != v2)) {
                for fi39 := 0; fi39 < 12; fi39++ {
                    v3 += v3
                    fmt.Println(fi39, v1)
                    v40 := GenS2{F0: ("foo bar" + ("" + "foo bar")), F1: -319}
                    _ = fi39
                    _ = v40
                }
                v1 += fi35
                var v41 bool = ("hello" != v6)
                _ = v41
            }
            var v42 float64 = 3.14159
            _ = v42
        }
        _ = fi35
    }
    if true {
        fmt.Println(v6, v6)
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        for fi44 := 0; fi44 < 18; fi44++ {
            if (2 < 2) {
                v45 := GenS4{F0: -9223372036854775808, F1: 456, F2: 32767}
                v46 := GenS5{F0: v2}
                var v47 bool = false
                var v48 int8 = 0
                _ = v45
                _ = v46
                _ = v47
                _ = v48
            }
            v6 = (v6 + (v6 + v6))
            _ = fi44
        }
        v49 := GenS6{F0: "hello"}
        if (!false) {
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            var v50 uint8 = 255
            if len(v5) > 0 {
                v5[0] = "hello"
            }
            _ = v50
        } else {
            v6 = (("foo bar" + v6) + v6)
            v4 = v4
            var v51 uint64 = 4294967295
            _ = v51
        }
        _ = fi43
        _ = v49
    }
    fmt.Println(v2, v3, v1)
    var v52 uint64 = 0
    v3 -= (v3 / (((v3 % ((1) | 1))) | 1))
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
