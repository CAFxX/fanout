package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint32; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8; F1 int8; F2 uint32 }
func genFn0() uint16 {
    if ("foo bar" < "\n\t") {
        var v6 uint16 = 0
        var v7 string = "foo bar"
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 532, F1: 255, F2: 67}
        var v9 uint32 = 47
        var v10 int32 = 2147483647
        _ = v10
        _ = v8
        _ = v9
    }
    switch "a" {
    case "hello":
        v11 := GenS2{F0: 873, F1: 979, F2: false}
        _ = v11
    default:
        var v12 uint8 = 31
        _ = v12
    }
    var v13 uint32 = 2147483648
    _ = v13
    return 0
}

func main() {
    var v1 int = -686
    v2 := make(map[int]uint32)
    var v3 uint = 636
    v4 := make(map[string]uint32)
    v5 := []string{("\n\t" + "hello"), "\n\t", ((("\n\t" + "a") + "a") + "foo bar"), "\n\t"}
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS3{F0: "hello", F1: uint16(v3)}
        v2[(v1 ^ fi14)] = 4294967295
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            v3 = (^(v3))
            if false {
                var v18 float64 = float64(1.5)
                var v19 int16 = -564
                var v20 uint16 = 255
                v19 += (v19 - 1)
                v21 := GenS2{F0: 2, F1: 663, F2: (594 < 573)}
                _ = v18
                _ = v19
                _ = v20
                _ = v21
            }
            switch (fi17 + fi17) {
            case 4294967295:
                switch fi16 {
                case 1:
                    v1 -= (^(fi17))
                    v3 += v3
                }
            case -9223372036854775808:
                var v22 bool = false
                _ = v22
            }
            _ = fi17
        }
        v23 := GenS3{F0: "", F1: 255}
        for fi24 := 0; fi24 < 8; fi24++ {
            v3 = 255
            v4[""] = 65535
            v3 = 4294967295
            _ = fi24
        }
        _ = fi16
        _ = v23
    }
    fmt.Println(v3, v3)
    var v25 string = ("foo bar" + ("foo bar" + (("a" + "a") + "a")))
    for fi26 := 0; fi26 < 19; fi26++ {
        _ = len(v5)
        _ = fi26
    }
    if false {
        if false {
            v3 = v3
            fmt.Println(v25)
            for fi27 := 0; fi27 < 13; fi27++ {
                v28 := GenS4{F0: 0, F1: 42, F2: 0}
                _ = fi27
                _ = v28
            }
        } else {
            var v29 uint64 = 1
            for fi30 := 0; fi30 < 7; fi30++ {
                if false {
                    v29 += (v29 | (v29 + v29))
                    if true {
                        var v31 int = (fi30 + -9223372036854775808)
                        v32 := GenS4{F0: 42, F1: -1, F2: 1}
                        _ = v31
                        _ = v32
                    } else {
                        v25 = v25
                    }
                    for fi33 := 0; fi33 < 13; fi33++ {
                        switch v25 {
                        case "hello":
                            switch (v25 + (v25 + v25)) {
                            case "foo bar":
                                v5 = append(v5, (v25 + (v25 + (v25 + ((v25 + "\n\t") + v25)))))
                                v34 := GenS1{F0: 130, F1: 2, F2: 35}
                                _ = v34
                            }
                        case "\n\t":
                            var v35 int8 = 127
                            _ = v35
                        case "a":
                            fmt.Println(fi33, fi30, v29)
                        default:
                            if false {
                                var v36 int32 = 856
                                if (false || false) {
                                    var v37 int8 = 100
                                    _ = v37
                                } else {
                                    if false {
                                        v1 += (fi33 - fi30)
                                        v1 -= v1
                                        if len(v5) > 0 {
                                            v5[0] = "a"
                                        }
                                    }
                                    var v38 uint64 = (^(v29))
                                    v29 = 2
                                    _ = v38
                                }
                                _ = v36
                            }
                        }
                        var v39 bool = false
                        _ = fi33
                        _ = v39
                    }
                }
                _ = fi30
            }
            var v40 int = v1
            _ = v29
            _ = v40
        }
        v25 = ""
    } else {
        if false {
            fmt.Println(v1)
        }
        var v41 float64 = -2.25
        _ = v41
    }
    if false {
        v42 := GenS4{F0: 0, F1: -9, F2: 2147483647}
        var v43 uint = (v3 + 18446744073709551615)
        _ = v42
        _ = v43
    }
    v25 = ((v25 + "") + v25)
    _ = len(v2)
    var v44 bool = true
    var v45 string = ("hello" + (v25 + ("a" + v25)))
    v46 := GenS3{F0: v45, F1: 42}
    fmt.Println(v1)
    var v47 uint8 = 42
    var v48 uint = 2
    fmt.Println(v3, v45)
    v49 := GenS2{F0: 4294967295, F1: 1, F2: ((1e10 != 1e10) && (1000 == 120))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
