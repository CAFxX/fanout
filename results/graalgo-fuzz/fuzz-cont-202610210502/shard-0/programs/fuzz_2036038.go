package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 bool; F2 bool }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 int64; F1 uint32; F2 int8 }
type GenS9 struct { F0 int16 }
func genFn0(p0 string, p1 string, p2 uint16) uint16 {
    if (335 < 65535) {
        fmt.Println(p2, p2)
        switch 127 {
        case 4294967295:
            v4 := []int{127, -2147483648, 9223372036854775807}
            _ = v4
        default:
            for fi5 := 0; fi5 < 10; fi5++ {
                if true {
                    if ((p1 <= p1) && (4294967295 <= 0)) {
                        if true {
                            v6 := GenS1{F0: 1}
                            var v7 uint = 255
                            _ = v6
                            _ = v7
                        }
                        var v8 int32 = -442
                        _ = v8
                    }
                } else {
                    p1 = ("hello" + (p1 + p0))
                }
                fmt.Println(p1, p2, p0)
                v9 := []int{(fi5 - fi5)}
                _ = fi5
                _ = v9
            }
        }
    } else {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := []int{fi10, fi10}
            _ = fi10
            _ = v11
        }
    }
    var v12 int32 = 1
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p2
}
func genFn1() int {
    if (("a" + (("hello" + "a") + "foo bar")) == "hello") {
        var v13 float64 = (1.5 * ((3.14159 - 1e10) * (3.14159 - -2.25)))
        _ = v13
    }
    switch "hello" {
    case "foo bar":
        v14 := []int{}
        _ = v14
    case "":
        if ((3.14159 != 0.1) && false) {
            if (-128 <= -128) {
                if (!(false || true)) {
                    v15 := GenS2{F0: 32767, F1: 1}
                    if (-866 >= -880) {
                        var v16 int16 = 32767
                        var v17 uint32 = 2147483648
                        _ = v16
                        _ = v17
                    }
                    _ = v15
                }
            } else {
                var v18 int16 = -598
                v19 := GenS3{F0: 543}
                _ = v18
                _ = v19
            }
            v20 := []int{32768, 2147483647, 4294967295, 0}
            var v21 float64 = 1e10
            _ = v20
            _ = v21
        } else {
            switch ("\n\t" + ("foo bar" + "a")) {
            case "a":
                v22 := GenS4{F0: 887, F1: 1, F2: 65535}
                _ = v22
            case "hello":
                var v23 int = 32768
                _ = v23
            }
        }
    default:
        switch (("\n\t" + "hello") + ("a" + "\n\t")) {
        case "a":
            if false {
                if false {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        if false {
                            v25 := []int{fi24}
                            for fi26 := 0; fi26 < 3; fi26++ {
                                if (true || (552 >= 42)) {
                                    var v27 int32 = -315
                                    _ = v27
                                } else {
                                    var v28 float32 = 1.5
                                    _ = v28
                                }
                                if len(v25) > 0 {
                                    v25[0] = fi24
                                }
                                var v29 int = fi26
                                var v30 int32 = 5
                                var v31 bool = (!("\n\t" != "\n\t"))
                                if (fi26 < fi24) {
                                    v25 = append(v25, v29)
                                    if v31 {
                                        var v32 string = ("foo bar" + ((("\n\t" + "a") + (("\n\t" + ("a" + "a")) + "")) + ""))
                                        _ = v32
                                    } else {
                                        var v33 uint64 = 475
                                        _ = v33
                                    }
                                }
                                v30 = -1
                                var v34 string = "a"
                                var v35 int64 = 0
                                fmt.Println(v31)
                                _ = fi26
                                _ = v29
                                _ = v30
                                _ = v31
                                _ = v34
                                _ = v35
                            }
                            var v36 uint32 = 2147483648
                            _ = v25
                            _ = v36
                        }
                        if true {
                            if false {
                                fmt.Println(fi24)
                                for fi37 := 0; fi37 < 18; fi37++ {
                                    var v38 int16 = 0
                                    _ = fi37
                                    _ = v38
                                }
                                fmt.Println(fi24)
                            }
                            var v39 string = "foo bar"
                            _ = v39
                        }
                        v40 := GenS5{F0: 1}
                        _ = fi24
                        _ = v40
                    }
                }
            }
        default:
            var v41 int8 = 1
            _ = v41
        }
    }
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 int = 128
    var v3 uint32 = 468
    switch genFn1() {
    case 32768:
        fmt.Println(v2, v3)
    default:
        _ = len(v1)
    }
    var v42 int8 = -1
    var v43 uint16 = genFn0("", "foo bar", 32767)
    v44 := GenS6{F0: v2, F1: true, F2: true}
    for fi45 := 0; fi45 < 14; fi45++ {
        var v46 uint32 = (v3 - (v3 & v3))
        v42 += (v42 << ((4294967295) & 63))
        var v47 int8 = -128
        _ = fi45
        _ = v46
        _ = v47
    }
    var v48 uint8 = 42
    var v49 int = int(v3)
    v50 := GenS7{F0: (v49 & genFn1()), F1: uint(v43)}
    fmt.Println(v42)
    for fi51 := 0; fi51 < 13; fi51++ {
        for fi52 := 0; fi52 < 20; fi52++ {
            if (!false) {
                if true {
                    if ((-291 <= -1) || (v3 > (v3 ^ 0))) {
                        if (float32(3.14159) > 0.0) {
                            if ("a" >= "") {
                                if (true && (!true)) {
                                    v43 = v43
                                } else {
                                    v53 := GenS8{F0: -9223372036854775808, F1: v3, F2: (v42 * (v42 + v42))}
                                    _ = v53
                                }
                            } else {
                                var v54 int32 = 727
                                var v55 int64 = -9223372036854775808
                                _ = v54
                                _ = v55
                            }
                        } else {
                            v56 := GenS9{F0: -32768}
                            v1 = append(v1, v3)
                            switch fi51 {
                            case 518:
                                var v57 int32 = 42
                                _ = v57
                            case 42:
                                v3 = 2
                            case 127:
                                v58 := GenS1{F0: 592}
                                _ = v58
                            default:
                                if len(v1) > 0 {
                                    v1[0] = 0
                                }
                            }
                            _ = v56
                        }
                    }
                    var v59 uint32 = (v3 / ((65535) | 1))
                    _ = v59
                }
            } else {
                v43 += genFn0("hello", "hello", v43)
            }
            var v60 float32 = (3.14159 / 0.1)
            _ = fi52
            _ = v60
        }
        v1 = append(v1, 65535)
        _ = fi51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v48
    _ = v49
    _ = v50
}
