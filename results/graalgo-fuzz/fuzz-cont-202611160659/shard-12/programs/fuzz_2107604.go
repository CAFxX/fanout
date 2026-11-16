package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0() string {
    var v6 int16 = -1
    v6 += -1
    if (!(42 >= 345)) {
        var v7 uint = 785
        _ = v7
    } else {
        switch "hello" {
        case "a":
            fmt.Println(v6)
        }
    }
    v8 := []int{128, 32768}
    _ = v6
    _ = v8
    return ((("hello" + "foo bar") + ("" + "a")) + "")
}

func main() {
    var v1 uint16 = 32767
    var v2 uint64 = 0
    v3 := make(map[int]uint32)
    var v4 string = ("foo bar" + ("" + ""))
    var v5 uint8 = 255
    for fi9 := 0; fi9 < 17; fi9++ {
        if (v4 > (v4 + v4)) {
            var v10 int64 = 9223372036854775807
            v11 := GenS1{F0: (v2 >> ((4) & 63))}
            _ = v10
            _ = v11
        } else {
            var v12 bool = (8 > 2)
            if ((v12 || v12) && (168 <= 2)) {
                v1 -= v1
            } else {
                switch fi9 {
                case 32768:
                    switch 152 {
                    case 42:
                        v13 := GenS2{F0: 2147483647, F1: (^(fi9)), F2: 2}
                        _ = v13
                    }
                case -9223372036854775808:
                    switch (^(fi9)) {
                    case -473:
                        for fi14 := 0; fi14 < 5; fi14++ {
                            fmt.Println(v12, v5)
                            _ = fi14
                        }
                    case -642:
                        for fi15 := 0; fi15 < 18; fi15++ {
                            v16 := []int{42, 678, fi9}
                            var v17 int64 = 1
                            fmt.Println(fi15, v4, fi15)
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case 127:
                        if ((v5 | 128) > 167) {
                            v3[fi9] = 65535
                            var v18 bool = (v12 || false)
                            _ = v18
                        }
                    }
                }
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 uint16 = 2
                v1 = 2
                if v12 {
                    v20 = (^(v20))
                    v21 := GenS3{F0: 21, F1: (v5 | 42)}
                    var v22 int = 741
                    var v23 int64 = 319
                    v24 := GenS4{F0: (v1 | (v20 / (((v1 - (^(v20)))) | 1)))}
                    if (753 <= 0) {
                        v25 := GenS5{F0: v2, F1: v23}
                        _ = v25
                    }
                    var v26 uint = 1
                    var v27 float64 = -2.25
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v26
                    _ = v27
                }
                _ = fi19
                _ = v20
            }
            _ = v12
        }
        _ = len(v3)
        _ = fi9
    }
    v3[127] = 0
    v28 := v3[9223372036854775807]
    v1 -= v1
    var v29 int32 = 0
    var v30 uint = 255
    v28 = uint32(v5)
    for fi31 := 0; fi31 < 6; fi31++ {
        v32 := GenS6{F0: (^(v5)), F1: (v2 + uint64(v28))}
        var v33 string = v4
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 bool = (v28 > 552)
        fmt.Println(v1, v28, v30)
        _ = fi34
        _ = v35
    }
    v36 := []int{32767}
    switch "" {
    case "":
        var v37 float32 = float32(((1e10 * float64(1e10)) * (1.5 - 0.0)))
        _ = v37
    default:
        for fi38 := 0; fi38 < 3; fi38++ {
            var v39 uint8 = uint8(v30)
            switch -9223372036854775808 {
            case -2147483648:
                v40 := GenS3{F0: uint32(v30), F1: (^(v5))}
                _ = v40
            case 2147483647:
                for fi41 := 0; fi41 < 19; fi41++ {
                    var v42 int8 = -1
                    var v43 int16 = 42
                    if len(v36) > 0 {
                        v36[0] = (fi41 & 42)
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            case 614:
                var v44 uint8 = (v5 & v5)
                _ = v44
            }
            if len(v36) > 0 {
                v36[0] = 4294967295
            }
            _ = fi38
            _ = v39
        }
    }
    if (32767 < -1) {
        v1 += (^(v1))
        switch ("foo bar" + v4) {
        case "foo bar":
            switch 1 {
            case 32767:
                v30 += v30
            case 4294967295:
                for fi45 := 0; fi45 < 2; fi45++ {
                    var v46 int64 = int64(v1)
                    switch 127 {
                    case 42:
                        fmt.Println(v46, fi45, v2)
                    default:
                        var v47 float64 = 1.5
                        _ = v47
                    }
                    v48 := v3[-9223372036854775808]
                    _ = fi45
                    _ = v46
                    _ = v48
                }
            case 128:
                switch v4 {
                case "hello":
                    delete(v3, 0)
                }
            }
        case "a":
            fmt.Println(v28)
        case "\n\t":
            var v49 float32 = float32(3.14159)
            _ = v49
        }
    }
    fmt.Println(v28, v4)
    for fi50 := 0; fi50 < 10; fi50++ {
        fmt.Println(v29, v4)
        for fi51 := 0; fi51 < 18; fi51++ {
            var v52 int = 2147483647
            _ = fi51
            _ = v52
        }
        _ = fi50
    }
    if true {
        fmt.Println(v29, v5, v28)
        for fi53 := 0; fi53 < 8; fi53++ {
            var v54 bool = (float64(1.5) != -2.25)
            for fi55 := 0; fi55 < 20; fi55++ {
                v54 = v54
                _ = fi55
            }
            _ = fi53
            _ = v54
        }
        fmt.Println(v1, v4, v4)
    }
    var v56 uint8 = 255
    v57 := v3[-1]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v56
    _ = v57
}
