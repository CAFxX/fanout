package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int32; F1 uint16; F2 uint }
type GenS3 struct { F0 uint8; F1 string; F2 uint }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64, p1 bool) uint16 {
    v7 := GenS1{F0: 0, F1: "\n\t"}
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        var v8 bool = p1
        _ = v8
    case "hello":
        var v9 uint8 = 42
        _ = v9
    default:
        if true {
            var v10 string = ((("" + "foo bar") + ("foo bar" + "hello")) + ("\n\t" + "\n\t"))
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := []int{fi11, fi11}
                var v13 uint16 = 32767
                if p1 {
                    fmt.Println(p1)
                    v12 = append(v12, fi11)
                } else {
                    switch v10 {
                    case "foo bar":
                        v12 = append(v12, 2147483647)
                    case "\n\t":
                        v14 := GenS2{F0: -2147483648, F1: (v13 | 122), F2: 976}
                        _ = v14
                    case "hello":
                        var v15 int16 = 377
                        _ = v15
                    default:
                        var v16 uint64 = 2
                        _ = v16
                    }
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := []int{}
    var v2 int32 = -1
    var v3 float32 = -2.25
    v4 := []string{}
    var v5 uint8 = 190
    var v6 uint64 = 18446744073709551615
    v4 = append(v4, "\n\t")
    var v17 uint64 = v6
    v18 := GenS2{F0: -956, F1: 794, F2: 1}
    v17 = 0
    if true {
        if false {
            var v19 uint16 = 2
            if true {
                var v20 int16 = -1
                _ = v20
            } else {
                if (true || true) {
                    fmt.Println(v19, v2)
                }
                var v21 int16 = -32768
                _ = v21
            }
            if true {
                for fi22 := 0; fi22 < 7; fi22++ {
                    var v23 int16 = -1
                    var v24 int8 = 127
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                if false {
                    v25 := GenS3{F0: v5, F1: "foo bar", F2: 42}
                    fmt.Println(v6, v17, v5)
                    _ = v25
                }
            } else {
                var v26 int32 = v2
                switch "\n\t" {
                case "a":
                    fmt.Println(v19, v19)
                }
                for fi27 := 0; fi27 < 16; fi27++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v28 float32 = v3
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
            _ = v19
        }
        v5 -= (^(v5))
        var v29 uint64 = v17
        _ = v29
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 int64 = 42
        if true {
            v17 = v17
            var v32 string = "hello"
            _ = v32
        }
        _ = fi30
        _ = v31
    }
    switch (("foo bar" + "") + "\n\t") {
    case "a":
        if (!true) {
            for fi33 := 0; fi33 < 18; fi33++ {
                switch ("" + ("a" + ("hello" + "foo bar"))) {
                case "\n\t":
                    switch "" {
                    case "":
                        _ = len(v1)
                    }
                case "foo bar":
                    for fi34 := 0; fi34 < 7; fi34++ {
                        var v35 bool = false
                        fmt.Println(fi33, fi34)
                        _ = fi34
                        _ = v35
                    }
                case "a":
                    v3 = 1e10
                }
                _ = fi33
            }
            fmt.Println(v17)
        } else {
            fmt.Println(v2, v2)
            var v36 uint64 = (v6 % ((v6) | 1))
            if false {
                var v37 uint32 = 1
                v38 := GenS1{F0: (-(v2)), F1: "hello"}
                _ = v37
                _ = v38
            } else {
                var v39 float32 = (v3 * (1e10 * v3))
                v40 := GenS4{F0: 611}
                _ = v39
                _ = v40
            }
            _ = v36
        }
    case "\n\t":
        if ((!false) && true) {
            for fi41 := 0; fi41 < 5; fi41++ {
                if false {
                    if (true && true) {
                        v3 -= 0.0
                        fmt.Println(v3, v3, v3)
                        var v42 string = "hello"
                        _ = v42
                    } else {
                        var v43 string = "hello"
                        var v44 float32 = v3
                        _ = v43
                        _ = v44
                    }
                } else {
                    v2 = (^(v2))
                    if false {
                        v5 += v5
                        for fi45 := 0; fi45 < 11; fi45++ {
                            v4 = append(v4, "")
                            _ = fi45
                        }
                    }
                    fmt.Println(v6)
                }
                _ = fi41
            }
            if (v17 > v17) {
                var v46 float64 = (1e10 / 1.0)
                _ = v46
            }
            if len(v1) > 0 {
                v1[0] = 42
            }
        } else {
            var v47 bool = true
            _ = v47
        }
    }
    var v48 uint32 = 65535
    var v49 bool = true
    v17 = v17
    for fi50 := 0; fi50 < 3; fi50++ {
        for fi51 := 0; fi51 < 10; fi51++ {
            v52 := GenS2{F0: 0, F1: 487, F2: 4294967295}
            _ = fi51
            _ = v52
        }
        v6 += 2
        v53 := GenS2{F0: (-(v2)), F1: 504, F2: 0}
        _ = fi50
        _ = v53
    }
    fmt.Println(v48, v2, v5)
    for fi54 := 0; fi54 < 5; fi54++ {
        fmt.Println(v49, fi54, v49)
        v5 -= (v5 & v5)
        v6 = v6
        _ = fi54
    }
    for fi55 := 0; fi55 < 10; fi55++ {
        for fi56 := 0; fi56 < 2; fi56++ {
            if len(v1) > 0 {
                v1[0] = fi55
            }
            if v49 {
                v57 := GenS1{F0: v2, F1: "\n\t"}
                _ = v57
            } else {
                if false {
                    if (true || (v49 && (v3 < v3))) {
                        if ((fi56 < (fi56 - 32768)) || ((v49 && (v49 || (2147483647 <= v48))) && (-278 > v2))) {
                            v3 = -2.25
                        } else {
                            _ = len(v1)
                            switch "\n\t" {
                            case "a":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "\n\t":
                                switch fi56 {
                                case 9223372036854775807:
                                    v6 = 1
                                case 371:
                                    if true {
                                        v58 := GenS4{F0: (v48 / ((v48) | 1))}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v58
                                    }
                                    var v59 int = (fi55 & fi56)
                                    v6 -= v17
                                    _ = v59
                                case -9223372036854775808:
                                    v2 = int32(v6)
                                }
                            }
                            v60 := GenS1{F0: (v2 & v2), F1: "hello"}
                            _ = v60
                        }
                    }
                    v17 = v17
                    v1 = append(v1, (fi55 & fi56))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v61 := GenS4{F0: v48}
                _ = v61
            }
            fmt.Println(fi55)
            _ = fi56
        }
        _ = len(v4)
        _ = fi55
    }
    v62 := GenS4{F0: v48}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v62
}
