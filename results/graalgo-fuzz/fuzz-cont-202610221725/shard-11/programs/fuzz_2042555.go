package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int32; F2 int64 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint16 }
func genFn0() uint16 {
    var v5 int8 = 18
    fmt.Println(v5)
    if true {
        var v6 string = "\n\t"
        _ = v6
    } else {
        v5 -= v5
    }
    if false {
        if true {
            if false {
                fmt.Println(v5)
                fmt.Println(v5)
                v7 := GenS1{F0: 4294967295}
                _ = v7
            }
        } else {
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 float64 = (1e10 / 1.0)
                _ = fi8
                _ = v9
            }
        }
    } else {
        switch ("hello" + "") {
        case "":
            switch "hello" {
            case "\n\t":
                fmt.Println(v5)
            case "a":
                v5 += (v5 >> ((42) & 63))
            case "foo bar":
                for fi10 := 0; fi10 < 16; fi10++ {
                    var v11 int32 = 0
                    for fi12 := 0; fi12 < 13; fi12++ {
                        switch -936 {
                        case 32768:
                            v13 := []int{fi12, -617}
                            _ = v13
                        }
                        switch (fi10 << ((4294967295) & 63)) {
                        case 128:
                            var v14 uint = 255
                            _ = v14
                        case 549:
                            var v15 bool = false
                            _ = v15
                        }
                        v5 = v5
                        _ = fi12
                    }
                    v16 := GenS2{F0: (^(v5)), F1: (v11 + v11), F2: -697}
                    _ = fi10
                    _ = v11
                    _ = v16
                }
            }
        }
    }
    _ = v5
    return 1
}
func genFn1(p0 uint16, p1 uint64) uint64 {
    p1 += 0
    v17 := GenS3{F0: (p1 | (p1 << ((42) & 63))), F1: ("" == ""), F2: (^(p0))}
    for fi18 := 0; fi18 < 15; fi18++ {
        p0 = genFn0()
        if (!true) {
            fmt.Println(p1, fi18, p0)
            switch "hello" {
            case "foo bar":
                v19 := []int{(fi18 & (^(fi18))), fi18, 4294967295, (fi18 << ((714) & 63))}
                _ = v19
            default:
                if (fi18 == fi18) {
                    switch "" {
                    case "\n\t":
                        v20 := []int{fi18, -1, 128}
                        _ = v20
                    default:
                        var v21 uint = 824
                        _ = v21
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v22 := GenS3{F0: p1, F1: ("a" != ("a" + "")), F2: (p0 & p0)}
                        _ = v22
                    case "hello":
                        var v23 uint32 = 4294967295
                        _ = v23
                    }
                    var v24 string = "foo bar"
                    _ = v24
                } else {
                    v25 := []int{(^(fi18)), 32767, (fi18 << ((255) & 63)), 2147483647}
                    v26 := GenS3{F0: p1, F1: true, F2: p0}
                    var v27 float64 = 1e10
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
            if true {
                p1 += p1
                v28 := GenS1{F0: 65535}
                p0 = uint16(p1)
                _ = v28
            }
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = v17
    return p1
}

func main() {
    v1 := []int{-1, 4294967295, 32767, 790}
    v2 := make(map[string]int)
    v3 := []int{2147483647, -2147483648}
    var v4 int8 = -1
    fmt.Println(v4)
    var v29 uint = 65535
    var v30 int64 = -1
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 int32 = -96
        _ = fi31
        _ = v32
    }
    v4 -= v4
    fmt.Println(v29, v29)
    if len(v1) > 0 {
        v1[0] = 949
    }
    var v33 bool = ((314 >= v29) || (2 == genFn0()))
    var v34 int32 = -1
    for fi35 := 0; fi35 < 3; fi35++ {
        delete(v2, "hello")
        _ = fi35
    }
    v4 = (-(v4))
    var v36 bool = (!(16 <= 127))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v36
    _ = v4
}
