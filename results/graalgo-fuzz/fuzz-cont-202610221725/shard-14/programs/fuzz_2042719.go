package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16; F1 int64; F2 uint }
type GenS4 struct { F0 uint32; F1 bool; F2 bool }

func main() {
    var v1 uint16 = 32767
    var v2 int64 = -114
    var v3 int16 = -32768
    var v4 int16 = (^(v3))
    fmt.Println(v3)
    fmt.Println(v1, v2)
    v4 = (v3 - v4)
    v5 := GenS1{F0: 18446744073709551615}
    switch ("\n\t" + ("" + "foo bar")) {
    case "a":
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 string = "foo bar"
            v8 := []int{42, 2147483647, (fi6 ^ fi6)}
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 bool = (4294967295 >= 65535)
            _ = fi9
            _ = v10
        }
    }
    var v11 int8 = 78
    for fi12 := 0; fi12 < 10; fi12++ {
        switch ("" + ("foo bar" + "\n\t")) {
        case "hello":
            fmt.Println(v1, fi12, fi12)
        case "":
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 uint8 = 128
                _ = fi13
                _ = v14
            }
        case "foo bar":
            v15 := GenS2{F0: 432, F1: 255}
            _ = v15
        default:
            switch fi12 {
            case 9223372036854775807:
                fmt.Println(v4, v2, v1)
            case -9223372036854775808:
                var v16 int8 = v11
                _ = v16
            default:
                switch -1 {
                case -273:
                    v17 := GenS2{F0: 42, F1: 4294967295}
                    _ = v17
                case 4294967295:
                    var v18 int = fi12
                    _ = v18
                case 42:
                    var v19 uint = 1
                    _ = v19
                }
            }
        }
        _ = fi12
    }
    if ((((-2.25 * (1.5 / 3.14159)) <= -2.25) && ("" > ("" + "\n\t"))) || true) {
        v2 += (v2 - 554)
    } else {
        fmt.Println(v11, v2, v1)
    }
    if (1 == -608) {
        switch "\n\t" {
        case "foo bar":
            v4 -= v4
        case "a":
            v20 := GenS2{F0: 1, F1: 1}
            _ = v20
        case "\n\t":
            var v21 uint64 = 18446744073709551615
            _ = v21
        }
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 uint16 = v1
            v3 = int16(v23)
            _ = fi22
            _ = v23
        }
    } else {
        switch "\n\t" {
        case "hello":
            for fi24 := 0; fi24 < 12; fi24++ {
                v2 = -9223372036854775808
                v11 += -49
                _ = fi24
            }
        case "foo bar":
            for fi25 := 0; fi25 < 12; fi25++ {
                switch (("" + "a") + "a") {
                case "":
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v1 = (v1 * v1)
                        if true {
                            v27 := GenS3{F0: v3, F1: 1, F2: 2}
                            if (397 >= 4294967295) {
                                var v28 int16 = (v3 >> ((uint(v1)) & 63))
                                v29 := GenS3{F0: -32768, F1: (v2 | (v2 & 0)), F2: 2}
                                _ = v28
                                _ = v29
                            }
                            switch "a" {
                            case "foo bar":
                                v3 -= v3
                            default:
                                var v30 int64 = (v2 % ((v2) | 1))
                                _ = v30
                            }
                            _ = v27
                        } else {
                            switch (fi26 << ((255) & 63)) {
                            case 32768:
                                var v31 int16 = 32767
                                _ = v31
                            default:
                                v32 := []int{(fi25 ^ (-(fi25))), 179, (^(fi25))}
                                _ = v32
                            }
                            var v33 int = -2147483648
                            v3 -= -980
                            _ = v33
                        }
                        fmt.Println(fi26, fi26)
                        v34 := []int{}
                        var v35 int8 = 42
                        if len(v34) > 0 {
                            v34[0] = 32767
                        }
                        v36 := GenS4{F0: 65535, F1: (!(!true)), F2: false}
                        v3 = (v3 | v4)
                        _ = fi26
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                case "a":
                    var v37 int16 = (-(v4))
                    _ = v37
                }
                _ = fi25
            }
        case "":
            v38 := GenS3{F0: -32768, F1: (v2 >> ((18446744073709551615) & 63)), F2: 4294967295}
            _ = v38
        }
        var v39 int64 = v2
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
