package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int32; F2 int16 }
type GenS5 struct { F0 string; F1 int32; F2 int64 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 int32; F1 int16; F2 bool }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{0}
    v3 := []int{2147483647}
    if false {
        for fi4 := 0; fi4 < 12; fi4++ {
            switch "hello" {
            case "foo bar":
                _ = len(v1)
            default:
                var v5 uint16 = 608
                _ = v5
            }
            _ = fi4
        }
        var v6 bool = ((-1 > 1) || (42 <= -107))
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 bool = (true && v6)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v6 = v6
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 string = "a"
                switch (("foo bar" + ("hello" + v11)) + (v11 + "\n\t")) {
                case "a":
                    if v6 {
                        v6 = v6
                        var v12 uint32 = 408
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 134, F1: 946}
                        _ = v13
                    }
                case "":
                    if (42 < -1) {
                        if true {
                            fmt.Println(fi7, fi10, v6)
                        }
                    } else {
                        v14 := GenS2{F0: 2, F1: 2147483647}
                        _ = v14
                    }
                }
                _ = fi10
                _ = v11
            }
            _ = fi7
        }
        _ = v6
    }
    var v15 uint64 = 1
    v15 = v15
    if (!false) {
        var v16 uint16 = 77
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v15 = 630
        if false {
            v18 := GenS1{F0: 4294967295, F1: v15}
            _ = v18
        } else {
            if false {
                if false {
                    v15 = v15
                    if true {
                        v15 -= (v15 & (v15 | 4294967295))
                    }
                } else {
                    if false {
                        var v19 string = "foo bar"
                        for fi20 := 0; fi20 < 2; fi20++ {
                            fmt.Println(v15, fi17)
                            var v21 uint32 = 2147483647
                            fmt.Println(v15, fi20)
                            v22 := GenS3{F0: 880}
                            var v23 int = (fi20 % ((fi20) | 1))
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        v15 -= 18446744073709551615
                        _ = v19
                    } else {
                        if true {
                            v24 := GenS4{F0: 18446744073709551615, F1: 1, F2: -250}
                            _ = v24
                        }
                    }
                    v25 := GenS5{F0: "a", F1: 1, F2: 76}
                    _ = v25
                }
                for fi26 := 0; fi26 < 14; fi26++ {
                    fmt.Println(fi26, fi17)
                    _ = fi26
                }
                v15 -= (v15 - v15)
            }
        }
        _ = fi17
    }
    var v27 int64 = 497
    var v28 int32 = -2147483648
    v15 -= (v15 ^ (v15 ^ v15))
    v15 = v15
    var v29 int = 1
    if ((v29 | (v29 * (-(v29)))) <= 32768) {
        v30 := GenS1{F0: 751, F1: (^(v15))}
        _ = len(v1)
        v2 = append(v2, (v29 ^ 1))
        _ = v30
    } else {
        if true {
            for fi31 := 0; fi31 < 1; fi31++ {
                v1["foo bar"] = 65535
                v32 := GenS1{F0: 652, F1: v15}
                var v33 uint16 = 255
                _ = fi31
                _ = v32
                _ = v33
            }
            switch ((("a" + ("hello" + "foo bar")) + "a") + ("hello" + ("" + ("hello" + ("a" + "\n\t"))))) {
            case "\n\t":
                if true {
                    for fi34 := 0; fi34 < 18; fi34++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v35 bool = (0.1 > (1e10 / 1e10))
                        v36 := GenS6{F0: 109, F1: "a"}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    var v37 bool = false
                    switch v29 {
                    case -141:
                        var v38 int = (-(v29))
                        _ = v38
                    case 128:
                        var v39 int64 = (^(v27))
                        _ = v39
                    }
                    if ((-656 <= 32767) && (v28 > v28)) {
                        for fi40 := 0; fi40 < 2; fi40++ {
                            var v41 uint32 = 65535
                            var v42 float32 = 0.0
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    } else {
                        var v43 uint16 = 32767
                        _ = v43
                    }
                    _ = v37
                }
            case "hello":
                v29 -= (v29 % (((v29 ^ -901)) | 1))
            case "foo bar":
                var v44 float64 = 1e10
                _ = v44
            }
        }
    }
    var v45 int16 = 39
    for fi46 := 0; fi46 < 16; fi46++ {
        if true {
            for fi47 := 0; fi47 < 1; fi47++ {
                v48 := GenS7{F0: v28, F1: v45, F2: (float64(0.1) > 1.5)}
                _ = fi47
                _ = v48
            }
        }
        var v49 int16 = v45
        _ = fi46
        _ = v49
    }
    v27 -= 42
    fmt.Println(v28)
    v28 = v28
    switch v29 {
    case 128:
        var v50 uint32 = 65535
        _ = v50
    case -408:
        var v51 string = ""
        _ = v51
    case 32768:
        v52 := GenS5{F0: "foo bar", F1: v28, F2: -1}
        _ = v52
    }
    v15 += (v15 | v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v45
}
