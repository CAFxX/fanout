package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint32; F1 string; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 float64 = -2.25
    var v2 string = ("a" + ("\n\t" + ((("a" + (("foo bar" + "a") + "")) + "\n\t") + (("\n\t" + ("a" + ("" + "a"))) + ((("hello" + "") + "") + "\n\t")))))
    v3 := make(map[string]string)
    v4 := []int{679, 0, -9223372036854775808, 2147483647}
    v5 := []uint32{149, 80, 65535, 652}
    _ = len(v4)
    if false {
        v1 += float64(1e10)
        v2 = v2
    }
    var v6 uint = 0
    if true {
        var v7 string = ""
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 1, F1: 0}
            _ = fi8
            _ = v9
        }
        var v10 float64 = (v1 + ((v1 - (1.5 + 1e10)) + (float64((3.14159 / 1.5)) * (0.1 + 3.14159))))
        _ = v10
        _ = v7
    } else {
        var v11 float64 = 0.0
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (-9223372036854775808 >= 0) {
            for fi12 := 0; fi12 < 19; fi12++ {
                switch (v2 + v2) {
                case "\n\t":
                    var v13 int16 = -700
                    _ = v13
                case "foo bar":
                    fmt.Println(v2, v2, v2)
                case "hello":
                    v2 = ""
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v6 = 18446744073709551615
                _ = fi12
            }
        }
        _ = v11
    }
    if (!false) {
        fmt.Println(v6, v1)
    }
    switch 132 {
    case 9223372036854775807:
        if (-341 <= 9223372036854775807) {
            var v14 int8 = 1
            _ = v14
        }
    case -2147483648:
        v6 -= (v6 ^ v6)
    default:
        switch (((v2 + "\n\t") + v2) + v2) {
        case "foo bar":
            v3[v2] = (v2 + v2)
        case "":
            if false {
                switch -368 {
                case 0:
                    var v15 uint8 = 42
                    _ = v15
                case -9223372036854775808:
                    if true {
                        if (true || false) {
                            v6 += v6
                            fmt.Println(v6)
                        }
                    }
                case 128:
                    var v16 int16 = -566
                    _ = v16
                }
                switch 801 {
                case -1:
                    if (v1 != v1) {
                        v6 -= (v6 | v6)
                    }
                case 42:
                    fmt.Println(v1)
                }
            } else {
                fmt.Println(v1, v1, v1)
            }
        default:
            v2 = v2
        }
    }
    var v17 int32 = -2147483648
    switch -352 {
    case 4294967295:
        var v18 uint64 = uint64(v6)
        _ = v18
    case -904:
        delete(v3, ("hello" + ""))
    case 695:
        fmt.Println(v17)
    }
    var v19 string = (v2 + ("\n\t" + v2))
    _ = len(v5)
    fmt.Println(v17)
    switch 70 {
    case 42:
        for fi20 := 0; fi20 < 14; fi20++ {
            for fi21 := 0; fi21 < 18; fi21++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v22 := GenS1{F0: 128, F1: (fi20 & 332)}
                _ = fi21
                _ = v22
            }
            v6 = (v6 * v6)
            if true {
                for fi23 := 0; fi23 < 16; fi23++ {
                    if (!(703 >= 0)) {
                        var v24 int32 = (v17 >> ((v6) & 63))
                        _ = v24
                    }
                    var v25 uint = (^(v6))
                    var v26 uint = 1
                    switch (fi20 / ((-622) | 1)) {
                    case 4294967295:
                        v27 := GenS1{F0: 0, F1: 41}
                        _ = v27
                    case -116:
                        switch fi23 {
                        case 42:
                            v17 = v17
                        case -1:
                            v28 := GenS2{F0: 311, F1: ((((v2 + v2) + v2) + (v2 + v2)) + v2), F2: "a"}
                            _ = v28
                        default:
                            if false {
                                v29 := GenS1{F0: 127, F1: fi20}
                                _ = v29
                            } else {
                                v17 += (v17 << (((^(v6))) & 63))
                            }
                        }
                    case 9223372036854775807:
                        v19 = "foo bar"
                    default:
                        fmt.Println(v26, v26)
                    }
                    v30 := GenS3{F0: 571}
                    if (!false) {
                        v31 := GenS2{F0: 2147483647, F1: v19, F2: (((("a" + "foo bar") + "\n\t") + v2) + "")}
                        var v32 float64 = -2.25
                        v1 = -2.25
                        _ = v31
                        _ = v32
                    } else {
                        v6 = (v6 ^ (v26 + 18446744073709551615))
                        switch v19 {
                        case "hello":
                            var v33 int64 = 0
                            _ = v33
                        }
                        v5 = append(v5, 42)
                    }
                    _ = fi23
                    _ = v25
                    _ = v26
                    _ = v30
                }
                v4 = append(v4, (-(fi20)))
            } else {
                v34 := GenS4{F0: 2, F1: 1}
                for fi35 := 0; fi35 < 10; fi35++ {
                    var v36 string = "hello"
                    if (true && (fi20 < 841)) {
                        v2 = v36
                    }
                    v19 = (v19 + (v19 + (v2 + ("" + (v2 + v36)))))
                    _ = fi35
                    _ = v36
                }
                _ = len(v3)
                _ = v34
            }
            _ = fi20
        }
    case 4294967295:
        v19 = (v19 + ((v19 + (v19 + v19)) + v19))
    case 421:
        switch 4294967295 {
        case 1:
            v37 := GenS4{F0: 0, F1: 32767}
            _ = v37
        default:
            v3[v19] = "a"
        }
    }
    v1 -= 1.5
    v17 = (v17 ^ v17)
    var v38 bool = (42 <= 65535)
    var v39 uint32 = 42
    if (int64(v39) >= -1) {
        var v40 float64 = v1
        v41 := GenS1{F0: 1, F1: 0}
        _ = v40
        _ = v41
    } else {
        fmt.Println(v2, v38)
        switch 32768 {
        case 32767:
            var v42 float64 = (-2.25 - v1)
            _ = v42
        case 799:
            v4 = append(v4, 770)
        default:
            v2 = v2
        }
        v43 := GenS2{F0: v39, F1: (v2 + v19), F2: v2}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
