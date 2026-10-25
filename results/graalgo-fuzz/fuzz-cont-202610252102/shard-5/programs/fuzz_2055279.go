package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string; F1 uint32 }
func genFn0() uint32 {
    var v4 int32 = 2147483647
    v5 := GenS1{F0: 3}
    switch ((("hello" + "a") + (("" + (("a" + "\n\t") + ("hello" + "a"))) + "\n\t")) + (("hello" + "hello") + "a")) {
    case "\n\t":
        fmt.Println(v4)
    case "a":
        v4 += 2147483647
    }
    if false {
        v6 := GenS1{F0: 255}
        _ = v6
    }
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint32, p1 int8, p2 bool) uint {
    p0 = (^(p0))
    for fi7 := 0; fi7 < 10; fi7++ {
        if false {
            v8 := GenS2{F0: (p0 * 65535), F1: 128}
            if (!false) {
                v9 := []int{}
                switch "hello" {
                case "foo bar":
                    if len(v9) > 0 {
                        v9[0] = fi7
                    }
                }
                var v10 uint = 255
                _ = v10
                _ = v9
            }
            var v11 float64 = (1e10 - 3.14159)
            _ = v11
            _ = v8
        } else {
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS2{F0: (p0 * genFn0()), F1: 80}
                _ = fi12
                _ = v13
            }
            v14 := []int{(fi7 * (^(fi7))), fi7, (fi7 / ((fi7) | 1)), (-(fi7))}
            _ = v14
        }
        v15 := []int{fi7}
        switch 9223372036854775807 {
        case 4294967295:
            for fi16 := 0; fi16 < 8; fi16++ {
                p2 = true
                var v17 uint16 = 246
                v18 := GenS3{F0: 42}
                v19 := GenS4{F0: (v17 | (v17 - (v17 / (((v17 % ((v17) | 1))) | 1)))), F1: 214, F2: 255}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        case 32767:
            var v20 int32 = 42
            _ = v20
        }
        _ = fi7
        _ = v15
    }
    switch "a" {
    case "\n\t":
        p1 -= (p1 + (p1 & (p1 | 127)))
    case "a":
        if (!(!p2)) {
            v21 := GenS5{F0: uint64(p0), F1: ((("a" + ("\n\t" + "hello")) + "a") + "foo bar")}
            _ = v21
        } else {
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 string = (("foo bar" + "") + (((((("hello" + "") + "\n\t") + (("hello" + "hello") + "a")) + "") + ("foo bar" + (("\n\t" + "a") + "foo bar"))) + ("hello" + ("a" + ""))))
                switch v23 {
                case "hello":
                    v24 := []int{fi22}
                    _ = v24
                default:
                    v23 = v23
                }
                if p2 {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        var v26 bool = false
                        if true {
                            p0 -= genFn0()
                        } else {
                            var v27 uint64 = 11
                            fmt.Println(fi25, v23, p0)
                            p1 -= p1
                            v28 := GenS6{F0: 666}
                            _ = v27
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                    }
                }
                _ = fi22
                _ = v23
            }
        }
    case "foo bar":
        if true {
            p1 = 0
            if ((p1 << ((791) & 63)) <= p1) {
                switch 32767 {
                case 128:
                    if (32767 != 68) {
                        if p2 {
                            var v29 int64 = 832
                            var v30 int8 = -77
                            _ = v29
                            _ = v30
                        }
                        if true {
                            p2 = (-32768 != 42)
                            var v31 uint16 = 483
                            _ = v31
                        } else {
                            var v32 float32 = 1.5
                            var v33 int32 = -1
                            _ = v32
                            _ = v33
                        }
                        fmt.Println(p0, p0)
                    }
                case -102:
                    var v34 int64 = 9223372036854775807
                    _ = v34
                case 91:
                    if (-2147483648 >= 42) {
                        v35 := []int{-1, 42}
                        _ = v35
                    } else {
                        v36 := []int{32767, -1}
                        fmt.Println(p0, p1)
                        switch -1 {
                        case 42:
                            for fi37 := 0; fi37 < 8; fi37++ {
                                for fi38 := 0; fi38 < 14; fi38++ {
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(fi38)
                                        switch (fi38 - fi37) {
                                        case 837:
                                            switch fi38 {
                                            case 32768:
                                                switch ((("a" + ("foo bar" + "")) + "\n\t") + "hello") {
                                                case "":
                                                    p0 = (p0 & (^(p0)))
                                                case "a":
                                                    v39 := GenS6{F0: 18446744073709551615}
                                                    _ = v39
                                                }
                                            }
                                        case -9223372036854775808:
                                            var v40 uint64 = 2
                                            _ = v40
                                        }
                                    default:
                                        var v41 int64 = 9223372036854775807
                                        _ = v41
                                    }
                                    _ = fi38
                                }
                                var v42 float32 = 1e10
                                _ = fi37
                                _ = v42
                            }
                        case 32767:
                            p1 = (p1 * p1)
                        }
                        _ = v36
                    }
                }
                for fi43 := 0; fi43 < 17; fi43++ {
                    var v44 float64 = (-2.25 * 1e10)
                    _ = fi43
                    _ = v44
                }
            }
        }
    default:
        for fi45 := 0; fi45 < 19; fi45++ {
            p0 = p0
            var v46 uint32 = (^(p0))
            v47 := GenS7{F0: "foo bar", F1: (v46 ^ 705)}
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 356
}

func main() {
    var v1 uint16 = 65535
    var v2 int = -1
    var v3 int64 = -920
    fmt.Println(v1, v1)
    v48 := []int{}
    if false {
        if len(v48) > 0 {
            v48[0] = (v2 * v2)
        }
    }
    var v49 bool = false
    if len(v48) > 0 {
        v48[0] = -1
    }
    v3 = v3
    v50 := GenS5{F0: 0, F1: ((("" + "") + "hello") + (((("\n\t" + "hello") + "foo bar") + "foo bar") + ("" + ("foo bar" + ((("foo bar" + "") + "foo bar") + "hello")))))}
    if false {
        for fi51 := 0; fi51 < 8; fi51++ {
            switch 42 {
            case 32768:
                if len(v48) >= 2 {
                    v48 = v48[1:len(v48)]
                }
            case -583:
                var v52 uint = genFn1(65535, 127, ((-111 <= 42) && (v49 && v49)))
                _ = v52
            default:
                fmt.Println(v3, v49)
            }
            v3 = 42
            if (!false) {
                var v53 int16 = 647
                _ = v53
            }
            _ = fi51
        }
    }
    switch v2 {
    case 987:
        var v54 uint = 18446744073709551615
        _ = v54
    }
    switch v2 {
    case 219:
        for fi55 := 0; fi55 < 5; fi55++ {
            for fi56 := 0; fi56 < 12; fi56++ {
                v49 = ((v3 + (-(v3))) < v3)
                _ = fi56
            }
            fmt.Println(v1, v3)
            _ = fi55
        }
    case 32768:
        v57 := GenS3{F0: (v3 % ((v3) | 1))}
        _ = v57
    case 657:
        v49 = (!((v1 * v1) > 243))
    default:
        var v58 int16 = 32767
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v50
}
