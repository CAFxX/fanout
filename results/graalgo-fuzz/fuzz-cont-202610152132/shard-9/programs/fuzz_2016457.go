package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint; F2 int }
type GenS4 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]string)
    var v3 int8 = 1
    for fi4 := 0; fi4 < 16; fi4++ {
        if ((!(42 <= 1)) && true) {
            v1 = v1
        }
        _ = fi4
    }
    v5 := GenS1{F0: 128, F1: 42}
    v6 := GenS1{F0: -2147483648, F1: 423}
    v1 = v1
    if (!false) {
        v7 := GenS1{F0: 128, F1: 1000}
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(v3, fi9)
                _ = fi9
            }
            fmt.Println(v1, fi8)
            _ = fi8
        }
        _ = v7
    }
    var v10 int16 = 0
    if (true || (("hello" + ((("\n\t" + "\n\t") + "") + ("hello" + ((("foo bar" + "hello") + "hello") + "")))) < "")) {
        var v11 uint = 834
        if (false && true) {
            v12 := GenS2{F0: 42}
            _ = v12
        }
        _ = v11
    } else {
        v13 := GenS1{F0: -2147483648, F1: -42}
        _ = v13
    }
    var v14 int16 = v10
    var v15 uint = 0
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(v3)
            _ = fi17
        }
        v3 = v3
        delete(v2, fi16)
        _ = fi16
    }
    switch 685 {
    case 642:
        switch "" {
        case "foo bar":
            var v18 int16 = 32767
            _ = v18
        case "hello":
            v3 -= v3
        default:
            for fi19 := 0; fi19 < 15; fi19++ {
                v1 += (v1 - v1)
                v20 := []int{}
                _ = fi19
                _ = v20
            }
        }
    case 611:
        v21 := GenS3{F0: 184, F1: v15, F2: 0}
        _ = v21
    case -2147483648:
        if (0.1 <= (((float32(1e10) * 1.5) + 3.14159) / 3.14159)) {
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 uint32 = 42
                _ = fi22
                _ = v23
            }
        }
    default:
        v15 = (v15 >> ((18446744073709551615) & 63))
    }
    v15 = v15
    if ("" > "a") {
        v24 := GenS4{F0: 226, F1: 1, F2: 42}
        for fi25 := 0; fi25 < 7; fi25++ {
            for fi26 := 0; fi26 < 20; fi26++ {
                if (!(true || true)) {
                    var v27 int = fi25
                    v28 := GenS3{F0: 1, F1: 61, F2: -573}
                    _ = v27
                    _ = v28
                }
                v15 = v15
                _ = fi26
            }
            v29 := GenS5{F0: v3}
            _ = fi25
            _ = v29
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            switch (fi30 / ((794) | 1)) {
            case -2147483648:
                fmt.Println(v3, v14)
            case 9223372036854775807:
                var v31 uint32 = 525
                _ = v31
            case 127:
                fmt.Println(v15, v3, v1)
            }
            _ = fi30
        }
        _ = v24
    } else {
        for fi32 := 0; fi32 < 15; fi32++ {
            for fi33 := 0; fi33 < 11; fi33++ {
                v10 += (v10 + v14)
                var v34 uint16 = v1
                _ = fi33
                _ = v34
            }
            fmt.Println(fi32, v3, v1)
            switch fi32 {
            case 2147483647:
                fmt.Println(fi32, v3, v3)
            case -784:
                var v35 bool = (!true)
                _ = v35
            }
            _ = fi32
        }
        var v36 int32 = 1
        _ = v36
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        switch "hello" {
        case "foo bar":
            var v38 uint8 = 68
            _ = v38
        case "hello":
            if (!false) {
                for fi39 := 0; fi39 < 16; fi39++ {
                    v40 := GenS3{F0: 226, F1: v15, F2: 127}
                    _ = fi39
                    _ = v40
                }
                if (127 < 2) {
                    v15 -= 0
                }
                fmt.Println(v14, fi37, v1)
            }
        case "":
            var v41 string = "hello"
            _ = v41
        default:
            var v42 int64 = 42
            _ = v42
        }
        switch "foo bar" {
        case "":
            var v43 uint8 = 188
            _ = v43
        case "\n\t":
            if (2 != 4294967295) {
                var v44 float64 = -2.25
                if (42 < -1) {
                    v45 := GenS2{F0: 0}
                    v15 = v15
                    v46 := GenS1{F0: 353, F1: v10}
                    _ = v45
                    _ = v46
                } else {
                    v3 = (v3 * (v3 & 37))
                }
                _ = v44
            }
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
