package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint8; F1 int64; F2 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 int16 }

func main() {
    var v1 string = "hello"
    var v2 uint32 = 4294967295
    var v3 int32 = -2147483648
    v4 := make(map[string]int)
    var v5 uint = 4294967295
    var v6 float64 = (3.14159 + 0.1)
    for fi7 := 0; fi7 < 9; fi7++ {
        if false {
            var v8 uint64 = 660
            var v9 bool = false
            fmt.Println(v5, v5, fi7)
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    var v10 uint16 = 32767
    v11 := v4["hello"]
    var v12 float32 = 3.14159
    _ = len(v4)
    switch v11 {
    case 32767:
        v12 += float32(v6)
    case 32768:
        v3 -= (v3 + 42)
    default:
        if (!(3.14159 < (v12 / float32(float64((0.0 + 1e10)))))) {
            if (!(1 == (v5 * (v5 | 65535)))) {
                switch (v11 ^ v11) {
                case 298:
                    v13 := GenS1{F0: (v2 >> (((v5 << ((v5) & 63))) & 63)), F1: v5, F2: 1}
                    _ = v13
                case 4294967295:
                    switch 4294967295 {
                    case -570:
                        switch ("\n\t" + v1) {
                        case "a":
                            var v14 bool = (!true)
                            _ = v14
                        }
                    default:
                        if false {
                            if ((float64(v12) + (0.1 - (-2.25 / v6))) <= (3.14159 * v6)) {
                                var v15 int64 = -1
                                var v16 int32 = (v3 % (((-(v3))) | 1))
                                _ = len(v4)
                                _ = v15
                                _ = v16
                            } else {
                                v17 := v4[v1]
                                _ = v17
                            }
                        } else {
                            v18 := GenS1{F0: v2, F1: (v5 & 159), F2: -116}
                            if (v5 > (^(v5))) {
                                switch v1 {
                                case "hello":
                                    v19 := GenS1{F0: 42, F1: uint(v2), F2: 127}
                                    _ = v19
                                case "a":
                                    v2 += v2
                                }
                                for fi20 := 0; fi20 < 5; fi20++ {
                                    if false {
                                        fmt.Println(v2, v10, v3)
                                        v21 := GenS1{F0: (^(v2)), F1: 0, F2: 1}
                                        var v22 int64 = 42
                                        _ = v21
                                        _ = v22
                                    }
                                    v3 -= (v3 / (((v3 + (v3 * v3))) | 1))
                                    v23 := GenS2{F0: (v10 << ((v5) & 63)), F1: 42}
                                    _ = fi20
                                    _ = v23
                                }
                            }
                            v2 = 266
                            _ = v18
                        }
                    }
                default:
                    v12 = v12
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    v1 = "a"
                    _ = fi24
                }
                if (!true) {
                    for fi25 := 0; fi25 < 4; fi25++ {
                        if true {
                            fmt.Println(v12, fi25, v5)
                        } else {
                            var v26 uint8 = 171
                            switch fi25 {
                            case 209:
                                var v27 uint16 = v10
                                _ = v27
                            }
                            _ = v26
                        }
                        fmt.Println(v12, v11, v2)
                        fmt.Println(v11, v6)
                        _ = fi25
                    }
                    var v28 uint32 = v2
                    v29 := GenS3{F0: -1, F1: (v3 & (^(v3)))}
                    _ = v28
                    _ = v29
                }
            }
        } else {
            fmt.Println(v12, v11)
            v30 := v4["\n\t"]
            _ = v30
        }
    }
    var v31 float32 = -2.25
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 string = ((v1 + v1) + v1)
        _ = fi32
        _ = v33
    }
    v2 -= v2
    for fi34 := 0; fi34 < 12; fi34++ {
        if (!true) {
            v35 := GenS4{F0: 2, F1: 9223372036854775807, F2: int8(v3)}
            switch "a" {
            case "":
                v36 := GenS5{F0: 0}
                _ = v36
            case "hello":
                v37 := v4[v1]
                _ = v37
            }
            v5 = v5
            _ = v35
        }
        switch v1 {
        case "hello":
            if true {
                for fi38 := 0; fi38 < 11; fi38++ {
                    v3 += (v3 % (((v3 & (v3 * 2147483647))) | 1))
                    switch 4294967295 {
                    case 0:
                        var v39 int64 = 42
                        _ = v39
                    case 127:
                        v6 = float64(v31)
                    case -85:
                        v3 = (^(v3))
                    default:
                        var v40 int32 = v3
                        switch v1 {
                        case "\n\t":
                            v41 := GenS6{F0: -240, F1: 0}
                            _ = v41
                        case "":
                            if (v3 != v3) {
                                var v42 uint8 = 12
                                _ = v42
                            }
                        default:
                            switch 708 {
                            case -724:
                                v43 := GenS7{F0: 1}
                                _ = v43
                            case 186:
                                v10 = (v10 + (v10 - (v10 >> ((v5) & 63))))
                            }
                        }
                        _ = v40
                    }
                    _ = fi38
                }
                v44 := []int{fi34, -115}
                v11 -= 4294967295
                _ = v44
            }
        }
        _ = fi34
    }
    if false {
        v12 -= v31
        switch "a" {
        case "hello":
            v1 = ((v1 + v1) + "\n\t")
        }
    }
    for fi45 := 0; fi45 < 2; fi45++ {
        var v46 int16 = -721
        switch ("foo bar" + "a") {
        case "foo bar":
            var v47 int32 = v3
            _ = v47
        case "a":
            var v48 uint = (v5 >> (((v5 >> ((v5) & 63))) & 63))
            _ = v48
        case "hello":
            v49 := GenS4{F0: 165, F1: int64(v3), F2: 42}
            _ = v49
        }
        v50 := v4["\n\t"]
        _ = fi45
        _ = v46
        _ = v50
    }
    v1 = ("hello" + ((v1 + (v1 + v1)) + ""))
    fmt.Println(v11, v31, v31)
    fmt.Println(v12, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
