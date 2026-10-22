package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int64; F1 uint32; F2 int64 }
type GenS3 struct { F0 int8; F1 int32 }
func genFn0(p0 int, p1 uint16) string {
    var v6 int16 = -32768
    v7 := GenS1{F0: (^(v6)), F1: 42}
    if false {
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS2{F0: 42, F1: 4294967295, F2: -1}
                _ = fi9
                _ = v10
            }
            fmt.Println(p0)
            _ = fi8
        }
        p1 = p1
    }
    var v11 string = "hello"
    fmt.Println(v11, p0)
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    _ = v7
    return v11
}
func genFn1(p0 uint64, p1 int16) uint {
    p0 = (p0 << ((4294967295) & 63))
    switch ("foo bar" + "") {
    case "\n\t":
        if (true || ("hello" < "")) {
            switch ((("a" + (((("a" + "a") + "foo bar") + "") + "\n\t")) + "\n\t") + "\n\t") {
            case "a":
                v12 := []int{9223372036854775807}
                _ = v12
            default:
                switch ("\n\t" + "") {
                case "":
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(p0)
                    case "hello":
                        if false {
                            var v13 int8 = 42
                            _ = v13
                        } else {
                            var v14 int16 = p1
                            for fi15 := 0; fi15 < 12; fi15++ {
                                fmt.Println(fi15)
                                v16 := GenS3{F0: -62, F1: -187}
                                _ = fi15
                                _ = v16
                            }
                            switch 2147483647 {
                            case 684:
                                var v17 float64 = (float64(-2.25) * 0.0)
                                _ = v17
                            case -9223372036854775808:
                                var v18 int16 = -833
                                _ = v18
                            case 128:
                                if ((42 <= 515) && true) {
                                    var v19 int8 = 1
                                    if (210 != 692) {
                                        fmt.Println(v19)
                                    } else {
                                        v20 := []int{2147483647, 1, 128, -706}
                                        _ = v20
                                    }
                                    _ = v19
                                } else {
                                    v21 := []int{-1}
                                    _ = v21
                                }
                            default:
                                fmt.Println(p1, v14, p1)
                            }
                            _ = v14
                        }
                    case "foo bar":
                        p0 -= 395
                    default:
                        p0 -= p0
                    }
                }
            }
        }
    case "hello":
        switch ("hello" + ("\n\t" + "a")) {
        case "hello":
            for fi22 := 0; fi22 < 8; fi22++ {
                v23 := []int{(fi22 - 32767), -9223372036854775808, (fi22 | fi22), (fi22 % ((fi22) | 1))}
                var v24 int = -2147483648
                if true {
                    fmt.Println(p0)
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        case "":
            var v25 uint = 1
            _ = v25
        }
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := []uint32{2, 0, 1, 685}
    var v3 int = -1
    var v4 int = v3
    var v5 uint16 = 771
    v26 := GenS3{F0: 42, F1: 2147483647}
    fmt.Println(v3, v4)
    var v27 int = 42
    v4 += 127
    v1 = append(v1, -423)
    var v28 float64 = 1.5
    if (false || true) {
        var v29 uint32 = 285
        _ = v29
    }
    v3 -= v4
    var v30 uint32 = 1
    if true {
        v1 = append(v1, v4)
        fmt.Println(v4, v5, v30)
        v30 = (^(v30))
    } else {
        var v31 bool = false
        _ = v31
    }
    var v32 int = v4
    for fi33 := 0; fi33 < 6; fi33++ {
        var v34 uint = 255
        var v35 uint64 = 2
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 uint8 = 74
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
