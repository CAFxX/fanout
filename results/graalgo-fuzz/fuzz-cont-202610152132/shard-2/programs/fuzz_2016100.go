package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64; F1 int; F2 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0() int {
    var v4 float64 = 3.14159
    var v5 string = "foo bar"
    var v6 int16 = 42
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-2147483648, -261}
    v3 := []string{"a", "\n\t", "", ""}
    var v7 int32 = -1
    delete(v1, 200)
    fmt.Println(v7)
    var v8 int = 669
    for fi9 := 0; fi9 < 4; fi9++ {
        if ((true && ((fi9 != (v8 - -1)) && false)) || true) {
            v8 = fi9
            fmt.Println(v8)
            v2 = append(v2, (-(fi9)))
        }
        _ = fi9
    }
    v8 = (v8 << ((98) & 63))
    switch "foo bar" {
    case "hello":
        switch v8 {
        case -701:
            switch ("a" + "a") {
            case "hello":
                v2 = append(v2, genFn0())
            }
        case 0:
            for fi10 := 0; fi10 < 12; fi10++ {
                if false {
                    if (false && true) {
                        v8 = (v8 - (fi10 & (^(fi10))))
                        v7 = -2147483648
                        var v11 uint = 2
                        _ = v11
                    }
                    if true {
                        _ = len(v1)
                    }
                }
                var v12 uint16 = 255
                v13 := GenS1{F0: (-(v7)), F1: (v12 % (((v12 + v12)) | 1))}
                _ = fi10
                _ = v12
                _ = v13
            }
        case 9223372036854775807:
            v8 = (v8 + v8)
        }
    case "":
        v14 := GenS1{F0: -730, F1: 284}
        _ = v14
    case "foo bar":
        for fi15 := 0; fi15 < 17; fi15++ {
            v7 += v7
            v8 = (v8 * v8)
            _ = fi15
        }
    }
    var v16 float64 = (0.0 / 1.0)
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v16, v16)
        switch fi17 {
        case 32768:
            if ((0.1 / v16) <= v16) {
                var v18 string = ("\n\t" + "\n\t")
                var v19 int64 = 42
                v3 = append(v3, v18)
                _ = v18
                _ = v19
            } else {
                var v20 uint32 = 2147483647
                _ = v20
            }
        case 186:
            fmt.Println(v7, v8, v7)
        case 719:
            v7 -= 0
        default:
            if (127 != 128) {
                v21 := GenS1{F0: -1, F1: 42}
                _ = v21
            }
        }
        var v22 uint64 = 4294967295
        _ = fi17
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            fmt.Println(v8, v16)
            var v25 uint32 = 618
            switch "foo bar" {
            case "hello":
                if (-1 <= 0) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v26 int32 = v7
                _ = v26
            }
            _ = len(v1)
            var v27 int = (fi24 ^ (-(fi24)))
            _ = fi24
            _ = v25
            _ = v27
        }
        v28 := GenS1{F0: (v7 & (v7 & (v7 | (v7 | (v7 ^ v7))))), F1: 32767}
        switch "foo bar" {
        case "a":
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 int64 = 1
                fmt.Println(v30, v30)
                v8 += fi29
                switch (fi23 ^ (fi29 * fi29)) {
                case 1:
                    if ((fi29 >> ((255) & 63)) != (fi23 >> ((18446744073709551615) & 63))) {
                        var v31 int8 = -115
                        _ = v31
                    }
                case 778:
                    v32 := GenS2{F0: 820, F1: genFn0(), F2: 2147483648}
                    _ = v32
                case 42:
                    switch "hello" {
                    case "\n\t":
                        v33 := GenS2{F0: 966, F1: genFn0(), F2: 0}
                        _ = v33
                    case "hello":
                        switch ("foo bar" + (("\n\t" + "") + "foo bar")) {
                        case "foo bar":
                            _ = len(v1)
                        case "\n\t":
                            var v34 int64 = v30
                            _ = v34
                        case "a":
                            fmt.Println(fi29)
                        default:
                            v16 = v16
                        }
                        v35 := GenS3{F0: 108}
                        _ = v35
                    }
                }
                _ = fi29
                _ = v30
            }
        case "":
            v16 += 0.1
        case "foo bar":
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v36 := GenS3{F0: 4294967295}
                _ = v36
            } else {
                var v37 bool = true
                var v38 bool = (false || false)
                _ = v37
                _ = v38
            }
        default:
            switch "\n\t" {
            case "":
                for fi39 := 0; fi39 < 13; fi39++ {
                    if true {
                        v16 += v16
                        var v40 string = ("" + "\n\t")
                        _ = v40
                    }
                    _ = fi39
                }
            case "hello":
                v7 = (v7 ^ (v7 << ((18446744073709551615) & 63)))
            }
        }
        _ = fi23
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
