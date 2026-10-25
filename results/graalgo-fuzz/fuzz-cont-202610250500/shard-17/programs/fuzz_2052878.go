package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 uint64; F2 string }
func genFn0() uint64 {
    var v4 uint16 = 32767
    v4 += (v4 % ((v4) | 1))
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{fi5, (fi5 ^ (fi5 * (fi5 >> ((239) & 63))))}
        for fi7 := 0; fi7 < 2; fi7++ {
            switch (fi5 ^ (fi5 | (fi5 ^ fi7))) {
            case 60:
                fmt.Println(v4)
            case 42:
                switch -1 {
                case 9223372036854775807:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case 32768:
                    var v8 bool = true
                    _ = v8
                case 942:
                    v4 += v4
                default:
                    v9 := GenS1{F0: 179, F1: 1}
                    _ = v9
                }
            default:
                switch "a" {
                case "a":
                    var v10 bool = true
                    _ = v10
                case "hello":
                    fmt.Println(fi7)
                case "\n\t":
                    v4 -= v4
                }
            }
            if true {
                if true {
                    v4 += (^(v4))
                    v4 -= v4
                }
                var v11 float32 = 0.1
                _ = v11
            } else {
                _ = len(v6)
                var v12 uint16 = v4
                v12 += (v12 % ((v4) | 1))
                _ = v12
            }
            v4 = 2
            _ = fi7
        }
        if true {
            v13 := GenS1{F0: 160, F1: 42}
            _ = v13
        }
        _ = fi5
        _ = v6
    }
    v4 = (^(v4))
    if false {
        if (!true) {
            v14 := []int{}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 13; fi15++ {
                v4 = 0
                for fi16 := 0; fi16 < 10; fi16++ {
                    v4 -= (v4 / (((v4 % ((v4) | 1))) | 1))
                    _ = fi16
                }
                v17 := []int{}
                _ = fi15
                _ = v17
            }
        }
        v18 := GenS2{F0: 1, F1: -702, F2: uint64(v4)}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 1; fi20++ {
                if false {
                    v4 += v4
                    v4 += v4
                } else {
                    v21 := []int{1, 4294967295}
                    _ = v21
                }
                _ = fi20
            }
            switch "foo bar" {
            case "foo bar":
                var v22 uint32 = 1
                _ = v22
            }
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := []int{fi23}
                if true {
                    if (false || false) {
                        var v25 float64 = (1e10 + 3.14159)
                        switch -1 {
                        case -1:
                            switch fi19 {
                            case 2147483647:
                                v26 := GenS2{F0: 127, F1: 4294967295, F2: 2}
                                _ = v26
                            case 128:
                                if (4294967295 < 265) {
                                    switch ("\n\t" + "\n\t") {
                                    case "":
                                        v4 = v4
                                    case "foo bar":
                                        v4 = (v4 + 2)
                                    case "a":
                                        var v27 int16 = -604
                                        _ = v27
                                    default:
                                        var v28 int8 = 25
                                        _ = v28
                                    }
                                    v29 := GenS2{F0: 0, F1: (fi19 >> ((42) & 63)), F2: 130}
                                    _ = v29
                                } else {
                                    v30 := GenS3{F0: 32767, F1: (v4 & 65535)}
                                    _ = v30
                                }
                            default:
                                v25 += float64((0.0 / 1.5))
                            }
                        case 2147483647:
                            v31 := GenS4{F0: 4294967295}
                            _ = v31
                        }
                        _ = v25
                    }
                }
                _ = fi23
                _ = v24
            }
            _ = fi19
        }
    }
    _ = v4
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 840
    var v3 int64 = -1
    switch -9223372036854775808 {
    case 4294967295:
        v3 += -1
    case 382:
        v2 -= v2
    case 127:
        switch 32767 {
        case 7:
            v3 -= v3
        default:
            for fi32 := 0; fi32 < 11; fi32++ {
                v33 := GenS5{F0: false, F1: 1, F2: "hello"}
                _ = fi32
                _ = v33
            }
        }
    }
    v3 += v3
    var v34 uint16 = 280
    var v35 int8 = 127
    for fi36 := 0; fi36 < 4; fi36++ {
        if true {
            v37 := []int{(fi36 % (((fi36 >> ((65535) & 63))) | 1)), 127, -971, (fi36 << ((42) & 63))}
            var v38 uint = (v2 / (((v2 << (((v2 + 42)) & 63))) | 1))
            v1[""] = ("\n\t" + "foo bar")
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 17; fi39++ {
            var v40 string = "hello"
            v41 := GenS5{F0: (fi36 <= (fi36 << (((v2 / ((v2) | 1))) & 63))), F1: uint64(v34), F2: v40}
            fmt.Println(fi39)
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi36
    }
    v2 = v2
    var v42 uint = (v2 << (((v2 | 42)) & 63))
    var v43 int8 = -1
    v43 += -128
    v44 := v1[(("foo bar" + ("\n\t" + "foo bar")) + "")]
    var v45 string = v44
    var v46 uint16 = v34
    v44 = "foo bar"
    var v47 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
