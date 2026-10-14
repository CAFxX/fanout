package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 int8 }
type GenS4 struct { F0 uint16; F1 int64; F2 int }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int32; F1 int16; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 65535
    v3 := make(map[int]uint32)
    if true {
        var v4 bool = true
        _ = v4
    }
    var v5 float64 = float64(0.0)
    switch (("hello" + "") + "") {
    case "a":
        fmt.Println(v5)
    default:
        if (1e10 > 1.5) {
            v5 = (0.0 * v5)
        }
    }
    v6 := GenS1{F0: 1, F1: 0}
    fmt.Println(v5, v2)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 bool = ((false && (1 > 1)) || false)
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int16 = 720
            switch "foo bar" {
            case "a":
                var v11 int8 = 90
                _ = v11
            case "foo bar":
                if v8 {
                    v5 -= -2.25
                    if (v8 && (v10 <= -568)) {
                        var v12 bool = ("\n\t" == ("a" + "hello"))
                        _ = v12
                    }
                }
            }
            fmt.Println(fi7)
            _ = fi9
            _ = v10
        }
        switch (fi7 >> ((2) & 63)) {
        case 128:
            v13 := GenS2{F0: 65535, F1: -1}
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    v5 = 0.0
    v14 := GenS3{F0: (v5 != float64((1e10 - float32((3.14159 + (v5 - (float64(-2.25) / 3.14159))))))), F1: 42, F2: 100}
    v15 := GenS3{F0: false, F1: 42, F2: -81}
    var v16 string = "foo bar"
    for fi17 := 0; fi17 < 7; fi17++ {
        if (!true) {
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := []int{0, (fi17 / ((2147483647) | 1))}
                switch ("foo bar" + v16) {
                case "":
                    var v20 float32 = float32(v5)
                    _ = v20
                case "foo bar":
                    if (!(((v16 + (v16 + "\n\t")) + "\n\t") <= ((v16 + (v16 + v16)) + "\n\t"))) {
                        switch (-(fi17)) {
                        case 32768:
                            v5 = 1e10
                        case 1:
                            var v21 float32 = (float32(v5) * float32(float64(-2.25)))
                            _ = v21
                        case 266:
                            if true {
                                var v22 uint64 = 42
                                _ = v22
                            }
                        default:
                            switch v16 {
                            case "":
                                fmt.Println(fi18)
                            case "foo bar":
                                fmt.Println(v16)
                            default:
                                if true {
                                    v16 = v16
                                    if false {
                                        delete(v3, fi18)
                                        switch v16 {
                                        case "a":
                                            v5 += 1.5
                                        case "hello":
                                            v23 := v3[(fi18 * (fi17 | fi17))]
                                            _ = v23
                                        }
                                    } else {
                                        v5 -= v5
                                        var v24 uint16 = 255
                                        var v25 uint16 = v24
                                        _ = v24
                                        _ = v25
                                    }
                                } else {
                                    var v26 int = (fi18 % ((fi17) | 1))
                                    var v27 uint64 = 0
                                    var v28 bool = (true || true)
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            }
                        }
                        var v29 string = (("a" + (v16 + "hello")) + v16)
                        _ = v29
                    }
                }
                v30 := GenS1{F0: 1, F1: 890}
                _ = fi18
                _ = v19
                _ = v30
            }
            for fi31 := 0; fi31 < 17; fi31++ {
                v3[32768] = v2
                fmt.Println(v5, v16)
                _ = fi31
            }
            var v32 uint16 = 1
            _ = v32
        } else {
            fmt.Println(v16, fi17, fi17)
            var v33 int32 = 0
            _ = v33
        }
        v34 := GenS1{F0: -1, F1: 1}
        _ = fi17
        _ = v34
    }
    v35 := []int{-1, 32768}
    if false {
        if (((!(716 >= (v2 - (v2 | v2)))) || true) && true) {
            for fi36 := 0; fi36 < 12; fi36++ {
                fmt.Println(fi36)
                _ = fi36
            }
            if false {
                var v37 int16 = 1
                v5 = (((v5 + ((0.0 + (v5 * -2.25)) + ((float64((0.1 * -2.25)) - v5) + 1.5))) / (float64(1e10) / v5)) + 0.0)
                _ = v37
            }
        }
        v16 = ((v16 + v16) + v16)
        v3[-316] = 65535
    } else {
        v16 = "foo bar"
    }
    fmt.Println(v16, v2, v5)
    for fi38 := 0; fi38 < 19; fi38++ {
        v16 = v16
        _ = fi38
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        if ("" != "\n\t") {
            v16 = v16
            if true {
                fmt.Println(v2)
                v40 := GenS3{F0: false, F1: 1, F2: 100}
                _ = v40
            }
            v16 = (v16 + "\n\t")
        }
        v41 := GenS4{F0: 42, F1: 9223372036854775807, F2: fi39}
        v42 := GenS1{F0: 959, F1: 678}
        _ = fi39
        _ = v41
        _ = v42
    }
    v43 := GenS5{F0: false, F1: 59}
    for fi44 := 0; fi44 < 2; fi44++ {
        v45 := GenS6{F0: -2147483648, F1: -716, F2: v2}
        v46 := v3[fi44]
        if true {
            if ((-2.25 + (((0.1 * (((3.14159 - (3.14159 + 3.14159)) * 1e10) / (0.0 - -2.25))) / -2.25) * 1.5)) >= (((-2.25 / 1.5) - (0.0 * 0.1)) + 3.14159)) {
                v5 = ((float64((((1e10 * 1e10) * ((-2.25 + 0.0) - -2.25)) * -2.25)) + (v5 * v5)) * (v5 / v5))
                fmt.Println(v46)
                switch 2147483647 {
                case 213:
                    v5 += (v5 * (3.14159 / v5))
                }
            }
            v2 += v46
        } else {
            v5 = 0.1
            var v47 uint8 = 176
            v48 := GenS4{F0: 65535, F1: -1, F2: 77}
            _ = v47
            _ = v48
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v43
    _ = v5
    _ = v6
}
