package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint32 }
func genFn0() int {
    for fi4 := 0; fi4 < 10; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    switch -511 {
    case -852:
        var v5 string = "hello"
        _ = v5
    case 4294967295:
        var v6 int = 127
        _ = v6
    case 9223372036854775807:
        switch 32767 {
        case 1:
            var v7 float32 = 0.0
            _ = v7
        }
    }
    v8 := GenS1{F0: 9223372036854775807, F1: (true || false), F2: 927}
    _ = v8
    return 32767
}

func main() {
    var v1 float32 = 0.0
    var v2 bool = true
    var v3 string = (("hello" + ("a" + "foo bar")) + "hello")
    var v9 uint8 = 1
    var v10 uint64 = 42
    var v11 int = 11
    v2 = ((^(v9)) >= (v9 | v9))
    v3 = v3
    v12 := GenS1{F0: 1, F1: v2, F2: 286}
    if (0 > 0) {
        if v2 {
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 int = (-(v11))
                _ = fi13
                _ = v14
            }
            v15 := []int{v11}
            _ = v15
        } else {
            v16 := GenS1{F0: 1, F1: v2, F2: 4294967295}
            v17 := []int{}
            for fi18 := 0; fi18 < 18; fi18++ {
                if (!v2) {
                    fmt.Println(v1, fi18, v1)
                } else {
                    switch fi18 {
                    case -1:
                        if true {
                            if true {
                                for fi19 := 0; fi19 < 20; fi19++ {
                                    var v20 uint64 = (v10 & v10)
                                    var v21 int64 = 9223372036854775807
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                            }
                            if (false || false) {
                                v11 += fi18
                            }
                        } else {
                            v9 += (v9 & (v9 << ((4294967295) & 63)))
                        }
                    default:
                        switch ("" + v3) {
                        case "a":
                            var v22 float32 = 1.5
                            _ = v22
                        case "\n\t":
                            v9 = v9
                        case "hello":
                            var v23 string = ((v3 + ("foo bar" + "hello")) + "hello")
                            _ = v23
                        default:
                            if false {
                                v9 = v9
                                if v2 {
                                    fmt.Println(v1)
                                    var v24 int32 = 42
                                    _ = v24
                                }
                                for fi25 := 0; fi25 < 4; fi25++ {
                                    var v26 string = v3
                                    _ = fi25
                                    _ = v26
                                }
                            }
                        }
                    }
                    v17 = append(v17, (v11 << ((0) & 63)))
                }
                _ = fi18
            }
            _ = v16
            _ = v17
        }
    } else {
        fmt.Println(v9, v9, v3)
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 float32 = 1e10
            var v29 uint = 42
            _ = fi27
            _ = v28
            _ = v29
        }
        var v30 float32 = v1
        _ = v30
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        v3 = "a"
        _ = fi31
    }
    if ((^(v9)) <= (^(v9))) {
        var v32 int = 858
        _ = v32
    }
    v33 := []int{127, v11, -878, (v11 * genFn0())}
    switch v3 {
    case "":
        for fi34 := 0; fi34 < 3; fi34++ {
            v1 = 0.0
            _ = fi34
        }
    case "\n\t":
        v2 = (!(!v2))
    default:
        for fi35 := 0; fi35 < 7; fi35++ {
            var v36 int16 = 32767
            v33 = append(v33, (fi35 % (((-(fi35))) | 1)))
            v11 -= v11
            _ = fi35
            _ = v36
        }
    }
    fmt.Println(v9, v1, v1)
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 int = 127
        v9 = 192
        var v39 float32 = 1.5
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v40:", v40)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v40
    _ = v9
}
