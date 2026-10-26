package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := []uint32{0, 42, 461, 65535}
    var v2 int = -900
    v3 := []string{}
    v2 -= v2
    v2 = v2
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 9; fi5++ {
            if false {
                if true {
                    v6 := GenS1{F0: 65535, F1: -1, F2: 255}
                    v7 := GenS2{F0: 42}
                    _ = v6
                    _ = v7
                }
                if false {
                    v2 = fi4
                } else {
                    v1 = append(v1, 0)
                    if true {
                        if (true || (true && false)) {
                            v2 += 32767
                        }
                        v8 := GenS2{F0: 65535}
                        v9 := GenS2{F0: 922}
                        _ = v8
                        _ = v9
                    }
                    if true {
                        if (false && false) {
                            var v10 uint16 = 255
                            _ = v10
                        } else {
                            v2 = fi4
                        }
                    } else {
                        v11 := GenS2{F0: 1}
                        v12 := GenS3{F0: 1000, F1: 181}
                        _ = v11
                        _ = v12
                    }
                    v2 = (fi5 & v2)
                    var v13 bool = (3.14159 <= 1.5)
                    v1 = append(v1, 655)
                    var v14 uint = 4294967295
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v13 = true
                    _ = v13
                    _ = v14
                }
            }
            _ = fi5
        }
        _ = fi4
    }
    v15 := GenS4{F0: (false || true), F1: true}
    v2 -= (v2 * -1)
    var v16 float64 = 1.5
    v2 += (-(v2))
    var v17 bool = (!false)
    for fi18 := 0; fi18 < 20; fi18++ {
        if v17 {
            v19 := GenS5{F0: v17, F1: 18446744073709551615}
            _ = v19
        } else {
            if v17 {
                v20 := GenS2{F0: 972}
                fmt.Println(v17, v17, v16)
                fmt.Println(fi18)
                _ = v20
            } else {
                switch "foo bar" {
                case "a":
                    if v17 {
                        v16 -= (v16 * -2.25)
                        switch "\n\t" {
                        case "a":
                            v17 = ("foo bar" < "foo bar")
                        default:
                            if v17 {
                                v21 := GenS6{F0: 176, F1: 0}
                                if v17 {
                                    var v22 uint64 = 4294967295
                                    fmt.Println(v22, v16)
                                    v23 := GenS1{F0: 2147483648, F1: -182, F2: 852}
                                    _ = v22
                                    _ = v23
                                } else {
                                    for fi24 := 0; fi24 < 5; fi24++ {
                                        v17 = ((39 != 287) || (!(!false)))
                                        switch (fi18 << ((625) & 63)) {
                                        case 0:
                                            switch (("a" + ("a" + "\n\t")) + ("foo bar" + "a")) {
                                            case "hello":
                                                if v17 {
                                                    _ = len(v3)
                                                } else {
                                                    var v25 int64 = 0
                                                    _ = v25
                                                }
                                            case "a":
                                                if v17 {
                                                    if ("hello" >= ("foo bar" + "a")) {
                                                        if v17 {
                                                            v3 = append(v3, ("hello" + "hello"))
                                                            v17 = v17
                                                        } else {
                                                            v17 = true
                                                            fmt.Println(v2, v2)
                                                            fmt.Println(v16)
                                                        }
                                                        v26 := GenS2{F0: 42}
                                                        _ = v26
                                                    }
                                                    var v27 float32 = float32(v16)
                                                    _ = v27
                                                } else {
                                                    v16 = v16
                                                    v2 -= 32767
                                                    switch fi18 {
                                                    case 799:
                                                        _ = len(v3)
                                                        var v28 uint32 = 190
                                                        var v29 int32 = -812
                                                        _ = v28
                                                        _ = v29
                                                    case 1:
                                                        v30 := GenS6{F0: 42, F1: 128}
                                                        _ = v30
                                                    }
                                                    if v17 {
                                                        fmt.Println(v17, fi24, v2)
                                                        var v31 int = fi24
                                                        fmt.Println(v16, fi24)
                                                        _ = v31
                                                    }
                                                    v2 -= fi18
                                                }
                                                var v32 int64 = -9223372036854775808
                                                _ = v32
                                            case "":
                                                v33 := GenS3{F0: 1, F1: 9223372036854775807}
                                                _ = v33
                                            default:
                                                if true {
                                                    var v34 uint = 42
                                                    _ = v34
                                                } else {
                                                    var v35 int32 = 1
                                                    var v36 uint16 = 65535
                                                    var v37 uint32 = 2
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                }
                                            }
                                        case 30:
                                            var v38 int16 = 1000
                                            _ = v38
                                        default:
                                            v17 = (v17 || v17)
                                        }
                                        var v39 int = (v2 % ((32767) | 1))
                                        v2 -= -1
                                        _ = fi24
                                        _ = v39
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                }
                                v3 = append(v3, "hello")
                                _ = v21
                            } else {
                                if v17 {
                                    v17 = (127 < 125)
                                }
                            }
                        }
                        for fi40 := 0; fi40 < 17; fi40++ {
                            var v41 float32 = 0.1
                            _ = fi40
                            _ = v41
                        }
                    } else {
                        v16 -= -2.25
                    }
                }
                var v42 int = fi18
                for fi43 := 0; fi43 < 10; fi43++ {
                    var v44 int8 = -40
                    _ = fi43
                    _ = v44
                }
                _ = v42
            }
            var v45 int32 = -1
            for fi46 := 0; fi46 < 13; fi46++ {
                fmt.Println(fi18, fi46, v16)
                fmt.Println(v45)
                if false {
                    var v47 uint16 = 965
                    switch (^(v2)) {
                    case -14:
                        v16 += ((v16 * float64(float32(v16))) / (float64(float32(v16)) / (v16 * v16)))
                    case 42:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                    v48 := GenS2{F0: 313}
                    _ = v47
                    _ = v48
                } else {
                    var v49 int32 = -1
                    _ = v49
                }
                fmt.Println(fi18)
                _ = fi46
            }
            _ = v45
        }
        _ = fi18
    }
    var v50 uint = 255
    v16 += (float64(0.1) / 1.0)
    v17 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v50
}
