package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int16 }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint; F1 int64; F2 uint32 }
func genFn0(p0 uint, p1 string) uint32 {
    switch 42 {
    case -1:
        v4 := []int{127, -100, 999, 128}
        _ = v4
    case 530:
        p0 += 65535
    case 2147483647:
        p0 = p0
    }
    p0 -= 2
    _ = p0
    _ = p1
    return uint32(p0)
}
func genFn1(p0 int16, p1 uint64) uint16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint32 = 2147483647
        for fi7 := 0; fi7 < 10; fi7++ {
            if false {
                p1 = p1
            } else {
                if false {
                    if true {
                        fmt.Println(p1, p0, fi5)
                        var v8 bool = (!false)
                        _ = v8
                    }
                    switch (-(fi5)) {
                    case 0:
                        var v9 uint16 = 0
                        _ = v9
                    default:
                        fmt.Println(v6, fi5)
                    }
                } else {
                    var v10 uint8 = uint8(p1)
                    if true {
                        v11 := []int{(fi7 | fi7), (fi5 * (fi5 % (((fi5 ^ 1)) | 1))), fi7, fi7}
                        _ = v11
                    } else {
                        v10 -= (^(v10))
                        v10 -= (v10 >> ((4) & 63))
                    }
                    if (!true) {
                        fmt.Println(p0, fi5)
                    } else {
                        v12 := GenS1{F0: (!(65535 == 0)), F1: (-(p0)), F2: 32767}
                        v13 := []int{817, -919, fi5, fi5}
                        p1 = p1
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                }
                v14 := GenS1{F0: (!(!false)), F1: 42, F2: p0}
                p1 += p1
                _ = v14
            }
            switch ("foo bar" + ("\n\t" + "foo bar")) {
            case "":
                var v15 uint = 4294967295
                _ = v15
            case "foo bar":
                v16 := []int{}
                fmt.Println(fi7, p1)
                v6 = v6
                _ = v16
            }
            p1 -= p1
            _ = fi7
        }
        var v17 uint = 255
        _ = fi5
        _ = v17
        _ = v6
    }
    v18 := GenS2{F0: (p0 | p0), F1: 0, F2: -9223372036854775808}
    var v19 uint32 = 990
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 int8 = 0
        v21 -= 127
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 65535
}

func main() {
    var v1 int64 = -972
    var v2 uint64 = 1
    var v3 uint8 = 234
    var v22 int = -2147483648
    var v23 float32 = -2.25
    if false {
        var v24 uint64 = v2
        v24 = v2
        _ = v24
    } else {
        v25 := GenS3{F0: 127, F1: 4294967295}
        for fi26 := 0; fi26 < 8; fi26++ {
            v27 := []int{v22}
            var v28 string = "hello"
            _ = fi26
            _ = v27
            _ = v28
        }
        fmt.Println(v22, v23, v3)
        _ = v25
    }
    v29 := GenS3{F0: 14, F1: (v2 << ((65535) & 63))}
    v30 := GenS4{F0: 18446744073709551615, F1: (v3 >> ((615) & 63)), F2: 0}
    v31 := []int{(v22 ^ 2147483647), v22, int(v2)}
    var v32 bool = (228 >= (v3 * (v3 % ((235) | 1))))
    if len(v31) > 0 {
        v31[0] = v22
    }
    v32 = false
    var v33 string = ("a" + "\n\t")
    var v34 uint64 = v2
    if true {
        var v35 uint8 = v3
        var v36 int = v22
        switch v33 {
        case "a":
            v35 = v35
        case "hello":
            switch ("a" + ((v33 + "hello") + v33)) {
            case "\n\t":
                if true {
                    var v37 int32 = 1
                    _ = v37
                } else {
                    v33 = v33
                    fmt.Println(v32, v35, v23)
                }
            }
        case "\n\t":
            if len(v31) > 0 {
                v31[0] = (v36 + (v36 * -1))
            }
        }
        _ = v35
        _ = v36
    }
    if (97 == int8(v2)) {
        var v38 uint8 = 42
        v39 := GenS3{F0: -71, F1: uint64(v38)}
        _ = v38
        _ = v39
    } else {
        var v40 bool = v32
        _ = v40
    }
    v41 := GenS4{F0: 268, F1: v3, F2: -86}
    for fi42 := 0; fi42 < 18; fi42++ {
        switch fi42 {
        case 1:
            if (!v32) {
                for fi43 := 0; fi43 < 8; fi43++ {
                    var v44 uint = 4294967295
                    _ = fi43
                    _ = v44
                }
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
                if v32 {
                    fmt.Println(v1, v3, v23)
                }
            } else {
                switch (v33 + v33) {
                case "foo bar":
                    var v45 float64 = (float64((((0.0 / -2.25) - v23) + float32(-2.25))) + 1.5)
                    _ = v45
                case "a":
                    var v46 uint = 255
                    _ = v46
                case "":
                    fmt.Println(v3, v32)
                }
                var v47 uint64 = v2
                if (v32 || (-2147483648 == 1)) {
                    v48 := GenS4{F0: uint(v2), F1: v3, F2: 100}
                    _ = v48
                } else {
                    var v49 int32 = -2147483648
                    _ = v49
                }
                _ = v47
            }
        case 32767:
            v32 = v32
        default:
            v33 = "hello"
        }
        _ = fi42
    }
    v34 -= 4294967295
    v50 := GenS3{F0: 100, F1: v34}
    switch v33 {
    case "a":
        for fi51 := 0; fi51 < 14; fi51++ {
            v23 = v23
            fmt.Println(v1, fi51, v23)
            v2 = 18446744073709551615
            _ = fi51
        }
    case "\n\t":
        var v52 uint64 = 0
        _ = v52
    case "hello":
        switch (v22 % ((v22) | 1)) {
        case 2147483647:
            v53 := GenS5{F0: 873, F1: -927, F2: uint32(v3)}
            _ = v53
        }
    default:
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v41
    _ = v50
}
