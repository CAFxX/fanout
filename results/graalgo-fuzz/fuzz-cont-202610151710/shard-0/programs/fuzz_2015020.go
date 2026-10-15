package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint }
func genFn0(p0 int8, p1 uint64) int64 {
    v6 := GenS1{F0: p1, F1: ("foo bar" + "a")}
    var v7 int8 = (^(p0))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 int16, p1 string, p2 int8) uint16 {
    for fi8 := 0; fi8 < 3; fi8++ {
        p1 = "a"
        _ = fi8
    }
    switch 571 {
    case 982:
        v9 := []int{}
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 float32 = 1.5
            p2 -= int8(p0)
            _ = fi11
            _ = v12
        }
        v13 := GenS2{F0: p2, F1: (p0 ^ 0), F2: 32767}
        p1 = "hello"
        _ = fi10
        _ = v13
    }
    var v14 uint = 865
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 1
}

func main() {
    var v1 float32 = ((0.0 * 0.1) / (3.14159 + float32(3.14159)))
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v4 uint8 = 2
    var v5 uint8 = (^(v4))
    if (true || (65535 == 32767)) {
        var v15 uint16 = uint16(v4)
        _ = v15
    } else {
        if false {
            _ = len(v3)
        } else {
            if (1e10 == v1) {
                switch (("" + "\n\t") + "\n\t") {
                case "hello":
                    var v16 uint16 = 2
                    _ = v16
                case "\n\t":
                    fmt.Println(v5)
                }
                switch 417 {
                case 127:
                    var v17 uint = 18446744073709551615
                    _ = v17
                case 0:
                    var v18 int64 = 1
                    _ = v18
                }
                v4 += 255
            } else {
                v19 := []int{}
                fmt.Println(v1, v5)
                v20 := GenS3{F0: -32768}
                _ = v19
                _ = v20
            }
        }
        var v21 uint8 = v5
        _ = v21
    }
    v4 = 0
    if true {
        v22 := GenS4{F0: 154, F1: 4294967295, F2: 32767}
        _ = v22
    } else {
        if true {
            if true {
                for fi23 := 0; fi23 < 14; fi23++ {
                    if true {
                        switch ((("hello" + ("a" + (("\n\t" + "a") + "\n\t"))) + "a") + ("\n\t" + "")) {
                        case "\n\t":
                            v24 := GenS5{F0: 2, F1: 2, F2: 1}
                            _ = v24
                        case "":
                            if false {
                                var v25 int32 = 42
                                v26 := []int{(fi23 + (fi23 % ((32768) | 1))), 32767, fi23}
                                _ = v25
                                _ = v26
                            }
                        }
                    } else {
                        fmt.Println(v5)
                        fmt.Println(v5, v5)
                        switch "a" {
                        case "hello":
                            for fi27 := 0; fi27 < 20; fi27++ {
                                var v28 float64 = 1e10
                                _ = fi27
                                _ = v28
                            }
                        }
                    }
                    fmt.Println(v1, fi23)
                    var v29 uint8 = 2
                    _ = fi23
                    _ = v29
                }
                if false {
                    v5 = v4
                    switch "foo bar" {
                    case "":
                        v5 += (v4 ^ (v5 % ((v4) | 1)))
                    default:
                        v4 = (v5 << ((18446744073709551615) & 63))
                    }
                    var v30 bool = true
                    _ = v30
                } else {
                    v1 = 0.1
                    v31 := GenS2{F0: -128, F1: 42, F2: 32767}
                    v1 = v1
                    _ = v31
                }
            }
            switch "hello" {
            case "foo bar":
                v1 = v1
            case "\n\t":
                for fi32 := 0; fi32 < 2; fi32++ {
                    v33 := v2[4294967295]
                    _ = fi32
                    _ = v33
                }
            }
        }
    }
    switch 424 {
    case 128:
        var v34 bool = false
        _ = v34
    case 239:
        delete(v3, ("a" + ((("" + "\n\t") + ("a" + "\n\t")) + "\n\t")))
    default:
        fmt.Println(v5, v1, v4)
    }
    if (!false) {
        for fi35 := 0; fi35 < 2; fi35++ {
            v1 = ((v1 + ((float32(-2.25) / 3.14159) * (v1 - 1.5))) / 1.0)
            _ = fi35
        }
        switch 32768 {
        case 42:
            fmt.Println(v5, v1)
        default:
            if (!true) {
                _ = len(v3)
            }
        }
    } else {
        var v36 int32 = 42
        _ = v36
    }
    var v37 uint16 = uint16(v5)
    v4 += (v5 / (((^(v5))) | 1))
    var v38 uint8 = (v5 * (v4 - 127))
    fmt.Println(v38, v37, v5)
    for fi39 := 0; fi39 < 20; fi39++ {
        var v40 int64 = -237
        var v41 uint8 = v4
        switch "\n\t" {
        case "hello":
            var v42 string = ("hello" + "")
            _ = v42
        case "\n\t":
            v40 += (v40 % (((v40 * -9223372036854775808)) | 1))
        case "a":
            v38 = (v41 * (v5 % (((v5 >> ((255) & 63))) | 1)))
        }
        _ = fi39
        _ = v40
        _ = v41
    }
    if false {
        v4 += v38
    }
    if ((2147483647 != 42) && false) {
        if false {
            var v43 uint = 65535
            fmt.Println(v1, v43, v1)
            var v44 int16 = 0
            _ = v43
            _ = v44
        } else {
            v45 := GenS5{F0: 4294967295, F1: (v37 << ((874) & 63)), F2: 0}
            _ = v45
        }
    } else {
        var v46 int8 = 41
        v5 = v38
        var v47 uint32 = 2147483648
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
