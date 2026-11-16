package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 bool; F2 uint8 }
type GenS5 struct { F0 int8; F1 int }
func genFn0(p0 int, p1 uint32, p2 int16) int64 {
    var v5 int32 = 2147483647
    switch (p0 / ((-1) | 1)) {
    case 250:
        var v6 int = 983
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -264
}
func genFn1() uint64 {
    var v7 float64 = float64(float32((1.5 * 3.14159)))
    var v8 uint32 = 0
    if false {
        v7 = v7
        for fi9 := 0; fi9 < 13; fi9++ {
            switch "foo bar" {
            case "foo bar":
                v10 := []int{(fi9 << ((65535) & 63)), int(v8), 127, (fi9 * fi9)}
                _ = v10
            case "\n\t":
                v11 := GenS1{F0: 187, F1: 944, F2: "\n\t"}
                _ = v11
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := []int{128}
                switch (fi9 << ((4294967295) & 63)) {
                case 1:
                    fmt.Println(fi12, v7, fi12)
                }
                var v14 string = "foo bar"
                _ = fi12
                _ = v13
                _ = v14
            }
            v8 -= v8
            _ = fi9
        }
        if true {
            if (false || (65535 >= 255)) {
                v15 := GenS1{F0: 42, F1: -1, F2: "\n\t"}
                v7 = v7
                v16 := GenS1{F0: 181, F1: 4294967295, F2: "foo bar"}
                _ = v15
                _ = v16
            } else {
                v17 := []int{-768}
                switch ("" + "\n\t") {
                case "":
                    fmt.Println(v8, v7)
                case "foo bar":
                    fmt.Println(v8, v8)
                }
                _ = v17
            }
            var v18 uint16 = 0
            _ = v18
        }
    } else {
        if false {
            v19 := GenS1{F0: 128, F1: 637, F2: "foo bar"}
            _ = v19
        } else {
            if ((-2147483648 <= 977) || (-9223372036854775808 < genFn0(1, v8, 1))) {
                fmt.Println(v8, v7)
                var v20 float32 = 3.14159
                v21 := GenS2{F0: 42, F1: -69, F2: -613}
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS1{F0: 2, F1: fi22, F2: "a"}
                var v24 int = fi22
                if false {
                    for fi25 := 0; fi25 < 16; fi25++ {
                        var v26 int = (v24 - -2147483648)
                        v7 = (0.1 + v7)
                        _ = fi25
                        _ = v26
                    }
                    v7 = v7
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    }
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint32 = 1
    v2 := []uint32{2147483648, (^(v1))}
    var v3 int = 127
    var v4 uint = 2
    if (!(255 != 14)) {
        v3 += (v3 - -1)
        switch ((("a" + (("hello" + "a") + "a")) + "foo bar") + "hello") {
        case "hello":
            _ = len(v2)
        }
    } else {
        fmt.Println(v1)
        var v27 string = ""
        _ = v27
    }
    if (!false) {
        v3 += v3
        var v28 uint8 = 42
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        if false {
            var v30 uint64 = genFn1()
            _ = v30
        } else {
            var v31 float64 = (1.5 - ((((float64(-2.25) - float64(0.1)) - (((3.14159 / 1e10) * 1e10) * 1e10)) + 3.14159) + ((((float64(-2.25) + (3.14159 / 0.1)) * 1e10) * -2.25) * (1e10 / float64(-2.25)))))
            _ = v31
        }
        var v32 string = "foo bar"
        var v33 float64 = 1.5
        _ = fi29
        _ = v32
        _ = v33
    }
    var v34 float64 = (((3.14159 / 1.0) / ((0.0 - ((float64(1.5) / ((float64(0.1) - 1e10) * float64((1.5 - float32(3.14159))))) * -2.25)) / 1e10)) + 1e10)
    var v35 uint8 = 128
    v34 = 1e10
    var v36 int32 = -2147483648
    v4 += 42
    for fi37 := 0; fi37 < 4; fi37++ {
        v38 := GenS3{F0: true}
        for fi39 := 0; fi39 < 20; fi39++ {
            var v40 int16 = 32767
            _ = fi39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    for fi41 := 0; fi41 < 11; fi41++ {
        v42 := GenS2{F0: 1, F1: int(v35), F2: 31}
        _ = fi41
        _ = v42
    }
    var v43 uint8 = 128
    switch (v3 % ((v3) | 1)) {
    case 4294967295:
        v44 := GenS4{F0: 1, F1: (float32(v34) == (0.1 - (0.0 * (0.1 * 1.5)))), F2: v43}
        _ = v44
    case -9223372036854775808:
        var v45 int = (v3 / (((v3 | v3)) | 1))
        _ = v45
    default:
        if (1 >= 688) {
            fmt.Println(v4, v35)
            if true {
                v46 := GenS5{F0: 1, F1: v3}
                v35 += (v35 >> ((v4) & 63))
                _ = v46
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v43
}
