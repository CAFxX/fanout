package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint, p1 int16, p2 uint) uint8 {
    v7 := []int{-377}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 uint64 = 18446744073709551615
    if false {
        var v9 float32 = 3.14159
        switch ("hello" + ("a" + ("a" + "hello"))) {
        case "\n\t":
            if (-2.25 <= 1.5) {
                var v10 int32 = -801
                fmt.Println(p0, p2, p0)
                var v11 float32 = float32(0.0)
                _ = v10
                _ = v11
            }
        case "foo bar":
            var v12 int8 = 1
            _ = v12
        case "":
            if false {
                for fi13 := 0; fi13 < 9; fi13++ {
                    v9 -= 0.1
                    v7 = append(v7, (fi13 | (fi13 % ((-747) | 1))))
                    _ = fi13
                }
                switch 0 {
                case -408:
                    if (0 == 2147483647) {
                        var v14 string = ((("" + "a") + ((("foo bar" + (("\n\t" + "foo bar") + ("foo bar" + ("foo bar" + "")))) + "hello") + (("foo bar" + "hello") + "a"))) + ("a" + ("a" + "foo bar")))
                        p1 = (p1 * (p1 + p1))
                        _ = v14
                    }
                }
                if (100 == 1) {
                    var v15 uint32 = 42
                    v8 = v8
                    v16 := GenS1{F0: 42, F1: -692, F2: 900}
                    _ = v15
                    _ = v16
                }
            } else {
                for fi17 := 0; fi17 < 3; fi17++ {
                    p0 = (^(p2))
                    p2 = 255
                    p0 += p0
                    _ = fi17
                }
                v9 -= v9
            }
        }
        if len(v7) > 0 {
            v7[0] = 32768
        }
        _ = v9
    }
    p0 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 bool = (false || true)
    v2 := []int{128, 127}
    v3 := []int{-286, 121, -540, 450}
    v4 := []uint32{42}
    v5 := make(map[string]uint32)
    v6 := []uint32{936, 2147483647}
    if len(v4) > 0 {
        v4[0] = 65535
    }
    v6 = append(v6, 966)
    fmt.Println(v1)
    var v18 string = ((("" + ("" + ("" + ""))) + "\n\t") + "foo bar")
    var v19 uint16 = 32767
    v20 := GenS2{F0: 2, F1: -1}
    fmt.Println(v18, v1)
    if true {
        v19 = 2
        var v21 int32 = -1
        if v1 {
            v21 -= v21
            v19 -= (v19 ^ v19)
        } else {
            if (!(!(v1 && true))) {
                for fi22 := 0; fi22 < 17; fi22++ {
                    _ = len(v5)
                    if v1 {
                        v23 := v5["hello"]
                        var v24 uint32 = 4294967295
                        if true {
                            _ = len(v3)
                            for fi25 := 0; fi25 < 18; fi25++ {
                                if v1 {
                                    var v26 string = "hello"
                                    _ = v26
                                } else {
                                    if (!false) {
                                        _ = len(v5)
                                    } else {
                                        switch v18 {
                                        case "":
                                            v1 = (24 >= 2)
                                        case "hello":
                                            var v27 uint8 = 42
                                            _ = v27
                                        case "\n\t":
                                            v18 = "hello"
                                        }
                                        var v28 uint32 = (v23 + v23)
                                        if v1 {
                                            if true {
                                                v29 := GenS1{F0: 52, F1: fi22, F2: 9223372036854775807}
                                                var v30 uint = 0
                                                _ = len(v3)
                                                _ = v29
                                                _ = v30
                                            } else {
                                                if true {
                                                    fmt.Println(v24, v23)
                                                    var v31 int16 = 42
                                                    _ = v31
                                                } else {
                                                    if true {
                                                        var v32 int8 = -87
                                                        var v33 int8 = 1
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    v34 := GenS2{F0: 42, F1: -1}
                                                    if (!(57 <= 127)) {
                                                        delete(v5, (v18 + ""))
                                                    }
                                                    _ = v34
                                                }
                                            }
                                            v19 -= 25
                                            switch (((v18 + v18) + v18) + "") {
                                            case "a":
                                                if false {
                                                    v23 = v23
                                                    v19 = v19
                                                }
                                            }
                                        }
                                        _ = v28
                                    }
                                    v35 := v5[v18]
                                    _ = v35
                                }
                                if true {
                                    switch fi22 {
                                    case -997:
                                        switch v18 {
                                        case "":
                                            var v36 float32 = 0.1
                                            _ = v36
                                        }
                                    }
                                    switch int(v23) {
                                    case -1:
                                        if (!(!(912 <= 0))) {
                                            var v37 float32 = 0.1
                                            _ = v37
                                        } else {
                                            v3 = append(v3, fi25)
                                            v38 := GenS2{F0: 903, F1: 42}
                                            _ = v38
                                        }
                                    case 2147483647:
                                        var v39 uint32 = v24
                                        _ = v39
                                    default:
                                        var v40 int = fi22
                                        var v41 int32 = (v21 / ((v21) | 1))
                                        v42 := GenS3{F0: 405}
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                    v43 := GenS1{F0: 1, F1: 983, F2: -9223372036854775808}
                                    _ = v43
                                } else {
                                    var v44 uint8 = 127
                                    v45 := GenS4{F0: 1}
                                    _ = v44
                                    _ = v45
                                }
                                _ = fi25
                            }
                        } else {
                            switch "" {
                            case "a":
                                for fi46 := 0; fi46 < 5; fi46++ {
                                    var v47 int64 = 232
                                    var v48 float64 = 1.5
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                            }
                            v49 := GenS2{F0: 4294967295, F1: 0}
                            v1 = (65535 < 1)
                            _ = v49
                        }
                        _ = v23
                        _ = v24
                    } else {
                        v50 := GenS2{F0: 0, F1: 0}
                        _ = v50
                    }
                    _ = fi22
                }
                for fi51 := 0; fi51 < 2; fi51++ {
                    var v52 int64 = 0
                    fmt.Println(v1)
                    v21 = v21
                    _ = fi51
                    _ = v52
                }
            }
            var v53 int = -741
            _ = v53
        }
        _ = v21
    } else {
        fmt.Println(v18, v1)
        v18 = "a"
        v18 = (("hello" + (v18 + (("a" + v18) + ""))) + ("hello" + (("hello" + v18) + v18)))
    }
    v1 = v1
    fmt.Println(v1, v18)
    if false {
        for fi54 := 0; fi54 < 14; fi54++ {
            delete(v5, (v18 + v18))
            var v55 float32 = float32((3.14159 / (1.5 + (3.14159 - 0.1))))
            v56 := v5["foo bar"]
            _ = fi54
            _ = v55
            _ = v56
        }
        v57 := GenS5{F0: 127, F1: -572}
        var v58 int64 = -1
        _ = v57
        _ = v58
    }
    v19 = v19
    switch -9223372036854775808 {
    case -2147483648:
        for fi59 := 0; fi59 < 1; fi59++ {
            var v60 int16 = 0
            _ = fi59
            _ = v60
        }
    }
    if false {
        var v61 uint64 = 18446744073709551615
        var v62 int64 = 0
        _ = v61
        _ = v62
    } else {
        var v63 float64 = 3.14159
        v64 := GenS2{F0: 4294967295, F1: 0}
        v18 = v18
        _ = v63
        _ = v64
    }
    if v1 {
        var v65 float64 = 3.14159
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
