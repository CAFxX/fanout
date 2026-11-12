package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int8 }
func genFn0(p0 int8, p1 int32, p2 uint32) int64 {
    fmt.Println(p1, p2, p0)
    if (9223372036854775807 < -1) {
        p0 += p0
    } else {
        for fi5 := 0; fi5 < 3; fi5++ {
            if (-937 < int64(p0)) {
                p2 += p2
                for fi6 := 0; fi6 < 17; fi6++ {
                    p1 = -1
                    p1 = p1
                    var v7 uint64 = 18446744073709551615
                    _ = fi6
                    _ = v7
                }
                v8 := GenS1{F0: 125}
                _ = v8
            }
            _ = fi5
        }
        switch 5 {
        case 0:
            p2 = p2
        case 287:
            p0 -= 42
        default:
            v9 := GenS2{F0: (p2 << ((4294967295) & 63)), F1: uint8(p2), F2: 884}
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 725
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ((("foo bar" + ("a" + "\n\t")) + "") + "")
    var v3 uint8 = 1
    v4 := make(map[int]string)
    v3 -= (^(v3))
    var v10 int = 641
    for fi11 := 0; fi11 < 5; fi11++ {
        v10 += fi11
        _ = fi11
    }
    var v12 int16 = 42
    fmt.Println(v12)
    switch v2 {
    case "foo bar":
        v2 = "a"
    case "\n\t":
        var v13 int = (v10 ^ int(v12))
        _ = v13
    case "hello":
        var v14 uint32 = 114
        _ = v14
    default:
        v15 := v1[(v10 ^ v10)]
        _ = v15
    }
    switch v2 {
    case "a":
        v16 := GenS1{F0: 42}
        _ = v16
    case "hello":
        var v17 string = "foo bar"
        _ = v17
    }
    if false {
        var v18 string = v2
        _ = v18
    } else {
        if true {
            switch (v10 | v10) {
            case -584:
                for fi19 := 0; fi19 < 10; fi19++ {
                    if false {
                        var v20 bool = ((true && (42 <= 1)) && false)
                        v21 := v4[(^(fi19))]
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint64 = 1
                        var v23 int = v10
                        _ = v22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 6; fi24++ {
                        v25 := GenS2{F0: 2, F1: v3, F2: 619}
                        if (!(779 >= 793)) {
                            var v26 int16 = v12
                            v12 += v12
                            _ = len(v4)
                            _ = v26
                        } else {
                            var v27 int32 = 0
                            v1[127] = 2
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                    }
                    _ = fi19
                }
            case 358:
                switch ("a" + ((((v2 + (v2 + v2)) + v2) + v2) + "a")) {
                case "foo bar":
                    if true {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            for fi29 := 0; fi29 < 14; fi29++ {
                                var v30 int32 = int32(fi28)
                                fmt.Println(fi29)
                                if true {
                                    var v31 int64 = 0
                                    _ = v31
                                } else {
                                    v32 := GenS2{F0: 2147483647, F1: (v3 / (((v3 >> ((273) & 63))) | 1)), F2: 1}
                                    _ = v32
                                }
                                fmt.Println(v10)
                                if false {
                                    v33 := GenS3{F0: (fi29 % ((fi28) | 1))}
                                    v34 := []int{127, fi29}
                                    switch v2 {
                                    case "\n\t":
                                        switch 127 {
                                        case -2147483648:
                                            v10 += (^(v10))
                                        }
                                    default:
                                        fmt.Println(fi28, v12, fi28)
                                    }
                                    v35 := GenS4{F0: (v2 + v2)}
                                    v1[fi28] = 2147483648
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                _ = fi29
                                _ = v30
                            }
                            var v36 bool = false
                            var v37 uint = 65535
                            _ = fi28
                            _ = v36
                            _ = v37
                        }
                    } else {
                        var v38 uint = 1
                        _ = v38
                    }
                case "a":
                    if (!true) {
                        v1[(v10 / ((v10) | 1))] = 1
                        delete(v1, (^(v10)))
                    } else {
                        v39 := v1[(-(v10))]
                        for fi40 := 0; fi40 < 16; fi40++ {
                            v41 := GenS4{F0: (v2 + "")}
                            if (-2147483648 >= 42) {
                                v42 := GenS5{F0: 2, F1: -1}
                                _ = v42
                            }
                            _ = fi40
                            _ = v41
                        }
                        _ = v39
                    }
                default:
                    var v43 int64 = -1
                    _ = v43
                }
            default:
                var v44 bool = (true || true)
                _ = v44
            }
            v2 = ""
            v10 += v10
        } else {
            v2 = v2
            v3 -= v3
            switch (v2 + "foo bar") {
            case "foo bar":
                v45 := GenS1{F0: 2}
                _ = v45
            }
        }
    }
    var v46 uint = 65535
    v3 -= (v3 % (((v3 ^ v3)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v46
}
