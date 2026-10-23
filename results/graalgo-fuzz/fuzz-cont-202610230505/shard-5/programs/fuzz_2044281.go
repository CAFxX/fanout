package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 string; F2 uint16 }
type GenS8 struct { F0 uint32; F1 int16 }
func genFn0(p0 int32, p1 int) uint16 {
    var v4 float64 = 0.0
    v5 := GenS1{F0: 63}
    if (0 <= 2147483648) {
        v6 := []int{(p1 & 128), 1, (p1 / (((p1 & p1)) | 1)), (p1 / ((9223372036854775807) | 1))}
        v7 := GenS1{F0: 1000}
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v8
    return 255
}
func genFn1(p0 int16) string {
    if false {
        var v9 int64 = 1
        _ = v9
    }
    switch "\n\t" {
    case "foo bar":
        var v10 int8 = 42
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                fmt.Println(fi11, fi12, fi12)
                var v13 string = "foo bar"
                _ = fi12
                _ = v13
            }
            p0 -= -32768
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 string = ""
                var v16 float32 = 1e10
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi11
        }
    }
    v17 := []int{9223372036854775807}
    _ = p0
    _ = v17
    return "a"
}

func main() {
    var v1 uint8 = 132
    var v2 uint16 = 2
    v3 := make(map[string]int)
    switch "" {
    case "\n\t":
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := GenS1{F0: int16(fi18)}
            v20 := GenS2{F0: v1, F1: 864}
            var v21 int = int(v2)
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    case "hello":
        v22 := GenS3{F0: 0, F1: 9223372036854775807}
        _ = v22
    case "a":
        var v23 uint = 18446744073709551615
        _ = v23
    }
    v1 -= v1
    var v24 uint16 = 1
    if (((false && (false || true)) && (42 != -32768)) && (0 >= -9223372036854775808)) {
        var v25 uint8 = v1
        v24 = v2
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        if (1.5 == (0.0 + 1.5)) {
            v27 := GenS4{F0: v1, F1: -9223372036854775808}
            switch fi26 {
            case -1:
                v1 -= 1
            case 32767:
                v28 := GenS4{F0: v1, F1: -1}
                _ = v28
            default:
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    var v29 uint8 = 42
                    _ = v29
                case "hello":
                    var v30 uint = 437
                    _ = v30
                case "\n\t":
                    v24 += v2
                default:
                    _ = len(v3)
                }
            }
            v31 := GenS5{F0: v1}
            _ = v27
            _ = v31
        }
        var v32 string = ""
        v32 = ""
        _ = fi26
        _ = v32
    }
    if true {
        var v33 int8 = 0
        v34 := GenS5{F0: (v1 % ((2) | 1))}
        _ = v33
        _ = v34
    }
    v35 := []int{280, 2147483647, 1, 1}
    var v36 uint32 = 352
    if (0 <= 42) {
        v36 = 1
        v36 += v36
        switch "\n\t" {
        case "hello":
            v24 = (v2 % (((v24 * genFn0(0, 9223372036854775807))) | 1))
        }
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        for fi38 := 0; fi38 < 7; fi38++ {
            v2 = v24
            _ = fi38
        }
        _ = fi37
    }
    v39 := v3[(("foo bar" + "\n\t") + ((("hello" + "hello") + "hello") + ""))]
    v40 := GenS6{F0: false}
    fmt.Println(v24, v1, v36)
    switch v39 {
    case 32767:
        if false {
            v41 := v3["hello"]
            _ = len(v35)
            _ = v41
        } else {
            if (false && ((("" + "hello") != "") || false)) {
                var v42 string = (("" + ("\n\t" + "a")) + "\n\t")
                _ = v42
            } else {
                switch (v39 & v39) {
                case 9223372036854775807:
                    v2 = (v2 ^ v24)
                default:
                    switch -1 {
                    case 347:
                        for fi43 := 0; fi43 < 5; fi43++ {
                            for fi44 := 0; fi44 < 3; fi44++ {
                                v45 := GenS7{F0: 2147483647, F1: "foo bar", F2: 2}
                                _ = fi44
                                _ = v45
                            }
                            _ = fi43
                        }
                    case 42:
                        v46 := v3["a"]
                        _ = v46
                    case 0:
                        var v47 int64 = -594
                        _ = v47
                    default:
                        _ = len(v35)
                    }
                }
                if false {
                    if (false || (!false)) {
                        for fi48 := 0; fi48 < 4; fi48++ {
                            var v49 uint32 = (v36 & v36)
                            v2 = (v2 & (v24 ^ 2))
                            _ = fi48
                            _ = v49
                        }
                        if (4294967295 > 0) {
                            v50 := GenS5{F0: v1}
                            _ = v50
                        }
                        switch (v39 + v39) {
                        case 2147483647:
                            v51 := GenS5{F0: v1}
                            _ = v51
                        }
                    } else {
                        var v52 uint64 = 4294967295
                        if (!(42 == -9223372036854775808)) {
                            var v53 float32 = float32(1e10)
                            _ = v53
                        } else {
                            var v54 uint32 = (v36 | v36)
                            var v55 int8 = int8(v36)
                            v56 := GenS8{F0: v36, F1: 905}
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        switch v39 {
                        case -2147483648:
                            v52 = 960
                        }
                        _ = v52
                    }
                }
            }
        }
    case 826:
        for fi57 := 0; fi57 < 14; fi57++ {
            switch (v39 | 9223372036854775807) {
            case 4294967295:
                v58 := GenS3{F0: 0, F1: v39}
                _ = v58
            case 2147483647:
                fmt.Println(v2, v2)
            case 516:
                switch (("\n\t" + "hello") + ("\n\t" + "")) {
                case "":
                    fmt.Println(v1, v1)
                case "\n\t":
                    var v59 int16 = 42
                    _ = v59
                case "hello":
                    for fi60 := 0; fi60 < 7; fi60++ {
                        if (1 == 42) {
                            var v61 int = 4294967295
                            v39 += (fi60 >> ((65535) & 63))
                            switch 807 {
                            case 42:
                                fmt.Println(fi60, v39, v1)
                            case -9223372036854775808:
                                if true {
                                    v24 += (v2 / ((genFn0(-1, v39)) | 1))
                                    if false {
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                    } else {
                                        var v62 int64 = -289
                                        _ = v62
                                    }
                                }
                            default:
                                v3[""] = v39
                            }
                            _ = v61
                        }
                        _ = fi60
                    }
                }
            }
            v1 = v1
            var v63 int16 = 1
            _ = fi57
            _ = v63
        }
    default:
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v40
}
