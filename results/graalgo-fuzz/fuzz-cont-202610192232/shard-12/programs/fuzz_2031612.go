package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint64 = 0
        var v8 uint64 = (^(v7))
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "" {
    case "hello":
        v9 := GenS1{F0: 2147483648, F1: 2, F2: false}
        _ = v9
    case "foo bar":
        v10 := []int{938, 32768}
        _ = v10
    }
    v11 := GenS1{F0: 65535, F1: 204, F2: false}
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            switch (("hello" + ("foo bar" + "hello")) + "a") {
            case "hello":
                if (!(958 < 65535)) {
                    switch (fi13 * -466) {
                    case 2147483647:
                        var v14 float64 = -2.25
                        _ = v14
                    case 128:
                        switch "a" {
                        case "foo bar":
                            if (!(("a" + (("hello" + (("a" + ("foo bar" + (("foo bar" + "foo bar") + "foo bar"))) + "\n\t")) + "foo bar")) == ("foo bar" + "\n\t"))) {
                                var v15 uint32 = 42
                                v16 := GenS2{F0: 42}
                                _ = v15
                                _ = v16
                            }
                        default:
                            var v17 int32 = 1
                            _ = v17
                        }
                    default:
                        var v18 int8 = 127
                        var v19 int16 = -32768
                        _ = v18
                        _ = v19
                    }
                    var v20 int32 = 42
                    var v21 bool = true
                    _ = v20
                    _ = v21
                } else {
                    var v22 uint8 = 128
                    v22 -= v22
                    var v23 bool = ((42 != 0) || ((128 != (^(v22))) && false))
                    _ = v22
                    _ = v23
                }
            case "foo bar":
                var v24 int32 = 42
                _ = v24
            case "\n\t":
                if false {
                    var v25 int32 = 42
                    fmt.Println(v25, fi13)
                    _ = v25
                } else {
                    fmt.Println(fi13, fi12)
                    var v26 int64 = 28
                    v27 := []int{(fi13 ^ fi12), (fi13 & -584)}
                    _ = v26
                    _ = v27
                }
            default:
                var v28 int = (fi12 | 128)
                _ = v28
            }
            if false {
                var v29 uint64 = 4294967295
                v30 := []int{}
                var v31 uint32 = 2
                _ = v29
                _ = v30
                _ = v31
            }
            if (true && (-2.25 >= float64(((1e10 + 1e10) + 0.0)))) {
                v32 := []int{(^(fi13)), fi12, 128, -1}
                var v33 uint64 = 239
                v34 := GenS3{F0: 462, F1: -631, F2: 0}
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi13
        }
        var v35 float32 = (1.5 / 1e10)
        _ = fi12
        _ = v35
    }
    v36 := []int{}
    _ = v11
    _ = v36
    return 76
}
func genFn1() int32 {
    switch 1 {
    case 32767:
        v37 := []int{32767, 9223372036854775807, 127}
        _ = v37
    case 4294967295:
        v38 := GenS4{F0: 2, F1: 4294967295, F2: 142}
        _ = v38
    default:
        var v39 uint32 = genFn0()
        _ = v39
    }
    var v40 uint64 = 548
    for fi41 := 0; fi41 < 20; fi41++ {
        v40 += (v40 / ((v40) | 1))
        _ = fi41
    }
    fmt.Println(v40)
    var v42 int64 = -995
    _ = v40
    _ = v42
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 string = ("" + (("" + ("hello" + ("a" + ""))) + "a"))
    v3 := make(map[uint32]int)
    var v4 uint64 = 18446744073709551615
    var v5 float64 = (float64(1e10) * 1.5)
    var v43 int16 = 1
    v44 := GenS2{F0: -1}
    v45 := []int{}
    for fi46 := 0; fi46 < 10; fi46++ {
        var v47 uint64 = 518
        _ = fi46
        _ = v47
    }
    var v48 string = ("hello" + v2)
    for fi49 := 0; fi49 < 7; fi49++ {
        if ((v1 % (((v1 / (((^(v1))) | 1))) | 1)) != v1) {
            switch v2 {
            case "a":
                var v50 int = fi49
                _ = v50
            }
        } else {
            v45 = append(v45, -760)
        }
        switch (fi49 + -158) {
        case 128:
            v5 += v5
        case 1:
            var v51 float32 = float32((v5 + v5))
            _ = v51
        case 9223372036854775807:
            fmt.Println(v48, v1)
        default:
            fmt.Println(v2, v1)
        }
        v3[4294967295] = (fi49 ^ (fi49 * 2147483647))
        _ = fi49
    }
    if (genFn0() > 4294967295) {
        v5 += (v5 * v5)
    }
    if false {
        if (v4 != 518) {
            fmt.Println(v48)
        }
        v5 = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v48
    _ = v5
}
