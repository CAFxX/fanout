package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 string; F1 uint; F2 string }
type GenS4 struct { F0 uint32; F1 string; F2 uint8 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 float64 = float64(1e10)
    if (!true) {
        if false {
            var v4 uint16 = 1
            var v5 uint64 = 18446744073709551615
            var v6 uint64 = (v5 | 369)
            _ = v4
            _ = v5
            _ = v6
        }
        if false {
            if (1 >= 4294967295) {
                var v7 float64 = v3
                var v8 int32 = -2147483648
                _ = v7
                _ = v8
            } else {
                switch 42 {
                case 1:
                    var v9 int16 = -1
                    _ = v9
                case 787:
                    v3 = 1.5
                }
            }
        }
        fmt.Println(v3)
    } else {
        if true {
            v3 = 1.5
            switch ("\n\t" + "\n\t") {
            case "hello":
                switch 1 {
                case 4294967295:
                    v3 += v3
                case 1:
                    v3 += v3
                }
            default:
                var v10 int = 2147483647
                _ = v10
            }
            switch 990 {
            case -456:
                for fi11 := 0; fi11 < 12; fi11++ {
                    fmt.Println(v3, v3)
                    for fi12 := 0; fi12 < 3; fi12++ {
                        if true {
                            v13 := GenS1{F0: 0}
                            _ = v13
                        } else {
                            v14 := GenS2{F0: -1, F1: 956}
                            if true {
                                v15 := v1[1]
                                var v16 int32 = -646
                                v1[992] = (108 > -199)
                                _ = v15
                                _ = v16
                            }
                            _ = v14
                        }
                        _ = fi12
                    }
                    var v17 uint16 = 42
                    _ = fi11
                    _ = v17
                }
            case -623:
                var v18 uint = 4294967295
                _ = v18
            default:
                v3 += 1e10
            }
        } else {
            v3 = float64(3.14159)
        }
        v3 -= 0.0
    }
    v3 = 0.0
    if (!(true && true)) {
        fmt.Println(v3)
        var v19 uint = 255
        fmt.Println(v19)
        _ = v19
    }
    fmt.Println(v3)
    switch 2147483647 {
    case -9223372036854775808:
        var v20 float64 = v3
        _ = v20
    case -928:
        fmt.Println(v3)
    default:
        v3 -= 1.5
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        _ = len(v1)
        _ = fi21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            if true {
                var v24 int32 = 1
                v3 = v3
                _ = v24
            }
            if true {
                fmt.Println(fi23, fi23)
                var v25 int8 = 42
                var v26 uint = 42
                _ = v25
                _ = v26
            } else {
                var v27 int64 = 42
                v27 += v27
                var v28 int = 32768
                v29 := []int{32767, -378, (v28 % (((^(v28))) | 1)), (fi23 & (fi23 & (fi22 + (fi22 - fi23))))}
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi23
        }
        if false {
            v30 := GenS1{F0: 42}
            switch (fi22 + (-(fi22))) {
            case -2147483648:
                fmt.Println(fi22)
            case 2147483647:
                var v31 bool = (65535 == 4294967295)
                _ = v31
            default:
                for fi32 := 0; fi32 < 20; fi32++ {
                    v33 := GenS3{F0: "hello", F1: 1, F2: "\n\t"}
                    fmt.Println(v3, fi22, fi32)
                    v34 := GenS1{F0: 146}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            if (!false) {
                v35 := GenS4{F0: 2147483647, F1: "", F2: 127}
                _ = v35
            }
            _ = v30
        }
        _ = fi22
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        v37 := []int{}
        if len(v37) >= 2 {
            v37 = v37[1:len(v37)]
        }
        var v38 int16 = 0
        _ = fi36
        _ = v37
        _ = v38
    }
    if false {
        var v39 uint8 = 255
        _ = v39
    }
    switch -9223372036854775808 {
    case -218:
        v40 := GenS5{F0: 398, F1: ("" + "a")}
        _ = v40
    case -685:
        switch ("" + ("hello" + ("hello" + "foo bar"))) {
        case "":
            fmt.Println(v3)
        case "\n\t":
            switch -103 {
            case 42:
                v41 := []int{0, -637, -1, 1}
                _ = v41
            case 32767:
                var v42 bool = true
                _ = v42
            case 1:
                var v43 int32 = -877
                _ = v43
            default:
                var v44 float64 = (v3 / (v3 + v3))
                _ = v44
            }
        default:
            v3 = (0.0 + (v3 - v3))
        }
    default:
        var v45 int8 = 1
        _ = v45
    }
    if false {
        v3 += 1e10
        for fi46 := 0; fi46 < 3; fi46++ {
            v3 = (v3 * float64((3.14159 + 1.5)))
            _ = fi46
        }
    } else {
        if (32767 <= -505) {
            var v47 float32 = ((1e10 - float32(v3)) + 0.0)
            _ = v47
        }
    }
    switch "hello" {
    case "a":
        v48 := GenS6{F0: "foo bar"}
        _ = v48
    case "foo bar":
        var v49 int8 = -73
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
