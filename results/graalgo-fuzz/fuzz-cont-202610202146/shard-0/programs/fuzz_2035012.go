package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []uint32{42, 42, 879, 865}
    var v2 bool = (false && true)
    v3 := []string{"hello", ("" + "hello"), "\n\t"}
    var v4 uint16 = 1
    for fi5 := 0; fi5 < 3; fi5++ {
        switch fi5 {
        case -555:
            v4 = (v4 ^ 451)
        case -51:
            fmt.Println(v4, fi5)
        case 107:
            v6 := GenS1{F0: 0}
            _ = v6
        default:
            var v7 string = "hello"
            _ = v7
        }
        if len(v1) > 0 {
            v1[0] = 42
        }
        fmt.Println(v2)
        _ = fi5
    }
    v2 = v2
    switch (("foo bar" + ("" + "a")) + "a") {
    case "hello":
        v8 := GenS1{F0: (v4 / ((255) | 1))}
        _ = v8
    default:
        v9 := GenS1{F0: 65535}
        _ = v9
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    switch ("hello" + "foo bar") {
    case "hello":
        v2 = v2
    default:
        if v2 {
            v2 = (true || (1 > 0))
            v2 = (!true)
            switch "foo bar" {
            case "hello":
                v10 := GenS1{F0: v4}
                _ = v10
            case "a":
                for fi11 := 0; fi11 < 6; fi11++ {
                    fmt.Println(v4)
                    _ = fi11
                }
            case "foo bar":
                switch -148 {
                case 2147483647:
                    var v12 int8 = 127
                    _ = v12
                case -2147483648:
                    v13 := GenS1{F0: (v4 + (v4 ^ 42))}
                    _ = v13
                case 534:
                    var v14 uint16 = (^(v4))
                    _ = v14
                }
            }
        } else {
            if (!v2) {
                var v15 uint64 = 18446744073709551615
                v4 = v4
                _ = v15
            } else {
                if (true && false) {
                    v2 = ((-2147483648 < -1) && (3.14159 < 0.1))
                    v4 += (^(v4))
                } else {
                    if (206 <= 127) {
                        v16 := GenS2{F0: 65535}
                        v2 = (0.1 != 1e10)
                        _ = v16
                    } else {
                        fmt.Println(v4)
                        var v17 float32 = 1e10
                        _ = v17
                    }
                }
                var v18 float32 = -2.25
                if ((v18 - (1.5 * v18)) <= (float32(1.5) / v18)) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v19 := GenS1{F0: 2}
                    v20 := GenS2{F0: 42}
                    _ = v19
                    _ = v20
                } else {
                    var v21 uint64 = 18446744073709551615
                    _ = v21
                }
                _ = v18
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                if v2 {
                    v2 = (v2 && false)
                    var v23 uint64 = 1
                    v23 -= 769
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        if v2 {
                            v2 = (!v2)
                        } else {
                            if ((1 != 2) || (true || (v2 && (v4 != 32767)))) {
                                fmt.Println(fi22, v2, v2)
                            } else {
                                var v25 uint = 545
                                fmt.Println(fi24, v2)
                                var v26 uint8 = 151
                                _ = v25
                                _ = v26
                            }
                        }
                        var v27 uint = 18446744073709551615
                        _ = fi24
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 9; fi28++ {
                        if v2 {
                            var v29 int = (fi28 + fi28)
                            if (!v2) {
                                var v30 int16 = 42
                                v31 := GenS2{F0: 427}
                                if ((115 >= 2) && (42 <= 1)) {
                                    if (-210 != 0) {
                                        var v32 bool = (!v2)
                                        _ = v32
                                    }
                                    v29 = v29
                                }
                                _ = v30
                                _ = v31
                            } else {
                                var v33 bool = (-1 != int32(v4))
                                fmt.Println(fi28, fi28)
                                switch fi22 {
                                case 350:
                                    var v34 int8 = 42
                                    _ = v34
                                case 2147483647:
                                    var v35 uint32 = 2
                                    if v33 {
                                        v2 = (!(-887 != -2147483648))
                                    } else {
                                        v36 := GenS3{F0: 42}
                                        var v37 int32 = -2147483648
                                        var v38 float64 = -2.25
                                        if v33 {
                                            if v33 {
                                                v2 = v2
                                                var v39 uint64 = 4294967295
                                                _ = v39
                                            }
                                        }
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = v35
                                case 645:
                                    v4 -= 65535
                                default:
                                    v33 = ((50 == 255) || (v33 || true))
                                }
                                v2 = (v2 && ((!v2) && false))
                                v1 = append(v1, 65535)
                                var v40 int = fi28
                                if (v2 && v2) {
                                    switch (fi28 * fi22) {
                                    case -1:
                                        if false {
                                            fmt.Println(v2)
                                            fmt.Println(v29, fi28, v2)
                                        } else {
                                            var v41 uint64 = 0
                                            var v42 int8 = -1
                                            _ = v41
                                            _ = v42
                                        }
                                    case 128:
                                        var v43 uint64 = 18446744073709551615
                                        _ = v43
                                    case -9223372036854775808:
                                        switch (fi22 * v29) {
                                        case 0:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            if false {
                                                var v44 bool = (3.14159 < (0.1 * float64(((-2.25 * float32(-2.25)) - (0.1 - 0.1)))))
                                                fmt.Println(v33, v40, fi22)
                                                var v45 uint8 = 127
                                                _ = len(v1)
                                                _ = len(v1)
                                                _ = v44
                                                _ = v45
                                            } else {
                                                _ = len(v3)
                                                var v46 int32 = 2147483647
                                                _ = v46
                                            }
                                        }
                                    default:
                                        v40 += -2147483648
                                    }
                                    switch (fi22 * v40) {
                                    case 128:
                                        if (fi22 != (v40 ^ v29)) {
                                            if ("a" > "a") {
                                                var v47 int64 = 0
                                                v4 = 32767
                                                _ = v47
                                            }
                                            v29 += 42
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v48 float64 = float64(1e10)
                                            _ = v48
                                        }
                                    case 1:
                                        switch (v29 & v40) {
                                        case -338:
                                            v4 = (v4 << ((1) & 63))
                                            if v2 {
                                                switch "\n\t" {
                                                case "":
                                                    v49 := GenS2{F0: 255}
                                                    _ = v49
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                if (v4 <= v4) {
                                                    switch "" {
                                                    case "":
                                                        _ = len(v1)
                                                    default:
                                                        fmt.Println(v33, v40)
                                                    }
                                                    var v50 int64 = -9223372036854775808
                                                    _ = v50
                                                }
                                                _ = len(v3)
                                                if v33 {
                                                    v1 = append(v1, 42)
                                                }
                                            }
                                            v29 += fi22
                                        default:
                                            v51 := GenS1{F0: (v4 ^ v4)}
                                            _ = v51
                                        }
                                    default:
                                        v52 := GenS2{F0: 255}
                                        _ = v52
                                    }
                                    var v53 int16 = 915
                                    _ = v53
                                } else {
                                    var v54 float64 = 3.14159
                                    _ = v54
                                }
                                _ = v33
                                _ = v40
                            }
                            _ = v29
                        }
                        fmt.Println(fi22, v4)
                        _ = fi28
                    }
                    v55 := GenS2{F0: 0}
                    _ = v55
                }
                if v2 {
                    var v56 uint8 = 1
                    _ = v56
                }
                _ = fi22
            }
            var v57 int32 = 1
            _ = v57
        }
    }
    v2 = v2
    if (false && (v2 || v2)) {
        var v58 uint16 = 42
        v4 = 255
        _ = v58
    } else {
        v4 += 945
        var v59 int8 = 1
        _ = v59
    }
    v60 := GenS3{F0: 1000}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
}
