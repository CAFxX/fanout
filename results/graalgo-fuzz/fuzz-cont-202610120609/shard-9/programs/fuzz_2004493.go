package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }

func main() {
    v1 := []string{"a", "\n\t", "foo bar", "hello"}
    v2 := make(map[uint32]bool)
    var v3 float32 = float32(1.5)
    var v4 int = 127
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 0
    v7 := v2[(v6 / (((v6 * 603)) | 1))]
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v3 -= ((0.0 - float32(0.0)) + v3)
            _ = fi9
        }
        _ = fi8
    }
    var v10 float32 = float32((((0.0 - -2.25) + -2.25) / (((float64(v3) - 0.1) + 3.14159) + 0.1)))
    for fi11 := 0; fi11 < 3; fi11++ {
        _ = len(v5)
        v12 := GenS1{F0: 255}
        var v13 float64 = 0.0
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 127}
    switch "hello" {
    case "foo bar":
        v15 := GenS1{F0: 0}
        _ = v15
    }
    var v16 float32 = v3
    var v17 uint8 = 128
    v18 := GenS2{F0: -404, F1: 65535}
    switch (("\n\t" + "a") + "\n\t") {
    case "":
        if (!v7) {
            var v19 uint8 = (v17 * 255)
            _ = v19
        } else {
            fmt.Println(v17)
            delete(v5, 65535)
            for fi20 := 0; fi20 < 7; fi20++ {
                if false {
                    if (934 == 2) {
                        v21 := GenS2{F0: -479, F1: 2}
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 bool = (-128 < 0)
                            if v7 {
                                v1 = append(v1, "hello")
                            } else {
                                var v24 int64 = -333
                                var v25 int32 = 190
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS1{F0: (v17 / ((v17) | 1))}
                            _ = fi22
                            _ = v23
                            _ = v26
                        }
                        v10 = v10
                    }
                    var v27 uint32 = (v6 % (((v6 << ((2) & 63))) | 1))
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 7; fi28++ {
                        fmt.Println(v10)
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        var v29 float32 = ((((0.0 + (v3 / v10)) * v10) / v10) / v3)
                        v30 := GenS1{F0: v17}
                        var v31 int8 = 127
                        var v32 int64 = 1
                        switch (("hello" + ("" + ("a" + ""))) + ("\n\t" + "hello")) {
                        case "":
                            fmt.Println(v32)
                        default:
                            v3 += 0.0
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            switch -265 {
                            case 42:
                                switch ("\n\t" + "") {
                                case "":
                                    switch fi28 {
                                    case 78:
                                        _ = len(v5)
                                    }
                                    var v33 uint16 = 0
                                    v4 = v4
                                    _ = v33
                                case "hello":
                                    if v7 {
                                        v6 = v6
                                        if (v7 || v7) {
                                            v4 = fi20
                                            if (((v29 + 1e10) < (float32(1e10) + float32(1.5))) || (!true)) {
                                                switch 212 {
                                                case 32768:
                                                    var v34 int32 = 1
                                                    _ = v34
                                                case 42:
                                                    var v35 int8 = (v31 % ((v31) | 1))
                                                    v4 += 32767
                                                    delete(v5, (v6 >> ((1) & 63)))
                                                    fmt.Println(v35, fi28)
                                                    var v36 int32 = 395
                                                    _ = v35
                                                    _ = v36
                                                case 9223372036854775807:
                                                    var v37 string = ("" + (((("a" + "foo bar") + "foo bar") + "a") + "\n\t"))
                                                    _ = v37
                                                }
                                                _ = len(v1)
                                            }
                                        }
                                        v10 += 0.1
                                    } else {
                                        v17 = (^(v17))
                                        var v38 float32 = ((((v16 - 1e10) * 3.14159) - v10) - v3)
                                        v32 = (v32 << ((42) & 63))
                                        _ = v38
                                    }
                                }
                                var v39 float64 = (0.0 - 1e10)
                                var v40 int8 = v31
                                var v41 int8 = (v40 ^ v40)
                                v42 := GenS3{F0: 18446744073709551615, F1: v6, F2: "a"}
                                v16 += (-2.25 + float32(v39))
                                var v43 uint = 18446744073709551615
                                _ = v39
                                _ = v40
                                _ = v41
                                _ = v42
                                _ = v43
                            case -10:
                                fmt.Println(v10)
                            case -888:
                                v29 = 0.0
                            default:
                                var v44 uint = 0
                                _ = v44
                            }
                        }
                        v45 := GenS1{F0: v17}
                        var v46 int32 = -2147483648
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v31
                        _ = v32
                        _ = v45
                        _ = v46
                    }
                }
                for fi47 := 0; fi47 < 2; fi47++ {
                    v48 := GenS1{F0: (v17 ^ v17)}
                    fmt.Println(fi20)
                    _ = fi47
                    _ = v48
                }
                fmt.Println(v4, v6, v3)
                _ = fi20
            }
        }
    }
    v2[2147483647] = (true || (0 >= 4294967295))
    if false {
        var v49 float64 = 1e10
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
