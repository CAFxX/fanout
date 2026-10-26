package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 bool }
func genFn0(p0 uint64, p1 int32, p2 int8) int8 {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int8 = p2
        if true {
            if ((65535 == 2147483648) || ((-(p1)) > p1)) {
                var v6 uint32 = 561
                _ = v6
            } else {
                v7 := GenS1{F0: v5}
                _ = v7
            }
            fmt.Println(fi4, p0)
        }
        p0 -= p0
        _ = fi4
        _ = v5
    }
    var v8 uint64 = p0
    if false {
        var v9 int16 = 1000
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 127
}
func genFn1(p0 int) uint {
    fmt.Println(p0)
    if false {
        v10 := []int{}
        p0 += (p0 << ((18446744073709551615) & 63))
        _ = v10
    } else {
        p0 -= p0
        var v11 uint = 65535
        switch 1 {
        case 4294967295:
            v11 = 0
        case 32768:
            var v12 int16 = 660
            _ = v12
        default:
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS2{F0: 127}
                _ = fi13
                _ = v14
            }
        }
        _ = v11
    }
    v15 := GenS3{F0: 397, F1: 222, F2: 32767}
    var v16 int = -1
    _ = p0
    _ = v15
    _ = v16
    return 2
}

func main() {
    var v1 string = "hello"
    var v2 bool = ((((1e10 / 1.5) / (float64(1.5) - ((((1.5 - 1e10) / (0.0 + 1.5)) / 1.5) - 3.14159))) + 0.0) < 0.1)
    var v3 string = v1
    switch "hello" {
    case "hello":
        var v17 float32 = ((1e10 / 3.14159) * 3.14159)
        _ = v17
    case "":
        var v18 float32 = ((float32(3.14159) + float32((((3.14159 - 1.5) - (3.14159 + 1e10)) / 1.5))) / 0.1)
        _ = v18
    case "\n\t":
        v1 = ""
    default:
        if v2 {
            var v19 bool = (3.14159 != 0.0)
            if false {
                fmt.Println(v19, v19)
                v1 = "hello"
                var v20 int32 = 97
                _ = v20
            }
            v21 := []int{9223372036854775807, 9223372036854775807, 42}
            _ = v19
            _ = v21
        } else {
            v2 = v2
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 float64 = (float64(-2.25) / 1.0)
                v24 := []int{}
                _ = fi22
                _ = v23
                _ = v24
            }
            if v2 {
                v2 = (1 != 0)
            }
        }
    }
    var v25 int = 4294967295
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS1{F0: 1}
        v1 = (v1 + (v1 + "a"))
        switch (v25 % (((fi26 + v25)) | 1)) {
        case 32768:
            v2 = (!(267 > 255))
        case -901:
            v28 := []int{fi26, v25, (v25 >> ((4294967295) & 63)), v25}
            _ = v28
        case -2147483648:
            v25 = (v25 ^ (fi26 - -778))
        }
        _ = fi26
        _ = v27
    }
    var v29 uint = 4294967295
    fmt.Println(v3)
    var v30 int8 = genFn0(0, -802, -128)
    if v2 {
        if (((42 != uint16(v29)) && true) && (0 > 0)) {
            if false {
                v3 = v1
                for fi31 := 0; fi31 < 5; fi31++ {
                    v32 := GenS1{F0: v30}
                    v33 := GenS2{F0: 112}
                    var v34 int = fi31
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                switch (v3 + v3) {
                case "a":
                    if true {
                        var v35 int32 = 1
                        _ = v35
                    } else {
                        var v36 int16 = 42
                        _ = v36
                    }
                case "foo bar":
                    var v37 float64 = -2.25
                    _ = v37
                case "hello":
                    var v38 int8 = genFn0(755, 2147483647, v30)
                    _ = v38
                default:
                    var v39 int8 = genFn0(2, -1, -1)
                    _ = v39
                }
            } else {
                v40 := []int{}
                _ = v40
            }
            for fi41 := 0; fi41 < 4; fi41++ {
                for fi42 := 0; fi42 < 10; fi42++ {
                    var v43 int = fi41
                    _ = fi42
                    _ = v43
                }
                for fi44 := 0; fi44 < 7; fi44++ {
                    if v2 {
                        switch ("foo bar" + (v3 + (v1 + v3))) {
                        case "a":
                            v1 = ""
                        case "foo bar":
                            if (!v2) {
                                v30 = v30
                                if v2 {
                                    var v45 int = (-(fi44))
                                    v1 = (v3 + v1)
                                    if true {
                                        fmt.Println(fi41, fi41)
                                    }
                                    switch (v1 + (v3 + ((v1 + (v1 + "\n\t")) + ((v3 + "a") + v1)))) {
                                    case "a":
                                        switch v1 {
                                        case "hello":
                                            if (32767 != 9223372036854775807) {
                                                v46 := GenS3{F0: 204, F1: 944, F2: 2}
                                                _ = v46
                                            }
                                        }
                                    default:
                                        v47 := []int{(-(v45)), (-(v25)), v25}
                                        _ = v47
                                    }
                                    v1 = (((("hello" + (v1 + ((v3 + (((v3 + ("foo bar" + v1)) + ("hello" + "hello")) + "a")) + v3))) + "hello") + "") + "foo bar")
                                    if v2 {
                                        v48 := []int{fi44, 32767}
                                        switch "a" {
                                        case "foo bar":
                                            v49 := GenS2{F0: (^(v30))}
                                            _ = v49
                                        case "\n\t":
                                            var v50 int64 = 779
                                            _ = v50
                                        }
                                        _ = v48
                                    } else {
                                        if (!v2) {
                                            v1 = ((v1 + v3) + ("hello" + v3))
                                            var v51 int64 = -1
                                            _ = v51
                                        }
                                        v3 = v1
                                    }
                                    _ = v45
                                }
                                if (!((v29 | v29) > 4294967295)) {
                                    if false {
                                        v52 := GenS3{F0: 65535, F1: 18446744073709551615, F2: 685}
                                        switch fi41 {
                                        case -1:
                                            v25 = (v25 % ((v25) | 1))
                                        case 42:
                                            v53 := GenS4{F0: (v29 ^ 2)}
                                            _ = v53
                                        case 32768:
                                            v29 -= v29
                                        default:
                                            if ((!v2) || (((0.0 + 1.5) * -2.25) >= 0.0)) {
                                                v30 = -128
                                                v54 := GenS3{F0: 463, F1: 18446744073709551615, F2: 2}
                                                switch "foo bar" {
                                                case "":
                                                    var v55 uint = (v29 & genFn1(fi41))
                                                    _ = v55
                                                case "hello":
                                                    var v56 bool = (0.1 == 0.1)
                                                    _ = v56
                                                case "\n\t":
                                                    fmt.Println(v2, v30)
                                                default:
                                                    switch ("a" + "hello") {
                                                    case "":
                                                        v57 := []int{}
                                                        _ = v57
                                                    case "hello":
                                                        switch ("foo bar" + "") {
                                                        case "hello":
                                                            var v58 float64 = (0.1 / (1e10 + 1e10))
                                                            _ = v58
                                                        }
                                                    }
                                                    var v59 uint16 = 847
                                                    v60 := GenS2{F0: (v30 - (v30 + (v30 * 0)))}
                                                    var v61 int16 = 64
                                                    var v62 uint32 = uint32(v59)
                                                    var v63 int64 = 42
                                                    var v64 float64 = 0.1
                                                    var v65 bool = (v62 < v62)
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                    _ = v65
                                                }
                                                _ = v54
                                            }
                                        }
                                        fmt.Println(fi44, v3, v2)
                                        _ = v52
                                    } else {
                                        var v66 uint = v29
                                        v67 := []int{(fi41 / ((v25) | 1)), fi44, (v25 * 4294967295), (fi44 | v25)}
                                        _ = v66
                                        _ = v67
                                    }
                                    v68 := GenS2{F0: -1}
                                    _ = v68
                                }
                            }
                        case "":
                            v69 := GenS5{F0: v29, F1: false}
                            _ = v69
                        }
                    } else {
                        var v70 string = "hello"
                        v3 = (v3 + v3)
                        v29 = v29
                        _ = v70
                    }
                    v71 := GenS4{F0: v29}
                    v72 := []int{v25, -2147483648, 42}
                    _ = fi44
                    _ = v71
                    _ = v72
                }
                _ = fi41
            }
            v73 := GenS4{F0: (v29 << ((846) & 63))}
            _ = v73
        } else {
            v3 = v3
            fmt.Println(v2)
            for fi74 := 0; fi74 < 7; fi74++ {
                switch (fi74 & v25) {
                case 4294967295:
                    var v75 bool = v2
                    _ = v75
                case 941:
                    switch (v1 + v1) {
                    case "a":
                        v76 := GenS4{F0: (v29 >> ((v29) & 63))}
                        _ = v76
                    default:
                        v3 = "a"
                    }
                default:
                    v77 := GenS2{F0: (v30 | (v30 % ((v30) | 1)))}
                    _ = v77
                }
                _ = fi74
            }
        }
        var v78 bool = false
        _ = v78
    } else {
        v3 = "foo bar"
        if v2 {
            v79 := []int{(v25 >> ((309) & 63)), v25}
            v1 = ("a" + v3)
            _ = v79
        } else {
            switch 767 {
            case -88:
                v2 = v2
            }
            v80 := []int{v25}
            fmt.Println(v2, v29)
            _ = v80
        }
        for fi81 := 0; fi81 < 20; fi81++ {
            var v82 uint = (v29 & v29)
            v83 := []int{}
            _ = fi81
            _ = v82
            _ = v83
        }
    }
    var v84 float64 = (3.14159 + float64(((float32(0.1) + 1.5) + (float32((0.1 * -2.25)) + (1.5 - ((0.1 - ((3.14159 / float32((3.14159 / 3.14159))) - 1e10)) * (float32(1.5) - (1e10 / 1.5))))))))
    if v2 {
        v3 = "hello"
    } else {
        v2 = (42 > (v25 << ((v29) & 63)))
        for fi85 := 0; fi85 < 12; fi85++ {
            v1 = "a"
            if ((0.1 * 3.14159) == 3.14159) {
                v29 += v29
                var v86 uint64 = 42
                fmt.Println(v2)
                _ = v86
            }
            _ = fi85
        }
        v87 := []int{v25, v25, v25, 353}
        _ = v87
    }
    v29 += (v29 - (v29 - (v29 ^ genFn1((v25 ^ (v25 | v25))))))
    var v88 string = v1
    var v89 int = v25
    switch ("a" + "hello") {
    case "hello":
        var v90 uint64 = 18446744073709551615
        _ = v90
    case "a":
        v89 = 665
    case "foo bar":
        fmt.Println(v1, v88, v1)
    default:
        fmt.Println(v25, v88)
    }
    v30 = (^(v30))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v84:", v84)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v84
    _ = v88
    _ = v89
}
