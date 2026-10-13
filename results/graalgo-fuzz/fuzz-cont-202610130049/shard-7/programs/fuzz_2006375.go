package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint; F2 int }
type GenS5 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 string = ("a" + "hello")
    var v3 uint16 = 65535
    var v4 bool = (!(!(391 != 2)))
    v5 := []int{4294967295}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 bool = v4
    fmt.Println(v4)
    switch v2 {
    case "":
        v6 = (223 <= 587)
    case "hello":
        _ = len(v5)
    default:
        fmt.Println(v3, v4, v1)
    }
    if false {
        if len(v5) > 0 {
            v5[0] = -254
        }
        if len(v5) > 0 {
            v5[0] = -580
        }
    }
    var v7 int64 = 1
    v7 = 0
    if v4 {
        if (0 != 255) {
            _ = len(v5)
            var v8 string = ""
            _ = v8
        } else {
            v1 -= (v1 ^ (v1 + v1))
        }
    }
    fmt.Println(v4, v6)
    for fi9 := 0; fi9 < 18; fi9++ {
        v3 = (^(v3))
        for fi10 := 0; fi10 < 20; fi10++ {
            v5 = append(v5, (fi10 ^ (fi10 | (-(fi10)))))
            _ = fi10
        }
        _ = fi9
    }
    fmt.Println(v4, v6)
    v11 := GenS1{F0: -1}
    var v12 uint32 = 474
    for fi13 := 0; fi13 < 3; fi13++ {
        switch "hello" {
        case "a":
            v12 -= (v12 | v12)
        case "":
            if v6 {
                switch (v2 + "hello") {
                case "hello":
                    var v14 string = "foo bar"
                    _ = v14
                case "a":
                    if ((((1.5 * -2.25) * 1.5) * ((1e10 / ((0.1 - -2.25) + -2.25)) * 1.5)) != -2.25) {
                        var v15 uint32 = 279
                        v2 = (v2 + v2)
                        var v16 int64 = v7
                        _ = v15
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            v7 = v7
                            v18 := GenS1{F0: 74}
                            v1 = (^(v1))
                            if true {
                                v1 += v1
                                if v4 {
                                    v5 = append(v5, (fi13 & 481))
                                    if v6 {
                                        var v19 uint16 = v3
                                        _ = len(v5)
                                        v20 := GenS1{F0: 0}
                                        _ = v19
                                        _ = v20
                                    }
                                    if (1 != -59) {
                                        v6 = true
                                        if v4 {
                                            v5 = append(v5, 32767)
                                        }
                                        fmt.Println(v1, fi17)
                                        if len(v5) > 0 {
                                            v5[0] = (fi17 >> ((4294967295) & 63))
                                        }
                                    }
                                } else {
                                    fmt.Println(v1, v7)
                                    var v21 int8 = -2
                                    _ = v21
                                }
                                var v22 int64 = v7
                                if (!v6) {
                                    if v6 {
                                        if v6 {
                                            v1 = (v1 ^ v1)
                                            fmt.Println(fi13)
                                        }
                                        v5 = append(v5, fi17)
                                    } else {
                                        v5 = append(v5, -764)
                                        if ((fi17 ^ fi17) == fi13) {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    }
                                    switch v2 {
                                    case "":
                                        var v23 uint = 0
                                        v24 := GenS2{F0: -1}
                                        _ = v23
                                        _ = v24
                                    default:
                                        var v25 int8 = int8(fi17)
                                        _ = v25
                                    }
                                    v26 := GenS2{F0: 2147483647}
                                    fmt.Println(fi13, v3, fi17)
                                    v3 += (v3 << ((203) & 63))
                                    _ = v26
                                }
                                _ = v22
                            } else {
                                _ = len(v5)
                            }
                            _ = fi17
                            _ = v18
                        }
                        _ = len(v5)
                    }
                case "\n\t":
                    var v27 int = fi13
                    _ = v27
                }
            }
        case "hello":
            v7 -= -946
        default:
            switch (v2 + ((v2 + "foo bar") + "foo bar")) {
            case "a":
                if v6 {
                    if (false && v4) {
                        if (!(0 < 94)) {
                            v6 = v4
                        } else {
                            if (v2 == v2) {
                                switch int(v7) {
                                case -88:
                                    v28 := GenS1{F0: -119}
                                    _ = v28
                                default:
                                    v4 = false
                                }
                                v29 := GenS2{F0: fi13}
                                if v4 {
                                    var v30 uint32 = 773
                                    if (fi13 >= fi13) {
                                        v6 = false
                                    } else {
                                        v30 = (v12 % ((v30) | 1))
                                    }
                                    _ = v30
                                } else {
                                    if ((0.0 * ((1.5 / 0.1) + (3.14159 - -2.25))) < 1e10) {
                                        v2 = "foo bar"
                                        fmt.Println(v2, v7)
                                    }
                                    v31 := GenS1{F0: 1}
                                    _ = len(v5)
                                    _ = v31
                                }
                                _ = v29
                            } else {
                                switch ("\n\t" + v2) {
                                case "":
                                    var v32 int8 = 1
                                    _ = v32
                                default:
                                    v3 = v3
                                }
                                for fi33 := 0; fi33 < 15; fi33++ {
                                    v3 = (v3 ^ v3)
                                    fmt.Println(v12, fi13, v4)
                                    var v34 float64 = 0.1
                                    if v6 {
                                        fmt.Println(v12, v2, v12)
                                        v35 := GenS3{F0: -1, F1: 0, F2: (v12 ^ v12)}
                                        var v36 uint16 = v3
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v37 := GenS4{F0: 4294967295, F1: 42, F2: (fi33 << ((42) & 63))}
                                        if (1 == 42) {
                                            v3 = v3
                                        }
                                        _ = len(v5)
                                        _ = v37
                                    }
                                    switch (v2 + v2) {
                                    case "\n\t":
                                        v38 := GenS5{F0: 2, F1: (v3 - v3), F2: 0}
                                        _ = v38
                                    case "foo bar":
                                        var v39 uint8 = 201
                                        _ = v39
                                    case "":
                                        v2 = "a"
                                    }
                                    if true {
                                        var v40 uint16 = v3
                                        if ((128 > 27) && v4) {
                                            fmt.Println(v6, fi33, v2)
                                            switch v2 {
                                            case "foo bar":
                                                v1 = v1
                                            default:
                                                v41 := GenS4{F0: 18446744073709551615, F1: 65535, F2: fi33}
                                                _ = v41
                                            }
                                        }
                                        _ = v40
                                    }
                                    _ = fi33
                                    _ = v34
                                }
                            }
                        }
                        v2 = (((v2 + v2) + v2) + v2)
                    } else {
                        var v42 int64 = v7
                        _ = v42
                    }
                    switch v2 {
                    case "a":
                        v43 := GenS6{F0: (v1 & v1)}
                        _ = v43
                    case "hello":
                        var v44 uint64 = v1
                        _ = v44
                    }
                    if (!v6) {
                        v45 := GenS5{F0: 2, F1: 342, F2: v12}
                        _ = v45
                    } else {
                        var v46 int64 = v7
                        if (true && v6) {
                            v46 += (-(v46))
                        }
                        _ = v46
                    }
                }
            case "hello":
                var v47 int = (fi13 >> ((42) & 63))
                _ = v47
            }
        }
        var v48 int64 = 171
        _ = fi13
        _ = v48
    }
    v49 := GenS7{F0: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
