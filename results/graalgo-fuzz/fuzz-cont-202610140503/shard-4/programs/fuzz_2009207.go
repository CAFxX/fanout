package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 string }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 string; F1 int8; F2 string }
func genFn0() uint16 {
    if false {
        var v4 bool = (-1 == -9223372036854775808)
        _ = v4
    } else {
        switch "hello" {
        case "":
            v5 := []int{}
            _ = v5
        case "\n\t":
            var v6 uint8 = 128
            _ = v6
        case "foo bar":
            var v7 int64 = 9223372036854775807
            _ = v7
        default:
            var v8 int32 = 795
            _ = v8
        }
    }
    var v9 bool = true
    _ = v9
    return 1
}
func genFn1() int16 {
    var v10 int = 916
    v11 := GenS1{F0: 1, F1: genFn0(), F2: "hello"}
    var v12 int32 = -2147483648
    var v13 int8 = 127
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 float32 = 3.14159
        if false {
            fmt.Println(v15)
            if false {
                for fi16 := 0; fi16 < 18; fi16++ {
                    v17 := GenS1{F0: 83, F1: genFn0(), F2: "a"}
                    _ = fi16
                    _ = v17
                }
            }
        }
        v18 := []int{125, fi14}
        _ = fi14
        _ = v15
        _ = v18
    }
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 888
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 18446744073709551615
    var v3 int16 = -877
    var v19 int32 = -2147483648
    var v20 int = -9223372036854775808
    if (0 <= 42) {
        var v21 int32 = v19
        for fi22 := 0; fi22 < 1; fi22++ {
            switch (^(v20)) {
            case -9223372036854775808:
                v21 += (-(v21))
            }
            _ = fi22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            fmt.Println(v2, v2, v2)
            if (!true) {
                if false {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        var v25 string = ("hello" + ("hello" + "hello"))
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 1; fi26++ {
                        switch 0 {
                        case -9223372036854775808:
                            v27 := GenS2{F0: "a", F1: (fi26 << ((1) & 63))}
                            _ = v27
                        case 128:
                            var v28 uint64 = (v2 - (v2 - v2))
                            _ = v28
                        case 0:
                            switch ("a" + "a") {
                            case "a":
                                v29 := GenS3{F0: v2, F1: 28, F2: -1}
                                _ = v29
                            case "hello":
                                var v30 uint16 = 1
                                _ = v30
                            default:
                                var v31 float32 = 1e10
                                _ = v31
                            }
                        }
                        if false {
                            v21 -= v21
                            v32 := GenS3{F0: v2, F1: -128, F2: 979}
                            _ = v32
                        }
                        v33 := GenS3{F0: v2, F1: 0, F2: 42}
                        v34 := GenS2{F0: ("\n\t" + ("\n\t" + "hello")), F1: (fi26 % (((v20 / ((int(v2)) | 1))) | 1))}
                        _ = fi26
                        _ = v33
                        _ = v34
                    }
                    v21 -= -2147483648
                }
                switch "a" {
                case "a":
                    v35 := v1[(fi23 >> ((140) & 63))]
                    _ = v35
                default:
                    var v36 int64 = 489
                    _ = v36
                }
                if (255 != 0) {
                    fmt.Println(v19, v3)
                    if false {
                        var v37 uint16 = genFn0()
                        if false {
                            v38 := []int{fi23, (-(fi23)), (v20 ^ (fi23 | v20)), fi23}
                            _ = v38
                        } else {
                            var v39 uint32 = 42
                            _ = v39
                        }
                        _ = v37
                    }
                    v21 = (v19 >> ((4294967295) & 63))
                }
            }
            for fi40 := 0; fi40 < 2; fi40++ {
                var v41 bool = ((v21 << ((42) & 63)) >= v19)
                v42 := []int{(fi23 % (((^(fi23))) | 1))}
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = fi23
        }
        _ = v21
    } else {
        var v43 string = "\n\t"
        v43 = (v43 + v43)
        _ = v43
    }
    v19 += v19
    v20 -= v20
    var v44 float32 = (float32(0.1) * 1e10)
    v45 := GenS4{F0: v20}
    for fi46 := 0; fi46 < 14; fi46++ {
        var v47 int32 = v19
        v48 := v1[v20]
        var v49 int8 = -21
        _ = fi46
        _ = v47
        _ = v48
        _ = v49
    }
    if (242 == 1) {
        v50 := GenS5{F0: 128, F1: ""}
        var v51 uint = 154
        _ = v50
        _ = v51
    }
    v3 = (v3 / ((genFn1()) | 1))
    if true {
        var v52 int16 = (v3 ^ genFn1())
        v53 := []int{int(v2), v20, -1}
        _ = v52
        _ = v53
    } else {
        v54 := GenS6{F0: "", F1: 91, F2: ""}
        var v55 string = "foo bar"
        _ = v54
        _ = v55
    }
    if (2147483647 > (v20 << ((1) & 63))) {
        v56 := v1[v20]
        var v57 uint64 = 4294967295
        var v58 int16 = v3
        _ = v56
        _ = v57
        _ = v58
    }
    v59 := GenS2{F0: "\n\t", F1: (v20 | v20)}
    switch -482 {
    case 32768:
        v3 = (v3 | (v3 >> ((42) & 63)))
    case 548:
        v60 := GenS3{F0: 1, F1: 0, F2: 42}
        _ = v60
    }
    if (false && ((v19 << ((61) & 63)) > int32(v2))) {
        var v61 int = (-(v20))
        fmt.Println(v20, v2)
        _ = v61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v44
    _ = v45
    _ = v59
}
