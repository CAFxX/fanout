package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 int16 }
func genFn0() uint {
    var v6 int32 = 2147483647
    switch 32768 {
    case -377:
        switch 0 {
        case -566:
            var v7 int16 = -32768
            _ = v7
        case 2147483647:
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 string = ("foo bar" + "foo bar")
                _ = fi8
                _ = v9
            }
        }
    case 128:
        v10 := []int{2147483647, 2147483647}
        _ = v10
    case -1:
        v11 := []int{-1, 2147483647, -1}
        _ = v11
    }
    v12 := []int{}
    _ = v12
    _ = v6
    return 2
}
func genFn1(p0 bool, p1 uint8, p2 bool) uint64 {
    switch (("a" + "a") + "") {
    case "":
        p1 = 237
    case "\n\t":
        if (!true) {
            v13 := GenS1{F0: "hello", F1: 164, F2: 18446744073709551615}
            v14 := []int{312, 196, -9223372036854775808}
            v15 := GenS1{F0: ("a" + (("foo bar" + "a") + "a")), F1: 619, F2: 0}
            _ = v13
            _ = v14
            _ = v15
        }
    default:
        v16 := []int{32768, -25}
        _ = v16
    }
    fmt.Println(p0)
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(p2, fi17)
        var v18 uint16 = 675
        v19 := []int{fi17, fi17}
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int16 = -919
    v21 := []int{-801}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 255
    v3 := []string{}
    var v4 uint16 = 42
    var v5 string = "hello"
    v5 = (v5 + v5)
    v1[65535] = (genFn0() < 1)
    fmt.Println(v5, v5, v5)
    for fi22 := 0; fi22 < 18; fi22++ {
        switch fi22 {
        case 876:
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 float64 = 3.14159
                switch -9223372036854775808 {
                case -477:
                    v4 = v4
                case -544:
                    _ = len(v1)
                case 128:
                    v5 = (v5 + (v5 + "\n\t"))
                }
                v4 = 42
                _ = fi23
                _ = v24
            }
        case -2147483648:
            v5 = "hello"
        case 591:
            v3 = append(v3, (v5 + v5))
        }
        if true {
            delete(v1, 311)
            v5 = v5
            if (0 == 0) {
                for fi25 := 0; fi25 < 8; fi25++ {
                    v5 = v5
                    var v26 uint64 = 35
                    v27 := GenS1{F0: v5, F1: 32767, F2: v26}
                    v4 -= (v4 * 2)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 2; fi28++ {
                    v4 += v4
                    if (!((!true) || true)) {
                        var v29 uint = 4294967295
                        var v30 uint = v29
                        v2 += v2
                        _ = v29
                        _ = v30
                    } else {
                        if ((1.5 == 1e10) || (false && true)) {
                            var v31 uint = 293
                            v32 := v1[0]
                            if true {
                                v33 := GenS2{F0: 622, F1: (v4 ^ (^(v4))), F2: 42}
                                v31 += 2
                                _ = v33
                            }
                            fmt.Println(fi28)
                            v34 := GenS2{F0: 916, F1: v4, F2: -2147483648}
                            v2 += (v2 % (((^(v2))) | 1))
                            _ = v31
                            _ = v32
                            _ = v34
                        }
                    }
                    if (-2.25 <= ((1.5 / -2.25) / float32(1.5))) {
                        _ = len(v3)
                    } else {
                        v35 := GenS2{F0: 0, F1: v4, F2: -2147483648}
                        var v36 uint = 2
                        _ = v35
                        _ = v36
                    }
                    _ = fi28
                }
                v2 = (v2 >> ((2) & 63))
            } else {
                var v37 uint8 = v2
                _ = v37
            }
        } else {
            var v38 int16 = 0
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v39 := GenS2{F0: 1, F1: (v4 - (v4 + v4)), F2: 596}
                v38 = v38
                _ = v39
            } else {
                v2 = 1
                var v40 int = fi22
                _ = v40
            }
            switch (v5 + ("a" + "\n\t")) {
            case "foo bar":
                var v41 float32 = -2.25
                _ = v41
            }
            _ = v38
        }
        _ = fi22
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        v43 := v1[65535]
        fmt.Println(v2, v2)
        v44 := GenS1{F0: "\n\t", F1: (^(v4)), F2: 42}
        _ = fi42
        _ = v43
        _ = v44
    }
    if true {
        v2 -= 233
        v5 = v5
    } else {
        for fi45 := 0; fi45 < 3; fi45++ {
            switch v5 {
            case "foo bar":
                for fi46 := 0; fi46 < 6; fi46++ {
                    v47 := GenS3{F0: fi46, F1: 42, F2: 128}
                    var v48 uint16 = (v4 << ((1) & 63))
                    if false {
                        v4 = v48
                        v49 := GenS3{F0: (fi45 * fi46), F1: 447, F2: v2}
                        fmt.Println(fi46)
                        _ = v49
                    }
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            case "hello":
                if true {
                    switch "hello" {
                    case "a":
                        v2 += (^(v2))
                    case "hello":
                        for fi50 := 0; fi50 < 20; fi50++ {
                            var v51 bool = false
                            _ = fi50
                            _ = v51
                        }
                    }
                } else {
                    fmt.Println(v5, v2, fi45)
                    switch ((((v5 + v5) + (v5 + v5)) + (("\n\t" + (("a" + "\n\t") + (v5 + "hello"))) + (v5 + ("foo bar" + ("foo bar" + v5))))) + "foo bar") {
                    case "":
                        fmt.Println(v5, v2)
                    case "hello":
                        v5 = ("foo bar" + "hello")
                    case "a":
                        for fi52 := 0; fi52 < 8; fi52++ {
                            var v53 string = "hello"
                            _ = fi52
                            _ = v53
                        }
                    }
                }
            case "\n\t":
                v4 += 65535
            }
            if true {
                v4 += 337
            } else {
                v5 = ((v5 + v5) + v5)
                switch fi45 {
                case 201:
                    switch "a" {
                    case "hello":
                        var v54 int16 = -926
                        _ = v54
                    }
                case 673:
                    fmt.Println(fi45)
                }
                for fi55 := 0; fi55 < 8; fi55++ {
                    v5 = v5
                    var v56 uint = 18446744073709551615
                    switch 0 {
                    case 801:
                        switch fi55 {
                        case 2147483647:
                            var v57 uint = 221
                            _ = v57
                        case 32768:
                            v58 := GenS1{F0: "", F1: v4, F2: uint64(v2)}
                            _ = v58
                        case 32767:
                            v2 += 27
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case 32767:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 688:
                        if false {
                            v59 := GenS4{F0: 42}
                            _ = v59
                        }
                    default:
                        v1[2147483647] = (true && ((v4 << ((65535) & 63)) >= (v4 - (^(v4)))))
                    }
                    _ = fi55
                    _ = v56
                }
            }
            for fi60 := 0; fi60 < 17; fi60++ {
                if (459 < v4) {
                    var v61 int32 = 0
                    v62 := GenS1{F0: v5, F1: 65535, F2: 42}
                    _ = v61
                    _ = v62
                }
                _ = fi60
            }
            _ = fi45
        }
        var v63 int16 = 1000
        var v64 float64 = 1e10
        _ = v63
        _ = v64
    }
    fmt.Println(v5, v2, v2)
    var v65 uint32 = 4294967295
    v65 = v65
    var v66 bool = (float64((0.1 * 0.0)) < 0.1)
    v67 := GenS2{F0: 1, F1: (v4 ^ 65535), F2: 1}
    v5 = v5
    v2 -= (v2 & (v2 / ((v2) | 1)))
    v68 := GenS2{F0: 642, F1: (v4 | v4), F2: -818}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
