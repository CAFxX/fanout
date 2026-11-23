package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 int8; F1 int32; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int32 = 2147483647
    var v2 int = -9223372036854775808
    var v3 int = (v2 - v2)
    v4 := []int{v2}
    for fi5 := 0; fi5 < 5; fi5++ {
        _ = len(v4)
        for fi6 := 0; fi6 < 19; fi6++ {
            switch ("a" + "hello") {
            case "foo bar":
                var v7 uint32 = 4294967295
                _ = v7
            case "hello":
                var v8 int16 = -1
                _ = v8
            case "\n\t":
                var v9 float32 = 0.1
                _ = v9
            default:
                v3 = int(v1)
            }
            var v10 float64 = 0.0
            var v11 uint8 = 42
            _ = fi6
            _ = v10
            _ = v11
        }
        if (65535 <= 255) {
            if (false || true) {
                var v12 uint8 = 42
                _ = v12
            } else {
                v13 := GenS1{F0: ("\n\t" + ("foo bar" + "\n\t")), F1: 127, F2: 0}
                _ = v13
            }
        }
        _ = fi5
    }
    switch (v3 / ((v2) | 1)) {
    case 32768:
        switch v2 {
        case 42:
            var v14 uint16 = 664
            _ = v14
        default:
            var v15 uint32 = 4294967295
            _ = v15
        }
    case 127:
        for fi16 := 0; fi16 < 12; fi16++ {
            fmt.Println(v2, v3)
            _ = fi16
        }
    case 909:
        v17 := GenS1{F0: ("a" + "a"), F1: 241, F2: 654}
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := GenS2{F0: 0, F1: (v2 % (((v2 | fi19)) | 1)), F2: 183}
            switch "" {
            case "a":
                fmt.Println(v3)
            }
            switch "hello" {
            case "hello":
                fmt.Println(v1, v1, fi18)
            case "\n\t":
                v4 = append(v4, fi19)
            case "":
                var v21 bool = (true && (!(fi19 == fi18)))
                _ = v21
            }
            _ = fi19
            _ = v20
        }
        fmt.Println(v3, v2, v3)
        _ = fi18
    }
    var v22 uint16 = 762
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = (v1 + v1)
    switch v2 {
    case -2147483648:
        switch 42 {
        case 128:
            if (32767 < v22) {
                var v23 uint32 = 65535
                for fi24 := 0; fi24 < 4; fi24++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v2)
                    _ = fi24
                }
                switch ("a" + "\n\t") {
                case "hello":
                    for fi25 := 0; fi25 < 12; fi25++ {
                        switch 1 {
                        case 128:
                            _ = len(v4)
                        case 4294967295:
                            var v26 uint16 = v22
                            _ = v26
                        case 0:
                            var v27 float64 = 3.14159
                            _ = v27
                        }
                        _ = fi25
                    }
                }
                _ = v23
            } else {
                if true {
                    var v28 uint = 4294967295
                    _ = v28
                } else {
                    v2 += v2
                    switch -614 {
                    case -2147483648:
                        if len(v4) > 0 {
                            v4[0] = (v3 + v3)
                        }
                    case 947:
                        var v29 int64 = -753
                        _ = v29
                    }
                }
                fmt.Println(v2, v3, v2)
                var v30 int = v2
                _ = v30
            }
        default:
            var v31 int = v2
            _ = v31
        }
    case -1:
        switch (("\n\t" + "hello") + "\n\t") {
        case "\n\t":
            var v32 int64 = 9223372036854775807
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v1 = 42
        if (((((0.0 - (-2.25 - 0.1)) - ((0.0 * 1.5) / (1e10 / -2.25))) / 0.1) <= -2.25) && false) {
            v3 -= fi33
        }
        v34 := GenS3{F0: 127, F1: v1, F2: (202 < 2)}
        _ = fi33
        _ = v34
    }
    if len(v4) > 0 {
        v4[0] = v2
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        v36 := GenS4{F0: v22}
        if true {
            switch fi35 {
            case 0:
                if (!(2 != 2)) {
                    var v37 uint32 = 65535
                    for fi38 := 0; fi38 < 16; fi38++ {
                        var v39 int8 = -1
                        v40 := GenS5{F0: 838}
                        if len(v4) > 0 {
                            v4[0] = fi38
                        }
                        var v41 int16 = -32768
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    _ = v37
                } else {
                    if ((^(v22)) > v22) {
                        _ = len(v4)
                    } else {
                        _ = len(v4)
                        v4 = append(v4, (v2 - (fi35 - (v2 << ((65535) & 63)))))
                    }
                    var v42 uint32 = 861
                    var v43 int8 = -42
                    _ = v42
                    _ = v43
                }
            case -9223372036854775808:
                var v44 string = "\n\t"
                _ = v44
            case 42:
                v45 := GenS2{F0: 569, F1: v2, F2: 0}
                _ = v45
            default:
                for fi46 := 0; fi46 < 11; fi46++ {
                    var v47 uint = 65535
                    var v48 float64 = (0.0 / ((1.5 * -2.25) * 1e10))
                    var v49 int16 = 0
                    _ = fi46
                    _ = v47
                    _ = v48
                    _ = v49
                }
            }
        } else {
            _ = len(v4)
            switch 128 {
            case 32768:
                var v50 uint16 = v22
                _ = v50
            default:
                v2 -= v2
            }
            switch (v3 % ((v2) | 1)) {
            case 127:
                v2 = (v3 - 1)
            }
        }
        var v51 int16 = -32768
        _ = fi35
        _ = v36
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
