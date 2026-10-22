package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint16; F2 string }
type GenS4 struct { F0 uint64 }
func genFn0() uint32 {
    var v5 uint32 = 2147483647
    v6 := []int{1, 9223372036854775807, 635, -719}
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 string) uint8 {
    p0 = "hello"
    if false {
        v7 := GenS1{F0: -32768, F1: 665, F2: 466}
        _ = v7
    } else {
        var v8 int = 736
        v9 := []int{132, v8}
        for fi10 := 0; fi10 < 19; fi10++ {
            switch (v8 - fi10) {
            case 0:
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
            fmt.Println(fi10, v8, fi10)
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 uint64 = 42
                switch v8 {
                case 952:
                    var v13 float32 = -2.25
                    _ = v13
                case 4294967295:
                    v14 := GenS1{F0: 32767, F1: -2147483648, F2: v12}
                    _ = v14
                case -1:
                    fmt.Println(fi11, fi10, fi10)
                }
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    switch "foo bar" {
    case "hello":
        switch p0 {
        case "foo bar":
            var v15 float32 = 3.14159
            _ = v15
        default:
            v16 := []int{32767}
            _ = v16
        }
    case "":
        v17 := []int{32767}
        _ = v17
    case "foo bar":
        if false {
            v18 := []int{973, 161}
            var v19 uint32 = 758
            _ = v18
            _ = v19
        } else {
            if (false || (1.5 < float32(0.0))) {
                var v20 string = (p0 + "hello")
                _ = v20
            }
            fmt.Println(p0)
            var v21 int = 32768
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 11; fi22++ {
            p0 = p0
            if true {
                var v23 int = (fi22 | (fi22 & fi22))
                if false {
                    v24 := []int{}
                    var v25 int = v23
                    for fi26 := 0; fi26 < 7; fi26++ {
                        if len(v24) > 0 {
                            v24[0] = 1
                        }
                        _ = len(v24)
                        var v27 int8 = 100
                        _ = fi26
                        _ = v27
                    }
                    _ = v24
                    _ = v25
                } else {
                    v28 := []int{(v23 ^ v23), fi22, fi22}
                    v23 = (-(fi22))
                    _ = v28
                }
                v23 -= (fi22 ^ fi22)
                _ = v23
            } else {
                for fi29 := 0; fi29 < 9; fi29++ {
                    p0 = "foo bar"
                    var v30 int32 = 1
                    v31 := []int{(fi29 ^ (fi29 << ((255) & 63)))}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                p0 = ("foo bar" + "")
            }
            _ = fi22
        }
    }
    _ = p0
    return 42
}

func main() {
    v1 := []int{4294967295}
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 831
    var v4 float64 = (1.5 - 1e10)
    var v32 int64 = v2
    var v33 string = "hello"
    var v34 float32 = 0.0
    for fi35 := 0; fi35 < 3; fi35++ {
        v36 := GenS1{F0: 1000, F1: 435, F2: v3}
        var v37 uint = 4294967295
        fmt.Println(fi35)
        _ = fi35
        _ = v36
        _ = v37
    }
    if ((2 >= 251) || false) {
        for fi38 := 0; fi38 < 9; fi38++ {
            var v39 int64 = v2
            _ = fi38
            _ = v39
        }
        var v40 uint = 843
        _ = v40
    } else {
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
    }
    var v41 uint64 = v3
    var v42 uint64 = (v41 - v3)
    _ = len(v1)
    var v43 string = ("foo bar" + "\n\t")
    _ = len(v1)
    var v44 uint = 1
    switch v33 {
    case "foo bar":
        for fi45 := 0; fi45 < 16; fi45++ {
            fmt.Println(v2, v42, v33)
            for fi46 := 0; fi46 < 17; fi46++ {
                var v47 uint32 = 1
                switch ("hello" + "a") {
                case "\n\t":
                    var v48 int = 32767
                    _ = v48
                case "hello":
                    switch -1 {
                    case 127:
                        v49 := GenS2{F0: true}
                        _ = v49
                    }
                }
                if (65535 < 367) {
                    switch fi45 {
                    case 9223372036854775807:
                        if true {
                            var v50 uint32 = v47
                            var v51 int = (fi46 << (((^(v44))) & 63))
                            v44 = (v44 & v44)
                            _ = v50
                            _ = v51
                        } else {
                            var v52 int = (fi45 << ((1) & 63))
                            if len(v1) > 0 {
                                v1[0] = fi46
                            }
                            _ = v52
                        }
                    default:
                        v2 = (v2 - 42)
                    }
                    var v53 bool = true
                    _ = v53
                }
                _ = fi46
                _ = v47
            }
            v54 := GenS3{F0: 1, F1: 521, F2: ((((("foo bar" + v33) + v43) + (((((v33 + v33) + ((("hello" + v33) + v33) + v43)) + "\n\t") + (v33 + ((v43 + "\n\t") + ((v43 + "foo bar") + (v43 + v43))))) + v43)) + ("a" + "\n\t")) + "hello")}
            _ = fi45
            _ = v54
        }
    }
    if (1 >= 42) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi55 := 0; fi55 < 3; fi55++ {
            v56 := GenS4{F0: (v3 ^ (v42 - (v41 - v41)))}
            _ = fi55
            _ = v56
        }
    }
    fmt.Println(v33, v43, v33)
    var v57 float32 = v34
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v57
}
