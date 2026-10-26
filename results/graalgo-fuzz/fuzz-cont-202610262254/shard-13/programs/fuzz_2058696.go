package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 int8; F1 string; F2 int }
type GenS3 struct { F0 bool; F1 uint32; F2 uint }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0() uint8 {
    switch -492 {
    case 128:
        switch (((("a" + "foo bar") + "hello") + ("foo bar" + "\n\t")) + ((("a" + "foo bar") + "foo bar") + "hello")) {
        case "\n\t":
            if (true || (32767 > 1)) {
                v7 := GenS1{F0: 938, F1: 1, F2: 9223372036854775807}
                var v8 int16 = 42
                _ = v7
                _ = v8
            } else {
                var v9 int = -1
                for fi10 := 0; fi10 < 4; fi10++ {
                    v9 -= fi10
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v9 -= fi10
                        v12 := GenS2{F0: 100, F1: ("foo bar" + "foo bar"), F2: v9}
                        _ = fi11
                        _ = v12
                    }
                    fmt.Println(v9)
                    _ = fi10
                }
                v13 := GenS3{F0: false, F1: 42, F2: 65535}
                _ = v13
                _ = v9
            }
        case "foo bar":
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(fi14)
                if (122 == -1) {
                    var v15 int16 = -32768
                    v16 := []int{fi14}
                    v16 = append(v16, fi14)
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{(fi14 / (((fi14 - fi14)) | 1)), (fi14 ^ 9223372036854775807), -9223372036854775808}
                    switch fi14 {
                    case 0:
                        if false {
                            fmt.Println(fi14)
                            var v18 string = ""
                            _ = v18
                        } else {
                            fmt.Println(fi14)
                        }
                    case -423:
                        fmt.Println(fi14)
                    case -9223372036854775808:
                        if len(v17) > 0 {
                            v17[0] = (fi14 / (((fi14 - fi14)) | 1))
                        }
                    }
                    _ = v17
                }
                _ = fi14
            }
        }
    }
    var v19 uint64 = 706
    var v20 int32 = -671
    v21 := GenS1{F0: -666, F1: v19, F2: 4294967295}
    _ = v19
    _ = v20
    _ = v21
    return 255
}
func genFn1(p0 uint32, p1 uint16) uint {
    fmt.Println(p1)
    v22 := []int{2147483647, 0, 42}
    var v23 string = "a"
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 255
}

func main() {
    var v1 uint16 = 1
    v2 := []uint32{2147483647, 42}
    var v3 uint32 = 2147483647
    var v4 uint8 = 63
    v5 := []int{-2147483648, -580, -190, -345}
    var v6 float64 = 1e10
    if (!(!(42 != 127))) {
        _ = len(v5)
    }
    switch 670 {
    case 2147483647:
        fmt.Println(v3, v3)
    case 894:
        v24 := GenS2{F0: 1, F1: "a", F2: 4294967295}
        _ = v24
    }
    v4 = v4
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v1 -= uint16(v3)
    _ = len(v5)
    if (false || ("\n\t" != "hello")) {
        v1 = (v1 ^ (v1 % ((v1) | 1)))
        for fi25 := 0; fi25 < 15; fi25++ {
            for fi26 := 0; fi26 < 8; fi26++ {
                fmt.Println(v6, fi26)
                _ = fi26
            }
            v4 += v4
            _ = fi25
        }
    } else {
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 uint64 = 316
            var v29 int8 = 100
            if (!true) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch (fi27 - (fi27 | (fi27 % ((int(v29)) | 1)))) {
                case -1:
                    v5 = append(v5, fi27)
                case 1:
                    fmt.Println(v3, fi27, fi27)
                }
            } else {
                for fi30 := 0; fi30 < 8; fi30++ {
                    v3 = 2147483647
                    var v31 int32 = 1
                    v32 := GenS2{F0: (v29 - v29), F1: "\n\t", F2: fi27}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                if (((fi27 | fi27) == int(v3)) || (42 > v1)) {
                    for fi33 := 0; fi33 < 11; fi33++ {
                        var v34 uint8 = 255
                        var v35 uint = 767
                        var v36 uint = (v35 * 1)
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                }
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        var v37 string = "\n\t"
        v4 += v4
        _ = v37
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        for fi39 := 0; fi39 < 1; fi39++ {
            switch ("" + "a") {
            case "foo bar":
                var v40 int8 = -1
                _ = v40
            case "":
                v4 = v4
            }
            v3 = v3
            _ = fi39
        }
        _ = fi38
    }
    var v41 uint8 = (v4 & v4)
    var v42 int = 127
    var v43 int32 = 42
    var v44 uint = 18446744073709551615
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = (v42 & 42)
    }
    v45 := GenS4{F0: (v4 >> ((v44) & 63)), F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
