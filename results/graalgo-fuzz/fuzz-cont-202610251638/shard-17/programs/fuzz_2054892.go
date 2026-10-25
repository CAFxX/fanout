package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint32 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
func genFn0(p0 int, p1 uint, p2 uint8) uint64 {
    p0 += (p0 - 127)
    v5 := []int{int(p2), (-(p0)), p0, -9223372036854775808}
    fmt.Println(p0, p2, p2)
    var v6 int8 = -106
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 641
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]uint32)
    var v3 float64 = 1e10
    var v4 uint = 255
    for fi7 := 0; fi7 < 2; fi7++ {
        v3 += v3
        _ = fi7
    }
    var v8 uint32 = 211
    var v9 uint64 = 42
    if true {
        switch ("a" + "\n\t") {
        case "":
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint8 = 128
                var v12 float32 = 3.14159
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            v3 = v3
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(v8, v8)
        _ = fi13
    }
    if true {
        var v14 int8 = 100
        v1 = 450
        _ = v14
    } else {
        fmt.Println(v8, v8, v8)
        for fi15 := 0; fi15 < 13; fi15++ {
            v8 -= v8
            if false {
                fmt.Println(fi15, v4, v4)
            }
            _ = fi15
        }
        fmt.Println(v4, v1)
    }
    switch 9223372036854775807 {
    case 4294967295:
        v3 = 0.1
    case -966:
        if (127 >= 4294967295) {
            if true {
                v16 := GenS1{F0: 65535, F1: "", F2: 2}
                v1 = v4
                _ = v16
            }
            delete(v2, ("\n\t" + (("" + "\n\t") + "a")))
            if ((1 > -1) || false) {
                switch ((("\n\t" + "foo bar") + "") + ("foo bar" + "foo bar")) {
                case "foo bar":
                    v3 = 1.5
                case "\n\t":
                    if (217 != 1) {
                        var v17 float32 = float32(v3)
                        _ = v17
                    } else {
                        v18 := []int{42, 32768, 2147483647}
                        v19 := GenS2{F0: (("foo bar" + "a") + "foo bar"), F1: 160}
                        _ = v18
                        _ = v19
                    }
                case "":
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v21 := []int{(^(fi20))}
                        var v22 uint8 = 255
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            } else {
                for fi23 := 0; fi23 < 13; fi23++ {
                    v24 := []int{}
                    for fi25 := 0; fi25 < 16; fi25++ {
                        v24 = append(v24, 0)
                        v26 := GenS3{F0: (fi25 ^ fi25)}
                        _ = fi25
                        _ = v26
                    }
                    var v27 int16 = 0
                    _ = fi23
                    _ = v24
                    _ = v27
                }
            }
        } else {
            _ = len(v2)
            v28 := GenS3{F0: 42}
            _ = v28
        }
    case -1:
        fmt.Println(v1, v8, v8)
    }
    switch "foo bar" {
    case "foo bar":
        if true {
            switch -2147483648 {
            case 9223372036854775807:
                v9 = (v9 >> (((v4 + (v1 >> ((4294967295) & 63)))) & 63))
            case -1:
                var v29 int = 127
                _ = v29
            case 42:
                switch -1 {
                case -2147483648:
                    v4 -= v1
                case 127:
                    switch "" {
                    case "foo bar":
                        fmt.Println(v4, v3)
                    }
                case 4294967295:
                    if false {
                        v1 += (^(v1))
                        if false {
                            for fi30 := 0; fi30 < 2; fi30++ {
                                v9 -= v9
                                v8 += 2147483648
                                _ = fi30
                            }
                            v31 := v2[("" + ("hello" + "a"))]
                            _ = v31
                        } else {
                            var v32 int64 = -326
                            if true {
                                var v33 uint = 445
                                _ = v33
                            } else {
                                var v34 bool = (21 < 128)
                                if false {
                                    delete(v2, "\n\t")
                                }
                                v8 = 42
                                _ = v34
                            }
                            _ = v32
                        }
                    } else {
                        if (2147483647 < -1) {
                            v3 += 0.0
                            _ = len(v2)
                        } else {
                            v35 := GenS3{F0: -469}
                            _ = v35
                        }
                        var v36 uint8 = uint8(v9)
                        v36 -= v36
                        _ = v36
                    }
                }
            default:
                v37 := GenS2{F0: ("a" + ("a" + ((("" + ("hello" + ("" + "hello"))) + ("a" + "")) + "hello"))), F1: -871}
                _ = v37
            }
            for fi38 := 0; fi38 < 7; fi38++ {
                var v39 bool = false
                v40 := GenS3{F0: (fi38 / ((fi38) | 1))}
                _ = fi38
                _ = v39
                _ = v40
            }
        }
    case "hello":
        v2["hello"] = (v8 + (v8 | 4294967295))
    }
    switch "a" {
    case "\n\t":
        v1 -= v4
    }
    if (genFn0(-595, v4, 1) != v9) {
        fmt.Println(v3)
        if (4294967295 < 127) {
            switch 127 {
            case -1:
                v8 = 2
            }
            switch (("\n\t" + "hello") + (("" + "foo bar") + ("\n\t" + ""))) {
            case "a":
                v4 -= 2
            case "hello":
                var v41 uint16 = 255
                _ = v41
            case "\n\t":
                v42 := []int{0, -1, 4294967295, 426}
                _ = v42
            }
        } else {
            v1 -= (v4 - 65535)
            v43 := GenS1{F0: 1, F1: (("a" + (("\n\t" + "a") + ("a" + ""))) + "hello"), F2: 2}
            _ = v43
        }
        if true {
            if true {
                v1 = 1
                var v44 uint8 = 0
                var v45 uint64 = 18446744073709551615
                _ = v44
                _ = v45
            }
            v46 := GenS2{F0: "\n\t", F1: -569}
            for fi47 := 0; fi47 < 17; fi47++ {
                if true {
                    v48 := []int{32768, -1, (fi47 ^ fi47), 127}
                    _ = v48
                } else {
                    for fi49 := 0; fi49 < 7; fi49++ {
                        v4 -= v4
                        var v50 uint = 4294967295
                        if (true && (false && true)) {
                            var v51 uint32 = 365
                            _ = v51
                        }
                        var v52 uint8 = uint8(v9)
                        switch fi49 {
                        case 4294967295:
                            v8 = v8
                        case -1:
                            v53 := []int{-2147483648, fi47, fi49}
                            _ = v53
                        case 128:
                            var v54 float32 = -2.25
                            _ = v54
                        }
                        _ = fi49
                        _ = v50
                        _ = v52
                    }
                }
                _ = fi47
            }
            _ = v46
        }
    }
    var v55 float64 = v3
    if (!(231 > 354)) {
        v8 = (^(v8))
        v8 = v8
        var v56 float32 = 3.14159
        _ = v56
    }
    delete(v2, "foo bar")
    var v57 uint32 = 2147483648
    v58 := GenS4{F0: 0}
    delete(v2, ((("hello" + ("a" + ((("hello" + "foo bar") + "a") + "\n\t"))) + "a") + (("a" + "") + "a")))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v57
    _ = v58
    _ = v8
    _ = v9
}
