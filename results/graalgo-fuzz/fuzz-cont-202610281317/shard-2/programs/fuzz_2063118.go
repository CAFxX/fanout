package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 bool; F1 int16; F2 int }
type GenS6 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0(p0 int) uint32 {
    for fi4 := 0; fi4 < 10; fi4++ {
        if (-344 > 640) {
            for fi5 := 0; fi5 < 5; fi5++ {
                v6 := []int{32768, (-(p0)), (^(p0))}
                p0 = (p0 & (fi5 + fi5))
                _ = fi5
                _ = v6
            }
        } else {
            if true {
                if true {
                    v7 := GenS1{F0: 255}
                    _ = v7
                }
            }
        }
        _ = fi4
    }
    switch (p0 * (p0 >> ((255) & 63))) {
    case 18:
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := []int{}
            if false {
                if false {
                    v10 := GenS1{F0: 255}
                    switch "" {
                    case "":
                        if false {
                            fmt.Println(fi8, p0)
                            p0 = (^(p0))
                            var v11 int8 = -128
                            _ = v11
                        }
                    case "foo bar":
                        v12 := GenS1{F0: 0}
                        _ = v12
                    }
                    _ = v10
                }
            }
            _ = fi8
            _ = v9
        }
    default:
        var v13 int16 = 32767
        _ = v13
    }
    var v14 float64 = (0.1 + (0.1 + 1e10))
    fmt.Println(v14)
    _ = p0
    _ = v14
    return 2
}
func genFn1() string {
    v15 := GenS2{F0: -32768, F1: -108}
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint64 = 42
        _ = fi16
        _ = v17
    }
    _ = v15
    return "a"
}

func main() {
    var v1 int16 = 42
    var v2 uint8 = 1
    v3 := make(map[string]bool)
    for fi18 := 0; fi18 < 4; fi18++ {
        v1 = 42
        for fi19 := 0; fi19 < 5; fi19++ {
            switch fi18 {
            case 32767:
                fmt.Println(fi19, fi19)
            default:
                v20 := []int{}
                _ = v20
            }
            v21 := []int{fi19, 42, 1}
            var v22 string = "a"
            _ = fi19
            _ = v21
            _ = v22
        }
        _ = fi18
    }
    if ((false && true) && (!false)) {
        v23 := GenS2{F0: -1, F1: -81}
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        _ = len(v3)
        var v25 bool = false
        _ = fi24
        _ = v25
    }
    if false {
        for fi26 := 0; fi26 < 4; fi26++ {
            for fi27 := 0; fi27 < 2; fi27++ {
                v28 := []int{fi27}
                if (false && (801 != 1)) {
                    if (("a" + "") != "") {
                        v2 -= 30
                        switch 2147483647 {
                        case 1:
                            v2 = (v2 ^ (^(v2)))
                        default:
                            fmt.Println(v2, v2, v2)
                        }
                    } else {
                        delete(v3, "")
                        v1 += v1
                        var v29 int = (^(fi27))
                        _ = v29
                    }
                    fmt.Println(v2, fi26, v1)
                } else {
                    var v30 int64 = 42
                    var v31 int8 = -26
                    if true {
                        v31 -= (v31 / (((v31 - v31)) | 1))
                        if false {
                            v3[("hello" + "foo bar")] = false
                        } else {
                            v32 := GenS3{F0: "foo bar"}
                            v30 = 0
                            _ = v32
                        }
                        if false {
                            var v33 int = fi26
                            _ = v33
                        }
                    } else {
                        v34 := GenS4{F0: ((!(int32(v2) > 0)) && (255 < 879)), F1: 1}
                        if ((v1 / ((v1) | 1)) == (-(v1))) {
                            v1 = (^(v1))
                            fmt.Println(v31)
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                            if ((374 == 630) || true) {
                                var v35 int32 = 1
                                var v36 int8 = v31
                                _ = v35
                                _ = v36
                            } else {
                                v30 -= v30
                                v2 += 0
                                var v37 int32 = -2147483648
                                _ = v37
                            }
                        } else {
                            if true {
                                v31 -= (^(v31))
                                v3[("" + "hello")] = ((0.1 - 3.14159) == 1e10)
                            }
                            _ = len(v28)
                            var v38 int32 = -1
                            _ = v38
                        }
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                        _ = v34
                    }
                    _ = v30
                    _ = v31
                }
                v1 += v1
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        for fi39 := 0; fi39 < 10; fi39++ {
            var v40 int8 = 34
            switch (fi39 | 87) {
            case 32768:
                if (!(true || true)) {
                    v2 -= (v2 - (v2 ^ (v2 | v2)))
                    var v41 int8 = 42
                    _ = v41
                } else {
                    var v42 string = "\n\t"
                    v43 := GenS3{F0: "hello"}
                    _ = v42
                    _ = v43
                }
            }
            switch "" {
            case "foo bar":
                v40 -= v40
            }
            _ = fi39
            _ = v40
        }
    } else {
        if true {
            fmt.Println(v2)
        }
    }
    if true {
        switch "\n\t" {
        case "a":
            var v44 int16 = (v1 >> ((981) & 63))
            _ = v44
        case "\n\t":
            v2 -= (v2 + v2)
        case "foo bar":
            switch 128 {
            case 9223372036854775807:
                v45 := GenS5{F0: (4294967295 == 39), F1: 112, F2: 33}
                _ = v45
            case 127:
                var v46 int64 = -9223372036854775808
                _ = v46
            default:
                v2 -= 42
            }
        default:
            v2 += 127
        }
        for fi47 := 0; fi47 < 12; fi47++ {
            var v48 uint = 1
            if ((true && false) && (!true)) {
                if true {
                    var v49 uint16 = 568
                    _ = v49
                } else {
                    v50 := v3["a"]
                    var v51 bool = v50
                    v1 -= (v1 << ((v48) & 63))
                    _ = v50
                    _ = v51
                }
                fmt.Println(fi47)
            } else {
                switch ("hello" + ("foo bar" + "foo bar")) {
                case "":
                    v52 := []int{fi47, (fi47 ^ fi47), fi47, -786}
                    _ = v52
                case "hello":
                    for fi53 := 0; fi53 < 11; fi53++ {
                        switch 32768 {
                        case 128:
                            var v54 int = (fi47 << (((v48 & (v48 + v48))) & 63))
                            _ = v54
                        case -285:
                            var v55 uint32 = 447
                            _ = v55
                        }
                        v3[("foo bar" + "foo bar")] = (true || true)
                        _ = fi53
                    }
                }
            }
            v1 -= 1000
            _ = fi47
            _ = v48
        }
        var v56 float64 = (0.0 + 3.14159)
        _ = v56
    }
    var v57 uint16 = 545
    var v58 uint8 = 0
    v59 := GenS1{F0: v58}
    if true {
        for fi60 := 0; fi60 < 15; fi60++ {
            v61 := GenS6{F0: v1, F1: 42, F2: 125}
            _ = fi60
            _ = v61
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v57
    _ = v58
    _ = v59
}
