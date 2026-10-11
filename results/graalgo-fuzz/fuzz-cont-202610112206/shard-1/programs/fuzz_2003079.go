package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 int8; F1 bool }
func genFn0() int64 {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int32 = int32(fi5)
        for fi7 := 0; fi7 < 3; fi7++ {
            v6 = (v6 / (((^(v6))) | 1))
            v6 += v6
            _ = fi7
        }
        fmt.Println(v6, fi5)
        _ = fi5
        _ = v6
    }
    var v8 uint16 = 65535
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 uint8) uint8 {
    var v9 uint = 42
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 bool = false
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := []int{fi12, fi10, fi12}
            _ = fi12
            _ = v13
        }
        var v14 uint32 = 2
        _ = fi10
        _ = v11
        _ = v14
    }
    v15 := GenS1{F0: (v9 | 65535)}
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int8 = 0
        _ = fi16
        _ = v17
    }
    var v18 int64 = 0
    _ = p0
    _ = p1
    _ = v15
    _ = v18
    _ = v9
    return p1
}

func main() {
    var v1 bool = (false || true)
    var v2 int64 = 895
    var v3 float64 = ((1.5 * 0.0) + 1.5)
    var v4 uint64 = 42
    if ((4294967295 >= 992) && (2 <= 4294967295)) {
        if v1 {
            v1 = (408 != 847)
            for fi19 := 0; fi19 < 14; fi19++ {
                v1 = false
                _ = fi19
            }
            v20 := GenS2{F0: 1}
            _ = v20
        }
    } else {
        v4 = v4
        var v21 uint32 = 42
        _ = v21
    }
    fmt.Println(v3, v4)
    if false {
        v22 := GenS1{F0: 971}
        switch (((("hello" + ("foo bar" + ("" + ""))) + ("a" + (("foo bar" + ("hello" + "\n\t")) + ""))) + "\n\t") + "") {
        case "\n\t":
            v23 := GenS3{F0: 2147483647, F1: "a"}
            _ = v23
        case "hello":
            switch ("hello" + "") {
            case "a":
                v24 := GenS1{F0: 42}
                _ = v24
            case "":
                v25 := []int{-2147483648}
                _ = v25
            case "foo bar":
                for fi26 := 0; fi26 < 7; fi26++ {
                    if false {
                        switch (fi26 + fi26) {
                        case 0:
                            switch ("\n\t" + ("" + "hello")) {
                            case "a":
                                if (!((v2 ^ genFn0()) >= v2)) {
                                    if false {
                                        v27 := GenS1{F0: 65535}
                                        var v28 int32 = -317
                                        _ = v27
                                        _ = v28
                                    } else {
                                        if ("foo bar" == "a") {
                                            v1 = v1
                                            var v29 uint64 = 4294967295
                                            _ = v29
                                        } else {
                                            v30 := GenS3{F0: 2147483647, F1: ((("a" + "a") + "foo bar") + "\n\t")}
                                            _ = v30
                                        }
                                    }
                                    v4 -= 18446744073709551615
                                }
                            case "":
                                v31 := GenS4{F0: true}
                                _ = v31
                            default:
                                var v32 uint64 = v4
                                _ = v32
                            }
                        case 2147483647:
                            v33 := []int{}
                            _ = v33
                        case 866:
                            v4 -= (v4 - 4294967295)
                        default:
                            if false {
                                v34 := []int{279, (fi26 | (fi26 >> ((311) & 63))), fi26, fi26}
                                _ = v34
                            }
                        }
                    } else {
                        v2 += (v2 % ((genFn0()) | 1))
                        v3 -= -2.25
                    }
                    _ = fi26
                }
            }
        case "":
            if v1 {
                v3 -= 1.5
                v35 := GenS5{F0: 1}
                _ = v35
            } else {
                var v36 uint64 = v4
                fmt.Println(v1)
                v4 = 2
                _ = v36
            }
        }
        _ = v22
    }
    v37 := GenS6{F0: 1, F1: 32767}
    v1 = v1
    fmt.Println(v3, v3)
    v38 := GenS5{F0: 0}
    fmt.Println(v3, v3)
    v39 := []int{574, -948}
    v2 -= v2
    if (((0.0 / 1e10) / (-2.25 + (v3 + 3.14159))) < v3) {
        if (!true) {
            switch 4294967295 {
            case 128:
                v39 = append(v39, -9223372036854775808)
            case 32767:
                v3 -= v3
            default:
                var v40 uint32 = 824
                _ = v40
            }
            v41 := GenS3{F0: 973, F1: "\n\t"}
            var v42 int64 = v2
            _ = v41
            _ = v42
        }
        fmt.Println(v3)
        v4 = v4
    } else {
        v43 := GenS7{F0: 100, F1: (0.1 > v3)}
        if len(v39) > 0 {
            v39[0] = -431
        }
        var v44 uint8 = 127
        _ = v43
        _ = v44
    }
    if len(v39) >= 2 {
        v39 = v39[1:len(v39)]
    }
    var v45 int64 = 1
    var v46 int = 128
    v39 = append(v39, (v46 >> ((18446744073709551615) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v45
    _ = v46
}
