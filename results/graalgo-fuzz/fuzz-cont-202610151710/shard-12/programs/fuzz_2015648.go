package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 int8 }
type GenS5 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0(p0 uint, p1 bool, p2 uint32) uint64 {
    p1 = false
    fmt.Println(p2, p0, p1)
    if p1 {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 float64 = -2.25
            v7 := GenS1{F0: -123}
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        if p1 {
            p0 = p0
            var v8 uint = 18446744073709551615
            _ = v8
        } else {
            v9 := []int{2147483647, 888}
            var v10 int16 = 1000
            _ = v10
            _ = v9
        }
        var v11 uint16 = 67
        switch (("a" + "\n\t") + ("" + "\n\t")) {
        case "a":
            var v12 uint32 = p2
            _ = v12
        case "":
            var v13 float64 = (0.0 / 3.14159)
            _ = v13
        default:
            v14 := []int{}
            _ = v14
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int64 = 42
    v2 := []int{127, 631, 733}
    v3 := make(map[int]int)
    var v4 string = "a"
    var v15 int64 = (v1 - v1)
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            if false {
                v4 = v4
                var v18 uint8 = 171
                v1 -= v15
                _ = v18
            } else {
                v1 -= v15
                v19 := GenS1{F0: 100}
                _ = v19
            }
            var v20 int64 = (v1 & v15)
            if (!(!true)) {
                v4 = ("a" + v4)
            } else {
                if (false && (true && (!("a" <= (v4 + ""))))) {
                    var v21 string = v4
                    var v22 bool = true
                    if len(v2) > 0 {
                        v2[0] = -1
                    }
                    _ = v21
                    _ = v22
                }
                switch ("\n\t" + "hello") {
                case "":
                    switch fi17 {
                    case 127:
                        _ = len(v3)
                    case 1:
                        switch fi17 {
                        case 2147483647:
                            if (true && ((0 == (v1 << ((42) & 63))) || (42 < 105))) {
                                var v23 int64 = -573
                                v24 := GenS1{F0: -128}
                                v25 := v3[(fi17 % ((9223372036854775807) | 1))]
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS1{F0: 127}
                            _ = v26
                        case 9223372036854775807:
                            if len(v2) > 0 {
                                v2[0] = fi16
                            }
                            var v27 uint8 = 2
                            _ = v27
                        default:
                            fmt.Println(v20)
                        }
                    }
                case "foo bar":
                    if false {
                        if (v4 == v4) {
                            switch (("hello" + "hello") + v4) {
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "hello":
                                v15 = v1
                            default:
                                v28 := GenS1{F0: 42}
                                _ = v28
                            }
                        } else {
                            v29 := v3[fi17]
                            var v30 int64 = (v15 % (((v1 + v20)) | 1))
                            _ = v29
                            _ = v30
                        }
                    }
                case "\n\t":
                    v31 := GenS2{F0: v15}
                    _ = v31
                }
                var v32 int32 = 0
                _ = v32
            }
            _ = fi17
            _ = v20
        }
        _ = fi16
    }
    if true {
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 int16 = 499
            var v35 uint = 553
            _ = fi33
            _ = v34
            _ = v35
        }
        var v36 float32 = float32(0.0)
        _ = v36
    }
    switch 2147483647 {
    case 67:
        _ = len(v3)
    default:
        if (((v4 + "a") == "a") || false) {
            var v37 float32 = float32(-2.25)
            if (127 >= 42) {
                var v38 float64 = 3.14159
                _ = len(v2)
                _ = v38
            } else {
                v37 = v37
            }
            _ = v37
        } else {
            fmt.Println(v1, v15)
            for fi39 := 0; fi39 < 3; fi39++ {
                v4 = "\n\t"
                _ = fi39
            }
            var v40 int8 = 101
            _ = v40
        }
    }
    var v41 uint16 = 65535
    fmt.Println(v41, v1)
    v42 := GenS3{F0: 42, F1: -128}
    switch v4 {
    case "a":
        v41 = (v41 | v41)
    case "\n\t":
        v41 -= (^(v41))
    case "foo bar":
        v43 := GenS4{F0: 255, F1: 65535, F2: -1}
        _ = v43
    default:
        v44 := GenS1{F0: -122}
        _ = v44
    }
    fmt.Println(v4)
    var v45 int = 23
    v46 := GenS1{F0: -58}
    var v47 uint64 = 4294967295
    var v48 uint32 = 7
    var v49 int8 = 67
    switch v4 {
    case "a":
        v50 := GenS5{F0: v4, F1: (v41 >> ((18446744073709551615) & 63)), F2: v41}
        _ = v50
    default:
        var v51 int16 = 0
        _ = v51
    }
    var v52 int64 = v15
    v53 := GenS4{F0: 42, F1: uint32(v41), F2: v49}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v52
    _ = v53
}
