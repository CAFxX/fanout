package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 string }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]string)
    v3 := []uint32{42, 42, 42, 285}
    v4 := make(map[int]uint32)
    fmt.Println(v1)
    _ = len(v2)
    v5 := GenS1{F0: 0, F1: 42, F2: -1}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int = (fi6 >> ((255) & 63))
        if true {
            var v8 float64 = v1
            v9 := GenS2{F0: 850, F1: -772}
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := v2["a"]
                if true {
                    var v12 float32 = 0.0
                    _ = v12
                }
                v13 := GenS1{F0: 0, F1: (^(fi6)), F2: 0}
                _ = fi10
                _ = v11
                _ = v13
            }
            _ = v8
            _ = v9
        }
        if (-125 > -104) {
            if false {
                var v14 uint = 2
                fmt.Println(fi6)
                _ = v14
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi6
        _ = v7
    }
    v15 := GenS3{F0: 0, F1: 57}
    v1 = v1
    v1 = 0.1
    v2[(("a" + "hello") + "a")] = "\n\t"
    if (496 != -1) {
        fmt.Println(v1)
    } else {
        var v16 int8 = -84
        if true {
            _ = len(v4)
        }
        v17 := GenS3{F0: -9223372036854775808, F1: 65535}
        _ = v16
        _ = v17
    }
    switch 0 {
    case 127:
        switch "" {
        case "a":
            v18 := GenS3{F0: -1, F1: 382}
            _ = v18
        case "":
            v19 := GenS4{F0: 0, F1: 1}
            _ = v19
        }
    case -947:
        var v20 int16 = 32767
        _ = v20
    case 9223372036854775807:
        v21 := GenS2{F0: 2, F1: -1}
        _ = v21
    default:
        switch 1 {
        case -585:
            v1 = v1
        case 702:
            switch 32767 {
            case -1:
                var v22 int = 737
                _ = v22
            }
        case 32767:
            for fi23 := 0; fi23 < 3; fi23++ {
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                v1 -= 3.14159
                _ = fi23
            }
        }
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        if false {
            for fi25 := 0; fi25 < 16; fi25++ {
                v3 = append(v3, 1)
                switch (fi25 * (fi25 / ((fi25) | 1))) {
                case -966:
                    if false {
                        var v26 int32 = -843
                        delete(v4, -2147483648)
                        v27 := GenS5{F0: 0}
                        _ = v26
                        _ = v27
                    } else {
                        var v28 uint16 = 65535
                        v29 := GenS2{F0: 922, F1: -1}
                        _ = v28
                        _ = v29
                    }
                case 163:
                    var v30 uint = 0
                    _ = v30
                case 9223372036854775807:
                    v31 := GenS6{F0: 128, F1: "foo bar"}
                    _ = v31
                default:
                    v1 -= (v1 + 0.1)
                }
                var v32 int16 = -32768
                _ = fi25
                _ = v32
            }
            switch "hello" {
            case "a":
                switch fi24 {
                case -657:
                    v33 := GenS5{F0: 0}
                    _ = v33
                case 42:
                    var v34 uint16 = 65535
                    _ = v34
                }
            default:
                v35 := v2[("foo bar" + ("foo bar" + "a"))]
                _ = v35
            }
        } else {
            v36 := GenS3{F0: -1, F1: 32767}
            var v37 string = "a"
            _ = v36
            _ = v37
        }
        _ = fi24
    }
    v1 = ((v1 / (v1 + float64((0.1 * (3.14159 - -2.25))))) - (v1 - v1))
    var v38 uint16 = 0
    if true {
        var v39 int16 = -1
        _ = v39
    } else {
        _ = len(v2)
    }
    v40 := GenS6{F0: 255, F1: "a"}
    if (42 <= 18446744073709551615) {
        v41 := GenS5{F0: 4294967295}
        _ = v41
    }
    var v42 int64 = 1
    fmt.Println(v1, v38, v42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
