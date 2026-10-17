package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int32; F2 uint }

func main() {
    v1 := []uint32{2}
    var v2 uint64 = 2
    var v3 float64 = 1e10
    v4 := []int{127, 9223372036854775807, 32767, -9223372036854775808}
    v5 := make(map[uint32]string)
    v6 := []uint32{722, 2, 0}
    if ((!false) && true) {
        var v7 int8 = 100
        switch 638 {
        case -2147483648:
            if false {
                if true {
                    if true {
                        delete(v5, 2)
                        var v8 string = ((("" + ("foo bar" + "")) + "") + "foo bar")
                        _ = v8
                    }
                }
            } else {
                if true {
                    var v9 int16 = 0
                    _ = v9
                } else {
                    if true {
                        if (true || true) {
                            var v10 int32 = -1
                            v11 := GenS1{F0: 2147483647, F1: (("\n\t" + "hello") + ("foo bar" + "")), F2: v2}
                            _ = v10
                            _ = v11
                        } else {
                            for fi12 := 0; fi12 < 10; fi12++ {
                                switch 961 {
                                case -9223372036854775808:
                                    v13 := GenS2{F0: 2147483648}
                                    _ = v13
                                case 705:
                                    v3 = v3
                                case 1:
                                    for fi14 := 0; fi14 < 11; fi14++ {
                                        var v15 string = ""
                                        _ = fi14
                                        _ = v15
                                    }
                                }
                                var v16 uint64 = v2
                                for fi17 := 0; fi17 < 20; fi17++ {
                                    var v18 int32 = 4
                                    if false {
                                        switch ("foo bar" + "\n\t") {
                                        case "foo bar":
                                            delete(v5, 2)
                                        case "hello":
                                            var v19 uint16 = 42
                                            _ = v19
                                        }
                                        v3 = 3.14159
                                    }
                                    if true {
                                        var v20 uint32 = 4294967295
                                        v7 = v7
                                        v21 := GenS1{F0: v20, F1: "\n\t", F2: 42}
                                        _ = v20
                                        _ = v21
                                    }
                                    _ = fi17
                                    _ = v18
                                }
                                _ = fi12
                                _ = v16
                            }
                        }
                    } else {
                        if false {
                            v22 := GenS3{F0: 150}
                            _ = v22
                        }
                    }
                }
            }
        case 32767:
            v1 = append(v1, 688)
        }
        fmt.Println(v2)
        _ = v7
    } else {
        var v23 float64 = v3
        _ = v23
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v24 uint = 4294967295
    if (-73 > 0) {
        var v25 string = "\n\t"
        _ = v25
    }
    v26 := GenS4{F0: ((v24 << ((4294967295) & 63)) < (v24 * (v24 / ((v24) | 1)))), F1: 328, F2: 42}
    var v27 uint64 = (v2 ^ (v2 ^ (v2 & v2)))
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 int32 = -2147483648
        _ = fi28
        _ = v29
    }
    if true {
        var v30 float64 = (v3 - -2.25)
        var v31 uint8 = 0
        for fi32 := 0; fi32 < 10; fi32++ {
            if true {
                switch "\n\t" {
                case "":
                    var v33 uint8 = (v31 % ((128) | 1))
                    _ = v33
                }
                var v34 int = (fi32 % ((fi32) | 1))
                _ = v34
            } else {
                _ = len(v5)
                for fi35 := 0; fi35 < 5; fi35++ {
                    v5[2] = "foo bar"
                    fmt.Println(v3, fi35, v27)
                    _ = fi35
                }
            }
            v36 := GenS5{F0: 830, F1: 42, F2: v24}
            _ = fi32
            _ = v36
        }
        _ = v30
        _ = v31
    } else {
        var v37 uint16 = 2
        switch "" {
        case "":
            var v38 uint64 = (v27 + 846)
            _ = v38
        }
        for fi39 := 0; fi39 < 12; fi39++ {
            v37 += v37
            for fi40 := 0; fi40 < 6; fi40++ {
                var v41 int = fi39
                v41 = (v41 ^ v41)
                v27 = v2
                _ = fi40
                _ = v41
            }
            _ = fi39
        }
        _ = v37
    }
    if (false || (127 == 48)) {
        var v42 int64 = 1
        for fi43 := 0; fi43 < 11; fi43++ {
            var v44 string = "\n\t"
            _ = fi43
            _ = v44
        }
        var v45 uint64 = 0
        _ = v42
        _ = v45
    } else {
        var v46 string = ""
        v47 := GenS1{F0: 0, F1: ((v46 + v46) + v46), F2: v27}
        fmt.Println(v24)
        _ = v46
        _ = v47
    }
    var v48 uint8 = 94
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
