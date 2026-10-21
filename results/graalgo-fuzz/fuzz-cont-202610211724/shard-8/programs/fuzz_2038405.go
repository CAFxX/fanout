package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 bool) bool {
    var v6 int32 = -628
    v7 := []int{2147483647, 798, 42, p0}
    if p1 {
        v6 += v6
        v8 := GenS1{F0: -31, F1: 32767, F2: ("hello" + (((("foo bar" + (("a" + "a") + "\n\t")) + "foo bar") + (("foo bar" + "") + "hello")) + ("a" + "\n\t")))}
        _ = v8
    } else {
        p0 -= -290
    }
    var v9 uint8 = 2
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v9
    return false
}

func main() {
    v1 := []string{"hello", "\n\t", ""}
    var v2 float64 = ((0.0 - 0.1) / 0.1)
    var v3 bool = (true && true)
    v4 := make(map[string]string)
    var v5 int = 42
    for fi10 := 0; fi10 < 15; fi10++ {
        if v3 {
            var v11 uint16 = 65535
            var v12 uint32 = 0
            for fi13 := 0; fi13 < 14; fi13++ {
                if (!(!(v3 && (!((int64(v11) <= -842) || false))))) {
                    var v14 int8 = 42
                    v1 = append(v1, "\n\t")
                    v5 -= (-(fi10))
                    var v15 uint32 = v12
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: -1, F1: 42, F2: ("" + ("a" + "a"))}
                _ = fi13
                _ = v16
            }
            _ = v11
            _ = v12
        } else {
            if ((((float32(((v2 - 0.0) * (v2 * v2))) * 0.1) + 1e10) - 3.14159) < 3.14159) {
                fmt.Println(v3, v2)
            }
        }
        _ = fi10
    }
    delete(v4, "\n\t")
    v1 = append(v1, "")
    v4["hello"] = ("\n\t" + ("" + "a"))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 int64 = -1
    var v18 int8 = 1
    v3 = (2 <= 528)
    v2 = 3.14159
    if len(v1) > 0 {
        v1[0] = "a"
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            if v3 {
                v18 += v18
            }
            _ = fi20
        }
        var v21 uint32 = 721
        _ = fi19
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 string = ""
        _ = fi22
        _ = v23
    }
    v18 -= (v18 | v18)
    var v24 string = "a"
    if (42 == 42) {
        var v25 int = 297
        v26 := GenS2{F0: v24, F1: v3}
        for fi27 := 0; fi27 < 15; fi27++ {
            if (18446744073709551615 > 630) {
                v28 := GenS2{F0: v24, F1: (4294967295 != 0)}
                if (v18 == v18) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = v2
                } else {
                    v24 = "a"
                    if true {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            switch fi29 {
                            case -92:
                                var v30 string = v24
                                _ = v30
                            case -934:
                                v31 := GenS1{F0: 0, F1: 0, F2: v24}
                                _ = v31
                            case 127:
                                switch (-(v5)) {
                                case 9223372036854775807:
                                    _ = len(v4)
                                }
                            }
                            var v32 bool = (300 > 1)
                            v33 := GenS3{F0: -32768, F1: 255}
                            _ = fi29
                            _ = v32
                            _ = v33
                        }
                    } else {
                        v34 := v4[v24]
                        _ = v34
                    }
                }
                var v35 int64 = v17
                _ = v28
                _ = v35
            }
            _ = len(v1)
            var v36 int8 = 0
            _ = fi27
            _ = v36
        }
        _ = v25
        _ = v26
    } else {
        v17 += v17
        v37 := v4[v24]
        var v38 int64 = (v17 & v17)
        _ = v37
        _ = v38
    }
    v5 -= (v5 << ((4294967295) & 63))
    switch "" {
    case "hello":
        if v3 {
            switch v5 {
            case 2147483647:
                v18 -= -128
            case 42:
                if false {
                    var v39 int32 = 0
                    if len(v1) > 0 {
                        v1[0] = (("\n\t" + v24) + ((v24 + "foo bar") + v24))
                    }
                    _ = v39
                }
            case 127:
                _ = len(v1)
            }
            v3 = v3
            var v40 int = (v5 | 128)
            _ = v40
        }
    case "\n\t":
        v2 = v2
    default:
        var v41 uint16 = 42
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
