package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }

func main() {
    v1 := []int{128}
    var v2 float32 = ((float32((0.1 + -2.25)) - 3.14159) * ((3.14159 * 3.14159) / 0.1))
    var v3 int = 1
    switch "" {
    case "":
        for fi4 := 0; fi4 < 14; fi4++ {
            fmt.Println(v2, v3, v2)
            if (false && false) {
                v3 += (v3 + 9223372036854775807)
            } else {
                v5 := GenS1{F0: true, F1: (42 > 971), F2: 4294967295}
                for fi6 := 0; fi6 < 18; fi6++ {
                    fmt.Println(fi6, fi6)
                    switch 32767 {
                    case 128:
                        v7 := GenS1{F0: (!true), F1: (v2 > (v2 - v2)), F2: 2}
                        _ = v7
                    default:
                        var v8 bool = true
                        _ = v8
                    }
                    _ = fi6
                }
                fmt.Println(fi4, fi4)
                _ = v5
            }
            v3 = (fi4 % ((363) | 1))
            _ = fi4
        }
    case "a":
        v3 = 32767
    case "foo bar":
        var v9 int16 = 0
        _ = v9
    }
    switch ("a" + (("" + "") + ((("foo bar" + ("foo bar" + "hello")) + "\n\t") + ("foo bar" + ("foo bar" + "a"))))) {
    case "hello":
        switch (v3 << ((1) & 63)) {
        case 0:
            var v10 uint64 = 18446744073709551615
            _ = v10
        default:
            if (!true) {
                for fi11 := 0; fi11 < 19; fi11++ {
                    fmt.Println(fi11, v3)
                    _ = fi11
                }
                v2 = 3.14159
                if ((!false) || ((("\n\t" + "\n\t") + ("a" + "\n\t")) <= "")) {
                    _ = len(v1)
                }
            }
        }
    case "":
        var v12 int16 = -1
        _ = v12
    case "foo bar":
        v13 := GenS2{F0: 65535, F1: 42}
        _ = v13
    default:
        var v14 uint8 = 7
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 float64 = -2.25
        var v17 bool = false
        var v18 float64 = v16
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    if ((("hello" + "a") + ("hello" + "foo bar")) <= ((("a" + ((("hello" + ("\n\t" + "")) + "foo bar") + "foo bar")) + ("" + "foo bar")) + "")) {
        var v19 int = (v3 % (((v3 & v3)) | 1))
        _ = v19
    } else {
        v1 = append(v1, (v3 ^ -9223372036854775808))
    }
    v2 = v2
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint = 65535
        switch "hello" {
        case "a":
            switch v3 {
            case 0:
                for fi22 := 0; fi22 < 19; fi22++ {
                    switch v3 {
                    case -1:
                        v23 := GenS1{F0: false, F1: true, F2: (v21 >> (((v21 & v21)) & 63))}
                        _ = v23
                    case 127:
                        v3 += (-(fi20))
                    case 637:
                        var v24 uint = 18446744073709551615
                        _ = v24
                    default:
                        v25 := GenS3{F0: 779, F1: 343}
                        _ = v25
                    }
                    _ = fi22
                }
            case -2147483648:
                v1 = append(v1, (-(fi20)))
            case 9223372036854775807:
                var v26 bool = false
                _ = v26
            }
        case "foo bar":
            v2 += ((v2 * v2) + v2)
        }
        _ = fi20
        _ = v21
    }
    var v27 uint32 = 864
    var v28 int64 = -1
    switch (v3 << ((255) & 63)) {
    case 2147483647:
        if ((true || true) || (901 > 0)) {
            v1 = append(v1, v3)
            switch "" {
            case "":
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 uint = 42
                    _ = fi29
                    _ = v30
                }
            case "hello":
                if (32767 < -202) {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        var v32 int = (-(fi31))
                        _ = fi31
                        _ = v32
                    }
                }
            }
            _ = len(v1)
        }
    }
    switch v3 {
    case 193:
        if (!true) {
            var v33 uint8 = 42
            v27 += v27
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        var v35 int16 = 42
        if false {
            v35 = v35
            var v36 int8 = 104
            _ = v36
        } else {
            var v37 int = (^(fi34))
            var v38 int64 = (v28 & int64(v27))
            _ = v37
            _ = v38
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
