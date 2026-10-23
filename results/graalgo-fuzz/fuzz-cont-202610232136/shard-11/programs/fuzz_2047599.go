package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2147483648
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    var v5 int = 4294967295
    v5 = (v5 * (v5 * (^(v5))))
    for fi6 := 0; fi6 < 4; fi6++ {
        v2 -= 4294967295
        if false {
            for fi7 := 0; fi7 < 7; fi7++ {
                var v8 uint16 = 65535
                _ = fi7
                _ = v8
            }
            v9 := []int{(v5 + (v5 << ((4294967295) & 63)))}
            var v10 int8 = 100
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := []int{}
        switch "foo bar" {
        case "hello":
            v2 += v2
        case "\n\t":
            var v13 uint32 = v2
            _ = v13
        case "a":
            if false {
                for fi14 := 0; fi14 < 19; fi14++ {
                    v5 = fi11
                    v15 := GenS1{F0: v2}
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = fi14
                    _ = v15
                }
                v2 += v2
            }
        default:
            switch -9223372036854775808 {
            case 1:
                _ = len(v12)
            }
        }
        _ = fi11
        _ = v12
    }
    v16 := []int{}
    v17 := GenS1{F0: v2}
    v5 = v5
    var v18 bool = false
    v2 -= v2
    if len(v16) > 0 {
        v16[0] = (^(v5))
    }
    var v19 int32 = -56
    v5 = v5
    v18 = ((!(v18 || v18)) && v18)
    for fi20 := 0; fi20 < 18; fi20++ {
        if false {
            for fi21 := 0; fi21 < 20; fi21++ {
                if true {
                    if (!(!(fi20 >= fi20))) {
                        switch ("hello" + "") {
                        case "foo bar":
                            v18 = (!(!v18))
                        case "\n\t":
                            v18 = (0 > 353)
                        case "hello":
                            v2 -= v2
                        default:
                            var v22 int16 = int16(v19)
                            _ = v22
                        }
                        v16 = append(v16, fi20)
                        if true {
                            v19 += v19
                            var v23 string = ("a" + "")
                            v16 = append(v16, (-(v5)))
                            switch fi21 {
                            case -60:
                                fmt.Println(v2, fi21)
                            default:
                                fmt.Println(v19, v23, v5)
                            }
                            var v24 float64 = (0.1 / 1.0)
                            _ = v23
                            _ = v24
                        }
                        v25 := GenS1{F0: 977}
                        _ = v25
                    } else {
                        v18 = false
                        var v26 uint = 33
                        fmt.Println(v5)
                        _ = v26
                    }
                    v27 := GenS1{F0: (v2 / ((v2) | 1))}
                    if true {
                        v5 -= (fi20 / (((^(fi21))) | 1))
                        var v28 uint16 = 32767
                        if v18 {
                            switch 1 {
                            case 32768:
                                v28 = 601
                            case 4294967295:
                                fmt.Println(v2, v28, v5)
                            }
                            switch "foo bar" {
                            case "hello":
                                var v29 uint16 = (v28 % ((0) | 1))
                                _ = v29
                            case "foo bar":
                                var v30 uint64 = 639
                                _ = v30
                            case "a":
                                var v31 uint64 = 2
                                _ = v31
                            }
                            fmt.Println(fi20, fi20, fi20)
                        }
                        if len(v16) > 0 {
                            v16[0] = (fi21 - fi21)
                        }
                        _ = v28
                    }
                    _ = v27
                } else {
                    v5 = (fi21 ^ fi20)
                    var v32 int = -63
                    _ = v32
                }
                v33 := GenS2{F0: 4294967295, F1: (v18 || v18)}
                _ = fi21
                _ = v33
            }
            switch v5 {
            case 1:
                var v34 int = v5
                _ = v34
            case 42:
                var v35 bool = (504 == 482)
                _ = v35
            case -9223372036854775808:
                v36 := GenS1{F0: v2}
                _ = v36
            }
        }
        _ = fi20
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 float32 = 3.14159
        _ = fi37
        _ = v38
    }
    switch v5 {
    case 32768:
        fmt.Println(v19, v19)
    case 128:
        var v39 int32 = -997
        _ = v39
    case 1:
        switch v5 {
        case -2147483648:
            _ = len(v1)
        default:
            if (!(!false)) {
                var v40 uint8 = 0
                _ = v40
            } else {
                switch (v5 % (((v5 & (v5 % ((v5) | 1)))) | 1)) {
                case -216:
                    if v18 {
                        for fi41 := 0; fi41 < 1; fi41++ {
                            v42 := GenS1{F0: (v2 & 2147483648)}
                            var v43 int16 = 1000
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        var v44 uint64 = 203
                        for fi45 := 0; fi45 < 16; fi45++ {
                            if true {
                                v2 += 0
                            }
                            _ = fi45
                        }
                        _ = v44
                    }
                case 127:
                    v18 = v18
                }
                var v46 float32 = 3.14159
                _ = v46
            }
        }
    }
    v19 = 0
    var v47 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
