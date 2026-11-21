package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int64; F1 uint }
func genFn0() string {
    var v7 string = "\n\t"
    fmt.Println(v7)
    if false {
        v8 := GenS1{F0: -896, F1: 0}
        var v9 uint = 42
        _ = v8
        _ = v9
    } else {
        if true {
            if (true || (251 != -2147483648)) {
                v10 := GenS1{F0: 301, F1: 42}
                _ = v10
            }
            var v11 uint16 = 89
            _ = v11
        }
        v12 := GenS1{F0: 32767, F1: 2}
        _ = v12
    }
    if true {
        if (30 < 9) {
            var v13 uint = 65535
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 float64 = 1.5
                _ = fi14
                _ = v15
            }
            if ((0 != 2147483647) || true) {
                v7 = v7
            } else {
                v13 = 262
                switch (v7 + "hello") {
                case "hello":
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v17 := GenS1{F0: 0, F1: 2}
                        _ = fi16
                        _ = v17
                    }
                case "a":
                    if true {
                        var v18 int16 = 1000
                        v19 := GenS2{F0: 100}
                        v20 := []int{127, 127}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                case "foo bar":
                    v13 = v13
                }
            }
            _ = v13
        }
    }
    v21 := GenS2{F0: -32}
    _ = v21
    _ = v7
    return ("foo bar" + (v7 + "a"))
}
func genFn1(p0 string, p1 int8) bool {
    var v22 int32 = 204
    fmt.Println(p0, p0, p1)
    v23 := GenS2{F0: 127}
    var v24 int16 = -1
    for fi25 := 0; fi25 < 9; fi25++ {
        switch (("" + ((((p0 + p0) + p0) + p0) + (p0 + ""))) + p0) {
        case "foo bar":
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 int64 = 1
                _ = fi26
                _ = v27
            }
        case "hello":
            if (3 < 255) {
                for fi28 := 0; fi28 < 11; fi28++ {
                    v29 := []int{fi25}
                    v29 = append(v29, (^(fi28)))
                    fmt.Println(p0, fi25)
                    _ = fi28
                    _ = v29
                }
                if true {
                    for fi30 := 0; fi30 < 19; fi30++ {
                        p0 = "a"
                        switch fi30 {
                        case -640:
                            v31 := []int{(fi30 & fi25), 2147483647, (fi30 << ((4294967295) & 63))}
                            _ = v31
                        case 128:
                            if true {
                                v22 += -450
                                v32 := GenS1{F0: (v24 + (v24 ^ 0)), F1: 18446744073709551615}
                                _ = v32
                            } else {
                                var v33 uint64 = 2
                                var v34 uint = 934
                                fmt.Println(fi25)
                                _ = v33
                                _ = v34
                            }
                        default:
                            v35 := []int{(fi30 * fi30), (-(fi30)), fi25}
                            _ = v35
                        }
                        if false {
                            v22 = -703
                            v36 := []int{}
                            var v37 uint16 = 32767
                            _ = v36
                            _ = v37
                        }
                        _ = fi30
                    }
                    p1 = p1
                } else {
                    if (p0 >= ("\n\t" + ((("\n\t" + p0) + p0) + ""))) {
                        for fi38 := 0; fi38 < 6; fi38++ {
                            fmt.Println(fi38, v24, fi38)
                            v22 -= (^(v22))
                            v24 -= v24
                            _ = fi38
                        }
                    } else {
                        v39 := []int{fi25}
                        fmt.Println(fi25)
                        _ = v39
                    }
                }
            }
        default:
            var v40 int16 = (v24 >> ((4294967295) & 63))
            _ = v40
        }
        v41 := []int{(fi25 % (((fi25 / ((fi25) | 1))) | 1)), 32767}
        var v42 float64 = (float64(-2.25) - (3.14159 - ((float64(0.0) + (0.0 / 0.1)) - 3.14159)))
        _ = fi25
        _ = v41
        _ = v42
    }
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    _ = v24
    return false
}

func main() {
    var v1 uint16 = 1
    var v2 uint32 = 0
    var v3 int16 = -32768
    var v4 int8 = -1
    var v5 int = int(v1)
    var v6 float32 = 0.1
    v43 := []int{694, (v5 << ((2) & 63)), (^(v5)), -289}
    var v44 uint64 = 668
    for fi45 := 0; fi45 < 15; fi45++ {
        fmt.Println(fi45)
        v46 := GenS3{F0: 979, F1: (v4 & 42)}
        _ = fi45
        _ = v46
    }
    var v47 string = ((("" + (("foo bar" + "hello") + "\n\t")) + "") + ("" + "foo bar"))
    if false {
        var v48 uint32 = uint32(v44)
        v1 -= (v1 >> ((255) & 63))
        if true {
            v49 := GenS1{F0: v3, F1: v44}
            if (106 > (v44 + v44)) {
                var v50 int8 = int8(v3)
                v51 := GenS4{F0: 386, F1: 4294967295}
                var v52 uint8 = uint8(v1)
                _ = v50
                _ = v51
                _ = v52
            }
            if len(v43) > 0 {
                v43[0] = (v5 & v5)
            }
            _ = v49
        } else {
            v53 := GenS3{F0: 42, F1: -128}
            fmt.Println(v5, v4, v2)
            _ = v53
        }
        _ = v48
    }
    v3 = v3
    v5 -= v5
    v43 = append(v43, v5)
    var v54 bool = (v4 < v4)
    v44 -= (v44 + v44)
    v1 -= v1
    fmt.Println(v54)
    v44 = uint64(v1)
    var v55 float64 = 1.5
    for fi56 := 0; fi56 < 5; fi56++ {
        v43 = append(v43, int(v3))
        _ = fi56
    }
    if len(v43) > 0 {
        v43[0] = v5
    }
    var v57 int64 = -855
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v47
    _ = v5
    _ = v54
    _ = v55
    _ = v57
    _ = v6
}
