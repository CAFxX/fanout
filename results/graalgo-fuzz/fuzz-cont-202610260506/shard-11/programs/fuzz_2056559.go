package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 bool; F2 int }

func main() {
    var v1 uint16 = 0
    var v2 int16 = 0
    v3 := []uint32{}
    var v4 float32 = 1.5
    v1 -= 255
    for fi5 := 0; fi5 < 18; fi5++ {
        v1 += v1
        _ = fi5
    }
    if (("\n\t" + ("a" + "\n\t")) <= ("hello" + ("\n\t" + "\n\t"))) {
        v2 -= 42
        v3 = append(v3, 65535)
    }
    v6 := GenS1{F0: "a"}
    v1 = (v1 - (v1 * (v1 << ((713) & 63))))
    for fi7 := 0; fi7 < 1; fi7++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (-2.25 > float64(v4)) {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 uint = 18446744073709551615
                switch "foo bar" {
                case "\n\t":
                    if ((!false) || (true && false)) {
                        switch "a" {
                        case "a":
                            if true {
                                var v10 uint16 = v1
                                var v11 string = "a"
                                _ = v10
                                _ = v11
                            } else {
                                fmt.Println(v2, v2, v9)
                                v12 := GenS1{F0: "foo bar"}
                                var v13 string = "\n\t"
                                var v14 uint = (v9 >> (((v9 % (((v9 - v9)) | 1))) & 63))
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                        case "hello":
                            if false {
                                fmt.Println(v2)
                                var v15 uint16 = v1
                                var v16 int8 = 0
                                _ = v15
                                _ = v16
                            }
                        }
                        if false {
                            v2 = (^(v2))
                            v17 := GenS1{F0: ""}
                            _ = v17
                        }
                    } else {
                        var v18 int8 = 1
                        var v19 float32 = v4
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi8
                _ = v9
            }
        } else {
            _ = len(v3)
            v20 := GenS2{F0: 1}
            _ = v20
        }
        var v21 uint16 = 32767
        _ = fi7
        _ = v21
    }
    v22 := GenS3{F0: (v1 % ((v1) | 1)), F1: true, F2: -330}
    if (v2 == v2) {
        if (-2147483648 == -946) {
            v23 := GenS3{F0: (^(v1)), F1: false, F2: 32767}
            var v24 uint16 = (v1 & v1)
            _ = v23
            _ = v24
        } else {
            switch 601 {
            case 4294967295:
                if true {
                    var v25 int = -11
                    _ = v25
                }
            default:
                if false {
                    if (!(0.1 != float64((v4 + (3.14159 - v4))))) {
                        v1 += (v1 ^ 42)
                        v26 := GenS2{F0: 127}
                        _ = v26
                    }
                } else {
                    var v27 int32 = -982
                    var v28 uint8 = 1
                    _ = v27
                    _ = v28
                }
            }
            v29 := GenS1{F0: "foo bar"}
            var v30 int64 = -9223372036854775808
            _ = v29
            _ = v30
        }
        var v31 uint64 = 0
        v32 := GenS2{F0: 100}
        _ = v31
        _ = v32
    } else {
        v33 := GenS3{F0: v1, F1: (!("" > "a")), F2: 156}
        if false {
            for fi34 := 0; fi34 < 8; fi34++ {
                for fi35 := 0; fi35 < 6; fi35++ {
                    v36 := GenS1{F0: ("foo bar" + "")}
                    _ = fi35
                    _ = v36
                }
                v37 := GenS3{F0: v1, F1: (1 < -128), F2: -1}
                _ = fi34
                _ = v37
            }
            v4 += (0.0 * v4)
            v38 := GenS3{F0: (v1 & 65535), F1: (false || true), F2: -51}
            _ = v38
        }
        var v39 int8 = 105
        _ = v33
        _ = v39
    }
    if (false && true) {
        var v40 string = "hello"
        fmt.Println(v2, v40)
        v2 = 32767
        _ = v40
    }
    v1 = v1
    for fi41 := 0; fi41 < 17; fi41++ {
        v2 = (v2 << ((744) & 63))
        fmt.Println(v4)
        v1 += (v1 / (((v1 & v1)) | 1))
        _ = fi41
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4, v1)
    v2 = 720
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v6
}
