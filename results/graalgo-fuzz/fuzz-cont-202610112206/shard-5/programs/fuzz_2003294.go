package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int32; F2 bool }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0(p0 bool) int32 {
    v4 := GenS1{F0: 42, F1: "foo bar"}
    if ("a" < "foo bar") {
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 float32 = -2.25
            for fi7 := 0; fi7 < 9; fi7++ {
                if (false || (89 > 0)) {
                    var v8 bool = false
                    var v9 uint8 = 1
                    var v10 string = (((("foo bar" + "a") + ("foo bar" + "a")) + "hello") + (("\n\t" + "hello") + ("a" + "foo bar")))
                    _ = v10
                    _ = v8
                    _ = v9
                } else {
                    switch "\n\t" {
                    case "hello":
                        var v11 int16 = -785
                        _ = v11
                    case "\n\t":
                        v12 := []int{(^(fi5))}
                        _ = len(v12)
                        v13 := GenS2{F0: 638}
                        _ = v12
                        _ = v13
                    }
                    fmt.Println(fi7, v6, p0)
                }
                _ = fi7
            }
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := []int{(fi14 >> ((2) & 63)), fi5, (fi14 - 895)}
                    v15 = append(v15, (fi5 % ((fi14) | 1)))
                    var v16 uint16 = 255
                    var v17 uint16 = v16
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            case "hello":
                v18 := GenS3{F0: -1, F1: -894, F2: p0}
                _ = v18
            case "a":
                for fi19 := 0; fi19 < 6; fi19++ {
                    if p0 {
                        switch fi19 {
                        case -9223372036854775808:
                            v6 += ((((-2.25 + v6) - (-2.25 + v6)) / 0.1) + -2.25)
                        case -2147483648:
                            if ((v6 - v6) >= v6) {
                                switch ((("a" + "\n\t") + "a") + "a") {
                                case "":
                                    var v20 uint32 = 65535
                                    _ = v20
                                }
                                var v21 uint = 2
                                v22 := []int{220, (fi19 + fi5)}
                                _ = v21
                                _ = v22
                            } else {
                                var v23 int16 = 1000
                                _ = v23
                            }
                        case 0:
                            var v24 int64 = -1
                            _ = v24
                        default:
                            switch ("a" + (("" + (("foo bar" + ("foo bar" + ("a" + (("a" + "hello") + "a")))) + "foo bar")) + "hello")) {
                            case "\n\t":
                                p0 = true
                            }
                        }
                        if p0 {
                            var v25 uint32 = 1
                            var v26 int8 = 0
                            var v27 int = fi5
                            var v28 bool = (1 == 1)
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            fmt.Println(fi5, v6, v6)
                            v29 := []int{32768, fi19, (fi5 | fi5)}
                            _ = v29
                        }
                        v6 -= v6
                        fmt.Println(v6, fi19, p0)
                    }
                    _ = fi19
                }
            }
            _ = fi5
            _ = v6
        }
    } else {
        p0 = p0
        var v30 int8 = -94
        v31 := GenS4{F0: 4294967295, F1: 255}
        _ = v30
        _ = v31
    }
    p0 = (884 < 255)
    p0 = false
    var v32 uint = 2
    _ = p0
    _ = v32
    _ = v4
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    var v33 float64 = (1e10 + 0.0)
    delete(v1, "\n\t")
    switch ((("hello" + "a") + ("foo bar" + "")) + "\n\t") {
    case "foo bar":
        v34 := GenS5{F0: 289, F1: 65535, F2: 2}
        _ = v34
    case "a":
        v35 := GenS5{F0: 464, F1: 32767, F2: 2}
        _ = v35
    }
    var v36 int8 = 100
    var v37 int = 42
    for fi38 := 0; fi38 < 8; fi38++ {
        v36 = (v36 & (v36 * v36))
        _ = len(v1)
        v2[1] = ((!(v33 == v33)) && true)
        _ = fi38
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        if true {
            v36 = v36
        }
        if ((true || ((((127 != 2) || false) || false) && false)) || ("hello" < "hello")) {
            v37 = (v37 << ((2) & 63))
        } else {
            var v40 int32 = genFn0(true)
            fmt.Println(v36, fi39, v36)
            _ = v40
        }
        v41 := GenS6{F0: v36, F1: 4294967295}
        _ = fi39
        _ = v41
    }
    v42 := []int{(v37 & (-(v37)))}
    fmt.Println(v36)
    fmt.Println(v37, v37)
    fmt.Println(v33, v36)
    switch (v37 << ((920) & 63)) {
    case 2147483647:
        for fi43 := 0; fi43 < 13; fi43++ {
            var v44 uint8 = 0
            fmt.Println(fi43, v33)
            _ = fi43
            _ = v44
        }
    case 32768:
        for fi45 := 0; fi45 < 10; fi45++ {
            v46 := v2[2147483647]
            _ = fi45
            _ = v46
        }
    }
    v33 += (v33 + v33)
    var v47 int32 = 42
    var v48 int32 = v47
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v42
    _ = v47
    _ = v48
}
