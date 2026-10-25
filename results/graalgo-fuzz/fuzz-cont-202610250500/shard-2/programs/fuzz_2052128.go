package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0(p0 bool) string {
    var v6 int16 = -1
    v6 = 1
    _ = p0
    _ = v6
    return ((("a" + ("foo bar" + "\n\t")) + "hello") + "\n\t")
}
func genFn1(p0 bool) uint16 {
    fmt.Println(p0)
    fmt.Println(p0)
    v7 := []int{32767, -1}
    if true {
        if (0 < 1) {
            var v8 float64 = ((float64(0.0) - 1e10) + (3.14159 / (0.0 + 1e10)))
            fmt.Println(p0)
            _ = v8
        } else {
            p0 = (p0 || p0)
            var v9 int32 = -1
            _ = v9
        }
        v10 := GenS1{F0: 225}
        var v11 uint = 255
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v7
    return 65535
}

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    var v3 int32 = 42
    var v4 int8 = -128
    var v5 int = -9223372036854775808
    delete(v2, (v5 ^ (v5 % ((v5) | 1))))
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v5, fi12, v3)
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int32 = (v3 | (v3 % ((1) | 1)))
                var v15 uint32 = 1
                var v16 int16 = -1
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        if false {
            if (1.5 >= 0.0) {
                var v17 int16 = -250
                _ = v17
            }
            v18 := GenS1{F0: 1}
            _ = v18
        }
    }
    var v19 uint64 = 4294967295
    if (1 >= -32768) {
        if true {
            v20 := GenS1{F0: 62}
            _ = v20
        } else {
            if (1 >= 1) {
                var v21 uint8 = 42
                var v22 float64 = 1e10
                _ = v21
                _ = v22
            } else {
                var v23 int16 = 202
                var v24 int64 = 42
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 16; fi25++ {
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 int32 = v3
                    _ = fi26
                    _ = v27
                }
                fmt.Println(v5, v19, fi25)
                _ = fi25
            }
        }
    }
    if (1 != 42) {
        var v28 uint = 500
        for fi29 := 0; fi29 < 7; fi29++ {
            if (((0.1 > 1e10) && false) || (false && true)) {
                if false {
                    v5 = -1
                    if ((v5 / (((fi29 - -1)) | 1)) > (^(fi29))) {
                        var v30 uint16 = 741
                        delete(v2, (^(fi29)))
                        for fi31 := 0; fi31 < 10; fi31++ {
                            v28 += v28
                            v28 -= 490
                            if true {
                                var v32 bool = true
                                fmt.Println(fi31)
                                _ = v32
                            } else {
                                fmt.Println(v30)
                            }
                            var v33 int = v5
                            _ = fi31
                            _ = v33
                        }
                        _ = v30
                    }
                    v34 := GenS2{F0: 4294967295, F1: 42}
                    _ = v34
                }
                v19 -= v19
                v3 -= (v3 * (v3 | (v3 ^ v3)))
            }
            _ = fi29
        }
        _ = v28
    } else {
        switch "\n\t" {
        case "foo bar":
            v35 := GenS1{F0: 42}
            _ = v35
        case "hello":
            if (514 != 0) {
                var v36 uint32 = 2
                for fi37 := 0; fi37 < 17; fi37++ {
                    var v38 int8 = -45
                    _ = fi37
                    _ = v38
                }
                _ = v36
            }
        case "":
            v3 = (^(v3))
        }
        switch ("hello" + "hello") {
        case "foo bar":
            v3 = (v3 * 598)
        case "":
            var v39 uint16 = 255
            _ = v39
        case "a":
            v4 += (v4 - (v4 >> ((18446744073709551615) & 63)))
        }
        v1 = append(v1, "hello")
    }
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
