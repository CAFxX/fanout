package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint32; F2 string }
type GenS5 struct { F0 uint32; F1 int8; F2 uint }
func genFn0(p0 uint32, p1 uint32, p2 int) uint32 {
    v5 := GenS1{F0: 2}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 935
}
func genFn1(p0 int) uint16 {
    var v6 uint32 = 741
    if ((false && true) && false) {
        if true {
            var v7 string = "hello"
            var v8 uint = 4294967295
            var v9 float64 = 3.14159
            _ = v7
            _ = v8
            _ = v9
        } else {
            if (165 > 128) {
                p0 -= (p0 - (p0 << ((2) & 63)))
            }
            fmt.Println(p0)
            v10 := GenS2{F0: -144}
            _ = v10
        }
        var v11 bool = true
        _ = v11
    }
    v12 := GenS2{F0: -627}
    _ = p0
    _ = v12
    _ = v6
    return 255
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -833
    var v3 int64 = 499
    var v4 int = 32768
    if true {
        v13 := []int{(v4 * 750), (v4 | v4), (v4 & -402), v4}
        _ = v13
    } else {
        v14 := []int{79, (-(v4)), -1, -2147483648}
        _ = len(v1)
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := GenS3{F0: -9223372036854775808}
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    var v17 uint32 = genFn0(826, 1, (v4 + v4))
    var v18 string = (("a" + "foo bar") + (("" + ("a" + (((("\n\t" + (("a" + "") + "a")) + "") + "") + ("foo bar" + ("foo bar" + "foo bar"))))) + (((("foo bar" + ("foo bar" + "\n\t")) + ("a" + "a")) + "hello") + ("" + "hello"))))
    switch v18 {
    case "foo bar":
        var v19 int16 = v2
        _ = v19
    case "a":
        if (!(v4 < v4)) {
            switch v18 {
            case "hello":
                switch (v4 >> ((0) & 63)) {
                case -592:
                    v20 := GenS4{F0: 1, F1: 65535, F2: "foo bar"}
                    _ = v20
                }
            default:
                v1[((v18 + "\n\t") + v18)] = v4
            }
            v21 := GenS1{F0: 18446744073709551615}
            _ = v21
        }
    default:
        v22 := GenS5{F0: genFn0((v17 * (v17 >> ((42) & 63))), v17, (v4 >> ((65535) & 63))), F1: 127, F2: 845}
        _ = v22
    }
    fmt.Println(v2, v4, v17)
    var v23 uint64 = 23
    if (!true) {
        switch (v18 + (v18 + v18)) {
        case "foo bar":
            v1[v18] = -611
        case "a":
            if true {
                if true {
                    fmt.Println(v4, v2, v4)
                } else {
                    v2 = (v2 * v2)
                    fmt.Println(v18)
                }
            }
        case "\n\t":
            if false {
                if (true || (!(!((!true) && true)))) {
                    switch 32767 {
                    case 998:
                        var v24 int = v4
                        _ = v24
                    case -601:
                        if true {
                            fmt.Println(v3)
                        }
                    case -9223372036854775808:
                        if false {
                            var v25 float32 = (1e10 / float32((0.1 / 0.1)))
                            v4 = v4
                            var v26 string = v18
                            _ = v25
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 5; fi27++ {
                                var v28 int = v4
                                var v29 bool = (!false)
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        }
                    default:
                        for fi30 := 0; fi30 < 18; fi30++ {
                            v3 += 9223372036854775807
                            _ = fi30
                        }
                    }
                }
            }
        default:
            v3 = (v3 & (v3 * 9223372036854775807))
        }
        if (false && (!((v18 > "hello") && true))) {
            var v31 bool = (0.1 >= 0.0)
            fmt.Println(v2, v3)
            var v32 uint8 = 255
            _ = v31
            _ = v32
        } else {
            var v33 int64 = v3
            var v34 float64 = 0.1
            _ = v33
            _ = v34
        }
    }
    var v35 uint8 = 255
    v36 := []int{}
    if true {
        var v37 uint8 = v35
        for fi38 := 0; fi38 < 15; fi38++ {
            if false {
                fmt.Println(v3, fi38)
            }
            _ = fi38
        }
        _ = v37
    }
    var v39 uint16 = 1
    if true {
        v4 = 4294967295
        var v40 string = "hello"
        var v41 uint64 = (v23 - v23)
        _ = v40
        _ = v41
    } else {
        fmt.Println(v39, v4)
        switch "\n\t" {
        case "foo bar":
            if false {
                fmt.Println(v4, v4, v39)
                fmt.Println(v23)
            }
        case "a":
            var v42 float64 = (1.5 + ((3.14159 + 1.5) / 1.0))
            _ = v42
        default:
            switch (v4 << ((4294967295) & 63)) {
            case 692:
                v1[(v18 + v18)] = 32767
            case -294:
                var v43 string = "a"
                _ = v43
            default:
                var v44 uint = 2
                _ = v44
            }
        }
        v17 -= v17
    }
    delete(v1, ("" + (((v18 + (v18 + ((("" + (((v18 + v18) + "") + v18)) + v18) + v18))) + v18) + v18)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
}
