package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (9223372036854775807 != 127)
    var v3 uint16 = 1
    var v4 uint16 = (v3 & 65535)
    switch -1 {
    case 32768:
        v3 -= 1
    case -309:
        for fi5 := 0; fi5 < 3; fi5++ {
            switch "foo bar" {
            case "a":
                v2 = v2
            default:
                v3 = (v4 / ((v4) | 1))
            }
            var v6 int8 = -128
            var v7 uint8 = 127
            _ = fi5
            _ = v6
            _ = v7
        }
    case 4294967295:
        _ = len(v1)
    }
    switch -832 {
    case -9223372036854775808:
        switch "foo bar" {
        case "hello":
            v4 = v4
        case "a":
            if true {
                for fi8 := 0; fi8 < 9; fi8++ {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        if (!(!(-2147483648 != 1))) {
                            if v2 {
                                v4 = 2
                                v3 += v3
                                v4 -= v4
                            } else {
                                if ((v4 * v4) != 0) {
                                    var v10 int8 = 42
                                    if (!v2) {
                                        if false {
                                            v11 := GenS1{F0: -782, F1: v10, F2: 42}
                                            v12 := v1[715]
                                            v12 -= v12
                                            _ = v11
                                            _ = v12
                                        }
                                    } else {
                                        var v13 int16 = 1000
                                        if (0 == -9223372036854775808) {
                                            var v14 int32 = -1
                                            _ = v14
                                        }
                                        v10 += v10
                                        _ = v13
                                    }
                                    if (!v2) {
                                        var v15 uint8 = 42
                                        v15 -= (v15 / ((255) | 1))
                                        v15 = 127
                                        _ = v15
                                    } else {
                                        v10 -= -128
                                        if ("foo bar" >= "") {
                                            v16 := GenS1{F0: 42, F1: v10, F2: -341}
                                            _ = v16
                                        } else {
                                            fmt.Println(fi8)
                                        }
                                    }
                                    _ = v10
                                } else {
                                    v17 := v1[2]
                                    _ = v17
                                }
                            }
                            var v18 uint = 65535
                            var v19 int32 = 79
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS1{F0: 177, F1: 0, F2: 276}
                        var v21 int16 = 1
                        _ = fi9
                        _ = v20
                        _ = v21
                    }
                    if (31 != 861) {
                        for fi22 := 0; fi22 < 13; fi22++ {
                            v23 := GenS2{F0: ("foo bar" + ""), F1: -2147483648}
                            _ = fi22
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v25 := GenS2{F0: "hello", F1: 42}
                            _ = fi24
                            _ = v25
                        }
                        var v26 int32 = -2147483648
                        _ = v26
                    }
                    if (v4 != (v4 << ((255) & 63))) {
                        fmt.Println(v3, v2)
                        v27 := []int{(fi8 ^ 9223372036854775807), 42, fi8, (fi8 << ((453) & 63))}
                        _ = v27
                    }
                    _ = fi8
                }
                var v28 uint16 = v3
                v29 := v1[2147483647]
                _ = v28
                _ = v29
            }
        case "foo bar":
            if (0 != 2147483647) {
                for fi30 := 0; fi30 < 8; fi30++ {
                    v4 = 0
                    _ = fi30
                }
                for fi31 := 0; fi31 < 8; fi31++ {
                    v2 = true
                    var v32 int8 = -128
                    for fi33 := 0; fi33 < 20; fi33++ {
                        v3 = v3
                        _ = fi33
                    }
                    _ = fi31
                    _ = v32
                }
            } else {
                for fi34 := 0; fi34 < 5; fi34++ {
                    var v35 int64 = -1
                    var v36 bool = v2
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        }
    case 731:
        var v37 uint64 = 936
        _ = v37
    case 2147483647:
        v38 := GenS1{F0: 449, F1: -77, F2: 832}
        _ = v38
    default:
        v39 := GenS2{F0: "", F1: 42}
        _ = v39
    }
    if (9223372036854775807 > 877) {
        if v2 {
            v40 := GenS2{F0: ("foo bar" + ((("hello" + "foo bar") + "") + "hello")), F1: 0}
            var v41 int32 = 0
            v4 = v3
            _ = v40
            _ = v41
        }
        v4 = (^(v3))
    }
    v42 := []int{-834}
    var v43 int64 = 0
    var v44 uint8 = 253
    var v45 float64 = 0.1
    switch ("hello" + "") {
    case "a":
        fmt.Println(v3, v44)
    case "hello":
        _ = len(v1)
    case "":
        v4 = v3
    default:
        fmt.Println(v3, v4, v44)
    }
    switch (("a" + ("hello" + "a")) + ("\n\t" + ("" + "hello"))) {
    case "\n\t":
        fmt.Println(v45)
    case "hello":
        var v46 int32 = 42
        _ = v46
    }
    var v47 bool = (4294967295 <= 0)
    var v48 int32 = 42
    var v49 int = 657
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v49
}
