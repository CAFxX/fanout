package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32; F1 int16; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 int32; F2 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 int32 }
func genFn0(p0 string, p1 int8, p2 uint8) uint64 {
    p2 = p2
    v4 := GenS1{F0: 1, F1: -1}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 250
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = 65535
    v3 := make(map[string]bool)
    var v5 string = "a"
    v6 := GenS1{F0: (v2 / ((v2) | 1)), F1: 1}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int32 = 2147483647
        if true {
            switch "hello" {
            case "\n\t":
                v9 := []int{(fi7 & fi7), fi7, fi7}
                _ = v9
            case "foo bar":
                fmt.Println(v1)
            case "a":
                if (true || true) {
                    v10 := GenS1{F0: v2, F1: 141}
                    _ = v10
                }
            }
            if true {
                v11 := GenS1{F0: 593, F1: (v8 - v8)}
                v8 = 775
                _ = v11
            } else {
                if true {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v2 += 0
                        _ = fi12
                    }
                    v13 := []int{(fi7 ^ 437)}
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 int32 = v8
                if true {
                    v1 = 94
                } else {
                    if (!((^(v15)) == v8)) {
                        v1 = v1
                    } else {
                        var v16 string = ("" + ("hello" + v5))
                        delete(v3, v5)
                        var v17 uint = v1
                        v17 -= v17
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS2{F0: (^(v15)), F1: -1, F2: (false && ((!((v1 % ((v1) | 1)) >= v1)) && true))}
                    _ = v18
                }
                v5 = (("a" + (("a" + (v5 + v5)) + ((v5 + v5) + v5))) + (((v5 + ("foo bar" + "a")) + v5) + v5))
                _ = fi14
                _ = v15
            }
        } else {
            var v19 string = (((v5 + (v5 + "\n\t")) + "a") + v5)
            var v20 int32 = v8
            v2 += (v2 | v2)
            _ = v19
            _ = v20
        }
        v21 := GenS3{F0: 42}
        _ = fi7
        _ = v21
        _ = v8
    }
    v22 := GenS4{F0: 332}
    v23 := GenS4{F0: -1}
    v24 := GenS5{F0: 127, F1: -1, F2: (v1 << ((v1) & 63))}
    var v25 float32 = 3.14159
    for fi26 := 0; fi26 < 10; fi26++ {
        v1 += v1
        v27 := GenS4{F0: 255}
        fmt.Println(v5)
        _ = fi26
        _ = v27
    }
    if false {
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 16; fi29++ {
                v2 += (v2 + (^(v2)))
                _ = fi29
            }
            var v30 int32 = -2147483648
            _ = fi28
            _ = v30
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 uint64 = 541
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 uint8 = uint8(v2)
            fmt.Println(fi33, v1)
            _ = fi33
            _ = v34
        }
    }
    var v35 float32 = v25
    var v36 int = -889
    v36 = 42
    switch (v36 * v36) {
    case 0:
        var v37 uint64 = 62
        _ = v37
    case 2147483647:
        if false {
            var v38 int8 = 100
            _ = v38
        } else {
            v1 += 18446744073709551615
            if false {
                switch (^(v36)) {
                case 0:
                    var v39 uint64 = 18446744073709551615
                    _ = v39
                case -9223372036854775808:
                    fmt.Println(v36, v36)
                case 9223372036854775807:
                    v35 -= float32((0.0 - 0.0))
                }
            } else {
                switch v5 {
                case "\n\t":
                    switch ((v5 + (v5 + (v5 + v5))) + v5) {
                    case "hello":
                        if false {
                            v5 = ((("\n\t" + v5) + v5) + (v5 + "\n\t"))
                        } else {
                            if (2 < 1) {
                                v40 := GenS4{F0: -2147483648}
                                v1 = v1
                                for fi41 := 0; fi41 < 7; fi41++ {
                                    if true {
                                        var v42 float64 = (0.1 * float64((v25 + (v25 - v35))))
                                        for fi43 := 0; fi43 < 3; fi43++ {
                                            v36 = (fi41 << ((v1) & 63))
                                            _ = fi43
                                        }
                                        v44 := GenS6{F0: 1}
                                        _ = v42
                                        _ = v44
                                    }
                                    _ = fi41
                                }
                                _ = v40
                            }
                        }
                    case "":
                        fmt.Println(v25, v2, v2)
                    case "a":
                        if false {
                            if (true || ((v2 ^ v2) < (^(v2)))) {
                                fmt.Println(v2)
                                v5 = ("" + "")
                                for fi45 := 0; fi45 < 5; fi45++ {
                                    var v46 int64 = 1
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            for fi47 := 0; fi47 < 11; fi47++ {
                                if (1 > 42) {
                                    v48 := v3[v5]
                                    _ = v48
                                }
                                v49 := []int{(fi47 / (((-(fi47))) | 1)), fi47}
                                _ = fi47
                                _ = v49
                            }
                            v50 := []int{149, (v36 / ((v36) | 1))}
                            _ = v50
                        } else {
                            v36 -= v36
                            var v51 int16 = -1
                            _ = v51
                        }
                    default:
                        var v52 int64 = 0
                        _ = v52
                    }
                case "":
                    var v53 float32 = v25
                    _ = v53
                case "hello":
                    v54 := []int{v36, v36, v36}
                    _ = v54
                }
                switch "a" {
                case "a":
                    for fi55 := 0; fi55 < 4; fi55++ {
                        v56 := GenS2{F0: 432, F1: 1, F2: true}
                        v57 := GenS7{F0: v5, F1: 42}
                        fmt.Println(fi55, v1, v25)
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                }
            }
        }
    case -292:
        v58 := GenS4{F0: 751}
        _ = v58
    default:
        if ((32767 < 42) || true) {
            delete(v3, (((((v5 + v5) + v5) + v5) + v5) + (("\n\t" + v5) + ((("\n\t" + v5) + "foo bar") + ((v5 + "\n\t") + (v5 + "foo bar"))))))
            fmt.Println(v35, v36)
        } else {
            var v59 int32 = -1
            _ = v59
        }
    }
    var v60 int = 42
    v36 = (^(v36))
    v61 := GenS3{F0: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
