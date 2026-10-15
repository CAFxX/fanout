package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint32 }
func genFn0() int8 {
    if (!false) {
        var v7 int16 = -26
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := []int{}
            var v10 uint16 = 255
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        var v11 int32 = 42
        _ = v11
    }
    if true {
        if false {
            v12 := []int{}
            _ = v12
        } else {
            var v13 uint8 = 46
            for fi14 := 0; fi14 < 20; fi14++ {
                v15 := GenS1{F0: ("a" + ("" + ((("foo bar" + "\n\t") + "foo bar") + "\n\t"))), F1: fi14}
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        v16 := []int{128, 128}
        var v17 uint32 = 4294967295
        _ = v16
        _ = v17
    }
    var v18 float32 = 1e10
    _ = v18
    return 1
}

func main() {
    var v1 bool = false
    v2 := make(map[string]bool)
    var v3 uint32 = 681
    var v4 float32 = 1e10
    v5 := []uint32{}
    var v6 int16 = 600
    _ = len(v2)
    v19 := v2["foo bar"]
    v6 = -366
    v20 := GenS1{F0: "\n\t", F1: 29}
    v19 = (((-2147483648 != -2147483648) || (!v19)) || (0 <= 255))
    if v19 {
        if false {
            if (1 != 42) {
                v6 = (v6 | (v6 ^ -459))
                v21 := GenS1{F0: "\n\t", F1: 32767}
                for fi22 := 0; fi22 < 20; fi22++ {
                    for fi23 := 0; fi23 < 6; fi23++ {
                        var v24 uint = 2
                        switch (fi23 - (^(fi22))) {
                        case -1:
                            if v1 {
                                v3 += v3
                                if (v1 && v19) {
                                    var v25 int = fi23
                                    _ = v25
                                }
                                v19 = true
                            }
                        default:
                            switch (fi23 ^ fi22) {
                            case 127:
                                v1 = ((42 != 813) || v19)
                            case 571:
                                _ = len(v5)
                            case 128:
                                v3 += v3
                            }
                        }
                        v26 := GenS1{F0: "hello", F1: (fi22 % ((fi23) | 1))}
                        _ = fi23
                        _ = v24
                        _ = v26
                    }
                    switch fi22 {
                    case 1:
                        v27 := v2["\n\t"]
                        _ = v27
                    case 2147483647:
                        fmt.Println(v6, v19)
                    }
                    _ = fi22
                }
                _ = v21
            }
            for fi28 := 0; fi28 < 7; fi28++ {
                fmt.Println(v1)
                for fi29 := 0; fi29 < 11; fi29++ {
                    fmt.Println(v6, v1)
                    v5 = append(v5, v3)
                    if false {
                        if (-250 > -1) {
                            var v30 bool = v19
                            _ = v30
                        }
                        v31 := GenS2{F0: 127}
                        v32 := GenS3{F0: (^(v3))}
                        _ = v31
                        _ = v32
                    } else {
                        if ((v1 && ((fi28 <= 293) && v1)) && (v6 >= v6)) {
                            v33 := v2["\n\t"]
                            fmt.Println(v1)
                            var v34 string = "\n\t"
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = fi29
                }
                _ = fi28
            }
        } else {
            for fi35 := 0; fi35 < 15; fi35++ {
                v36 := v2[("a" + "")]
                v19 = false
                _ = fi35
                _ = v36
            }
        }
        v2[("\n\t" + ("foo bar" + "\n\t"))] = v19
    } else {
        if true {
            v6 = -227
            var v37 uint32 = v3
            var v38 float32 = v4
            _ = v37
            _ = v38
        }
        v19 = false
        v19 = v19
    }
    var v39 int16 = v6
    for fi40 := 0; fi40 < 7; fi40++ {
        var v41 uint64 = 187
        for fi42 := 0; fi42 < 10; fi42++ {
            switch "a" {
            case "":
                var v43 uint32 = 4294967295
                _ = v43
            case "hello":
                var v44 int8 = genFn0()
                _ = v44
            case "a":
                v19 = (v4 == v4)
            }
            _ = fi42
        }
        var v45 uint32 = 228
        _ = fi40
        _ = v41
        _ = v45
    }
    for fi46 := 0; fi46 < 19; fi46++ {
        v47 := GenS4{F0: -2147483648, F1: 787}
        v39 = -1
        if (v1 && (42 < 340)) {
            if len(v5) > 0 {
                v5[0] = (^(v3))
            }
            var v48 int = fi46
            switch ("hello" + "foo bar") {
            case "\n\t":
                fmt.Println(v3)
            case "hello":
                delete(v2, "a")
            case "":
                v49 := GenS5{F0: 127, F1: (v3 * v3)}
                _ = v49
            default:
                for fi50 := 0; fi50 < 20; fi50++ {
                    if v19 {
                        switch (^(v48)) {
                        case -596:
                            _ = len(v2)
                        }
                        v6 = v6
                    }
                    _ = fi50
                }
            }
            _ = v48
        }
        _ = fi46
        _ = v47
    }
    var v51 int32 = -370
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
