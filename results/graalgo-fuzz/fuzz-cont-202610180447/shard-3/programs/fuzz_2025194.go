package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint8, p1 int16, p2 int8) bool {
    for fi5 := 0; fi5 < 19; fi5++ {
        switch "\n\t" {
        case "\n\t":
            p0 = 224
        case "hello":
            if false {
                p1 = 42
                var v6 float32 = 0.0
                _ = v6
            } else {
                if false {
                    for fi7 := 0; fi7 < 5; fi7++ {
                        v8 := []int{fi5}
                        p0 = (p0 >> ((897) & 63))
                        _ = fi7
                        _ = v8
                    }
                    if false {
                        if ((true && true) || true) {
                            var v9 float32 = 1e10
                            v10 := []int{(fi5 % ((int(p0)) | 1)), (fi5 * (fi5 >> ((168) & 63)))}
                            var v11 float64 = 1e10
                            _ = v10
                            _ = v11
                            _ = v9
                        }
                        for fi12 := 0; fi12 < 1; fi12++ {
                            if true {
                                v13 := GenS1{F0: ("\n\t" + ("a" + "\n\t"))}
                                _ = v13
                            } else {
                                p0 = (p0 | 178)
                                var v14 float32 = 3.14159
                                var v15 uint8 = (p0 ^ (p0 ^ (^(p0))))
                                p2 = p2
                                _ = v14
                                _ = v15
                            }
                            p0 = (p0 >> ((4294967295) & 63))
                            _ = fi12
                        }
                    }
                }
                fmt.Println(p1, p2, p1)
                fmt.Println(p0, fi5, p2)
            }
        }
        _ = fi5
    }
    p2 -= (p2 * (p2 * p2))
    v16 := []int{}
    for fi17 := 0; fi17 < 9; fi17++ {
        switch fi17 {
        case -654:
            var v18 uint = 255
            _ = v18
        }
        _ = fi17
    }
    if len(v16) > 0 {
        v16[0] = int(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return true
}
func genFn1() uint64 {
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 int8 = -128
        v21 := GenS1{F0: ""}
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS1{F0: "a"}
            v20 += v20
            _ = fi22
            _ = v23
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    if ("" >= "\n\t") {
        if true {
            var v24 int32 = 1
            switch 958 {
            case -2147483648:
                for fi25 := 0; fi25 < 14; fi25++ {
                    v26 := GenS1{F0: "\n\t"}
                    _ = fi25
                    _ = v26
                }
            case -82:
                v24 = 723
            default:
                var v27 bool = false
                _ = v27
            }
            _ = v24
        }
        var v28 float64 = 0.1
        fmt.Println(v28)
        _ = v28
    } else {
        var v29 int8 = 42
        _ = v29
    }
    var v30 int8 = -128
    _ = v30
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 255
    v3 := make(map[string]string)
    var v4 uint8 = uint8(v2)
    switch 1 {
    case 42:
        switch 128 {
        case 127:
            v31 := []int{-538, 0, -1}
            _ = v31
        }
    case -1:
        switch "foo bar" {
        case "":
            v32 := GenS1{F0: "hello"}
            _ = v32
        case "foo bar":
            var v33 uint16 = 0
            _ = v33
        case "a":
            v34 := GenS2{F0: v2, F1: 761}
            _ = v34
        default:
            v35 := GenS2{F0: (^(v2)), F1: 849}
            _ = v35
        }
    }
    if (0 <= 43) {
        var v36 int8 = -128
        _ = v36
    }
    switch ("" + "hello") {
    case "":
        var v37 int = 2147483647
        _ = v37
    case "foo bar":
        for fi38 := 0; fi38 < 18; fi38++ {
            for fi39 := 0; fi39 < 19; fi39++ {
                var v40 uint64 = 0
                _ = fi39
                _ = v40
            }
            var v41 float32 = 0.0
            _ = fi38
            _ = v41
        }
    default:
        var v42 uint32 = 42
        _ = v42
    }
    v4 = (v4 ^ 22)
    v1["foo bar"] = true
    fmt.Println(v4, v4)
    var v43 int = 128
    var v44 string = "\n\t"
    v44 = v44
    fmt.Println(v43, v43, v43)
    for fi45 := 0; fi45 < 17; fi45++ {
        v46 := GenS1{F0: v44}
        _ = fi45
        _ = v46
    }
    v47 := GenS1{F0: (("" + (v44 + v44)) + "hello")}
    for fi48 := 0; fi48 < 7; fi48++ {
        v43 += v43
        var v49 bool = true
        _ = fi48
        _ = v49
    }
    v50 := []int{v43}
    v51 := GenS1{F0: (((v44 + "hello") + (v44 + v44)) + v44)}
    switch v43 {
    case 491:
        v44 = ("hello" + ("foo bar" + (v44 + v44)))
    case 128:
        fmt.Println(v2)
    default:
        v52 := v3["hello"]
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v47
    _ = v50
    _ = v51
}
