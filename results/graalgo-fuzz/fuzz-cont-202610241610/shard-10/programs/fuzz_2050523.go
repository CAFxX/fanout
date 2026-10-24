package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0() uint16 {
    switch 128 {
    case 1:
        for fi5 := 0; fi5 < 16; fi5++ {
            switch "hello" {
            case "a":
                var v6 int8 = 1
                _ = v6
            }
            if true {
                var v7 float32 = 1e10
                for fi8 := 0; fi8 < 10; fi8++ {
                    v9 := []int{9223372036854775807, fi5, (fi8 % (((fi5 - fi8)) | 1))}
                    v10 := GenS1{F0: 4294967295}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                for fi11 := 0; fi11 < 4; fi11++ {
                    if (false && true) {
                        var v12 uint64 = 42
                        _ = v12
                    }
                    v13 := GenS1{F0: -471}
                    _ = fi11
                    _ = v13
                }
                _ = v7
            } else {
                fmt.Println(fi5)
                var v14 bool = (42 > 0)
                _ = v14
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 int = fi15
                var v17 uint8 = 128
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi5
        }
    case -211:
        if (!(3.14159 <= 3.14159)) {
            var v18 int16 = 32767
            var v19 string = ("\n\t" + ((("a" + "foo bar") + ("hello" + "")) + "\n\t"))
            _ = v18
            _ = v19
        } else {
            if true {
                v20 := []int{32768, -647, 128}
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint32 = 1
                    var v23 int64 = -364
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v20
            } else {
                var v24 string = "\n\t"
                v24 = "a"
                _ = v24
            }
            var v25 int8 = 1
            _ = v25
        }
    case 552:
        var v26 uint8 = 228
        _ = v26
    default:
        v27 := []int{-2147483648, 42, 4294967295, 9223372036854775807}
        _ = v27
    }
    var v28 int8 = -1
    v28 = v28
    _ = v28
    return 42
}

func main() {
    v1 := []uint32{42}
    v2 := make(map[uint32]uint32)
    var v3 int32 = 971
    v4 := make(map[uint32]bool)
    var v29 int64 = 330
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 int = (fi31 >> ((305) & 63))
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 float64 = (1.5 * float64(float32(0.0)))
            if (v34 != 0.1) {
                var v35 uint = 1
                _ = v35
            } else {
                v34 = ((((v34 * (float64(float32(3.14159)) * (v34 / 0.1))) / v34) * ((v34 / v34) * -2.25)) + (v34 / 1.5))
            }
            v36 := GenS2{F0: -1, F1: "hello"}
            _ = fi33
            _ = v34
            _ = v36
        }
        _ = fi30
    }
    var v37 uint = 18446744073709551615
    var v38 uint32 = 2147483648
    if true {
        if (435 != 9223372036854775807) {
            v37 = (v37 + (v37 + (v37 / ((195) | 1))))
        }
        v29 += (v29 / ((v29) | 1))
    } else {
        for fi39 := 0; fi39 < 9; fi39++ {
            v29 = (^(v29))
            if false {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v1 = append(v1, v38)
                } else {
                    if len(v1) > 0 {
                        v1[0] = (v38 * (v38 % ((v38) | 1)))
                    }
                }
            }
            if (2 != genFn0()) {
                switch ("foo bar" + (("\n\t" + ("\n\t" + (("a" + "hello") + "a"))) + "\n\t")) {
                case "\n\t":
                    var v40 string = ("hello" + "\n\t")
                    _ = v40
                }
                fmt.Println(v3, v38, v29)
            }
            _ = fi39
        }
        v38 -= v38
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        var v42 uint = v37
        if ((v29 | int64(fi41)) < (v29 / ((9223372036854775807) | 1))) {
            v1 = append(v1, v38)
            v43 := v4[v38]
            _ = v43
        }
        var v44 float64 = (1e10 * ((0.1 + (0.0 / 1.0)) * (float64(3.14159) * (((0.1 / 1.5) / ((float64(0.0) + ((1e10 + (0.1 - 3.14159)) - 0.1)) / 1.5)) * 1.5))))
        _ = fi41
        _ = v42
        _ = v44
    }
    if (((v29 * (v29 >> (((v37 * (v37 % ((v37) | 1)))) & 63))) > (-(v29))) || true) {
        v45 := GenS1{F0: 1}
        switch ("\n\t" + "hello") {
        case "foo bar":
            var v46 bool = (486 > 42)
            _ = v46
        case "\n\t":
            v47 := GenS2{F0: 0, F1: "\n\t"}
            _ = v47
        case "a":
            switch -488 {
            case 128:
                v48 := GenS2{F0: -1, F1: "hello"}
                _ = v48
            case 127:
                var v49 uint = 65535
                _ = v49
            }
        }
        _ = v45
    }
    var v50 int64 = v29
    for fi51 := 0; fi51 < 6; fi51++ {
        var v52 string = "a"
        fmt.Println(v3, v52, v38)
        switch int(v37) {
        case 1:
            for fi53 := 0; fi53 < 4; fi53++ {
                v54 := GenS3{F0: "a"}
                fmt.Println(v52, v37, v52)
                var v55 string = "hello"
                v56 := GenS4{F0: 0, F1: 0, F2: v38}
                _ = fi53
                _ = v54
                _ = v55
                _ = v56
            }
        }
        _ = fi51
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v50
}
