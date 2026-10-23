package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 bool; F1 int16; F2 uint16 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(fi7)
            var v8 int16 = 32767
            _ = fi7
            _ = v8
        }
        var v9 float64 = 1.5
        _ = fi6
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: 42, F2: 395}
    _ = v10
    return 128
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    v4 := []int{2147483647, 0, 32768}
    var v5 int64 = -1
    var v11 float32 = float32(-2.25)
    var v12 int8 = -1
    var v13 uint8 = v1
    switch "" {
    case "a":
        v12 += v12
    case "foo bar":
        v14 := GenS1{F0: 88, F1: uint32(v13), F2: 388}
        _ = v14
    case "":
        switch 4294967295 {
        case 1:
            v15 := GenS2{F0: 65535, F1: 2147483647, F2: (true && false)}
            _ = v15
        case 32768:
            for fi16 := 0; fi16 < 10; fi16++ {
                v13 += 0
                var v17 uint64 = 1
                if (true || ((65535 == 1) && true)) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    for fi18 := 0; fi18 < 14; fi18++ {
                        fmt.Println(v13, v13, v5)
                        _ = fi18
                    }
                    var v19 uint = 283
                    _ = v19
                } else {
                    switch ("hello" + "\n\t") {
                    case "":
                        v3[2147483647] = ("hello" + ("" + "a"))
                    default:
                        switch fi16 {
                        case 235:
                            fmt.Println(v17)
                        }
                    }
                    v12 -= v12
                    var v20 float64 = 0.0
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
        case 255:
            fmt.Println(v5, v13, v12)
        default:
            v12 = (v12 >> ((1) & 63))
        }
    }
    var v21 int32 = 39
    fmt.Println(v13, v11)
    var v22 uint32 = 42
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 int32 = (-(v21))
        for fi25 := 0; fi25 < 12; fi25++ {
            v24 -= (v24 ^ v21)
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    v26 := GenS3{F0: 65535, F1: false, F2: 183}
    fmt.Println(v21, v13, v21)
    v27 := GenS3{F0: 42, F1: (!false), F2: -538}
    v3[(v22 % (((v22 * v22)) | 1))] = "foo bar"
    if (((v11 - 0.0) + 3.14159) >= v11) {
        v2[v22] = 2147483648
        var v28 float32 = 0.1
        if false {
            var v29 int16 = 731
            fmt.Println(v21)
            switch 4294967295 {
            case 277:
                v5 = int64(v22)
            case 32768:
                var v30 uint32 = 141
                _ = v30
            default:
                var v31 uint64 = 0
                _ = v31
            }
            _ = v29
        } else {
            if true {
                v4 = append(v4, 85)
                switch ("\n\t" + ("" + ("foo bar" + ("" + "")))) {
                case "a":
                    v12 -= v12
                default:
                    var v32 float64 = 1.5
                    _ = v32
                }
                switch "foo bar" {
                case "a":
                    var v33 int = 0
                    _ = v33
                case "hello":
                    _ = len(v3)
                }
            }
            _ = len(v2)
        }
        _ = v28
    } else {
        for fi34 := 0; fi34 < 3; fi34++ {
            switch (fi34 >> ((454) & 63)) {
            case 467:
                v21 -= v21
            case 9223372036854775807:
                fmt.Println(v13, v5)
            }
            var v35 float64 = 3.14159
            _ = fi34
            _ = v35
        }
        var v36 uint8 = (v1 / ((genFn0()) | 1))
        if true {
            var v37 uint16 = uint16(v36)
            for fi38 := 0; fi38 < 13; fi38++ {
                fmt.Println(v1)
                for fi39 := 0; fi39 < 14; fi39++ {
                    var v40 int8 = v12
                    _ = fi39
                    _ = v40
                }
                var v41 int64 = (v5 | v5)
                _ = fi38
                _ = v41
            }
            _ = v37
        } else {
            fmt.Println(v21, v22, v1)
            v11 = v11
        }
        _ = v36
    }
    v42 := GenS4{F0: v22, F1: -1}
    v43 := GenS5{F0: true, F1: 991, F2: 1}
    delete(v2, (v22 >> ((1) & 63)))
    v4 = append(v4, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
