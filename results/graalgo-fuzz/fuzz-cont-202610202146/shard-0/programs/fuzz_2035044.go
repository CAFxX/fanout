package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (!false)
    v3 := []uint32{221, 4294967295, 2, 4294967295}
    var v4 uint64 = 18446744073709551615
    v4 = v4
    if v2 {
        v4 += 18446744073709551615
        fmt.Println(v2)
    }
    _ = len(v3)
    if (42 != -1) {
        fmt.Println(v2)
        v5 := GenS1{F0: 285, F1: 6, F2: 42}
        _ = v5
    } else {
        for fi6 := 0; fi6 < 13; fi6++ {
            v2 = (!(!(922 != 42)))
            _ = fi6
        }
        switch "" {
        case "foo bar":
            if false {
                v7 := GenS2{F0: 209}
                _ = v7
            }
        case "a":
            v8 := GenS3{F0: 127}
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint32 = 42
                if (0 > 2147483647) {
                    switch fi9 {
                    case 729:
                        var v11 string = ("\n\t" + "hello")
                        _ = v11
                    case 32767:
                        for fi12 := 0; fi12 < 8; fi12++ {
                            v4 += 42
                            _ = fi12
                        }
                    case -958:
                        v13 := GenS4{F0: -74}
                        _ = v13
                    default:
                        if v2 {
                            fmt.Println(v2, v4)
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                v1[v10] = v10
                            }
                        } else {
                            switch "" {
                            case "hello":
                                v14 := GenS5{F0: "\n\t"}
                                _ = v14
                            case "\n\t":
                                var v15 int64 = -9223372036854775808
                                _ = v15
                            case "":
                                switch ("a" + "hello") {
                                case "foo bar":
                                    for fi16 := 0; fi16 < 20; fi16++ {
                                        v17 := GenS5{F0: "hello"}
                                        v18 := GenS5{F0: ""}
                                        switch fi16 {
                                        case 32768:
                                            var v19 string = (("hello" + "foo bar") + ((("" + "hello") + "foo bar") + ("\n\t" + "\n\t")))
                                            _ = v19
                                        case 0:
                                            v10 = (v10 - 42)
                                        default:
                                            if v2 {
                                                switch ("a" + "foo bar") {
                                                case "a":
                                                    switch 32767 {
                                                    case -2147483648:
                                                        if len(v3) > 0 {
                                                            v3[0] = uint32(v4)
                                                        }
                                                        var v20 uint8 = 2
                                                        _ = v20
                                                    case 42:
                                                        v2 = v2
                                                        var v21 uint64 = v4
                                                        v10 = v10
                                                        var v22 uint16 = 1
                                                        _ = v21
                                                        _ = v22
                                                    case 32768:
                                                        _ = len(v1)
                                                    }
                                                default:
                                                    if v2 {
                                                        if (-429 <= 0) {
                                                            v10 -= v10
                                                            if v2 {
                                                                delete(v1, 2)
                                                                switch fi9 {
                                                                case 1:
                                                                    v10 -= uint32(v4)
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                                var v23 bool = (1e10 >= 1e10)
                                                                _ = v23
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        } else {
                                                            v2 = false
                                                            fmt.Println(v2)
                                                        }
                                                        v2 = false
                                                        v2 = false
                                                        switch 2147483647 {
                                                        case -2147483648:
                                                            fmt.Println(v4)
                                                        }
                                                    } else {
                                                        var v24 int8 = -1
                                                        fmt.Println(fi9)
                                                        _ = v24
                                                    }
                                                }
                                                var v25 int8 = -1
                                                v4 = v4
                                                v10 = v10
                                                _ = v25
                                            }
                                            fmt.Println(v10, v4, fi16)
                                            var v26 uint = 65535
                                            _ = v26
                                        }
                                        _ = fi16
                                        _ = v17
                                        _ = v18
                                    }
                                case "":
                                    v27 := GenS5{F0: "\n\t"}
                                    _ = v27
                                }
                            }
                            for fi28 := 0; fi28 < 3; fi28++ {
                                delete(v1, (v10 - 160))
                                _ = fi28
                            }
                            v4 += v4
                        }
                    }
                    var v29 int = (^(fi9))
                    _ = v29
                }
                for fi30 := 0; fi30 < 4; fi30++ {
                    var v31 float64 = (3.14159 / float64((((1e10 - -2.25) - -2.25) - (0.1 / 0.1))))
                    _ = fi30
                    _ = v31
                }
                _ = fi9
                _ = v10
            }
        }
        switch 0 {
        case 9223372036854775807:
            fmt.Println(v4, v2)
        case 298:
            v4 -= v4
        }
    }
    v2 = false
    v32 := GenS5{F0: ("foo bar" + ("\n\t" + ""))}
    fmt.Println(v2)
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 uint = 583
        if true {
            v34 = (v34 % (((v34 ^ v34)) | 1))
            switch "hello" {
            case "hello":
                v2 = true
            case "":
                v35 := GenS6{F0: (1 <= (fi33 / ((384) | 1))), F1: 166}
                _ = v35
            default:
                v4 += (v4 / (((v4 / ((v4) | 1))) | 1))
            }
            for fi36 := 0; fi36 < 15; fi36++ {
                switch "" {
                case "":
                    v37 := GenS7{F0: int64(v4), F1: (("hello" + "foo bar") + "")}
                    fmt.Println(fi33, fi36)
                    switch (fi36 | fi36) {
                    case 0:
                        var v38 float64 = 0.0
                        _ = v38
                    }
                    _ = v37
                case "hello":
                    var v39 uint8 = 0
                    _ = v39
                case "foo bar":
                    v40 := GenS1{F0: 0, F1: -90, F2: -1}
                    _ = v40
                default:
                    fmt.Println(fi36)
                }
                _ = fi36
            }
        } else {
            var v41 bool = (!v2)
            _ = v41
        }
        _ = fi33
        _ = v34
    }
    v2 = false
    _ = len(v1)
    switch "a" {
    case "a":
        if ((int8(v4) >= 42) && v2) {
            v2 = v2
        } else {
            var v42 bool = v2
            v4 -= v4
            for fi43 := 0; fi43 < 13; fi43++ {
                v4 = v4
                switch ("" + "hello") {
                case "hello":
                    var v44 uint8 = 179
                    _ = v44
                case "foo bar":
                    v3 = append(v3, 42)
                }
                _ = fi43
            }
            _ = v42
        }
    }
    fmt.Println(v4)
    v45 := v1[1]
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v45
}
