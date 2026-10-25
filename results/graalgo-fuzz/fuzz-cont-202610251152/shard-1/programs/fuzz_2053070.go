package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint16; F1 uint; F2 int32 }
type GenS4 struct { F0 int; F1 string }

func main() {
    var v1 uint16 = 1
    var v2 string = "hello"
    var v3 float64 = 1.5
    var v4 float64 = (-2.25 * v3)
    v5 := make(map[string]uint32)
    switch v2 {
    case "hello":
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 int16 = 0
            _ = fi6
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 4; fi8++ {
            if true {
                var v9 int32 = 1
                var v10 int8 = 42
                _ = v10
                _ = v9
            }
            var v11 int64 = 0
            _ = fi8
            _ = v11
        }
    }
    v12 := GenS1{F0: 42}
    for fi13 := 0; fi13 < 1; fi13++ {
        v4 = (3.14159 / 1.0)
        switch fi13 {
        case 1:
            var v14 uint16 = (v1 & (v1 + 65535))
            _ = v14
        case -742:
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 uint64 = 730
                var v17 uint64 = 16
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = fi13
    }
    switch 2147483647 {
    case -2147483648:
        if (96 <= 144) {
            for fi18 := 0; fi18 < 15; fi18++ {
                v19 := GenS2{F0: v1, F1: (fi18 - 127)}
                v20 := []int{(fi18 % (((fi18 & -9223372036854775808)) | 1))}
                _ = fi18
                _ = v19
                _ = v20
            }
            delete(v5, v2)
        }
    case 1:
        for fi21 := 0; fi21 < 5; fi21++ {
            if (!(!true)) {
                var v22 int16 = 1000
                if true {
                    v23 := v5["\n\t"]
                    fmt.Println(v4, v2, v1)
                    if false {
                        var v24 int8 = 127
                        for fi25 := 0; fi25 < 18; fi25++ {
                            if (v1 != v1) {
                                if true {
                                    v2 = "foo bar"
                                } else {
                                    v1 = (^(v1))
                                }
                                var v26 int = fi21
                                switch fi21 {
                                case 4294967295:
                                    if false {
                                        var v27 uint = 255
                                        v22 -= int16(v1)
                                        _ = v27
                                    } else {
                                        fmt.Println(v1)
                                        v22 = v22
                                        v28 := GenS1{F0: 42}
                                        _ = v28
                                    }
                                case 921:
                                    fmt.Println(fi21)
                                case 9223372036854775807:
                                    switch ((v2 + "\n\t") + "\n\t") {
                                    case "":
                                        v24 -= v24
                                    case "\n\t":
                                        var v29 int16 = v22
                                        _ = v29
                                    default:
                                        _ = len(v5)
                                    }
                                default:
                                    var v30 int32 = 1
                                    _ = v30
                                }
                                v31 := GenS2{F0: (v1 << ((0) & 63)), F1: (fi25 & v26)}
                                _ = v26
                                _ = v31
                            } else {
                                v32 := GenS2{F0: v1, F1: fi25}
                                if false {
                                    v4 = 1e10
                                    switch "\n\t" {
                                    case "a":
                                        if false {
                                            v33 := []int{fi21, (fi25 & fi21), 128, 2147483647}
                                            _ = v33
                                        }
                                    }
                                    v34 := []int{(fi25 - (-(fi25)))}
                                    _ = v34
                                }
                                _ = v32
                            }
                            v35 := v5["hello"]
                            _ = fi25
                            _ = v35
                        }
                        _ = v24
                    }
                    _ = v23
                } else {
                    v36 := v5[("" + (v2 + ("foo bar" + v2)))]
                    switch v2 {
                    case "a":
                        v2 = (v2 + "")
                    default:
                        v37 := GenS2{F0: (v1 & v1), F1: 42}
                        _ = v37
                    }
                    _ = v36
                }
                _ = v22
            } else {
                v4 += (3.14159 + 0.1)
                if true {
                    v4 += 1e10
                    var v38 int = -2147483648
                    for fi39 := 0; fi39 < 6; fi39++ {
                        fmt.Println(fi21, v38)
                        var v40 string = v2
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                } else {
                    v41 := []int{fi21, -2147483648, -2147483648}
                    v4 = float64(1.5)
                    v42 := GenS3{F0: (v1 | v1), F1: 286, F2: -2}
                    _ = v41
                    _ = v42
                }
                var v43 uint16 = v1
                _ = v43
            }
            var v44 float32 = 1.5
            _ = fi21
            _ = v44
        }
    default:
        if true {
            v45 := GenS1{F0: -968}
            var v46 int16 = 791
            v47 := []int{315}
            _ = v45
            _ = v46
            _ = v47
        }
    }
    v48 := GenS1{F0: 2147483647}
    fmt.Println(v2, v4, v4)
    v49 := GenS4{F0: -1, F1: (v2 + v2)}
    var v50 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
