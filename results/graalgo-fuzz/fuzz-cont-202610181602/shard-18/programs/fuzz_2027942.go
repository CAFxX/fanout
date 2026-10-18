package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int16 }

func main() {
    var v1 bool = true
    var v2 bool = ((0 <= 255) && (((42 < 65535) || (!true)) || v1))
    var v3 string = ("" + "foo bar")
    v4 := []uint32{1, 4294967295}
    v1 = v2
    for fi5 := 0; fi5 < 14; fi5++ {
        if v2 {
            var v6 int16 = -32768
            switch (fi5 >> ((255) & 63)) {
            case 127:
                switch fi5 {
                case 798:
                    var v7 uint16 = 65535
                    _ = v7
                default:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case 768:
                switch ("hello" + "") {
                case "\n\t":
                    fmt.Println(fi5, v1)
                case "a":
                    v1 = v2
                default:
                    if v1 {
                        if false {
                            for fi8 := 0; fi8 < 6; fi8++ {
                                v1 = v2
                                _ = fi8
                            }
                            v6 += 963
                            if v2 {
                                if len(v4) > 0 {
                                    v4[0] = 2147483648
                                }
                                var v9 int64 = 1
                                _ = v9
                            }
                        } else {
                            var v10 uint8 = 0
                            v1 = ((v6 / ((v6) | 1)) == v6)
                            _ = v10
                        }
                        v6 = 32767
                        v11 := GenS1{F0: "hello"}
                        _ = v11
                    } else {
                        var v12 string = v3
                        _ = v12
                    }
                }
            default:
                v13 := GenS2{F0: -1, F1: 1}
                _ = v13
            }
            _ = v6
        } else {
            v14 := GenS2{F0: 1, F1: -39}
            _ = v14
        }
        var v15 float32 = ((0.0 * 0.0) - (1e10 * 1e10))
        v3 = (("a" + "a") + "hello")
        _ = fi5
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        v2 = (128 >= 255)
        v17 := GenS3{F0: 42, F1: 2}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi16
        _ = v17
    }
    switch "\n\t" {
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            var v18 int64 = 1
            _ = v18
        case "a":
            if (v1 || (v2 && true)) {
                var v19 float32 = 3.14159
                _ = v19
            } else {
                if (!(-1 >= 0)) {
                    v1 = false
                    if false {
                        v4 = append(v4, 2)
                    } else {
                        var v20 float32 = 0.0
                        switch -401 {
                        case 570:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            var v21 uint8 = 127
                            _ = v21
                        }
                        _ = v20
                    }
                } else {
                    var v22 bool = (((-9223372036854775808 >= 1) && (((((v3 + ((v3 + "foo bar") + "foo bar")) + "\n\t") + "") != v3) && (!v1))) && (-325 > 42))
                    _ = v22
                }
            }
        default:
            _ = len(v4)
        }
    case "":
        var v23 uint16 = 65535
        _ = v23
    }
    fmt.Println(v3, v1, v3)
    var v24 uint64 = 1
    var v25 float32 = ((((float32((0.1 + ((3.14159 * 0.0) + 3.14159))) + -2.25) + 0.0) * (0.1 * 1.5)) * (0.1 * (0.0 + 0.0)))
    var v26 int64 = 0
    switch "a" {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        _ = len(v4)
    }
    if len(v4) > 0 {
        v4[0] = 0
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        if false {
            v4 = append(v4, 65535)
            var v28 int32 = -2147483648
            _ = v28
        } else {
            fmt.Println(v1, v3, v26)
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
            if true {
                switch "foo bar" {
                case "a":
                    fmt.Println(fi27)
                case "hello":
                    v24 = (v24 << ((75) & 63))
                case "foo bar":
                    switch (v3 + "foo bar") {
                    case "hello":
                        switch -368 {
                        case 1:
                            if (!v2) {
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    v4 = append(v4, 2)
                                    _ = fi29
                                }
                                var v30 int8 = -1
                                _ = v30
                            } else {
                                v31 := GenS4{F0: 199, F1: 42}
                                var v32 int16 = -32768
                                fmt.Println(v1, v26)
                                _ = v31
                                _ = v32
                            }
                        default:
                            v24 = v24
                        }
                    case "a":
                        for fi33 := 0; fi33 < 16; fi33++ {
                            if (v1 && (v2 || v1)) {
                                var v34 int64 = 699
                                _ = v34
                            } else {
                                switch ("foo bar" + v3) {
                                case "a":
                                    fmt.Println(fi27, v1, v25)
                                }
                                v35 := GenS3{F0: 330, F1: (v24 * 1)}
                                _ = v35
                            }
                            _ = fi33
                        }
                    default:
                        var v36 uint32 = 42
                        _ = v36
                    }
                default:
                    v24 -= (v24 >> ((65535) & 63))
                }
            }
        }
        fmt.Println(v3, v24, v3)
        v24 += v24
        _ = fi27
    }
    var v37 bool = v2
    var v38 string = "\n\t"
    _ = len(v4)
    v37 = v37
    var v39 uint = 18446744073709551615
    var v40 float32 = (v25 - (3.14159 + float32((((float64(v25) + (0.0 - float64(1.5))) + 1.5) / 0.1))))
    for fi41 := 0; fi41 < 16; fi41++ {
        fmt.Println(fi41, fi41, v3)
        switch fi41 {
        case 42:
            if v37 {
                v4 = append(v4, 42)
            } else {
                v25 = 0.1
                if (((v40 * v40) <= v25) || ((v24 >> ((v39) & 63)) <= v24)) {
                    fmt.Println(v2, v37, v2)
                } else {
                    var v42 uint8 = 42
                    var v43 uint16 = 255
                    _ = v42
                    _ = v43
                }
            }
        case 32767:
            if (1 < 1) {
                if ((uint8(v39) <= 2) || v1) {
                    var v44 int16 = -741
                    if (true || true) {
                        fmt.Println(v25)
                        switch fi41 {
                        case 2147483647:
                            if len(v4) > 0 {
                                v4[0] = 2
                            }
                        }
                        v40 = v40
                    } else {
                        v45 := GenS4{F0: 128, F1: (v44 ^ (v44 % ((v44) | 1)))}
                        _ = v45
                    }
                    _ = v44
                }
                switch (v38 + ("" + "foo bar")) {
                case "":
                    if true {
                        v3 = v38
                        var v46 float32 = ((v25 + float32((0.0 - (3.14159 / -2.25)))) * v40)
                        fmt.Println(v37, v46, v37)
                        _ = v46
                    }
                case "\n\t":
                    var v47 uint16 = uint16(v39)
                    _ = v47
                case "hello":
                    v3 = ""
                }
                v25 = v40
            }
        case 170:
            var v48 uint = (v39 ^ (v39 << (((v39 & v39)) & 63)))
            _ = v48
        }
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
