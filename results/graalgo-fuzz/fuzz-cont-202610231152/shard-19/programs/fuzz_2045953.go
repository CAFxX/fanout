package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int; F2 int16 }
func genFn0() int {
    for fi5 := 0; fi5 < 10; fi5++ {
        if false {
            var v6 uint8 = 5
            v6 = v6
            v6 = 2
            _ = v6
        } else {
            if (false && true) {
                var v7 uint16 = 255
                _ = v7
            }
        }
        _ = fi5
    }
    var v8 int8 = 100
    var v9 uint64 = 254
    v10 := GenS1{F0: 347}
    _ = v10
    _ = v8
    _ = v9
    return 62
}
func genFn1(p0 bool, p1 uint, p2 int16) string {
    if true {
        fmt.Println(p0)
        v11 := GenS1{F0: 127}
        _ = v11
    }
    var v12 int8 = -62
    p0 = (0 > (p1 ^ (p1 & p1)))
    var v13 float64 = 1.5
    v14 := GenS2{F0: 255, F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return ("\n\t" + "a")
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]uint32)
    var v3 float32 = 1e10
    v4 := []uint32{65535, 65535, 2147483647, 0}
    v1 += v1
    var v15 int64 = v1
    switch (("" + "hello") + "hello") {
    case "foo bar":
        switch 42 {
        case -243:
            fmt.Println(v3, v1, v1)
        default:
            v16 := GenS3{F0: -2147483648}
            _ = v16
        }
    case "hello":
        if false {
            var v17 float64 = -2.25
            _ = v17
        } else {
            var v18 uint = 1
            _ = v18
        }
    case "\n\t":
        v19 := GenS1{F0: 371}
        _ = v19
    default:
        var v20 int64 = -386
        _ = v20
    }
    if (850 != -793) {
        var v21 int32 = 41
        v2[127] = 2147483647
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := GenS4{F0: 4294967295, F1: genFn0(), F2: 1000}
        _ = fi22
        _ = v23
    }
    switch (("hello" + ("hello" + ("a" + ("\n\t" + "a")))) + (("" + "\n\t") + ("foo bar" + "\n\t"))) {
    case "hello":
        var v24 int64 = (v15 << ((255) & 63))
        _ = v24
    }
    delete(v2, genFn0())
    switch ("" + "\n\t") {
    case "\n\t":
        v15 = 0
    case "":
        if (!(18446744073709551615 > 227)) {
            if (0 >= 2) {
                if ((float32((1e10 + ((1e10 / (-2.25 + (((float64(v3) + float64((v3 - v3))) - ((float64(v3) / float64(float32(1e10))) * 1.5)) * 1e10))) + 1.5))) + 0.1) != (v3 * v3)) {
                    v15 = v15
                } else {
                    switch 0 {
                    case 42:
                        var v25 uint8 = 1
                        _ = v25
                    case 2147483647:
                        switch 2147483647 {
                        case 42:
                            var v26 uint64 = 18446744073709551615
                            _ = v26
                        case -1:
                            switch ("a" + ("" + (("\n\t" + ("a" + ("a" + "foo bar"))) + "\n\t"))) {
                            case "\n\t":
                                v27 := GenS1{F0: 0}
                                _ = v27
                            case "foo bar":
                                v28 := GenS2{F0: 2, F1: 492}
                                _ = v28
                            case "a":
                                v29 := GenS3{F0: -2147483648}
                                _ = v29
                            }
                        default:
                            var v30 string = "\n\t"
                            _ = v30
                        }
                    case -926:
                        if (!true) {
                            var v31 string = ""
                            var v32 string = "hello"
                            _ = v31
                            _ = v32
                        } else {
                            v33 := GenS2{F0: 133, F1: 964}
                            _ = v33
                        }
                    default:
                        v2[42] = 562
                    }
                    v34 := GenS3{F0: -2147483648}
                    switch (("a" + ("" + "a")) + ("" + "")) {
                    case "":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "\n\t":
                        var v35 float64 = (-2.25 * 0.1)
                        _ = v35
                    case "hello":
                        v3 += ((((v3 * -2.25) / 1.5) / (((v3 / (float32(3.14159) / (v3 / (v3 * float32(0.0))))) + v3) / v3)) + -2.25)
                    default:
                        v36 := GenS4{F0: 65535, F1: -776, F2: -186}
                        _ = v36
                    }
                    _ = v34
                }
                v15 += v1
            }
            var v37 uint = 2
            _ = v37
        }
    default:
        _ = len(v2)
    }
    v15 = (v1 ^ -232)
    for fi38 := 0; fi38 < 18; fi38++ {
        switch fi38 {
        case 4294967295:
            if false {
                fmt.Println(v1)
                v1 = 848
            } else {
                for fi39 := 0; fi39 < 19; fi39++ {
                    v40 := GenS1{F0: (fi38 + (^(fi38)))}
                    _ = fi39
                    _ = v40
                }
                var v41 int64 = (-(v1))
                _ = v41
            }
        }
        var v42 uint64 = 19
        switch ("hello" + "a") {
        case "":
            fmt.Println(v15)
        }
        _ = fi38
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
