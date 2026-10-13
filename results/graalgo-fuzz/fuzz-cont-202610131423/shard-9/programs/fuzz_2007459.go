package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int8; F2 int8 }
func genFn0() int16 {
    var v4 uint8 = 2
    var v5 bool = true
    v6 := []int{42, -2147483648, 0, 282}
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 bool = (-2147483648 != 1)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: (("foo bar" + "hello") + "\n\t")}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: (("hello" + "foo bar") + "foo bar"), F1: 2, F2: 415}
    if (-882 <= -474) {
        v1 = v1
        switch "a" {
        case "foo bar":
            v10 := GenS2{F0: "foo bar", F1: 65535, F2: -1}
            _ = v10
        case "\n\t":
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 int8 = 1
                var v13 uint8 = 255
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            var v14 uint = 0
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            switch "foo bar" {
            case "hello":
                var v16 uint = 849
                _ = v16
            case "":
                v17 := GenS3{F0: 0}
                _ = v17
            case "foo bar":
                v18 := GenS1{F0: ("hello" + "hello")}
                _ = v18
            default:
                var v19 int = fi15
                _ = v19
            }
            var v20 uint8 = 255
            v21 := v2[584]
            _ = fi15
            _ = v20
            _ = v21
        }
        switch 42 {
        case 32768:
            v22 := []int{127, -1}
            _ = v22
        case -1:
            switch "foo bar" {
            case "a":
                var v23 uint8 = 1
                _ = v23
            case "hello":
                fmt.Println(v1)
            }
        default:
            var v24 int16 = -32768
            _ = v24
        }
        v1 = false
    }
    if (3.14159 == (3.14159 / 1e10)) {
        var v25 int16 = -1
        switch "foo bar" {
        case "":
            switch ((("foo bar" + "foo bar") + "foo bar") + "a") {
            case "":
                v1 = (v1 || v1)
            case "a":
                v25 -= v25
            default:
                fmt.Println(v1, v1)
            }
        default:
            v3[4294967295] = "hello"
        }
        if v1 {
            if (142 != 129) {
                for fi26 := 0; fi26 < 14; fi26++ {
                    var v27 uint32 = 65535
                    var v28 int8 = 1
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v29 := GenS4{F0: 42, F1: -34, F2: -1}
                var v30 uint32 = 804
                _ = v29
                _ = v30
            } else {
                v25 += 829
                v31 := GenS3{F0: 1}
                switch "" {
                case "a":
                    v32 := []int{int(v25), -514, 4294967295}
                    _ = v32
                case "hello":
                    var v33 uint = 2
                    _ = v33
                }
                _ = v31
            }
            for fi34 := 0; fi34 < 18; fi34++ {
                switch fi34 {
                case -968:
                    if (0 >= 638) {
                        var v35 uint = 65535
                        for fi36 := 0; fi36 < 19; fi36++ {
                            var v37 uint32 = 2
                            var v38 uint32 = (^(v37))
                            fmt.Println(v38, v37, v1)
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = v35
                    } else {
                        var v39 int32 = 1
                        _ = v39
                    }
                case 807:
                    v40 := GenS2{F0: ("" + "hello"), F1: 753, F2: 42}
                    _ = v40
                case 9223372036854775807:
                    for fi41 := 0; fi41 < 11; fi41++ {
                        v25 += 800
                        var v42 uint8 = 15
                        _ = fi41
                        _ = v42
                    }
                }
                v1 = false
                _ = fi34
            }
        }
        _ = v25
    } else {
        var v43 int32 = -753
        if ((379 >= 1000) || (v1 || true)) {
            v44 := v3[521]
            var v45 string = v44
            _ = v44
            _ = v45
        }
        _ = v43
    }
    v1 = (128 >= -533)
    v1 = v1
    switch ("hello" + "\n\t") {
    case "":
        var v46 uint8 = 0
        _ = v46
    case "\n\t":
        v1 = v1
    case "hello":
        var v47 bool = (1.5 > 1.5)
        _ = v47
    default:
        for fi48 := 0; fi48 < 16; fi48++ {
            var v49 int = (fi48 - 32767)
            _ = fi48
            _ = v49
        }
    }
    v1 = v1
    var v50 uint16 = 1
    v1 = ((v50 / (((v50 & v50)) | 1)) == (v50 & (v50 >> ((840) & 63))))
    var v51 float64 = 0.0
    delete(v2, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v9
}
