package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 int8, p1 uint) string {
    var v4 uint32 = 1
    v5 := GenS1{F0: (p0 * (^(p0)))}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "a"
}

func main() {
    v1 := []int{}
    v2 := []int{2147483647, 631}
    v3 := make(map[string]int)
    var v6 bool = (!true)
    fmt.Println(v6)
    v6 = v6
    if v6 {
        for fi7 := 0; fi7 < 18; fi7++ {
            if true {
                for fi8 := 0; fi8 < 6; fi8++ {
                    v9 := GenS1{F0: 0}
                    fmt.Println(fi7)
                    _ = fi8
                    _ = v9
                }
            } else {
                if v6 {
                    for fi10 := 0; fi10 < 6; fi10++ {
                        v11 := GenS2{F0: (("" + "") + ("hello" + "a")), F1: 72}
                        v6 = false
                        _ = fi10
                        _ = v11
                    }
                    if v6 {
                        switch "a" {
                        case "\n\t":
                            var v12 float64 = ((1.5 - 1.5) * 1.5)
                            _ = v12
                        case "foo bar":
                            for fi13 := 0; fi13 < 20; fi13++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi13
                            }
                        }
                        if (1000 <= 1000) {
                            v14 := GenS1{F0: -11}
                            v6 = (true || (!(true && (4294967295 > 2147483647))))
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 3; fi15++ {
                            if ("a" < "\n\t") {
                                fmt.Println(fi15)
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                                fmt.Println(fi7, fi7, fi15)
                            } else {
                                v6 = v6
                                var v16 int32 = 2147483647
                                v16 -= (v16 % (((v16 & -2147483648)) | 1))
                                var v17 float64 = 1e10
                                switch fi7 {
                                case 128:
                                    v17 = v17
                                default:
                                    v16 = (v16 & (v16 | (-(v16))))
                                }
                                var v18 int = (fi15 ^ fi7)
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            _ = fi15
                        }
                    } else {
                        for fi19 := 0; fi19 < 5; fi19++ {
                            var v20 uint8 = 2
                            v6 = v6
                            _ = fi19
                            _ = v20
                        }
                        v21 := GenS2{F0: "a", F1: -1}
                        var v22 int32 = 2147483647
                        _ = v21
                        _ = v22
                    }
                } else {
                    var v23 int32 = 1
                    _ = v23
                }
            }
            _ = fi7
        }
    }
    switch 32767 {
    case -2147483648:
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 uint64 = 665
            switch (fi24 - fi24) {
            case 1:
                v6 = v6
            default:
                v25 -= (v25 >> ((38) & 63))
            }
            _ = fi24
            _ = v25
        }
    case 128:
        v6 = v6
    default:
        switch "a" {
        case "":
            v26 := GenS3{F0: -118, F1: 855}
            _ = v26
        case "a":
            var v27 uint16 = 42
            _ = v27
        }
    }
    v6 = false
    fmt.Println(v6)
    for fi28 := 0; fi28 < 3; fi28++ {
        fmt.Println(fi28, v6)
        for fi29 := 0; fi29 < 3; fi29++ {
            var v30 float64 = 1.5
            var v31 bool = true
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    v6 = v6
    v6 = v6
    if (1 == 127) {
        switch ("\n\t" + "foo bar") {
        case "a":
            for fi32 := 0; fi32 < 15; fi32++ {
                if false {
                    switch fi32 {
                    case -9223372036854775808:
                        switch "\n\t" {
                        case "":
                            if true {
                                var v33 string = "hello"
                                _ = v33
                            }
                        }
                    }
                }
                v6 = v6
                _ = fi32
            }
        case "":
            v34 := GenS2{F0: "a", F1: -1}
            _ = v34
        }
        fmt.Println(v6)
        v3["\n\t"] = 9223372036854775807
    } else {
        var v35 bool = v6
        _ = v35
    }
    v36 := GenS1{F0: 3}
    v37 := GenS1{F0: -80}
    v6 = v6
    for fi38 := 0; fi38 < 5; fi38++ {
        v6 = v6
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v6
}
