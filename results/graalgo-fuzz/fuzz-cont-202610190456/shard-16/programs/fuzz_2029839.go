package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }
type GenS2 struct { F0 int; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 1
    var v3 int64 = -659
    switch -2147483648 {
    case 2147483647:
        v4 := GenS1{F0: "hello", F1: 1, F2: -426}
        _ = v4
    case 699:
        v2 = (v2 + 42)
    default:
        v3 = 0
    }
    if (!false) {
        switch 32767 {
        case -2147483648:
            v3 -= v3
        case 32767:
            for fi5 := 0; fi5 < 12; fi5++ {
                fmt.Println(v2, fi5)
                _ = fi5
            }
        default:
            v6 := []int{402, -1, 32767, 9223372036854775807}
            _ = v6
        }
    } else {
        var v7 uint64 = (^(v2))
        if false {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: "a", F1: 65535, F2: -32768}
                v7 = (v7 / ((v7) | 1))
                _ = fi8
                _ = v9
            }
            var v10 uint32 = 626
            v11 := []int{}
            _ = v10
            _ = v11
        } else {
            var v12 bool = (!true)
            if v12 {
                v13 := GenS2{F0: 0, F1: 14, F2: 1}
                _ = v13
            }
            var v14 uint = 255
            _ = v12
            _ = v14
        }
        _ = v7
    }
    v15 := []int{}
    v3 -= v3
    switch ("a" + "foo bar") {
    case "foo bar":
        if false {
            switch "\n\t" {
            case "\n\t":
                if ((true && true) || false) {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            var v18 int64 = v3
                            v3 = v18
                            _ = len(v1)
                            if len(v15) > 0 {
                                v15[0] = fi16
                            }
                            var v19 int8 = -128
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        var v20 string = ""
                        _ = fi16
                        _ = v20
                    }
                    var v21 float64 = 0.0
                    var v22 uint64 = v2
                    _ = v21
                    _ = v22
                } else {
                    switch "foo bar" {
                    case "a":
                        if ((1e10 + 1e10) == (3.14159 / (-2.25 / float32(1e10)))) {
                            var v23 uint8 = 237
                            var v24 uint64 = v2
                            v3 += v3
                            _ = v23
                            _ = v24
                        }
                    case "hello":
                        var v25 string = "hello"
                        _ = v25
                    case "\n\t":
                        if (-2147483648 >= -2147483648) {
                            for fi26 := 0; fi26 < 17; fi26++ {
                                var v27 uint64 = 771
                                _ = fi26
                                _ = v27
                            }
                            if (false && false) {
                                v2 = v2
                                for fi28 := 0; fi28 < 13; fi28++ {
                                    v3 -= v3
                                    var v29 uint8 = 128
                                    _ = fi28
                                    _ = v29
                                }
                            }
                        } else {
                            for fi30 := 0; fi30 < 3; fi30++ {
                                v3 -= -1
                                v31 := GenS3{F0: -128}
                                _ = fi30
                                _ = v31
                            }
                            for fi32 := 0; fi32 < 12; fi32++ {
                                for fi33 := 0; fi33 < 14; fi33++ {
                                    v3 = v3
                                    v3 += (v3 & 0)
                                    if (((3.14159 / 0.1) * 1e10) >= 0.0) {
                                        _ = len(v15)
                                    } else {
                                        switch fi33 {
                                        case 4294967295:
                                            switch "\n\t" {
                                            case "\n\t":
                                                if (float64(3.14159) >= 1e10) {
                                                    fmt.Println(v2)
                                                    var v34 int8 = -1
                                                    v34 += 9
                                                    if ((42 < 255) || true) {
                                                        v34 += v34
                                                        v35 := GenS2{F0: (fi32 % ((fi33) | 1)), F1: 42, F2: 154}
                                                        _ = v35
                                                    }
                                                    _ = v34
                                                }
                                            }
                                        }
                                        v3 -= 1
                                        var v36 float32 = float32(1.5)
                                        _ = v36
                                    }
                                    _ = fi33
                                }
                                _ = fi32
                            }
                            switch "foo bar" {
                            case "hello":
                                var v37 uint32 = 2147483648
                                _ = v37
                            case "foo bar":
                                var v38 bool = false
                                _ = v38
                            }
                        }
                    }
                    for fi39 := 0; fi39 < 16; fi39++ {
                        var v40 string = ""
                        _ = fi39
                        _ = v40
                    }
                    var v41 int64 = 42
                    _ = v41
                }
            case "a":
                v42 := GenS3{F0: 0}
                _ = v42
            }
        } else {
            v2 = (v2 + v2)
        }
    case "\n\t":
        fmt.Println(v2, v3)
    case "":
        v2 -= v2
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        var v44 bool = false
        v2 -= 0
        _ = fi43
        _ = v44
    }
    var v45 bool = (!true)
    for fi46 := 0; fi46 < 3; fi46++ {
        v45 = (v45 || false)
        _ = len(v15)
        _ = fi46
    }
    var v47 string = "a"
    v3 = v3
    v45 = v45
    v1[2] = (v47 + ("" + v47))
    switch v47 {
    case "foo bar":
        switch 4294967295 {
        case 911:
            var v48 bool = true
            _ = v48
        case 1:
            for fi49 := 0; fi49 < 12; fi49++ {
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi49
            }
        default:
            var v50 float32 = (0.1 / ((1.5 + (1e10 - -2.25)) / (-2.25 + float32((3.14159 / 1.0)))))
            _ = v50
        }
    default:
        switch (v47 + v47) {
        case "hello":
            var v51 string = (v47 + (v47 + ""))
            _ = v51
        }
    }
    delete(v1, 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v45
    _ = v47
}
